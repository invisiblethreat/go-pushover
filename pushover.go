// Package pushover provides methods for sending messages using the http://pushover.net API.
package pushover

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
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
