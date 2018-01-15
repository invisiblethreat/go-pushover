package main

import (
	"io/ioutil"
	"os"

	"github.com/invisiblethreat/go-pushover/pushover"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	var file, title, message string
	pflag.StringVarP(&file, "file", "f", "pushover.yaml",
		"YAML config file location")
	pflag.StringVarP(&title,
		"title", "t", "",
		"Title for message. If empty, default name for the token will be used")
	pflag.StringVarP(&message, "message", "m", "",
		"Message to send. Omit if piping from STDIN")
	pflag.Parse()
	var config pushover.Config
	// try to get env vars first
	config, err := pushover.GetConfigEnv()

	if err != nil {
		envErr := err
		// fall back to a config file
		config, err = pushover.GetConfigFile(file)
		if err != nil {
			logrus.Fatalf(
				"Errors using env vars and config file: \n\t%s\n\t%s\n",
				envErr.Error(), err.Error())
		}

	}

	m := pushover.NewMessageConfig(config)
	if title != "" {
		m.SetTitle(title)
	}

	if message == "" {

		// Read the message from stdin and send
		stdin, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			logrus.Fatal("Error getting input from STDIN")
		}
		message = string(stdin)
	}
	// Send the message
	_, err = m.Push(message)
	if err != nil {
		logrus.WithError(err).Fatal("Error while sending message")
	}
}
