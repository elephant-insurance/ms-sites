package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// AddressStatusID uniquely identifies a particular AddressStatus
type AddressStatusID string

// Clone creates a safe, independent copy of a AddressStatusID
func (i *AddressStatusID) Clone() *AddressStatusID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two AddressStatusIds are equivalent
func (i *AddressStatusID) Equals(j *AddressStatusID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *AddressStatusID that is either valid or nil
func (i *AddressStatusID) ID() *AddressStatusID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *AddressStatusID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the AddressStatusID corresponds to a recognized AddressStatus
func (i *AddressStatusID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return AddressStatus.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *AddressStatusID) ValidatedID() *ValidatedAddressStatusID {
	if i != nil {
		return &ValidatedAddressStatusID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *AddressStatusID) MarshalJSON() ([]byte, error) {
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

func (i *AddressStatusID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := AddressStatusID(dataString)
	item := AddressStatus.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	addressStatusCurrentID        AddressStatusID = "current"
	addressStatusPreviousID       AddressStatusID = "previous"
	addressStatusSecondPreviousID AddressStatusID = "secondPrevious"
	addressStatusFutureID         AddressStatusID = "future"
	addressStatusMailingID        AddressStatusID = "mailing"
)

// EnumAddressStatusItem describes an entry in an enumeration of AddressStatus
type EnumAddressStatusItem struct {
	ID        AddressStatusID   `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	addressStatusCurrent        = EnumAddressStatusItem{addressStatusCurrentID, "Current", nil, "Current", 1}
	addressStatusPrevious       = EnumAddressStatusItem{addressStatusPreviousID, "Previous", nil, "Previous", 2}
	addressStatusSecondPrevious = EnumAddressStatusItem{addressStatusSecondPreviousID, "SecondPrevious", nil, "SecondPrevious", 3}
	addressStatusFuture         = EnumAddressStatusItem{addressStatusFutureID, "Future", nil, "Future", 4}
	addressStatusMailing        = EnumAddressStatusItem{addressStatusMailingID, "Mailing", nil, "Mailing", 5}
)

// EnumAddressStatus is a collection of AddressStatus items
type EnumAddressStatus struct {
	Description string
	Items       []*EnumAddressStatusItem
	Name        string

	Current        *EnumAddressStatusItem
	Previous       *EnumAddressStatusItem
	SecondPrevious *EnumAddressStatusItem
	Future         *EnumAddressStatusItem
	Mailing        *EnumAddressStatusItem

	itemDict map[string]*EnumAddressStatusItem
}

// AddressStatus is a public singleton instance of EnumAddressStatus
// representing address status
var AddressStatus = &EnumAddressStatus{
	Description: "address status",
	Items: []*EnumAddressStatusItem{
		&addressStatusCurrent,
		&addressStatusPrevious,
		&addressStatusSecondPrevious,
		&addressStatusFuture,
		&addressStatusMailing,
	},
	Name:           "EnumAddressStatus",
	Current:        &addressStatusCurrent,
	Previous:       &addressStatusPrevious,
	SecondPrevious: &addressStatusSecondPrevious,
	Future:         &addressStatusFuture,
	Mailing:        &addressStatusMailing,

	itemDict: map[string]*EnumAddressStatusItem{
		strings.ToLower(string(addressStatusCurrentID)):        &addressStatusCurrent,
		strings.ToLower(string(addressStatusPreviousID)):       &addressStatusPrevious,
		strings.ToLower(string(addressStatusSecondPreviousID)): &addressStatusSecondPrevious,
		strings.ToLower(string(addressStatusFutureID)):         &addressStatusFuture,
		strings.ToLower(string(addressStatusMailingID)):        &addressStatusMailing,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumAddressStatus) ByID(id AddressStatusIdentifier) *EnumAddressStatusItem {
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
func (e *EnumAddressStatus) ByIDString(idx string) *EnumAddressStatusItem {
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
func (e *EnumAddressStatus) ByIndex(idx int) *EnumAddressStatusItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedAddressStatusID is a struct that is designed to replace a *AddressStatusID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *AddressStatusID it contains while being a better JSON citizen.
type ValidatedAddressStatusID struct {
	// id will point to a valid AddressStatusID, if possible
	// If id is nil, then ValidatedAddressStatusID.Valid() will return false.
	id *AddressStatusID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedAddressStatusID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedAddressStatusID
func (vi *ValidatedAddressStatusID) Clone() *ValidatedAddressStatusID {
	if vi == nil {
		return nil
	}

	var cid *AddressStatusID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedAddressStatusID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedAddressStatusIds represent the same AddressStatus
func (vi *ValidatedAddressStatusID) Equals(vj *ValidatedAddressStatusID) bool {
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

// Valid returns true if and only if the ValidatedAddressStatusID corresponds to a recognized AddressStatus
func (vi *ValidatedAddressStatusID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedAddressStatusID) ID() *AddressStatusID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedAddressStatusID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedAddressStatusID) ValidatedID() *ValidatedAddressStatusID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedAddressStatusID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedAddressStatusID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedAddressStatusID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedAddressStatusID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedAddressStatusID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := AddressStatusID(capString)
	item := AddressStatus.ByID(&id)
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

func (vi ValidatedAddressStatusID) String() string {
	return vi.ToIDString()
}

type AddressStatusIdentifier interface {
	ID() *AddressStatusID
	Valid() bool
}
