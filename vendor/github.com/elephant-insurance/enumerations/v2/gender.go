package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// GenderID uniquely identifies a particular Gender
type GenderID string

// Clone creates a safe, independent copy of a GenderID
func (i *GenderID) Clone() *GenderID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two GenderIds are equivalent
func (i *GenderID) Equals(j *GenderID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *GenderID that is either valid or nil
func (i *GenderID) ID() *GenderID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *GenderID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the GenderID corresponds to a recognized Gender
func (i *GenderID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return Gender.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *GenderID) ValidatedID() *ValidatedGenderID {
	if i != nil {
		return &ValidatedGenderID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *GenderID) MarshalJSON() ([]byte, error) {
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

func (i *GenderID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := GenderID(dataString)
	item := Gender.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	genderMaleID   GenderID = "M"
	genderFemaleID GenderID = "F"
)

// EnumGenderItem describes an entry in an enumeration of Gender
type EnumGenderItem struct {
	ID        GenderID          `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	genderMale   = EnumGenderItem{genderMaleID, "Male", nil, "Male", 1}
	genderFemale = EnumGenderItem{genderFemaleID, "Female", nil, "Female", 2}
)

// EnumGender is a collection of Gender items
type EnumGender struct {
	Description string
	Items       []*EnumGenderItem
	Name        string

	Male   *EnumGenderItem
	Female *EnumGenderItem

	itemDict map[string]*EnumGenderItem
}

// Gender is a public singleton instance of EnumGender
// representing genders
var Gender = &EnumGender{
	Description: "genders",
	Items: []*EnumGenderItem{
		&genderMale,
		&genderFemale,
	},
	Name:   "EnumGender",
	Male:   &genderMale,
	Female: &genderFemale,

	itemDict: map[string]*EnumGenderItem{
		strings.ToLower(string(genderMaleID)):   &genderMale,
		strings.ToLower(string(genderFemaleID)): &genderFemale,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumGender) ByID(id GenderIdentifier) *EnumGenderItem {
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
func (e *EnumGender) ByIDString(idx string) *EnumGenderItem {
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
func (e *EnumGender) ByIndex(idx int) *EnumGenderItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedGenderID is a struct that is designed to replace a *GenderID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *GenderID it contains while being a better JSON citizen.
type ValidatedGenderID struct {
	// id will point to a valid GenderID, if possible
	// If id is nil, then ValidatedGenderID.Valid() will return false.
	id *GenderID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedGenderID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedGenderID
func (vi *ValidatedGenderID) Clone() *ValidatedGenderID {
	if vi == nil {
		return nil
	}

	var cid *GenderID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedGenderID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedGenderIds represent the same Gender
func (vi *ValidatedGenderID) Equals(vj *ValidatedGenderID) bool {
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

// Valid returns true if and only if the ValidatedGenderID corresponds to a recognized Gender
func (vi *ValidatedGenderID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedGenderID) ID() *GenderID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedGenderID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedGenderID) ValidatedID() *ValidatedGenderID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedGenderID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedGenderID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedGenderID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedGenderID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedGenderID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := GenderID(capString)
	item := Gender.ByID(&id)
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

func (vi ValidatedGenderID) String() string {
	return vi.ToIDString()
}

type GenderIdentifier interface {
	ID() *GenderID
	Valid() bool
}
