package ctoai

import (
	"fmt"
	"time"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

// Prompt is the object that contains the prompt methods
type Prompt struct{}

func NewPrompt() *Prompt {
	return &Prompt{}
}

// InputOption is an option for the Input prompt function
type InputOption func(*daemon.InputPromptBody)

// OptInputFlag sets the flag value for the input prompt.
//
// The flag value is used to match command line arguments to prompts.
func OptInputFlag(flag string) InputOption {
	return func(definition *daemon.InputPromptBody) {
		definition.Flag = flag
	}
}

// OptInputDefault sets the default value for the input prompt.
func OptInputDefault(defaultValue string) InputOption {
	return func(definition *daemon.InputPromptBody) {
		definition.Default = defaultValue
	}
}

// OptInputAllowEmpty sets whether the input prompt should accept an empty line.
//
// Has no effect if a default is set.
func OptInputAllowEmpty(allowEmpty bool) InputOption {
	return func(definition *daemon.InputPromptBody) {
		definition.AllowEmpty = allowEmpty
	}
}

// Input presents an input (single-line text) prompt on the interface
// (i.e. terminal or slack).
//
// The method returns the user's response as string.
//
// Example:
//
//  p := ctoai.NewPrompt()
//  resp, err := p.Input("opinion" ,"What do you think of Go?", "good") // user responds with "good"
//  if err != nil {
//      panic(err)
//  }
//
//  fmt.Println(resp)
//
// Output:
// good
func (*Prompt) Input(name, msg string, options ...InputOption) (string, error) {
	definition := daemon.InputPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       name,
			PromptType: "input",
			Message:    msg,
		},
	}
	for _, option := range options {
		option(&definition)
	}

	body, err := daemon.AsyncRequest("prompt", definition)
	if err != nil {
		return "", err
	}

	if value, ok := body[name]; ok {
		if str, ok := value.(string); ok {
			return str, nil
		}
		return "", fmt.Errorf("Daemon returned non-string value %v", value)
	}
	return "", fmt.Errorf("Daemon returned incorrect JSON %v", body)
}

// NumberOption is a functional option type for the Number method.
type NumberOption func(*daemon.NumberPromptBody)

// OptNumberFlag sets the flag value for the number prompt.
//
// The flag value is used to match command line arguments to prompts.
func OptNumberFlag(flag string) NumberOption {
	return func(definition *daemon.NumberPromptBody) {
		definition.Flag = flag
	}
}

func OptNumberDefault(defaultValue int) NumberOption {
	return func(definition *daemon.NumberPromptBody) {
		definition.DefaultValue = defaultValue
		definition.DefaultSet = true
	}
}

func OptNumberMaximum(maximumValue int) NumberOption {
	return func(definition *daemon.NumberPromptBody) {
		definition.MaximumValue = maximumValue
		definition.MaximumSet = true
	}
}

func OptNumberMinimum(minimumValue int) NumberOption {
	return func(definition *daemon.NumberPromptBody) {
		definition.MinimumValue = minimumValue
		definition.MinimumSet = true
	}
}

// Number presents a prompt for a numeric value to the interface
// (i.e. terminal or slack).
//
// The method returns the user's response as int.
//
// Example:
//
//  p := ctoai.NewPrompt()
//  resp, err := p.Number("count", "How many hoorays do you want?", OptNumberFlag("n"), OptNumberDefault(10), OptNumberMinimum(0), OptNumberMaximum(10)) // user responds with 7
//  if err != nil {
//      panic(err)
//  }
//
//  fmt.Println(resp)
//
// Output:
// 7
func (*Prompt) Number(name, msg string, options ...NumberOption) (int, error) {
	definition := daemon.NumberPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:    name,
			Message: msg,
		},
	}
	for _, option := range options {
		option(&definition)
	}

	body, err := daemon.AsyncRequest("prompt", definition)
	if err != nil {
		return 0, err
	}

	if value, ok := body[name]; ok {
		if num, ok := value.(float64); ok {
			return int(num), nil
		}
		return 0, fmt.Errorf("Daemon returned non-numeric value %v", value)
	}
	return 0, fmt.Errorf("Daemon returned incorrect JSON %v", body)
}

type SecretOption func(*daemon.SecretPromptBody)

// OptSecretFlag sets the flag value for the secret prompt.
//
// The flag value is used to match command line arguments to prompts.
func OptSecretFlag(flag string) SecretOption {
	return func(definition *daemon.SecretPromptBody) {
		definition.Flag = flag
	}
}

