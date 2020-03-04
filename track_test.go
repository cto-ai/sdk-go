package ctoai

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_TrackRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Content-Type"][0] != "application/json" {
			t.Errorf("Headers incorrect: %v", r.Header["Content-Type"])
		}

		if r.Method != "POST" {
			t.Errorf("Method incorrect: %v", r.Method)
		}

		if r.URL.Path != "/track" {
			t.Errorf("Method incorrect: %v", r.URL.Path)
		}

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

	_, port, err := net.SplitHostPort(ts.URL[7:])
	if err != nil {
		t.Errorf("Error splitting host port: %s", err)
	}

	err = os.Setenv("SDK_SPEAK_PORT", port)
	if err != nil {
		t.Errorf("Error setting test env variable: %s", err)
	}

	s := NewSdk()
	s.Track([]string{"tag1", "tag2"}, "testEvent", map[string]interface{}{
		"testKey": "testValue",
	})
}
