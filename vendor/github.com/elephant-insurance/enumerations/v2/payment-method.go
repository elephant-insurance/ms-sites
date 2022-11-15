package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// PaymentMethodID uniquely identifies a particular PaymentMethod
type PaymentMethodID string

// Clone creates a safe, independent copy of a PaymentMethodID
func (i *PaymentMethodID) Clone() *PaymentMethodID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two PaymentMethodIds are equivalent
func (i *PaymentMethodID) Equals(j *PaymentMethodID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *PaymentMethodID that is either valid or nil
func (i *PaymentMethodID) ID() *PaymentMethodID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *PaymentMethodID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the PaymentMethodID corresponds to a recognized PaymentMethod
func (i *PaymentMethodID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return PaymentMethod.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *PaymentMethodID) ValidatedID() *ValidatedPaymentMethodID {
	if i != nil {
		return &ValidatedPaymentMethodID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *PaymentMethodID) MarshalJSON() ([]byte, error) {
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

func (i *PaymentMethodID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := PaymentMethodID(dataString)
	item := PaymentMethod.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	paymentMethodACHID        PaymentMethodID = "ACH"
	paymentMethodCreditCardID PaymentMethodID = "CreditCard"
	paymentMethodResponsiveID PaymentMethodID = "Responsive"
	paymentMethodWireID       PaymentMethodID = "Wire"
	paymentMethodCashID       PaymentMethodID = "Cash"
)

// EnumPaymentMethodItem describes an entry in an enumeration of PaymentMethod
type EnumPaymentMethodItem struct {
	ID        PaymentMethodID   `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	paymentMethodACH        = EnumPaymentMethodItem{paymentMethodACHID, "ACH", nil, "ACH", 1}
	paymentMethodCreditCard = EnumPaymentMethodItem{paymentMethodCreditCardID, "CreditCard", nil, "CreditCard", 2}
	paymentMethodResponsive = EnumPaymentMethodItem{paymentMethodResponsiveID, "Responsive", nil, "Responsive", 3}
	paymentMethodWire       = EnumPaymentMethodItem{paymentMethodWireID, "Wire", nil, "Wire", 4}
	paymentMethodCash       = EnumPaymentMethodItem{paymentMethodCashID, "Cash", nil, "Cash", 5}
)

// EnumPaymentMethod is a collection of PaymentMethod items
type EnumPaymentMethod struct {
	Description string
	Items       []*EnumPaymentMethodItem
	Name        string

	ACH        *EnumPaymentMethodItem
	CreditCard *EnumPaymentMethodItem
	Responsive *EnumPaymentMethodItem
	Wire       *EnumPaymentMethodItem
	Cash       *EnumPaymentMethodItem

	itemDict map[string]*EnumPaymentMethodItem
}

// PaymentMethod is a public singleton instance of EnumPaymentMethod
// representing methods of paying insurance premium
var PaymentMethod = &EnumPaymentMethod{
	Description: "methods of paying insurance premium",
	Items: []*EnumPaymentMethodItem{
		&paymentMethodACH,
		&paymentMethodCreditCard,
		&paymentMethodResponsive,
		&paymentMethodWire,
		&paymentMethodCash,
	},
	Name:       "EnumPaymentMethod",
	ACH:        &paymentMethodACH,
	CreditCard: &paymentMethodCreditCard,
	Responsive: &paymentMethodResponsive,
	Wire:       &paymentMethodWire,
	Cash:       &paymentMethodCash,

	itemDict: map[string]*EnumPaymentMethodItem{
		strings.ToLower(string(paymentMethodACHID)):        &paymentMethodACH,
		strings.ToLower(string(paymentMethodCreditCardID)): &paymentMethodCreditCard,
		strings.ToLower(string(paymentMethodResponsiveID)): &paymentMethodResponsive,
		strings.ToLower(string(paymentMethodWireID)):       &paymentMethodWire,
		strings.ToLower(string(paymentMethodCashID)):       &paymentMethodCash,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumPaymentMethod) ByID(id PaymentMethodIdentifier) *EnumPaymentMethodItem {
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
func (e *EnumPaymentMethod) ByIDString(idx string) *EnumPaymentMethodItem {
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
func (e *EnumPaymentMethod) ByIndex(idx int) *EnumPaymentMethodItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedPaymentMethodID is a struct that is designed to replace a *PaymentMethodID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *PaymentMethodID it contains while being a better JSON citizen.
type ValidatedPaymentMethodID struct {
	// id will point to a valid PaymentMethodID, if possible
	// If id is nil, then ValidatedPaymentMethodID.Valid() will return false.
	id *PaymentMethodID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedPaymentMethodID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedPaymentMethodID
func (vi *ValidatedPaymentMethodID) Clone() *ValidatedPaymentMethodID {
	if vi == nil {
		return nil
	}

	var cid *PaymentMethodID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedPaymentMethodID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedPaymentMethodIds represent the same PaymentMethod
func (vi *ValidatedPaymentMethodID) Equals(vj *ValidatedPaymentMethodID) bool {
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

// Valid returns true if and only if the ValidatedPaymentMethodID corresponds to a recognized PaymentMethod
func (vi *ValidatedPaymentMethodID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedPaymentMethodID) ID() *PaymentMethodID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedPaymentMethodID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedPaymentMethodID) ValidatedID() *ValidatedPaymentMethodID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedPaymentMethodID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedPaymentMethodID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedPaymentMethodID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedPaymentMethodID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedPaymentMethodID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := PaymentMethodID(capString)
	item := PaymentMethod.ByID(&id)
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

func (vi ValidatedPaymentMethodID) String() string {
	return vi.ToIDString()
}

type PaymentMethodIdentifier interface {
	ID() *PaymentMethodID
	Valid() bool
}
