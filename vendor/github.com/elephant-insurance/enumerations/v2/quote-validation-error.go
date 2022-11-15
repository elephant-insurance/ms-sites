package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// QuoteValidationErrorID uniquely identifies a particular QuoteValidationError
type QuoteValidationErrorID string

// Clone creates a safe, independent copy of a QuoteValidationErrorID
func (i *QuoteValidationErrorID) Clone() *QuoteValidationErrorID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two QuoteValidationErrorIds are equivalent
func (i *QuoteValidationErrorID) Equals(j *QuoteValidationErrorID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *QuoteValidationErrorID that is either valid or nil
func (i *QuoteValidationErrorID) ID() *QuoteValidationErrorID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *QuoteValidationErrorID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the QuoteValidationErrorID corresponds to a recognized QuoteValidationError
func (i *QuoteValidationErrorID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return QuoteValidationError.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *QuoteValidationErrorID) ValidatedID() *ValidatedQuoteValidationErrorID {
	if i != nil {
		return &ValidatedQuoteValidationErrorID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *QuoteValidationErrorID) MarshalJSON() ([]byte, error) {
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

func (i *QuoteValidationErrorID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := QuoteValidationErrorID(dataString)
	item := QuoteValidationError.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	quoteValidationErrorApplicantAddressCityNilID                              QuoteValidationErrorID = "ApplicantAddressCityNil"
	quoteValidationErrorApplicantAddressCountyNilID                            QuoteValidationErrorID = "ApplicantAddressCountyNil"
	quoteValidationErrorApplicantAddressIsPOBoxID                              QuoteValidationErrorID = "ApplicantAddressIsPOBox"
	quoteValidationErrorApplicantAddressLine1NilID                             QuoteValidationErrorID = "ApplicantAddressLine1Nil"
	quoteValidationErrorApplicantAddressNilID                                  QuoteValidationErrorID = "ApplicantAddressNil"
	quoteValidationErrorApplicantAddressLine1ExceedsMaxLengthID                QuoteValidationErrorID = "ApplicantAddressLine1ExceedsMaxLength"
	quoteValidationErrorApplicantAddressNotRatableID                           QuoteValidationErrorID = "ApplicantAddressNotRatable"
	quoteValidationErrorApplicantAddressPostalCodeNilID                        QuoteValidationErrorID = "ApplicantAddressPostalCodeNil"
	quoteValidationErrorApplicantAddressStateInvalidID                         QuoteValidationErrorID = "ApplicantAddressStateInvalid"
	quoteValidationErrorApplicantAddressStateNilID                             QuoteValidationErrorID = "ApplicantAddressStateNil"
	quoteValidationErrorApplicantDateOfBirthNilID                              QuoteValidationErrorID = "ApplicantDateOfBirthNil"
	quoteValidationErrorApplicantLicenseStatusInvalidID                        QuoteValidationErrorID = "ApplicantLicenseStatusInvalid"
	quoteValidationErrorApplicantLicenseStatusPermitID                         QuoteValidationErrorID = "ApplicantLicenseStatusPermit"
	quoteValidationErrorApplicantMaritalStatusMarriedNoSpouseID                QuoteValidationErrorID = "ApplicantMaritalStatusMarriedNoSpouse"
	quoteValidationErrorApplicantMaritalStatusMultipleSpousesID                QuoteValidationErrorID = "ApplicantMaritalStatusMultipleSpouses"
	quoteValidationErrorApplicantMaritalStatusNilID                            QuoteValidationErrorID = "ApplicantMaritalStatusNil"
	quoteValidationErrorApplicantMaritalStatusNotValidID                       QuoteValidationErrorID = "ApplicantMaritalStatusNotValid"
	quoteValidationErrorApplicantNilID                                         QuoteValidationErrorID = "ApplicantNil"
	quoteValidationErrorApplicantNotDriverID                                   QuoteValidationErrorID = "ApplicantNotDriver"
	quoteValidationErrorApplicantPhoneNumberInvalidID                          QuoteValidationErrorID = "ApplicantPhoneNumberInvalid"
	quoteValidationErrorApplicantPhoneNumberNilID                              QuoteValidationErrorID = "ApplicantPhoneNumberNil"
	quoteValidationErrorApplicantUnder18ID                                     QuoteValidationErrorID = "ApplicantUnder18"
	quoteValidationErrorApplicantOccupationStatusNilID                         QuoteValidationErrorID = "ApplicantOccupationStatusNil"
	quoteValidationErrorApplicantOccupationStatusUnemployedID                  QuoteValidationErrorID = "ApplicantOccupationStatusUnemployed"
	quoteValidationErrorCannotCarryBothUMPDAndCollID                           QuoteValidationErrorID = "CannotCarryBothUMPDAndColl"
	quoteValidationErrorCantHoldLoanLeasePayOffOnOlderVehicleID                QuoteValidationErrorID = "CantHoldLoanLeasePayOffOnOlderVehicle"
	quoteValidationErrorCollRequiredForRentID                                  QuoteValidationErrorID = "CollRequiredForRent"
	quoteValidationErrorCompCantExceedCollID                                   QuoteValidationErrorID = "CompCantExceedColl"
	quoteValidationErrorCompCollNotAllowedOnOlderVehicleID                     QuoteValidationErrorID = "CompCollNotAllowedOnOlderVehicle"
	quoteValidationErrorCompCollRequiredForCustomEquipmentID                   QuoteValidationErrorID = "CompCollRequiredForCustomEquipment"
	quoteValidationErrorCompCollRequiredForLoanID                              QuoteValidationErrorID = "CompCollRequiredForLoan"
	quoteValidationErrorCompCollRequiredForRentID                              QuoteValidationErrorID = "CompCollRequiredForRent"
	quoteValidationErrorCompRequiredForRentID                                  QuoteValidationErrorID = "CompRequiredForRent"
	quoteValidationErrorCompRequiredToHoldCollID                               QuoteValidationErrorID = "CompRequiredToHoldColl"
	quoteValidationErrorDDRequiresColl500OrBetterID                            QuoteValidationErrorID = "DDRequiresColl500OrBetter"
	quoteValidationErrorDriverArrayCannotBeEmptyID                             QuoteValidationErrorID = "DriverArrayCannotBeEmpty"
	quoteValidationErrorDriverHasBadLicenseStatusID                            QuoteValidationErrorID = "DriverHasBadLicenseStatus"
	quoteValidationErrorDriverIDNotUniqueID                                    QuoteValidationErrorID = "DriverIDNotUnique"
	quoteValidationErrorDriverFirstNameExceedsMaxLengthID                      QuoteValidationErrorID = "DriverFirstNameExceedsMaxLength"
	quoteValidationErrorDriverLastNameExceedsMaxLengthID                       QuoteValidationErrorID = "DriverLastNameExceedsMaxLength"
	quoteValidationErrorDriverIncidentsMoreThan1MinorConvictionID              QuoteValidationErrorID = "DriverIncidentsMoreThan1MinorConviction"
	quoteValidationErrorDriverIncidentsMoreThan2MajorConvictionsID             QuoteValidationErrorID = "DriverIncidentsMoreThan2MajorConvictions"
	quoteValidationErrorDriverIncidentsMoreThan2MajorConvictionsIncludingDUIID QuoteValidationErrorID = "DriverIncidentsMoreThan2MajorConvictionsIncludingDUI"
	quoteValidationErrorDriverIncidentsMoreThan4AtFaultID                      QuoteValidationErrorID = "DriverIncidentsMoreThan4AtFault"
	quoteValidationErrorDriverIncidentsMoreThan4NotAtFaultID                   QuoteValidationErrorID = "DriverIncidentsMoreThan4NotAtFault"
	quoteValidationErrorDriverIncidentsMoreThanOneDUIID                        QuoteValidationErrorID = "DriverIncidentsMoreThanOneDUI"
	quoteValidationErrorDriverPrimaryVehicleInvalidID                          QuoteValidationErrorID = "DriverPrimaryVehicleInvalid"
	quoteValidationErrorDriverPrimaryVehicleNilID                              QuoteValidationErrorID = "DriverPrimaryVehicleNil"
	quoteValidationErrorEveryDriverRequiresDriverIDID                          QuoteValidationErrorID = "EveryDriverRequiresDriverID"
	quoteValidationErrorEveryVehicleRequiresVehicleIDID                        QuoteValidationErrorID = "EveryVehicleRequiresVehicleID"
	quoteValidationErrorHasActivePolicyID                                      QuoteValidationErrorID = "HasActivePolicy"
	quoteValidationErrorHasBadDebtsFlagID                                      QuoteValidationErrorID = "HasBadDebtsFlag"
	quoteValidationErrorHasMaterialMisRepID                                    QuoteValidationErrorID = "HasMaterialMisRep"
	quoteValidationErrorNoLoanLeasePayOffWithoutLienHolderID                   QuoteValidationErrorID = "NoLoanLeasePayOffWithoutLienHolder"
	quoteValidationErrorPDCantExceedBIID                                       QuoteValidationErrorID = "PDCantExceedBI"
	quoteValidationErrorPDCantExceedBIExceptForLowestLimitID                   QuoteValidationErrorID = "PDCantExceedBIExceptForLowestLimit"
	quoteValidationErrorSameDayBindRuleErrorID                                 QuoteValidationErrorID = "SameDayBindRuleError"
	quoteValidationErrorPolicyCurrentInsuranceCarrierNilID                     QuoteValidationErrorID = "PolicyCurrentInsuranceCarrierNil"
	quoteValidationErrorPolicyCurrentInsuranceLimitsNilID                      QuoteValidationErrorID = "PolicyCurrentInsuranceLimitsNil"
	quoteValidationErrorPolicyCurrentInsurancePriorLapseNilID                  QuoteValidationErrorID = "PolicyCurrentInsurancePriorLapseNil"
	quoteValidationErrorPolicyCurrentInsuranceStatusNilID                      QuoteValidationErrorID = "PolicyCurrentInsuranceStatusNil"
	quoteValidationErrorPolicyCurrentInsuranceYearsWithNilID                   QuoteValidationErrorID = "PolicyCurrentInsuranceYearsWithNil"
	quoteValidationErrorPolicyDriversEmptyID                                   QuoteValidationErrorID = "PolicyDriversEmpty"
	quoteValidationErrorPolicyEffectiveDateNilID                               QuoteValidationErrorID = "PolicyEffectiveDateNil"
	quoteValidationErrorPolicyEffectiveDateNotMoreThan60DaysInFutureID         QuoteValidationErrorID = "policyEffectiveDateNotMoreThan60DaysInFuture"
	quoteValidationErrorPolicyEffectiveDatePastID                              QuoteValidationErrorID = "PolicyEffectiveDatePast"
	quoteValidationErrorPolicyHasDriverWithSR22ID                              QuoteValidationErrorID = "PolicyHasDriverWithSR22"
	quoteValidationErrorPolicyNilID                                            QuoteValidationErrorID = "PolicyNil"
	quoteValidationErrorPolicyVehiclesEmptyID                                  QuoteValidationErrorID = "PolicyVehiclesEmpty"
	quoteValidationErrorTotalIncidentsMoreThan4AtFaultID                       QuoteValidationErrorID = "TotalIncidentsMoreThan4AtFault"
	quoteValidationErrorUMBIAvailableOnlyWithUMPDID                            QuoteValidationErrorID = "UMBIAvailableOnlyWithUMPD"
	quoteValidationErrorUMBICantExceedBIID                                     QuoteValidationErrorID = "UMBICantExceedBI"
	quoteValidationErrorUMBIMustBeSameAsBIID                                   QuoteValidationErrorID = "UMBIMustBeSameAsBI"
	quoteValidationErrorUMPDAvailableOnlyWithUMBIID                            QuoteValidationErrorID = "UMPDAvailableOnlyWithUMBI"
	quoteValidationErrorUMPDCantExceedPDID                                     QuoteValidationErrorID = "UMPDCantExceedPD"
	quoteValidationErrorUMPDMustBeSameAsPDID                                   QuoteValidationErrorID = "UMPDMustBeSameAsPD"
	quoteValidationErrorVehicleIDNotUniqueID                                   QuoteValidationErrorID = "VehicleIDNotUnique"
	quoteValidationErrorVehicleNotInsurableID                                  QuoteValidationErrorID = "VehicleNotInsurable"
	quoteValidationErrorVehicleParkedAtMoreThanOneAddressID                    QuoteValidationErrorID = "VehicleParkedAtMoreThanOneAddress"
	quoteValidationErrorVehiclePrimaryDriverInvalidID                          QuoteValidationErrorID = "VehiclePrimaryDriverInvalid"
	quoteValidationErrorVehiclePrimaryDriverNilID                              QuoteValidationErrorID = "VehiclePrimaryDriverNil"
	quoteValidationErrorVehicleVINInvalidID                                    QuoteValidationErrorID = "VehicleVINInvalid"
	quoteValidationErrorVehicleVINNilID                                        QuoteValidationErrorID = "VehicleVINNil"
	quoteValidationErrorCoverageValueNotValidID                                QuoteValidationErrorID = "CoverageValueNotValid"
	quoteValidationErrorEISPABodilyInjuryCovCoverageValueNotValidID            QuoteValidationErrorID = "EISPABodilyInjuryCovCoverageValueNotValid"
	quoteValidationErrorEISPAPropertyDamageCovCoverageValueNotValidID          QuoteValidationErrorID = "EISPAPropertyDamageCovCoverageValueNotValid"
	quoteValidationErrorPAUMBICovCoverageValueNotValidID                       QuoteValidationErrorID = "PAUMBICovCoverageValueNotValid"
	quoteValidationErrorPAMedPayCovCoverageValueNotValidID                     QuoteValidationErrorID = "PAMedPayCovCoverageValueNotValid"
	quoteValidationErrorPAUMPDCovCoverageValueNotValidID                       QuoteValidationErrorID = "PAUMPDCovCoverageValueNotValid"
	quoteValidationErrorPAPIP_MDCoverageValueNotValidID                        QuoteValidationErrorID = "PAPIP_MDCoverageValueNotValid"
	quoteValidationErrorPAPIP_TXCoverageValueNotValidID                        QuoteValidationErrorID = "PAPIP_TXCoverageValueNotValid"
	quoteValidationErrorEISPALegalPlanCovCoverageValueNotValidID               QuoteValidationErrorID = "EISPALegalPlanCovCoverageValueNotValid"
	quoteValidationErrorEISPAIncomeLossCovCoverageValueNotValidID              QuoteValidationErrorID = "EISPAIncomeLossCovCoverageValueNotValid"
	quoteValidationErrorPAUM_GACovCoverageValueNotValidID                      QuoteValidationErrorID = "PAUM_GACovCoverageValueNotValid"
	quoteValidationErrorPAComprehensiveCovCoverageValueNotValidID              QuoteValidationErrorID = "PAComprehensiveCovCoverageValueNotValid"
	quoteValidationErrorPACollisionCovCoverageValueNotValidID                  QuoteValidationErrorID = "PACollisionCovCoverageValueNotValid"
	quoteValidationErrorEISPAUMPD_ILCovCoverageValueNotValidID                 QuoteValidationErrorID = "EISPAUMPD_ILCovCoverageValueNotValid"
	quoteValidationErrorEISPAUMPD_OHCovCoverageValueNotValidID                 QuoteValidationErrorID = "EISPAUMPD_OHCovCoverageValueNotValid"
	quoteValidationErrorPATowingLaborCovCoverageValueNotValidID                QuoteValidationErrorID = "PATowingLaborCovCoverageValueNotValid"
	quoteValidationErrorPARentalCovCoverageValueNotValidID                     QuoteValidationErrorID = "PARentalCovCoverageValueNotValid"
	quoteValidationErrorEISPALoanGapCovCoverageValueNotValidID                 QuoteValidationErrorID = "EISPALoanGapCovCoverageValueNotValid"
	quoteValidationErrorEISPACustEquipCovCoverageValueNotValidID               QuoteValidationErrorID = "EISPACustEquipCovCoverageValueNotValid"
	quoteValidationErrorPAVehicleMonitorCovCoverageValueNotValidID             QuoteValidationErrorID = "PAVehicleMonitorCovCoverageValueNotValid"
	quoteValidationErrorCoveragesNotValidForFr44ID                             QuoteValidationErrorID = "CoveragesNotValidForFr44"
	quoteValidationErrorAllDriversMustHaveGenderID                             QuoteValidationErrorID = "AllDriversMustHaveGender"
	quoteValidationErrorRestrictedLicenceStateForPolicyID                      QuoteValidationErrorID = "RestrictedLicenceStateForPolicy"
)

// EnumQuoteValidationErrorItem describes an entry in an enumeration of QuoteValidationError
type EnumQuoteValidationErrorItem struct {
	ID        QuoteValidationErrorID `json:"Value"`
	Desc      string                 `json:"Description,omitempty"`
	Meta      map[string]string      `json:",omitempty"`
	Name      string                 `json:"Name"`
	SortOrder int
}

var (
	quoteValidationErrorApplicantAddressCityNil                              = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantAddressCityNilID, "the applicant's address must have a city", nil, "ApplicantAddressCityNil", 1}
	quoteValidationErrorApplicantAddressCountyNil                            = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantAddressCountyNilID, "the applicant's address must have a valid county", nil, "ApplicantAddressCountyNil", 2}
	quoteValidationErrorApplicantAddressIsPOBox                              = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantAddressIsPOBoxID, "the applicant's address must not be in a postal code that only serves post office boxes", nil, "ApplicantAddressIsPOBox", 3}
	quoteValidationErrorApplicantAddressLine1Nil                             = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantAddressLine1NilID, "the applicant's address must have a line 1", nil, "ApplicantAddressLine1Nil", 4}
	quoteValidationErrorApplicantAddressNil                                  = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantAddressNilID, "the applicant must have an address", nil, "ApplicantAddressNil", 5}
	quoteValidationErrorApplicantAddressLine1ExceedsMaxLength                = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantAddressLine1ExceedsMaxLengthID, "the applicant address line 1 must not have more than 60 characters", nil, "ApplicantAddressLine1ExceedsMaxLength", 6}
	quoteValidationErrorApplicantAddressNotRatable                           = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantAddressNotRatableID, "the applicant's address must not be in a postal code that is ratable", nil, "ApplicantAddressNotRatable", 7}
	quoteValidationErrorApplicantAddressPostalCodeNil                        = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantAddressPostalCodeNilID, "the applicant's address must have a postal code", nil, "ApplicantAddressPostalCodeNil", 8}
	quoteValidationErrorApplicantAddressStateInvalid                         = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantAddressStateInvalidID, "the applicant's address must be in a state where we provide coverage", nil, "ApplicantAddressStateInvalid", 9}
	quoteValidationErrorApplicantAddressStateNil                             = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantAddressStateNilID, "the applicant's address must have a state", nil, "ApplicantAddressStateNil", 10}
	quoteValidationErrorApplicantDateOfBirthNil                              = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantDateOfBirthNilID, "the applicant must have a date of birth", nil, "ApplicantDateOfBirthNil", 11}
	quoteValidationErrorApplicantLicenseStatusInvalid                        = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantLicenseStatusInvalidID, "the applicant must have a valid driver's license", nil, "ApplicantLicenseStatusInvalid", 12}
	quoteValidationErrorApplicantLicenseStatusPermit                         = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantLicenseStatusPermitID, "the applicant must not have a driver's permit", nil, "ApplicantLicenseStatusPermit", 13}
	quoteValidationErrorApplicantMaritalStatusMarriedNoSpouse                = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantMaritalStatusMarriedNoSpouseID, "a married applicant must list a spouse", nil, "ApplicantMaritalStatusMarriedNoSpouse", 14}
	quoteValidationErrorApplicantMaritalStatusMultipleSpouses                = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantMaritalStatusMultipleSpousesID, "the applicant must not have more than one spouse", nil, "ApplicantMaritalStatusMultipleSpouses", 15}
	quoteValidationErrorApplicantMaritalStatusNil                            = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantMaritalStatusNilID, "the applicant must have a marital status", nil, "ApplicantMaritalStatusNil", 16}
	quoteValidationErrorApplicantMaritalStatusNotValid                       = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantMaritalStatusNotValidID, "the applicant must have a valid marital status", nil, "ApplicantMaritalStatusNotValid", 17}
	quoteValidationErrorApplicantNil                                         = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantNilID, "quote applicant cannot be nil", nil, "ApplicantNil", 18}
	quoteValidationErrorApplicantNotDriver                                   = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantNotDriverID, "the applicant must be a driver on the policy", nil, "ApplicantNotDriver", 19}
	quoteValidationErrorApplicantPhoneNumberInvalid                          = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantPhoneNumberInvalidID, "the applicant's phone number must be a 10-digit number without leading zeros", nil, "ApplicantPhoneNumberInvalid", 20}
	quoteValidationErrorApplicantPhoneNumberNil                              = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantPhoneNumberNilID, "the applicant must list a phone number", nil, "ApplicantPhoneNumberNil", 21}
	quoteValidationErrorApplicantUnder18                                     = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantUnder18ID, "the applicant must be at least 18 years of age", nil, "ApplicantUnder18", 22}
	quoteValidationErrorApplicantOccupationStatusNil                         = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantOccupationStatusNilID, "the applicant must have an occupation status", nil, "ApplicantOccupationStatusNil", 23}
	quoteValidationErrorApplicantOccupationStatusUnemployed                  = EnumQuoteValidationErrorItem{quoteValidationErrorApplicantOccupationStatusUnemployedID, "the applicant's occupation title and/or occupation status cannot be unemployed", nil, "ApplicantOccupationStatusUnemployed", 24}
	quoteValidationErrorCannotCarryBothUMPDAndColl                           = EnumQuoteValidationErrorItem{quoteValidationErrorCannotCarryBothUMPDAndCollID, "each vehicle may carry either Collision coverage or Uninsured/Underinsured Motorist Property Damage coverage, not both", nil, "CannotCarryBothUMPDAndColl", 25}
	quoteValidationErrorCantHoldLoanLeasePayOffOnOlderVehicle                = EnumQuoteValidationErrorItem{quoteValidationErrorCantHoldLoanLeasePayOffOnOlderVehicleID, "cannot hold Loan/Lease Payoff if vehicle is more than 8 years old", nil, "CantHoldLoanLeasePayOffOnOlderVehicle", 26}
	quoteValidationErrorCollRequiredForRent                                  = EnumQuoteValidationErrorItem{quoteValidationErrorCollRequiredForRentID, "Collision coverage is required for Rental Insurance to be added to a vehicle", nil, "CollRequiredForRent", 27}
	quoteValidationErrorCompCantExceedColl                                   = EnumQuoteValidationErrorItem{quoteValidationErrorCompCantExceedCollID, "comprehensive deductible cannot exceed collision deductible", nil, "CompCantExceedColl", 28}
	quoteValidationErrorCompCollNotAllowedOnOlderVehicle                     = EnumQuoteValidationErrorItem{quoteValidationErrorCompCollNotAllowedOnOlderVehicleID, "collision cannot be added for vehicles older than 1981", nil, "CompCollNotAllowedOnOlderVehicle", 29}
	quoteValidationErrorCompCollRequiredForCustomEquipment                   = EnumQuoteValidationErrorItem{quoteValidationErrorCompCollRequiredForCustomEquipmentID, "comprehensive and collision coverages are required to carry Additional Custom Parts or Equipment coverage", nil, "CompCollRequiredForCustomEquipment", 30}
	quoteValidationErrorCompCollRequiredForLoan                              = EnumQuoteValidationErrorItem{quoteValidationErrorCompCollRequiredForLoanID, "comprehensive and collision coverages are required to carry Loan/Lease Payoff Coverage", nil, "CompCollRequiredForLoan", 31}
	quoteValidationErrorCompCollRequiredForRent                              = EnumQuoteValidationErrorItem{quoteValidationErrorCompCollRequiredForRentID, "comprehensive and collision coverages are required to carry Rental coverage.", nil, "CompCollRequiredForRent", 32}
	quoteValidationErrorCompRequiredForRent                                  = EnumQuoteValidationErrorItem{quoteValidationErrorCompRequiredForRentID, "comprehensive coverage is required to carry Rental coverage", nil, "CompRequiredForRent", 33}
	quoteValidationErrorCompRequiredToHoldColl                               = EnumQuoteValidationErrorItem{quoteValidationErrorCompRequiredToHoldCollID, "each vehicle with collision coverage must also carry comprehensive coverage", nil, "CompRequiredToHoldColl", 34}
	quoteValidationErrorDDRequiresColl500OrBetter                            = EnumQuoteValidationErrorItem{quoteValidationErrorDDRequiresColl500OrBetterID, "a vehicle must have collision coverage with a deductible of $500 or less to qualify for Diminishing Deductible", nil, "DDRequiresColl500OrBetter", 35}
	quoteValidationErrorDriverArrayCannotBeEmpty                             = EnumQuoteValidationErrorItem{quoteValidationErrorDriverArrayCannotBeEmptyID, "Driver array can not be empty", nil, "DriverArrayCannotBeEmpty", 36}
	quoteValidationErrorDriverHasBadLicenseStatus                            = EnumQuoteValidationErrorItem{quoteValidationErrorDriverHasBadLicenseStatusID, "no driver on the policy can have a suspended, revoked, expired, surrendered, or non-licensed license status", nil, "DriverHasBadLicenseStatus", 37}
	quoteValidationErrorDriverIDNotUnique                                    = EnumQuoteValidationErrorItem{quoteValidationErrorDriverIDNotUniqueID, "no two drivers on the policy can have same driver id", nil, "DriverIDNotUnique", 38}
	quoteValidationErrorDriverFirstNameExceedsMaxLength                      = EnumQuoteValidationErrorItem{quoteValidationErrorDriverFirstNameExceedsMaxLengthID, "Driver first name has exceeded the maximum length of 30 characters", nil, "DriverFirstNameExceedsMaxLength", 39}
	quoteValidationErrorDriverLastNameExceedsMaxLength                       = EnumQuoteValidationErrorItem{quoteValidationErrorDriverLastNameExceedsMaxLengthID, "Driver last name has exceeded the maximum length of 30 characters", nil, "DriverLastNameExceedsMaxLength", 40}
	quoteValidationErrorDriverIncidentsMoreThan1MinorConviction              = EnumQuoteValidationErrorItem{quoteValidationErrorDriverIncidentsMoreThan1MinorConvictionID, "no driver on the policy can have more than one minor conviction in the past 3 years", nil, "DriverIncidentsMoreThan1MinorConviction", 41}
	quoteValidationErrorDriverIncidentsMoreThan2MajorConvictions             = EnumQuoteValidationErrorItem{quoteValidationErrorDriverIncidentsMoreThan2MajorConvictionsID, "no driver on the policy can have more than two major convictions in the past 3 years", nil, "DriverIncidentsMoreThan2MajorConvictions", 42}
	quoteValidationErrorDriverIncidentsMoreThan2MajorConvictionsIncludingDUI = EnumQuoteValidationErrorItem{quoteValidationErrorDriverIncidentsMoreThan2MajorConvictionsIncludingDUIID, "no rated driver can have more than two major convictions including DUI in the past 3 years", nil, "DriverIncidentsMoreThan2MajorConvictionsIncludingDUI", 43}
	quoteValidationErrorDriverIncidentsMoreThan4AtFault                      = EnumQuoteValidationErrorItem{quoteValidationErrorDriverIncidentsMoreThan4AtFaultID, "no driver on the policy can have more than four at fault incidents in the past 3 years", nil, "DriverIncidentsMoreThan4AtFault", 44}
	quoteValidationErrorDriverIncidentsMoreThan4NotAtFault                   = EnumQuoteValidationErrorItem{quoteValidationErrorDriverIncidentsMoreThan4NotAtFaultID, "no driver on the policy can have more than four not at fault incidents in the past 3 years", nil, "DriverIncidentsMoreThan4NotAtFault", 45}
	quoteValidationErrorDriverIncidentsMoreThanOneDUI                        = EnumQuoteValidationErrorItem{quoteValidationErrorDriverIncidentsMoreThanOneDUIID, "no rated driver can have more than one DUI in the past 3 years", nil, "DriverIncidentsMoreThanOneDUI", 46}
	quoteValidationErrorDriverPrimaryVehicleInvalid                          = EnumQuoteValidationErrorItem{quoteValidationErrorDriverPrimaryVehicleInvalidID, "every driver's primary vehicle id must correspond to a vehicle on the policy", nil, "DriverPrimaryVehicleInvalid", 47}
	quoteValidationErrorDriverPrimaryVehicleNil                              = EnumQuoteValidationErrorItem{quoteValidationErrorDriverPrimaryVehicleNilID, "every driver must be assigned a primary vehicle", nil, "DriverPrimaryVehicleNil", 48}
	quoteValidationErrorEveryDriverRequiresDriverID                          = EnumQuoteValidationErrorItem{quoteValidationErrorEveryDriverRequiresDriverIDID, "each driver on the policy must be assigned a unique DriverID", nil, "EveryDriverRequiresDriverID", 49}
	quoteValidationErrorEveryVehicleRequiresVehicleID                        = EnumQuoteValidationErrorItem{quoteValidationErrorEveryVehicleRequiresVehicleIDID, "each vehicle on the policy must be assigned a unique VehicleID", nil, "EveryVehicleRequiresVehicleID", 50}
	quoteValidationErrorHasActivePolicy                                      = EnumQuoteValidationErrorItem{quoteValidationErrorHasActivePolicyID, "the customer already has a policy with us", nil, "HasActivePolicy", 51}
	quoteValidationErrorHasBadDebtsFlag                                      = EnumQuoteValidationErrorItem{quoteValidationErrorHasBadDebtsFlagID, "the customer has been flagged for bad debts", nil, "HasBadDebtsFlag", 52}
	quoteValidationErrorHasMaterialMisRep                                    = EnumQuoteValidationErrorItem{quoteValidationErrorHasMaterialMisRepID, "the customer has been flagged for material misrepresentation", nil, "HasMaterialMisRep", 53}
	quoteValidationErrorNoLoanLeasePayOffWithoutLienHolder                   = EnumQuoteValidationErrorItem{quoteValidationErrorNoLoanLeasePayOffWithoutLienHolderID, "cannot have Loan/Lease Payoff without lienholder on vehicle", nil, "NoLoanLeasePayOffWithoutLienHolder", 54}
	quoteValidationErrorPDCantExceedBI                                       = EnumQuoteValidationErrorItem{quoteValidationErrorPDCantExceedBIID, "Property Damage limit cannot exceed the per person Bodily Injury limit", nil, "PDCantExceedBI", 55}
	quoteValidationErrorPDCantExceedBIExceptForLowestLimit                   = EnumQuoteValidationErrorItem{quoteValidationErrorPDCantExceedBIExceptForLowestLimitID, "PD cant Exceed BI except if PD is 30", nil, "PDCantExceedBIExceptForLowestLimit", 56}
	quoteValidationErrorSameDayBindRuleError                                 = EnumQuoteValidationErrorItem{quoteValidationErrorSameDayBindRuleErrorID, "Currently not insured and not allowed to do same day bind", nil, "SameDayBindRuleError", 57}
	quoteValidationErrorPolicyCurrentInsuranceCarrierNil                     = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyCurrentInsuranceCarrierNilID, "current insurance carrier must be supplied", nil, "PolicyCurrentInsuranceCarrierNil", 58}
	quoteValidationErrorPolicyCurrentInsuranceLimitsNil                      = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyCurrentInsuranceLimitsNilID, "current insurance injury limits must be supplied", nil, "PolicyCurrentInsuranceLimitsNil", 59}
	quoteValidationErrorPolicyCurrentInsurancePriorLapseNil                  = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyCurrentInsurancePriorLapseNilID, "prior insurance lapse information must be supplied", nil, "PolicyCurrentInsurancePriorLapseNil", 60}
	quoteValidationErrorPolicyCurrentInsuranceStatusNil                      = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyCurrentInsuranceStatusNilID, "current insurance status must be supplied", nil, "PolicyCurrentInsuranceStatusNil", 61}
	quoteValidationErrorPolicyCurrentInsuranceYearsWithNil                   = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyCurrentInsuranceYearsWithNilID, "years with current insurance carrier must be supplied", nil, "PolicyCurrentInsuranceYearsWithNil", 62}
	quoteValidationErrorPolicyDriversEmpty                                   = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyDriversEmptyID, "the policy must have at least one driver", nil, "PolicyDriversEmpty", 63}
	quoteValidationErrorPolicyEffectiveDateNil                               = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyEffectiveDateNilID, "the policy must have an effective date", nil, "PolicyEffectiveDateNil", 64}
	quoteValidationErrorPolicyEffectiveDateNotMoreThan60DaysInFuture         = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyEffectiveDateNotMoreThan60DaysInFutureID, "the policy can not be more than 60 days in future", nil, "PolicyEffectiveDateNotMoreThan60DaysInFuture", 65}
	quoteValidationErrorPolicyEffectiveDatePast                              = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyEffectiveDatePastID, "the policy effective date cannot be in the past", nil, "PolicyEffectiveDatePast", 66}
	quoteValidationErrorPolicyHasDriverWithSR22                              = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyHasDriverWithSR22ID, "no driver on the policy can have an SR22", nil, "PolicyHasDriverWithSR22", 67}
	quoteValidationErrorPolicyNil                                            = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyNilID, "the quote must have a policy", nil, "PolicyNil", 68}
	quoteValidationErrorPolicyVehiclesEmpty                                  = EnumQuoteValidationErrorItem{quoteValidationErrorPolicyVehiclesEmptyID, "the policy must have at least one vehicle", nil, "PolicyVehiclesEmpty", 69}
	quoteValidationErrorTotalIncidentsMoreThan4AtFault                       = EnumQuoteValidationErrorItem{quoteValidationErrorTotalIncidentsMoreThan4AtFaultID, "total at fault incidents on policy should be less than four in past 3 years", nil, "TotalIncidentsMoreThan4AtFault", 70}
	quoteValidationErrorUMBIAvailableOnlyWithUMPD                            = EnumQuoteValidationErrorItem{quoteValidationErrorUMBIAvailableOnlyWithUMPDID, "Uninsured/Underinsured motorist bodily injury coverage is only available with uninsured motorist property damage coverage", nil, "UMBIAvailableOnlyWithUMPD", 71}
	quoteValidationErrorUMBICantExceedBI                                     = EnumQuoteValidationErrorItem{quoteValidationErrorUMBICantExceedBIID, "Uninsured/Underinsured Motorist Bodily Injury limits may not exceed the Bodily Injury limits", nil, "UMBICantExceedBI", 72}
	quoteValidationErrorUMBIMustBeSameAsBI                                   = EnumQuoteValidationErrorItem{quoteValidationErrorUMBIMustBeSameAsBIID, "Uninsured/Underinsured Motorist Bodily Injury limits must be the same as the Bodily Injury limits", nil, "UMBIMustBeSameAsBI", 73}
	quoteValidationErrorUMPDAvailableOnlyWithUMBI                            = EnumQuoteValidationErrorItem{quoteValidationErrorUMPDAvailableOnlyWithUMBIID, "Uninsured/Underinsured motorist property damage coverage is only available with uninsured motorist bodily injury coverage", nil, "UMPDAvailableOnlyWithUMBI", 74}
	quoteValidationErrorUMPDCantExceedPD                                     = EnumQuoteValidationErrorItem{quoteValidationErrorUMPDCantExceedPDID, "Uninsured/Underinsured Motorist Property Damage limits may not exceed Property Damage Limits", nil, "UMPDCantExceedPD", 75}
	quoteValidationErrorUMPDMustBeSameAsPD                                   = EnumQuoteValidationErrorItem{quoteValidationErrorUMPDMustBeSameAsPDID, "Uninsured/Underinsured Motorist Property Damage limits must be the same as the Property Damage limits", nil, "UMPDMustBeSameAsPD", 76}
	quoteValidationErrorVehicleIDNotUnique                                   = EnumQuoteValidationErrorItem{quoteValidationErrorVehicleIDNotUniqueID, "no two vehicles on the policy can have same vehicle id", nil, "VehicleIDNotUnique", 77}
	quoteValidationErrorVehicleNotInsurable                                  = EnumQuoteValidationErrorItem{quoteValidationErrorVehicleNotInsurableID, "no vehicle on the policy can be marked doNotInsure in our VIN rating system", nil, "VehicleNotInsurable", 78}
	quoteValidationErrorVehicleParkedAtMoreThanOneAddress                    = EnumQuoteValidationErrorItem{quoteValidationErrorVehicleParkedAtMoreThanOneAddressID, "Vehicles Parked at more than 1 Address", nil, "VehicleParkedAtMoreThanOneAddress", 79}
	quoteValidationErrorVehiclePrimaryDriverInvalid                          = EnumQuoteValidationErrorItem{quoteValidationErrorVehiclePrimaryDriverInvalidID, "every vehicle's primary driver id must correspond to a driver on the policy", nil, "VehiclePrimaryDriverInvalid", 80}
	quoteValidationErrorVehiclePrimaryDriverNil                              = EnumQuoteValidationErrorItem{quoteValidationErrorVehiclePrimaryDriverNilID, "every vehicle must be assigned a primary driver", nil, "VehiclePrimaryDriverNil", 81}
	quoteValidationErrorVehicleVINInvalid                                    = EnumQuoteValidationErrorItem{quoteValidationErrorVehicleVINInvalidID, "each vehicle on the policy must have a valid VIN", nil, "VehicleVINInvalid", 82}
	quoteValidationErrorVehicleVINNil                                        = EnumQuoteValidationErrorItem{quoteValidationErrorVehicleVINNilID, "every vehicle on the policy must have a VIN", nil, "VehicleVINNil", 83}
	quoteValidationErrorCoverageValueNotValid                                = EnumQuoteValidationErrorItem{quoteValidationErrorCoverageValueNotValidID, "value provided for coverage is not a valid value", nil, "CoverageValueNotValid", 84}
	quoteValidationErrorEISPABodilyInjuryCovCoverageValueNotValid            = EnumQuoteValidationErrorItem{quoteValidationErrorEISPABodilyInjuryCovCoverageValueNotValidID, "value provided for for EISPABodilyInjuryCov coverage is not a valid value", nil, "EISPABodilyInjuryCovCoverageValueNotValid", 85}
	quoteValidationErrorEISPAPropertyDamageCovCoverageValueNotValid          = EnumQuoteValidationErrorItem{quoteValidationErrorEISPAPropertyDamageCovCoverageValueNotValidID, "value provided for EISPAPropertyDamageCov coverage is not a valid value", nil, "EISPAPropertyDamageCovCoverageValueNotValid", 86}
	quoteValidationErrorPAUMBICovCoverageValueNotValid                       = EnumQuoteValidationErrorItem{quoteValidationErrorPAUMBICovCoverageValueNotValidID, "value provided for PAUMBICov coverage is not a valid value", nil, "PAUMBICovCoverageValueNotValid", 87}
	quoteValidationErrorPAMedPayCovCoverageValueNotValid                     = EnumQuoteValidationErrorItem{quoteValidationErrorPAMedPayCovCoverageValueNotValidID, "value provided for PAMedPayCov coverage is not a valid value", nil, "PAMedPayCovCoverageValueNotValid", 88}
	quoteValidationErrorPAUMPDCovCoverageValueNotValid                       = EnumQuoteValidationErrorItem{quoteValidationErrorPAUMPDCovCoverageValueNotValidID, "value provided for PAUMPDCov coverage is not a valid value", nil, "PAUMPDCovCoverageValueNotValid", 89}
	quoteValidationErrorPAPIP_MDCoverageValueNotValid                        = EnumQuoteValidationErrorItem{quoteValidationErrorPAPIP_MDCoverageValueNotValidID, "value provided for PAPIP_MD coverage is not a valid value", nil, "PAPIP_MDCoverageValueNotValid", 90}
	quoteValidationErrorPAPIP_TXCoverageValueNotValid                        = EnumQuoteValidationErrorItem{quoteValidationErrorPAPIP_TXCoverageValueNotValidID, "value provided for PAPIP_TXCoverage coverage is not a valid value", nil, "PAPIP_TXCoverageValueNotValid", 91}
	quoteValidationErrorEISPALegalPlanCovCoverageValueNotValid               = EnumQuoteValidationErrorItem{quoteValidationErrorEISPALegalPlanCovCoverageValueNotValidID, "value provided for EISPALegalPlanCov coverage is not a valid value", nil, "EISPALegalPlanCovCoverageValueNotValid", 92}
	quoteValidationErrorEISPAIncomeLossCovCoverageValueNotValid              = EnumQuoteValidationErrorItem{quoteValidationErrorEISPAIncomeLossCovCoverageValueNotValidID, "value provided for EISPAIncomeLossCov coverage is not a valid value", nil, "EISPAIncomeLossCovCoverageValueNotValid", 93}
	quoteValidationErrorPAUM_GACovCoverageValueNotValid                      = EnumQuoteValidationErrorItem{quoteValidationErrorPAUM_GACovCoverageValueNotValidID, "value provided for PAUM_GACov coverage is not a valid value", nil, "PAUM_GACovCoverageValueNotValid", 94}
	quoteValidationErrorPAComprehensiveCovCoverageValueNotValid              = EnumQuoteValidationErrorItem{quoteValidationErrorPAComprehensiveCovCoverageValueNotValidID, "value provided for PAComprehensiveCov coverage is not a valid value", nil, "PAComprehensiveCovCoverageValueNotValid", 95}
	quoteValidationErrorPACollisionCovCoverageValueNotValid                  = EnumQuoteValidationErrorItem{quoteValidationErrorPACollisionCovCoverageValueNotValidID, "value provided for PACollisionCov coverage is not a valid value", nil, "PACollisionCovCoverageValueNotValid", 96}
	quoteValidationErrorEISPAUMPD_ILCovCoverageValueNotValid                 = EnumQuoteValidationErrorItem{quoteValidationErrorEISPAUMPD_ILCovCoverageValueNotValidID, "value provided for EISPAUMPD_ILCov coverage is not a valid value", nil, "EISPAUMPD_ILCovCoverageValueNotValid", 97}
	quoteValidationErrorEISPAUMPD_OHCovCoverageValueNotValid                 = EnumQuoteValidationErrorItem{quoteValidationErrorEISPAUMPD_OHCovCoverageValueNotValidID, "value provided for EISPAUMPD_OHCov coverage is not a valid value", nil, "EISPAUMPD_OHCovCoverageValueNotValid", 98}
	quoteValidationErrorPATowingLaborCovCoverageValueNotValid                = EnumQuoteValidationErrorItem{quoteValidationErrorPATowingLaborCovCoverageValueNotValidID, "value provided for PATowingLaborCov coverage is not a valid value", nil, "PATowingLaborCovCoverageValueNotValid", 99}
	quoteValidationErrorPARentalCovCoverageValueNotValid                     = EnumQuoteValidationErrorItem{quoteValidationErrorPARentalCovCoverageValueNotValidID, "value provided for PARentalCov coverage is not a valid value", nil, "PARentalCovCoverageValueNotValid", 100}
	quoteValidationErrorEISPALoanGapCovCoverageValueNotValid                 = EnumQuoteValidationErrorItem{quoteValidationErrorEISPALoanGapCovCoverageValueNotValidID, "value provided for EISPALoanGapCov coverage is not a valid value", nil, "EISPALoanGapCovCoverageValueNotValid", 101}
	quoteValidationErrorEISPACustEquipCovCoverageValueNotValid               = EnumQuoteValidationErrorItem{quoteValidationErrorEISPACustEquipCovCoverageValueNotValidID, "value provided for EISPACustEquipCov coverage is not a valid value", nil, "EISPACustEquipCovCoverageValueNotValid", 102}
	quoteValidationErrorPAVehicleMonitorCovCoverageValueNotValid             = EnumQuoteValidationErrorItem{quoteValidationErrorPAVehicleMonitorCovCoverageValueNotValidID, "value provided for PAVehicleMonitorCov coverage is not a valid value", nil, "PAVehicleMonitorCovCoverageValueNotValid", 103}
	quoteValidationErrorCoveragesNotValidForFr44                             = EnumQuoteValidationErrorItem{quoteValidationErrorCoveragesNotValidForFr44ID, "coverages that you selected is not acceptable for FR44 Policy", nil, "CoveragesNotValidForFr44", 104}
	quoteValidationErrorAllDriversMustHaveGender                             = EnumQuoteValidationErrorItem{quoteValidationErrorAllDriversMustHaveGenderID, "gender value can't be null", nil, "AllDriversMustHaveGender", 105}
	quoteValidationErrorRestrictedLicenceStateForPolicy                      = EnumQuoteValidationErrorItem{quoteValidationErrorRestrictedLicenceStateForPolicyID, "unfortunately, one of your driver's license does not meet the qualifications for a policy with Elephant. We require your driver's license to be issued by your state of residence to be eligible for a policy with us.", nil, "RestrictedLicenceStateForPolicy", 106}
)

