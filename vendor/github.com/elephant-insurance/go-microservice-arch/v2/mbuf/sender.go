package mbuf

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// msgSender abstracts away the send method, which handles the details of actually forwarding messages to their destination
// at runtime this method is ALWAYS called in a new goroutine, making it difficult to debug
// specifying an interface lets us prevent forking when we test
type msgSender interface {
	send(msgs []interface{})
	sendnow(msgs []interface{}) (status *int, respBody *string, err error)
	sendnownoretry(msgs []interface{})
}

// realMsgSender is the implementation we use when NOT testing
type realMsgSender struct {
	b *messageRelayType
}

// send just wraps the sendroutine method in a goroutine
// override this implementation to keep everything single-threaded in testing
func (r realMsgSender) send(msgs []interface{}) {
	go r.sendroutine(msgs, true)
}

// sendnow sends the messages on the calling thread with retries and returns basic call status info
func (r realMsgSender) sendnow(msgs []interface{}) (status *int, respBody *string, err error) {
	return r.sendroutine(msgs, true)
}

// sendnownoretry calls sendroutine in the calling thread with no retries
func (r realMsgSender) sendnownoretry(msgs []interface{}) {
	r.sendroutine(msgs, false)
}

// send forwards buffered messages to the log relayer
// it is intended to run in its own goroutine
// since log messages as data structures aren't exposed
// outside this package, we don't need to worry about threading here
// if RelayRetryWait is positive and send fails,
// it waits this many seconds then calls itself recursively for a retry.
// Returns true if it succeeds; this is only used for recursive retries
// Returning false means we're done trying with these messages,
// so be sure to call flush() to possibly send them to stdout
// THIS FUNCTION MUST NEVER BE CALLED WITH THE BUFFER SYNC LOCKED!
// Reason: this function potentially sync locks the relay manager
// If a thread locks both the app could stop running
// THIS METHOD MUST NEVER REFER TO b.messages BECAUSE IT IS NOT SYNCHRONIZED
// Instead it receives the slice of messages it needs to send as a parameter
func (r realMsgSender) sendroutine(msgs []interface{}, retry bool) (status *int, respBody *string, err error) {
	if len(msgs) == 0 {
		return
	}

	if r.b == nil {
		err = errors.New(`nil message relay`)
		return
	}

	b := r.b
	var (
		req *http.Request
		ok  bool
	)

	b.sends.Click(1)
	// should NEVER happen, but I'm superstitious
	if b.manager == nil {
		b.manager = newRelayManager()
	}

	if !r.b.manager.sending() {
		// we're not relaying for now
		msg := "send to relay canceled due to relay connection failures, flushing message(s)"
		debugLog(msg)
		err = errors.New(msg)

		// we're not going to try again with these, so flush them
		b.flush(msgs)

		return
	}

	debugLog("send relay URL is: " + r.b.relayURL)
	// debugLog(fmt.Sprintf(" the relay format is: %v", theRelayExchangeFormat))

	// format the messages
	var msgSlices = make([]interface{}, len(msgs))
	for i := 0; i < len(msgs); i++ {
		msgSlices[i] = b.renderer.Render(msgs[i])
	}

	// TODO: nil-check this so that we don't have to use a marshaler unless we actually need one
	req, err = b.marshaler.Marshal(msgSlices)

	if err != nil {
		// Never seen this happen, but if it does, let's just bounce
		debugLog("error creating http request: " + err.Error())
		b.flush(msgs)
		return
	}

	if req == nil {
		msg := "JSON marshaler returned nil with no error"
		err = errors.New(msg)
		b.flush(msgs)
		return
	}

	debugLog("attempting to send to: " + req.URL.String())

	if b.client == nil {
		b.client = &http.Client{}
	}

	// sendMsg will recursively retry until its retry count param reaches 0
	// if we set startingRetryCount = 0, then we will not retry after a single failed send
	startingRetryCount := 0

	// if we do want to retry, then set the count to b.maxRetries OR absoluteMaxRelayRetries, whichever is lower
	if retry {
		if b.maxRetries > absoluteMaxRelayRetries {
			startingRetryCount = absoluteMaxRelayRetries
		} else {
			startingRetryCount = b.maxRetries
		}
	}

	debugLog(fmt.Sprintf("sending with %v retries", startingRetryCount))
	ok, status, respBody, err = r.sendMsg(req, startingRetryCount)
	if !ok {
		debugLog("failed to send message buffer, flushing")
		b.flush(msgs)
		return
	}

	debugLog("buffer send successful")

	return
}

