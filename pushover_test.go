package pushover

import (
	"testing"

	"github.com/TV4/env"
	pushover "github.com/bdenning/go-pushover"
)

// TestPush runs through a number of test cases (testCases) and ensures that API responses are as expected.
func TestPush(t *testing.T) {

	// Run tests that are intended to test network connectivity and other non-API failures against the real API

	config := Config{
		Token:   env.String("PUSHOVER_TOKEN", ""),
		UserKey: env.String("PUSHOVER_USER", ""),
	}

	if !config.AllSet() {
		t.Error("Credentials are not set to actively push messages")
	}
	// Create a fresh new message object for each test case
	m := NewMessageConfig(config)

	// Send a message and check for errors
	m.SetTitle("Testing message")
	m.SetSound(SoundAlien)
	m.SetPriority(PriorityEmergency)
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

func ExampleMessage_Push() {
	// You'll need to configure these by logging in to https://pushover.net.
	token := "KzGDORePKggMaC0QOYAMyEEuZJnyUi"
	user := "e9e1495ec75826de5983cd1abc8031"

	// Send a new message using the Push method.
	m := pushover.NewMessage(token, user)
	m.Push("Test message contents")
}
