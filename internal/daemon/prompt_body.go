package daemon

import (
	"encoding/json"
	"time"
)

// PromptEnvelope is the common fields for all prompt type JSON bodies
type PromptEnvelope struct {
	Name       string `json:"name"`
	PromptType string `json:"type"`
	Message    string `json:"message"`
	Flag       string `json:"flag,omitempty"`
}

// InputPromptBody is the JSON body for an input prompt
type InputPromptBody struct {
	PromptEnvelope
	Default    string `json:"default,omitempty"`
	AllowEmpty bool   `json:"allowEmpty"`
}

// NumberPromptBody is the JSON body for a number prompt
type NumberPromptBody struct {
	PromptEnvelope
	DefaultValue int
	DefaultSet   bool
	MaximumValue int
	MaximumSet   bool
	MinimumValue int
	MinimumSet   bool
}

// MarshalJSON marshals the correct set of fields for the NumberPromptBody
func (n NumberPromptBody) MarshalJSON() ([]byte, error) {
	output := make(map[string]interface{})
	output["name"] = n.Name
	output["type"] = "number"
	output["message"] = n.Message

	if n.Flag != "" {
		output["flag"] = n.Flag
	}

	if n.DefaultSet {
		output["default"] = n.DefaultValue
	}
	if n.MaximumSet {
		output["maximum"] = n.MaximumValue
	}
	if n.MinimumSet {
		output["minimum"] = n.MinimumValue
	}

	return json.Marshal(output)
}

// SecretPromptBody is the JSON body for a secret prompt
type SecretPromptBody = PromptEnvelope

// PasswordPromptBody is the JSON body for a password prompt
type PasswordPromptBody struct {
	PromptEnvelope
	Confirm bool `json:"confirm"`
}

// ConfirmPromptBody is the JSON body for a confirm prompt
type ConfirmPromptBody struct {
	PromptEnvelope
	Default bool `json:"default"`
}

// ListPromptBody is the JSON body for a list or autocomplete prompt
type ListPromptBody struct {
	PromptEnvelope
	Choices        []string
	DefaultSet     bool
	DefaultIndex   int
	DefaultValue   string
	DefaultIsValue bool
}

// MarshalJSON marshals the correct set of fields for the ListPromptBody
func (n ListPromptBody) MarshalJSON() ([]byte, error) {
	output := make(map[string]interface{})
	output["name"] = n.Name
	output["type"] = n.PromptType
	output["message"] = n.Message
	output["choices"] = n.Choices

	if n.Flag != "" {
		output["flag"] = n.Flag
	}

	if n.DefaultSet {
		if n.DefaultIsValue {
			output["default"] = n.DefaultValue
		} else {
			output["default"] = n.DefaultIndex
		}
	}

	return json.Marshal(output)
}

// CheckboxPromptBody is the JSON body for a checkbox prompt
type CheckboxPromptBody struct {
	PromptEnvelope
	Choices        []string
	DefaultSet     bool
	DefaultIndex   []int
	DefaultValue   []string
	DefaultIsValue bool
}

// MarshalJSON marshals the correct set of fields for the CheckboxPromptBody
func (n CheckboxPromptBody) MarshalJSON() ([]byte, error) {
	output := make(map[string]interface{})
	output["name"] = n.Name
	output["type"] = n.PromptType
	output["message"] = n.Message
	output["choices"] = n.Choices

	if n.Flag != "" {
		output["flag"] = n.Flag
	}

	if n.DefaultSet {
		if n.DefaultIsValue {
			output["default"] = n.DefaultValue
		} else {
			output["default"] = n.DefaultIndex
		}
	}

	return json.Marshal(output)
}

// EditorPromptBody is the JSON body for a editor prompt
type EditorPromptBody struct {
	PromptEnvelope
	Default string `json:"default"`
}

// DatetimePromptBody is the JSON body for a datetime prompt
type DatetimePromptBody struct {
	PromptEnvelope
	Variant string    `json:"variant"`
	Default time.Time `json:"default,omitempty"`
	Maximum time.Time `json:"maximum,omitempty"`
	Minimum time.Time `json:"minimum,omitempty"`
}
