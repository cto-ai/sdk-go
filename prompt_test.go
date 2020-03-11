package ctoai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func Test_PromptRequest_PromptInput(t *testing.T) {
	expectedResponse := `{"replyFilename": "/tmp/response-mocktest"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.InputPromptBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Name != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.PromptType != "input" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Message != "type test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Default != "error" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Flag != "I" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.AllowEmpty != true {
			t.Errorf("Error unexpected request body: %v", tmp)
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp["name"] != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["type"] != "number" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["message"] != "type 2" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["flag"] != "N" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["default"] != float64(0) {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["minimum"] != float64(1) {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if _, ok := tmp["maximum"]; ok {
			t.Errorf("Error unexpected request body: %v", tmp)
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.SecretPromptBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Name != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.PromptType != "secret" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Message != "what is secret" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Flag != "S" {
			t.Errorf("Error unexpected request body: %v", tmp)
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.PasswordPromptBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Name != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.PromptType != "password" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Message != "what is password" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Flag != "P" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Confirm != true {
			t.Errorf("Error unexpected request body: %v", tmp)
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.ConfirmPromptBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Name != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.PromptType != "confirm" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Message != "confirm?" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Flag != "C" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Default != true {
			t.Errorf("Error unexpected request body: %v", tmp)
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp["name"] != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["type"] != "list" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["message"] != "choose" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["choices"].([]interface{})[0] != choices[0] {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["choices"].([]interface{})[1] != choices[1] {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["default"] != "aws" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["flag"] != "L" {
			t.Errorf("Error unexpected request body: %v", tmp)
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp["name"] != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["type"] != "autocomplete" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["message"] != "choose" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["choices"].([]interface{})[0] != choices[0] {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["choices"].([]interface{})[1] != choices[1] {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["default"] != 1.0 {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["flag"] != "A" {
			t.Errorf("Error unexpected request body: %v", tmp)
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp["name"] != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["type"] != "checkbox" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["message"] != "choose" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["flag"] != "C" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["choices"].([]interface{})[0] != choices[0] {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["choices"].([]interface{})[1] != choices[1] {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp["choices"].([]interface{})[2] != choices[2] {
			t.Errorf("Error unexpected request body: %v", tmp)
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.EditorPromptBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Name != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.PromptType != "editor" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Message != "edit" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Flag != "E" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Default != "default" {
			t.Errorf("Error unexpected request body: %v", tmp)
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ValidateRequest(t, r, "/prompt")

		var tmp daemon.DatetimePromptBody
		err := json.NewDecoder(r.Body).Decode(&tmp)
		if err != nil {
			t.Errorf("Error in decoding response body: %s", err)
		}

		if tmp.Name != "test" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.PromptType != "datetime" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Message != "what date" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Variant != "datetime" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Flag != "D" {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Minimum != input {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Maximum != input {
			t.Errorf("Error unexpected request body: %v", tmp)
		}

		if tmp.Default != input {
			t.Errorf("Error unexpected request body: %v", tmp)
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
