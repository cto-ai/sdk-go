package daemon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const badPortMsg = "The CTO.ai Ops SDK requires a daemon process to be running; this does not appear to be the case."

func port() int {
	portStr := os.Getenv("SDK_SPEAK_PORT")
	if portStr == "" {
		panic(badPortMsg)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(badPortMsg)
	}

	return port
}

func daemonRequest(endpoint string, body interface{}) (*http.Response, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON body: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://127.0.0.1:%d/%s", port(), endpoint),
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return nil, fmt.Errorf("Error in daemon request: %w", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		// TODO: can this be a more specific type?
		var responseBody interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		if err != nil {
			return nil, fmt.Errorf("Status code %d, with JSON decode error %w on body", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("Status code %d with JSON error %v", resp.StatusCode, responseBody)
	}

	return resp, nil
}

func SimpleRequest(endpoint string, body interface{}) error {
	_, err := daemonRequest(endpoint, body)
	return err
}

func SyncRequest(endpoint string, body interface{}) (interface{}, error) {
	resp, err := daemonRequest(endpoint, body)
	if err != nil {
		return nil, err
	}

	var responseBody struct {
		Value interface{} `json:"value"`
	}

	err = json.NewDecoder(resp.body).Decode(&responseBody)
	if err != nil {
		return nil, fmt.Errorf("Error decoding daemon response %w", err)
	}

	return responseBody.Value, nil
}

func AsyncRequest(endpoint string, body interface{}) (map[string]interface{}, error) {
	resp, err := daemonRequest(endpoint, body)
	if err != nil {
		return nil, err
	}

	var responseBody struct {
		Filename string `json:"replyFilename"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return nil, fmt.Errorf("Error decoding daemon response %w", err)
	}

	bytes, err := ioutil.ReadFile(responseBody.Filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading daemon response %w", err)
	}

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(bytes, &responseMap)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling daemon response %w", err)
	}

	return responseMap, nil
}
