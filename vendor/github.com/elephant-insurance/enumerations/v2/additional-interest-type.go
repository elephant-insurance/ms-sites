package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// AdditionalInterestTypeID uniquely identifies a particular AdditionalInterestType
type AdditionalInterestTypeID string

// Clone creates a safe, independent copy of a AdditionalInterestTypeID
func (i *AdditionalInterestTypeID) Clone() *AdditionalInterestTypeID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two AdditionalInterestTypeIds are equivalent
func (i *AdditionalInterestTypeID) Equals(j *AdditionalInterestTypeID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *AdditionalInterestTypeID that is either valid or nil
func (i *AdditionalInterestTypeID) ID() *AdditionalInterestTypeID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *AdditionalInterestTypeID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the AdditionalInterestTypeID corresponds to a recognized AdditionalInterestType
func (i *AdditionalInterestTypeID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return AdditionalInterestType.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *AdditionalInterestTypeID) ValidatedID() *ValidatedAdditionalInterestTypeID {
	if i != nil {
		return &ValidatedAdditionalInterestTypeID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *AdditionalInterestTypeID) MarshalJSON() ([]byte, error) {
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

func (i *AdditionalInterestTypeID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := AdditionalInterestTypeID(dataString)
	item := AdditionalInterestType.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	additionalInterestTypeLessorID     AdditionalInterestTypeID = "LESSOR"
	additionalInterestTypeLienholderID AdditionalInterestTypeID = "LIEN"
	additionalInterestTypeLOSSPID      AdditionalInterestTypeID = "LOSSP"
)

// EnumAdditionalInterestTypeItem describes an entry in an enumeration of AdditionalInterestType
type EnumAdditionalInterestTypeItem struct {
	ID        AdditionalInterestTypeID `json:"Value"`
	Desc      string                   `json:"Description,omitempty"`
	Meta      map[string]string        `json:",omitempty"`
	Name      string                   `json:"Name"`
	SortOrder int
}

var (
	additionalInterestTypeLessor     = EnumAdditionalInterestTypeItem{additionalInterestTypeLessorID, "Lessor", nil, "Lessor", 1}
	additionalInterestTypeLienholder = EnumAdditionalInterestTypeItem{additionalInterestTypeLienholderID, "Lienholder", nil, "Lienholder", 2}
	additionalInterestTypeLOSSP      = EnumAdditionalInterestTypeItem{additionalInterestTypeLOSSPID, "LOSSP", nil, "LOSSP", 3}
)

// EnumAdditionalInterestType is a collection of AdditionalInterestType items
type EnumAdditionalInterestType struct {
	Description string
	Items       []*EnumAdditionalInterestTypeItem
	Name        string

	Lessor     *EnumAdditionalInterestTypeItem
	Lienholder *EnumAdditionalInterestTypeItem
	LOSSP      *EnumAdditionalInterestTypeItem

	itemDict map[string]*EnumAdditionalInterestTypeItem
}

// AdditionalInterestType is a public singleton instance of EnumAdditionalInterestType
// representing types of legal financial interest in a vehicle
var AdditionalInterestType = &EnumAdditionalInterestType{
	Description: "types of legal financial interest in a vehicle",
	Items: []*EnumAdditionalInterestTypeItem{
		&additionalInterestTypeLessor,
		&additionalInterestTypeLienholder,
		&additionalInterestTypeLOSSP,
	},
	Name:       "EnumAdditionalInterestType",
	Lessor:     &additionalInterestTypeLessor,
	Lienholder: &additionalInterestTypeLienholder,
	LOSSP:      &additionalInterestTypeLOSSP,

	itemDict: map[string]*EnumAdditionalInterestTypeItem{
		strings.ToLower(string(additionalInterestTypeLessorID)):     &additionalInterestTypeLessor,
		strings.ToLower(string(additionalInterestTypeLienholderID)): &additionalInterestTypeLienholder,
		strings.ToLower(string(additionalInterestTypeLOSSPID)):      &additionalInterestTypeLOSSP,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumAdditionalInterestType) ByID(id AdditionalInterestTypeIdentifier) *EnumAdditionalInterestTypeItem {
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
func (e *EnumAdditionalInterestType) ByIDString(idx string) *EnumAdditionalInterestTypeItem {
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
func (e *EnumAdditionalInterestType) ByIndex(idx int) *EnumAdditionalInterestTypeItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedAdditionalInterestTypeID is a struct that is designed to replace a *AdditionalInterestTypeID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *AdditionalInterestTypeID it contains while being a better JSON citizen.
type ValidatedAdditionalInterestTypeID struct {
	// id will point to a valid AdditionalInterestTypeID, if possible
	// If id is nil, then ValidatedAdditionalInterestTypeID.Valid() will return false.
	id *AdditionalInterestTypeID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedAdditionalInterestTypeID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedAdditionalInterestTypeID
func (vi *ValidatedAdditionalInterestTypeID) Clone() *ValidatedAdditionalInterestTypeID {
	if vi == nil {
		return nil
	}

	var cid *AdditionalInterestTypeID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedAdditionalInterestTypeID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedAdditionalInterestTypeIds represent the same AdditionalInterestType
func (vi *ValidatedAdditionalInterestTypeID) Equals(vj *ValidatedAdditionalInterestTypeID) bool {
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

// Valid returns true if and only if the ValidatedAdditionalInterestTypeID corresponds to a recognized AdditionalInterestType
func (vi *ValidatedAdditionalInterestTypeID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedAdditionalInterestTypeID) ID() *AdditionalInterestTypeID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedAdditionalInterestTypeID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedAdditionalInterestTypeID) ValidatedID() *ValidatedAdditionalInterestTypeID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedAdditionalInterestTypeID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedAdditionalInterestTypeID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedAdditionalInterestTypeID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedAdditionalInterestTypeID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedAdditionalInterestTypeID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := AdditionalInterestTypeID(capString)
	item := AdditionalInterestType.ByID(&id)
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

func (vi ValidatedAdditionalInterestTypeID) String() string {
	return vi.ToIDString()
}

type AdditionalInterestTypeIdentifier interface {
	ID() *AdditionalInterestTypeID
	Valid() bool
}
