package ctoai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_TrackRequest(t *testing.T) {
	expectedBody := map[string]interface{}{
		"event":   "testEvent",
		"tags":    []interface{}{"tag1", "tag2"},
		"testKey": "testValue",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/track")

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
	err := s.Track([]string{"tag1", "tag2"}, "testEvent", map[string]interface{}{
		"testKey": "testValue",
	})
	if err != nil {
		t.Errorf("Error running track: %v", err)
	}
}

func TestEvents_Success(t *testing.T) {
	expectedResponse := `{"value": [{"event": "test"}, {"data": "code"}]}`
	expectedBody := map[string]interface{}{
		"start": "2020-01-01",
		"end":   "2020-06-01",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/events")

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
	output, err := s.Events("2020-01-01", "2020-06-01")
	if err != nil {
		t.Errorf("Error in config request: %v", err)
	}

	if !reflect.DeepEqual(output, []map[string]interface{}{
		map[string]interface{}{
			"event": "test",
		},
		map[string]interface{}{
			"data": "code",
		},
	}) {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func TestEvents_NoArray(t *testing.T) {
	expectedResponse := `{"value": "result"}`
	expectedBody := map[string]interface{}{
		"start": "2020-01-01",
		"end":   "2020-06-01",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/events")

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
	_, err := s.Events("2020-01-01", "2020-06-01")
	if err == nil {
		t.Errorf("No error in config request")
	}
}

func TestEvents_NotObjects(t *testing.T) {
	expectedResponse := `{"value": ["result"]}`
	expectedBody := map[string]interface{}{
		"start": "2020-01-01",
		"end":   "2020-06-01",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/events")

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
	_, err := s.Events("2020-01-01", "2020-06-01")
	if err == nil {
		t.Errorf("No error in config request")
	}
}
