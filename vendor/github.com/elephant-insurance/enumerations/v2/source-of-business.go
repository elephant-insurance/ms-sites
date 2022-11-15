package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// SourceOfBusinessID uniquely identifies a particular SourceOfBusiness
type SourceOfBusinessID string

// Clone creates a safe, independent copy of a SourceOfBusinessID
func (i *SourceOfBusinessID) Clone() *SourceOfBusinessID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two SourceOfBusinessIds are equivalent
func (i *SourceOfBusinessID) Equals(j *SourceOfBusinessID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *SourceOfBusinessID that is either valid or nil
func (i *SourceOfBusinessID) ID() *SourceOfBusinessID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *SourceOfBusinessID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the SourceOfBusinessID corresponds to a recognized SourceOfBusiness
func (i *SourceOfBusinessID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return SourceOfBusiness.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *SourceOfBusinessID) ValidatedID() *ValidatedSourceOfBusinessID {
	if i != nil {
		return &ValidatedSourceOfBusinessID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *SourceOfBusinessID) MarshalJSON() ([]byte, error) {
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

func (i *SourceOfBusinessID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := SourceOfBusinessID(dataString)
	item := SourceOfBusiness.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	sourceOfBusinessAdharmonicsID             SourceOfBusinessID = "adharmonics"
	sourceOfBusinessAgentID                   SourceOfBusinessID = "agent"
	sourceOfBusinessAggregatorOutboundID      SourceOfBusinessID = "aggregator_outbound"
	sourceOfBusinessAGWorkersID               SourceOfBusinessID = "agWorkers"
	sourceOfBusinessAirportID                 SourceOfBusinessID = "airport"
	sourceOfBusinessAggregatorAllWebID        SourceOfBusinessID = "aggregator_all_web"
	sourceOfBusinessApparentID                SourceOfBusinessID = "apparent"
	sourceOfBusinessBankrateID                SourceOfBusinessID = "bankrate"
	sourceOfBusinessBillboardID               SourceOfBusinessID = "billboard"
	sourceOfBusinessBoltID                    SourceOfBusinessID = "bolt"
	sourceOfBusinessBrightwayID               SourceOfBusinessID = "brightway"
	sourceOfBusinessCertainlyID               SourceOfBusinessID = "certainly"
	sourceOfBusinessCompareNowID              SourceOfBusinessID = "inspop"
	sourceOfBusinessCoverhoundID              SourceOfBusinessID = "coverhound"
	sourceOfBusinessDatalotID                 SourceOfBusinessID = "datalot"
	sourceOfBusinessFetchID                   SourceOfBusinessID = "Fetch"
	sourceOfBusinessElephantID                SourceOfBusinessID = "elephant"
	sourceOfBusinessFirstConnectID            SourceOfBusinessID = "firstConnect"
	sourceOfBusinessGabiID                    SourceOfBusinessID = "gabi"
	sourceOfBusinessGetJerryID                SourceOfBusinessID = "getJerry"
	sourceOfBusinessHippoInsuranceID          SourceOfBusinessID = "hippo"
	sourceOfBusinessInsuranceComID            SourceOfBusinessID = "insurancedotcom"
	sourceOfBusinessInsuraMatchID             SourceOfBusinessID = "insuraMatch"
	sourceOfBusinessInsurifyID                SourceOfBusinessID = "insurify"
	sourceOfBusinessInternetID                SourceOfBusinessID = "internet"
	sourceOfBusinessKatchID                   SourceOfBusinessID = "katch"
	sourceOfBusinessKemperID                  SourceOfBusinessID = "kemperInsurance"
	sourceOfBusinessKissterraID               SourceOfBusinessID = "kissterra"
	sourceOfBusinessLibertyMutualID           SourceOfBusinessID = "liberty_mutual"
	sourceOfBusinessMediaAlphaID              SourceOfBusinessID = "mediaalpha"
	sourceOfBusinessMossID                    SourceOfBusinessID = "moss"
	sourceOfBusinessMovieID                   SourceOfBusinessID = "movie"
	sourceOfBusinessOnlineTransferID          SourceOfBusinessID = "online_transfer"
	sourceOfBusinessOtherID                   SourceOfBusinessID = "other"
	sourceOfBusinessPolicyPilotID             SourceOfBusinessID = "policyPilot"
	sourceOfBusinessPrintID                   SourceOfBusinessID = "print"
	sourceOfBusinessProsperInsuranceID        SourceOfBusinessID = "prosperInsurance"
	sourceOfBusinessQuinStreetID              SourceOfBusinessID = "quinstreet"
	sourceOfBusinessQuoteLabID                SourceOfBusinessID = "quotelab"
	sourceOfBusinessQuoteWizardID             SourceOfBusinessID = "quotewizard"
	sourceOfBusinessRadioID                   SourceOfBusinessID = "radio"
	sourceOfBusinessRateForceID               SourceOfBusinessID = "RateForce"
	sourceOfBusinessReferralID                SourceOfBusinessID = "referral"
	sourceOfBusinessRepeatCustomerID          SourceOfBusinessID = "repeat_customer"
	sourceOfBusinessReviID                    SourceOfBusinessID = "revi"
	sourceOfBusinessRocklandID                SourceOfBusinessID = "rockland"
	sourceOfBusinessSafeAutoID                SourceOfBusinessID = "safeAuto"
	sourceOfBusinessSmartFinancialID          SourceOfBusinessID = "smartfinancial"
	sourceOfBusinessSuitedConnectorID         SourceOfBusinessID = "suitedConnector"
	sourceOfBusinessSuperiorAccessID          SourceOfBusinessID = "superiorAccess"
	sourceOfBusinessTheLeadCoID               SourceOfBusinessID = "theleadco"
	sourceOfBusinessTVID                      SourceOfBusinessID = "tv"
	sourceOfBusinessUndergroundElephantID     SourceOfBusinessID = "underground_elephant"
	sourceOfBusinessVantageID                 SourceOfBusinessID = "vantage"
	sourceOfBusinessWebID                     SourceOfBusinessID = "web"
	sourceOfBusinessZebraID                   SourceOfBusinessID = "zebra"
	sourceOfBusinessAstoriaID                 SourceOfBusinessID = "astoria"
	sourceOfBusinessMediaForceID              SourceOfBusinessID = "mediaForce"
	sourceOfBusinessAgileRatesID              SourceOfBusinessID = "agileRates"
	sourceOfBusinessNextGenLeadsID            SourceOfBusinessID = "nextGenLeads"
	sourceOfBusinessTransparentlyID           SourceOfBusinessID = "transparently"
	sourceOfBusinessAllWebID                  SourceOfBusinessID = "allWeb"
	sourceOfBusinessAssuranceID               SourceOfBusinessID = "assurance"
	sourceOfBusinessGooseheadID               SourceOfBusinessID = "goosehead"
	sourceOfBusinessAgenteroID                SourceOfBusinessID = "agentero"
	sourceOfBusinessCoverageComID             SourceOfBusinessID = "coverageDotCom"
	sourceOfBusinessBindrightID               SourceOfBusinessID = "bindRight"
	sourceOfBusinessNationWideID              SourceOfBusinessID = "nationWide"
	sourceOfBusinessFinderID                  SourceOfBusinessID = "finder"
	sourceOfBusinessAppalachianUnderwritersID SourceOfBusinessID = "appalachianUnderwriters"
	sourceOfBusinessRenegadeID                SourceOfBusinessID = "renegade"
	sourceOfBusinessMaticID                   SourceOfBusinessID = "matic"
	sourceOfBusinessMianID                    SourceOfBusinessID = "mian"
	sourceOfBusinessSmartChoiceID             SourceOfBusinessID = "smartChoice"
	sourceOfBusinessWayID                     SourceOfBusinessID = "way"
	sourceOfBusinessWoodlandFinancialID       SourceOfBusinessID = "woodlandFinancial"
	sourceOfBusinessCreditKarmaID             SourceOfBusinessID = "creditKarma"
	sourceOfBusinessAgentsAllianceID          SourceOfBusinessID = "agentsAlliance"
	sourceOfBusinessSavvyID                   SourceOfBusinessID = "savvy"
	sourceOfBusinessIntegraID                 SourceOfBusinessID = "integra"
	sourceOfBusinessCaribouID                 SourceOfBusinessID = "caribou"
)

// EnumSourceOfBusinessItem describes an entry in an enumeration of SourceOfBusiness
type EnumSourceOfBusinessItem struct {
	ID        SourceOfBusinessID `json:"Value"`
	Desc      string             `json:"Description,omitempty"`
	Meta      map[string]string  `json:",omitempty"`
	Name      string             `json:"Name"`
	SortOrder int

	// Meta Properties
	IsSuperClick       string
	IsPriceComparision string
}

var (
	sourceOfBusinessAdharmonics             = EnumSourceOfBusinessItem{sourceOfBusinessAdharmonicsID, "Adharmonics", map[string]string{"IsSuperClick": "true", "IsPriceComparision": "false"}, "Adharmonics", 1, "true", "false"}
	sourceOfBusinessAgent                   = EnumSourceOfBusinessItem{sourceOfBusinessAgentID, "Agent", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Agent", 2, "false", "false"}
	sourceOfBusinessAggregatorOutbound      = EnumSourceOfBusinessItem{sourceOfBusinessAggregatorOutboundID, "Aggregator Outbound", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "AggregatorOutbound", 3, "false", "false"}
	sourceOfBusinessAGWorkers               = EnumSourceOfBusinessItem{sourceOfBusinessAGWorkersID, "AGWorkers", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "AGWorkers", 4, "false", "true"}
	sourceOfBusinessAirport                 = EnumSourceOfBusinessItem{sourceOfBusinessAirportID, "Airport", map[string]string{"IsSuperClick": "", "IsPriceComparision": ""}, "Airport", 5, "", ""}
	sourceOfBusinessAggregatorAllWeb        = EnumSourceOfBusinessItem{sourceOfBusinessAggregatorAllWebID, "Aggregator AllWeb", map[string]string{"IsSuperClick": "true", "IsPriceComparision": "false"}, "AggregatorAllWeb", 6, "true", "false"}
	sourceOfBusinessApparent                = EnumSourceOfBusinessItem{sourceOfBusinessApparentID, "Apparent", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Apparent", 7, "false", "false"}
	sourceOfBusinessBankrate                = EnumSourceOfBusinessItem{sourceOfBusinessBankrateID, "Bankrate", map[string]string{"IsSuperClick": "true", "IsPriceComparision": "false"}, "Bankrate", 8, "true", "false"}
	sourceOfBusinessBillboard               = EnumSourceOfBusinessItem{sourceOfBusinessBillboardID, "Billboard", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Billboard", 9, "false", "false"}
	sourceOfBusinessBolt                    = EnumSourceOfBusinessItem{sourceOfBusinessBoltID, "Bolt", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Bolt", 10, "false", "true"}
	sourceOfBusinessBrightway               = EnumSourceOfBusinessItem{sourceOfBusinessBrightwayID, "Brightway", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Brightway", 11, "false", "true"}
	sourceOfBusinessCertainly               = EnumSourceOfBusinessItem{sourceOfBusinessCertainlyID, "Certainly", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Certainly", 12, "false", "true"}
	sourceOfBusinessCompareNow              = EnumSourceOfBusinessItem{sourceOfBusinessCompareNowID, "CompareNow", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "CompareNow", 13, "false", "true"}
	sourceOfBusinessCoverhound              = EnumSourceOfBusinessItem{sourceOfBusinessCoverhoundID, "Coverhound", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Coverhound", 14, "false", "true"}
	sourceOfBusinessDatalot                 = EnumSourceOfBusinessItem{sourceOfBusinessDatalotID, "Datalot", map[string]string{"IsSuperClick": "true", "IsPriceComparision": "false"}, "Datalot", 15, "true", "false"}
	sourceOfBusinessFetch                   = EnumSourceOfBusinessItem{sourceOfBusinessFetchID, "Fetch", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Fetch", 16, "false", "true"}
	sourceOfBusinessElephant                = EnumSourceOfBusinessItem{sourceOfBusinessElephantID, "Elephant", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Elephant", 17, "false", "true"}
	sourceOfBusinessFirstConnect            = EnumSourceOfBusinessItem{sourceOfBusinessFirstConnectID, "First Connect", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "FirstConnect", 18, "false", "true"}
	sourceOfBusinessGabi                    = EnumSourceOfBusinessItem{sourceOfBusinessGabiID, "Gabi", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Gabi", 19, "false", "true"}
	sourceOfBusinessGetJerry                = EnumSourceOfBusinessItem{sourceOfBusinessGetJerryID, "GetJerry", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "GetJerry", 20, "false", "true"}
	sourceOfBusinessHippoInsurance          = EnumSourceOfBusinessItem{sourceOfBusinessHippoInsuranceID, "Hippo Insurance", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "HippoInsurance", 21, "false", "true"}
	sourceOfBusinessInsuranceCom            = EnumSourceOfBusinessItem{sourceOfBusinessInsuranceComID, "Insurance.com", map[string]string{"IsSuperClick": "true", "IsPriceComparision": "true"}, "InsuranceCom", 22, "true", "true"}
	sourceOfBusinessInsuraMatch             = EnumSourceOfBusinessItem{sourceOfBusinessInsuraMatchID, "InsuraMatch", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "InsuraMatch", 23, "false", "false"}
	sourceOfBusinessInsurify                = EnumSourceOfBusinessItem{sourceOfBusinessInsurifyID, "Insurify", map[string]string{"IsSuperClick": "true", "IsPriceComparision": "true"}, "Insurify", 24, "true", "true"}
	sourceOfBusinessInternet                = EnumSourceOfBusinessItem{sourceOfBusinessInternetID, "Internet", map[string]string{"IsSuperClick": "", "IsPriceComparision": ""}, "Internet", 25, "", ""}
	sourceOfBusinessKatch                   = EnumSourceOfBusinessItem{sourceOfBusinessKatchID, "Katch", map[string]string{"IsSuperClick": "true", "IsPriceComparision": "false"}, "Katch", 26, "true", "false"}
	sourceOfBusinessKemper                  = EnumSourceOfBusinessItem{sourceOfBusinessKemperID, "Kemper", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Kemper", 27, "false", "true"}
	sourceOfBusinessKissterra               = EnumSourceOfBusinessItem{sourceOfBusinessKissterraID, "Kissterra", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Kissterra", 28, "false", "false"}
	sourceOfBusinessLibertyMutual           = EnumSourceOfBusinessItem{sourceOfBusinessLibertyMutualID, "Liberty Mutual", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "LibertyMutual", 29, "false", "true"}
	sourceOfBusinessMediaAlpha              = EnumSourceOfBusinessItem{sourceOfBusinessMediaAlphaID, "MediaAlpha", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "MediaAlpha", 30, "false", "false"}
	sourceOfBusinessMoss                    = EnumSourceOfBusinessItem{sourceOfBusinessMossID, "Moss", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Moss", 31, "false", "false"}
	sourceOfBusinessMovie                   = EnumSourceOfBusinessItem{sourceOfBusinessMovieID, "Movie", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Movie", 32, "false", "false"}
	sourceOfBusinessOnlineTransfer          = EnumSourceOfBusinessItem{sourceOfBusinessOnlineTransferID, "Online Transfer", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "OnlineTransfer", 33, "false", "false"}
	sourceOfBusinessOther                   = EnumSourceOfBusinessItem{sourceOfBusinessOtherID, "Other", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Other", 34, "false", "false"}
	sourceOfBusinessPolicyPilot             = EnumSourceOfBusinessItem{sourceOfBusinessPolicyPilotID, "Policy Pilot", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "PolicyPilot", 35, "false", "false"}
	sourceOfBusinessPrint                   = EnumSourceOfBusinessItem{sourceOfBusinessPrintID, "Print", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Print", 36, "false", "false"}
	sourceOfBusinessProsperInsurance        = EnumSourceOfBusinessItem{sourceOfBusinessProsperInsuranceID, "Prosper Insurance", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "ProsperInsurance", 37, "false", "true"}
	sourceOfBusinessQuinStreet              = EnumSourceOfBusinessItem{sourceOfBusinessQuinStreetID, "QuinStreet", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "QuinStreet", 38, "false", "false"}
	sourceOfBusinessQuoteLab                = EnumSourceOfBusinessItem{sourceOfBusinessQuoteLabID, "Quote Lab", map[string]string{"IsSuperClick": "true", "IsPriceComparision": "false"}, "QuoteLab", 39, "true", "false"}
	sourceOfBusinessQuoteWizard             = EnumSourceOfBusinessItem{sourceOfBusinessQuoteWizardID, "QuoteWizard", map[string]string{"IsSuperClick": "true", "IsPriceComparision": "false"}, "QuoteWizard", 40, "true", "false"}
	sourceOfBusinessRadio                   = EnumSourceOfBusinessItem{sourceOfBusinessRadioID, "Radio", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Radio", 41, "false", "false"}
	sourceOfBusinessRateForce               = EnumSourceOfBusinessItem{sourceOfBusinessRateForceID, "RateForce", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "RateForce", 42, "false", "true"}
	sourceOfBusinessReferral                = EnumSourceOfBusinessItem{sourceOfBusinessReferralID, "Referral", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Referral", 43, "false", "false"}
	sourceOfBusinessRepeatCustomer          = EnumSourceOfBusinessItem{sourceOfBusinessRepeatCustomerID, "Repeat Customer", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "RepeatCustomer", 44, "false", "false"}
	sourceOfBusinessRevi                    = EnumSourceOfBusinessItem{sourceOfBusinessReviID, "Revi", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Revi", 45, "false", "false"}
	sourceOfBusinessRockland                = EnumSourceOfBusinessItem{sourceOfBusinessRocklandID, "Rockland", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Rockland", 46, "false", "true"}
	sourceOfBusinessSafeAuto                = EnumSourceOfBusinessItem{sourceOfBusinessSafeAutoID, "safeAuto", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "SafeAuto", 47, "false", "false"}
	sourceOfBusinessSmartFinancial          = EnumSourceOfBusinessItem{sourceOfBusinessSmartFinancialID, "Smart Financial", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "SmartFinancial", 48, "false", "false"}
	sourceOfBusinessSuitedConnector         = EnumSourceOfBusinessItem{sourceOfBusinessSuitedConnectorID, "Suited Connector", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "SuitedConnector", 49, "false", "false"}
	sourceOfBusinessSuperiorAccess          = EnumSourceOfBusinessItem{sourceOfBusinessSuperiorAccessID, "Superior Access", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "SuperiorAccess", 50, "false", "true"}
	sourceOfBusinessTheLeadCo               = EnumSourceOfBusinessItem{sourceOfBusinessTheLeadCoID, "TheLeadCo", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "TheLeadCo", 51, "false", "false"}
	sourceOfBusinessTV                      = EnumSourceOfBusinessItem{sourceOfBusinessTVID, "TV", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "TV", 52, "false", "false"}
	sourceOfBusinessUndergroundElephant     = EnumSourceOfBusinessItem{sourceOfBusinessUndergroundElephantID, "Underground Elephant", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "UndergroundElephant", 53, "false", "false"}
	sourceOfBusinessVantage                 = EnumSourceOfBusinessItem{sourceOfBusinessVantageID, "Vantage", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Vantage", 54, "false", "false"}
	sourceOfBusinessWeb                     = EnumSourceOfBusinessItem{sourceOfBusinessWebID, "Web", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Web", 55, "false", "false"}
	sourceOfBusinessZebra                   = EnumSourceOfBusinessItem{sourceOfBusinessZebraID, "Zebra", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Zebra", 56, "false", "true"}
	sourceOfBusinessAstoria                 = EnumSourceOfBusinessItem{sourceOfBusinessAstoriaID, "Astoria", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Astoria", 57, "false", "false"}
	sourceOfBusinessMediaForce              = EnumSourceOfBusinessItem{sourceOfBusinessMediaForceID, "Media Force", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "MediaForce", 58, "false", "false"}
	sourceOfBusinessAgileRates              = EnumSourceOfBusinessItem{sourceOfBusinessAgileRatesID, "Agile Rates", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "AgileRates", 59, "false", "false"}
	sourceOfBusinessNextGenLeads            = EnumSourceOfBusinessItem{sourceOfBusinessNextGenLeadsID, "NextGenLeads", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "NextGenLeads", 60, "false", "false"}
	sourceOfBusinessTransparently           = EnumSourceOfBusinessItem{sourceOfBusinessTransparentlyID, "Transparently", map[string]string{"IsSuperClick": "true", "IsPriceComparision": "false"}, "Transparently", 61, "true", "false"}
	sourceOfBusinessAllWeb                  = EnumSourceOfBusinessItem{sourceOfBusinessAllWebID, "AllWeb", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "AllWeb", 62, "false", "false"}
	sourceOfBusinessAssurance               = EnumSourceOfBusinessItem{sourceOfBusinessAssuranceID, "Assurance", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Assurance", 63, "false", "false"}
	sourceOfBusinessGoosehead               = EnumSourceOfBusinessItem{sourceOfBusinessGooseheadID, "Goosehead", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Goosehead", 64, "false", "false"}
	sourceOfBusinessAgentero                = EnumSourceOfBusinessItem{sourceOfBusinessAgenteroID, "Agentero", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Agentero", 65, "false", "false"}
	sourceOfBusinessCoverageCom             = EnumSourceOfBusinessItem{sourceOfBusinessCoverageComID, "Coverage.com", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "CoverageCom", 66, "false", "true"}
	sourceOfBusinessBindright               = EnumSourceOfBusinessItem{sourceOfBusinessBindrightID, "Bindright", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Bindright", 67, "false", "true"}
	sourceOfBusinessNationWide              = EnumSourceOfBusinessItem{sourceOfBusinessNationWideID, "NationWide", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "NationWide", 68, "false", "true"}
	sourceOfBusinessFinder                  = EnumSourceOfBusinessItem{sourceOfBusinessFinderID, "Finder", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Finder", 69, "false", "true"}
	sourceOfBusinessAppalachianUnderwriters = EnumSourceOfBusinessItem{sourceOfBusinessAppalachianUnderwritersID, "AppalachianUnderwriters", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "AppalachianUnderwriters", 70, "false", "false"}
	sourceOfBusinessRenegade                = EnumSourceOfBusinessItem{sourceOfBusinessRenegadeID, "Renegade", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Renegade", 71, "false", "true"}
	sourceOfBusinessMatic                   = EnumSourceOfBusinessItem{sourceOfBusinessMaticID, "Matic", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Matic", 72, "false", "false"}
	sourceOfBusinessMian                    = EnumSourceOfBusinessItem{sourceOfBusinessMianID, "Mian", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "Mian", 73, "false", "false"}
	sourceOfBusinessSmartChoice             = EnumSourceOfBusinessItem{sourceOfBusinessSmartChoiceID, "SmartChoice", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "SmartChoice", 74, "false", "false"}
	sourceOfBusinessWay                     = EnumSourceOfBusinessItem{sourceOfBusinessWayID, "Way", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Way", 75, "false", "true"}
	sourceOfBusinessWoodlandFinancial       = EnumSourceOfBusinessItem{sourceOfBusinessWoodlandFinancialID, "WoodlandFinancial", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "WoodlandFinancial", 76, "false", "true"}
	sourceOfBusinessCreditKarma             = EnumSourceOfBusinessItem{sourceOfBusinessCreditKarmaID, "CreditKarma", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "CreditKarma", 77, "false", "true"}
	sourceOfBusinessAgentsAlliance          = EnumSourceOfBusinessItem{sourceOfBusinessAgentsAllianceID, "AgentsAlliance", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "false"}, "AgentsAlliance", 78, "false", "false"}
	sourceOfBusinessSavvy                   = EnumSourceOfBusinessItem{sourceOfBusinessSavvyID, "Savvy", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Savvy", 79, "false", "true"}
	sourceOfBusinessIntegra                 = EnumSourceOfBusinessItem{sourceOfBusinessIntegraID, "Integra", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Integra", 80, "false", "true"}
	sourceOfBusinessCaribou                 = EnumSourceOfBusinessItem{sourceOfBusinessCaribouID, "Caribou", map[string]string{"IsSuperClick": "false", "IsPriceComparision": "true"}, "Caribou", 81, "false", "true"}
)

// EnumSourceOfBusiness is a collection of SourceOfBusiness items
type EnumSourceOfBusiness struct {
	Description string
	Items       []*EnumSourceOfBusinessItem
	Name        string

	Adharmonics             *EnumSourceOfBusinessItem
	Agent                   *EnumSourceOfBusinessItem
	AggregatorOutbound      *EnumSourceOfBusinessItem
	AGWorkers               *EnumSourceOfBusinessItem
	Airport                 *EnumSourceOfBusinessItem
	AggregatorAllWeb        *EnumSourceOfBusinessItem
	Apparent                *EnumSourceOfBusinessItem
	Bankrate                *EnumSourceOfBusinessItem
	Billboard               *EnumSourceOfBusinessItem
	Bolt                    *EnumSourceOfBusinessItem
	Brightway               *EnumSourceOfBusinessItem
	Certainly               *EnumSourceOfBusinessItem
	CompareNow              *EnumSourceOfBusinessItem
	Coverhound              *EnumSourceOfBusinessItem
	Datalot                 *EnumSourceOfBusinessItem
	Fetch                   *EnumSourceOfBusinessItem
	Elephant                *EnumSourceOfBusinessItem
	FirstConnect            *EnumSourceOfBusinessItem
	Gabi                    *EnumSourceOfBusinessItem
	GetJerry                *EnumSourceOfBusinessItem
	HippoInsurance          *EnumSourceOfBusinessItem
	InsuranceCom            *EnumSourceOfBusinessItem
	InsuraMatch             *EnumSourceOfBusinessItem
	Insurify                *EnumSourceOfBusinessItem
	Internet                *EnumSourceOfBusinessItem
	Katch                   *EnumSourceOfBusinessItem
	Kemper                  *EnumSourceOfBusinessItem
	Kissterra               *EnumSourceOfBusinessItem
	LibertyMutual           *EnumSourceOfBusinessItem
	MediaAlpha              *EnumSourceOfBusinessItem
	Moss                    *EnumSourceOfBusinessItem
	Movie                   *EnumSourceOfBusinessItem
	OnlineTransfer          *EnumSourceOfBusinessItem
	Other                   *EnumSourceOfBusinessItem
	PolicyPilot             *EnumSourceOfBusinessItem
	Print                   *EnumSourceOfBusinessItem
	ProsperInsurance        *EnumSourceOfBusinessItem
	QuinStreet              *EnumSourceOfBusinessItem
	QuoteLab                *EnumSourceOfBusinessItem
	QuoteWizard             *EnumSourceOfBusinessItem
	Radio                   *EnumSourceOfBusinessItem
	RateForce               *EnumSourceOfBusinessItem
	Referral                *EnumSourceOfBusinessItem
	RepeatCustomer          *EnumSourceOfBusinessItem
	Revi                    *EnumSourceOfBusinessItem
	Rockland                *EnumSourceOfBusinessItem
	SafeAuto                *EnumSourceOfBusinessItem
	SmartFinancial          *EnumSourceOfBusinessItem
	SuitedConnector         *EnumSourceOfBusinessItem
	SuperiorAccess          *EnumSourceOfBusinessItem
	TheLeadCo               *EnumSourceOfBusinessItem
	TV                      *EnumSourceOfBusinessItem
	UndergroundElephant     *EnumSourceOfBusinessItem
	Vantage                 *EnumSourceOfBusinessItem
	Web                     *EnumSourceOfBusinessItem
	Zebra                   *EnumSourceOfBusinessItem
	Astoria                 *EnumSourceOfBusinessItem
	MediaForce              *EnumSourceOfBusinessItem
	AgileRates              *EnumSourceOfBusinessItem
	NextGenLeads            *EnumSourceOfBusinessItem
	Transparently           *EnumSourceOfBusinessItem
	AllWeb                  *EnumSourceOfBusinessItem
	Assurance               *EnumSourceOfBusinessItem
	Goosehead               *EnumSourceOfBusinessItem
	Agentero                *EnumSourceOfBusinessItem
	CoverageCom             *EnumSourceOfBusinessItem
	Bindright               *EnumSourceOfBusinessItem
	NationWide              *EnumSourceOfBusinessItem
	Finder                  *EnumSourceOfBusinessItem
	AppalachianUnderwriters *EnumSourceOfBusinessItem
	Renegade                *EnumSourceOfBusinessItem
	Matic                   *EnumSourceOfBusinessItem
	Mian                    *EnumSourceOfBusinessItem
	SmartChoice             *EnumSourceOfBusinessItem
	Way                     *EnumSourceOfBusinessItem
	WoodlandFinancial       *EnumSourceOfBusinessItem
	CreditKarma             *EnumSourceOfBusinessItem
	AgentsAlliance          *EnumSourceOfBusinessItem
	Savvy                   *EnumSourceOfBusinessItem
	Integra                 *EnumSourceOfBusinessItem
	Caribou                 *EnumSourceOfBusinessItem

	itemDict map[string]*EnumSourceOfBusinessItem
}

// SourceOfBusiness is a public singleton instance of EnumSourceOfBusiness
// representing sources of quote business
var SourceOfBusiness = &EnumSourceOfBusiness{
	Description: "sources of quote business",
	Items: []*EnumSourceOfBusinessItem{
		&sourceOfBusinessAdharmonics,
		&sourceOfBusinessAgent,
		&sourceOfBusinessAggregatorOutbound,
		&sourceOfBusinessAGWorkers,
		&sourceOfBusinessAirport,
		&sourceOfBusinessAggregatorAllWeb,
		&sourceOfBusinessApparent,
		&sourceOfBusinessBankrate,
		&sourceOfBusinessBillboard,
		&sourceOfBusinessBolt,
		&sourceOfBusinessBrightway,
		&sourceOfBusinessCertainly,
		&sourceOfBusinessCompareNow,
		&sourceOfBusinessCoverhound,
		&sourceOfBusinessDatalot,
		&sourceOfBusinessFetch,
		&sourceOfBusinessElephant,
		&sourceOfBusinessFirstConnect,
		&sourceOfBusinessGabi,
		&sourceOfBusinessGetJerry,
		&sourceOfBusinessHippoInsurance,
		&sourceOfBusinessInsuranceCom,
		&sourceOfBusinessInsuraMatch,
		&sourceOfBusinessInsurify,
		&sourceOfBusinessInternet,
		&sourceOfBusinessKatch,
		&sourceOfBusinessKemper,
		&sourceOfBusinessKissterra,
		&sourceOfBusinessLibertyMutual,
		&sourceOfBusinessMediaAlpha,
		&sourceOfBusinessMoss,
		&sourceOfBusinessMovie,
		&sourceOfBusinessOnlineTransfer,
		&sourceOfBusinessOther,
		&sourceOfBusinessPolicyPilot,
		&sourceOfBusinessPrint,
		&sourceOfBusinessProsperInsurance,
		&sourceOfBusinessQuinStreet,
		&sourceOfBusinessQuoteLab,
		&sourceOfBusinessQuoteWizard,
		&sourceOfBusinessRadio,
		&sourceOfBusinessRateForce,
		&sourceOfBusinessReferral,
		&sourceOfBusinessRepeatCustomer,
		&sourceOfBusinessRevi,
		&sourceOfBusinessRockland,
		&sourceOfBusinessSafeAuto,
		&sourceOfBusinessSmartFinancial,
		&sourceOfBusinessSuitedConnector,
		&sourceOfBusinessSuperiorAccess,
		&sourceOfBusinessTheLeadCo,
		&sourceOfBusinessTV,
		&sourceOfBusinessUndergroundElephant,
		&sourceOfBusinessVantage,
		&sourceOfBusinessWeb,
		&sourceOfBusinessZebra,
		&sourceOfBusinessAstoria,
		&sourceOfBusinessMediaForce,
		&sourceOfBusinessAgileRates,
		&sourceOfBusinessNextGenLeads,
		&sourceOfBusinessTransparently,
		&sourceOfBusinessAllWeb,
		&sourceOfBusinessAssurance,
		&sourceOfBusinessGoosehead,
		&sourceOfBusinessAgentero,
		&sourceOfBusinessCoverageCom,
		&sourceOfBusinessBindright,
		&sourceOfBusinessNationWide,
		&sourceOfBusinessFinder,
		&sourceOfBusinessAppalachianUnderwriters,
		&sourceOfBusinessRenegade,
		&sourceOfBusinessMatic,
		&sourceOfBusinessMian,
		&sourceOfBusinessSmartChoice,
		&sourceOfBusinessWay,
		&sourceOfBusinessWoodlandFinancial,
		&sourceOfBusinessCreditKarma,
		&sourceOfBusinessAgentsAlliance,
		&sourceOfBusinessSavvy,
		&sourceOfBusinessIntegra,
		&sourceOfBusinessCaribou,
	},
	Name:                    "EnumSourceOfBusiness",
	Adharmonics:             &sourceOfBusinessAdharmonics,
	Agent:                   &sourceOfBusinessAgent,
	AggregatorOutbound:      &sourceOfBusinessAggregatorOutbound,
	AGWorkers:               &sourceOfBusinessAGWorkers,
	Airport:                 &sourceOfBusinessAirport,
	AggregatorAllWeb:        &sourceOfBusinessAggregatorAllWeb,
	Apparent:                &sourceOfBusinessApparent,
	Bankrate:                &sourceOfBusinessBankrate,
	Billboard:               &sourceOfBusinessBillboard,
	Bolt:                    &sourceOfBusinessBolt,
	Brightway:               &sourceOfBusinessBrightway,
	Certainly:               &sourceOfBusinessCertainly,
	CompareNow:              &sourceOfBusinessCompareNow,
	Coverhound:              &sourceOfBusinessCoverhound,
	Datalot:                 &sourceOfBusinessDatalot,
	Fetch:                   &sourceOfBusinessFetch,
	Elephant:                &sourceOfBusinessElephant,
	FirstConnect:            &sourceOfBusinessFirstConnect,
	Gabi:                    &sourceOfBusinessGabi,
	GetJerry:                &sourceOfBusinessGetJerry,
	HippoInsurance:          &sourceOfBusinessHippoInsurance,
	InsuranceCom:            &sourceOfBusinessInsuranceCom,
	InsuraMatch:             &sourceOfBusinessInsuraMatch,
	Insurify:                &sourceOfBusinessInsurify,
	Internet:                &sourceOfBusinessInternet,
	Katch:                   &sourceOfBusinessKatch,
	Kemper:                  &sourceOfBusinessKemper,
	Kissterra:               &sourceOfBusinessKissterra,
	LibertyMutual:           &sourceOfBusinessLibertyMutual,
	MediaAlpha:              &sourceOfBusinessMediaAlpha,
	Moss:                    &sourceOfBusinessMoss,
	Movie:                   &sourceOfBusinessMovie,
	OnlineTransfer:          &sourceOfBusinessOnlineTransfer,
	Other:                   &sourceOfBusinessOther,
	PolicyPilot:             &sourceOfBusinessPolicyPilot,
	Print:                   &sourceOfBusinessPrint,
	ProsperInsurance:        &sourceOfBusinessProsperInsurance,
	QuinStreet:              &sourceOfBusinessQuinStreet,
	QuoteLab:                &sourceOfBusinessQuoteLab,
	QuoteWizard:             &sourceOfBusinessQuoteWizard,
	Radio:                   &sourceOfBusinessRadio,
	RateForce:               &sourceOfBusinessRateForce,
	Referral:                &sourceOfBusinessReferral,
	RepeatCustomer:          &sourceOfBusinessRepeatCustomer,
	Revi:                    &sourceOfBusinessRevi,
	Rockland:                &sourceOfBusinessRockland,
	SafeAuto:                &sourceOfBusinessSafeAuto,
	SmartFinancial:          &sourceOfBusinessSmartFinancial,
	SuitedConnector:         &sourceOfBusinessSuitedConnector,
	SuperiorAccess:          &sourceOfBusinessSuperiorAccess,
	TheLeadCo:               &sourceOfBusinessTheLeadCo,
	TV:                      &sourceOfBusinessTV,
	UndergroundElephant:     &sourceOfBusinessUndergroundElephant,
	Vantage:                 &sourceOfBusinessVantage,
	Web:                     &sourceOfBusinessWeb,
	Zebra:                   &sourceOfBusinessZebra,
	Astoria:                 &sourceOfBusinessAstoria,
	MediaForce:              &sourceOfBusinessMediaForce,
	AgileRates:              &sourceOfBusinessAgileRates,
	NextGenLeads:            &sourceOfBusinessNextGenLeads,
	Transparently:           &sourceOfBusinessTransparently,
	AllWeb:                  &sourceOfBusinessAllWeb,
	Assurance:               &sourceOfBusinessAssurance,
	Goosehead:               &sourceOfBusinessGoosehead,
	Agentero:                &sourceOfBusinessAgentero,
	CoverageCom:             &sourceOfBusinessCoverageCom,
	Bindright:               &sourceOfBusinessBindright,
	NationWide:              &sourceOfBusinessNationWide,
	Finder:                  &sourceOfBusinessFinder,
	AppalachianUnderwriters: &sourceOfBusinessAppalachianUnderwriters,
	Renegade:                &sourceOfBusinessRenegade,
	Matic:                   &sourceOfBusinessMatic,
	Mian:                    &sourceOfBusinessMian,
	SmartChoice:             &sourceOfBusinessSmartChoice,
	Way:                     &sourceOfBusinessWay,
	WoodlandFinancial:       &sourceOfBusinessWoodlandFinancial,
	CreditKarma:             &sourceOfBusinessCreditKarma,
	AgentsAlliance:          &sourceOfBusinessAgentsAlliance,
	Savvy:                   &sourceOfBusinessSavvy,
	Integra:                 &sourceOfBusinessIntegra,
	Caribou:                 &sourceOfBusinessCaribou,

	itemDict: map[string]*EnumSourceOfBusinessItem{
		strings.ToLower(string(sourceOfBusinessAdharmonicsID)):             &sourceOfBusinessAdharmonics,
		strings.ToLower(string(sourceOfBusinessAgentID)):                   &sourceOfBusinessAgent,
		strings.ToLower(string(sourceOfBusinessAggregatorOutboundID)):      &sourceOfBusinessAggregatorOutbound,
		strings.ToLower(string(sourceOfBusinessAGWorkersID)):               &sourceOfBusinessAGWorkers,
		strings.ToLower(string(sourceOfBusinessAirportID)):                 &sourceOfBusinessAirport,
		strings.ToLower(string(sourceOfBusinessAggregatorAllWebID)):        &sourceOfBusinessAggregatorAllWeb,
		strings.ToLower(string(sourceOfBusinessApparentID)):                &sourceOfBusinessApparent,
		strings.ToLower(string(sourceOfBusinessBankrateID)):                &sourceOfBusinessBankrate,
		strings.ToLower(string(sourceOfBusinessBillboardID)):               &sourceOfBusinessBillboard,
		strings.ToLower(string(sourceOfBusinessBoltID)):                    &sourceOfBusinessBolt,
		strings.ToLower(string(sourceOfBusinessBrightwayID)):               &sourceOfBusinessBrightway,
		strings.ToLower(string(sourceOfBusinessCertainlyID)):               &sourceOfBusinessCertainly,
		strings.ToLower(string(sourceOfBusinessCompareNowID)):              &sourceOfBusinessCompareNow,
		strings.ToLower(string(sourceOfBusinessCoverhoundID)):              &sourceOfBusinessCoverhound,
		strings.ToLower(string(sourceOfBusinessDatalotID)):                 &sourceOfBusinessDatalot,
		strings.ToLower(string(sourceOfBusinessFetchID)):                   &sourceOfBusinessFetch,
		strings.ToLower(string(sourceOfBusinessElephantID)):                &sourceOfBusinessElephant,
		strings.ToLower(string(sourceOfBusinessFirstConnectID)):            &sourceOfBusinessFirstConnect,
		strings.ToLower(string(sourceOfBusinessGabiID)):                    &sourceOfBusinessGabi,
		strings.ToLower(string(sourceOfBusinessGetJerryID)):                &sourceOfBusinessGetJerry,
		strings.ToLower(string(sourceOfBusinessHippoInsuranceID)):          &sourceOfBusinessHippoInsurance,
		strings.ToLower(string(sourceOfBusinessInsuranceComID)):            &sourceOfBusinessInsuranceCom,
		strings.ToLower(string(sourceOfBusinessInsuraMatchID)):             &sourceOfBusinessInsuraMatch,
		strings.ToLower(string(sourceOfBusinessInsurifyID)):                &sourceOfBusinessInsurify,
		strings.ToLower(string(sourceOfBusinessInternetID)):                &sourceOfBusinessInternet,
		strings.ToLower(string(sourceOfBusinessKatchID)):                   &sourceOfBusinessKatch,
		strings.ToLower(string(sourceOfBusinessKemperID)):                  &sourceOfBusinessKemper,
		strings.ToLower(string(sourceOfBusinessKissterraID)):               &sourceOfBusinessKissterra,
		strings.ToLower(string(sourceOfBusinessLibertyMutualID)):           &sourceOfBusinessLibertyMutual,
		strings.ToLower(string(sourceOfBusinessMediaAlphaID)):              &sourceOfBusinessMediaAlpha,
		strings.ToLower(string(sourceOfBusinessMossID)):                    &sourceOfBusinessMoss,
		strings.ToLower(string(sourceOfBusinessMovieID)):                   &sourceOfBusinessMovie,
		strings.ToLower(string(sourceOfBusinessOnlineTransferID)):          &sourceOfBusinessOnlineTransfer,
		strings.ToLower(string(sourceOfBusinessOtherID)):                   &sourceOfBusinessOther,
		strings.ToLower(string(sourceOfBusinessPolicyPilotID)):             &sourceOfBusinessPolicyPilot,
		strings.ToLower(string(sourceOfBusinessPrintID)):                   &sourceOfBusinessPrint,
		strings.ToLower(string(sourceOfBusinessProsperInsuranceID)):        &sourceOfBusinessProsperInsurance,
		strings.ToLower(string(sourceOfBusinessQuinStreetID)):              &sourceOfBusinessQuinStreet,
		strings.ToLower(string(sourceOfBusinessQuoteLabID)):                &sourceOfBusinessQuoteLab,
		strings.ToLower(string(sourceOfBusinessQuoteWizardID)):             &sourceOfBusinessQuoteWizard,
		strings.ToLower(string(sourceOfBusinessRadioID)):                   &sourceOfBusinessRadio,
		strings.ToLower(string(sourceOfBusinessRateForceID)):               &sourceOfBusinessRateForce,
		strings.ToLower(string(sourceOfBusinessReferralID)):                &sourceOfBusinessReferral,
		strings.ToLower(string(sourceOfBusinessRepeatCustomerID)):          &sourceOfBusinessRepeatCustomer,
		strings.ToLower(string(sourceOfBusinessReviID)):                    &sourceOfBusinessRevi,
		strings.ToLower(string(sourceOfBusinessRocklandID)):                &sourceOfBusinessRockland,
		strings.ToLower(string(sourceOfBusinessSafeAutoID)):                &sourceOfBusinessSafeAuto,
		strings.ToLower(string(sourceOfBusinessSmartFinancialID)):          &sourceOfBusinessSmartFinancial,
		strings.ToLower(string(sourceOfBusinessSuitedConnectorID)):         &sourceOfBusinessSuitedConnector,
		strings.ToLower(string(sourceOfBusinessSuperiorAccessID)):          &sourceOfBusinessSuperiorAccess,
		strings.ToLower(string(sourceOfBusinessTheLeadCoID)):               &sourceOfBusinessTheLeadCo,
		strings.ToLower(string(sourceOfBusinessTVID)):                      &sourceOfBusinessTV,
		strings.ToLower(string(sourceOfBusinessUndergroundElephantID)):     &sourceOfBusinessUndergroundElephant,
		strings.ToLower(string(sourceOfBusinessVantageID)):                 &sourceOfBusinessVantage,
		strings.ToLower(string(sourceOfBusinessWebID)):                     &sourceOfBusinessWeb,
		strings.ToLower(string(sourceOfBusinessZebraID)):                   &sourceOfBusinessZebra,
		strings.ToLower(string(sourceOfBusinessAstoriaID)):                 &sourceOfBusinessAstoria,
		strings.ToLower(string(sourceOfBusinessMediaForceID)):              &sourceOfBusinessMediaForce,
		strings.ToLower(string(sourceOfBusinessAgileRatesID)):              &sourceOfBusinessAgileRates,
		strings.ToLower(string(sourceOfBusinessNextGenLeadsID)):            &sourceOfBusinessNextGenLeads,
		strings.ToLower(string(sourceOfBusinessTransparentlyID)):           &sourceOfBusinessTransparently,
		strings.ToLower(string(sourceOfBusinessAllWebID)):                  &sourceOfBusinessAllWeb,
		strings.ToLower(string(sourceOfBusinessAssuranceID)):               &sourceOfBusinessAssurance,
		strings.ToLower(string(sourceOfBusinessGooseheadID)):               &sourceOfBusinessGoosehead,
		strings.ToLower(string(sourceOfBusinessAgenteroID)):                &sourceOfBusinessAgentero,
		strings.ToLower(string(sourceOfBusinessCoverageComID)):             &sourceOfBusinessCoverageCom,
		strings.ToLower(string(sourceOfBusinessBindrightID)):               &sourceOfBusinessBindright,
		strings.ToLower(string(sourceOfBusinessNationWideID)):              &sourceOfBusinessNationWide,
		strings.ToLower(string(sourceOfBusinessFinderID)):                  &sourceOfBusinessFinder,
		strings.ToLower(string(sourceOfBusinessAppalachianUnderwritersID)): &sourceOfBusinessAppalachianUnderwriters,
		strings.ToLower(string(sourceOfBusinessRenegadeID)):                &sourceOfBusinessRenegade,
		strings.ToLower(string(sourceOfBusinessMaticID)):                   &sourceOfBusinessMatic,
		strings.ToLower(string(sourceOfBusinessMianID)):                    &sourceOfBusinessMian,
		strings.ToLower(string(sourceOfBusinessSmartChoiceID)):             &sourceOfBusinessSmartChoice,
		strings.ToLower(string(sourceOfBusinessWayID)):                     &sourceOfBusinessWay,
		strings.ToLower(string(sourceOfBusinessWoodlandFinancialID)):       &sourceOfBusinessWoodlandFinancial,
		strings.ToLower(string(sourceOfBusinessCreditKarmaID)):             &sourceOfBusinessCreditKarma,
		strings.ToLower(string(sourceOfBusinessAgentsAllianceID)):          &sourceOfBusinessAgentsAlliance,
		strings.ToLower(string(sourceOfBusinessSavvyID)):                   &sourceOfBusinessSavvy,
		strings.ToLower(string(sourceOfBusinessIntegraID)):                 &sourceOfBusinessIntegra,
		strings.ToLower(string(sourceOfBusinessCaribouID)):                 &sourceOfBusinessCaribou,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumSourceOfBusiness) ByID(id SourceOfBusinessIdentifier) *EnumSourceOfBusinessItem {
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
func (e *EnumSourceOfBusiness) ByIDString(idx string) *EnumSourceOfBusinessItem {
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
func (e *EnumSourceOfBusiness) ByIndex(idx int) *EnumSourceOfBusinessItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedSourceOfBusinessID is a struct that is designed to replace a *SourceOfBusinessID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *SourceOfBusinessID it contains while being a better JSON citizen.
type ValidatedSourceOfBusinessID struct {
	// id will point to a valid SourceOfBusinessID, if possible
	// If id is nil, then ValidatedSourceOfBusinessID.Valid() will return false.
	id *SourceOfBusinessID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedSourceOfBusinessID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedSourceOfBusinessID
func (vi *ValidatedSourceOfBusinessID) Clone() *ValidatedSourceOfBusinessID {
	if vi == nil {
		return nil
	}

	var cid *SourceOfBusinessID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedSourceOfBusinessID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedSourceOfBusinessIds represent the same SourceOfBusiness
func (vi *ValidatedSourceOfBusinessID) Equals(vj *ValidatedSourceOfBusinessID) bool {
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

// Valid returns true if and only if the ValidatedSourceOfBusinessID corresponds to a recognized SourceOfBusiness
func (vi *ValidatedSourceOfBusinessID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedSourceOfBusinessID) ID() *SourceOfBusinessID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedSourceOfBusinessID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedSourceOfBusinessID) ValidatedID() *ValidatedSourceOfBusinessID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedSourceOfBusinessID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedSourceOfBusinessID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedSourceOfBusinessID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedSourceOfBusinessID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedSourceOfBusinessID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := SourceOfBusinessID(capString)
	item := SourceOfBusiness.ByID(&id)
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

func (vi ValidatedSourceOfBusinessID) String() string {
	return vi.ToIDString()
}

type SourceOfBusinessIdentifier interface {
	ID() *SourceOfBusinessID
	Valid() bool
}
