package ctoai

import (
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func ValidateRequest(t *testing.T, r *http.Request, path string) {
	if r.Header["Content-Type"][0] != "application/json" {
		t.Errorf("Headers incorrect: %v", r.Header["Content-Type"])
	}

	if r.Method != "POST" {
		t.Errorf("Method incorrect: %v", r.Method)
	}

	if r.URL.Path != path {
		t.Errorf("Method incorrect: %v", r.URL.Path)
	}
}

func SetPortVar(t *testing.T, ts *httptest.Server) {
	_, port, err := net.SplitHostPort(ts.URL[7:])
	if err != nil {
		t.Errorf("Error splitting host port: %s", err)
	}

	err = os.Setenv("SDK_SPEAK_PORT", port)
	if err != nil {
		t.Errorf("Error setting test env variable: %s", err)
	}
}
