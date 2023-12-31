import { assertExists, assertEquals } from 'https://deno.land/std/assert/mod.ts';
import Helper from './helper.ts';

let h = new Helper();
await h.init();
let fetch = h.wrapFetch();
let rootURL = h.rootURL;

Deno.test("formWorkflow: currentStep person designated AND group requirements", async () => {
    let res = await fetch(rootURL + `api/formWorkflow/484/currentStep`);
    let data = await res.json()

    let got = data[9].description;
    let want = 'Group A';
    assertEquals(data[9].description, want, 
        `description = ${got}, want = ${want}`);

    got = data[-1].description;
    want = 'Step 1 (Omar Marvin)';
    assertEquals(got, want, 
        `description = ${got}, want = ${want}`);

    got = data[9].approverName;
    assertExists(!got, 
        `approver exists = ${!got}, want = false`); // approverName should not exist for depID 9
});
