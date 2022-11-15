package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// StudentEnrollmentID uniquely identifies a particular StudentEnrollment
type StudentEnrollmentID string

// Clone creates a safe, independent copy of a StudentEnrollmentID
func (i *StudentEnrollmentID) Clone() *StudentEnrollmentID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two StudentEnrollmentIds are equivalent
func (i *StudentEnrollmentID) Equals(j *StudentEnrollmentID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *StudentEnrollmentID that is either valid or nil
func (i *StudentEnrollmentID) ID() *StudentEnrollmentID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *StudentEnrollmentID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the StudentEnrollmentID corresponds to a recognized StudentEnrollment
func (i *StudentEnrollmentID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return StudentEnrollment.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *StudentEnrollmentID) ValidatedID() *ValidatedStudentEnrollmentID {
	if i != nil {
		return &ValidatedStudentEnrollmentID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *StudentEnrollmentID) MarshalJSON() ([]byte, error) {
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

func (i *StudentEnrollmentID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := StudentEnrollmentID(dataString)
	item := StudentEnrollment.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	studentEnrollmentHighSchoolID             StudentEnrollmentID = "HighSchool"
	studentEnrollmentTradeSchoolID            StudentEnrollmentID = "TradeSchool"
	studentEnrollmentCollegeDegreeFor2YearsID StudentEnrollmentID = "CollegeDegreeFor2Years"
	studentEnrollmentCollegeDegreeFor4YearsID StudentEnrollmentID = "CollegeDegreeFor4Years"
	studentEnrollmentGraduateSchoolID         StudentEnrollmentID = "GraduateSchool"
)

// EnumStudentEnrollmentItem describes an entry in an enumeration of StudentEnrollment
type EnumStudentEnrollmentItem struct {
	ID        StudentEnrollmentID `json:"Value"`
	Desc      string              `json:"Description,omitempty"`
	Meta      map[string]string   `json:",omitempty"`
	Name      string              `json:"Name"`
	SortOrder int
}

var (
	studentEnrollmentHighSchool             = EnumStudentEnrollmentItem{studentEnrollmentHighSchoolID, "High School", nil, "HighSchool", 1}
	studentEnrollmentTradeSchool            = EnumStudentEnrollmentItem{studentEnrollmentTradeSchoolID, "Trade School", nil, "TradeSchool", 2}
	studentEnrollmentCollegeDegreeFor2Years = EnumStudentEnrollmentItem{studentEnrollmentCollegeDegreeFor2YearsID, "College (2 Year Degree)", nil, "CollegeDegreeFor2Years", 3}
	studentEnrollmentCollegeDegreeFor4Years = EnumStudentEnrollmentItem{studentEnrollmentCollegeDegreeFor4YearsID, "College (4 Year Degree)", nil, "CollegeDegreeFor4Years", 4}
	studentEnrollmentGraduateSchool         = EnumStudentEnrollmentItem{studentEnrollmentGraduateSchoolID, "Graduate School", nil, "GraduateSchool", 5}
)

// EnumStudentEnrollment is a collection of StudentEnrollment items
type EnumStudentEnrollment struct {
	Description string
	Items       []*EnumStudentEnrollmentItem
	Name        string

	HighSchool             *EnumStudentEnrollmentItem
	TradeSchool            *EnumStudentEnrollmentItem
	CollegeDegreeFor2Years *EnumStudentEnrollmentItem
	CollegeDegreeFor4Years *EnumStudentEnrollmentItem
	GraduateSchool         *EnumStudentEnrollmentItem

	itemDict map[string]*EnumStudentEnrollmentItem
}

// StudentEnrollment is a public singleton instance of EnumStudentEnrollment
// representing Types of school a student may be enrolled in
var StudentEnrollment = &EnumStudentEnrollment{
	Description: "Types of school a student may be enrolled in",
	Items: []*EnumStudentEnrollmentItem{
		&studentEnrollmentHighSchool,
		&studentEnrollmentTradeSchool,
		&studentEnrollmentCollegeDegreeFor2Years,
		&studentEnrollmentCollegeDegreeFor4Years,
		&studentEnrollmentGraduateSchool,
	},
	Name:                   "EnumStudentEnrollment",
	HighSchool:             &studentEnrollmentHighSchool,
	TradeSchool:            &studentEnrollmentTradeSchool,
	CollegeDegreeFor2Years: &studentEnrollmentCollegeDegreeFor2Years,
	CollegeDegreeFor4Years: &studentEnrollmentCollegeDegreeFor4Years,
	GraduateSchool:         &studentEnrollmentGraduateSchool,

	itemDict: map[string]*EnumStudentEnrollmentItem{
		strings.ToLower(string(studentEnrollmentHighSchoolID)):             &studentEnrollmentHighSchool,
		strings.ToLower(string(studentEnrollmentTradeSchoolID)):            &studentEnrollmentTradeSchool,
		strings.ToLower(string(studentEnrollmentCollegeDegreeFor2YearsID)): &studentEnrollmentCollegeDegreeFor2Years,
		strings.ToLower(string(studentEnrollmentCollegeDegreeFor4YearsID)): &studentEnrollmentCollegeDegreeFor4Years,
		strings.ToLower(string(studentEnrollmentGraduateSchoolID)):         &studentEnrollmentGraduateSchool,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumStudentEnrollment) ByID(id StudentEnrollmentIdentifier) *EnumStudentEnrollmentItem {
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
func (e *EnumStudentEnrollment) ByIDString(idx string) *EnumStudentEnrollmentItem {
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
func (e *EnumStudentEnrollment) ByIndex(idx int) *EnumStudentEnrollmentItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedStudentEnrollmentID is a struct that is designed to replace a *StudentEnrollmentID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *StudentEnrollmentID it contains while being a better JSON citizen.
type ValidatedStudentEnrollmentID struct {
	// id will point to a valid StudentEnrollmentID, if possible
	// If id is nil, then ValidatedStudentEnrollmentID.Valid() will return false.
	id *StudentEnrollmentID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedStudentEnrollmentID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedStudentEnrollmentID
func (vi *ValidatedStudentEnrollmentID) Clone() *ValidatedStudentEnrollmentID {
	if vi == nil {
		return nil
	}

	var cid *StudentEnrollmentID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedStudentEnrollmentID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedStudentEnrollmentIds represent the same StudentEnrollment
func (vi *ValidatedStudentEnrollmentID) Equals(vj *ValidatedStudentEnrollmentID) bool {
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

// Valid returns true if and only if the ValidatedStudentEnrollmentID corresponds to a recognized StudentEnrollment
func (vi *ValidatedStudentEnrollmentID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedStudentEnrollmentID) ID() *StudentEnrollmentID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedStudentEnrollmentID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedStudentEnrollmentID) ValidatedID() *ValidatedStudentEnrollmentID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedStudentEnrollmentID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedStudentEnrollmentID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedStudentEnrollmentID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedStudentEnrollmentID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedStudentEnrollmentID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := StudentEnrollmentID(capString)
	item := StudentEnrollment.ByID(&id)
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

func (vi ValidatedStudentEnrollmentID) String() string {
	return vi.ToIDString()
}

type StudentEnrollmentIdentifier interface {
	ID() *StudentEnrollmentID
	Valid() bool
}
