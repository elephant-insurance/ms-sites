package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// BusinessTypeID uniquely identifies a particular BusinessType
type BusinessTypeID string

// Clone creates a safe, independent copy of a BusinessTypeID
func (i *BusinessTypeID) Clone() *BusinessTypeID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two BusinessTypeIds are equivalent
func (i *BusinessTypeID) Equals(j *BusinessTypeID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *BusinessTypeID that is either valid or nil
func (i *BusinessTypeID) ID() *BusinessTypeID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *BusinessTypeID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the BusinessTypeID corresponds to a recognized BusinessType
func (i *BusinessTypeID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return BusinessType.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *BusinessTypeID) ValidatedID() *ValidatedBusinessTypeID {
	if i != nil {
		return &ValidatedBusinessTypeID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *BusinessTypeID) MarshalJSON() ([]byte, error) {
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

func (i *BusinessTypeID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := BusinessTypeID(dataString)
	item := BusinessType.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	businessTypeAgencyRateCallID BusinessTypeID = "agencyRateCall"
	businessTypeBrokerClickID    BusinessTypeID = "brokerClick"
	businessTypeEmailRateCallID  BusinessTypeID = "emailRateCall"
	businessTypeInternetID       BusinessTypeID = "internet"
	businessTypeRateCallID       BusinessTypeID = "rateCall"
	businessTypeSharedLeadID     BusinessTypeID = "sharedLead"
	businessTypeSMSRateCallID    BusinessTypeID = "smsRateCall"
	businessTypeSuperClickID     BusinessTypeID = "superClick"
	businessTypeWarmTransferID   BusinessTypeID = "warmTransfer"
	businessTypeWebAgencyID      BusinessTypeID = "webAgency"
	businessTypeADRateCallID     BusinessTypeID = "adRateCall"
)

// EnumBusinessTypeItem describes an entry in an enumeration of BusinessType
type EnumBusinessTypeItem struct {
	ID        BusinessTypeID    `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	businessTypeAgencyRateCall = EnumBusinessTypeItem{businessTypeAgencyRateCallID, "Agency rate call", nil, "AgencyRateCall", 1}
	businessTypeBrokerClick    = EnumBusinessTypeItem{businessTypeBrokerClickID, "Broker click", nil, "BrokerClick", 2}
	businessTypeEmailRateCall  = EnumBusinessTypeItem{businessTypeEmailRateCallID, "Email based rate Call", nil, "EmailRateCall", 3}
	businessTypeInternet       = EnumBusinessTypeItem{businessTypeInternetID, "Internet", nil, "Internet", 4}
	businessTypeRateCall       = EnumBusinessTypeItem{businessTypeRateCallID, "Rate call", nil, "RateCall", 5}
	businessTypeSharedLead     = EnumBusinessTypeItem{businessTypeSharedLeadID, "Shared lead", nil, "SharedLead", 6}
	businessTypeSMSRateCall    = EnumBusinessTypeItem{businessTypeSMSRateCallID, "SMS based rate call", nil, "SMSRateCall", 7}
	businessTypeSuperClick     = EnumBusinessTypeItem{businessTypeSuperClickID, "Super click", nil, "SuperClick", 8}
	businessTypeWarmTransfer   = EnumBusinessTypeItem{businessTypeWarmTransferID, "Warm transfer", nil, "WarmTransfer", 9}
	businessTypeWebAgency      = EnumBusinessTypeItem{businessTypeWebAgencyID, "Web agency", nil, "WebAgency", 10}
	businessTypeADRateCall     = EnumBusinessTypeItem{businessTypeADRateCallID, "AD Rate Call", nil, "ADRateCall", 11}
)

// EnumBusinessType is a collection of BusinessType items
type EnumBusinessType struct {
	Description string
	Items       []*EnumBusinessTypeItem
	Name        string

	AgencyRateCall *EnumBusinessTypeItem
	BrokerClick    *EnumBusinessTypeItem
	EmailRateCall  *EnumBusinessTypeItem
	Internet       *EnumBusinessTypeItem
	RateCall       *EnumBusinessTypeItem
	SharedLead     *EnumBusinessTypeItem
	SMSRateCall    *EnumBusinessTypeItem
	SuperClick     *EnumBusinessTypeItem
	WarmTransfer   *EnumBusinessTypeItem
	WebAgency      *EnumBusinessTypeItem
	ADRateCall     *EnumBusinessTypeItem

	itemDict map[string]*EnumBusinessTypeItem
}

// BusinessType is a public singleton instance of EnumBusinessType
// representing channels of quote business
var BusinessType = &EnumBusinessType{
	Description: "channels of quote business",
	Items: []*EnumBusinessTypeItem{
		&businessTypeAgencyRateCall,
		&businessTypeBrokerClick,
		&businessTypeEmailRateCall,
		&businessTypeInternet,
		&businessTypeRateCall,
		&businessTypeSharedLead,
		&businessTypeSMSRateCall,
		&businessTypeSuperClick,
		&businessTypeWarmTransfer,
		&businessTypeWebAgency,
		&businessTypeADRateCall,
	},
	Name:           "EnumBusinessType",
	AgencyRateCall: &businessTypeAgencyRateCall,
	BrokerClick:    &businessTypeBrokerClick,
	EmailRateCall:  &businessTypeEmailRateCall,
	Internet:       &businessTypeInternet,
	RateCall:       &businessTypeRateCall,
	SharedLead:     &businessTypeSharedLead,
	SMSRateCall:    &businessTypeSMSRateCall,
	SuperClick:     &businessTypeSuperClick,
	WarmTransfer:   &businessTypeWarmTransfer,
	WebAgency:      &businessTypeWebAgency,
	ADRateCall:     &businessTypeADRateCall,

	itemDict: map[string]*EnumBusinessTypeItem{
		strings.ToLower(string(businessTypeAgencyRateCallID)): &businessTypeAgencyRateCall,
		strings.ToLower(string(businessTypeBrokerClickID)):    &businessTypeBrokerClick,
		strings.ToLower(string(businessTypeEmailRateCallID)):  &businessTypeEmailRateCall,
		strings.ToLower(string(businessTypeInternetID)):       &businessTypeInternet,
		strings.ToLower(string(businessTypeRateCallID)):       &businessTypeRateCall,
		strings.ToLower(string(businessTypeSharedLeadID)):     &businessTypeSharedLead,
		strings.ToLower(string(businessTypeSMSRateCallID)):    &businessTypeSMSRateCall,
		strings.ToLower(string(businessTypeSuperClickID)):     &businessTypeSuperClick,
		strings.ToLower(string(businessTypeWarmTransferID)):   &businessTypeWarmTransfer,
		strings.ToLower(string(businessTypeWebAgencyID)):      &businessTypeWebAgency,
		strings.ToLower(string(businessTypeADRateCallID)):     &businessTypeADRateCall,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumBusinessType) ByID(id BusinessTypeIdentifier) *EnumBusinessTypeItem {
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
func (e *EnumBusinessType) ByIDString(idx string) *EnumBusinessTypeItem {
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
func (e *EnumBusinessType) ByIndex(idx int) *EnumBusinessTypeItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedBusinessTypeID is a struct that is designed to replace a *BusinessTypeID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *BusinessTypeID it contains while being a better JSON citizen.
type ValidatedBusinessTypeID struct {
	// id will point to a valid BusinessTypeID, if possible
	// If id is nil, then ValidatedBusinessTypeID.Valid() will return false.
	id *BusinessTypeID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedBusinessTypeID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedBusinessTypeID
func (vi *ValidatedBusinessTypeID) Clone() *ValidatedBusinessTypeID {
	if vi == nil {
		return nil
	}

	var cid *BusinessTypeID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedBusinessTypeID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedBusinessTypeIds represent the same BusinessType
func (vi *ValidatedBusinessTypeID) Equals(vj *ValidatedBusinessTypeID) bool {
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

// Valid returns true if and only if the ValidatedBusinessTypeID corresponds to a recognized BusinessType
func (vi *ValidatedBusinessTypeID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedBusinessTypeID) ID() *BusinessTypeID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedBusinessTypeID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedBusinessTypeID) ValidatedID() *ValidatedBusinessTypeID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedBusinessTypeID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedBusinessTypeID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedBusinessTypeID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedBusinessTypeID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedBusinessTypeID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := BusinessTypeID(capString)
	item := BusinessType.ByID(&id)
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

func (vi ValidatedBusinessTypeID) String() string {
	return vi.ToIDString()
}

type BusinessTypeIdentifier interface {
	ID() *BusinessTypeID
	Valid() bool
}