// EnumQuoteValidationError is a collection of QuoteValidationError items
type EnumQuoteValidationError struct {
	Description string
	Items       []*EnumQuoteValidationErrorItem
	Name        string

	ApplicantAddressCityNil                              *EnumQuoteValidationErrorItem
	ApplicantAddressCountyNil                            *EnumQuoteValidationErrorItem
	ApplicantAddressIsPOBox                              *EnumQuoteValidationErrorItem
	ApplicantAddressLine1Nil                             *EnumQuoteValidationErrorItem
	ApplicantAddressNil                                  *EnumQuoteValidationErrorItem
	ApplicantAddressLine1ExceedsMaxLength                *EnumQuoteValidationErrorItem
	ApplicantAddressNotRatable                           *EnumQuoteValidationErrorItem
	ApplicantAddressPostalCodeNil                        *EnumQuoteValidationErrorItem
	ApplicantAddressStateInvalid                         *EnumQuoteValidationErrorItem
	ApplicantAddressStateNil                             *EnumQuoteValidationErrorItem
	ApplicantDateOfBirthNil                              *EnumQuoteValidationErrorItem
	ApplicantLicenseStatusInvalid                        *EnumQuoteValidationErrorItem
	ApplicantLicenseStatusPermit                         *EnumQuoteValidationErrorItem
	ApplicantMaritalStatusMarriedNoSpouse                *EnumQuoteValidationErrorItem
	ApplicantMaritalStatusMultipleSpouses                *EnumQuoteValidationErrorItem
	ApplicantMaritalStatusNil                            *EnumQuoteValidationErrorItem
	ApplicantMaritalStatusNotValid                       *EnumQuoteValidationErrorItem
	ApplicantNil                                         *EnumQuoteValidationErrorItem
	ApplicantNotDriver                                   *EnumQuoteValidationErrorItem
	ApplicantPhoneNumberInvalid                          *EnumQuoteValidationErrorItem
	ApplicantPhoneNumberNil                              *EnumQuoteValidationErrorItem
	ApplicantUnder18                                     *EnumQuoteValidationErrorItem
	ApplicantOccupationStatusNil                         *EnumQuoteValidationErrorItem
	ApplicantOccupationStatusUnemployed                  *EnumQuoteValidationErrorItem
	CannotCarryBothUMPDAndColl                           *EnumQuoteValidationErrorItem
	CantHoldLoanLeasePayOffOnOlderVehicle                *EnumQuoteValidationErrorItem
	CollRequiredForRent                                  *EnumQuoteValidationErrorItem
	CompCantExceedColl                                   *EnumQuoteValidationErrorItem
	CompCollNotAllowedOnOlderVehicle                     *EnumQuoteValidationErrorItem
	CompCollRequiredForCustomEquipment                   *EnumQuoteValidationErrorItem
	CompCollRequiredForLoan                              *EnumQuoteValidationErrorItem
	CompCollRequiredForRent                              *EnumQuoteValidationErrorItem
	CompRequiredForRent                                  *EnumQuoteValidationErrorItem
	CompRequiredToHoldColl                               *EnumQuoteValidationErrorItem
	DDRequiresColl500OrBetter                            *EnumQuoteValidationErrorItem
	DriverArrayCannotBeEmpty                             *EnumQuoteValidationErrorItem
	DriverHasBadLicenseStatus                            *EnumQuoteValidationErrorItem
	DriverIDNotUnique                                    *EnumQuoteValidationErrorItem
	DriverFirstNameExceedsMaxLength                      *EnumQuoteValidationErrorItem
	DriverLastNameExceedsMaxLength                       *EnumQuoteValidationErrorItem
	DriverIncidentsMoreThan1MinorConviction              *EnumQuoteValidationErrorItem
	DriverIncidentsMoreThan2MajorConvictions             *EnumQuoteValidationErrorItem
	DriverIncidentsMoreThan2MajorConvictionsIncludingDUI *EnumQuoteValidationErrorItem
	DriverIncidentsMoreThan4AtFault                      *EnumQuoteValidationErrorItem
	DriverIncidentsMoreThan4NotAtFault                   *EnumQuoteValidationErrorItem
	DriverIncidentsMoreThanOneDUI                        *EnumQuoteValidationErrorItem
	DriverPrimaryVehicleInvalid                          *EnumQuoteValidationErrorItem
	DriverPrimaryVehicleNil                              *EnumQuoteValidationErrorItem
	EveryDriverRequiresDriverID                          *EnumQuoteValidationErrorItem
	EveryVehicleRequiresVehicleID                        *EnumQuoteValidationErrorItem
	HasActivePolicy                                      *EnumQuoteValidationErrorItem
	HasBadDebtsFlag                                      *EnumQuoteValidationErrorItem
	HasMaterialMisRep                                    *EnumQuoteValidationErrorItem
	NoLoanLeasePayOffWithoutLienHolder                   *EnumQuoteValidationErrorItem
	PDCantExceedBI                                       *EnumQuoteValidationErrorItem
	PDCantExceedBIExceptForLowestLimit                   *EnumQuoteValidationErrorItem
	SameDayBindRuleError                                 *EnumQuoteValidationErrorItem
	PolicyCurrentInsuranceCarrierNil                     *EnumQuoteValidationErrorItem
	PolicyCurrentInsuranceLimitsNil                      *EnumQuoteValidationErrorItem
	PolicyCurrentInsurancePriorLapseNil                  *EnumQuoteValidationErrorItem
	PolicyCurrentInsuranceStatusNil                      *EnumQuoteValidationErrorItem
	PolicyCurrentInsuranceYearsWithNil                   *EnumQuoteValidationErrorItem
	PolicyDriversEmpty                                   *EnumQuoteValidationErrorItem
	PolicyEffectiveDateNil                               *EnumQuoteValidationErrorItem
	PolicyEffectiveDateNotMoreThan60DaysInFuture         *EnumQuoteValidationErrorItem
	PolicyEffectiveDatePast                              *EnumQuoteValidationErrorItem
	PolicyHasDriverWithSR22                              *EnumQuoteValidationErrorItem
	PolicyNil                                            *EnumQuoteValidationErrorItem
	PolicyVehiclesEmpty                                  *EnumQuoteValidationErrorItem
	TotalIncidentsMoreThan4AtFault                       *EnumQuoteValidationErrorItem
	UMBIAvailableOnlyWithUMPD                            *EnumQuoteValidationErrorItem
	UMBICantExceedBI                                     *EnumQuoteValidationErrorItem
	UMBIMustBeSameAsBI                                   *EnumQuoteValidationErrorItem
	UMPDAvailableOnlyWithUMBI                            *EnumQuoteValidationErrorItem
	UMPDCantExceedPD                                     *EnumQuoteValidationErrorItem
	UMPDMustBeSameAsPD                                   *EnumQuoteValidationErrorItem
	VehicleIDNotUnique                                   *EnumQuoteValidationErrorItem
	VehicleNotInsurable                                  *EnumQuoteValidationErrorItem
	VehicleParkedAtMoreThanOneAddress                    *EnumQuoteValidationErrorItem
	VehiclePrimaryDriverInvalid                          *EnumQuoteValidationErrorItem
	VehiclePrimaryDriverNil                              *EnumQuoteValidationErrorItem
	VehicleVINInvalid                                    *EnumQuoteValidationErrorItem
	VehicleVINNil                                        *EnumQuoteValidationErrorItem
	CoverageValueNotValid                                *EnumQuoteValidationErrorItem
	EISPABodilyInjuryCovCoverageValueNotValid            *EnumQuoteValidationErrorItem
	EISPAPropertyDamageCovCoverageValueNotValid          *EnumQuoteValidationErrorItem
	PAUMBICovCoverageValueNotValid                       *EnumQuoteValidationErrorItem
	PAMedPayCovCoverageValueNotValid                     *EnumQuoteValidationErrorItem
	PAUMPDCovCoverageValueNotValid                       *EnumQuoteValidationErrorItem
	PAPIP_MDCoverageValueNotValid                        *EnumQuoteValidationErrorItem
	PAPIP_TXCoverageValueNotValid                        *EnumQuoteValidationErrorItem
	EISPALegalPlanCovCoverageValueNotValid               *EnumQuoteValidationErrorItem
	EISPAIncomeLossCovCoverageValueNotValid              *EnumQuoteValidationErrorItem
	PAUM_GACovCoverageValueNotValid                      *EnumQuoteValidationErrorItem
	PAComprehensiveCovCoverageValueNotValid              *EnumQuoteValidationErrorItem
	PACollisionCovCoverageValueNotValid                  *EnumQuoteValidationErrorItem
	EISPAUMPD_ILCovCoverageValueNotValid                 *EnumQuoteValidationErrorItem
	EISPAUMPD_OHCovCoverageValueNotValid                 *EnumQuoteValidationErrorItem
	PATowingLaborCovCoverageValueNotValid                *EnumQuoteValidationErrorItem
	PARentalCovCoverageValueNotValid                     *EnumQuoteValidationErrorItem
	EISPALoanGapCovCoverageValueNotValid                 *EnumQuoteValidationErrorItem
	EISPACustEquipCovCoverageValueNotValid               *EnumQuoteValidationErrorItem
	PAVehicleMonitorCovCoverageValueNotValid             *EnumQuoteValidationErrorItem
	CoveragesNotValidForFr44                             *EnumQuoteValidationErrorItem
	AllDriversMustHaveGender                             *EnumQuoteValidationErrorItem
	RestrictedLicenceStateForPolicy                      *EnumQuoteValidationErrorItem

	itemDict map[string]*EnumQuoteValidationErrorItem
}

