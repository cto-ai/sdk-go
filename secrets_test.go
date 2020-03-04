package ctoai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func Test_SecretsRequest_GetSecret(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Content-Type"][0] != "application/json" {
			t.Errorf("Headers incorrect: %v", r.Header["Content-Type"])
		}

		if r.Method != "POST" {
			t.Errorf("Method incorrect: %v", r.Method)
		}

		if r.URL.Path != "/secret/get" {
			t.Errorf("Method incorrect: %v", r.URL.Path)
		}

		var tmp daemon.GetSecretBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Key != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		fmt.Fprintf(w, expectedResponse)
	}))

	defer ts.Close()

	// write a fake file
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": "secret"}`), 0777)

	_, port, err := net.SplitHostPort(ts.URL[7:])
	if err != nil {
		t.Errorf("Error splitting host port: %s", err)
	}

	err = os.Setenv("SDK_SPEAK_PORT", port)
	if err != nil {
		t.Errorf("Error setting test env variable: %s", err)
	}

	s := NewSdk()
	output, err := s.GetSecret("test")
	if err != nil {
		t.Errorf("Error in secret request: %v", err)
	}

	if output != "secret" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_SecretsRequest_SetSecret(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Content-Type"][0] != "application/json" {
			t.Errorf("Headers incorrect: %v", r.Header["Content-Type"])
		}

		if r.Method != "POST" {
			t.Errorf("Method incorrect: %v", r.Method)
		}

		if r.URL.Path != "/secret/set" {
			t.Errorf("Method incorrect: %v", r.URL.Path)
		}

		var tmp daemon.SetSecretBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Key != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		fmt.Fprintf(w, expectedResponse)
	}))

	defer ts.Close()

	// write a fake file
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"key": "test"}`), 0777)

	_, port, err := net.SplitHostPort(ts.URL[7:])
	if err != nil {
		t.Errorf("Error splitting host port: %s", err)
	}

	err = os.Setenv("SDK_SPEAK_PORT", port)
	if err != nil {
		t.Errorf("Error setting test env variable: %s", err)
	}

	s := NewSdk()
	key, err := s.SetSecret("test", "secret")
	if err != nil {
		t.Errorf("Error in secret request: %v", err)
	}

	if key != "test" {
		t.Errorf("Error unexpected returned key: %v", key)
	}
}
