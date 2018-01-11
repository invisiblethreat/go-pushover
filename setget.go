package pushover

// These are all courtsey functions that make interacting wiht a message very
// easy, and fairly uniform.

// SetTitle sets the title of a message. If unset, the default Pushover name
// will be used
func (m *Message) SetTitle(title string) {
	m.Title = title
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
func (m *Message) SetPriority() {}

// GetPriority returns the priority of the message
func (m *Message) GetPriority() int {
	return m.Priority
}

// SetSound sets the sound for the alert. See vars.go for the list.
func (m *Message) SetSound(s string) {
	m.Sound = s
}

// GetSound returns the set value of the sound set for the message
func (m *Message) GetSound() string {
	return m.Sound
}
