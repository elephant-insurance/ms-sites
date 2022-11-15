package mbuf

import (
	"bytes"
	"errors"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
)

func NewDefaultHTTPMarshaler(serviceURL string) HTTPMarshaler {
	return &defaultHTTPMarshaler{
		relayURL:          serviceURL,
		failuresToMarshal: &clicker.Clicker{},
	}
}

// defaultHTTPMarshaler is the default implementation of the HTTPMarshaler interface
// It configures the message body as an array of JSON objects
// and sends messages to the specified endpoint via POST
type defaultHTTPMarshaler struct {
	relayURL          string
	failuresToMarshal *clicker.Clicker
}

func (m *defaultHTTPMarshaler) Marshal(msgs []interface{}) (*http.Request, error) {
	body, err := buildJSONArrayBody(msgs)
	if err != nil {
		m.failuresToMarshal.Click(1)
		return nil, err
	}

	rq, err := http.NewRequest(http.MethodPost, m.relayURL, bytes.NewReader(body))
	if err != nil {
		m.failuresToMarshal.Click(1)
		msg := errMsgFailedToCreateRequest
		debugLog(msg + `: ` + err.Error())
		return nil, errors.New(msg)
	}

	// rq.Header.Add automatically uppercases the key
	rq.Header.Add("Content-Type", "application/json")

	return rq, nil
}

// buildJSONArrayBody takes an array of individually-marshaled ([]byte) JSON items
// and returns a marshaled JSON array of those items
// This enables us to reuse a message we have already marshaled for use as an individual message
// instead of creating the array of (unmarshaled) stucts in RAM and re-marshaling as an array, saving CPU time.
func buildJSONArrayBody(entries []interface{}) ([]byte, error) {
	dataBuf := strings.Builder{}
	dataBuf.WriteString("[")
	firstEntry := true

	for i := 0; i < len(entries); i++ {
		thisEntry := entries[i]
		if thisEntry == nil {
			// shouldn't happen but no reason to get fussed over it
			continue
		}
		// turn the interface into a []byte, or die
		if thisba, ok := thisEntry.([]byte); ok {
			// we now have a byte slice
			if len(thisba) == 0 {
				// also shouldn't happen but also no problem
				continue
			}
			// we have something to write now
			if firstEntry {
				// don't write a comma this time, but do it next time
				firstEntry = false
			} else {
				// we've written before, so append a comma
				dataBuf.WriteString(",\n")
			}
			// append the msg itself
			dataBuf.Write(thisba)
		} else {
			// this indicates a coding bug, so error out
			msg := "message could not be cast to []byte"
			debugLog(msg + ": " + spew.Sdump(thisEntry))
			return nil, errors.New(msg)
		}
	}

	dataBuf.WriteString("]")

	return []byte(dataBuf.String()), nil
}

// Diagnostics returns a count of any failures to marshal we have experienced
func (m *defaultHTTPMarshaler) Diagnostics() map[string]interface{} {
	return map[string]interface{}{
		diagnosticsFieldFailuresToMarshal: m.failuresToMarshal.Clicks,
	}
}
