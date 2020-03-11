package ctoai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func Test_ProgressbarRequest_ProgressbarStart(t *testing.T) {
	expectedBody := daemon.ProgressBarStartBody{
		Text:    "start",
		Length:  2,
		Initial: 1,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/progress-bar/start")

		var tmp daemon.ProgressBarStartBody
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

	u := NewUx()
	err := u.ProgressBarStart(2, 1, "start")
	if err != nil {
		t.Errorf("Error starting progress bar: %v", err)
	}

}

func Test_ProgressbarRequest_ProgressbarAdvance(t *testing.T) {
	expectedBody := daemon.ProgressBarAdvanceBody{
		Increment: 1,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/progress-bar/advance")

		var tmp daemon.ProgressBarAdvanceBody
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

	u := NewUx()
	err := u.ProgressBarAdvance(1)
	if err != nil {
		t.Errorf("Error advancing progress bar: %v", err)
	}
}

func Test_ProgressbarRequest_ProgressbarStop(t *testing.T) {
	expectedBody := daemon.ProgressBarStopBody{
		Text: "Done",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/progress-bar/stop")

		var tmp daemon.ProgressBarStopBody
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

	u := NewUx()
	err := u.ProgressBarStop("Done")
	if err != nil {
		t.Errorf("Error stopping progress bar: %v", err)
	}
}
