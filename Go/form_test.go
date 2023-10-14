package main

import (
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestFormVersion(t *testing.T) {
	apitest.New().
		EnableNetworking(cli).
		Get(rootURL + "api/form/version").
		Expect(t).Body(`"1"`).
		Status(200).
		End()
}

func TestFormHomepageQuery(t *testing.T) {
	apitest.New().
		EnableNetworking(cli).
		Get(rootURL+`api/form/query`).
		Query("q", `{"terms":[{"id":"title","operator":"LIKE","match":"***","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`).
		Expect(t).
		Assert(checkExists("460.recordID")).
		Status(200).
		End()
}
