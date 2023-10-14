import Helper from './helper.ts';

let h = new Helper();
await h.init();
let fetch = h.wrapFetch();
let rootURL = h.rootURL;

Deno.test("api/form/version", async () => {
    let res = await fetch(rootURL + `api/form/version`);
    let val = await res.json()

    if(val != 1) {
        throw Error("Version doesn't match");
    }
});

Deno.test("api/form/query - homepage", async () => {
    let res = await fetch(rootURL + `api/form/query?q={"terms":[{"id":"title","operator":"LIKE","match":"***","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`);
    let val = await res.json()

    if(val[504].recordID != 504) {
        throw Error("Wrong record ID");
    }
});