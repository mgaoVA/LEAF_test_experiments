import { assertExists, assertEquals } from 'https://deno.land/std@0.204.0/assert/mod.ts';
import Helper from './helper.ts';

let h = new Helper();
await h.init();
let fetch = h.wrapFetch();
let rootURL = h.rootURL;

Deno.test("form: version", async () => {
    let res = await fetch(rootURL + `api/form/version`);
    let val = await res.json()

    assertEquals(val, "1");
});

Deno.test("form: query homepage", async () => {
    let res = await fetch(rootURL + `api/form/query?q={"terms":[{"id":"title","operator":"LIKE","match":"***","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`);
    let val = await res.json()

    assertEquals(val[504].recordID, 504);
});

Deno.test("form: non-admin query for actionable records", async () => {
    let res = await fetch(rootURL + `api/form/query?q={"terms":[{"id":"stepID","operator":"=","match":"actionable","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service"],"sort":{},"limit":1000,"limitOffset":0}&x-filterData=recordID,title&masquerade=nonAdmin`);
    let val = await res.json()

    assertExists(val[503]); // tester is backup of person designated
    assertExists(val[504]); // tester is backup of initiator
    assertExists(!val[505]); // tester is not the tester
});
