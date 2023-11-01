package main

import (
	"io"
	"net/http"
)

func httpGet(url string) (string, *http.Response) {
	res, _ := client.Get(url)
	bodyBytes, _ := io.ReadAll(res.Body)
	return string(bodyBytes), res
}
