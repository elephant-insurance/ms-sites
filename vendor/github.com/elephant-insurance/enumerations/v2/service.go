package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// ServiceID uniquely identifies a particular Service
type ServiceID string

// Clone creates a safe, independent copy of a ServiceID
func (i *ServiceID) Clone() *ServiceID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two ServiceIds are equivalent
func (i *ServiceID) Equals(j *ServiceID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *ServiceID that is either valid or nil
func (i *ServiceID) ID() *ServiceID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

func (i *ServiceID) Parent() *ServiceID {
	if this := Service.ByIDString(string(*i)); this != nil {
		if this.Parent != nil {
			return &this.Parent.ID
		}

		return i
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *ServiceID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the ServiceID corresponds to a recognized Service
func (i *ServiceID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return Service.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *ServiceID) ValidatedID() *ValidatedServiceID {
	if i != nil {
		return &ValidatedServiceID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *ServiceID) MarshalJSON() ([]byte, error) {
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

func (i *ServiceID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := ServiceID(dataString)
	item := Service.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	serviceMsAccountStatusID          ServiceID = "ms-account-status"
	serviceMsAccountVerifyID          ServiceID = "ms-account-verify"
	serviceMsAGProxyID                ServiceID = "ms-ag-proxy"
	serviceMsAgentIdentityID          ServiceID = "ms-agent-identity"
	serviceMsANSRID                   ServiceID = "ms-ansr"
	serviceMsAVSID                    ServiceID = "ms-avs"
	serviceMsBCBlobID                 ServiceID = "ms-bcblob"
	serviceMsBoltQTPushID             ServiceID = "ms-bolt-qt-push"
	serviceMsChannelSourceID          ServiceID = "ms-channel-source"
	serviceMsChatbotID                ServiceID = "ms-chatbot"
	serviceMsCoveragePackageID        ServiceID = "ms-coverage-package"
	serviceMsCoverageVerifyID         ServiceID = "ms-coverage-verify"
	serviceMsCreditScoreFilterID      ServiceID = "ms-credit-score-filter"
	serviceMsCustomerPortalID         ServiceID = "ms-customer-portal"
	serviceMsDataPrefillID            ServiceID = "ms-data-prefill"
	serviceMsDataPrefillWebID         ServiceID = "ms-data-prefill-web"
	serviceMsDHIID                    ServiceID = "ms-dhi"
	serviceMsDNCID                    ServiceID = "ms-dnc"
	serviceMsDocSigningID             ServiceID = "ms-doc-signing"
	serviceMsDriverRecordID           ServiceID = "ms-driver-record"
	serviceMsEnumLookupID             ServiceID = "ms-enum-lookup"
	serviceMsFNOLID                   ServiceID = "ms-fnol"
	serviceMsGeoDataID                ServiceID = "ms-geo-data"
	serviceMsGWBlobID                 ServiceID = "ms-gwblob"
	serviceMsHomesiteID               ServiceID = "ms-homesite"
	serviceMsLogRelayID               ServiceID = "ms-log-relay"
	serviceMsLoginID                  ServiceID = "ms-login"
	serviceMsMessageID                ServiceID = "ms-message"
	serviceMsOktaScimBCID             ServiceID = "ms-okta-scim-bc"
	serviceMsOktaScimCCID             ServiceID = "ms-okta-scim-cc"
	serviceMsOktaScimPCID             ServiceID = "ms-okta-scim-pc"
	serviceMsPaymentID                ServiceID = "ms-payment"
	serviceMsPCBlobID                 ServiceID = "ms-pcblob"
	serviceMsPolkDataID               ServiceID = "ms-polk-data"
	serviceMsQIComID                  ServiceID = "ms-qi-com"
	serviceMsQuoteIntentID            ServiceID = "ms-quote-intent"
	serviceMsQuoteIntentAggregatorsID ServiceID = "ms-quote-intent-aggregators"
	serviceMsQuoteIntentWebID         ServiceID = "ms-quote-intent-web"
	serviceMsRatabaseID               ServiceID = "ms-ratabase"
	serviceMsRCComID                  ServiceID = "ms-rc-com"
	serviceMsRCDBID                   ServiceID = "ms-rc-db"
	serviceMsSemafoneID               ServiceID = "ms-semafone"
	serviceMsShiftID                  ServiceID = "ms-shift"
	serviceMsTemplateID               ServiceID = "ms-template"
	serviceMsTitleRecordID            ServiceID = "ms-title-record"
	serviceMsTransUnionID             ServiceID = "ms-transunion"
	serviceMsVehicleScoreID           ServiceID = "ms-vehicle-score"
	serviceMsVinRatingID              ServiceID = "ms-vin-rating"
	serviceMsRCPosID                  ServiceID = "ms-rcpos"
	serviceMsMessageContactID         ServiceID = "ms-message-contact"
	serviceMsRootAdminID              ServiceID = "ms-root-admin"
	serviceMsEverQuoteBidID           ServiceID = "ms-everquotebid"
	serviceMsMediaAlphaBidID          ServiceID = "ms-mediaalphabid"
	serviceMsQuinStreetBidID          ServiceID = "ms-quinstreetbid"
	serviceMsQuoteWizardBidID         ServiceID = "ms-quotewizardbid"
	serviceMsRTBConfigID              ServiceID = "ms-rtb-config"
	serviceMsSinglesearchID           ServiceID = "ms-singlesearch"
	serviceMsCallQueuesID             ServiceID = "ms-call-queues"
	serviceMsDisplayID                ServiceID = "ms-display"
	serviceMsOperationsID             ServiceID = "ms-operations"
	serviceMsPhoneQueuesID            ServiceID = "ms-phone-queues"
	serviceML_CPVYID                  ServiceID = "ml-cpvy"
	serviceML_VehicleScoreID          ServiceID = "ml-vehicle-score"
	serviceAncillaryApiID             ServiceID = "ancillary-api"
	serviceBillingCenterID            ServiceID = "billing-center"
	serviceClaimsCenterID             ServiceID = "claims-center"
	serviceContactManagerID           ServiceID = "contact-manager"
	serviceExternalApiID              ServiceID = "external-api"
	servicePolicyCenterID             ServiceID = "policy-center"
	serviceRatabaseID                 ServiceID = "ratabase"
	serviceTestID                     ServiceID = "test"
	serviceElvisID                    ServiceID = "elvis"
	serviceAzureBlobStorageID         ServiceID = "az-blob-storage"
	serviceAzureLogAnalyticsID        ServiceID = "az-log-analytics"
	serviceAzureServiceBusQueueID     ServiceID = "az-service-bus-queue"
	serviceAdobeSignID                ServiceID = "adobesign"
	serviceMelissaID                  ServiceID = "melissa"
	serviceDBIPIntegrationID          ServiceID = "db-ip-integration"
	serviceDBRateCallID               ServiceID = "db-rate-call"
	serviceDBExternalLeadsID          ServiceID = "db-external-leads"
	serviceDBVINISOID                 ServiceID = "db-viniso"
	serviceDBQuoteIntentID            ServiceID = "db-quote-intent"
)

// EnumServiceItem describes an entry in an enumeration of Service
type EnumServiceItem struct {
	ID        ServiceID         `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	Parent    *EnumServiceItem  `json:"-"`
	SortOrder int

	// Meta Properties
	LogArea string
}

var (
	serviceMsAccountStatus          = EnumServiceItem{serviceMsAccountStatusID, "ms-account-status", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsAccountStatus", nil, 1, ServiceLogArea.Microservice}
	serviceMsAccountVerify          = EnumServiceItem{serviceMsAccountVerifyID, "ms-account-verify", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsAccountVerify", nil, 2, ServiceLogArea.Microservice}
	serviceMsAGProxy                = EnumServiceItem{serviceMsAGProxyID, "ms-ag-proxy", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsAGProxy", nil, 3, ServiceLogArea.Microservice}
	serviceMsAgentIdentity          = EnumServiceItem{serviceMsAgentIdentityID, "ms-agent-identity", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsAgentIdentity", nil, 4, ServiceLogArea.Microservice}
	serviceMsANSR                   = EnumServiceItem{serviceMsANSRID, "ms-ansr", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsANSR", nil, 5, ServiceLogArea.Microservice}
	serviceMsAVS                    = EnumServiceItem{serviceMsAVSID, "ms-avs", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsAVS", nil, 6, ServiceLogArea.Microservice}
	serviceMsBCBlob                 = EnumServiceItem{serviceMsBCBlobID, "Billing Center instance of ms-gwblob", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsBCBlob", &serviceMsGWBlob, 7, ServiceLogArea.Microservice}
	serviceMsBoltQTPush             = EnumServiceItem{serviceMsBoltQTPushID, "ms-bolt-qt-push", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsBoltQTPush", nil, 8, ServiceLogArea.Microservice}
	serviceMsChannelSource          = EnumServiceItem{serviceMsChannelSourceID, "ms-channel-source", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsChannelSource", nil, 9, ServiceLogArea.Microservice}
	serviceMsChatbot                = EnumServiceItem{serviceMsChatbotID, "ms-chatbot", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsChatbot", nil, 10, ServiceLogArea.Microservice}
	serviceMsCoveragePackage        = EnumServiceItem{serviceMsCoveragePackageID, "ms-coverage-package", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsCoveragePackage", nil, 11, ServiceLogArea.Microservice}
	serviceMsCoverageVerify         = EnumServiceItem{serviceMsCoverageVerifyID, "ms-coverage-verify", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsCoverageVerify", nil, 12, ServiceLogArea.Microservice}
	serviceMsCreditScoreFilter      = EnumServiceItem{serviceMsCreditScoreFilterID, "ms-credit-score-filter", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsCreditScoreFilter", nil, 13, ServiceLogArea.Microservice}
	serviceMsCustomerPortal         = EnumServiceItem{serviceMsCustomerPortalID, "ms-customer-portal", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsCustomerPortal", nil, 14, ServiceLogArea.Microservice}
	serviceMsDataPrefill            = EnumServiceItem{serviceMsDataPrefillID, "ms-data-prefill", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsDataPrefill", nil, 15, ServiceLogArea.Microservice}
	serviceMsDataPrefillWeb         = EnumServiceItem{serviceMsDataPrefillWebID, "ms-data-prefill-web", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsDataPrefillWeb", &serviceMsDataPrefill, 16, ServiceLogArea.Microservice}
	serviceMsDHI                    = EnumServiceItem{serviceMsDHIID, "ms-dhi", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsDHI", nil, 17, ServiceLogArea.Microservice}
	serviceMsDNC                    = EnumServiceItem{serviceMsDNCID, "ms-dnc", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsDNC", nil, 18, ServiceLogArea.Microservice}
	serviceMsDocSigning             = EnumServiceItem{serviceMsDocSigningID, "ms-doc-signing", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsDocSigning", nil, 19, ServiceLogArea.Microservice}
	serviceMsDriverRecord           = EnumServiceItem{serviceMsDriverRecordID, "ms-driver-record", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsDriverRecord", nil, 20, ServiceLogArea.Microservice}
	serviceMsEnumLookup             = EnumServiceItem{serviceMsEnumLookupID, "ms-enum-lookup", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsEnumLookup", nil, 21, ServiceLogArea.Microservice}
	serviceMsFNOL                   = EnumServiceItem{serviceMsFNOLID, "ms-fnol", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsFNOL", nil, 22, ServiceLogArea.Microservice}
	serviceMsGeoData                = EnumServiceItem{serviceMsGeoDataID, "ms-geo-data", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsGeoData", nil, 23, ServiceLogArea.Microservice}
	serviceMsGWBlob                 = EnumServiceItem{serviceMsGWBlobID, "Guidewire Blob Storage service", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsGWBlob", nil, 24, ServiceLogArea.Microservice}
	serviceMsHomesite               = EnumServiceItem{serviceMsHomesiteID, "ms-homesite", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsHomesite", nil, 25, ServiceLogArea.Microservice}
	serviceMsLogRelay               = EnumServiceItem{serviceMsLogRelayID, "ms-log-relay", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsLogRelay", nil, 26, ServiceLogArea.Microservice}
	serviceMsLogin                  = EnumServiceItem{serviceMsLoginID, "ms-login", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsLogin", nil, 27, ServiceLogArea.Microservice}
	serviceMsMessage                = EnumServiceItem{serviceMsMessageID, "ms-message", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsMessage", nil, 28, ServiceLogArea.Microservice}
	serviceMsOktaScimBC             = EnumServiceItem{serviceMsOktaScimBCID, "ms-okta-scim-bc", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsOktaScimBC", nil, 29, ServiceLogArea.Microservice}
	serviceMsOktaScimCC             = EnumServiceItem{serviceMsOktaScimCCID, "ms-okta-scim-cc", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsOktaScimCC", nil, 30, ServiceLogArea.Microservice}
	serviceMsOktaScimPC             = EnumServiceItem{serviceMsOktaScimPCID, "ms-okta-scim-pc", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsOktaScimPC", nil, 31, ServiceLogArea.Microservice}
	serviceMsPayment                = EnumServiceItem{serviceMsPaymentID, "ms-payment", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsPayment", nil, 32, ServiceLogArea.Microservice}
	serviceMsPCBlob                 = EnumServiceItem{serviceMsPCBlobID, "Policy Center instance of ms-gwblob service", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsPCBlob", &serviceMsGWBlob, 33, ServiceLogArea.Microservice}
	serviceMsPolkData               = EnumServiceItem{serviceMsPolkDataID, "ms-polk-data", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsPolkData", nil, 34, ServiceLogArea.Microservice}
	serviceMsQICom                  = EnumServiceItem{serviceMsQIComID, "ms-qi-com", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsQICom", nil, 35, ServiceLogArea.Microservice}
	serviceMsQuoteIntent            = EnumServiceItem{serviceMsQuoteIntentID, "The ms-quote-intent microservice", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsQuoteIntent", nil, 36, ServiceLogArea.Microservice}
	serviceMsQuoteIntentAggregators = EnumServiceItem{serviceMsQuoteIntentAggregatorsID, "Aggregators instance of ms-quote-intent", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsQuoteIntentAggregators", &serviceMsQuoteIntent, 37, ServiceLogArea.Microservice}
	serviceMsQuoteIntentWeb         = EnumServiceItem{serviceMsQuoteIntentWebID, "Web instance of ms-quote-intent", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsQuoteIntentWeb", &serviceMsQuoteIntent, 38, ServiceLogArea.Microservice}
	serviceMsRatabase               = EnumServiceItem{serviceMsRatabaseID, "ms-ratabase", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsRatabase", nil, 39, ServiceLogArea.Microservice}
	serviceMsRCCom                  = EnumServiceItem{serviceMsRCComID, "ms-rc-com", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsRCCom", nil, 40, ServiceLogArea.Microservice}
	serviceMsRCDB                   = EnumServiceItem{serviceMsRCDBID, "ms-rc-db", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsRCDB", nil, 41, ServiceLogArea.Microservice}
	serviceMsSemafone               = EnumServiceItem{serviceMsSemafoneID, "ms-semafone", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsSemafone", nil, 42, ServiceLogArea.Microservice}
	serviceMsShift                  = EnumServiceItem{serviceMsShiftID, "ms-shift", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsShift", nil, 43, ServiceLogArea.Microservice}
	serviceMsTemplate               = EnumServiceItem{serviceMsTemplateID, "Elephant microservice template", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsTemplate", nil, 44, ServiceLogArea.Microservice}
	serviceMsTitleRecord            = EnumServiceItem{serviceMsTitleRecordID, "ms-title-record", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsTitleRecord", nil, 45, ServiceLogArea.Microservice}
	serviceMsTransUnion             = EnumServiceItem{serviceMsTransUnionID, "ms-transunion", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsTransUnion", nil, 46, ServiceLogArea.Microservice}
	serviceMsVehicleScore           = EnumServiceItem{serviceMsVehicleScoreID, "Vehicle Score microservice", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsVehicleScore", nil, 47, ServiceLogArea.Microservice}
	serviceMsVinRating              = EnumServiceItem{serviceMsVinRatingID, "ms-vin-rating", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsVinRating", nil, 48, ServiceLogArea.Microservice}
	serviceMsRCPos                  = EnumServiceItem{serviceMsRCPosID, "ms-rcpos", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsRCPos", nil, 49, ServiceLogArea.Microservice}
	serviceMsMessageContact         = EnumServiceItem{serviceMsMessageContactID, "ms-message-contact", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsMessageContact", nil, 50, ServiceLogArea.Microservice}
	serviceMsRootAdmin              = EnumServiceItem{serviceMsRootAdminID, "ms-root-admin", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Microservice}, "MsRootAdmin", nil, 51, ServiceLogArea.Microservice}
	serviceMsEverQuoteBid           = EnumServiceItem{serviceMsEverQuoteBidID, "EverQuote RTB service", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.RealTimeBidding}, "MsEverQuoteBid", nil, 52, ServiceLogArea.RealTimeBidding}
	serviceMsMediaAlphaBid          = EnumServiceItem{serviceMsMediaAlphaBidID, "MediaAlpha RTB service", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.RealTimeBidding}, "MsMediaAlphaBid", nil, 53, ServiceLogArea.RealTimeBidding}
	serviceMsQuinStreetBid          = EnumServiceItem{serviceMsQuinStreetBidID, "QuinStreet RTB service", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.RealTimeBidding}, "MsQuinStreetBid", nil, 54, ServiceLogArea.RealTimeBidding}
	serviceMsQuoteWizardBid         = EnumServiceItem{serviceMsQuoteWizardBidID, "QuoteWizard RTB service", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.RealTimeBidding}, "MsQuoteWizardBid", nil, 55, ServiceLogArea.RealTimeBidding}
	serviceMsRTBConfig              = EnumServiceItem{serviceMsRTBConfigID, "RTB configuration service", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.RealTimeBidding}, "MsRTBConfig", nil, 56, ServiceLogArea.RealTimeBidding}
	serviceMsSinglesearch           = EnumServiceItem{serviceMsSinglesearchID, "ms-singlesearch", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.SingleSearch}, "MsSinglesearch", nil, 57, ServiceLogArea.SingleSearch}
	serviceMsCallQueues             = EnumServiceItem{serviceMsCallQueuesID, "ms-call-queues", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Wallboard}, "MsCallQueues", nil, 58, ServiceLogArea.Wallboard}
	serviceMsDisplay                = EnumServiceItem{serviceMsDisplayID, "ms-display", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Wallboard}, "MsDisplay", nil, 59, ServiceLogArea.Wallboard}
	serviceMsOperations             = EnumServiceItem{serviceMsOperationsID, "ms-operations", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Wallboard}, "MsOperations", nil, 60, ServiceLogArea.Wallboard}
	serviceMsPhoneQueues            = EnumServiceItem{serviceMsPhoneQueuesID, "ms-phone-queues", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.Wallboard}, "MsPhoneQueues", nil, 61, ServiceLogArea.Wallboard}
	serviceML_CPVY                  = EnumServiceItem{serviceML_CPVYID, "Machine Learning: CPVY", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "ML_CPVY", nil, 62, ServiceLogArea.None}
	serviceML_VehicleScore          = EnumServiceItem{serviceML_VehicleScoreID, "Machine Learning: Vehicle Score", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "ML_VehicleScore", nil, 63, ServiceLogArea.None}
	serviceAncillaryApi             = EnumServiceItem{serviceAncillaryApiID, "Ancillary API", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "AncillaryApi", nil, 64, ServiceLogArea.None}
	serviceBillingCenter            = EnumServiceItem{serviceBillingCenterID, "Billing Center", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "BillingCenter", nil, 65, ServiceLogArea.None}
	serviceClaimsCenter             = EnumServiceItem{serviceClaimsCenterID, "Claims Center", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "ClaimsCenter", nil, 66, ServiceLogArea.None}
	serviceContactManager           = EnumServiceItem{serviceContactManagerID, "Contact Manager", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "ContactManager", nil, 67, ServiceLogArea.None}
	serviceExternalApi              = EnumServiceItem{serviceExternalApiID, "External API", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "ExternalApi", nil, 68, ServiceLogArea.None}
	servicePolicyCenter             = EnumServiceItem{servicePolicyCenterID, "Policy Center", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "PolicyCenter", nil, 69, ServiceLogArea.None}
	serviceRatabase                 = EnumServiceItem{serviceRatabaseID, "Ratabase", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "Ratabase", nil, 70, ServiceLogArea.None}
	serviceTest                     = EnumServiceItem{serviceTestID, "A service ID for testing", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "Test", nil, 71, ServiceLogArea.None}
	serviceElvis                    = EnumServiceItem{serviceElvisID, "ELVIS Database Manager", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "Elvis", nil, 72, ServiceLogArea.None}
	serviceAzureBlobStorage         = EnumServiceItem{serviceAzureBlobStorageID, "Azure Blob Storage", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "AzureBlobStorage", nil, 73, ServiceLogArea.None}
	serviceAzureLogAnalytics        = EnumServiceItem{serviceAzureLogAnalyticsID, "Azure Log Analytics", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "AzureLogAnalytics", nil, 74, ServiceLogArea.None}
	serviceAzureServiceBusQueue     = EnumServiceItem{serviceAzureServiceBusQueueID, "Azure Service Bus Queue", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "AzureServiceBusQueue", nil, 75, ServiceLogArea.None}
	serviceAdobeSign                = EnumServiceItem{serviceAdobeSignID, "AdobeSign", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "AdobeSign", nil, 76, ServiceLogArea.None}
	serviceMelissa                  = EnumServiceItem{serviceMelissaID, "Melissa Data Service", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "Melissa", nil, 77, ServiceLogArea.None}
	serviceDBIPIntegration          = EnumServiceItem{serviceDBIPIntegrationID, "IP Integration Database", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "DBIPIntegration", nil, 78, ServiceLogArea.None}
	serviceDBRateCall               = EnumServiceItem{serviceDBRateCallID, "Ratecall Database", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "DBRateCall", nil, 79, ServiceLogArea.None}
	serviceDBExternalLeads          = EnumServiceItem{serviceDBExternalLeadsID, "External Leads Database", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "DBExternalLeads", nil, 80, ServiceLogArea.None}
	serviceDBVINISO                 = EnumServiceItem{serviceDBVINISOID, "VINISO Database", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "DBVINISO", nil, 81, ServiceLogArea.None}
	serviceDBQuoteIntent            = EnumServiceItem{serviceDBQuoteIntentID, "Quote Intent Database", map[string]string{ServiceMetaLogAreaKey: ServiceLogArea.None}, "DBQuoteIntent", nil, 82, ServiceLogArea.None}
)

// EnumService is a collection of Service items
type EnumService struct {
	Description string
	Items       []*EnumServiceItem
	Name        string

	MsAccountStatus          *EnumServiceItem
	MsAccountVerify          *EnumServiceItem
	MsAGProxy                *EnumServiceItem
	MsAgentIdentity          *EnumServiceItem
	MsANSR                   *EnumServiceItem
	MsAVS                    *EnumServiceItem
	MsBCBlob                 *EnumServiceItem
	MsBoltQTPush             *EnumServiceItem
	MsChannelSource          *EnumServiceItem
	MsChatbot                *EnumServiceItem
	MsCoveragePackage        *EnumServiceItem
	MsCoverageVerify         *EnumServiceItem
	MsCreditScoreFilter      *EnumServiceItem
	MsCustomerPortal         *EnumServiceItem
	MsDataPrefill            *EnumServiceItem
	MsDataPrefillWeb         *EnumServiceItem
	MsDHI                    *EnumServiceItem
	MsDNC                    *EnumServiceItem
	MsDocSigning             *EnumServiceItem
	MsDriverRecord           *EnumServiceItem
	MsEnumLookup             *EnumServiceItem
	MsFNOL                   *EnumServiceItem
	MsGeoData                *EnumServiceItem
	MsGWBlob                 *EnumServiceItem
	MsHomesite               *EnumServiceItem
	MsLogRelay               *EnumServiceItem
	MsLogin                  *EnumServiceItem
	MsMessage                *EnumServiceItem
	MsOktaScimBC             *EnumServiceItem
	MsOktaScimCC             *EnumServiceItem
	MsOktaScimPC             *EnumServiceItem
	MsPayment                *EnumServiceItem
	MsPCBlob                 *EnumServiceItem
	MsPolkData               *EnumServiceItem
	MsQICom                  *EnumServiceItem
	MsQuoteIntent            *EnumServiceItem
	MsQuoteIntentAggregators *EnumServiceItem
	MsQuoteIntentWeb         *EnumServiceItem
	MsRatabase               *EnumServiceItem
	MsRCCom                  *EnumServiceItem
	MsRCDB                   *EnumServiceItem
	MsSemafone               *EnumServiceItem
	MsShift                  *EnumServiceItem
	MsTemplate               *EnumServiceItem
	MsTitleRecord            *EnumServiceItem
	MsTransUnion             *EnumServiceItem
	MsVehicleScore           *EnumServiceItem
	MsVinRating              *EnumServiceItem
	MsRCPos                  *EnumServiceItem
	MsMessageContact         *EnumServiceItem
	MsRootAdmin              *EnumServiceItem
	MsEverQuoteBid           *EnumServiceItem
	MsMediaAlphaBid          *EnumServiceItem
	MsQuinStreetBid          *EnumServiceItem
	MsQuoteWizardBid         *EnumServiceItem
	MsRTBConfig              *EnumServiceItem
	MsSinglesearch           *EnumServiceItem
	MsCallQueues             *EnumServiceItem
	MsDisplay                *EnumServiceItem
	MsOperations             *EnumServiceItem
	MsPhoneQueues            *EnumServiceItem
	ML_CPVY                  *EnumServiceItem
	ML_VehicleScore          *EnumServiceItem
	AncillaryApi             *EnumServiceItem
	BillingCenter            *EnumServiceItem
	ClaimsCenter             *EnumServiceItem
	ContactManager           *EnumServiceItem
	ExternalApi              *EnumServiceItem
	PolicyCenter             *EnumServiceItem
	Ratabase                 *EnumServiceItem
	Test                     *EnumServiceItem
	Elvis                    *EnumServiceItem
	AzureBlobStorage         *EnumServiceItem
	AzureLogAnalytics        *EnumServiceItem
	AzureServiceBusQueue     *EnumServiceItem
	AdobeSign                *EnumServiceItem
	Melissa                  *EnumServiceItem
	DBIPIntegration          *EnumServiceItem
	DBRateCall               *EnumServiceItem
	DBExternalLeads          *EnumServiceItem
	DBVINISO                 *EnumServiceItem
	DBQuoteIntent            *EnumServiceItem

	itemDict map[string]*EnumServiceItem
}

// Service is a public singleton instance of EnumService
// representing all services called by our microservices
var Service = &EnumService{
	Description: "all services called by our microservices",
	Items: []*EnumServiceItem{
		&serviceMsAccountStatus,
		&serviceMsAccountVerify,
		&serviceMsAGProxy,
		&serviceMsAgentIdentity,
		&serviceMsANSR,
		&serviceMsAVS,
		&serviceMsBCBlob,
		&serviceMsBoltQTPush,
		&serviceMsChannelSource,
		&serviceMsChatbot,
		&serviceMsCoveragePackage,
		&serviceMsCoverageVerify,
		&serviceMsCreditScoreFilter,
		&serviceMsCustomerPortal,
		&serviceMsDataPrefill,
		&serviceMsDataPrefillWeb,
		&serviceMsDHI,
		&serviceMsDNC,
		&serviceMsDocSigning,
		&serviceMsDriverRecord,
		&serviceMsEnumLookup,
		&serviceMsFNOL,
		&serviceMsGeoData,
		&serviceMsGWBlob,
		&serviceMsHomesite,
		&serviceMsLogRelay,
		&serviceMsLogin,
		&serviceMsMessage,
		&serviceMsOktaScimBC,
		&serviceMsOktaScimCC,
		&serviceMsOktaScimPC,
		&serviceMsPayment,
		&serviceMsPCBlob,
		&serviceMsPolkData,
		&serviceMsQICom,
		&serviceMsQuoteIntent,
		&serviceMsQuoteIntentAggregators,
		&serviceMsQuoteIntentWeb,
		&serviceMsRatabase,
		&serviceMsRCCom,
		&serviceMsRCDB,
		&serviceMsSemafone,
		&serviceMsShift,
		&serviceMsTemplate,
		&serviceMsTitleRecord,
		&serviceMsTransUnion,
		&serviceMsVehicleScore,
		&serviceMsVinRating,
		&serviceMsRCPos,
		&serviceMsMessageContact,
		&serviceMsRootAdmin,
		&serviceMsEverQuoteBid,
		&serviceMsMediaAlphaBid,
		&serviceMsQuinStreetBid,
		&serviceMsQuoteWizardBid,
		&serviceMsRTBConfig,
		&serviceMsSinglesearch,
		&serviceMsCallQueues,
		&serviceMsDisplay,
		&serviceMsOperations,
		&serviceMsPhoneQueues,
		&serviceML_CPVY,
		&serviceML_VehicleScore,
		&serviceAncillaryApi,
		&serviceBillingCenter,
		&serviceClaimsCenter,
		&serviceContactManager,
		&serviceExternalApi,
		&servicePolicyCenter,
		&serviceRatabase,
		&serviceTest,
		&serviceElvis,
		&serviceAzureBlobStorage,
		&serviceAzureLogAnalytics,
		&serviceAzureServiceBusQueue,
		&serviceAdobeSign,
		&serviceMelissa,
		&serviceDBIPIntegration,
		&serviceDBRateCall,
		&serviceDBExternalLeads,
		&serviceDBVINISO,
		&serviceDBQuoteIntent,
	},
	Name:                     "EnumService",
	MsAccountStatus:          &serviceMsAccountStatus,
	MsAccountVerify:          &serviceMsAccountVerify,
	MsAGProxy:                &serviceMsAGProxy,
	MsAgentIdentity:          &serviceMsAgentIdentity,
	MsANSR:                   &serviceMsANSR,
	MsAVS:                    &serviceMsAVS,
	MsBCBlob:                 &serviceMsBCBlob,
	MsBoltQTPush:             &serviceMsBoltQTPush,
	MsChannelSource:          &serviceMsChannelSource,
	MsChatbot:                &serviceMsChatbot,
	MsCoveragePackage:        &serviceMsCoveragePackage,
	MsCoverageVerify:         &serviceMsCoverageVerify,
	MsCreditScoreFilter:      &serviceMsCreditScoreFilter,
	MsCustomerPortal:         &serviceMsCustomerPortal,
	MsDataPrefill:            &serviceMsDataPrefill,
	MsDataPrefillWeb:         &serviceMsDataPrefillWeb,
	MsDHI:                    &serviceMsDHI,
	MsDNC:                    &serviceMsDNC,
	MsDocSigning:             &serviceMsDocSigning,
	MsDriverRecord:           &serviceMsDriverRecord,
	MsEnumLookup:             &serviceMsEnumLookup,
	MsFNOL:                   &serviceMsFNOL,
	MsGeoData:                &serviceMsGeoData,
	MsGWBlob:                 &serviceMsGWBlob,
	MsHomesite:               &serviceMsHomesite,
	MsLogRelay:               &serviceMsLogRelay,
	MsLogin:                  &serviceMsLogin,
	MsMessage:                &serviceMsMessage,
	MsOktaScimBC:             &serviceMsOktaScimBC,
	MsOktaScimCC:             &serviceMsOktaScimCC,
	MsOktaScimPC:             &serviceMsOktaScimPC,
	MsPayment:                &serviceMsPayment,
	MsPCBlob:                 &serviceMsPCBlob,
	MsPolkData:               &serviceMsPolkData,
	MsQICom:                  &serviceMsQICom,
	MsQuoteIntent:            &serviceMsQuoteIntent,
	MsQuoteIntentAggregators: &serviceMsQuoteIntentAggregators,
	MsQuoteIntentWeb:         &serviceMsQuoteIntentWeb,
	MsRatabase:               &serviceMsRatabase,
	MsRCCom:                  &serviceMsRCCom,
	MsRCDB:                   &serviceMsRCDB,
	MsSemafone:               &serviceMsSemafone,
	MsShift:                  &serviceMsShift,
	MsTemplate:               &serviceMsTemplate,
	MsTitleRecord:            &serviceMsTitleRecord,
	MsTransUnion:             &serviceMsTransUnion,
	MsVehicleScore:           &serviceMsVehicleScore,
	MsVinRating:              &serviceMsVinRating,
	MsRCPos:                  &serviceMsRCPos,
	MsMessageContact:         &serviceMsMessageContact,
	MsRootAdmin:              &serviceMsRootAdmin,
	MsEverQuoteBid:           &serviceMsEverQuoteBid,
	MsMediaAlphaBid:          &serviceMsMediaAlphaBid,
	MsQuinStreetBid:          &serviceMsQuinStreetBid,
	MsQuoteWizardBid:         &serviceMsQuoteWizardBid,
	MsRTBConfig:              &serviceMsRTBConfig,
	MsSinglesearch:           &serviceMsSinglesearch,
	MsCallQueues:             &serviceMsCallQueues,
	MsDisplay:                &serviceMsDisplay,
	MsOperations:             &serviceMsOperations,
	MsPhoneQueues:            &serviceMsPhoneQueues,
	ML_CPVY:                  &serviceML_CPVY,
	ML_VehicleScore:          &serviceML_VehicleScore,
	AncillaryApi:             &serviceAncillaryApi,
	BillingCenter:            &serviceBillingCenter,
	ClaimsCenter:             &serviceClaimsCenter,
	ContactManager:           &serviceContactManager,
	ExternalApi:              &serviceExternalApi,
	PolicyCenter:             &servicePolicyCenter,
	Ratabase:                 &serviceRatabase,
	Test:                     &serviceTest,
	Elvis:                    &serviceElvis,
	AzureBlobStorage:         &serviceAzureBlobStorage,
	AzureLogAnalytics:        &serviceAzureLogAnalytics,
	AzureServiceBusQueue:     &serviceAzureServiceBusQueue,
	AdobeSign:                &serviceAdobeSign,
	Melissa:                  &serviceMelissa,
	DBIPIntegration:          &serviceDBIPIntegration,
	DBRateCall:               &serviceDBRateCall,
	DBExternalLeads:          &serviceDBExternalLeads,
	DBVINISO:                 &serviceDBVINISO,
	DBQuoteIntent:            &serviceDBQuoteIntent,

	itemDict: map[string]*EnumServiceItem{
		strings.ToLower(string(serviceMsAccountStatusID)):          &serviceMsAccountStatus,
		strings.ToLower(string(serviceMsAccountVerifyID)):          &serviceMsAccountVerify,
		strings.ToLower(string(serviceMsAGProxyID)):                &serviceMsAGProxy,
		strings.ToLower(string(serviceMsAgentIdentityID)):          &serviceMsAgentIdentity,
		strings.ToLower(string(serviceMsANSRID)):                   &serviceMsANSR,
		strings.ToLower(string(serviceMsAVSID)):                    &serviceMsAVS,
		strings.ToLower(string(serviceMsBCBlobID)):                 &serviceMsBCBlob,
		strings.ToLower(string(serviceMsBoltQTPushID)):             &serviceMsBoltQTPush,
		strings.ToLower(string(serviceMsChannelSourceID)):          &serviceMsChannelSource,
		strings.ToLower(string(serviceMsChatbotID)):                &serviceMsChatbot,
		strings.ToLower(string(serviceMsCoveragePackageID)):        &serviceMsCoveragePackage,
		strings.ToLower(string(serviceMsCoverageVerifyID)):         &serviceMsCoverageVerify,
		strings.ToLower(string(serviceMsCreditScoreFilterID)):      &serviceMsCreditScoreFilter,
		strings.ToLower(string(serviceMsCustomerPortalID)):         &serviceMsCustomerPortal,
		strings.ToLower(string(serviceMsDataPrefillID)):            &serviceMsDataPrefill,
		strings.ToLower(string(serviceMsDataPrefillWebID)):         &serviceMsDataPrefillWeb,
		strings.ToLower(string(serviceMsDHIID)):                    &serviceMsDHI,
		strings.ToLower(string(serviceMsDNCID)):                    &serviceMsDNC,
		strings.ToLower(string(serviceMsDocSigningID)):             &serviceMsDocSigning,
		strings.ToLower(string(serviceMsDriverRecordID)):           &serviceMsDriverRecord,
		strings.ToLower(string(serviceMsEnumLookupID)):             &serviceMsEnumLookup,
		strings.ToLower(string(serviceMsFNOLID)):                   &serviceMsFNOL,
		strings.ToLower(string(serviceMsGeoDataID)):                &serviceMsGeoData,
		strings.ToLower(string(serviceMsGWBlobID)):                 &serviceMsGWBlob,
		strings.ToLower(string(serviceMsHomesiteID)):               &serviceMsHomesite,
		strings.ToLower(string(serviceMsLogRelayID)):               &serviceMsLogRelay,
		strings.ToLower(string(serviceMsLoginID)):                  &serviceMsLogin,
		strings.ToLower(string(serviceMsMessageID)):                &serviceMsMessage,
		strings.ToLower(string(serviceMsOktaScimBCID)):             &serviceMsOktaScimBC,
		strings.ToLower(string(serviceMsOktaScimCCID)):             &serviceMsOktaScimCC,
		strings.ToLower(string(serviceMsOktaScimPCID)):             &serviceMsOktaScimPC,
		strings.ToLower(string(serviceMsPaymentID)):                &serviceMsPayment,
		strings.ToLower(string(serviceMsPCBlobID)):                 &serviceMsPCBlob,
		strings.ToLower(string(serviceMsPolkDataID)):               &serviceMsPolkData,
		strings.ToLower(string(serviceMsQIComID)):                  &serviceMsQICom,
		strings.ToLower(string(serviceMsQuoteIntentID)):            &serviceMsQuoteIntent,
		strings.ToLower(string(serviceMsQuoteIntentAggregatorsID)): &serviceMsQuoteIntentAggregators,
		strings.ToLower(string(serviceMsQuoteIntentWebID)):         &serviceMsQuoteIntentWeb,
		strings.ToLower(string(serviceMsRatabaseID)):               &serviceMsRatabase,
		strings.ToLower(string(serviceMsRCComID)):                  &serviceMsRCCom,
		strings.ToLower(string(serviceMsRCDBID)):                   &serviceMsRCDB,
		strings.ToLower(string(serviceMsSemafoneID)):               &serviceMsSemafone,
		strings.ToLower(string(serviceMsShiftID)):                  &serviceMsShift,
		strings.ToLower(string(serviceMsTemplateID)):               &serviceMsTemplate,
		strings.ToLower(string(serviceMsTitleRecordID)):            &serviceMsTitleRecord,
		strings.ToLower(string(serviceMsTransUnionID)):             &serviceMsTransUnion,
		strings.ToLower(string(serviceMsVehicleScoreID)):           &serviceMsVehicleScore,
		strings.ToLower(string(serviceMsVinRatingID)):              &serviceMsVinRating,
		strings.ToLower(string(serviceMsRCPosID)):                  &serviceMsRCPos,
		strings.ToLower(string(serviceMsMessageContactID)):         &serviceMsMessageContact,
		strings.ToLower(string(serviceMsRootAdminID)):              &serviceMsRootAdmin,
		strings.ToLower(string(serviceMsEverQuoteBidID)):           &serviceMsEverQuoteBid,
		strings.ToLower(string(serviceMsMediaAlphaBidID)):          &serviceMsMediaAlphaBid,
		strings.ToLower(string(serviceMsQuinStreetBidID)):          &serviceMsQuinStreetBid,
		strings.ToLower(string(serviceMsQuoteWizardBidID)):         &serviceMsQuoteWizardBid,
		strings.ToLower(string(serviceMsRTBConfigID)):              &serviceMsRTBConfig,
		strings.ToLower(string(serviceMsSinglesearchID)):           &serviceMsSinglesearch,
		strings.ToLower(string(serviceMsCallQueuesID)):             &serviceMsCallQueues,
		strings.ToLower(string(serviceMsDisplayID)):                &serviceMsDisplay,
		strings.ToLower(string(serviceMsOperationsID)):             &serviceMsOperations,
		strings.ToLower(string(serviceMsPhoneQueuesID)):            &serviceMsPhoneQueues,
		strings.ToLower(string(serviceML_CPVYID)):                  &serviceML_CPVY,
		strings.ToLower(string(serviceML_VehicleScoreID)):          &serviceML_VehicleScore,
		strings.ToLower(string(serviceAncillaryApiID)):             &serviceAncillaryApi,
		strings.ToLower(string(serviceBillingCenterID)):            &serviceBillingCenter,
		strings.ToLower(string(serviceClaimsCenterID)):             &serviceClaimsCenter,
		strings.ToLower(string(serviceContactManagerID)):           &serviceContactManager,
		strings.ToLower(string(serviceExternalApiID)):              &serviceExternalApi,
		strings.ToLower(string(servicePolicyCenterID)):             &servicePolicyCenter,
		strings.ToLower(string(serviceRatabaseID)):                 &serviceRatabase,
		strings.ToLower(string(serviceTestID)):                     &serviceTest,
		strings.ToLower(string(serviceElvisID)):                    &serviceElvis,
		strings.ToLower(string(serviceAzureBlobStorageID)):         &serviceAzureBlobStorage,
		strings.ToLower(string(serviceAzureLogAnalyticsID)):        &serviceAzureLogAnalytics,
		strings.ToLower(string(serviceAzureServiceBusQueueID)):     &serviceAzureServiceBusQueue,
		strings.ToLower(string(serviceAdobeSignID)):                &serviceAdobeSign,
		strings.ToLower(string(serviceMelissaID)):                  &serviceMelissa,
		strings.ToLower(string(serviceDBIPIntegrationID)):          &serviceDBIPIntegration,
		strings.ToLower(string(serviceDBRateCallID)):               &serviceDBRateCall,
		strings.ToLower(string(serviceDBExternalLeadsID)):          &serviceDBExternalLeads,
		strings.ToLower(string(serviceDBVINISOID)):                 &serviceDBVINISO,
		strings.ToLower(string(serviceDBQuoteIntentID)):            &serviceDBQuoteIntent,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumService) ByID(id ServiceIdentifier) *EnumServiceItem {
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
func (e *EnumService) ByIDString(idx string) *EnumServiceItem {
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
func (e *EnumService) ByIndex(idx int) *EnumServiceItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedServiceID is a struct that is designed to replace a *ServiceID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *ServiceID it contains while being a better JSON citizen.
type ValidatedServiceID struct {
	// id will point to a valid ServiceID, if possible
	// If id is nil, then ValidatedServiceID.Valid() will return false.
	id *ServiceID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedServiceID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedServiceID
func (vi *ValidatedServiceID) Clone() *ValidatedServiceID {
	if vi == nil {
		return nil
	}

	var cid *ServiceID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedServiceID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedServiceIds represent the same Service
func (vi *ValidatedServiceID) Equals(vj *ValidatedServiceID) bool {
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

func (vi *ValidatedServiceID) Parent() *ValidatedServiceID {
	if vi == nil || vi.id == nil {
		pid := vi.id.Parent()
		if pid != nil {
			return pid.ValidatedID()
		}
	}

	return nil
}

// Valid returns true if and only if the ValidatedServiceID corresponds to a recognized Service
func (vi *ValidatedServiceID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedServiceID) ID() *ServiceID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedServiceID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedServiceID) ValidatedID() *ValidatedServiceID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedServiceID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedServiceID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedServiceID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedServiceID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedServiceID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := ServiceID(capString)
	item := Service.ByID(&id)
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

func (vi ValidatedServiceID) String() string {
	return vi.ToIDString()
}

type ServiceIdentifier interface {
	ID() *ServiceID
	Valid() bool
}