// Secret presents an input prompt for secrets in the interface
// (i.e. terminal or slack).
//
// The secret can be entered by the user or selected from their team's
// secret store. The value will be displayed on the terminal/screen as
// it is entered for verification
//
// For secrets, the name doubles as the default secret store key for
// the desired secret. If this key is present, it will be selected as
// the default option for the user.
//
// The method returns the user's response as string.
//
// Example:
//
//  p := ctoai.NewPrompt()
//  resp, err := p.Secret("SSH_KEY", "What SSH private key do you want to use?", OptSecretFlag("s")) // user responds with 1234asdf
//  if err != nil {
//      panic(err)
//  }
//
//  fmt.Println(resp)
//
// Output:
// 1234asdf
func (*Prompt) Secret(name, msg string, options ...SecretOption) (string, error) {
	definition := daemon.SecretPromptBody{
		Name:       name,
		PromptType: "secret",
		Message:    msg,
	}
	for _, option := range options {
		option(&definition)
	}

	body, err := daemon.AsyncRequest("prompt", definition)
	if err != nil {
		return "", err
	}

	if value, ok := body[name]; ok {
		if str, ok := value.(string); ok {
			return str, nil
		}
		return "", fmt.Errorf("Daemon returned non-string value %v", value)
	}
	return "", fmt.Errorf("Daemon returned incorrect JSON %v", body)
}

type PasswordOption func(*daemon.PasswordPromptBody)

// OptPasswordFlag sets the flag value for the password prompt.
//
// The flag value is used to match command line arguments to prompts.
func OptPasswordFlag(flag string) PasswordOption {
	return func(definition *daemon.PasswordPromptBody) {
		definition.Flag = flag
	}
}

func OptPasswordConfirm(confirm bool) PasswordOption {
	return func(definition *daemon.PasswordPromptBody) {
		definition.Confirm = confirm
	}
}

// Password presents an input prompt for passwords in the interface
// (i.e. terminal or slack).
//
// The password can be entered by the user or selected from their
// team's secret store. If the value is entered directly, it will be
// obscured on the screen.
//
// For passwords, the name doubles as the default secret store key for
// the desired secret. If this key is present, it will be selected as
// the default option for the user.
//
// The method returns the user's response as string.
//
// Example:
//
//  p := ctoai.NewPrompt()
//  resp, err := p.Password("password", "What new password would you like to use?", OptPasswordFlag("p")) // user responds with 1234asdf
//  if err != nil {
//      panic(err)
//  }
//
//  fmt.Println(resp)
//
// Output:
// 1234asdf
func (*Prompt) Password(name, msg string, options ...PasswordOption) (string, error) {
	definition := daemon.PasswordPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       name,
			PromptType: "password",
			Message:    msg,
		},
	}
	for _, option := range options {
		option(&definition)
	}

	body, err := daemon.AsyncRequest("prompt", definition)
	if err != nil {
		return "", err
	}

	if value, ok := body[name]; ok {
		if str, ok := value.(string); ok {
			return str, nil
		}
		return "", fmt.Errorf("Daemon returned non-string value %v", value)
	}
	return "", fmt.Errorf("Daemon returned incorrect JSON %v", body)
}

type ConfirmOption func(*daemon.ConfirmPromptBody)

// OptConfirmFlag sets the flag value for the confirm prompt.
//
// The flag value is used to match command line arguments to prompts.
func OptConfirmFlag(flag string) ConfirmOption {
	return func(definition *daemon.ConfirmPromptBody) {
		definition.Flag = flag
	}
}

func OptConfirmDefault(defaultValue bool) ConfirmOption {
	return func(definition *daemon.ConfirmPromptBody) {
		definition.Default = defaultValue
	}
}

// Confirm presents a yes/no question to the user in the interface
// (i.e. terminal or slack).
//
// Example:
//
//  p := ctoai.NewPrompt()
//  resp, err := p.PromptConfirm("verbose", "Do you want to run in verbose mode?", OptConfirmFlag("c"), OptConfirmDefault(false)) // user responds with y
//  if err != nil {
//      panic(err)
//  }
//
//  fmt.Println(resp)
//
// Output:
// true
func (*Prompt) Confirm(name, msg string, options ...ConfirmOption) (bool, error) {
	definition := daemon.ConfirmPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       name,
			PromptType: "confirm",
			Message:    msg,
		},
	}
	for _, option := range options {
		option(&definition)
	}

	body, err := daemon.AsyncRequest("prompt", definition)
	if err != nil {
		return false, err
	}

	if value, ok := body[name]; ok {
		if b, ok := value.(bool); ok {
			return b, nil
		}
		return false, fmt.Errorf("Daemon returned non-boolean value %v", value)
	}
	return false, fmt.Errorf("Daemon returned incorrect JSON %v", body)
}

type ListOption func(*daemon.ListPromptBody)

