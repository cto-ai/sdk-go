package ctoai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func Test_SpinnerRequest_spinnerStart(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/start-spinner")

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

	SetPortVar(t, ts)

	u := NewUx()
	err := u.SpinnerStart("start")
	if err != nil {
		t.Errorf("Error starting spinner: %v", err)
	}
}

func Test_SpinnerRequest_spinnerStop(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/stop-spinner")

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

	SetPortVar(t, ts)

	u := NewUx()
	err := u.SpinnerStop("stop")
	if err != nil {
		t.Errorf("Error stopping spinner: %v", err)
	}
}
