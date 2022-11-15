package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// IdentityTypeID uniquely identifies a particular IdentityType
type IdentityTypeID string

// Clone creates a safe, independent copy of a IdentityTypeID
func (i *IdentityTypeID) Clone() *IdentityTypeID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two IdentityTypeIds are equivalent
func (i *IdentityTypeID) Equals(j *IdentityTypeID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *IdentityTypeID that is either valid or nil
func (i *IdentityTypeID) ID() *IdentityTypeID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *IdentityTypeID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the IdentityTypeID corresponds to a recognized IdentityType
func (i *IdentityTypeID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return IdentityType.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *IdentityTypeID) ValidatedID() *ValidatedIdentityTypeID {
	if i != nil {
		return &ValidatedIdentityTypeID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *IdentityTypeID) MarshalJSON() ([]byte, error) {
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

func (i *IdentityTypeID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := IdentityTypeID(dataString)
	item := IdentityType.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	identityTypeTransactionTrackingID IdentityTypeID = "transactionTracking"
	identityTypeQuoteIDID             IdentityTypeID = "quoteID"
	identityTypePartnerIDID           IdentityTypeID = "partnerID"
	identityTypeProgramIDID           IdentityTypeID = "programID"
	identityTypeDealerIDID            IdentityTypeID = "dealerID"
	identityTypeVendorUserIDID        IdentityTypeID = "vendorUserID"
	identityTypeOtherIDID             IdentityTypeID = "otherID"
)

// EnumIdentityTypeItem describes an entry in an enumeration of IdentityType
type EnumIdentityTypeItem struct {
	ID        IdentityTypeID    `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	identityTypeTransactionTracking = EnumIdentityTypeItem{identityTypeTransactionTrackingID, "TransactionTracking", nil, "TransactionTracking", 1}
	identityTypeQuoteID             = EnumIdentityTypeItem{identityTypeQuoteIDID, "QuoteID", nil, "QuoteID", 2}
	identityTypePartnerID           = EnumIdentityTypeItem{identityTypePartnerIDID, "PartnerID", nil, "PartnerID", 3}
	identityTypeProgramID           = EnumIdentityTypeItem{identityTypeProgramIDID, "ProgramID", nil, "ProgramID", 4}
	identityTypeDealerID            = EnumIdentityTypeItem{identityTypeDealerIDID, "DealerID", nil, "DealerID", 5}
	identityTypeVendorUserID        = EnumIdentityTypeItem{identityTypeVendorUserIDID, "VendorUserID", nil, "VendorUserID", 6}
	identityTypeOtherID             = EnumIdentityTypeItem{identityTypeOtherIDID, "OtherID", nil, "OtherID", 7}
)

// EnumIdentityType is a collection of IdentityType items
type EnumIdentityType struct {
	Description string
	Items       []*EnumIdentityTypeItem
	Name        string

	TransactionTracking *EnumIdentityTypeItem
	QuoteID             *EnumIdentityTypeItem
	PartnerID           *EnumIdentityTypeItem
	ProgramID           *EnumIdentityTypeItem
	DealerID            *EnumIdentityTypeItem
	VendorUserID        *EnumIdentityTypeItem
	OtherID             *EnumIdentityTypeItem

	itemDict map[string]*EnumIdentityTypeItem
}

// IdentityType is a public singleton instance of EnumIdentityType
// representing Identity Type
var IdentityType = &EnumIdentityType{
	Description: "Identity Type",
	Items: []*EnumIdentityTypeItem{
		&identityTypeTransactionTracking,
		&identityTypeQuoteID,
		&identityTypePartnerID,
		&identityTypeProgramID,
		&identityTypeDealerID,
		&identityTypeVendorUserID,
		&identityTypeOtherID,
	},
	Name:                "EnumIdentityType",
	TransactionTracking: &identityTypeTransactionTracking,
	QuoteID:             &identityTypeQuoteID,
	PartnerID:           &identityTypePartnerID,
	ProgramID:           &identityTypeProgramID,
	DealerID:            &identityTypeDealerID,
	VendorUserID:        &identityTypeVendorUserID,
	OtherID:             &identityTypeOtherID,

	itemDict: map[string]*EnumIdentityTypeItem{
		strings.ToLower(string(identityTypeTransactionTrackingID)): &identityTypeTransactionTracking,
		strings.ToLower(string(identityTypeQuoteIDID)):             &identityTypeQuoteID,
		strings.ToLower(string(identityTypePartnerIDID)):           &identityTypePartnerID,
		strings.ToLower(string(identityTypeProgramIDID)):           &identityTypeProgramID,
		strings.ToLower(string(identityTypeDealerIDID)):            &identityTypeDealerID,
		strings.ToLower(string(identityTypeVendorUserIDID)):        &identityTypeVendorUserID,
		strings.ToLower(string(identityTypeOtherIDID)):             &identityTypeOtherID,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumIdentityType) ByID(id IdentityTypeIdentifier) *EnumIdentityTypeItem {
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
func (e *EnumIdentityType) ByIDString(idx string) *EnumIdentityTypeItem {
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
func (e *EnumIdentityType) ByIndex(idx int) *EnumIdentityTypeItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedIdentityTypeID is a struct that is designed to replace a *IdentityTypeID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *IdentityTypeID it contains while being a better JSON citizen.
type ValidatedIdentityTypeID struct {
	// id will point to a valid IdentityTypeID, if possible
	// If id is nil, then ValidatedIdentityTypeID.Valid() will return false.
	id *IdentityTypeID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedIdentityTypeID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedIdentityTypeID
func (vi *ValidatedIdentityTypeID) Clone() *ValidatedIdentityTypeID {
	if vi == nil {
		return nil
	}

	var cid *IdentityTypeID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedIdentityTypeID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedIdentityTypeIds represent the same IdentityType
func (vi *ValidatedIdentityTypeID) Equals(vj *ValidatedIdentityTypeID) bool {
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

// Valid returns true if and only if the ValidatedIdentityTypeID corresponds to a recognized IdentityType
func (vi *ValidatedIdentityTypeID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedIdentityTypeID) ID() *IdentityTypeID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedIdentityTypeID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedIdentityTypeID) ValidatedID() *ValidatedIdentityTypeID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedIdentityTypeID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedIdentityTypeID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedIdentityTypeID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedIdentityTypeID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedIdentityTypeID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := IdentityTypeID(capString)
	item := IdentityType.ByID(&id)
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

func (vi ValidatedIdentityTypeID) String() string {
	return vi.ToIDString()
}

type IdentityTypeIdentifier interface {
	ID() *IdentityTypeID
	Valid() bool
}
