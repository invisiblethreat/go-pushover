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
		Device    string `json:"device"`
		Title     string `json:"title"`
		URL       string `json:"url"`
		URLTitle  string `json:"url_title"`
		Priority  int    `json:"priority"`
		Timestamp int32  `json:"timestamp"`
		Sound     string `json:"sound"`
	}

	// Response contains the JSON response returned by the pushover.net API
	Response struct {
		Request string   `json:"request"`
		Status  int      `json:"status"`
		Errors  []string `json:"errors"`
	}
)
