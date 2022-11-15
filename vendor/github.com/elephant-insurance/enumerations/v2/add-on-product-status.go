package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// AddOnProductStatusID uniquely identifies a particular AddOnProductStatus
type AddOnProductStatusID string

// Clone creates a safe, independent copy of a AddOnProductStatusID
func (i *AddOnProductStatusID) Clone() *AddOnProductStatusID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two AddOnProductStatusIds are equivalent
func (i *AddOnProductStatusID) Equals(j *AddOnProductStatusID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *AddOnProductStatusID that is either valid or nil
func (i *AddOnProductStatusID) ID() *AddOnProductStatusID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *AddOnProductStatusID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the AddOnProductStatusID corresponds to a recognized AddOnProductStatus
func (i *AddOnProductStatusID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return AddOnProductStatus.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *AddOnProductStatusID) ValidatedID() *ValidatedAddOnProductStatusID {
	if i != nil {
		return &ValidatedAddOnProductStatusID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *AddOnProductStatusID) MarshalJSON() ([]byte, error) {
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

func (i *AddOnProductStatusID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := AddOnProductStatusID(dataString)
	item := AddOnProductStatus.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	addOnProductStatusDeliveredID                          AddOnProductStatusID = "delivered"
	addOnProductStatusNotAvailableID                       AddOnProductStatusID = "notAvailable"
	addOnProductStatusNotAuthorizedID                      AddOnProductStatusID = "notAuthorized"
	addOnProductStatusDefaultDeliveredID                   AddOnProductStatusID = "defaultDelivered"
	addOnProductStatusInvalidRequestID                     AddOnProductStatusID = "invalidRequest"
	addOnProductStatusDataUnavailableID                    AddOnProductStatusID = "dataUnavailable"
	addOnProductStatusInvalidProcessingRequestID           AddOnProductStatusID = "invalidProcessingRequest"
	addOnProductStatusDefaultServiceCanceledID             AddOnProductStatusID = "defaultServiceCanceled"
	addOnProductStatusProcessingRequestCannotBeEvaluatedID AddOnProductStatusID = "processingRequestCannotBeEvaluated"
	addOnProductStatusInvalidInputModelAttributeID         AddOnProductStatusID = "invalidInputModelAttribute"
	addOnProductStatusTransactionIDTooLongID               AddOnProductStatusID = "transactionIDTooLong"
	addOnProductStatusInputDataMissingOrInvalidID          AddOnProductStatusID = "inputDataMissingOrInvalid"
	addOnProductStatusExceededMaximumConsumerThresholdID   AddOnProductStatusID = "exceededMaximumConsumerThreshold"
	addOnProductStatusRequestReceivedID                    AddOnProductStatusID = "requestReceived"
)

// EnumAddOnProductStatusItem describes an entry in an enumeration of AddOnProductStatus
type EnumAddOnProductStatusItem struct {
	ID        AddOnProductStatusID `json:"Value"`
	Desc      string               `json:"Description,omitempty"`
	Meta      map[string]string    `json:",omitempty"`
	Name      string               `json:"Name"`
	SortOrder int
}

var (
	addOnProductStatusDelivered                          = EnumAddOnProductStatusItem{addOnProductStatusDeliveredID, "Delivered", nil, "Delivered", 1}
	addOnProductStatusNotAvailable                       = EnumAddOnProductStatusItem{addOnProductStatusNotAvailableID, "NotAvailable", nil, "NotAvailable", 2}
	addOnProductStatusNotAuthorized                      = EnumAddOnProductStatusItem{addOnProductStatusNotAuthorizedID, "NotAuthorized", nil, "NotAuthorized", 3}
	addOnProductStatusDefaultDelivered                   = EnumAddOnProductStatusItem{addOnProductStatusDefaultDeliveredID, "DefaultDelivered", nil, "DefaultDelivered", 4}
	addOnProductStatusInvalidRequest                     = EnumAddOnProductStatusItem{addOnProductStatusInvalidRequestID, "InvalidRequest", nil, "InvalidRequest", 5}
	addOnProductStatusDataUnavailable                    = EnumAddOnProductStatusItem{addOnProductStatusDataUnavailableID, "DataUnavailable", nil, "DataUnavailable", 6}
	addOnProductStatusInvalidProcessingRequest           = EnumAddOnProductStatusItem{addOnProductStatusInvalidProcessingRequestID, "InvalidProcessingRequest", nil, "InvalidProcessingRequest", 7}
	addOnProductStatusDefaultServiceCanceled             = EnumAddOnProductStatusItem{addOnProductStatusDefaultServiceCanceledID, "DefaultServiceCanceled", nil, "DefaultServiceCanceled", 8}
	addOnProductStatusProcessingRequestCannotBeEvaluated = EnumAddOnProductStatusItem{addOnProductStatusProcessingRequestCannotBeEvaluatedID, "ProcessingRequestCannotBeEvaluated", nil, "ProcessingRequestCannotBeEvaluated", 9}
	addOnProductStatusInvalidInputModelAttribute         = EnumAddOnProductStatusItem{addOnProductStatusInvalidInputModelAttributeID, "InvalidInputModelAttribute", nil, "InvalidInputModelAttribute", 10}
	addOnProductStatusTransactionIDTooLong               = EnumAddOnProductStatusItem{addOnProductStatusTransactionIDTooLongID, "TransactionIDTooLong", nil, "TransactionIDTooLong", 11}
	addOnProductStatusInputDataMissingOrInvalid          = EnumAddOnProductStatusItem{addOnProductStatusInputDataMissingOrInvalidID, "InputDataMissingOrInvalid", nil, "InputDataMissingOrInvalid", 12}
	addOnProductStatusExceededMaximumConsumerThreshold   = EnumAddOnProductStatusItem{addOnProductStatusExceededMaximumConsumerThresholdID, "ExceededMaximumConsumerThreshold", nil, "ExceededMaximumConsumerThreshold", 13}
	addOnProductStatusRequestReceived                    = EnumAddOnProductStatusItem{addOnProductStatusRequestReceivedID, "RequestReceived", nil, "RequestReceived", 14}
)

// EnumAddOnProductStatus is a collection of AddOnProductStatus items
type EnumAddOnProductStatus struct {
	Description string
	Items       []*EnumAddOnProductStatusItem
	Name        string

	Delivered                          *EnumAddOnProductStatusItem
	NotAvailable                       *EnumAddOnProductStatusItem
	NotAuthorized                      *EnumAddOnProductStatusItem
	DefaultDelivered                   *EnumAddOnProductStatusItem
	InvalidRequest                     *EnumAddOnProductStatusItem
	DataUnavailable                    *EnumAddOnProductStatusItem
	InvalidProcessingRequest           *EnumAddOnProductStatusItem
	DefaultServiceCanceled             *EnumAddOnProductStatusItem
	ProcessingRequestCannotBeEvaluated *EnumAddOnProductStatusItem
	InvalidInputModelAttribute         *EnumAddOnProductStatusItem
	TransactionIDTooLong               *EnumAddOnProductStatusItem
	InputDataMissingOrInvalid          *EnumAddOnProductStatusItem
	ExceededMaximumConsumerThreshold   *EnumAddOnProductStatusItem
	RequestReceived                    *EnumAddOnProductStatusItem

	itemDict map[string]*EnumAddOnProductStatusItem
}

// AddOnProductStatus is a public singleton instance of EnumAddOnProductStatus
// representing add on product status for DHI
var AddOnProductStatus = &EnumAddOnProductStatus{
	Description: "add on product status for DHI",
	Items: []*EnumAddOnProductStatusItem{
		&addOnProductStatusDelivered,
		&addOnProductStatusNotAvailable,
		&addOnProductStatusNotAuthorized,
		&addOnProductStatusDefaultDelivered,
		&addOnProductStatusInvalidRequest,
		&addOnProductStatusDataUnavailable,
		&addOnProductStatusInvalidProcessingRequest,
		&addOnProductStatusDefaultServiceCanceled,
		&addOnProductStatusProcessingRequestCannotBeEvaluated,
		&addOnProductStatusInvalidInputModelAttribute,
		&addOnProductStatusTransactionIDTooLong,
		&addOnProductStatusInputDataMissingOrInvalid,
		&addOnProductStatusExceededMaximumConsumerThreshold,
		&addOnProductStatusRequestReceived,
	},
	Name:                               "EnumAddOnProductStatus",
	Delivered:                          &addOnProductStatusDelivered,
	NotAvailable:                       &addOnProductStatusNotAvailable,
	NotAuthorized:                      &addOnProductStatusNotAuthorized,
	DefaultDelivered:                   &addOnProductStatusDefaultDelivered,
	InvalidRequest:                     &addOnProductStatusInvalidRequest,
	DataUnavailable:                    &addOnProductStatusDataUnavailable,
	InvalidProcessingRequest:           &addOnProductStatusInvalidProcessingRequest,
	DefaultServiceCanceled:             &addOnProductStatusDefaultServiceCanceled,
	ProcessingRequestCannotBeEvaluated: &addOnProductStatusProcessingRequestCannotBeEvaluated,
	InvalidInputModelAttribute:         &addOnProductStatusInvalidInputModelAttribute,
	TransactionIDTooLong:               &addOnProductStatusTransactionIDTooLong,
	InputDataMissingOrInvalid:          &addOnProductStatusInputDataMissingOrInvalid,
	ExceededMaximumConsumerThreshold:   &addOnProductStatusExceededMaximumConsumerThreshold,
	RequestReceived:                    &addOnProductStatusRequestReceived,

	itemDict: map[string]*EnumAddOnProductStatusItem{
		strings.ToLower(string(addOnProductStatusDeliveredID)):                          &addOnProductStatusDelivered,
		strings.ToLower(string(addOnProductStatusNotAvailableID)):                       &addOnProductStatusNotAvailable,
		strings.ToLower(string(addOnProductStatusNotAuthorizedID)):                      &addOnProductStatusNotAuthorized,
		strings.ToLower(string(addOnProductStatusDefaultDeliveredID)):                   &addOnProductStatusDefaultDelivered,
		strings.ToLower(string(addOnProductStatusInvalidRequestID)):                     &addOnProductStatusInvalidRequest,
		strings.ToLower(string(addOnProductStatusDataUnavailableID)):                    &addOnProductStatusDataUnavailable,
		strings.ToLower(string(addOnProductStatusInvalidProcessingRequestID)):           &addOnProductStatusInvalidProcessingRequest,
		strings.ToLower(string(addOnProductStatusDefaultServiceCanceledID)):             &addOnProductStatusDefaultServiceCanceled,
		strings.ToLower(string(addOnProductStatusProcessingRequestCannotBeEvaluatedID)): &addOnProductStatusProcessingRequestCannotBeEvaluated,
		strings.ToLower(string(addOnProductStatusInvalidInputModelAttributeID)):         &addOnProductStatusInvalidInputModelAttribute,
		strings.ToLower(string(addOnProductStatusTransactionIDTooLongID)):               &addOnProductStatusTransactionIDTooLong,
		strings.ToLower(string(addOnProductStatusInputDataMissingOrInvalidID)):          &addOnProductStatusInputDataMissingOrInvalid,
		strings.ToLower(string(addOnProductStatusExceededMaximumConsumerThresholdID)):   &addOnProductStatusExceededMaximumConsumerThreshold,
		strings.ToLower(string(addOnProductStatusRequestReceivedID)):                    &addOnProductStatusRequestReceived,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumAddOnProductStatus) ByID(id AddOnProductStatusIdentifier) *EnumAddOnProductStatusItem {
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
func (e *EnumAddOnProductStatus) ByIDString(idx string) *EnumAddOnProductStatusItem {
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
func (e *EnumAddOnProductStatus) ByIndex(idx int) *EnumAddOnProductStatusItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedAddOnProductStatusID is a struct that is designed to replace a *AddOnProductStatusID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *AddOnProductStatusID it contains while being a better JSON citizen.
type ValidatedAddOnProductStatusID struct {
	// id will point to a valid AddOnProductStatusID, if possible
	// If id is nil, then ValidatedAddOnProductStatusID.Valid() will return false.
	id *AddOnProductStatusID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedAddOnProductStatusID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedAddOnProductStatusID
func (vi *ValidatedAddOnProductStatusID) Clone() *ValidatedAddOnProductStatusID {
	if vi == nil {
		return nil
	}

	var cid *AddOnProductStatusID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedAddOnProductStatusID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedAddOnProductStatusIds represent the same AddOnProductStatus
func (vi *ValidatedAddOnProductStatusID) Equals(vj *ValidatedAddOnProductStatusID) bool {
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

// Valid returns true if and only if the ValidatedAddOnProductStatusID corresponds to a recognized AddOnProductStatus
func (vi *ValidatedAddOnProductStatusID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedAddOnProductStatusID) ID() *AddOnProductStatusID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedAddOnProductStatusID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedAddOnProductStatusID) ValidatedID() *ValidatedAddOnProductStatusID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedAddOnProductStatusID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedAddOnProductStatusID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedAddOnProductStatusID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedAddOnProductStatusID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedAddOnProductStatusID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := AddOnProductStatusID(capString)
	item := AddOnProductStatus.ByID(&id)
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

func (vi ValidatedAddOnProductStatusID) String() string {
	return vi.ToIDString()
}

type AddOnProductStatusIdentifier interface {
	ID() *AddOnProductStatusID
	Valid() bool
}
