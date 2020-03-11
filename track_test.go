package ctoai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_TrackRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/track")

		var tmp map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp["event"] != "testEvent" {
			t.Errorf("Error unexpected request body: %+v", tmp)
		}

		if tmp["tag"].([]interface{})[0] != "tag1" {
			t.Errorf("Error unexpected request body: %+v", tmp)
		}

		if tmp["tag"].([]interface{})[1] != "tag2" {
			t.Errorf("Error unexpected request body: %+v", tmp)
		}

		if tmp["testKey"] != "testValue" {
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
