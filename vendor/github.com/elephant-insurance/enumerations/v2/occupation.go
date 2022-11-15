package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// OccupationID uniquely identifies a particular Occupation
type OccupationID string

// Clone creates a safe, independent copy of a OccupationID
func (i *OccupationID) Clone() *OccupationID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two OccupationIds are equivalent
func (i *OccupationID) Equals(j *OccupationID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *OccupationID that is either valid or nil
func (i *OccupationID) ID() *OccupationID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *OccupationID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the OccupationID corresponds to a recognized Occupation
func (i *OccupationID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return Occupation.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *OccupationID) ValidatedID() *ValidatedOccupationID {
	if i != nil {
		return &ValidatedOccupationID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *OccupationID) MarshalJSON() ([]byte, error) {
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

func (i *OccupationID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := OccupationID(dataString)
	item := Occupation.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	occupationAccountExecutiveID                           OccupationID = "1"
	occupationAccountantorCPAID                            OccupationID = "2"
	occupationActorActressID                               OccupationID = "3"
	occupationActuaryID                                    OccupationID = "4"
	occupationAcupuncturistID                              OccupationID = "5"
	occupationAdjudicatorID                                OccupationID = "6"
	occupationAdministrativeAssistantID                    OccupationID = "7"
	occupationAdministrativeClerkorReceptionistID          OccupationID = "8"
	occupationAdministratorID                              OccupationID = "9"
	occupationAdministratorHealthorHospitalID              OccupationID = "10"
	occupationAdministratorPublicID                        OccupationID = "11"
	occupationAdvisorFinancialorCreditID                   OccupationID = "12"
	occupationAgentAdvertisingID                           OccupationID = "13"
	occupationAgentBorderPatrolID                          OccupationID = "14"
	occupationAgentEmploymentID                            OccupationID = "15"
	occupationAgentFBIID                                   OccupationID = "16"
	occupationAgentImportExportID                          OccupationID = "17"
	occupationAgentInsuranceID                             OccupationID = "18"
	occupationAgentIRSID                                   OccupationID = "19"
	occupationAgentLeasingID                               OccupationID = "20"
	occupationAgentRealEstateID                            OccupationID = "21"
	occupationAgentSecretServiceID                         OccupationID = "22"
	occupationAgentTheatricalID                            OccupationID = "23"
	occupationAgentTravelTicketID                          OccupationID = "24"
	occupationAgentTreasuryID                              OccupationID = "25"
	occupationAideAllOtherID                               OccupationID = "26"
	occupationAirTrafficControllerID                       OccupationID = "27"
	occupationAmbassadorID                                 OccupationID = "28"
	occupationAnalystAllOtherID                            OccupationID = "29"
	occupationAnalystEngineeringID                         OccupationID = "30"
	occupationAnalystFinancialID                           OccupationID = "31"
	occupationAnesthesiologistID                           OccupationID = "32"
	occupationAntiqueDealerID                              OccupationID = "33"
	occupationAppraiserAllOtherID                          OccupationID = "34"
	occupationAppraiserArtID                               OccupationID = "35"
	occupationAppraiserRealEstateID                        OccupationID = "36"
	occupationApprenticeAllOtherID                         OccupationID = "37"
	occupationArchitectID                                  OccupationID = "38"
	occupationArchivistLibraryorMuseumID                   OccupationID = "39"
	occupationArtistCommercialID                           OccupationID = "40"
	occupationArtistNoncommercialID                        OccupationID = "41"
	occupationAssistantAllOtherID                          OccupationID = "42"
	occupationAssistantEditorialID                         OccupationID = "43"
	occupationAssistantLegislativeorLegalID                OccupationID = "44"
	occupationAssistantMedicalID                           OccupationID = "45"
	occupationAssistantPrintingPressID                     OccupationID = "46"
	occupationAthleteID                                    OccupationID = "47"
	occupationAttorneyorLawyerID                           OccupationID = "48"
	occupationAuctioneerID                                 OccupationID = "49"
	occupationAudiologistID                                OccupationID = "50"
	occupationAuditorFinancialID                           OccupationID = "51"
	occupationAuditorNonFinancialID                        OccupationID = "52"
	occupationBakerID                                      OccupationID = "53"
	occupationBankTellerID                                 OccupationID = "54"
	occupationBartenderID                                  OccupationID = "55"
	occupationBoatCaptainID                                OccupationID = "56"
	occupationBoatmanID                                    OccupationID = "57"
	occupationBondsmanID                                   OccupationID = "58"
	occupationBookbinderID                                 OccupationID = "59"
	occupationBookkeeperID                                 OccupationID = "60"
	occupationBroadcasterID                                OccupationID = "61"
	occupationBrokerInsuranceID                            OccupationID = "62"
	occupationBrokerMortgageID                             OccupationID = "63"
	occupationBrokerPawnID                                 OccupationID = "64"
	occupationBrokerRealEstateID                           OccupationID = "65"
	occupationBrokerStockBondsID                           OccupationID = "66"
	occupationBuyerPurchasingAgentID                       OccupationID = "67"
	occupationCameramanID                                  OccupationID = "68"
	occupationCarpenterorCabinetmakerID                    OccupationID = "69"
	occupationCashierFoodID                                OccupationID = "70"
	occupationCashierOfficeID                              OccupationID = "71"
	occupationCashierRetailID                              OccupationID = "72"
	occupationCatererID                                    OccupationID = "73"
	occupationCensusTakerID                                OccupationID = "74"
	occupationChiropractorID                               OccupationID = "75"
	occupationClaimsAdjusterID                             OccupationID = "76"
	occupationClaimsExaminerID                             OccupationID = "77"
	occupationClerkAccountingorFinancialID                 OccupationID = "78"
	occupationClerkAllOtherID                              OccupationID = "79"
	occupationCollegeDeanID                                OccupationID = "80"
	occupationCommunicationSpecialistID                    OccupationID = "81"
	occupationComputerProgrammerID                         OccupationID = "82"
	occupationComputerTechnicalSupportRepID                OccupationID = "83"
	occupationConductorID                                  OccupationID = "84"
	occupationConservationistID                            OccupationID = "85"
	occupationConstructionWorkerID                         OccupationID = "86"
	occupationConsultantorAdvisorID                        OccupationID = "87"
	occupationContractororDeveloperAllOtherID              OccupationID = "88"
	occupationContractororDeveloperArtisanSkilledID        OccupationID = "89"
	occupationContractororDeveloperBlueCollar50EmployeesID OccupationID = "90"
	occupationContractororDeveloperWhiteCollarID           OccupationID = "91"
	occupationControllerorComptrollerID                    OccupationID = "92"
	occupationControllerorCoordinatorProductionID          OccupationID = "93"
	occupationCookorChefID                                 OccupationID = "94"
	occupationCoordinatorOfficeID                          OccupationID = "95"
	occupationCopywriterID                                 OccupationID = "96"
	occupationCoronerID                                    OccupationID = "97"
	occupationCosmetologistBeauticianID                    OccupationID = "98"
	occupationCounselorAllOtherID                          OccupationID = "99"
	occupationCounselorEducationID                         OccupationID = "100"
	occupationCounselorFamilyAndChildID                    OccupationID = "101"
	occupationCounselorMentalHealthID                      OccupationID = "102"
	occupationCourtReporterID                              OccupationID = "103"
	occupationCraftsmanID                                  OccupationID = "104"
	occupationCriminologistID                              OccupationID = "105"
	occupationCuratorID                                    OccupationID = "106"
	occupationCustodianorJanitorID                         OccupationID = "107"
	occupationCustomerServiceRepresentativeID              OccupationID = "108"
	occupationDayCareWorkerID                              OccupationID = "109"
	occupationDecoratorInteriorID                          OccupationID = "110"
	occupationDeliveryPersonorMailCarrierID                OccupationID = "111"
	occupationDentalHygienistID                            OccupationID = "112"
	occupationDentistID                                    OccupationID = "113"
	occupationDesignerComputerWebsiteID                    OccupationID = "114"
	occupationDesignerFloralID                             OccupationID = "115"
	occupationDesignerGraphicorTechnicalID                 OccupationID = "116"
	occupationDesignerProfessionalID                       OccupationID = "117"
	occupationDesignerWindowID                             OccupationID = "118"
	occupationDietitianNutritionistID                      OccupationID = "119"
	occupationDirectororExecutiveID                        OccupationID = "120"
	occupationDiscJockeyID                                 OccupationID = "121"
	occupationDispatcherID                                 OccupationID = "122"
	occupationDogBreederID                                 OccupationID = "123"
	occupationDrafterorCartographerID                      OccupationID = "124"
	occupationDriverAllOtherID                             OccupationID = "125"
	occupationDriverTruckID                                OccupationID = "126"
	occupationEconomistID                                  OccupationID = "127"
	occupationEditorAllOtherID                             OccupationID = "128"
	occupationEditorFilmID                                 OccupationID = "129"
	occupationElectricianID                                OccupationID = "130"
	occupationEmbalmerID                                   OccupationID = "131"
	occupationEngineerAllOtherID                           OccupationID = "132"
	occupationEngineerCertifiedNetworkID                   OccupationID = "133"
	occupationEngineerComputerSoftwareID                   OccupationID = "134"
	occupationEngineerComputerSystemsID                    OccupationID = "135"
	occupationEngineerConstructionID                       OccupationID = "136"
	occupationEngineerElectricalElectronicID               OccupationID = "137"
	occupationEngineerEquipmentID                          OccupationID = "138"
	occupationEngineerFacilitiesID                         OccupationID = "139"
	occupationEngineerFlightID                             OccupationID = "140"
	occupationEngineerMechanicalID                         OccupationID = "141"
	occupationEngineerOperatingID                          OccupationID = "142"
	occupationEngineerPetroleumorMiningID                  OccupationID = "143"
	occupationEngineerSafetyID                             OccupationID = "144"
	occupationEngineerSalesID                              OccupationID = "145"
	occupationEntertainerPerformerID                       OccupationID = "146"
	occupationExpediterID                                  OccupationID = "147"
	occupationFactoryWorkerID                              OccupationID = "148"
	occupationFiremanWomanChiefCaptLtID                    OccupationID = "149"
	occupationFiremanWomanNonChiefID                       OccupationID = "150"
	occupationFishermanID                                  OccupationID = "151"
	occupationFlightAttendantID                            OccupationID = "152"
	occupationFloormenSupervisorID                         OccupationID = "153"
	occupationFloristID                                    OccupationID = "154"
	occupationForemanForewomanID                           OccupationID = "155"
	occupationForesterID                                   OccupationID = "156"
	occupationFundraiserID                                 OccupationID = "157"
	occupationGeographerID                                 OccupationID = "158"
	occupationGovtOfficialElectedID                        OccupationID = "159"
	occupationGraderID                                     OccupationID = "160"
	occupationGuardEmbassyID                               OccupationID = "161"
	occupationGuardSecurityorPrisonID                      OccupationID = "162"
	occupationGunsmithID                                   OccupationID = "163"
	occupationHairdresserBarberID                          OccupationID = "164"
	occupationHistorianID                                  OccupationID = "165"
	occupationHostorHostessRestaurantID                    OccupationID = "166"
	occupationHousekeeperorMaidID                          OccupationID = "167"
	occupationHumanResourcesRepresentativeID               OccupationID = "168"
	occupationIllustratorID                                OccupationID = "169"
	occupationInspectorAgriculturalID                      OccupationID = "170"
	occupationInspectorAircraftAccessoriesID               OccupationID = "171"
	occupationInspectorAllOtherID                          OccupationID = "172"
	occupationInspectorConstructionID                      OccupationID = "173"
	occupationInspectorPostalID                            OccupationID = "174"
	occupationInspectorWhiteCollarID                       OccupationID = "175"
	occupationInvestigatorPrivateID                        OccupationID = "176"
	occupationInvestmentBankerID                           OccupationID = "177"
	occupationInvestorPrivateID                            OccupationID = "178"
	occupationJournalistID                                 OccupationID = "179"
	occupationJourneymanID                                 OccupationID = "180"
	occupationJudgeID                                      OccupationID = "181"
	occupationLaborRelationsWorkerID                       OccupationID = "182"
	occupationLandscaperID                                 OccupationID = "183"
	occupationLibrarianID                                  OccupationID = "184"
	occupationLifeGuardID                                  OccupationID = "185"
	occupationLinguistID                                   OccupationID = "186"
	occupationLithographerID                               OccupationID = "187"
	occupationLobbyistID                                   OccupationID = "188"
	occupationLocksmithID                                  OccupationID = "189"
	occupationLongshoremenID                               OccupationID = "190"
	occupationMachinistID                                  OccupationID = "191"
	occupationManagerAirportID                             OccupationID = "192"
	occupationManagerAllOtherDegreedID                     OccupationID = "193"
	occupationManagerCafeteriaID                           OccupationID = "194"
	occupationManagerCityID                                OccupationID = "195"
	occupationManagerClericalStaffID                       OccupationID = "196"
	occupationManagerConvenienceorGasStationStoreID        OccupationID = "197"
	occupationManagerDepartmentStoreID                     OccupationID = "198"
	occupationManagerFinancialorCreditID                   OccupationID = "199"
	occupationManagerGeneralID                             OccupationID = "200"
	occupationManagerHealthClubID                          OccupationID = "201"
	occupationManagerHotelID                               OccupationID = "202"
	occupationManagerHumanResourcesID                      OccupationID = "203"
	occupationManagerMerchandiseID                         OccupationID = "204"
	occupationManagerOfficeID                              OccupationID = "205"
	occupationManagerOperationsID                          OccupationID = "206"
	occupationManagerProductionID                          OccupationID = "207"
	occupationManagerProfessionalTechStaffID               OccupationID = "208"
	occupationManagerProjectID                             OccupationID = "209"
	occupationManagerPropertyNonResidentID                 OccupationID = "210"
	occupationManagerPropertyResidentID                    OccupationID = "211"
	occupationManagerRestaurantFastFoodID                  OccupationID = "212"
	occupationManagerRestaurantNonFastFoodID               OccupationID = "213"
	occupationManagerSalesID                               OccupationID = "214"
	occupationManagerSecurityScreenerID                    OccupationID = "215"
	occupationManagerShippingReceivingID                   OccupationID = "216"
	occupationManagerStageID                               OccupationID = "217"
	occupationManagerSupermarketID                         OccupationID = "218"
	occupationManagerorOwnerSandwichShopID                 OccupationID = "219"
	occupationManicuristID                                 OccupationID = "220"
	occupationMarketingRepresentativeID                    OccupationID = "221"
	occupationMarshalFireID                                OccupationID = "222"
	occupationMarshalUSDeputyID                            OccupationID = "223"
	occupationMasseuseID                                   OccupationID = "224"
	occupationMathematicianID                              OccupationID = "225"
	occupationMeatcutterButcherID                          OccupationID = "226"
	occupationMechanicorServicemanAutoID                   OccupationID = "227"
	occupationMechanicorServicemanBoatID                   OccupationID = "228"
	occupationMechanicorServicemanDieselID                 OccupationID = "229"
	occupationMerchantID                                   OccupationID = "230"
	occupationMillwrightID                                 OccupationID = "231"
	occupationMorticianID                                  OccupationID = "232"
	occupationMusicianClassicalID                          OccupationID = "233"
	occupationMusicianOtherID                              OccupationID = "234"
	occupationNurseCNACertifiedNursingAssistantID          OccupationID = "235"
	occupationNurseLVNorLPNID                              OccupationID = "236"
	occupationNurseRNID                                    OccupationID = "237"
	occupationNursePractitionerID                          OccupationID = "238"
	occupationOceanographerID                              OccupationID = "239"
	occupationOfficerCorrectionalID                        OccupationID = "240"
	occupationOfficerCourtID                               OccupationID = "241"
	occupationOfficerForeignServiceID                      OccupationID = "242"
	occupationOfficerLoanID                                OccupationID = "243"
	occupationOfficerPoliceID                              OccupationID = "244"
	occupationOfficerPoliceChiefCaptainID                  OccupationID = "245"
	occupationOfficerPoliceDetectiveSgtLtID                OccupationID = "246"
	occupationOfficerProbationParoleID                     OccupationID = "247"
	occupationOfficerTelecommunicationsID                  OccupationID = "248"
	occupationOfficerWarrantID                             OccupationID = "249"
	occupationOfficerorManagerBankID                       OccupationID = "250"
	occupationOperatorAllOtherID                           OccupationID = "251"
	occupationOperatorBusinessID                           OccupationID = "252"
	occupationOperatorControlRoomID                        OccupationID = "253"
	occupationOperatorDataEntryID                          OccupationID = "254"
	occupationOperatorForkLiftID                           OccupationID = "255"
	occupationOperatorHeavyEquipmentID                     OccupationID = "256"
	occupationOperatorMachinePrecisionID                   OccupationID = "257"
	occupationOperatorNuclearReactorID                     OccupationID = "258"
	occupationOperatorTelephoneID                          OccupationID = "259"
	occupationOperatorWastewaterTreatmentPlantClassIVID    OccupationID = "260"
	occupationOpticianID                                   OccupationID = "261"
	occupationOptometristID                                OccupationID = "262"
	occupationOrthodontistID                               OccupationID = "263"
	occupationOwnerAllOtherID                              OccupationID = "264"
	occupationOwnerBarID                                   OccupationID = "265"
	occupationOwnerBeautyBarberShopID                      OccupationID = "266"
	occupationOwnerDealershipAutoDealerID                  OccupationID = "267"
	occupationOwnerorManagerFarmOrRanchID                  OccupationID = "268"
	occupationPainterID                                    OccupationID = "269"
	occupationParalegalID                                  OccupationID = "270"
	occupationParamedicorEMTID                             OccupationID = "271"
	occupationParkForestRangerID                           OccupationID = "272"
	occupationPathologistSpeechID                          OccupationID = "273"
	occupationPersonnelManagementSpecialistID              OccupationID = "274"
	occupationPestControlWorkerorExterminatorID            OccupationID = "275"
	occupationPharmacistID                                 OccupationID = "276"
	occupationPharmacologistID                             OccupationID = "277"
	occupationPhlebotomistID                               OccupationID = "278"
	occupationPhotographerID                               OccupationID = "279"
	occupationPhotographicProcessorID                      OccupationID = "280"
	occupationPhysicalTherapistAPTAMemberID                OccupationID = "281"
	occupationPhysicalTherapistNonAPTAMemberID             OccupationID = "282"
	occupationPhysicianorDoctorID                          OccupationID = "283"
	occupationPilotID                                      OccupationID = "284"
	occupationPilotCropBushID                              OccupationID = "285"
	occupationPipefitterOtherFitterID                      OccupationID = "286"
	occupationPlannerAllOtherID                            OccupationID = "287"
	occupationPlannerProductionorPrinterID                 OccupationID = "288"
	occupationPlumberID                                    OccupationID = "289"
	occupationPodiatristID                                 OccupationID = "290"
	occupationPoliticianID                                 OccupationID = "291"
	occupationPoolServiceCleanerID                         OccupationID = "292"
	occupationPostalExecutiveGradesPcesIIIID               OccupationID = "293"
	occupationPostmasterRuralID                            OccupationID = "294"
	occupationPostmasterUrbanSuburbanID                    OccupationID = "295"
	occupationPresidentBlueCollar50EmplID                  OccupationID = "296"
	occupationPresidentSkilledBlueCollarLessThan50EmpID    OccupationID = "297"
	occupationPresidentWhiteCollarID                       OccupationID = "298"
	occupationPrincipalorAssistantPrincipalID              OccupationID = "299"
	occupationPrinterID                                    OccupationID = "300"
	occupationProducerID                                   OccupationID = "301"
	occupationProfessorID                                  OccupationID = "302"
	occupationProgramManagementExpertID                    OccupationID = "303"
	occupationProofreaderID                                OccupationID = "304"
	occupationPsychiatristID                               OccupationID = "305"
	occupationPsychologistID                               OccupationID = "306"
	occupationPublicRelationsID                            OccupationID = "307"
	occupationPublisherID                                  OccupationID = "308"
	occupationQualityControlManufacturingID                OccupationID = "309"
	occupationQualityControlProfessionalID                 OccupationID = "310"
	occupationRadiologistID                                OccupationID = "311"
	occupationRanchHelperCowboyID                          OccupationID = "312"
	occupationRecruiterID                                  OccupationID = "313"
	occupationRegistrarID                                  OccupationID = "314"
	occupationReligiousClergyOrdainedorLicensedID          OccupationID = "315"
	occupationReligiousLaypersonNonClergyID                OccupationID = "316"
	occupationRepairServiceInstallACHeatingID              OccupationID = "317"
	occupationRepairServiceInstallAllOtherID               OccupationID = "318"
	occupationRepairServiceInstallJewelryWatchmakerID      OccupationID = "319"
	occupationRepairServiceInstallLineID                   OccupationID = "320"
	occupationRepairServiceInstallTrainedID                OccupationID = "321"
	occupationReporterID                                   OccupationID = "322"
	occupationResearcherAllOtherID                         OccupationID = "323"
	occupationRespiratoryTherapistID                       OccupationID = "324"
	occupationRoutemanRoutewomanID                         OccupationID = "325"
	occupationSalespersonAllOtherID                        OccupationID = "326"
	occupationSalespersonCarID                             OccupationID = "327"
	occupationSalespersonDoorToDoorID                      OccupationID = "328"
	occupationSalespersonHighTechID                        OccupationID = "329"
	occupationSalespersonNonHighTechID                     OccupationID = "330"
	occupationSalespersonPharmaceuticalID                  OccupationID = "331"
	occupationSalespersonRetailID                          OccupationID = "332"
	occupationSalespersonWholesaleID                       OccupationID = "333"
	occupationSanitarianID                                 OccupationID = "334"
	occupationSchedulerID                                  OccupationID = "335"
	occupationScientistAllOtherID                          OccupationID = "336"
	occupationSeamstressTailorID                           OccupationID = "337"
	occupationSecurityScreenerID                           OccupationID = "338"
	occupationShoeShinerRepairmanID                        OccupationID = "339"
	occupationSingerSongwriterID                           OccupationID = "340"
	occupationStaffingSpecialistID                         OccupationID = "341"
	occupationStateExaminerID                              OccupationID = "342"
	occupationSuperintendentAllOtherID                     OccupationID = "343"
	occupationSuperintendentDrillerID                      OccupationID = "344"
	occupationSuperintendentSchoolID                       OccupationID = "345"
	occupationSuperintendentorSupervisorBuildingMaintID    OccupationID = "346"
	occupationSupervisorAccountingID                       OccupationID = "347"
	occupationSupervisorAllOtherDegreedID                  OccupationID = "348"
	occupationSupervisorDataSystemsID                      OccupationID = "349"
	occupationSupervisorHumanResourcePersonnelID           OccupationID = "350"
	occupationSupervisorOfficeID                           OccupationID = "351"
	occupationSupervisorOperationsID                       OccupationID = "352"
	occupationSupervisorOtherDegreedID                     OccupationID = "353"
	occupationSupervisorPostalID                           OccupationID = "354"
	occupationSupervisorProductionID                       OccupationID = "355"
	occupationSupervisorRestaurantNonFastFoodID            OccupationID = "356"
	occupationSurgeonID                                    OccupationID = "357"
	occupationSurveyorLicensedID                           OccupationID = "358"
	occupationSurveyorNonLicensedID                        OccupationID = "359"
	occupationTaxExaminerNotClericalID                     OccupationID = "360"
	occupationTaxPreparerNotAccountantID                   OccupationID = "361"
	occupationTaxidermistID                                OccupationID = "362"
	occupationTeachersorCoachesorInstructorsID             OccupationID = "363"
	occupationTechnicianAllOtherID                         OccupationID = "364"
	occupationTechnicianElectricalorElectronicID           OccupationID = "365"
	occupationTechnicianFoodID                             OccupationID = "366"
	occupationTechnicianInstrumentationID                  OccupationID = "367"
	occupationTechnicianLabID                              OccupationID = "368"
	occupationTechnicianMedicalID                          OccupationID = "369"
	occupationTechnicianRadiologicalID                     OccupationID = "370"
	occupationTechnicianScienceID                          OccupationID = "371"
	occupationTechnicianTestingID                          OccupationID = "372"
	occupationTechnicianUltrasoundID                       OccupationID = "373"
	occupationTechnicianXRayID                             OccupationID = "374"
	occupationTechnicianorAssistantEngineeringID           OccupationID = "375"
	occupationTelemarketerID                               OccupationID = "376"
	occupationTherapistID                                  OccupationID = "377"
	occupationTrainerAerobicsFitnessID                     OccupationID = "378"
	occupationTrainerAthleticNataMemberID                  OccupationID = "379"
	occupationTrainerAthleticNonNataMemberID               OccupationID = "380"
	occupationTrainerCaretakerAnimalID                     OccupationID = "381"
	occupationTranslatororInterpreterID                    OccupationID = "382"
	occupationTreasurerID                                  OccupationID = "383"
	occupationTutorID                                      OccupationID = "384"
	occupationUnderwriterInsuranceID                       OccupationID = "385"
	occupationVendorID                                     OccupationID = "386"
	occupationVeterinarianID                               OccupationID = "387"
	occupationVicePresBusinessID                           OccupationID = "388"
	occupationWaiterWaitressID                             OccupationID = "389"
	occupationWardenAllOtherID                             OccupationID = "390"
	occupationWardenGameID                                 OccupationID = "391"
	occupationWorkerMetalNotSteelID                        OccupationID = "392"
	occupationWorkerRailroadID                             OccupationID = "393"
	occupationWorkerSocialCaseID                           OccupationID = "394"
	occupationWriterAllOtherID                             OccupationID = "395"
	occupationWriterCommercialID                           OccupationID = "396"
	occupationE1ID                                         OccupationID = "397"
	occupationE2ID                                         OccupationID = "398"
	occupationE3ID                                         OccupationID = "399"
	occupationE4ID                                         OccupationID = "400"
	occupationE4PID                                        OccupationID = "401"
	occupationE5ID                                         OccupationID = "402"
	occupationE5PID                                        OccupationID = "403"
	occupationE6ID                                         OccupationID = "404"
	occupationE6PID                                        OccupationID = "405"
	occupationE7ID                                         OccupationID = "406"
	occupationE8ID                                         OccupationID = "407"
	occupationE9ID                                         OccupationID = "408"
	occupationO1ID                                         OccupationID = "409"
	occupationO10ID                                        OccupationID = "410"
	occupationO2ID                                         OccupationID = "411"
	occupationO3ID                                         OccupationID = "412"
	occupationO4ID                                         OccupationID = "413"
	occupationO5ID                                         OccupationID = "414"
	occupationO6ID                                         OccupationID = "415"
	occupationO7ID                                         OccupationID = "416"
	occupationO8ID                                         OccupationID = "417"
	occupationO9ID                                         OccupationID = "418"
	occupationW1ID                                         OccupationID = "419"
	occupationW2ID                                         OccupationID = "420"
	occupationW3ID                                         OccupationID = "421"
	occupationW4ID                                         OccupationID = "422"
	occupationW5ID                                         OccupationID = "423"
	occupationHomemakerID                                  OccupationID = "424"
	occupationUnemployedID                                 OccupationID = "425"
	occupationGraduateSchoolID                             OccupationID = "426"
	occupationHighSchoolID                                 OccupationID = "427"
	occupationTradeSchoolorAssociateDegreeID               OccupationID = "428"
	occupationUndergraduate4yeardegreeID                   OccupationID = "429"
	occupationDisabledID                                   OccupationID = "430"
)

// EnumOccupationItem describes an entry in an enumeration of Occupation
type EnumOccupationItem struct {
	ID        OccupationID      `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	Keywords string
}

var (
	occupationAccountExecutive                           = EnumOccupationItem{occupationAccountExecutiveID, "Account Executive", map[string]string{"Keywords": "AAAAA,Sales,Account,Representative"}, "AccountExecutive", 1, "AAAAA,Sales,Account,Representative"}
	occupationAccountantorCPA                            = EnumOccupationItem{occupationAccountantorCPAID, "Accountant or CPA", map[string]string{"Keywords": "Certified,CPA,Accountant,Public"}, "AccountantorCPA", 2, "Certified,CPA,Accountant,Public"}
	occupationActorActress                               = EnumOccupationItem{occupationActorActressID, "Actor / Actress", map[string]string{"Keywords": "Extra,Thespian,Pantomimist,Actor,Comedian,Impersonator,Actress,Ventriloquist,Mime"}, "ActorActress", 3, "Extra,Thespian,Pantomimist,Actor,Comedian,Impersonator,Actress,Ventriloquist,Mime"}
	occupationActuary                                    = EnumOccupationItem{occupationActuaryID, "Actuary", map[string]string{"Keywords": "Actuary,Statistical,Risk Management"}, "Actuary", 4, "Actuary,Statistical,Risk Management"}
	occupationAcupuncturist                              = EnumOccupationItem{occupationAcupuncturistID, "Acupuncturist", map[string]string{"Keywords": "Acupuncturist,Therapy,Stress"}, "Acupuncturist", 5, "Acupuncturist,Therapy,Stress"}
	occupationAdjudicator                                = EnumOccupationItem{occupationAdjudicatorID, "Adjudicator", map[string]string{"Keywords": "Arbitrator,Adjudicator"}, "Adjudicator", 6, "Arbitrator,Adjudicator"}
	occupationAdministrativeAssistant                    = EnumOccupationItem{occupationAdministrativeAssistantID, "Administrative Assistant", map[string]string{"Keywords": "Secretary,Administrative Assistant,Office"}, "AdministrativeAssistant", 7, "Secretary,Administrative Assistant,Office"}
	occupationAdministrativeClerkorReceptionist          = EnumOccupationItem{occupationAdministrativeClerkorReceptionistID, "Administrative Clerk or Receptionist", map[string]string{"Keywords": "Receptionist,Administrative,Clerk"}, "AdministrativeClerkorReceptionist", 8, "Receptionist,Administrative,Clerk"}
	occupationAdministrator                              = EnumOccupationItem{occupationAdministratorID, "Administrator", map[string]string{"Keywords": "Administrator,Chairperson,Board"}, "Administrator", 9, "Administrator,Chairperson,Board"}
	occupationAdministratorHealthorHospital              = EnumOccupationItem{occupationAdministratorHealthorHospitalID, "Administrator - Health or Hospital", map[string]string{"Keywords": "Administrator,Hospital,Health"}, "AdministratorHealthorHospital", 10, "Administrator,Hospital,Health"}
	occupationAdministratorPublic                        = EnumOccupationItem{occupationAdministratorPublicID, "Administrator - Public", map[string]string{"Keywords": "Administrator,Public"}, "AdministratorPublic", 11, "Administrator,Public"}
	occupationAdvisorFinancialorCredit                   = EnumOccupationItem{occupationAdvisorFinancialorCreditID, "Advisor - Financial or Credit", map[string]string{"Keywords": "Financial,Credit,Advisor,Personal"}, "AdvisorFinancialorCredit", 12, "Financial,Credit,Advisor,Personal"}
	occupationAgentAdvertising                           = EnumOccupationItem{occupationAgentAdvertisingID, "Agent - Advertising", map[string]string{"Keywords": "Media,Agent,Advertising"}, "AgentAdvertising", 13, "Media,Agent,Advertising"}
	occupationAgentBorderPatrol                          = EnumOccupationItem{occupationAgentBorderPatrolID, "Agent - Border Patrol", map[string]string{"Keywords": "Agent,Border patrol"}, "AgentBorderPatrol", 14, "Agent,Border patrol"}
	occupationAgentEmployment                            = EnumOccupationItem{occupationAgentEmploymentID, "Agent - Employment", map[string]string{"Keywords": "Temporary,Hire,Agent,Employment"}, "AgentEmployment", 15, "Temporary,Hire,Agent,Employment"}
	occupationAgentFBI                                   = EnumOccupationItem{occupationAgentFBIID, "Agent - FBI", map[string]string{"Keywords": "DEA,Alcohol,FBI,Federal Bureau of Investigation,Firearms,ATF,Agent,Drug Enforcement Agency,Department of Homeland Security,ICE,Immigration,DHS,Tobbacco"}, "AgentFBI", 16, "DEA,Alcohol,FBI,Federal Bureau of Investigation,Firearms,ATF,Agent,Drug Enforcement Agency,Department of Homeland Security,ICE,Immigration,DHS,Tobbacco"}
	occupationAgentImportExport                          = EnumOccupationItem{occupationAgentImportExportID, "Agent - Import / Export", map[string]string{"Keywords": "Import,Agent,Customs,Export"}, "AgentImportExport", 17, "Import,Agent,Customs,Export"}
	occupationAgentInsurance                             = EnumOccupationItem{occupationAgentInsuranceID, "Agent - Insurance", map[string]string{"Keywords": "Annuities,Business,Agent,Workers compensation,Automobile,Health,Life,Insurance,Homeowners"}, "AgentInsurance", 18, "Annuities,Business,Agent,Workers compensation,Automobile,Health,Life,Insurance,Homeowners"}
	occupationAgentIRS                                   = EnumOccupationItem{occupationAgentIRSID, "Agent - IRS", map[string]string{"Keywords": "IRS,Agent,Internal Revenue Service"}, "AgentIRS", 19, "IRS,Agent,Internal Revenue Service"}
	occupationAgentLeasing                               = EnumOccupationItem{occupationAgentLeasingID, "Agent - Leasing", map[string]string{"Keywords": "Leasing,Agent,Rental"}, "AgentLeasing", 20, "Leasing,Agent,Rental"}
	occupationAgentRealEstate                            = EnumOccupationItem{occupationAgentRealEstateID, "Agent - Real Estate", map[string]string{"Keywords": "Real estate,Agent"}, "AgentRealEstate", 21, "Real estate,Agent"}
	occupationAgentSecretService                         = EnumOccupationItem{occupationAgentSecretServiceID, "Agent - Secret Service", map[string]string{"Keywords": "Agent,Secret Service"}, "AgentSecretService", 22, "Agent,Secret Service"}
	occupationAgentTheatrical                            = EnumOccupationItem{occupationAgentTheatricalID, "Agent - Theatrical", map[string]string{"Keywords": "Model,Agent,Talent,Theatrical"}, "AgentTheatrical", 23, "Model,Agent,Talent,Theatrical"}
	occupationAgentTravelTicket                          = EnumOccupationItem{occupationAgentTravelTicketID, "Agent - Travel / Ticket", map[string]string{"Keywords": "Travel,Ticket,Airline,Cruise,Agent,Vacation"}, "AgentTravelTicket", 24, "Travel,Ticket,Airline,Cruise,Agent,Vacation"}
	occupationAgentTreasury                              = EnumOccupationItem{occupationAgentTreasuryID, "Agent - Treasury", map[string]string{"Keywords": "Treasury,Agent"}, "AgentTreasury", 25, "Treasury,Agent"}
	occupationAideAllOther                               = EnumOccupationItem{occupationAideAllOtherID, "Aide - All Other", map[string]string{"Keywords": "Teacher,Home,Nurse,Other,Health,Aide"}, "AideAllOther", 26, "Teacher,Home,Nurse,Other,Health,Aide"}
	occupationAirTrafficController                       = EnumOccupationItem{occupationAirTrafficControllerID, "Air Traffic Controller", map[string]string{"Keywords": "Air Traffic Controller"}, "AirTrafficController", 27, "Air Traffic Controller"}
	occupationAmbassador                                 = EnumOccupationItem{occupationAmbassadorID, "Ambassador", map[string]string{"Keywords": "Emissary,Diplomat,Ambassador"}, "Ambassador", 28, "Emissary,Diplomat,Ambassador"}
	occupationAnalystAllOther                            = EnumOccupationItem{occupationAnalystAllOtherID, "Analyst - All Other", map[string]string{"Keywords": "Intelligence,Business,Analyst,Other"}, "AnalystAllOther", 29, "Intelligence,Business,Analyst,Other"}
	occupationAnalystEngineering                         = EnumOccupationItem{occupationAnalystEngineeringID, "Analyst - Engineering", map[string]string{"Keywords": "Engineering,Analyst,Technical"}, "AnalystEngineering", 30, "Engineering,Analyst,Technical"}
	occupationAnalystFinancial                           = EnumOccupationItem{occupationAnalystFinancialID, "Analyst - Financial", map[string]string{"Keywords": "Financial,Analyst,Credit,Invest"}, "AnalystFinancial", 31, "Financial,Analyst,Credit,Invest"}
	occupationAnesthesiologist                           = EnumOccupationItem{occupationAnesthesiologistID, "Anesthesiologist", map[string]string{"Keywords": "Anesthesiologist"}, "Anesthesiologist", 32, "Anesthesiologist"}
	occupationAntiqueDealer                              = EnumOccupationItem{occupationAntiqueDealerID, "Antique Dealer", map[string]string{"Keywords": "Sales,Dealer,Antique"}, "AntiqueDealer", 33, "Sales,Dealer,Antique"}
	occupationAppraiserAllOther                          = EnumOccupationItem{occupationAppraiserAllOtherID, "Appraiser - All Other", map[string]string{"Keywords": "Cost,Arbitrator,Value,Other,Appraiser"}, "AppraiserAllOther", 34, "Cost,Arbitrator,Value,Other,Appraiser"}
	occupationAppraiserArt                               = EnumOccupationItem{occupationAppraiserArtID, "Appraiser - Art", map[string]string{"Keywords": "Art,Arbitrator,Value,Appraiser"}, "AppraiserArt", 35, "Art,Arbitrator,Value,Appraiser"}
	occupationAppraiserRealEstate                        = EnumOccupationItem{occupationAppraiserRealEstateID, "Appraiser - Real Estate", map[string]string{"Keywords": "Value,Arbitrator,Home,Appraiser,House,Property,Real Estate"}, "AppraiserRealEstate", 36, "Value,Arbitrator,Home,Appraiser,House,Property,Real Estate"}
	occupationApprenticeAllOther                         = EnumOccupationItem{occupationApprenticeAllOtherID, "Apprentice - All Other", map[string]string{"Keywords": "Apprentice,Electrical,Plumber,Pipe,Other,Skilled,Mason"}, "ApprenticeAllOther", 37, "Apprentice,Electrical,Plumber,Pipe,Other,Skilled,Mason"}
	occupationArchitect                                  = EnumOccupationItem{occupationArchitectID, "Architect", map[string]string{"Keywords": "Draft,Architect,Building,Design"}, "Architect", 38, "Draft,Architect,Building,Design"}
	occupationArchivistLibraryorMuseum                   = EnumOccupationItem{occupationArchivistLibraryorMuseumID, "Archivist Library or Museum", map[string]string{"Keywords": "Library,Museum,Curator,Archivist"}, "ArchivistLibraryorMuseum", 39, "Library,Museum,Curator,Archivist"}
	occupationArtistCommercial                           = EnumOccupationItem{occupationArtistCommercialID, "Artist - Commercial", map[string]string{"Keywords": "Commercial,Craft,Artist"}, "ArtistCommercial", 40, "Commercial,Craft,Artist"}
	occupationArtistNoncommercial                        = EnumOccupationItem{occupationArtistNoncommercialID, "Artist - Noncommercial", map[string]string{"Keywords": "Craft,Noncommercial,Artist"}, "ArtistNoncommercial", 41, "Craft,Noncommercial,Artist"}
	occupationAssistantAllOther                          = EnumOccupationItem{occupationAssistantAllOtherID, "Assistant - All Other", map[string]string{"Keywords": "Assistant,Other"}, "AssistantAllOther", 42, "Assistant,Other"}
	occupationAssistantEditorial                         = EnumOccupationItem{occupationAssistantEditorialID, "Assistant - Editorial", map[string]string{"Keywords": "Editorial,Assistant"}, "AssistantEditorial", 43, "Editorial,Assistant"}
	occupationAssistantLegislativeorLegal                = EnumOccupationItem{occupationAssistantLegislativeorLegalID, "Assistant - Legislative or Legal", map[string]string{"Keywords": "Legal,Assistant,Legislative"}, "AssistantLegislativeorLegal", 44, "Legal,Assistant,Legislative"}
	occupationAssistantMedical                           = EnumOccupationItem{occupationAssistantMedicalID, "Assistant - Medical", map[string]string{"Keywords": "Medical,Assistant,Dental"}, "AssistantMedical", 45, "Medical,Assistant,Dental"}
	occupationAssistantPrintingPress                     = EnumOccupationItem{occupationAssistantPrintingPressID, "Assistant - Printing Press", map[string]string{"Keywords": "Assistant,Jogger,Printing press"}, "AssistantPrintingPress", 46, "Assistant,Jogger,Printing press"}
	occupationAthlete                                    = EnumOccupationItem{occupationAthleteID, "Athlete", map[string]string{"Keywords": "Ball,Base,Athlete,Tennis,Ski,Gymnast,Foot,Soccer,Basket,Volley,Rugby,Hockey"}, "Athlete", 47, "Ball,Base,Athlete,Tennis,Ski,Gymnast,Foot,Soccer,Basket,Volley,Rugby,Hockey"}
	occupationAttorneyorLawyer                           = EnumOccupationItem{occupationAttorneyorLawyerID, "Attorney or Lawyer", map[string]string{"Keywords": "Attorney,District,Legal,Lawyer"}, "AttorneyorLawyer", 48, "Attorney,District,Legal,Lawyer"}
	occupationAuctioneer                                 = EnumOccupationItem{occupationAuctioneerID, "Auctioneer", map[string]string{"Keywords": "Auctioneer"}, "Auctioneer", 49, "Auctioneer"}
	occupationAudiologist                                = EnumOccupationItem{occupationAudiologistID, "Audiologist", map[string]string{"Keywords": "Vestibular,Audiologist,Ear Doctor"}, "Audiologist", 50, "Vestibular,Audiologist,Ear Doctor"}
	occupationAuditorFinancial                           = EnumOccupationItem{occupationAuditorFinancialID, "Auditor - Financial", map[string]string{"Keywords": "Financial,Auditor,Bank,Account"}, "AuditorFinancial", 51, "Financial,Auditor,Bank,Account"}
	occupationAuditorNonFinancial                        = EnumOccupationItem{occupationAuditorNonFinancialID, "Auditor - Non Financial", map[string]string{"Keywords": "Auditor,Business"}, "AuditorNonFinancial", 52, "Auditor,Business"}
	occupationBaker                                      = EnumOccupationItem{occupationBakerID, "Baker", map[string]string{"Keywords": "Chef,Desert,Cake,Cook,Baker"}, "Baker", 53, "Chef,Desert,Cake,Cook,Baker"}
	occupationBankTeller                                 = EnumOccupationItem{occupationBankTellerID, "Bank Teller", map[string]string{"Keywords": "Bank Teller,Clerk"}, "BankTeller", 54, "Bank Teller,Clerk"}
	occupationBartender                                  = EnumOccupationItem{occupationBartenderID, "Bartender", map[string]string{"Keywords": "Bartender,Mixologist,Server"}, "Bartender", 55, "Bartender,Mixologist,Server"}
	occupationBoatCaptain                                = EnumOccupationItem{occupationBoatCaptainID, "Boat Captain", map[string]string{"Keywords": "Ship,Sea,Sail,Boat Captain"}, "BoatCaptain", 56, "Ship,Sea,Sail,Boat Captain"}
	occupationBoatman                                    = EnumOccupationItem{occupationBoatmanID, "Boatman", map[string]string{"Keywords": "Boatman,ship,sea,dock"}, "Boatman", 57, "Boatman,ship,sea,dock"}
	occupationBondsman                                   = EnumOccupationItem{occupationBondsmanID, "Bondsman", map[string]string{"Keywords": "Loan,Bail,Bounty Hunter,Bondsman"}, "Bondsman", 58, "Loan,Bail,Bounty Hunter,Bondsman"}
	occupationBookbinder                                 = EnumOccupationItem{occupationBookbinderID, "Bookbinder", map[string]string{"Keywords": "Bookbinder,Print"}, "Bookbinder", 59, "Bookbinder,Print"}
	occupationBookkeeper                                 = EnumOccupationItem{occupationBookkeeperID, "Bookkeeper", map[string]string{"Keywords": "Data,Folder,Records,Bookkeeper,Files"}, "Bookkeeper", 60, "Data,Folder,Records,Bookkeeper,Files"}
	occupationBroadcaster                                = EnumOccupationItem{occupationBroadcasterID, "Broadcaster", map[string]string{"Keywords": "Broadcaster,Anchor,Television,Radio,Mass Media"}, "Broadcaster", 61, "Broadcaster,Anchor,Television,Radio,Mass Media"}
	occupationBrokerInsurance                            = EnumOccupationItem{occupationBrokerInsuranceID, "Broker - Insurance", map[string]string{"Keywords": "Home,Broker,Auto,Health,Life,Insurance,Annuity"}, "BrokerInsurance", 62, "Home,Broker,Auto,Health,Life,Insurance,Annuity"}
	occupationBrokerMortgage                             = EnumOccupationItem{occupationBrokerMortgageID, "Broker - Mortgage", map[string]string{"Keywords": "Mortgage,Broker,Bank"}, "BrokerMortgage", 63, "Mortgage,Broker,Bank"}
	occupationBrokerPawn                                 = EnumOccupationItem{occupationBrokerPawnID, "Broker - Pawn", map[string]string{"Keywords": "Loan,Broker,Pawn"}, "BrokerPawn", 64, "Loan,Broker,Pawn"}
	occupationBrokerRealEstate                           = EnumOccupationItem{occupationBrokerRealEstateID, "Broker - Real Estate", map[string]string{"Keywords": "Broker,Home sales,Real Estate"}, "BrokerRealEstate", 65, "Broker,Home sales,Real Estate"}
	occupationBrokerStockBonds                           = EnumOccupationItem{occupationBrokerStockBondsID, "Broker - Stock & Bonds", map[string]string{"Keywords": "Bonds,Series 7,Broker,Stock"}, "BrokerStockBonds", 66, "Bonds,Series 7,Broker,Stock"}
	occupationBuyerPurchasingAgent                       = EnumOccupationItem{occupationBuyerPurchasingAgentID, "Buyer - Purchasing Agent", map[string]string{"Keywords": "Media,Buyer,Purchasing Agent,Equiptment"}, "BuyerPurchasingAgent", 67, "Media,Buyer,Purchasing Agent,Equiptment"}
	occupationCameraman                                  = EnumOccupationItem{occupationCameramanID, "Cameraman", map[string]string{"Keywords": "Cameraman,Cinematography,Director"}, "Cameraman", 68, "Cameraman,Cinematography,Director"}
	occupationCarpenterorCabinetmaker                    = EnumOccupationItem{occupationCarpenterorCabinetmakerID, "Carpenter or Cabinetmaker", map[string]string{"Keywords": "Cabinetmaker,Furniture,Frame,Carpenter,Dry wall"}, "CarpenterorCabinetmaker", 69, "Cabinetmaker,Furniture,Frame,Carpenter,Dry wall"}
	occupationCashierFood                                = EnumOccupationItem{occupationCashierFoodID, "Cashier - Food", map[string]string{"Keywords": "Food,Attendant,Gas station,Cashier"}, "CashierFood", 70, "Food,Attendant,Gas station,Cashier"}
	occupationCashierOffice                              = EnumOccupationItem{occupationCashierOfficeID, "Cashier - Office", map[string]string{"Keywords": "Cashier,Office"}, "CashierOffice", 71, "Cashier,Office"}
	occupationCashierRetail                              = EnumOccupationItem{occupationCashierRetailID, "Cashier - Retail", map[string]string{"Keywords": "Grocery,Gas station,Cashier,Retail,Department"}, "CashierRetail", 72, "Grocery,Gas station,Cashier,Retail,Department"}
	occupationCaterer                                    = EnumOccupationItem{occupationCatererID, "Caterer", map[string]string{"Keywords": "Food service,Caterer,Server"}, "Caterer", 73, "Food service,Caterer,Server"}
	occupationCensusTaker                                = EnumOccupationItem{occupationCensusTakerID, "Census Taker", map[string]string{"Keywords": "Survey,Census Taker"}, "CensusTaker", 74, "Survey,Census Taker"}
	occupationChiropractor                               = EnumOccupationItem{occupationChiropractorID, "Chiropractor", map[string]string{"Keywords": "Injury,Spine,Ortho,DC,Chiropractor,Back"}, "Chiropractor", 75, "Injury,Spine,Ortho,DC,Chiropractor,Back"}
	occupationClaimsAdjuster                             = EnumOccupationItem{occupationClaimsAdjusterID, "Claims Adjuster", map[string]string{"Keywords": "Calculate,Evaluate,Claims Adjuster"}, "ClaimsAdjuster", 76, "Calculate,Evaluate,Claims Adjuster"}
	occupationClaimsExaminer                             = EnumOccupationItem{occupationClaimsExaminerID, "Claims Examiner", map[string]string{"Keywords": "Calculate,Evaluate,Claims Examiner,Recovery specialist"}, "ClaimsExaminer", 77, "Calculate,Evaluate,Claims Examiner,Recovery specialist"}
	occupationClerkAccountingorFinancial                 = EnumOccupationItem{occupationClerkAccountingorFinancialID, "Clerk - Accounting or Financial", map[string]string{"Keywords": "Financial,Clerk,Accounting"}, "ClerkAccountingorFinancial", 78, "Financial,Clerk,Accounting"}
	occupationClerkAllOther                              = EnumOccupationItem{occupationClerkAllOtherID, "Clerk - All Other", map[string]string{"Keywords": "Receiving,Other,Clerk,Shipping"}, "ClerkAllOther", 79, "Receiving,Other,Clerk,Shipping"}
	occupationCollegeDean                                = EnumOccupationItem{occupationCollegeDeanID, "College Dean", map[string]string{"Keywords": "College Dean,Advisor"}, "CollegeDean", 80, "College Dean,Advisor"}
	occupationCommunicationSpecialist                    = EnumOccupationItem{occupationCommunicationSpecialistID, "Communication Specialist", map[string]string{"Keywords": "Media,Consumer,Communication Specialist,Language,Interpreter,Translator"}, "CommunicationSpecialist", 81, "Media,Consumer,Communication Specialist,Language,Interpreter,Translator"}
	occupationComputerProgrammer                         = EnumOccupationItem{occupationComputerProgrammerID, "Computer Programmer", map[string]string{"Keywords": "Windows,Web,DBA,Database,Computer,Programmer,Java,Internet,Developer"}, "ComputerProgrammer", 82, "Windows,Web,DBA,Database,Computer,Programmer,Java,Internet,Developer"}
	occupationComputerTechnicalSupportRep                = EnumOccupationItem{occupationComputerTechnicalSupportRepID, "Computer Technical Support Rep", map[string]string{"Keywords": "Phone,Call center,Computer,Representative,Technical Support"}, "ComputerTechnicalSupportRep", 83, "Phone,Call center,Computer,Representative,Technical Support"}
	occupationConductor                                  = EnumOccupationItem{occupationConductorID, "Conductor", map[string]string{"Keywords": "Frieght,Train,Conductor,Railroad Engineer"}, "Conductor", 84, "Frieght,Train,Conductor,Railroad Engineer"}
	occupationConservationist                            = EnumOccupationItem{occupationConservationistID, "Conservationist", map[string]string{"Keywords": "Environmental Activist,Preservation,Conservationist"}, "Conservationist", 85, "Environmental Activist,Preservation,Conservationist"}
	occupationConstructionWorker                         = EnumOccupationItem{occupationConstructionWorkerID, "Construction Worker", map[string]string{"Keywords": "Brick layer,Dock,Mortar,Construction Worker,Pour,Frame,Roofing,Concrete,Mason,Dry wall"}, "ConstructionWorker", 86, "Brick layer,Dock,Mortar,Construction Worker,Pour,Frame,Roofing,Concrete,Mason,Dry wall"}
	occupationConsultantorAdvisor                        = EnumOccupationItem{occupationConsultantorAdvisorID, "Consultant or Advisor", map[string]string{"Keywords": "Technical,Consultant,Advisor,Specialist"}, "ConsultantorAdvisor", 87, "Technical,Consultant,Advisor,Specialist"}
	occupationContractororDeveloperAllOther              = EnumOccupationItem{occupationContractororDeveloperAllOtherID, "Contractor or Developer - All Other", map[string]string{"Keywords": "Contractor,Other,Developer"}, "ContractororDeveloperAllOther", 88, "Contractor,Other,Developer"}
	occupationContractororDeveloperArtisanSkilled        = EnumOccupationItem{occupationContractororDeveloperArtisanSkilledID, "Contractor or Developer - Artisan / Skilled", map[string]string{"Keywords": "Contractor,Skilled,Artisan,Developer"}, "ContractororDeveloperArtisanSkilled", 89, "Contractor,Skilled,Artisan,Developer"}
	occupationContractororDeveloperBlueCollar50Employees = EnumOccupationItem{occupationContractororDeveloperBlueCollar50EmployeesID, "Contractor or Developer - Blue Collar 50+ Employees", map[string]string{"Keywords": "Blue Collar,Contractor,Employees,Fifty,Developer"}, "ContractororDeveloperBlueCollar50Employees", 90, "Blue Collar,Contractor,Employees,Fifty,Developer"}
	occupationContractororDeveloperWhiteCollar           = EnumOccupationItem{occupationContractororDeveloperWhiteCollarID, "Contractor or Developer - White Collar", map[string]string{"Keywords": "White Collar,Contractor,Technical,Developer"}, "ContractororDeveloperWhiteCollar", 91, "White Collar,Contractor,Technical,Developer"}
	occupationControllerorComptroller                    = EnumOccupationItem{occupationControllerorComptrollerID, "Controller or Comptroller", map[string]string{"Keywords": "Controller,Finance officer,Comptroller"}, "ControllerorComptroller", 92, "Controller,Finance officer,Comptroller"}
	occupationControllerorCoordinatorProduction          = EnumOccupationItem{occupationControllerorCoordinatorProductionID, "Controller or Coordinator - Production", map[string]string{"Keywords": "Controller,Production,Coordinator"}, "ControllerorCoordinatorProduction", 93, "Controller,Production,Coordinator"}
	occupationCookorChef                                 = EnumOccupationItem{occupationCookorChefID, "Cook or Chef", map[string]string{"Keywords": "Gourmet,Chef,Food prep,Cook,Culinary"}, "CookorChef", 94, "Gourmet,Chef,Food prep,Cook,Culinary"}
	occupationCoordinatorOffice                          = EnumOccupationItem{occupationCoordinatorOfficeID, "Coordinator - Office", map[string]string{"Keywords": "Liaison,Marketing,Coordinator,Office"}, "CoordinatorOffice", 95, "Liaison,Marketing,Coordinator,Office"}
	occupationCopywriter                                 = EnumOccupationItem{occupationCopywriterID, "Copywriter", map[string]string{"Keywords": "Copywriter"}, "Copywriter", 96, "Copywriter"}
	occupationCoroner                                    = EnumOccupationItem{occupationCoronerID, "Coroner", map[string]string{"Keywords": "Autopsy,Coroner"}, "Coroner", 97, "Autopsy,Coroner"}
	occupationCosmetologistBeautician                    = EnumOccupationItem{occupationCosmetologistBeauticianID, "Cosmetologist / Beautician", map[string]string{"Keywords": "Beautician,Cosmetologist,Make up artist"}, "CosmetologistBeautician", 98, "Beautician,Cosmetologist,Make up artist"}
	occupationCounselorAllOther                          = EnumOccupationItem{occupationCounselorAllOtherID, "Counselor - All Other", map[string]string{"Keywords": "Counselor,Alcohol,Drug,Other,Substance abuse"}, "CounselorAllOther", 99, "Counselor,Alcohol,Drug,Other,Substance abuse"}
	occupationCounselorEducation                         = EnumOccupationItem{occupationCounselorEducationID, "Counselor - Education", map[string]string{"Keywords": "Counselor,Education,Advisor"}, "CounselorEducation", 100, "Counselor,Education,Advisor"}
	occupationCounselorFamilyAndChild                    = EnumOccupationItem{occupationCounselorFamilyAndChildID, "Counselor - Family And Child", map[string]string{"Keywords": "Counselor,Family,Child,Marriage"}, "CounselorFamilyAndChild", 101, "Counselor,Family,Child,Marriage"}
	occupationCounselorMentalHealth                      = EnumOccupationItem{occupationCounselorMentalHealthID, "Counselor - Mental Health", map[string]string{"Keywords": "Counselor,Mental Health"}, "CounselorMentalHealth", 102, "Counselor,Mental Health"}
	occupationCourtReporter                              = EnumOccupationItem{occupationCourtReporterID, "Court Reporter", map[string]string{"Keywords": "Paper,Court Reporter,News"}, "CourtReporter", 103, "Paper,Court Reporter,News"}
	occupationCraftsman                                  = EnumOccupationItem{occupationCraftsmanID, "Craftsman", map[string]string{"Keywords": "Designer,Furniture maker,Craftsman,Woodworker"}, "Craftsman", 104, "Designer,Furniture maker,Craftsman,Woodworker"}
	occupationCriminologist                              = EnumOccupationItem{occupationCriminologistID, "Criminologist", map[string]string{"Keywords": "Criminologist,Profiler"}, "Criminologist", 105, "Criminologist,Profiler"}
	occupationCurator                                    = EnumOccupationItem{occupationCuratorID, "Curator", map[string]string{"Keywords": "Museum,Curator"}, "Curator", 106, "Museum,Curator"}
	occupationCustodianorJanitor                         = EnumOccupationItem{occupationCustodianorJanitorID, "Custodian or Janitor", map[string]string{"Keywords": "Custodian,Janitor,Clean,Maintenance,Groundskeeper"}, "CustodianorJanitor", 107, "Custodian,Janitor,Clean,Maintenance,Groundskeeper"}
	occupationCustomerServiceRepresentative              = EnumOccupationItem{occupationCustomerServiceRepresentativeID, "Customer Service Representative", map[string]string{"Keywords": "Customer,Service,Representative,Customer care"}, "CustomerServiceRepresentative", 108, "Customer,Service,Representative,Customer care"}
	occupationDayCareWorker                              = EnumOccupationItem{occupationDayCareWorkerID, "Day Care Worker", map[string]string{"Keywords": "Nanny,Babysitter,Day Care Worker"}, "DayCareWorker", 109, "Nanny,Babysitter,Day Care Worker"}
	occupationDecoratorInterior                          = EnumOccupationItem{occupationDecoratorInteriorID, "Decorator - Interior", map[string]string{"Keywords": "Interior,Designer,Decorator"}, "DecoratorInterior", 110, "Interior,Designer,Decorator"}
	occupationDeliveryPersonorMailCarrier                = EnumOccupationItem{occupationDeliveryPersonorMailCarrierID, "Delivery Person or Mail Carrier", map[string]string{"Keywords": "Mail Carrier,USPS,Delivery Person,Fed ex,UPS"}, "DeliveryPersonorMailCarrier", 111, "Mail Carrier,USPS,Delivery Person,Fed ex,UPS"}
	occupationDentalHygienist                            = EnumOccupationItem{occupationDentalHygienistID, "Dental Hygienist", map[string]string{"Keywords": "Dental Hygienist"}, "DentalHygienist", 112, "Dental Hygienist"}
	occupationDentist                                    = EnumOccupationItem{occupationDentistID, "Dentist", map[string]string{"Keywords": "Dentist"}, "Dentist", 113, "Dentist"}
	occupationDesignerComputerWebsite                    = EnumOccupationItem{occupationDesignerComputerWebsiteID, "Designer - Computer Website", map[string]string{"Keywords": "Computer Website,Designer,Developer,Internet"}, "DesignerComputerWebsite", 114, "Computer Website,Designer,Developer,Internet"}
	occupationDesignerFloral                             = EnumOccupationItem{occupationDesignerFloralID, "Designer - Floral", map[string]string{"Keywords": "Designer - Floral,Boquet,Flowers"}, "DesignerFloral", 115, "Designer - Floral,Boquet,Flowers"}
	occupationDesignerGraphicorTechnical                 = EnumOccupationItem{occupationDesignerGraphicorTechnicalID, "Designer - Graphic or Technical", map[string]string{"Keywords": "Designer,Graphic,Web,Technical,Computer,Internet"}, "DesignerGraphicorTechnical", 116, "Designer,Graphic,Web,Technical,Computer,Internet"}
	occupationDesignerProfessional                       = EnumOccupationItem{occupationDesignerProfessionalID, "Designer - Professional", map[string]string{"Keywords": "Professional,Designer,Product,Usability"}, "DesignerProfessional", 117, "Professional,Designer,Product,Usability"}
	occupationDesignerWindow                             = EnumOccupationItem{occupationDesignerWindowID, "Designer - Window", map[string]string{"Keywords": "Designer,Stainded,Frame,Glass,Window"}, "DesignerWindow", 118, "Designer,Stainded,Frame,Glass,Window"}
	occupationDietitianNutritionist                      = EnumOccupationItem{occupationDietitianNutritionistID, "Dietitian / Nutritionist", map[string]string{"Keywords": "Nutritionist,Dietitian,Hospital,Jail,Prison,Personal"}, "DietitianNutritionist", 119, "Nutritionist,Dietitian,Hospital,Jail,Prison,Personal"}
	occupationDirectororExecutive                        = EnumOccupationItem{occupationDirectororExecutiveID, "Director or Executive", map[string]string{"Keywords": "Executive,Board,Director"}, "DirectororExecutive", 120, "Executive,Board,Director"}
	occupationDiscJockey                                 = EnumOccupationItem{occupationDiscJockeyID, "Disc Jockey", map[string]string{"Keywords": "Disc Jockey,Music,DJ"}, "DiscJockey", 121, "Disc Jockey,Music,DJ"}
	occupationDispatcher                                 = EnumOccupationItem{occupationDispatcherID, "Dispatcher", map[string]string{"Keywords": "Dispatcher,Logistics,Emergency,Truck,Radio,Police,911"}, "Dispatcher", 122, "Dispatcher,Logistics,Emergency,Truck,Radio,Police,911"}
	occupationDogBreeder                                 = EnumOccupationItem{occupationDogBreederID, "Dog Breeder", map[string]string{"Keywords": "Bird,Horse,Dog Breeder,Animal"}, "DogBreeder", 123, "Bird,Horse,Dog Breeder,Animal"}
	occupationDrafterorCartographer                      = EnumOccupationItem{occupationDrafterorCartographerID, "Drafter or Cartographer", map[string]string{"Keywords": "Architecture,CAD,Cartographer,Drafter"}, "DrafterorCartographer", 124, "Architecture,CAD,Cartographer,Drafter"}
	occupationDriverAllOther                             = EnumOccupationItem{occupationDriverAllOtherID, "Driver - All Other", map[string]string{"Keywords": "Driver,Car,Taxi,Rally,Limousine,Other,Race,Shofer,Private"}, "DriverAllOther", 125, "Driver,Car,Taxi,Rally,Limousine,Other,Race,Shofer,Private"}
	occupationDriverTruck                                = EnumOccupationItem{occupationDriverTruckID, "Driver - Truck", map[string]string{"Keywords": "Driver,18 wheeler,Eighteen,Semi,Tractor trailer,Truck,Haul"}, "DriverTruck", 126, "Driver,18 wheeler,Eighteen,Semi,Tractor trailer,Truck,Haul"}
	occupationEconomist                                  = EnumOccupationItem{occupationEconomistID, "Economist", map[string]string{"Keywords": "Analyst,Economist,Market"}, "Economist", 127, "Analyst,Economist,Market"}
	occupationEditorAllOther                             = EnumOccupationItem{occupationEditorAllOtherID, "Editor - All Other", map[string]string{"Keywords": "Newspaper,Book,Editor,Magazine,Other,Online,Print,Blog,Article"}, "EditorAllOther", 128, "Newspaper,Book,Editor,Magazine,Other,Online,Print,Blog,Article"}
	occupationEditorFilm                                 = EnumOccupationItem{occupationEditorFilmID, "Editor - Film", map[string]string{"Keywords": "Editor,Film,Director"}, "EditorFilm", 129, "Editor,Film,Director"}
	occupationElectrician                                = EnumOccupationItem{occupationElectricianID, "Electrician", map[string]string{"Keywords": "Electrician,Utility"}, "Electrician", 130, "Electrician,Utility"}
	occupationEmbalmer                                   = EnumOccupationItem{occupationEmbalmerID, "Embalmer", map[string]string{"Keywords": "Coroner,Embalmer"}, "Embalmer", 131, "Coroner,Embalmer"}
	occupationEngineerAllOther                           = EnumOccupationItem{occupationEngineerAllOtherID, "Engineer - All Other", map[string]string{"Keywords": "Engineer,Other"}, "EngineerAllOther", 132, "Engineer,Other"}
	occupationEngineerCertifiedNetwork                   = EnumOccupationItem{occupationEngineerCertifiedNetworkID, "Engineer - Certified Network", map[string]string{"Keywords": "Engineer,Certified Network"}, "EngineerCertifiedNetwork", 133, "Engineer,Certified Network"}
	occupationEngineerComputerSoftware                   = EnumOccupationItem{occupationEngineerComputerSoftwareID, "Engineer - Computer Software", map[string]string{"Keywords": "Engineer,Computer Software"}, "EngineerComputerSoftware", 134, "Engineer,Computer Software"}
	occupationEngineerComputerSystems                    = EnumOccupationItem{occupationEngineerComputerSystemsID, "Engineer - Computer Systems", map[string]string{"Keywords": "Engineer,Computer Systems"}, "EngineerComputerSystems", 135, "Engineer,Computer Systems"}
	occupationEngineerConstruction                       = EnumOccupationItem{occupationEngineerConstructionID, "Engineer - Construction", map[string]string{"Keywords": "Engineer,Construction"}, "EngineerConstruction", 136, "Engineer,Construction"}
	occupationEngineerElectricalElectronic               = EnumOccupationItem{occupationEngineerElectricalElectronicID, "Engineer - Electrical / Electronic", map[string]string{"Keywords": "Engineer,Electronic,Electrical"}, "EngineerElectricalElectronic", 137, "Engineer,Electronic,Electrical"}
	occupationEngineerEquipment                          = EnumOccupationItem{occupationEngineerEquipmentID, "Engineer - Equipment", map[string]string{"Keywords": "Engineer,Equipment"}, "EngineerEquipment", 138, "Engineer,Equipment"}
	occupationEngineerFacilities                         = EnumOccupationItem{occupationEngineerFacilitiesID, "Engineer - Facilities", map[string]string{"Keywords": "Engineer,Facilities"}, "EngineerFacilities", 139, "Engineer,Facilities"}
	occupationEngineerFlight                             = EnumOccupationItem{occupationEngineerFlightID, "Engineer - Flight", map[string]string{"Keywords": "Engineer,Flight,Airplane,Mechanic"}, "EngineerFlight", 140, "Engineer,Flight,Airplane,Mechanic"}
	occupationEngineerMechanical                         = EnumOccupationItem{occupationEngineerMechanicalID, "Engineer - Mechanical", map[string]string{"Keywords": "Engineer,Mechanical"}, "EngineerMechanical", 141, "Engineer,Mechanical"}
	occupationEngineerOperating                          = EnumOccupationItem{occupationEngineerOperatingID, "Engineer - Operating", map[string]string{"Keywords": "Engineer,Operating"}, "EngineerOperating", 142, "Engineer,Operating"}
	occupationEngineerPetroleumorMining                  = EnumOccupationItem{occupationEngineerPetroleumorMiningID, "Engineer - Petroleum or Mining", map[string]string{"Keywords": "Engineer,Petroleum,Mining"}, "EngineerPetroleumorMining", 143, "Engineer,Petroleum,Mining"}
	occupationEngineerSafety                             = EnumOccupationItem{occupationEngineerSafetyID, "Engineer - Safety", map[string]string{"Keywords": "Engineer,Safety"}, "EngineerSafety", 144, "Engineer,Safety"}
	occupationEngineerSales                              = EnumOccupationItem{occupationEngineerSalesID, "Engineer - Sales", map[string]string{"Keywords": "Engineer,Sales"}, "EngineerSales", 145, "Engineer,Sales"}
	occupationEntertainerPerformer                       = EnumOccupationItem{occupationEntertainerPerformerID, "Entertainer / Performer", map[string]string{"Keywords": "Magician,Exotic,Clown,Acrobat,Performer,Dancer,Entertainer,Stripper,Comic"}, "EntertainerPerformer", 146, "Magician,Exotic,Clown,Acrobat,Performer,Dancer,Entertainer,Stripper,Comic"}
	occupationExpediter                                  = EnumOccupationItem{occupationExpediterID, "Expediter", map[string]string{"Keywords": "Expediter,Accelerator"}, "Expediter", 147, "Expediter,Accelerator"}
	occupationFactoryWorker                              = EnumOccupationItem{occupationFactoryWorkerID, "Factory Worker", map[string]string{"Keywords": "Assembler,Warehouse,Plant,Production,Factory Worker,Technician,Line,Weld"}, "FactoryWorker", 148, "Assembler,Warehouse,Plant,Production,Factory Worker,Technician,Line,Weld"}
	occupationFiremanWomanChiefCaptLt                    = EnumOccupationItem{occupationFiremanWomanChiefCaptLtID, "Fireman / Woman - Chief / Capt / Lt", map[string]string{"Keywords": "Woman,Lieutenant,Captain,Fighter,Fireman,Chief"}, "FiremanWomanChiefCaptLt", 149, "Woman,Lieutenant,Captain,Fighter,Fireman,Chief"}
	occupationFiremanWomanNonChief                       = EnumOccupationItem{occupationFiremanWomanNonChiefID, "Fireman / Woman - Non Chief", map[string]string{"Keywords": "Woman,Fighter,Fireman"}, "FiremanWomanNonChief", 150, "Woman,Fighter,Fireman"}
	occupationFisherman                                  = EnumOccupationItem{occupationFishermanID, "Fisherman", map[string]string{"Keywords": "Sea,Shark,Crab,Oysters,Muscles,Fisherman,Fish,Sport,Clams"}, "Fisherman", 151, "Sea,Shark,Crab,Oysters,Muscles,Fisherman,Fish,Sport,Clams"}
	occupationFlightAttendant                            = EnumOccupationItem{occupationFlightAttendantID, "Flight Attendant", map[string]string{"Keywords": "Air hosts,Stewards,Flight Attendant"}, "FlightAttendant", 152, "Air hosts,Stewards,Flight Attendant"}
	occupationFloormenSupervisor                         = EnumOccupationItem{occupationFloormenSupervisorID, "Floormen (Supervisor)", map[string]string{"Keywords": "Plant,Floormen,Production,Supervisor,Factory,Line"}, "FloormenSupervisor", 153, "Plant,Floormen,Production,Supervisor,Factory,Line"}
	occupationFlorist                                    = EnumOccupationItem{occupationFloristID, "Florist", map[string]string{"Keywords": "Décor,Flowers,Florist,Floral"}, "Florist", 154, "Décor,Flowers,Florist,Floral"}
	occupationForemanForewoman                           = EnumOccupationItem{occupationForemanForewomanID, "Foreman / Forewoman", map[string]string{"Keywords": "Foreman,Warehouse,Plant,Construction,Factory"}, "ForemanForewoman", 155, "Foreman,Warehouse,Plant,Construction,Factory"}
	occupationForester                                   = EnumOccupationItem{occupationForesterID, "Forester", map[string]string{"Keywords": "Logs,Chainsaw,Tree Trimmer,Forester"}, "Forester", 156, "Logs,Chainsaw,Tree Trimmer,Forester"}
	occupationFundraiser                                 = EnumOccupationItem{occupationFundraiserID, "Fundraiser", map[string]string{"Keywords": "Volenteer,Not for profit,Fundraiser"}, "Fundraiser", 157, "Volenteer,Not for profit,Fundraiser"}
	occupationGeographer                                 = EnumOccupationItem{occupationGeographerID, "Geographer", map[string]string{"Keywords": "Geographer,Surveyor,Maps"}, "Geographer", 158, "Geographer,Surveyor,Maps"}
	occupationGovtOfficialElected                        = EnumOccupationItem{occupationGovtOfficialElectedID, "Govt. Official - Elected", map[string]string{"Keywords": "Government,Elected,Official"}, "GovtOfficialElected", 159, "Government,Elected,Official"}
	occupationGrader                                     = EnumOccupationItem{occupationGraderID, "Grader", map[string]string{"Keywords": "Soil,Grader"}, "Grader", 160, "Soil,Grader"}
	occupationGuardEmbassy                               = EnumOccupationItem{occupationGuardEmbassyID, "Guard - Embassy", map[string]string{"Keywords": "Security,Embassy,Guard"}, "GuardEmbassy", 161, "Security,Embassy,Guard"}
	occupationGuardSecurityorPrison                      = EnumOccupationItem{occupationGuardSecurityorPrisonID, "Guard - Security or Prison", map[string]string{"Keywords": "Security,TSA,Guard,Prison,Bouncer"}, "GuardSecurityorPrison", 162, "Security,TSA,Guard,Prison,Bouncer"}
	occupationGunsmith                                   = EnumOccupationItem{occupationGunsmithID, "Gunsmith", map[string]string{"Keywords": "Pistol,Gunsmith,Firearm,Rifle,Repair"}, "Gunsmith", 163, "Pistol,Gunsmith,Firearm,Rifle,Repair"}
	occupationHairdresserBarber                          = EnumOccupationItem{occupationHairdresserBarberID, "Hairdresser / Barber", map[string]string{"Keywords": "Styleist,Barber,Hairdresser,Braider"}, "HairdresserBarber", 164, "Styleist,Barber,Hairdresser,Braider"}
	occupationHistorian                                  = EnumOccupationItem{occupationHistorianID, "Historian", map[string]string{"Keywords": "Museum,Historian,Research,College"}, "Historian", 165, "Museum,Historian,Research,College"}
	occupationHostorHostessRestaurant                    = EnumOccupationItem{occupationHostorHostessRestaurantID, "Host or Hostess Restaurant", map[string]string{"Keywords": "Restaurant,Greeter,Hostess"}, "HostorHostessRestaurant", 166, "Restaurant,Greeter,Hostess"}
	occupationHousekeeperorMaid                          = EnumOccupationItem{occupationHousekeeperorMaidID, "Housekeeper or Maid", map[string]string{"Keywords": "Maid,Butler,Housekeeper or Maid"}, "HousekeeperorMaid", 167, "Maid,Butler,Housekeeper or Maid"}
	occupationHumanResourcesRepresentative               = EnumOccupationItem{occupationHumanResourcesRepresentativeID, "Human Resources Representative", map[string]string{"Keywords": "Training,Human resources representative,Coordinate,People Services,Benefits,Recruiter"}, "HumanResourcesRepresentative", 168, "Training,Human resources representative,Coordinate,People Services,Benefits,Recruiter"}
	occupationIllustrator                                = EnumOccupationItem{occupationIllustratorID, "Illustrator", map[string]string{"Keywords": "Illustrator,Draw,Cartoonist,Artist"}, "Illustrator", 169, "Illustrator,Draw,Cartoonist,Artist"}
	occupationInspectorAgricultural                      = EnumOccupationItem{occupationInspectorAgriculturalID, "Inspector - Agricultural", map[string]string{"Keywords": "Inspector,Agricultural,Food,Farm"}, "InspectorAgricultural", 170, "Inspector,Agricultural,Food,Farm"}
	occupationInspectorAircraftAccessories               = EnumOccupationItem{occupationInspectorAircraftAccessoriesID, "Inspector - Aircraft Accessories", map[string]string{"Keywords": "Inspector,Aircraft,Accessories"}, "InspectorAircraftAccessories", 171, "Inspector,Aircraft,Accessories"}
	occupationInspectorAllOther                          = EnumOccupationItem{occupationInspectorAllOtherID, "Inspector - All Other", map[string]string{"Keywords": "Inspector,Other,Vehicle"}, "InspectorAllOther", 172, "Inspector,Other,Vehicle"}
	occupationInspectorConstruction                      = EnumOccupationItem{occupationInspectorConstructionID, "Inspector - Construction", map[string]string{"Keywords": "Inspector,Construction"}, "InspectorConstruction", 173, "Inspector,Construction"}
	occupationInspectorPostal                            = EnumOccupationItem{occupationInspectorPostalID, "Inspector - Postal", map[string]string{"Keywords": "Inspector,Delivery,Postal"}, "InspectorPostal", 174, "Inspector,Delivery,Postal"}
	occupationInspectorWhiteCollar                       = EnumOccupationItem{occupationInspectorWhiteCollarID, "Inspector - White Collar", map[string]string{"Keywords": "Inspector,White Collar"}, "InspectorWhiteCollar", 175, "Inspector,White Collar"}
	occupationInvestigatorPrivate                        = EnumOccupationItem{occupationInvestigatorPrivateID, "Investigator - Private", map[string]string{"Keywords": "Investigator,Insurance,Private"}, "InvestigatorPrivate", 176, "Investigator,Insurance,Private"}
	occupationInvestmentBanker                           = EnumOccupationItem{occupationInvestmentBankerID, "Investment Banker", map[string]string{"Keywords": "Investment Banker,Bond,Stock"}, "InvestmentBanker", 177, "Investment Banker,Bond,Stock"}
	occupationInvestorPrivate                            = EnumOccupationItem{occupationInvestorPrivateID, "Investor - Private", map[string]string{"Keywords": "Investor,Private,Personal"}, "InvestorPrivate", 178, "Investor,Private,Personal"}
	occupationJournalist                                 = EnumOccupationItem{occupationJournalistID, "Journalist", map[string]string{"Keywords": "Journalist,Newspaper,Writer,Web,Magazine,Internet"}, "Journalist", 179, "Journalist,Newspaper,Writer,Web,Magazine,Internet"}
	occupationJourneyman                                 = EnumOccupationItem{occupationJourneymanID, "Journeyman", map[string]string{"Keywords": "Apprentice,Electrician,Tradesman,Journeyman,Carpenter"}, "Journeyman", 180, "Apprentice,Electrician,Tradesman,Journeyman,Carpenter"}
	occupationJudge                                      = EnumOccupationItem{occupationJudgeID, "Judge", map[string]string{"Keywords": "Judge,Judicial,Magistrate"}, "Judge", 181, "Judge,Judicial,Magistrate"}
	occupationLaborRelationsWorker                       = EnumOccupationItem{occupationLaborRelationsWorkerID, "Labor Relations Worker", map[string]string{"Keywords": "Union,Labor relations worker,Human Resources"}, "LaborRelationsWorker", 182, "Union,Labor relations worker,Human Resources"}
	occupationLandscaper                                 = EnumOccupationItem{occupationLandscaperID, "Landscaper", map[string]string{"Keywords": "Mow,Tree,Landscaper,Fertitlize,Plants,Lawn care"}, "Landscaper", 183, "Mow,Tree,Landscaper,Fertitlize,Plants,Lawn care"}
	occupationLibrarian                                  = EnumOccupationItem{occupationLibrarianID, "Librarian", map[string]string{"Keywords": "Books,Librarian,Research"}, "Librarian", 184, "Books,Librarian,Research"}
	occupationLifeGuard                                  = EnumOccupationItem{occupationLifeGuardID, "Life Guard", map[string]string{"Keywords": "Certified,Life guard,Pool,Beach,Lake,First aid"}, "LifeGuard", 185, "Certified,Life guard,Pool,Beach,Lake,First aid"}
	occupationLinguist                                   = EnumOccupationItem{occupationLinguistID, "Linguist", map[string]string{"Keywords": "Grammar,Language,Linguist,Interpreter,Translator"}, "Linguist", 186, "Grammar,Language,Linguist,Interpreter,Translator"}
	occupationLithographer                               = EnumOccupationItem{occupationLithographerID, "Lithographer", map[string]string{"Keywords": "Lithographer,Photographer,Artist,Copy"}, "Lithographer", 187, "Lithographer,Photographer,Artist,Copy"}
	occupationLobbyist                                   = EnumOccupationItem{occupationLobbyistID, "Lobbyist", map[string]string{"Keywords": "Politcal,Medical,Lobbyist"}, "Lobbyist", 188, "Politcal,Medical,Lobbyist"}
	occupationLocksmith                                  = EnumOccupationItem{occupationLocksmithID, "Locksmith", map[string]string{"Keywords": "Handyman,Locksmith"}, "Locksmith", 189, "Handyman,Locksmith"}
	occupationLongshoremen                               = EnumOccupationItem{occupationLongshoremenID, "Longshoremen", map[string]string{"Keywords": "Longshoremen,Boat"}, "Longshoremen", 190, "Longshoremen,Boat"}
	occupationMachinist                                  = EnumOccupationItem{occupationMachinistID, "Machinist", map[string]string{"Keywords": "Press,Machinist,Lathe"}, "Machinist", 191, "Press,Machinist,Lathe"}
	occupationManagerAirport                             = EnumOccupationItem{occupationManagerAirportID, "Manager - Airport", map[string]string{"Keywords": "Airport,Manager"}, "ManagerAirport", 192, "Airport,Manager"}
	occupationManagerAllOtherDegreed                     = EnumOccupationItem{occupationManagerAllOtherDegreedID, "Manager - All Other (Degreed)", map[string]string{"Keywords": "Degreed,Manager,Other"}, "ManagerAllOtherDegreed", 193, "Degreed,Manager,Other"}
	occupationManagerCafeteria                           = EnumOccupationItem{occupationManagerCafeteriaID, "Manager - Cafeteria", map[string]string{"Keywords": "Cafeteria,Manager,Food service"}, "ManagerCafeteria", 194, "Cafeteria,Manager,Food service"}
	occupationManagerCity                                = EnumOccupationItem{occupationManagerCityID, "Manager - City", map[string]string{"Keywords": "Manager,Planner,City"}, "ManagerCity", 195, "Manager,Planner,City"}
	occupationManagerClericalStaff                       = EnumOccupationItem{occupationManagerClericalStaffID, "Manager - Clerical Staff", map[string]string{"Keywords": "Manager,Clerical staff"}, "ManagerClericalStaff", 196, "Manager,Clerical staff"}
	occupationManagerConvenienceorGasStationStore        = EnumOccupationItem{occupationManagerConvenienceorGasStationStoreID, "Manager - Convenience or Gas Station Store", map[string]string{"Keywords": "Convenience,Store,Manager,Gas station,Shop"}, "ManagerConvenienceorGasStationStore", 197, "Convenience,Store,Manager,Gas station,Shop"}
	occupationManagerDepartmentStore                     = EnumOccupationItem{occupationManagerDepartmentStoreID, "Manager - Department Store", map[string]string{"Keywords": "Target,Store,Manager,Lowes,Home Depot,Department"}, "ManagerDepartmentStore", 198, "Target,Store,Manager,Lowes,Home Depot,Department"}
	occupationManagerFinancialorCredit                   = EnumOccupationItem{occupationManagerFinancialorCreditID, "Manager - Financial or Credit", map[string]string{"Keywords": "Financial,Manager,Credit"}, "ManagerFinancialorCredit", 199, "Financial,Manager,Credit"}
	occupationManagerGeneral                             = EnumOccupationItem{occupationManagerGeneralID, "Manager - General", map[string]string{"Keywords": "General,Manager"}, "ManagerGeneral", 200, "General,Manager"}
	occupationManagerHealthClub                          = EnumOccupationItem{occupationManagerHealthClubID, "Manager - Health Club", map[string]string{"Keywords": "Gym,Health Club,Manager"}, "ManagerHealthClub", 201, "Gym,Health Club,Manager"}
	occupationManagerHotel                               = EnumOccupationItem{occupationManagerHotelID, "Manager - Hotel", map[string]string{"Keywords": "Manager,Hotel"}, "ManagerHotel", 202, "Manager,Hotel"}
	occupationManagerHumanResources                      = EnumOccupationItem{occupationManagerHumanResourcesID, "Manager - Human Resources", map[string]string{"Keywords": "Human resources,Manager"}, "ManagerHumanResources", 203, "Human resources,Manager"}
	occupationManagerMerchandise                         = EnumOccupationItem{occupationManagerMerchandiseID, "Manager - Merchandise", map[string]string{"Keywords": "Manager,Merchandise"}, "ManagerMerchandise", 204, "Manager,Merchandise"}
	occupationManagerOffice                              = EnumOccupationItem{occupationManagerOfficeID, "Manager - Office", map[string]string{"Keywords": "Manager,Office"}, "ManagerOffice", 205, "Manager,Office"}
	occupationManagerOperations                          = EnumOccupationItem{occupationManagerOperationsID, "Manager - Operations", map[string]string{"Keywords": "Operations,Manager"}, "ManagerOperations", 206, "Operations,Manager"}
	occupationManagerProduction                          = EnumOccupationItem{occupationManagerProductionID, "Manager - Production", map[string]string{"Keywords": "Manager,Production,Assembly,Line"}, "ManagerProduction", 207, "Manager,Production,Assembly,Line"}
	occupationManagerProfessionalTechStaff               = EnumOccupationItem{occupationManagerProfessionalTechStaffID, "Manager - Professional & Tech Staff", map[string]string{"Keywords": "Professional,Staff,Manager,Tech"}, "ManagerProfessionalTechStaff", 208, "Professional,Staff,Manager,Tech"}
	occupationManagerProject                             = EnumOccupationItem{occupationManagerProjectID, "Manager - Project", map[string]string{"Keywords": "Manager,Project"}, "ManagerProject", 209, "Manager,Project"}
	occupationManagerPropertyNonResident                 = EnumOccupationItem{occupationManagerPropertyNonResidentID, "Manager - Property Non-Resident", map[string]string{"Keywords": "Manager,Non-Resident,Property"}, "ManagerPropertyNonResident", 210, "Manager,Non-Resident,Property"}
	occupationManagerPropertyResident                    = EnumOccupationItem{occupationManagerPropertyResidentID, "Manager - Property Resident", map[string]string{"Keywords": "Property resident,Manager"}, "ManagerPropertyResident", 211, "Property resident,Manager"}
	occupationManagerRestaurantFastFood                  = EnumOccupationItem{occupationManagerRestaurantFastFoodID, "Manager - Restaurant Fast Food", map[string]string{"Keywords": "Restaurant,Manager,Fast food"}, "ManagerRestaurantFastFood", 212, "Restaurant,Manager,Fast food"}
	occupationManagerRestaurantNonFastFood               = EnumOccupationItem{occupationManagerRestaurantNonFastFoodID, "Manager - Restaurant Non Fast Food", map[string]string{"Keywords": "Restaurant,Manager,Non fast food"}, "ManagerRestaurantNonFastFood", 213, "Restaurant,Manager,Non fast food"}
	occupationManagerSales                               = EnumOccupationItem{occupationManagerSalesID, "Manager - Sales", map[string]string{"Keywords": "Sales,Service,Claims,Manager"}, "ManagerSales", 214, "Sales,Service,Claims,Manager"}
	occupationManagerSecurityScreener                    = EnumOccupationItem{occupationManagerSecurityScreenerID, "Manager - Security Screener", map[string]string{"Keywords": "Security,Screener,Manager"}, "ManagerSecurityScreener", 215, "Security,Screener,Manager"}
	occupationManagerShippingReceiving                   = EnumOccupationItem{occupationManagerShippingReceivingID, "Manager - Shipping / Receiving", map[string]string{"Keywords": "Receiving,Manager,Shipping"}, "ManagerShippingReceiving", 216, "Receiving,Manager,Shipping"}
	occupationManagerStage                               = EnumOccupationItem{occupationManagerStageID, "Manager - Stage", map[string]string{"Keywords": "Manager,Stage,Theater"}, "ManagerStage", 217, "Manager,Stage,Theater"}
	occupationManagerSupermarket                         = EnumOccupationItem{occupationManagerSupermarketID, "Manager - Supermarket", map[string]string{"Keywords": "Supermarket,Manager"}, "ManagerSupermarket", 218, "Supermarket,Manager"}
	occupationManagerorOwnerSandwichShop                 = EnumOccupationItem{occupationManagerorOwnerSandwichShopID, "Manager or Owner - Sandwich Shop", map[string]string{"Keywords": "Owner,Sandwich shop,Firehouse,Manager,Jersey Mikes,Jimmy Johns,Subway"}, "ManagerorOwnerSandwichShop", 219, "Owner,Sandwich shop,Firehouse,Manager,Jersey Mikes,Jimmy Johns,Subway"}
	occupationManicurist                                 = EnumOccupationItem{occupationManicuristID, "Manicurist", map[string]string{"Keywords": "Manicurist,Nails"}, "Manicurist", 220, "Manicurist,Nails"}
	occupationMarketingRepresentative                    = EnumOccupationItem{occupationMarketingRepresentativeID, "Marketing Representative", map[string]string{"Keywords": "Marketing,Representative,Coordinator,Public"}, "MarketingRepresentative", 221, "Marketing,Representative,Coordinator,Public"}
	occupationMarshalFire                                = EnumOccupationItem{occupationMarshalFireID, "Marshal - Fire", map[string]string{"Keywords": "Fire,Marshal"}, "MarshalFire", 222, "Fire,Marshal"}
	occupationMarshalUSDeputy                            = EnumOccupationItem{occupationMarshalUSDeputyID, "Marshal - U.S. Deputy", map[string]string{"Keywords": "US,Deputy,Marshal"}, "MarshalUSDeputy", 223, "US,Deputy,Marshal"}
	occupationMasseuse                                   = EnumOccupationItem{occupationMasseuseID, "Masseuse", map[string]string{"Keywords": "Therapy,Masseuse"}, "Masseuse", 224, "Therapy,Masseuse"}
	occupationMathematician                              = EnumOccupationItem{occupationMathematicianID, "Mathematician", map[string]string{"Keywords": "Statistics,Mathematician"}, "Mathematician", 225, "Statistics,Mathematician"}
	occupationMeatcutterButcher                          = EnumOccupationItem{occupationMeatcutterButcherID, "Meatcutter / Butcher", map[string]string{"Keywords": "Deli,Butcher,Food service,Meatcutter"}, "MeatcutterButcher", 226, "Deli,Butcher,Food service,Meatcutter"}
	occupationMechanicorServicemanAuto                   = EnumOccupationItem{occupationMechanicorServicemanAutoID, "Mechanic or Serviceman - Auto", map[string]string{"Keywords": "Mechanic or Serviceman - Auto,car repair"}, "MechanicorServicemanAuto", 227, "Mechanic or Serviceman - Auto,car repair"}
	occupationMechanicorServicemanBoat                   = EnumOccupationItem{occupationMechanicorServicemanBoatID, "Mechanic or Serviceman - Boat", map[string]string{"Keywords": "Serviceman,Boat,Jet Ski,Mechanic"}, "MechanicorServicemanBoat", 228, "Serviceman,Boat,Jet Ski,Mechanic"}
	occupationMechanicorServicemanDiesel                 = EnumOccupationItem{occupationMechanicorServicemanDieselID, "Mechanic or Serviceman - Diesel", map[string]string{"Keywords": "Truck repair,Serviceman,Auto,Diesel,Mechanic"}, "MechanicorServicemanDiesel", 229, "Truck repair,Serviceman,Auto,Diesel,Mechanic"}
	occupationMerchant                                   = EnumOccupationItem{occupationMerchantID, "Merchant", map[string]string{"Keywords": "Vendor,Merchant,Whole sale"}, "Merchant", 230, "Vendor,Merchant,Whole sale"}
	occupationMillwright                                 = EnumOccupationItem{occupationMillwrightID, "Millwright", map[string]string{"Keywords": "Machine,Lathe,Millwright"}, "Millwright", 231, "Machine,Lathe,Millwright"}
	occupationMortician                                  = EnumOccupationItem{occupationMorticianID, "Mortician", map[string]string{"Keywords": "Funeral,Mortician"}, "Mortician", 232, "Funeral,Mortician"}
	occupationMusicianClassical                          = EnumOccupationItem{occupationMusicianClassicalID, "Musician - Classical", map[string]string{"Keywords": "Country,Musician,Bass,Drum,Jazz,Pop,Rap,Sing,Rock,Pianist,Classical,Flute,Songwriter,MC,Compose,Guitarist,DJ"}, "MusicianClassical", 233, "Country,Musician,Bass,Drum,Jazz,Pop,Rap,Sing,Rock,Pianist,Classical,Flute,Songwriter,MC,Compose,Guitarist,DJ"}
	occupationMusicianOther                              = EnumOccupationItem{occupationMusicianOtherID, "Musician - Other", map[string]string{"Keywords": "Other,Country,Musician,Bass,Drum,Jazz,Pop,Rap,Sing,Rock,Pianist,Composer,Flute,Songwriter,MC,Guitarist,DJ"}, "MusicianOther", 234, "Other,Country,Musician,Bass,Drum,Jazz,Pop,Rap,Sing,Rock,Pianist,Composer,Flute,Songwriter,MC,Guitarist,DJ"}
	occupationNurseCNACertifiedNursingAssistant          = EnumOccupationItem{occupationNurseCNACertifiedNursingAssistantID, "Nurse - CNA (Certified Nursing Assistant)", map[string]string{"Keywords": "Certified,Assistant,Nurse,CNA"}, "NurseCNACertifiedNursingAssistant", 235, "Certified,Assistant,Nurse,CNA"}
	occupationNurseLVNorLPN                              = EnumOccupationItem{occupationNurseLVNorLPNID, "Nurse - LVN or LPN", map[string]string{"Keywords": "LPN,LVN,Nurse"}, "NurseLVNorLPN", 236, "LPN,LVN,Nurse"}
	occupationNurseRN                                    = EnumOccupationItem{occupationNurseRNID, "Nurse - RN", map[string]string{"Keywords": "Registered,Nurse,RN"}, "NurseRN", 237, "Registered,Nurse,RN"}
	occupationNursePractitioner                          = EnumOccupationItem{occupationNursePractitionerID, "Nurse Practitioner", map[string]string{"Keywords": "Nurse,Practitioner"}, "NursePractitioner", 238, "Nurse,Practitioner"}
	occupationOceanographer                              = EnumOccupationItem{occupationOceanographerID, "Oceanographer", map[string]string{"Keywords": "Oceanographer"}, "Oceanographer", 239, "Oceanographer"}
	occupationOfficerCorrectional                        = EnumOccupationItem{occupationOfficerCorrectionalID, "Officer - Correctional", map[string]string{"Keywords": "Juvenile,Correctional,Officer,Prison"}, "OfficerCorrectional", 240, "Juvenile,Correctional,Officer,Prison"}
	occupationOfficerCourt                               = EnumOccupationItem{occupationOfficerCourtID, "Officer - Court", map[string]string{"Keywords": "Officer,Court"}, "OfficerCourt", 241, "Officer,Court"}
	occupationOfficerForeignService                      = EnumOccupationItem{occupationOfficerForeignServiceID, "Officer - Foreign Service", map[string]string{"Keywords": "Foreign,Service,Officer,Military"}, "OfficerForeignService", 242, "Foreign,Service,Officer,Military"}
	occupationOfficerLoan                                = EnumOccupationItem{occupationOfficerLoanID, "Officer - Loan", map[string]string{"Keywords": "Financial,Loan,Lend,Officer"}, "OfficerLoan", 243, "Financial,Loan,Lend,Officer"}
	occupationOfficerPolice                              = EnumOccupationItem{occupationOfficerPoliceID, "Officer - Police", map[string]string{"Keywords": "Cop,Officer,Deputy,Sheriff,Police"}, "OfficerPolice", 244, "Cop,Officer,Deputy,Sheriff,Police"}
	occupationOfficerPoliceChiefCaptain                  = EnumOccupationItem{occupationOfficerPoliceChiefCaptainID, "Officer - Police Chief & Captain", map[string]string{"Keywords": "Cop,Captain,Officer,Deputy,Sheriff,Police,Chief"}, "OfficerPoliceChiefCaptain", 245, "Cop,Captain,Officer,Deputy,Sheriff,Police,Chief"}
	occupationOfficerPoliceDetectiveSgtLt                = EnumOccupationItem{occupationOfficerPoliceDetectiveSgtLtID, "Officer - Police Detective / Sgt / Lt", map[string]string{"Keywords": "Detective,Lieutenant,Sergeant,Officer,Police"}, "OfficerPoliceDetectiveSgtLt", 246, "Detective,Lieutenant,Sergeant,Officer,Police"}
	occupationOfficerProbationParole                     = EnumOccupationItem{occupationOfficerProbationParoleID, "Officer - Probation / Parole", map[string]string{"Keywords": "Probation,Cop,Enforcer,Child support,Officer,Parole"}, "OfficerProbationParole", 247, "Probation,Cop,Enforcer,Child support,Officer,Parole"}
	occupationOfficerTelecommunications                  = EnumOccupationItem{occupationOfficerTelecommunicationsID, "Officer - Telecommunications", map[string]string{"Keywords": "Officer,Telecommunications"}, "OfficerTelecommunications", 248, "Officer,Telecommunications"}
	occupationOfficerWarrant                             = EnumOccupationItem{occupationOfficerWarrantID, "Officer - Warrant", map[string]string{"Keywords": "Warrant,Serve,Arrest,Process,Officer"}, "OfficerWarrant", 249, "Warrant,Serve,Arrest,Process,Officer"}
	occupationOfficerorManagerBank                       = EnumOccupationItem{occupationOfficerorManagerBankID, "Officer or Manager - Bank", map[string]string{"Keywords": "Financial,Manager,Bank,Officer"}, "OfficerorManagerBank", 250, "Financial,Manager,Bank,Officer"}
	occupationOperatorAllOther                           = EnumOccupationItem{occupationOperatorAllOtherID, "Operator - All Other", map[string]string{"Keywords": "Operator,Other"}, "OperatorAllOther", 251, "Operator,Other"}
	occupationOperatorBusiness                           = EnumOccupationItem{occupationOperatorBusinessID, "Operator - Business", map[string]string{"Keywords": "Business,Operator"}, "OperatorBusiness", 252, "Business,Operator"}
	occupationOperatorControlRoom                        = EnumOccupationItem{occupationOperatorControlRoomID, "Operator - Control Room", map[string]string{"Keywords": "Plant,Operator,Control room"}, "OperatorControlRoom", 253, "Plant,Operator,Control room"}
	occupationOperatorDataEntry                          = EnumOccupationItem{occupationOperatorDataEntryID, "Operator - Data Entry", map[string]string{"Keywords": "Computer,Data entry,Operator"}, "OperatorDataEntry", 254, "Computer,Data entry,Operator"}
	occupationOperatorForkLift                           = EnumOccupationItem{occupationOperatorForkLiftID, "Operator - Fork Lift", map[string]string{"Keywords": "Fork lift,Operator,Skid loader"}, "OperatorForkLift", 255, "Fork lift,Operator,Skid loader"}
	occupationOperatorHeavyEquipment                     = EnumOccupationItem{occupationOperatorHeavyEquipmentID, "Operator - Heavy Equipment", map[string]string{"Keywords": "Oil,Miner,Crane,Heavy equipment,Riggor,Operator,Coal,Skid,Tractor"}, "OperatorHeavyEquipment", 256, "Oil,Miner,Crane,Heavy equipment,Riggor,Operator,Coal,Skid,Tractor"}
	occupationOperatorMachinePrecision                   = EnumOccupationItem{occupationOperatorMachinePrecisionID, "Operator - Machine, Precision", map[string]string{"Keywords": "Machine,Operator,Precision"}, "OperatorMachinePrecision", 257, "Machine,Operator,Precision"}
	occupationOperatorNuclearReactor                     = EnumOccupationItem{occupationOperatorNuclearReactorID, "Operator - Nuclear Reactor", map[string]string{"Keywords": "Plant,Operator,Nuclear reactor"}, "OperatorNuclearReactor", 258, "Plant,Operator,Nuclear reactor"}
	occupationOperatorTelephone                          = EnumOccupationItem{occupationOperatorTelephoneID, "Operator - Telephone", map[string]string{"Keywords": "Telephone,Operator,Clerk,Relay"}, "OperatorTelephone", 259, "Telephone,Operator,Clerk,Relay"}
	occupationOperatorWastewaterTreatmentPlantClassIV    = EnumOccupationItem{occupationOperatorWastewaterTreatmentPlantClassIVID, "Operator - Wastewater Treatment Plant Class IV", map[string]string{"Keywords": "4,Operator,Plant Class IV,Wastewater treatment,Four"}, "OperatorWastewaterTreatmentPlantClassIV", 260, "4,Operator,Plant Class IV,Wastewater treatment,Four"}
	occupationOptician                                   = EnumOccupationItem{occupationOpticianID, "Optician", map[string]string{"Keywords": "Glasses,Optician,Eyes"}, "Optician", 261, "Glasses,Optician,Eyes"}
	occupationOptometrist                                = EnumOccupationItem{occupationOptometristID, "Optometrist", map[string]string{"Keywords": "Optometrist,Glasses,Eyes"}, "Optometrist", 262, "Optometrist,Glasses,Eyes"}
	occupationOrthodontist                               = EnumOccupationItem{occupationOrthodontistID, "Orthodontist", map[string]string{"Keywords": "Teeth,Dentist,Surgeon,Orthodontist"}, "Orthodontist", 263, "Teeth,Dentist,Surgeon,Orthodontist"}
	occupationOwnerAllOther                              = EnumOccupationItem{occupationOwnerAllOtherID, "Owner - All Other", map[string]string{"Keywords": "Owner,Business,Other"}, "OwnerAllOther", 264, "Owner,Business,Other"}
	occupationOwnerBar                                   = EnumOccupationItem{occupationOwnerBarID, "Owner - Bar", map[string]string{"Keywords": "Owner,Alcohol,Bar"}, "OwnerBar", 265, "Owner,Alcohol,Bar"}
	occupationOwnerBeautyBarberShop                      = EnumOccupationItem{occupationOwnerBeautyBarberShopID, "Owner - Beauty / Barber Shop", map[string]string{"Keywords": "Owner,Barber shop,Beauty"}, "OwnerBeautyBarberShop", 266, "Owner,Barber shop,Beauty"}
	occupationOwnerDealershipAutoDealer                  = EnumOccupationItem{occupationOwnerDealershipAutoDealerID, "Owner - Dealership / Auto Dealer", map[string]string{"Keywords": "Owner,Auto,Dealership"}, "OwnerDealershipAutoDealer", 267, "Owner,Auto,Dealership"}
	occupationOwnerorManagerFarmOrRanch                  = EnumOccupationItem{occupationOwnerorManagerFarmOrRanchID, "Owner or Manager - Farm Or Ranch", map[string]string{"Keywords": "Rancher,Owner,Manager,Farmer"}, "OwnerorManagerFarmOrRanch", 268, "Rancher,Owner,Manager,Farmer"}
	occupationPainter                                    = EnumOccupationItem{occupationPainterID, "Painter", map[string]string{"Keywords": "Painter"}, "Painter", 269, "Painter"}
	occupationParalegal                                  = EnumOccupationItem{occupationParalegalID, "Paralegal", map[string]string{"Keywords": "Legal,Lawyer assistant,Para"}, "Paralegal", 270, "Legal,Lawyer assistant,Para"}
	occupationParamedicorEMT                             = EnumOccupationItem{occupationParamedicorEMTID, "Paramedic or EMT", map[string]string{"Keywords": "Emergency services technician,Paramedic,EMT"}, "ParamedicorEMT", 271, "Emergency services technician,Paramedic,EMT"}
	occupationParkForestRanger                           = EnumOccupationItem{occupationParkForestRangerID, "Park / Forest Ranger", map[string]string{"Keywords": "Forest,Park,Guide,Ranger,Tour"}, "ParkForestRanger", 272, "Forest,Park,Guide,Ranger,Tour"}
	occupationPathologistSpeech                          = EnumOccupationItem{occupationPathologistSpeechID, "Pathologist - Speech", map[string]string{"Keywords": "Teacher,Pathologist,Language,Speech"}, "PathologistSpeech", 273, "Teacher,Pathologist,Language,Speech"}
	occupationPersonnelManagementSpecialist              = EnumOccupationItem{occupationPersonnelManagementSpecialistID, "Personnel Management Specialist", map[string]string{"Keywords": "Consultant,Personnel Management Specialist"}, "PersonnelManagementSpecialist", 274, "Consultant,Personnel Management Specialist"}
	occupationPestControlWorkerorExterminator            = EnumOccupationItem{occupationPestControlWorkerorExterminatorID, "Pest Control Worker or Exterminator", map[string]string{"Keywords": "Exterminator,Pest Control"}, "PestControlWorkerorExterminator", 275, "Exterminator,Pest Control"}
	occupationPharmacist                                 = EnumOccupationItem{occupationPharmacistID, "Pharmacist", map[string]string{"Keywords": "Perscription,Pharmacist,Medicine,Fill order"}, "Pharmacist", 276, "Perscription,Pharmacist,Medicine,Fill order"}
	occupationPharmacologist                             = EnumOccupationItem{occupationPharmacologistID, "Pharmacologist", map[string]string{"Keywords": "Perscription,Medicine,Pharmacologist,Fill order"}, "Pharmacologist", 277, "Perscription,Medicine,Pharmacologist,Fill order"}
	occupationPhlebotomist                               = EnumOccupationItem{occupationPhlebotomistID, "Phlebotomist", map[string]string{"Keywords": "Blood,Test,Shots,Needle,Donate,Phlebotomist"}, "Phlebotomist", 278, "Blood,Test,Shots,Needle,Donate,Phlebotomist"}
	occupationPhotographer                               = EnumOccupationItem{occupationPhotographerID, "Photographer", map[string]string{"Keywords": "Photographer,Camera,Collage,Pictures"}, "Photographer", 279, "Photographer,Camera,Collage,Pictures"}
	occupationPhotographicProcessor                      = EnumOccupationItem{occupationPhotographicProcessorID, "Photographic Processor", map[string]string{"Keywords": "Photographic,Film developer,Photo,Pictures,I hour,Processor"}, "PhotographicProcessor", 280, "Photographic,Film developer,Photo,Pictures,I hour,Processor"}
	occupationPhysicalTherapistAPTAMember                = EnumOccupationItem{occupationPhysicalTherapistAPTAMemberID, "Physical Therapist- APTA Member", map[string]string{"Keywords": "APTA Member,Physical Therapist,Stress"}, "PhysicalTherapistAPTAMember", 281, "APTA Member,Physical Therapist,Stress"}
	occupationPhysicalTherapistNonAPTAMember             = EnumOccupationItem{occupationPhysicalTherapistNonAPTAMemberID, "Physical Therapist- Non APTA Member", map[string]string{"Keywords": "Physical Therapist,Non APTA Member,Stress"}, "PhysicalTherapistNonAPTAMember", 282, "Physical Therapist,Non APTA Member,Stress"}
	occupationPhysicianorDoctor                          = EnumOccupationItem{occupationPhysicianorDoctorID, "Physician or Doctor", map[string]string{"Keywords": "Medical doctor,Physician,Doctor"}, "PhysicianorDoctor", 283, "Medical doctor,Physician,Doctor"}
	occupationPilot                                      = EnumOccupationItem{occupationPilotID, "Pilot", map[string]string{"Keywords": "Pilot,Airplane,Bus,Passenger,Public,Private"}, "Pilot", 284, "Pilot,Airplane,Bus,Passenger,Public,Private"}
	occupationPilotCropBush                              = EnumOccupationItem{occupationPilotCropBushID, "Pilot - Crop, Bush", map[string]string{"Keywords": "Pilot,Duster,Crop,Bush,Farm,Water"}, "PilotCropBush", 285, "Pilot,Duster,Crop,Bush,Farm,Water"}
	occupationPipefitterOtherFitter                      = EnumOccupationItem{occupationPipefitterOtherFitterID, "Pipefitter / Other Fitter", map[string]string{"Keywords": "Other,Pipefitter"}, "PipefitterOtherFitter", 286, "Other,Pipefitter"}
	occupationPlannerAllOther                            = EnumOccupationItem{occupationPlannerAllOtherID, "Planner - All Other", map[string]string{"Keywords": "Wedding,Event,Other,Project,Planner"}, "PlannerAllOther", 287, "Wedding,Event,Other,Project,Planner"}
	occupationPlannerProductionorPrinter                 = EnumOccupationItem{occupationPlannerProductionorPrinterID, "Planner - Production or Printer", map[string]string{"Keywords": "Production,Printer,Planner"}, "PlannerProductionorPrinter", 288, "Production,Printer,Planner"}
	occupationPlumber                                    = EnumOccupationItem{occupationPlumberID, "Plumber", map[string]string{"Keywords": "Clogg,Plumber,Sink,Toilet"}, "Plumber", 289, "Clogg,Plumber,Sink,Toilet"}
	occupationPodiatrist                                 = EnumOccupationItem{occupationPodiatristID, "Podiatrist", map[string]string{"Keywords": "Leg,Podiatrist,Diagnosis,Foot doctor,Ankle"}, "Podiatrist", 290, "Leg,Podiatrist,Diagnosis,Foot doctor,Ankle"}
	occupationPolitician                                 = EnumOccupationItem{occupationPoliticianID, "Politician", map[string]string{"Keywords": "Executive,Legaslative,Senate,Politician,Congress,Judicial,Mayor,Board,President,Chair"}, "Politician", 291, "Executive,Legaslative,Senate,Politician,Congress,Judicial,Mayor,Board,President,Chair"}
	occupationPoolServiceCleaner                         = EnumOccupationItem{occupationPoolServiceCleanerID, "Pool Service / Cleaner", map[string]string{"Keywords": "Cleaner,Pool service"}, "PoolServiceCleaner", 292, "Cleaner,Pool service"}
	occupationPostalExecutiveGradesPcesIII               = EnumOccupationItem{occupationPostalExecutiveGradesPcesIIIID, "Postal Executive (Grades Pces I & II)", map[string]string{"Keywords": "Executive,Postal,Pces I & II,Grades"}, "PostalExecutiveGradesPcesIII", 293, "Executive,Postal,Pces I & II,Grades"}
	occupationPostmasterRural                            = EnumOccupationItem{occupationPostmasterRuralID, "Postmaster - Rural", map[string]string{"Keywords": "Country,Postmaster,Rural"}, "PostmasterRural", 294, "Country,Postmaster,Rural"}
	occupationPostmasterUrbanSuburban                    = EnumOccupationItem{occupationPostmasterUrbanSuburbanID, "Postmaster - Urban & Suburban", map[string]string{"Keywords": "Sub,Urban,Postmaster,City"}, "PostmasterUrbanSuburban", 295, "Sub,Urban,Postmaster,City"}
	occupationPresidentBlueCollar50Empl                  = EnumOccupationItem{occupationPresidentBlueCollar50EmplID, "President - Blue Collar -50+ Empl", map[string]string{"Keywords": "Fifty or more,50+ Employees,Blue collar,President"}, "PresidentBlueCollar50Empl", 296, "Fifty or more,50+ Employees,Blue collar,President"}
	occupationPresidentSkilledBlueCollarLessThan50Emp    = EnumOccupationItem{occupationPresidentSkilledBlueCollarLessThan50EmpID, "President - Skilled / Blue Collar Less Than 50 Emp", map[string]string{"Keywords": "Blue Collar,Less than 50 employees,Skilled,President,Fifty or less"}, "PresidentSkilledBlueCollarLessThan50Emp", 297, "Blue Collar,Less than 50 employees,Skilled,President,Fifty or less"}
	occupationPresidentWhiteCollar                       = EnumOccupationItem{occupationPresidentWhiteCollarID, "President - White Collar", map[string]string{"Keywords": "White collar,CFO,CEO,President,Corperate,CIO"}, "PresidentWhiteCollar", 298, "White collar,CFO,CEO,President,Corperate,CIO"}
	occupationPrincipalorAssistantPrincipal              = EnumOccupationItem{occupationPrincipalorAssistantPrincipalID, "Principal or Assistant Principal", map[string]string{"Keywords": "Principal,Assistant"}, "PrincipalorAssistantPrincipal", 299, "Principal,Assistant"}
	occupationPrinter                                    = EnumOccupationItem{occupationPrinterID, "Printer", map[string]string{"Keywords": "Copier,Fax,Mailroom,Printer"}, "Printer", 300, "Copier,Fax,Mailroom,Printer"}
	occupationProducer                                   = EnumOccupationItem{occupationProducerID, "Producer", map[string]string{"Keywords": "Music,Producer,Film,Director"}, "Producer", 301, "Music,Producer,Film,Director"}
	occupationProfessor                                  = EnumOccupationItem{occupationProfessorID, "Professor", map[string]string{"Keywords": "Teacher,Professor,Higher education,College"}, "Professor", 302, "Teacher,Professor,Higher education,College"}
	occupationProgramManagementExpert                    = EnumOccupationItem{occupationProgramManagementExpertID, "Program Management Expert", map[string]string{"Keywords": "Expert,Program management,Consultant,Contract"}, "ProgramManagementExpert", 303, "Expert,Program management,Consultant,Contract"}
	occupationProofreader                                = EnumOccupationItem{occupationProofreaderID, "Proofreader", map[string]string{"Keywords": "Editor,Tutor,Proofreader"}, "Proofreader", 304, "Editor,Tutor,Proofreader"}
	occupationPsychiatrist                               = EnumOccupationItem{occupationPsychiatristID, "Psychiatrist", map[string]string{"Keywords": "Counselor,Psychiatrist,Doctor"}, "Psychiatrist", 305, "Counselor,Psychiatrist,Doctor"}
	occupationPsychologist                               = EnumOccupationItem{occupationPsychologistID, "Psychologist", map[string]string{"Keywords": "Psychologist,Research"}, "Psychologist", 306, "Psychologist,Research"}
	occupationPublicRelations                            = EnumOccupationItem{occupationPublicRelationsID, "Public Relations", map[string]string{"Keywords": "VA Blood services,Public Relations,PR,Recruiter"}, "PublicRelations", 307, "VA Blood services,Public Relations,PR,Recruiter"}
	occupationPublisher                                  = EnumOccupationItem{occupationPublisherID, "Publisher", map[string]string{"Keywords": "Publisher,Copywrighter,Reasearcher"}, "Publisher", 308, "Publisher,Copywrighter,Reasearcher"}
	occupationQualityControlManufacturing                = EnumOccupationItem{occupationQualityControlManufacturingID, "Quality Control - Manufacturing", map[string]string{"Keywords": "Assurance,QA,Analyst,Manufacturing,Quality Control"}, "QualityControlManufacturing", 309, "Assurance,QA,Analyst,Manufacturing,Quality Control"}
	occupationQualityControlProfessional                 = EnumOccupationItem{occupationQualityControlProfessionalID, "Quality Control - Professional", map[string]string{"Keywords": "Professional,Assurance,QA,Analyst,Quality Control"}, "QualityControlProfessional", 310, "Professional,Assurance,QA,Analyst,Quality Control"}
	occupationRadiologist                                = EnumOccupationItem{occupationRadiologistID, "Radiologist", map[string]string{"Keywords": "Bones,Scan,Cats,Radiologist,X ray"}, "Radiologist", 311, "Bones,Scan,Cats,Radiologist,X ray"}
	occupationRanchHelperCowboy                          = EnumOccupationItem{occupationRanchHelperCowboyID, "Ranch Helper / Cowboy", map[string]string{"Keywords": "Farm hand,Ranch,Care taker,Detassel,Helper,Cowboy"}, "RanchHelperCowboy", 312, "Farm hand,Ranch,Care taker,Detassel,Helper,Cowboy"}
	occupationRecruiter                                  = EnumOccupationItem{occupationRecruiterID, "Recruiter", map[string]string{"Keywords": "Athletic,Human Resources,Corperate,Recruiter,College"}, "Recruiter", 313, "Athletic,Human Resources,Corperate,Recruiter,College"}
	occupationRegistrar                                  = EnumOccupationItem{occupationRegistrarID, "Registrar", map[string]string{"Keywords": "Archive,Registrar,Records keeper"}, "Registrar", 314, "Archive,Registrar,Records keeper"}
	occupationReligiousClergyOrdainedorLicensed          = EnumOccupationItem{occupationReligiousClergyOrdainedorLicensedID, "Religious - Clergy (Ordained or Licensed)", map[string]string{"Keywords": "Preacher,Licensed,Clergy,Pastor,Religious,Minister,Ordained,Priest"}, "ReligiousClergyOrdainedorLicensed", 315, "Preacher,Licensed,Clergy,Pastor,Religious,Minister,Ordained,Priest"}
	occupationReligiousLaypersonNonClergy                = EnumOccupationItem{occupationReligiousLaypersonNonClergyID, "Religious - Layperson (Non-Clergy)", map[string]string{"Keywords": "Preacher,Licensed,Pastor,Layperson,Religious,Ordained,Minister,Priest"}, "ReligiousLaypersonNonClergy", 316, "Preacher,Licensed,Pastor,Layperson,Religious,Ordained,Minister,Priest"}
	occupationRepairServiceInstallACHeating              = EnumOccupationItem{occupationRepairServiceInstallACHeatingID, "Repair / Service / Install - AC / Heating", map[string]string{"Keywords": "AC,Service,Install,Heating,HVAC,Air conditioning,Repair"}, "RepairServiceInstallACHeating", 317, "AC,Service,Install,Heating,HVAC,Air conditioning,Repair"}
	occupationRepairServiceInstallAllOther               = EnumOccupationItem{occupationRepairServiceInstallAllOtherID, "Repair / Service / Install - All Other", map[string]string{"Keywords": "Verizon,Service,Fios,Comcast,Cable,Other,Shed installer,Internet,Repair"}, "RepairServiceInstallAllOther", 318, "Verizon,Service,Fios,Comcast,Cable,Other,Shed installer,Internet,Repair"}
	occupationRepairServiceInstallJewelryWatchmaker      = EnumOccupationItem{occupationRepairServiceInstallJewelryWatchmakerID, "Repair / Service / Install - Jewelry & Watchmaker", map[string]string{"Keywords": "Service,Watchmaker,Install,Jewelry,Repair"}, "RepairServiceInstallJewelryWatchmaker", 319, "Service,Watchmaker,Install,Jewelry,Repair"}
	occupationRepairServiceInstallLine                   = EnumOccupationItem{occupationRepairServiceInstallLineID, "Repair / Service / Install - Line", map[string]string{"Keywords": "Service,Install,Line,Repair"}, "RepairServiceInstallLine", 320, "Service,Install,Line,Repair"}
	occupationRepairServiceInstallTrained                = EnumOccupationItem{occupationRepairServiceInstallTrainedID, "Repair / Service / Install - Trained", map[string]string{"Keywords": "Verizon,Service,Fios,Comcast,Install,Trained,Cable,Internet,Repair"}, "RepairServiceInstallTrained", 321, "Verizon,Service,Fios,Comcast,Install,Trained,Cable,Internet,Repair"}
	occupationReporter                                   = EnumOccupationItem{occupationReporterID, "Reporter", map[string]string{"Keywords": "Newspaper,Television,Reporter,Court,News"}, "Reporter", 322, "Newspaper,Television,Reporter,Court,News"}
	occupationResearcherAllOther                         = EnumOccupationItem{occupationResearcherAllOtherID, "Researcher - All Other", map[string]string{"Keywords": "Other,Researcher"}, "ResearcherAllOther", 323, "Other,Researcher"}
	occupationRespiratoryTherapist                       = EnumOccupationItem{occupationRespiratoryTherapistID, "Respiratory Therapist", map[string]string{"Keywords": "Therapist,Respiratory"}, "RespiratoryTherapist", 324, "Therapist,Respiratory"}
	occupationRoutemanRoutewoman                         = EnumOccupationItem{occupationRoutemanRoutewomanID, "Routeman / Routewoman", map[string]string{"Keywords": "Man,Sales,Woman,Deliveries,Route"}, "RoutemanRoutewoman", 325, "Man,Sales,Woman,Deliveries,Route"}
	occupationSalespersonAllOther                        = EnumOccupationItem{occupationSalespersonAllOtherID, "Salesperson - All Other", map[string]string{"Keywords": "Other,Salesperson"}, "SalespersonAllOther", 326, "Other,Salesperson"}
	occupationSalespersonCar                             = EnumOccupationItem{occupationSalespersonCarID, "Salesperson - Car", map[string]string{"Keywords": "Car,Salesperson"}, "SalespersonCar", 327, "Car,Salesperson"}
	occupationSalespersonDoorToDoor                      = EnumOccupationItem{occupationSalespersonDoorToDoorID, "Salesperson - Door To Door", map[string]string{"Keywords": "Door to door,Salesperson"}, "SalespersonDoorToDoor", 328, "Door to door,Salesperson"}
	occupationSalespersonHighTech                        = EnumOccupationItem{occupationSalespersonHighTechID, "Salesperson - High Tech", map[string]string{"Keywords": "High Tech,Salesperson"}, "SalespersonHighTech", 329, "High Tech,Salesperson"}
	occupationSalespersonNonHighTech                     = EnumOccupationItem{occupationSalespersonNonHighTechID, "Salesperson - Non High Tech", map[string]string{"Keywords": "Non high tech,Salesperson"}, "SalespersonNonHighTech", 330, "Non high tech,Salesperson"}
	occupationSalespersonPharmaceutical                  = EnumOccupationItem{occupationSalespersonPharmaceuticalID, "Salesperson - Pharmaceutical", map[string]string{"Keywords": "Pharmaceutical,Salesperson"}, "SalespersonPharmaceutical", 331, "Pharmaceutical,Salesperson"}
	occupationSalespersonRetail                          = EnumOccupationItem{occupationSalespersonRetailID, "Salesperson - Retail", map[string]string{"Keywords": "Retail,Salesperson"}, "SalespersonRetail", 332, "Retail,Salesperson"}
	occupationSalespersonWholesale                       = EnumOccupationItem{occupationSalespersonWholesaleID, "Salesperson - Wholesale", map[string]string{"Keywords": "Whole,Salesperson"}, "SalespersonWholesale", 333, "Whole,Salesperson"}
	occupationSanitarian                                 = EnumOccupationItem{occupationSanitarianID, "Sanitarian", map[string]string{"Keywords": "Sanitation,Disposal,Garbage,Waste"}, "Sanitarian", 334, "Sanitation,Disposal,Garbage,Waste"}
	occupationScheduler                                  = EnumOccupationItem{occupationSchedulerID, "Scheduler", map[string]string{"Keywords": "Scheduler,Assistant,Corperate,Calendar,Planner"}, "Scheduler", 335, "Scheduler,Assistant,Corperate,Calendar,Planner"}
	occupationScientistAllOther                          = EnumOccupationItem{occupationScientistAllOtherID, "Scientist - All Other", map[string]string{"Keywords": "Study,Chemist,Scientist,Biologist,Other,Research,Chemical,Biology"}, "ScientistAllOther", 336, "Study,Chemist,Scientist,Biologist,Other,Research,Chemical,Biology"}
	occupationSeamstressTailor                           = EnumOccupationItem{occupationSeamstressTailorID, "Seamstress / Tailor", map[string]string{"Keywords": "Patch,Seamstress,Tailor,Stuffed animals,Sew,Clothing"}, "SeamstressTailor", 337, "Patch,Seamstress,Tailor,Stuffed animals,Sew,Clothing"}
	occupationSecurityScreener                           = EnumOccupationItem{occupationSecurityScreenerID, "Security Screener", map[string]string{"Keywords": "Security Screener,Scanner,TSA"}, "SecurityScreener", 338, "Security Screener,Scanner,TSA"}
	occupationShoeShinerRepairman                        = EnumOccupationItem{occupationShoeShinerRepairmanID, "Shoe Shiner / Repairman", map[string]string{"Keywords": "Repairman,Shoes,Street vendor,Shoe Shiner"}, "ShoeShinerRepairman", 339, "Repairman,Shoes,Street vendor,Shoe Shiner"}
	occupationSingerSongwriter                           = EnumOccupationItem{occupationSingerSongwriterID, "Singer / Songwriter", map[string]string{"Keywords": "Jazz,Pop,Rap,Rock,Singer,Songwriter,Country,MC,Musician"}, "SingerSongwriter", 340, "Jazz,Pop,Rap,Rock,Singer,Songwriter,Country,MC,Musician"}
	occupationStaffingSpecialist                         = EnumOccupationItem{occupationStaffingSpecialistID, "Staffing Specialist", map[string]string{"Keywords": "Human resources,Staffing specialist,Recruiter"}, "StaffingSpecialist", 341, "Human resources,Staffing specialist,Recruiter"}
	occupationStateExaminer                              = EnumOccupationItem{occupationStateExaminerID, "State Examiner", map[string]string{"Keywords": "Audit,Government,State Examiner,Market conduct"}, "StateExaminer", 342, "Audit,Government,State Examiner,Market conduct"}
	occupationSuperintendentAllOther                     = EnumOccupationItem{occupationSuperintendentAllOtherID, "Superintendent - All Other", map[string]string{"Keywords": "Construction,Other,Supervisor,Superintendent"}, "SuperintendentAllOther", 343, "Construction,Other,Supervisor,Superintendent"}
	occupationSuperintendentDriller                      = EnumOccupationItem{occupationSuperintendentDrillerID, "Superintendent - Driller", map[string]string{"Keywords": "Oil,Driller,Superintendent"}, "SuperintendentDriller", 344, "Oil,Driller,Superintendent"}
	occupationSuperintendentSchool                       = EnumOccupationItem{occupationSuperintendentSchoolID, "Superintendent - School", map[string]string{"Keywords": "District,School,Superintendent"}, "SuperintendentSchool", 345, "District,School,Superintendent"}
	occupationSuperintendentorSupervisorBuildingMaint    = EnumOccupationItem{occupationSuperintendentorSupervisorBuildingMaintID, "Superintendent or Supervisor-Building & Maint.", map[string]string{"Keywords": "Maintenance,Building,Supervisor,Superintendent"}, "SuperintendentorSupervisorBuildingMaint", 346, "Maintenance,Building,Supervisor,Superintendent"}
	occupationSupervisorAccounting                       = EnumOccupationItem{occupationSupervisorAccountingID, "Supervisor - Accounting", map[string]string{"Keywords": "Records,Supervisor,Tax,Accounting"}, "SupervisorAccounting", 347, "Records,Supervisor,Tax,Accounting"}
	occupationSupervisorAllOtherDegreed                  = EnumOccupationItem{occupationSupervisorAllOtherDegreedID, "Supervisor - All Other (Degreed)", map[string]string{"Keywords": "Travel,Agriculture,Facturing,Sales,Trades,Executive,Degreed,Management,Ranch,Construction,Food,Other,Hotel,Advertising,Owners,Electric,Oil,Restaurant,Transportation,Farm,Supervisor,Retail"}, "SupervisorAllOtherDegreed", 348, "Travel,Agriculture,Facturing,Sales,Trades,Executive,Degreed,Management,Ranch,Construction,Food,Other,Hotel,Advertising,Owners,Electric,Oil,Restaurant,Transportation,Farm,Supervisor,Retail"}
	occupationSupervisorDataSystems                      = EnumOccupationItem{occupationSupervisorDataSystemsID, "Supervisor - Data Systems", map[string]string{"Keywords": "DBA,Supervisor,Data systems,IT,Developer"}, "SupervisorDataSystems", 349, "DBA,Supervisor,Data systems,IT,Developer"}
	occupationSupervisorHumanResourcePersonnel           = EnumOccupationItem{occupationSupervisorHumanResourcePersonnelID, "Supervisor - Human Resource / Personnel", map[string]string{"Keywords": "Personnel,Human resource,Supervisor"}, "SupervisorHumanResourcePersonnel", 350, "Personnel,Human resource,Supervisor"}
	occupationSupervisorOffice                           = EnumOccupationItem{occupationSupervisorOfficeID, "Supervisor - Office", map[string]string{"Keywords": "Supervisor,Office"}, "SupervisorOffice", 351, "Supervisor,Office"}
	occupationSupervisorOperations                       = EnumOccupationItem{occupationSupervisorOperationsID, "Supervisor - Operations", map[string]string{"Keywords": "Operations,Supervisor"}, "SupervisorOperations", 352, "Operations,Supervisor"}
	occupationSupervisorOtherDegreed                     = EnumOccupationItem{occupationSupervisorOtherDegreedID, "Supervisor - Other (Degreed)", map[string]string{"Keywords": "Financial,Degreed,Other,Supervisor,Office,Accounting,Insurance"}, "SupervisorOtherDegreed", 353, "Financial,Degreed,Other,Supervisor,Office,Accounting,Insurance"}
	occupationSupervisorPostal                           = EnumOccupationItem{occupationSupervisorPostalID, "Supervisor - Postal", map[string]string{"Keywords": "Postal,Supervisor"}, "SupervisorPostal", 354, "Postal,Supervisor"}
	occupationSupervisorProduction                       = EnumOccupationItem{occupationSupervisorProductionID, "Supervisor - Production", map[string]string{"Keywords": "Production,Supervisor,Line,Lead"}, "SupervisorProduction", 355, "Production,Supervisor,Line,Lead"}
	occupationSupervisorRestaurantNonFastFood            = EnumOccupationItem{occupationSupervisorRestaurantNonFastFoodID, "Supervisor - Restaurant Non Fast Food", map[string]string{"Keywords": "Guest,Restaurant,Supervisor,Lead"}, "SupervisorRestaurantNonFastFood", 356, "Guest,Restaurant,Supervisor,Lead"}
	occupationSurgeon                                    = EnumOccupationItem{occupationSurgeonID, "Surgeon", map[string]string{"Keywords": "Medical doctor,Surgeon"}, "Surgeon", 357, "Medical doctor,Surgeon"}
	occupationSurveyorLicensed                           = EnumOccupationItem{occupationSurveyorLicensedID, "Surveyor (Licensed)", map[string]string{"Keywords": "Construction,Surveyor,Licensed"}, "SurveyorLicensed", 358, "Construction,Surveyor,Licensed"}
	occupationSurveyorNonLicensed                        = EnumOccupationItem{occupationSurveyorNonLicensedID, "Surveyor (Non-Licensed)", map[string]string{"Keywords": "Construction,Surveyor"}, "SurveyorNonLicensed", 359, "Construction,Surveyor"}
	occupationTaxExaminerNotClerical                     = EnumOccupationItem{occupationTaxExaminerNotClericalID, "Tax Examiner (Not Clerical)", map[string]string{"Keywords": "Examiner,Tax"}, "TaxExaminerNotClerical", 360, "Examiner,Tax"}
	occupationTaxPreparerNotAccountant                   = EnumOccupationItem{occupationTaxPreparerNotAccountantID, "Tax Preparer (Not Accountant)", map[string]string{"Keywords": "Taxes,File,Preparer"}, "TaxPreparerNotAccountant", 361, "Taxes,File,Preparer"}
	occupationTaxidermist                                = EnumOccupationItem{occupationTaxidermistID, "Taxidermist", map[string]string{"Keywords": "Pets,Stuffed animals,Mount,Fish,Taxidermist"}, "Taxidermist", 362, "Pets,Stuffed animals,Mount,Fish,Taxidermist"}
	occupationTeachersorCoachesorInstructors             = EnumOccupationItem{occupationTeachersorCoachesorInstructorsID, "Teachers or Coaches or Instructors", map[string]string{"Keywords": "Base,Gym,Swimming,Teacher,Substitute,Lacrosse,Assistant,Math,Soccer,Basket,Volley,Science,Hockey,Coach,Instruct,Tennis,Social studies,Foot"}, "TeachersorCoachesorInstructors", 363, "Base,Gym,Swimming,Teacher,Substitute,Lacrosse,Assistant,Math,Soccer,Basket,Volley,Science,Hockey,Coach,Instruct,Tennis,Social studies,Foot"}
	occupationTechnicianAllOther                         = EnumOccupationItem{occupationTechnicianAllOtherID, "Technician - All Other", map[string]string{"Keywords": "Cataloger,Technician,Other,Pharmacy tech,Special event,Maintenance worker,Set up,Maintenance"}, "TechnicianAllOther", 364, "Cataloger,Technician,Other,Pharmacy tech,Special event,Maintenance worker,Set up,Maintenance"}
	occupationTechnicianElectricalorElectronic           = EnumOccupationItem{occupationTechnicianElectricalorElectronicID, "Technician - Electrical or Electronic", map[string]string{"Keywords": "Electrical,Technician"}, "TechnicianElectricalorElectronic", 365, "Electrical,Technician"}
	occupationTechnicianFood                             = EnumOccupationItem{occupationTechnicianFoodID, "Technician - Food", map[string]string{"Keywords": "Food,Technician"}, "TechnicianFood", 366, "Food,Technician"}
	occupationTechnicianInstrumentation                  = EnumOccupationItem{occupationTechnicianInstrumentationID, "Technician - Instrumentation", map[string]string{"Keywords": "Instrumentation,Technician"}, "TechnicianInstrumentation", 367, "Instrumentation,Technician"}
	occupationTechnicianLab                              = EnumOccupationItem{occupationTechnicianLabID, "Technician - Lab", map[string]string{"Keywords": "Technician,Medical,Lab"}, "TechnicianLab", 368, "Technician,Medical,Lab"}
	occupationTechnicianMedical                          = EnumOccupationItem{occupationTechnicianMedicalID, "Technician - Medical", map[string]string{"Keywords": "Technician,Medical"}, "TechnicianMedical", 369, "Technician,Medical"}
	occupationTechnicianRadiological                     = EnumOccupationItem{occupationTechnicianRadiologicalID, "Technician - Radiological", map[string]string{"Keywords": "Bones,Technician,Radiological,X ray"}, "TechnicianRadiological", 370, "Bones,Technician,Radiological,X ray"}
	occupationTechnicianScience                          = EnumOccupationItem{occupationTechnicianScienceID, "Technician - Science", map[string]string{"Keywords": "Technician,Science"}, "TechnicianScience", 371, "Technician,Science"}
	occupationTechnicianTesting                          = EnumOccupationItem{occupationTechnicianTestingID, "Technician - Testing", map[string]string{"Keywords": "Technician,Testing"}, "TechnicianTesting", 372, "Technician,Testing"}
	occupationTechnicianUltrasound                       = EnumOccupationItem{occupationTechnicianUltrasoundID, "Technician - Ultrasound", map[string]string{"Keywords": "Technician,Medical,Ultrasound"}, "TechnicianUltrasound", 373, "Technician,Medical,Ultrasound"}
	occupationTechnicianXRay                             = EnumOccupationItem{occupationTechnicianXRayID, "Technician - X Ray", map[string]string{"Keywords": "Technician,Medical,X Ray"}, "TechnicianXRay", 374, "Technician,Medical,X Ray"}
	occupationTechnicianorAssistantEngineering           = EnumOccupationItem{occupationTechnicianorAssistantEngineeringID, "Technician or Assistant - Engineering", map[string]string{"Keywords": "Engineering,Technician,Assistant"}, "TechnicianorAssistantEngineering", 375, "Engineering,Technician,Assistant"}
	occupationTelemarketer                               = EnumOccupationItem{occupationTelemarketerID, "Telemarketer", map[string]string{"Keywords": "Telemarketer,Consumer direct"}, "Telemarketer", 376, "Telemarketer,Consumer direct"}
	occupationTherapist                                  = EnumOccupationItem{occupationTherapistID, "Therapist", map[string]string{"Keywords": "Counselor,Therapist,Family,Marriage"}, "Therapist", 377, "Counselor,Therapist,Family,Marriage"}
	occupationTrainerAerobicsFitness                     = EnumOccupationItem{occupationTrainerAerobicsFitnessID, "Trainer - Aerobics/Fitness", map[string]string{"Keywords": "Fitness,Physical,Aerobics,Trainer"}, "TrainerAerobicsFitness", 378, "Fitness,Physical,Aerobics,Trainer"}
	occupationTrainerAthleticNataMember                  = EnumOccupationItem{occupationTrainerAthleticNataMemberID, "Trainer - Athletic Nata Member", map[string]string{"Keywords": "Physical,Fitness,Athletic,Nata member,Trainer"}, "TrainerAthleticNataMember", 379, "Physical,Fitness,Athletic,Nata member,Trainer"}
	occupationTrainerAthleticNonNataMember               = EnumOccupationItem{occupationTrainerAthleticNonNataMemberID, "Trainer - Athletic Non-Nata Member", map[string]string{"Keywords": "Physical,Fitness,Athletic,Trainer,Non nata"}, "TrainerAthleticNonNataMember", 380, "Physical,Fitness,Athletic,Trainer,Non nata"}
	occupationTrainerCaretakerAnimal                     = EnumOccupationItem{occupationTrainerCaretakerAnimalID, "Trainer / Caretaker - Animal", map[string]string{"Keywords": "Dog walk,Pet spa,Sitter,Trainer,Caretaker,Groomer,Animal,Boarder,Obedient"}, "TrainerCaretakerAnimal", 381, "Dog walk,Pet spa,Sitter,Trainer,Caretaker,Groomer,Animal,Boarder,Obedient"}
	occupationTranslatororInterpreter                    = EnumOccupationItem{occupationTranslatororInterpreterID, "Translator or Interpreter", map[string]string{"Keywords": "Communication,Language,Interpreter,Linguist,Translator"}, "TranslatororInterpreter", 382, "Communication,Language,Interpreter,Linguist,Translator"}
	occupationTreasurer                                  = EnumOccupationItem{occupationTreasurerID, "Treasurer", map[string]string{"Keywords": "Money,Treasurer,Currency"}, "Treasurer", 383, "Money,Treasurer,Currency"}
	occupationTutor                                      = EnumOccupationItem{occupationTutorID, "Tutor", map[string]string{"Keywords": "Teacher,Tutor"}, "Tutor", 384, "Teacher,Tutor"}
	occupationUnderwriterInsurance                       = EnumOccupationItem{occupationUnderwriterInsuranceID, "Underwriter - Insurance", map[string]string{"Keywords": "Risk,Underwriter,Insurance"}, "UnderwriterInsurance", 385, "Risk,Underwriter,Insurance"}
	occupationVendor                                     = EnumOccupationItem{occupationVendorID, "Vendor", map[string]string{"Keywords": "Street,Machine,Drink,Vendor,Food,Game,Market"}, "Vendor", 386, "Street,Machine,Drink,Vendor,Food,Game,Market"}
	occupationVeterinarian                               = EnumOccupationItem{occupationVeterinarianID, "Veterinarian", map[string]string{"Keywords": "Animal doctor,Veterinarian"}, "Veterinarian", 387, "Animal doctor,Veterinarian"}
	occupationVicePresBusiness                           = EnumOccupationItem{occupationVicePresBusinessID, "Vice Pres - Business", map[string]string{"Keywords": "Business,Vice president"}, "VicePresBusiness", 388, "Business,Vice president"}
	occupationWaiterWaitress                             = EnumOccupationItem{occupationWaiterWaitressID, "Waiter / Waitress", map[string]string{"Keywords": "Restaurant,Waiter,Food,Server"}, "WaiterWaitress", 389, "Restaurant,Waiter,Food,Server"}
	occupationWardenAllOther                             = EnumOccupationItem{occupationWardenAllOtherID, "Warden - All Other", map[string]string{"Keywords": "Other,Warden,Prison"}, "WardenAllOther", 390, "Other,Warden,Prison"}
	occupationWardenGame                                 = EnumOccupationItem{occupationWardenGameID, "Warden - Game", map[string]string{"Keywords": "Fishing,Game,Warden,Hunting"}, "WardenGame", 391, "Fishing,Game,Warden,Hunting"}
	occupationWorkerMetalNotSteel                        = EnumOccupationItem{occupationWorkerMetalNotSteelID, "Worker - Metal (Not Steel)", map[string]string{"Keywords": "Worker,Welder,Metal"}, "WorkerMetalNotSteel", 392, "Worker,Welder,Metal"}
	occupationWorkerRailroad                             = EnumOccupationItem{occupationWorkerRailroadID, "Worker - Railroad", map[string]string{"Keywords": "Worker,Railroad"}, "WorkerRailroad", 393, "Worker,Railroad"}
	occupationWorkerSocialCase                           = EnumOccupationItem{occupationWorkerSocialCaseID, "Worker - Social / Case", map[string]string{"Keywords": "Worker,Case,Social"}, "WorkerSocialCase", 394, "Worker,Case,Social"}
	occupationWriterAllOther                             = EnumOccupationItem{occupationWriterAllOtherID, "Writer - All Other", map[string]string{"Keywords": "Book,Newspaper,Writer,Magazine,Other,Internet,Blog"}, "WriterAllOther", 395, "Book,Newspaper,Writer,Magazine,Other,Internet,Blog"}
	occupationWriterCommercial                           = EnumOccupationItem{occupationWriterCommercialID, "Writer - Commercial", map[string]string{"Keywords": "Book,Newspaper,Writer,Magazine,Other,Internet,Blog"}, "WriterCommercial", 396, "Book,Newspaper,Writer,Magazine,Other,Internet,Blog"}
	occupationE1                                         = EnumOccupationItem{occupationE1ID, "E-1", map[string]string{"Keywords": ""}, "E1", 397, ""}
	occupationE2                                         = EnumOccupationItem{occupationE2ID, "E-2", map[string]string{"Keywords": ""}, "E2", 398, ""}
	occupationE3                                         = EnumOccupationItem{occupationE3ID, "E-3", map[string]string{"Keywords": ""}, "E3", 399, ""}
	occupationE4                                         = EnumOccupationItem{occupationE4ID, "E-4", map[string]string{"Keywords": ""}, "E4", 400, ""}
	occupationE4P                                        = EnumOccupationItem{occupationE4PID, "E-4P", map[string]string{"Keywords": ""}, "E4P", 401, ""}
	occupationE5                                         = EnumOccupationItem{occupationE5ID, "E-5", map[string]string{"Keywords": ""}, "E5", 402, ""}
	occupationE5P                                        = EnumOccupationItem{occupationE5PID, "E-5P", map[string]string{"Keywords": ""}, "E5P", 403, ""}
	occupationE6                                         = EnumOccupationItem{occupationE6ID, "E-6", map[string]string{"Keywords": ""}, "E6", 404, ""}
	occupationE6P                                        = EnumOccupationItem{occupationE6PID, "E-6P", map[string]string{"Keywords": ""}, "E6P", 405, ""}
	occupationE7                                         = EnumOccupationItem{occupationE7ID, "E-7", map[string]string{"Keywords": ""}, "E7", 406, ""}
	occupationE8                                         = EnumOccupationItem{occupationE8ID, "E-8", map[string]string{"Keywords": ""}, "E8", 407, ""}
	occupationE9                                         = EnumOccupationItem{occupationE9ID, "E-9", map[string]string{"Keywords": ""}, "E9", 408, ""}
	occupationO1                                         = EnumOccupationItem{occupationO1ID, "O-1", map[string]string{"Keywords": ""}, "O1", 409, ""}
	occupationO10                                        = EnumOccupationItem{occupationO10ID, "O-10", map[string]string{"Keywords": ""}, "O10", 410, ""}
	occupationO2                                         = EnumOccupationItem{occupationO2ID, "O-2", map[string]string{"Keywords": ""}, "O2", 411, ""}
	occupationO3                                         = EnumOccupationItem{occupationO3ID, "O-3", map[string]string{"Keywords": ""}, "O3", 412, ""}
	occupationO4                                         = EnumOccupationItem{occupationO4ID, "O-4", map[string]string{"Keywords": ""}, "O4", 413, ""}
	occupationO5                                         = EnumOccupationItem{occupationO5ID, "O-5", map[string]string{"Keywords": ""}, "O5", 414, ""}
	occupationO6                                         = EnumOccupationItem{occupationO6ID, "O-6", map[string]string{"Keywords": ""}, "O6", 415, ""}
	occupationO7                                         = EnumOccupationItem{occupationO7ID, "O-7", map[string]string{"Keywords": ""}, "O7", 416, ""}
	occupationO8                                         = EnumOccupationItem{occupationO8ID, "O-8", map[string]string{"Keywords": ""}, "O8", 417, ""}
	occupationO9                                         = EnumOccupationItem{occupationO9ID, "O-9", map[string]string{"Keywords": ""}, "O9", 418, ""}
	occupationW1                                         = EnumOccupationItem{occupationW1ID, "W-1", map[string]string{"Keywords": ""}, "W1", 419, ""}
	occupationW2                                         = EnumOccupationItem{occupationW2ID, "W-2", map[string]string{"Keywords": ""}, "W2", 420, ""}
	occupationW3                                         = EnumOccupationItem{occupationW3ID, "W-3", map[string]string{"Keywords": ""}, "W3", 421, ""}
	occupationW4                                         = EnumOccupationItem{occupationW4ID, "W-4", map[string]string{"Keywords": ""}, "W4", 422, ""}
	occupationW5                                         = EnumOccupationItem{occupationW5ID, "W-5", map[string]string{"Keywords": ""}, "W5", 423, ""}
	occupationHomemaker                                  = EnumOccupationItem{occupationHomemakerID, "Homemaker", map[string]string{"Keywords": ""}, "Homemaker", 424, ""}
	occupationUnemployed                                 = EnumOccupationItem{occupationUnemployedID, "Unemployed", map[string]string{"Keywords": ""}, "Unemployed", 425, ""}
	occupationGraduateSchool                             = EnumOccupationItem{occupationGraduateSchoolID, "Graduate School", map[string]string{"Keywords": ""}, "GraduateSchool", 426, ""}
	occupationHighSchool                                 = EnumOccupationItem{occupationHighSchoolID, "High School", map[string]string{"Keywords": ""}, "HighSchool", 427, ""}
	occupationTradeSchoolorAssociateDegree               = EnumOccupationItem{occupationTradeSchoolorAssociateDegreeID, "Trade School or Associate Degree", map[string]string{"Keywords": ""}, "TradeSchoolorAssociateDegree", 428, ""}
	occupationUndergraduate4yeardegree                   = EnumOccupationItem{occupationUndergraduate4yeardegreeID, "Undergraduate (4 year degree)", map[string]string{"Keywords": ""}, "Undergraduate4yeardegree", 429, ""}
	occupationDisabled                                   = EnumOccupationItem{occupationDisabledID, "Disabled", map[string]string{"Keywords": ""}, "Disabled", 430, ""}
)

// EnumOccupation is a collection of Occupation items
type EnumOccupation struct {
	Description string
	Items       []*EnumOccupationItem
	Name        string

	AccountExecutive                           *EnumOccupationItem
	AccountantorCPA                            *EnumOccupationItem
	ActorActress                               *EnumOccupationItem
	Actuary                                    *EnumOccupationItem
	Acupuncturist                              *EnumOccupationItem
	Adjudicator                                *EnumOccupationItem
	AdministrativeAssistant                    *EnumOccupationItem
	AdministrativeClerkorReceptionist          *EnumOccupationItem
	Administrator                              *EnumOccupationItem
	AdministratorHealthorHospital              *EnumOccupationItem
	AdministratorPublic                        *EnumOccupationItem
	AdvisorFinancialorCredit                   *EnumOccupationItem
	AgentAdvertising                           *EnumOccupationItem
	AgentBorderPatrol                          *EnumOccupationItem
	AgentEmployment                            *EnumOccupationItem
	AgentFBI                                   *EnumOccupationItem
	AgentImportExport                          *EnumOccupationItem
	AgentInsurance                             *EnumOccupationItem
	AgentIRS                                   *EnumOccupationItem
	AgentLeasing                               *EnumOccupationItem
	AgentRealEstate                            *EnumOccupationItem
	AgentSecretService                         *EnumOccupationItem
	AgentTheatrical                            *EnumOccupationItem
	AgentTravelTicket                          *EnumOccupationItem
	AgentTreasury                              *EnumOccupationItem
	AideAllOther                               *EnumOccupationItem
	AirTrafficController                       *EnumOccupationItem
	Ambassador                                 *EnumOccupationItem
	AnalystAllOther                            *EnumOccupationItem
	AnalystEngineering                         *EnumOccupationItem
	AnalystFinancial                           *EnumOccupationItem
	Anesthesiologist                           *EnumOccupationItem
	AntiqueDealer                              *EnumOccupationItem
	AppraiserAllOther                          *EnumOccupationItem
	AppraiserArt                               *EnumOccupationItem
	AppraiserRealEstate                        *EnumOccupationItem
	ApprenticeAllOther                         *EnumOccupationItem
	Architect                                  *EnumOccupationItem
	ArchivistLibraryorMuseum                   *EnumOccupationItem
	ArtistCommercial                           *EnumOccupationItem
	ArtistNoncommercial                        *EnumOccupationItem
	AssistantAllOther                          *EnumOccupationItem
	AssistantEditorial                         *EnumOccupationItem
	AssistantLegislativeorLegal                *EnumOccupationItem
	AssistantMedical                           *EnumOccupationItem
	AssistantPrintingPress                     *EnumOccupationItem
	Athlete                                    *EnumOccupationItem
	AttorneyorLawyer                           *EnumOccupationItem
	Auctioneer                                 *EnumOccupationItem
	Audiologist                                *EnumOccupationItem
	AuditorFinancial                           *EnumOccupationItem
	AuditorNonFinancial                        *EnumOccupationItem
	Baker                                      *EnumOccupationItem
	BankTeller                                 *EnumOccupationItem
	Bartender                                  *EnumOccupationItem
	BoatCaptain                                *EnumOccupationItem
	Boatman                                    *EnumOccupationItem
	Bondsman                                   *EnumOccupationItem
	Bookbinder                                 *EnumOccupationItem
	Bookkeeper                                 *EnumOccupationItem
	Broadcaster                                *EnumOccupationItem
	BrokerInsurance                            *EnumOccupationItem
	BrokerMortgage                             *EnumOccupationItem
	BrokerPawn                                 *EnumOccupationItem
	BrokerRealEstate                           *EnumOccupationItem
	BrokerStockBonds                           *EnumOccupationItem
	BuyerPurchasingAgent                       *EnumOccupationItem
	Cameraman                                  *EnumOccupationItem
	CarpenterorCabinetmaker                    *EnumOccupationItem
	CashierFood                                *EnumOccupationItem
	CashierOffice                              *EnumOccupationItem
	CashierRetail                              *EnumOccupationItem
	Caterer                                    *EnumOccupationItem
	CensusTaker                                *EnumOccupationItem
	Chiropractor                               *EnumOccupationItem
	ClaimsAdjuster                             *EnumOccupationItem
	ClaimsExaminer                             *EnumOccupationItem
	ClerkAccountingorFinancial                 *EnumOccupationItem
	ClerkAllOther                              *EnumOccupationItem
	CollegeDean                                *EnumOccupationItem
	CommunicationSpecialist                    *EnumOccupationItem
	ComputerProgrammer                         *EnumOccupationItem
	ComputerTechnicalSupportRep                *EnumOccupationItem
	Conductor                                  *EnumOccupationItem
	Conservationist                            *EnumOccupationItem
	ConstructionWorker                         *EnumOccupationItem
	ConsultantorAdvisor                        *EnumOccupationItem
	ContractororDeveloperAllOther              *EnumOccupationItem
	ContractororDeveloperArtisanSkilled        *EnumOccupationItem
	ContractororDeveloperBlueCollar50Employees *EnumOccupationItem
	ContractororDeveloperWhiteCollar           *EnumOccupationItem
	ControllerorComptroller                    *EnumOccupationItem
	ControllerorCoordinatorProduction          *EnumOccupationItem
	CookorChef                                 *EnumOccupationItem
	CoordinatorOffice                          *EnumOccupationItem
	Copywriter                                 *EnumOccupationItem
	Coroner                                    *EnumOccupationItem
	CosmetologistBeautician                    *EnumOccupationItem
	CounselorAllOther                          *EnumOccupationItem
	CounselorEducation                         *EnumOccupationItem
	CounselorFamilyAndChild                    *EnumOccupationItem
	CounselorMentalHealth                      *EnumOccupationItem
	CourtReporter                              *EnumOccupationItem
	Craftsman                                  *EnumOccupationItem
	Criminologist                              *EnumOccupationItem
	Curator                                    *EnumOccupationItem
	CustodianorJanitor                         *EnumOccupationItem
	CustomerServiceRepresentative              *EnumOccupationItem
	DayCareWorker                              *EnumOccupationItem
	DecoratorInterior                          *EnumOccupationItem
	DeliveryPersonorMailCarrier                *EnumOccupationItem
	DentalHygienist                            *EnumOccupationItem
	Dentist                                    *EnumOccupationItem
	DesignerComputerWebsite                    *EnumOccupationItem
	DesignerFloral                             *EnumOccupationItem
	DesignerGraphicorTechnical                 *EnumOccupationItem
	DesignerProfessional                       *EnumOccupationItem
	DesignerWindow                             *EnumOccupationItem
	DietitianNutritionist                      *EnumOccupationItem
	DirectororExecutive                        *EnumOccupationItem
	DiscJockey                                 *EnumOccupationItem
	Dispatcher                                 *EnumOccupationItem
	DogBreeder                                 *EnumOccupationItem
	DrafterorCartographer                      *EnumOccupationItem
	DriverAllOther                             *EnumOccupationItem
	DriverTruck                                *EnumOccupationItem
	Economist                                  *EnumOccupationItem
	EditorAllOther                             *EnumOccupationItem
	EditorFilm                                 *EnumOccupationItem
	Electrician                                *EnumOccupationItem
	Embalmer                                   *EnumOccupationItem
	EngineerAllOther                           *EnumOccupationItem
	EngineerCertifiedNetwork                   *EnumOccupationItem
	EngineerComputerSoftware                   *EnumOccupationItem
	EngineerComputerSystems                    *EnumOccupationItem
	EngineerConstruction                       *EnumOccupationItem
	EngineerElectricalElectronic               *EnumOccupationItem
	EngineerEquipment                          *EnumOccupationItem
	EngineerFacilities                         *EnumOccupationItem
	EngineerFlight                             *EnumOccupationItem
	EngineerMechanical                         *EnumOccupationItem
	EngineerOperating                          *EnumOccupationItem
	EngineerPetroleumorMining                  *EnumOccupationItem
	EngineerSafety                             *EnumOccupationItem
	EngineerSales                              *EnumOccupationItem
	EntertainerPerformer                       *EnumOccupationItem
	Expediter                                  *EnumOccupationItem
	FactoryWorker                              *EnumOccupationItem
	FiremanWomanChiefCaptLt                    *EnumOccupationItem
	FiremanWomanNonChief                       *EnumOccupationItem
	Fisherman                                  *EnumOccupationItem
	FlightAttendant                            *EnumOccupationItem
	FloormenSupervisor                         *EnumOccupationItem
	Florist                                    *EnumOccupationItem
	ForemanForewoman                           *EnumOccupationItem
	Forester                                   *EnumOccupationItem
	Fundraiser                                 *EnumOccupationItem
	Geographer                                 *EnumOccupationItem
	GovtOfficialElected                        *EnumOccupationItem
	Grader                                     *EnumOccupationItem
	GuardEmbassy                               *EnumOccupationItem
	GuardSecurityorPrison                      *EnumOccupationItem
	Gunsmith                                   *EnumOccupationItem
	HairdresserBarber                          *EnumOccupationItem
	Historian                                  *EnumOccupationItem
	HostorHostessRestaurant                    *EnumOccupationItem
	HousekeeperorMaid                          *EnumOccupationItem
	HumanResourcesRepresentative               *EnumOccupationItem
	Illustrator                                *EnumOccupationItem
	InspectorAgricultural                      *EnumOccupationItem
	InspectorAircraftAccessories               *EnumOccupationItem
	InspectorAllOther                          *EnumOccupationItem
	InspectorConstruction                      *EnumOccupationItem
	InspectorPostal                            *EnumOccupationItem
	InspectorWhiteCollar                       *EnumOccupationItem
	InvestigatorPrivate                        *EnumOccupationItem
	InvestmentBanker                           *EnumOccupationItem
	InvestorPrivate                            *EnumOccupationItem
	Journalist                                 *EnumOccupationItem
	Journeyman                                 *EnumOccupationItem
	Judge                                      *EnumOccupationItem
	LaborRelationsWorker                       *EnumOccupationItem
	Landscaper                                 *EnumOccupationItem
	Librarian                                  *EnumOccupationItem
	LifeGuard                                  *EnumOccupationItem
	Linguist                                   *EnumOccupationItem
	Lithographer                               *EnumOccupationItem
	Lobbyist                                   *EnumOccupationItem
	Locksmith                                  *EnumOccupationItem
	Longshoremen                               *EnumOccupationItem
	Machinist                                  *EnumOccupationItem
	ManagerAirport                             *EnumOccupationItem
	ManagerAllOtherDegreed                     *EnumOccupationItem
	ManagerCafeteria                           *EnumOccupationItem
	ManagerCity                                *EnumOccupationItem
	ManagerClericalStaff                       *EnumOccupationItem
	ManagerConvenienceorGasStationStore        *EnumOccupationItem
	ManagerDepartmentStore                     *EnumOccupationItem
	ManagerFinancialorCredit                   *EnumOccupationItem
	ManagerGeneral                             *EnumOccupationItem
	ManagerHealthClub                          *EnumOccupationItem
	ManagerHotel                               *EnumOccupationItem
	ManagerHumanResources                      *EnumOccupationItem
	ManagerMerchandise                         *EnumOccupationItem
	ManagerOffice                              *EnumOccupationItem
	ManagerOperations                          *EnumOccupationItem
	ManagerProduction                          *EnumOccupationItem
	ManagerProfessionalTechStaff               *EnumOccupationItem
	ManagerProject                             *EnumOccupationItem
	ManagerPropertyNonResident                 *EnumOccupationItem
	ManagerPropertyResident                    *EnumOccupationItem
	ManagerRestaurantFastFood                  *EnumOccupationItem
	ManagerRestaurantNonFastFood               *EnumOccupationItem
	ManagerSales                               *EnumOccupationItem
	ManagerSecurityScreener                    *EnumOccupationItem
	ManagerShippingReceiving                   *EnumOccupationItem
	ManagerStage                               *EnumOccupationItem
	ManagerSupermarket                         *EnumOccupationItem
	ManagerorOwnerSandwichShop                 *EnumOccupationItem
	Manicurist                                 *EnumOccupationItem
	MarketingRepresentative                    *EnumOccupationItem
	MarshalFire                                *EnumOccupationItem
	MarshalUSDeputy                            *EnumOccupationItem
	Masseuse                                   *EnumOccupationItem
	Mathematician                              *EnumOccupationItem
	MeatcutterButcher                          *EnumOccupationItem
	MechanicorServicemanAuto                   *EnumOccupationItem
	MechanicorServicemanBoat                   *EnumOccupationItem
	MechanicorServicemanDiesel                 *EnumOccupationItem
	Merchant                                   *EnumOccupationItem
	Millwright                                 *EnumOccupationItem
	Mortician                                  *EnumOccupationItem
	MusicianClassical                          *EnumOccupationItem
	MusicianOther                              *EnumOccupationItem
	NurseCNACertifiedNursingAssistant          *EnumOccupationItem
	NurseLVNorLPN                              *EnumOccupationItem
	NurseRN                                    *EnumOccupationItem
	NursePractitioner                          *EnumOccupationItem
	Oceanographer                              *EnumOccupationItem
	OfficerCorrectional                        *EnumOccupationItem
	OfficerCourt                               *EnumOccupationItem
	OfficerForeignService                      *EnumOccupationItem
	OfficerLoan                                *EnumOccupationItem
	OfficerPolice                              *EnumOccupationItem
	OfficerPoliceChiefCaptain                  *EnumOccupationItem
	OfficerPoliceDetectiveSgtLt                *EnumOccupationItem
	OfficerProbationParole                     *EnumOccupationItem
	OfficerTelecommunications                  *EnumOccupationItem
	OfficerWarrant                             *EnumOccupationItem
	OfficerorManagerBank                       *EnumOccupationItem
	OperatorAllOther                           *EnumOccupationItem
	OperatorBusiness                           *EnumOccupationItem
	OperatorControlRoom                        *EnumOccupationItem
	OperatorDataEntry                          *EnumOccupationItem
	OperatorForkLift                           *EnumOccupationItem
	OperatorHeavyEquipment                     *EnumOccupationItem
	OperatorMachinePrecision                   *EnumOccupationItem
	OperatorNuclearReactor                     *EnumOccupationItem
	OperatorTelephone                          *EnumOccupationItem
	OperatorWastewaterTreatmentPlantClassIV    *EnumOccupationItem
	Optician                                   *EnumOccupationItem
	Optometrist                                *EnumOccupationItem
	Orthodontist                               *EnumOccupationItem
	OwnerAllOther                              *EnumOccupationItem
	OwnerBar                                   *EnumOccupationItem
	OwnerBeautyBarberShop                      *EnumOccupationItem
	OwnerDealershipAutoDealer                  *EnumOccupationItem
	OwnerorManagerFarmOrRanch                  *EnumOccupationItem
	Painter                                    *EnumOccupationItem
	Paralegal                                  *EnumOccupationItem
	ParamedicorEMT                             *EnumOccupationItem
	ParkForestRanger                           *EnumOccupationItem
	PathologistSpeech                          *EnumOccupationItem
	PersonnelManagementSpecialist              *EnumOccupationItem
	PestControlWorkerorExterminator            *EnumOccupationItem
	Pharmacist                                 *EnumOccupationItem
	Pharmacologist                             *EnumOccupationItem
	Phlebotomist                               *EnumOccupationItem
	Photographer                               *EnumOccupationItem
	PhotographicProcessor                      *EnumOccupationItem
	PhysicalTherapistAPTAMember                *EnumOccupationItem
	PhysicalTherapistNonAPTAMember             *EnumOccupationItem
	PhysicianorDoctor                          *EnumOccupationItem
	Pilot                                      *EnumOccupationItem
	PilotCropBush                              *EnumOccupationItem
	PipefitterOtherFitter                      *EnumOccupationItem
	PlannerAllOther                            *EnumOccupationItem
	PlannerProductionorPrinter                 *EnumOccupationItem
	Plumber                                    *EnumOccupationItem
	Podiatrist                                 *EnumOccupationItem
	Politician                                 *EnumOccupationItem
	PoolServiceCleaner                         *EnumOccupationItem
	PostalExecutiveGradesPcesIII               *EnumOccupationItem
	PostmasterRural                            *EnumOccupationItem
	PostmasterUrbanSuburban                    *EnumOccupationItem
	PresidentBlueCollar50Empl                  *EnumOccupationItem
	PresidentSkilledBlueCollarLessThan50Emp    *EnumOccupationItem
	PresidentWhiteCollar                       *EnumOccupationItem
	PrincipalorAssistantPrincipal              *EnumOccupationItem
	Printer                                    *EnumOccupationItem
	Producer                                   *EnumOccupationItem
	Professor                                  *EnumOccupationItem
	ProgramManagementExpert                    *EnumOccupationItem
	Proofreader                                *EnumOccupationItem
	Psychiatrist                               *EnumOccupationItem
	Psychologist                               *EnumOccupationItem
	PublicRelations                            *EnumOccupationItem
	Publisher                                  *EnumOccupationItem
	QualityControlManufacturing                *EnumOccupationItem
	QualityControlProfessional                 *EnumOccupationItem
	Radiologist                                *EnumOccupationItem
	RanchHelperCowboy                          *EnumOccupationItem
	Recruiter                                  *EnumOccupationItem
	Registrar                                  *EnumOccupationItem
	ReligiousClergyOrdainedorLicensed          *EnumOccupationItem
	ReligiousLaypersonNonClergy                *EnumOccupationItem
	RepairServiceInstallACHeating              *EnumOccupationItem
	RepairServiceInstallAllOther               *EnumOccupationItem
	RepairServiceInstallJewelryWatchmaker      *EnumOccupationItem
	RepairServiceInstallLine                   *EnumOccupationItem
	RepairServiceInstallTrained                *EnumOccupationItem
	Reporter                                   *EnumOccupationItem
	ResearcherAllOther                         *EnumOccupationItem
	RespiratoryTherapist                       *EnumOccupationItem
	RoutemanRoutewoman                         *EnumOccupationItem
	SalespersonAllOther                        *EnumOccupationItem
	SalespersonCar                             *EnumOccupationItem
	SalespersonDoorToDoor                      *EnumOccupationItem
	SalespersonHighTech                        *EnumOccupationItem
	SalespersonNonHighTech                     *EnumOccupationItem
	SalespersonPharmaceutical                  *EnumOccupationItem
	SalespersonRetail                          *EnumOccupationItem
	SalespersonWholesale                       *EnumOccupationItem
	Sanitarian                                 *EnumOccupationItem
	Scheduler                                  *EnumOccupationItem
	ScientistAllOther                          *EnumOccupationItem
	SeamstressTailor                           *EnumOccupationItem
	SecurityScreener                           *EnumOccupationItem
	ShoeShinerRepairman                        *EnumOccupationItem
	SingerSongwriter                           *EnumOccupationItem
	StaffingSpecialist                         *EnumOccupationItem
	StateExaminer                              *EnumOccupationItem
	SuperintendentAllOther                     *EnumOccupationItem
	SuperintendentDriller                      *EnumOccupationItem
	SuperintendentSchool                       *EnumOccupationItem
	SuperintendentorSupervisorBuildingMaint    *EnumOccupationItem
	SupervisorAccounting                       *EnumOccupationItem
	SupervisorAllOtherDegreed                  *EnumOccupationItem
	SupervisorDataSystems                      *EnumOccupationItem
	SupervisorHumanResourcePersonnel           *EnumOccupationItem
	SupervisorOffice                           *EnumOccupationItem
	SupervisorOperations                       *EnumOccupationItem
	SupervisorOtherDegreed                     *EnumOccupationItem
	SupervisorPostal                           *EnumOccupationItem
	SupervisorProduction                       *EnumOccupationItem
	SupervisorRestaurantNonFastFood            *EnumOccupationItem
	Surgeon                                    *EnumOccupationItem
	SurveyorLicensed                           *EnumOccupationItem
	SurveyorNonLicensed                        *EnumOccupationItem
	TaxExaminerNotClerical                     *EnumOccupationItem
	TaxPreparerNotAccountant                   *EnumOccupationItem
	Taxidermist                                *EnumOccupationItem
	TeachersorCoachesorInstructors             *EnumOccupationItem
	TechnicianAllOther                         *EnumOccupationItem
	TechnicianElectricalorElectronic           *EnumOccupationItem
	TechnicianFood                             *EnumOccupationItem
	TechnicianInstrumentation                  *EnumOccupationItem
	TechnicianLab                              *EnumOccupationItem
	TechnicianMedical                          *EnumOccupationItem
	TechnicianRadiological                     *EnumOccupationItem
	TechnicianScience                          *EnumOccupationItem
	TechnicianTesting                          *EnumOccupationItem
	TechnicianUltrasound                       *EnumOccupationItem
	TechnicianXRay                             *EnumOccupationItem
	TechnicianorAssistantEngineering           *EnumOccupationItem
	Telemarketer                               *EnumOccupationItem
	Therapist                                  *EnumOccupationItem
	TrainerAerobicsFitness                     *EnumOccupationItem
	TrainerAthleticNataMember                  *EnumOccupationItem
	TrainerAthleticNonNataMember               *EnumOccupationItem
	TrainerCaretakerAnimal                     *EnumOccupationItem
	TranslatororInterpreter                    *EnumOccupationItem
	Treasurer                                  *EnumOccupationItem
	Tutor                                      *EnumOccupationItem
	UnderwriterInsurance                       *EnumOccupationItem
	Vendor                                     *EnumOccupationItem
	Veterinarian                               *EnumOccupationItem
	VicePresBusiness                           *EnumOccupationItem
	WaiterWaitress                             *EnumOccupationItem
	WardenAllOther                             *EnumOccupationItem
	WardenGame                                 *EnumOccupationItem
	WorkerMetalNotSteel                        *EnumOccupationItem
	WorkerRailroad                             *EnumOccupationItem
	WorkerSocialCase                           *EnumOccupationItem
	WriterAllOther                             *EnumOccupationItem
	WriterCommercial                           *EnumOccupationItem
	E1                                         *EnumOccupationItem
	E2                                         *EnumOccupationItem
	E3                                         *EnumOccupationItem
	E4                                         *EnumOccupationItem
	E4P                                        *EnumOccupationItem
	E5                                         *EnumOccupationItem
	E5P                                        *EnumOccupationItem
	E6                                         *EnumOccupationItem
	E6P                                        *EnumOccupationItem
	E7                                         *EnumOccupationItem
	E8                                         *EnumOccupationItem
	E9                                         *EnumOccupationItem
	O1                                         *EnumOccupationItem
	O10                                        *EnumOccupationItem
	O2                                         *EnumOccupationItem
	O3                                         *EnumOccupationItem
	O4                                         *EnumOccupationItem
	O5                                         *EnumOccupationItem
	O6                                         *EnumOccupationItem
	O7                                         *EnumOccupationItem
	O8                                         *EnumOccupationItem
	O9                                         *EnumOccupationItem
	W1                                         *EnumOccupationItem
	W2                                         *EnumOccupationItem
	W3                                         *EnumOccupationItem
	W4                                         *EnumOccupationItem
	W5                                         *EnumOccupationItem
	Homemaker                                  *EnumOccupationItem
	Unemployed                                 *EnumOccupationItem
	GraduateSchool                             *EnumOccupationItem
	HighSchool                                 *EnumOccupationItem
	TradeSchoolorAssociateDegree               *EnumOccupationItem
	Undergraduate4yeardegree                   *EnumOccupationItem
	Disabled                                   *EnumOccupationItem

	itemDict map[string]*EnumOccupationItem
}

// Occupation is a public singleton instance of EnumOccupation
// representing driver occupations
var Occupation = &EnumOccupation{
	Description: "driver occupations",
	Items: []*EnumOccupationItem{
		&occupationAccountExecutive,
		&occupationAccountantorCPA,
		&occupationActorActress,
		&occupationActuary,
		&occupationAcupuncturist,
		&occupationAdjudicator,
		&occupationAdministrativeAssistant,
		&occupationAdministrativeClerkorReceptionist,
		&occupationAdministrator,
		&occupationAdministratorHealthorHospital,
		&occupationAdministratorPublic,
		&occupationAdvisorFinancialorCredit,
		&occupationAgentAdvertising,
		&occupationAgentBorderPatrol,
		&occupationAgentEmployment,
		&occupationAgentFBI,
		&occupationAgentImportExport,
		&occupationAgentInsurance,
		&occupationAgentIRS,
		&occupationAgentLeasing,
		&occupationAgentRealEstate,
		&occupationAgentSecretService,
		&occupationAgentTheatrical,
		&occupationAgentTravelTicket,
		&occupationAgentTreasury,
		&occupationAideAllOther,
		&occupationAirTrafficController,
		&occupationAmbassador,
		&occupationAnalystAllOther,
		&occupationAnalystEngineering,
		&occupationAnalystFinancial,
		&occupationAnesthesiologist,
		&occupationAntiqueDealer,
		&occupationAppraiserAllOther,
		&occupationAppraiserArt,
		&occupationAppraiserRealEstate,
		&occupationApprenticeAllOther,
		&occupationArchitect,
		&occupationArchivistLibraryorMuseum,
		&occupationArtistCommercial,
		&occupationArtistNoncommercial,
		&occupationAssistantAllOther,
		&occupationAssistantEditorial,
		&occupationAssistantLegislativeorLegal,
		&occupationAssistantMedical,
		&occupationAssistantPrintingPress,
		&occupationAthlete,
		&occupationAttorneyorLawyer,
		&occupationAuctioneer,
		&occupationAudiologist,
		&occupationAuditorFinancial,
		&occupationAuditorNonFinancial,
		&occupationBaker,
		&occupationBankTeller,
		&occupationBartender,
		&occupationBoatCaptain,
		&occupationBoatman,
		&occupationBondsman,
		&occupationBookbinder,
		&occupationBookkeeper,
		&occupationBroadcaster,
		&occupationBrokerInsurance,
		&occupationBrokerMortgage,
		&occupationBrokerPawn,
		&occupationBrokerRealEstate,
		&occupationBrokerStockBonds,
		&occupationBuyerPurchasingAgent,
		&occupationCameraman,
		&occupationCarpenterorCabinetmaker,
		&occupationCashierFood,
		&occupationCashierOffice,
		&occupationCashierRetail,
		&occupationCaterer,
		&occupationCensusTaker,
		&occupationChiropractor,
		&occupationClaimsAdjuster,
		&occupationClaimsExaminer,
		&occupationClerkAccountingorFinancial,
		&occupationClerkAllOther,
		&occupationCollegeDean,
		&occupationCommunicationSpecialist,
		&occupationComputerProgrammer,
		&occupationComputerTechnicalSupportRep,
		&occupationConductor,
		&occupationConservationist,
		&occupationConstructionWorker,
		&occupationConsultantorAdvisor,
		&occupationContractororDeveloperAllOther,
		&occupationContractororDeveloperArtisanSkilled,
		&occupationContractororDeveloperBlueCollar50Employees,
		&occupationContractororDeveloperWhiteCollar,
		&occupationControllerorComptroller,
		&occupationControllerorCoordinatorProduction,
		&occupationCookorChef,
		&occupationCoordinatorOffice,
		&occupationCopywriter,
		&occupationCoroner,
		&occupationCosmetologistBeautician,
		&occupationCounselorAllOther,
		&occupationCounselorEducation,
		&occupationCounselorFamilyAndChild,
		&occupationCounselorMentalHealth,
		&occupationCourtReporter,
		&occupationCraftsman,
		&occupationCriminologist,
		&occupationCurator,
		&occupationCustodianorJanitor,
		&occupationCustomerServiceRepresentative,
		&occupationDayCareWorker,
		&occupationDecoratorInterior,
		&occupationDeliveryPersonorMailCarrier,
		&occupationDentalHygienist,
		&occupationDentist,
		&occupationDesignerComputerWebsite,
		&occupationDesignerFloral,
		&occupationDesignerGraphicorTechnical,
		&occupationDesignerProfessional,
		&occupationDesignerWindow,
		&occupationDietitianNutritionist,
		&occupationDirectororExecutive,
		&occupationDiscJockey,
		&occupationDispatcher,
		&occupationDogBreeder,
		&occupationDrafterorCartographer,
		&occupationDriverAllOther,
		&occupationDriverTruck,
		&occupationEconomist,
		&occupationEditorAllOther,
		&occupationEditorFilm,
		&occupationElectrician,
		&occupationEmbalmer,
		&occupationEngineerAllOther,
		&occupationEngineerCertifiedNetwork,
		&occupationEngineerComputerSoftware,
		&occupationEngineerComputerSystems,
		&occupationEngineerConstruction,
		&occupationEngineerElectricalElectronic,
		&occupationEngineerEquipment,
		&occupationEngineerFacilities,
		&occupationEngineerFlight,
		&occupationEngineerMechanical,
		&occupationEngineerOperating,
		&occupationEngineerPetroleumorMining,
		&occupationEngineerSafety,
		&occupationEngineerSales,
		&occupationEntertainerPerformer,
		&occupationExpediter,
		&occupationFactoryWorker,
		&occupationFiremanWomanChiefCaptLt,
		&occupationFiremanWomanNonChief,
		&occupationFisherman,
		&occupationFlightAttendant,
		&occupationFloormenSupervisor,
		&occupationFlorist,
		&occupationForemanForewoman,
		&occupationForester,
		&occupationFundraiser,
		&occupationGeographer,
		&occupationGovtOfficialElected,
		&occupationGrader,
		&occupationGuardEmbassy,
		&occupationGuardSecurityorPrison,
		&occupationGunsmith,
		&occupationHairdresserBarber,
		&occupationHistorian,
		&occupationHostorHostessRestaurant,
		&occupationHousekeeperorMaid,
		&occupationHumanResourcesRepresentative,
		&occupationIllustrator,
		&occupationInspectorAgricultural,
		&occupationInspectorAircraftAccessories,
		&occupationInspectorAllOther,
		&occupationInspectorConstruction,
		&occupationInspectorPostal,
		&occupationInspectorWhiteCollar,
		&occupationInvestigatorPrivate,
		&occupationInvestmentBanker,
		&occupationInvestorPrivate,
		&occupationJournalist,
		&occupationJourneyman,
		&occupationJudge,
		&occupationLaborRelationsWorker,
		&occupationLandscaper,
		&occupationLibrarian,
		&occupationLifeGuard,
		&occupationLinguist,
		&occupationLithographer,
		&occupationLobbyist,
		&occupationLocksmith,
		&occupationLongshoremen,
		&occupationMachinist,
		&occupationManagerAirport,
		&occupationManagerAllOtherDegreed,
		&occupationManagerCafeteria,
		&occupationManagerCity,
		&occupationManagerClericalStaff,
		&occupationManagerConvenienceorGasStationStore,
		&occupationManagerDepartmentStore,
		&occupationManagerFinancialorCredit,
		&occupationManagerGeneral,
		&occupationManagerHealthClub,
		&occupationManagerHotel,
		&occupationManagerHumanResources,
		&occupationManagerMerchandise,
		&occupationManagerOffice,
		&occupationManagerOperations,
		&occupationManagerProduction,
		&occupationManagerProfessionalTechStaff,
		&occupationManagerProject,
		&occupationManagerPropertyNonResident,
		&occupationManagerPropertyResident,
		&occupationManagerRestaurantFastFood,
		&occupationManagerRestaurantNonFastFood,
		&occupationManagerSales,
		&occupationManagerSecurityScreener,
		&occupationManagerShippingReceiving,
		&occupationManagerStage,
		&occupationManagerSupermarket,
		&occupationManagerorOwnerSandwichShop,
		&occupationManicurist,
		&occupationMarketingRepresentative,
		&occupationMarshalFire,
		&occupationMarshalUSDeputy,
		&occupationMasseuse,
		&occupationMathematician,
		&occupationMeatcutterButcher,
		&occupationMechanicorServicemanAuto,
		&occupationMechanicorServicemanBoat,
		&occupationMechanicorServicemanDiesel,
		&occupationMerchant,
		&occupationMillwright,
		&occupationMortician,
		&occupationMusicianClassical,
		&occupationMusicianOther,
		&occupationNurseCNACertifiedNursingAssistant,
		&occupationNurseLVNorLPN,
		&occupationNurseRN,
		&occupationNursePractitioner,
		&occupationOceanographer,
		&occupationOfficerCorrectional,
		&occupationOfficerCourt,
		&occupationOfficerForeignService,
		&occupationOfficerLoan,
		&occupationOfficerPolice,
		&occupationOfficerPoliceChiefCaptain,
		&occupationOfficerPoliceDetectiveSgtLt,
		&occupationOfficerProbationParole,
		&occupationOfficerTelecommunications,
		&occupationOfficerWarrant,
		&occupationOfficerorManagerBank,
		&occupationOperatorAllOther,
		&occupationOperatorBusiness,
		&occupationOperatorControlRoom,
		&occupationOperatorDataEntry,
		&occupationOperatorForkLift,
		&occupationOperatorHeavyEquipment,
		&occupationOperatorMachinePrecision,
		&occupationOperatorNuclearReactor,
		&occupationOperatorTelephone,
		&occupationOperatorWastewaterTreatmentPlantClassIV,
		&occupationOptician,
		&occupationOptometrist,
		&occupationOrthodontist,
		&occupationOwnerAllOther,
		&occupationOwnerBar,
		&occupationOwnerBeautyBarberShop,
		&occupationOwnerDealershipAutoDealer,
		&occupationOwnerorManagerFarmOrRanch,
		&occupationPainter,
		&occupationParalegal,
		&occupationParamedicorEMT,
		&occupationParkForestRanger,
		&occupationPathologistSpeech,
		&occupationPersonnelManagementSpecialist,
		&occupationPestControlWorkerorExterminator,
		&occupationPharmacist,
		&occupationPharmacologist,
		&occupationPhlebotomist,
		&occupationPhotographer,
		&occupationPhotographicProcessor,
		&occupationPhysicalTherapistAPTAMember,
		&occupationPhysicalTherapistNonAPTAMember,
		&occupationPhysicianorDoctor,
		&occupationPilot,
		&occupationPilotCropBush,
		&occupationPipefitterOtherFitter,
		&occupationPlannerAllOther,
		&occupationPlannerProductionorPrinter,
		&occupationPlumber,
		&occupationPodiatrist,
		&occupationPolitician,
		&occupationPoolServiceCleaner,
		&occupationPostalExecutiveGradesPcesIII,
		&occupationPostmasterRural,
		&occupationPostmasterUrbanSuburban,
		&occupationPresidentBlueCollar50Empl,
		&occupationPresidentSkilledBlueCollarLessThan50Emp,
		&occupationPresidentWhiteCollar,
		&occupationPrincipalorAssistantPrincipal,
		&occupationPrinter,
		&occupationProducer,
		&occupationProfessor,
		&occupationProgramManagementExpert,
		&occupationProofreader,
		&occupationPsychiatrist,
		&occupationPsychologist,
		&occupationPublicRelations,
		&occupationPublisher,
		&occupationQualityControlManufacturing,
		&occupationQualityControlProfessional,
		&occupationRadiologist,
		&occupationRanchHelperCowboy,
		&occupationRecruiter,
		&occupationRegistrar,
		&occupationReligiousClergyOrdainedorLicensed,
		&occupationReligiousLaypersonNonClergy,
		&occupationRepairServiceInstallACHeating,
		&occupationRepairServiceInstallAllOther,
		&occupationRepairServiceInstallJewelryWatchmaker,
		&occupationRepairServiceInstallLine,
		&occupationRepairServiceInstallTrained,
		&occupationReporter,
		&occupationResearcherAllOther,
		&occupationRespiratoryTherapist,
		&occupationRoutemanRoutewoman,
		&occupationSalespersonAllOther,
		&occupationSalespersonCar,
		&occupationSalespersonDoorToDoor,
		&occupationSalespersonHighTech,
		&occupationSalespersonNonHighTech,
		&occupationSalespersonPharmaceutical,
		&occupationSalespersonRetail,
		&occupationSalespersonWholesale,
		&occupationSanitarian,
		&occupationScheduler,
		&occupationScientistAllOther,
		&occupationSeamstressTailor,
		&occupationSecurityScreener,
		&occupationShoeShinerRepairman,
		&occupationSingerSongwriter,
		&occupationStaffingSpecialist,
		&occupationStateExaminer,
		&occupationSuperintendentAllOther,
		&occupationSuperintendentDriller,
		&occupationSuperintendentSchool,
		&occupationSuperintendentorSupervisorBuildingMaint,
		&occupationSupervisorAccounting,
		&occupationSupervisorAllOtherDegreed,
		&occupationSupervisorDataSystems,
		&occupationSupervisorHumanResourcePersonnel,
		&occupationSupervisorOffice,
		&occupationSupervisorOperations,
		&occupationSupervisorOtherDegreed,
		&occupationSupervisorPostal,
		&occupationSupervisorProduction,
		&occupationSupervisorRestaurantNonFastFood,
		&occupationSurgeon,
		&occupationSurveyorLicensed,
		&occupationSurveyorNonLicensed,
		&occupationTaxExaminerNotClerical,
		&occupationTaxPreparerNotAccountant,
		&occupationTaxidermist,
		&occupationTeachersorCoachesorInstructors,
		&occupationTechnicianAllOther,
		&occupationTechnicianElectricalorElectronic,
		&occupationTechnicianFood,
		&occupationTechnicianInstrumentation,
		&occupationTechnicianLab,
		&occupationTechnicianMedical,
		&occupationTechnicianRadiological,
		&occupationTechnicianScience,
		&occupationTechnicianTesting,
		&occupationTechnicianUltrasound,
		&occupationTechnicianXRay,
		&occupationTechnicianorAssistantEngineering,
		&occupationTelemarketer,
		&occupationTherapist,
		&occupationTrainerAerobicsFitness,
		&occupationTrainerAthleticNataMember,
		&occupationTrainerAthleticNonNataMember,
		&occupationTrainerCaretakerAnimal,
		&occupationTranslatororInterpreter,
		&occupationTreasurer,
		&occupationTutor,
		&occupationUnderwriterInsurance,
		&occupationVendor,
		&occupationVeterinarian,
		&occupationVicePresBusiness,
		&occupationWaiterWaitress,
		&occupationWardenAllOther,
		&occupationWardenGame,
		&occupationWorkerMetalNotSteel,
		&occupationWorkerRailroad,
		&occupationWorkerSocialCase,
		&occupationWriterAllOther,
		&occupationWriterCommercial,
		&occupationE1,
		&occupationE2,
		&occupationE3,
		&occupationE4,
		&occupationE4P,
		&occupationE5,
		&occupationE5P,
		&occupationE6,
		&occupationE6P,
		&occupationE7,
		&occupationE8,
		&occupationE9,
		&occupationO1,
		&occupationO10,
		&occupationO2,
		&occupationO3,
		&occupationO4,
		&occupationO5,
		&occupationO6,
		&occupationO7,
		&occupationO8,
		&occupationO9,
		&occupationW1,
		&occupationW2,
		&occupationW3,
		&occupationW4,
		&occupationW5,
		&occupationHomemaker,
		&occupationUnemployed,
		&occupationGraduateSchool,
		&occupationHighSchool,
		&occupationTradeSchoolorAssociateDegree,
		&occupationUndergraduate4yeardegree,
		&occupationDisabled,
	},
	Name:                                "EnumOccupation",
	AccountExecutive:                    &occupationAccountExecutive,
	AccountantorCPA:                     &occupationAccountantorCPA,
	ActorActress:                        &occupationActorActress,
	Actuary:                             &occupationActuary,
	Acupuncturist:                       &occupationAcupuncturist,
	Adjudicator:                         &occupationAdjudicator,
	AdministrativeAssistant:             &occupationAdministrativeAssistant,
	AdministrativeClerkorReceptionist:   &occupationAdministrativeClerkorReceptionist,
	Administrator:                       &occupationAdministrator,
	AdministratorHealthorHospital:       &occupationAdministratorHealthorHospital,
	AdministratorPublic:                 &occupationAdministratorPublic,
	AdvisorFinancialorCredit:            &occupationAdvisorFinancialorCredit,
	AgentAdvertising:                    &occupationAgentAdvertising,
	AgentBorderPatrol:                   &occupationAgentBorderPatrol,
	AgentEmployment:                     &occupationAgentEmployment,
	AgentFBI:                            &occupationAgentFBI,
	AgentImportExport:                   &occupationAgentImportExport,
	AgentInsurance:                      &occupationAgentInsurance,
	AgentIRS:                            &occupationAgentIRS,
	AgentLeasing:                        &occupationAgentLeasing,
	AgentRealEstate:                     &occupationAgentRealEstate,
	AgentSecretService:                  &occupationAgentSecretService,
	AgentTheatrical:                     &occupationAgentTheatrical,
	AgentTravelTicket:                   &occupationAgentTravelTicket,
	AgentTreasury:                       &occupationAgentTreasury,
	AideAllOther:                        &occupationAideAllOther,
	AirTrafficController:                &occupationAirTrafficController,
	Ambassador:                          &occupationAmbassador,
	AnalystAllOther:                     &occupationAnalystAllOther,
	AnalystEngineering:                  &occupationAnalystEngineering,
	AnalystFinancial:                    &occupationAnalystFinancial,
	Anesthesiologist:                    &occupationAnesthesiologist,
	AntiqueDealer:                       &occupationAntiqueDealer,
	AppraiserAllOther:                   &occupationAppraiserAllOther,
	AppraiserArt:                        &occupationAppraiserArt,
	AppraiserRealEstate:                 &occupationAppraiserRealEstate,
	ApprenticeAllOther:                  &occupationApprenticeAllOther,
	Architect:                           &occupationArchitect,
	ArchivistLibraryorMuseum:            &occupationArchivistLibraryorMuseum,
	ArtistCommercial:                    &occupationArtistCommercial,
	ArtistNoncommercial:                 &occupationArtistNoncommercial,
	AssistantAllOther:                   &occupationAssistantAllOther,
	AssistantEditorial:                  &occupationAssistantEditorial,
	AssistantLegislativeorLegal:         &occupationAssistantLegislativeorLegal,
	AssistantMedical:                    &occupationAssistantMedical,
	AssistantPrintingPress:              &occupationAssistantPrintingPress,
	Athlete:                             &occupationAthlete,
	AttorneyorLawyer:                    &occupationAttorneyorLawyer,
	Auctioneer:                          &occupationAuctioneer,
	Audiologist:                         &occupationAudiologist,
	AuditorFinancial:                    &occupationAuditorFinancial,
	AuditorNonFinancial:                 &occupationAuditorNonFinancial,
	Baker:                               &occupationBaker,
	BankTeller:                          &occupationBankTeller,
	Bartender:                           &occupationBartender,
	BoatCaptain:                         &occupationBoatCaptain,
	Boatman:                             &occupationBoatman,
	Bondsman:                            &occupationBondsman,
	Bookbinder:                          &occupationBookbinder,
	Bookkeeper:                          &occupationBookkeeper,
	Broadcaster:                         &occupationBroadcaster,
	BrokerInsurance:                     &occupationBrokerInsurance,
	BrokerMortgage:                      &occupationBrokerMortgage,
	BrokerPawn:                          &occupationBrokerPawn,
	BrokerRealEstate:                    &occupationBrokerRealEstate,
	BrokerStockBonds:                    &occupationBrokerStockBonds,
	BuyerPurchasingAgent:                &occupationBuyerPurchasingAgent,
	Cameraman:                           &occupationCameraman,
	CarpenterorCabinetmaker:             &occupationCarpenterorCabinetmaker,
	CashierFood:                         &occupationCashierFood,
	CashierOffice:                       &occupationCashierOffice,
	CashierRetail:                       &occupationCashierRetail,
	Caterer:                             &occupationCaterer,
	CensusTaker:                         &occupationCensusTaker,
	Chiropractor:                        &occupationChiropractor,
	ClaimsAdjuster:                      &occupationClaimsAdjuster,
	ClaimsExaminer:                      &occupationClaimsExaminer,
	ClerkAccountingorFinancial:          &occupationClerkAccountingorFinancial,
	ClerkAllOther:                       &occupationClerkAllOther,
	CollegeDean:                         &occupationCollegeDean,
	CommunicationSpecialist:             &occupationCommunicationSpecialist,
	ComputerProgrammer:                  &occupationComputerProgrammer,
	ComputerTechnicalSupportRep:         &occupationComputerTechnicalSupportRep,
	Conductor:                           &occupationConductor,
	Conservationist:                     &occupationConservationist,
	ConstructionWorker:                  &occupationConstructionWorker,
	ConsultantorAdvisor:                 &occupationConsultantorAdvisor,
	ContractororDeveloperAllOther:       &occupationContractororDeveloperAllOther,
	ContractororDeveloperArtisanSkilled: &occupationContractororDeveloperArtisanSkilled,
	ContractororDeveloperBlueCollar50Employees: &occupationContractororDeveloperBlueCollar50Employees,
	ContractororDeveloperWhiteCollar:           &occupationContractororDeveloperWhiteCollar,
	ControllerorComptroller:                    &occupationControllerorComptroller,
	ControllerorCoordinatorProduction:          &occupationControllerorCoordinatorProduction,
	CookorChef:                                 &occupationCookorChef,
	CoordinatorOffice:                          &occupationCoordinatorOffice,
	Copywriter:                                 &occupationCopywriter,
	Coroner:                                    &occupationCoroner,
	CosmetologistBeautician:                    &occupationCosmetologistBeautician,
	CounselorAllOther:                          &occupationCounselorAllOther,
	CounselorEducation:                         &occupationCounselorEducation,
	CounselorFamilyAndChild:                    &occupationCounselorFamilyAndChild,
	CounselorMentalHealth:                      &occupationCounselorMentalHealth,
	CourtReporter:                              &occupationCourtReporter,
	Craftsman:                                  &occupationCraftsman,
	Criminologist:                              &occupationCriminologist,
	Curator:                                    &occupationCurator,
	CustodianorJanitor:                         &occupationCustodianorJanitor,
	CustomerServiceRepresentative:              &occupationCustomerServiceRepresentative,
	DayCareWorker:                              &occupationDayCareWorker,
	DecoratorInterior:                          &occupationDecoratorInterior,
	DeliveryPersonorMailCarrier:                &occupationDeliveryPersonorMailCarrier,
	DentalHygienist:                            &occupationDentalHygienist,
	Dentist:                                    &occupationDentist,
	DesignerComputerWebsite:                    &occupationDesignerComputerWebsite,
	DesignerFloral:                             &occupationDesignerFloral,
	DesignerGraphicorTechnical:                 &occupationDesignerGraphicorTechnical,
	DesignerProfessional:                       &occupationDesignerProfessional,
	DesignerWindow:                             &occupationDesignerWindow,
	DietitianNutritionist:                      &occupationDietitianNutritionist,
	DirectororExecutive:                        &occupationDirectororExecutive,
	DiscJockey:                                 &occupationDiscJockey,
	Dispatcher:                                 &occupationDispatcher,
	DogBreeder:                                 &occupationDogBreeder,
	DrafterorCartographer:                      &occupationDrafterorCartographer,
	DriverAllOther:                             &occupationDriverAllOther,
	DriverTruck:                                &occupationDriverTruck,
	Economist:                                  &occupationEconomist,
	EditorAllOther:                             &occupationEditorAllOther,
	EditorFilm:                                 &occupationEditorFilm,
	Electrician:                                &occupationElectrician,
	Embalmer:                                   &occupationEmbalmer,
	EngineerAllOther:                           &occupationEngineerAllOther,
	EngineerCertifiedNetwork:                   &occupationEngineerCertifiedNetwork,
	EngineerComputerSoftware:                   &occupationEngineerComputerSoftware,
	EngineerComputerSystems:                    &occupationEngineerComputerSystems,
	EngineerConstruction:                       &occupationEngineerConstruction,
	EngineerElectricalElectronic:               &occupationEngineerElectricalElectronic,
	EngineerEquipment:                          &occupationEngineerEquipment,
	EngineerFacilities:                         &occupationEngineerFacilities,
	EngineerFlight:                             &occupationEngineerFlight,
	EngineerMechanical:                         &occupationEngineerMechanical,
	EngineerOperating:                          &occupationEngineerOperating,
	EngineerPetroleumorMining:                  &occupationEngineerPetroleumorMining,
	EngineerSafety:                             &occupationEngineerSafety,
	EngineerSales:                              &occupationEngineerSales,
	EntertainerPerformer:                       &occupationEntertainerPerformer,
	Expediter:                                  &occupationExpediter,
	FactoryWorker:                              &occupationFactoryWorker,
	FiremanWomanChiefCaptLt:                    &occupationFiremanWomanChiefCaptLt,
	FiremanWomanNonChief:                       &occupationFiremanWomanNonChief,
	Fisherman:                                  &occupationFisherman,
	FlightAttendant:                            &occupationFlightAttendant,
	FloormenSupervisor:                         &occupationFloormenSupervisor,
	Florist:                                    &occupationFlorist,
	ForemanForewoman:                           &occupationForemanForewoman,
	Forester:                                   &occupationForester,
	Fundraiser:                                 &occupationFundraiser,
	Geographer:                                 &occupationGeographer,
	GovtOfficialElected:                        &occupationGovtOfficialElected,
	Grader:                                     &occupationGrader,
	GuardEmbassy:                               &occupationGuardEmbassy,
	GuardSecurityorPrison:                      &occupationGuardSecurityorPrison,
	Gunsmith:                                   &occupationGunsmith,
	HairdresserBarber:                          &occupationHairdresserBarber,
	Historian:                                  &occupationHistorian,
	HostorHostessRestaurant:                    &occupationHostorHostessRestaurant,
	HousekeeperorMaid:                          &occupationHousekeeperorMaid,
	HumanResourcesRepresentative:               &occupationHumanResourcesRepresentative,
	Illustrator:                                &occupationIllustrator,
	InspectorAgricultural:                      &occupationInspectorAgricultural,
	InspectorAircraftAccessories:               &occupationInspectorAircraftAccessories,
	InspectorAllOther:                          &occupationInspectorAllOther,
	InspectorConstruction:                      &occupationInspectorConstruction,
	InspectorPostal:                            &occupationInspectorPostal,
	InspectorWhiteCollar:                       &occupationInspectorWhiteCollar,
	InvestigatorPrivate:                        &occupationInvestigatorPrivate,
	InvestmentBanker:                           &occupationInvestmentBanker,
	InvestorPrivate:                            &occupationInvestorPrivate,
	Journalist:                                 &occupationJournalist,
	Journeyman:                                 &occupationJourneyman,
	Judge:                                      &occupationJudge,
	LaborRelationsWorker:                       &occupationLaborRelationsWorker,
	Landscaper:                                 &occupationLandscaper,
	Librarian:                                  &occupationLibrarian,
	LifeGuard:                                  &occupationLifeGuard,
	Linguist:                                   &occupationLinguist,
	Lithographer:                               &occupationLithographer,
	Lobbyist:                                   &occupationLobbyist,
	Locksmith:                                  &occupationLocksmith,
	Longshoremen:                               &occupationLongshoremen,
	Machinist:                                  &occupationMachinist,
	ManagerAirport:                             &occupationManagerAirport,
	ManagerAllOtherDegreed:                     &occupationManagerAllOtherDegreed,
	ManagerCafeteria:                           &occupationManagerCafeteria,
	ManagerCity:                                &occupationManagerCity,
	ManagerClericalStaff:                       &occupationManagerClericalStaff,
	ManagerConvenienceorGasStationStore:        &occupationManagerConvenienceorGasStationStore,
	ManagerDepartmentStore:                     &occupationManagerDepartmentStore,
	ManagerFinancialorCredit:                   &occupationManagerFinancialorCredit,
	ManagerGeneral:                             &occupationManagerGeneral,
	ManagerHealthClub:                          &occupationManagerHealthClub,
	ManagerHotel:                               &occupationManagerHotel,
	ManagerHumanResources:                      &occupationManagerHumanResources,
	ManagerMerchandise:                         &occupationManagerMerchandise,
	ManagerOffice:                              &occupationManagerOffice,
	ManagerOperations:                          &occupationManagerOperations,
	ManagerProduction:                          &occupationManagerProduction,
	ManagerProfessionalTechStaff:               &occupationManagerProfessionalTechStaff,
	ManagerProject:                             &occupationManagerProject,
	ManagerPropertyNonResident:                 &occupationManagerPropertyNonResident,
	ManagerPropertyResident:                    &occupationManagerPropertyResident,
	ManagerRestaurantFastFood:                  &occupationManagerRestaurantFastFood,
	ManagerRestaurantNonFastFood:               &occupationManagerRestaurantNonFastFood,
	ManagerSales:                               &occupationManagerSales,
	ManagerSecurityScreener:                    &occupationManagerSecurityScreener,
	ManagerShippingReceiving:                   &occupationManagerShippingReceiving,
	ManagerStage:                               &occupationManagerStage,
	ManagerSupermarket:                         &occupationManagerSupermarket,
	ManagerorOwnerSandwichShop:                 &occupationManagerorOwnerSandwichShop,
	Manicurist:                                 &occupationManicurist,
	MarketingRepresentative:                    &occupationMarketingRepresentative,
	MarshalFire:                                &occupationMarshalFire,
	MarshalUSDeputy:                            &occupationMarshalUSDeputy,
	Masseuse:                                   &occupationMasseuse,
	Mathematician:                              &occupationMathematician,
	MeatcutterButcher:                          &occupationMeatcutterButcher,
	MechanicorServicemanAuto:                   &occupationMechanicorServicemanAuto,
	MechanicorServicemanBoat:                   &occupationMechanicorServicemanBoat,
	MechanicorServicemanDiesel:                 &occupationMechanicorServicemanDiesel,
	Merchant:                                   &occupationMerchant,
	Millwright:                                 &occupationMillwright,
	Mortician:                                  &occupationMortician,
	MusicianClassical:                          &occupationMusicianClassical,
	MusicianOther:                              &occupationMusicianOther,
	NurseCNACertifiedNursingAssistant:          &occupationNurseCNACertifiedNursingAssistant,
	NurseLVNorLPN:                              &occupationNurseLVNorLPN,
	NurseRN:                                    &occupationNurseRN,
	NursePractitioner:                          &occupationNursePractitioner,
	Oceanographer:                              &occupationOceanographer,
	OfficerCorrectional:                        &occupationOfficerCorrectional,
	OfficerCourt:                               &occupationOfficerCourt,
	OfficerForeignService:                      &occupationOfficerForeignService,
	OfficerLoan:                                &occupationOfficerLoan,
	OfficerPolice:                              &occupationOfficerPolice,
	OfficerPoliceChiefCaptain:                  &occupationOfficerPoliceChiefCaptain,
	OfficerPoliceDetectiveSgtLt:                &occupationOfficerPoliceDetectiveSgtLt,
	OfficerProbationParole:                     &occupationOfficerProbationParole,
	OfficerTelecommunications:                  &occupationOfficerTelecommunications,
	OfficerWarrant:                             &occupationOfficerWarrant,
	OfficerorManagerBank:                       &occupationOfficerorManagerBank,
	OperatorAllOther:                           &occupationOperatorAllOther,
	OperatorBusiness:                           &occupationOperatorBusiness,
	OperatorControlRoom:                        &occupationOperatorControlRoom,
	OperatorDataEntry:                          &occupationOperatorDataEntry,
	OperatorForkLift:                           &occupationOperatorForkLift,
	OperatorHeavyEquipment:                     &occupationOperatorHeavyEquipment,
	OperatorMachinePrecision:                   &occupationOperatorMachinePrecision,
	OperatorNuclearReactor:                     &occupationOperatorNuclearReactor,
	OperatorTelephone:                          &occupationOperatorTelephone,
	OperatorWastewaterTreatmentPlantClassIV:    &occupationOperatorWastewaterTreatmentPlantClassIV,
	Optician:                                   &occupationOptician,
	Optometrist:                                &occupationOptometrist,
	Orthodontist:                               &occupationOrthodontist,
	OwnerAllOther:                              &occupationOwnerAllOther,
	OwnerBar:                                   &occupationOwnerBar,
	OwnerBeautyBarberShop:                      &occupationOwnerBeautyBarberShop,
	OwnerDealershipAutoDealer:                  &occupationOwnerDealershipAutoDealer,
	OwnerorManagerFarmOrRanch:                  &occupationOwnerorManagerFarmOrRanch,
	Painter:                                    &occupationPainter,
	Paralegal:                                  &occupationParalegal,
	ParamedicorEMT:                             &occupationParamedicorEMT,
	ParkForestRanger:                           &occupationParkForestRanger,
	PathologistSpeech:                          &occupationPathologistSpeech,
	PersonnelManagementSpecialist:              &occupationPersonnelManagementSpecialist,
	PestControlWorkerorExterminator:            &occupationPestControlWorkerorExterminator,
	Pharmacist:                                 &occupationPharmacist,
	Pharmacologist:                             &occupationPharmacologist,
	Phlebotomist:                               &occupationPhlebotomist,
	Photographer:                               &occupationPhotographer,
	PhotographicProcessor:                      &occupationPhotographicProcessor,
	PhysicalTherapistAPTAMember:                &occupationPhysicalTherapistAPTAMember,
	PhysicalTherapistNonAPTAMember:             &occupationPhysicalTherapistNonAPTAMember,
	PhysicianorDoctor:                          &occupationPhysicianorDoctor,
	Pilot:                                      &occupationPilot,
	PilotCropBush:                              &occupationPilotCropBush,
	PipefitterOtherFitter:                      &occupationPipefitterOtherFitter,
	PlannerAllOther:                            &occupationPlannerAllOther,
	PlannerProductionorPrinter:                 &occupationPlannerProductionorPrinter,
	Plumber:                                    &occupationPlumber,
	Podiatrist:                                 &occupationPodiatrist,
	Politician:                                 &occupationPolitician,
	PoolServiceCleaner:                         &occupationPoolServiceCleaner,
	PostalExecutiveGradesPcesIII:               &occupationPostalExecutiveGradesPcesIII,
	PostmasterRural:                            &occupationPostmasterRural,
	PostmasterUrbanSuburban:                    &occupationPostmasterUrbanSuburban,
	PresidentBlueCollar50Empl:                  &occupationPresidentBlueCollar50Empl,
	PresidentSkilledBlueCollarLessThan50Emp:    &occupationPresidentSkilledBlueCollarLessThan50Emp,
	PresidentWhiteCollar:                       &occupationPresidentWhiteCollar,
	PrincipalorAssistantPrincipal:              &occupationPrincipalorAssistantPrincipal,
	Printer:                                    &occupationPrinter,
	Producer:                                   &occupationProducer,
	Professor:                                  &occupationProfessor,
	ProgramManagementExpert:                    &occupationProgramManagementExpert,
	Proofreader:                                &occupationProofreader,
	Psychiatrist:                               &occupationPsychiatrist,
	Psychologist:                               &occupationPsychologist,
	PublicRelations:                            &occupationPublicRelations,
	Publisher:                                  &occupationPublisher,
	QualityControlManufacturing:                &occupationQualityControlManufacturing,
	QualityControlProfessional:                 &occupationQualityControlProfessional,
	Radiologist:                                &occupationRadiologist,
	RanchHelperCowboy:                          &occupationRanchHelperCowboy,
	Recruiter:                                  &occupationRecruiter,
	Registrar:                                  &occupationRegistrar,
	ReligiousClergyOrdainedorLicensed:          &occupationReligiousClergyOrdainedorLicensed,
	ReligiousLaypersonNonClergy:                &occupationReligiousLaypersonNonClergy,
	RepairServiceInstallACHeating:              &occupationRepairServiceInstallACHeating,
	RepairServiceInstallAllOther:               &occupationRepairServiceInstallAllOther,
	RepairServiceInstallJewelryWatchmaker:      &occupationRepairServiceInstallJewelryWatchmaker,
	RepairServiceInstallLine:                   &occupationRepairServiceInstallLine,
	RepairServiceInstallTrained:                &occupationRepairServiceInstallTrained,
	Reporter:                                   &occupationReporter,
	ResearcherAllOther:                         &occupationResearcherAllOther,
	RespiratoryTherapist:                       &occupationRespiratoryTherapist,
	RoutemanRoutewoman:                         &occupationRoutemanRoutewoman,
	SalespersonAllOther:                        &occupationSalespersonAllOther,
	SalespersonCar:                             &occupationSalespersonCar,
	SalespersonDoorToDoor:                      &occupationSalespersonDoorToDoor,
	SalespersonHighTech:                        &occupationSalespersonHighTech,
	SalespersonNonHighTech:                     &occupationSalespersonNonHighTech,
	SalespersonPharmaceutical:                  &occupationSalespersonPharmaceutical,
	SalespersonRetail:                          &occupationSalespersonRetail,
	SalespersonWholesale:                       &occupationSalespersonWholesale,
	Sanitarian:                                 &occupationSanitarian,
	Scheduler:                                  &occupationScheduler,
	ScientistAllOther:                          &occupationScientistAllOther,
	SeamstressTailor:                           &occupationSeamstressTailor,
	SecurityScreener:                           &occupationSecurityScreener,
	ShoeShinerRepairman:                        &occupationShoeShinerRepairman,
	SingerSongwriter:                           &occupationSingerSongwriter,
	StaffingSpecialist:                         &occupationStaffingSpecialist,
	StateExaminer:                              &occupationStateExaminer,
	SuperintendentAllOther:                     &occupationSuperintendentAllOther,
	SuperintendentDriller:                      &occupationSuperintendentDriller,
	SuperintendentSchool:                       &occupationSuperintendentSchool,
	SuperintendentorSupervisorBuildingMaint:    &occupationSuperintendentorSupervisorBuildingMaint,
	SupervisorAccounting:                       &occupationSupervisorAccounting,
	SupervisorAllOtherDegreed:                  &occupationSupervisorAllOtherDegreed,
	SupervisorDataSystems:                      &occupationSupervisorDataSystems,
	SupervisorHumanResourcePersonnel:           &occupationSupervisorHumanResourcePersonnel,
	SupervisorOffice:                           &occupationSupervisorOffice,
	SupervisorOperations:                       &occupationSupervisorOperations,
	SupervisorOtherDegreed:                     &occupationSupervisorOtherDegreed,
	SupervisorPostal:                           &occupationSupervisorPostal,
	SupervisorProduction:                       &occupationSupervisorProduction,
	SupervisorRestaurantNonFastFood:            &occupationSupervisorRestaurantNonFastFood,
	Surgeon:                                    &occupationSurgeon,
	SurveyorLicensed:                           &occupationSurveyorLicensed,
	SurveyorNonLicensed:                        &occupationSurveyorNonLicensed,
	TaxExaminerNotClerical:                     &occupationTaxExaminerNotClerical,
	TaxPreparerNotAccountant:                   &occupationTaxPreparerNotAccountant,
	Taxidermist:                                &occupationTaxidermist,
	TeachersorCoachesorInstructors:             &occupationTeachersorCoachesorInstructors,
	TechnicianAllOther:                         &occupationTechnicianAllOther,
	TechnicianElectricalorElectronic:           &occupationTechnicianElectricalorElectronic,
	TechnicianFood:                             &occupationTechnicianFood,
	TechnicianInstrumentation:                  &occupationTechnicianInstrumentation,
	TechnicianLab:                              &occupationTechnicianLab,
	TechnicianMedical:                          &occupationTechnicianMedical,
	TechnicianRadiological:                     &occupationTechnicianRadiological,
	TechnicianScience:                          &occupationTechnicianScience,
	TechnicianTesting:                          &occupationTechnicianTesting,
	TechnicianUltrasound:                       &occupationTechnicianUltrasound,
	TechnicianXRay:                             &occupationTechnicianXRay,
	TechnicianorAssistantEngineering:           &occupationTechnicianorAssistantEngineering,
	Telemarketer:                               &occupationTelemarketer,
	Therapist:                                  &occupationTherapist,
	TrainerAerobicsFitness:                     &occupationTrainerAerobicsFitness,
	TrainerAthleticNataMember:                  &occupationTrainerAthleticNataMember,
	TrainerAthleticNonNataMember:               &occupationTrainerAthleticNonNataMember,
	TrainerCaretakerAnimal:                     &occupationTrainerCaretakerAnimal,
	TranslatororInterpreter:                    &occupationTranslatororInterpreter,
	Treasurer:                                  &occupationTreasurer,
	Tutor:                                      &occupationTutor,
	UnderwriterInsurance:                       &occupationUnderwriterInsurance,
	Vendor:                                     &occupationVendor,
	Veterinarian:                               &occupationVeterinarian,
	VicePresBusiness:                           &occupationVicePresBusiness,
	WaiterWaitress:                             &occupationWaiterWaitress,
	WardenAllOther:                             &occupationWardenAllOther,
	WardenGame:                                 &occupationWardenGame,
	WorkerMetalNotSteel:                        &occupationWorkerMetalNotSteel,
	WorkerRailroad:                             &occupationWorkerRailroad,
	WorkerSocialCase:                           &occupationWorkerSocialCase,
	WriterAllOther:                             &occupationWriterAllOther,
	WriterCommercial:                           &occupationWriterCommercial,
	E1:                                         &occupationE1,
	E2:                                         &occupationE2,
	E3:                                         &occupationE3,
	E4:                                         &occupationE4,
	E4P:                                        &occupationE4P,
	E5:                                         &occupationE5,
	E5P:                                        &occupationE5P,
	E6:                                         &occupationE6,
	E6P:                                        &occupationE6P,
	E7:                                         &occupationE7,
	E8:                                         &occupationE8,
	E9:                                         &occupationE9,
	O1:                                         &occupationO1,
	O10:                                        &occupationO10,
	O2:                                         &occupationO2,
	O3:                                         &occupationO3,
	O4:                                         &occupationO4,
	O5:                                         &occupationO5,
	O6:                                         &occupationO6,
	O7:                                         &occupationO7,
	O8:                                         &occupationO8,
	O9:                                         &occupationO9,
	W1:                                         &occupationW1,
	W2:                                         &occupationW2,
	W3:                                         &occupationW3,
	W4:                                         &occupationW4,
	W5:                                         &occupationW5,
	Homemaker:                                  &occupationHomemaker,
	Unemployed:                                 &occupationUnemployed,
	GraduateSchool:                             &occupationGraduateSchool,
	HighSchool:                                 &occupationHighSchool,
	TradeSchoolorAssociateDegree:               &occupationTradeSchoolorAssociateDegree,
	Undergraduate4yeardegree:                   &occupationUndergraduate4yeardegree,
	Disabled:                                   &occupationDisabled,

	itemDict: map[string]*EnumOccupationItem{
		strings.ToLower(string(occupationAccountExecutiveID)):                           &occupationAccountExecutive,
		strings.ToLower(string(occupationAccountantorCPAID)):                            &occupationAccountantorCPA,
		strings.ToLower(string(occupationActorActressID)):                               &occupationActorActress,
		strings.ToLower(string(occupationActuaryID)):                                    &occupationActuary,
		strings.ToLower(string(occupationAcupuncturistID)):                              &occupationAcupuncturist,
		strings.ToLower(string(occupationAdjudicatorID)):                                &occupationAdjudicator,
		strings.ToLower(string(occupationAdministrativeAssistantID)):                    &occupationAdministrativeAssistant,
		strings.ToLower(string(occupationAdministrativeClerkorReceptionistID)):          &occupationAdministrativeClerkorReceptionist,
		strings.ToLower(string(occupationAdministratorID)):                              &occupationAdministrator,
		strings.ToLower(string(occupationAdministratorHealthorHospitalID)):              &occupationAdministratorHealthorHospital,
		strings.ToLower(string(occupationAdministratorPublicID)):                        &occupationAdministratorPublic,
		strings.ToLower(string(occupationAdvisorFinancialorCreditID)):                   &occupationAdvisorFinancialorCredit,
		strings.ToLower(string(occupationAgentAdvertisingID)):                           &occupationAgentAdvertising,
		strings.ToLower(string(occupationAgentBorderPatrolID)):                          &occupationAgentBorderPatrol,
		strings.ToLower(string(occupationAgentEmploymentID)):                            &occupationAgentEmployment,
		strings.ToLower(string(occupationAgentFBIID)):                                   &occupationAgentFBI,
		strings.ToLower(string(occupationAgentImportExportID)):                          &occupationAgentImportExport,
		strings.ToLower(string(occupationAgentInsuranceID)):                             &occupationAgentInsurance,
		strings.ToLower(string(occupationAgentIRSID)):                                   &occupationAgentIRS,
		strings.ToLower(string(occupationAgentLeasingID)):                               &occupationAgentLeasing,
		strings.ToLower(string(occupationAgentRealEstateID)):                            &occupationAgentRealEstate,
		strings.ToLower(string(occupationAgentSecretServiceID)):                         &occupationAgentSecretService,
		strings.ToLower(string(occupationAgentTheatricalID)):                            &occupationAgentTheatrical,
		strings.ToLower(string(occupationAgentTravelTicketID)):                          &occupationAgentTravelTicket,
		strings.ToLower(string(occupationAgentTreasuryID)):                              &occupationAgentTreasury,
		strings.ToLower(string(occupationAideAllOtherID)):                               &occupationAideAllOther,
		strings.ToLower(string(occupationAirTrafficControllerID)):                       &occupationAirTrafficController,
		strings.ToLower(string(occupationAmbassadorID)):                                 &occupationAmbassador,
		strings.ToLower(string(occupationAnalystAllOtherID)):                            &occupationAnalystAllOther,
		strings.ToLower(string(occupationAnalystEngineeringID)):                         &occupationAnalystEngineering,
		strings.ToLower(string(occupationAnalystFinancialID)):                           &occupationAnalystFinancial,
		strings.ToLower(string(occupationAnesthesiologistID)):                           &occupationAnesthesiologist,
		strings.ToLower(string(occupationAntiqueDealerID)):                              &occupationAntiqueDealer,
		strings.ToLower(string(occupationAppraiserAllOtherID)):                          &occupationAppraiserAllOther,
		strings.ToLower(string(occupationAppraiserArtID)):                               &occupationAppraiserArt,
		strings.ToLower(string(occupationAppraiserRealEstateID)):                        &occupationAppraiserRealEstate,
		strings.ToLower(string(occupationApprenticeAllOtherID)):                         &occupationApprenticeAllOther,
		strings.ToLower(string(occupationArchitectID)):                                  &occupationArchitect,
		strings.ToLower(string(occupationArchivistLibraryorMuseumID)):                   &occupationArchivistLibraryorMuseum,
		strings.ToLower(string(occupationArtistCommercialID)):                           &occupationArtistCommercial,
		strings.ToLower(string(occupationArtistNoncommercialID)):                        &occupationArtistNoncommercial,
		strings.ToLower(string(occupationAssistantAllOtherID)):                          &occupationAssistantAllOther,
		strings.ToLower(string(occupationAssistantEditorialID)):                         &occupationAssistantEditorial,
		strings.ToLower(string(occupationAssistantLegislativeorLegalID)):                &occupationAssistantLegislativeorLegal,
		strings.ToLower(string(occupationAssistantMedicalID)):                           &occupationAssistantMedical,
		strings.ToLower(string(occupationAssistantPrintingPressID)):                     &occupationAssistantPrintingPress,
		strings.ToLower(string(occupationAthleteID)):                                    &occupationAthlete,
		strings.ToLower(string(occupationAttorneyorLawyerID)):                           &occupationAttorneyorLawyer,
		strings.ToLower(string(occupationAuctioneerID)):                                 &occupationAuctioneer,
		strings.ToLower(string(occupationAudiologistID)):                                &occupationAudiologist,
		strings.ToLower(string(occupationAuditorFinancialID)):                           &occupationAuditorFinancial,
		strings.ToLower(string(occupationAuditorNonFinancialID)):                        &occupationAuditorNonFinancial,
		strings.ToLower(string(occupationBakerID)):                                      &occupationBaker,
		strings.ToLower(string(occupationBankTellerID)):                                 &occupationBankTeller,
		strings.ToLower(string(occupationBartenderID)):                                  &occupationBartender,
		strings.ToLower(string(occupationBoatCaptainID)):                                &occupationBoatCaptain,
		strings.ToLower(string(occupationBoatmanID)):                                    &occupationBoatman,
		strings.ToLower(string(occupationBondsmanID)):                                   &occupationBondsman,
		strings.ToLower(string(occupationBookbinderID)):                                 &occupationBookbinder,
		strings.ToLower(string(occupationBookkeeperID)):                                 &occupationBookkeeper,
		strings.ToLower(string(occupationBroadcasterID)):                                &occupationBroadcaster,
		strings.ToLower(string(occupationBrokerInsuranceID)):                            &occupationBrokerInsurance,
		strings.ToLower(string(occupationBrokerMortgageID)):                             &occupationBrokerMortgage,
		strings.ToLower(string(occupationBrokerPawnID)):                                 &occupationBrokerPawn,
		strings.ToLower(string(occupationBrokerRealEstateID)):                           &occupationBrokerRealEstate,
		strings.ToLower(string(occupationBrokerStockBondsID)):                           &occupationBrokerStockBonds,
		strings.ToLower(string(occupationBuyerPurchasingAgentID)):                       &occupationBuyerPurchasingAgent,
		strings.ToLower(string(occupationCameramanID)):                                  &occupationCameraman,
		strings.ToLower(string(occupationCarpenterorCabinetmakerID)):                    &occupationCarpenterorCabinetmaker,
		strings.ToLower(string(occupationCashierFoodID)):                                &occupationCashierFood,
		strings.ToLower(string(occupationCashierOfficeID)):                              &occupationCashierOffice,
		strings.ToLower(string(occupationCashierRetailID)):                              &occupationCashierRetail,
		strings.ToLower(string(occupationCatererID)):                                    &occupationCaterer,
		strings.ToLower(string(occupationCensusTakerID)):                                &occupationCensusTaker,
		strings.ToLower(string(occupationChiropractorID)):                               &occupationChiropractor,
		strings.ToLower(string(occupationClaimsAdjusterID)):                             &occupationClaimsAdjuster,
		strings.ToLower(string(occupationClaimsExaminerID)):                             &occupationClaimsExaminer,
		strings.ToLower(string(occupationClerkAccountingorFinancialID)):                 &occupationClerkAccountingorFinancial,
		strings.ToLower(string(occupationClerkAllOtherID)):                              &occupationClerkAllOther,
		strings.ToLower(string(occupationCollegeDeanID)):                                &occupationCollegeDean,
		strings.ToLower(string(occupationCommunicationSpecialistID)):                    &occupationCommunicationSpecialist,
		strings.ToLower(string(occupationComputerProgrammerID)):                         &occupationComputerProgrammer,
		strings.ToLower(string(occupationComputerTechnicalSupportRepID)):                &occupationComputerTechnicalSupportRep,
		strings.ToLower(string(occupationConductorID)):                                  &occupationConductor,
		strings.ToLower(string(occupationConservationistID)):                            &occupationConservationist,
		strings.ToLower(string(occupationConstructionWorkerID)):                         &occupationConstructionWorker,
		strings.ToLower(string(occupationConsultantorAdvisorID)):                        &occupationConsultantorAdvisor,
		strings.ToLower(string(occupationContractororDeveloperAllOtherID)):              &occupationContractororDeveloperAllOther,
		strings.ToLower(string(occupationContractororDeveloperArtisanSkilledID)):        &occupationContractororDeveloperArtisanSkilled,
		strings.ToLower(string(occupationContractororDeveloperBlueCollar50EmployeesID)): &occupationContractororDeveloperBlueCollar50Employees,
		strings.ToLower(string(occupationContractororDeveloperWhiteCollarID)):           &occupationContractororDeveloperWhiteCollar,
		strings.ToLower(string(occupationControllerorComptrollerID)):                    &occupationControllerorComptroller,
		strings.ToLower(string(occupationControllerorCoordinatorProductionID)):          &occupationControllerorCoordinatorProduction,
		strings.ToLower(string(occupationCookorChefID)):                                 &occupationCookorChef,
		strings.ToLower(string(occupationCoordinatorOfficeID)):                          &occupationCoordinatorOffice,
		strings.ToLower(string(occupationCopywriterID)):                                 &occupationCopywriter,
		strings.ToLower(string(occupationCoronerID)):                                    &occupationCoroner,
		strings.ToLower(string(occupationCosmetologistBeauticianID)):                    &occupationCosmetologistBeautician,
		strings.ToLower(string(occupationCounselorAllOtherID)):                          &occupationCounselorAllOther,
		strings.ToLower(string(occupationCounselorEducationID)):                         &occupationCounselorEducation,
		strings.ToLower(string(occupationCounselorFamilyAndChildID)):                    &occupationCounselorFamilyAndChild,
		strings.ToLower(string(occupationCounselorMentalHealthID)):                      &occupationCounselorMentalHealth,
		strings.ToLower(string(occupationCourtReporterID)):                              &occupationCourtReporter,
		strings.ToLower(string(occupationCraftsmanID)):                                  &occupationCraftsman,
		strings.ToLower(string(occupationCriminologistID)):                              &occupationCriminologist,
		strings.ToLower(string(occupationCuratorID)):                                    &occupationCurator,
		strings.ToLower(string(occupationCustodianorJanitorID)):                         &occupationCustodianorJanitor,
		strings.ToLower(string(occupationCustomerServiceRepresentativeID)):              &occupationCustomerServiceRepresentative,
		strings.ToLower(string(occupationDayCareWorkerID)):                              &occupationDayCareWorker,
		strings.ToLower(string(occupationDecoratorInteriorID)):                          &occupationDecoratorInterior,
		strings.ToLower(string(occupationDeliveryPersonorMailCarrierID)):                &occupationDeliveryPersonorMailCarrier,
		strings.ToLower(string(occupationDentalHygienistID)):                            &occupationDentalHygienist,
		strings.ToLower(string(occupationDentistID)):                                    &occupationDentist,
		strings.ToLower(string(occupationDesignerComputerWebsiteID)):                    &occupationDesignerComputerWebsite,
		strings.ToLower(string(occupationDesignerFloralID)):                             &occupationDesignerFloral,
		strings.ToLower(string(occupationDesignerGraphicorTechnicalID)):                 &occupationDesignerGraphicorTechnical,
		strings.ToLower(string(occupationDesignerProfessionalID)):                       &occupationDesignerProfessional,
		strings.ToLower(string(occupationDesignerWindowID)):                             &occupationDesignerWindow,
		strings.ToLower(string(occupationDietitianNutritionistID)):                      &occupationDietitianNutritionist,
		strings.ToLower(string(occupationDirectororExecutiveID)):                        &occupationDirectororExecutive,
		strings.ToLower(string(occupationDiscJockeyID)):                                 &occupationDiscJockey,
		strings.ToLower(string(occupationDispatcherID)):                                 &occupationDispatcher,
		strings.ToLower(string(occupationDogBreederID)):                                 &occupationDogBreeder,
		strings.ToLower(string(occupationDrafterorCartographerID)):                      &occupationDrafterorCartographer,
		strings.ToLower(string(occupationDriverAllOtherID)):                             &occupationDriverAllOther,
		strings.ToLower(string(occupationDriverTruckID)):                                &occupationDriverTruck,
		strings.ToLower(string(occupationEconomistID)):                                  &occupationEconomist,
		strings.ToLower(string(occupationEditorAllOtherID)):                             &occupationEditorAllOther,
		strings.ToLower(string(occupationEditorFilmID)):                                 &occupationEditorFilm,
		strings.ToLower(string(occupationElectricianID)):                                &occupationElectrician,
		strings.ToLower(string(occupationEmbalmerID)):                                   &occupationEmbalmer,
		strings.ToLower(string(occupationEngineerAllOtherID)):                           &occupationEngineerAllOther,
		strings.ToLower(string(occupationEngineerCertifiedNetworkID)):                   &occupationEngineerCertifiedNetwork,
		strings.ToLower(string(occupationEngineerComputerSoftwareID)):                   &occupationEngineerComputerSoftware,
		strings.ToLower(string(occupationEngineerComputerSystemsID)):                    &occupationEngineerComputerSystems,
		strings.ToLower(string(occupationEngineerConstructionID)):                       &occupationEngineerConstruction,
		strings.ToLower(string(occupationEngineerElectricalElectronicID)):               &occupationEngineerElectricalElectronic,
		strings.ToLower(string(occupationEngineerEquipmentID)):                          &occupationEngineerEquipment,
		strings.ToLower(string(occupationEngineerFacilitiesID)):                         &occupationEngineerFacilities,
		strings.ToLower(string(occupationEngineerFlightID)):                             &occupationEngineerFlight,
		strings.ToLower(string(occupationEngineerMechanicalID)):                         &occupationEngineerMechanical,
		strings.ToLower(string(occupationEngineerOperatingID)):                          &occupationEngineerOperating,
		strings.ToLower(string(occupationEngineerPetroleumorMiningID)):                  &occupationEngineerPetroleumorMining,
		strings.ToLower(string(occupationEngineerSafetyID)):                             &occupationEngineerSafety,
		strings.ToLower(string(occupationEngineerSalesID)):                              &occupationEngineerSales,
		strings.ToLower(string(occupationEntertainerPerformerID)):                       &occupationEntertainerPerformer,
		strings.ToLower(string(occupationExpediterID)):                                  &occupationExpediter,
		strings.ToLower(string(occupationFactoryWorkerID)):                              &occupationFactoryWorker,
		strings.ToLower(string(occupationFiremanWomanChiefCaptLtID)):                    &occupationFiremanWomanChiefCaptLt,
		strings.ToLower(string(occupationFiremanWomanNonChiefID)):                       &occupationFiremanWomanNonChief,
		strings.ToLower(string(occupationFishermanID)):                                  &occupationFisherman,
		strings.ToLower(string(occupationFlightAttendantID)):                            &occupationFlightAttendant,
		strings.ToLower(string(occupationFloormenSupervisorID)):                         &occupationFloormenSupervisor,
		strings.ToLower(string(occupationFloristID)):                                    &occupationFlorist,
		strings.ToLower(string(occupationForemanForewomanID)):                           &occupationForemanForewoman,
		strings.ToLower(string(occupationForesterID)):                                   &occupationForester,
		strings.ToLower(string(occupationFundraiserID)):                                 &occupationFundraiser,
		strings.ToLower(string(occupationGeographerID)):                                 &occupationGeographer,
		strings.ToLower(string(occupationGovtOfficialElectedID)):                        &occupationGovtOfficialElected,
		strings.ToLower(string(occupationGraderID)):                                     &occupationGrader,
		strings.ToLower(string(occupationGuardEmbassyID)):                               &occupationGuardEmbassy,
		strings.ToLower(string(occupationGuardSecurityorPrisonID)):                      &occupationGuardSecurityorPrison,
		strings.ToLower(string(occupationGunsmithID)):                                   &occupationGunsmith,
		strings.ToLower(string(occupationHairdresserBarberID)):                          &occupationHairdresserBarber,
		strings.ToLower(string(occupationHistorianID)):                                  &occupationHistorian,
		strings.ToLower(string(occupationHostorHostessRestaurantID)):                    &occupationHostorHostessRestaurant,
		strings.ToLower(string(occupationHousekeeperorMaidID)):                          &occupationHousekeeperorMaid,
		strings.ToLower(string(occupationHumanResourcesRepresentativeID)):               &occupationHumanResourcesRepresentative,
		strings.ToLower(string(occupationIllustratorID)):                                &occupationIllustrator,
		strings.ToLower(string(occupationInspectorAgriculturalID)):                      &occupationInspectorAgricultural,
		strings.ToLower(string(occupationInspectorAircraftAccessoriesID)):               &occupationInspectorAircraftAccessories,
		strings.ToLower(string(occupationInspectorAllOtherID)):                          &occupationInspectorAllOther,
		strings.ToLower(string(occupationInspectorConstructionID)):                      &occupationInspectorConstruction,
		strings.ToLower(string(occupationInspectorPostalID)):                            &occupationInspectorPostal,
		strings.ToLower(string(occupationInspectorWhiteCollarID)):                       &occupationInspectorWhiteCollar,
		strings.ToLower(string(occupationInvestigatorPrivateID)):                        &occupationInvestigatorPrivate,
		strings.ToLower(string(occupationInvestmentBankerID)):                           &occupationInvestmentBanker,
		strings.ToLower(string(occupationInvestorPrivateID)):                            &occupationInvestorPrivate,
		strings.ToLower(string(occupationJournalistID)):                                 &occupationJournalist,
		strings.ToLower(string(occupationJourneymanID)):                                 &occupationJourneyman,
		strings.ToLower(string(occupationJudgeID)):                                      &occupationJudge,
		strings.ToLower(string(occupationLaborRelationsWorkerID)):                       &occupationLaborRelationsWorker,
		strings.ToLower(string(occupationLandscaperID)):                                 &occupationLandscaper,
		strings.ToLower(string(occupationLibrarianID)):                                  &occupationLibrarian,
		strings.ToLower(string(occupationLifeGuardID)):                                  &occupationLifeGuard,
		strings.ToLower(string(occupationLinguistID)):                                   &occupationLinguist,
		strings.ToLower(string(occupationLithographerID)):                               &occupationLithographer,
		strings.ToLower(string(occupationLobbyistID)):                                   &occupationLobbyist,
		strings.ToLower(string(occupationLocksmithID)):                                  &occupationLocksmith,
		strings.ToLower(string(occupationLongshoremenID)):                               &occupationLongshoremen,
		strings.ToLower(string(occupationMachinistID)):                                  &occupationMachinist,
		strings.ToLower(string(occupationManagerAirportID)):                             &occupationManagerAirport,
		strings.ToLower(string(occupationManagerAllOtherDegreedID)):                     &occupationManagerAllOtherDegreed,
		strings.ToLower(string(occupationManagerCafeteriaID)):                           &occupationManagerCafeteria,
		strings.ToLower(string(occupationManagerCityID)):                                &occupationManagerCity,
		strings.ToLower(string(occupationManagerClericalStaffID)):                       &occupationManagerClericalStaff,
		strings.ToLower(string(occupationManagerConvenienceorGasStationStoreID)):        &occupationManagerConvenienceorGasStationStore,
		strings.ToLower(string(occupationManagerDepartmentStoreID)):                     &occupationManagerDepartmentStore,
		strings.ToLower(string(occupationManagerFinancialorCreditID)):                   &occupationManagerFinancialorCredit,
		strings.ToLower(string(occupationManagerGeneralID)):                             &occupationManagerGeneral,
		strings.ToLower(string(occupationManagerHealthClubID)):                          &occupationManagerHealthClub,
		strings.ToLower(string(occupationManagerHotelID)):                               &occupationManagerHotel,
		strings.ToLower(string(occupationManagerHumanResourcesID)):                      &occupationManagerHumanResources,
		strings.ToLower(string(occupationManagerMerchandiseID)):                         &occupationManagerMerchandise,
		strings.ToLower(string(occupationManagerOfficeID)):                              &occupationManagerOffice,
		strings.ToLower(string(occupationManagerOperationsID)):                          &occupationManagerOperations,
		strings.ToLower(string(occupationManagerProductionID)):                          &occupationManagerProduction,
		strings.ToLower(string(occupationManagerProfessionalTechStaffID)):               &occupationManagerProfessionalTechStaff,
		strings.ToLower(string(occupationManagerProjectID)):                             &occupationManagerProject,
		strings.ToLower(string(occupationManagerPropertyNonResidentID)):                 &occupationManagerPropertyNonResident,
		strings.ToLower(string(occupationManagerPropertyResidentID)):                    &occupationManagerPropertyResident,
		strings.ToLower(string(occupationManagerRestaurantFastFoodID)):                  &occupationManagerRestaurantFastFood,
		strings.ToLower(string(occupationManagerRestaurantNonFastFoodID)):               &occupationManagerRestaurantNonFastFood,
		strings.ToLower(string(occupationManagerSalesID)):                               &occupationManagerSales,
		strings.ToLower(string(occupationManagerSecurityScreenerID)):                    &occupationManagerSecurityScreener,
		strings.ToLower(string(occupationManagerShippingReceivingID)):                   &occupationManagerShippingReceiving,
		strings.ToLower(string(occupationManagerStageID)):                               &occupationManagerStage,
		strings.ToLower(string(occupationManagerSupermarketID)):                         &occupationManagerSupermarket,
		strings.ToLower(string(occupationManagerorOwnerSandwichShopID)):                 &occupationManagerorOwnerSandwichShop,
		strings.ToLower(string(occupationManicuristID)):                                 &occupationManicurist,
		strings.ToLower(string(occupationMarketingRepresentativeID)):                    &occupationMarketingRepresentative,
		strings.ToLower(string(occupationMarshalFireID)):                                &occupationMarshalFire,
		strings.ToLower(string(occupationMarshalUSDeputyID)):                            &occupationMarshalUSDeputy,
		strings.ToLower(string(occupationMasseuseID)):                                   &occupationMasseuse,
		strings.ToLower(string(occupationMathematicianID)):                              &occupationMathematician,
		strings.ToLower(string(occupationMeatcutterButcherID)):                          &occupationMeatcutterButcher,
		strings.ToLower(string(occupationMechanicorServicemanAutoID)):                   &occupationMechanicorServicemanAuto,
		strings.ToLower(string(occupationMechanicorServicemanBoatID)):                   &occupationMechanicorServicemanBoat,
		strings.ToLower(string(occupationMechanicorServicemanDieselID)):                 &occupationMechanicorServicemanDiesel,
		strings.ToLower(string(occupationMerchantID)):                                   &occupationMerchant,
		strings.ToLower(string(occupationMillwrightID)):                                 &occupationMillwright,
		strings.ToLower(string(occupationMorticianID)):                                  &occupationMortician,
		strings.ToLower(string(occupationMusicianClassicalID)):                          &occupationMusicianClassical,
		strings.ToLower(string(occupationMusicianOtherID)):                              &occupationMusicianOther,
		strings.ToLower(string(occupationNurseCNACertifiedNursingAssistantID)):          &occupationNurseCNACertifiedNursingAssistant,
		strings.ToLower(string(occupationNurseLVNorLPNID)):                              &occupationNurseLVNorLPN,
		strings.ToLower(string(occupationNurseRNID)):                                    &occupationNurseRN,
		strings.ToLower(string(occupationNursePractitionerID)):                          &occupationNursePractitioner,
		strings.ToLower(string(occupationOceanographerID)):                              &occupationOceanographer,
		strings.ToLower(string(occupationOfficerCorrectionalID)):                        &occupationOfficerCorrectional,
		strings.ToLower(string(occupationOfficerCourtID)):                               &occupationOfficerCourt,
		strings.ToLower(string(occupationOfficerForeignServiceID)):                      &occupationOfficerForeignService,
		strings.ToLower(string(occupationOfficerLoanID)):                                &occupationOfficerLoan,
		strings.ToLower(string(occupationOfficerPoliceID)):                              &occupationOfficerPolice,
		strings.ToLower(string(occupationOfficerPoliceChiefCaptainID)):                  &occupationOfficerPoliceChiefCaptain,
		strings.ToLower(string(occupationOfficerPoliceDetectiveSgtLtID)):                &occupationOfficerPoliceDetectiveSgtLt,
		strings.ToLower(string(occupationOfficerProbationParoleID)):                     &occupationOfficerProbationParole,
		strings.ToLower(string(occupationOfficerTelecommunicationsID)):                  &occupationOfficerTelecommunications,
		strings.ToLower(string(occupationOfficerWarrantID)):                             &occupationOfficerWarrant,
		strings.ToLower(string(occupationOfficerorManagerBankID)):                       &occupationOfficerorManagerBank,
		strings.ToLower(string(occupationOperatorAllOtherID)):                           &occupationOperatorAllOther,
		strings.ToLower(string(occupationOperatorBusinessID)):                           &occupationOperatorBusiness,
		strings.ToLower(string(occupationOperatorControlRoomID)):                        &occupationOperatorControlRoom,
		strings.ToLower(string(occupationOperatorDataEntryID)):                          &occupationOperatorDataEntry,
		strings.ToLower(string(occupationOperatorForkLiftID)):                           &occupationOperatorForkLift,
		strings.ToLower(string(occupationOperatorHeavyEquipmentID)):                     &occupationOperatorHeavyEquipment,
		strings.ToLower(string(occupationOperatorMachinePrecisionID)):                   &occupationOperatorMachinePrecision,
		strings.ToLower(string(occupationOperatorNuclearReactorID)):                     &occupationOperatorNuclearReactor,
		strings.ToLower(string(occupationOperatorTelephoneID)):                          &occupationOperatorTelephone,
		strings.ToLower(string(occupationOperatorWastewaterTreatmentPlantClassIVID)):    &occupationOperatorWastewaterTreatmentPlantClassIV,
		strings.ToLower(string(occupationOpticianID)):                                   &occupationOptician,
		strings.ToLower(string(occupationOptometristID)):                                &occupationOptometrist,
		strings.ToLower(string(occupationOrthodontistID)):                               &occupationOrthodontist,
		strings.ToLower(string(occupationOwnerAllOtherID)):                              &occupationOwnerAllOther,
		strings.ToLower(string(occupationOwnerBarID)):                                   &occupationOwnerBar,
		strings.ToLower(string(occupationOwnerBeautyBarberShopID)):                      &occupationOwnerBeautyBarberShop,
		strings.ToLower(string(occupationOwnerDealershipAutoDealerID)):                  &occupationOwnerDealershipAutoDealer,
		strings.ToLower(string(occupationOwnerorManagerFarmOrRanchID)):                  &occupationOwnerorManagerFarmOrRanch,
		strings.ToLower(string(occupationPainterID)):                                    &occupationPainter,
		strings.ToLower(string(occupationParalegalID)):                                  &occupationParalegal,
		strings.ToLower(string(occupationParamedicorEMTID)):                             &occupationParamedicorEMT,
		strings.ToLower(string(occupationParkForestRangerID)):                           &occupationParkForestRanger,
		strings.ToLower(string(occupationPathologistSpeechID)):                          &occupationPathologistSpeech,
		strings.ToLower(string(occupationPersonnelManagementSpecialistID)):              &occupationPersonnelManagementSpecialist,
		strings.ToLower(string(occupationPestControlWorkerorExterminatorID)):            &occupationPestControlWorkerorExterminator,
		strings.ToLower(string(occupationPharmacistID)):                                 &occupationPharmacist,
		strings.ToLower(string(occupationPharmacologistID)):                             &occupationPharmacologist,
		strings.ToLower(string(occupationPhlebotomistID)):                               &occupationPhlebotomist,
		strings.ToLower(string(occupationPhotographerID)):                               &occupationPhotographer,
		strings.ToLower(string(occupationPhotographicProcessorID)):                      &occupationPhotographicProcessor,
		strings.ToLower(string(occupationPhysicalTherapistAPTAMemberID)):                &occupationPhysicalTherapistAPTAMember,
		strings.ToLower(string(occupationPhysicalTherapistNonAPTAMemberID)):             &occupationPhysicalTherapistNonAPTAMember,
		strings.ToLower(string(occupationPhysicianorDoctorID)):                          &occupationPhysicianorDoctor,
		strings.ToLower(string(occupationPilotID)):                                      &occupationPilot,
		strings.ToLower(string(occupationPilotCropBushID)):                              &occupationPilotCropBush,
		strings.ToLower(string(occupationPipefitterOtherFitterID)):                      &occupationPipefitterOtherFitter,
		strings.ToLower(string(occupationPlannerAllOtherID)):                            &occupationPlannerAllOther,
		strings.ToLower(string(occupationPlannerProductionorPrinterID)):                 &occupationPlannerProductionorPrinter,
		strings.ToLower(string(occupationPlumberID)):                                    &occupationPlumber,
		strings.ToLower(string(occupationPodiatristID)):                                 &occupationPodiatrist,
		strings.ToLower(string(occupationPoliticianID)):                                 &occupationPolitician,
		strings.ToLower(string(occupationPoolServiceCleanerID)):                         &occupationPoolServiceCleaner,
		strings.ToLower(string(occupationPostalExecutiveGradesPcesIIIID)):               &occupationPostalExecutiveGradesPcesIII,
		strings.ToLower(string(occupationPostmasterRuralID)):                            &occupationPostmasterRural,
		strings.ToLower(string(occupationPostmasterUrbanSuburbanID)):                    &occupationPostmasterUrbanSuburban,
		strings.ToLower(string(occupationPresidentBlueCollar50EmplID)):                  &occupationPresidentBlueCollar50Empl,
		strings.ToLower(string(occupationPresidentSkilledBlueCollarLessThan50EmpID)):    &occupationPresidentSkilledBlueCollarLessThan50Emp,
		strings.ToLower(string(occupationPresidentWhiteCollarID)):                       &occupationPresidentWhiteCollar,
		strings.ToLower(string(occupationPrincipalorAssistantPrincipalID)):              &occupationPrincipalorAssistantPrincipal,
		strings.ToLower(string(occupationPrinterID)):                                    &occupationPrinter,
		strings.ToLower(string(occupationProducerID)):                                   &occupationProducer,
		strings.ToLower(string(occupationProfessorID)):                                  &occupationProfessor,
		strings.ToLower(string(occupationProgramManagementExpertID)):                    &occupationProgramManagementExpert,
		strings.ToLower(string(occupationProofreaderID)):                                &occupationProofreader,
		strings.ToLower(string(occupationPsychiatristID)):                               &occupationPsychiatrist,
		strings.ToLower(string(occupationPsychologistID)):                               &occupationPsychologist,
		strings.ToLower(string(occupationPublicRelationsID)):                            &occupationPublicRelations,
		strings.ToLower(string(occupationPublisherID)):                                  &occupationPublisher,
		strings.ToLower(string(occupationQualityControlManufacturingID)):                &occupationQualityControlManufacturing,
		strings.ToLower(string(occupationQualityControlProfessionalID)):                 &occupationQualityControlProfessional,
		strings.ToLower(string(occupationRadiologistID)):                                &occupationRadiologist,
		strings.ToLower(string(occupationRanchHelperCowboyID)):                          &occupationRanchHelperCowboy,
		strings.ToLower(string(occupationRecruiterID)):                                  &occupationRecruiter,
		strings.ToLower(string(occupationRegistrarID)):                                  &occupationRegistrar,
		strings.ToLower(string(occupationReligiousClergyOrdainedorLicensedID)):          &occupationReligiousClergyOrdainedorLicensed,
		strings.ToLower(string(occupationReligiousLaypersonNonClergyID)):                &occupationReligiousLaypersonNonClergy,
		strings.ToLower(string(occupationRepairServiceInstallACHeatingID)):              &occupationRepairServiceInstallACHeating,
		strings.ToLower(string(occupationRepairServiceInstallAllOtherID)):               &occupationRepairServiceInstallAllOther,
		strings.ToLower(string(occupationRepairServiceInstallJewelryWatchmakerID)):      &occupationRepairServiceInstallJewelryWatchmaker,
		strings.ToLower(string(occupationRepairServiceInstallLineID)):                   &occupationRepairServiceInstallLine,
		strings.ToLower(string(occupationRepairServiceInstallTrainedID)):                &occupationRepairServiceInstallTrained,
		strings.ToLower(string(occupationReporterID)):                                   &occupationReporter,
		strings.ToLower(string(occupationResearcherAllOtherID)):                         &occupationResearcherAllOther,
		strings.ToLower(string(occupationRespiratoryTherapistID)):                       &occupationRespiratoryTherapist,
		strings.ToLower(string(occupationRoutemanRoutewomanID)):                         &occupationRoutemanRoutewoman,
		strings.ToLower(string(occupationSalespersonAllOtherID)):                        &occupationSalespersonAllOther,
		strings.ToLower(string(occupationSalespersonCarID)):                             &occupationSalespersonCar,
		strings.ToLower(string(occupationSalespersonDoorToDoorID)):                      &occupationSalespersonDoorToDoor,
		strings.ToLower(string(occupationSalespersonHighTechID)):                        &occupationSalespersonHighTech,
		strings.ToLower(string(occupationSalespersonNonHighTechID)):                     &occupationSalespersonNonHighTech,
		strings.ToLower(string(occupationSalespersonPharmaceuticalID)):                  &occupationSalespersonPharmaceutical,
		strings.ToLower(string(occupationSalespersonRetailID)):                          &occupationSalespersonRetail,
		strings.ToLower(string(occupationSalespersonWholesaleID)):                       &occupationSalespersonWholesale,
		strings.ToLower(string(occupationSanitarianID)):                                 &occupationSanitarian,
		strings.ToLower(string(occupationSchedulerID)):                                  &occupationScheduler,
		strings.ToLower(string(occupationScientistAllOtherID)):                          &occupationScientistAllOther,
		strings.ToLower(string(occupationSeamstressTailorID)):                           &occupationSeamstressTailor,
		strings.ToLower(string(occupationSecurityScreenerID)):                           &occupationSecurityScreener,
		strings.ToLower(string(occupationShoeShinerRepairmanID)):                        &occupationShoeShinerRepairman,
		strings.ToLower(string(occupationSingerSongwriterID)):                           &occupationSingerSongwriter,
		strings.ToLower(string(occupationStaffingSpecialistID)):                         &occupationStaffingSpecialist,
		strings.ToLower(string(occupationStateExaminerID)):                              &occupationStateExaminer,
		strings.ToLower(string(occupationSuperintendentAllOtherID)):                     &occupationSuperintendentAllOther,
		strings.ToLower(string(occupationSuperintendentDrillerID)):                      &occupationSuperintendentDriller,
		strings.ToLower(string(occupationSuperintendentSchoolID)):                       &occupationSuperintendentSchool,
		strings.ToLower(string(occupationSuperintendentorSupervisorBuildingMaintID)):    &occupationSuperintendentorSupervisorBuildingMaint,
		strings.ToLower(string(occupationSupervisorAccountingID)):                       &occupationSupervisorAccounting,
		strings.ToLower(string(occupationSupervisorAllOtherDegreedID)):                  &occupationSupervisorAllOtherDegreed,
		strings.ToLower(string(occupationSupervisorDataSystemsID)):                      &occupationSupervisorDataSystems,
		strings.ToLower(string(occupationSupervisorHumanResourcePersonnelID)):           &occupationSupervisorHumanResourcePersonnel,
		strings.ToLower(string(occupationSupervisorOfficeID)):                           &occupationSupervisorOffice,
		strings.ToLower(string(occupationSupervisorOperationsID)):                       &occupationSupervisorOperations,
		strings.ToLower(string(occupationSupervisorOtherDegreedID)):                     &occupationSupervisorOtherDegreed,
		strings.ToLower(string(occupationSupervisorPostalID)):                           &occupationSupervisorPostal,
		strings.ToLower(string(occupationSupervisorProductionID)):                       &occupationSupervisorProduction,
		strings.ToLower(string(occupationSupervisorRestaurantNonFastFoodID)):            &occupationSupervisorRestaurantNonFastFood,
		strings.ToLower(string(occupationSurgeonID)):                                    &occupationSurgeon,
		strings.ToLower(string(occupationSurveyorLicensedID)):                           &occupationSurveyorLicensed,
		strings.ToLower(string(occupationSurveyorNonLicensedID)):                        &occupationSurveyorNonLicensed,
		strings.ToLower(string(occupationTaxExaminerNotClericalID)):                     &occupationTaxExaminerNotClerical,
		strings.ToLower(string(occupationTaxPreparerNotAccountantID)):                   &occupationTaxPreparerNotAccountant,
		strings.ToLower(string(occupationTaxidermistID)):                                &occupationTaxidermist,
		strings.ToLower(string(occupationTeachersorCoachesorInstructorsID)):             &occupationTeachersorCoachesorInstructors,
		strings.ToLower(string(occupationTechnicianAllOtherID)):                         &occupationTechnicianAllOther,
		strings.ToLower(string(occupationTechnicianElectricalorElectronicID)):           &occupationTechnicianElectricalorElectronic,
		strings.ToLower(string(occupationTechnicianFoodID)):                             &occupationTechnicianFood,
		strings.ToLower(string(occupationTechnicianInstrumentationID)):                  &occupationTechnicianInstrumentation,
		strings.ToLower(string(occupationTechnicianLabID)):                              &occupationTechnicianLab,
		strings.ToLower(string(occupationTechnicianMedicalID)):                          &occupationTechnicianMedical,
		strings.ToLower(string(occupationTechnicianRadiologicalID)):                     &occupationTechnicianRadiological,
		strings.ToLower(string(occupationTechnicianScienceID)):                          &occupationTechnicianScience,
		strings.ToLower(string(occupationTechnicianTestingID)):                          &occupationTechnicianTesting,
		strings.ToLower(string(occupationTechnicianUltrasoundID)):                       &occupationTechnicianUltrasound,
		strings.ToLower(string(occupationTechnicianXRayID)):                             &occupationTechnicianXRay,
		strings.ToLower(string(occupationTechnicianorAssistantEngineeringID)):           &occupationTechnicianorAssistantEngineering,
		strings.ToLower(string(occupationTelemarketerID)):                               &occupationTelemarketer,
		strings.ToLower(string(occupationTherapistID)):                                  &occupationTherapist,
		strings.ToLower(string(occupationTrainerAerobicsFitnessID)):                     &occupationTrainerAerobicsFitness,
		strings.ToLower(string(occupationTrainerAthleticNataMemberID)):                  &occupationTrainerAthleticNataMember,
		strings.ToLower(string(occupationTrainerAthleticNonNataMemberID)):               &occupationTrainerAthleticNonNataMember,
		strings.ToLower(string(occupationTrainerCaretakerAnimalID)):                     &occupationTrainerCaretakerAnimal,
		strings.ToLower(string(occupationTranslatororInterpreterID)):                    &occupationTranslatororInterpreter,
		strings.ToLower(string(occupationTreasurerID)):                                  &occupationTreasurer,
		strings.ToLower(string(occupationTutorID)):                                      &occupationTutor,
		strings.ToLower(string(occupationUnderwriterInsuranceID)):                       &occupationUnderwriterInsurance,
		strings.ToLower(string(occupationVendorID)):                                     &occupationVendor,
		strings.ToLower(string(occupationVeterinarianID)):                               &occupationVeterinarian,
		strings.ToLower(string(occupationVicePresBusinessID)):                           &occupationVicePresBusiness,
		strings.ToLower(string(occupationWaiterWaitressID)):                             &occupationWaiterWaitress,
		strings.ToLower(string(occupationWardenAllOtherID)):                             &occupationWardenAllOther,
		strings.ToLower(string(occupationWardenGameID)):                                 &occupationWardenGame,
		strings.ToLower(string(occupationWorkerMetalNotSteelID)):                        &occupationWorkerMetalNotSteel,
		strings.ToLower(string(occupationWorkerRailroadID)):                             &occupationWorkerRailroad,
		strings.ToLower(string(occupationWorkerSocialCaseID)):                           &occupationWorkerSocialCase,
		strings.ToLower(string(occupationWriterAllOtherID)):                             &occupationWriterAllOther,
		strings.ToLower(string(occupationWriterCommercialID)):                           &occupationWriterCommercial,
		strings.ToLower(string(occupationE1ID)):                                         &occupationE1,
		strings.ToLower(string(occupationE2ID)):                                         &occupationE2,
		strings.ToLower(string(occupationE3ID)):                                         &occupationE3,
		strings.ToLower(string(occupationE4ID)):                                         &occupationE4,
		strings.ToLower(string(occupationE4PID)):                                        &occupationE4P,
		strings.ToLower(string(occupationE5ID)):                                         &occupationE5,
		strings.ToLower(string(occupationE5PID)):                                        &occupationE5P,
		strings.ToLower(string(occupationE6ID)):                                         &occupationE6,
		strings.ToLower(string(occupationE6PID)):                                        &occupationE6P,
		strings.ToLower(string(occupationE7ID)):                                         &occupationE7,
		strings.ToLower(string(occupationE8ID)):                                         &occupationE8,
		strings.ToLower(string(occupationE9ID)):                                         &occupationE9,
		strings.ToLower(string(occupationO1ID)):                                         &occupationO1,
		strings.ToLower(string(occupationO10ID)):                                        &occupationO10,
		strings.ToLower(string(occupationO2ID)):                                         &occupationO2,
		strings.ToLower(string(occupationO3ID)):                                         &occupationO3,
		strings.ToLower(string(occupationO4ID)):                                         &occupationO4,
		strings.ToLower(string(occupationO5ID)):                                         &occupationO5,
		strings.ToLower(string(occupationO6ID)):                                         &occupationO6,
		strings.ToLower(string(occupationO7ID)):                                         &occupationO7,
		strings.ToLower(string(occupationO8ID)):                                         &occupationO8,
		strings.ToLower(string(occupationO9ID)):                                         &occupationO9,
		strings.ToLower(string(occupationW1ID)):                                         &occupationW1,
		strings.ToLower(string(occupationW2ID)):                                         &occupationW2,
		strings.ToLower(string(occupationW3ID)):                                         &occupationW3,
		strings.ToLower(string(occupationW4ID)):                                         &occupationW4,
		strings.ToLower(string(occupationW5ID)):                                         &occupationW5,
		strings.ToLower(string(occupationHomemakerID)):                                  &occupationHomemaker,
		strings.ToLower(string(occupationUnemployedID)):                                 &occupationUnemployed,
		strings.ToLower(string(occupationGraduateSchoolID)):                             &occupationGraduateSchool,
		strings.ToLower(string(occupationHighSchoolID)):                                 &occupationHighSchool,
		strings.ToLower(string(occupationTradeSchoolorAssociateDegreeID)):               &occupationTradeSchoolorAssociateDegree,
		strings.ToLower(string(occupationUndergraduate4yeardegreeID)):                   &occupationUndergraduate4yeardegree,
		strings.ToLower(string(occupationDisabledID)):                                   &occupationDisabled,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumOccupation) ByID(id OccupationIdentifier) *EnumOccupationItem {
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
func (e *EnumOccupation) ByIDString(idx string) *EnumOccupationItem {
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
func (e *EnumOccupation) ByIndex(idx int) *EnumOccupationItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedOccupationID is a struct that is designed to replace a *OccupationID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *OccupationID it contains while being a better JSON citizen.
type ValidatedOccupationID struct {
	// id will point to a valid OccupationID, if possible
	// If id is nil, then ValidatedOccupationID.Valid() will return false.
	id *OccupationID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedOccupationID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedOccupationID
func (vi *ValidatedOccupationID) Clone() *ValidatedOccupationID {
	if vi == nil {
		return nil
	}

	var cid *OccupationID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedOccupationID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedOccupationIds represent the same Occupation
func (vi *ValidatedOccupationID) Equals(vj *ValidatedOccupationID) bool {
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

// Valid returns true if and only if the ValidatedOccupationID corresponds to a recognized Occupation
func (vi *ValidatedOccupationID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedOccupationID) ID() *OccupationID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedOccupationID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedOccupationID) ValidatedID() *ValidatedOccupationID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedOccupationID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedOccupationID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedOccupationID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedOccupationID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedOccupationID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := OccupationID(capString)
	item := Occupation.ByID(&id)
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

func (vi ValidatedOccupationID) String() string {
	return vi.ToIDString()
}

type OccupationIdentifier interface {
	ID() *OccupationID
	Valid() bool
}
