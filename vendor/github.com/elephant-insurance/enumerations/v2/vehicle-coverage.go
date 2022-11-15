package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// VehicleCoverageID uniquely identifies a particular VehicleCoverage
type VehicleCoverageID string

// Clone creates a safe, independent copy of a VehicleCoverageID
func (i *VehicleCoverageID) Clone() *VehicleCoverageID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two VehicleCoverageIds are equivalent
func (i *VehicleCoverageID) Equals(j *VehicleCoverageID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *VehicleCoverageID that is either valid or nil
func (i *VehicleCoverageID) ID() *VehicleCoverageID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *VehicleCoverageID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the VehicleCoverageID corresponds to a recognized VehicleCoverage
func (i *VehicleCoverageID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return VehicleCoverage.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *VehicleCoverageID) ValidatedID() *ValidatedVehicleCoverageID {
	if i != nil {
		return &ValidatedVehicleCoverageID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *VehicleCoverageID) MarshalJSON() ([]byte, error) {
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

func (i *VehicleCoverageID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := VehicleCoverageID(dataString)
	item := VehicleCoverage.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	vehicleCoverageVehicleMonitoringID                       VehicleCoverageID = "PAVehicleMonitorCov"
	vehicleCoverageComprehensiveID                           VehicleCoverageID = "PAComprehensiveCov"
	vehicleCoverageCollisionID                               VehicleCoverageID = "PACollisionCov"
	vehicleCoverageUninsuredMotoristPropertyDamageIllinoisID VehicleCoverageID = "EISPAUMPD_ILCov"
	vehicleCoverageUninsuredMotoristPropertyDamageOhioID     VehicleCoverageID = "EISPAUMPD_OHCov"
	vehicleCoverageTowingAndLaborID                          VehicleCoverageID = "PATowingLaborCov"
	vehicleCoverageRentalReimbursementID                     VehicleCoverageID = "PARentalCov"
	vehicleCoverageLoanGapID                                 VehicleCoverageID = "EISPALoanGapCov"
	vehicleCoverageCustomEquipmentID                         VehicleCoverageID = "EISPACustEquipCov"
)

// EnumVehicleCoverageItem describes an entry in an enumeration of VehicleCoverage
type EnumVehicleCoverageItem struct {
	ID        VehicleCoverageID `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	vehicleCoverageVehicleMonitoring                       = EnumVehicleCoverageItem{vehicleCoverageVehicleMonitoringID, "Teen Driver App", nil, "VehicleMonitoring", 1}
	vehicleCoverageComprehensive                           = EnumVehicleCoverageItem{vehicleCoverageComprehensiveID, "Other Than Collision (Comprehensive)", nil, "Comprehensive", 2}
	vehicleCoverageCollision                               = EnumVehicleCoverageItem{vehicleCoverageCollisionID, "Collision", nil, "Collision", 3}
	vehicleCoverageUninsuredMotoristPropertyDamageIllinois = EnumVehicleCoverageItem{vehicleCoverageUninsuredMotoristPropertyDamageIllinoisID, "Uninsured Motorist - Property Damage", nil, "UninsuredMotoristPropertyDamageIllinois", 4}
	vehicleCoverageUninsuredMotoristPropertyDamageOhio     = EnumVehicleCoverageItem{vehicleCoverageUninsuredMotoristPropertyDamageOhioID, "Uninsured Motorist - Property Damage", nil, "UninsuredMotoristPropertyDamageOhio", 5}
	vehicleCoverageTowingAndLabor                          = EnumVehicleCoverageItem{vehicleCoverageTowingAndLaborID, "Roadside Assistance", nil, "TowingAndLabor", 6}
	vehicleCoverageRentalReimbursement                     = EnumVehicleCoverageItem{vehicleCoverageRentalReimbursementID, "Rental Reimbursement", nil, "RentalReimbursement", 7}
	vehicleCoverageLoanGap                                 = EnumVehicleCoverageItem{vehicleCoverageLoanGapID, "Loan Lease Payoff", nil, "LoanGap", 8}
	vehicleCoverageCustomEquipment                         = EnumVehicleCoverageItem{vehicleCoverageCustomEquipmentID, "Custom Equipment Coverage", nil, "CustomEquipment", 9}
)

// EnumVehicleCoverage is a collection of VehicleCoverage items
type EnumVehicleCoverage struct {
	Description string
	Items       []*EnumVehicleCoverageItem
	Name        string

	VehicleMonitoring                       *EnumVehicleCoverageItem
	Comprehensive                           *EnumVehicleCoverageItem
	Collision                               *EnumVehicleCoverageItem
	UninsuredMotoristPropertyDamageIllinois *EnumVehicleCoverageItem
	UninsuredMotoristPropertyDamageOhio     *EnumVehicleCoverageItem
	TowingAndLabor                          *EnumVehicleCoverageItem
	RentalReimbursement                     *EnumVehicleCoverageItem
	LoanGap                                 *EnumVehicleCoverageItem
	CustomEquipment                         *EnumVehicleCoverageItem

	itemDict map[string]*EnumVehicleCoverageItem
}

// VehicleCoverage is a public singleton instance of EnumVehicleCoverage
// representing codes for vehicle coverages
var VehicleCoverage = &EnumVehicleCoverage{
	Description: "codes for vehicle coverages",
	Items: []*EnumVehicleCoverageItem{
		&vehicleCoverageVehicleMonitoring,
		&vehicleCoverageComprehensive,
		&vehicleCoverageCollision,
		&vehicleCoverageUninsuredMotoristPropertyDamageIllinois,
		&vehicleCoverageUninsuredMotoristPropertyDamageOhio,
		&vehicleCoverageTowingAndLabor,
		&vehicleCoverageRentalReimbursement,
		&vehicleCoverageLoanGap,
		&vehicleCoverageCustomEquipment,
	},
	Name:                                    "EnumVehicleCoverage",
	VehicleMonitoring:                       &vehicleCoverageVehicleMonitoring,
	Comprehensive:                           &vehicleCoverageComprehensive,
	Collision:                               &vehicleCoverageCollision,
	UninsuredMotoristPropertyDamageIllinois: &vehicleCoverageUninsuredMotoristPropertyDamageIllinois,
	UninsuredMotoristPropertyDamageOhio:     &vehicleCoverageUninsuredMotoristPropertyDamageOhio,
	TowingAndLabor:                          &vehicleCoverageTowingAndLabor,
	RentalReimbursement:                     &vehicleCoverageRentalReimbursement,
	LoanGap:                                 &vehicleCoverageLoanGap,
	CustomEquipment:                         &vehicleCoverageCustomEquipment,

	itemDict: map[string]*EnumVehicleCoverageItem{
		strings.ToLower(string(vehicleCoverageVehicleMonitoringID)):                       &vehicleCoverageVehicleMonitoring,
		strings.ToLower(string(vehicleCoverageComprehensiveID)):                           &vehicleCoverageComprehensive,
		strings.ToLower(string(vehicleCoverageCollisionID)):                               &vehicleCoverageCollision,
		strings.ToLower(string(vehicleCoverageUninsuredMotoristPropertyDamageIllinoisID)): &vehicleCoverageUninsuredMotoristPropertyDamageIllinois,
		strings.ToLower(string(vehicleCoverageUninsuredMotoristPropertyDamageOhioID)):     &vehicleCoverageUninsuredMotoristPropertyDamageOhio,
		strings.ToLower(string(vehicleCoverageTowingAndLaborID)):                          &vehicleCoverageTowingAndLabor,
		strings.ToLower(string(vehicleCoverageRentalReimbursementID)):                     &vehicleCoverageRentalReimbursement,
		strings.ToLower(string(vehicleCoverageLoanGapID)):                                 &vehicleCoverageLoanGap,
		strings.ToLower(string(vehicleCoverageCustomEquipmentID)):                         &vehicleCoverageCustomEquipment,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumVehicleCoverage) ByID(id VehicleCoverageIdentifier) *EnumVehicleCoverageItem {
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
func (e *EnumVehicleCoverage) ByIDString(idx string) *EnumVehicleCoverageItem {
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
func (e *EnumVehicleCoverage) ByIndex(idx int) *EnumVehicleCoverageItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedVehicleCoverageID is a struct that is designed to replace a *VehicleCoverageID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *VehicleCoverageID it contains while being a better JSON citizen.
type ValidatedVehicleCoverageID struct {
	// id will point to a valid VehicleCoverageID, if possible
	// If id is nil, then ValidatedVehicleCoverageID.Valid() will return false.
	id *VehicleCoverageID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedVehicleCoverageID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedVehicleCoverageID
func (vi *ValidatedVehicleCoverageID) Clone() *ValidatedVehicleCoverageID {
	if vi == nil {
		return nil
	}

	var cid *VehicleCoverageID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedVehicleCoverageID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedVehicleCoverageIds represent the same VehicleCoverage
func (vi *ValidatedVehicleCoverageID) Equals(vj *ValidatedVehicleCoverageID) bool {
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

// Valid returns true if and only if the ValidatedVehicleCoverageID corresponds to a recognized VehicleCoverage
func (vi *ValidatedVehicleCoverageID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedVehicleCoverageID) ID() *VehicleCoverageID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedVehicleCoverageID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedVehicleCoverageID) ValidatedID() *ValidatedVehicleCoverageID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedVehicleCoverageID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedVehicleCoverageID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedVehicleCoverageID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedVehicleCoverageID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedVehicleCoverageID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := VehicleCoverageID(capString)
	item := VehicleCoverage.ByID(&id)
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

func (vi ValidatedVehicleCoverageID) String() string {
	return vi.ToIDString()
}

type VehicleCoverageIdentifier interface {
	ID() *VehicleCoverageID
	Valid() bool
}