// OptListFlag sets the flag value for the list prompt.
//
// The flag value is used to match command line arguments to prompts.
func OptListFlag(flag string) ListOption {
	return func(definition *daemon.ListPromptBody) {
		definition.Flag = flag
	}
}

func OptListDefaultValue(defaultValue string) ListOption {
	return func(definition *daemon.ListPromptBody) {
		definition.DefaultValue = defaultValue
		definition.DefaultSet = true
		definition.DefaultIsValue = true
	}
}

func OptListDefaultIndex(defaultIndex int) ListOption {
	return func(definition *daemon.ListPromptBody) {
		definition.DefaultIndex = defaultIndex
		definition.DefaultSet = true
		definition.DefaultIsValue = false
	}
}

func OptListAutocomplete(autocomplete bool) ListOption {
	return func(definition *daemon.ListPromptBody) {
		if autocomplete {
			definition.PromptType = "autocomplete"
		} else {
			definition.PromptType = "list"
		}
	}
}

// List presents a list of options to the user to select one item from in
// the interface (i.e. terminal or slack).
//
// choices is the list of string options that can be selected
//
// Example:
//
//  p := ctoai.NewPrompt()
//  resp, err := p.List("platform", "What cloud platform would you like to deploy to?", []string{"AWS", "Google Cloud", "Azure"}, OptListDefault("AWS"), OptListFlag("L"), OptListAutocomplete(true)) // user selects Azure
//  if err != nil {
//      panic(err)
//  }
//
//  fmt.Println(resp)
//
// Output:
// Azure
func (*Prompt) List(name, msg string, choices []string, options ...ListOption) (string, error) {
	definition := daemon.ListPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       name,
			PromptType: "list",
			Message:    msg,
		},
		Choices: choices,
	}
	for _, option := range options {
		option(&definition)
	}

	body, err := daemon.AsyncRequest("prompt", definition)
	if err != nil {
		return "", err
	}

	if value, ok := body[name]; ok {
		if b, ok := value.(string); ok {
			return b, nil
		}
		return "", fmt.Errorf("Daemon returned non-string value %v", value)
	}
	return "", fmt.Errorf("Daemon returned incorrect JSON %v", body)
}

type CheckboxOption func(*daemon.CheckboxPromptBody)

// OptCheckboxFlag sets the flag value for the checkbox prompt.
//
// The flag value is used to match command line arguments to prompts.
func OptCheckboxFlag(flag string) CheckboxOption {
	return func(definition *daemon.CheckboxPromptBody) {
		definition.Flag = flag
	}
}

func OptCheckboxDefaultValues(defaultValues []string) CheckboxOption {
	return func(definition *daemon.CheckboxPromptBody) {
		definition.DefaultValue = defaultValues
		definition.DefaultSet = true
		definition.DefaultIsValue = true
	}
}

func OptCheckboxDefaultIndex(defaultIndexes []int) CheckboxOption {
	return func(definition *daemon.CheckboxPromptBody) {
		definition.DefaultIndex = defaultIndexes
		definition.DefaultSet = true
		definition.DefaultIsValue = false
	}
}

// Checkbox presents a list of options to the user, who can select multiple
// items in the interface (i.e. terminal or slack).
//
// choices is the list of string options that can be selected
//
// Example:
//
//  p := ctoai.NewPrompt()
//  resp, err := s.PromptCheckbox("tools", "Which interpreters would you like to have included in your OS image?", []string{"Lua", "Node.js", "Perl", "Python 2", "Python 3", "Raku", "Ruby"}, OptCheckboxDefaultValues([]string{"Lua"}), OptCheckboxFlag("C")) // user selects Lua
//  if err != nil {
//      panic(err)
//  }
//
//  fmt.Println(resp)
//
// Output:
// Lua
func (*Prompt) Checkbox(name, msg string, choices []string, options ...CheckboxOption) ([]string, error) {
	definition := daemon.CheckboxPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       name,
			PromptType: "checkbox",
			Message:    msg,
		},
		Choices: choices,
	}
	for _, option := range options {
		option(&definition)
	}

	body, err := daemon.AsyncRequest("prompt", definition)
	if err != nil {
		return nil, err
	}

	if value, ok := body[name]; ok {
		if values, ok := value.([]interface{}); ok {
			strings := make([]string, len(values))
			for i, v := range values {
				if s, ok := v.(string); ok {
					strings[i] = s
				} else {
					return nil, fmt.Errorf("Daemon returned non-string value %v", v)
				}
			}
			return strings, nil
		}
		return nil, fmt.Errorf("Daemon returned non-array value %v", value)
	}
	return nil, fmt.Errorf("Daemon returned incorrect JSON %v", body)
}

