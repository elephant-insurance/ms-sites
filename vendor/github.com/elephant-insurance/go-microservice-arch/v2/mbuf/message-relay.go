package mbuf

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
)

/*

This is a buffer designed to reduce network overhead
by reducing the number of http sends to a given log-handling service.

The message relay holds everything necessary for sending a buffered slice of messages
to a given log processor. It caches messages until they are ready to send,
creates the HTTP request, handles filures and resends, and flushes messages that could not be sent.

The relayBuffer type contains a sync.Mutex for restricting access to a single thread.
The relayBuffer has two synchronized operations: add and clear.
Add adds a new message struct to the buffer. This happens on the main thread and returns instantly.
Clear sends the current buffer to a new thread for sending, then resets the buffer.

Add should be called to queue up a new message for sending.
Clear should only be called by a timer thread that flushes the buffer every n seconds.

The relayBuffer.send([]interface{}) operation is NOT synchronized, but it does not not need to be
so long as it does not refer to b.messages.
Instead, it receives messages to send as a function parameter.

The design is intended to reduce thread contention by keeping synched operations very
simple and very fast.
*/

// messageRelayType is the "real" implementation of the relayBuffer interface
type messageRelayType struct {
	sync.Mutex
	badStatusFromRelay *clicker.Clicker
	client             *http.Client
	clientTimeout      time.Duration
	errorsLogged       *clicker.Clicker
	flushers           []Flusher
	marshaler          HTTPMarshaler
	manager            relayManager
	maxLength          int
	maxRetries         int
	messages           []interface{}
	pollInterval       time.Duration
	renderer           MessageRenderer
	relayURL           string
	retries            *clicker.Clicker
	retryWaitSeconds   int
	sender             msgSender
	sends              *clicker.Clicker
	successfulSends    *clicker.Clicker
	synchronousMode    bool
}

// Add implements the MessageRelay interface
// it adds a new message to the buffer thread-safely
// if the new message pushes the buffer to capacity,
// we fire off a new thread to send the messages
func (b *messageRelayType) Add(msg interface{}) {
	if msg == nil {
		return
	}

	// no need to lock if we're in single-message mode, just send
	if b.maxLength < 2 {
		go b.send([]interface{}{msg})
		return
	}

	b.Lock()
	defer b.Unlock()
	b.messages = append(b.messages, msg)

	if len(b.messages) >= b.maxLength {
		// this message pushed us to the limit
		// so while we're locked we're going to fire off a clear and send
		// can't call clear() directly as it would try to re-lock (bad!)
		msgs := b.messages
		b.messages = []interface{}{}

		// send is not mutex-locked, so no worries
		go b.send(msgs)
		debugLog(fmt.Sprintf(" sending full buffer of %v messages to relay", len(msgs)))
	}
}

func (b *messageRelayType) SendNow(msgs []interface{}) (status *int, respBody *string, err error) {
	return b.sender.sendnow(msgs)
}

// SendNowNoRetry sends any cached messages immediately on the calling thread without retries.
// Use this to ensure that everything is sent before exiting the program.
// This method does not lock the buffer before clearing it,
// to prevent a potential thread-lock when flushing unsendable messages.
// Therefore it could in very rare cases panic due to a race condition
// and should only be called right before the program exits.
func (b *messageRelayType) SendNowNoRetry() {
	// always have to check again inside the lock
	if len(b.messages) > 0 {
		msgs := b.messages
		b.messages = []interface{}{}
		b.sender.sendnownoretry(msgs)
	}
}

// clear returns the currently-buffered messages and resets the buffer thread-safely
func (b *messageRelayType) clear() []interface{} {
	// probably best to send on this calling timer thread
	// since no client is waiting on it

	// no need to lock if we have no messages
	// even if another thread is currently adding one, we will get
	// back around to it next time the timer fires
	if len(b.messages) == 0 {
		return nil
	}

	// we have messages, so lock and clean up as quickly as possible
	b.Lock()
	defer b.Unlock()

	// check again, now that we have the lock
	if len(b.messages) == 0 {
		return nil
	}

	rtn := b.messages
	b.messages = make([]interface{}, 0, b.maxLength)

	return rtn
}

// send forwards buffered messages to the log relayer
// for testing, provide an alternate implementattion of sender
func (b *messageRelayType) send(msgs []interface{}) {
	b.sender.send(msgs)
}

