package mbuf

// Settings holds all of the available optional settings for a MessageRelay.
type Settings struct {

	// ClientTimeoutSecs is how many seconds to wait before timng out a request to the log processor
	ClientTimeoutSecs *int `yaml:"ClientTimeoutSecs" config:"optional"`

	// Flushers sets the Flushers for this buffer
	Flushers []Flusher `yaml:"-" config:"optional"`

	// Marshaler sets the HTTPMarshaler for this buffer
	Marshaler HTTPMarshaler `yaml:"-" config:"optional"`

	// MaxLength is the number of messages we cache before sending
	MaxLength *int `yaml:"MaxLength" config:"optional"`

	// MaxRetries is how many times we retry sending to the log processor after a network faliure
	MaxRetries *int `yaml:"MaxRetries" config:"optional"`

	// PollIntervalSecs is how many seconds to wait before automatically sending cached messages
	PollIntervalSecs *int `yaml:"PollIntervalSecs" config:"optional"`

	// Renderer sets the Renderer for this buffer
	Renderer MessageRenderer `yaml:"-" config:"optional"`

	// RelayURL is the web address of the log processor. This setting is required and connot be empty.
	RelayURL string `yaml:"RelayURL" config:"optional"`

	// RetryWaitSeconds is how many seconds we wait before attempting to retry a failed message
	RetryWaitSeconds *int `yaml:"RetryWaitSeconds" config:"optional"`
}
