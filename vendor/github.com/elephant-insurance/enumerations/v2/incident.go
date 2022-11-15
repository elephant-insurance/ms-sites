package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// IncidentID uniquely identifies a particular Incident
type IncidentID string

// Clone creates a safe, independent copy of a IncidentID
func (i *IncidentID) Clone() *IncidentID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two IncidentIds are equivalent
func (i *IncidentID) Equals(j *IncidentID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *IncidentID that is either valid or nil
func (i *IncidentID) ID() *IncidentID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *IncidentID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the IncidentID corresponds to a recognized Incident
func (i *IncidentID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return Incident.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *IncidentID) ValidatedID() *ValidatedIncidentID {
	if i != nil {
		return &ValidatedIncidentID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *IncidentID) MarshalJSON() ([]byte, error) {
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

func (i *IncidentID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := IncidentID(dataString)
	item := Incident.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	incidentAccidentAtFaultID                IncidentID = "AtFaultAccident"
	incidentAccidentNotAtFaultID             IncidentID = "NonChargeableAccident"
	incidentAdministrativeNoteID             IncidentID = "AdministrativeNote"
	incidentDrivingOnRevokedRegistrationID   IncidentID = "DrivingOnARevokedOrSuspendedLicense"
	incidentDUIDWIID                         IncidentID = "DrivingUnderTheInfluence"
	incidentEquipmentCargoLitteringID        IncidentID = "EquipmentCargoLittering"
	incidentExpiredOrImproperLicensePlatesID IncidentID = "ExpiredOrImproperLicensePlates"
	incidentFailureToObeyOfficerID           IncidentID = "UseofRadarFailureToObeyOfficer"
	incidentFailureToStopID                  IncidentID = "FailureToStopFollowingAccident"
	incidentGlassID                          IncidentID = "glass"
	incidentHomicideOrManslaughterID         IncidentID = "HomicideOrManslaughter"
	incidentImproperDrivingID                IncidentID = "ImproperDriving"
	incidentIncidentOTCClaimID               IncidentID = "Comprehensive"
	incidentLicenseSuspensionID              IncidentID = "LicenseSuspension"
	incidentMinorCoverageClaimID             IncidentID = "MinorCoverageClaim"
	incidentMinorTrafficViolationID          IncidentID = "FailureToYieldStopSignal"
	incidentNoPayoutClaimID                  IncidentID = "NoPayoutClaim"
	incidentOtherMajorViolationID            IncidentID = "OtherMajorViolations"
	incidentOutOfStateConvictionID           IncidentID = "OutofStateConviction"
	incidentPassingSchoolBusID               IncidentID = "PassedStopppedSchoolBus"
	incidentPoliceReportFiledID              IncidentID = "MVRAccidentDescription"
	incidentRacingID                         IncidentID = "Racing"
	incidentRecklessDrivingID                IncidentID = "RecklessOrNegligentDriving"
	incidentRentalID                         IncidentID = "rental"
	incidentRoadsideID                       IncidentID = "ers"
	incidentSeatBeltOrCellPhoneViolationID   IncidentID = "SeatBeltCellPhoneViolation"
	incidentSmallComprehensiveClaimID        IncidentID = "SmallComprehensiveClaim"
	incidentSpeeding30OrMoreID               IncidentID = "Speed30MphOrMore"
	incidentSpeedLessThan30OverID            IncidentID = "SpeedLessThan30Over"
	incidentTheftOfVehicleID                 IncidentID = "CriminalUseTheftofVehicle"
)

// EnumIncidentItem describes an entry in an enumeration of Incident
type EnumIncidentItem struct {
	ID        IncidentID        `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	Category       string
	Classification string
}

var (
	incidentAccidentAtFault                = EnumIncidentItem{incidentAccidentAtFaultID, "Accident – At Fault", map[string]string{IncidentMetaCategoryKey: IncidentCategory.AccidentOrClaim, IncidentMetaClassificationKey: IncidentClass.AtFault}, "AccidentAtFault", 1, IncidentCategory.AccidentOrClaim, IncidentClass.AtFault}
	incidentAccidentNotAtFault             = EnumIncidentItem{incidentAccidentNotAtFaultID, "Accident – Not At Fault", map[string]string{IncidentMetaCategoryKey: IncidentCategory.AccidentOrClaim, IncidentMetaClassificationKey: IncidentClass.NonChargeableAccident}, "AccidentNotAtFault", 2, IncidentCategory.AccidentOrClaim, IncidentClass.NonChargeableAccident}
	incidentAdministrativeNote             = EnumIncidentItem{incidentAdministrativeNoteID, "Administrative Note", map[string]string{IncidentMetaCategoryKey: "Other Violations", IncidentMetaClassificationKey: "NON"}, "AdministrativeNote", 3, "Other Violations", "NON"}
	incidentDrivingOnRevokedRegistration   = EnumIncidentItem{incidentDrivingOnRevokedRegistrationID, "Driving on Revoked or Suspended License", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.Major}, "DrivingOnRevokedRegistration", 4, IncidentCategory.MajorViolation, IncidentClass.Major}
	incidentDUIDWI                         = EnumIncidentItem{incidentDUIDWIID, "DUI / DWI", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.DUI}, "DUIDWI", 5, IncidentCategory.MajorViolation, IncidentClass.DUI}
	incidentEquipmentCargoLittering        = EnumIncidentItem{incidentEquipmentCargoLitteringID, "Equipment, Cargo or Littering Violation", map[string]string{IncidentMetaCategoryKey: "Other Violations", IncidentMetaClassificationKey: IncidentClass.NonChargeableConviction}, "EquipmentCargoLittering", 6, "Other Violations", IncidentClass.NonChargeableConviction}
	incidentExpiredOrImproperLicensePlates = EnumIncidentItem{incidentExpiredOrImproperLicensePlatesID, "Expired or Improper License, Registration or Plates", map[string]string{IncidentMetaCategoryKey: "Other Violations", IncidentMetaClassificationKey: IncidentClass.NonChargeableConviction}, "ExpiredOrImproperLicensePlates", 7, "Other Violations", IncidentClass.NonChargeableConviction}
	incidentFailureToObeyOfficer           = EnumIncidentItem{incidentFailureToObeyOfficerID, "Failure to Obey Officer / User of Radar", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.Major}, "FailureToObeyOfficer", 8, IncidentCategory.MajorViolation, IncidentClass.Major}
	incidentFailureToStop                  = EnumIncidentItem{incidentFailureToStopID, "Failure to Stop After Accident", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.Major}, "FailureToStop", 9, IncidentCategory.MajorViolation, IncidentClass.Major}
	incidentGlass                          = EnumIncidentItem{incidentGlassID, "Glass Claim", map[string]string{IncidentMetaCategoryKey: IncidentCategory.AccidentOrClaim, IncidentMetaClassificationKey: IncidentClass.NonChargeableAccident}, "Glass", 10, IncidentCategory.AccidentOrClaim, IncidentClass.NonChargeableAccident}
	incidentHomicideOrManslaughter         = EnumIncidentItem{incidentHomicideOrManslaughterID, "Homicide or Manslaughter", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.Major}, "HomicideOrManslaughter", 11, IncidentCategory.MajorViolation, IncidentClass.Major}
	incidentImproperDriving                = EnumIncidentItem{incidentImproperDrivingID, "Improper Driving / Other Minor Moving Violation", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MinorViolation, IncidentMetaClassificationKey: IncidentClass.Minor}, "ImproperDriving", 12, IncidentCategory.MinorViolation, IncidentClass.Minor}
	incidentIncidentOTCClaim               = EnumIncidentItem{incidentIncidentOTCClaimID, "Comprehensive (OTC Claim)", map[string]string{IncidentMetaCategoryKey: IncidentCategory.AccidentOrClaim, IncidentMetaClassificationKey: "NON"}, "IncidentOTCClaim", 13, IncidentCategory.AccidentOrClaim, "NON"}
	incidentLicenseSuspension              = EnumIncidentItem{incidentLicenseSuspensionID, "License Suspension", map[string]string{IncidentMetaCategoryKey: "Other Violations", IncidentMetaClassificationKey: "NON:"}, "LicenseSuspension", 14, "Other Violations", "NON:"}
	incidentMinorCoverageClaim             = EnumIncidentItem{incidentMinorCoverageClaimID, "Minor Coverage Claim", map[string]string{IncidentMetaCategoryKey: IncidentCategory.AccidentOrClaim, IncidentMetaClassificationKey: "NON"}, "MinorCoverageClaim", 15, IncidentCategory.AccidentOrClaim, "NON"}
	incidentMinorTrafficViolation          = EnumIncidentItem{incidentMinorTrafficViolationID, "Failure to Yield / Stop / Signal", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MinorViolation, IncidentMetaClassificationKey: IncidentClass.Minor}, "MinorTrafficViolation", 16, IncidentCategory.MinorViolation, IncidentClass.Minor}
	incidentNoPayoutClaim                  = EnumIncidentItem{incidentNoPayoutClaimID, "No Payout Claim", map[string]string{IncidentMetaCategoryKey: IncidentCategory.AccidentOrClaim, IncidentMetaClassificationKey: "NON"}, "NoPayoutClaim", 17, IncidentCategory.AccidentOrClaim, "NON"}
	incidentOtherMajorViolation            = EnumIncidentItem{incidentOtherMajorViolationID, "Other Major Violation", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.Major}, "OtherMajorViolation", 18, IncidentCategory.MajorViolation, IncidentClass.Major}
	incidentOutOfStateConviction           = EnumIncidentItem{incidentOutOfStateConvictionID, "Out of State Conviction", map[string]string{IncidentMetaCategoryKey: "Other Violations", IncidentMetaClassificationKey: "NON"}, "OutOfStateConviction", 19, "Other Violations", "NON"}
	incidentPassingSchoolBus               = EnumIncidentItem{incidentPassingSchoolBusID, "Passing Stopped School Bus", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.Major}, "PassingSchoolBus", 20, IncidentCategory.MajorViolation, IncidentClass.Major}
	incidentPoliceReportFiled              = EnumIncidentItem{incidentPoliceReportFiledID, "Police Report Filed", map[string]string{IncidentMetaCategoryKey: "Other Violations", IncidentMetaClassificationKey: "NON"}, "PoliceReportFiled", 21, "Other Violations", "NON"}
	incidentRacing                         = EnumIncidentItem{incidentRacingID, "Racing", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.Major}, "Racing", 22, IncidentCategory.MajorViolation, IncidentClass.Major}
	incidentRecklessDriving                = EnumIncidentItem{incidentRecklessDrivingID, "Reckless or Negligent Driving", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.Major}, "RecklessDriving", 23, IncidentCategory.MajorViolation, IncidentClass.Major}
	incidentRental                         = EnumIncidentItem{incidentRentalID, "Rental Incident", map[string]string{IncidentMetaCategoryKey: IncidentCategory.AccidentOrClaim, IncidentMetaClassificationKey: IncidentClass.NonChargeableAccident}, "Rental", 24, IncidentCategory.AccidentOrClaim, IncidentClass.NonChargeableAccident}
	incidentRoadside                       = EnumIncidentItem{incidentRoadsideID, "Roadside", map[string]string{IncidentMetaCategoryKey: IncidentCategory.AccidentOrClaim, IncidentMetaClassificationKey: IncidentClass.NonChargeableAccident}, "Roadside", 25, IncidentCategory.AccidentOrClaim, IncidentClass.NonChargeableAccident}
	incidentSeatBeltOrCellPhoneViolation   = EnumIncidentItem{incidentSeatBeltOrCellPhoneViolationID, "Seat Belt or Cell Phone Violation", map[string]string{IncidentMetaCategoryKey: "Other Violations", IncidentMetaClassificationKey: IncidentClass.NonChargeableConviction}, "SeatBeltOrCellPhoneViolation", 26, "Other Violations", IncidentClass.NonChargeableConviction}
	incidentSmallComprehensiveClaim        = EnumIncidentItem{incidentSmallComprehensiveClaimID, "Small Comprehensive Claim", map[string]string{IncidentMetaCategoryKey: IncidentCategory.AccidentOrClaim, IncidentMetaClassificationKey: "NON"}, "SmallComprehensiveClaim", 27, IncidentCategory.AccidentOrClaim, "NON"}
	incidentSpeeding30OrMore               = EnumIncidentItem{incidentSpeeding30OrMoreID, "Speeding – 30 MPH Or More Over Limit", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.Major}, "Speeding30OrMore", 28, IncidentCategory.MajorViolation, IncidentClass.Major}
	incidentSpeedLessThan30Over            = EnumIncidentItem{incidentSpeedLessThan30OverID, "Speeding – Less than 30 MPH Over Limit", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MinorViolation, IncidentMetaClassificationKey: IncidentClass.Minor}, "SpeedLessThan30Over", 29, IncidentCategory.MinorViolation, IncidentClass.Minor}
	incidentTheftOfVehicle                 = EnumIncidentItem{incidentTheftOfVehicleID, "Theft of Vehicle / Criminal Use", map[string]string{IncidentMetaCategoryKey: IncidentCategory.MajorViolation, IncidentMetaClassificationKey: IncidentClass.Major}, "TheftOfVehicle", 30, IncidentCategory.MajorViolation, IncidentClass.Major}
)

// EnumIncident is a collection of Incident items
type EnumIncident struct {
	Description string
	Items       []*EnumIncidentItem
	Name        string

	AccidentAtFault                *EnumIncidentItem
	AccidentNotAtFault             *EnumIncidentItem
	AdministrativeNote             *EnumIncidentItem
	DrivingOnRevokedRegistration   *EnumIncidentItem
	DUIDWI                         *EnumIncidentItem
	EquipmentCargoLittering        *EnumIncidentItem
	ExpiredOrImproperLicensePlates *EnumIncidentItem
	FailureToObeyOfficer           *EnumIncidentItem
	FailureToStop                  *EnumIncidentItem
	Glass                          *EnumIncidentItem
	HomicideOrManslaughter         *EnumIncidentItem
	ImproperDriving                *EnumIncidentItem
	IncidentOTCClaim               *EnumIncidentItem
	LicenseSuspension              *EnumIncidentItem
	MinorCoverageClaim             *EnumIncidentItem
	MinorTrafficViolation          *EnumIncidentItem
	NoPayoutClaim                  *EnumIncidentItem
	OtherMajorViolation            *EnumIncidentItem
	OutOfStateConviction           *EnumIncidentItem
	PassingSchoolBus               *EnumIncidentItem
	PoliceReportFiled              *EnumIncidentItem
	Racing                         *EnumIncidentItem
	RecklessDriving                *EnumIncidentItem
	Rental                         *EnumIncidentItem
	Roadside                       *EnumIncidentItem
	SeatBeltOrCellPhoneViolation   *EnumIncidentItem
	SmallComprehensiveClaim        *EnumIncidentItem
	Speeding30OrMore               *EnumIncidentItem
	SpeedLessThan30Over            *EnumIncidentItem
	TheftOfVehicle                 *EnumIncidentItem

	itemDict map[string]*EnumIncidentItem
}

// Incident is a public singleton instance of EnumIncident
// representing types of incident
var Incident = &EnumIncident{
	Description: "types of incident",
	Items: []*EnumIncidentItem{
		&incidentAccidentAtFault,
		&incidentAccidentNotAtFault,
		&incidentAdministrativeNote,
		&incidentDrivingOnRevokedRegistration,
		&incidentDUIDWI,
		&incidentEquipmentCargoLittering,
		&incidentExpiredOrImproperLicensePlates,
		&incidentFailureToObeyOfficer,
		&incidentFailureToStop,
		&incidentGlass,
		&incidentHomicideOrManslaughter,
		&incidentImproperDriving,
		&incidentIncidentOTCClaim,
		&incidentLicenseSuspension,
		&incidentMinorCoverageClaim,
		&incidentMinorTrafficViolation,
		&incidentNoPayoutClaim,
		&incidentOtherMajorViolation,
		&incidentOutOfStateConviction,
		&incidentPassingSchoolBus,
		&incidentPoliceReportFiled,
		&incidentRacing,
		&incidentRecklessDriving,
		&incidentRental,
		&incidentRoadside,
		&incidentSeatBeltOrCellPhoneViolation,
		&incidentSmallComprehensiveClaim,
		&incidentSpeeding30OrMore,
		&incidentSpeedLessThan30Over,
		&incidentTheftOfVehicle,
	},
	Name:                           "EnumIncident",
	AccidentAtFault:                &incidentAccidentAtFault,
	AccidentNotAtFault:             &incidentAccidentNotAtFault,
	AdministrativeNote:             &incidentAdministrativeNote,
	DrivingOnRevokedRegistration:   &incidentDrivingOnRevokedRegistration,
	DUIDWI:                         &incidentDUIDWI,
	EquipmentCargoLittering:        &incidentEquipmentCargoLittering,
	ExpiredOrImproperLicensePlates: &incidentExpiredOrImproperLicensePlates,
	FailureToObeyOfficer:           &incidentFailureToObeyOfficer,
	FailureToStop:                  &incidentFailureToStop,
	Glass:                          &incidentGlass,
	HomicideOrManslaughter:         &incidentHomicideOrManslaughter,
	ImproperDriving:                &incidentImproperDriving,
	IncidentOTCClaim:               &incidentIncidentOTCClaim,
	LicenseSuspension:              &incidentLicenseSuspension,
	MinorCoverageClaim:             &incidentMinorCoverageClaim,
	MinorTrafficViolation:          &incidentMinorTrafficViolation,
	NoPayoutClaim:                  &incidentNoPayoutClaim,
	OtherMajorViolation:            &incidentOtherMajorViolation,
	OutOfStateConviction:           &incidentOutOfStateConviction,
	PassingSchoolBus:               &incidentPassingSchoolBus,
	PoliceReportFiled:              &incidentPoliceReportFiled,
	Racing:                         &incidentRacing,
	RecklessDriving:                &incidentRecklessDriving,
	Rental:                         &incidentRental,
	Roadside:                       &incidentRoadside,
	SeatBeltOrCellPhoneViolation:   &incidentSeatBeltOrCellPhoneViolation,
	SmallComprehensiveClaim:        &incidentSmallComprehensiveClaim,
	Speeding30OrMore:               &incidentSpeeding30OrMore,
	SpeedLessThan30Over:            &incidentSpeedLessThan30Over,
	TheftOfVehicle:                 &incidentTheftOfVehicle,

	itemDict: map[string]*EnumIncidentItem{
		strings.ToLower(string(incidentAccidentAtFaultID)):                &incidentAccidentAtFault,
		strings.ToLower(string(incidentAccidentNotAtFaultID)):             &incidentAccidentNotAtFault,
		strings.ToLower(string(incidentAdministrativeNoteID)):             &incidentAdministrativeNote,
		strings.ToLower(string(incidentDrivingOnRevokedRegistrationID)):   &incidentDrivingOnRevokedRegistration,
		strings.ToLower(string(incidentDUIDWIID)):                         &incidentDUIDWI,
		strings.ToLower(string(incidentEquipmentCargoLitteringID)):        &incidentEquipmentCargoLittering,
		strings.ToLower(string(incidentExpiredOrImproperLicensePlatesID)): &incidentExpiredOrImproperLicensePlates,
		strings.ToLower(string(incidentFailureToObeyOfficerID)):           &incidentFailureToObeyOfficer,
		strings.ToLower(string(incidentFailureToStopID)):                  &incidentFailureToStop,
		strings.ToLower(string(incidentGlassID)):                          &incidentGlass,
		strings.ToLower(string(incidentHomicideOrManslaughterID)):         &incidentHomicideOrManslaughter,
		strings.ToLower(string(incidentImproperDrivingID)):                &incidentImproperDriving,
		strings.ToLower(string(incidentIncidentOTCClaimID)):               &incidentIncidentOTCClaim,
		strings.ToLower(string(incidentLicenseSuspensionID)):              &incidentLicenseSuspension,
		strings.ToLower(string(incidentMinorCoverageClaimID)):             &incidentMinorCoverageClaim,
		strings.ToLower(string(incidentMinorTrafficViolationID)):          &incidentMinorTrafficViolation,
		strings.ToLower(string(incidentNoPayoutClaimID)):                  &incidentNoPayoutClaim,
		strings.ToLower(string(incidentOtherMajorViolationID)):            &incidentOtherMajorViolation,
		strings.ToLower(string(incidentOutOfStateConvictionID)):           &incidentOutOfStateConviction,
		strings.ToLower(string(incidentPassingSchoolBusID)):               &incidentPassingSchoolBus,
		strings.ToLower(string(incidentPoliceReportFiledID)):              &incidentPoliceReportFiled,
		strings.ToLower(string(incidentRacingID)):                         &incidentRacing,
		strings.ToLower(string(incidentRecklessDrivingID)):                &incidentRecklessDriving,
		strings.ToLower(string(incidentRentalID)):                         &incidentRental,
		strings.ToLower(string(incidentRoadsideID)):                       &incidentRoadside,
		strings.ToLower(string(incidentSeatBeltOrCellPhoneViolationID)):   &incidentSeatBeltOrCellPhoneViolation,
		strings.ToLower(string(incidentSmallComprehensiveClaimID)):        &incidentSmallComprehensiveClaim,
		strings.ToLower(string(incidentSpeeding30OrMoreID)):               &incidentSpeeding30OrMore,
		strings.ToLower(string(incidentSpeedLessThan30OverID)):            &incidentSpeedLessThan30Over,
		strings.ToLower(string(incidentTheftOfVehicleID)):                 &incidentTheftOfVehicle,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumIncident) ByID(id IncidentIdentifier) *EnumIncidentItem {
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
func (e *EnumIncident) ByIDString(idx string) *EnumIncidentItem {
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
func (e *EnumIncident) ByIndex(idx int) *EnumIncidentItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedIncidentID is a struct that is designed to replace a *IncidentID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *IncidentID it contains while being a better JSON citizen.
type ValidatedIncidentID struct {
	// id will point to a valid IncidentID, if possible
	// If id is nil, then ValidatedIncidentID.Valid() will return false.
	id *IncidentID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedIncidentID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedIncidentID
func (vi *ValidatedIncidentID) Clone() *ValidatedIncidentID {
	if vi == nil {
		return nil
	}

	var cid *IncidentID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedIncidentID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedIncidentIds represent the same Incident
func (vi *ValidatedIncidentID) Equals(vj *ValidatedIncidentID) bool {
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

// Valid returns true if and only if the ValidatedIncidentID corresponds to a recognized Incident
func (vi *ValidatedIncidentID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedIncidentID) ID() *IncidentID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedIncidentID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedIncidentID) ValidatedID() *ValidatedIncidentID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedIncidentID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedIncidentID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedIncidentID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedIncidentID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedIncidentID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := IncidentID(capString)
	item := Incident.ByID(&id)
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

func (vi ValidatedIncidentID) String() string {
	return vi.ToIDString()
}

type IncidentIdentifier interface {
	ID() *IncidentID
	Valid() bool
}
