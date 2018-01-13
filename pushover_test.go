package pushover

import (
	"os"
	"testing"
)

// TestPush runs through a number of test cases (testCases) and ensures that API responses are as expected.
func TestPush(t *testing.T) {

	// Run tests that are intended to test network connectivity and other non-API failures against the real API

	config, err := GetConfigFile("pushover.yaml")

	if err != nil {
		t.Error("Credentials are not set to actively push messages")
	}
	// Create a fresh new message object for each test case
	m := NewMessageConfig(config)

	// Send a message and check for errors
	m.SetTitle("Testing message")
	m.SetSound(SoundAlien)
	m.SetPriority(PriorityNormal)
	m.SetEmergencyDefault()

	resp, err := m.Push("Testing message")
	if err != nil {
		t.Errorf("Error sending message: %s", err.Error())
	}

	// Check for failures that did not result in Push() returning an error
	if resp.Status != StatusSuccess {
		t.Errorf("Response was not successful: %s", err.Error())
	}

	// Check that the the status returned by the API is what we were expecting.
	if resp.Status != 1 {
		t.Errorf("The test returned an unexpected status code: %d", resp.Status)
	}

}

func TestGetConfigFile(t *testing.T) {
	config, err := GetConfigFile("data/test_creds.yaml")

	if err != nil {
		t.Errorf("Error reading config file: %s", err.Error())
	}

	expected := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	if config.UserKey != expected {
		t.Errorf("Expected %s, but got %s", expected, config.UserKey)
	}
	expected = "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"
	if config.Token != expected {
		t.Errorf("Expected %s, but got %s", expected, config.UserKey)
	}
}

func TestGetConfigEnv(t *testing.T) {

	err := os.Setenv("PUSHOVER_USER", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		t.Errorf("Error seting env var: %s", err.Error())
	}
	err = os.Setenv("PUSHOVER_TOKEN", "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyy")
	if err != nil {
		t.Errorf("Error seting env var: %s", err.Error())
	}

	config, err := GetConfigEnv()
	if err != nil {
		t.Errorf("Error getting config from envars: %s", err.Error())
	}

	expected := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	if config.UserKey != expected {
		t.Errorf("Expected %s, but got %s", expected, config.UserKey)
	}
	expected = "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"
	if config.Token != expected {
		t.Errorf("Expected %s, but got %s", expected, config.UserKey)
	}
}
