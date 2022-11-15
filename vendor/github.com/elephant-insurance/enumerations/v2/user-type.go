package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// UserTypeID uniquely identifies a particular UserType
type UserTypeID string

// Clone creates a safe, independent copy of a UserTypeID
func (i *UserTypeID) Clone() *UserTypeID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two UserTypeIds are equivalent
func (i *UserTypeID) Equals(j *UserTypeID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *UserTypeID that is either valid or nil
func (i *UserTypeID) ID() *UserTypeID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *UserTypeID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the UserTypeID corresponds to a recognized UserType
func (i *UserTypeID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return UserType.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *UserTypeID) ValidatedID() *ValidatedUserTypeID {
	if i != nil {
		return &ValidatedUserTypeID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *UserTypeID) MarshalJSON() ([]byte, error) {
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

func (i *UserTypeID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := UserTypeID(dataString)
	item := UserType.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	userTypeUnknownID  UserTypeID = "Unknown"
	userTypeCustomerID UserTypeID = "Customer"
	userTypeAgentID    UserTypeID = "Agent"
)

// EnumUserTypeItem describes an entry in an enumeration of UserType
type EnumUserTypeItem struct {
	ID        UserTypeID        `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	userTypeUnknown  = EnumUserTypeItem{userTypeUnknownID, "Unknown", nil, "Unknown", 1}
	userTypeCustomer = EnumUserTypeItem{userTypeCustomerID, "Customer", nil, "Customer", 2}
	userTypeAgent    = EnumUserTypeItem{userTypeAgentID, "Agent", nil, "Agent", 3}
)

// EnumUserType is a collection of UserType items
type EnumUserType struct {
	Description string
	Items       []*EnumUserTypeItem
	Name        string

	Unknown  *EnumUserTypeItem
	Customer *EnumUserTypeItem
	Agent    *EnumUserTypeItem

	itemDict map[string]*EnumUserTypeItem
}

// UserType is a public singleton instance of EnumUserType
// representing types of quote service users
var UserType = &EnumUserType{
	Description: "types of quote service users",
	Items: []*EnumUserTypeItem{
		&userTypeUnknown,
		&userTypeCustomer,
		&userTypeAgent,
	},
	Name:     "EnumUserType",
	Unknown:  &userTypeUnknown,
	Customer: &userTypeCustomer,
	Agent:    &userTypeAgent,

	itemDict: map[string]*EnumUserTypeItem{
		strings.ToLower(string(userTypeUnknownID)):  &userTypeUnknown,
		strings.ToLower(string(userTypeCustomerID)): &userTypeCustomer,
		strings.ToLower(string(userTypeAgentID)):    &userTypeAgent,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumUserType) ByID(id UserTypeIdentifier) *EnumUserTypeItem {
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
func (e *EnumUserType) ByIDString(idx string) *EnumUserTypeItem {
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
func (e *EnumUserType) ByIndex(idx int) *EnumUserTypeItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedUserTypeID is a struct that is designed to replace a *UserTypeID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *UserTypeID it contains while being a better JSON citizen.
type ValidatedUserTypeID struct {
	// id will point to a valid UserTypeID, if possible
	// If id is nil, then ValidatedUserTypeID.Valid() will return false.
	id *UserTypeID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedUserTypeID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedUserTypeID
func (vi *ValidatedUserTypeID) Clone() *ValidatedUserTypeID {
	if vi == nil {
		return nil
	}

	var cid *UserTypeID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedUserTypeID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedUserTypeIds represent the same UserType
func (vi *ValidatedUserTypeID) Equals(vj *ValidatedUserTypeID) bool {
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

// Valid returns true if and only if the ValidatedUserTypeID corresponds to a recognized UserType
func (vi *ValidatedUserTypeID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedUserTypeID) ID() *UserTypeID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedUserTypeID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedUserTypeID) ValidatedID() *ValidatedUserTypeID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedUserTypeID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedUserTypeID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedUserTypeID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedUserTypeID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedUserTypeID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := UserTypeID(capString)
	item := UserType.ByID(&id)
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

func (vi ValidatedUserTypeID) String() string {
	return vi.ToIDString()
}

type UserTypeIdentifier interface {
	ID() *UserTypeID
	Valid() bool
}
