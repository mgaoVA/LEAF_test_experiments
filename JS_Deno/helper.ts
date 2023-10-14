import * as anotherCookiejar from "https://deno.land/x/another_cookiejar@v5.0.3/mod.ts";

export default class Helper {
    public cookieJar = new anotherCookiejar.CookieJar();
    public rootURL = "https://localhost/LEAF_Request_Portal/";

    constructor() {
    }

    // init performs initial setup and logs into the dev environment.
    // In dev, the current username is set via REMOTE_USER docker environment
    async init() {
        let fetch = this.wrapFetch();
        let res = await fetch(this.rootURL, {
            headers: {
                "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46"
            }
        });
        await res.text();
        if(res.status != 200) {
            throw Error("Can't log in");
        }
    }

    wrapFetch() {
        return anotherCookiejar.wrapFetch({
            fetch: fetch,
            cookieJar: this.cookieJar
        });
    }
}