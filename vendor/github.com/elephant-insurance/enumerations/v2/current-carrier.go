package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// CurrentCarrierID uniquely identifies a particular CurrentCarrier
type CurrentCarrierID string

// Clone creates a safe, independent copy of a CurrentCarrierID
func (i *CurrentCarrierID) Clone() *CurrentCarrierID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two CurrentCarrierIds are equivalent
func (i *CurrentCarrierID) Equals(j *CurrentCarrierID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *CurrentCarrierID that is either valid or nil
func (i *CurrentCarrierID) ID() *CurrentCarrierID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *CurrentCarrierID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the CurrentCarrierID corresponds to a recognized CurrentCarrier
func (i *CurrentCarrierID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return CurrentCarrier.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *CurrentCarrierID) ValidatedID() *ValidatedCurrentCarrierID {
	if i != nil {
		return &ValidatedCurrentCarrierID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *CurrentCarrierID) MarshalJSON() ([]byte, error) {
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

func (i *CurrentCarrierID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := CurrentCarrierID(dataString)
	item := CurrentCarrier.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	currentCarrierAAAID                     CurrentCarrierID = "AAA_Auto_Ins"
	currentCarrierAccessID                  CurrentCarrierID = "ACCESS_INSURANCE"
	currentCarrierACEGroupID                CurrentCarrierID = "ACE_Group"
	currentCarrierAffinityMutualID          CurrentCarrierID = "AFFINITY_MUT_INS"
	currentCarrierAffirmativeID             CurrentCarrierID = "AFFIRMATIVE_INS_GRP"
	currentCarrierAGICID                    CurrentCarrierID = "AGIC_INC"
	currentCarrierAIGID                     CurrentCarrierID = "AIG"
	currentCarrierAlaskaNationalID          CurrentCarrierID = "ALASKA_NATL_CORP"
	currentCarrierAlfaID                    CurrentCarrierID = "ALFA_INS_GROUP"
	currentCarrierAlleghenyID               CurrentCarrierID = "ALLEGHANY_INS_HDG_GP"
	currentCarrierAllianzID                 CurrentCarrierID = "ALLIANZ_OF_AMER_COMP"
	currentCarrierAlliedWorldID             CurrentCarrierID = "ALLIED_WORLD_GROUP"
	currentCarrierAllstateID                CurrentCarrierID = "Allstate"
	currentCarrierAmericanEuropeanID        CurrentCarrierID = "AMER_EUROPEAN_GROUP"
	currentCarrierAmericanFamilyID          CurrentCarrierID = "American_Fam_Ins"
	currentCarrierAmericanFarmAndRanchID    CurrentCarrierID = "AMER_FARM_RANCH_GR"
	currentCarrierAmericanIndependentID     CurrentCarrierID = "AMERICAN_INDEP_COS"
	currentCarrierAmericanInternationalID   CurrentCarrierID = "AMERICAN_INT_GRP"
	currentCarrierAmericanNationalID        CurrentCarrierID = "AMER_NAT_P_C_GROUP"
	currentCarrierAmericanSterlingID        CurrentCarrierID = "AMER_STERL_INS_COMP"
	currentCarrierAmeripriseID              CurrentCarrierID = "AMERIPRISE_FNC_INC"
	currentCarrierAutoClubEnterprisesID     CurrentCarrierID = "AUTO_CLUB_ENT_INS_GR"
	currentCarrierAutoClubOfFloridaID       CurrentCarrierID = "AUTO_CLUB_INS_CO_FL"
	currentCarrierAutoClubSouthID           CurrentCarrierID = "AUTO_CLUB_SOUTH_INS"
	currentCarrierAutoOwnersID              CurrentCarrierID = "Auto_Owners_Ins"
	currentCarrierAxisID                    CurrentCarrierID = "AXIS_INSURANCE_GROUP"
	currentCarrierBaldwinAndLyonsID         CurrentCarrierID = "BALDWIN_LYONS_GRP"
	currentCarrierBrethrenMutualID          CurrentCarrierID = "BRETHREN_MUT_INS_CO"
	currentCarrierCaliforniaCasualtyID      CurrentCarrierID = "CALIFORNIA_CAS_GRP"
	currentCarrierCenturyNationalID         CurrentCarrierID = "CENTURYNAT_INS_CO"
	currentCarrierChubbID                   CurrentCarrierID = "CHUBB_GRP_OF_INS_COS"
	currentCarrierChurchMutualID            CurrentCarrierID = "CHURCH_MUT_INS_CO"
	currentCarrierCivilServiceEmployeesID   CurrentCarrierID = "CIVIL_SERVICE_EMP_GR"
	currentCarrierCNAID                     CurrentCarrierID = "CNA_INS_COMPANIES"
	currentCarrierCNAFinancialGroupID       CurrentCarrierID = "CNA_Financ_Group"
	currentCarrierCompanionID               CurrentCarrierID = "COMPANION_P_C_GRP"
	currentCarrierConsumersID               CurrentCarrierID = "CONSUMERS_INS_USA"
	currentCarrierCountryFinancialID        CurrentCarrierID = "Country_Financial"
	currentCarrierCountryFinancialPCID      CurrentCarrierID = "COUNTRY_FINL_PC_GRP"
	currentCarrierDirectAutoID              CurrentCarrierID = "Direct_Auto"
	currentCarrierDirectGeneralID           CurrentCarrierID = "DIRECT_GENERAL_GROUP"
	currentCarrierDonegalGroupID            CurrentCarrierID = "Donegal_group"
	currentCarrierEInsID                    CurrentCarrierID = "E_Ins"
	currentCarrierElectricID                CurrentCarrierID = "ELECTRIC_INS_CO"
	currentCarrierElephantID                CurrentCarrierID = "Elephant"
	currentCarrierEMCID                     CurrentCarrierID = "EMC_INS_COS"
	currentCarrierErieID                    CurrentCarrierID = "Erie_Ins"
	currentCarrierEverestREID               CurrentCarrierID = "EVEREST_RE_US_GROUP"
	currentCarrierEvergreenUSAID            CurrentCarrierID = "EVRGREEN_USA_RET_GRP"
	currentCarrierFairfaxFinancialID        CurrentCarrierID = "FAIRFAX_FIN_USA_GRP"
	currentCarrierFarmBureauID              CurrentCarrierID = "FARM_BUREAU_PC_GRP"
	currentCarrierFarmersID                 CurrentCarrierID = "Farmer_Ins"
	currentCarrierFarmersMutualID           CurrentCarrierID = "FARMERS_MUT_INS_NEB"
	currentCarrierFidelityNationalID        CurrentCarrierID = "FIDELITY_NATL_PC_GP"
	currentCarrierFirstAcceptanceID         CurrentCarrierID = "FIRST_ACC_INS_GRP"
	currentCarrierFrankenmuthID             CurrentCarrierID = "FRANKENMUTH_INS_GRP"
	currentCarrierFredLoyaID                CurrentCarrierID = "Fred_Loya_Ins"
	currentCarrierGeicoID                   CurrentCarrierID = "Geico"
	currentCarrierGeorgiaFarmBureauID       CurrentCarrierID = "GEORGIA_FARM_BUR_GRP"
	currentCarrierGoodvilleAndGermanID      CurrentCarrierID = "GVILLE_GERM_MUT_POOL"
	currentCarrierGrangeID                  CurrentCarrierID = "GRANGE_INS_GROUP"
	currentCarrierGrangeMutualID            CurrentCarrierID = "GRANGE_MUT_CAS_POOL"
	currentCarrierGreatAmericanID           CurrentCarrierID = "GREAT_AMER_PC_GROUP"
	currentCarrierGuideOneID                CurrentCarrierID = "GUIDEONE_INS_COMPS"
	currentCarrierHallmarkID                CurrentCarrierID = "HALLMARK_INS_GRP"
	currentCarrierHastingsMutualID          CurrentCarrierID = "HASTINGS_MUT_INS_CO"
	currentCarrierHiscoxUSAID               CurrentCarrierID = "HISCOX_USA_GROUP"
	currentCarrierHochheimPrairieID         CurrentCarrierID = "HOCHHEIM_PRAIRIE_GRP"
	currentCarrierHomeStateID               CurrentCarrierID = "HOME_STATE_INS_GROUP"
	currentCarrierHoraceMannID              CurrentCarrierID = "HORACE_MANN_INS_GRP"
	currentCarrierIATID                     CurrentCarrierID = "IAT_INSURANCE_GROUP"
	currentCarrierIMTID                     CurrentCarrierID = "IMT_INS_COS"
	currentCarrierInfinityID                CurrentCarrierID = "INF_PROP_CAS_GRP"
	currentCarrierIntegonNationalID         CurrentCarrierID = "INTEGON_NATIONAL_GRP"
	currentCarrierIronshoreID               CurrentCarrierID = "IRONSHORE_INS_GRP"
	currentCarrierKemperID                  CurrentCarrierID = "KEMPER_PC_GROUP"
	currentCarrierKeyID                     CurrentCarrierID = "KEY_INSURANCE_CO"
	currentCarrierKingswayFinancialID       CurrentCarrierID = "Kingsway_Finan"
	currentCarrierLibertyMutualID           CurrentCarrierID = "Liberty_Mutual"
	currentCarrierMAPFREID                  CurrentCarrierID = "MAPFRE_NA_GROUP"
	currentCarrierMerchantsID               CurrentCarrierID = "MERCHANTS_INS_GROUP"
	currentCarrierMercuryGeneralID          CurrentCarrierID = "MERCURY_GENERAL_GRP"
	currentCarrierMetLifeID                 CurrentCarrierID = "MetLife"
	currentCarrierMichiganFarmBureauID      CurrentCarrierID = "MI_FARM_BUREAU_GROUP"
	currentCarrierMichiganMillersID         CurrentCarrierID = "MICH_MILLERS_MUT"
	currentCarrierMidwestFamilyMutualID     CurrentCarrierID = "MIDWEST_FAMILY_MUT"
	currentCarrierMMGID                     CurrentCarrierID = "MMG_INSURANCE_CO"
	currentCarrierMonumentID                CurrentCarrierID = "Monument"
	currentCarrierMSADUSID                  CurrentCarrierID = "MSAD_US_INS_GROUP"
	currentCarrierMunichAmericanID          CurrentCarrierID = "MUNICHAMERICAN_COS"
	currentCarrierNationalGeneralHoldingsID CurrentCarrierID = "NATL_GEN_HOLD_CORP"
	currentCarrierNationwideID              CurrentCarrierID = "Nationwide"
	currentCarrierNLCID                     CurrentCarrierID = "NLC_INSURANCE_POOL"
	currentCarrierNodakMutualID             CurrentCarrierID = "NODAK_MUTUAL_GROUP"
	currentCarrierNonProfitsID              CurrentCarrierID = "NONPROF_INS_ALL_GRP"
	currentCarrierNorfolkAndDedhamID        CurrentCarrierID = "NORFOLK_DEDHAM_PL"
	currentCarrierNYCMID                    CurrentCarrierID = "NYCM_INSURANCE_GROUP"
	currentCarrierOkalahomaFarmBureauID     CurrentCarrierID = "OKLAHOMA_F_B_GROUP"
	currentCarrierOldAmericaCapitalID       CurrentCarrierID = "OLD_AMER_CAPT_GRP"
	currentCarrierOregonMutualID            CurrentCarrierID = "OREGON_MUT_GROUP"
	currentCarrierOtherID                   CurrentCarrierID = "Other"
	currentCarrierPacificSpecialtyID        CurrentCarrierID = "PACIFIC_SPECIALTY_GR"
	currentCarrierPekinID                   CurrentCarrierID = "Pekin_Ins_Group"
	currentCarrierPennLumbermensID          CurrentCarrierID = "PENN_LUMBERMENS_MUT"
	currentCarrierPharmacistsID             CurrentCarrierID = "PHARMACISTS_INS_GRP"
	currentCarrierPhiladlphiaTokioID        CurrentCarrierID = "PHILADLPHIA_TOKIO_GP"
	currentCarrierPioneerStateMutualID      CurrentCarrierID = "PIONEER_STATE_MUT"
	currentCarrierPreferredMutualID         CurrentCarrierID = "PREFERRED_MUTUAL_GRP"
	currentCarrierProgressiveID             CurrentCarrierID = "Progressive"
	currentCarrierProsightSpecialtyID       CurrentCarrierID = "PROSIGHT_SPEC_GROUP"
	currentCarrierPublicServiceMutualID     CurrentCarrierID = "PUB_SERV_MUT_HOLD"
	currentCarrierPureID                    CurrentCarrierID = "PURE_GRP_OF_INS_COS"
	currentCarrierQBEAmericasID             CurrentCarrierID = "QBE_AMERICAS_GROUP"
	currentCarrierRAMMutualID               CurrentCarrierID = "RAM_MUT_INS_CO"
	currentCarrierRepublicGroupID           CurrentCarrierID = "REPUBLIC_GROUP"
	currentCarrierRLIID                     CurrentCarrierID = "RLI_GROUP"
	currentCarrierSafeAutoID                CurrentCarrierID = "Safeauto"
	currentCarrierSamsungFireAndMarineID    CurrentCarrierID = "SAMSUNG_FIREMAR_USB"
	currentCarrierSecuraID                  CurrentCarrierID = "SECURA_INS_COMPANIES"
	currentCarrierSentryID                  CurrentCarrierID = "Sentry_Ins"
	currentCarrierSompoJapanID              CurrentCarrierID = "SOMPO_JAPAN_US_GROUP"
	currentCarrierSouthportLaneID           CurrentCarrierID = "SOUTHPORT_LANE_GROUP"
	currentCarrierStarAndShieldID           CurrentCarrierID = "STAR_SHIELD_INS_EX"
	currentCarrierStarCasualtyID            CurrentCarrierID = "STAR_CASUALTY_INS_CO"
	currentCarrierStarrInternationalID      CurrentCarrierID = "STARR_INT_COMP_INC"
	currentCarrierStateAutoID               CurrentCarrierID = "State_Auto"
	currentCarrierStateFarmID               CurrentCarrierID = "State_Farm"
	currentCarrierStateNationalID           CurrentCarrierID = "STATE_NATIONAL_GROUP"
	currentCarrierSterlingID                CurrentCarrierID = "STERLING_INS_CO"
	currentCarrierTheGeneralID              CurrentCarrierID = "The_General"
	currentCarrierTheHartfordID             CurrentCarrierID = "The_Hartford"
	currentCarrierTheRepublicGroupID        CurrentCarrierID = "The_Republic_group"
	currentCarrierTitusID                   CurrentCarrierID = "TITUS_GROUP"
	currentCarrierTopaID                    CurrentCarrierID = "TOPA_INSURANCE_GROUP"
	currentCarrierTowerID                   CurrentCarrierID = "TOWER_GRP_COMPANIES"
	currentCarrierTradersGeneralID          CurrentCarrierID = "TRADERS_GEN_AGENCY"
	currentCarrierTravelersID               CurrentCarrierID = "Travelers"
	currentCarrierUnitedHomeID              CurrentCarrierID = "UNITED_HOME_INS_COMP"
	currentCarrierUnitrinID                 CurrentCarrierID = "Unitrin"
	currentCarrierUSAAID                    CurrentCarrierID = "USAA"
	currentCarrierUticaMutualID             CurrentCarrierID = "UTICA_MUT_INS_COMP"
	currentCarrierVAFarmBureauID            CurrentCarrierID = "VA_Farm_Bureau"
	currentCarrierWawanesaID                CurrentCarrierID = "WAWANESA_GENERAL_INS"
	currentCarrierWestBendMutualID          CurrentCarrierID = "WEST_BEND_MUT_INS_CO"
	currentCarrierWesternGeneralID          CurrentCarrierID = "WESTERN_GENERAL_INS"
	currentCarrierWesternReservePoolID      CurrentCarrierID = "WESTERN_RESERVE_POOL"
	currentCarrierWestfieldID               CurrentCarrierID = "WESTFIELD_GROUP"
	currentCarrierWhiteMountainsID          CurrentCarrierID = "WHITE_MTNS_INS_GRP"
	currentCarrierWisconsinMutualID         CurrentCarrierID = "WISCONSIN_MUT_INS_CO"
	currentCarrierWolverineMutualID         CurrentCarrierID = "WOLVERINE_MUT_INS_CO"
	currentCarrierWorkmensAutoID            CurrentCarrierID = "WORKMENS_AUTO_INS"
	currentCarrierWRMAmericaID              CurrentCarrierID = "WRM_AMERICA_GROUP"
	currentCarrierZurichInsuranceID         CurrentCarrierID = "ZURICH_INS_GRP_LTD"
)

// EnumCurrentCarrierItem describes an entry in an enumeration of CurrentCarrier
type EnumCurrentCarrierItem struct {
	ID        CurrentCarrierID  `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	AmBest      string
	DisplayonUI string
	Parent      string
}

var (
	currentCarrierAAA                     = EnumCurrentCarrierItem{currentCarrierAAAID, "AAA Auto Insurance", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "AAA", 1, "[]", "false", ""}
	currentCarrierAccess                  = EnumCurrentCarrierItem{currentCarrierAccessID, "Access Insurance Co", map[string]string{"AmBest": "[11651]", "DisplayonUI": "false", "Parent": "ACCESS_INSURANCE"}, "Access", 2, "[11651]", "false", "ACCESS_INSURANCE"}
	currentCarrierACEGroup                = EnumCurrentCarrierItem{currentCarrierACEGroupID, "ACE Group", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "ACEGroup", 3, "[]", "false", ""}
	currentCarrierAffinityMutual          = EnumCurrentCarrierItem{currentCarrierAffinityMutualID, "Affinity Mutual Insurance", map[string]string{"AmBest": "[4811]", "DisplayonUI": "false", "Parent": "AFFINITY_MUT_INS"}, "AffinityMutual", 4, "[4811]", "false", "AFFINITY_MUT_INS"}
	currentCarrierAffirmative             = EnumCurrentCarrierItem{currentCarrierAffirmativeID, "Affirmative Insurance Group", map[string]string{"AmBest": "[1795]", "DisplayonUI": "false", "Parent": "AFFIRMATIVE_INS_GRP"}, "Affirmative", 5, "[1795]", "false", "AFFIRMATIVE_INS_GRP"}
	currentCarrierAGIC                    = EnumCurrentCarrierItem{currentCarrierAGICID, "Agic Inc", map[string]string{"AmBest": "[14128]", "DisplayonUI": "false", "Parent": "AGIC_INC"}, "AGIC", 6, "[14128]", "false", "AGIC_INC"}
	currentCarrierAIG                     = EnumCurrentCarrierItem{currentCarrierAIGID, "AIG", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "AIG", 7, "[]", "false", ""}
	currentCarrierAlaskaNational          = EnumCurrentCarrierItem{currentCarrierAlaskaNationalID, "Alaska National Corporationc", map[string]string{"AmBest": "[2648]", "DisplayonUI": "false", "Parent": "ALASKA_NATL_CORP"}, "AlaskaNational", 8, "[2648]", "false", "ALASKA_NATL_CORP"}
	currentCarrierAlfa                    = EnumCurrentCarrierItem{currentCarrierAlfaID, "Alfa Insurance Group", map[string]string{"AmBest": "[12333, 10042]", "DisplayonUI": "false", "Parent": "ALFA_INS_GROUP"}, "Alfa", 9, "[12333, 10042]", "false", "ALFA_INS_GROUP"}
	currentCarrierAllegheny               = EnumCurrentCarrierItem{currentCarrierAlleghenyID, "Alleghany Insurance Holdings Group", map[string]string{"AmBest": "[235]", "DisplayonUI": "false", "Parent": "ALLEGHANY_INS_HDG_GP"}, "Allegheny", 10, "[235]", "false", "ALLEGHANY_INS_HDG_GP"}
	currentCarrierAllianz                 = EnumCurrentCarrierItem{currentCarrierAllianzID, "Allianz Of America Companies", map[string]string{"AmBest": "[3682]", "DisplayonUI": "false", "Parent": "ALLIANZ_OF_AMER_COMP"}, "Allianz", 11, "[3682]", "false", "ALLIANZ_OF_AMER_COMP"}
	currentCarrierAlliedWorld             = EnumCurrentCarrierItem{currentCarrierAlliedWorldID, "Allied World Group", map[string]string{"AmBest": "[12525, 13865, 12526, 12699]", "DisplayonUI": "false", "Parent": "ALLIED_WORLD_GROUP"}, "AlliedWorld", 12, "[12525, 13865, 12526, 12699]", "false", "ALLIED_WORLD_GROUP"}
	currentCarrierAllstate                = EnumCurrentCarrierItem{currentCarrierAllstateID, "Allstate", map[string]string{"AmBest": "[542, 764, 1978, 2017, 2018, 3652, 3791, 11559, 11702, 11703, 11794, 12106, 12535, 12536, 12612, 12710, 12711, 13080, 13082]", "DisplayonUI": "true", "Parent": "ALLSTATE_INSURANCE_GROUP"}, "Allstate", 13, "[542, 764, 1978, 2017, 2018, 3652, 3791, 11559, 11702, 11703, 11794, 12106, 12535, 12536, 12612, 12710, 12711, 13080, 13082]", "true", "ALLSTATE_INSURANCE_GROUP"}
	currentCarrierAmericanEuropean        = EnumCurrentCarrierItem{currentCarrierAmericanEuropeanID, "American European Group", map[string]string{"AmBest": "[2317]", "DisplayonUI": "false", "Parent": "AMER_EUROPEAN_GROUP"}, "AmericanEuropean", 14, "[2317]", "false", "AMER_EUROPEAN_GROUP"}
	currentCarrierAmericanFamily          = EnumCurrentCarrierItem{currentCarrierAmericanFamilyID, "American Family Insurance", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "AmericanFamily", 15, "[]", "false", ""}
	currentCarrierAmericanFarmAndRanch    = EnumCurrentCarrierItem{currentCarrierAmericanFarmAndRanchID, "American Farm & Ranch Group", map[string]string{"AmBest": "[1726]", "DisplayonUI": "false", "Parent": "AMER_FARM_RANCH_GR"}, "AmericanFarmAndRanch", 16, "[1726]", "false", "AMER_FARM_RANCH_GR"}
	currentCarrierAmericanIndependent     = EnumCurrentCarrierItem{currentCarrierAmericanIndependentID, "American Independent Insurance", map[string]string{"AmBest": "[10831, 10642]", "DisplayonUI": "false", "Parent": "AMERICAN_INDEP_COS"}, "AmericanIndependent", 17, "[10831, 10642]", "false", "AMERICAN_INDEP_COS"}
	currentCarrierAmericanInternational   = EnumCurrentCarrierItem{currentCarrierAmericanInternationalID, "American International Group", map[string]string{"AmBest": "[2389, 2034, 10255, 2349, 4000, 2035, 2351, 2363]", "DisplayonUI": "false", "Parent": "AMERICAN_INT_GRP"}, "AmericanInternational", 18, "[2389, 2034, 10255, 2349, 4000, 2035, 2351, 2363]", "false", "AMERICAN_INT_GRP"}
	currentCarrierAmericanNational        = EnumCurrentCarrierItem{currentCarrierAmericanNationalID, "American National P & C GROUP", map[string]string{"AmBest": "[2803, 11700, 3533, 12472, 362, 12411]", "DisplayonUI": "false", "Parent": "AMER_NAT_P_C_GROUP"}, "AmericanNational", 19, "[2803, 11700, 3533, 12472, 362, 12411]", "false", "AMER_NAT_P_C_GROUP"}
	currentCarrierAmericanSterling        = EnumCurrentCarrierItem{currentCarrierAmericanSterlingID, "American Sterling Insurance Company", map[string]string{"AmBest": "[1915]", "DisplayonUI": "false", "Parent": "AMER_STERL_INS_COMP"}, "AmericanSterling", 20, "[1915]", "false", "AMER_STERL_INS_COMP"}
	currentCarrierAmeriprise              = EnumCurrentCarrierItem{currentCarrierAmeripriseID, "Ameriprise Financial", map[string]string{"AmBest": "[13104, 3563]", "DisplayonUI": "false", "Parent": "AMERIPRISE_FNC_INC"}, "Ameriprise", 21, "[13104, 3563]", "false", "AMERIPRISE_FNC_INC"}
	currentCarrierAutoClubEnterprises     = EnumCurrentCarrierItem{currentCarrierAutoClubEnterprisesID, "Auto Club Enterprises Insurance Group", map[string]string{"AmBest": "[185, 186, 1758, 4089, 2139, 650, 405, 4435,2140]", "DisplayonUI": "false", "Parent": "AUTO_CLUB_ENT_INS_GR"}, "AutoClubEnterprises", 22, "[185, 186, 1758, 4089, 2139, 650, 405, 4435,2140]", "false", "AUTO_CLUB_ENT_INS_GR"}
	currentCarrierAutoClubOfFlorida       = EnumCurrentCarrierItem{currentCarrierAutoClubOfFloridaID, "Auto Club Insurance Company of Florida", map[string]string{"AmBest": "[13083]", "DisplayonUI": "false", "Parent": "AUTO_CLUB_INS_CO_FL"}, "AutoClubOfFlorida", 23, "[13083]", "false", "AUTO_CLUB_INS_CO_FL"}
	currentCarrierAutoClubSouth           = EnumCurrentCarrierItem{currentCarrierAutoClubSouthID, "Auto Club South Insurance", map[string]string{"AmBest": "[11209]", "DisplayonUI": "false", "Parent": "AUTO_CLUB_SOUTH_INS"}, "AutoClubSouth", 24, "[11209]", "false", "AUTO_CLUB_SOUTH_INS"}
	currentCarrierAutoOwners              = EnumCurrentCarrierItem{currentCarrierAutoOwnersID, "Auto-Owners Insurance", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "AutoOwners", 25, "[]", "false", ""}
	currentCarrierAxis                    = EnumCurrentCarrierItem{currentCarrierAxisID, "Axis Insurance Group", map[string]string{"AmBest": "[13034, 12515]", "DisplayonUI": "false", "Parent": "AXIS_INSURANCE_GROUP"}, "Axis", 26, "[13034, 12515]", "false", "AXIS_INSURANCE_GROUP"}
	currentCarrierBaldwinAndLyons         = EnumCurrentCarrierItem{currentCarrierBaldwinAndLyonsID, "Baldwin & Lyons Group", map[string]string{"AmBest": "[784, 1840]", "DisplayonUI": "false", "Parent": "BALDWIN_LYONS_GRP"}, "BaldwinAndLyons", 27, "[784, 1840]", "false", "BALDWIN_LYONS_GRP"}
	currentCarrierBrethrenMutual          = EnumCurrentCarrierItem{currentCarrierBrethrenMutualID, "Brethren Mutual Insurance Co", map[string]string{"AmBest": "[220]", "DisplayonUI": "false", "Parent": "BRETHREN_MUT_INS_CO"}, "BrethrenMutual", 28, "[220]", "false", "BRETHREN_MUT_INS_CO"}
	currentCarrierCaliforniaCasualty      = EnumCurrentCarrierItem{currentCarrierCaliforniaCasualtyID, "California Casualty Insurance Co", map[string]string{"AmBest": "[3576, 3809, 222, 3336]", "DisplayonUI": "false", "Parent": "CALIFORNIA_CAS_GRP"}, "CaliforniaCasualty", 29, "[3576, 3809, 222, 3336]", "false", "CALIFORNIA_CAS_GRP"}
	currentCarrierCenturyNational         = EnumCurrentCarrierItem{currentCarrierCenturyNationalID, "Century-National Insurance Co", map[string]string{"AmBest": "[3090]", "DisplayonUI": "false", "Parent": "CENTURYNAT_INS_CO"}, "CenturyNational", 30, "[3090]", "false", "CENTURYNAT_INS_CO"}
	currentCarrierChubb                   = EnumCurrentCarrierItem{currentCarrierChubbID, "Chubb Group Of Insurance COS", map[string]string{"AmBest": "[11578, 3566, 11560, 2084, 2085,2385, 2086]", "DisplayonUI": "false", "Parent": "CHUBB_GRP_OF_INS_COS"}, "Chubb", 31, "[11578, 3566, 11560, 2084, 2085,2385, 2086]", "false", "CHUBB_GRP_OF_INS_COS"}
	currentCarrierChurchMutual            = EnumCurrentCarrierItem{currentCarrierChurchMutualID, "Church Mutual Insurance Company", map[string]string{"AmBest": "[259]", "DisplayonUI": "false", "Parent": "CHURCH_MUT_INS_CO"}, "ChurchMutual", 32, "[259]", "false", "CHURCH_MUT_INS_CO"}
	currentCarrierCivilServiceEmployees   = EnumCurrentCarrierItem{currentCarrierCivilServiceEmployeesID, "Civil Service Emp Group", map[string]string{"AmBest": "[274, 1963]", "DisplayonUI": "false", "Parent": "CIVIL_SERVICE_EMP_GR"}, "CivilServiceEmployees", 33, "[274, 1963]", "false", "CIVIL_SERVICE_EMP_GR"}
	currentCarrierCNA                     = EnumCurrentCarrierItem{currentCarrierCNAID, "Cna Insurance companies", map[string]string{"AmBest": "[2118]", "DisplayonUI": "false", "Parent": "CNA_INS_COMPANIES"}, "CNA", 34, "[2118]", "false", "CNA_INS_COMPANIES"}
	currentCarrierCNAFinancialGroup       = EnumCurrentCarrierItem{currentCarrierCNAFinancialGroupID, "CNA Financial Group", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "CNAFinancialGroup", 35, "[]", "false", ""}
	currentCarrierCompanion               = EnumCurrentCarrierItem{currentCarrierCompanionID, "Companion P and C Group", map[string]string{"AmBest": "[12069, 1979]", "DisplayonUI": "false", "Parent": "COMPANION_P_C_GRP"}, "Companion", 36, "[12069, 1979]", "false", "COMPANION_P_C_GRP"}
	currentCarrierConsumers               = EnumCurrentCarrierItem{currentCarrierConsumersID, "Consumers Insurance USA", map[string]string{"AmBest": "[11775]", "DisplayonUI": "false", "Parent": "CONSUMERS_INS_USA"}, "Consumers", 37, "[11775]", "false", "CONSUMERS_INS_USA"}
	currentCarrierCountryFinancial        = EnumCurrentCarrierItem{currentCarrierCountryFinancialID, "Country Financial", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "CountryFinancial", 38, "[]", "false", ""}
	currentCarrierCountryFinancialPC      = EnumCurrentCarrierItem{currentCarrierCountryFinancialPCID, "Country Finl PC Group", map[string]string{"AmBest": "[2248, 2249, 2250]", "DisplayonUI": "false", "Parent": "COUNTRY_FINL_PC_GRP"}, "CountryFinancialPC", 39, "[2248, 2249, 2250]", "false", "COUNTRY_FINL_PC_GRP"}
	currentCarrierDirectAuto              = EnumCurrentCarrierItem{currentCarrierDirectAutoID, "Direct Auto", map[string]string{"AmBest": "[681, 11320, 11947, 12040, 12130]", "DisplayonUI": "false", "Parent": ""}, "DirectAuto", 40, "[681, 11320, 11947, 12040, 12130]", "false", ""}
	currentCarrierDirectGeneral           = EnumCurrentCarrierItem{currentCarrierDirectGeneralID, "Direct General Group", map[string]string{"AmBest": "[12040, 11947, 681, 11320, 12130]", "DisplayonUI": "false", "Parent": "DIRECT_GENERAL_GROUP"}, "DirectGeneral", 41, "[12040, 11947, 681, 11320, 12130]", "false", "DIRECT_GENERAL_GROUP"}
	currentCarrierDonegalGroup            = EnumCurrentCarrierItem{currentCarrierDonegalGroupID, "Donegal Group", map[string]string{"AmBest": "[2223, 318, 556, 12137, 12137, 4205, 831, 4200, 850]", "DisplayonUI": "true", "Parent": "DONEGAL_INSURANCE_GROUP"}, "DonegalGroup", 42, "[2223, 318, 556, 12137, 12137, 4205, 831, 4200, 850]", "true", "DONEGAL_INSURANCE_GROUP"}
	currentCarrierEIns                    = EnumCurrentCarrierItem{currentCarrierEInsID, "E-Surance", map[string]string{"AmBest": "[273, 466, 2526]", "DisplayonUI": "false", "Parent": ""}, "EIns", 43, "[273, 466, 2526]", "false", ""}
	currentCarrierElectric                = EnumCurrentCarrierItem{currentCarrierElectricID, "Electric Insurance Company", map[string]string{"AmBest": "[2146]", "DisplayonUI": "false", "Parent": "ELECTRIC_INS_CO"}, "Electric", 44, "[2146]", "false", "ELECTRIC_INS_CO"}
	currentCarrierElephant                = EnumCurrentCarrierItem{currentCarrierElephantID, "Elephant Insurance", map[string]string{"AmBest": "[80091, 14111]", "DisplayonUI": "true", "Parent": "ADMIRAL_GROUP_PLC"}, "Elephant", 45, "[80091, 14111]", "true", "ADMIRAL_GROUP_PLC"}
	currentCarrierEMC                     = EnumCurrentCarrierItem{currentCarrierEMCID, "EMC Insurance Companies", map[string]string{"AmBest": "[311, 2039, 2160, 2161, 448, 3638, 2346]", "DisplayonUI": "false", "Parent": "EMC_INS_COS"}, "EMC", 46, "[311, 2039, 2160, 2161, 448, 3638, 2346]", "false", "EMC_INS_COS"}
	currentCarrierErie                    = EnumCurrentCarrierItem{currentCarrierErieID, "Erie Insurance", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "Erie", 47, "[]", "false", ""}
	currentCarrierEverestRE               = EnumCurrentCarrierItem{currentCarrierEverestREID, "Everest RE US Group", map[string]string{"AmBest": "[11197]", "DisplayonUI": "false", "Parent": "EVEREST_RE_US_GROUP"}, "EverestRE", 48, "[11197]", "false", "EVEREST_RE_US_GROUP"}
	currentCarrierEvergreenUSA            = EnumCurrentCarrierItem{currentCarrierEvergreenUSAID, "Evergreen USA Risk Retention Group Inc", map[string]string{"AmBest": "[11155]", "DisplayonUI": "false", "Parent": "EVRGREEN_USA_RET_GRP"}, "EvergreenUSA", 49, "[11155]", "false", "EVRGREEN_USA_RET_GRP"}
	currentCarrierFairfaxFinancial        = EnumCurrentCarrierItem{currentCarrierFairfaxFinancialID, "FairFax Financial (USA) Group", map[string]string{"AmBest": "[3081, 984]", "DisplayonUI": "false", "Parent": "FAIRFAX_FIN_USA_GRP"}, "FairfaxFinancial", 50, "[3081, 984]", "false", "FAIRFAX_FIN_USA_GRP"}
	currentCarrierFarmBureau              = EnumCurrentCarrierItem{currentCarrierFarmBureauID, "Farm Bureau Property & Casualty Group", map[string]string{"AmBest": "[354, 3600]", "DisplayonUI": "false", "Parent": "FARM_BUREAU_PC_GRP"}, "FarmBureau", 51, "[354, 3600]", "false", "FARM_BUREAU_PC_GRP"}
	currentCarrierFarmers                 = EnumCurrentCarrierItem{currentCarrierFarmersID, "Farmers Insurance", map[string]string{"AmBest": "[11587, 11109, 10614, 3247, 3786, 3641, 2359, 12021, 2796, 577, 11402, 12039, 12404, 12461, 13761, 11778, 177, 3100, 4158, 4159, 4305, 4306, 2171, 2577, 12089, 4097, 3634, 2189, 4083, 2173, 270, 10796, 3312, 2174]", "DisplayonUI": "true", "Parent": "FARMERS_INS_GROUP"}, "Farmers", 52, "[11587, 11109, 10614, 3247, 3786, 3641, 2359, 12021, 2796, 577, 11402, 12039, 12404, 12461, 13761, 11778, 177, 3100, 4158, 4159, 4305, 4306, 2171, 2577, 12089, 4097, 3634, 2189, 4083, 2173, 270, 10796, 3312, 2174]", "true", "FARMERS_INS_GROUP"}
	currentCarrierFarmersMutual           = EnumCurrentCarrierItem{currentCarrierFarmersMutualID, "Farmers Mutal Insurance NEB", map[string]string{"AmBest": "[371]", "DisplayonUI": "false", "Parent": "FARMERS_MUT_INS_NEB"}, "FarmersMutual", 53, "[371]", "false", "FARMERS_MUT_INS_NEB"}
	currentCarrierFidelityNational        = EnumCurrentCarrierItem{currentCarrierFidelityNationalID, "Fidelity Natinal P&C Group", map[string]string{"AmBest": "[12478, 4496]", "DisplayonUI": "false", "Parent": "FIDELITY_NATL_PC_GP"}, "FidelityNational", 54, "[12478, 4496]", "false", "FIDELITY_NATL_PC_GP"}
	currentCarrierFirstAcceptance         = EnumCurrentCarrierItem{currentCarrierFirstAcceptanceID, "First Acceptance Insurance Group", map[string]string{"AmBest": "[13595, 12544, 11832]", "DisplayonUI": "false", "Parent": "FIRST_ACC_INS_GRP"}, "FirstAcceptance", 55, "[13595, 12544, 11832]", "false", "FIRST_ACC_INS_GRP"}
	currentCarrierFrankenmuth             = EnumCurrentCarrierItem{currentCarrierFrankenmuthID, "Frankenmuth Insurance group", map[string]string{"AmBest": "[402, 11495]", "DisplayonUI": "false", "Parent": "FRANKENMUTH_INS_GRP"}, "Frankenmuth", 56, "[402, 11495]", "false", "FRANKENMUTH_INS_GRP"}
	currentCarrierFredLoya                = EnumCurrentCarrierItem{currentCarrierFredLoyaID, "Fred Loya Insurance", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "FredLoya", 57, "[]", "false", ""}
	currentCarrierGeico                   = EnumCurrentCarrierItem{currentCarrierGeicoID, "Geico", map[string]string{"AmBest": "[1737, 1852, 2204, 2205, 14301, 14302, 14303, 22059, 93021, 864, 646, 2540, 2541, 3736]", "DisplayonUI": "true", "Parent": "BERKSHIRE_HATHAWAY"}, "Geico", 58, "[1737, 1852, 2204, 2205, 14301, 14302, 14303, 22059, 93021, 864, 646, 2540, 2541, 3736]", "true", "BERKSHIRE_HATHAWAY"}
	currentCarrierGeorgiaFarmBureau       = EnumCurrentCarrierItem{currentCarrierGeorgiaFarmBureauID, "Georgia Farm Bureau Group", map[string]string{"AmBest": "[10746, 412]", "DisplayonUI": "false", "Parent": "GEORGIA_FARM_BUR_GRP"}, "GeorgiaFarmBureau", 59, "[10746, 412]", "false", "GEORGIA_FARM_BUR_GRP"}
	currentCarrierGoodvilleAndGerman      = EnumCurrentCarrierItem{currentCarrierGoodvilleAndGermanID, "GoodVille & German Mutual Pool", map[string]string{"AmBest": "[10510]", "DisplayonUI": "false", "Parent": "GVILLE_GERM_MUT_POOL"}, "GoodvilleAndGerman", 60, "[10510]", "false", "GVILLE_GERM_MUT_POOL"}
	currentCarrierGrange                  = EnumCurrentCarrierItem{currentCarrierGrangeID, "Grange Insurance Group", map[string]string{"AmBest": "[2208]", "DisplayonUI": "false", "Parent": "GRANGE_INS_GROUP"}, "Grange", 61, "[2208]", "false", "GRANGE_INS_GROUP"}
	currentCarrierGrangeMutual            = EnumCurrentCarrierItem{currentCarrierGrangeMutualID, "Grange Mutual Casualty Pool", map[string]string{"AmBest": "[11777, 12470, 422, 12717, 516, 13841, 10778]", "DisplayonUI": "false", "Parent": "GRANGE_MUT_CAS_POOL"}, "GrangeMutual", 62, "[11777, 12470, 422, 12717, 516, 13841, 10778]", "false", "GRANGE_MUT_CAS_POOL"}
	currentCarrierGreatAmerican           = EnumCurrentCarrierItem{currentCarrierGreatAmericanID, "Great amer P&C Group", map[string]string{"AmBest": "[3521, 2004, 11873, 3837, 3293, 2213, 2210, 10937, 173, 10618]", "DisplayonUI": "false", "Parent": "GREAT_AMER_PC_GROUP"}, "GreatAmerican", 63, "[3521, 2004, 11873, 3837, 3293, 2213, 2210, 10937, 173, 10618]", "false", "GREAT_AMER_PC_GROUP"}
	currentCarrierGuideOne                = EnumCurrentCarrierItem{currentCarrierGuideOneID, "GuideOne Insurance Companies", map[string]string{"AmBest": "[1854, 1870, 2404, 2403]", "DisplayonUI": "false", "Parent": "GUIDEONE_INS_COMPS"}, "GuideOne", 64, "[1854, 1870, 2404, 2403]", "false", "GUIDEONE_INS_COMPS"}
	currentCarrierHallmark                = EnumCurrentCarrierItem{currentCarrierHallmarkID, "Hallmark Insurance Group", map[string]string{"AmBest": "[11746, 11212, 1728, 264, 12023, 11747, 11679, 10445, 10612, 14154, 10784, 2225, 12626, 2226]", "DisplayonUI": "false", "Parent": "HALLMARK_INS_GRP"}, "Hallmark", 65, "[11746, 11212, 1728, 264, 12023, 11747, 11679, 10445, 10612, 14154, 10784, 2225, 12626, 2226]", "false", "HALLMARK_INS_GRP"}
	currentCarrierHastingsMutual          = EnumCurrentCarrierItem{currentCarrierHastingsMutualID, "Hastings Mutual Insurance Company", map[string]string{"AmBest": "[464]", "DisplayonUI": "false", "Parent": "HASTINGS_MUT_INS_CO"}, "HastingsMutual", 66, "[464]", "false", "HASTINGS_MUT_INS_CO"}
	currentCarrierHiscoxUSA               = EnumCurrentCarrierItem{currentCarrierHiscoxUSAID, "Hiscox USA Group", map[string]string{"AmBest": "[3030]", "DisplayonUI": "false", "Parent": "HISCOX_USA_GROUP"}, "HiscoxUSA", 67, "[3030]", "false", "HISCOX_USA_GROUP"}
	currentCarrierHochheimPrairie         = EnumCurrentCarrierItem{currentCarrierHochheimPrairieID, "Hochheim Prairie Group", map[string]string{"AmBest": "[3817]", "DisplayonUI": "false", "Parent": "HOCHHEIM_PRAIRIE_GRP"}, "HochheimPrairie", 68, "[3817]", "false", "HOCHHEIM_PRAIRIE_GRP"}
	currentCarrierHomeState               = EnumCurrentCarrierItem{currentCarrierHomeStateID, "Home State Insurance Group", map[string]string{"AmBest": "[10311]", "DisplayonUI": "false", "Parent": "HOME_STATE_INS_GROUP"}, "HomeState", 69, "[10311]", "false", "HOME_STATE_INS_GROUP"}
	currentCarrierHoraceMann              = EnumCurrentCarrierItem{currentCarrierHoraceMannID, "Horace Mann Insurance Group", map[string]string{"AmBest": "[3359, 12314, 4148, 884]", "DisplayonUI": "false", "Parent": "HORACE_MANN_INS_GRP"}, "HoraceMann", 70, "[3359, 12314, 4148, 884]", "false", "HORACE_MANN_INS_GRP"}
	currentCarrierIAT                     = EnumCurrentCarrierItem{currentCarrierIATID, "IAT Insurance Group", map[string]string{"AmBest": "[10611, 960, 2312, 327, 975]", "DisplayonUI": "false", "Parent": "IAT_INSURANCE_GROUP"}, "IAT", 71, "[10611, 960, 2312, 327, 975]", "false", "IAT_INSURANCE_GROUP"}
	currentCarrierIMT                     = EnumCurrentCarrierItem{currentCarrierIMTID, "IMT Insurance Group", map[string]string{"AmBest": "[530, 13117]", "DisplayonUI": "false", "Parent": "IMT_INS_COS"}, "IMT", 72, "[530, 13117]", "false", "IMT_INS_COS"}
	currentCarrierInfinity                = EnumCurrentCarrierItem{currentCarrierInfinityID, "Infinity Property & Casualty Group", map[string]string{"AmBest": "[11334, 2515, 555, 4661, 3572, 11669, 2217, 11745, 12288, 4941, 2710, 11252, 843]", "DisplayonUI": "false", "Parent": "INF_PROP_CAS_GRP"}, "Infinity", 73, "[11334, 2515, 555, 4661, 3572, 11669, 2217, 11745, 12288, 4941, 2710, 11252, 843]", "false", "INF_PROP_CAS_GRP"}
	currentCarrierIntegonNational         = EnumCurrentCarrierItem{currentCarrierIntegonNationalID, "Integon National Group", map[string]string{"AmBest": "[2387, 2669, 3366]", "DisplayonUI": "false", "Parent": "INTEGON_NATIONAL_GRP"}, "IntegonNational", 74, "[2387, 2669, 3366]", "false", "INTEGON_NATIONAL_GRP"}
	currentCarrierIronshore               = EnumCurrentCarrierItem{currentCarrierIronshoreID, "IronShore Insurance group", map[string]string{"AmBest": "[13847, 13866]", "DisplayonUI": "false", "Parent": "IRONSHORE_INS_GRP"}, "Ironshore", 75, "[13847, 13866]", "false", "IRONSHORE_INS_GRP"}
	currentCarrierKemper                  = EnumCurrentCarrierItem{currentCarrierKemperID, "Kemper P&C Group", map[string]string{"AmBest": "[2634, 10419, 391, 12213, 3596, 2701, 12149, 11946, 3045, 609, 2523, 12560, 11055, 11762, 12212, 12561, 3289, 1873, 11979, 2028]", "DisplayonUI": "false", "Parent": "KEMPER_PC_GROUP"}, "Kemper", 76, "[2634, 10419, 391, 12213, 3596, 2701, 12149, 11946, 3045, 609, 2523, 12560, 11055, 11762, 12212, 12561, 3289, 1873, 11979, 2028]", "false", "KEMPER_PC_GROUP"}
	currentCarrierKey                     = EnumCurrentCarrierItem{currentCarrierKeyID, "Key Insurance Company", map[string]string{"AmBest": "[13897]", "DisplayonUI": "false", "Parent": "KEY_INSURANCE_CO"}, "Key", 77, "[13897]", "false", "KEY_INSURANCE_CO"}
	currentCarrierKingswayFinancial       = EnumCurrentCarrierItem{currentCarrierKingswayFinancialID, "Kingsway Financial", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "KingswayFinancial", 78, "[]", "false", ""}
	currentCarrierLibertyMutual           = EnumCurrentCarrierItem{currentCarrierLibertyMutualID, "Liberty Mutual", map[string]string{"AmBest": "[588, 1812, 1896, 2087, 2253, 2282, 2283, 2286, 2287, 2289, 2290, 2446, 2447, 2448, 2709, 3028, 3610, 3794, 3795, 4240, 4938, 10764, 10765, 12053, 12425, 12569, 80115]", "DisplayonUI": "true", "Parent": "LIBERTY_MUTUAL_INSURANCE_COMPANIES"}, "LibertyMutual", 79, "[588, 1812, 1896, 2087, 2253, 2282, 2283, 2286, 2287, 2289, 2290, 2446, 2447, 2448, 2709, 3028, 3610, 3794, 3795, 4240, 4938, 10764, 10765, 12053, 12425, 12569, 80115]", "true", "LIBERTY_MUTUAL_INSURANCE_COMPANIES"}
	currentCarrierMAPFRE                  = EnumCurrentCarrierItem{currentCarrierMAPFREID, "MAPFRE North America Group", map[string]string{"AmBest": "[3131, 2365, 10805]", "DisplayonUI": "false", "Parent": "MAPFRE_NA_GROUP"}, "MAPFRE", 80, "[3131, 2365, 10805]", "false", "MAPFRE_NA_GROUP"}
	currentCarrierMerchants               = EnumCurrentCarrierItem{currentCarrierMerchantsID, "Merchants Insurance Group", map[string]string{"AmBest": "[2316, 13775]", "DisplayonUI": "false", "Parent": "MERCHANTS_INS_GROUP"}, "Merchants", 81, "[2316, 13775]", "false", "MERCHANTS_INS_GROUP"}
	currentCarrierMercuryGeneral          = EnumCurrentCarrierItem{currentCarrierMercuryGeneralID, "Mercury General Group", map[string]string{"AmBest": "[3371, 231, 2646, 233, 595, 4286, 12489, 11564, 3574, 12490, 10786, 10787, 11279]", "DisplayonUI": "false", "Parent": "MERCURY_GENERAL_GRP"}, "MercuryGeneral", 82, "[3371, 231, 2646, 233, 595, 4286, 12489, 11564, 3574, 12490, 10786, 10787, 11279]", "false", "MERCURY_GENERAL_GRP"}
	currentCarrierMetLife                 = EnumCurrentCarrierItem{currentCarrierMetLifeID, "MetLife", map[string]string{"AmBest": "[2276, 2496, 2761, 2866, 3288, 3733, 3831, 4675, 11417]", "DisplayonUI": "true", "Parent": "METLIFE_PERSONAL_GRP"}, "MetLife", 83, "[2276, 2496, 2761, 2866, 3288, 3733, 3831, 4675, 11417]", "true", "METLIFE_PERSONAL_GRP"}
	currentCarrierMichiganFarmBureau      = EnumCurrentCarrierItem{currentCarrierMichiganFarmBureauID, "Michigan Farm Bureau Group", map[string]string{"AmBest": "[2341, 2342]", "DisplayonUI": "false", "Parent": "MI_FARM_BUREAU_GROUP"}, "MichiganFarmBureau", 84, "[2341, 2342]", "false", "MI_FARM_BUREAU_GROUP"}
	currentCarrierMichiganMillers         = EnumCurrentCarrierItem{currentCarrierMichiganMillersID, "Michigan Millers Mutual", map[string]string{"AmBest": "[600]", "DisplayonUI": "false", "Parent": "MICH_MILLERS_MUT"}, "MichiganMillers", 85, "[600]", "false", "MICH_MILLERS_MUT"}
	currentCarrierMidwestFamilyMutual     = EnumCurrentCarrierItem{currentCarrierMidwestFamilyMutualID, "MidWest Family Mutual", map[string]string{"AmBest": "[2327]", "DisplayonUI": "false", "Parent": "MIDWEST_FAMILY_MUT"}, "MidwestFamilyMutual", 86, "[2327]", "false", "MIDWEST_FAMILY_MUT"}
	currentCarrierMMG                     = EnumCurrentCarrierItem{currentCarrierMMGID, "MMG Insurance Company", map[string]string{"AmBest": "[4692]", "DisplayonUI": "false", "Parent": "MMG_INSURANCE_CO"}, "MMG", 87, "[4692]", "false", "MMG_INSURANCE_CO"}
	currentCarrierMonument                = EnumCurrentCarrierItem{currentCarrierMonumentID, "Monument", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "Monument", 88, "[]", "false", ""}
	currentCarrierMSADUS                  = EnumCurrentCarrierItem{currentCarrierMSADUSID, "MS&AD US Insurance Group", map[string]string{"AmBest": "[4377, 883]", "DisplayonUI": "false", "Parent": "MSAD_US_INS_GROUP"}, "MSADUS", 89, "[4377, 883]", "false", "MSAD_US_INS_GROUP"}
	currentCarrierMunichAmerican          = EnumCurrentCarrierItem{currentCarrierMunichAmericanID, "Munich American Companies", map[string]string{"AmBest": "[465, 11074]", "DisplayonUI": "false", "Parent": "MUNICHAMERICAN_COS"}, "MunichAmerican", 90, "[465, 11074]", "false", "MUNICHAMERICAN_COS"}
	currentCarrierNationalGeneralHoldings = EnumCurrentCarrierItem{currentCarrierNationalGeneralHoldingsID, "National General Holdings Corp", map[string]string{"AmBest": "[1822]", "DisplayonUI": "false", "Parent": "NATL_GEN_HOLD_CORP"}, "NationalGeneralHoldings", 91, "[1822]", "false", "NATL_GEN_HOLD_CORP"}
	currentCarrierNationwide              = EnumCurrentCarrierItem{currentCarrierNationwideID, "Nationwide", map[string]string{"AmBest": "[1772, 80029, 80037, 80045, 2014, 80028, 80036, 80044, 10346, 91261, 91641, 1987, 1872, 80030, 80038, 80046, 80169, 80170, 366, 80012, 11802, 91271, 277, 91426, 2356, 2513, 80027, 80035, 80043, 90445, 90450, 12238, 2357, 80024, 80032, 80040, 2358, 80007, 80023, 80031, 80039, 80171, 80172, 80173, 80174, 2594, 80025, 80033, 80041, 91431, 12121, 548, 91621, 10709, 11688, 671, 91636, 12059, 11689, 12058, 3539]", "DisplayonUI": "true", "Parent": "NATIONWIDE_GROUP"}, "Nationwide", 92, "[1772, 80029, 80037, 80045, 2014, 80028, 80036, 80044, 10346, 91261, 91641, 1987, 1872, 80030, 80038, 80046, 80169, 80170, 366, 80012, 11802, 91271, 277, 91426, 2356, 2513, 80027, 80035, 80043, 90445, 90450, 12238, 2357, 80024, 80032, 80040, 2358, 80007, 80023, 80031, 80039, 80171, 80172, 80173, 80174, 2594, 80025, 80033, 80041, 91431, 12121, 548, 91621, 10709, 11688, 671, 91636, 12059, 11689, 12058, 3539]", "true", "NATIONWIDE_GROUP"}
	currentCarrierNLC                     = EnumCurrentCarrierItem{currentCarrierNLCID, "NLC Insurance Pool", map[string]string{"AmBest": "[11046, 478, 696, 3774]", "DisplayonUI": "false", "Parent": "NLC_INSURANCE_POOL"}, "NLC", 93, "[11046, 478, 696, 3774]", "false", "NLC_INSURANCE_POOL"}
	currentCarrierNodakMutual             = EnumCurrentCarrierItem{currentCarrierNodakMutualID, "Nodak Mutal Group", map[string]string{"AmBest": "[12426, 4187, 705]", "DisplayonUI": "false", "Parent": "NODAK_MUTUAL_GROUP"}, "NodakMutual", 94, "[12426, 4187, 705]", "false", "NODAK_MUTUAL_GROUP"}
	currentCarrierNonProfits              = EnumCurrentCarrierItem{currentCarrierNonProfitsID, "Non Profits Insurance Alliance Group", map[string]string{"AmBest": "[12419]", "DisplayonUI": "false", "Parent": "NONPROF_INS_ALL_GRP"}, "NonProfits", 95, "[12419]", "false", "NONPROF_INS_ALL_GRP"}
	currentCarrierNorfolkAndDedham        = EnumCurrentCarrierItem{currentCarrierNorfolkAndDedhamID, "Norfolk & Dedham PL", map[string]string{"AmBest": "[319, 399, 2367]", "DisplayonUI": "false", "Parent": "NORFOLK_DEDHAM_PL"}, "NorfolkAndDedham", 96, "[319, 399, 2367]", "false", "NORFOLK_DEDHAM_PL"}
	currentCarrierNYCM                    = EnumCurrentCarrierItem{currentCarrierNYCMID, "NYCM Insurance Group", map[string]string{"AmBest": "[12475, 700]", "DisplayonUI": "false", "Parent": "NYCM_INSURANCE_GROUP"}, "NYCM", 97, "[12475, 700]", "false", "NYCM_INSURANCE_GROUP"}
	currentCarrierOkalahomaFarmBureau     = EnumCurrentCarrierItem{currentCarrierOkalahomaFarmBureauID, "Oklahoma Farm Bureau Group", map[string]string{"AmBest": "[4099, 732]", "DisplayonUI": "false", "Parent": "OKLAHOMA_F_B_GROUP"}, "OkalahomaFarmBureau", 98, "[4099, 732]", "false", "OKLAHOMA_F_B_GROUP"}
	currentCarrierOldAmericaCapital       = EnumCurrentCarrierItem{currentCarrierOldAmericaCapitalID, "Old American Capital Group", map[string]string{"AmBest": "[10360]", "DisplayonUI": "false", "Parent": "OLD_AMER_CAPT_GRP"}, "OldAmericaCapital", 99, "[10360]", "false", "OLD_AMER_CAPT_GRP"}
	currentCarrierOregonMutual            = EnumCurrentCarrierItem{currentCarrierOregonMutualID, "Oregon Mutual Group", map[string]string{"AmBest": "[738, 105]", "DisplayonUI": "false", "Parent": "OREGON_MUT_GROUP"}, "OregonMutual", 100, "[738, 105]", "false", "OREGON_MUT_GROUP"}
	currentCarrierOther                   = EnumCurrentCarrierItem{currentCarrierOtherID, "Other", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "Other", 101, "[]", "false", ""}
	currentCarrierPacificSpecialty        = EnumCurrentCarrierItem{currentCarrierPacificSpecialtyID, "Pacific Specialty Group", map[string]string{"AmBest": "[11148]", "DisplayonUI": "false", "Parent": "PACIFIC_SPECIALTY_GR"}, "PacificSpecialty", 102, "[11148]", "false", "PACIFIC_SPECIALTY_GR"}
	currentCarrierPekin                   = EnumCurrentCarrierItem{currentCarrierPekinID, "Pekin Insurance Group", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "Pekin", 103, "[]", "false", ""}
	currentCarrierPennLumbermens          = EnumCurrentCarrierItem{currentCarrierPennLumbermensID, "Penn Lumbermens Mutual", map[string]string{"AmBest": "[756]", "DisplayonUI": "false", "Parent": "PENN_LUMBERMENS_MUT"}, "PennLumbermens", 104, "[756]", "false", "PENN_LUMBERMENS_MUT"}
	currentCarrierPharmacists             = EnumCurrentCarrierItem{currentCarrierPharmacistsID, "Pharmacists Insurance Group", map[string]string{"AmBest": "[320]", "DisplayonUI": "false", "Parent": "PHARMACISTS_INS_GRP"}, "Pharmacists", 105, "[320]", "false", "PHARMACISTS_INS_GRP"}
	currentCarrierPhiladlphiaTokio        = EnumCurrentCarrierItem{currentCarrierPhiladlphiaTokioID, "Philadlphia/Tokio Group", map[string]string{"AmBest": "[3616, 18564, 763]", "DisplayonUI": "false", "Parent": "PHILADLPHIA_TOKIO_GP"}, "PhiladlphiaTokio", 106, "[3616, 18564, 763]", "false", "PHILADLPHIA_TOKIO_GP"}
	currentCarrierPioneerStateMutual      = EnumCurrentCarrierItem{currentCarrierPioneerStateMutualID, "Pioneer State Mutual", map[string]string{"AmBest": "[860]", "DisplayonUI": "false", "Parent": "PIONEER_STATE_MUT"}, "PioneerStateMutual", 107, "[860]", "false", "PIONEER_STATE_MUT"}
	currentCarrierPreferredMutual         = EnumCurrentCarrierItem{currentCarrierPreferredMutualID, "Preferred Mutual Group", map[string]string{"AmBest": "[774]", "DisplayonUI": "false", "Parent": "PREFERRED_MUTUAL_GRP"}, "PreferredMutual", 108, "[774]", "false", "PREFERRED_MUTUAL_GRP"}
	currentCarrierProgressive             = EnumCurrentCarrierItem{currentCarrierProgressiveID, "Progressive", map[string]string{"AmBest": "[90510, 90515, 90520, 90525, 11759, 1891, 11150, 547, 586, 4456, 638, 2407, 11088, 1864, 3690, 649, 11760, 11441, 1839, 1762, 11665, 10745, 2408, 11761, 10724, 2650, 1853, 11246, 2609, 11247, 11698, 11758, 2640, 3645, 11248, 4287, 1900]", "DisplayonUI": "true", "Parent": "PROGRESSIVE_GROUP_OF_COMPANIES"}, "Progressive", 109, "[90510, 90515, 90520, 90525, 11759, 1891, 11150, 547, 586, 4456, 638, 2407, 11088, 1864, 3690, 649, 11760, 11441, 1839, 1762, 11665, 10745, 2408, 11761, 10724, 2650, 1853, 11246, 2609, 11247, 11698, 11758, 2640, 3645, 11248, 4287, 1900]", "true", "PROGRESSIVE_GROUP_OF_COMPANIES"}
	currentCarrierProsightSpecialty       = EnumCurrentCarrierItem{currentCarrierProsightSpecialtyID, "Prosight Specialty Group", map[string]string{"AmBest": "[728, 4676, 13309]", "DisplayonUI": "false", "Parent": "PROSIGHT_SPEC_GROUP"}, "ProsightSpecialty", 110, "[728, 4676, 13309]", "false", "PROSIGHT_SPEC_GROUP"}
	currentCarrierPublicServiceMutual     = EnumCurrentCarrierItem{currentCarrierPublicServiceMutualID, "Public Service Mutual Holding Company", map[string]string{"AmBest": "[2816, 792, 12272]", "DisplayonUI": "false", "Parent": "PUB_SERV_MUT_HOLD"}, "PublicServiceMutual", 111, "[2816, 792, 12272]", "false", "PUB_SERV_MUT_HOLD"}
	currentCarrierPure                    = EnumCurrentCarrierItem{currentCarrierPureID, "Pure Group of Insurance Companies", map[string]string{"AmBest": "[13816]", "DisplayonUI": "false", "Parent": "PURE_GRP_OF_INS_COS"}, "Pure", 112, "[13816]", "false", "PURE_GRP_OF_INS_COS"}
	currentCarrierQBEAmericas             = EnumCurrentCarrierItem{currentCarrierQBEAmericasID, "Qbe Americas Group", map[string]string{"AmBest": "[2416, 4242, 543, 676, 2643, 2739, 2418, 2402, 770, 2370]", "DisplayonUI": "false", "Parent": "QBE_AMERICAS_GROUP"}, "QBEAmericas", 113, "[2416, 4242, 543, 676, 2643, 2739, 2418, 2402, 770, 2370]", "false", "QBE_AMERICAS_GROUP"}
	currentCarrierRAMMutual               = EnumCurrentCarrierItem{currentCarrierRAMMutualID, "RAM Mutal Insurance Company", map[string]string{"AmBest": "[4814]", "DisplayonUI": "false", "Parent": "RAM_MUT_INS_CO"}, "RAMMutual", 114, "[4814]", "false", "RAM_MUT_INS_CO"}
	currentCarrierRepublicGroup           = EnumCurrentCarrierItem{currentCarrierRepublicGroupID, "Republic Group", map[string]string{"AmBest": "[3382]", "DisplayonUI": "false", "Parent": "REPUBLIC_GROUP"}, "RepublicGroup", 115, "[3382]", "false", "REPUBLIC_GROUP"}
	currentCarrierRLI                     = EnumCurrentCarrierItem{currentCarrierRLIID, "RLI group", map[string]string{"AmBest": "[4210, 2719]", "DisplayonUI": "false", "Parent": "RLI_GROUP"}, "RLI", 116, "[4210, 2719]", "false", "RLI_GROUP"}
	currentCarrierSafeAuto                = EnumCurrentCarrierItem{currentCarrierSafeAutoID, "SafeAuto", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "SafeAuto", 117, "[]", "false", ""}
	currentCarrierSamsungFireAndMarine    = EnumCurrentCarrierItem{currentCarrierSamsungFireAndMarineID, "Samsung Fire & Mar USB", map[string]string{"AmBest": "[11021]", "DisplayonUI": "false", "Parent": "SAMSUNG_FIREMAR_USB"}, "SamsungFireAndMarine", 118, "[11021]", "false", "SAMSUNG_FIREMAR_USB"}
	currentCarrierSecura                  = EnumCurrentCarrierItem{currentCarrierSecuraID, "Secura Insurance Companies", map[string]string{"AmBest": "[483, 11792]", "DisplayonUI": "false", "Parent": "SECURA_INS_COMPANIES"}, "Secura", 119, "[483, 11792]", "false", "SECURA_INS_COMPANIES"}
	currentCarrierSentry                  = EnumCurrentCarrierItem{currentCarrierSentryID, "Sentry Insurance", map[string]string{"AmBest": "[3277, 309, 2321, 4039, 11414, 3588, 2466, 2224, 4410]", "DisplayonUI": "true", "Parent": "SENTRY_INSURANCE_A_MUTUAL_COMPANY"}, "Sentry", 120, "[3277, 309, 2321, 4039, 11414, 3588, 2466, 2224, 4410]", "true", "SENTRY_INSURANCE_A_MUTUAL_COMPANY"}
	currentCarrierSompoJapan              = EnumCurrentCarrierItem{currentCarrierSompoJapanID, "Sompo Japan US Group", map[string]string{"AmBest": "[2693, 3060]", "DisplayonUI": "false", "Parent": "SOMPO_JAPAN_US_GROUP"}, "SompoJapan", 121, "[2693, 3060]", "false", "SOMPO_JAPAN_US_GROUP"}
	currentCarrierSouthportLane           = EnumCurrentCarrierItem{currentCarrierSouthportLaneID, "Southport Lane Group", map[string]string{"AmBest": "[1790]", "DisplayonUI": "false", "Parent": "SOUTHPORT_LANE_GROUP"}, "SouthportLane", 122, "[1790]", "false", "SOUTHPORT_LANE_GROUP"}
	currentCarrierStarAndShield           = EnumCurrentCarrierItem{currentCarrierStarAndShieldID, "Star & Shield Insurance Ex", map[string]string{"AmBest": "[14034]", "DisplayonUI": "false", "Parent": "STAR_SHIELD_INS_EX"}, "StarAndShield", 123, "[14034]", "false", "STAR_SHIELD_INS_EX"}
	currentCarrierStarCasualty            = EnumCurrentCarrierItem{currentCarrierStarCasualtyID, "Star Casualty Insurance Company", map[string]string{"AmBest": "[11198]", "DisplayonUI": "false", "Parent": "STAR_CASUALTY_INS_CO"}, "StarCasualty", 124, "[11198]", "false", "STAR_CASUALTY_INS_CO"}
	currentCarrierStarrInternational      = EnumCurrentCarrierItem{currentCarrierStarrInternationalID, "Starr International Company Inc", map[string]string{"AmBest": "[13853]", "DisplayonUI": "false", "Parent": "STARR_INT_COMP_INC"}, "StarrInternational", 125, "[13853]", "false", "STARR_INT_COMP_INC"}
	currentCarrierStateAuto               = EnumCurrentCarrierItem{currentCarrierStateAutoID, "State Auto", map[string]string{"AmBest": "[268, 3333, 628, 753, 12383, 3591, 2475, 855]", "DisplayonUI": "false", "Parent": "STATE_AUTO_INS_COS"}, "StateAuto", 126, "[268, 3333, 628, 753, 12383, 3591, 2475, 855]", "false", "STATE_AUTO_INS_COS"}
	currentCarrierStateFarm               = EnumCurrentCarrierItem{currentCarrierStateFarmID, "State Farm", map[string]string{"AmBest": "[88433, 88533, 88012, 88022, 88032, 88042, 88052, 88062, 88072, 88082, 88092, 88112, 88122, 88132, 88142, 88152, 88162, 88172, 88182, 88192, 88202, 88222, 88232, 88242, 88252, 88262, 88272, 88282, 88292, 88312, 88322, 88342, 88352, 88362, 88372, 88382, 88402, 88412, 88422, 88442, 88452, 88462, 88472, 88482, 88492, 88502, 88512, 88522, 88552, 88592, 88752, 88306, 88305, 88011, 88021, 88031, 88041, 88051, 88061, 88071, 88081, 88091, 88111, 88121, 88131, 88141, 88151, 88161, 88171, 88181, 88191, 88201, 88211, 88221, 88231, 88241, 88251, 88261, 88271, 88281, 88291, 88311, 88321, 88331, 88341, 88351, 88361, 88371, 88381, 88391, 88401, 88411, 88421, 88431, 88441, 88451, 88461, 88471, 88481, 88491, 88501, 88511, 88521, 88531, 88551, 88591, 88751, 2476, 2477, 13016, 11224, 2479]", "DisplayonUI": "true", "Parent": "STATE_FARM_GROUP"}, "StateFarm", 127, "[88433, 88533, 88012, 88022, 88032, 88042, 88052, 88062, 88072, 88082, 88092, 88112, 88122, 88132, 88142, 88152, 88162, 88172, 88182, 88192, 88202, 88222, 88232, 88242, 88252, 88262, 88272, 88282, 88292, 88312, 88322, 88342, 88352, 88362, 88372, 88382, 88402, 88412, 88422, 88442, 88452, 88462, 88472, 88482, 88492, 88502, 88512, 88522, 88552, 88592, 88752, 88306, 88305, 88011, 88021, 88031, 88041, 88051, 88061, 88071, 88081, 88091, 88111, 88121, 88131, 88141, 88151, 88161, 88171, 88181, 88191, 88201, 88211, 88221, 88231, 88241, 88251, 88261, 88271, 88281, 88291, 88311, 88321, 88331, 88341, 88351, 88361, 88371, 88381, 88391, 88401, 88411, 88421, 88431, 88441, 88451, 88461, 88471, 88481, 88491, 88501, 88511, 88521, 88531, 88551, 88591, 88751, 2476, 2477, 13016, 11224, 2479]", "true", "STATE_FARM_GROUP"}
	currentCarrierStateNational           = EnumCurrentCarrierItem{currentCarrierStateNationalID, "State National Group", map[string]string{"AmBest": "[10681]", "DisplayonUI": "false", "Parent": "STATE_NATIONAL_GROUP"}, "StateNational", 128, "[10681]", "false", "STATE_NATIONAL_GROUP"}
	currentCarrierSterling                = EnumCurrentCarrierItem{currentCarrierSterlingID, "Sterling Insurance Company", map[string]string{"AmBest": "[872]", "DisplayonUI": "false", "Parent": "STERLING_INS_CO"}, "Sterling", 129, "[872]", "false", "STERLING_INS_CO"}
	currentCarrierTheGeneral              = EnumCurrentCarrierItem{currentCarrierTheGeneralID, "The General", map[string]string{"AmBest": "[11303, 1894, 14118]", "DisplayonUI": "true", "Parent": "AMERICAN_FAMILY_INSURANCE_GROUP"}, "TheGeneral", 130, "[11303, 1894, 14118]", "true", "AMERICAN_FAMILY_INSURANCE_GROUP"}
	currentCarrierTheHartford             = EnumCurrentCarrierItem{currentCarrierTheHartfordID, "The Hartford", map[string]string{"AmBest": "[1907, 2187, 2229, 2230, 2231, 2232, 2234, 2235, 2610, 2611, 2612, 2613, 2614, 2706, 10777, 11654]", "DisplayonUI": "true", "Parent": "HARTFORD_INSURANCE_GROUP"}, "TheHartford", 131, "[1907, 2187, 2229, 2230, 2231, 2232, 2234, 2235, 2610, 2611, 2612, 2613, 2614, 2706, 10777, 11654]", "true", "HARTFORD_INSURANCE_GROUP"}
	currentCarrierTheRepublicGroup        = EnumCurrentCarrierItem{currentCarrierTheRepublicGroupID, "The Republic Group", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "TheRepublicGroup", 132, "[]", "false", ""}
	currentCarrierTitus                   = EnumCurrentCarrierItem{currentCarrierTitusID, "Titus Group", map[string]string{"AmBest": "[4682]", "DisplayonUI": "false", "Parent": "TITUS_GROUP"}, "Titus", 133, "[4682]", "false", "TITUS_GROUP"}
	currentCarrierTopa                    = EnumCurrentCarrierItem{currentCarrierTopaID, "Topa Insurance Group", map[string]string{"AmBest": "[2749]", "DisplayonUI": "false", "Parent": "TOPA_INSURANCE_GROUP"}, "Topa", 134, "[2749]", "false", "TOPA_INSURANCE_GROUP"}
	currentCarrierTower                   = EnumCurrentCarrierItem{currentCarrierTowerID, "Tower Group Companies", map[string]string{"AmBest": "[13114, 14012, 13055, 10109, 2821, 12541, 11352, 11356, 11487, 269]", "DisplayonUI": "false", "Parent": "TOWER_GRP_COMPANIES"}, "Tower", 135, "[13114, 14012, 13055, 10109, 2821, 12541, 11352, 11356, 11487, 269]", "false", "TOWER_GRP_COMPANIES"}
	currentCarrierTradersGeneral          = EnumCurrentCarrierItem{currentCarrierTradersGeneralID, "Traders General Agency Inc", map[string]string{"AmBest": "[1875]", "DisplayonUI": "false", "Parent": "TRADERS_GEN_AGENCY"}, "TradersGeneral", 136, "[1875]", "false", "TRADERS_GEN_AGENCY"}
	currentCarrierTravelers               = EnumCurrentCarrierItem{currentCarrierTravelersID, "Travelers", map[string]string{"AmBest": "[712, 1744, 2002, 2016, 2452, 2453, 2516, 2517, 2518, 2520, 2538, 3792, 4003, 4025, 4046, 4311, 4461, 4869, 11020, 11024, 11025, 11026, 11027, 11300, 11302, 11421, 11872, 12061, 80067, 4465]", "DisplayonUI": "true", "Parent": "TRAVELERS_GROUP"}, "Travelers", 137, "[712, 1744, 2002, 2016, 2452, 2453, 2516, 2517, 2518, 2520, 2538, 3792, 4003, 4025, 4046, 4311, 4461, 4869, 11020, 11024, 11025, 11026, 11027, 11300, 11302, 11421, 11872, 12061, 80067, 4465]", "true", "TRAVELERS_GROUP"}
	currentCarrierUnitedHome              = EnumCurrentCarrierItem{currentCarrierUnitedHomeID, "United Home Insurance Company", map[string]string{"AmBest": "[11444]", "DisplayonUI": "false", "Parent": "UNITED_HOME_INS_COMP"}, "UnitedHome", 138, "[11444]", "false", "UNITED_HOME_INS_COMP"}
	currentCarrierUnitrin                 = EnumCurrentCarrierItem{currentCarrierUnitrinID, "Unitrin", map[string]string{"AmBest": "[]", "DisplayonUI": "false", "Parent": ""}, "Unitrin", 139, "[]", "false", ""}
	currentCarrierUSAA                    = EnumCurrentCarrierItem{currentCarrierUSAAID, "USAA", map[string]string{"AmBest": "[934, 4049, 4865, 11699, 12120]", "DisplayonUI": "true", "Parent": "USAA_GROUP"}, "USAA", 140, "[934, 4049, 4865, 11699, 12120]", "true", "USAA_GROUP"}
	currentCarrierUticaMutual             = EnumCurrentCarrierItem{currentCarrierUticaMutualID, "Utica Mutual Insurance Company", map[string]string{"AmBest": "[946, 428, 798,2825, 11953, 4332]", "DisplayonUI": "false", "Parent": "UTICA_MUT_INS_COMP"}, "UticaMutual", 141, "[946, 428, 798,2825, 11953, 4332]", "false", "UTICA_MUT_INS_COMP"}
	currentCarrierVAFarmBureau            = EnumCurrentCarrierItem{currentCarrierVAFarmBureauID, "Virginia Farm Bureau", map[string]string{"AmBest": "[2548, 2549, 11667]", "DisplayonUI": "true", "Parent": "VIRGINIA_F_B_GROUP"}, "VAFarmBureau", 142, "[2548, 2549, 11667]", "true", "VIRGINIA_F_B_GROUP"}
	currentCarrierWawanesa                = EnumCurrentCarrierItem{currentCarrierWawanesaID, "Wawanesa General Insurance", map[string]string{"AmBest": "[11976]", "DisplayonUI": "false", "Parent": "WAWANESA_GENERAL_INS"}, "Wawanesa", 143, "[11976]", "false", "WAWANESA_GENERAL_INS"}
	currentCarrierWestBendMutual          = EnumCurrentCarrierItem{currentCarrierWestBendMutualID, "West Bend Mutual Insurance Company", map[string]string{"AmBest": "[964]", "DisplayonUI": "false", "Parent": "WEST_BEND_MUT_INS_CO"}, "WestBendMutual", 144, "[964]", "false", "WEST_BEND_MUT_INS_CO"}
	currentCarrierWesternGeneral          = EnumCurrentCarrierItem{currentCarrierWesternGeneralID, "Western General Insurance", map[string]string{"AmBest": "[3560]", "DisplayonUI": "false", "Parent": "WESTERN_GENERAL_INS"}, "WesternGeneral", 145, "[3560]", "false", "WESTERN_GENERAL_INS"}
	currentCarrierWesternReservePool      = EnumCurrentCarrierItem{currentCarrierWesternReservePoolID, "Western Reserve Pool", map[string]string{"AmBest": "[2284, 2285]", "DisplayonUI": "false", "Parent": "WESTERN_RESERVE_POOL"}, "WesternReservePool", 146, "[2284, 2285]", "false", "WESTERN_RESERVE_POOL"}
	currentCarrierWestfield               = EnumCurrentCarrierItem{currentCarrierWestfieldID, "WestField Group", map[string]string{"AmBest": "[2069, 2381, 2382, 4043]", "DisplayonUI": "false", "Parent": "WESTFIELD_GROUP"}, "Westfield", 147, "[2069, 2381, 2382, 4043]", "false", "WESTFIELD_GROUP"}
	currentCarrierWhiteMountains          = EnumCurrentCarrierItem{currentCarrierWhiteMountainsID, "White Mountains Insurance group", map[string]string{"AmBest": "[12666, 10157, 2106, 2196]", "DisplayonUI": "false", "Parent": "WHITE_MTNS_INS_GRP"}, "WhiteMountains", 148, "[12666, 10157, 2106, 2196]", "false", "WHITE_MTNS_INS_GRP"}
	currentCarrierWisconsinMutual         = EnumCurrentCarrierItem{currentCarrierWisconsinMutualID, "Wisconsin Mutual Insurance Company", map[string]string{"AmBest": "[3524]", "DisplayonUI": "false", "Parent": "WISCONSIN_MUT_INS_CO"}, "WisconsinMutual", 149, "[3524]", "false", "WISCONSIN_MUT_INS_CO"}
	currentCarrierWolverineMutual         = EnumCurrentCarrierItem{currentCarrierWolverineMutualID, "Wolverine Mutual Insurance Company", map[string]string{"AmBest": "[3237]", "DisplayonUI": "false", "Parent": "WOLVERINE_MUT_INS_CO"}, "WolverineMutual", 150, "[3237]", "false", "WOLVERINE_MUT_INS_CO"}
	currentCarrierWorkmensAuto            = EnumCurrentCarrierItem{currentCarrierWorkmensAutoID, "Workmen's Auto Insurance", map[string]string{"AmBest": "[980]", "DisplayonUI": "false", "Parent": "WORKMENS_AUTO_INS"}, "WorkmensAuto", 151, "[980]", "false", "WORKMENS_AUTO_INS"}
	currentCarrierWRMAmerica              = EnumCurrentCarrierItem{currentCarrierWRMAmericaID, "Wrm America Group", map[string]string{"AmBest": "[12582]", "DisplayonUI": "false", "Parent": "WRM_AMERICA_GROUP"}, "WRMAmerica", 152, "[12582]", "false", "WRM_AMERICA_GROUP"}
	currentCarrierZurichInsurance         = EnumCurrentCarrierItem{currentCarrierZurichInsuranceID, "Zurich Insurance Group Ltd", map[string]string{"AmBest": "[1754]", "DisplayonUI": "false", "Parent": "ZURICH_INS_GRP_LTD"}, "ZurichInsurance", 153, "[1754]", "false", "ZURICH_INS_GRP_LTD"}
)

// EnumCurrentCarrier is a collection of CurrentCarrier items
type EnumCurrentCarrier struct {
	Description string
	Items       []*EnumCurrentCarrierItem
	Name        string

	AAA                     *EnumCurrentCarrierItem
	Access                  *EnumCurrentCarrierItem
	ACEGroup                *EnumCurrentCarrierItem
	AffinityMutual          *EnumCurrentCarrierItem
	Affirmative             *EnumCurrentCarrierItem
	AGIC                    *EnumCurrentCarrierItem
	AIG                     *EnumCurrentCarrierItem
	AlaskaNational          *EnumCurrentCarrierItem
	Alfa                    *EnumCurrentCarrierItem
	Allegheny               *EnumCurrentCarrierItem
	Allianz                 *EnumCurrentCarrierItem
	AlliedWorld             *EnumCurrentCarrierItem
	Allstate                *EnumCurrentCarrierItem
	AmericanEuropean        *EnumCurrentCarrierItem
	AmericanFamily          *EnumCurrentCarrierItem
	AmericanFarmAndRanch    *EnumCurrentCarrierItem
	AmericanIndependent     *EnumCurrentCarrierItem
	AmericanInternational   *EnumCurrentCarrierItem
	AmericanNational        *EnumCurrentCarrierItem
	AmericanSterling        *EnumCurrentCarrierItem
	Ameriprise              *EnumCurrentCarrierItem
	AutoClubEnterprises     *EnumCurrentCarrierItem
	AutoClubOfFlorida       *EnumCurrentCarrierItem
	AutoClubSouth           *EnumCurrentCarrierItem
	AutoOwners              *EnumCurrentCarrierItem
	Axis                    *EnumCurrentCarrierItem
	BaldwinAndLyons         *EnumCurrentCarrierItem
	BrethrenMutual          *EnumCurrentCarrierItem
	CaliforniaCasualty      *EnumCurrentCarrierItem
	CenturyNational         *EnumCurrentCarrierItem
	Chubb                   *EnumCurrentCarrierItem
	ChurchMutual            *EnumCurrentCarrierItem
	CivilServiceEmployees   *EnumCurrentCarrierItem
	CNA                     *EnumCurrentCarrierItem
	CNAFinancialGroup       *EnumCurrentCarrierItem
	Companion               *EnumCurrentCarrierItem
	Consumers               *EnumCurrentCarrierItem
	CountryFinancial        *EnumCurrentCarrierItem
	CountryFinancialPC      *EnumCurrentCarrierItem
	DirectAuto              *EnumCurrentCarrierItem
	DirectGeneral           *EnumCurrentCarrierItem
	DonegalGroup            *EnumCurrentCarrierItem
	EIns                    *EnumCurrentCarrierItem
	Electric                *EnumCurrentCarrierItem
	Elephant                *EnumCurrentCarrierItem
	EMC                     *EnumCurrentCarrierItem
	Erie                    *EnumCurrentCarrierItem
	EverestRE               *EnumCurrentCarrierItem
	EvergreenUSA            *EnumCurrentCarrierItem
	FairfaxFinancial        *EnumCurrentCarrierItem
	FarmBureau              *EnumCurrentCarrierItem
	Farmers                 *EnumCurrentCarrierItem
	FarmersMutual           *EnumCurrentCarrierItem
	FidelityNational        *EnumCurrentCarrierItem
	FirstAcceptance         *EnumCurrentCarrierItem
	Frankenmuth             *EnumCurrentCarrierItem
	FredLoya                *EnumCurrentCarrierItem
	Geico                   *EnumCurrentCarrierItem
	GeorgiaFarmBureau       *EnumCurrentCarrierItem
	GoodvilleAndGerman      *EnumCurrentCarrierItem
	Grange                  *EnumCurrentCarrierItem
	GrangeMutual            *EnumCurrentCarrierItem
	GreatAmerican           *EnumCurrentCarrierItem
	GuideOne                *EnumCurrentCarrierItem
	Hallmark                *EnumCurrentCarrierItem
	HastingsMutual          *EnumCurrentCarrierItem
	HiscoxUSA               *EnumCurrentCarrierItem
	HochheimPrairie         *EnumCurrentCarrierItem
	HomeState               *EnumCurrentCarrierItem
	HoraceMann              *EnumCurrentCarrierItem
	IAT                     *EnumCurrentCarrierItem
	IMT                     *EnumCurrentCarrierItem
	Infinity                *EnumCurrentCarrierItem
	IntegonNational         *EnumCurrentCarrierItem
	Ironshore               *EnumCurrentCarrierItem
	Kemper                  *EnumCurrentCarrierItem
	Key                     *EnumCurrentCarrierItem
	KingswayFinancial       *EnumCurrentCarrierItem
	LibertyMutual           *EnumCurrentCarrierItem
	MAPFRE                  *EnumCurrentCarrierItem
	Merchants               *EnumCurrentCarrierItem
	MercuryGeneral          *EnumCurrentCarrierItem
	MetLife                 *EnumCurrentCarrierItem
	MichiganFarmBureau      *EnumCurrentCarrierItem
	MichiganMillers         *EnumCurrentCarrierItem
	MidwestFamilyMutual     *EnumCurrentCarrierItem
	MMG                     *EnumCurrentCarrierItem
	Monument                *EnumCurrentCarrierItem
	MSADUS                  *EnumCurrentCarrierItem
	MunichAmerican          *EnumCurrentCarrierItem
	NationalGeneralHoldings *EnumCurrentCarrierItem
	Nationwide              *EnumCurrentCarrierItem
	NLC                     *EnumCurrentCarrierItem
	NodakMutual             *EnumCurrentCarrierItem
	NonProfits              *EnumCurrentCarrierItem
	NorfolkAndDedham        *EnumCurrentCarrierItem
	NYCM                    *EnumCurrentCarrierItem
	OkalahomaFarmBureau     *EnumCurrentCarrierItem
	OldAmericaCapital       *EnumCurrentCarrierItem
	OregonMutual            *EnumCurrentCarrierItem
	Other                   *EnumCurrentCarrierItem
	PacificSpecialty        *EnumCurrentCarrierItem
	Pekin                   *EnumCurrentCarrierItem
	PennLumbermens          *EnumCurrentCarrierItem
	Pharmacists             *EnumCurrentCarrierItem
	PhiladlphiaTokio        *EnumCurrentCarrierItem
	PioneerStateMutual      *EnumCurrentCarrierItem
	PreferredMutual         *EnumCurrentCarrierItem
	Progressive             *EnumCurrentCarrierItem
	ProsightSpecialty       *EnumCurrentCarrierItem
	PublicServiceMutual     *EnumCurrentCarrierItem
	Pure                    *EnumCurrentCarrierItem
	QBEAmericas             *EnumCurrentCarrierItem
	RAMMutual               *EnumCurrentCarrierItem
	RepublicGroup           *EnumCurrentCarrierItem
	RLI                     *EnumCurrentCarrierItem
	SafeAuto                *EnumCurrentCarrierItem
	SamsungFireAndMarine    *EnumCurrentCarrierItem
	Secura                  *EnumCurrentCarrierItem
	Sentry                  *EnumCurrentCarrierItem
	SompoJapan              *EnumCurrentCarrierItem
	SouthportLane           *EnumCurrentCarrierItem
	StarAndShield           *EnumCurrentCarrierItem
	StarCasualty            *EnumCurrentCarrierItem
	StarrInternational      *EnumCurrentCarrierItem
	StateAuto               *EnumCurrentCarrierItem
	StateFarm               *EnumCurrentCarrierItem
	StateNational           *EnumCurrentCarrierItem
	Sterling                *EnumCurrentCarrierItem
	TheGeneral              *EnumCurrentCarrierItem
	TheHartford             *EnumCurrentCarrierItem
	TheRepublicGroup        *EnumCurrentCarrierItem
	Titus                   *EnumCurrentCarrierItem
	Topa                    *EnumCurrentCarrierItem
	Tower                   *EnumCurrentCarrierItem
	TradersGeneral          *EnumCurrentCarrierItem
	Travelers               *EnumCurrentCarrierItem
	UnitedHome              *EnumCurrentCarrierItem
	Unitrin                 *EnumCurrentCarrierItem
	USAA                    *EnumCurrentCarrierItem
	UticaMutual             *EnumCurrentCarrierItem
	VAFarmBureau            *EnumCurrentCarrierItem
	Wawanesa                *EnumCurrentCarrierItem
	WestBendMutual          *EnumCurrentCarrierItem
	WesternGeneral          *EnumCurrentCarrierItem
	WesternReservePool      *EnumCurrentCarrierItem
	Westfield               *EnumCurrentCarrierItem
	WhiteMountains          *EnumCurrentCarrierItem
	WisconsinMutual         *EnumCurrentCarrierItem
	WolverineMutual         *EnumCurrentCarrierItem
	WorkmensAuto            *EnumCurrentCarrierItem
	WRMAmerica              *EnumCurrentCarrierItem
	ZurichInsurance         *EnumCurrentCarrierItem

	itemDict map[string]*EnumCurrentCarrierItem
}

// CurrentCarrier is a public singleton instance of EnumCurrentCarrier
// representing auto insurance companies
var CurrentCarrier = &EnumCurrentCarrier{
	Description: "auto insurance companies",
	Items: []*EnumCurrentCarrierItem{
		&currentCarrierAAA,
		&currentCarrierAccess,
		&currentCarrierACEGroup,
		&currentCarrierAffinityMutual,
		&currentCarrierAffirmative,
		&currentCarrierAGIC,
		&currentCarrierAIG,
		&currentCarrierAlaskaNational,
		&currentCarrierAlfa,
		&currentCarrierAllegheny,
		&currentCarrierAllianz,
		&currentCarrierAlliedWorld,
		&currentCarrierAllstate,
		&currentCarrierAmericanEuropean,
		&currentCarrierAmericanFamily,
		&currentCarrierAmericanFarmAndRanch,
		&currentCarrierAmericanIndependent,
		&currentCarrierAmericanInternational,
		&currentCarrierAmericanNational,
		&currentCarrierAmericanSterling,
		&currentCarrierAmeriprise,
		&currentCarrierAutoClubEnterprises,
		&currentCarrierAutoClubOfFlorida,
		&currentCarrierAutoClubSouth,
		&currentCarrierAutoOwners,
		&currentCarrierAxis,
		&currentCarrierBaldwinAndLyons,
		&currentCarrierBrethrenMutual,
		&currentCarrierCaliforniaCasualty,
		&currentCarrierCenturyNational,
		&currentCarrierChubb,
		&currentCarrierChurchMutual,
		&currentCarrierCivilServiceEmployees,
		&currentCarrierCNA,
		&currentCarrierCNAFinancialGroup,
		&currentCarrierCompanion,
		&currentCarrierConsumers,
		&currentCarrierCountryFinancial,
		&currentCarrierCountryFinancialPC,
		&currentCarrierDirectAuto,
		&currentCarrierDirectGeneral,
		&currentCarrierDonegalGroup,
		&currentCarrierEIns,
		&currentCarrierElectric,
		&currentCarrierElephant,
		&currentCarrierEMC,
		&currentCarrierErie,
		&currentCarrierEverestRE,
		&currentCarrierEvergreenUSA,
		&currentCarrierFairfaxFinancial,
		&currentCarrierFarmBureau,
		&currentCarrierFarmers,
		&currentCarrierFarmersMutual,
		&currentCarrierFidelityNational,
		&currentCarrierFirstAcceptance,
		&currentCarrierFrankenmuth,
		&currentCarrierFredLoya,
		&currentCarrierGeico,
		&currentCarrierGeorgiaFarmBureau,
		&currentCarrierGoodvilleAndGerman,
		&currentCarrierGrange,
		&currentCarrierGrangeMutual,
		&currentCarrierGreatAmerican,
		&currentCarrierGuideOne,
		&currentCarrierHallmark,
		&currentCarrierHastingsMutual,
		&currentCarrierHiscoxUSA,
		&currentCarrierHochheimPrairie,
		&currentCarrierHomeState,
		&currentCarrierHoraceMann,
		&currentCarrierIAT,
		&currentCarrierIMT,
		&currentCarrierInfinity,
		&currentCarrierIntegonNational,
		&currentCarrierIronshore,
		&currentCarrierKemper,
		&currentCarrierKey,
		&currentCarrierKingswayFinancial,
		&currentCarrierLibertyMutual,
		&currentCarrierMAPFRE,
		&currentCarrierMerchants,
		&currentCarrierMercuryGeneral,
		&currentCarrierMetLife,
		&currentCarrierMichiganFarmBureau,
		&currentCarrierMichiganMillers,
		&currentCarrierMidwestFamilyMutual,
		&currentCarrierMMG,
		&currentCarrierMonument,
		&currentCarrierMSADUS,
		&currentCarrierMunichAmerican,
		&currentCarrierNationalGeneralHoldings,
		&currentCarrierNationwide,
		&currentCarrierNLC,
		&currentCarrierNodakMutual,
		&currentCarrierNonProfits,
		&currentCarrierNorfolkAndDedham,
		&currentCarrierNYCM,
		&currentCarrierOkalahomaFarmBureau,
		&currentCarrierOldAmericaCapital,
		&currentCarrierOregonMutual,
		&currentCarrierOther,
		&currentCarrierPacificSpecialty,
		&currentCarrierPekin,
		&currentCarrierPennLumbermens,
		&currentCarrierPharmacists,
		&currentCarrierPhiladlphiaTokio,
		&currentCarrierPioneerStateMutual,
		&currentCarrierPreferredMutual,
		&currentCarrierProgressive,
		&currentCarrierProsightSpecialty,
		&currentCarrierPublicServiceMutual,
		&currentCarrierPure,
		&currentCarrierQBEAmericas,
		&currentCarrierRAMMutual,
		&currentCarrierRepublicGroup,
		&currentCarrierRLI,
		&currentCarrierSafeAuto,
		&currentCarrierSamsungFireAndMarine,
		&currentCarrierSecura,
		&currentCarrierSentry,
		&currentCarrierSompoJapan,
		&currentCarrierSouthportLane,
		&currentCarrierStarAndShield,
		&currentCarrierStarCasualty,
		&currentCarrierStarrInternational,
		&currentCarrierStateAuto,
		&currentCarrierStateFarm,
		&currentCarrierStateNational,
		&currentCarrierSterling,
		&currentCarrierTheGeneral,
		&currentCarrierTheHartford,
		&currentCarrierTheRepublicGroup,
		&currentCarrierTitus,
		&currentCarrierTopa,
		&currentCarrierTower,
		&currentCarrierTradersGeneral,
		&currentCarrierTravelers,
		&currentCarrierUnitedHome,
		&currentCarrierUnitrin,
		&currentCarrierUSAA,
		&currentCarrierUticaMutual,
		&currentCarrierVAFarmBureau,
		&currentCarrierWawanesa,
		&currentCarrierWestBendMutual,
		&currentCarrierWesternGeneral,
		&currentCarrierWesternReservePool,
		&currentCarrierWestfield,
		&currentCarrierWhiteMountains,
		&currentCarrierWisconsinMutual,
		&currentCarrierWolverineMutual,
		&currentCarrierWorkmensAuto,
		&currentCarrierWRMAmerica,
		&currentCarrierZurichInsurance,
	},
	Name:                    "EnumCurrentCarrier",
	AAA:                     &currentCarrierAAA,
	Access:                  &currentCarrierAccess,
	ACEGroup:                &currentCarrierACEGroup,
	AffinityMutual:          &currentCarrierAffinityMutual,
	Affirmative:             &currentCarrierAffirmative,
	AGIC:                    &currentCarrierAGIC,
	AIG:                     &currentCarrierAIG,
	AlaskaNational:          &currentCarrierAlaskaNational,
	Alfa:                    &currentCarrierAlfa,
	Allegheny:               &currentCarrierAllegheny,
	Allianz:                 &currentCarrierAllianz,
	AlliedWorld:             &currentCarrierAlliedWorld,
	Allstate:                &currentCarrierAllstate,
	AmericanEuropean:        &currentCarrierAmericanEuropean,
	AmericanFamily:          &currentCarrierAmericanFamily,
	AmericanFarmAndRanch:    &currentCarrierAmericanFarmAndRanch,
	AmericanIndependent:     &currentCarrierAmericanIndependent,
	AmericanInternational:   &currentCarrierAmericanInternational,
	AmericanNational:        &currentCarrierAmericanNational,
	AmericanSterling:        &currentCarrierAmericanSterling,
	Ameriprise:              &currentCarrierAmeriprise,
	AutoClubEnterprises:     &currentCarrierAutoClubEnterprises,
	AutoClubOfFlorida:       &currentCarrierAutoClubOfFlorida,
	AutoClubSouth:           &currentCarrierAutoClubSouth,
	AutoOwners:              &currentCarrierAutoOwners,
	Axis:                    &currentCarrierAxis,
	BaldwinAndLyons:         &currentCarrierBaldwinAndLyons,
	BrethrenMutual:          &currentCarrierBrethrenMutual,
	CaliforniaCasualty:      &currentCarrierCaliforniaCasualty,
	CenturyNational:         &currentCarrierCenturyNational,
	Chubb:                   &currentCarrierChubb,
	ChurchMutual:            &currentCarrierChurchMutual,
	CivilServiceEmployees:   &currentCarrierCivilServiceEmployees,
	CNA:                     &currentCarrierCNA,
	CNAFinancialGroup:       &currentCarrierCNAFinancialGroup,
	Companion:               &currentCarrierCompanion,
	Consumers:               &currentCarrierConsumers,
	CountryFinancial:        &currentCarrierCountryFinancial,
	CountryFinancialPC:      &currentCarrierCountryFinancialPC,
	DirectAuto:              &currentCarrierDirectAuto,
	DirectGeneral:           &currentCarrierDirectGeneral,
	DonegalGroup:            &currentCarrierDonegalGroup,
	EIns:                    &currentCarrierEIns,
	Electric:                &currentCarrierElectric,
	Elephant:                &currentCarrierElephant,
	EMC:                     &currentCarrierEMC,
	Erie:                    &currentCarrierErie,
	EverestRE:               &currentCarrierEverestRE,
	EvergreenUSA:            &currentCarrierEvergreenUSA,
	FairfaxFinancial:        &currentCarrierFairfaxFinancial,
	FarmBureau:              &currentCarrierFarmBureau,
	Farmers:                 &currentCarrierFarmers,
	FarmersMutual:           &currentCarrierFarmersMutual,
	FidelityNational:        &currentCarrierFidelityNational,
	FirstAcceptance:         &currentCarrierFirstAcceptance,
	Frankenmuth:             &currentCarrierFrankenmuth,
	FredLoya:                &currentCarrierFredLoya,
	Geico:                   &currentCarrierGeico,
	GeorgiaFarmBureau:       &currentCarrierGeorgiaFarmBureau,
	GoodvilleAndGerman:      &currentCarrierGoodvilleAndGerman,
	Grange:                  &currentCarrierGrange,
	GrangeMutual:            &currentCarrierGrangeMutual,
	GreatAmerican:           &currentCarrierGreatAmerican,
	GuideOne:                &currentCarrierGuideOne,
	Hallmark:                &currentCarrierHallmark,
	HastingsMutual:          &currentCarrierHastingsMutual,
	HiscoxUSA:               &currentCarrierHiscoxUSA,
	HochheimPrairie:         &currentCarrierHochheimPrairie,
	HomeState:               &currentCarrierHomeState,
	HoraceMann:              &currentCarrierHoraceMann,
	IAT:                     &currentCarrierIAT,
	IMT:                     &currentCarrierIMT,
	Infinity:                &currentCarrierInfinity,
	IntegonNational:         &currentCarrierIntegonNational,
	Ironshore:               &currentCarrierIronshore,
	Kemper:                  &currentCarrierKemper,
	Key:                     &currentCarrierKey,
	KingswayFinancial:       &currentCarrierKingswayFinancial,
	LibertyMutual:           &currentCarrierLibertyMutual,
	MAPFRE:                  &currentCarrierMAPFRE,
	Merchants:               &currentCarrierMerchants,
	MercuryGeneral:          &currentCarrierMercuryGeneral,
	MetLife:                 &currentCarrierMetLife,
	MichiganFarmBureau:      &currentCarrierMichiganFarmBureau,
	MichiganMillers:         &currentCarrierMichiganMillers,
	MidwestFamilyMutual:     &currentCarrierMidwestFamilyMutual,
	MMG:                     &currentCarrierMMG,
	Monument:                &currentCarrierMonument,
	MSADUS:                  &currentCarrierMSADUS,
	MunichAmerican:          &currentCarrierMunichAmerican,
	NationalGeneralHoldings: &currentCarrierNationalGeneralHoldings,
	Nationwide:              &currentCarrierNationwide,
	NLC:                     &currentCarrierNLC,
	NodakMutual:             &currentCarrierNodakMutual,
	NonProfits:              &currentCarrierNonProfits,
	NorfolkAndDedham:        &currentCarrierNorfolkAndDedham,
	NYCM:                    &currentCarrierNYCM,
	OkalahomaFarmBureau:     &currentCarrierOkalahomaFarmBureau,
	OldAmericaCapital:       &currentCarrierOldAmericaCapital,
	OregonMutual:            &currentCarrierOregonMutual,
	Other:                   &currentCarrierOther,
	PacificSpecialty:        &currentCarrierPacificSpecialty,
	Pekin:                   &currentCarrierPekin,
	PennLumbermens:          &currentCarrierPennLumbermens,
	Pharmacists:             &currentCarrierPharmacists,
	PhiladlphiaTokio:        &currentCarrierPhiladlphiaTokio,
	PioneerStateMutual:      &currentCarrierPioneerStateMutual,
	PreferredMutual:         &currentCarrierPreferredMutual,
	Progressive:             &currentCarrierProgressive,
	ProsightSpecialty:       &currentCarrierProsightSpecialty,
	PublicServiceMutual:     &currentCarrierPublicServiceMutual,
	Pure:                    &currentCarrierPure,
	QBEAmericas:             &currentCarrierQBEAmericas,
	RAMMutual:               &currentCarrierRAMMutual,
	RepublicGroup:           &currentCarrierRepublicGroup,
	RLI:                     &currentCarrierRLI,
	SafeAuto:                &currentCarrierSafeAuto,
	SamsungFireAndMarine:    &currentCarrierSamsungFireAndMarine,
	Secura:                  &currentCarrierSecura,
	Sentry:                  &currentCarrierSentry,
	SompoJapan:              &currentCarrierSompoJapan,
	SouthportLane:           &currentCarrierSouthportLane,
	StarAndShield:           &currentCarrierStarAndShield,
	StarCasualty:            &currentCarrierStarCasualty,
	StarrInternational:      &currentCarrierStarrInternational,
	StateAuto:               &currentCarrierStateAuto,
	StateFarm:               &currentCarrierStateFarm,
	StateNational:           &currentCarrierStateNational,
	Sterling:                &currentCarrierSterling,
	TheGeneral:              &currentCarrierTheGeneral,
	TheHartford:             &currentCarrierTheHartford,
	TheRepublicGroup:        &currentCarrierTheRepublicGroup,
	Titus:                   &currentCarrierTitus,
	Topa:                    &currentCarrierTopa,
	Tower:                   &currentCarrierTower,
	TradersGeneral:          &currentCarrierTradersGeneral,
	Travelers:               &currentCarrierTravelers,
	UnitedHome:              &currentCarrierUnitedHome,
	Unitrin:                 &currentCarrierUnitrin,
	USAA:                    &currentCarrierUSAA,
	UticaMutual:             &currentCarrierUticaMutual,
	VAFarmBureau:            &currentCarrierVAFarmBureau,
	Wawanesa:                &currentCarrierWawanesa,
	WestBendMutual:          &currentCarrierWestBendMutual,
	WesternGeneral:          &currentCarrierWesternGeneral,
	WesternReservePool:      &currentCarrierWesternReservePool,
	Westfield:               &currentCarrierWestfield,
	WhiteMountains:          &currentCarrierWhiteMountains,
	WisconsinMutual:         &currentCarrierWisconsinMutual,
	WolverineMutual:         &currentCarrierWolverineMutual,
	WorkmensAuto:            &currentCarrierWorkmensAuto,
	WRMAmerica:              &currentCarrierWRMAmerica,
	ZurichInsurance:         &currentCarrierZurichInsurance,

	itemDict: map[string]*EnumCurrentCarrierItem{
		strings.ToLower(string(currentCarrierAAAID)):                     &currentCarrierAAA,
		strings.ToLower(string(currentCarrierAccessID)):                  &currentCarrierAccess,
		strings.ToLower(string(currentCarrierACEGroupID)):                &currentCarrierACEGroup,
		strings.ToLower(string(currentCarrierAffinityMutualID)):          &currentCarrierAffinityMutual,
		strings.ToLower(string(currentCarrierAffirmativeID)):             &currentCarrierAffirmative,
		strings.ToLower(string(currentCarrierAGICID)):                    &currentCarrierAGIC,
		strings.ToLower(string(currentCarrierAIGID)):                     &currentCarrierAIG,
		strings.ToLower(string(currentCarrierAlaskaNationalID)):          &currentCarrierAlaskaNational,
		strings.ToLower(string(currentCarrierAlfaID)):                    &currentCarrierAlfa,
		strings.ToLower(string(currentCarrierAlleghenyID)):               &currentCarrierAllegheny,
		strings.ToLower(string(currentCarrierAllianzID)):                 &currentCarrierAllianz,
		strings.ToLower(string(currentCarrierAlliedWorldID)):             &currentCarrierAlliedWorld,
		strings.ToLower(string(currentCarrierAllstateID)):                &currentCarrierAllstate,
		strings.ToLower(string(currentCarrierAmericanEuropeanID)):        &currentCarrierAmericanEuropean,
		strings.ToLower(string(currentCarrierAmericanFamilyID)):          &currentCarrierAmericanFamily,
		strings.ToLower(string(currentCarrierAmericanFarmAndRanchID)):    &currentCarrierAmericanFarmAndRanch,
		strings.ToLower(string(currentCarrierAmericanIndependentID)):     &currentCarrierAmericanIndependent,
		strings.ToLower(string(currentCarrierAmericanInternationalID)):   &currentCarrierAmericanInternational,
		strings.ToLower(string(currentCarrierAmericanNationalID)):        &currentCarrierAmericanNational,
		strings.ToLower(string(currentCarrierAmericanSterlingID)):        &currentCarrierAmericanSterling,
		strings.ToLower(string(currentCarrierAmeripriseID)):              &currentCarrierAmeriprise,
		strings.ToLower(string(currentCarrierAutoClubEnterprisesID)):     &currentCarrierAutoClubEnterprises,
		strings.ToLower(string(currentCarrierAutoClubOfFloridaID)):       &currentCarrierAutoClubOfFlorida,
		strings.ToLower(string(currentCarrierAutoClubSouthID)):           &currentCarrierAutoClubSouth,
		strings.ToLower(string(currentCarrierAutoOwnersID)):              &currentCarrierAutoOwners,
		strings.ToLower(string(currentCarrierAxisID)):                    &currentCarrierAxis,
		strings.ToLower(string(currentCarrierBaldwinAndLyonsID)):         &currentCarrierBaldwinAndLyons,
		strings.ToLower(string(currentCarrierBrethrenMutualID)):          &currentCarrierBrethrenMutual,
		strings.ToLower(string(currentCarrierCaliforniaCasualtyID)):      &currentCarrierCaliforniaCasualty,
		strings.ToLower(string(currentCarrierCenturyNationalID)):         &currentCarrierCenturyNational,
		strings.ToLower(string(currentCarrierChubbID)):                   &currentCarrierChubb,
		strings.ToLower(string(currentCarrierChurchMutualID)):            &currentCarrierChurchMutual,
		strings.ToLower(string(currentCarrierCivilServiceEmployeesID)):   &currentCarrierCivilServiceEmployees,
		strings.ToLower(string(currentCarrierCNAID)):                     &currentCarrierCNA,
		strings.ToLower(string(currentCarrierCNAFinancialGroupID)):       &currentCarrierCNAFinancialGroup,
		strings.ToLower(string(currentCarrierCompanionID)):               &currentCarrierCompanion,
		strings.ToLower(string(currentCarrierConsumersID)):               &currentCarrierConsumers,
		strings.ToLower(string(currentCarrierCountryFinancialID)):        &currentCarrierCountryFinancial,
		strings.ToLower(string(currentCarrierCountryFinancialPCID)):      &currentCarrierCountryFinancialPC,
		strings.ToLower(string(currentCarrierDirectAutoID)):              &currentCarrierDirectAuto,
		strings.ToLower(string(currentCarrierDirectGeneralID)):           &currentCarrierDirectGeneral,
		strings.ToLower(string(currentCarrierDonegalGroupID)):            &currentCarrierDonegalGroup,
		strings.ToLower(string(currentCarrierEInsID)):                    &currentCarrierEIns,
		strings.ToLower(string(currentCarrierElectricID)):                &currentCarrierElectric,
		strings.ToLower(string(currentCarrierElephantID)):                &currentCarrierElephant,
		strings.ToLower(string(currentCarrierEMCID)):                     &currentCarrierEMC,
		strings.ToLower(string(currentCarrierErieID)):                    &currentCarrierErie,
		strings.ToLower(string(currentCarrierEverestREID)):               &currentCarrierEverestRE,
		strings.ToLower(string(currentCarrierEvergreenUSAID)):            &currentCarrierEvergreenUSA,
		strings.ToLower(string(currentCarrierFairfaxFinancialID)):        &currentCarrierFairfaxFinancial,
		strings.ToLower(string(currentCarrierFarmBureauID)):              &currentCarrierFarmBureau,
		strings.ToLower(string(currentCarrierFarmersID)):                 &currentCarrierFarmers,
		strings.ToLower(string(currentCarrierFarmersMutualID)):           &currentCarrierFarmersMutual,
		strings.ToLower(string(currentCarrierFidelityNationalID)):        &currentCarrierFidelityNational,
		strings.ToLower(string(currentCarrierFirstAcceptanceID)):         &currentCarrierFirstAcceptance,
		strings.ToLower(string(currentCarrierFrankenmuthID)):             &currentCarrierFrankenmuth,
		strings.ToLower(string(currentCarrierFredLoyaID)):                &currentCarrierFredLoya,
		strings.ToLower(string(currentCarrierGeicoID)):                   &currentCarrierGeico,
		strings.ToLower(string(currentCarrierGeorgiaFarmBureauID)):       &currentCarrierGeorgiaFarmBureau,
		strings.ToLower(string(currentCarrierGoodvilleAndGermanID)):      &currentCarrierGoodvilleAndGerman,
		strings.ToLower(string(currentCarrierGrangeID)):                  &currentCarrierGrange,
		strings.ToLower(string(currentCarrierGrangeMutualID)):            &currentCarrierGrangeMutual,
		strings.ToLower(string(currentCarrierGreatAmericanID)):           &currentCarrierGreatAmerican,
		strings.ToLower(string(currentCarrierGuideOneID)):                &currentCarrierGuideOne,
		strings.ToLower(string(currentCarrierHallmarkID)):                &currentCarrierHallmark,
		strings.ToLower(string(currentCarrierHastingsMutualID)):          &currentCarrierHastingsMutual,
		strings.ToLower(string(currentCarrierHiscoxUSAID)):               &currentCarrierHiscoxUSA,
		strings.ToLower(string(currentCarrierHochheimPrairieID)):         &currentCarrierHochheimPrairie,
		strings.ToLower(string(currentCarrierHomeStateID)):               &currentCarrierHomeState,
		strings.ToLower(string(currentCarrierHoraceMannID)):              &currentCarrierHoraceMann,
		strings.ToLower(string(currentCarrierIATID)):                     &currentCarrierIAT,
		strings.ToLower(string(currentCarrierIMTID)):                     &currentCarrierIMT,
		strings.ToLower(string(currentCarrierInfinityID)):                &currentCarrierInfinity,
		strings.ToLower(string(currentCarrierIntegonNationalID)):         &currentCarrierIntegonNational,
		strings.ToLower(string(currentCarrierIronshoreID)):               &currentCarrierIronshore,
		strings.ToLower(string(currentCarrierKemperID)):                  &currentCarrierKemper,
		strings.ToLower(string(currentCarrierKeyID)):                     &currentCarrierKey,
		strings.ToLower(string(currentCarrierKingswayFinancialID)):       &currentCarrierKingswayFinancial,
		strings.ToLower(string(currentCarrierLibertyMutualID)):           &currentCarrierLibertyMutual,
		strings.ToLower(string(currentCarrierMAPFREID)):                  &currentCarrierMAPFRE,
		strings.ToLower(string(currentCarrierMerchantsID)):               &currentCarrierMerchants,
		strings.ToLower(string(currentCarrierMercuryGeneralID)):          &currentCarrierMercuryGeneral,
		strings.ToLower(string(currentCarrierMetLifeID)):                 &currentCarrierMetLife,
		strings.ToLower(string(currentCarrierMichiganFarmBureauID)):      &currentCarrierMichiganFarmBureau,
		strings.ToLower(string(currentCarrierMichiganMillersID)):         &currentCarrierMichiganMillers,
		strings.ToLower(string(currentCarrierMidwestFamilyMutualID)):     &currentCarrierMidwestFamilyMutual,
		strings.ToLower(string(currentCarrierMMGID)):                     &currentCarrierMMG,
		strings.ToLower(string(currentCarrierMonumentID)):                &currentCarrierMonument,
		strings.ToLower(string(currentCarrierMSADUSID)):                  &currentCarrierMSADUS,
		strings.ToLower(string(currentCarrierMunichAmericanID)):          &currentCarrierMunichAmerican,
		strings.ToLower(string(currentCarrierNationalGeneralHoldingsID)): &currentCarrierNationalGeneralHoldings,
		strings.ToLower(string(currentCarrierNationwideID)):              &currentCarrierNationwide,
		strings.ToLower(string(currentCarrierNLCID)):                     &currentCarrierNLC,
		strings.ToLower(string(currentCarrierNodakMutualID)):             &currentCarrierNodakMutual,
		strings.ToLower(string(currentCarrierNonProfitsID)):              &currentCarrierNonProfits,
		strings.ToLower(string(currentCarrierNorfolkAndDedhamID)):        &currentCarrierNorfolkAndDedham,
		strings.ToLower(string(currentCarrierNYCMID)):                    &currentCarrierNYCM,
		strings.ToLower(string(currentCarrierOkalahomaFarmBureauID)):     &currentCarrierOkalahomaFarmBureau,
		strings.ToLower(string(currentCarrierOldAmericaCapitalID)):       &currentCarrierOldAmericaCapital,
		strings.ToLower(string(currentCarrierOregonMutualID)):            &currentCarrierOregonMutual,
		strings.ToLower(string(currentCarrierOtherID)):                   &currentCarrierOther,
		strings.ToLower(string(currentCarrierPacificSpecialtyID)):        &currentCarrierPacificSpecialty,
		strings.ToLower(string(currentCarrierPekinID)):                   &currentCarrierPekin,
		strings.ToLower(string(currentCarrierPennLumbermensID)):          &currentCarrierPennLumbermens,
		strings.ToLower(string(currentCarrierPharmacistsID)):             &currentCarrierPharmacists,
		strings.ToLower(string(currentCarrierPhiladlphiaTokioID)):        &currentCarrierPhiladlphiaTokio,
		strings.ToLower(string(currentCarrierPioneerStateMutualID)):      &currentCarrierPioneerStateMutual,
		strings.ToLower(string(currentCarrierPreferredMutualID)):         &currentCarrierPreferredMutual,
		strings.ToLower(string(currentCarrierProgressiveID)):             &currentCarrierProgressive,
		strings.ToLower(string(currentCarrierProsightSpecialtyID)):       &currentCarrierProsightSpecialty,
		strings.ToLower(string(currentCarrierPublicServiceMutualID)):     &currentCarrierPublicServiceMutual,
		strings.ToLower(string(currentCarrierPureID)):                    &currentCarrierPure,
		strings.ToLower(string(currentCarrierQBEAmericasID)):             &currentCarrierQBEAmericas,
		strings.ToLower(string(currentCarrierRAMMutualID)):               &currentCarrierRAMMutual,
		strings.ToLower(string(currentCarrierRepublicGroupID)):           &currentCarrierRepublicGroup,
		strings.ToLower(string(currentCarrierRLIID)):                     &currentCarrierRLI,
		strings.ToLower(string(currentCarrierSafeAutoID)):                &currentCarrierSafeAuto,
		strings.ToLower(string(currentCarrierSamsungFireAndMarineID)):    &currentCarrierSamsungFireAndMarine,
		strings.ToLower(string(currentCarrierSecuraID)):                  &currentCarrierSecura,
		strings.ToLower(string(currentCarrierSentryID)):                  &currentCarrierSentry,
		strings.ToLower(string(currentCarrierSompoJapanID)):              &currentCarrierSompoJapan,
		strings.ToLower(string(currentCarrierSouthportLaneID)):           &currentCarrierSouthportLane,
		strings.ToLower(string(currentCarrierStarAndShieldID)):           &currentCarrierStarAndShield,
		strings.ToLower(string(currentCarrierStarCasualtyID)):            &currentCarrierStarCasualty,
		strings.ToLower(string(currentCarrierStarrInternationalID)):      &currentCarrierStarrInternational,
		strings.ToLower(string(currentCarrierStateAutoID)):               &currentCarrierStateAuto,
		strings.ToLower(string(currentCarrierStateFarmID)):               &currentCarrierStateFarm,
		strings.ToLower(string(currentCarrierStateNationalID)):           &currentCarrierStateNational,
		strings.ToLower(string(currentCarrierSterlingID)):                &currentCarrierSterling,
		strings.ToLower(string(currentCarrierTheGeneralID)):              &currentCarrierTheGeneral,
		strings.ToLower(string(currentCarrierTheHartfordID)):             &currentCarrierTheHartford,
		strings.ToLower(string(currentCarrierTheRepublicGroupID)):        &currentCarrierTheRepublicGroup,
		strings.ToLower(string(currentCarrierTitusID)):                   &currentCarrierTitus,
		strings.ToLower(string(currentCarrierTopaID)):                    &currentCarrierTopa,
		strings.ToLower(string(currentCarrierTowerID)):                   &currentCarrierTower,
		strings.ToLower(string(currentCarrierTradersGeneralID)):          &currentCarrierTradersGeneral,
		strings.ToLower(string(currentCarrierTravelersID)):               &currentCarrierTravelers,
		strings.ToLower(string(currentCarrierUnitedHomeID)):              &currentCarrierUnitedHome,
		strings.ToLower(string(currentCarrierUnitrinID)):                 &currentCarrierUnitrin,
		strings.ToLower(string(currentCarrierUSAAID)):                    &currentCarrierUSAA,
		strings.ToLower(string(currentCarrierUticaMutualID)):             &currentCarrierUticaMutual,
		strings.ToLower(string(currentCarrierVAFarmBureauID)):            &currentCarrierVAFarmBureau,
		strings.ToLower(string(currentCarrierWawanesaID)):                &currentCarrierWawanesa,
		strings.ToLower(string(currentCarrierWestBendMutualID)):          &currentCarrierWestBendMutual,
		strings.ToLower(string(currentCarrierWesternGeneralID)):          &currentCarrierWesternGeneral,
		strings.ToLower(string(currentCarrierWesternReservePoolID)):      &currentCarrierWesternReservePool,
		strings.ToLower(string(currentCarrierWestfieldID)):               &currentCarrierWestfield,
		strings.ToLower(string(currentCarrierWhiteMountainsID)):          &currentCarrierWhiteMountains,
		strings.ToLower(string(currentCarrierWisconsinMutualID)):         &currentCarrierWisconsinMutual,
		strings.ToLower(string(currentCarrierWolverineMutualID)):         &currentCarrierWolverineMutual,
		strings.ToLower(string(currentCarrierWorkmensAutoID)):            &currentCarrierWorkmensAuto,
		strings.ToLower(string(currentCarrierWRMAmericaID)):              &currentCarrierWRMAmerica,
		strings.ToLower(string(currentCarrierZurichInsuranceID)):         &currentCarrierZurichInsurance,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumCurrentCarrier) ByID(id CurrentCarrierIdentifier) *EnumCurrentCarrierItem {
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
func (e *EnumCurrentCarrier) ByIDString(idx string) *EnumCurrentCarrierItem {
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
func (e *EnumCurrentCarrier) ByIndex(idx int) *EnumCurrentCarrierItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedCurrentCarrierID is a struct that is designed to replace a *CurrentCarrierID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *CurrentCarrierID it contains while being a better JSON citizen.
type ValidatedCurrentCarrierID struct {
	// id will point to a valid CurrentCarrierID, if possible
	// If id is nil, then ValidatedCurrentCarrierID.Valid() will return false.
	id *CurrentCarrierID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedCurrentCarrierID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedCurrentCarrierID
func (vi *ValidatedCurrentCarrierID) Clone() *ValidatedCurrentCarrierID {
	if vi == nil {
		return nil
	}

	var cid *CurrentCarrierID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedCurrentCarrierID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedCurrentCarrierIds represent the same CurrentCarrier
func (vi *ValidatedCurrentCarrierID) Equals(vj *ValidatedCurrentCarrierID) bool {
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

// Valid returns true if and only if the ValidatedCurrentCarrierID corresponds to a recognized CurrentCarrier
func (vi *ValidatedCurrentCarrierID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedCurrentCarrierID) ID() *CurrentCarrierID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedCurrentCarrierID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedCurrentCarrierID) ValidatedID() *ValidatedCurrentCarrierID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedCurrentCarrierID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedCurrentCarrierID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedCurrentCarrierID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedCurrentCarrierID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedCurrentCarrierID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := CurrentCarrierID(capString)
	item := CurrentCarrier.ByID(&id)
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

func (vi ValidatedCurrentCarrierID) String() string {
	return vi.ToIDString()
}

type CurrentCarrierIdentifier interface {
	ID() *CurrentCarrierID
	Valid() bool
}
