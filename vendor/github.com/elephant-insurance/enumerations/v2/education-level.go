package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// EducationLevelID uniquely identifies a particular EducationLevel
type EducationLevelID string

// Clone creates a safe, independent copy of a EducationLevelID
func (i *EducationLevelID) Clone() *EducationLevelID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two EducationLevelIds are equivalent
func (i *EducationLevelID) Equals(j *EducationLevelID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *EducationLevelID that is either valid or nil
func (i *EducationLevelID) ID() *EducationLevelID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *EducationLevelID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the EducationLevelID corresponds to a recognized EducationLevel
func (i *EducationLevelID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return EducationLevel.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *EducationLevelID) ValidatedID() *ValidatedEducationLevelID {
	if i != nil {
		return &ValidatedEducationLevelID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *EducationLevelID) MarshalJSON() ([]byte, error) {
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

func (i *EducationLevelID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := EducationLevelID(dataString)
	item := EducationLevel.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	educationLevelLowerThanHighSchoolID  EducationLevelID = "lowerthanhighschool"
	educationLevelHighSchoolGEDID        EducationLevelID = "highschool"
	educationLevelSomeCollegeID          EducationLevelID = "somecollege"
	educationLevelAssociatesVocationalID EducationLevelID = "associatevocational"
	educationLevelBachelorsID            EducationLevelID = "bachelors"
	educationLevelMastersID              EducationLevelID = "masters"
	educationLevelDoctorateID            EducationLevelID = "phdmdjd"
)

// EnumEducationLevelItem describes an entry in an enumeration of EducationLevel
type EnumEducationLevelItem struct {
	ID        EducationLevelID  `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	educationLevelLowerThanHighSchool  = EnumEducationLevelItem{educationLevelLowerThanHighSchoolID, "Lower than High School", nil, "LowerThanHighSchool", 1}
	educationLevelHighSchoolGED        = EnumEducationLevelItem{educationLevelHighSchoolGEDID, "High School/GED", nil, "HighSchoolGED", 2}
	educationLevelSomeCollege          = EnumEducationLevelItem{educationLevelSomeCollegeID, "Some College", nil, "SomeCollege", 3}
	educationLevelAssociatesVocational = EnumEducationLevelItem{educationLevelAssociatesVocationalID, "Associates or Vocational", nil, "AssociatesVocational", 4}
	educationLevelBachelors            = EnumEducationLevelItem{educationLevelBachelorsID, "Bachelors", nil, "Bachelors", 5}
	educationLevelMasters              = EnumEducationLevelItem{educationLevelMastersID, "Masters", nil, "Masters", 6}
	educationLevelDoctorate            = EnumEducationLevelItem{educationLevelDoctorateID, "Doctorate", nil, "Doctorate", 7}
)

// EnumEducationLevel is a collection of EducationLevel items
type EnumEducationLevel struct {
	Description string
	Items       []*EnumEducationLevelItem
	Name        string

	LowerThanHighSchool  *EnumEducationLevelItem
	HighSchoolGED        *EnumEducationLevelItem
	SomeCollege          *EnumEducationLevelItem
	AssociatesVocational *EnumEducationLevelItem
	Bachelors            *EnumEducationLevelItem
	Masters              *EnumEducationLevelItem
	Doctorate            *EnumEducationLevelItem

	itemDict map[string]*EnumEducationLevelItem
}

// EducationLevel is a public singleton instance of EnumEducationLevel
// representing levels of educational achievement
var EducationLevel = &EnumEducationLevel{
	Description: "levels of educational achievement",
	Items: []*EnumEducationLevelItem{
		&educationLevelLowerThanHighSchool,
		&educationLevelHighSchoolGED,
		&educationLevelSomeCollege,
		&educationLevelAssociatesVocational,
		&educationLevelBachelors,
		&educationLevelMasters,
		&educationLevelDoctorate,
	},
	Name:                 "EnumEducationLevel",
	LowerThanHighSchool:  &educationLevelLowerThanHighSchool,
	HighSchoolGED:        &educationLevelHighSchoolGED,
	SomeCollege:          &educationLevelSomeCollege,
	AssociatesVocational: &educationLevelAssociatesVocational,
	Bachelors:            &educationLevelBachelors,
	Masters:              &educationLevelMasters,
	Doctorate:            &educationLevelDoctorate,

	itemDict: map[string]*EnumEducationLevelItem{
		strings.ToLower(string(educationLevelLowerThanHighSchoolID)):  &educationLevelLowerThanHighSchool,
		strings.ToLower(string(educationLevelHighSchoolGEDID)):        &educationLevelHighSchoolGED,
		strings.ToLower(string(educationLevelSomeCollegeID)):          &educationLevelSomeCollege,
		strings.ToLower(string(educationLevelAssociatesVocationalID)): &educationLevelAssociatesVocational,
		strings.ToLower(string(educationLevelBachelorsID)):            &educationLevelBachelors,
		strings.ToLower(string(educationLevelMastersID)):              &educationLevelMasters,
		strings.ToLower(string(educationLevelDoctorateID)):            &educationLevelDoctorate,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumEducationLevel) ByID(id EducationLevelIdentifier) *EnumEducationLevelItem {
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
func (e *EnumEducationLevel) ByIDString(idx string) *EnumEducationLevelItem {
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
func (e *EnumEducationLevel) ByIndex(idx int) *EnumEducationLevelItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedEducationLevelID is a struct that is designed to replace a *EducationLevelID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *EducationLevelID it contains while being a better JSON citizen.
type ValidatedEducationLevelID struct {
	// id will point to a valid EducationLevelID, if possible
	// If id is nil, then ValidatedEducationLevelID.Valid() will return false.
	id *EducationLevelID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedEducationLevelID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedEducationLevelID
func (vi *ValidatedEducationLevelID) Clone() *ValidatedEducationLevelID {
	if vi == nil {
		return nil
	}

	var cid *EducationLevelID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedEducationLevelID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedEducationLevelIds represent the same EducationLevel
func (vi *ValidatedEducationLevelID) Equals(vj *ValidatedEducationLevelID) bool {
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

// Valid returns true if and only if the ValidatedEducationLevelID corresponds to a recognized EducationLevel
func (vi *ValidatedEducationLevelID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedEducationLevelID) ID() *EducationLevelID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedEducationLevelID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedEducationLevelID) ValidatedID() *ValidatedEducationLevelID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedEducationLevelID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedEducationLevelID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedEducationLevelID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedEducationLevelID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedEducationLevelID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := EducationLevelID(capString)
	item := EducationLevel.ByID(&id)
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

func (vi ValidatedEducationLevelID) String() string {
	return vi.ToIDString()
}

type EducationLevelIdentifier interface {
	ID() *EducationLevelID
	Valid() bool
}
