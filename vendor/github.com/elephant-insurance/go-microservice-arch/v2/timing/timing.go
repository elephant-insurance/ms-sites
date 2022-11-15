package timing

import (
	"fmt"
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
)

// Timing is a record of a single timed event
// It records the time, duration, status, and route for statistical analysis
type Timing struct {
	// SequenceNumber is just a serial number for the record
	// With this, the results from multiple calls can be correlated and de-duped
	SequenceNumber uint64 `json:"SequenceNumber"`
	// Label is a simple description of the event for filtering
	Label *string `json:"Label"`
	// Path is the API path for the request, as submitted to the router
	Path *string `json:"Path"`
	// Status is the status code returned by the API for this call
	Status int `json:"Status"`
	// Target captures the identity of the service or database we are timing
	Target *enum.ServiceID
	// TimeInitiated is the exact time the request reached the controller
	TimeInitiated time.Time `json:"TimeInitiated"`
	// Duration is how long the request took to complete, in microseconds
	DurationMicroseconds int64 `json:"DurationMicroseconds"`
	// Type is the broad type of timing, for easier filtering
	Type TimingType `json:"TimingType"`
	// TransactionID records the transaction id for the timing, if applicable
	TransactionID *string `json:"TransactionID"`
}

const (
	// TimingCSVHeader is returned as the first row when timings are requested in CSV format
	TimingCSVHeader = "SequenceNumber,Label,Path,Status,TimeInitiated,DurationMicroseconds,Type"

	dummyTimingLabel = `dummy-timing-label`
)

type TimingType string

const (
	// TimingTypeUnknown is the default, probably should not use
	TimingTypeUnknown TimingType = `unknown`
	// TimingTypeWebService is a timing of an http call TO this app
	TimingTypeWebService TimingType = `web-service`
	// TimingTypeWebClient is a timing of an http call FROM this app
	TimingTypeWebClient TimingType = `web-client`
	// TimingTypeDatabaseQuery is a timing of a database call from this app
	TimingTypeDatabaseQuery TimingType = `db-client`
	// TimingTypeTestWebClient is a timing of a TEST http call FROM this app
	TimingTypeTestWebClient TimingType = `test-web-client`
	// TimingTypeTestDatabaseQuery is a timing of a TEST database call from this app
	TimingTypeTestDatabaseQuery TimingType = `test-db-client`
)

func (tt TimingType) MarshalJSON() ([]byte, error) {
	rtn := TimingTypeUnknown
	switch tt {
	case TimingTypeWebService, TimingTypeWebClient, TimingTypeDatabaseQuery, TimingTypeTestWebClient, TimingTypeTestDatabaseQuery:
		rtn = tt
	}

	return []byte(`"` + rtn + `"`), nil
}

// ToCSV turns a timing into a string we can import into Excel
func (t *Timing) ToCSV() string {
	if t == nil {
		return ""
	}
	label := ""
	if t.Label != nil {
		label = *t.Label
	}
	path := ""
	if t.Path != nil {
		path = *t.Path
	}
	target := ""
	if t.Target != nil {
		target = string(*t.Target)
	}

	return fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v", t.SequenceNumber, label, target, path, t.Status, t.TimeInitiated.Format(time.RFC3339Nano), t.DurationMicroseconds, t.Type)
}

func newDummyTiming() interface{} {
	foo := dummyTimingLabel
	return &Timing{
		Label: &foo,
		Type:  TimingTypeUnknown,
	}
}
