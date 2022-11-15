package timing

import (
	"net/http"
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
)

// timer is a simple struct used to create a Timing
type Timer struct {
	Label         *string
	Path          *string
	Target        *enum.ServiceID
	TimeInitiated *time.Time
	TimingType    TimingType
}

// stop signals the end of the event and turns a timer object into a Timing that we can store
func (t *Timer) stop(status int) *Timing {
	if t == nil || t.TimeInitiated == nil {
		return nil
	}

	return &Timing{
		DurationMicroseconds: time.Since(*t.TimeInitiated).Microseconds(),
		Label:                t.Label,
		Path:                 t.Path,
		Status:               status,
		Target:               t.Target,
		TimeInitiated:        *t.TimeInitiated,
		Type:                 t.TimingType,
	}
}

// NewTimer creates a timer without starting it (setting the time initiated field)
func NewTimer(ttype TimingType, label, path *string) *Timer {
	return &Timer{
		Label:      label,
		Path:       path,
		TimingType: ttype,
	}
}

// StartServiceTiming starts a timer for a request that we are fulfilling
// For these timings, path should be the route pattern chosen by the router (e.g. "GET /ms-my-srv/products/:productId")
// label should uniquely identify the action (e.g., "GetCustomer")
func StartServiceTiming(label, path *string) *Timer {
	rightNow := time.Now()

	rtn := NewTimer(TimingTypeWebService, label, path)

	return rtn.Start(&rightNow)
}

// StartClientTiming starts a timer for a request that we are making to another service
// For these timings, path should be the URL to the service we're calling
func StartClientTiming(label, path *string, target *enum.ServiceID) *Timer {
	rightNow := time.Now()

	rtn := NewTimer(TimingTypeWebClient, label, path)
	rtn.Target = target

	return rtn.Start(&rightNow)
}

// StartDatabaseTiming starts a timer for a database query
// For these timings, path should be the fully-qualified name of the table or proc we're hitting
func StartDatabaseTiming(label, path *string, target *enum.ServiceID) *Timer {
	rightNow := time.Now()

	rtn := NewTimer(TimingTypeDatabaseQuery, label, path)
	rtn.Target = target

	return rtn.Start(&rightNow)
}

func (t *Timer) Start(timeInitiated *time.Time) *Timer {
	if t == nil || t.TimeInitiated != nil {
		return t
	}

	var started time.Time
	if timeInitiated != nil {
		started = *timeInitiated
	} else {
		started = time.Now()
	}

	t.TimeInitiated = &started

	return t
}

// Stop calls stop on the timer and stores the timing it returns.
// Status can be an HTTP status code or any other int that signals the outcome of the event.
func (t *Timer) Stop(status int) *Timing {
	if t == nil {
		return nil
	}

	newTiming := t.stop(status)
	ApplicationTimings.Add(newTiming)

	return newTiming
}

// StopResponse is a convenience func that allows us to stop a timer before we know whether we got back a nil response.
func (t *Timer) StopResponse(resp *http.Response) *Timing {
	if t == nil {
		return nil
	}

	if resp == nil {
		return t.Stop(0)
	}

	return t.Stop(resp.StatusCode)
}
