package ctoai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func Test_SecretsRequest_GetSecret(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`
	expectedBody := daemon.GetSecretBody{
		Key: "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/secret/get")

		var tmp daemon.GetSecretBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if !reflect.DeepEqual(tmp, expectedBody) {
			t.Errorf("Error unexpected request body: %+v", tmp)
		}

		fmt.Fprintf(w, expectedResponse)
	}))

	defer ts.Close()

	// write a fake file
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": "secret"}`), 0777)

	SetPortVar(t, ts)

	s := NewSdk()
	output, err := s.GetSecret("test")
	if err != nil {
		t.Errorf("Error in secret request: %v", err)
	}

	if output != "secret" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_SecretsRequest_GetSecretHidden(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`
	expectedBody := daemon.GetSecretBody{
		Key:    "test",
		Hidden: true,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/secret/get")

		var tmp daemon.GetSecretBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if !reflect.DeepEqual(tmp, expectedBody) {
			t.Errorf("Error unexpected request body: %+v", tmp)
		}

		fmt.Fprintf(w, expectedResponse)
	}))

	defer ts.Close()

	// write a fake file
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": "secret"}`), 0777)

	SetPortVar(t, ts)

	s := NewSdk()
	output, err := s.GetSecret("test", OptGetSecretHidden())
	if err != nil {
		t.Errorf("Error in secret request: %v", err)
	}

	if output != "secret" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_SecretsRequest_SetSecret(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`
	expectedBody := daemon.SetSecretBody{
		Key:   "test",
		Value: "secret",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/secret/set")

		var tmp daemon.SetSecretBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if !reflect.DeepEqual(tmp, expectedBody) {
			t.Errorf("Error unexpected request body: %+v", tmp)
		}

		fmt.Fprintf(w, expectedResponse)
	}))

	defer ts.Close()

	// write a fake file
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"key": "test"}`), 0777)

	SetPortVar(t, ts)

	s := NewSdk()
	key, err := s.SetSecret("test", "secret")
	if err != nil {
		t.Errorf("Error in secret request: %v", err)
	}

	if key != "test" {
		t.Errorf("Error unexpected returned key: %v", key)
	}
}
