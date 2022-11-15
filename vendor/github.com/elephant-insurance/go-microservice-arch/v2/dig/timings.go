package dig

import (
	"net/http"
	"time"

	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
	"github.com/elephant-insurance/go-microservice-arch/v2/timing"
)

// these are types and functions for tracking timings
// the underlying functionality is provided by the timing package
// here we wrap that functionality in types and funcs that are also
// aware of our logging and diagnostics namespaces

type timer struct {
	context msrqc.Context
	t       *timing.Timer
}

// Timer is a dumb interface to wrap the timer struct.
type Timer interface {
	// Start starts the timer, if the timer was created in a stopped state.
	// We can pass in a start time so that timings are perfectly synchronized
	// or we can leave the parameter blank, setting it to Now().
	Start(timeInitiated *time.Time) Timer

	// Stop stops the timing, records the result, and returns the Timing with elapsed time in microseconds.
	Stop(status int) *timing.Timing

	// StopResponse will log a zero status for a nil response, resp.StatusCode otherwise.
	StopResponse(resp *http.Response) *timing.Timing
}

func (t *timer) Start(timeInitiated *time.Time) Timer {
	if t == nil || t.t == nil {
		return t
	}

	t.t.Start(timeInitiated)

	return t
}

// Stop calls Stop(code) on the embedded Timer and also logs a timing record
func (t *timer) Stop(code int) *timing.Timing {
	if t == nil || t.t == nil {
		return nil
	}

	rtn := t.t.Stop(code)
	log.WriteTiming(t.context, rtn)

	return rtn
}

func (t *timer) StopResponse(resp *http.Response) *timing.Timing {
	if t == nil || t.t == nil {
		return nil
	}

	if resp == nil {
		return t.Stop(0)
	}

	return t.Stop(resp.StatusCode)
}

// StartServiceTiming starts a timer for a request that we are fulfilling
// For these timings, path should be the route pattern chosen by the router (e.g. "GET /ms-my-srv/products/:productId")
// label should uniquely identify the action (e.g., "GetCustomer")
func StartServiceTiming(c msrqc.Context, label, path *string) Timer {
	return StartTiming(c, timing.TimingTypeWebService, label, path)
}

// StartClientTiming starts a timer for a request that we are making to another service
// For these timings, path should be the URL to the service we're calling
func StartClientTiming(c msrqc.Context, label, path *string) Timer {
	return StartTiming(c, timing.TimingTypeWebClient, label, path)
}

// StartDatabaseTiming starts a timer for a database query
// For these timings, path should be the fully-qualified name of the table or proc we're hitting
func StartDatabaseTiming(c msrqc.Context, label, path *string) Timer {
	return StartTiming(c, timing.TimingTypeDatabaseQuery, label, path)
}

// StartTiming starts a timing of the designated TimingType
func StartTiming(c msrqc.Context, ttype timing.TimingType, label, path *string) Timer {
	tt := timing.NewTimer(ttype, label, path)
	tt.Start(nil)

	return &timer{
		t:       tt,
		context: c,
	}
}
