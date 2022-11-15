package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// PCErrorID uniquely identifies a particular PCError
type PCErrorID string

// Clone creates a safe, independent copy of a PCErrorID
func (i *PCErrorID) Clone() *PCErrorID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two PCErrorIds are equivalent
func (i *PCErrorID) Equals(j *PCErrorID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *PCErrorID that is either valid or nil
func (i *PCErrorID) ID() *PCErrorID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *PCErrorID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the PCErrorID corresponds to a recognized PCError
func (i *PCErrorID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return PCError.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *PCErrorID) ValidatedID() *ValidatedPCErrorID {
	if i != nil {
		return &ValidatedPCErrorID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *PCErrorID) MarshalJSON() ([]byte, error) {
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

func (i *PCErrorID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := PCErrorID(dataString)
	item := PCError.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	pCErrorMissingModelYearID                            PCErrorID = "5003Year"
	pCErrorInvalidDateFormatID                           PCErrorID = "100"
	pCErrorInvalidStateID                                PCErrorID = "1001"
	pCErrorAccountNumberMissingID                        PCErrorID = "1002"
	pCErrorAccountNotFoundID                             PCErrorID = "1003"
	pCErrorMapAccountErrorID                             PCErrorID = "1004"
	pCErrorSaveAccountErrorID                            PCErrorID = "1005"
	pCErrorHasExistingAccountErrorID                     PCErrorID = "1006"
	pCErrorSaveAccountHolderErrorID                      PCErrorID = "1007"
	pCErrorGenerateAccountErrorID                        PCErrorID = "1008"
	pCErrorCreateAccountErrorID                          PCErrorID = "1009"
	pCErrorGetRecentAccountErrorID                       PCErrorID = "1010"
	pCErrorAccountHasBadDebtID                           PCErrorID = "1011"
	pCErrorAccountLookupErrorID                          PCErrorID = "1012"
	pCErrorFirstNameMissingID                            PCErrorID = "2001"
	pCErrorInvalidEffectiveDateID                        PCErrorID = "2002"
	pCErrorDobMissingID                                  PCErrorID = "2003"
	pCErrorHomePhoneMissingID                            PCErrorID = "2004"
	pCErrorAddressLine1MissingID                         PCErrorID = "2005"
	pCErrorCityMissingID                                 PCErrorID = "2006"
	pCErrorStateMissingID                                PCErrorID = "2007"
	pCErrorPostalMissingID                               PCErrorID = "2008"
	pCErrorGenderMissingID                               PCErrorID = "2009"
	pCErrorLicenseStateOrNumberMissingID                 PCErrorID = "2010"
	pCErrorHighestEducationLevelMissingID                PCErrorID = "2011"
	pCErrorOccupationTitleMissingID                      PCErrorID = "2012"
	pCErrorCannotUpdateSameDayBindQuestionID             PCErrorID = "2013"
	pCErrorInvalidPhoneNumberID                          PCErrorID = "2014"
	pCErrorInvalidEmailID                                PCErrorID = "2015"
	pCErrorInvalidOnlineDiscountID                       PCErrorID = "2016"
	pCErrorInvalidWebToPhoneID                           PCErrorID = "2017"
	pCErrorValidatePaperlessDiscountID                   PCErrorID = "2018"
	pCErrorInvalidPolicyEffectiveDateID                  PCErrorID = "2019"
	pCErrorInvalidTransactiveEffecitveID                 PCErrorID = "2020"
	pCErrorMarriedWithNoSpouseID                         PCErrorID = "2021"
	pCErrorMarriedWithMultpleSpousesID                   PCErrorID = "2022"
	pCErrorEarlyBirdDiscountID                           PCErrorID = "2023"
	pCErrorCountyMissingID                               PCErrorID = "2024"
	pCErrorJobNumberMissingID                            PCErrorID = "3001"
	pCErrorPolicynumberNumberMissingID                   PCErrorID = "3003"
	pCErrorMapQualifiedDiscountsID                       PCErrorID = "3004"
	pCErrorCannotCreateSavePolicyperiodID                PCErrorID = "3005"
	pCErrorMapPolicyErrorID                              PCErrorID = "3006"
	pCErrorMapPolicyInitDataErrorID                      PCErrorID = "3007"
	pCErrorMapPolicyLineDetailsErrorID                   PCErrorID = "3008"
	pCErrorMapQuotedValuesErrorID                        PCErrorID = "3009"
	pCErrorSavePolicyErrorID                             PCErrorID = "3010"
	pCErrorRetrieveQuoteErrorID                          PCErrorID = "3013"
	pCErrorInvalidJobStatusID                            PCErrorID = "3101"
	pCErrorCannotEditJobID                               PCErrorID = "3102"
	pCErrorDeleteDriverErrorID                           PCErrorID = "3201"
	pCErrorDeleteManualIncidentsErrorID                  PCErrorID = "3202"
	pCErrorDeleteVehicleErrorID                          PCErrorID = "3203"
	pCErrorRatingStatusMissingID                         PCErrorID = "4001"
	pCErrorDriverUnacceptableRiskID                      PCErrorID = "4002"
	pCErrorCreditPullErrorID                             PCErrorID = "4003"
	pCErrorAdditionalDriversNeededID                     PCErrorID = "4004"
	pCErrorFr19DataMissingID                             PCErrorID = "4005"
	pCErrorAllDriversRequirePrimaryVehicleID             PCErrorID = "4006"
	pCErrorAdditionalDriverNeededID                      PCErrorID = "4007"
	pCErrorInvalidLicenseStatusID                        PCErrorID = "4008"
	pCErrorNoForeignLicenseID                            PCErrorID = "4009"
	pCErrorOutOfStateLicenseID                           PCErrorID = "4010"
	pCErrorPrimaryDriverRequiredID                       PCErrorID = "4011"
	pCErrorMissingDriverID                               PCErrorID = "4012"
	pCErrorMissingCreditPullID                           PCErrorID = "4013"
	pCErrorInvalidDriverIDID                             PCErrorID = "4014"
	pCErrorSaveDriverOccupationID                        PCErrorID = "4015"
	pCErrorMapDriverOccupationID                         PCErrorID = "4016"
	pCErrorMapDriverBasicInfoID                          PCErrorID = "4017"
	pCErrorMapDriverIncidentsID                          PCErrorID = "4018"
	pCErrorMapDriverErrorID                              PCErrorID = "4019"
	pCErrorMapDriverLicenseDetailsID                     PCErrorID = "4020"
	pCErrorMapDriversErrorID                             PCErrorID = "4021"
	pCErrorInvalidIncidentIDID                           PCErrorID = "4022"
	pCErrorSaveDriverBasicInfoID                         PCErrorID = "4023"
	pCErrorSaveDriverIncidentsID                         PCErrorID = "4024"
	pCErrorSaveDriverErrorID                             PCErrorID = "4025"
	pCErrorSaveDriverLicenseDetailsID                    PCErrorID = "4026"
	pCErrorMinConvictionExceededID                       PCErrorID = "4027"
	pCErrorMissingMakeID                                 PCErrorID = "5002"
	pCErrorMissingModelID                                PCErrorID = "5003"
	pCErrorMissingPrimaryUseID                           PCErrorID = "5004"
	pCErrorInvalidVehicleDriverIDID                      PCErrorID = "5005"
	pCErrorCannotAddVehicleToJobID                       PCErrorID = "5006"
	pCErrorInvalidVehicleIDID                            PCErrorID = "5007"
	pCErrorCannotRemoveVehicleID                         PCErrorID = "5008"
	pCErrorCannotFindVehicleForDriverID                  PCErrorID = "5009"
	pCErrorLienLeaseRequiredID                           PCErrorID = "5010"
	pCErrorAllVehicleRequirePrimaryDriverID              PCErrorID = "5011"
	pCErrorExoticVehicleID                               PCErrorID = "5012"
	pCErrorInvalidVinID                                  PCErrorID = "5013"
	pCErrorVehicleTypeRequiredID                         PCErrorID = "5014"
	pCErrorVehicleNewCostRequiredID                      PCErrorID = "5015"
	pCErrorVehiclePositiveCostRequiredID                 PCErrorID = "5016"
	pCErrorAtLeastOneVehicleRequiredID                   PCErrorID = "5017"
	pCErrorVinsMustBeUniqueForAllVehiclesID              PCErrorID = "5018"
	pCErrorVehiclePrimaryUseRequiredID                   PCErrorID = "5019"
	pCErrorSaveLoanLeinholderID                          PCErrorID = "5020"
	pCErrorSaveVehicleCoveragesID                        PCErrorID = "5021"
	pCErrorMapVehiclesErrorID                            PCErrorID = "5022"
	pCErrorSaveVehicleErrorID                            PCErrorID = "5023"
	pCErrorVehicleBrandedID                              PCErrorID = "5024"
	pCErrorVinIsoID                                      PCErrorID = "5025"
	pCErrorMissingPaperlessDiscountID                    PCErrorID = "6001"
	pCErrorMissingResidencyID                            PCErrorID = "6002"
	pCErrorMissingESignatureDiscountID                   PCErrorID = "6003"
	pCErrorMissingDoYouCurrentlyHaveInsureanceQuestionID PCErrorID = "6004"
	pCErrorMissingSourceOfBusinessID                     PCErrorID = "6005"
	pCErrorMissingCurrentCarrierNameID                   PCErrorID = "6006"
	pCErrorMissingCurrentCoverageEndDateID               PCErrorID = "6007"
	pCErrorMissingCurrentInjuryLimitsID                  PCErrorID = "6008"
	pCErrorRelationshipToInsuredMissingID                PCErrorID = "6009"
	pCErrorMissingLapseInCoverageQuestionID              PCErrorID = "6010"
	pCErrorMissingYearWithCurrentProviderQuestionID      PCErrorID = "6011"
	pCErrorCannotRemoveApplicantID                       PCErrorID = "6012"
	pCErrorSaveCoverageTermsID                           PCErrorID = "7002"
	pCErrorMapCoveragesErrorID                           PCErrorID = "7003"
	pCErrorInvalidCollisionDeductibleID                  PCErrorID = "7004"
	pCErrorDiminishingDeductibleOnVehiclesID             PCErrorID = "7005"
	pCErrorSavingManualIncidentsErrorID                  PCErrorID = "8000"
	pCErrorSavingClueIncidentsErrorID                    PCErrorID = "8001"
	pCErrorSavingMvrIncidentsErrorID                     PCErrorID = "8002"
	pCErrorUnknownSaveIncidentsErrorID                   PCErrorID = "8003"
	pCErrorQuoteRequiresPhotoInspectionID                PCErrorID = "8004"
	pCErrorMissingMvrReportID                            PCErrorID = "8005"
	pCErrorDuiLimitExceededID                            PCErrorID = "8006"
	pCErrorAtFaultLimitExceededID                        PCErrorID = "8007"
	pCErrorMinMajorConvictionLimitExceededID             PCErrorID = "8008"
	pCErrorMapClueIncidentsErrorID                       PCErrorID = "8009"
	pCErrorMapMvrIncidentsErrorID                        PCErrorID = "8010"
	pCErrorMapManualIncidentsErrorID                     PCErrorID = "8011"
	pCErrorMapIncidentsUnknownErrorID                    PCErrorID = "8012"
	pCErrorSavingIncidentsErrorID                        PCErrorID = "8013"
	pCErrorShouldPullCreditErrorID                       PCErrorID = "8014"
	pCErrorPullReportsErrorID                            PCErrorID = "8015"
	pCErrorMajorConvictionsLimitExceededID               PCErrorID = "8016"
	pCErrorPipClaimsLimitExceededID                      PCErrorID = "8017"
	pCErrorChargeableAccidentsLimitExceededID            PCErrorID = "8018"
	pCErrorJobNotInQuoteStatusID                         PCErrorID = "9000"
	pCErrorAccountHasActivePolicyID                      PCErrorID = "9001"
	pCErrorUnderwritingBlockingIssueID                   PCErrorID = "9002"
	pCErrorClueReportRequiredID                          PCErrorID = "9004"
	pCErrorBackDateErrorID                               PCErrorID = "9005"
	pCErrorExcludedDriverSendDocumentID                  PCErrorID = "9006"
	pCErrorExcludedDriverResendDocumentID                PCErrorID = "9007"
	pCErrorExcludedDriverFormNotSignedID                 PCErrorID = "9008"
	pCErrorAccidentDayOfBindID                           PCErrorID = "9009"
	pCErrorMissingPaymentplanID                          PCErrorID = "9010"
	pCErrorPolicyHasPhotoReviewID                        PCErrorID = "9011"
	pCErrorDocusignIntegrationErrorID                    PCErrorID = "9012"
	pCErrorCannotBindAccidentDayOfBindID                 PCErrorID = "9014"
	pCErrorDriverNotAssignedForLegalID                   PCErrorID = "9015"
	pCErrorIneligibleForFamilyLegalID                    PCErrorID = "9016"
	pCErrorNullDmsReasonID                               PCErrorID = "9017"
	pCErrorTerritoryRestrictionID                        PCErrorID = "9018"
	pCErrorAccountHasMaterialMisrepID                    PCErrorID = "9019"
	pCErrorIllegalBundleTransferExceptionID              PCErrorID = "9998"
	pCErrorConcurrentDataChangeExceptionID               PCErrorID = "9999"
	pCErrorGenerateProductsErrorID                       PCErrorID = "10001"
	pCErrorQuotePolicyErrorID                            PCErrorID = "10002"
	pCErrorBindPolicyErrorID                             PCErrorID = "10003"
	pCErrorFetchDocumentsErrorID                         PCErrorID = "10004"
	pCErrorInvalidPaymentPlanID                          PCErrorID = "10005"
	pCErrorInvalidPolicyDataID                           PCErrorID = "10010"
	pCErrorNoDocumentsID                                 PCErrorID = "10020"
	pCErrorNoJobFoundID                                  PCErrorID = "10021"
	pCErrorDeleteItemErrorID                             PCErrorID = "10022"
	pCErrorWorkorderNotQuotableID                        PCErrorID = "10023"
	pCErrorUnknownPaymentErrorID                         PCErrorID = "20000"
	pCErrorNoCreditCardNumberID                          PCErrorID = "20001"
	pCErrorFeePaidWithDownPaymentID                      PCErrorID = "20002"
	pCErrorCardDeclinedID                                PCErrorID = "20012"
	pCErrorInvalidCardNumberID                           PCErrorID = "20023"
	pCErrorInvalidExpirationID                           PCErrorID = "20024"
	pCErrorInvalidCardTypeID                             PCErrorID = "20025"
	pCErrorInsufficientFundsID                           PCErrorID = "20050"
	pCErrorInvalidAbaNumberID                            PCErrorID = "20101"
	pCErrorInvalidBankAccountNumberID                    PCErrorID = "20102"
	pCErrorInvalidBankNameID                             PCErrorID = "20103"
	pCErrorBankAccountTypeID                             PCErrorID = "20104"
	pCErrorLastNameMissingID                             PCErrorID = "200200"
	pCErrorMissingAccountContactID                       PCErrorID = "201313"
	pCErrorValidateFcraNoticeID                          PCErrorID = "201818"
	pCErrorPolicyHolderWithPermitLicenseID               PCErrorID = "2048"
)

// EnumPCErrorItem describes an entry in an enumeration of PCError
type EnumPCErrorItem struct {
	ID        PCErrorID         `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	DisplayName string
}

var (
	pCErrorMissingModelYear                            = EnumPCErrorItem{pCErrorMissingModelYearID, "MissingModelYear", map[string]string{"DisplayName": "Missing Model Year"}, "MissingModelYear", 0, "Missing Model Year"}
	pCErrorInvalidDateFormat                           = EnumPCErrorItem{pCErrorInvalidDateFormatID, "InvalidDateFormat", map[string]string{"DisplayName": "Invalid Date Format"}, "InvalidDateFormat", 100, "Invalid Date Format"}
	pCErrorInvalidState                                = EnumPCErrorItem{pCErrorInvalidStateID, "InvalidState", map[string]string{"DisplayName": "Invalid State"}, "InvalidState", 1001, "Invalid State"}
	pCErrorAccountNumberMissing                        = EnumPCErrorItem{pCErrorAccountNumberMissingID, "AccountNumberMissing", map[string]string{"DisplayName": "Account Number Missing"}, "AccountNumberMissing", 1002, "Account Number Missing"}
	pCErrorAccountNotFound                             = EnumPCErrorItem{pCErrorAccountNotFoundID, "AccountNotFound", map[string]string{"DisplayName": "Account Not Found"}, "AccountNotFound", 1003, "Account Not Found"}
	pCErrorMapAccountError                             = EnumPCErrorItem{pCErrorMapAccountErrorID, "MapAccountError", map[string]string{"DisplayName": "Map Account Error"}, "MapAccountError", 1004, "Map Account Error"}
	pCErrorSaveAccountError                            = EnumPCErrorItem{pCErrorSaveAccountErrorID, "SaveAccountError", map[string]string{"DisplayName": "Save Account Error"}, "SaveAccountError", 1005, "Save Account Error"}
	pCErrorHasExistingAccountError                     = EnumPCErrorItem{pCErrorHasExistingAccountErrorID, "HasExistingAccountError", map[string]string{"DisplayName": "Has Existing Account Error"}, "HasExistingAccountError", 1006, "Has Existing Account Error"}
	pCErrorSaveAccountHolderError                      = EnumPCErrorItem{pCErrorSaveAccountHolderErrorID, "SaveAccountHolderError", map[string]string{"DisplayName": "Save Account Holder Error"}, "SaveAccountHolderError", 1007, "Save Account Holder Error"}
	pCErrorGenerateAccountError                        = EnumPCErrorItem{pCErrorGenerateAccountErrorID, "GenerateAccountError", map[string]string{"DisplayName": "Generate Account Error"}, "GenerateAccountError", 1008, "Generate Account Error"}
	pCErrorCreateAccountError                          = EnumPCErrorItem{pCErrorCreateAccountErrorID, "CreateAccountError", map[string]string{"DisplayName": "Create Account Error"}, "CreateAccountError", 1009, "Create Account Error"}
	pCErrorGetRecentAccountError                       = EnumPCErrorItem{pCErrorGetRecentAccountErrorID, "GetRecentAccountError", map[string]string{"DisplayName": "Get Recent Account Error"}, "GetRecentAccountError", 1010, "Get Recent Account Error"}
	pCErrorAccountHasBadDebt                           = EnumPCErrorItem{pCErrorAccountHasBadDebtID, "AccountHasBadDebt", map[string]string{"DisplayName": "Account Has Bad Debt"}, "AccountHasBadDebt", 1011, "Account Has Bad Debt"}
	pCErrorAccountLookupError                          = EnumPCErrorItem{pCErrorAccountLookupErrorID, "AccountLookupError", map[string]string{"DisplayName": "Account Lookup Error"}, "AccountLookupError", 1012, "Account Lookup Error"}
	pCErrorFirstNameMissing                            = EnumPCErrorItem{pCErrorFirstNameMissingID, "FirstNameMissing", map[string]string{"DisplayName": "First Name Missing"}, "FirstNameMissing", 2001, "First Name Missing"}
	pCErrorInvalidEffectiveDate                        = EnumPCErrorItem{pCErrorInvalidEffectiveDateID, "InvalidEffectiveDate", map[string]string{"DisplayName": "Invalid Effective Date"}, "InvalidEffectiveDate", 2002, "Invalid Effective Date"}
	pCErrorDobMissing                                  = EnumPCErrorItem{pCErrorDobMissingID, "DobMissing", map[string]string{"DisplayName": "Dob Missing"}, "DobMissing", 2003, "Dob Missing"}
	pCErrorHomePhoneMissing                            = EnumPCErrorItem{pCErrorHomePhoneMissingID, "HomePhoneMissing", map[string]string{"DisplayName": "Home Phone Missing"}, "HomePhoneMissing", 2004, "Home Phone Missing"}
	pCErrorAddressLine1Missing                         = EnumPCErrorItem{pCErrorAddressLine1MissingID, "AddressLine1Missing", map[string]string{"DisplayName": "Address Line1 Missing"}, "AddressLine1Missing", 2005, "Address Line1 Missing"}
	pCErrorCityMissing                                 = EnumPCErrorItem{pCErrorCityMissingID, "CityMissing", map[string]string{"DisplayName": "City Missing"}, "CityMissing", 2006, "City Missing"}
	pCErrorStateMissing                                = EnumPCErrorItem{pCErrorStateMissingID, "StateMissing", map[string]string{"DisplayName": "State Missing"}, "StateMissing", 2007, "State Missing"}
	pCErrorPostalMissing                               = EnumPCErrorItem{pCErrorPostalMissingID, "PostalMissing", map[string]string{"DisplayName": "Postal Missing"}, "PostalMissing", 2008, "Postal Missing"}
	pCErrorGenderMissing                               = EnumPCErrorItem{pCErrorGenderMissingID, "GenderMissing", map[string]string{"DisplayName": "Gender Missing"}, "GenderMissing", 2009, "Gender Missing"}
	pCErrorLicenseStateOrNumberMissing                 = EnumPCErrorItem{pCErrorLicenseStateOrNumberMissingID, "LicenseStateOrNumberMissing", map[string]string{"DisplayName": "License State Or Number Missing"}, "LicenseStateOrNumberMissing", 2010, "License State Or Number Missing"}
	pCErrorHighestEducationLevelMissing                = EnumPCErrorItem{pCErrorHighestEducationLevelMissingID, "HighestEducationLevelMissing", map[string]string{"DisplayName": "Highest Education Level Missing"}, "HighestEducationLevelMissing", 2011, "Highest Education Level Missing"}
	pCErrorOccupationTitleMissing                      = EnumPCErrorItem{pCErrorOccupationTitleMissingID, "OccupationTitleMissing", map[string]string{"DisplayName": "Occupation Title Missing"}, "OccupationTitleMissing", 2012, "Occupation Title Missing"}
	pCErrorCannotUpdateSameDayBindQuestion             = EnumPCErrorItem{pCErrorCannotUpdateSameDayBindQuestionID, "CannotUpdateSameDayBindQuestion", map[string]string{"DisplayName": "Cannot Update Same Day Bind Question"}, "CannotUpdateSameDayBindQuestion", 2013, "Cannot Update Same Day Bind Question"}
	pCErrorInvalidPhoneNumber                          = EnumPCErrorItem{pCErrorInvalidPhoneNumberID, "InvalidPhoneNumber", map[string]string{"DisplayName": "Invalid Phone Number"}, "InvalidPhoneNumber", 2014, "Invalid Phone Number"}
	pCErrorInvalidEmail                                = EnumPCErrorItem{pCErrorInvalidEmailID, "InvalidEmail", map[string]string{"DisplayName": "Invalid Email"}, "InvalidEmail", 2015, "Invalid Email"}
	pCErrorInvalidOnlineDiscount                       = EnumPCErrorItem{pCErrorInvalidOnlineDiscountID, "InvalidOnlineDiscount", map[string]string{"DisplayName": "Invalid Online Discount"}, "InvalidOnlineDiscount", 2016, "Invalid Online Discount"}
	pCErrorInvalidWebToPhone                           = EnumPCErrorItem{pCErrorInvalidWebToPhoneID, "InvalidWebToPhone", map[string]string{"DisplayName": "Invalid Web To Phone"}, "InvalidWebToPhone", 2017, "Invalid Web To Phone"}
	pCErrorValidatePaperlessDiscount                   = EnumPCErrorItem{pCErrorValidatePaperlessDiscountID, "ValidatePaperlessDiscount", map[string]string{"DisplayName": "Validate Paperless Discount"}, "ValidatePaperlessDiscount", 2018, "Validate Paperless Discount"}
	pCErrorInvalidPolicyEffectiveDate                  = EnumPCErrorItem{pCErrorInvalidPolicyEffectiveDateID, "InvalidPolicyEffectiveDate", map[string]string{"DisplayName": "Invalid Policy Effective Date"}, "InvalidPolicyEffectiveDate", 2019, "Invalid Policy Effective Date"}
	pCErrorInvalidTransactiveEffecitve                 = EnumPCErrorItem{pCErrorInvalidTransactiveEffecitveID, "InvalidTransactiveEffecitve", map[string]string{"DisplayName": "Invalid Transactive Effecitve"}, "InvalidTransactiveEffecitve", 2020, "Invalid Transactive Effecitve"}
	pCErrorMarriedWithNoSpouse                         = EnumPCErrorItem{pCErrorMarriedWithNoSpouseID, "MarriedWithNoSpouse", map[string]string{"DisplayName": "Married With No Spouse"}, "MarriedWithNoSpouse", 2021, "Married With No Spouse"}
	pCErrorMarriedWithMultpleSpouses                   = EnumPCErrorItem{pCErrorMarriedWithMultpleSpousesID, "MarriedWithMultpleSpouses", map[string]string{"DisplayName": "Married With Multple Spouses"}, "MarriedWithMultpleSpouses", 2022, "Married With Multple Spouses"}
	pCErrorEarlyBirdDiscount                           = EnumPCErrorItem{pCErrorEarlyBirdDiscountID, "EarlyBirdDiscount", map[string]string{"DisplayName": "Early Bird Discount"}, "EarlyBirdDiscount", 2023, "Early Bird Discount"}
	pCErrorCountyMissing                               = EnumPCErrorItem{pCErrorCountyMissingID, "CountyMissing", map[string]string{"DisplayName": "County Missing"}, "CountyMissing", 2024, "County Missing"}
	pCErrorJobNumberMissing                            = EnumPCErrorItem{pCErrorJobNumberMissingID, "JobNumberMissing", map[string]string{"DisplayName": "Job Number Missing"}, "JobNumberMissing", 3001, "Job Number Missing"}
	pCErrorPolicynumberNumberMissing                   = EnumPCErrorItem{pCErrorPolicynumberNumberMissingID, "PolicynumberNumberMissing", map[string]string{"DisplayName": "Policynumber Number Missing"}, "PolicynumberNumberMissing", 3003, "Policynumber Number Missing"}
	pCErrorMapQualifiedDiscounts                       = EnumPCErrorItem{pCErrorMapQualifiedDiscountsID, "MapQualifiedDiscounts", map[string]string{"DisplayName": "Map Qualified Discounts"}, "MapQualifiedDiscounts", 3004, "Map Qualified Discounts"}
	pCErrorCannotCreateSavePolicyperiod                = EnumPCErrorItem{pCErrorCannotCreateSavePolicyperiodID, "CannotCreateSavePolicyperiod", map[string]string{"DisplayName": "Cannot Create Save Policyperiod"}, "CannotCreateSavePolicyperiod", 3005, "Cannot Create Save Policyperiod"}
	pCErrorMapPolicyError                              = EnumPCErrorItem{pCErrorMapPolicyErrorID, "MapPolicyError", map[string]string{"DisplayName": "Map Policy Error"}, "MapPolicyError", 3006, "Map Policy Error"}
	pCErrorMapPolicyInitDataError                      = EnumPCErrorItem{pCErrorMapPolicyInitDataErrorID, "MapPolicyInitDataError", map[string]string{"DisplayName": "Map Policy Init Data Error"}, "MapPolicyInitDataError", 3007, "Map Policy Init Data Error"}
	pCErrorMapPolicyLineDetailsError                   = EnumPCErrorItem{pCErrorMapPolicyLineDetailsErrorID, "MapPolicyLineDetailsError", map[string]string{"DisplayName": "Map Policy Line Details Error"}, "MapPolicyLineDetailsError", 3008, "Map Policy Line Details Error"}
	pCErrorMapQuotedValuesError                        = EnumPCErrorItem{pCErrorMapQuotedValuesErrorID, "MapQuotedValuesError", map[string]string{"DisplayName": "Map Quoted Values Error"}, "MapQuotedValuesError", 3009, "Map Quoted Values Error"}
	pCErrorSavePolicyError                             = EnumPCErrorItem{pCErrorSavePolicyErrorID, "SavePolicyError", map[string]string{"DisplayName": "Save Policy Error"}, "SavePolicyError", 3010, "Save Policy Error"}
	pCErrorRetrieveQuoteError                          = EnumPCErrorItem{pCErrorRetrieveQuoteErrorID, "RetrieveQuoteError", map[string]string{"DisplayName": "Retrieve Quote Error"}, "RetrieveQuoteError", 3013, "Retrieve Quote Error"}
	pCErrorInvalidJobStatus                            = EnumPCErrorItem{pCErrorInvalidJobStatusID, "InvalidJobStatus", map[string]string{"DisplayName": "Invalid Job Status"}, "InvalidJobStatus", 3101, "Invalid Job Status"}
	pCErrorCannotEditJob                               = EnumPCErrorItem{pCErrorCannotEditJobID, "CannotEditJob", map[string]string{"DisplayName": "Cannot Edit Job"}, "CannotEditJob", 3102, "Cannot Edit Job"}
	pCErrorDeleteDriverError                           = EnumPCErrorItem{pCErrorDeleteDriverErrorID, "DeleteDriverError", map[string]string{"DisplayName": "Delete Driver Error"}, "DeleteDriverError", 3201, "Delete Driver Error"}
	pCErrorDeleteManualIncidentsError                  = EnumPCErrorItem{pCErrorDeleteManualIncidentsErrorID, "DeleteManualIncidentsError", map[string]string{"DisplayName": "Delete Manual Incidents Error"}, "DeleteManualIncidentsError", 3202, "Delete Manual Incidents Error"}
	pCErrorDeleteVehicleError                          = EnumPCErrorItem{pCErrorDeleteVehicleErrorID, "DeleteVehicleError", map[string]string{"DisplayName": "Delete Vehicle Error"}, "DeleteVehicleError", 3203, "Delete Vehicle Error"}
	pCErrorRatingStatusMissing                         = EnumPCErrorItem{pCErrorRatingStatusMissingID, "RatingStatusMissing", map[string]string{"DisplayName": "Rating Status Missing"}, "RatingStatusMissing", 4001, "Rating Status Missing"}
	pCErrorDriverUnacceptableRisk                      = EnumPCErrorItem{pCErrorDriverUnacceptableRiskID, "DriverUnacceptableRisk", map[string]string{"DisplayName": "Driver Unacceptable Risk"}, "DriverUnacceptableRisk", 4002, "Driver Unacceptable Risk"}
	pCErrorCreditPullError                             = EnumPCErrorItem{pCErrorCreditPullErrorID, "CreditPullError", map[string]string{"DisplayName": "Credit Pull Error"}, "CreditPullError", 4003, "Credit Pull Error"}
	pCErrorAdditionalDriversNeeded                     = EnumPCErrorItem{pCErrorAdditionalDriversNeededID, "AdditionalDriversNeeded", map[string]string{"DisplayName": "Additional Drivers Needed"}, "AdditionalDriversNeeded", 4004, "Additional Drivers Needed"}
	pCErrorFr19DataMissing                             = EnumPCErrorItem{pCErrorFr19DataMissingID, "Fr19DataMissing", map[string]string{"DisplayName": "Fr19 Data Missing"}, "Fr19DataMissing", 4005, "Fr19 Data Missing"}
	pCErrorAllDriversRequirePrimaryVehicle             = EnumPCErrorItem{pCErrorAllDriversRequirePrimaryVehicleID, "AllDriversRequirePrimaryVehicle", map[string]string{"DisplayName": "All Drivers Require Primary Vehicle"}, "AllDriversRequirePrimaryVehicle", 4006, "All Drivers Require Primary Vehicle"}
	pCErrorAdditionalDriverNeeded                      = EnumPCErrorItem{pCErrorAdditionalDriverNeededID, "AdditionalDriverNeeded", map[string]string{"DisplayName": "Additional Driver Needed"}, "AdditionalDriverNeeded", 4007, "Additional Driver Needed"}
	pCErrorInvalidLicenseStatus                        = EnumPCErrorItem{pCErrorInvalidLicenseStatusID, "InvalidLicenseStatus", map[string]string{"DisplayName": "Invalid License Status"}, "InvalidLicenseStatus", 4008, "Invalid License Status"}
	pCErrorNoForeignLicense                            = EnumPCErrorItem{pCErrorNoForeignLicenseID, "NoForeignLicense", map[string]string{"DisplayName": "No Foreign License"}, "NoForeignLicense", 4009, "No Foreign License"}
	pCErrorOutOfStateLicense                           = EnumPCErrorItem{pCErrorOutOfStateLicenseID, "OutOfStateLicense", map[string]string{"DisplayName": "Out Of State License"}, "OutOfStateLicense", 4010, "Out Of State License"}
	pCErrorPrimaryDriverRequired                       = EnumPCErrorItem{pCErrorPrimaryDriverRequiredID, "PrimaryDriverRequired", map[string]string{"DisplayName": "Primary Driver Required"}, "PrimaryDriverRequired", 4011, "Primary Driver Required"}
	pCErrorMissingDriver                               = EnumPCErrorItem{pCErrorMissingDriverID, "MissingDriver", map[string]string{"DisplayName": "Missing Driver"}, "MissingDriver", 4012, "Missing Driver"}
	pCErrorMissingCreditPull                           = EnumPCErrorItem{pCErrorMissingCreditPullID, "MissingCreditPull", map[string]string{"DisplayName": "Missing Credit Pull"}, "MissingCreditPull", 4013, "Missing Credit Pull"}
	pCErrorInvalidDriverID                             = EnumPCErrorItem{pCErrorInvalidDriverIDID, "InvalidDriverId", map[string]string{"DisplayName": "Invalid Driver Id"}, "InvalidDriverID", 4014, "Invalid Driver Id"}
	pCErrorSaveDriverOccupation                        = EnumPCErrorItem{pCErrorSaveDriverOccupationID, "SaveDriverOccupation", map[string]string{"DisplayName": "Save Driver Occupation"}, "SaveDriverOccupation", 4015, "Save Driver Occupation"}
	pCErrorMapDriverOccupation                         = EnumPCErrorItem{pCErrorMapDriverOccupationID, "MapDriverOccupation", map[string]string{"DisplayName": "Map Driver Occupation"}, "MapDriverOccupation", 4016, "Map Driver Occupation"}
	pCErrorMapDriverBasicInfo                          = EnumPCErrorItem{pCErrorMapDriverBasicInfoID, "MapDriverBasicInfo", map[string]string{"DisplayName": "Map Driver Basic Info"}, "MapDriverBasicInfo", 4017, "Map Driver Basic Info"}
	pCErrorMapDriverIncidents                          = EnumPCErrorItem{pCErrorMapDriverIncidentsID, "MapDriverIncidents", map[string]string{"DisplayName": "Map Driver Incidents"}, "MapDriverIncidents", 4018, "Map Driver Incidents"}
	pCErrorMapDriverError                              = EnumPCErrorItem{pCErrorMapDriverErrorID, "MapDriverError", map[string]string{"DisplayName": "Map Driver Error"}, "MapDriverError", 4019, "Map Driver Error"}
	pCErrorMapDriverLicenseDetails                     = EnumPCErrorItem{pCErrorMapDriverLicenseDetailsID, "MapDriverLicenseDetails", map[string]string{"DisplayName": "Map Driver License Details"}, "MapDriverLicenseDetails", 4020, "Map Driver License Details"}
	pCErrorMapDriversError                             = EnumPCErrorItem{pCErrorMapDriversErrorID, "MapDriversError", map[string]string{"DisplayName": "Map Drivers Error"}, "MapDriversError", 4021, "Map Drivers Error"}
	pCErrorInvalidIncidentID                           = EnumPCErrorItem{pCErrorInvalidIncidentIDID, "InvalidIncidentId", map[string]string{"DisplayName": "Invalid Incident Id"}, "InvalidIncidentID", 4022, "Invalid Incident Id"}
	pCErrorSaveDriverBasicInfo                         = EnumPCErrorItem{pCErrorSaveDriverBasicInfoID, "SaveDriverBasicInfo", map[string]string{"DisplayName": "Save Driver Basic Info"}, "SaveDriverBasicInfo", 4023, "Save Driver Basic Info"}
	pCErrorSaveDriverIncidents                         = EnumPCErrorItem{pCErrorSaveDriverIncidentsID, "SaveDriverIncidents", map[string]string{"DisplayName": "Save Driver Incidents"}, "SaveDriverIncidents", 4024, "Save Driver Incidents"}
	pCErrorSaveDriverError                             = EnumPCErrorItem{pCErrorSaveDriverErrorID, "SaveDriverError", map[string]string{"DisplayName": "Save Driver Error"}, "SaveDriverError", 4025, "Save Driver Error"}
	pCErrorSaveDriverLicenseDetails                    = EnumPCErrorItem{pCErrorSaveDriverLicenseDetailsID, "SaveDriverLicenseDetails", map[string]string{"DisplayName": "Save Driver License Details"}, "SaveDriverLicenseDetails", 4026, "Save Driver License Details"}
	pCErrorMinConvictionExceeded                       = EnumPCErrorItem{pCErrorMinConvictionExceededID, "MinConvictionExceeded", map[string]string{"DisplayName": "Min Conviction Exceeded"}, "MinConvictionExceeded", 4027, "Min Conviction Exceeded"}
	pCErrorMissingMake                                 = EnumPCErrorItem{pCErrorMissingMakeID, "MissingMake", map[string]string{"DisplayName": "Missing Make"}, "MissingMake", 5002, "Missing Make"}
	pCErrorMissingModel                                = EnumPCErrorItem{pCErrorMissingModelID, "MissingModel", map[string]string{"DisplayName": "Missing Model"}, "MissingModel", 5003, "Missing Model"}
	pCErrorMissingPrimaryUse                           = EnumPCErrorItem{pCErrorMissingPrimaryUseID, "MissingPrimaryUse", map[string]string{"DisplayName": "Missing Primary Use"}, "MissingPrimaryUse", 5004, "Missing Primary Use"}
	pCErrorInvalidVehicleDriverID                      = EnumPCErrorItem{pCErrorInvalidVehicleDriverIDID, "InvalidVehicleDriverId", map[string]string{"DisplayName": "Invalid Vehicle Driver Id"}, "InvalidVehicleDriverID", 5005, "Invalid Vehicle Driver Id"}
	pCErrorCannotAddVehicleToJob                       = EnumPCErrorItem{pCErrorCannotAddVehicleToJobID, "CannotAddVehicleToJob", map[string]string{"DisplayName": "Cannot Add Vehicle To Job"}, "CannotAddVehicleToJob", 5006, "Cannot Add Vehicle To Job"}
	pCErrorInvalidVehicleID                            = EnumPCErrorItem{pCErrorInvalidVehicleIDID, "InvalidVehicleId", map[string]string{"DisplayName": "Invalid Vehicle Id"}, "InvalidVehicleID", 5007, "Invalid Vehicle Id"}
	pCErrorCannotRemoveVehicle                         = EnumPCErrorItem{pCErrorCannotRemoveVehicleID, "CannotRemoveVehicle", map[string]string{"DisplayName": "Cannot Remove Vehicle"}, "CannotRemoveVehicle", 5008, "Cannot Remove Vehicle"}
	pCErrorCannotFindVehicleForDriver                  = EnumPCErrorItem{pCErrorCannotFindVehicleForDriverID, "CannotFindVehicleForDriver", map[string]string{"DisplayName": "Cannot Find Vehicle For Driver"}, "CannotFindVehicleForDriver", 5009, "Cannot Find Vehicle For Driver"}
	pCErrorLienLeaseRequired                           = EnumPCErrorItem{pCErrorLienLeaseRequiredID, "LienLeaseRequired", map[string]string{"DisplayName": "Lien Lease Required"}, "LienLeaseRequired", 5010, "Lien Lease Required"}
	pCErrorAllVehicleRequirePrimaryDriver              = EnumPCErrorItem{pCErrorAllVehicleRequirePrimaryDriverID, "AllVehicleRequirePrimaryDriver", map[string]string{"DisplayName": "All Vehicle Require Primary Driver"}, "AllVehicleRequirePrimaryDriver", 5011, "All Vehicle Require Primary Driver"}
	pCErrorExoticVehicle                               = EnumPCErrorItem{pCErrorExoticVehicleID, "ExoticVehicle", map[string]string{"DisplayName": "Exotic Vehicle"}, "ExoticVehicle", 5012, "Exotic Vehicle"}
	pCErrorInvalidVin                                  = EnumPCErrorItem{pCErrorInvalidVinID, "InvalidVin", map[string]string{"DisplayName": "Invalid Vin"}, "InvalidVin", 5013, "Invalid Vin"}
	pCErrorVehicleTypeRequired                         = EnumPCErrorItem{pCErrorVehicleTypeRequiredID, "VehicleTypeRequired", map[string]string{"DisplayName": "Vehicle Type Required"}, "VehicleTypeRequired", 5014, "Vehicle Type Required"}
	pCErrorVehicleNewCostRequired                      = EnumPCErrorItem{pCErrorVehicleNewCostRequiredID, "VehicleNewCostRequired", map[string]string{"DisplayName": "Vehicle New Cost Required"}, "VehicleNewCostRequired", 5015, "Vehicle New Cost Required"}
	pCErrorVehiclePositiveCostRequired                 = EnumPCErrorItem{pCErrorVehiclePositiveCostRequiredID, "VehiclePositiveCostRequired", map[string]string{"DisplayName": "Vehicle Positive Cost Required"}, "VehiclePositiveCostRequired", 5016, "Vehicle Positive Cost Required"}
	pCErrorAtLeastOneVehicleRequired                   = EnumPCErrorItem{pCErrorAtLeastOneVehicleRequiredID, "AtLeastOneVehicleRequired", map[string]string{"DisplayName": "At Least One Vehicle Required"}, "AtLeastOneVehicleRequired", 5017, "At Least One Vehicle Required"}
	pCErrorVinsMustBeUniqueForAllVehicles              = EnumPCErrorItem{pCErrorVinsMustBeUniqueForAllVehiclesID, "VinsMustBeUniqueForAllVehicles", map[string]string{"DisplayName": "Vins Must Be Unique For All Vehicles"}, "VinsMustBeUniqueForAllVehicles", 5018, "Vins Must Be Unique For All Vehicles"}
	pCErrorVehiclePrimaryUseRequired                   = EnumPCErrorItem{pCErrorVehiclePrimaryUseRequiredID, "VehiclePrimaryUseRequired", map[string]string{"DisplayName": "Vehicle Primary Use Required"}, "VehiclePrimaryUseRequired", 5019, "Vehicle Primary Use Required"}
	pCErrorSaveLoanLeinholder                          = EnumPCErrorItem{pCErrorSaveLoanLeinholderID, "SaveLoanLeinholder", map[string]string{"DisplayName": "Save Loan Leinholder"}, "SaveLoanLeinholder", 5020, "Save Loan Leinholder"}
	pCErrorSaveVehicleCoverages                        = EnumPCErrorItem{pCErrorSaveVehicleCoveragesID, "SaveVehicleCoverages", map[string]string{"DisplayName": "Save Vehicle Coverages"}, "SaveVehicleCoverages", 5021, "Save Vehicle Coverages"}
	pCErrorMapVehiclesError                            = EnumPCErrorItem{pCErrorMapVehiclesErrorID, "MapVehiclesError", map[string]string{"DisplayName": "Map Vehicles Error"}, "MapVehiclesError", 5022, "Map Vehicles Error"}
	pCErrorSaveVehicleError                            = EnumPCErrorItem{pCErrorSaveVehicleErrorID, "SaveVehicleError", map[string]string{"DisplayName": "Save Vehicle Error"}, "SaveVehicleError", 5023, "Save Vehicle Error"}
	pCErrorVehicleBranded                              = EnumPCErrorItem{pCErrorVehicleBrandedID, "VehicleBranded", map[string]string{"DisplayName": "Vehicle Branded"}, "VehicleBranded", 5024, "Vehicle Branded"}
	pCErrorVinIso                                      = EnumPCErrorItem{pCErrorVinIsoID, "VinIso", map[string]string{"DisplayName": "Vin Iso"}, "VinIso", 5025, "Vin Iso"}
	pCErrorMissingPaperlessDiscount                    = EnumPCErrorItem{pCErrorMissingPaperlessDiscountID, "MissingPaperlessDiscount", map[string]string{"DisplayName": "Missing Paperless Discount"}, "MissingPaperlessDiscount", 6001, "Missing Paperless Discount"}
	pCErrorMissingResidency                            = EnumPCErrorItem{pCErrorMissingResidencyID, "MissingResidency", map[string]string{"DisplayName": "Missing Residency"}, "MissingResidency", 6002, "Missing Residency"}
	pCErrorMissingESignatureDiscount                   = EnumPCErrorItem{pCErrorMissingESignatureDiscountID, "MissingESignatureDiscount", map[string]string{"DisplayName": "Missing E Signature Discount"}, "MissingESignatureDiscount", 6003, "Missing E Signature Discount"}
	pCErrorMissingDoYouCurrentlyHaveInsureanceQuestion = EnumPCErrorItem{pCErrorMissingDoYouCurrentlyHaveInsureanceQuestionID, "MissingDoYouCurrentlyHaveInsureanceQuestion", map[string]string{"DisplayName": "Missing Do You Currently Have Insureance Question"}, "MissingDoYouCurrentlyHaveInsureanceQuestion", 6004, "Missing Do You Currently Have Insureance Question"}
	pCErrorMissingSourceOfBusiness                     = EnumPCErrorItem{pCErrorMissingSourceOfBusinessID, "MissingSourceOfBusiness", map[string]string{"DisplayName": "Missing Source Of Business"}, "MissingSourceOfBusiness", 6005, "Missing Source Of Business"}
	pCErrorMissingCurrentCarrierName                   = EnumPCErrorItem{pCErrorMissingCurrentCarrierNameID, "MissingCurrentCarrierName", map[string]string{"DisplayName": "Missing Current Carrier Name"}, "MissingCurrentCarrierName", 6006, "Missing Current Carrier Name"}
	pCErrorMissingCurrentCoverageEndDate               = EnumPCErrorItem{pCErrorMissingCurrentCoverageEndDateID, "MissingCurrentCoverageEndDate", map[string]string{"DisplayName": "Missing Current Coverage End Date"}, "MissingCurrentCoverageEndDate", 6007, "Missing Current Coverage End Date"}
	pCErrorMissingCurrentInjuryLimits                  = EnumPCErrorItem{pCErrorMissingCurrentInjuryLimitsID, "MissingCurrentInjuryLimits", map[string]string{"DisplayName": "Missing Current Injury Limits"}, "MissingCurrentInjuryLimits", 6008, "Missing Current Injury Limits"}
	pCErrorRelationshipToInsuredMissing                = EnumPCErrorItem{pCErrorRelationshipToInsuredMissingID, "RelationshipToInsuredMissing", map[string]string{"DisplayName": "Relationship To Insured Missing"}, "RelationshipToInsuredMissing", 6009, "Relationship To Insured Missing"}
	pCErrorMissingLapseInCoverageQuestion              = EnumPCErrorItem{pCErrorMissingLapseInCoverageQuestionID, "MissingLapseInCoverageQuestion", map[string]string{"DisplayName": "Missing Lapse In Coverage Question"}, "MissingLapseInCoverageQuestion", 6010, "Missing Lapse In Coverage Question"}
	pCErrorMissingYearWithCurrentProviderQuestion      = EnumPCErrorItem{pCErrorMissingYearWithCurrentProviderQuestionID, "MissingYearWithCurrentProviderQuestion", map[string]string{"DisplayName": "Missing Year With Current Provider Question"}, "MissingYearWithCurrentProviderQuestion", 6011, "Missing Year With Current Provider Question"}
	pCErrorCannotRemoveApplicant                       = EnumPCErrorItem{pCErrorCannotRemoveApplicantID, "CannotRemoveApplicant", map[string]string{"DisplayName": "Cannot Remove Applicant"}, "CannotRemoveApplicant", 6012, "Cannot Remove Applicant"}
	pCErrorSaveCoverageTerms                           = EnumPCErrorItem{pCErrorSaveCoverageTermsID, "SaveCoverageTerms", map[string]string{"DisplayName": "Save Coverage Terms"}, "SaveCoverageTerms", 7002, "Save Coverage Terms"}
	pCErrorMapCoveragesError                           = EnumPCErrorItem{pCErrorMapCoveragesErrorID, "MapCoveragesError", map[string]string{"DisplayName": "Map Coverages Error"}, "MapCoveragesError", 7003, "Map Coverages Error"}
	pCErrorInvalidCollisionDeductible                  = EnumPCErrorItem{pCErrorInvalidCollisionDeductibleID, "InvalidCollisionDeductible", map[string]string{"DisplayName": "Invalid Collision Deductible"}, "InvalidCollisionDeductible", 7004, "Invalid Collision Deductible"}
	pCErrorDiminishingDeductibleOnVehicles             = EnumPCErrorItem{pCErrorDiminishingDeductibleOnVehiclesID, "DiminishingDeductibleOnVehicles", map[string]string{"DisplayName": "Diminishing Deductible On Vehicles"}, "DiminishingDeductibleOnVehicles", 7005, "Diminishing Deductible On Vehicles"}
	pCErrorSavingManualIncidentsError                  = EnumPCErrorItem{pCErrorSavingManualIncidentsErrorID, "SavingManualIncidentsError", map[string]string{"DisplayName": "Saving Manual Incidents Error"}, "SavingManualIncidentsError", 8000, "Saving Manual Incidents Error"}
	pCErrorSavingClueIncidentsError                    = EnumPCErrorItem{pCErrorSavingClueIncidentsErrorID, "SavingClueIncidentsError", map[string]string{"DisplayName": "Saving Clue Incidents Error"}, "SavingClueIncidentsError", 8001, "Saving Clue Incidents Error"}
	pCErrorSavingMvrIncidentsError                     = EnumPCErrorItem{pCErrorSavingMvrIncidentsErrorID, "SavingMvrIncidentsError", map[string]string{"DisplayName": "Saving Mvr Incidents Error"}, "SavingMvrIncidentsError", 8002, "Saving Mvr Incidents Error"}
	pCErrorUnknownSaveIncidentsError                   = EnumPCErrorItem{pCErrorUnknownSaveIncidentsErrorID, "UnknownSaveIncidentsError", map[string]string{"DisplayName": "Unknown Save Incidents Error"}, "UnknownSaveIncidentsError", 8003, "Unknown Save Incidents Error"}
	pCErrorQuoteRequiresPhotoInspection                = EnumPCErrorItem{pCErrorQuoteRequiresPhotoInspectionID, "QuoteRequiresPhotoInspection", map[string]string{"DisplayName": "Quote Requires Photo Inspection"}, "QuoteRequiresPhotoInspection", 8004, "Quote Requires Photo Inspection"}
	pCErrorMissingMvrReport                            = EnumPCErrorItem{pCErrorMissingMvrReportID, "MissingMvrReport", map[string]string{"DisplayName": "Missing Mvr Report"}, "MissingMvrReport", 8005, "Missing Mvr Report"}
	pCErrorDuiLimitExceeded                            = EnumPCErrorItem{pCErrorDuiLimitExceededID, "DuiLimitExceeded", map[string]string{"DisplayName": "Dui Limit Exceeded"}, "DuiLimitExceeded", 8006, "Dui Limit Exceeded"}
	pCErrorAtFaultLimitExceeded                        = EnumPCErrorItem{pCErrorAtFaultLimitExceededID, "AtFaultLimitExceeded", map[string]string{"DisplayName": "At Fault Limit Exceeded"}, "AtFaultLimitExceeded", 8007, "At Fault Limit Exceeded"}
	pCErrorMinMajorConvictionLimitExceeded             = EnumPCErrorItem{pCErrorMinMajorConvictionLimitExceededID, "MinMajorConvictionLimitExceeded", map[string]string{"DisplayName": "Min Major Conviction Limit Exceeded"}, "MinMajorConvictionLimitExceeded", 8008, "Min Major Conviction Limit Exceeded"}
	pCErrorMapClueIncidentsError                       = EnumPCErrorItem{pCErrorMapClueIncidentsErrorID, "MapClueIncidentsError", map[string]string{"DisplayName": "Map Clue Incidents Error"}, "MapClueIncidentsError", 8009, "Map Clue Incidents Error"}
	pCErrorMapMvrIncidentsError                        = EnumPCErrorItem{pCErrorMapMvrIncidentsErrorID, "MapMvrIncidentsError", map[string]string{"DisplayName": "Map Mvr Incidents Error"}, "MapMvrIncidentsError", 8010, "Map Mvr Incidents Error"}
	pCErrorMapManualIncidentsError                     = EnumPCErrorItem{pCErrorMapManualIncidentsErrorID, "MapManualIncidentsError", map[string]string{"DisplayName": "Map Manual Incidents Error"}, "MapManualIncidentsError", 8011, "Map Manual Incidents Error"}
	pCErrorMapIncidentsUnknownError                    = EnumPCErrorItem{pCErrorMapIncidentsUnknownErrorID, "MapIncidentsUnknownError", map[string]string{"DisplayName": "Map Incidents Unknown Error"}, "MapIncidentsUnknownError", 8012, "Map Incidents Unknown Error"}
	pCErrorSavingIncidentsError                        = EnumPCErrorItem{pCErrorSavingIncidentsErrorID, "SavingIncidentsError", map[string]string{"DisplayName": "Saving Incidents Error"}, "SavingIncidentsError", 8013, "Saving Incidents Error"}
	pCErrorShouldPullCreditError                       = EnumPCErrorItem{pCErrorShouldPullCreditErrorID, "ShouldPullCreditError", map[string]string{"DisplayName": "Should Pull Credit Error"}, "ShouldPullCreditError", 8014, "Should Pull Credit Error"}
	pCErrorPullReportsError                            = EnumPCErrorItem{pCErrorPullReportsErrorID, "PullReportsError", map[string]string{"DisplayName": "Pull Reports Error"}, "PullReportsError", 8015, "Pull Reports Error"}
	pCErrorMajorConvictionsLimitExceeded               = EnumPCErrorItem{pCErrorMajorConvictionsLimitExceededID, "MajorConvictionsLimitExceeded", map[string]string{"DisplayName": "Major Convictions Limit Exceeded"}, "MajorConvictionsLimitExceeded", 8016, "Major Convictions Limit Exceeded"}
	pCErrorPipClaimsLimitExceeded                      = EnumPCErrorItem{pCErrorPipClaimsLimitExceededID, "PipClaimsLimitExceeded", map[string]string{"DisplayName": "Pip Claims Limit Exceeded"}, "PipClaimsLimitExceeded", 8017, "Pip Claims Limit Exceeded"}
	pCErrorChargeableAccidentsLimitExceeded            = EnumPCErrorItem{pCErrorChargeableAccidentsLimitExceededID, "ChargeableAccidentsLimitExceeded", map[string]string{"DisplayName": "Chargeable Accidents Limit Exceeded"}, "ChargeableAccidentsLimitExceeded", 8018, "Chargeable Accidents Limit Exceeded"}
	pCErrorJobNotInQuoteStatus                         = EnumPCErrorItem{pCErrorJobNotInQuoteStatusID, "JobNotInQuoteStatus", map[string]string{"DisplayName": "Job Not In Quote Status"}, "JobNotInQuoteStatus", 9000, "Job Not In Quote Status"}
	pCErrorAccountHasActivePolicy                      = EnumPCErrorItem{pCErrorAccountHasActivePolicyID, "AccountHasActivePolicy", map[string]string{"DisplayName": "Account Has Active Policy"}, "AccountHasActivePolicy", 9001, "Account Has Active Policy"}
	pCErrorUnderwritingBlockingIssue                   = EnumPCErrorItem{pCErrorUnderwritingBlockingIssueID, "UnderwritingBlockingIssue", map[string]string{"DisplayName": "Underwriting Blocking Issue"}, "UnderwritingBlockingIssue", 9002, "Underwriting Blocking Issue"}
	pCErrorClueReportRequired                          = EnumPCErrorItem{pCErrorClueReportRequiredID, "ClueReportRequired", map[string]string{"DisplayName": "Clue Report Required"}, "ClueReportRequired", 9004, "Clue Report Required"}
	pCErrorBackDateError                               = EnumPCErrorItem{pCErrorBackDateErrorID, "BackDateError", map[string]string{"DisplayName": "Back Date Error"}, "BackDateError", 9005, "Back Date Error"}
	pCErrorExcludedDriverSendDocument                  = EnumPCErrorItem{pCErrorExcludedDriverSendDocumentID, "ExcludedDriverSendDocument", map[string]string{"DisplayName": "Excluded Driver Send Document"}, "ExcludedDriverSendDocument", 9006, "Excluded Driver Send Document"}
	pCErrorExcludedDriverResendDocument                = EnumPCErrorItem{pCErrorExcludedDriverResendDocumentID, "ExcludedDriverResendDocument", map[string]string{"DisplayName": "Excluded Driver Resend Document"}, "ExcludedDriverResendDocument", 9007, "Excluded Driver Resend Document"}
	pCErrorExcludedDriverFormNotSigned                 = EnumPCErrorItem{pCErrorExcludedDriverFormNotSignedID, "ExcludedDriverFormNotSigned", map[string]string{"DisplayName": "Excluded Driver Form Not Signed"}, "ExcludedDriverFormNotSigned", 9008, "Excluded Driver Form Not Signed"}
	pCErrorAccidentDayOfBind                           = EnumPCErrorItem{pCErrorAccidentDayOfBindID, "AccidentDayOfBind", map[string]string{"DisplayName": "Accident Day Of Bind"}, "AccidentDayOfBind", 9009, "Accident Day Of Bind"}
	pCErrorMissingPaymentplan                          = EnumPCErrorItem{pCErrorMissingPaymentplanID, "MissingPaymentplan", map[string]string{"DisplayName": "Missing Paymentplan"}, "MissingPaymentplan", 9010, "Missing Paymentplan"}
	pCErrorPolicyHasPhotoReview                        = EnumPCErrorItem{pCErrorPolicyHasPhotoReviewID, "PolicyHasPhotoReview", map[string]string{"DisplayName": "Policy Has Photo Review"}, "PolicyHasPhotoReview", 9011, "Policy Has Photo Review"}
	pCErrorDocusignIntegrationError                    = EnumPCErrorItem{pCErrorDocusignIntegrationErrorID, "DocusignIntegrationError", map[string]string{"DisplayName": "Docusign Integration Error"}, "DocusignIntegrationError", 9012, "Docusign Integration Error"}
	pCErrorCannotBindAccidentDayOfBind                 = EnumPCErrorItem{pCErrorCannotBindAccidentDayOfBindID, "CannotBindAccidentDayOfBind", map[string]string{"DisplayName": "Cannot Bind Accident Day Of Bind"}, "CannotBindAccidentDayOfBind", 9014, "Cannot Bind Accident Day Of Bind"}
	pCErrorDriverNotAssignedForLegal                   = EnumPCErrorItem{pCErrorDriverNotAssignedForLegalID, "DriverNotAssignedForLegal", map[string]string{"DisplayName": "Driver Not Assigned For Legal"}, "DriverNotAssignedForLegal", 9015, "Driver Not Assigned For Legal"}
	pCErrorIneligibleForFamilyLegal                    = EnumPCErrorItem{pCErrorIneligibleForFamilyLegalID, "IneligibleForFamilyLegal", map[string]string{"DisplayName": "Ineligible For Family Legal"}, "IneligibleForFamilyLegal", 9016, "Ineligible For Family Legal"}
	pCErrorNullDmsReason                               = EnumPCErrorItem{pCErrorNullDmsReasonID, "NullDmsReason", map[string]string{"DisplayName": "Null Dms Reason"}, "NullDmsReason", 9017, "Null Dms Reason"}
	pCErrorTerritoryRestriction                        = EnumPCErrorItem{pCErrorTerritoryRestrictionID, "TerritoryRestriction", map[string]string{"DisplayName": "Territory Restriction"}, "TerritoryRestriction", 9018, "Territory Restriction"}
	pCErrorAccountHasMaterialMisrep                    = EnumPCErrorItem{pCErrorAccountHasMaterialMisrepID, "AccountHasMaterialMisrep", map[string]string{"DisplayName": "Account Has Material Misrep"}, "AccountHasMaterialMisrep", 9019, "Account Has Material Misrep"}
	pCErrorIllegalBundleTransferException              = EnumPCErrorItem{pCErrorIllegalBundleTransferExceptionID, "IllegalBundleTransferException", map[string]string{"DisplayName": "Illegal Bundle Transfer Exception"}, "IllegalBundleTransferException", 9998, "Illegal Bundle Transfer Exception"}
	pCErrorConcurrentDataChangeException               = EnumPCErrorItem{pCErrorConcurrentDataChangeExceptionID, "ConcurrentDataChangeException", map[string]string{"DisplayName": "Concurrent Data Change Exception"}, "ConcurrentDataChangeException", 9999, "Concurrent Data Change Exception"}
	pCErrorGenerateProductsError                       = EnumPCErrorItem{pCErrorGenerateProductsErrorID, "GenerateProductsError", map[string]string{"DisplayName": "Generate Products Error"}, "GenerateProductsError", 10001, "Generate Products Error"}
	pCErrorQuotePolicyError                            = EnumPCErrorItem{pCErrorQuotePolicyErrorID, "QuotePolicyError", map[string]string{"DisplayName": "Quote Policy Error"}, "QuotePolicyError", 10002, "Quote Policy Error"}
	pCErrorBindPolicyError                             = EnumPCErrorItem{pCErrorBindPolicyErrorID, "BindPolicyError", map[string]string{"DisplayName": "Bind Policy Error"}, "BindPolicyError", 10003, "Bind Policy Error"}
	pCErrorFetchDocumentsError                         = EnumPCErrorItem{pCErrorFetchDocumentsErrorID, "FetchDocumentsError", map[string]string{"DisplayName": "Fetch Documents Error"}, "FetchDocumentsError", 10004, "Fetch Documents Error"}
	pCErrorInvalidPaymentPlan                          = EnumPCErrorItem{pCErrorInvalidPaymentPlanID, "InvalidPaymentPlan", map[string]string{"DisplayName": "Invalid Payment Plan"}, "InvalidPaymentPlan", 10005, "Invalid Payment Plan"}
	pCErrorInvalidPolicyData                           = EnumPCErrorItem{pCErrorInvalidPolicyDataID, "InvalidPolicyData", map[string]string{"DisplayName": "Invalid Policy Data"}, "InvalidPolicyData", 10010, "Invalid Policy Data"}
	pCErrorNoDocuments                                 = EnumPCErrorItem{pCErrorNoDocumentsID, "NoDocuments", map[string]string{"DisplayName": "No Documents"}, "NoDocuments", 10020, "No Documents"}
	pCErrorNoJobFound                                  = EnumPCErrorItem{pCErrorNoJobFoundID, "NoJobFound", map[string]string{"DisplayName": "No Job Found"}, "NoJobFound", 10021, "No Job Found"}
	pCErrorDeleteItemError                             = EnumPCErrorItem{pCErrorDeleteItemErrorID, "DeleteItemError", map[string]string{"DisplayName": "Delete Item Error"}, "DeleteItemError", 10022, "Delete Item Error"}
	pCErrorWorkorderNotQuotable                        = EnumPCErrorItem{pCErrorWorkorderNotQuotableID, "WorkorderNotQuotable", map[string]string{"DisplayName": "Workorder Not Quotable"}, "WorkorderNotQuotable", 10023, "Workorder Not Quotable"}
	pCErrorUnknownPaymentError                         = EnumPCErrorItem{pCErrorUnknownPaymentErrorID, "UnknownPaymentError", map[string]string{"DisplayName": "Unknown Payment Error"}, "UnknownPaymentError", 20000, "Unknown Payment Error"}
	pCErrorNoCreditCardNumber                          = EnumPCErrorItem{pCErrorNoCreditCardNumberID, "NoCreditCardNumber", map[string]string{"DisplayName": "No Credit Card Number"}, "NoCreditCardNumber", 20001, "No Credit Card Number"}
	pCErrorFeePaidWithDownPayment                      = EnumPCErrorItem{pCErrorFeePaidWithDownPaymentID, "FeePaidWithDownPayment", map[string]string{"DisplayName": "Fee Paid With Down Payment"}, "FeePaidWithDownPayment", 20002, "Fee Paid With Down Payment"}
	pCErrorCardDeclined                                = EnumPCErrorItem{pCErrorCardDeclinedID, "CardDeclined", map[string]string{"DisplayName": "Card Declined"}, "CardDeclined", 20012, "Card Declined"}
	pCErrorInvalidCardNumber                           = EnumPCErrorItem{pCErrorInvalidCardNumberID, "InvalidCardNumber", map[string]string{"DisplayName": "Invalid Card Number"}, "InvalidCardNumber", 20023, "Invalid Card Number"}
	pCErrorInvalidExpiration                           = EnumPCErrorItem{pCErrorInvalidExpirationID, "InvalidExpiration", map[string]string{"DisplayName": "Invalid Expiration"}, "InvalidExpiration", 20024, "Invalid Expiration"}
	pCErrorInvalidCardType                             = EnumPCErrorItem{pCErrorInvalidCardTypeID, "InvalidCardType", map[string]string{"DisplayName": "Invalid Card Type"}, "InvalidCardType", 20025, "Invalid Card Type"}
	pCErrorInsufficientFunds                           = EnumPCErrorItem{pCErrorInsufficientFundsID, "InsufficientFunds", map[string]string{"DisplayName": "Insufficient Funds"}, "InsufficientFunds", 20050, "Insufficient Funds"}
	pCErrorInvalidAbaNumber                            = EnumPCErrorItem{pCErrorInvalidAbaNumberID, "InvalidAbaNumber", map[string]string{"DisplayName": "Invalid Aba Number"}, "InvalidAbaNumber", 20101, "Invalid Aba Number"}
	pCErrorInvalidBankAccountNumber                    = EnumPCErrorItem{pCErrorInvalidBankAccountNumberID, "InvalidBankAccountNumber", map[string]string{"DisplayName": "Invalid Bank Account Number"}, "InvalidBankAccountNumber", 20102, "Invalid Bank Account Number"}
	pCErrorInvalidBankName                             = EnumPCErrorItem{pCErrorInvalidBankNameID, "InvalidBankName", map[string]string{"DisplayName": "Invalid Bank Name"}, "InvalidBankName", 20103, "Invalid Bank Name"}
	pCErrorBankAccountType                             = EnumPCErrorItem{pCErrorBankAccountTypeID, "BankAccountType", map[string]string{"DisplayName": "Bank Account Type"}, "BankAccountType", 20104, "Bank Account Type"}
	pCErrorLastNameMissing                             = EnumPCErrorItem{pCErrorLastNameMissingID, "LastNameMissing", map[string]string{"DisplayName": "Last Name Missing"}, "LastNameMissing", 200200, "Last Name Missing"}
	pCErrorMissingAccountContact                       = EnumPCErrorItem{pCErrorMissingAccountContactID, "MissingAccountContact", map[string]string{"DisplayName": "Missing Account Contact"}, "MissingAccountContact", 201313, "Missing Account Contact"}
	pCErrorValidateFcraNotice                          = EnumPCErrorItem{pCErrorValidateFcraNoticeID, "ValidateFcraNotice", map[string]string{"DisplayName": "Validate Fcra Notice"}, "ValidateFcraNotice", 201818, "Validate Fcra Notice"}
	pCErrorPolicyHolderWithPermitLicense               = EnumPCErrorItem{pCErrorPolicyHolderWithPermitLicenseID, "PolicyHolderWithPermitLicense", map[string]string{"DisplayName": "Policy Holder With Permit License Status"}, "PolicyHolderWithPermitLicense", 2048, "Policy Holder With Permit License Status"}
)

// EnumPCError is a collection of PCError items
type EnumPCError struct {
	Description string
	Items       []*EnumPCErrorItem
	Name        string

	MissingModelYear                            *EnumPCErrorItem
	InvalidDateFormat                           *EnumPCErrorItem
	InvalidState                                *EnumPCErrorItem
	AccountNumberMissing                        *EnumPCErrorItem
	AccountNotFound                             *EnumPCErrorItem
	MapAccountError                             *EnumPCErrorItem
	SaveAccountError                            *EnumPCErrorItem
	HasExistingAccountError                     *EnumPCErrorItem
	SaveAccountHolderError                      *EnumPCErrorItem
	GenerateAccountError                        *EnumPCErrorItem
	CreateAccountError                          *EnumPCErrorItem
	GetRecentAccountError                       *EnumPCErrorItem
	AccountHasBadDebt                           *EnumPCErrorItem
	AccountLookupError                          *EnumPCErrorItem
	FirstNameMissing                            *EnumPCErrorItem
	InvalidEffectiveDate                        *EnumPCErrorItem
	DobMissing                                  *EnumPCErrorItem
	HomePhoneMissing                            *EnumPCErrorItem
	AddressLine1Missing                         *EnumPCErrorItem
	CityMissing                                 *EnumPCErrorItem
	StateMissing                                *EnumPCErrorItem
	PostalMissing                               *EnumPCErrorItem
	GenderMissing                               *EnumPCErrorItem
	LicenseStateOrNumberMissing                 *EnumPCErrorItem
	HighestEducationLevelMissing                *EnumPCErrorItem
	OccupationTitleMissing                      *EnumPCErrorItem
	CannotUpdateSameDayBindQuestion             *EnumPCErrorItem
	InvalidPhoneNumber                          *EnumPCErrorItem
	InvalidEmail                                *EnumPCErrorItem
	InvalidOnlineDiscount                       *EnumPCErrorItem
	InvalidWebToPhone                           *EnumPCErrorItem
	ValidatePaperlessDiscount                   *EnumPCErrorItem
	InvalidPolicyEffectiveDate                  *EnumPCErrorItem
	InvalidTransactiveEffecitve                 *EnumPCErrorItem
	MarriedWithNoSpouse                         *EnumPCErrorItem
	MarriedWithMultpleSpouses                   *EnumPCErrorItem
	EarlyBirdDiscount                           *EnumPCErrorItem
	CountyMissing                               *EnumPCErrorItem
	JobNumberMissing                            *EnumPCErrorItem
	PolicynumberNumberMissing                   *EnumPCErrorItem
	MapQualifiedDiscounts                       *EnumPCErrorItem
	CannotCreateSavePolicyperiod                *EnumPCErrorItem
	MapPolicyError                              *EnumPCErrorItem
	MapPolicyInitDataError                      *EnumPCErrorItem
	MapPolicyLineDetailsError                   *EnumPCErrorItem
	MapQuotedValuesError                        *EnumPCErrorItem
	SavePolicyError                             *EnumPCErrorItem
	RetrieveQuoteError                          *EnumPCErrorItem
	InvalidJobStatus                            *EnumPCErrorItem
	CannotEditJob                               *EnumPCErrorItem
	DeleteDriverError                           *EnumPCErrorItem
	DeleteManualIncidentsError                  *EnumPCErrorItem
	DeleteVehicleError                          *EnumPCErrorItem
	RatingStatusMissing                         *EnumPCErrorItem
	DriverUnacceptableRisk                      *EnumPCErrorItem
	CreditPullError                             *EnumPCErrorItem
	AdditionalDriversNeeded                     *EnumPCErrorItem
	Fr19DataMissing                             *EnumPCErrorItem
	AllDriversRequirePrimaryVehicle             *EnumPCErrorItem
	AdditionalDriverNeeded                      *EnumPCErrorItem
	InvalidLicenseStatus                        *EnumPCErrorItem
	NoForeignLicense                            *EnumPCErrorItem
	OutOfStateLicense                           *EnumPCErrorItem
	PrimaryDriverRequired                       *EnumPCErrorItem
	MissingDriver                               *EnumPCErrorItem
	MissingCreditPull                           *EnumPCErrorItem
	InvalidDriverID                             *EnumPCErrorItem
	SaveDriverOccupation                        *EnumPCErrorItem
	MapDriverOccupation                         *EnumPCErrorItem
	MapDriverBasicInfo                          *EnumPCErrorItem
	MapDriverIncidents                          *EnumPCErrorItem
	MapDriverError                              *EnumPCErrorItem
	MapDriverLicenseDetails                     *EnumPCErrorItem
	MapDriversError                             *EnumPCErrorItem
	InvalidIncidentID                           *EnumPCErrorItem
	SaveDriverBasicInfo                         *EnumPCErrorItem
	SaveDriverIncidents                         *EnumPCErrorItem
	SaveDriverError                             *EnumPCErrorItem
	SaveDriverLicenseDetails                    *EnumPCErrorItem
	MinConvictionExceeded                       *EnumPCErrorItem
	MissingMake                                 *EnumPCErrorItem
	MissingModel                                *EnumPCErrorItem
	MissingPrimaryUse                           *EnumPCErrorItem
	InvalidVehicleDriverID                      *EnumPCErrorItem
	CannotAddVehicleToJob                       *EnumPCErrorItem
	InvalidVehicleID                            *EnumPCErrorItem
	CannotRemoveVehicle                         *EnumPCErrorItem
	CannotFindVehicleForDriver                  *EnumPCErrorItem
	LienLeaseRequired                           *EnumPCErrorItem
	AllVehicleRequirePrimaryDriver              *EnumPCErrorItem
	ExoticVehicle                               *EnumPCErrorItem
	InvalidVin                                  *EnumPCErrorItem
	VehicleTypeRequired                         *EnumPCErrorItem
	VehicleNewCostRequired                      *EnumPCErrorItem
	VehiclePositiveCostRequired                 *EnumPCErrorItem
	AtLeastOneVehicleRequired                   *EnumPCErrorItem
	VinsMustBeUniqueForAllVehicles              *EnumPCErrorItem
	VehiclePrimaryUseRequired                   *EnumPCErrorItem
	SaveLoanLeinholder                          *EnumPCErrorItem
	SaveVehicleCoverages                        *EnumPCErrorItem
	MapVehiclesError                            *EnumPCErrorItem
	SaveVehicleError                            *EnumPCErrorItem
	VehicleBranded                              *EnumPCErrorItem
	VinIso                                      *EnumPCErrorItem
	MissingPaperlessDiscount                    *EnumPCErrorItem
	MissingResidency                            *EnumPCErrorItem
	MissingESignatureDiscount                   *EnumPCErrorItem
	MissingDoYouCurrentlyHaveInsureanceQuestion *EnumPCErrorItem
	MissingSourceOfBusiness                     *EnumPCErrorItem
	MissingCurrentCarrierName                   *EnumPCErrorItem
	MissingCurrentCoverageEndDate               *EnumPCErrorItem
	MissingCurrentInjuryLimits                  *EnumPCErrorItem
	RelationshipToInsuredMissing                *EnumPCErrorItem
	MissingLapseInCoverageQuestion              *EnumPCErrorItem
	MissingYearWithCurrentProviderQuestion      *EnumPCErrorItem
	CannotRemoveApplicant                       *EnumPCErrorItem
	SaveCoverageTerms                           *EnumPCErrorItem
	MapCoveragesError                           *EnumPCErrorItem
	InvalidCollisionDeductible                  *EnumPCErrorItem
	DiminishingDeductibleOnVehicles             *EnumPCErrorItem
	SavingManualIncidentsError                  *EnumPCErrorItem
	SavingClueIncidentsError                    *EnumPCErrorItem
	SavingMvrIncidentsError                     *EnumPCErrorItem
	UnknownSaveIncidentsError                   *EnumPCErrorItem
	QuoteRequiresPhotoInspection                *EnumPCErrorItem
	MissingMvrReport                            *EnumPCErrorItem
	DuiLimitExceeded                            *EnumPCErrorItem
	AtFaultLimitExceeded                        *EnumPCErrorItem
	MinMajorConvictionLimitExceeded             *EnumPCErrorItem
	MapClueIncidentsError                       *EnumPCErrorItem
	MapMvrIncidentsError                        *EnumPCErrorItem
	MapManualIncidentsError                     *EnumPCErrorItem
	MapIncidentsUnknownError                    *EnumPCErrorItem
	SavingIncidentsError                        *EnumPCErrorItem
	ShouldPullCreditError                       *EnumPCErrorItem
	PullReportsError                            *EnumPCErrorItem
	MajorConvictionsLimitExceeded               *EnumPCErrorItem
	PipClaimsLimitExceeded                      *EnumPCErrorItem
	ChargeableAccidentsLimitExceeded            *EnumPCErrorItem
	JobNotInQuoteStatus                         *EnumPCErrorItem
	AccountHasActivePolicy                      *EnumPCErrorItem
	UnderwritingBlockingIssue                   *EnumPCErrorItem
	ClueReportRequired                          *EnumPCErrorItem
	BackDateError                               *EnumPCErrorItem
	ExcludedDriverSendDocument                  *EnumPCErrorItem
	ExcludedDriverResendDocument                *EnumPCErrorItem
	ExcludedDriverFormNotSigned                 *EnumPCErrorItem
	AccidentDayOfBind                           *EnumPCErrorItem
	MissingPaymentplan                          *EnumPCErrorItem
	PolicyHasPhotoReview                        *EnumPCErrorItem
	DocusignIntegrationError                    *EnumPCErrorItem
	CannotBindAccidentDayOfBind                 *EnumPCErrorItem
	DriverNotAssignedForLegal                   *EnumPCErrorItem
	IneligibleForFamilyLegal                    *EnumPCErrorItem
	NullDmsReason                               *EnumPCErrorItem
	TerritoryRestriction                        *EnumPCErrorItem
	AccountHasMaterialMisrep                    *EnumPCErrorItem
	IllegalBundleTransferException              *EnumPCErrorItem
	ConcurrentDataChangeException               *EnumPCErrorItem
	GenerateProductsError                       *EnumPCErrorItem
	QuotePolicyError                            *EnumPCErrorItem
	BindPolicyError                             *EnumPCErrorItem
	FetchDocumentsError                         *EnumPCErrorItem
	InvalidPaymentPlan                          *EnumPCErrorItem
	InvalidPolicyData                           *EnumPCErrorItem
	NoDocuments                                 *EnumPCErrorItem
	NoJobFound                                  *EnumPCErrorItem
	DeleteItemError                             *EnumPCErrorItem
	WorkorderNotQuotable                        *EnumPCErrorItem
	UnknownPaymentError                         *EnumPCErrorItem
	NoCreditCardNumber                          *EnumPCErrorItem
	FeePaidWithDownPayment                      *EnumPCErrorItem
	CardDeclined                                *EnumPCErrorItem
	InvalidCardNumber                           *EnumPCErrorItem
	InvalidExpiration                           *EnumPCErrorItem
	InvalidCardType                             *EnumPCErrorItem
	InsufficientFunds                           *EnumPCErrorItem
	InvalidAbaNumber                            *EnumPCErrorItem
	InvalidBankAccountNumber                    *EnumPCErrorItem
	InvalidBankName                             *EnumPCErrorItem
	BankAccountType                             *EnumPCErrorItem
	LastNameMissing                             *EnumPCErrorItem
	MissingAccountContact                       *EnumPCErrorItem
	ValidateFcraNotice                          *EnumPCErrorItem
	PolicyHolderWithPermitLicense               *EnumPCErrorItem

	itemDict map[string]*EnumPCErrorItem
}

// PCError is a public singleton instance of EnumPCError
// representing Policy Center error codes
var PCError = &EnumPCError{
	Description: "Policy Center error codes",
	Items: []*EnumPCErrorItem{
		&pCErrorMissingModelYear,
		&pCErrorInvalidDateFormat,
		&pCErrorInvalidState,
		&pCErrorAccountNumberMissing,
		&pCErrorAccountNotFound,
		&pCErrorMapAccountError,
		&pCErrorSaveAccountError,
		&pCErrorHasExistingAccountError,
		&pCErrorSaveAccountHolderError,
		&pCErrorGenerateAccountError,
		&pCErrorCreateAccountError,
		&pCErrorGetRecentAccountError,
		&pCErrorAccountHasBadDebt,
		&pCErrorAccountLookupError,
		&pCErrorFirstNameMissing,
		&pCErrorInvalidEffectiveDate,
		&pCErrorDobMissing,
		&pCErrorHomePhoneMissing,
		&pCErrorAddressLine1Missing,
		&pCErrorCityMissing,
		&pCErrorStateMissing,
		&pCErrorPostalMissing,
		&pCErrorGenderMissing,
		&pCErrorLicenseStateOrNumberMissing,
		&pCErrorHighestEducationLevelMissing,
		&pCErrorOccupationTitleMissing,
		&pCErrorCannotUpdateSameDayBindQuestion,
		&pCErrorInvalidPhoneNumber,
		&pCErrorInvalidEmail,
		&pCErrorInvalidOnlineDiscount,
		&pCErrorInvalidWebToPhone,
		&pCErrorValidatePaperlessDiscount,
		&pCErrorInvalidPolicyEffectiveDate,
		&pCErrorInvalidTransactiveEffecitve,
		&pCErrorMarriedWithNoSpouse,
		&pCErrorMarriedWithMultpleSpouses,
		&pCErrorEarlyBirdDiscount,
		&pCErrorCountyMissing,
		&pCErrorJobNumberMissing,
		&pCErrorPolicynumberNumberMissing,
		&pCErrorMapQualifiedDiscounts,
		&pCErrorCannotCreateSavePolicyperiod,
		&pCErrorMapPolicyError,
		&pCErrorMapPolicyInitDataError,
		&pCErrorMapPolicyLineDetailsError,
		&pCErrorMapQuotedValuesError,
		&pCErrorSavePolicyError,
		&pCErrorRetrieveQuoteError,
		&pCErrorInvalidJobStatus,
		&pCErrorCannotEditJob,
		&pCErrorDeleteDriverError,
		&pCErrorDeleteManualIncidentsError,
		&pCErrorDeleteVehicleError,
		&pCErrorRatingStatusMissing,
		&pCErrorDriverUnacceptableRisk,
		&pCErrorCreditPullError,
		&pCErrorAdditionalDriversNeeded,
		&pCErrorFr19DataMissing,
		&pCErrorAllDriversRequirePrimaryVehicle,
		&pCErrorAdditionalDriverNeeded,
		&pCErrorInvalidLicenseStatus,
		&pCErrorNoForeignLicense,
		&pCErrorOutOfStateLicense,
		&pCErrorPrimaryDriverRequired,
		&pCErrorMissingDriver,
		&pCErrorMissingCreditPull,
		&pCErrorInvalidDriverID,
		&pCErrorSaveDriverOccupation,
		&pCErrorMapDriverOccupation,
		&pCErrorMapDriverBasicInfo,
		&pCErrorMapDriverIncidents,
		&pCErrorMapDriverError,
		&pCErrorMapDriverLicenseDetails,
		&pCErrorMapDriversError,
		&pCErrorInvalidIncidentID,
		&pCErrorSaveDriverBasicInfo,
		&pCErrorSaveDriverIncidents,
		&pCErrorSaveDriverError,
		&pCErrorSaveDriverLicenseDetails,
		&pCErrorMinConvictionExceeded,
		&pCErrorMissingMake,
		&pCErrorMissingModel,
		&pCErrorMissingPrimaryUse,
		&pCErrorInvalidVehicleDriverID,
		&pCErrorCannotAddVehicleToJob,
		&pCErrorInvalidVehicleID,
		&pCErrorCannotRemoveVehicle,
		&pCErrorCannotFindVehicleForDriver,
		&pCErrorLienLeaseRequired,
		&pCErrorAllVehicleRequirePrimaryDriver,
		&pCErrorExoticVehicle,
		&pCErrorInvalidVin,
		&pCErrorVehicleTypeRequired,
		&pCErrorVehicleNewCostRequired,
		&pCErrorVehiclePositiveCostRequired,
		&pCErrorAtLeastOneVehicleRequired,
		&pCErrorVinsMustBeUniqueForAllVehicles,
		&pCErrorVehiclePrimaryUseRequired,
		&pCErrorSaveLoanLeinholder,
		&pCErrorSaveVehicleCoverages,
		&pCErrorMapVehiclesError,
		&pCErrorSaveVehicleError,
		&pCErrorVehicleBranded,
		&pCErrorVinIso,
		&pCErrorMissingPaperlessDiscount,
		&pCErrorMissingResidency,
		&pCErrorMissingESignatureDiscount,
		&pCErrorMissingDoYouCurrentlyHaveInsureanceQuestion,
		&pCErrorMissingSourceOfBusiness,
		&pCErrorMissingCurrentCarrierName,
		&pCErrorMissingCurrentCoverageEndDate,
		&pCErrorMissingCurrentInjuryLimits,
		&pCErrorRelationshipToInsuredMissing,
		&pCErrorMissingLapseInCoverageQuestion,
		&pCErrorMissingYearWithCurrentProviderQuestion,
		&pCErrorCannotRemoveApplicant,
		&pCErrorSaveCoverageTerms,
		&pCErrorMapCoveragesError,
		&pCErrorInvalidCollisionDeductible,
		&pCErrorDiminishingDeductibleOnVehicles,
		&pCErrorSavingManualIncidentsError,
		&pCErrorSavingClueIncidentsError,
		&pCErrorSavingMvrIncidentsError,
		&pCErrorUnknownSaveIncidentsError,
		&pCErrorQuoteRequiresPhotoInspection,
		&pCErrorMissingMvrReport,
		&pCErrorDuiLimitExceeded,
		&pCErrorAtFaultLimitExceeded,
		&pCErrorMinMajorConvictionLimitExceeded,
		&pCErrorMapClueIncidentsError,
		&pCErrorMapMvrIncidentsError,
		&pCErrorMapManualIncidentsError,
		&pCErrorMapIncidentsUnknownError,
		&pCErrorSavingIncidentsError,
		&pCErrorShouldPullCreditError,
		&pCErrorPullReportsError,
		&pCErrorMajorConvictionsLimitExceeded,
		&pCErrorPipClaimsLimitExceeded,
		&pCErrorChargeableAccidentsLimitExceeded,
		&pCErrorJobNotInQuoteStatus,
		&pCErrorAccountHasActivePolicy,
		&pCErrorUnderwritingBlockingIssue,
		&pCErrorClueReportRequired,
		&pCErrorBackDateError,
		&pCErrorExcludedDriverSendDocument,
		&pCErrorExcludedDriverResendDocument,
		&pCErrorExcludedDriverFormNotSigned,
		&pCErrorAccidentDayOfBind,
		&pCErrorMissingPaymentplan,
		&pCErrorPolicyHasPhotoReview,
		&pCErrorDocusignIntegrationError,
		&pCErrorCannotBindAccidentDayOfBind,
		&pCErrorDriverNotAssignedForLegal,
		&pCErrorIneligibleForFamilyLegal,
		&pCErrorNullDmsReason,
		&pCErrorTerritoryRestriction,
		&pCErrorAccountHasMaterialMisrep,
		&pCErrorIllegalBundleTransferException,
		&pCErrorConcurrentDataChangeException,
		&pCErrorGenerateProductsError,
		&pCErrorQuotePolicyError,
		&pCErrorBindPolicyError,
		&pCErrorFetchDocumentsError,
		&pCErrorInvalidPaymentPlan,
		&pCErrorInvalidPolicyData,
		&pCErrorNoDocuments,
		&pCErrorNoJobFound,
		&pCErrorDeleteItemError,
		&pCErrorWorkorderNotQuotable,
		&pCErrorUnknownPaymentError,
		&pCErrorNoCreditCardNumber,
		&pCErrorFeePaidWithDownPayment,
		&pCErrorCardDeclined,
		&pCErrorInvalidCardNumber,
		&pCErrorInvalidExpiration,
		&pCErrorInvalidCardType,
		&pCErrorInsufficientFunds,
		&pCErrorInvalidAbaNumber,
		&pCErrorInvalidBankAccountNumber,
		&pCErrorInvalidBankName,
		&pCErrorBankAccountType,
		&pCErrorLastNameMissing,
		&pCErrorMissingAccountContact,
		&pCErrorValidateFcraNotice,
		&pCErrorPolicyHolderWithPermitLicense,
	},
	Name:                                        "EnumPCError",
	MissingModelYear:                            &pCErrorMissingModelYear,
	InvalidDateFormat:                           &pCErrorInvalidDateFormat,
	InvalidState:                                &pCErrorInvalidState,
	AccountNumberMissing:                        &pCErrorAccountNumberMissing,
	AccountNotFound:                             &pCErrorAccountNotFound,
	MapAccountError:                             &pCErrorMapAccountError,
	SaveAccountError:                            &pCErrorSaveAccountError,
	HasExistingAccountError:                     &pCErrorHasExistingAccountError,
	SaveAccountHolderError:                      &pCErrorSaveAccountHolderError,
	GenerateAccountError:                        &pCErrorGenerateAccountError,
	CreateAccountError:                          &pCErrorCreateAccountError,
	GetRecentAccountError:                       &pCErrorGetRecentAccountError,
	AccountHasBadDebt:                           &pCErrorAccountHasBadDebt,
	AccountLookupError:                          &pCErrorAccountLookupError,
	FirstNameMissing:                            &pCErrorFirstNameMissing,
	InvalidEffectiveDate:                        &pCErrorInvalidEffectiveDate,
	DobMissing:                                  &pCErrorDobMissing,
	HomePhoneMissing:                            &pCErrorHomePhoneMissing,
	AddressLine1Missing:                         &pCErrorAddressLine1Missing,
	CityMissing:                                 &pCErrorCityMissing,
	StateMissing:                                &pCErrorStateMissing,
	PostalMissing:                               &pCErrorPostalMissing,
	GenderMissing:                               &pCErrorGenderMissing,
	LicenseStateOrNumberMissing:                 &pCErrorLicenseStateOrNumberMissing,
	HighestEducationLevelMissing:                &pCErrorHighestEducationLevelMissing,
	OccupationTitleMissing:                      &pCErrorOccupationTitleMissing,
	CannotUpdateSameDayBindQuestion:             &pCErrorCannotUpdateSameDayBindQuestion,
	InvalidPhoneNumber:                          &pCErrorInvalidPhoneNumber,
	InvalidEmail:                                &pCErrorInvalidEmail,
	InvalidOnlineDiscount:                       &pCErrorInvalidOnlineDiscount,
	InvalidWebToPhone:                           &pCErrorInvalidWebToPhone,
	ValidatePaperlessDiscount:                   &pCErrorValidatePaperlessDiscount,
	InvalidPolicyEffectiveDate:                  &pCErrorInvalidPolicyEffectiveDate,
	InvalidTransactiveEffecitve:                 &pCErrorInvalidTransactiveEffecitve,
	MarriedWithNoSpouse:                         &pCErrorMarriedWithNoSpouse,
	MarriedWithMultpleSpouses:                   &pCErrorMarriedWithMultpleSpouses,
	EarlyBirdDiscount:                           &pCErrorEarlyBirdDiscount,
	CountyMissing:                               &pCErrorCountyMissing,
	JobNumberMissing:                            &pCErrorJobNumberMissing,
	PolicynumberNumberMissing:                   &pCErrorPolicynumberNumberMissing,
	MapQualifiedDiscounts:                       &pCErrorMapQualifiedDiscounts,
	CannotCreateSavePolicyperiod:                &pCErrorCannotCreateSavePolicyperiod,
	MapPolicyError:                              &pCErrorMapPolicyError,
	MapPolicyInitDataError:                      &pCErrorMapPolicyInitDataError,
	MapPolicyLineDetailsError:                   &pCErrorMapPolicyLineDetailsError,
	MapQuotedValuesError:                        &pCErrorMapQuotedValuesError,
	SavePolicyError:                             &pCErrorSavePolicyError,
	RetrieveQuoteError:                          &pCErrorRetrieveQuoteError,
	InvalidJobStatus:                            &pCErrorInvalidJobStatus,
	CannotEditJob:                               &pCErrorCannotEditJob,
	DeleteDriverError:                           &pCErrorDeleteDriverError,
	DeleteManualIncidentsError:                  &pCErrorDeleteManualIncidentsError,
	DeleteVehicleError:                          &pCErrorDeleteVehicleError,
	RatingStatusMissing:                         &pCErrorRatingStatusMissing,
	DriverUnacceptableRisk:                      &pCErrorDriverUnacceptableRisk,
	CreditPullError:                             &pCErrorCreditPullError,
	AdditionalDriversNeeded:                     &pCErrorAdditionalDriversNeeded,
	Fr19DataMissing:                             &pCErrorFr19DataMissing,
	AllDriversRequirePrimaryVehicle:             &pCErrorAllDriversRequirePrimaryVehicle,
	AdditionalDriverNeeded:                      &pCErrorAdditionalDriverNeeded,
	InvalidLicenseStatus:                        &pCErrorInvalidLicenseStatus,
	NoForeignLicense:                            &pCErrorNoForeignLicense,
	OutOfStateLicense:                           &pCErrorOutOfStateLicense,
	PrimaryDriverRequired:                       &pCErrorPrimaryDriverRequired,
	MissingDriver:                               &pCErrorMissingDriver,
	MissingCreditPull:                           &pCErrorMissingCreditPull,
	InvalidDriverID:                             &pCErrorInvalidDriverID,
	SaveDriverOccupation:                        &pCErrorSaveDriverOccupation,
	MapDriverOccupation:                         &pCErrorMapDriverOccupation,
	MapDriverBasicInfo:                          &pCErrorMapDriverBasicInfo,
	MapDriverIncidents:                          &pCErrorMapDriverIncidents,
	MapDriverError:                              &pCErrorMapDriverError,
	MapDriverLicenseDetails:                     &pCErrorMapDriverLicenseDetails,
	MapDriversError:                             &pCErrorMapDriversError,
	InvalidIncidentID:                           &pCErrorInvalidIncidentID,
	SaveDriverBasicInfo:                         &pCErrorSaveDriverBasicInfo,
	SaveDriverIncidents:                         &pCErrorSaveDriverIncidents,
	SaveDriverError:                             &pCErrorSaveDriverError,
	SaveDriverLicenseDetails:                    &pCErrorSaveDriverLicenseDetails,
	MinConvictionExceeded:                       &pCErrorMinConvictionExceeded,
	MissingMake:                                 &pCErrorMissingMake,
	MissingModel:                                &pCErrorMissingModel,
	MissingPrimaryUse:                           &pCErrorMissingPrimaryUse,
	InvalidVehicleDriverID:                      &pCErrorInvalidVehicleDriverID,
	CannotAddVehicleToJob:                       &pCErrorCannotAddVehicleToJob,
	InvalidVehicleID:                            &pCErrorInvalidVehicleID,
	CannotRemoveVehicle:                         &pCErrorCannotRemoveVehicle,
	CannotFindVehicleForDriver:                  &pCErrorCannotFindVehicleForDriver,
	LienLeaseRequired:                           &pCErrorLienLeaseRequired,
	AllVehicleRequirePrimaryDriver:              &pCErrorAllVehicleRequirePrimaryDriver,
	ExoticVehicle:                               &pCErrorExoticVehicle,
	InvalidVin:                                  &pCErrorInvalidVin,
	VehicleTypeRequired:                         &pCErrorVehicleTypeRequired,
	VehicleNewCostRequired:                      &pCErrorVehicleNewCostRequired,
	VehiclePositiveCostRequired:                 &pCErrorVehiclePositiveCostRequired,
	AtLeastOneVehicleRequired:                   &pCErrorAtLeastOneVehicleRequired,
	VinsMustBeUniqueForAllVehicles:              &pCErrorVinsMustBeUniqueForAllVehicles,
	VehiclePrimaryUseRequired:                   &pCErrorVehiclePrimaryUseRequired,
	SaveLoanLeinholder:                          &pCErrorSaveLoanLeinholder,
	SaveVehicleCoverages:                        &pCErrorSaveVehicleCoverages,
	MapVehiclesError:                            &pCErrorMapVehiclesError,
	SaveVehicleError:                            &pCErrorSaveVehicleError,
	VehicleBranded:                              &pCErrorVehicleBranded,
	VinIso:                                      &pCErrorVinIso,
	MissingPaperlessDiscount:                    &pCErrorMissingPaperlessDiscount,
	MissingResidency:                            &pCErrorMissingResidency,
	MissingESignatureDiscount:                   &pCErrorMissingESignatureDiscount,
	MissingDoYouCurrentlyHaveInsureanceQuestion: &pCErrorMissingDoYouCurrentlyHaveInsureanceQuestion,
	MissingSourceOfBusiness:                     &pCErrorMissingSourceOfBusiness,
	MissingCurrentCarrierName:                   &pCErrorMissingCurrentCarrierName,
	MissingCurrentCoverageEndDate:               &pCErrorMissingCurrentCoverageEndDate,
	MissingCurrentInjuryLimits:                  &pCErrorMissingCurrentInjuryLimits,
	RelationshipToInsuredMissing:                &pCErrorRelationshipToInsuredMissing,
	MissingLapseInCoverageQuestion:              &pCErrorMissingLapseInCoverageQuestion,
	MissingYearWithCurrentProviderQuestion:      &pCErrorMissingYearWithCurrentProviderQuestion,
	CannotRemoveApplicant:                       &pCErrorCannotRemoveApplicant,
	SaveCoverageTerms:                           &pCErrorSaveCoverageTerms,
	MapCoveragesError:                           &pCErrorMapCoveragesError,
	InvalidCollisionDeductible:                  &pCErrorInvalidCollisionDeductible,
	DiminishingDeductibleOnVehicles:             &pCErrorDiminishingDeductibleOnVehicles,
	SavingManualIncidentsError:                  &pCErrorSavingManualIncidentsError,
	SavingClueIncidentsError:                    &pCErrorSavingClueIncidentsError,
	SavingMvrIncidentsError:                     &pCErrorSavingMvrIncidentsError,
	UnknownSaveIncidentsError:                   &pCErrorUnknownSaveIncidentsError,
	QuoteRequiresPhotoInspection:                &pCErrorQuoteRequiresPhotoInspection,
	MissingMvrReport:                            &pCErrorMissingMvrReport,
	DuiLimitExceeded:                            &pCErrorDuiLimitExceeded,
	AtFaultLimitExceeded:                        &pCErrorAtFaultLimitExceeded,
	MinMajorConvictionLimitExceeded:             &pCErrorMinMajorConvictionLimitExceeded,
	MapClueIncidentsError:                       &pCErrorMapClueIncidentsError,
	MapMvrIncidentsError:                        &pCErrorMapMvrIncidentsError,
	MapManualIncidentsError:                     &pCErrorMapManualIncidentsError,
	MapIncidentsUnknownError:                    &pCErrorMapIncidentsUnknownError,
	SavingIncidentsError:                        &pCErrorSavingIncidentsError,
	ShouldPullCreditError:                       &pCErrorShouldPullCreditError,
	PullReportsError:                            &pCErrorPullReportsError,
	MajorConvictionsLimitExceeded:               &pCErrorMajorConvictionsLimitExceeded,
	PipClaimsLimitExceeded:                      &pCErrorPipClaimsLimitExceeded,
	ChargeableAccidentsLimitExceeded:            &pCErrorChargeableAccidentsLimitExceeded,
	JobNotInQuoteStatus:                         &pCErrorJobNotInQuoteStatus,
	AccountHasActivePolicy:                      &pCErrorAccountHasActivePolicy,
	UnderwritingBlockingIssue:                   &pCErrorUnderwritingBlockingIssue,
	ClueReportRequired:                          &pCErrorClueReportRequired,
	BackDateError:                               &pCErrorBackDateError,
	ExcludedDriverSendDocument:                  &pCErrorExcludedDriverSendDocument,
	ExcludedDriverResendDocument:                &pCErrorExcludedDriverResendDocument,
	ExcludedDriverFormNotSigned:                 &pCErrorExcludedDriverFormNotSigned,
	AccidentDayOfBind:                           &pCErrorAccidentDayOfBind,
	MissingPaymentplan:                          &pCErrorMissingPaymentplan,
	PolicyHasPhotoReview:                        &pCErrorPolicyHasPhotoReview,
	DocusignIntegrationError:                    &pCErrorDocusignIntegrationError,
	CannotBindAccidentDayOfBind:                 &pCErrorCannotBindAccidentDayOfBind,
	DriverNotAssignedForLegal:                   &pCErrorDriverNotAssignedForLegal,
	IneligibleForFamilyLegal:                    &pCErrorIneligibleForFamilyLegal,
	NullDmsReason:                               &pCErrorNullDmsReason,
	TerritoryRestriction:                        &pCErrorTerritoryRestriction,
	AccountHasMaterialMisrep:                    &pCErrorAccountHasMaterialMisrep,
	IllegalBundleTransferException:              &pCErrorIllegalBundleTransferException,
	ConcurrentDataChangeException:               &pCErrorConcurrentDataChangeException,
	GenerateProductsError:                       &pCErrorGenerateProductsError,
	QuotePolicyError:                            &pCErrorQuotePolicyError,
	BindPolicyError:                             &pCErrorBindPolicyError,
	FetchDocumentsError:                         &pCErrorFetchDocumentsError,
	InvalidPaymentPlan:                          &pCErrorInvalidPaymentPlan,
	InvalidPolicyData:                           &pCErrorInvalidPolicyData,
	NoDocuments:                                 &pCErrorNoDocuments,
	NoJobFound:                                  &pCErrorNoJobFound,
	DeleteItemError:                             &pCErrorDeleteItemError,
	WorkorderNotQuotable:                        &pCErrorWorkorderNotQuotable,
	UnknownPaymentError:                         &pCErrorUnknownPaymentError,
	NoCreditCardNumber:                          &pCErrorNoCreditCardNumber,
	FeePaidWithDownPayment:                      &pCErrorFeePaidWithDownPayment,
	CardDeclined:                                &pCErrorCardDeclined,
	InvalidCardNumber:                           &pCErrorInvalidCardNumber,
	InvalidExpiration:                           &pCErrorInvalidExpiration,
	InvalidCardType:                             &pCErrorInvalidCardType,
	InsufficientFunds:                           &pCErrorInsufficientFunds,
	InvalidAbaNumber:                            &pCErrorInvalidAbaNumber,
	InvalidBankAccountNumber:                    &pCErrorInvalidBankAccountNumber,
	InvalidBankName:                             &pCErrorInvalidBankName,
	BankAccountType:                             &pCErrorBankAccountType,
	LastNameMissing:                             &pCErrorLastNameMissing,
	MissingAccountContact:                       &pCErrorMissingAccountContact,
	ValidateFcraNotice:                          &pCErrorValidateFcraNotice,
	PolicyHolderWithPermitLicense:               &pCErrorPolicyHolderWithPermitLicense,

	itemDict: map[string]*EnumPCErrorItem{
		strings.ToLower(string(pCErrorMissingModelYearID)):                            &pCErrorMissingModelYear,
		strings.ToLower(string(pCErrorInvalidDateFormatID)):                           &pCErrorInvalidDateFormat,
		strings.ToLower(string(pCErrorInvalidStateID)):                                &pCErrorInvalidState,
		strings.ToLower(string(pCErrorAccountNumberMissingID)):                        &pCErrorAccountNumberMissing,
		strings.ToLower(string(pCErrorAccountNotFoundID)):                             &pCErrorAccountNotFound,
		strings.ToLower(string(pCErrorMapAccountErrorID)):                             &pCErrorMapAccountError,
		strings.ToLower(string(pCErrorSaveAccountErrorID)):                            &pCErrorSaveAccountError,
		strings.ToLower(string(pCErrorHasExistingAccountErrorID)):                     &pCErrorHasExistingAccountError,
		strings.ToLower(string(pCErrorSaveAccountHolderErrorID)):                      &pCErrorSaveAccountHolderError,
		strings.ToLower(string(pCErrorGenerateAccountErrorID)):                        &pCErrorGenerateAccountError,
		strings.ToLower(string(pCErrorCreateAccountErrorID)):                          &pCErrorCreateAccountError,
		strings.ToLower(string(pCErrorGetRecentAccountErrorID)):                       &pCErrorGetRecentAccountError,
		strings.ToLower(string(pCErrorAccountHasBadDebtID)):                           &pCErrorAccountHasBadDebt,
		strings.ToLower(string(pCErrorAccountLookupErrorID)):                          &pCErrorAccountLookupError,
		strings.ToLower(string(pCErrorFirstNameMissingID)):                            &pCErrorFirstNameMissing,
		strings.ToLower(string(pCErrorInvalidEffectiveDateID)):                        &pCErrorInvalidEffectiveDate,
		strings.ToLower(string(pCErrorDobMissingID)):                                  &pCErrorDobMissing,
		strings.ToLower(string(pCErrorHomePhoneMissingID)):                            &pCErrorHomePhoneMissing,
		strings.ToLower(string(pCErrorAddressLine1MissingID)):                         &pCErrorAddressLine1Missing,
		strings.ToLower(string(pCErrorCityMissingID)):                                 &pCErrorCityMissing,
		strings.ToLower(string(pCErrorStateMissingID)):                                &pCErrorStateMissing,
		strings.ToLower(string(pCErrorPostalMissingID)):                               &pCErrorPostalMissing,
		strings.ToLower(string(pCErrorGenderMissingID)):                               &pCErrorGenderMissing,
		strings.ToLower(string(pCErrorLicenseStateOrNumberMissingID)):                 &pCErrorLicenseStateOrNumberMissing,
		strings.ToLower(string(pCErrorHighestEducationLevelMissingID)):                &pCErrorHighestEducationLevelMissing,
		strings.ToLower(string(pCErrorOccupationTitleMissingID)):                      &pCErrorOccupationTitleMissing,
		strings.ToLower(string(pCErrorCannotUpdateSameDayBindQuestionID)):             &pCErrorCannotUpdateSameDayBindQuestion,
		strings.ToLower(string(pCErrorInvalidPhoneNumberID)):                          &pCErrorInvalidPhoneNumber,
		strings.ToLower(string(pCErrorInvalidEmailID)):                                &pCErrorInvalidEmail,
		strings.ToLower(string(pCErrorInvalidOnlineDiscountID)):                       &pCErrorInvalidOnlineDiscount,
		strings.ToLower(string(pCErrorInvalidWebToPhoneID)):                           &pCErrorInvalidWebToPhone,
		strings.ToLower(string(pCErrorValidatePaperlessDiscountID)):                   &pCErrorValidatePaperlessDiscount,
		strings.ToLower(string(pCErrorInvalidPolicyEffectiveDateID)):                  &pCErrorInvalidPolicyEffectiveDate,
		strings.ToLower(string(pCErrorInvalidTransactiveEffecitveID)):                 &pCErrorInvalidTransactiveEffecitve,
		strings.ToLower(string(pCErrorMarriedWithNoSpouseID)):                         &pCErrorMarriedWithNoSpouse,
		strings.ToLower(string(pCErrorMarriedWithMultpleSpousesID)):                   &pCErrorMarriedWithMultpleSpouses,
		strings.ToLower(string(pCErrorEarlyBirdDiscountID)):                           &pCErrorEarlyBirdDiscount,
		strings.ToLower(string(pCErrorCountyMissingID)):                               &pCErrorCountyMissing,
		strings.ToLower(string(pCErrorJobNumberMissingID)):                            &pCErrorJobNumberMissing,
		strings.ToLower(string(pCErrorPolicynumberNumberMissingID)):                   &pCErrorPolicynumberNumberMissing,
		strings.ToLower(string(pCErrorMapQualifiedDiscountsID)):                       &pCErrorMapQualifiedDiscounts,
		strings.ToLower(string(pCErrorCannotCreateSavePolicyperiodID)):                &pCErrorCannotCreateSavePolicyperiod,
		strings.ToLower(string(pCErrorMapPolicyErrorID)):                              &pCErrorMapPolicyError,
		strings.ToLower(string(pCErrorMapPolicyInitDataErrorID)):                      &pCErrorMapPolicyInitDataError,
		strings.ToLower(string(pCErrorMapPolicyLineDetailsErrorID)):                   &pCErrorMapPolicyLineDetailsError,
		strings.ToLower(string(pCErrorMapQuotedValuesErrorID)):                        &pCErrorMapQuotedValuesError,
		strings.ToLower(string(pCErrorSavePolicyErrorID)):                             &pCErrorSavePolicyError,
		strings.ToLower(string(pCErrorRetrieveQuoteErrorID)):                          &pCErrorRetrieveQuoteError,
		strings.ToLower(string(pCErrorInvalidJobStatusID)):                            &pCErrorInvalidJobStatus,
		strings.ToLower(string(pCErrorCannotEditJobID)):                               &pCErrorCannotEditJob,
		strings.ToLower(string(pCErrorDeleteDriverErrorID)):                           &pCErrorDeleteDriverError,
		strings.ToLower(string(pCErrorDeleteManualIncidentsErrorID)):                  &pCErrorDeleteManualIncidentsError,
		strings.ToLower(string(pCErrorDeleteVehicleErrorID)):                          &pCErrorDeleteVehicleError,
		strings.ToLower(string(pCErrorRatingStatusMissingID)):                         &pCErrorRatingStatusMissing,
		strings.ToLower(string(pCErrorDriverUnacceptableRiskID)):                      &pCErrorDriverUnacceptableRisk,
		strings.ToLower(string(pCErrorCreditPullErrorID)):                             &pCErrorCreditPullError,
		strings.ToLower(string(pCErrorAdditionalDriversNeededID)):                     &pCErrorAdditionalDriversNeeded,
		strings.ToLower(string(pCErrorFr19DataMissingID)):                             &pCErrorFr19DataMissing,
		strings.ToLower(string(pCErrorAllDriversRequirePrimaryVehicleID)):             &pCErrorAllDriversRequirePrimaryVehicle,
		strings.ToLower(string(pCErrorAdditionalDriverNeededID)):                      &pCErrorAdditionalDriverNeeded,
		strings.ToLower(string(pCErrorInvalidLicenseStatusID)):                        &pCErrorInvalidLicenseStatus,
		strings.ToLower(string(pCErrorNoForeignLicenseID)):                            &pCErrorNoForeignLicense,
		strings.ToLower(string(pCErrorOutOfStateLicenseID)):                           &pCErrorOutOfStateLicense,
		strings.ToLower(string(pCErrorPrimaryDriverRequiredID)):                       &pCErrorPrimaryDriverRequired,
		strings.ToLower(string(pCErrorMissingDriverID)):                               &pCErrorMissingDriver,
		strings.ToLower(string(pCErrorMissingCreditPullID)):                           &pCErrorMissingCreditPull,
		strings.ToLower(string(pCErrorInvalidDriverIDID)):                             &pCErrorInvalidDriverID,
		strings.ToLower(string(pCErrorSaveDriverOccupationID)):                        &pCErrorSaveDriverOccupation,
		strings.ToLower(string(pCErrorMapDriverOccupationID)):                         &pCErrorMapDriverOccupation,
		strings.ToLower(string(pCErrorMapDriverBasicInfoID)):                          &pCErrorMapDriverBasicInfo,
		strings.ToLower(string(pCErrorMapDriverIncidentsID)):                          &pCErrorMapDriverIncidents,
		strings.ToLower(string(pCErrorMapDriverErrorID)):                              &pCErrorMapDriverError,
		strings.ToLower(string(pCErrorMapDriverLicenseDetailsID)):                     &pCErrorMapDriverLicenseDetails,
		strings.ToLower(string(pCErrorMapDriversErrorID)):                             &pCErrorMapDriversError,
		strings.ToLower(string(pCErrorInvalidIncidentIDID)):                           &pCErrorInvalidIncidentID,
		strings.ToLower(string(pCErrorSaveDriverBasicInfoID)):                         &pCErrorSaveDriverBasicInfo,
		strings.ToLower(string(pCErrorSaveDriverIncidentsID)):                         &pCErrorSaveDriverIncidents,
		strings.ToLower(string(pCErrorSaveDriverErrorID)):                             &pCErrorSaveDriverError,
		strings.ToLower(string(pCErrorSaveDriverLicenseDetailsID)):                    &pCErrorSaveDriverLicenseDetails,
		strings.ToLower(string(pCErrorMinConvictionExceededID)):                       &pCErrorMinConvictionExceeded,
		strings.ToLower(string(pCErrorMissingMakeID)):                                 &pCErrorMissingMake,
		strings.ToLower(string(pCErrorMissingModelID)):                                &pCErrorMissingModel,
		strings.ToLower(string(pCErrorMissingPrimaryUseID)):                           &pCErrorMissingPrimaryUse,
		strings.ToLower(string(pCErrorInvalidVehicleDriverIDID)):                      &pCErrorInvalidVehicleDriverID,
		strings.ToLower(string(pCErrorCannotAddVehicleToJobID)):                       &pCErrorCannotAddVehicleToJob,
		strings.ToLower(string(pCErrorInvalidVehicleIDID)):                            &pCErrorInvalidVehicleID,
		strings.ToLower(string(pCErrorCannotRemoveVehicleID)):                         &pCErrorCannotRemoveVehicle,
		strings.ToLower(string(pCErrorCannotFindVehicleForDriverID)):                  &pCErrorCannotFindVehicleForDriver,
		strings.ToLower(string(pCErrorLienLeaseRequiredID)):                           &pCErrorLienLeaseRequired,
		strings.ToLower(string(pCErrorAllVehicleRequirePrimaryDriverID)):              &pCErrorAllVehicleRequirePrimaryDriver,
		strings.ToLower(string(pCErrorExoticVehicleID)):                               &pCErrorExoticVehicle,
		strings.ToLower(string(pCErrorInvalidVinID)):                                  &pCErrorInvalidVin,
		strings.ToLower(string(pCErrorVehicleTypeRequiredID)):                         &pCErrorVehicleTypeRequired,
		strings.ToLower(string(pCErrorVehicleNewCostRequiredID)):                      &pCErrorVehicleNewCostRequired,
		strings.ToLower(string(pCErrorVehiclePositiveCostRequiredID)):                 &pCErrorVehiclePositiveCostRequired,
		strings.ToLower(string(pCErrorAtLeastOneVehicleRequiredID)):                   &pCErrorAtLeastOneVehicleRequired,
		strings.ToLower(string(pCErrorVinsMustBeUniqueForAllVehiclesID)):              &pCErrorVinsMustBeUniqueForAllVehicles,
		strings.ToLower(string(pCErrorVehiclePrimaryUseRequiredID)):                   &pCErrorVehiclePrimaryUseRequired,
		strings.ToLower(string(pCErrorSaveLoanLeinholderID)):                          &pCErrorSaveLoanLeinholder,
		strings.ToLower(string(pCErrorSaveVehicleCoveragesID)):                        &pCErrorSaveVehicleCoverages,
		strings.ToLower(string(pCErrorMapVehiclesErrorID)):                            &pCErrorMapVehiclesError,
		strings.ToLower(string(pCErrorSaveVehicleErrorID)):                            &pCErrorSaveVehicleError,
		strings.ToLower(string(pCErrorVehicleBrandedID)):                              &pCErrorVehicleBranded,
		strings.ToLower(string(pCErrorVinIsoID)):                                      &pCErrorVinIso,
		strings.ToLower(string(pCErrorMissingPaperlessDiscountID)):                    &pCErrorMissingPaperlessDiscount,
		strings.ToLower(string(pCErrorMissingResidencyID)):                            &pCErrorMissingResidency,
		strings.ToLower(string(pCErrorMissingESignatureDiscountID)):                   &pCErrorMissingESignatureDiscount,
		strings.ToLower(string(pCErrorMissingDoYouCurrentlyHaveInsureanceQuestionID)): &pCErrorMissingDoYouCurrentlyHaveInsureanceQuestion,
		strings.ToLower(string(pCErrorMissingSourceOfBusinessID)):                     &pCErrorMissingSourceOfBusiness,
		strings.ToLower(string(pCErrorMissingCurrentCarrierNameID)):                   &pCErrorMissingCurrentCarrierName,
		strings.ToLower(string(pCErrorMissingCurrentCoverageEndDateID)):               &pCErrorMissingCurrentCoverageEndDate,
		strings.ToLower(string(pCErrorMissingCurrentInjuryLimitsID)):                  &pCErrorMissingCurrentInjuryLimits,
		strings.ToLower(string(pCErrorRelationshipToInsuredMissingID)):                &pCErrorRelationshipToInsuredMissing,
		strings.ToLower(string(pCErrorMissingLapseInCoverageQuestionID)):              &pCErrorMissingLapseInCoverageQuestion,
		strings.ToLower(string(pCErrorMissingYearWithCurrentProviderQuestionID)):      &pCErrorMissingYearWithCurrentProviderQuestion,
		strings.ToLower(string(pCErrorCannotRemoveApplicantID)):                       &pCErrorCannotRemoveApplicant,
		strings.ToLower(string(pCErrorSaveCoverageTermsID)):                           &pCErrorSaveCoverageTerms,
		strings.ToLower(string(pCErrorMapCoveragesErrorID)):                           &pCErrorMapCoveragesError,
		strings.ToLower(string(pCErrorInvalidCollisionDeductibleID)):                  &pCErrorInvalidCollisionDeductible,
		strings.ToLower(string(pCErrorDiminishingDeductibleOnVehiclesID)):             &pCErrorDiminishingDeductibleOnVehicles,
		strings.ToLower(string(pCErrorSavingManualIncidentsErrorID)):                  &pCErrorSavingManualIncidentsError,
		strings.ToLower(string(pCErrorSavingClueIncidentsErrorID)):                    &pCErrorSavingClueIncidentsError,
		strings.ToLower(string(pCErrorSavingMvrIncidentsErrorID)):                     &pCErrorSavingMvrIncidentsError,
		strings.ToLower(string(pCErrorUnknownSaveIncidentsErrorID)):                   &pCErrorUnknownSaveIncidentsError,
		strings.ToLower(string(pCErrorQuoteRequiresPhotoInspectionID)):                &pCErrorQuoteRequiresPhotoInspection,
		strings.ToLower(string(pCErrorMissingMvrReportID)):                            &pCErrorMissingMvrReport,
		strings.ToLower(string(pCErrorDuiLimitExceededID)):                            &pCErrorDuiLimitExceeded,
		strings.ToLower(string(pCErrorAtFaultLimitExceededID)):                        &pCErrorAtFaultLimitExceeded,
		strings.ToLower(string(pCErrorMinMajorConvictionLimitExceededID)):             &pCErrorMinMajorConvictionLimitExceeded,
		strings.ToLower(string(pCErrorMapClueIncidentsErrorID)):                       &pCErrorMapClueIncidentsError,
		strings.ToLower(string(pCErrorMapMvrIncidentsErrorID)):                        &pCErrorMapMvrIncidentsError,
		strings.ToLower(string(pCErrorMapManualIncidentsErrorID)):                     &pCErrorMapManualIncidentsError,
		strings.ToLower(string(pCErrorMapIncidentsUnknownErrorID)):                    &pCErrorMapIncidentsUnknownError,
		strings.ToLower(string(pCErrorSavingIncidentsErrorID)):                        &pCErrorSavingIncidentsError,
		strings.ToLower(string(pCErrorShouldPullCreditErrorID)):                       &pCErrorShouldPullCreditError,
		strings.ToLower(string(pCErrorPullReportsErrorID)):                            &pCErrorPullReportsError,
		strings.ToLower(string(pCErrorMajorConvictionsLimitExceededID)):               &pCErrorMajorConvictionsLimitExceeded,
		strings.ToLower(string(pCErrorPipClaimsLimitExceededID)):                      &pCErrorPipClaimsLimitExceeded,
		strings.ToLower(string(pCErrorChargeableAccidentsLimitExceededID)):            &pCErrorChargeableAccidentsLimitExceeded,
		strings.ToLower(string(pCErrorJobNotInQuoteStatusID)):                         &pCErrorJobNotInQuoteStatus,
		strings.ToLower(string(pCErrorAccountHasActivePolicyID)):                      &pCErrorAccountHasActivePolicy,
		strings.ToLower(string(pCErrorUnderwritingBlockingIssueID)):                   &pCErrorUnderwritingBlockingIssue,
		strings.ToLower(string(pCErrorClueReportRequiredID)):                          &pCErrorClueReportRequired,
		strings.ToLower(string(pCErrorBackDateErrorID)):                               &pCErrorBackDateError,
		strings.ToLower(string(pCErrorExcludedDriverSendDocumentID)):                  &pCErrorExcludedDriverSendDocument,
		strings.ToLower(string(pCErrorExcludedDriverResendDocumentID)):                &pCErrorExcludedDriverResendDocument,
		strings.ToLower(string(pCErrorExcludedDriverFormNotSignedID)):                 &pCErrorExcludedDriverFormNotSigned,
		strings.ToLower(string(pCErrorAccidentDayOfBindID)):                           &pCErrorAccidentDayOfBind,
		strings.ToLower(string(pCErrorMissingPaymentplanID)):                          &pCErrorMissingPaymentplan,
		strings.ToLower(string(pCErrorPolicyHasPhotoReviewID)):                        &pCErrorPolicyHasPhotoReview,
		strings.ToLower(string(pCErrorDocusignIntegrationErrorID)):                    &pCErrorDocusignIntegrationError,
		strings.ToLower(string(pCErrorCannotBindAccidentDayOfBindID)):                 &pCErrorCannotBindAccidentDayOfBind,
		strings.ToLower(string(pCErrorDriverNotAssignedForLegalID)):                   &pCErrorDriverNotAssignedForLegal,
		strings.ToLower(string(pCErrorIneligibleForFamilyLegalID)):                    &pCErrorIneligibleForFamilyLegal,
		strings.ToLower(string(pCErrorNullDmsReasonID)):                               &pCErrorNullDmsReason,
		strings.ToLower(string(pCErrorTerritoryRestrictionID)):                        &pCErrorTerritoryRestriction,
		strings.ToLower(string(pCErrorAccountHasMaterialMisrepID)):                    &pCErrorAccountHasMaterialMisrep,
		strings.ToLower(string(pCErrorIllegalBundleTransferExceptionID)):              &pCErrorIllegalBundleTransferException,
		strings.ToLower(string(pCErrorConcurrentDataChangeExceptionID)):               &pCErrorConcurrentDataChangeException,
		strings.ToLower(string(pCErrorGenerateProductsErrorID)):                       &pCErrorGenerateProductsError,
		strings.ToLower(string(pCErrorQuotePolicyErrorID)):                            &pCErrorQuotePolicyError,
		strings.ToLower(string(pCErrorBindPolicyErrorID)):                             &pCErrorBindPolicyError,
		strings.ToLower(string(pCErrorFetchDocumentsErrorID)):                         &pCErrorFetchDocumentsError,
		strings.ToLower(string(pCErrorInvalidPaymentPlanID)):                          &pCErrorInvalidPaymentPlan,
		strings.ToLower(string(pCErrorInvalidPolicyDataID)):                           &pCErrorInvalidPolicyData,
		strings.ToLower(string(pCErrorNoDocumentsID)):                                 &pCErrorNoDocuments,
		strings.ToLower(string(pCErrorNoJobFoundID)):                                  &pCErrorNoJobFound,
		strings.ToLower(string(pCErrorDeleteItemErrorID)):                             &pCErrorDeleteItemError,
		strings.ToLower(string(pCErrorWorkorderNotQuotableID)):                        &pCErrorWorkorderNotQuotable,
		strings.ToLower(string(pCErrorUnknownPaymentErrorID)):                         &pCErrorUnknownPaymentError,
		strings.ToLower(string(pCErrorNoCreditCardNumberID)):                          &pCErrorNoCreditCardNumber,
		strings.ToLower(string(pCErrorFeePaidWithDownPaymentID)):                      &pCErrorFeePaidWithDownPayment,
		strings.ToLower(string(pCErrorCardDeclinedID)):                                &pCErrorCardDeclined,
		strings.ToLower(string(pCErrorInvalidCardNumberID)):                           &pCErrorInvalidCardNumber,
		strings.ToLower(string(pCErrorInvalidExpirationID)):                           &pCErrorInvalidExpiration,
		strings.ToLower(string(pCErrorInvalidCardTypeID)):                             &pCErrorInvalidCardType,
		strings.ToLower(string(pCErrorInsufficientFundsID)):                           &pCErrorInsufficientFunds,
		strings.ToLower(string(pCErrorInvalidAbaNumberID)):                            &pCErrorInvalidAbaNumber,
		strings.ToLower(string(pCErrorInvalidBankAccountNumberID)):                    &pCErrorInvalidBankAccountNumber,
		strings.ToLower(string(pCErrorInvalidBankNameID)):                             &pCErrorInvalidBankName,
		strings.ToLower(string(pCErrorBankAccountTypeID)):                             &pCErrorBankAccountType,
		strings.ToLower(string(pCErrorLastNameMissingID)):                             &pCErrorLastNameMissing,
		strings.ToLower(string(pCErrorMissingAccountContactID)):                       &pCErrorMissingAccountContact,
		strings.ToLower(string(pCErrorValidateFcraNoticeID)):                          &pCErrorValidateFcraNotice,
		strings.ToLower(string(pCErrorPolicyHolderWithPermitLicenseID)):               &pCErrorPolicyHolderWithPermitLicense,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumPCError) ByID(id PCErrorIdentifier) *EnumPCErrorItem {
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
func (e *EnumPCError) ByIDString(idx string) *EnumPCErrorItem {
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
func (e *EnumPCError) ByIndex(idx int) *EnumPCErrorItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedPCErrorID is a struct that is designed to replace a *PCErrorID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *PCErrorID it contains while being a better JSON citizen.
type ValidatedPCErrorID struct {
	// id will point to a valid PCErrorID, if possible
	// If id is nil, then ValidatedPCErrorID.Valid() will return false.
	id *PCErrorID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedPCErrorID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedPCErrorID
func (vi *ValidatedPCErrorID) Clone() *ValidatedPCErrorID {
	if vi == nil {
		return nil
	}

	var cid *PCErrorID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedPCErrorID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedPCErrorIds represent the same PCError
func (vi *ValidatedPCErrorID) Equals(vj *ValidatedPCErrorID) bool {
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

// Valid returns true if and only if the ValidatedPCErrorID corresponds to a recognized PCError
func (vi *ValidatedPCErrorID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedPCErrorID) ID() *PCErrorID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedPCErrorID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedPCErrorID) ValidatedID() *ValidatedPCErrorID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedPCErrorID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedPCErrorID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedPCErrorID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedPCErrorID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedPCErrorID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := PCErrorID(capString)
	item := PCError.ByID(&id)
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

func (vi ValidatedPCErrorID) String() string {
	return vi.ToIDString()
}

type PCErrorIdentifier interface {
	ID() *PCErrorID
	Valid() bool
}
