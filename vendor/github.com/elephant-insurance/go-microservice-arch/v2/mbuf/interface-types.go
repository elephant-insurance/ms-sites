package mbuf

import "net/http"

// MessageRelay is the primary interface of the messagerelay package.
// It represents a completely self-contained "fire and forget" message sender
type MessageRelay interface {
	// Add adds a single message object to this MessageRelay's buffer.
	Add(interface{})

	// Diagnostics provides counters and other feedback about the health of the buffer.
	// Since most error conditions are not returned to the caller, this is the caller's primary means of checking the health of the MessageRelay and its children.
	Diagnostics() map[string]interface{}

	// SendNow sends the messages on the calling thread with retries and returns basic call status info
	SendNow(msgs []interface{}) (status *int, respBody *string, err error)

	// SendNowNoRetry causes all messages to be sent immediately on the calling thread.
	// Use this to clean up before exiting.
	SendNowNoRetry()
}

// HTTPMarshaler turns a slice of messages into a sendable HTTP request.
// You may provide your own implementation of HTTPMarshaler to handle special requirements.
// Each element of the array argument to Marshal must be castable to []byte for building the body of the request.
// In addition to building the HTTPRequest itself, Marshal is responsible for building the array structure for multi-message sends.
// For simple JSON array of messages, for example, this means separating individual messages by commas, and wrapping them in [].
type HTTPMarshaler interface {

	// Marshaler turns a slice of messages into a sendable HTTP request.
	Marshal([]interface{}) (*http.Request, error)

	// Diagnostics provides counters and other feedback about the health of the marshaler.
	// These values are added to the Diagnostics output for the MessageRelay.
	Diagnostics() map[string]interface{}
}

// MessageRenderer turns a single message object into a byte array appropriate for adding to the body of a message.
// In general, a MessageRenderer knows about a certain type of message (microservice_log, realtimebid_log, etc.),
// but a default MessageRenderer is provided which simply serializes the message using the encoding/json package.
// You may provide your own implementation of MessageRenderer to meet your specific requirements.

// An implementation of MessageRenderer specific to a certain type of log is responsible for ensuring that each message is formatted appropriately for its receiver,
// and also for limiting the fields forwarded to fields or columns that we agree should appear in our logs.
// For services like Azure Log Analytics and Loggly, which accept any arbitrary fields, use MessageRenderer to prevent the creation of new columns or indexed fields.
type MessageRenderer interface {

	// Render turns a single message into a byte array appropriate for adding to a message body.
	Render(interface{}) []byte

	// Diagnostics provides counters and other feedback about the health of the renderer
	Diagnostics() map[string]interface{}
}

// Flusher handles messages that we have given up trying to send.
// The default flusher simply prints messages to stdout as JSON.
// You may add additional Flushers to a MessageRelay to enable additional behaviors.
// Additional flushers could forward messages to other handlers, write them to a file, etc.
type Flusher interface {

	// Flush is the last thing we do with messages before throwing them out after failing to send them.
	Flush([]interface{})

	// Diagnostics provides counters and other feedback about the health of the Flusher.
	Diagnostics() map[string]interface{}
}