// sendMsg handles the low-level details of sending a request to the log processor
// if it gets an error back from sending, it increments the retry count and calls itself recursively after a wait
// returns true if finally successful, false otherwise
// this method DOES NOT call flush(), so the caller must flush messages if this returns false
// this method will recurse to a max depth of b.maxRetries or absoluteMaxRelayRetries, whichever is lower
func (r realMsgSender) sendMsg(rq *http.Request, retriesRemaining int) (success bool, status *int, respBody *string, err error) {
	success = false
	var resp *http.Response

	debugLog(fmt.Sprintf("sendMsg called with %v retries", retriesRemaining))
	b := r.b

	if b == nil || rq == nil || retriesRemaining < 0 {
		debugLog(`bad call to sendMsg, something is nil or we are out of retries`)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), b.clientTimeout)
	defer cancel()

	debugLog(`client sending...`)
	resp, err = b.client.Do(rq.WithContext(ctx))
	debugLog("...client sent")

	// an error back typically indicates a connection problem, which may be temporary
	if err != nil {
		b.errorsLogged.Click(1)
		debugLog(fmt.Sprintf("error connecting to relay: %v", err.Error()))

		b.manager.failure()
		if b.retryWaitSeconds > 0 && retriesRemaining > 0 {
			debugLog(fmt.Sprintf(" waiting %v seconds to attempt resend", b.retryWaitSeconds))

			time.Sleep(time.Duration(b.retryWaitSeconds) * time.Second)
			debugLog("attempting to resend after wait")
			b.retries.Click(1)

			// retry recursively until we either get back a good response (return true) or run out of retries (false)
			return r.sendMsg(rq, retriesRemaining-1)
		}

		// we're not either done retrying (retries == b.maxRetries) or not retrying at all
		return
	}

	// we didn't get back an error, but did we get a response?
	if resp == nil {
		// not much we can do with a nil response so let these go
		debugLog("NIL response")
		b.manager.failure()
		return
	}

	debugLog("non-nil response")
	defer resp.Body.Close()
	bodyBytes, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		debugLog("error reading response body: " + rerr.Error())
	} else {
		bodyString := string(bodyBytes)
		respBody = &bodyString
		debugLog("Response body:\n" + bodyString)
	}

	statusCode := resp.StatusCode
	debugLog(fmt.Sprintf(" status from relay: %v", statusCode))
	if statusCode > 0 {
		status = &statusCode
	}

	// This means we had a successful connection, but the recipient did not like the message
	// there is no point trying to send it again because it will just return the same thing again
	if statusCode > 299 {
		debugLog(fmt.Sprintf("bad status from log relay: %v (%v)", statusCode, resp.Status))
		b.manager.failure()
		b.badStatusFromRelay.Click(1)

		return
	}

	success = true
	debugLog("successful send to relay")
	b.success()

	return
}

// consoleSender simulates a messagerelay by printing its output to stdout
type consoleSender struct{}

func (s *consoleSender) send(msgs []interface{}) {
	s.sendnow(msgs)
}

func (s *consoleSender) sendnow(msgs []interface{}) {
	buf := bytes.Buffer{}
	buf.WriteString(`[`)

	if len(msgs) > 0 {
		buf.WriteString("\n")
		for _, thisMsg := range msgs {
			if thisMsg != nil {
				if jsonBytes, err := json.Marshal(thisMsg); err != nil {
					buf.WriteString("\t")
					buf.WriteString(string(jsonBytes))
					buf.WriteString(",\n")
				}
			}
		}
	}
	buf.WriteString(`]`)

	fmt.Println(buf.String())
}
