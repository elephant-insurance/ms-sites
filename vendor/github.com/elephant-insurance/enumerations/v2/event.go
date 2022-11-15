package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// EventID uniquely identifies a particular Event
type EventID string

// Clone creates a safe, independent copy of a EventID
func (i *EventID) Clone() *EventID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two EventIds are equivalent
func (i *EventID) Equals(j *EventID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *EventID that is either valid or nil
func (i *EventID) ID() *EventID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *EventID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the EventID corresponds to a recognized Event
func (i *EventID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return Event.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *EventID) ValidatedID() *ValidatedEventID {
	if i != nil {
		return &ValidatedEventID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *EventID) MarshalJSON() ([]byte, error) {
	if i == nil || *i == "" {
		return []byte("null"), nil
	}

	if !i.Valid() {
		err := errors.New(errorMarshalInvalidID)
		return nil, err
	}

	istring := string(*i)

	return []byte(`"` + istring + `"`), nil
}

func (i *EventID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := EventID(dataString)
	item := Event.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	eventApplicationDiagnosticID            EventID = "app-dig"
	eventApplicationPanicID                 EventID = "app-panic"
	eventApplicationStartupID               EventID = "app-start"
	eventServiceRequestTimeoutID            EventID = "service-timeout"
	eventServiceRequestErrorID              EventID = "service-error"
	eventServiceRequestNilResponseID        EventID = "service-nil-response"
	eventServiceRequestBadStatusID          EventID = "service-bad-status"
	eventDatabaseRequestTimeoutID           EventID = "database-timeout"
	eventDatabaseClientRequestErrorID       EventID = "database-error"
	eventDatabaseClientRequestNilResponseID EventID = "database-nil-response"
	eventDebugEventID                       EventID = "debug"
	eventServiceRequestInvalidID            EventID = "server-request-invalid"
	eventServiceRequestFailureID            EventID = "server-request-failure"
	eventServiceRequestFullSuccessID        EventID = "server-request-success"
	eventServiceRequestPartialSuccessID     EventID = "server-request-partial-success"
)

// EnumEventItem describes an entry in an enumeration of Event
type EnumEventItem struct {
	ID        EventID           `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	Severity *EnumEventSeverityItem
}

var (
	eventApplicationDiagnostic            = EnumEventItem{eventApplicationDiagnosticID, "the application fired an event to record its memory usage and other statistics", map[string]string{EventMetaSeverityKey: string(eventSeverityDiagnosticID)}, "ApplicationDiagnostic", 1, &eventSeverityDiagnostic}
	eventApplicationPanic                 = EnumEventItem{eventApplicationPanicID, "the application has recovered from a panic error", map[string]string{EventMetaSeverityKey: string(eventSeverityErrorID)}, "ApplicationPanic", 2, &eventSeverityError}
	eventApplicationStartup               = EnumEventItem{eventApplicationStartupID, "the application has started", map[string]string{EventMetaSeverityKey: string(eventSeverityInformationalID)}, "ApplicationStartup", 3, &eventSeverityInformational}
	eventServiceRequestTimeout            = EnumEventItem{eventServiceRequestTimeoutID, "an HTTP request from this app to TARGET timed out", map[string]string{EventMetaSeverityKey: string(eventSeverityErrorID)}, "ServiceRequestTimeout", 4, &eventSeverityError}
	eventServiceRequestError              = EnumEventItem{eventServiceRequestErrorID, "an HTTP request from this app to TARGET raised an error", map[string]string{EventMetaSeverityKey: string(eventSeverityErrorID)}, "ServiceRequestError", 5, &eventSeverityError}
	eventServiceRequestNilResponse        = EnumEventItem{eventServiceRequestNilResponseID, "an HTTP request from this app to TARGET received a nil response", map[string]string{EventMetaSeverityKey: string(eventSeverityErrorID)}, "ServiceRequestNilResponse", 6, &eventSeverityError}
	eventServiceRequestBadStatus          = EnumEventItem{eventServiceRequestBadStatusID, "an HTTP request from this app to TARGET received a bad status code", map[string]string{EventMetaSeverityKey: string(eventSeverityErrorID)}, "ServiceRequestBadStatus", 7, &eventSeverityError}
	eventDatabaseRequestTimeout           = EnumEventItem{eventDatabaseRequestTimeoutID, "a database request from this app to TARGET timed out", map[string]string{EventMetaSeverityKey: string(eventSeverityErrorID)}, "DatabaseRequestTimeout", 8, &eventSeverityError}
	eventDatabaseClientRequestError       = EnumEventItem{eventDatabaseClientRequestErrorID, "a database request from this app to TARGET raised an error", map[string]string{EventMetaSeverityKey: string(eventSeverityErrorID)}, "DatabaseClientRequestError", 9, &eventSeverityError}
	eventDatabaseClientRequestNilResponse = EnumEventItem{eventDatabaseClientRequestNilResponseID, "a database request from this app to TARGET received a nil response", map[string]string{EventMetaSeverityKey: string(eventSeverityErrorID)}, "DatabaseClientRequestNilResponse", 10, &eventSeverityError}
	eventDebugEvent                       = EnumEventItem{eventDebugEventID, "an event triggered solely for testing event log connectivity and function", map[string]string{EventMetaSeverityKey: string(eventSeverityDevelopmentalID)}, "DebugEvent", 11, &eventSeverityDevelopmental}
	eventServiceRequestInvalid            = EnumEventItem{eventServiceRequestInvalidID, "an HTTP request to this service was invalid", map[string]string{EventMetaSeverityKey: string(eventSeverityCautionaryID)}, "ServiceRequestInvalid", 12, &eventSeverityCautionary}
	eventServiceRequestFailure            = EnumEventItem{eventServiceRequestFailureID, "an HTTP request from this app to TARGET failed for unspecified reasons", map[string]string{EventMetaSeverityKey: string(eventSeverityCautionaryID)}, "ServiceRequestFailure", 13, &eventSeverityCautionary}
	eventServiceRequestFullSuccess        = EnumEventItem{eventServiceRequestFullSuccessID, "an HTTP request from this app to TARGET service was fully successful", map[string]string{EventMetaSeverityKey: string(eventSeverityInformationalID)}, "ServiceRequestFullSuccess", 14, &eventSeverityInformational}
	eventServiceRequestPartialSuccess     = EnumEventItem{eventServiceRequestPartialSuccessID, "an HTTP request from this app to TARGET service was partly successful", map[string]string{EventMetaSeverityKey: string(eventSeverityCautionaryID)}, "ServiceRequestPartialSuccess", 15, &eventSeverityCautionary}
)

// EnumEvent is a collection of Event items
type EnumEvent struct {
	Description string
	Items       []*EnumEventItem
	Name        string

	ApplicationDiagnostic            *EnumEventItem
	ApplicationPanic                 *EnumEventItem
	ApplicationStartup               *EnumEventItem
	ServiceRequestTimeout            *EnumEventItem
	ServiceRequestError              *EnumEventItem
	ServiceRequestNilResponse        *EnumEventItem
	ServiceRequestBadStatus          *EnumEventItem
	DatabaseRequestTimeout           *EnumEventItem
	DatabaseClientRequestError       *EnumEventItem
	DatabaseClientRequestNilResponse *EnumEventItem
	DebugEvent                       *EnumEventItem
	ServiceRequestInvalid            *EnumEventItem
	ServiceRequestFailure            *EnumEventItem
	ServiceRequestFullSuccess        *EnumEventItem
	ServiceRequestPartialSuccess     *EnumEventItem

	itemDict map[string]*EnumEventItem
}

// Event is a public singleton instance of EnumEvent
// representing microservice application events
var Event = &EnumEvent{
	Description: "microservice application events",
	Items: []*EnumEventItem{
		&eventApplicationDiagnostic,
		&eventApplicationPanic,
		&eventApplicationStartup,
		&eventServiceRequestTimeout,
		&eventServiceRequestError,
		&eventServiceRequestNilResponse,
		&eventServiceRequestBadStatus,
		&eventDatabaseRequestTimeout,
		&eventDatabaseClientRequestError,
		&eventDatabaseClientRequestNilResponse,
		&eventDebugEvent,
		&eventServiceRequestInvalid,
		&eventServiceRequestFailure,
		&eventServiceRequestFullSuccess,
		&eventServiceRequestPartialSuccess,
	},
	Name:                             "EnumEvent",
	ApplicationDiagnostic:            &eventApplicationDiagnostic,
	ApplicationPanic:                 &eventApplicationPanic,
	ApplicationStartup:               &eventApplicationStartup,
	ServiceRequestTimeout:            &eventServiceRequestTimeout,
	ServiceRequestError:              &eventServiceRequestError,
	ServiceRequestNilResponse:        &eventServiceRequestNilResponse,
	ServiceRequestBadStatus:          &eventServiceRequestBadStatus,
	DatabaseRequestTimeout:           &eventDatabaseRequestTimeout,
	DatabaseClientRequestError:       &eventDatabaseClientRequestError,
	DatabaseClientRequestNilResponse: &eventDatabaseClientRequestNilResponse,
	DebugEvent:                       &eventDebugEvent,
	ServiceRequestInvalid:            &eventServiceRequestInvalid,
	ServiceRequestFailure:            &eventServiceRequestFailure,
	ServiceRequestFullSuccess:        &eventServiceRequestFullSuccess,
	ServiceRequestPartialSuccess:     &eventServiceRequestPartialSuccess,

	itemDict: map[string]*EnumEventItem{
		strings.ToLower(string(eventApplicationDiagnosticID)):            &eventApplicationDiagnostic,
		strings.ToLower(string(eventApplicationPanicID)):                 &eventApplicationPanic,
		strings.ToLower(string(eventApplicationStartupID)):               &eventApplicationStartup,
		strings.ToLower(string(eventServiceRequestTimeoutID)):            &eventServiceRequestTimeout,
		strings.ToLower(string(eventServiceRequestErrorID)):              &eventServiceRequestError,
		strings.ToLower(string(eventServiceRequestNilResponseID)):        &eventServiceRequestNilResponse,
		strings.ToLower(string(eventServiceRequestBadStatusID)):          &eventServiceRequestBadStatus,
		strings.ToLower(string(eventDatabaseRequestTimeoutID)):           &eventDatabaseRequestTimeout,
		strings.ToLower(string(eventDatabaseClientRequestErrorID)):       &eventDatabaseClientRequestError,
		strings.ToLower(string(eventDatabaseClientRequestNilResponseID)): &eventDatabaseClientRequestNilResponse,
		strings.ToLower(string(eventDebugEventID)):                       &eventDebugEvent,
		strings.ToLower(string(eventServiceRequestInvalidID)):            &eventServiceRequestInvalid,
		strings.ToLower(string(eventServiceRequestFailureID)):            &eventServiceRequestFailure,
		strings.ToLower(string(eventServiceRequestFullSuccessID)):        &eventServiceRequestFullSuccess,
		strings.ToLower(string(eventServiceRequestPartialSuccessID)):     &eventServiceRequestPartialSuccess,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumEvent) ByID(id EventIdentifier) *EnumEventItem {
	if e == nil || id == nil {
		return nil
	}

	if idx := id.ID(); idx != nil {
		idxString := strings.ToLower(string(*idx))

		if rtn, ok := e.itemDict[idxString]; ok {
			return rtn
		}
	}

	return nil
}

// ByIDString retrieves an entry by a string representation of its ID
func (e *EnumEvent) ByIDString(idx string) *EnumEventItem {
	if e == nil || len(e.itemDict) == 0 || idx == "" {
		return nil
	}

	normIdx := strings.ToLower(idx)
	if rtn, ok := e.itemDict[normIdx]; ok {
		return rtn
	}

	return nil
}

// ByIndex retrieves an entry based on its index (NOT NECESSARILY sort order) value
func (e *EnumEvent) ByIndex(idx int) *EnumEventItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedEventID is a struct that is designed to replace a *EventID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *EventID it contains while being a better JSON citizen.
type ValidatedEventID struct {
	// id will point to a valid EventID, if possible
	// If id is nil, then ValidatedEventID.Valid() will return false.
	id *EventID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedEventID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedEventID
func (vi *ValidatedEventID) Clone() *ValidatedEventID {
	if vi == nil {
		return nil
	}

	var cid *EventID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedEventID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedEventIds represent the same Event
func (vi *ValidatedEventID) Equals(vj *ValidatedEventID) bool {
	if vi == nil && vj == nil {
		return true
	}

	if vi == nil || vj == nil {
		return false
	}

	if vi.id == nil && vj.id == nil {
		return true
	}

	if vi.id == nil || vj.id == nil {
		return false
	}

	return vi.id.Equals(vj.id)
}

// Valid returns true if and only if the ValidatedEventID corresponds to a recognized Event
func (vi *ValidatedEventID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedEventID) ID() *EventID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedEventID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedEventID) ValidatedID() *ValidatedEventID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedEventID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedEventID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedEventID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedEventID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedEventID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := EventID(capString)
	item := Event.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		if vi.Errors == nil {
			vi.Errors = []error{}
		}
		vi.Errors = append(vi.Errors, err)
		return nil
	}

	vi.id = item.ID.Clone()

	return nil
}

func (vi ValidatedEventID) String() string {
	return vi.ToIDString()
}

type EventIdentifier interface {
	ID() *EventID
	Valid() bool
}
