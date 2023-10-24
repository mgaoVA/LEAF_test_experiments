package main

import (
	"encoding/json"
	"io"
	"log"
	"testing"
)

func getFormQuery(url string) FormQueryResponse {
	res, _ := client.Get(url)
	b, _ := io.ReadAll(res.Body)

	var m FormQueryResponse
	err := json.Unmarshal(b, &m)
	if err != nil {
		log.Printf("JSON parsing error, couldn't parse: %v", string(b))
		log.Printf("JSON parsing error: %v", err.Error())
	}
	return m
}

func TestForm_HomepageQuery(t *testing.T) {
	res := getFormQuery(rootURL + `api/form/query?q={"terms":[{"id":"title","operator":"LIKE","match":"***","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`)

	got := res[460].RecordID
	want := 460

	if got != want {
		t.Errorf("RecordID = %v, want = %v", got, want)
	}
}

func TestForm_NonadminQueryActionable(t *testing.T) {
	res := getFormQuery(rootURL + `api/form/query?q={"terms":[{"id":"stepID","operator":"=","match":"actionable","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service"],"sort":{},"limit":1000,"limitOffset":0}&x-filterData=recordID,title&masquerade=nonAdmin`)

	if _, exists := res[503]; !exists {
		t.Errorf("Record 503 should be readable because tester is backup of person designated")
	}

	if _, exists := res[504]; !exists {
		t.Errorf("Record 504 should be readable because tester is backup of initiator")
	}

	if _, exists := res[505]; exists {
		t.Errorf("Record 505 should not be readable because tester is not the requestor")
	}
}
