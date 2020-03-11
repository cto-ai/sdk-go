package ctoai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func Test_ProgressbarRequest_ProgressbarStart(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/progress-bar/start")

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

	SetPortVar(t, ts)

	u := NewUx()
	err := u.ProgressBarStart(2, 1, "start")
	if err != nil {
		t.Errorf("Error starting progress bar: %v", err)
	}

}

func Test_ProgressbarRequest_ProgressbarAdvance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/progress-bar/advance")

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

	SetPortVar(t, ts)

	u := NewUx()
	err := u.ProgressBarAdvance(1)
	if err != nil {
		t.Errorf("Error advancing progress bar: %v", err)
	}
}

func Test_ProgressbarRequest_ProgressbarStop(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/progress-bar/stop")

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

	SetPortVar(t, ts)

	u := NewUx()
	err := u.ProgressBarStop("Done")
	if err != nil {
		t.Errorf("Error stopping progress bar: %v", err)
	}
}
