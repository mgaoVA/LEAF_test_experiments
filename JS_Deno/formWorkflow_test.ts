import { assertExists, assertEquals } from 'https://deno.land/std@0.204.0/assert/mod.ts';
import Helper from './helper.ts';

let h = new Helper();
await h.init();
let fetch = h.wrapFetch();
let rootURL = h.rootURL;

Deno.test("formWorkflow: currentStep person designated AND group requirements", async () => {
    let res = await fetch(rootURL + `api/formWorkflow/484/currentStep`);
    let val = await res.json()

    assertEquals(val[9].description, 'Group A');
    assertEquals(val[-1].description, 'Step 1 (Omar Marvin)');

    assertExists(!val[9].approverName); // approverName should not exist for depID 9
});
