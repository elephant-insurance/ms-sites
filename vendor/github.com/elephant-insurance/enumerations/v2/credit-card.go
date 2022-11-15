package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// CreditCardID uniquely identifies a particular CreditCard
type CreditCardID string

// Clone creates a safe, independent copy of a CreditCardID
func (i *CreditCardID) Clone() *CreditCardID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two CreditCardIds are equivalent
func (i *CreditCardID) Equals(j *CreditCardID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *CreditCardID that is either valid or nil
func (i *CreditCardID) ID() *CreditCardID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *CreditCardID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the CreditCardID corresponds to a recognized CreditCard
func (i *CreditCardID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return CreditCard.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *CreditCardID) ValidatedID() *ValidatedCreditCardID {
	if i != nil {
		return &ValidatedCreditCardID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *CreditCardID) MarshalJSON() ([]byte, error) {
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

func (i *CreditCardID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := CreditCardID(dataString)
	item := CreditCard.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	creditCardVisaID            CreditCardID = "visa"
	creditCardMasterCardID      CreditCardID = "mastercard"
	creditCardDiscoverID        CreditCardID = "discover"
	creditCardAmericanExpressID CreditCardID = "amex"
)

// EnumCreditCardItem describes an entry in an enumeration of CreditCard
type EnumCreditCardItem struct {
	ID        CreditCardID      `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	RexgexCCType string
	Class        string
}

var (
	creditCardVisa            = EnumCreditCardItem{creditCardVisaID, "Visa", map[string]string{"RexgexCCType": "^4", "Class": "pf-visa"}, "Visa", 1, "^4", "pf-visa"}
	creditCardMasterCard      = EnumCreditCardItem{creditCardMasterCardID, "MasterCard", map[string]string{"RexgexCCType": "^5[1-5]|^2[1-7]", "Class": "pf-mastercard"}, "MasterCard", 2, "^5[1-5]|^2[1-7]", "pf-mastercard"}
	creditCardDiscover        = EnumCreditCardItem{creditCardDiscoverID, "Discover", map[string]string{"RexgexCCType": "^(6011)|^(622(1(2[6-9]|[3-9][0-9])|[2-8][0-9]{2}|9([01][0-9]|2[0-5])))|^(64[4-9])|^65", "Class": "pf-discover"}, "Discover", 3, "^(6011)|^(622(1(2[6-9]|[3-9][0-9])|[2-8][0-9]{2}|9([01][0-9]|2[0-5])))|^(64[4-9])|^65", "pf-discover"}
	creditCardAmericanExpress = EnumCreditCardItem{creditCardAmericanExpressID, "American Express", map[string]string{"RexgexCCType": "^(34)|^(37)", "Class": "pf-american-express"}, "AmericanExpress", 4, "^(34)|^(37)", "pf-american-express"}
)

// EnumCreditCard is a collection of CreditCard items
type EnumCreditCard struct {
	Description string
	Items       []*EnumCreditCardItem
	Name        string

	Visa            *EnumCreditCardItem
	MasterCard      *EnumCreditCardItem
	Discover        *EnumCreditCardItem
	AmericanExpress *EnumCreditCardItem

	itemDict map[string]*EnumCreditCardItem
}

// CreditCard is a public singleton instance of EnumCreditCard
// representing credit cards
var CreditCard = &EnumCreditCard{
	Description: "credit cards",
	Items: []*EnumCreditCardItem{
		&creditCardVisa,
		&creditCardMasterCard,
		&creditCardDiscover,
		&creditCardAmericanExpress,
	},
	Name:            "EnumCreditCard",
	Visa:            &creditCardVisa,
	MasterCard:      &creditCardMasterCard,
	Discover:        &creditCardDiscover,
	AmericanExpress: &creditCardAmericanExpress,

	itemDict: map[string]*EnumCreditCardItem{
		strings.ToLower(string(creditCardVisaID)):            &creditCardVisa,
		strings.ToLower(string(creditCardMasterCardID)):      &creditCardMasterCard,
		strings.ToLower(string(creditCardDiscoverID)):        &creditCardDiscover,
		strings.ToLower(string(creditCardAmericanExpressID)): &creditCardAmericanExpress,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumCreditCard) ByID(id CreditCardIdentifier) *EnumCreditCardItem {
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
func (e *EnumCreditCard) ByIDString(idx string) *EnumCreditCardItem {
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
func (e *EnumCreditCard) ByIndex(idx int) *EnumCreditCardItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedCreditCardID is a struct that is designed to replace a *CreditCardID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *CreditCardID it contains while being a better JSON citizen.
type ValidatedCreditCardID struct {
	// id will point to a valid CreditCardID, if possible
	// If id is nil, then ValidatedCreditCardID.Valid() will return false.
	id *CreditCardID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedCreditCardID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedCreditCardID
func (vi *ValidatedCreditCardID) Clone() *ValidatedCreditCardID {
	if vi == nil {
		return nil
	}

	var cid *CreditCardID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedCreditCardID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedCreditCardIds represent the same CreditCard
func (vi *ValidatedCreditCardID) Equals(vj *ValidatedCreditCardID) bool {
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

// Valid returns true if and only if the ValidatedCreditCardID corresponds to a recognized CreditCard
func (vi *ValidatedCreditCardID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedCreditCardID) ID() *CreditCardID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedCreditCardID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedCreditCardID) ValidatedID() *ValidatedCreditCardID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedCreditCardID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedCreditCardID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedCreditCardID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedCreditCardID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedCreditCardID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := CreditCardID(capString)
	item := CreditCard.ByID(&id)
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

func (vi ValidatedCreditCardID) String() string {
	return vi.ToIDString()
}

type CreditCardIdentifier interface {
	ID() *CreditCardID
	Valid() bool
}
