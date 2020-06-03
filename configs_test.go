package ctoai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	expectedResponse := `{"value": "config-value"}`
	expectedBody := map[string]interface{}{
		"key": "test-key",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/config/get")

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
	output, err := s.GetConfig("test-key")
	if err != nil {
		t.Errorf("Error in config request: %v", err)
	}

	if output != "config-value" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func TestGetConfig_Null(t *testing.T) {
	expectedResponse := `{"value": null}`
	expectedBody := map[string]interface{}{
		"key": "test-key",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/config/get")

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
	output, err := s.GetConfig("test-key")
	if err != nil {
		t.Errorf("Error in config request: %v", err)
	}

	if output != "" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func TestSetConfig(t *testing.T) {
	expectedBody := map[string]interface{}{
		"key":   "key-of-value",
		"value": "value-of-key",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/config/set")

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
	err := s.SetConfig("key-of-value", "value-of-key")
	if err != nil {
		t.Errorf("Error in config request: %v", err)
	}
}

func TestDeleteConfig(t *testing.T) {
	expectedResponse := `{"value": true}`
	expectedBody := map[string]interface{}{
		"key": "test-key",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/config/delete")

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
	output, err := s.DeleteConfig("test-key")
	if err != nil {
		t.Errorf("Error in config request: %v", err)
	}

	if output != true {
		t.Errorf("Error unexpected output: %v", output)
	}
}
