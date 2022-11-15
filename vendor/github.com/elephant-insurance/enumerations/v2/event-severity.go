package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// EventSeverityID uniquely identifies a particular EventSeverity
type EventSeverityID string

// Clone creates a safe, independent copy of a EventSeverityID
func (i *EventSeverityID) Clone() *EventSeverityID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two EventSeverityIds are equivalent
func (i *EventSeverityID) Equals(j *EventSeverityID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *EventSeverityID that is either valid or nil
func (i *EventSeverityID) ID() *EventSeverityID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *EventSeverityID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the EventSeverityID corresponds to a recognized EventSeverity
func (i *EventSeverityID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return EventSeverity.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *EventSeverityID) ValidatedID() *ValidatedEventSeverityID {
	if i != nil {
		return &ValidatedEventSeverityID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *EventSeverityID) MarshalJSON() ([]byte, error) {
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

func (i *EventSeverityID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := EventSeverityID(dataString)
	item := EventSeverity.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	eventSeverityDevelopmentalID EventSeverityID = "dev"
	eventSeverityDiagnosticID    EventSeverityID = "dig"
	eventSeverityInformationalID EventSeverityID = "inf"
	eventSeverityCautionaryID    EventSeverityID = "cau"
	eventSeverityErrorID         EventSeverityID = "err"
	eventSeverityCriticalID      EventSeverityID = "cri"
)

// EnumEventSeverityItem describes an entry in an enumeration of EventSeverity
type EnumEventSeverityItem struct {
	ID        EventSeverityID   `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	eventSeverityDevelopmental = EnumEventSeverityItem{eventSeverityDevelopmentalID, "this event is under development and should be ignored by production processes", nil, "Developmental", 1}
	eventSeverityDiagnostic    = EnumEventSeverityItem{eventSeverityDiagnosticID, "this event is triggered by the app itself to record health statistics for tracking", nil, "Diagnostic", 2}
	eventSeverityInformational = EnumEventSeverityItem{eventSeverityInformationalID, "this event is not necessarily a problem, just something worth tracking", nil, "Informational", 3}
	eventSeverityCautionary    = EnumEventSeverityItem{eventSeverityCautionaryID, "this event appears to indicate a potential problem", nil, "Cautionary", 4}
	eventSeverityError         = EnumEventSeverityItem{eventSeverityErrorID, "this event definitely should never happen", nil, "Error", 5}
	eventSeverityCritical      = EnumEventSeverityItem{eventSeverityCriticalID, "this event requires immediate attention", nil, "Critical", 6}
)

// EnumEventSeverity is a collection of EventSeverity items
type EnumEventSeverity struct {
	Description string
	Items       []*EnumEventSeverityItem
	Name        string

	Developmental *EnumEventSeverityItem
	Diagnostic    *EnumEventSeverityItem
	Informational *EnumEventSeverityItem
	Cautionary    *EnumEventSeverityItem
	Error         *EnumEventSeverityItem
	Critical      *EnumEventSeverityItem

	itemDict map[string]*EnumEventSeverityItem
}

// EventSeverity is a public singleton instance of EnumEventSeverity
// representing severity of microservice application events
var EventSeverity = &EnumEventSeverity{
	Description: "severity of microservice application events",
	Items: []*EnumEventSeverityItem{
		&eventSeverityDevelopmental,
		&eventSeverityDiagnostic,
		&eventSeverityInformational,
		&eventSeverityCautionary,
		&eventSeverityError,
		&eventSeverityCritical,
	},
	Name:          "EnumEventSeverity",
	Developmental: &eventSeverityDevelopmental,
	Diagnostic:    &eventSeverityDiagnostic,
	Informational: &eventSeverityInformational,
	Cautionary:    &eventSeverityCautionary,
	Error:         &eventSeverityError,
	Critical:      &eventSeverityCritical,

	itemDict: map[string]*EnumEventSeverityItem{
		strings.ToLower(string(eventSeverityDevelopmentalID)): &eventSeverityDevelopmental,
		strings.ToLower(string(eventSeverityDiagnosticID)):    &eventSeverityDiagnostic,
		strings.ToLower(string(eventSeverityInformationalID)): &eventSeverityInformational,
		strings.ToLower(string(eventSeverityCautionaryID)):    &eventSeverityCautionary,
		strings.ToLower(string(eventSeverityErrorID)):         &eventSeverityError,
		strings.ToLower(string(eventSeverityCriticalID)):      &eventSeverityCritical,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumEventSeverity) ByID(id EventSeverityIdentifier) *EnumEventSeverityItem {
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
func (e *EnumEventSeverity) ByIDString(idx string) *EnumEventSeverityItem {
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
func (e *EnumEventSeverity) ByIndex(idx int) *EnumEventSeverityItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedEventSeverityID is a struct that is designed to replace a *EventSeverityID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *EventSeverityID it contains while being a better JSON citizen.
type ValidatedEventSeverityID struct {
	// id will point to a valid EventSeverityID, if possible
	// If id is nil, then ValidatedEventSeverityID.Valid() will return false.
	id *EventSeverityID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedEventSeverityID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedEventSeverityID
func (vi *ValidatedEventSeverityID) Clone() *ValidatedEventSeverityID {
	if vi == nil {
		return nil
	}

	var cid *EventSeverityID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedEventSeverityID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedEventSeverityIds represent the same EventSeverity
func (vi *ValidatedEventSeverityID) Equals(vj *ValidatedEventSeverityID) bool {
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

// Valid returns true if and only if the ValidatedEventSeverityID corresponds to a recognized EventSeverity
func (vi *ValidatedEventSeverityID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedEventSeverityID) ID() *EventSeverityID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedEventSeverityID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedEventSeverityID) ValidatedID() *ValidatedEventSeverityID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedEventSeverityID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedEventSeverityID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedEventSeverityID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedEventSeverityID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedEventSeverityID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := EventSeverityID(capString)
	item := EventSeverity.ByID(&id)
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

func (vi ValidatedEventSeverityID) String() string {
	return vi.ToIDString()
}

type EventSeverityIdentifier interface {
	ID() *EventSeverityID
	Valid() bool
}
