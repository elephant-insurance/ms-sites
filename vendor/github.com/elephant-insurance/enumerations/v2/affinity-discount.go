package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// AffinityDiscountTypeID uniquely identifies a particular AffinityDiscountType
type AffinityDiscountTypeID string

// Clone creates a safe, independent copy of a AffinityDiscountTypeID
func (i *AffinityDiscountTypeID) Clone() *AffinityDiscountTypeID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two AffinityDiscountTypeIds are equivalent
func (i *AffinityDiscountTypeID) Equals(j *AffinityDiscountTypeID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *AffinityDiscountTypeID that is either valid or nil
func (i *AffinityDiscountTypeID) ID() *AffinityDiscountTypeID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *AffinityDiscountTypeID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the AffinityDiscountTypeID corresponds to a recognized AffinityDiscountType
func (i *AffinityDiscountTypeID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return AffinityDiscountType.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *AffinityDiscountTypeID) ValidatedID() *ValidatedAffinityDiscountTypeID {
	if i != nil {
		return &ValidatedAffinityDiscountTypeID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *AffinityDiscountTypeID) MarshalJSON() ([]byte, error) {
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

func (i *AffinityDiscountTypeID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := AffinityDiscountTypeID(dataString)
	item := AffinityDiscountType.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	affinityDiscountTypeAffinityDiscountNoneID    AffinityDiscountTypeID = "None"
	affinityDiscountTypeReferralDiscountID        AffinityDiscountTypeID = "ReferralCode"
	affinityDiscountTypeAceableID                 AffinityDiscountTypeID = "Aceable"
	affinityDiscountTypeAdmiralEmployeeDiscountID AffinityDiscountTypeID = "Admiral_Employee"
)

// EnumAffinityDiscountTypeItem describes an entry in an enumeration of AffinityDiscountType
type EnumAffinityDiscountTypeItem struct {
	ID        AffinityDiscountTypeID `json:"Value"`
	Desc      string                 `json:"Description,omitempty"`
	Meta      map[string]string      `json:",omitempty"`
	Name      string                 `json:"Name"`
	SortOrder int

	// Meta Properties
	StateCodes string
}

var (
	affinityDiscountTypeAffinityDiscountNone    = EnumAffinityDiscountTypeItem{affinityDiscountTypeAffinityDiscountNoneID, "None", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "AffinityDiscountNone", 1, "IN,TN,IL,MD,TX,VA,GA,OH"}
	affinityDiscountTypeReferralDiscount        = EnumAffinityDiscountTypeItem{affinityDiscountTypeReferralDiscountID, "Referral Discount", map[string]string{"StateCodes": "IL,MD,TX,VA"}, "ReferralDiscount", 2, "IL,MD,TX,VA"}
	affinityDiscountTypeAceable                 = EnumAffinityDiscountTypeItem{affinityDiscountTypeAceableID, "Aceable Discount", map[string]string{"StateCodes": "IN,IL,MD,TX,VA,GA,OH"}, "Aceable", 3, "IN,IL,MD,TX,VA,GA,OH"}
	affinityDiscountTypeAdmiralEmployeeDiscount = EnumAffinityDiscountTypeItem{affinityDiscountTypeAdmiralEmployeeDiscountID, "Admiral Employee Discount", map[string]string{"StateCodes": "IN,IL,MD,TX,VA,GA,OH"}, "AdmiralEmployeeDiscount", 4, "IN,IL,MD,TX,VA,GA,OH"}
)

// EnumAffinityDiscountType is a collection of AffinityDiscountType items
type EnumAffinityDiscountType struct {
	Description string
	Items       []*EnumAffinityDiscountTypeItem
	Name        string

	AffinityDiscountNone    *EnumAffinityDiscountTypeItem
	ReferralDiscount        *EnumAffinityDiscountTypeItem
	Aceable                 *EnumAffinityDiscountTypeItem
	AdmiralEmployeeDiscount *EnumAffinityDiscountTypeItem

	itemDict map[string]*EnumAffinityDiscountTypeItem
}

// AffinityDiscountType is a public singleton instance of EnumAffinityDiscountType
// representing types of affinity discounts
var AffinityDiscountType = &EnumAffinityDiscountType{
	Description: "types of affinity discounts",
	Items: []*EnumAffinityDiscountTypeItem{
		&affinityDiscountTypeAffinityDiscountNone,
		&affinityDiscountTypeReferralDiscount,
		&affinityDiscountTypeAceable,
		&affinityDiscountTypeAdmiralEmployeeDiscount,
	},
	Name:                    "EnumAffinityDiscountType",
	AffinityDiscountNone:    &affinityDiscountTypeAffinityDiscountNone,
	ReferralDiscount:        &affinityDiscountTypeReferralDiscount,
	Aceable:                 &affinityDiscountTypeAceable,
	AdmiralEmployeeDiscount: &affinityDiscountTypeAdmiralEmployeeDiscount,

	itemDict: map[string]*EnumAffinityDiscountTypeItem{
		strings.ToLower(string(affinityDiscountTypeAffinityDiscountNoneID)):    &affinityDiscountTypeAffinityDiscountNone,
		strings.ToLower(string(affinityDiscountTypeReferralDiscountID)):        &affinityDiscountTypeReferralDiscount,
		strings.ToLower(string(affinityDiscountTypeAceableID)):                 &affinityDiscountTypeAceable,
		strings.ToLower(string(affinityDiscountTypeAdmiralEmployeeDiscountID)): &affinityDiscountTypeAdmiralEmployeeDiscount,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumAffinityDiscountType) ByID(id AffinityDiscountTypeIdentifier) *EnumAffinityDiscountTypeItem {
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
func (e *EnumAffinityDiscountType) ByIDString(idx string) *EnumAffinityDiscountTypeItem {
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
func (e *EnumAffinityDiscountType) ByIndex(idx int) *EnumAffinityDiscountTypeItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedAffinityDiscountTypeID is a struct that is designed to replace a *AffinityDiscountTypeID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *AffinityDiscountTypeID it contains while being a better JSON citizen.
type ValidatedAffinityDiscountTypeID struct {
	// id will point to a valid AffinityDiscountTypeID, if possible
	// If id is nil, then ValidatedAffinityDiscountTypeID.Valid() will return false.
	id *AffinityDiscountTypeID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedAffinityDiscountTypeID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedAffinityDiscountTypeID
func (vi *ValidatedAffinityDiscountTypeID) Clone() *ValidatedAffinityDiscountTypeID {
	if vi == nil {
		return nil
	}

	var cid *AffinityDiscountTypeID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedAffinityDiscountTypeID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedAffinityDiscountTypeIds represent the same AffinityDiscountType
func (vi *ValidatedAffinityDiscountTypeID) Equals(vj *ValidatedAffinityDiscountTypeID) bool {
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

// Valid returns true if and only if the ValidatedAffinityDiscountTypeID corresponds to a recognized AffinityDiscountType
func (vi *ValidatedAffinityDiscountTypeID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedAffinityDiscountTypeID) ID() *AffinityDiscountTypeID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedAffinityDiscountTypeID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedAffinityDiscountTypeID) ValidatedID() *ValidatedAffinityDiscountTypeID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedAffinityDiscountTypeID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedAffinityDiscountTypeID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedAffinityDiscountTypeID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedAffinityDiscountTypeID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedAffinityDiscountTypeID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := AffinityDiscountTypeID(capString)
	item := AffinityDiscountType.ByID(&id)
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

func (vi ValidatedAffinityDiscountTypeID) String() string {
	return vi.ToIDString()
}

type AffinityDiscountTypeIdentifier interface {
	ID() *AffinityDiscountTypeID
	Valid() bool
}
