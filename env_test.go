package ctoai

import (
	"os"
	"testing"
)

func Test_GetHostOS(t *testing.T) {
	expected := "test"

	err := os.Setenv("OPS_HOST_PLATFORM", expected)
	if err != nil {
		t.Errorf("Error setting test env variable: %s", err)
	}

	s := NewSdk()
	output := s.GetHostOS()
	if output != "test" {
		t.Errorf("Error unexpected output: %v", output)
	}

	err = os.Setenv("OPS_HOST_PLATFORM", "")
	if err != nil {
		t.Errorf("Error clearing test env variable: %s", err)
	}

	output = s.GetHostOS()
	if output != "unknown" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_GetInterfaceType(t *testing.T) {
	expected := "test"

	err := os.Setenv("SDK_INTERFACE_TYPE", expected)
	if err != nil {
		t.Errorf("Error setting test env variable: %s", err)
	}

	s := NewSdk()
	output := s.GetInterfaceType()
	if output != "test" {
		t.Errorf("Error unexpected output: %v", output)
	}

	err = os.Setenv("SDK_INTERFACE_TYPE", "")
	if err != nil {
		t.Errorf("Error clearing test env variable: %s", err)
	}

	output = s.GetInterfaceType()
	if output != "terminal" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_StateDir(t *testing.T) {
	expected := "test"

	err := os.Setenv("SDK_STATE_DIR", expected)
	if err != nil {
		t.Errorf("Error setting test env variable: %s", err)
	}

	s := NewSdk()
	output := s.GetStatePath()

	if output != "test" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_HomeDir(t *testing.T) {
	expected := "test"

	err := os.Setenv("SDK_HOME_DIR", expected)
	if err != nil {
		t.Errorf("Error setting test env variable: %s", err)
	}

	s := NewSdk()
	output := s.HomeDir()
	if output != "test" {
		t.Errorf("Error unexpected output: %v", output)
	}

	err = os.Setenv("SDK_HOME_DIR", "")
	if err != nil {
		t.Errorf("Error clearing test env variable: %s", err)
	}

	output = s.HomeDir()
	if output != "/root" {
		t.Errorf("Error unexpected output: %v", output)
	}
}

func Test_ConfigDir(t *testing.T) {
	expected := "test"

	err := os.Setenv("SDK_CONFIG_DIR", expected)
	if err != nil {
		t.Errorf("Error setting test env variable: %s", err)
	}

	s := NewSdk()
	output := s.GetConfigPath()

	if output != "test" {
		t.Errorf("Error unexpected output: %v", output)
	}
}
