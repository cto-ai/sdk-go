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

func Test_SpinnerRequest_spinnerStart(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Content-Type"][0] != "application/json" {
			t.Errorf("Headers incorrect: %v", r.Header["Content-Type"])
		}

		if r.Method != "POST" {
			t.Errorf("Method incorrect: %v", r.Method)
		}

		if r.URL.Path != "/start-spinner" {
			t.Errorf("Method incorrect: %v", r.URL.Path)
		}

		var tmp daemon.SpinnerStartBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Text != "start" {
			t.Errorf("Error unexpected request body: %v", tmp)
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
	err = u.SpinnerStart("start")
	if err != nil {
		t.Errorf("Error starting spinner: %v", err)
	}
}

func Test_SpinnerRequest_spinnerStop(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Content-Type"][0] != "application/json" {
			t.Errorf("Headers incorrect: %v", r.Header["Content-Type"])
		}

		if r.Method != "POST" {
			t.Errorf("Method incorrect: %v", r.Method)
		}

		if r.URL.Path != "/stop-spinner" {
			t.Errorf("Method incorrect: %v", r.URL.Path)
		}

		var tmp daemon.SpinnerStopBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Text != "stop" {
			t.Errorf("Error unexpected request body: %v", tmp)
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
	err = u.SpinnerStop("stop")
	if err != nil {
		t.Errorf("Error stopping spinner: %v", err)
	}
}
