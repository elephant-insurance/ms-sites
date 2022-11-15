package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// EmploymentStatusID uniquely identifies a particular EmploymentStatus
type EmploymentStatusID string

// Clone creates a safe, independent copy of a EmploymentStatusID
func (i *EmploymentStatusID) Clone() *EmploymentStatusID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two EmploymentStatusIds are equivalent
func (i *EmploymentStatusID) Equals(j *EmploymentStatusID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *EmploymentStatusID that is either valid or nil
func (i *EmploymentStatusID) ID() *EmploymentStatusID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *EmploymentStatusID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the EmploymentStatusID corresponds to a recognized EmploymentStatus
func (i *EmploymentStatusID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return EmploymentStatus.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *EmploymentStatusID) ValidatedID() *ValidatedEmploymentStatusID {
	if i != nil {
		return &ValidatedEmploymentStatusID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *EmploymentStatusID) MarshalJSON() ([]byte, error) {
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

func (i *EmploymentStatusID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := EmploymentStatusID(dataString)
	item := EmploymentStatus.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	employmentStatusEmployedID        EmploymentStatusID = "EmployedPrivately"
	employmentStatusDisabledID        EmploymentStatusID = "Disabled"
	employmentStatusFullTimeStudentID EmploymentStatusID = "FullStudent"
	employmentStatusHomemakerID       EmploymentStatusID = "HomeMaker"
	employmentStatusMilitaryActiveID  EmploymentStatusID = "ServeMilitary"
	employmentStatusMilitaryRetiredID EmploymentStatusID = "RetiredMilitary"
	employmentStatusUnemployedID      EmploymentStatusID = "Unemployed"
	employmentStatusRetiredID         EmploymentStatusID = "Retired"
	employmentStatusSelfEmployedID    EmploymentStatusID = "SelfEmp"
)

// EnumEmploymentStatusItem describes an entry in an enumeration of EmploymentStatus
type EnumEmploymentStatusItem struct {
	ID        EmploymentStatusID `json:"Value"`
	Desc      string             `json:"Description,omitempty"`
	Meta      map[string]string  `json:",omitempty"`
	Name      string             `json:"Name"`
	SortOrder int
}

var (
	employmentStatusEmployed        = EnumEmploymentStatusItem{employmentStatusEmployedID, "Employed", nil, "Employed", 1}
	employmentStatusDisabled        = EnumEmploymentStatusItem{employmentStatusDisabledID, "Disabled", nil, "Disabled", 2}
	employmentStatusFullTimeStudent = EnumEmploymentStatusItem{employmentStatusFullTimeStudentID, "Full-Time Student", nil, "FullTimeStudent", 3}
	employmentStatusHomemaker       = EnumEmploymentStatusItem{employmentStatusHomemakerID, "Homemaker", nil, "Homemaker", 4}
	employmentStatusMilitaryActive  = EnumEmploymentStatusItem{employmentStatusMilitaryActiveID, "Military – Active", nil, "MilitaryActive", 5}
	employmentStatusMilitaryRetired = EnumEmploymentStatusItem{employmentStatusMilitaryRetiredID, "Military – Retired", nil, "MilitaryRetired", 6}
	employmentStatusUnemployed      = EnumEmploymentStatusItem{employmentStatusUnemployedID, "Not Currently Employed", nil, "Unemployed", 7}
	employmentStatusRetired         = EnumEmploymentStatusItem{employmentStatusRetiredID, "Retired", nil, "Retired", 8}
	employmentStatusSelfEmployed    = EnumEmploymentStatusItem{employmentStatusSelfEmployedID, "Self-Employed", nil, "SelfEmployed", 9}
)

// EnumEmploymentStatus is a collection of EmploymentStatus items
type EnumEmploymentStatus struct {
	Description string
	Items       []*EnumEmploymentStatusItem
	Name        string

	Employed        *EnumEmploymentStatusItem
	Disabled        *EnumEmploymentStatusItem
	FullTimeStudent *EnumEmploymentStatusItem
	Homemaker       *EnumEmploymentStatusItem
	MilitaryActive  *EnumEmploymentStatusItem
	MilitaryRetired *EnumEmploymentStatusItem
	Unemployed      *EnumEmploymentStatusItem
	Retired         *EnumEmploymentStatusItem
	SelfEmployed    *EnumEmploymentStatusItem

	itemDict map[string]*EnumEmploymentStatusItem
}

// EmploymentStatus is a public singleton instance of EnumEmploymentStatus
// representing employment statuses
var EmploymentStatus = &EnumEmploymentStatus{
	Description: "employment statuses",
	Items: []*EnumEmploymentStatusItem{
		&employmentStatusEmployed,
		&employmentStatusDisabled,
		&employmentStatusFullTimeStudent,
		&employmentStatusHomemaker,
		&employmentStatusMilitaryActive,
		&employmentStatusMilitaryRetired,
		&employmentStatusUnemployed,
		&employmentStatusRetired,
		&employmentStatusSelfEmployed,
	},
	Name:            "EnumEmploymentStatus",
	Employed:        &employmentStatusEmployed,
	Disabled:        &employmentStatusDisabled,
	FullTimeStudent: &employmentStatusFullTimeStudent,
	Homemaker:       &employmentStatusHomemaker,
	MilitaryActive:  &employmentStatusMilitaryActive,
	MilitaryRetired: &employmentStatusMilitaryRetired,
	Unemployed:      &employmentStatusUnemployed,
	Retired:         &employmentStatusRetired,
	SelfEmployed:    &employmentStatusSelfEmployed,

	itemDict: map[string]*EnumEmploymentStatusItem{
		strings.ToLower(string(employmentStatusEmployedID)):        &employmentStatusEmployed,
		strings.ToLower(string(employmentStatusDisabledID)):        &employmentStatusDisabled,
		strings.ToLower(string(employmentStatusFullTimeStudentID)): &employmentStatusFullTimeStudent,
		strings.ToLower(string(employmentStatusHomemakerID)):       &employmentStatusHomemaker,
		strings.ToLower(string(employmentStatusMilitaryActiveID)):  &employmentStatusMilitaryActive,
		strings.ToLower(string(employmentStatusMilitaryRetiredID)): &employmentStatusMilitaryRetired,
		strings.ToLower(string(employmentStatusUnemployedID)):      &employmentStatusUnemployed,
		strings.ToLower(string(employmentStatusRetiredID)):         &employmentStatusRetired,
		strings.ToLower(string(employmentStatusSelfEmployedID)):    &employmentStatusSelfEmployed,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumEmploymentStatus) ByID(id EmploymentStatusIdentifier) *EnumEmploymentStatusItem {
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
func (e *EnumEmploymentStatus) ByIDString(idx string) *EnumEmploymentStatusItem {
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
func (e *EnumEmploymentStatus) ByIndex(idx int) *EnumEmploymentStatusItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedEmploymentStatusID is a struct that is designed to replace a *EmploymentStatusID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *EmploymentStatusID it contains while being a better JSON citizen.
type ValidatedEmploymentStatusID struct {
	// id will point to a valid EmploymentStatusID, if possible
	// If id is nil, then ValidatedEmploymentStatusID.Valid() will return false.
	id *EmploymentStatusID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedEmploymentStatusID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedEmploymentStatusID
func (vi *ValidatedEmploymentStatusID) Clone() *ValidatedEmploymentStatusID {
	if vi == nil {
		return nil
	}

	var cid *EmploymentStatusID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedEmploymentStatusID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedEmploymentStatusIds represent the same EmploymentStatus
func (vi *ValidatedEmploymentStatusID) Equals(vj *ValidatedEmploymentStatusID) bool {
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

// Valid returns true if and only if the ValidatedEmploymentStatusID corresponds to a recognized EmploymentStatus
func (vi *ValidatedEmploymentStatusID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedEmploymentStatusID) ID() *EmploymentStatusID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedEmploymentStatusID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedEmploymentStatusID) ValidatedID() *ValidatedEmploymentStatusID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedEmploymentStatusID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedEmploymentStatusID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedEmploymentStatusID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedEmploymentStatusID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedEmploymentStatusID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := EmploymentStatusID(capString)
	item := EmploymentStatus.ByID(&id)
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

func (vi ValidatedEmploymentStatusID) String() string {
	return vi.ToIDString()
}

type EmploymentStatusIdentifier interface {
	ID() *EmploymentStatusID
	Valid() bool
}
