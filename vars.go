package pushover

const (
	// PushoverURL is the API endpoint that will be used for sending all messages.
	PushoverURL = "https://api.pushover.net/1/messages.json"
	// StatusSuccess is the expected status code when a message has been succesfully sent.
	StatusSuccess = 1

	// Sounds
	SoundPushover     = "pushover"
	SoundBike         = "bike"
	SoundBugle        = "bugle"
	SoundCashRegister = "cashregister"
	SoundClassical    = "classical"
	SoundCosmic       = "cosmic"
	SoundFalling      = "falling"
	SoundGamelan      = "gamelan"
	SoundIncoming     = "incoming"
	SoundIntermission = "intermission"
	SoundMagic        = "magic"
	SoundMechanical   = "mechanical"
	SoundPianobar     = "pianobar"
	SoundSiren        = "siren"
	SoundSpaceAlarm   = "spacealarm"
	SoundTugBoat      = "tugboat"
	SoundAlien        = "alien"
	SoundClimb        = "climb"
	SoundPersistent   = "persistent"
	SoundEcho         = "echo"
	SoundUpDown       = "updown"
	SoundNone         = "none"

	PriorityLowest    = -2
	PriorityLow       = -1
	PriorityNormal    = 0
	PriorityHigh      = 1
	PriorityEmergency = 2
)
