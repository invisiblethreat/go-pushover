// Package pushover provides methods for sending messages using the http://pushover.net API.
package pushover

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/TV4/env"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

// NewMessage returns a new Message with API token values and a user configured
// .
func NewMessage(token, user string) *Message {
	return &Message{Token: token, User: user}
}

// NewMessageConfig returns a new Message with API token values and a user
// configured. The argument of Config type allow for easier constraint checking
func NewMessageConfig(config Config) *Message {
	return &Message{Token: config.Token, User: config.UserKey}
}

// Push sends a message via the pushover.net API and returns the json response
func (m *Message) Push(message string) (r *Response, err error) {
	r = &Response{}

	if message == "" {
		return r, errors.New("Message can not be blank")
	}
	m.Message = message

	// Check that required items are set
	if m.Priority == PriorityEmergency {
		err = m.EmergencyParamsSet()
		if err != nil {
			return r, err
		}
	}

	msg, err := json.Marshal(m)
	if err != nil {
		return r, err
	}

	buf := bytes.NewReader(msg)
	// Send the message the the pushover.net API
	resp, err := http.Post(PushoverURL, "application/json", buf)
	//resp, err := http.PostForm(m.URL, msg)
	if err != nil {

		return r, err
	}
	defer resp.Body.Close()

	// Decode the json response from pushover.net in to our Response struct
	if err = json.NewDecoder(resp.Body).Decode(r); err != nil {
		return r, err
	}

	// Check to see if pushover.net set the status to indicate an error without providing and explanation
	if r.Status != StatusSuccess {
		if len(r.Errors) > 0 {
			joined := strings.Join(r.Errors, ", ")
			return r, errors.New(joined)
		}
	}
	return r, nil
}

// GetConfigFile reads YAML file and returns a Config struct.
func GetConfigFile(file string) (Config, error) {

	var config Config
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}

	if !config.AllSet() {
		return config, errors.New("YAML items in file were set to access API")
	}

	return config, nil
}

// GetConfigEnv reads environment variables returns a Config struct.
func GetConfigEnv() (Config, error) {
	config := Config{
		Token:   env.String("PUSHOVER_TOKEN", ""),
		UserKey: env.String("PUSHOVER_USER", ""),
	}

	if !config.AllSet() {
		return config, errors.New("Environmental items were set to access API")
	}

	return config, nil
}

func main() {
	var file, title, message string
	pflag.StringVarP(&file, "file", "f", "pushover.yaml", "YAML config file location")
	pflag.StringVarP(&title, "title", "t", "", "Title for message")
	pflag.StringVarP(&message, "message", "m", "", "Message to send")
	pflag.Parse()
	var config Config
	// try to get env vars first
	config, err := GetConfigEnv()

	if err != nil {
		envErr := err
		// fall back to a config file
		config, err = GetConfigFile(file)
		if err != nil {
			logrus.Fatalf(
				"Errors using env vars and config file: \n\t%s\n\t%s\n",
				envErr.Error(), err.Error())
		}

	}

	m := NewMessageConfig(config)
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
