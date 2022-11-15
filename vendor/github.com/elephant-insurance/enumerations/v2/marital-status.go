package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// MaritalStatusID uniquely identifies a particular MaritalStatus
type MaritalStatusID string

// Clone creates a safe, independent copy of a MaritalStatusID
func (i *MaritalStatusID) Clone() *MaritalStatusID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two MaritalStatusIds are equivalent
func (i *MaritalStatusID) Equals(j *MaritalStatusID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *MaritalStatusID that is either valid or nil
func (i *MaritalStatusID) ID() *MaritalStatusID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *MaritalStatusID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the MaritalStatusID corresponds to a recognized MaritalStatus
func (i *MaritalStatusID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return MaritalStatus.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *MaritalStatusID) ValidatedID() *ValidatedMaritalStatusID {
	if i != nil {
		return &ValidatedMaritalStatusID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *MaritalStatusID) MarshalJSON() ([]byte, error) {
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

func (i *MaritalStatusID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := MaritalStatusID(dataString)
	item := MaritalStatus.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	maritalStatusNeverMarriedID MaritalStatusID = "S"
	maritalStatusMarriedID      MaritalStatusID = "M"
	maritalStatusDivorcedID     MaritalStatusID = "D"
	maritalStatusSeparatedID    MaritalStatusID = "P"
	maritalStatusWidowedID      MaritalStatusID = "W"
	maritalStatusCivilUnionID   MaritalStatusID = "MCU"
)

// EnumMaritalStatusItem describes an entry in an enumeration of MaritalStatus
type EnumMaritalStatusItem struct {
	ID        MaritalStatusID   `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	StateCodes      string
	AlternativeKeys string
}

var (
	maritalStatusNeverMarried = EnumMaritalStatusItem{maritalStatusNeverMarriedID, "Single (Never Married)", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH", "AlternativeKeys": "single,unmarried"}, "NeverMarried", 1, "IN,TN,IL,MD,TX,VA,GA,OH", "single,unmarried"}
	maritalStatusMarried      = EnumMaritalStatusItem{maritalStatusMarriedID, "Married", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH", "AlternativeKeys": "married"}, "Married", 2, "IN,TN,IL,MD,TX,VA,GA,OH", "married"}
	maritalStatusDivorced     = EnumMaritalStatusItem{maritalStatusDivorcedID, "Divorced", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH", "AlternativeKeys": "divorced"}, "Divorced", 3, "IN,TN,IL,MD,TX,VA,GA,OH", "divorced"}
	maritalStatusSeparated    = EnumMaritalStatusItem{maritalStatusSeparatedID, "Separated", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH", "AlternativeKeys": "separated"}, "Separated", 4, "IN,TN,IL,MD,TX,VA,GA,OH", "separated"}
	maritalStatusWidowed      = EnumMaritalStatusItem{maritalStatusWidowedID, "Widowed", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH", "AlternativeKeys": "widowed"}, "Widowed", 5, "IN,TN,IL,MD,TX,VA,GA,OH", "widowed"}
	maritalStatusCivilUnion   = EnumMaritalStatusItem{maritalStatusCivilUnionID, "Civil Union", map[string]string{"StateCodes": "IL", "AlternativeKeys": "civil,union"}, "CivilUnion", 6, "IL", "civil,union"}
)

// EnumMaritalStatus is a collection of MaritalStatus items
type EnumMaritalStatus struct {
	Description string
	Items       []*EnumMaritalStatusItem
	Name        string

	NeverMarried *EnumMaritalStatusItem
	Married      *EnumMaritalStatusItem
	Divorced     *EnumMaritalStatusItem
	Separated    *EnumMaritalStatusItem
	Widowed      *EnumMaritalStatusItem
	CivilUnion   *EnumMaritalStatusItem

	itemDict map[string]*EnumMaritalStatusItem
}

// MaritalStatus is a public singleton instance of EnumMaritalStatus
// representing marital statuses
var MaritalStatus = &EnumMaritalStatus{
	Description: "marital statuses",
	Items: []*EnumMaritalStatusItem{
		&maritalStatusNeverMarried,
		&maritalStatusMarried,
		&maritalStatusDivorced,
		&maritalStatusSeparated,
		&maritalStatusWidowed,
		&maritalStatusCivilUnion,
	},
	Name:         "EnumMaritalStatus",
	NeverMarried: &maritalStatusNeverMarried,
	Married:      &maritalStatusMarried,
	Divorced:     &maritalStatusDivorced,
	Separated:    &maritalStatusSeparated,
	Widowed:      &maritalStatusWidowed,
	CivilUnion:   &maritalStatusCivilUnion,

	itemDict: map[string]*EnumMaritalStatusItem{
		strings.ToLower(string(maritalStatusNeverMarriedID)): &maritalStatusNeverMarried,
		strings.ToLower(string(maritalStatusMarriedID)):      &maritalStatusMarried,
		strings.ToLower(string(maritalStatusDivorcedID)):     &maritalStatusDivorced,
		strings.ToLower(string(maritalStatusSeparatedID)):    &maritalStatusSeparated,
		strings.ToLower(string(maritalStatusWidowedID)):      &maritalStatusWidowed,
		strings.ToLower(string(maritalStatusCivilUnionID)):   &maritalStatusCivilUnion,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumMaritalStatus) ByID(id MaritalStatusIdentifier) *EnumMaritalStatusItem {
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
func (e *EnumMaritalStatus) ByIDString(idx string) *EnumMaritalStatusItem {
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
func (e *EnumMaritalStatus) ByIndex(idx int) *EnumMaritalStatusItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedMaritalStatusID is a struct that is designed to replace a *MaritalStatusID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *MaritalStatusID it contains while being a better JSON citizen.
type ValidatedMaritalStatusID struct {
	// id will point to a valid MaritalStatusID, if possible
	// If id is nil, then ValidatedMaritalStatusID.Valid() will return false.
	id *MaritalStatusID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedMaritalStatusID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedMaritalStatusID
func (vi *ValidatedMaritalStatusID) Clone() *ValidatedMaritalStatusID {
	if vi == nil {
		return nil
	}

	var cid *MaritalStatusID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedMaritalStatusID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedMaritalStatusIds represent the same MaritalStatus
func (vi *ValidatedMaritalStatusID) Equals(vj *ValidatedMaritalStatusID) bool {
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

// Valid returns true if and only if the ValidatedMaritalStatusID corresponds to a recognized MaritalStatus
func (vi *ValidatedMaritalStatusID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedMaritalStatusID) ID() *MaritalStatusID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedMaritalStatusID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedMaritalStatusID) ValidatedID() *ValidatedMaritalStatusID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedMaritalStatusID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedMaritalStatusID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedMaritalStatusID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedMaritalStatusID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedMaritalStatusID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := MaritalStatusID(capString)
	item := MaritalStatus.ByID(&id)
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

func (vi ValidatedMaritalStatusID) String() string {
	return vi.ToIDString()
}

type MaritalStatusIdentifier interface {
	ID() *MaritalStatusID
	Valid() bool
}
