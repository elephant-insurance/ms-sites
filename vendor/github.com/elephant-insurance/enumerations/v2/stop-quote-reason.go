package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// StopQuoteReasonID uniquely identifies a particular StopQuoteReason
type StopQuoteReasonID string

// Clone creates a safe, independent copy of a StopQuoteReasonID
func (i *StopQuoteReasonID) Clone() *StopQuoteReasonID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two StopQuoteReasonIds are equivalent
func (i *StopQuoteReasonID) Equals(j *StopQuoteReasonID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *StopQuoteReasonID that is either valid or nil
func (i *StopQuoteReasonID) ID() *StopQuoteReasonID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *StopQuoteReasonID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the StopQuoteReasonID corresponds to a recognized StopQuoteReason
func (i *StopQuoteReasonID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return StopQuoteReason.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *StopQuoteReasonID) ValidatedID() *ValidatedStopQuoteReasonID {
	if i != nil {
		return &ValidatedStopQuoteReasonID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *StopQuoteReasonID) MarshalJSON() ([]byte, error) {
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

func (i *StopQuoteReasonID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := StopQuoteReasonID(dataString)
	item := StopQuoteReason.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	stopQuoteReasonBadDebtsFlagID                   StopQuoteReasonID = "1011"
	stopQuoteReasonUnlicensedAdditionalDriverID     StopQuoteReasonID = "1234"
	stopQuoteReasonNeedUWPhotoReviewID              StopQuoteReasonID = "1246"
	stopQuoteReasonBlockLicenseFromStateID          StopQuoteReasonID = "1247"
	stopQuoteReasonMaterialMisrepID                 StopQuoteReasonID = "1348"
	stopQuoteReasonIncidentViolationsID             StopQuoteReasonID = "4018"
	stopQuoteReasonExcludedDriverID                 StopQuoteReasonID = "2222"
	stopQuoteReasonOtherBlockID                     StopQuoteReasonID = "0"
	stopQuoteReasonHasSR22ID                        StopQuoteReasonID = "2255"
	stopQuoteReasonHasFR44ID                        StopQuoteReasonID = "2244"
	stopQuoteReasonBrandedVehicleID                 StopQuoteReasonID = "5024"
	stopQuoteReasonVehicleOwnershipID               StopQuoteReasonID = "6226"
	stopQuoteReasonHasBadVinID                      StopQuoteReasonID = "9019"
	stopQuoteReasonRestrictLicenseStatusSuspendedID StopQuoteReasonID = "9030"
	stopQuoteReasonRestrictLicenseStatusID          StopQuoteReasonID = "9040"
	stopQuoteReasonHasActivePolicyID                StopQuoteReasonID = "9031"
	stopQuoteReasonNotRatedLocationID               StopQuoteReasonID = "9999"
	stopQuoteReasonTerritoryRestrictionID           StopQuoteReasonID = "9018"
	stopQuoteReasonExceedLimitOfIncidentID          StopQuoteReasonID = "8007"
	stopQuoteReasonHighCollisionDeductibleID        StopQuoteReasonID = "7005"
	stopQuoteReasonSameVINID                        StopQuoteReasonID = "5018"
	stopQuoteReasonMissingPrimaryDriverID           StopQuoteReasonID = "5011"
	stopQuoteReasonInvalidLicenseStatusID           StopQuoteReasonID = "4008"
	stopQuoteReasonLicenseNotAllowedID              StopQuoteReasonID = "4002"
	stopQuoteReasonJobNumberExpiredID               StopQuoteReasonID = "3102"
	stopQuoteReasonRatabaseErrorID                  StopQuoteReasonID = "10002"
	stopQuoteReasonServicesDownID                   StopQuoteReasonID = "4000"
)

// EnumStopQuoteReasonItem describes an entry in an enumeration of StopQuoteReason
type EnumStopQuoteReasonItem struct {
	ID        StopQuoteReasonID `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	stopQuoteReasonBadDebtsFlag                   = EnumStopQuoteReasonItem{stopQuoteReasonBadDebtsFlagID, "Bad debts flag on account", nil, "BadDebtsFlag", 1}
	stopQuoteReasonUnlicensedAdditionalDriver     = EnumStopQuoteReasonItem{stopQuoteReasonUnlicensedAdditionalDriverID, "Unlicensed additional driver", nil, "UnlicensedAdditionalDriver", 2}
	stopQuoteReasonNeedUWPhotoReview              = EnumStopQuoteReasonItem{stopQuoteReasonNeedUWPhotoReviewID, "Need UWPhoto review", nil, "NeedUWPhotoReview", 3}
	stopQuoteReasonBlockLicenseFromState          = EnumStopQuoteReasonItem{stopQuoteReasonBlockLicenseFromStateID, "Block license from state", nil, "BlockLicenseFromState", 4}
	stopQuoteReasonMaterialMisrep                 = EnumStopQuoteReasonItem{stopQuoteReasonMaterialMisrepID, "Material Misrep flag", nil, "MaterialMisrep", 5}
	stopQuoteReasonIncidentViolations             = EnumStopQuoteReasonItem{stopQuoteReasonIncidentViolationsID, "Incident violations", nil, "IncidentViolations", 6}
	stopQuoteReasonExcludedDriver                 = EnumStopQuoteReasonItem{stopQuoteReasonExcludedDriverID, "Exclude driver not allowed", nil, "ExcludedDriver", 7}
	stopQuoteReasonOtherBlock                     = EnumStopQuoteReasonItem{stopQuoteReasonOtherBlockID, "Other block or unknown", nil, "OtherBlock", 8}
	stopQuoteReasonHasSR22                        = EnumStopQuoteReasonItem{stopQuoteReasonHasSR22ID, "Has SR 22 flag", nil, "HasSR22", 9}
	stopQuoteReasonHasFR44                        = EnumStopQuoteReasonItem{stopQuoteReasonHasFR44ID, "Has FR 44 flag", nil, "HasFR44", 10}
	stopQuoteReasonBrandedVehicle                 = EnumStopQuoteReasonItem{stopQuoteReasonBrandedVehicleID, "Has Branded vehicle", nil, "BrandedVehicle", 11}
	stopQuoteReasonVehicleOwnership               = EnumStopQuoteReasonItem{stopQuoteReasonVehicleOwnershipID, "Insurable interest", nil, "VehicleOwnership", 12}
	stopQuoteReasonHasBadVin                      = EnumStopQuoteReasonItem{stopQuoteReasonHasBadVinID, "Bad VIN", nil, "HasBadVin", 13}
	stopQuoteReasonRestrictLicenseStatusSuspended = EnumStopQuoteReasonItem{stopQuoteReasonRestrictLicenseStatusSuspendedID, "Suspended license", nil, "RestrictLicenseStatusSuspended", 14}
	stopQuoteReasonRestrictLicenseStatus          = EnumStopQuoteReasonItem{stopQuoteReasonRestrictLicenseStatusID, "Restricted license status", nil, "RestrictLicenseStatus", 15}
	stopQuoteReasonHasActivePolicy                = EnumStopQuoteReasonItem{stopQuoteReasonHasActivePolicyID, "Has active policy", nil, "HasActivePolicy", 16}
	stopQuoteReasonNotRatedLocation               = EnumStopQuoteReasonItem{stopQuoteReasonNotRatedLocationID, "Not rated location", nil, "NotRatedLocation", 17}
	stopQuoteReasonTerritoryRestriction           = EnumStopQuoteReasonItem{stopQuoteReasonTerritoryRestrictionID, "Binding restriction", nil, "TerritoryRestriction", 18}
	stopQuoteReasonExceedLimitOfIncident          = EnumStopQuoteReasonItem{stopQuoteReasonExceedLimitOfIncidentID, "Exceed the Limit of the Incident", nil, "ExceedLimitOfIncident", 19}
	stopQuoteReasonHighCollisionDeductible        = EnumStopQuoteReasonItem{stopQuoteReasonHighCollisionDeductibleID, "Veh has a Collision Deductible > $500.00", nil, "HighCollisionDeductible", 20}
	stopQuoteReasonSameVIN                        = EnumStopQuoteReasonItem{stopQuoteReasonSameVINID, "Multiple Veh have Same VIN", nil, "SameVIN", 21}
	stopQuoteReasonMissingPrimaryDriver           = EnumStopQuoteReasonItem{stopQuoteReasonMissingPrimaryDriverID, "All Veh Require Primary Driver", nil, "MissingPrimaryDriver", 22}
	stopQuoteReasonInvalidLicenseStatus           = EnumStopQuoteReasonItem{stopQuoteReasonInvalidLicenseStatusID, "Invalid License Status", nil, "InvalidLicenseStatus", 23}
	stopQuoteReasonLicenseNotAllowed              = EnumStopQuoteReasonItem{stopQuoteReasonLicenseNotAllowedID, "License not Allowed", nil, "LicenseNotAllowed", 24}
	stopQuoteReasonJobNumberExpired               = EnumStopQuoteReasonItem{stopQuoteReasonJobNumberExpiredID, "Job number expired", nil, "JobNumberExpired", 25}
	stopQuoteReasonRatabaseError                  = EnumStopQuoteReasonItem{stopQuoteReasonRatabaseErrorID, "Ratabase Error", nil, "RatabaseError", 26}
	stopQuoteReasonServicesDown                   = EnumStopQuoteReasonItem{stopQuoteReasonServicesDownID, "Services Down", nil, "ServicesDown", 27}
)

// EnumStopQuoteReason is a collection of StopQuoteReason items
type EnumStopQuoteReason struct {
	Description string
	Items       []*EnumStopQuoteReasonItem
	Name        string

	BadDebtsFlag                   *EnumStopQuoteReasonItem
	UnlicensedAdditionalDriver     *EnumStopQuoteReasonItem
	NeedUWPhotoReview              *EnumStopQuoteReasonItem
	BlockLicenseFromState          *EnumStopQuoteReasonItem
	MaterialMisrep                 *EnumStopQuoteReasonItem
	IncidentViolations             *EnumStopQuoteReasonItem
	ExcludedDriver                 *EnumStopQuoteReasonItem
	OtherBlock                     *EnumStopQuoteReasonItem
	HasSR22                        *EnumStopQuoteReasonItem
	HasFR44                        *EnumStopQuoteReasonItem
	BrandedVehicle                 *EnumStopQuoteReasonItem
	VehicleOwnership               *EnumStopQuoteReasonItem
	HasBadVin                      *EnumStopQuoteReasonItem
	RestrictLicenseStatusSuspended *EnumStopQuoteReasonItem
	RestrictLicenseStatus          *EnumStopQuoteReasonItem
	HasActivePolicy                *EnumStopQuoteReasonItem
	NotRatedLocation               *EnumStopQuoteReasonItem
	TerritoryRestriction           *EnumStopQuoteReasonItem
	ExceedLimitOfIncident          *EnumStopQuoteReasonItem
	HighCollisionDeductible        *EnumStopQuoteReasonItem
	SameVIN                        *EnumStopQuoteReasonItem
	MissingPrimaryDriver           *EnumStopQuoteReasonItem
	InvalidLicenseStatus           *EnumStopQuoteReasonItem
	LicenseNotAllowed              *EnumStopQuoteReasonItem
	JobNumberExpired               *EnumStopQuoteReasonItem
	RatabaseError                  *EnumStopQuoteReasonItem
	ServicesDown                   *EnumStopQuoteReasonItem

	itemDict map[string]*EnumStopQuoteReasonItem
}

// StopQuoteReason is a public singleton instance of EnumStopQuoteReason
// representing Reasons for stopping an online quote
var StopQuoteReason = &EnumStopQuoteReason{
	Description: "Reasons for stopping an online quote",
	Items: []*EnumStopQuoteReasonItem{
		&stopQuoteReasonBadDebtsFlag,
		&stopQuoteReasonUnlicensedAdditionalDriver,
		&stopQuoteReasonNeedUWPhotoReview,
		&stopQuoteReasonBlockLicenseFromState,
		&stopQuoteReasonMaterialMisrep,
		&stopQuoteReasonIncidentViolations,
		&stopQuoteReasonExcludedDriver,
		&stopQuoteReasonOtherBlock,
		&stopQuoteReasonHasSR22,
		&stopQuoteReasonHasFR44,
		&stopQuoteReasonBrandedVehicle,
		&stopQuoteReasonVehicleOwnership,
		&stopQuoteReasonHasBadVin,
		&stopQuoteReasonRestrictLicenseStatusSuspended,
		&stopQuoteReasonRestrictLicenseStatus,
		&stopQuoteReasonHasActivePolicy,
		&stopQuoteReasonNotRatedLocation,
		&stopQuoteReasonTerritoryRestriction,
		&stopQuoteReasonExceedLimitOfIncident,
		&stopQuoteReasonHighCollisionDeductible,
		&stopQuoteReasonSameVIN,
		&stopQuoteReasonMissingPrimaryDriver,
		&stopQuoteReasonInvalidLicenseStatus,
		&stopQuoteReasonLicenseNotAllowed,
		&stopQuoteReasonJobNumberExpired,
		&stopQuoteReasonRatabaseError,
		&stopQuoteReasonServicesDown,
	},
	Name:                           "EnumStopQuoteReason",
	BadDebtsFlag:                   &stopQuoteReasonBadDebtsFlag,
	UnlicensedAdditionalDriver:     &stopQuoteReasonUnlicensedAdditionalDriver,
	NeedUWPhotoReview:              &stopQuoteReasonNeedUWPhotoReview,
	BlockLicenseFromState:          &stopQuoteReasonBlockLicenseFromState,
	MaterialMisrep:                 &stopQuoteReasonMaterialMisrep,
	IncidentViolations:             &stopQuoteReasonIncidentViolations,
	ExcludedDriver:                 &stopQuoteReasonExcludedDriver,
	OtherBlock:                     &stopQuoteReasonOtherBlock,
	HasSR22:                        &stopQuoteReasonHasSR22,
	HasFR44:                        &stopQuoteReasonHasFR44,
	BrandedVehicle:                 &stopQuoteReasonBrandedVehicle,
	VehicleOwnership:               &stopQuoteReasonVehicleOwnership,
	HasBadVin:                      &stopQuoteReasonHasBadVin,
	RestrictLicenseStatusSuspended: &stopQuoteReasonRestrictLicenseStatusSuspended,
	RestrictLicenseStatus:          &stopQuoteReasonRestrictLicenseStatus,
	HasActivePolicy:                &stopQuoteReasonHasActivePolicy,
	NotRatedLocation:               &stopQuoteReasonNotRatedLocation,
	TerritoryRestriction:           &stopQuoteReasonTerritoryRestriction,
	ExceedLimitOfIncident:          &stopQuoteReasonExceedLimitOfIncident,
	HighCollisionDeductible:        &stopQuoteReasonHighCollisionDeductible,
	SameVIN:                        &stopQuoteReasonSameVIN,
	MissingPrimaryDriver:           &stopQuoteReasonMissingPrimaryDriver,
	InvalidLicenseStatus:           &stopQuoteReasonInvalidLicenseStatus,
	LicenseNotAllowed:              &stopQuoteReasonLicenseNotAllowed,
	JobNumberExpired:               &stopQuoteReasonJobNumberExpired,
	RatabaseError:                  &stopQuoteReasonRatabaseError,
	ServicesDown:                   &stopQuoteReasonServicesDown,

	itemDict: map[string]*EnumStopQuoteReasonItem{
		strings.ToLower(string(stopQuoteReasonBadDebtsFlagID)):                   &stopQuoteReasonBadDebtsFlag,
		strings.ToLower(string(stopQuoteReasonUnlicensedAdditionalDriverID)):     &stopQuoteReasonUnlicensedAdditionalDriver,
		strings.ToLower(string(stopQuoteReasonNeedUWPhotoReviewID)):              &stopQuoteReasonNeedUWPhotoReview,
		strings.ToLower(string(stopQuoteReasonBlockLicenseFromStateID)):          &stopQuoteReasonBlockLicenseFromState,
		strings.ToLower(string(stopQuoteReasonMaterialMisrepID)):                 &stopQuoteReasonMaterialMisrep,
		strings.ToLower(string(stopQuoteReasonIncidentViolationsID)):             &stopQuoteReasonIncidentViolations,
		strings.ToLower(string(stopQuoteReasonExcludedDriverID)):                 &stopQuoteReasonExcludedDriver,
		strings.ToLower(string(stopQuoteReasonOtherBlockID)):                     &stopQuoteReasonOtherBlock,
		strings.ToLower(string(stopQuoteReasonHasSR22ID)):                        &stopQuoteReasonHasSR22,
		strings.ToLower(string(stopQuoteReasonHasFR44ID)):                        &stopQuoteReasonHasFR44,
		strings.ToLower(string(stopQuoteReasonBrandedVehicleID)):                 &stopQuoteReasonBrandedVehicle,
		strings.ToLower(string(stopQuoteReasonVehicleOwnershipID)):               &stopQuoteReasonVehicleOwnership,
		strings.ToLower(string(stopQuoteReasonHasBadVinID)):                      &stopQuoteReasonHasBadVin,
		strings.ToLower(string(stopQuoteReasonRestrictLicenseStatusSuspendedID)): &stopQuoteReasonRestrictLicenseStatusSuspended,
		strings.ToLower(string(stopQuoteReasonRestrictLicenseStatusID)):          &stopQuoteReasonRestrictLicenseStatus,
		strings.ToLower(string(stopQuoteReasonHasActivePolicyID)):                &stopQuoteReasonHasActivePolicy,
		strings.ToLower(string(stopQuoteReasonNotRatedLocationID)):               &stopQuoteReasonNotRatedLocation,
		strings.ToLower(string(stopQuoteReasonTerritoryRestrictionID)):           &stopQuoteReasonTerritoryRestriction,
		strings.ToLower(string(stopQuoteReasonExceedLimitOfIncidentID)):          &stopQuoteReasonExceedLimitOfIncident,
		strings.ToLower(string(stopQuoteReasonHighCollisionDeductibleID)):        &stopQuoteReasonHighCollisionDeductible,
		strings.ToLower(string(stopQuoteReasonSameVINID)):                        &stopQuoteReasonSameVIN,
		strings.ToLower(string(stopQuoteReasonMissingPrimaryDriverID)):           &stopQuoteReasonMissingPrimaryDriver,
		strings.ToLower(string(stopQuoteReasonInvalidLicenseStatusID)):           &stopQuoteReasonInvalidLicenseStatus,
		strings.ToLower(string(stopQuoteReasonLicenseNotAllowedID)):              &stopQuoteReasonLicenseNotAllowed,
		strings.ToLower(string(stopQuoteReasonJobNumberExpiredID)):               &stopQuoteReasonJobNumberExpired,
		strings.ToLower(string(stopQuoteReasonRatabaseErrorID)):                  &stopQuoteReasonRatabaseError,
		strings.ToLower(string(stopQuoteReasonServicesDownID)):                   &stopQuoteReasonServicesDown,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumStopQuoteReason) ByID(id StopQuoteReasonIdentifier) *EnumStopQuoteReasonItem {
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
func (e *EnumStopQuoteReason) ByIDString(idx string) *EnumStopQuoteReasonItem {
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
func (e *EnumStopQuoteReason) ByIndex(idx int) *EnumStopQuoteReasonItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedStopQuoteReasonID is a struct that is designed to replace a *StopQuoteReasonID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *StopQuoteReasonID it contains while being a better JSON citizen.
type ValidatedStopQuoteReasonID struct {
	// id will point to a valid StopQuoteReasonID, if possible
	// If id is nil, then ValidatedStopQuoteReasonID.Valid() will return false.
	id *StopQuoteReasonID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedStopQuoteReasonID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedStopQuoteReasonID
func (vi *ValidatedStopQuoteReasonID) Clone() *ValidatedStopQuoteReasonID {
	if vi == nil {
		return nil
	}

	var cid *StopQuoteReasonID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedStopQuoteReasonID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedStopQuoteReasonIds represent the same StopQuoteReason
func (vi *ValidatedStopQuoteReasonID) Equals(vj *ValidatedStopQuoteReasonID) bool {
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

// Valid returns true if and only if the ValidatedStopQuoteReasonID corresponds to a recognized StopQuoteReason
func (vi *ValidatedStopQuoteReasonID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedStopQuoteReasonID) ID() *StopQuoteReasonID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedStopQuoteReasonID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedStopQuoteReasonID) ValidatedID() *ValidatedStopQuoteReasonID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedStopQuoteReasonID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedStopQuoteReasonID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedStopQuoteReasonID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedStopQuoteReasonID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedStopQuoteReasonID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := StopQuoteReasonID(capString)
	item := StopQuoteReason.ByID(&id)
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

func (vi ValidatedStopQuoteReasonID) String() string {
	return vi.ToIDString()
}

type StopQuoteReasonIdentifier interface {
	ID() *StopQuoteReasonID
	Valid() bool
}
