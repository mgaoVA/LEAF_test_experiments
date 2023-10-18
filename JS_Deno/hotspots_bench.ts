import Helper from './helper.ts';

let h = new Helper();
await h.init();
let fetch = h.wrapFetch();
let rootURL = h.rootURL;

Deno.bench("Homepage: default query", async () => {
    await fetch(rootURL + `api/form/query?q={"terms":[{"id":"title","operator":"LIKE","match":"***","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`);
});

Deno.bench("Inbox: non-admin query for actionable records", async () => {
    await fetch(rootURL + `api/form/query?q={"terms":[{"id":"stepID","operator":"=","match":"actionable","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service"],"sort":{},"limit":1000,"limitOffset":0}&x-filterData=recordID,title&masquerade=nonAdmin`);
});
