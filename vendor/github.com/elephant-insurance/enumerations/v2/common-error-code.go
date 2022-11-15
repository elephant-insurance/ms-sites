package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// CommonErrorCodeID uniquely identifies a particular CommonErrorCode
type CommonErrorCodeID string

// Clone creates a safe, independent copy of a CommonErrorCodeID
func (i *CommonErrorCodeID) Clone() *CommonErrorCodeID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two CommonErrorCodeIds are equivalent
func (i *CommonErrorCodeID) Equals(j *CommonErrorCodeID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *CommonErrorCodeID that is either valid or nil
func (i *CommonErrorCodeID) ID() *CommonErrorCodeID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *CommonErrorCodeID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the CommonErrorCodeID corresponds to a recognized CommonErrorCode
func (i *CommonErrorCodeID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return CommonErrorCode.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *CommonErrorCodeID) ValidatedID() *ValidatedCommonErrorCodeID {
	if i != nil {
		return &ValidatedCommonErrorCodeID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *CommonErrorCodeID) MarshalJSON() ([]byte, error) {
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

func (i *CommonErrorCodeID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := CommonErrorCodeID(dataString)
	item := CommonErrorCode.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	commonErrorCodeDBReadID                        CommonErrorCodeID = "DBRead"
	commonErrorCodeDBUnknownID                     CommonErrorCodeID = "DBUnknown"
	commonErrorCodeDBWriteID                       CommonErrorCodeID = "DBWrite"
	commonErrorCodeHTTPRequestID                   CommonErrorCodeID = "HTTPRequest"
	commonErrorCodeHTTPStatusID                    CommonErrorCodeID = "HTTPStatus"
	commonErrorCodeInvalidValueID                  CommonErrorCodeID = "InvalidValue"
	commonErrorCodeMarshalJSONID                   CommonErrorCodeID = "MarshalJSON"
	commonErrorCodeMarshalXMLID                    CommonErrorCodeID = "MarshalXML"
	commonErrorCodeNoInstallmentPreferredPayPlanID CommonErrorCodeID = "NoInstallmentPreferredPayPlan"
	commonErrorCodeNoPIFPayPlanErrorID             CommonErrorCodeID = "NoPIFPayPlanError"
	commonErrorCodeObjectCantBeBlankID             CommonErrorCodeID = "ObjectCantBeBlank"
	commonErrorCodePayPlansEmptyID                 CommonErrorCodeID = "PayPlansEmpty"
	commonErrorCodeRuleEngineErrorID               CommonErrorCodeID = "RuleEngineError"
	commonErrorCodeUnauthorizedID                  CommonErrorCodeID = "Unauthorized"
	commonErrorCodeUnmarshalJSONID                 CommonErrorCodeID = "UnmarshalJSON"
	commonErrorCodeUnmarshalXMLID                  CommonErrorCodeID = "UnmarshalXML"
	commonErrorCodeValueRequiredID                 CommonErrorCodeID = "ValueRequired"
)

// EnumCommonErrorCodeItem describes an entry in an enumeration of CommonErrorCode
type EnumCommonErrorCodeItem struct {
	ID        CommonErrorCodeID `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	commonErrorCodeDBRead                        = EnumCommonErrorCodeItem{commonErrorCodeDBReadID, "data read error from database", nil, "DBRead", 1}
	commonErrorCodeDBUnknown                     = EnumCommonErrorCodeItem{commonErrorCodeDBUnknownID, "unknown database error", nil, "DBUnknown", 2}
	commonErrorCodeDBWrite                       = EnumCommonErrorCodeItem{commonErrorCodeDBWriteID, "data write error from database", nil, "DBWrite", 3}
	commonErrorCodeHTTPRequest                   = EnumCommonErrorCodeItem{commonErrorCodeHTTPRequestID, "call to service failed during request", nil, "HTTPRequest", 4}
	commonErrorCodeHTTPStatus                    = EnumCommonErrorCodeItem{commonErrorCodeHTTPStatusID, "call to service returned non-success HTTP code", nil, "HTTPStatus", 5}
	commonErrorCodeInvalidValue                  = EnumCommonErrorCodeItem{commonErrorCodeInvalidValueID, "invalid field value", nil, "InvalidValue", 6}
	commonErrorCodeMarshalJSON                   = EnumCommonErrorCodeItem{commonErrorCodeMarshalJSONID, "object provided isn't in proper JSON or has some invalid values", nil, "MarshalJSON", 7}
	commonErrorCodeMarshalXML                    = EnumCommonErrorCodeItem{commonErrorCodeMarshalXMLID, "object provided isn't in proper XML or has some invalid values", nil, "MarshalXML", 8}
	commonErrorCodeNoInstallmentPreferredPayPlan = EnumCommonErrorCodeItem{commonErrorCodeNoInstallmentPreferredPayPlanID, "there are no installment preferred pay plans on the policy", nil, "NoInstallmentPreferredPayPlan", 9}
	commonErrorCodeNoPIFPayPlanError             = EnumCommonErrorCodeItem{commonErrorCodeNoPIFPayPlanErrorID, "missing paid in full play plan", nil, "NoPIFPayPlanError", 10}
	commonErrorCodeObjectCantBeBlank             = EnumCommonErrorCodeItem{commonErrorCodeObjectCantBeBlankID, "object is blank or null", nil, "ObjectCantBeBlank", 11}
	commonErrorCodePayPlansEmpty                 = EnumCommonErrorCodeItem{commonErrorCodePayPlansEmptyID, "there are no payment plans on the policy", nil, "PayPlansEmpty", 12}
	commonErrorCodeRuleEngineError               = EnumCommonErrorCodeItem{commonErrorCodeRuleEngineErrorID, "received error from rule engine", nil, "RuleEngineError", 13}
	commonErrorCodeUnauthorized                  = EnumCommonErrorCodeItem{commonErrorCodeUnauthorizedID, "unauthorized request or access", nil, "Unauthorized", 14}
	commonErrorCodeUnmarshalJSON                 = EnumCommonErrorCodeItem{commonErrorCodeUnmarshalJSONID, "data provided isn't in proper JSON or has some invalid values", nil, "UnmarshalJSON", 15}
	commonErrorCodeUnmarshalXML                  = EnumCommonErrorCodeItem{commonErrorCodeUnmarshalXMLID, "data provided isn't in proper XML or has some invalid values", nil, "UnmarshalXML", 16}
	commonErrorCodeValueRequired                 = EnumCommonErrorCodeItem{commonErrorCodeValueRequiredID, "required field", nil, "ValueRequired", 17}
)

// EnumCommonErrorCode is a collection of CommonErrorCode items
type EnumCommonErrorCode struct {
	Description string
	Items       []*EnumCommonErrorCodeItem
	Name        string

	DBRead                        *EnumCommonErrorCodeItem
	DBUnknown                     *EnumCommonErrorCodeItem
	DBWrite                       *EnumCommonErrorCodeItem
	HTTPRequest                   *EnumCommonErrorCodeItem
	HTTPStatus                    *EnumCommonErrorCodeItem
	InvalidValue                  *EnumCommonErrorCodeItem
	MarshalJSON                   *EnumCommonErrorCodeItem
	MarshalXML                    *EnumCommonErrorCodeItem
	NoInstallmentPreferredPayPlan *EnumCommonErrorCodeItem
	NoPIFPayPlanError             *EnumCommonErrorCodeItem
	ObjectCantBeBlank             *EnumCommonErrorCodeItem
	PayPlansEmpty                 *EnumCommonErrorCodeItem
	RuleEngineError               *EnumCommonErrorCodeItem
	Unauthorized                  *EnumCommonErrorCodeItem
	UnmarshalJSON                 *EnumCommonErrorCodeItem
	UnmarshalXML                  *EnumCommonErrorCodeItem
	ValueRequired                 *EnumCommonErrorCodeItem

	itemDict map[string]*EnumCommonErrorCodeItem
}

// CommonErrorCode is a public singleton instance of EnumCommonErrorCode
// representing error codes for common conditions
var CommonErrorCode = &EnumCommonErrorCode{
	Description: "error codes for common conditions",
	Items: []*EnumCommonErrorCodeItem{
		&commonErrorCodeDBRead,
		&commonErrorCodeDBUnknown,
		&commonErrorCodeDBWrite,
		&commonErrorCodeHTTPRequest,
		&commonErrorCodeHTTPStatus,
		&commonErrorCodeInvalidValue,
		&commonErrorCodeMarshalJSON,
		&commonErrorCodeMarshalXML,
		&commonErrorCodeNoInstallmentPreferredPayPlan,
		&commonErrorCodeNoPIFPayPlanError,
		&commonErrorCodeObjectCantBeBlank,
		&commonErrorCodePayPlansEmpty,
		&commonErrorCodeRuleEngineError,
		&commonErrorCodeUnauthorized,
		&commonErrorCodeUnmarshalJSON,
		&commonErrorCodeUnmarshalXML,
		&commonErrorCodeValueRequired,
	},
	Name:                          "EnumCommonErrorCode",
	DBRead:                        &commonErrorCodeDBRead,
	DBUnknown:                     &commonErrorCodeDBUnknown,
	DBWrite:                       &commonErrorCodeDBWrite,
	HTTPRequest:                   &commonErrorCodeHTTPRequest,
	HTTPStatus:                    &commonErrorCodeHTTPStatus,
	InvalidValue:                  &commonErrorCodeInvalidValue,
	MarshalJSON:                   &commonErrorCodeMarshalJSON,
	MarshalXML:                    &commonErrorCodeMarshalXML,
	NoInstallmentPreferredPayPlan: &commonErrorCodeNoInstallmentPreferredPayPlan,
	NoPIFPayPlanError:             &commonErrorCodeNoPIFPayPlanError,
	ObjectCantBeBlank:             &commonErrorCodeObjectCantBeBlank,
	PayPlansEmpty:                 &commonErrorCodePayPlansEmpty,
	RuleEngineError:               &commonErrorCodeRuleEngineError,
	Unauthorized:                  &commonErrorCodeUnauthorized,
	UnmarshalJSON:                 &commonErrorCodeUnmarshalJSON,
	UnmarshalXML:                  &commonErrorCodeUnmarshalXML,
	ValueRequired:                 &commonErrorCodeValueRequired,

	itemDict: map[string]*EnumCommonErrorCodeItem{
		strings.ToLower(string(commonErrorCodeDBReadID)):                        &commonErrorCodeDBRead,
		strings.ToLower(string(commonErrorCodeDBUnknownID)):                     &commonErrorCodeDBUnknown,
		strings.ToLower(string(commonErrorCodeDBWriteID)):                       &commonErrorCodeDBWrite,
		strings.ToLower(string(commonErrorCodeHTTPRequestID)):                   &commonErrorCodeHTTPRequest,
		strings.ToLower(string(commonErrorCodeHTTPStatusID)):                    &commonErrorCodeHTTPStatus,
		strings.ToLower(string(commonErrorCodeInvalidValueID)):                  &commonErrorCodeInvalidValue,
		strings.ToLower(string(commonErrorCodeMarshalJSONID)):                   &commonErrorCodeMarshalJSON,
		strings.ToLower(string(commonErrorCodeMarshalXMLID)):                    &commonErrorCodeMarshalXML,
		strings.ToLower(string(commonErrorCodeNoInstallmentPreferredPayPlanID)): &commonErrorCodeNoInstallmentPreferredPayPlan,
		strings.ToLower(string(commonErrorCodeNoPIFPayPlanErrorID)):             &commonErrorCodeNoPIFPayPlanError,
		strings.ToLower(string(commonErrorCodeObjectCantBeBlankID)):             &commonErrorCodeObjectCantBeBlank,
		strings.ToLower(string(commonErrorCodePayPlansEmptyID)):                 &commonErrorCodePayPlansEmpty,
		strings.ToLower(string(commonErrorCodeRuleEngineErrorID)):               &commonErrorCodeRuleEngineError,
		strings.ToLower(string(commonErrorCodeUnauthorizedID)):                  &commonErrorCodeUnauthorized,
		strings.ToLower(string(commonErrorCodeUnmarshalJSONID)):                 &commonErrorCodeUnmarshalJSON,
		strings.ToLower(string(commonErrorCodeUnmarshalXMLID)):                  &commonErrorCodeUnmarshalXML,
		strings.ToLower(string(commonErrorCodeValueRequiredID)):                 &commonErrorCodeValueRequired,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumCommonErrorCode) ByID(id CommonErrorCodeIdentifier) *EnumCommonErrorCodeItem {
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
func (e *EnumCommonErrorCode) ByIDString(idx string) *EnumCommonErrorCodeItem {
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
func (e *EnumCommonErrorCode) ByIndex(idx int) *EnumCommonErrorCodeItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedCommonErrorCodeID is a struct that is designed to replace a *CommonErrorCodeID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *CommonErrorCodeID it contains while being a better JSON citizen.
type ValidatedCommonErrorCodeID struct {
	// id will point to a valid CommonErrorCodeID, if possible
	// If id is nil, then ValidatedCommonErrorCodeID.Valid() will return false.
	id *CommonErrorCodeID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedCommonErrorCodeID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedCommonErrorCodeID
func (vi *ValidatedCommonErrorCodeID) Clone() *ValidatedCommonErrorCodeID {
	if vi == nil {
		return nil
	}

	var cid *CommonErrorCodeID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedCommonErrorCodeID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedCommonErrorCodeIds represent the same CommonErrorCode
func (vi *ValidatedCommonErrorCodeID) Equals(vj *ValidatedCommonErrorCodeID) bool {
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

// Valid returns true if and only if the ValidatedCommonErrorCodeID corresponds to a recognized CommonErrorCode
func (vi *ValidatedCommonErrorCodeID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedCommonErrorCodeID) ID() *CommonErrorCodeID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedCommonErrorCodeID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedCommonErrorCodeID) ValidatedID() *ValidatedCommonErrorCodeID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedCommonErrorCodeID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedCommonErrorCodeID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedCommonErrorCodeID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedCommonErrorCodeID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedCommonErrorCodeID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := CommonErrorCodeID(capString)
	item := CommonErrorCode.ByID(&id)
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

func (vi ValidatedCommonErrorCodeID) String() string {
	return vi.ToIDString()
}

type CommonErrorCodeIdentifier interface {
	ID() *CommonErrorCodeID
	Valid() bool
}
