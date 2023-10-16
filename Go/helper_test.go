package main

import (
	"io"
)

func getJson(url string) string {
	res, _ := client.Get(url)
	bodyBytes, _ := io.ReadAll(res.Body)
	return string(bodyBytes)
}
