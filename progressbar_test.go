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

func Test_ProgressbarRequest_ProgressbarStart(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Content-Type"][0] != "application/json" {
			t.Errorf("Headers incorrect: %v", r.Header["Content-Type"])
		}

		if r.Method != "POST" {
			t.Errorf("Method incorrect: %v", r.Method)
		}

		if r.URL.Path != "/progress-bar/start" {
			t.Errorf("Method incorrect: %v", r.URL.Path)
		}

		var tmp daemon.ProgressBarStartBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Text != "start" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Length != 2 {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Initial != 1 {
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
	err = u.ProgressBarStart(2, 1, "start")
	if err != nil {
		t.Errorf("Error starting progress bar: %v", err)
	}

}

func Test_ProgressbarRequest_ProgressbarAdvance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Content-Type"][0] != "application/json" {
			t.Errorf("Headers incorrect: %v", r.Header["Content-Type"])
		}

		if r.Method != "POST" {
			t.Errorf("Method incorrect: %v", r.Method)
		}

		if r.URL.Path != "/progress-bar/advance" {
			t.Errorf("Method incorrect: %v", r.URL.Path)
		}

		var tmp daemon.ProgressBarAdvanceBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Increment != 1 {
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
	err = u.ProgressBarAdvance(1)
	if err != nil {
		t.Errorf("Error advancing progress bar: %v", err)
	}
}

func Test_ProgressbarRequest_ProgressbarStop(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Content-Type"][0] != "application/json" {
			t.Errorf("Headers incorrect: %v", r.Header["Content-Type"])
		}

		if r.Method != "POST" {
			t.Errorf("Method incorrect: %v", r.Method)
		}

		if r.URL.Path != "/progress-bar/stop" {
			t.Errorf("Method incorrect: %v", r.URL.Path)
		}

		var tmp daemon.ProgressBarStopBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Text != "Done" {
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
	err = u.ProgressBarStop("Done")
	if err != nil {
		t.Errorf("Error stopping progress bar: %v", err)
	}
}
