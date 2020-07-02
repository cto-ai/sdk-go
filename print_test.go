package ctoai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func TestItalic(t *testing.T) {
	err := os.Setenv("SDK_INTERFACE_TYPE", "terminal")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ux := NewUx()
	result := ux.Italic("test string")
	expected := "\033[3mtest string\033[23m"
	if result != expected {
		t.Fatalf("in terminal expected %s, got %s", expected, result)
	}

	err = os.Setenv("SDK_INTERFACE_TYPE", "slack")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result = ux.Italic("test string")
	expected = "_test string_"
	if result != expected {
		t.Fatalf("in Slack expected %s, got %s", expected, result)
	}
}

func TestBold(t *testing.T) {
	err := os.Setenv("SDK_INTERFACE_TYPE", "terminal")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ux := NewUx()
	result := ux.Bold("test string")
	expected := "\033[1mtest string\033[0m"
	if result != expected {
		t.Fatalf("in terminal expected %s, got %s", expected, result)
	}

	err = os.Setenv("SDK_INTERFACE_TYPE", "slack")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result = ux.Bold("test string")
	expected = "*test string*"
	if result != expected {
		t.Fatalf("in Slack expected %s, got %s", expected, result)
	}
}

func Test_PrintRequest(t *testing.T) {
	expectedBody := daemon.PrintBody{
		Text: "test",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/print")

		var tmp daemon.PrintBody
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
	err := u.Print("test")
	if err != nil {
		t.Errorf("Error printing test value: %v", err)
	}
}
