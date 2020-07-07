package ctoai

import (
	"errors"
	"fmt"
	"os"

	"github.com/cto-ai/sdk-go/internal/daemon"
)

func getenv(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}

// Sdk is the object that contains the SDK methods
type Sdk struct{}

func NewSdk() *Sdk {
	return &Sdk{}
}

// GetHostOS returns the current host OS.
func (*Sdk) GetHostOS() string {
	return getenv("OPS_HOST_PLATFORM", "unknown")
}

// GetInterfaceType returns the interface type that the op is attached to
func (*Sdk) GetInterfaceType() string {
	return getenv("SDK_INTERFACE_TYPE", "terminal")
}

// HomeDir returns the location of the user home directory
func (*Sdk) HomeDir() string {
	return getenv("SDK_HOME_DIR", "/root")
}

// GetStatePath returns the path to the state directory (local to this particular workflow)
// DEPRECATED: use `HomeDir` instead.
func (*Sdk) GetStatePath() string {
	path := os.Getenv("SDK_STATE_DIR")
	if path == "" {
		panic("State directory not found in environment var SDK_STATE_DIR")
	}
	return path
}

// GetConfigPath returns the path to the config directory (local to this particular op)
// DEPRECATED: incompatible with current config API
func (*Sdk) GetConfigPath() string {
	path := os.Getenv("SDK_CONFIG_DIR")
	if path == "" {
		panic("Config directory not found in environment var SDK_CONFIG_DIR")
	}
	return path
}

// GetState returns a value from the state (workflow-local) key/value store
// DEPRECATED: state is used by deprecated workflows feature
func (s *Sdk) GetState(key string) (interface{}, error) {
	return daemon.SyncRequest("state/get", map[string]interface{}{"key": key})
}

// GetAllState returns a map of all keys to values in the state (workflow-local) key/value store
// DEPRECATED: state is used by deprecated workflows feature
func (s *Sdk) GetAllState() (map[string]interface{}, error) {
	value, err := daemon.SyncRequest("state/get-all", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	mapValue, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Received non-object JSON %v", value)
	}
	return mapValue, nil
}

// SetState sets a value in the state (workflow-local) key/value store
// DEPRECATED: state is used by deprecated workflows feature
func (s *Sdk) SetState(key string, value interface{}) error {
	return daemon.SimpleRequest("state/set", map[string]interface{}{
		"key":   key,
		"value": value,
	})
}

// GetConfig returns a value from the config (user-specific) key/value store
func (s *Sdk) GetConfig(key string) (string, error) {
	daemonValue, err := daemon.SyncRequest("config/get", map[string]string{"key": key})
	if err != nil {
		return "", err
	}

	if daemonValue == nil {
		return "", nil
	}

	stringValue, ok := daemonValue.(string)
	if !ok {
		return "", fmt.Errorf("Received non-string JSON %v", daemonValue)
	}
	return stringValue, nil
}

// GetAllConfig returns a map of all keys to values in the config (workflow-local) key/value store
func (s *Sdk) GetAllConfig() (map[string]string, error) {
	value, err := daemon.SyncRequest("config/get-all", map[string]string{})
	if err != nil {
		return nil, err
	}

	mapValue, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Received non-object JSON %v", value)
	}

	result := make(map[string]string)

	for k, v := range mapValue {
		strV, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("Received non-string JSON %v", v)
		}
		result[k] = strV
	}

	return result, nil
}

// SetConfig sets a value in the config (user-specific) key/value store
func (s *Sdk) SetConfig(key string, value string) error {
	return daemon.SimpleRequest("config/set", map[string]string{
		"key":   key,
		"value": value,
	})
}

// DeleteConfig deletes a value from the config (user-specific) key/value store
// Returns false if key not found, true if success
func (s *Sdk) DeleteConfig(key string) (bool, error) {
	daemonValue, err := daemon.SyncRequest("config/delete", map[string]string{"key": key})
	if err != nil {
		return false, err
	}

	valueDeleted, ok := daemonValue.(bool)
	if !ok {
		return false, fmt.Errorf("Received non-boolean JSON %v", daemonValue)
	}
	return valueDeleted, nil
}

// GetSecretOption is an option for the GetSecret function
type GetSecretOption func(*daemon.GetSecretBody)

// OptGetSecretHidden suppresses the notification that a secret has been retrieved for this request
func OptGetSecretHidden() GetSecretOption {
	return func(body *daemon.GetSecretBody) {
		body.Hidden = true
	}
}

// GetSecret requests a secret from the secret store by key.
//
// If the secret exists, it is returned, with the daemon notifying the user that it is in use.
// Otherwise, the user is prompted to provide a replacement.
func (*Sdk) GetSecret(key string, options ...GetSecretOption) (string, error) {

	requestBody := daemon.GetSecretBody{Key: key}
	for _, option := range options {
		option(&requestBody)
	}

	body, err := daemon.AsyncRequest("secret/get", requestBody)
	if err != nil {
		return "", err
	}

	if value, ok := body[key]; ok {
		return value.(string), nil
	}
	return "", fmt.Errorf("Body should include key %s", key)
}

// SetSecret sets a particular value into the secret store
//
// If the secret already exists, the user is prompted on whether to overwrite it.
func (*Sdk) SetSecret(key string, value string) (string, error) {
	body, err := daemon.AsyncRequest("secret/set", daemon.SetSecretBody{Key: key, Value: value})
	if err != nil {
		return "", err
	}

	if value, ok := body["key"]; ok && value != nil {
		return value.(string), nil
	}
	return "", fmt.Errorf("Secret set of %s failed", key)
}

// Track sends an event to the CTO.ai analytics system.
// This is the public facing API to enable developers to send data
// that they want to be tracked by cto.ai
//
// Example:
//
//  s := ctoai.NewSdk()
//  err := s.Track([]string{"sdk", "go", "tracked"}, "testing", map[string]interface{}{
//      "user": "name",
//  })
//
// The event, tags, and payload will be logged.
func (*Sdk) Track(tags []string, event string, metadata map[string]interface{}) error {
	requestBody := map[string]interface{}{
		"tags":  tags,
		"event": event,
	}
	for k, v := range metadata {
		requestBody[k] = v
	}

	// We suppress this error to be consistent with other languages
	_ = daemon.SimpleRequest("track", requestBody)

	return nil
}

func (*Sdk) Events(start, end string) ([]map[string]interface{}, error) {
	result, err := daemon.SyncRequest("events", daemon.EventsBody{
		Start: start,
		End:   end,
	})

	if err != nil {
		return nil, fmt.Errorf("error getting events from backend: %w", err)
	}

	slice, ok := result.([]interface{})
	if !ok {
		return nil, errors.New("backend returned non-array JSON")
	}

	events := make([]map[string]interface{}, len(slice))
	for i, entry := range slice {
		obj, ok := entry.(map[string]interface{})
		if !ok {
			return nil, errors.New("backend returned non-object JSON entries")
		}
		events[i] = obj
	}

	return events, nil
}