// QuoteValidationError is a public singleton instance of EnumQuoteValidationError
// representing validation errors that prevent us from returning a quote
var QuoteValidationError = &EnumQuoteValidationError{
	Description: "validation errors that prevent us from returning a quote",
	Items: []*EnumQuoteValidationErrorItem{
		&quoteValidationErrorApplicantAddressCityNil,
		&quoteValidationErrorApplicantAddressCountyNil,
		&quoteValidationErrorApplicantAddressIsPOBox,
		&quoteValidationErrorApplicantAddressLine1Nil,
		&quoteValidationErrorApplicantAddressNil,
		&quoteValidationErrorApplicantAddressLine1ExceedsMaxLength,
		&quoteValidationErrorApplicantAddressNotRatable,
		&quoteValidationErrorApplicantAddressPostalCodeNil,
		&quoteValidationErrorApplicantAddressStateInvalid,
		&quoteValidationErrorApplicantAddressStateNil,
		&quoteValidationErrorApplicantDateOfBirthNil,
		&quoteValidationErrorApplicantLicenseStatusInvalid,
		&quoteValidationErrorApplicantLicenseStatusPermit,
		&quoteValidationErrorApplicantMaritalStatusMarriedNoSpouse,
		&quoteValidationErrorApplicantMaritalStatusMultipleSpouses,
		&quoteValidationErrorApplicantMaritalStatusNil,
		&quoteValidationErrorApplicantMaritalStatusNotValid,
		&quoteValidationErrorApplicantNil,
		&quoteValidationErrorApplicantNotDriver,
		&quoteValidationErrorApplicantPhoneNumberInvalid,
		&quoteValidationErrorApplicantPhoneNumberNil,
		&quoteValidationErrorApplicantUnder18,
		&quoteValidationErrorApplicantOccupationStatusNil,
		&quoteValidationErrorApplicantOccupationStatusUnemployed,
		&quoteValidationErrorCannotCarryBothUMPDAndColl,
		&quoteValidationErrorCantHoldLoanLeasePayOffOnOlderVehicle,
		&quoteValidationErrorCollRequiredForRent,
		&quoteValidationErrorCompCantExceedColl,
		&quoteValidationErrorCompCollNotAllowedOnOlderVehicle,
		&quoteValidationErrorCompCollRequiredForCustomEquipment,
		&quoteValidationErrorCompCollRequiredForLoan,
		&quoteValidationErrorCompCollRequiredForRent,
		&quoteValidationErrorCompRequiredForRent,
		&quoteValidationErrorCompRequiredToHoldColl,
		&quoteValidationErrorDDRequiresColl500OrBetter,
		&quoteValidationErrorDriverArrayCannotBeEmpty,
		&quoteValidationErrorDriverHasBadLicenseStatus,
		&quoteValidationErrorDriverIDNotUnique,
		&quoteValidationErrorDriverFirstNameExceedsMaxLength,
		&quoteValidationErrorDriverLastNameExceedsMaxLength,
		&quoteValidationErrorDriverIncidentsMoreThan1MinorConviction,
		&quoteValidationErrorDriverIncidentsMoreThan2MajorConvictions,
		&quoteValidationErrorDriverIncidentsMoreThan2MajorConvictionsIncludingDUI,
		&quoteValidationErrorDriverIncidentsMoreThan4AtFault,
		&quoteValidationErrorDriverIncidentsMoreThan4NotAtFault,
		&quoteValidationErrorDriverIncidentsMoreThanOneDUI,
		&quoteValidationErrorDriverPrimaryVehicleInvalid,
		&quoteValidationErrorDriverPrimaryVehicleNil,
		&quoteValidationErrorEveryDriverRequiresDriverID,
		&quoteValidationErrorEveryVehicleRequiresVehicleID,
		&quoteValidationErrorHasActivePolicy,
		&quoteValidationErrorHasBadDebtsFlag,
		&quoteValidationErrorHasMaterialMisRep,
		&quoteValidationErrorNoLoanLeasePayOffWithoutLienHolder,
		&quoteValidationErrorPDCantExceedBI,
		&quoteValidationErrorPDCantExceedBIExceptForLowestLimit,
		&quoteValidationErrorSameDayBindRuleError,
		&quoteValidationErrorPolicyCurrentInsuranceCarrierNil,
		&quoteValidationErrorPolicyCurrentInsuranceLimitsNil,
		&quoteValidationErrorPolicyCurrentInsurancePriorLapseNil,
		&quoteValidationErrorPolicyCurrentInsuranceStatusNil,
		&quoteValidationErrorPolicyCurrentInsuranceYearsWithNil,
		&quoteValidationErrorPolicyDriversEmpty,
		&quoteValidationErrorPolicyEffectiveDateNil,
		&quoteValidationErrorPolicyEffectiveDateNotMoreThan60DaysInFuture,
		&quoteValidationErrorPolicyEffectiveDatePast,
		&quoteValidationErrorPolicyHasDriverWithSR22,
		&quoteValidationErrorPolicyNil,
		&quoteValidationErrorPolicyVehiclesEmpty,
		&quoteValidationErrorTotalIncidentsMoreThan4AtFault,
		&quoteValidationErrorUMBIAvailableOnlyWithUMPD,
		&quoteValidationErrorUMBICantExceedBI,
		&quoteValidationErrorUMBIMustBeSameAsBI,
		&quoteValidationErrorUMPDAvailableOnlyWithUMBI,
		&quoteValidationErrorUMPDCantExceedPD,
		&quoteValidationErrorUMPDMustBeSameAsPD,
		&quoteValidationErrorVehicleIDNotUnique,
		&quoteValidationErrorVehicleNotInsurable,
		&quoteValidationErrorVehicleParkedAtMoreThanOneAddress,
		&quoteValidationErrorVehiclePrimaryDriverInvalid,
		&quoteValidationErrorVehiclePrimaryDriverNil,
		&quoteValidationErrorVehicleVINInvalid,
		&quoteValidationErrorVehicleVINNil,
		&quoteValidationErrorCoverageValueNotValid,
		&quoteValidationErrorEISPABodilyInjuryCovCoverageValueNotValid,
		&quoteValidationErrorEISPAPropertyDamageCovCoverageValueNotValid,
		&quoteValidationErrorPAUMBICovCoverageValueNotValid,
		&quoteValidationErrorPAMedPayCovCoverageValueNotValid,
		&quoteValidationErrorPAUMPDCovCoverageValueNotValid,
		&quoteValidationErrorPAPIP_MDCoverageValueNotValid,
		&quoteValidationErrorPAPIP_TXCoverageValueNotValid,
		&quoteValidationErrorEISPALegalPlanCovCoverageValueNotValid,
		&quoteValidationErrorEISPAIncomeLossCovCoverageValueNotValid,
		&quoteValidationErrorPAUM_GACovCoverageValueNotValid,
		&quoteValidationErrorPAComprehensiveCovCoverageValueNotValid,
		&quoteValidationErrorPACollisionCovCoverageValueNotValid,
		&quoteValidationErrorEISPAUMPD_ILCovCoverageValueNotValid,
		&quoteValidationErrorEISPAUMPD_OHCovCoverageValueNotValid,
		&quoteValidationErrorPATowingLaborCovCoverageValueNotValid,
		&quoteValidationErrorPARentalCovCoverageValueNotValid,
		&quoteValidationErrorEISPALoanGapCovCoverageValueNotValid,
		&quoteValidationErrorEISPACustEquipCovCoverageValueNotValid,
		&quoteValidationErrorPAVehicleMonitorCovCoverageValueNotValid,
		&quoteValidationErrorCoveragesNotValidForFr44,
		&quoteValidationErrorAllDriversMustHaveGender,
		&quoteValidationErrorRestrictedLicenceStateForPolicy,
	},
	Name:                                                 "EnumQuoteValidationError",
	ApplicantAddressCityNil:                              &quoteValidationErrorApplicantAddressCityNil,
	ApplicantAddressCountyNil:                            &quoteValidationErrorApplicantAddressCountyNil,
	ApplicantAddressIsPOBox:                              &quoteValidationErrorApplicantAddressIsPOBox,
	ApplicantAddressLine1Nil:                             &quoteValidationErrorApplicantAddressLine1Nil,
	ApplicantAddressNil:                                  &quoteValidationErrorApplicantAddressNil,
	ApplicantAddressLine1ExceedsMaxLength:                &quoteValidationErrorApplicantAddressLine1ExceedsMaxLength,
	ApplicantAddressNotRatable:                           &quoteValidationErrorApplicantAddressNotRatable,
	ApplicantAddressPostalCodeNil:                        &quoteValidationErrorApplicantAddressPostalCodeNil,
	ApplicantAddressStateInvalid:                         &quoteValidationErrorApplicantAddressStateInvalid,
	ApplicantAddressStateNil:                             &quoteValidationErrorApplicantAddressStateNil,
	ApplicantDateOfBirthNil:                              &quoteValidationErrorApplicantDateOfBirthNil,
	ApplicantLicenseStatusInvalid:                        &quoteValidationErrorApplicantLicenseStatusInvalid,
	ApplicantLicenseStatusPermit:                         &quoteValidationErrorApplicantLicenseStatusPermit,
	ApplicantMaritalStatusMarriedNoSpouse:                &quoteValidationErrorApplicantMaritalStatusMarriedNoSpouse,
	ApplicantMaritalStatusMultipleSpouses:                &quoteValidationErrorApplicantMaritalStatusMultipleSpouses,
	ApplicantMaritalStatusNil:                            &quoteValidationErrorApplicantMaritalStatusNil,
	ApplicantMaritalStatusNotValid:                       &quoteValidationErrorApplicantMaritalStatusNotValid,
	ApplicantNil:                                         &quoteValidationErrorApplicantNil,
	ApplicantNotDriver:                                   &quoteValidationErrorApplicantNotDriver,
	ApplicantPhoneNumberInvalid:                          &quoteValidationErrorApplicantPhoneNumberInvalid,
	ApplicantPhoneNumberNil:                              &quoteValidationErrorApplicantPhoneNumberNil,
	ApplicantUnder18:                                     &quoteValidationErrorApplicantUnder18,
	ApplicantOccupationStatusNil:                         &quoteValidationErrorApplicantOccupationStatusNil,
	ApplicantOccupationStatusUnemployed:                  &quoteValidationErrorApplicantOccupationStatusUnemployed,
	CannotCarryBothUMPDAndColl:                           &quoteValidationErrorCannotCarryBothUMPDAndColl,
	CantHoldLoanLeasePayOffOnOlderVehicle:                &quoteValidationErrorCantHoldLoanLeasePayOffOnOlderVehicle,
	CollRequiredForRent:                                  &quoteValidationErrorCollRequiredForRent,
	CompCantExceedColl:                                   &quoteValidationErrorCompCantExceedColl,
	CompCollNotAllowedOnOlderVehicle:                     &quoteValidationErrorCompCollNotAllowedOnOlderVehicle,
	CompCollRequiredForCustomEquipment:                   &quoteValidationErrorCompCollRequiredForCustomEquipment,
	CompCollRequiredForLoan:                              &quoteValidationErrorCompCollRequiredForLoan,
	CompCollRequiredForRent:                              &quoteValidationErrorCompCollRequiredForRent,
	CompRequiredForRent:                                  &quoteValidationErrorCompRequiredForRent,
	CompRequiredToHoldColl:                               &quoteValidationErrorCompRequiredToHoldColl,
	DDRequiresColl500OrBetter:                            &quoteValidationErrorDDRequiresColl500OrBetter,
	DriverArrayCannotBeEmpty:                             &quoteValidationErrorDriverArrayCannotBeEmpty,
	DriverHasBadLicenseStatus:                            &quoteValidationErrorDriverHasBadLicenseStatus,
	DriverIDNotUnique:                                    &quoteValidationErrorDriverIDNotUnique,
	DriverFirstNameExceedsMaxLength:                      &quoteValidationErrorDriverFirstNameExceedsMaxLength,
	DriverLastNameExceedsMaxLength:                       &quoteValidationErrorDriverLastNameExceedsMaxLength,
	DriverIncidentsMoreThan1MinorConviction:              &quoteValidationErrorDriverIncidentsMoreThan1MinorConviction,
	DriverIncidentsMoreThan2MajorConvictions:             &quoteValidationErrorDriverIncidentsMoreThan2MajorConvictions,
	DriverIncidentsMoreThan2MajorConvictionsIncludingDUI: &quoteValidationErrorDriverIncidentsMoreThan2MajorConvictionsIncludingDUI,
	DriverIncidentsMoreThan4AtFault:                      &quoteValidationErrorDriverIncidentsMoreThan4AtFault,
	DriverIncidentsMoreThan4NotAtFault:                   &quoteValidationErrorDriverIncidentsMoreThan4NotAtFault,
	DriverIncidentsMoreThanOneDUI:                        &quoteValidationErrorDriverIncidentsMoreThanOneDUI,
	DriverPrimaryVehicleInvalid:                          &quoteValidationErrorDriverPrimaryVehicleInvalid,
	DriverPrimaryVehicleNil:                              &quoteValidationErrorDriverPrimaryVehicleNil,
	EveryDriverRequiresDriverID:                          &quoteValidationErrorEveryDriverRequiresDriverID,
	EveryVehicleRequiresVehicleID:                        &quoteValidationErrorEveryVehicleRequiresVehicleID,
	HasActivePolicy:                                      &quoteValidationErrorHasActivePolicy,
	HasBadDebtsFlag:                                      &quoteValidationErrorHasBadDebtsFlag,
	HasMaterialMisRep:                                    &quoteValidationErrorHasMaterialMisRep,
	NoLoanLeasePayOffWithoutLienHolder:                   &quoteValidationErrorNoLoanLeasePayOffWithoutLienHolder,
	PDCantExceedBI:                                       &quoteValidationErrorPDCantExceedBI,
	PDCantExceedBIExceptForLowestLimit:                   &quoteValidationErrorPDCantExceedBIExceptForLowestLimit,
	SameDayBindRuleError:                                 &quoteValidationErrorSameDayBindRuleError,
	PolicyCurrentInsuranceCarrierNil:                     &quoteValidationErrorPolicyCurrentInsuranceCarrierNil,
	PolicyCurrentInsuranceLimitsNil:                      &quoteValidationErrorPolicyCurrentInsuranceLimitsNil,
	PolicyCurrentInsurancePriorLapseNil:                  &quoteValidationErrorPolicyCurrentInsurancePriorLapseNil,
	PolicyCurrentInsuranceStatusNil:                      &quoteValidationErrorPolicyCurrentInsuranceStatusNil,
	PolicyCurrentInsuranceYearsWithNil:                   &quoteValidationErrorPolicyCurrentInsuranceYearsWithNil,
	PolicyDriversEmpty:                                   &quoteValidationErrorPolicyDriversEmpty,
	PolicyEffectiveDateNil:                               &quoteValidationErrorPolicyEffectiveDateNil,
	PolicyEffectiveDateNotMoreThan60DaysInFuture:         &quoteValidationErrorPolicyEffectiveDateNotMoreThan60DaysInFuture,
	PolicyEffectiveDatePast:                              &quoteValidationErrorPolicyEffectiveDatePast,
	PolicyHasDriverWithSR22:                              &quoteValidationErrorPolicyHasDriverWithSR22,
	PolicyNil:                                            &quoteValidationErrorPolicyNil,
	PolicyVehiclesEmpty:                                  &quoteValidationErrorPolicyVehiclesEmpty,
	TotalIncidentsMoreThan4AtFault:                       &quoteValidationErrorTotalIncidentsMoreThan4AtFault,
	UMBIAvailableOnlyWithUMPD:                            &quoteValidationErrorUMBIAvailableOnlyWithUMPD,
	UMBICantExceedBI:                                     &quoteValidationErrorUMBICantExceedBI,
	UMBIMustBeSameAsBI:                                   &quoteValidationErrorUMBIMustBeSameAsBI,
	UMPDAvailableOnlyWithUMBI:                            &quoteValidationErrorUMPDAvailableOnlyWithUMBI,
	UMPDCantExceedPD:                                     &quoteValidationErrorUMPDCantExceedPD,
	UMPDMustBeSameAsPD:                                   &quoteValidationErrorUMPDMustBeSameAsPD,
	VehicleIDNotUnique:                                   &quoteValidationErrorVehicleIDNotUnique,
	VehicleNotInsurable:                                  &quoteValidationErrorVehicleNotInsurable,
	VehicleParkedAtMoreThanOneAddress:                    &quoteValidationErrorVehicleParkedAtMoreThanOneAddress,
	VehiclePrimaryDriverInvalid:                          &quoteValidationErrorVehiclePrimaryDriverInvalid,
	VehiclePrimaryDriverNil:                              &quoteValidationErrorVehiclePrimaryDriverNil,
	VehicleVINInvalid:                                    &quoteValidationErrorVehicleVINInvalid,
	VehicleVINNil:                                        &quoteValidationErrorVehicleVINNil,
	CoverageValueNotValid:                                &quoteValidationErrorCoverageValueNotValid,
	EISPABodilyInjuryCovCoverageValueNotValid:            &quoteValidationErrorEISPABodilyInjuryCovCoverageValueNotValid,
	EISPAPropertyDamageCovCoverageValueNotValid:          &quoteValidationErrorEISPAPropertyDamageCovCoverageValueNotValid,
	PAUMBICovCoverageValueNotValid:                       &quoteValidationErrorPAUMBICovCoverageValueNotValid,
	PAMedPayCovCoverageValueNotValid:                     &quoteValidationErrorPAMedPayCovCoverageValueNotValid,
	PAUMPDCovCoverageValueNotValid:                       &quoteValidationErrorPAUMPDCovCoverageValueNotValid,
	PAPIP_MDCoverageValueNotValid:                        &quoteValidationErrorPAPIP_MDCoverageValueNotValid,
	PAPIP_TXCoverageValueNotValid:                        &quoteValidationErrorPAPIP_TXCoverageValueNotValid,
	EISPALegalPlanCovCoverageValueNotValid:               &quoteValidationErrorEISPALegalPlanCovCoverageValueNotValid,
	EISPAIncomeLossCovCoverageValueNotValid:              &quoteValidationErrorEISPAIncomeLossCovCoverageValueNotValid,
	PAUM_GACovCoverageValueNotValid:                      &quoteValidationErrorPAUM_GACovCoverageValueNotValid,
	PAComprehensiveCovCoverageValueNotValid:              &quoteValidationErrorPAComprehensiveCovCoverageValueNotValid,
	PACollisionCovCoverageValueNotValid:                  &quoteValidationErrorPACollisionCovCoverageValueNotValid,
	EISPAUMPD_ILCovCoverageValueNotValid:                 &quoteValidationErrorEISPAUMPD_ILCovCoverageValueNotValid,
	EISPAUMPD_OHCovCoverageValueNotValid:                 &quoteValidationErrorEISPAUMPD_OHCovCoverageValueNotValid,
	PATowingLaborCovCoverageValueNotValid:                &quoteValidationErrorPATowingLaborCovCoverageValueNotValid,
	PARentalCovCoverageValueNotValid:                     &quoteValidationErrorPARentalCovCoverageValueNotValid,
	EISPALoanGapCovCoverageValueNotValid:                 &quoteValidationErrorEISPALoanGapCovCoverageValueNotValid,
	EISPACustEquipCovCoverageValueNotValid:               &quoteValidationErrorEISPACustEquipCovCoverageValueNotValid,
	PAVehicleMonitorCovCoverageValueNotValid:             &quoteValidationErrorPAVehicleMonitorCovCoverageValueNotValid,
	CoveragesNotValidForFr44:                             &quoteValidationErrorCoveragesNotValidForFr44,
	AllDriversMustHaveGender:                             &quoteValidationErrorAllDriversMustHaveGender,
	RestrictedLicenceStateForPolicy:                      &quoteValidationErrorRestrictedLicenceStateForPolicy,

	itemDict: map[string]*EnumQuoteValidationErrorItem{
		strings.ToLower(string(quoteValidationErrorApplicantAddressCityNilID)):                              &quoteValidationErrorApplicantAddressCityNil,
		strings.ToLower(string(quoteValidationErrorApplicantAddressCountyNilID)):                            &quoteValidationErrorApplicantAddressCountyNil,
		strings.ToLower(string(quoteValidationErrorApplicantAddressIsPOBoxID)):                              &quoteValidationErrorApplicantAddressIsPOBox,
		strings.ToLower(string(quoteValidationErrorApplicantAddressLine1NilID)):                             &quoteValidationErrorApplicantAddressLine1Nil,
		strings.ToLower(string(quoteValidationErrorApplicantAddressNilID)):                                  &quoteValidationErrorApplicantAddressNil,
		strings.ToLower(string(quoteValidationErrorApplicantAddressLine1ExceedsMaxLengthID)):                &quoteValidationErrorApplicantAddressLine1ExceedsMaxLength,
		strings.ToLower(string(quoteValidationErrorApplicantAddressNotRatableID)):                           &quoteValidationErrorApplicantAddressNotRatable,
		strings.ToLower(string(quoteValidationErrorApplicantAddressPostalCodeNilID)):                        &quoteValidationErrorApplicantAddressPostalCodeNil,
		strings.ToLower(string(quoteValidationErrorApplicantAddressStateInvalidID)):                         &quoteValidationErrorApplicantAddressStateInvalid,
		strings.ToLower(string(quoteValidationErrorApplicantAddressStateNilID)):                             &quoteValidationErrorApplicantAddressStateNil,
		strings.ToLower(string(quoteValidationErrorApplicantDateOfBirthNilID)):                              &quoteValidationErrorApplicantDateOfBirthNil,
		strings.ToLower(string(quoteValidationErrorApplicantLicenseStatusInvalidID)):                        &quoteValidationErrorApplicantLicenseStatusInvalid,
		strings.ToLower(string(quoteValidationErrorApplicantLicenseStatusPermitID)):                         &quoteValidationErrorApplicantLicenseStatusPermit,
		strings.ToLower(string(quoteValidationErrorApplicantMaritalStatusMarriedNoSpouseID)):                &quoteValidationErrorApplicantMaritalStatusMarriedNoSpouse,
		strings.ToLower(string(quoteValidationErrorApplicantMaritalStatusMultipleSpousesID)):                &quoteValidationErrorApplicantMaritalStatusMultipleSpouses,
		strings.ToLower(string(quoteValidationErrorApplicantMaritalStatusNilID)):                            &quoteValidationErrorApplicantMaritalStatusNil,
		strings.ToLower(string(quoteValidationErrorApplicantMaritalStatusNotValidID)):                       &quoteValidationErrorApplicantMaritalStatusNotValid,
		strings.ToLower(string(quoteValidationErrorApplicantNilID)):                                         &quoteValidationErrorApplicantNil,
		strings.ToLower(string(quoteValidationErrorApplicantNotDriverID)):                                   &quoteValidationErrorApplicantNotDriver,
		strings.ToLower(string(quoteValidationErrorApplicantPhoneNumberInvalidID)):                          &quoteValidationErrorApplicantPhoneNumberInvalid,
		strings.ToLower(string(quoteValidationErrorApplicantPhoneNumberNilID)):                              &quoteValidationErrorApplicantPhoneNumberNil,
		strings.ToLower(string(quoteValidationErrorApplicantUnder18ID)):                                     &quoteValidationErrorApplicantUnder18,
		strings.ToLower(string(quoteValidationErrorApplicantOccupationStatusNilID)):                         &quoteValidationErrorApplicantOccupationStatusNil,
		strings.ToLower(string(quoteValidationErrorApplicantOccupationStatusUnemployedID)):                  &quoteValidationErrorApplicantOccupationStatusUnemployed,
		strings.ToLower(string(quoteValidationErrorCannotCarryBothUMPDAndCollID)):                           &quoteValidationErrorCannotCarryBothUMPDAndColl,
		strings.ToLower(string(quoteValidationErrorCantHoldLoanLeasePayOffOnOlderVehicleID)):                &quoteValidationErrorCantHoldLoanLeasePayOffOnOlderVehicle,
		strings.ToLower(string(quoteValidationErrorCollRequiredForRentID)):                                  &quoteValidationErrorCollRequiredForRent,
		strings.ToLower(string(quoteValidationErrorCompCantExceedCollID)):                                   &quoteValidationErrorCompCantExceedColl,
		strings.ToLower(string(quoteValidationErrorCompCollNotAllowedOnOlderVehicleID)):                     &quoteValidationErrorCompCollNotAllowedOnOlderVehicle,
		strings.ToLower(string(quoteValidationErrorCompCollRequiredForCustomEquipmentID)):                   &quoteValidationErrorCompCollRequiredForCustomEquipment,
		strings.ToLower(string(quoteValidationErrorCompCollRequiredForLoanID)):                              &quoteValidationErrorCompCollRequiredForLoan,
		strings.ToLower(string(quoteValidationErrorCompCollRequiredForRentID)):                              &quoteValidationErrorCompCollRequiredForRent,
		strings.ToLower(string(quoteValidationErrorCompRequiredForRentID)):                                  &quoteValidationErrorCompRequiredForRent,
		strings.ToLower(string(quoteValidationErrorCompRequiredToHoldCollID)):                               &quoteValidationErrorCompRequiredToHoldColl,
		strings.ToLower(string(quoteValidationErrorDDRequiresColl500OrBetterID)):                            &quoteValidationErrorDDRequiresColl500OrBetter,
		strings.ToLower(string(quoteValidationErrorDriverArrayCannotBeEmptyID)):                             &quoteValidationErrorDriverArrayCannotBeEmpty,
		strings.ToLower(string(quoteValidationErrorDriverHasBadLicenseStatusID)):                            &quoteValidationErrorDriverHasBadLicenseStatus,
		strings.ToLower(string(quoteValidationErrorDriverIDNotUniqueID)):                                    &quoteValidationErrorDriverIDNotUnique,
		strings.ToLower(string(quoteValidationErrorDriverFirstNameExceedsMaxLengthID)):                      &quoteValidationErrorDriverFirstNameExceedsMaxLength,
		strings.ToLower(string(quoteValidationErrorDriverLastNameExceedsMaxLengthID)):                       &quoteValidationErrorDriverLastNameExceedsMaxLength,
		strings.ToLower(string(quoteValidationErrorDriverIncidentsMoreThan1MinorConvictionID)):              &quoteValidationErrorDriverIncidentsMoreThan1MinorConviction,
		strings.ToLower(string(quoteValidationErrorDriverIncidentsMoreThan2MajorConvictionsID)):             &quoteValidationErrorDriverIncidentsMoreThan2MajorConvictions,
		strings.ToLower(string(quoteValidationErrorDriverIncidentsMoreThan2MajorConvictionsIncludingDUIID)): &quoteValidationErrorDriverIncidentsMoreThan2MajorConvictionsIncludingDUI,
		strings.ToLower(string(quoteValidationErrorDriverIncidentsMoreThan4AtFaultID)):                      &quoteValidationErrorDriverIncidentsMoreThan4AtFault,
		strings.ToLower(string(quoteValidationErrorDriverIncidentsMoreThan4NotAtFaultID)):                   &quoteValidationErrorDriverIncidentsMoreThan4NotAtFault,
		strings.ToLower(string(quoteValidationErrorDriverIncidentsMoreThanOneDUIID)):                        &quoteValidationErrorDriverIncidentsMoreThanOneDUI,
		strings.ToLower(string(quoteValidationErrorDriverPrimaryVehicleInvalidID)):                          &quoteValidationErrorDriverPrimaryVehicleInvalid,
		strings.ToLower(string(quoteValidationErrorDriverPrimaryVehicleNilID)):                              &quoteValidationErrorDriverPrimaryVehicleNil,
		strings.ToLower(string(quoteValidationErrorEveryDriverRequiresDriverIDID)):                          &quoteValidationErrorEveryDriverRequiresDriverID,
		strings.ToLower(string(quoteValidationErrorEveryVehicleRequiresVehicleIDID)):                        &quoteValidationErrorEveryVehicleRequiresVehicleID,
		strings.ToLower(string(quoteValidationErrorHasActivePolicyID)):                                      &quoteValidationErrorHasActivePolicy,
		strings.ToLower(string(quoteValidationErrorHasBadDebtsFlagID)):                                      &quoteValidationErrorHasBadDebtsFlag,
		strings.ToLower(string(quoteValidationErrorHasMaterialMisRepID)):                                    &quoteValidationErrorHasMaterialMisRep,
		strings.ToLower(string(quoteValidationErrorNoLoanLeasePayOffWithoutLienHolderID)):                   &quoteValidationErrorNoLoanLeasePayOffWithoutLienHolder,
		strings.ToLower(string(quoteValidationErrorPDCantExceedBIID)):                                       &quoteValidationErrorPDCantExceedBI,
		strings.ToLower(string(quoteValidationErrorPDCantExceedBIExceptForLowestLimitID)):                   &quoteValidationErrorPDCantExceedBIExceptForLowestLimit,
		strings.ToLower(string(quoteValidationErrorSameDayBindRuleErrorID)):                                 &quoteValidationErrorSameDayBindRuleError,
		strings.ToLower(string(quoteValidationErrorPolicyCurrentInsuranceCarrierNilID)):                     &quoteValidationErrorPolicyCurrentInsuranceCarrierNil,
		strings.ToLower(string(quoteValidationErrorPolicyCurrentInsuranceLimitsNilID)):                      &quoteValidationErrorPolicyCurrentInsuranceLimitsNil,
		strings.ToLower(string(quoteValidationErrorPolicyCurrentInsurancePriorLapseNilID)):                  &quoteValidationErrorPolicyCurrentInsurancePriorLapseNil,
		strings.ToLower(string(quoteValidationErrorPolicyCurrentInsuranceStatusNilID)):                      &quoteValidationErrorPolicyCurrentInsuranceStatusNil,
		strings.ToLower(string(quoteValidationErrorPolicyCurrentInsuranceYearsWithNilID)):                   &quoteValidationErrorPolicyCurrentInsuranceYearsWithNil,
		strings.ToLower(string(quoteValidationErrorPolicyDriversEmptyID)):                                   &quoteValidationErrorPolicyDriversEmpty,
		strings.ToLower(string(quoteValidationErrorPolicyEffectiveDateNilID)):                               &quoteValidationErrorPolicyEffectiveDateNil,
		strings.ToLower(string(quoteValidationErrorPolicyEffectiveDateNotMoreThan60DaysInFutureID)):         &quoteValidationErrorPolicyEffectiveDateNotMoreThan60DaysInFuture,
		strings.ToLower(string(quoteValidationErrorPolicyEffectiveDatePastID)):                              &quoteValidationErrorPolicyEffectiveDatePast,
		strings.ToLower(string(quoteValidationErrorPolicyHasDriverWithSR22ID)):                              &quoteValidationErrorPolicyHasDriverWithSR22,
		strings.ToLower(string(quoteValidationErrorPolicyNilID)):                                            &quoteValidationErrorPolicyNil,
		strings.ToLower(string(quoteValidationErrorPolicyVehiclesEmptyID)):                                  &quoteValidationErrorPolicyVehiclesEmpty,
		strings.ToLower(string(quoteValidationErrorTotalIncidentsMoreThan4AtFaultID)):                       &quoteValidationErrorTotalIncidentsMoreThan4AtFault,
		strings.ToLower(string(quoteValidationErrorUMBIAvailableOnlyWithUMPDID)):                            &quoteValidationErrorUMBIAvailableOnlyWithUMPD,
		strings.ToLower(string(quoteValidationErrorUMBICantExceedBIID)):                                     &quoteValidationErrorUMBICantExceedBI,
		strings.ToLower(string(quoteValidationErrorUMBIMustBeSameAsBIID)):                                   &quoteValidationErrorUMBIMustBeSameAsBI,
		strings.ToLower(string(quoteValidationErrorUMPDAvailableOnlyWithUMBIID)):                            &quoteValidationErrorUMPDAvailableOnlyWithUMBI,
		strings.ToLower(string(quoteValidationErrorUMPDCantExceedPDID)):                                     &quoteValidationErrorUMPDCantExceedPD,
		strings.ToLower(string(quoteValidationErrorUMPDMustBeSameAsPDID)):                                   &quoteValidationErrorUMPDMustBeSameAsPD,
		strings.ToLower(string(quoteValidationErrorVehicleIDNotUniqueID)):                                   &quoteValidationErrorVehicleIDNotUnique,
		strings.ToLower(string(quoteValidationErrorVehicleNotInsurableID)):                                  &quoteValidationErrorVehicleNotInsurable,
		strings.ToLower(string(quoteValidationErrorVehicleParkedAtMoreThanOneAddressID)):                    &quoteValidationErrorVehicleParkedAtMoreThanOneAddress,
		strings.ToLower(string(quoteValidationErrorVehiclePrimaryDriverInvalidID)):                          &quoteValidationErrorVehiclePrimaryDriverInvalid,
		strings.ToLower(string(quoteValidationErrorVehiclePrimaryDriverNilID)):                              &quoteValidationErrorVehiclePrimaryDriverNil,
		strings.ToLower(string(quoteValidationErrorVehicleVINInvalidID)):                                    &quoteValidationErrorVehicleVINInvalid,
		strings.ToLower(string(quoteValidationErrorVehicleVINNilID)):                                        &quoteValidationErrorVehicleVINNil,
		strings.ToLower(string(quoteValidationErrorCoverageValueNotValidID)):                                &quoteValidationErrorCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorEISPABodilyInjuryCovCoverageValueNotValidID)):            &quoteValidationErrorEISPABodilyInjuryCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorEISPAPropertyDamageCovCoverageValueNotValidID)):          &quoteValidationErrorEISPAPropertyDamageCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPAUMBICovCoverageValueNotValidID)):                       &quoteValidationErrorPAUMBICovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPAMedPayCovCoverageValueNotValidID)):                     &quoteValidationErrorPAMedPayCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPAUMPDCovCoverageValueNotValidID)):                       &quoteValidationErrorPAUMPDCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPAPIP_MDCoverageValueNotValidID)):                        &quoteValidationErrorPAPIP_MDCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPAPIP_TXCoverageValueNotValidID)):                        &quoteValidationErrorPAPIP_TXCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorEISPALegalPlanCovCoverageValueNotValidID)):               &quoteValidationErrorEISPALegalPlanCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorEISPAIncomeLossCovCoverageValueNotValidID)):              &quoteValidationErrorEISPAIncomeLossCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPAUM_GACovCoverageValueNotValidID)):                      &quoteValidationErrorPAUM_GACovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPAComprehensiveCovCoverageValueNotValidID)):              &quoteValidationErrorPAComprehensiveCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPACollisionCovCoverageValueNotValidID)):                  &quoteValidationErrorPACollisionCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorEISPAUMPD_ILCovCoverageValueNotValidID)):                 &quoteValidationErrorEISPAUMPD_ILCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorEISPAUMPD_OHCovCoverageValueNotValidID)):                 &quoteValidationErrorEISPAUMPD_OHCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPATowingLaborCovCoverageValueNotValidID)):                &quoteValidationErrorPATowingLaborCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPARentalCovCoverageValueNotValidID)):                     &quoteValidationErrorPARentalCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorEISPALoanGapCovCoverageValueNotValidID)):                 &quoteValidationErrorEISPALoanGapCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorEISPACustEquipCovCoverageValueNotValidID)):               &quoteValidationErrorEISPACustEquipCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorPAVehicleMonitorCovCoverageValueNotValidID)):             &quoteValidationErrorPAVehicleMonitorCovCoverageValueNotValid,
		strings.ToLower(string(quoteValidationErrorCoveragesNotValidForFr44ID)):                             &quoteValidationErrorCoveragesNotValidForFr44,
		strings.ToLower(string(quoteValidationErrorAllDriversMustHaveGenderID)):                             &quoteValidationErrorAllDriversMustHaveGender,
		strings.ToLower(string(quoteValidationErrorRestrictedLicenceStateForPolicyID)):                      &quoteValidationErrorRestrictedLicenceStateForPolicy,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumQuoteValidationError) ByID(id QuoteValidationErrorIdentifier) *EnumQuoteValidationErrorItem {
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
func (e *EnumQuoteValidationError) ByIDString(idx string) *EnumQuoteValidationErrorItem {
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
func (e *EnumQuoteValidationError) ByIndex(idx int) *EnumQuoteValidationErrorItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedQuoteValidationErrorID is a struct that is designed to replace a *QuoteValidationErrorID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *QuoteValidationErrorID it contains while being a better JSON citizen.
type ValidatedQuoteValidationErrorID struct {
	// id will point to a valid QuoteValidationErrorID, if possible
	// If id is nil, then ValidatedQuoteValidationErrorID.Valid() will return false.
	id *QuoteValidationErrorID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedQuoteValidationErrorID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedQuoteValidationErrorID
func (vi *ValidatedQuoteValidationErrorID) Clone() *ValidatedQuoteValidationErrorID {
	if vi == nil {
		return nil
	}

	var cid *QuoteValidationErrorID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedQuoteValidationErrorID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedQuoteValidationErrorIds represent the same QuoteValidationError
func (vi *ValidatedQuoteValidationErrorID) Equals(vj *ValidatedQuoteValidationErrorID) bool {
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

// Valid returns true if and only if the ValidatedQuoteValidationErrorID corresponds to a recognized QuoteValidationError
func (vi *ValidatedQuoteValidationErrorID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedQuoteValidationErrorID) ID() *QuoteValidationErrorID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedQuoteValidationErrorID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedQuoteValidationErrorID) ValidatedID() *ValidatedQuoteValidationErrorID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedQuoteValidationErrorID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedQuoteValidationErrorID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedQuoteValidationErrorID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedQuoteValidationErrorID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedQuoteValidationErrorID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := QuoteValidationErrorID(capString)
	item := QuoteValidationError.ByID(&id)
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

func (vi ValidatedQuoteValidationErrorID) String() string {
	return vi.ToIDString()
}

type QuoteValidationErrorIdentifier interface {
	ID() *QuoteValidationErrorID
	Valid() bool
}
