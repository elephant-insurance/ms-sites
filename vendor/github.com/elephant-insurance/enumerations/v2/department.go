package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// DepartmentID uniquely identifies a particular Department
type DepartmentID string

// Clone creates a safe, independent copy of a DepartmentID
func (i *DepartmentID) Clone() *DepartmentID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two DepartmentIds are equivalent
func (i *DepartmentID) Equals(j *DepartmentID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *DepartmentID that is either valid or nil
func (i *DepartmentID) ID() *DepartmentID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *DepartmentID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the DepartmentID corresponds to a recognized Department
func (i *DepartmentID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return Department.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *DepartmentID) ValidatedID() *ValidatedDepartmentID {
	if i != nil {
		return &ValidatedDepartmentID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *DepartmentID) MarshalJSON() ([]byte, error) {
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

func (i *DepartmentID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := DepartmentID(dataString)
	item := Department.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	departmentUnknownID      DepartmentID = "Unknown"
	departmentSalesID        DepartmentID = "Sales"
	departmentServiceID      DepartmentID = "Service"
	departmentClaimsID       DepartmentID = "Claims"
	departmentMarketingID    DepartmentID = "Marketing"
	departmentUnderwritingID DepartmentID = "Underwriting"
)

// EnumDepartmentItem describes an entry in an enumeration of Department
type EnumDepartmentItem struct {
	ID        DepartmentID      `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	departmentUnknown      = EnumDepartmentItem{departmentUnknownID, "Unknown", nil, "Unknown", 1}
	departmentSales        = EnumDepartmentItem{departmentSalesID, "Sales", nil, "Sales", 2}
	departmentService      = EnumDepartmentItem{departmentServiceID, "Service", nil, "Service", 3}
	departmentClaims       = EnumDepartmentItem{departmentClaimsID, "Claims", nil, "Claims", 4}
	departmentMarketing    = EnumDepartmentItem{departmentMarketingID, "Marketing", nil, "Marketing", 5}
	departmentUnderwriting = EnumDepartmentItem{departmentUnderwritingID, "Underwriting", nil, "Underwriting", 6}
)

// EnumDepartment is a collection of Department items
type EnumDepartment struct {
	Description string
	Items       []*EnumDepartmentItem
	Name        string

	Unknown      *EnumDepartmentItem
	Sales        *EnumDepartmentItem
	Service      *EnumDepartmentItem
	Claims       *EnumDepartmentItem
	Marketing    *EnumDepartmentItem
	Underwriting *EnumDepartmentItem

	itemDict map[string]*EnumDepartmentItem
}

// Department is a public singleton instance of EnumDepartment
// representing departments in our company
var Department = &EnumDepartment{
	Description: "departments in our company",
	Items: []*EnumDepartmentItem{
		&departmentUnknown,
		&departmentSales,
		&departmentService,
		&departmentClaims,
		&departmentMarketing,
		&departmentUnderwriting,
	},
	Name:         "EnumDepartment",
	Unknown:      &departmentUnknown,
	Sales:        &departmentSales,
	Service:      &departmentService,
	Claims:       &departmentClaims,
	Marketing:    &departmentMarketing,
	Underwriting: &departmentUnderwriting,

	itemDict: map[string]*EnumDepartmentItem{
		strings.ToLower(string(departmentUnknownID)):      &departmentUnknown,
		strings.ToLower(string(departmentSalesID)):        &departmentSales,
		strings.ToLower(string(departmentServiceID)):      &departmentService,
		strings.ToLower(string(departmentClaimsID)):       &departmentClaims,
		strings.ToLower(string(departmentMarketingID)):    &departmentMarketing,
		strings.ToLower(string(departmentUnderwritingID)): &departmentUnderwriting,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumDepartment) ByID(id DepartmentIdentifier) *EnumDepartmentItem {
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
func (e *EnumDepartment) ByIDString(idx string) *EnumDepartmentItem {
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
func (e *EnumDepartment) ByIndex(idx int) *EnumDepartmentItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedDepartmentID is a struct that is designed to replace a *DepartmentID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *DepartmentID it contains while being a better JSON citizen.
type ValidatedDepartmentID struct {
	// id will point to a valid DepartmentID, if possible
	// If id is nil, then ValidatedDepartmentID.Valid() will return false.
	id *DepartmentID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedDepartmentID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedDepartmentID
func (vi *ValidatedDepartmentID) Clone() *ValidatedDepartmentID {
	if vi == nil {
		return nil
	}

	var cid *DepartmentID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedDepartmentID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedDepartmentIds represent the same Department
func (vi *ValidatedDepartmentID) Equals(vj *ValidatedDepartmentID) bool {
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

// Valid returns true if and only if the ValidatedDepartmentID corresponds to a recognized Department
func (vi *ValidatedDepartmentID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedDepartmentID) ID() *DepartmentID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedDepartmentID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedDepartmentID) ValidatedID() *ValidatedDepartmentID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedDepartmentID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedDepartmentID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedDepartmentID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedDepartmentID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedDepartmentID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := DepartmentID(capString)
	item := Department.ByID(&id)
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

func (vi ValidatedDepartmentID) String() string {
	return vi.ToIDString()
}

type DepartmentIdentifier interface {
	ID() *DepartmentID
	Valid() bool
}
