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

func TestNewMesage(t *testing.T) {
	msg := NewMessage("x", "y")

	if msg.Token != "x" {
		t.Errorf("Expected %s, but got %s", "x", msg.Token)
	}

	if msg.User != "y" {
		t.Errorf("Expected %s, but got %s", "y", msg.User)
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

func TestAllSet(t *testing.T) {

	config := Config{Token: "x"}
	if config.AllSet() != false {
		t.Errorf("AllSet: Expected %t, but got %t", false, true)
	}

	config.UserKey = "y"
	if config.AllSet() != true {
		t.Errorf("AllSet: Expected %t, but got %t", true, false)
	}
}
func TestGettersSettters(t *testing.T) {
	m := Message{}

	expected := "test string"

	m.SetMessage(expected)
	if m.GetMessage() != expected {
		t.Errorf("Expected %s, but got %s", expected, m.GetMessage())
	}

	m.SetTitle(expected)
	if m.GetTitle() != expected {
		t.Errorf("Expected %s, but got %s", expected, m.GetTitle())
	}

	var expectedTS int32 = 1

	m.SetUnixTimestamp(expectedTS)
	if m.GetUnixTimestamp() != expectedTS {
		t.Errorf("Expected %d, but got %d", expectedTS, m.GetUnixTimestamp())
	}

	m.SetDevice(expected)
	if m.GetDevice() != expected {
		t.Errorf("Expected %s, but got %s", expected, m.GetDevice())
	}

	m.SetURL(expected)
	if m.GetURL() != expected {
		t.Errorf("Expected %s, but got %s", expected, m.GetURL())
	}

	m.SetURLTitle(expected)
	if m.GetURLTitle() != expected {
		t.Errorf("Expected %s, but got %s", expected, m.GetURLTitle())
	}

	expectedSound := SoundPushover

	m.SetSound(expectedSound)
	if m.GetSound() != expectedSound {
		t.Errorf("Expected %s, but got %s", expectedSound, m.GetSound())
	}

	expectedPriority := PriorityNormal

	m.SetPriority(expectedPriority)
	if m.GetPriority() != expectedPriority {
		t.Errorf("Expected %d, but got %d", expectedPriority, m.GetPriority())
	}

	expectedExpiry := 10800
	failExpiry := 10801

	err := m.SetExpiry(expectedExpiry)
	if err != nil {
		t.Errorf("Error setting expiry time: %s", err.Error())
	}

	if m.GetExpiry() != expectedExpiry {
		t.Errorf("Expected %d, but got %d", expectedExpiry, m.GetExpiry())
	}

	err = m.SetExpiry(failExpiry)
	if err == nil {
		t.Errorf("Expected failure condition not returned for SetExpiry")
	}

	expectedRetry := 30
	failRetry := 29

	err = m.SetRetry(expectedRetry)
	if err != nil {
		t.Errorf("Error setting expiry time: %s", err.Error())
	}

	if m.GetRetry() != expectedRetry {
		t.Errorf("Expected %d, but got %d", expectedRetry, m.GetRetry())
	}

	err = m.SetRetry(failRetry)
	if err == nil {
		t.Errorf("Expected failure condition not returned for SetRetry")
	}

}

func TestSetEmergencyDefault(t *testing.T) {
	config := Config{UserKey: "x", Token: "y"}

	m := NewMessageConfig(config)
	m.SetPriority(PriorityEmergency)
	m.SetEmergencyDefault()

	err := m.EmergencyParamsSet()
	if err != nil {
		t.Errorf("Error in validating emergency params: %s", err.Error())
	}

	if m.GetRetry() != 300 {
		t.Errorf("Expected %d, but got %d", 300, m.GetRetry())
	}
	if m.GetExpiry() != 3600 {
		t.Errorf("Expected %d, but got %d", 3600, m.GetExpiry())
	}

	m.Retry = 0
	err = m.EmergencyParamsSet()
	if err == nil {
		t.Error("Error condition should have been triggered for EmergencyParamSet")
	}

	err = m.EmergencyParamsSet()
	m.Expire = 0
	if err == nil {
		t.Error("Error condition should have been triggered for EmergencyParamSet")
	}

	m.Priority = PriorityHigh
	err = m.EmergencyParamsSet()
	if err != nil {
		t.Error("Error condition should not have been triggered for EmergencyParamSet")
	}
}
