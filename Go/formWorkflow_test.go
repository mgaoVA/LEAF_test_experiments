package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tidwall/gjson"
)

func TestFormWorkflow_currentStepPersonDesignatedAndGroup(t *testing.T) {
	data := getJson(rootURL + `api/formWorkflow/484/currentStep`)

	gotString := gjson.Get(data, "9.description").String()
	wantString := "Group A"
	if !cmp.Equal(gotString, wantString) {
		t.Errorf("Description = %v, want = %v", gotString, wantString)
	}

	gotString = gjson.Get(data, "-1.description").String()
	wantString = "Step 1 (Omar Marvin)"
	if !cmp.Equal(gotString, wantString) {
		t.Errorf("Description = %v, want = %v", gotString, wantString)
	}

	gotBool := gjson.Get(data, "9.approverName").Exists()
	wantBool := false
	if gotBool != wantBool {
		t.Errorf("approverName Exists = %v, want = %v", gotBool, wantBool)
	}
}
