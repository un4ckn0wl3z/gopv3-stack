package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_readJSON(t *testing.T) {
	// sample json
	sampleJSON := map[string]interface{}{
		"foo": "bar",
	}

	body, _ := json.Marshal(sampleJSON)

	// read into var
	var decodedJSON struct {
		FOO string `json:"foo"`
	}

	req, err := http.NewRequest("POST", "/", bytes.NewReader(body))
	if err != nil {
		t.Log(err)
	}
	// test rr
	rr := httptest.NewRecorder()
	defer req.Body.Close()
	err = testApp.readJSON(rr, req, &decodedJSON)

	if err != nil {
		t.Error("failed to decode json", err)
	}
}

func Test_writeJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := jsonResponse{
		Error:   false,
		Message: "Test",
	}

	headers := make(http.Header)
	headers.Add("FOO", "BAR")
	err := testApp.writeJSON(rr, http.StatusOK, payload, headers)

	if err != nil {
		t.Errorf("failed to write json %v", err)
	}
	testApp.environment = "production"

	err = testApp.writeJSON(rr, http.StatusOK, payload, headers)

	if err != nil {
		t.Errorf("failed to write json in production env %v", err)
	}

	testApp.environment = "development"

}
