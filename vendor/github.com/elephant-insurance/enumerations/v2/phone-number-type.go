package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// PhoneNumberTypeID uniquely identifies a particular PhoneNumberType
type PhoneNumberTypeID string

// Clone creates a safe, independent copy of a PhoneNumberTypeID
func (i *PhoneNumberTypeID) Clone() *PhoneNumberTypeID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two PhoneNumberTypeIds are equivalent
func (i *PhoneNumberTypeID) Equals(j *PhoneNumberTypeID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *PhoneNumberTypeID that is either valid or nil
func (i *PhoneNumberTypeID) ID() *PhoneNumberTypeID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *PhoneNumberTypeID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the PhoneNumberTypeID corresponds to a recognized PhoneNumberType
func (i *PhoneNumberTypeID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return PhoneNumberType.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *PhoneNumberTypeID) ValidatedID() *ValidatedPhoneNumberTypeID {
	if i != nil {
		return &ValidatedPhoneNumberTypeID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *PhoneNumberTypeID) MarshalJSON() ([]byte, error) {
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

func (i *PhoneNumberTypeID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := PhoneNumberTypeID(dataString)
	item := PhoneNumberType.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	phoneNumberTypeHomeID   PhoneNumberTypeID = "home"
	phoneNumberTypeWorkID   PhoneNumberTypeID = "work"
	phoneNumberTypeMobileID PhoneNumberTypeID = "mobile"
)

// EnumPhoneNumberTypeItem describes an entry in an enumeration of PhoneNumberType
type EnumPhoneNumberTypeItem struct {
	ID        PhoneNumberTypeID `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	phoneNumberTypeHome   = EnumPhoneNumberTypeItem{phoneNumberTypeHomeID, "Home", nil, "Home", 1}
	phoneNumberTypeWork   = EnumPhoneNumberTypeItem{phoneNumberTypeWorkID, "Work", nil, "Work", 2}
	phoneNumberTypeMobile = EnumPhoneNumberTypeItem{phoneNumberTypeMobileID, "Mobile", nil, "Mobile", 3}
)

// EnumPhoneNumberType is a collection of PhoneNumberType items
type EnumPhoneNumberType struct {
	Description string
	Items       []*EnumPhoneNumberTypeItem
	Name        string

	Home   *EnumPhoneNumberTypeItem
	Work   *EnumPhoneNumberTypeItem
	Mobile *EnumPhoneNumberTypeItem

	itemDict map[string]*EnumPhoneNumberTypeItem
}

// PhoneNumberType is a public singleton instance of EnumPhoneNumberType
// representing types of phone number
var PhoneNumberType = &EnumPhoneNumberType{
	Description: "types of phone number",
	Items: []*EnumPhoneNumberTypeItem{
		&phoneNumberTypeHome,
		&phoneNumberTypeWork,
		&phoneNumberTypeMobile,
	},
	Name:   "EnumPhoneNumberType",
	Home:   &phoneNumberTypeHome,
	Work:   &phoneNumberTypeWork,
	Mobile: &phoneNumberTypeMobile,

	itemDict: map[string]*EnumPhoneNumberTypeItem{
		strings.ToLower(string(phoneNumberTypeHomeID)):   &phoneNumberTypeHome,
		strings.ToLower(string(phoneNumberTypeWorkID)):   &phoneNumberTypeWork,
		strings.ToLower(string(phoneNumberTypeMobileID)): &phoneNumberTypeMobile,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumPhoneNumberType) ByID(id PhoneNumberTypeIdentifier) *EnumPhoneNumberTypeItem {
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
func (e *EnumPhoneNumberType) ByIDString(idx string) *EnumPhoneNumberTypeItem {
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
func (e *EnumPhoneNumberType) ByIndex(idx int) *EnumPhoneNumberTypeItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedPhoneNumberTypeID is a struct that is designed to replace a *PhoneNumberTypeID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *PhoneNumberTypeID it contains while being a better JSON citizen.
type ValidatedPhoneNumberTypeID struct {
	// id will point to a valid PhoneNumberTypeID, if possible
	// If id is nil, then ValidatedPhoneNumberTypeID.Valid() will return false.
	id *PhoneNumberTypeID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedPhoneNumberTypeID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedPhoneNumberTypeID
func (vi *ValidatedPhoneNumberTypeID) Clone() *ValidatedPhoneNumberTypeID {
	if vi == nil {
		return nil
	}

	var cid *PhoneNumberTypeID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedPhoneNumberTypeID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedPhoneNumberTypeIds represent the same PhoneNumberType
func (vi *ValidatedPhoneNumberTypeID) Equals(vj *ValidatedPhoneNumberTypeID) bool {
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

// Valid returns true if and only if the ValidatedPhoneNumberTypeID corresponds to a recognized PhoneNumberType
func (vi *ValidatedPhoneNumberTypeID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedPhoneNumberTypeID) ID() *PhoneNumberTypeID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedPhoneNumberTypeID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedPhoneNumberTypeID) ValidatedID() *ValidatedPhoneNumberTypeID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedPhoneNumberTypeID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedPhoneNumberTypeID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedPhoneNumberTypeID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedPhoneNumberTypeID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedPhoneNumberTypeID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := PhoneNumberTypeID(capString)
	item := PhoneNumberType.ByID(&id)
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

func (vi ValidatedPhoneNumberTypeID) String() string {
	return vi.ToIDString()
}

type PhoneNumberTypeIdentifier interface {
	ID() *PhoneNumberTypeID
	Valid() bool
}
