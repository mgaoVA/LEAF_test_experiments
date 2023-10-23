package main

import (
	"io"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tidwall/gjson"
)

func TestForm_Version(t *testing.T) {
	got := getJson(rootURL + "api/form/version")
	want := `"1"`

	if !cmp.Equal(got, want) {
		t.Errorf("form version = %v, want = %v", got, want)
	}
}

func TestForm_HomepageQuery(t *testing.T) {
	got := getJson(rootURL + `api/form/query?q={"terms":[{"id":"title","operator":"LIKE","match":"***","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service","status","categoryName"],"sort":{"column":"date","direction":"DESC"},"limit":50}`)

	if !gjson.Get(got, "460.recordID").Exists() {
		t.Errorf("RecordID 460 should be readable")
	}
}

func TestForm_NonadminQueryActionable(t *testing.T) {
	got := getJson(rootURL + `api/form/query?q={"terms":[{"id":"stepID","operator":"=","match":"actionable","gate":"AND"},{"id":"deleted","operator":"=","match":0,"gate":"AND"}],"joins":["service"],"sort":{},"limit":1000,"limitOffset":0}&x-filterData=recordID,title&masquerade=nonAdmin`)

	if !gjson.Valid(got) {
		t.Errorf("Invalid JSON. got = %v, want = valid json", got)
	}

	if !gjson.Get(got, "503").Exists() {
		t.Errorf("Record 504 should be readable because tester is backup of person designated")
	}

	if !gjson.Get(got, "504").Exists() {
		t.Errorf("Record 504 should be readable because tester is backup of initiator")
	}

	if gjson.Get(got, "505").Exists() {
		t.Errorf("Record 505 should not be readable because tester is not the requestor")
	}
}

func TestForm_AdminCanEditData(t *testing.T) {
	postData := url.Values{}
	postData.Set("CSRFToken", csrfToken)
	postData.Set("3", "12345")

	res, _ := client.PostForm(rootURL+`api/form/505`, postData)
	bodyBytes, _ := io.ReadAll(res.Body)
	got := string(bodyBytes)
	want := `"1"`

	if !cmp.Equal(got, want) {
		t.Errorf("Admin got = %v, want = %v", got, want)
	}
}

func TestForm_NonadminCannotEditData(t *testing.T) {
	postData := url.Values{}
	postData.Set("CSRFToken", csrfToken)
	postData.Set("3", "12345")

	res, _ := client.PostForm(rootURL+`api/form/505?masquerade=nonAdmin`, postData)
	bodyBytes, _ := io.ReadAll(res.Body)
	got := string(bodyBytes)
	want := `"0"`

	if !cmp.Equal(got, want) {
		t.Errorf("Non-admin got = %v, want = %v", got, want)
	}
}
