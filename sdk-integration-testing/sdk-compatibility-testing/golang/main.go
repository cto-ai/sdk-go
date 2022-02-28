package main

import (
	"fmt"
	"time"

	ctoai "github.com/cto-ai/sdk-go"
)

func main() {

	client := ctoai.NewClient()

	err := client.Ux.Print("Begin")
	if err != nil {
		panic(err)
	}

	err = client.Sdk.Track([]string{"informational"}, "sdk-data", map[string]interface{}{
		"osVisible": client.Sdk.GetHostOS(),
		"homeDir":   client.Sdk.HomeDir(),
		"statePath": client.Sdk.GetStatePath(),
		// TODO: Fix the SDK so that calling in doesn't crash
		"configPath": "/ops",
		// "configPath": client.Sdk.GetConfigPath(),
	})
	if err != nil {
		panic(err)
	}

	input, err := client.Prompt.Input("testInput", "Input prompt", ctoai.OptInputDefault("open"), ctoai.OptInputAllowEmpty(true))
	if err != nil {
		panic(err)
	}

	err = client.Ux.Print(fmt.Sprintf("Test Input: %s", input))
	if err != nil {
		panic(err)
	}

	confirm, err := client.Prompt.Confirm("testConfirm", "Confirm prompt")
	if err != nil {
		panic(err)
	}

	if confirm {
		err = client.Ux.Print("Confirm is true")
	}

	listResult, err := client.Prompt.List("testList", "List prompt", []string{"well", "hello", "there"})
	if err != nil {
		panic(err)
	}

	err = client.Ux.Print(listResult)
	if err != nil {
		panic(err)
	}

	autocompleteResult, err := client.Prompt.List("testAutocomplete", "Autocomplete prompt", []string{"1", "2", "3"}, ctoai.OptListAutocomplete(true), ctoai.OptListDefaultIndex(1))
	if err != nil {
		panic(err)
	}

	err = client.Ux.Print(autocompleteResult)
	if err != nil {
		panic(err)
	}

	password, err := client.Prompt.Password("testPassword", "Password prompt")
	if err != nil {
		panic(err)
	}

	if password == "passwordTest" {
		err = client.Ux.Print("Password matches")
		if err != nil {
			panic(err)
		}

	}

	minimum, err := time.Parse(time.RFC3339, "2020-02-20T22:38:07Z")
	if err != nil {
		panic(err)
	}

	dateTime, err := client.Prompt.Datetime("testDateTime", "DateTime prompt", ctoai.OptDatetimeMinimum(minimum))
	if err != nil {
		panic(err)
	}

	err = client.Ux.Print(dateTime.Format(time.RFC3339))
	if err != nil {
		panic(err)
	}

	minimum, err = time.Parse(time.RFC3339, "2019-01-01T00:47:28Z")
	if err != nil {
		panic(err)
	}

	timeResult, err := client.Prompt.Datetime("testTime", "Time only prompt", ctoai.OptDatetimeMinimum(minimum), ctoai.OptDatetimeVariant(ctoai.TIME))
	if err != nil {
		panic(err)
	}

	err = client.Ux.Print(timeResult.Format(time.RFC3339))
	if err != nil {
		panic(err)
	}

	minimum, err = time.Parse(time.RFC3339, "2020-02-21T00:00:00Z")
	if err != nil {
		panic(err)
	}

	dateResult, err := client.Prompt.Datetime("testDate", "Date only prompt", ctoai.OptDatetimeMinimum(minimum), ctoai.OptDatetimeVariant(ctoai.DATE))
	if err != nil {
		panic(err)
	}

	err = client.Ux.Print(dateResult.Format(time.RFC3339))
	if err != nil {
		panic(err)
	}

	err = client.Ux.SpinnerStart("Starting spinner")
	if err != nil {
		panic(err)
	}

	err = client.Ux.SpinnerStop("Stopping spinner")
	if err != nil {
		panic(err)
	}

	err = client.Ux.ProgressBarStart(5, 0, "Starting Progress Bar")
	if err != nil {
		panic(err)
	}
	err = client.Ux.ProgressBarAdvance(2)
	if err != nil {
		panic(err)
	}
	err = client.Ux.ProgressBarStop("Stopping Progress Bar")
	if err != nil {
		panic(err)
	}

	err = client.Sdk.SetConfig("CONFIG_KEY", "Some random value")
	if err != nil {
		panic(err)
	}

	// We need to simulate the behaviour of the Python and JS SDKs here
	_, err = client.Sdk.GetAllConfig()
	if err != nil {
		panic(err)
	}

	err = client.Ux.Print("CONFIG_KEY set: Some random value")
	if err != nil {
		panic(err)
	}

	configValue, err := client.Sdk.GetConfig("CONFIG_KEY")
	if err != nil {
		panic(err)
	}

	err = client.Ux.Print(fmt.Sprintf("CONFIG_KEY value retrieved: %s", configValue))
	if err != nil {
		panic(err)
	}

	deleted, err := client.Sdk.DeleteConfig("CONFIG_KEY")
	if err != nil {
		panic(err)
	}

	if deleted {
		err = client.Ux.Print("CONFIG_KEY value deleted")
		if err != nil {
			panic(err)
		}
	}

	newValue, err := client.Sdk.GetConfig("CONFIG_KEY")
	if err != nil {
		panic(err)
	}

	if newValue == "" {
		err = client.Ux.Print("CONFIG_KEY not found anymore")
		if err != nil {
			panic(err)
		}

	}

}
