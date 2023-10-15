import { assertExists, assertEquals } from 'https://deno.land/std/assert/mod.ts';
import Helper from './helper.ts';

let h = new Helper();
await h.init();
let fetch = h.wrapFetch();
let rootURL = h.rootURL;

Deno.test("formWorkflow: currentStep person designated AND group requirements", async () => {
    let res = await fetch(rootURL + `api/formWorkflow/484/currentStep`);
    let data = await res.json()

    assertEquals(data[9].description, 'Group A');
    assertEquals(data[-1].description, 'Step 1 (Omar Marvin)');

    assertExists(!data[9].approverName); // approverName should not exist for depID 9
});
