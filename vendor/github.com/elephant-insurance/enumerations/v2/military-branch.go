package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// MilitaryBranchID uniquely identifies a particular MilitaryBranch
type MilitaryBranchID string

// Clone creates a safe, independent copy of a MilitaryBranchID
func (i *MilitaryBranchID) Clone() *MilitaryBranchID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two MilitaryBranchIds are equivalent
func (i *MilitaryBranchID) Equals(j *MilitaryBranchID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *MilitaryBranchID that is either valid or nil
func (i *MilitaryBranchID) ID() *MilitaryBranchID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *MilitaryBranchID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the MilitaryBranchID corresponds to a recognized MilitaryBranch
func (i *MilitaryBranchID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return MilitaryBranch.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *MilitaryBranchID) ValidatedID() *ValidatedMilitaryBranchID {
	if i != nil {
		return &ValidatedMilitaryBranchID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *MilitaryBranchID) MarshalJSON() ([]byte, error) {
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

func (i *MilitaryBranchID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := MilitaryBranchID(dataString)
	item := MilitaryBranch.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	militaryBranchAirForceID   MilitaryBranchID = "AirForce"
	militaryBranchArmyID       MilitaryBranchID = "Army"
	militaryBranchCoastGuardID MilitaryBranchID = "CoastGuard"
	militaryBranchMarinesID    MilitaryBranchID = "MarineCorps"
	militaryBranchNavyID       MilitaryBranchID = "Navy"
)

// EnumMilitaryBranchItem describes an entry in an enumeration of MilitaryBranch
type EnumMilitaryBranchItem struct {
	ID        MilitaryBranchID  `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	militaryBranchAirForce   = EnumMilitaryBranchItem{militaryBranchAirForceID, "Air Force", nil, "AirForce", 1}
	militaryBranchArmy       = EnumMilitaryBranchItem{militaryBranchArmyID, "Army", nil, "Army", 2}
	militaryBranchCoastGuard = EnumMilitaryBranchItem{militaryBranchCoastGuardID, "Coast Guard", nil, "CoastGuard", 3}
	militaryBranchMarines    = EnumMilitaryBranchItem{militaryBranchMarinesID, "Marines", nil, "Marines", 4}
	militaryBranchNavy       = EnumMilitaryBranchItem{militaryBranchNavyID, "Navy", nil, "Navy", 5}
)

// EnumMilitaryBranch is a collection of MilitaryBranch items
type EnumMilitaryBranch struct {
	Description string
	Items       []*EnumMilitaryBranchItem
	Name        string

	AirForce   *EnumMilitaryBranchItem
	Army       *EnumMilitaryBranchItem
	CoastGuard *EnumMilitaryBranchItem
	Marines    *EnumMilitaryBranchItem
	Navy       *EnumMilitaryBranchItem

	itemDict map[string]*EnumMilitaryBranchItem
}

// MilitaryBranch is a public singleton instance of EnumMilitaryBranch
// representing branches of the US Military
var MilitaryBranch = &EnumMilitaryBranch{
	Description: "branches of the US Military",
	Items: []*EnumMilitaryBranchItem{
		&militaryBranchAirForce,
		&militaryBranchArmy,
		&militaryBranchCoastGuard,
		&militaryBranchMarines,
		&militaryBranchNavy,
	},
	Name:       "EnumMilitaryBranch",
	AirForce:   &militaryBranchAirForce,
	Army:       &militaryBranchArmy,
	CoastGuard: &militaryBranchCoastGuard,
	Marines:    &militaryBranchMarines,
	Navy:       &militaryBranchNavy,

	itemDict: map[string]*EnumMilitaryBranchItem{
		strings.ToLower(string(militaryBranchAirForceID)):   &militaryBranchAirForce,
		strings.ToLower(string(militaryBranchArmyID)):       &militaryBranchArmy,
		strings.ToLower(string(militaryBranchCoastGuardID)): &militaryBranchCoastGuard,
		strings.ToLower(string(militaryBranchMarinesID)):    &militaryBranchMarines,
		strings.ToLower(string(militaryBranchNavyID)):       &militaryBranchNavy,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumMilitaryBranch) ByID(id MilitaryBranchIdentifier) *EnumMilitaryBranchItem {
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
func (e *EnumMilitaryBranch) ByIDString(idx string) *EnumMilitaryBranchItem {
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
func (e *EnumMilitaryBranch) ByIndex(idx int) *EnumMilitaryBranchItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedMilitaryBranchID is a struct that is designed to replace a *MilitaryBranchID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *MilitaryBranchID it contains while being a better JSON citizen.
type ValidatedMilitaryBranchID struct {
	// id will point to a valid MilitaryBranchID, if possible
	// If id is nil, then ValidatedMilitaryBranchID.Valid() will return false.
	id *MilitaryBranchID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedMilitaryBranchID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedMilitaryBranchID
func (vi *ValidatedMilitaryBranchID) Clone() *ValidatedMilitaryBranchID {
	if vi == nil {
		return nil
	}

	var cid *MilitaryBranchID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedMilitaryBranchID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedMilitaryBranchIds represent the same MilitaryBranch
func (vi *ValidatedMilitaryBranchID) Equals(vj *ValidatedMilitaryBranchID) bool {
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

// Valid returns true if and only if the ValidatedMilitaryBranchID corresponds to a recognized MilitaryBranch
func (vi *ValidatedMilitaryBranchID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedMilitaryBranchID) ID() *MilitaryBranchID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedMilitaryBranchID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedMilitaryBranchID) ValidatedID() *ValidatedMilitaryBranchID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedMilitaryBranchID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedMilitaryBranchID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedMilitaryBranchID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedMilitaryBranchID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedMilitaryBranchID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := MilitaryBranchID(capString)
	item := MilitaryBranch.ByID(&id)
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

func (vi ValidatedMilitaryBranchID) String() string {
	return vi.ToIDString()
}

type MilitaryBranchIdentifier interface {
	ID() *MilitaryBranchID
	Valid() bool
}
