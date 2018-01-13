package pushover

type (
	// Message contains all the required settings for sending messages via the
	//pushover.net API. https://pushover.net/api
	Message struct {

		// Required
		Token   string `json:"token"`
		User    string `json:"user"`
		Message string `json:"message"`
		//Optional
		Device    string   `json:"device"`
		Title     string   `json:"title"`
		URL       string   `json:"url"`
		URLTitle  string   `json:"url_title"`
		Priority  Priority `json:"priority"`
		Timestamp int32    `json:"timestamp"`
		Sound     Sound    `json:"sound"`
		// For when we have PriorityEmergency messages
		Retry  int `json:"retry"`  // Must be greater than 30s
		Expire int `json:"expire"` // Must be less than 10800s (3h)
	}

	// Response contains the JSON response returned by the pushover.net API
	Response struct {
		Request string   `json:"request"`
		Status  int      `json:"status"`
		Errors  []string `json:"errors"`
	}

	// Config helps with managing the ability to enforce constraints more easily
	Config struct {
		UserKey string `yaml:"user_key"`
		Token   string `yaml:"api_token"`
	}

	// Sound is a an acceptable string for the Pushover API
	Sound string

	// Priority is an acceptable priority for the Pushover API
	Priority int
)
