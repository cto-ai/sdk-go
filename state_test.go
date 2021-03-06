package ctoai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetState(t *testing.T) {
	expectedResponse := `{"value": "state-value"}`
	expectedBody := map[string]interface{}{
		"key": "test-key",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/state/get")

		var tmp map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if !reflect.DeepEqual(tmp, expectedBody) {
			t.Errorf("Error unexpected request body: %+v", tmp)
		}

		fmt.Fprintf(w, expectedResponse)
	}))

	defer ts.Close()

	SetPortVar(t, ts)

	s := NewSdk()
	output, err := s.GetState("test-key")
	if err != nil {
		t.Errorf("Error in state request: %v", err)
	}

	if output != "state-value" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func TestSetState(t *testing.T) {
	expectedBody := map[string]interface{}{
		"key":   "key-of-value",
		"value": "value-of-key",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/state/set")

		var tmp map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if !reflect.DeepEqual(tmp, expectedBody) {
			t.Errorf("Error unexpected request body: %+v", tmp)
		}
	}))

	defer ts.Close()

	SetPortVar(t, ts)

	s := NewSdk()
	err := s.SetState("key-of-value", "value-of-key")
	if err != nil {
		t.Errorf("Error in state request: %v", err)
	}
}
