package ctoai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func Test_PromptRequest_PromptInput(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`
	expectedBody := daemon.InputPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       "test",
			PromptType: "input",
			Message:    "type test",
			Flag:       "I",
		},
		Default:    "error",
		AllowEmpty: true,
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.InputPromptBody
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
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": "test"}`), 0777)

	SetPortVar(t, ts)

	p := NewPrompt()
	output, err := p.Input("test", "type test", OptInputDefault("error"), OptInputFlag("I"), OptInputAllowEmpty(true))
	if err != nil {
		t.Errorf("Error in prompt request: %v", err)
	}

	if output != "test" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_PromptRequest_PromptNumber(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`
	expectedBody := map[string]interface{}{
		"name":    "test",
		"type":    "number",
		"message": "type 2",
		"flag":    "N",
		"default": float64(0),
		"minimum": float64(1),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp map[string]interface{}
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
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": 2}`), 0777)

	SetPortVar(t, ts)

	p := NewPrompt()
	output, err := p.Number("test", "type 2", OptNumberFlag("N"), OptNumberDefault(0), OptNumberMinimum(1))
	if err != nil {
		t.Errorf("Error in prompt request: %v", err)
	}

	if output != 2 {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_PromptRequest_PromptSecret(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`
	expectedBody := daemon.SecretPromptBody{
		Name:       "test",
		PromptType: "secret",
		Message:    "what is secret",
		Flag:       "S",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.SecretPromptBody
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

	p := NewPrompt()
	output, err := p.Secret("test", "what is secret", OptSecretFlag("S"))
	if err != nil {
		t.Errorf("Error in prompt request: %v", err)
	}

	if output != "secret" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_PromptRequest_PromptPassword(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`
	expectedBody := daemon.PasswordPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       "test",
			PromptType: "password",
			Message:    "what is password",
			Flag:       "P",
		},
		Confirm: true,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.PasswordPromptBody
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
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": "password"}`), 0777)

	SetPortVar(t, ts)

	p := NewPrompt()
	output, err := p.Password("test", "what is password", OptPasswordFlag("P"), OptPasswordConfirm(true))
	if err != nil {
		t.Errorf("Error in prompt request: %v", err)
	}

	if output != "password" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_PromptRequest_PromptConfirm(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`
	expectedBody := daemon.ConfirmPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       "test",
			PromptType: "confirm",
			Message:    "confirm?",
			Flag:       "C",
		},
		Default: true,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.ConfirmPromptBody
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
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": false}`), 0777)

	SetPortVar(t, ts)

	p := NewPrompt()
	output, err := p.Confirm("test", "confirm?", OptConfirmFlag("C"), OptConfirmDefault(true))
	if err != nil {
		t.Errorf("Error in prompt request: %v", err)
	}

	if output != false {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_PromptRequest_PromptList(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`

	choices := []string{"aws", "gcd"}
	expectedBody := map[string]interface{}{
		"name":    "test",
		"type":    "list",
		"message": "choose",
		"choices": []interface{}{"aws", "gcd"},
		"default": "aws",
		"flag":    "L",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp map[string]interface{}
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
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": "gcd"}`), 0777)

	SetPortVar(t, ts)

	p := NewPrompt()
	output, err := p.List("test", "choose", choices, OptListDefaultValue("aws"), OptListFlag("L"))
	if err != nil {
		t.Errorf("Error in prompt request: %v", err)
	}

	if output != "gcd" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_PromptRequest_PromptAutocomplete(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`

	choices := []string{"aws", "gcd"}
	expectedBody := map[string]interface{}{
		"name":    "test",
		"type":    "autocomplete",
		"message": "choose",
		"choices": []interface{}{"aws", "gcd"},
		"default": 1.0,
		"flag":    "A",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp map[string]interface{}
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
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": "gcd"}`), 0777)

	SetPortVar(t, ts)

	p := NewPrompt()
	output, err := p.List("test", "choose", choices, OptListAutocomplete(true), OptListDefaultIndex(1), OptListFlag("A"))
	if err != nil {
		t.Errorf("Error in prompt request: %v", err)
	}

	if output != "gcd" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_PromptRequest_PromptCheckbox(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`

	choices := []string{"aws", "gcd", "azure"}
	expectedBody := map[string]interface{}{
		"name":    "test",
		"type":    "checkbox",
		"message": "choose",
		"choices": []interface{}{"aws", "gcd", "azure"},
		"flag":    "C",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp map[string]interface{}
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
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": ["gcd", "azure"]}`), 0777)

	SetPortVar(t, ts)

	p := NewPrompt()
	output, err := p.Checkbox("test", "choose", choices, OptCheckboxFlag("C"))
	if err != nil {
		t.Errorf("Error in prompt request: %v", err)
	}

	if output[0] != "gcd" {
		t.Errorf("Error unexpected output: %v", output[0])
	}

	if output[1] != "azure" {
		t.Errorf("Error unexpected output: %v", output[1])
	}
}

func Test_PromptRequest_PromptEditor(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`
	expectedBody := daemon.EditorPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       "test",
			PromptType: "editor",
			Message:    "edit",
			Flag:       "E",
		},
		Default: "default",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.EditorPromptBody
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
	err := ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": "test"}`), 0777)

	SetPortVar(t, ts)

	p := NewPrompt()
	output, err := p.Editor("test", "edit", OptEditorDefault("default"), OptEditorFlag("E"))
	if err != nil {
		t.Errorf("Error in prompt request: %v", err)
	}

	if output != "test" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_PromptRequest_PromptDatetime(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`

	expectedRequest := "2006-01-02T15:04:05Z"
	input, err := time.Parse(time.RFC3339, expectedRequest)
	if err != nil {
		t.Errorf("Error parsing expected time: %v", err)
	}
	expectedBody := daemon.DatetimePromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       "test",
			PromptType: "datetime",
			Message:    "what date",
			Flag:       "D",
		},
		Variant: "datetime",
		Default: input,
		Minimum: input,
		Maximum: input,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.DatetimePromptBody
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
	err = ioutil.WriteFile("/tmp/response-mocktest", []byte(`{"test": "2006-01-02T15:04:05Z"}`), 0777)

	SetPortVar(t, ts)

	p := NewPrompt()
	output, err := p.Datetime("test", "what date", OptDatetimeFlag("D"), OptDatetimeDefault(input), OptDatetimeMaximum(input), OptDatetimeMinimum(input))
	if err != nil {
		t.Errorf("Error in prompt request: %v", err)
	}

	const format = "2006-01-02 15:04:05"

	expected, err := time.Parse(format, "2006-01-02 15:04:05")
	if err != nil {
		t.Errorf("Error parsing expected time: %v", err)
	}

	if output != expected {
		t.Errorf("Error unexpected output: %v", output)
	}
}
