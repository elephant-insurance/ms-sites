package mbuf

import (
	"fmt"
	"sync"
	"time"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
)

/*

The relay manager is a "traffic cop" for the log message relay.

This package has been designed to allow the client app to continue to run
when the log-relay is down, however this will cause some slowdown as requests
to the log relay server must time out before execution can continue.

The relay manager tracks how many bad responses (errors) have been received
from the log message relayer since the last good response. It also exposes a "serving"
func which returns true ("green" status, serving logs to the relay) or false ("red"
status, NOT sending to the relay). Before attempting to send messages to the log relay,
the package will check the return value of serving(); if false it will not send.

Note that requests that return HTTP error status codes (4xx, 5xx) MUST NOT be treated as
failures by the relay manager! The reason: an error status indicates that the message was
received by the log processor, but that the processor cannot handle it for some reason,
usually due to a formatting or protocol error. No number of retries will convince the
log processor to handle a message it can't understand, so there is no point to retrying these.

Errors returned from the HTTP send, on the other hand, usually indicate network problems,
which we may reasonably expect to resolve in a short period of time. Such messages SHOULD be
re-sent.

*/

// SETTINGS
// These values are defaults, but are exposed in case different behavior is desired

// FailuresToShutdownRelay is how many failed sends to the relayer in a row it takes to stop sending
var FailuresToShutdownRelay = 3

// RelayShutdown1Duration is how long the manager stops sending to the relayer the first time it stops sending
var RelayShutdown1Duration = time.Duration(12) * time.Second

// RelayShutdown2Duration is how long the manager stops sending to the relayer the second time it stops sending
var RelayShutdown2Duration = time.Duration(120) * time.Second

// RelayShutdown3Duration is how long the manager stops sending to the relayer the third and subsequent times it stops sending
var RelayShutdown3Duration = time.Duration(30*60) * time.Second

type relayManager interface {
	sending() bool
	success()
	failure()
	diagnostics() map[string]interface{}
}

type relayManagerType struct {
	sync.Mutex
	// allowSending is a private toggle for whether the manager is allowing sends to the relay now or not
	allowSending bool
	// disabledUntil is a nilable timestamp for recording when the manager will allow attempts to send again
	disabledUntil time.Time
	// failures is the number of consecutive failures to send since the last success
	failures int
	// failuresWhenDisabled is the number of failures in a row we had the last time we shut down
	failuresWhenDisabled int
	// relayFailuresSinceStart tracks how many times have we gotten back any bad result
	relayFailuresSinceStart *clicker.Clicker
	// relayShutdownsSinceStart tracks how many times have we had to shutdown the relay due to errors
	relayShutdownsSinceStart *clicker.Clicker
	// shutdowns is how many times we've shut down sending since the last successful send
	shutdowns int
}

// sending returns false when sending is disabled, true otherwise
func (m *relayManagerType) sending() bool {
	if m.allowSending {
		return true
	}

	// check to see whether it's time to try again
	if m.disabledUntil.Before(time.Now()) {
		m.Lock()
		defer m.Unlock()

		// now we're single-threaded, so check one more time before we do anything
		if !m.allowSending && m.disabledUntil.Before(time.Now()) {
			// time to try again
			debugLog("relay shutdown timed out, trying again to connect")

			m.allowSending = true
			m.disabledUntil = time.Time{}
			return true
		}
	}

	return false
}

// success is called to record a successful send
func (m *relayManagerType) success() {
	if m.allowSending && m.failures == 0 {
		return
	}

	m.Lock()
	defer m.Unlock()
	// one success resets everything
	m.allowSending = true
	m.disabledUntil = time.Time{}
	m.failures = 0
	m.failuresWhenDisabled = 0
	m.shutdowns = 0
}

// failure is called to record a failed attempt to send
func (m *relayManagerType) failure() {
	// nothing to do if we're already turned off
	if !m.allowSending {
		return
	}

	m.Lock()
	defer m.Unlock()
	if m.allowSending {
		// don't want to increment this if we've stopped sending already
		m.relayFailuresSinceStart.Click(1)
		m.failures++
		debugLog(fmt.Sprintf("relay manager logging failure %v", m.failures))
	}

	// have we crossed the threshold?
	if m.failures >= m.failuresWhenDisabled+FailuresToShutdownRelay {
		m.allowSending = false
		m.failuresWhenDisabled = m.failures
		m.relayShutdownsSinceStart.Click(1)
		m.shutdowns++
		switch m.shutdowns {
		case 1:
			m.disabledUntil = time.Now().Add(RelayShutdown1Duration)
			debugLog(fmt.Sprintf(" disabling relay for %v due to connection failures", RelayShutdown1Duration))

		case 2:
			m.disabledUntil = time.Now().Add(RelayShutdown2Duration)
			debugLog(fmt.Sprintf(" disabling relay for %v due to connection failures", RelayShutdown2Duration))

		default:
			m.disabledUntil = time.Now().Add(RelayShutdown3Duration)
			debugLog(fmt.Sprintf(" disabling relay for %v due to connection failures", RelayShutdown3Duration))
		}
	}
}

func newRelayManager() relayManager {
	return &relayManagerType{
		allowSending:             true,
		disabledUntil:            time.Time{},
		failures:                 0,
		failuresWhenDisabled:     0,
		relayFailuresSinceStart:  &clicker.Clicker{},
		relayShutdownsSinceStart: &clicker.Clicker{},
		shutdowns:                0,
	}
}

func (m *relayManagerType) diagnostics() map[string]interface{} {
	return map[string]interface{}{
		diagnosticsFieldTotalFailures:  m.relayFailuresSinceStart.Clicks,
		diagnosticsFieldTotalShutdowns: m.relayShutdownsSinceStart.Clicks,
	}
}
