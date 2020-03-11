package ctoai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func Test_PrintRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/print")

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

	SetPortVar(t, ts)

	u := NewUx()
	err := u.Print("test")
	if err != nil {
		t.Errorf("Error printing test value: %v", err)
	}
}