// EditorOption is an option for the Editor prompt function
type EditorOption func(*daemon.EditorPromptBody)

// OptEditorFlag sets the flag value for the editor prompt.
//
// The flag value is used to match command line arguments to prompts.
func OptEditorFlag(flag string) EditorOption {
	return func(definition *daemon.EditorPromptBody) {
		definition.Flag = flag
	}
}

func OptEditorDefault(defaultValue string) EditorOption {
	return func(definition *daemon.EditorPromptBody) {
		definition.Default = defaultValue
	}
}

// Editor presets a prompt requesting a multi-line response from the
// user. If used in a terminal interface, the nano editor will be
// presented.
//
// Example:
//
//  p := ctoai.NewPrompt()
//
//  template := `Features:
//
//  Fixes:
//
//  Chores:
//`
//
//  resp, err := p.Editor("notes", "Please enter your release notes", OptEditorDefault(template), OptEditorFlag("e"))
//  if err != nil {
//      panic(err)
//  }
//
// Output:
// [Nano will be brought up with the template in the editor]
func (*Prompt) Editor(name, msg string, options ...EditorOption) (string, error) {
	definition := daemon.EditorPromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       name,
			PromptType: "editor",
			Message:    msg,
		},
	}
	for _, option := range options {
		option(&definition)
	}

	body, err := daemon.AsyncRequest("prompt", definition)
	if err != nil {
		return "", err
	}

	if value, ok := body[name]; ok {
		if str, ok := value.(string); ok {
			return str, nil
		}
		return "", fmt.Errorf("Daemon returned non-string value %v", value)
	}
	return "", fmt.Errorf("Daemon returned incorrect JSON %v", body)
}

// DatetimeOption is an option for the Datetime prompt function
type DatetimeOption func(*daemon.DatetimePromptBody)

// OptDatetimeFlag sets the flag value for the datetime prompt.
//
// The flag value is used to match command line arguments to prompts.
func OptDatetimeFlag(flag string) DatetimeOption {
	return func(definition *daemon.DatetimePromptBody) {
		definition.Flag = flag
	}
}

const (
	DATETIME = "datetime"
	DATE     = "date"
	TIME     = "time"
)

func OptDatetimeVariant(variant string) DatetimeOption {
	return func(definition *daemon.DatetimePromptBody) {
		definition.Variant = variant
	}
}

func OptDatetimeDefault(defaultValue time.Time) DatetimeOption {
	return func(definition *daemon.DatetimePromptBody) {
		definition.Default = defaultValue
	}
}

func OptDatetimeMaximum(maximumValue time.Time) DatetimeOption {
	return func(definition *daemon.DatetimePromptBody) {
		definition.Maximum = maximumValue
	}
}

func OptDatetimeMinimum(minimumValue time.Time) DatetimeOption {
	return func(definition *daemon.DatetimePromptBody) {
		definition.Minimum = minimumValue
	}
}

// Datetime presents a date picker to the user that allows them to
// select a date and/or time.
//
// The method returns the user's response as a time.Time type.
//
// Example:
//  import "time"
//
//  p := ctoai.NewPrompt()
//  resp, err := p.Datetime("nextRun", "When do you want to run the code next?", OptDatetimeVariant(DATETIME), OptDatetimeFlag("T"), OptDatetimeDefault(time.Now()), OptDatetimeMaximum(time.Now().Add(time.Hour * 2)), OptDatetimeMinimum(time.Now())) // user selects default
//  if err != nil {
//      panic(err)
//  }
//
//  fmt.Println(resp)
//
// Output:
// [the output will equal time.Now() in 2006-01-02 15:04:05 format]
func (*Prompt) Datetime(name, msg string, options ...DatetimeOption) (time.Time, error) {
	definition := daemon.DatetimePromptBody{
		PromptEnvelope: daemon.PromptEnvelope{
			Name:       name,
			PromptType: "datetime",
			Message:    msg,
		},
		Variant: DATETIME,
	}
	for _, option := range options {
		option(&definition)
	}

	body, err := daemon.AsyncRequest("prompt", definition)
	if err != nil {
		return time.Unix(0, 0), err
	}

	if value, ok := body[name]; ok {
		if str, ok := value.(string); ok {
			t, err := time.Parse(time.RFC3339, str)
			if err != nil {
				return time.Unix(0, 0), fmt.Errorf("Daemon returned invalid timestamp %v", str)
			}
			return t, nil
		}
		return time.Unix(0, 0), fmt.Errorf("Daemon returned non-string value %v", value)
	}
	return time.Unix(0, 0), fmt.Errorf("Daemon returned incorrect JSON %v", body)
}
