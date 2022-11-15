package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// LocationReasonID uniquely identifies a particular LocationReason
type LocationReasonID string

// Clone creates a safe, independent copy of a LocationReasonID
func (i *LocationReasonID) Clone() *LocationReasonID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two LocationReasonIds are equivalent
func (i *LocationReasonID) Equals(j *LocationReasonID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *LocationReasonID that is either valid or nil
func (i *LocationReasonID) ID() *LocationReasonID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *LocationReasonID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the LocationReasonID corresponds to a recognized LocationReason
func (i *LocationReasonID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return LocationReason.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *LocationReasonID) ValidatedID() *ValidatedLocationReasonID {
	if i != nil {
		return &ValidatedLocationReasonID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *LocationReasonID) MarshalJSON() ([]byte, error) {
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

func (i *LocationReasonID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := LocationReasonID(dataString)
	item := LocationReason.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	locationReasonIncarceratedID LocationReasonID = "incarcerated"
	locationReasonOutOfCountryID LocationReasonID = "outofcountry"
	locationReasonOutOfStateID   LocationReasonID = "livesOutofState"
)

// EnumLocationReasonItem describes an entry in an enumeration of LocationReason
type EnumLocationReasonItem struct {
	ID        LocationReasonID  `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	locationReasonIncarcerated = EnumLocationReasonItem{locationReasonIncarceratedID, "Incarcerated", nil, "Incarcerated", 1}
	locationReasonOutOfCountry = EnumLocationReasonItem{locationReasonOutOfCountryID, "Out of country", nil, "OutOfCountry", 2}
	locationReasonOutOfState   = EnumLocationReasonItem{locationReasonOutOfStateID, "Lives out of state", nil, "OutOfState", 3}
)

// EnumLocationReason is a collection of LocationReason items
type EnumLocationReason struct {
	Description string
	Items       []*EnumLocationReasonItem
	Name        string

	Incarcerated *EnumLocationReasonItem
	OutOfCountry *EnumLocationReasonItem
	OutOfState   *EnumLocationReasonItem

	itemDict map[string]*EnumLocationReasonItem
}

// LocationReason is a public singleton instance of EnumLocationReason
// representing reasons for bing at a different location
var LocationReason = &EnumLocationReason{
	Description: "reasons for bing at a different location",
	Items: []*EnumLocationReasonItem{
		&locationReasonIncarcerated,
		&locationReasonOutOfCountry,
		&locationReasonOutOfState,
	},
	Name:         "EnumLocationReason",
	Incarcerated: &locationReasonIncarcerated,
	OutOfCountry: &locationReasonOutOfCountry,
	OutOfState:   &locationReasonOutOfState,

	itemDict: map[string]*EnumLocationReasonItem{
		strings.ToLower(string(locationReasonIncarceratedID)): &locationReasonIncarcerated,
		strings.ToLower(string(locationReasonOutOfCountryID)): &locationReasonOutOfCountry,
		strings.ToLower(string(locationReasonOutOfStateID)):   &locationReasonOutOfState,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumLocationReason) ByID(id LocationReasonIdentifier) *EnumLocationReasonItem {
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
func (e *EnumLocationReason) ByIDString(idx string) *EnumLocationReasonItem {
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
func (e *EnumLocationReason) ByIndex(idx int) *EnumLocationReasonItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedLocationReasonID is a struct that is designed to replace a *LocationReasonID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *LocationReasonID it contains while being a better JSON citizen.
type ValidatedLocationReasonID struct {
	// id will point to a valid LocationReasonID, if possible
	// If id is nil, then ValidatedLocationReasonID.Valid() will return false.
	id *LocationReasonID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedLocationReasonID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedLocationReasonID
func (vi *ValidatedLocationReasonID) Clone() *ValidatedLocationReasonID {
	if vi == nil {
		return nil
	}

	var cid *LocationReasonID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedLocationReasonID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedLocationReasonIds represent the same LocationReason
func (vi *ValidatedLocationReasonID) Equals(vj *ValidatedLocationReasonID) bool {
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

// Valid returns true if and only if the ValidatedLocationReasonID corresponds to a recognized LocationReason
func (vi *ValidatedLocationReasonID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedLocationReasonID) ID() *LocationReasonID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedLocationReasonID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedLocationReasonID) ValidatedID() *ValidatedLocationReasonID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedLocationReasonID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedLocationReasonID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedLocationReasonID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedLocationReasonID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedLocationReasonID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := LocationReasonID(capString)
	item := LocationReason.ByID(&id)
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

func (vi ValidatedLocationReasonID) String() string {
	return vi.ToIDString()
}

type LocationReasonIdentifier interface {
	ID() *LocationReasonID
	Valid() bool
}
