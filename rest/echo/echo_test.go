package echo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoBody(t *testing.T) {

	// given
	reqPayload := `{
		"field1": "value1",
		"field2": "value2"
	}`
	req, err := http.NewRequest("GET", "/", strings.NewReader(reqPayload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	respRecorder := httptest.NewRecorder()

	// when
	Body(respRecorder, req)

	// then
	if respRecorder.Code != http.StatusOK {
		t.Errorf("Invalid response status (expected: %d, actual: %d)", http.StatusOK, respRecorder.Code)
	}

	respPayload := respRecorder.Body.String()
	if respPayload != reqPayload {
		t.Errorf("Invalid response body (actual: '%s')", respPayload)
	}
}

func TestEchoHeaders(t *testing.T) {

	// given
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/html")
	req.Header.Set("Authorization", "Bearer terefere")
	respRecorder := httptest.NewRecorder()

	// when
	Headers(respRecorder, req)

	// then
	if respRecorder.Code != http.StatusOK {
		t.Errorf("Invalid response status (expected: %d, actual: %d)", http.StatusOK, respRecorder.Code)
	}

	respContentType := respRecorder.Header().Get("Content-Type")
	if respContentType != "application/json" {
		t.Errorf("Invalid response content type (expected: 'application/jso'n, actual: '%s')", respContentType)
	}

	var respJson map[string][]string
	respBody := respRecorder.Body.Bytes()
	if err := json.Unmarshal(respBody, &respJson); err != nil {
		t.Errorf("Unable to unmarshal response. Root cause: %s", err.Error())
		return
	}

	checkHeaderValue := func(headerName, expectedValue string) {
		if (len(respJson[headerName]) != 1) || (respJson[headerName][0] != expectedValue) {
			t.Errorf("Invalid echoed %s (expected: '%s', actual: '%v')", headerName, expectedValue, respJson[headerName])
		}
	}

	checkHeaderValue("Content-Type", "text/html")
	checkHeaderValue("Authorization", "Bearer terefere")
}