package ctoai

import (
	"encoding/json"
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
