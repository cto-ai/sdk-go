package ctoai

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func Test_PrintRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Content-Type"][0] != "application/json" {
			t.Errorf("Headers incorrect: %v", r.Header["Content-Type"])
		}

		if r.Method != "POST" {
			t.Errorf("Method incorrect: %v", r.Method)
		}

		if r.URL.Path != "/print" {
			t.Errorf("Method incorrect: %v", r.URL.Path)
		}

		var tmp daemon.PrintBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Text != "test" {
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

	u := NewUx()
	err = u.Print("test")
	if err != nil {
		t.Errorf("Error printing test value: %v", err)
	}
}
