package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// VehicleCoverageTermID uniquely identifies a particular VehicleCoverageTerm
type VehicleCoverageTermID string

// Clone creates a safe, independent copy of a VehicleCoverageTermID
func (i *VehicleCoverageTermID) Clone() *VehicleCoverageTermID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two VehicleCoverageTermIds are equivalent
func (i *VehicleCoverageTermID) Equals(j *VehicleCoverageTermID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *VehicleCoverageTermID that is either valid or nil
func (i *VehicleCoverageTermID) ID() *VehicleCoverageTermID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *VehicleCoverageTermID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the VehicleCoverageTermID corresponds to a recognized VehicleCoverageTerm
func (i *VehicleCoverageTermID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return VehicleCoverageTerm.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *VehicleCoverageTermID) ValidatedID() *ValidatedVehicleCoverageTermID {
	if i != nil {
		return &ValidatedVehicleCoverageTermID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *VehicleCoverageTermID) MarshalJSON() ([]byte, error) {
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

func (i *VehicleCoverageTermID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := VehicleCoverageTermID(dataString)
	item := VehicleCoverageTerm.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	vehicleCoverageTermComprehensiveDeductibleID     VehicleCoverageTermID = "PACompDeductible"
	vehicleCoverageTermCollisionDeductibleID         VehicleCoverageTermID = "PACollDeductible"
	vehicleCoverageTermUMPDLimitsIllinoisID          VehicleCoverageTermID = "PAUMPD_IL"
	vehicleCoverageTermUMPDLimitsOhioID              VehicleCoverageTermID = "PAUMPD_OH"
	vehicleCoverageTermTowingAndLaborLimitVirginiaID VehicleCoverageTermID = "TowingAndLaborLimit"
	vehicleCoverageTermRentalLimitsID                VehicleCoverageTermID = "PARental"
	vehicleCoverageTermCustomEquipmentLimitsID       VehicleCoverageTermID = "PACustEquip"
	vehicleCoverageTermVehicleMonitorID              VehicleCoverageTermID = "PAVMONNewDevice"
)

// EnumVehicleCoverageTermItem describes an entry in an enumeration of VehicleCoverageTerm
type EnumVehicleCoverageTermItem struct {
	ID        VehicleCoverageTermID `json:"Value"`
	Desc      string                `json:"Description,omitempty"`
	Meta      map[string]string     `json:",omitempty"`
	Name      string                `json:"Name"`
	SortOrder int
}

var (
	vehicleCoverageTermComprehensiveDeductible     = EnumVehicleCoverageTermItem{vehicleCoverageTermComprehensiveDeductibleID, "Comprehensive Deductible", nil, "ComprehensiveDeductible", 1}
	vehicleCoverageTermCollisionDeductible         = EnumVehicleCoverageTermItem{vehicleCoverageTermCollisionDeductibleID, "Collision Deductible", nil, "CollisionDeductible", 2}
	vehicleCoverageTermUMPDLimitsIllinois          = EnumVehicleCoverageTermItem{vehicleCoverageTermUMPDLimitsIllinoisID, "Illinois Uninsured Motorist Property Damage Limits", nil, "UMPDLimitsIllinois", 3}
	vehicleCoverageTermUMPDLimitsOhio              = EnumVehicleCoverageTermItem{vehicleCoverageTermUMPDLimitsOhioID, "Ohio Uninsured Motorist Property Damage Limits", nil, "UMPDLimitsOhio", 4}
	vehicleCoverageTermTowingAndLaborLimitVirginia = EnumVehicleCoverageTermItem{vehicleCoverageTermTowingAndLaborLimitVirginiaID, "Towing And Labor Limit", nil, "TowingAndLaborLimitVirginia", 5}
	vehicleCoverageTermRentalLimits                = EnumVehicleCoverageTermItem{vehicleCoverageTermRentalLimitsID, "Rental Reimbursement Limits", nil, "RentalLimits", 6}
	vehicleCoverageTermCustomEquipmentLimits       = EnumVehicleCoverageTermItem{vehicleCoverageTermCustomEquipmentLimitsID, "Custom Equipment Limits", nil, "CustomEquipmentLimits", 7}
	vehicleCoverageTermVehicleMonitor              = EnumVehicleCoverageTermItem{vehicleCoverageTermVehicleMonitorID, "Vehicle VehicleMonitor", nil, "VehicleMonitor", 8}
)

// EnumVehicleCoverageTerm is a collection of VehicleCoverageTerm items
type EnumVehicleCoverageTerm struct {
	Description string
	Items       []*EnumVehicleCoverageTermItem
	Name        string

	ComprehensiveDeductible     *EnumVehicleCoverageTermItem
	CollisionDeductible         *EnumVehicleCoverageTermItem
	UMPDLimitsIllinois          *EnumVehicleCoverageTermItem
	UMPDLimitsOhio              *EnumVehicleCoverageTermItem
	TowingAndLaborLimitVirginia *EnumVehicleCoverageTermItem
	RentalLimits                *EnumVehicleCoverageTermItem
	CustomEquipmentLimits       *EnumVehicleCoverageTermItem
	VehicleMonitor              *EnumVehicleCoverageTermItem

	itemDict map[string]*EnumVehicleCoverageTermItem
}

// VehicleCoverageTerm is a public singleton instance of EnumVehicleCoverageTerm
// representing term codes for vehicle coverages
var VehicleCoverageTerm = &EnumVehicleCoverageTerm{
	Description: "term codes for vehicle coverages",
	Items: []*EnumVehicleCoverageTermItem{
		&vehicleCoverageTermComprehensiveDeductible,
		&vehicleCoverageTermCollisionDeductible,
		&vehicleCoverageTermUMPDLimitsIllinois,
		&vehicleCoverageTermUMPDLimitsOhio,
		&vehicleCoverageTermTowingAndLaborLimitVirginia,
		&vehicleCoverageTermRentalLimits,
		&vehicleCoverageTermCustomEquipmentLimits,
		&vehicleCoverageTermVehicleMonitor,
	},
	Name:                        "EnumVehicleCoverageTerm",
	ComprehensiveDeductible:     &vehicleCoverageTermComprehensiveDeductible,
	CollisionDeductible:         &vehicleCoverageTermCollisionDeductible,
	UMPDLimitsIllinois:          &vehicleCoverageTermUMPDLimitsIllinois,
	UMPDLimitsOhio:              &vehicleCoverageTermUMPDLimitsOhio,
	TowingAndLaborLimitVirginia: &vehicleCoverageTermTowingAndLaborLimitVirginia,
	RentalLimits:                &vehicleCoverageTermRentalLimits,
	CustomEquipmentLimits:       &vehicleCoverageTermCustomEquipmentLimits,
	VehicleMonitor:              &vehicleCoverageTermVehicleMonitor,

	itemDict: map[string]*EnumVehicleCoverageTermItem{
		strings.ToLower(string(vehicleCoverageTermComprehensiveDeductibleID)):     &vehicleCoverageTermComprehensiveDeductible,
		strings.ToLower(string(vehicleCoverageTermCollisionDeductibleID)):         &vehicleCoverageTermCollisionDeductible,
		strings.ToLower(string(vehicleCoverageTermUMPDLimitsIllinoisID)):          &vehicleCoverageTermUMPDLimitsIllinois,
		strings.ToLower(string(vehicleCoverageTermUMPDLimitsOhioID)):              &vehicleCoverageTermUMPDLimitsOhio,
		strings.ToLower(string(vehicleCoverageTermTowingAndLaborLimitVirginiaID)): &vehicleCoverageTermTowingAndLaborLimitVirginia,
		strings.ToLower(string(vehicleCoverageTermRentalLimitsID)):                &vehicleCoverageTermRentalLimits,
		strings.ToLower(string(vehicleCoverageTermCustomEquipmentLimitsID)):       &vehicleCoverageTermCustomEquipmentLimits,
		strings.ToLower(string(vehicleCoverageTermVehicleMonitorID)):              &vehicleCoverageTermVehicleMonitor,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumVehicleCoverageTerm) ByID(id VehicleCoverageTermIdentifier) *EnumVehicleCoverageTermItem {
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
func (e *EnumVehicleCoverageTerm) ByIDString(idx string) *EnumVehicleCoverageTermItem {
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
func (e *EnumVehicleCoverageTerm) ByIndex(idx int) *EnumVehicleCoverageTermItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedVehicleCoverageTermID is a struct that is designed to replace a *VehicleCoverageTermID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *VehicleCoverageTermID it contains while being a better JSON citizen.
type ValidatedVehicleCoverageTermID struct {
	// id will point to a valid VehicleCoverageTermID, if possible
	// If id is nil, then ValidatedVehicleCoverageTermID.Valid() will return false.
	id *VehicleCoverageTermID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedVehicleCoverageTermID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedVehicleCoverageTermID
func (vi *ValidatedVehicleCoverageTermID) Clone() *ValidatedVehicleCoverageTermID {
	if vi == nil {
		return nil
	}

	var cid *VehicleCoverageTermID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedVehicleCoverageTermID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedVehicleCoverageTermIds represent the same VehicleCoverageTerm
func (vi *ValidatedVehicleCoverageTermID) Equals(vj *ValidatedVehicleCoverageTermID) bool {
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

// Valid returns true if and only if the ValidatedVehicleCoverageTermID corresponds to a recognized VehicleCoverageTerm
func (vi *ValidatedVehicleCoverageTermID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedVehicleCoverageTermID) ID() *VehicleCoverageTermID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedVehicleCoverageTermID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedVehicleCoverageTermID) ValidatedID() *ValidatedVehicleCoverageTermID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedVehicleCoverageTermID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedVehicleCoverageTermID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedVehicleCoverageTermID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedVehicleCoverageTermID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedVehicleCoverageTermID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := VehicleCoverageTermID(capString)
	item := VehicleCoverageTerm.ByID(&id)
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

func (vi ValidatedVehicleCoverageTermID) String() string {
	return vi.ToIDString()
}

type VehicleCoverageTermIdentifier interface {
	ID() *VehicleCoverageTermID
	Valid() bool
}
