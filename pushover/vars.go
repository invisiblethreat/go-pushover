package pushover

const (
	// PushoverURL is the API endpoint that will be used for sending all messages.
	PushoverURL = "https://api.pushover.net/1/messages.json"
	// StatusSuccess is the expected status code when a message has been succesfully sent.
	StatusSuccess = 1

	//Sounds for messages

	// SoundPushover is a Pushover API sound
	SoundPushover Sound = "pushover"
	// SoundClassical is a Pushover API sound
	SoundClassical Sound = "classical"

	// SoundCosmic is a Pushover API sound
	SoundCosmic Sound = "cosmic"

	// SoundFalling is a Pushover API sound
	SoundFalling Sound = "falling"

	// SoundGamelan is a Pushover API sound
	SoundGamelan Sound = "gamelan"

	// SoundIncoming is a Pushover API sound
	SoundIncoming Sound = "incoming"

	// SoundIntermission is a Pushover API sound
	SoundIntermission Sound = "intermission"

	// SoundMagic is a Pushover API sound
	SoundMagic Sound = "magic"

	// SoundMechanical is a Pushover API sound
	SoundMechanical Sound = "mechanical"

	// SoundPianobar is a Pushover API sound
	SoundPianobar Sound = "pianobar"

	// SoundSiren is a Pushover API sound
	SoundSiren Sound = "siren"

	// SoundSpaceAlarm is a Pushover API sound
	SoundSpaceAlarm Sound = "spacealarm"

	// SoundTugBoat is a Pushover API sound
	SoundTugBoat Sound = "tugboat"

	// SoundAlien is a Pushover API sound
	SoundAlien Sound = "alien"

	// SoundClimb is a Pushover API sound
	SoundClimb Sound = "climb"

	// SoundPersistent is a Pushover API sound
	SoundPersistent Sound = "persistent"

	// SoundEcho is a Pushover API sound
	SoundEcho Sound = "echo"

	// SoundUpDown is a Pushover API sound
	SoundUpDown Sound = "updown"
)

const (
	// PriorityLowest makes no sounds, alerts, or popups
	PriorityLowest Priority = -2

	// PriorityLow makes no sounds or alerts, but will make a popup
	PriorityLow Priority = -1

	// PriorityNormal will make a sound, alert, and popup. If this message is
	// received during quiet hours, it will be treated as PriorityLow
	PriorityNormal Priority = 0

	// PriorityHigh will bypass quiet hours, and will be highlighted in red
	PriorityHigh Priority = 1

	// PriorityEmergency are the same as PriorityHigh but require acknowledged.
	// Additionally, 'retry' and 'expiry' must be set in the message
	PriorityEmergency Priority = 2
)
