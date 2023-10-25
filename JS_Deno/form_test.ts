import { assertExists, assertEquals } from 'https://deno.land/std/assert/mod.ts';
import Helper from './helper.ts';

let h = new Helper();
await h.init();
let fetch = h.wrapFetch();
let rootURL = h.rootURL;
let CSRFToken = h.CSRFToken;

Deno.test("form: version", async () => {
    let res = await fetch(rootURL + `api/form/version`);
    let data = await res.json()

    assertEquals(data, "1", `form version = ${data}, want = 1`);
});

Deno.test("form: query homepage", async () => {
    let res = await fetch(rootURL + `api/form/query?q={"terms":[{"id":"title","operator":"LIKE","match":"***","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`);
    let data = await res.json()

    assertExists(data[460].recordID, `RecordID 460 should be readable`)
});

Deno.test("form: non-admin query for actionable records", async () => {
    let res = await fetch(rootURL + `api/form/query?q={"terms":[{"id":"stepID","operator":"=","match":"actionable","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service"],"sort":{},"limit":1000,"limitOffset":0}&x-filterData=recordID,title&masquerade=nonAdmin`);
    let data = await res.json()

    assertExists(data[503], `Record 503 should be readable because tester is backup of person designated`);
    assertExists(data[504], `Record 504 should be readable because tester is backup of initiator`);
    assertExists(!data[505], `Record 505 should not be readable because tester is not the requestor`);
    assertExists(data[500], `Record 500 should be readable because tester is the designated reviewer`);
});

Deno.test("form: admin edit record datafield", async () => {
    let formData = new FormData();
    formData.append('CSRFToken', CSRFToken);
    formData.append('3', '12345');

    let res = await fetch(rootURL + `api/form/505`, {
        method: "POST",
        body: formData
    });
    let data = await res.json()

    assertEquals(data, "1", `admin edit = ${data}, want = 1`);
});

Deno.test("form: non-admin cannot edit record datafield", async () => {
    let formData = new FormData();
    formData.append('CSRFToken', CSRFToken);
    formData.append('3', '12345');

    let res = await fetch(rootURL + `api/form/505?masquerade=nonAdmin`, {
        method: "POST",
        body: formData
    });
    let data = await res.json()

    assertEquals(data, "0", `non-admin edit = ${data}, want = 0`);
});