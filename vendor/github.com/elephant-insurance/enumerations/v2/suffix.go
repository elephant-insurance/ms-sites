package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// SuffixID uniquely identifies a particular Suffix
type SuffixID string

// Clone creates a safe, independent copy of a SuffixID
func (i *SuffixID) Clone() *SuffixID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two SuffixIds are equivalent
func (i *SuffixID) Equals(j *SuffixID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *SuffixID that is either valid or nil
func (i *SuffixID) ID() *SuffixID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *SuffixID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the SuffixID corresponds to a recognized Suffix
func (i *SuffixID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return Suffix.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *SuffixID) ValidatedID() *ValidatedSuffixID {
	if i != nil {
		return &ValidatedSuffixID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *SuffixID) MarshalJSON() ([]byte, error) {
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

func (i *SuffixID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := SuffixID(dataString)
	item := Suffix.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	suffixJuniorID    SuffixID = "jr"
	suffixSeniorID    SuffixID = "sr"
	suffixTheFirstID  SuffixID = "c_Ir"
	suffixTheSecondID SuffixID = "c_II"
	suffixTheThirdID  SuffixID = "c_III"
	suffixTheFourthID SuffixID = "c_IV"
)

// EnumSuffixItem describes an entry in an enumeration of Suffix
type EnumSuffixItem struct {
	ID        SuffixID          `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	suffixJunior    = EnumSuffixItem{suffixJuniorID, "Jr.", nil, "Junior", 1}
	suffixSenior    = EnumSuffixItem{suffixSeniorID, "Sr.", nil, "Senior", 2}
	suffixTheFirst  = EnumSuffixItem{suffixTheFirstID, "I", nil, "TheFirst", 3}
	suffixTheSecond = EnumSuffixItem{suffixTheSecondID, "II", nil, "TheSecond", 4}
	suffixTheThird  = EnumSuffixItem{suffixTheThirdID, "III", nil, "TheThird", 5}
	suffixTheFourth = EnumSuffixItem{suffixTheFourthID, "IV", nil, "TheFourth", 6}
)

// EnumSuffix is a collection of Suffix items
type EnumSuffix struct {
	Description string
	Items       []*EnumSuffixItem
	Name        string

	Junior    *EnumSuffixItem
	Senior    *EnumSuffixItem
	TheFirst  *EnumSuffixItem
	TheSecond *EnumSuffixItem
	TheThird  *EnumSuffixItem
	TheFourth *EnumSuffixItem

	itemDict map[string]*EnumSuffixItem
}

// Suffix is a public singleton instance of EnumSuffix
// representing suffixes for customer names
var Suffix = &EnumSuffix{
	Description: "suffixes for customer names",
	Items: []*EnumSuffixItem{
		&suffixJunior,
		&suffixSenior,
		&suffixTheFirst,
		&suffixTheSecond,
		&suffixTheThird,
		&suffixTheFourth,
	},
	Name:      "EnumSuffix",
	Junior:    &suffixJunior,
	Senior:    &suffixSenior,
	TheFirst:  &suffixTheFirst,
	TheSecond: &suffixTheSecond,
	TheThird:  &suffixTheThird,
	TheFourth: &suffixTheFourth,

	itemDict: map[string]*EnumSuffixItem{
		strings.ToLower(string(suffixJuniorID)):    &suffixJunior,
		strings.ToLower(string(suffixSeniorID)):    &suffixSenior,
		strings.ToLower(string(suffixTheFirstID)):  &suffixTheFirst,
		strings.ToLower(string(suffixTheSecondID)): &suffixTheSecond,
		strings.ToLower(string(suffixTheThirdID)):  &suffixTheThird,
		strings.ToLower(string(suffixTheFourthID)): &suffixTheFourth,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumSuffix) ByID(id SuffixIdentifier) *EnumSuffixItem {
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
func (e *EnumSuffix) ByIDString(idx string) *EnumSuffixItem {
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
func (e *EnumSuffix) ByIndex(idx int) *EnumSuffixItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedSuffixID is a struct that is designed to replace a *SuffixID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *SuffixID it contains while being a better JSON citizen.
type ValidatedSuffixID struct {
	// id will point to a valid SuffixID, if possible
	// If id is nil, then ValidatedSuffixID.Valid() will return false.
	id *SuffixID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedSuffixID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedSuffixID
func (vi *ValidatedSuffixID) Clone() *ValidatedSuffixID {
	if vi == nil {
		return nil
	}

	var cid *SuffixID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedSuffixID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedSuffixIds represent the same Suffix
func (vi *ValidatedSuffixID) Equals(vj *ValidatedSuffixID) bool {
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

// Valid returns true if and only if the ValidatedSuffixID corresponds to a recognized Suffix
func (vi *ValidatedSuffixID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedSuffixID) ID() *SuffixID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedSuffixID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedSuffixID) ValidatedID() *ValidatedSuffixID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedSuffixID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedSuffixID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedSuffixID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedSuffixID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedSuffixID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := SuffixID(capString)
	item := Suffix.ByID(&id)
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

func (vi ValidatedSuffixID) String() string {
	return vi.ToIDString()
}

type SuffixIdentifier interface {
	ID() *SuffixID
	Valid() bool
}
