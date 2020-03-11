package ctoai

import (
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
func (*Sdk) GetStatePath() string {
	path := os.Getenv("SDK_STATE_DIR")
	if path == "" {
		panic("State directory not found in environment var SDK_STATE_DIR")
	}
	return path
}

// GetConfigPath returns the path to the config directory (local to this particular op)
func (*Sdk) GetConfigPath() string {
	path := os.Getenv("SDK_CONFIG_DIR")
	if path == "" {
		panic("Config directory not found in environment var SDK_CONFIG_DIR")
	}
	return path
}

// GetState returns a value from the state (workflow-local) key/value store
func (s *Sdk) GetState(key string) (interface{}, error) {
	return daemon.SyncRequest("state/get", map[string]interface{}{"key": key})
}

// SetState sets a value in the state (workflow-local) key/value store
func (s *Sdk) SetState(key string, value interface{}) error {
	return daemon.SimpleRequest("state/set", map[string]interface{}{
		"key":   key,
		"value": value,
	})
}

// GetConfig returns a value from the config (user-specific) key/value store
func (s *Sdk) GetConfig(key string) (interface{}, error) {
	return daemon.SyncRequest("config/get", map[string]interface{}{"key": key})
}

// SetConfig sets a value in the config (user-specific) key/value store
func (s *Sdk) SetConfig(key string, value interface{}) error {
	return daemon.SimpleRequest("config/set", map[string]interface{}{
		"key":   key,
		"value": value,
	})
}

// GetSecret requests a secret from the secret store by key.
//
// If the secret exists, it is returned, with the daemon notifying the user that it is in use.
// Otherwise, the user is prompted to provide a replacement.
func (*Sdk) GetSecret(key string) (string, error) {
	body, err := daemon.AsyncRequest("secret/get", daemon.GetSecretBody{Key: key})
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
