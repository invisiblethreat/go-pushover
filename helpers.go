package pushover

import (
	"fmt"
	"strings"
)

// These are all courtsey functions that make interacting wiht a message very
// easy, and fairly uniform.

// SetTitle sets the title of a message. If unset, the default Pushover name
// will be used
func (m *Message) SetTitle(s string) {
	m.Title = s
}

// GetTitle returns the set title for the message.
func (m *Message) GetTitle() string {
	return m.Title
}

// SetUnixTimestamp will set a time to be used rather than the time the message
// was received at the Pushover server
func (m *Message) SetUnixTimestamp(t int32) {
	m.Timestamp = t
}

// GetUnixTimestamp returns the timestamp supplied to the record.
func (m *Message) GetUnixTimestamp() int32 {
	return m.Timestamp
}

// SetDevice to send message. This should also work for groups.
func (m *Message) SetDevice(s string) {
	m.Device = s
}

// GetDevice returns the value of the device
func (m *Message) GetDevice() string {
	return m.Device
}

// SetURL sets the clickable URL
func (m *Message) SetURL(s string) {
	m.URL = s
}

// GetURL returns the value of the URL
func (m *Message) GetURL() string {
	return m.URL
}

// SetURLTitle Alternate text for the clickable URL
func (m *Message) SetURLTitle(s string) {
	m.URLTitle = s
}

// GetURLTitle returns the value of the title for the clickable URL
func (m *Message) GetURLTitle() string {
	return m.URLTitle

}

// SetPriority sets the priority of the message. See vars.go for the levels.
func (m *Message) SetPriority(p Priority) {
	m.Priority = p
}

// GetPriority returns the priority of the message
func (m *Message) GetPriority() Priority {
	return m.Priority
}

// SetSound sets the sound for the alert. See vars.go for the list.
func (m *Message) SetSound(s Sound) {
	m.Sound = s
}

// GetSound returns the set value of the sound set for the message
func (m *Message) GetSound() Sound {
	return m.Sound
}

// SetExpiry sets the sound for the alert. See vars.go for the list.
func (m *Message) SetExpiry(i int) error {
	if i < 10800 {
		m.Expire = i
		return nil
	}
	return fmt.Errorf("Expiry %ds exceeds Pushover maximum of 10800s", i)
}

// GetExpiry returns the set value of the expiry time set for the message
func (m *Message) GetExpiry() int {
	return m.Expire
}

// SetRetry sets the retry interval for the alert.
func (m *Message) SetRetry(i int) error {
	if i >= 30 {
		m.Retry = i
		return nil
	}
	return fmt.Errorf("Retry of %ds is more agressive than 30s Pushover minimum", i)
}

// GetRetry returns the set value of the retry time set for the message
func (m *Message) GetRetry() int {
	return m.Retry
}

// SetEmergencyDefault sets the expiry to 1h and retry to 5m
func (m *Message) SetEmergencyDefault() {
	m.Expire = 3600 // 1h
	m.Retry = 300   //5m
}

// EmergencyParamsSet tests if required items are set for sending an emergency
// message
func (m *Message) EmergencyParamsSet() error {
	if m.Priority != PriorityEmergency {
		return nil
	}
	var items []string
	if m.Retry == 0 {
		items = append(items, "retry")
	}

	if m.Expire == 0 {
		items = append(items, "expire")
	}

	if len(items) > 0 {
		return fmt.Errorf("Required items not set: %s", strings.Join(items, ", "))
	}
	return nil
}

// AllSet is used to see if we need more items to be set for sending a message
func (c *Config) AllSet() bool {
	if c.Token != "" && c.UserKey != "" {
		return true
	}
	return false
}