// start fires off a thread to monitor this buffer for unsent messages
func (b *messageRelayType) start() {
	// never start a thread unless we need it
	// pollInterval == 0 means we don't want a polling thread
	// maxLength < 2 means we never have any meessages in the buffer, so no need to poll
	if b != nil && b.pollInterval > 0 && b.maxLength > 1 {
		go func() {
			for range time.Tick(b.pollInterval) {
				// clear only locks the buffer and returns messages if there are some available
				msgs := b.clear()
				lm := len(msgs)
				debugLog(fmt.Sprintf("fired for buffer with %v messages", lm))
				if lm > 0 {
					b.send(msgs)
				}
			}
		}()
	}
}

// success registers a successful send by incrementing our success counter and
// letting the relay manager know all is well
func (b *messageRelayType) success() {
	b.successfulSends.Click(1)
	b.manager.success()
}

// flush runs each Flusher that we have for any messages that we could not send
func (b *messageRelayType) flush(msgs []interface{}) {
	for i := 0; i < len(b.flushers); i++ {
		if b.flushers[i] != nil {
			b.flushers[i].Flush(msgs)
		}
	}
}

// Diagnostics cycles through our members and reports back interesting things.
func (b *messageRelayType) Diagnostics() map[string]interface{} {
	rtn := map[string]interface{}{
		diagnosticsFieldBadResponses:      b.badStatusFromRelay.Clicks,
		diagnosticsFieldErrorResponses:    b.errorsLogged.Clicks,
		diagnosticsFieldTotalRetries:      b.retries.Clicks,
		diagnosticsFieldTotalSendAttempts: b.sends.Clicks,
		diagnosticsFieldSuccessfulSends:   b.successfulSends.Clicks,
	}

	for k, v := range b.renderer.Diagnostics() {
		rtn[k] = v
	}

	for k, v := range b.marshaler.Diagnostics() {
		rtn[k] = v
	}

	for k, v := range b.manager.diagnostics() {
		rtn[k] = v
	}

	for i := 0; i < len(b.flushers); i++ {
		for k, v := range b.flushers[i].Diagnostics() {
			rtn[k+strconv.Itoa(i)] = v
		}
	}

	return rtn
}

// new sets up a basic message relay with defaults and overrides with the settings we pass in
// it DOES NOT attempt to start the relay
func newMBUF(s *Settings) *messageRelayType {
	ru := ``
	if s != nil {
		ru = s.RelayURL
	}

	rtn := messageRelayType{
		badStatusFromRelay: &clicker.Clicker{},
		client:             &http.Client{},
		clientTimeout:      time.Duration(defaultClientTimeoutSecs) * time.Second,
		errorsLogged:       &clicker.Clicker{},
		flushers:           []Flusher{NewDefaultFlusher()},
		manager:            newRelayManager(),
		maxLength:          defaultMaxBufferLength,
		maxRetries:         defaultMaxRetries,
		messages:           make([]interface{}, 0, defaultMaxBufferLength),
		pollInterval:       time.Duration(defaultPollIntervalSecs) * time.Second,
		relayURL:           ru,
		renderer:           NewDefaultMessageRenderer(),
		retries:            &clicker.Clicker{},
		retryWaitSeconds:   defaultRetryWaitSecs,
		sends:              &clicker.Clicker{},
		successfulSends:    &clicker.Clicker{},
	}

	rtn.sender = &realMsgSender{&rtn}

	if s == nil {
		return &rtn
	}

	if s.ClientTimeoutSecs != nil {
		rtn.clientTimeout = time.Duration(*s.ClientTimeoutSecs) * time.Second
	}

	if s.Flushers != nil {
		rtn.flushers = s.Flushers
	}

	if s.Marshaler != nil {
		rtn.marshaler = s.Marshaler
	} else {
		rtn.marshaler = NewDefaultHTTPMarshaler(s.RelayURL)
	}

	if s.MaxLength != nil {
		rtn.maxLength = *s.MaxLength
	}

	if s.MaxRetries != nil {
		rtn.maxRetries = *s.MaxRetries
	}

	if s.PollIntervalSecs != nil {
		rtn.pollInterval = time.Duration(*s.PollIntervalSecs) * time.Second
	}

	if s.Renderer != nil {
		rtn.renderer = s.Renderer
	}

	// RetryWaitSeconds is how many seconds we wait before attempting to retry a failed message
	if s.RetryWaitSeconds != nil {
		rtn.retryWaitSeconds = *s.RetryWaitSeconds
	}

	return &rtn
}
