package mbuf

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
)

// TestMessageRelay is an implementation of MessageRelay for testing.
// It remembers what has been Added to it and offers methods for querying its memory.
// Use this as the backend for a log to test exactly what is sent to the relay.
type TestMessageRelay struct {
	Messages   []*MessageItem
	AddedCount *clicker.Clicker
}

func NewTestMessageRelay() *TestMessageRelay {
	return &TestMessageRelay{
		Messages:   []*MessageItem{},
		AddedCount: &clicker.Clicker{},
	}
}

// Add adds a single message object to this MessageRelay's buffer.
func (tmr *TestMessageRelay) Add(msg interface{}) {
	if tmr == nil {
		return
	}

	tmr.Messages = append(tmr.Messages, newMessageItem(msg))
	tmr.AddedCount.Click(1)
}

// Diagnostics provides counters and other feedback about the health of the buffer.
// Since most error conditions are not returned to the caller, this is the caller's primary means of checking the health of the MessageRelay and its children.
func (tmr *TestMessageRelay) Diagnostics() map[string]interface{} {
	rtn := map[string]interface{}{
		`Added`:       tmr.AddedCount.Clicks,
		`LastMessage`: tmr.LastMessageString(),
	}

	return rtn
}

// SendNowNoRetry causes all meessages to be sent immediately on the calling thread.
// Use this to clean up before exiting.
func (tmr *TestMessageRelay) SendNowNoRetry() {}

func (tmr *TestMessageRelay) SendNow(msgs []interface{}) (status *int, respBody *string, err error) {
	return
}

// MessageItem is a record of a message submitted to a TestMessageRelay via Add(msg).
type MessageItem struct {
	Item       interface{}
	JSONMap    map[string]string
	JSONString string
}

func (tmr *TestMessageRelay) LastMessage() *MessageItem {
	if tmr != nil && len(tmr.Messages) > 0 {
		return tmr.Messages[len(tmr.Messages)-1]
	}
	return nil
}

func (tmr *TestMessageRelay) LastMessageHasKey(key string) bool {
	lm := tmr.LastMessage()
	if lm == nil {
		return false
	}

	return lm.HasKey(key)
}

func (tmr *TestMessageRelay) LastMessageHasKeyValue(key, value string) bool {
	lm := tmr.LastMessage()
	if lm == nil {
		return false
	}

	return lm.HasKeyValue(key, value)
}

func (tmr *TestMessageRelay) LastMessageString() string {
	lms := ``
	lm := tmr.LastMessage()
	if lm != nil {
		lms = lm.JSONString
	}

	return lms
}

// HasKey verifies that the message has the submitted key
func (mi *MessageItem) HasKey(key string) bool {
	if mi == nil {
		return false
	}

	v, found := mi.JSONMap[key]

	return found && v != ``
}

// HasKeyValue verifies that the message has the submitted key/value pair
func (mi *MessageItem) HasKeyValue(key, value string) bool {
	if mi == nil {
		return false
	}

	v, found := mi.JSONMap[key]

	return found && strings.Contains(v, value)
}

// HasValues verifies that the message has the submitted key/value pairs
func (mi *MessageItem) HasKeyValues(valMap map[string]string) (allFound bool, missing error) {
	if mi == nil {
		return false, errors.New(`attempted to check values in nil MessageItem`)
	}
	allFound = true
	errs := []string{}

	for k, v := range valMap {
		if !mi.HasKeyValue(k, v) {
			errs = append(errs, `field not matched: `+k)
			allFound = false
		}
	}

	if len(errs) > 0 {
		missing = errors.New(strings.Join(errs, `, `))
	}

	return
}

func newMessageItem(msg interface{}) *MessageItem {
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		return nil
	}

	jsonString := string(jsonBytes)
	jsonMap := map[string]json.RawMessage{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		return nil
	}

	outMap := make(map[string]string, len(jsonMap))
	for k, v := range jsonMap {
		outMap[k] = string(v)
	}

	return &MessageItem{
		Item:       msg,
		JSONMap:    outMap,
		JSONString: jsonString,
	}
}
