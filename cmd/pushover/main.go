package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/bdenning/pushover"
)

func getEnvSettings() (token string, user string, device string) {
	token = os.Getenv("PUSHOVER_TOKEN")
	user = os.Getenv("PUSHOVER_USER")
	device = os.Getenv("PUSHOVER_DEVICE")

	if token == "" {
		log.Fatal("Missing PUSHOVER_TOKEN environment variable")
	}

	if user == "" {
		log.Fatal("Missing PUSHOVER_USER environment variable")
	}

	if device == "" {
		log.Fatal("Missing PUSHOVER_DEVICE environment variable")
	}

	return
}

func main() {
	// Get the required configuration information
	token, user, device := getEnvSettings()
	title, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new pushover message object with our settings
	m := pushover.NewMessage(token, user, device)

	// Read the message from stdin and send via pushover.net
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Error getting input from stdin")
	}

	// Send the message
	_, err = m.Push(title, string(stdin))
	if err != nil {
		log.Fatal("Error while sending message")
	}
}
