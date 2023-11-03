package main

import (
	"encoding/json"
	"io"
	"strings"
	"testing"
)

func getFormQuery(url string) (FormQueryResponse, error) {
	url = strings.Replace(url, " ", "%20", -1)
	res, _ := client.Get(url)
	b, _ := io.ReadAll(res.Body)

	var m FormQueryResponse
	err := json.Unmarshal(b, &m)

	return m, err
}

func TestFormQuery_HomepageQuery(t *testing.T) {
	res, _ := getFormQuery(rootURL + `api/form/query?q={"terms":[{"id":"title","operator":"LIKE","match":"***","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`)

	// get first key
	var key int
	for k, _ := range res {
		key = k
		break
	}

	got := res[key].RecordID
	want := key

	if got != want {
		t.Errorf("RecordID = %v, want = %v", got, want)
	}
}

func TestFormQuery_NonadminQueryActionable(t *testing.T) {
	res, _ := getFormQuery(rootURL + `api/form/query?q={"terms":[{"id":"stepID","operator":"=","match":"actionable","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service"],"sort":{},"limit":1000,"limitOffset":0}&x-filterData=recordID,title&masquerade=nonAdmin`)

	if _, exists := res[503]; !exists {
		t.Errorf("Record 503 should be readable because tester is backup of person designated")
	}

	if _, exists := res[504]; !exists {
		t.Errorf("Record 504 should be readable because tester is backup of initiator")
	}

	if _, exists := res[505]; exists {
		t.Errorf("Record 505 should not be readable because tester is not the requestor")
	}

	if _, exists := res[500]; !exists {
		t.Errorf("Record 500 should be readable because tester is the designated reviewer")
	}
}

func TestFormQuery_FulltextSearch_ApplePearOrange(t *testing.T) {
	res, _ := getFormQuery(rootURL + `api/form/query?q={"terms":[{"id":"data","indicatorID":"3","operator":"MATCH","match":"apple pear orange","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`)

	if _, exists := res[499]; !exists {
		t.Errorf(`Record 499 should be returned because a data field contains either apple, pear, or orange`)
	}

	if _, exists := res[498]; !exists {
		t.Errorf(`Record 498 should be returned because a data field contains either apple, pear, or orange`)
	}
}

func TestFormQuery_FulltextSearch_ApplePear_RequireOrange(t *testing.T) {
	res, _ := getFormQuery(rootURL + `api/form/query?q={"terms":[{"id":"data","indicatorID":"3","operator":"MATCH","match":"apple pear %2Borange","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`)

	if _, exists := res[499]; !exists {
		t.Errorf(`Record 499 should be returned because a data field contains the word "orange"`)
	}
}

func TestFormQuery_FulltextSearch_ApplePearNoOrange(t *testing.T) {
	res, _ := getFormQuery(rootURL + `api/form/query?q={"terms":[{"id":"data","indicatorID":"3","operator":"MATCH","match":"apple pear %2Dorange","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`)

	if _, exists := res[499]; exists {
		t.Errorf(`Record 499 should not be returned because the data field contains the word "orange". want = no orange`)
	}

	if _, exists := res[498]; !exists {
		t.Errorf(`Record 498 should be returned because the data field does not contain the word "orange"`)
	}
}
