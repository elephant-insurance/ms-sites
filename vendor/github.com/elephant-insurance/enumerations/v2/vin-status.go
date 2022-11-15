package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// VinStatusID uniquely identifies a particular VinStatus
type VinStatusID string

// Clone creates a safe, independent copy of a VinStatusID
func (i *VinStatusID) Clone() *VinStatusID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two VinStatusIds are equivalent
func (i *VinStatusID) Equals(j *VinStatusID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *VinStatusID that is either valid or nil
func (i *VinStatusID) ID() *VinStatusID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *VinStatusID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the VinStatusID corresponds to a recognized VinStatus
func (i *VinStatusID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return VinStatus.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *VinStatusID) ValidatedID() *ValidatedVinStatusID {
	if i != nil {
		return &ValidatedVinStatusID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *VinStatusID) MarshalJSON() ([]byte, error) {
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

func (i *VinStatusID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := VinStatusID(dataString)
	item := VinStatus.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	vinStatusCleanID                VinStatusID = "clean"
	vinStatusBrandedID              VinStatusID = "branded"
	vinStatusInvalidVinID           VinStatusID = "invalidVin"
	vinStatusCarfaxNotNeededID      VinStatusID = "carfaxNotNeeded"
	vinStatusCarfaxErrorID          VinStatusID = "carfaxError"
	vinStatusBrandedWithinOneYearID VinStatusID = "brandedWithinOneYear"
)

// EnumVinStatusItem describes an entry in an enumeration of VinStatus
type EnumVinStatusItem struct {
	ID        VinStatusID       `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	vinStatusClean                = EnumVinStatusItem{vinStatusCleanID, "clean status", nil, "Clean", 0}
	vinStatusBranded              = EnumVinStatusItem{vinStatusBrandedID, "branded status", nil, "Branded", 1}
	vinStatusInvalidVin           = EnumVinStatusItem{vinStatusInvalidVinID, "invalidVin status", nil, "InvalidVin", 2}
	vinStatusCarfaxNotNeeded      = EnumVinStatusItem{vinStatusCarfaxNotNeededID, "carfaxNotNeeded status", nil, "CarfaxNotNeeded", 3}
	vinStatusCarfaxError          = EnumVinStatusItem{vinStatusCarfaxErrorID, "carfaxError status", nil, "CarfaxError", 4}
	vinStatusBrandedWithinOneYear = EnumVinStatusItem{vinStatusBrandedWithinOneYearID, "brandedWithinOneYear status", nil, "BrandedWithinOneYear", 5}
)

// EnumVinStatus is a collection of VinStatus items
type EnumVinStatus struct {
	Description string
	Items       []*EnumVinStatusItem
	Name        string

	Clean                *EnumVinStatusItem
	Branded              *EnumVinStatusItem
	InvalidVin           *EnumVinStatusItem
	CarfaxNotNeeded      *EnumVinStatusItem
	CarfaxError          *EnumVinStatusItem
	BrandedWithinOneYear *EnumVinStatusItem

	itemDict map[string]*EnumVinStatusItem
}

// VinStatus is a public singleton instance of EnumVinStatus
// representing status for vin
var VinStatus = &EnumVinStatus{
	Description: "status for vin",
	Items: []*EnumVinStatusItem{
		&vinStatusClean,
		&vinStatusBranded,
		&vinStatusInvalidVin,
		&vinStatusCarfaxNotNeeded,
		&vinStatusCarfaxError,
		&vinStatusBrandedWithinOneYear,
	},
	Name:                 "EnumVinStatus",
	Clean:                &vinStatusClean,
	Branded:              &vinStatusBranded,
	InvalidVin:           &vinStatusInvalidVin,
	CarfaxNotNeeded:      &vinStatusCarfaxNotNeeded,
	CarfaxError:          &vinStatusCarfaxError,
	BrandedWithinOneYear: &vinStatusBrandedWithinOneYear,

	itemDict: map[string]*EnumVinStatusItem{
		strings.ToLower(string(vinStatusCleanID)):                &vinStatusClean,
		strings.ToLower(string(vinStatusBrandedID)):              &vinStatusBranded,
		strings.ToLower(string(vinStatusInvalidVinID)):           &vinStatusInvalidVin,
		strings.ToLower(string(vinStatusCarfaxNotNeededID)):      &vinStatusCarfaxNotNeeded,
		strings.ToLower(string(vinStatusCarfaxErrorID)):          &vinStatusCarfaxError,
		strings.ToLower(string(vinStatusBrandedWithinOneYearID)): &vinStatusBrandedWithinOneYear,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumVinStatus) ByID(id VinStatusIdentifier) *EnumVinStatusItem {
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
func (e *EnumVinStatus) ByIDString(idx string) *EnumVinStatusItem {
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
func (e *EnumVinStatus) ByIndex(idx int) *EnumVinStatusItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedVinStatusID is a struct that is designed to replace a *VinStatusID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *VinStatusID it contains while being a better JSON citizen.
type ValidatedVinStatusID struct {
	// id will point to a valid VinStatusID, if possible
	// If id is nil, then ValidatedVinStatusID.Valid() will return false.
	id *VinStatusID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedVinStatusID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedVinStatusID
func (vi *ValidatedVinStatusID) Clone() *ValidatedVinStatusID {
	if vi == nil {
		return nil
	}

	var cid *VinStatusID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedVinStatusID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedVinStatusIds represent the same VinStatus
func (vi *ValidatedVinStatusID) Equals(vj *ValidatedVinStatusID) bool {
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

// Valid returns true if and only if the ValidatedVinStatusID corresponds to a recognized VinStatus
func (vi *ValidatedVinStatusID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedVinStatusID) ID() *VinStatusID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedVinStatusID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedVinStatusID) ValidatedID() *ValidatedVinStatusID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedVinStatusID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedVinStatusID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedVinStatusID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedVinStatusID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedVinStatusID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := VinStatusID(capString)
	item := VinStatus.ByID(&id)
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

func (vi ValidatedVinStatusID) String() string {
	return vi.ToIDString()
}

type VinStatusIdentifier interface {
	ID() *VinStatusID
	Valid() bool
}
