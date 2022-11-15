package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// VehicleUsageID uniquely identifies a particular VehicleUsage
type VehicleUsageID string

// Clone creates a safe, independent copy of a VehicleUsageID
func (i *VehicleUsageID) Clone() *VehicleUsageID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two VehicleUsageIds are equivalent
func (i *VehicleUsageID) Equals(j *VehicleUsageID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *VehicleUsageID that is either valid or nil
func (i *VehicleUsageID) ID() *VehicleUsageID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *VehicleUsageID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the VehicleUsageID corresponds to a recognized VehicleUsage
func (i *VehicleUsageID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return VehicleUsage.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *VehicleUsageID) ValidatedID() *ValidatedVehicleUsageID {
	if i != nil {
		return &ValidatedVehicleUsageID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *VehicleUsageID) MarshalJSON() ([]byte, error) {
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

func (i *VehicleUsageID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := VehicleUsageID(dataString)
	item := VehicleUsage.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	vehicleUsageCommutingID   VehicleUsageID = "commuting"
	vehicleUsagePleasureID    VehicleUsageID = "pleasure"
	vehicleUsageBusinessID    VehicleUsageID = "business"
	vehicleUsageRidesharingID VehicleUsageID = "businessForRide"
)

// EnumVehicleUsageItem describes an entry in an enumeration of VehicleUsage
type EnumVehicleUsageItem struct {
	ID        VehicleUsageID    `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	vehicleUsageCommuting   = EnumVehicleUsageItem{vehicleUsageCommutingID, "Commute to work or school", nil, "Commuting", 0}
	vehicleUsagePleasure    = EnumVehicleUsageItem{vehicleUsagePleasureID, "Pleasure", nil, "Pleasure", 1}
	vehicleUsageBusiness    = EnumVehicleUsageItem{vehicleUsageBusinessID, "Business", nil, "Business", 2}
	vehicleUsageRidesharing = EnumVehicleUsageItem{vehicleUsageRidesharingID, "Business including ridesharing", nil, "Ridesharing", 3}
)

// EnumVehicleUsage is a collection of VehicleUsage items
type EnumVehicleUsage struct {
	Description string
	Items       []*EnumVehicleUsageItem
	Name        string

	Commuting   *EnumVehicleUsageItem
	Pleasure    *EnumVehicleUsageItem
	Business    *EnumVehicleUsageItem
	Ridesharing *EnumVehicleUsageItem

	itemDict map[string]*EnumVehicleUsageItem
}

// VehicleUsage is a public singleton instance of EnumVehicleUsage
// representing usages for vehicles
var VehicleUsage = &EnumVehicleUsage{
	Description: "usages for vehicles",
	Items: []*EnumVehicleUsageItem{
		&vehicleUsageCommuting,
		&vehicleUsagePleasure,
		&vehicleUsageBusiness,
		&vehicleUsageRidesharing,
	},
	Name:        "EnumVehicleUsage",
	Commuting:   &vehicleUsageCommuting,
	Pleasure:    &vehicleUsagePleasure,
	Business:    &vehicleUsageBusiness,
	Ridesharing: &vehicleUsageRidesharing,

	itemDict: map[string]*EnumVehicleUsageItem{
		strings.ToLower(string(vehicleUsageCommutingID)):   &vehicleUsageCommuting,
		strings.ToLower(string(vehicleUsagePleasureID)):    &vehicleUsagePleasure,
		strings.ToLower(string(vehicleUsageBusinessID)):    &vehicleUsageBusiness,
		strings.ToLower(string(vehicleUsageRidesharingID)): &vehicleUsageRidesharing,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumVehicleUsage) ByID(id VehicleUsageIdentifier) *EnumVehicleUsageItem {
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
func (e *EnumVehicleUsage) ByIDString(idx string) *EnumVehicleUsageItem {
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
func (e *EnumVehicleUsage) ByIndex(idx int) *EnumVehicleUsageItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedVehicleUsageID is a struct that is designed to replace a *VehicleUsageID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *VehicleUsageID it contains while being a better JSON citizen.
type ValidatedVehicleUsageID struct {
	// id will point to a valid VehicleUsageID, if possible
	// If id is nil, then ValidatedVehicleUsageID.Valid() will return false.
	id *VehicleUsageID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedVehicleUsageID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedVehicleUsageID
func (vi *ValidatedVehicleUsageID) Clone() *ValidatedVehicleUsageID {
	if vi == nil {
		return nil
	}

	var cid *VehicleUsageID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedVehicleUsageID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedVehicleUsageIds represent the same VehicleUsage
func (vi *ValidatedVehicleUsageID) Equals(vj *ValidatedVehicleUsageID) bool {
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

// Valid returns true if and only if the ValidatedVehicleUsageID corresponds to a recognized VehicleUsage
func (vi *ValidatedVehicleUsageID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedVehicleUsageID) ID() *VehicleUsageID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedVehicleUsageID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedVehicleUsageID) ValidatedID() *ValidatedVehicleUsageID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedVehicleUsageID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedVehicleUsageID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedVehicleUsageID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedVehicleUsageID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedVehicleUsageID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := VehicleUsageID(capString)
	item := VehicleUsage.ByID(&id)
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

func (vi ValidatedVehicleUsageID) String() string {
	return vi.ToIDString()
}

type VehicleUsageIdentifier interface {
	ID() *VehicleUsageID
	Valid() bool
}
