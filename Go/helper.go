package main

import (
	"errors"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

func checkExists(input string) func(*http.Response, *http.Request) error {
	return func(res *http.Response, _ *http.Request) error {
		body, _ := io.ReadAll(res.Body)
		json := string(body)

		if gjson.Get(json, input).Exists() {
			return nil
		}
		return errors.New("JSON element doesn't exist: " + input)
	}
}
