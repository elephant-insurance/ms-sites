package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// YearsWithID uniquely identifies a particular YearsWith
type YearsWithID string

// Clone creates a safe, independent copy of a YearsWithID
func (i *YearsWithID) Clone() *YearsWithID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two YearsWithIds are equivalent
func (i *YearsWithID) Equals(j *YearsWithID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *YearsWithID that is either valid or nil
func (i *YearsWithID) ID() *YearsWithID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *YearsWithID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the YearsWithID corresponds to a recognized YearsWith
func (i *YearsWithID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return YearsWith.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *YearsWithID) ValidatedID() *ValidatedYearsWithID {
	if i != nil {
		return &ValidatedYearsWithID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *YearsWithID) MarshalJSON() ([]byte, error) {
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

func (i *YearsWithID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := YearsWithID(dataString)
	item := YearsWith.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	yearsWithLessThan1ID YearsWithID = "less_1year"
	yearsWithOneID       YearsWithID = "1year"
	yearsWithTwoID       YearsWithID = "2years"
	yearsWithThreeID     YearsWithID = "3years"
	yearsWithFourID      YearsWithID = "4years"
	yearsWithFivePlusID  YearsWithID = "5plus_years"
)

// EnumYearsWithItem describes an entry in an enumeration of YearsWith
type EnumYearsWithItem struct {
	ID        YearsWithID       `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	AlternativeKeys string
}

var (
	yearsWithLessThan1 = EnumYearsWithItem{yearsWithLessThan1ID, "Less Than 1 Year", map[string]string{"AlternativeKeys": "0"}, "LessThan1", 1, "0"}
	yearsWithOne       = EnumYearsWithItem{yearsWithOneID, "1 Year", map[string]string{"AlternativeKeys": "1"}, "One", 2, "1"}
	yearsWithTwo       = EnumYearsWithItem{yearsWithTwoID, "2 Years", map[string]string{"AlternativeKeys": "2"}, "Two", 3, "2"}
	yearsWithThree     = EnumYearsWithItem{yearsWithThreeID, "3 Years", map[string]string{"AlternativeKeys": "3"}, "Three", 4, "3"}
	yearsWithFour      = EnumYearsWithItem{yearsWithFourID, "4 Years", map[string]string{"AlternativeKeys": "4"}, "Four", 5, "4"}
	yearsWithFivePlus  = EnumYearsWithItem{yearsWithFivePlusID, "5 or More Years", map[string]string{"AlternativeKeys": "5"}, "FivePlus", 6, "5"}
)

// EnumYearsWith is a collection of YearsWith items
type EnumYearsWith struct {
	Description string
	Items       []*EnumYearsWithItem
	Name        string

	LessThan1 *EnumYearsWithItem
	One       *EnumYearsWithItem
	Two       *EnumYearsWithItem
	Three     *EnumYearsWithItem
	Four      *EnumYearsWithItem
	FivePlus  *EnumYearsWithItem

	itemDict map[string]*EnumYearsWithItem
}

// YearsWith is a public singleton instance of EnumYearsWith
// representing ranges of years with current carrier
var YearsWith = &EnumYearsWith{
	Description: "ranges of years with current carrier",
	Items: []*EnumYearsWithItem{
		&yearsWithLessThan1,
		&yearsWithOne,
		&yearsWithTwo,
		&yearsWithThree,
		&yearsWithFour,
		&yearsWithFivePlus,
	},
	Name:      "EnumYearsWith",
	LessThan1: &yearsWithLessThan1,
	One:       &yearsWithOne,
	Two:       &yearsWithTwo,
	Three:     &yearsWithThree,
	Four:      &yearsWithFour,
	FivePlus:  &yearsWithFivePlus,

	itemDict: map[string]*EnumYearsWithItem{
		strings.ToLower(string(yearsWithLessThan1ID)): &yearsWithLessThan1,
		strings.ToLower(string(yearsWithOneID)):       &yearsWithOne,
		strings.ToLower(string(yearsWithTwoID)):       &yearsWithTwo,
		strings.ToLower(string(yearsWithThreeID)):     &yearsWithThree,
		strings.ToLower(string(yearsWithFourID)):      &yearsWithFour,
		strings.ToLower(string(yearsWithFivePlusID)):  &yearsWithFivePlus,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumYearsWith) ByID(id YearsWithIdentifier) *EnumYearsWithItem {
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
func (e *EnumYearsWith) ByIDString(idx string) *EnumYearsWithItem {
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
func (e *EnumYearsWith) ByIndex(idx int) *EnumYearsWithItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedYearsWithID is a struct that is designed to replace a *YearsWithID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *YearsWithID it contains while being a better JSON citizen.
type ValidatedYearsWithID struct {
	// id will point to a valid YearsWithID, if possible
	// If id is nil, then ValidatedYearsWithID.Valid() will return false.
	id *YearsWithID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedYearsWithID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedYearsWithID
func (vi *ValidatedYearsWithID) Clone() *ValidatedYearsWithID {
	if vi == nil {
		return nil
	}

	var cid *YearsWithID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedYearsWithID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedYearsWithIds represent the same YearsWith
func (vi *ValidatedYearsWithID) Equals(vj *ValidatedYearsWithID) bool {
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

// Valid returns true if and only if the ValidatedYearsWithID corresponds to a recognized YearsWith
func (vi *ValidatedYearsWithID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedYearsWithID) ID() *YearsWithID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedYearsWithID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedYearsWithID) ValidatedID() *ValidatedYearsWithID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedYearsWithID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedYearsWithID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedYearsWithID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedYearsWithID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedYearsWithID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := YearsWithID(capString)
	item := YearsWith.ByID(&id)
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

func (vi ValidatedYearsWithID) String() string {
	return vi.ToIDString()
}

type YearsWithIdentifier interface {
	ID() *YearsWithID
	Valid() bool
}
