package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// PaymentPlanID uniquely identifies a particular PaymentPlan
type PaymentPlanID string

// Clone creates a safe, independent copy of a PaymentPlanID
func (i *PaymentPlanID) Clone() *PaymentPlanID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two PaymentPlanIds are equivalent
func (i *PaymentPlanID) Equals(j *PaymentPlanID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *PaymentPlanID that is either valid or nil
func (i *PaymentPlanID) ID() *PaymentPlanID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *PaymentPlanID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the PaymentPlanID corresponds to a recognized PaymentPlan
func (i *PaymentPlanID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return PaymentPlan.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *PaymentPlanID) ValidatedID() *ValidatedPaymentPlanID {
	if i != nil {
		return &ValidatedPaymentPlanID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *PaymentPlanID) MarshalJSON() ([]byte, error) {
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

func (i *PaymentPlanID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := PaymentPlanID(dataString)
	item := PaymentPlan.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	paymentPlanAnnualRenewalCatchUpID          PaymentPlanID = "PayPlan:CtchRnw"
	paymentPlanAnnualRenewalStandardID         PaymentPlanID = "PayPlan:StdRnw"
	paymentPlanAnnual1DownID                   PaymentPlanID = "PayPlan:OneMonthDown"
	paymentPlanAnnual1DownSelectDueDateID      PaymentPlanID = "PayPlan:1MnthDwnDue2"
	paymentPlanAnnual1AndHalfDownID            PaymentPlanID = "PayPlan:OneHalf"
	paymentPlanAnnual2DownID                   PaymentPlanID = "PayPlan:AutoMnth"
	paymentPlanAnnual6Down6PaymentID           PaymentPlanID = "PayPlan:SixMthly"
	paymentPlanAnnual6Down1PaymentBillMeID     PaymentPlanID = "PayPlan:TwoPays"
	paymentPlanAnnual6Down1PaymentAutoID       PaymentPlanID = "PayPlan:TwoPaysRecur"
	paymentPlanAnnualPayInFullID               PaymentPlanID = "PayPlan:Full"
	paymentPlanSemiHalfDown5PaymentID          PaymentPlanID = "PayPlan:HalfMnthDwn6M"
	paymentPlanSemi1DownID                     PaymentPlanID = "PayPlan:1MnthDwn6M"
	paymentPlanSemi1DownSelectDueDateID        PaymentPlanID = "PayPlan:1MnthDwnDue6M"
	paymentPlanSemi2DownID                     PaymentPlanID = "PayPlan:2MnthDwn6M"
	paymentPlanSemi3Down1PaymentID             PaymentPlanID = "PayPlan:TwoPay6M"
	paymentPlanSemiPayInFullID                 PaymentPlanID = "PayPlan:Full6M"
	paymentPlanSemi1DownSelectDefaultDueDateID PaymentPlanID = "PayPlan:1MDefaultDue"
	paymentPlanSemi2DownDefaultDueDateID       PaymentPlanID = "PayPlan:2MDefaultDue"
)

// EnumPaymentPlanItem describes an entry in an enumeration of PaymentPlan
type EnumPaymentPlanItem struct {
	ID        PaymentPlanID     `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	TermLength    string
	Installments  string
	DownPCT       string
	InstallPCT    string
	ValidInStates string
}

var (
	paymentPlanAnnualRenewalCatchUp          = EnumPaymentPlanItem{paymentPlanAnnualRenewalCatchUpID, "Renewal Catch - Up Plan - DP due day before renewal", map[string]string{"TermLength": "12", "Installments": "10", "DownPCT": "0    ", "InstallPCT": "9.09 ", "ValidInStates": "IL,IN,MD,TN,TX,VA"}, "AnnualRenewalCatchUp", 1, "12", "10", "0    ", "9.09 ", "IL,IN,MD,TN,TX,VA"}
	paymentPlanAnnualRenewalStandard         = EnumPaymentPlanItem{paymentPlanAnnualRenewalStandardID, "Standard Renewal Plan - DP due 1 month before renewal", map[string]string{"TermLength": "12", "Installments": "11", "DownPCT": "0    ", "InstallPCT": "8.33 ", "ValidInStates": "IL,IN,MD,TN,TX,VA"}, "AnnualRenewalStandard", 2, "12", "11", "0    ", "8.33 ", "IL,IN,MD,TN,TX,VA"}
	paymentPlanAnnual1Down                   = EnumPaymentPlanItem{paymentPlanAnnual1DownID, "1 Month Down, Auto Pay", map[string]string{"TermLength": "12", "Installments": "11", "DownPCT": "8.37 ", "InstallPCT": "8.33 ", "ValidInStates": "IL,IN,MD,TN,TX,VA"}, "Annual1Down", 3, "12", "11", "8.37 ", "8.33 ", "IL,IN,MD,TN,TX,VA"}
	paymentPlanAnnual1DownSelectDueDate      = EnumPaymentPlanItem{paymentPlanAnnual1DownSelectDueDateID, "1 Month Down, Auto Pay - Select Due Date", map[string]string{"TermLength": "12", "Installments": "10", "DownPCT": "11.6 ", "InstallPCT": "8.84 ", "ValidInStates": "IL,IN,MD,TN,TX,VA"}, "Annual1DownSelectDueDate", 4, "12", "10", "11.6 ", "8.84 ", "IL,IN,MD,TN,TX,VA"}
	paymentPlanAnnual1AndHalfDown            = EnumPaymentPlanItem{paymentPlanAnnual1AndHalfDownID, "1.5 Months Down with Monthly Recurring", map[string]string{"TermLength": "12", "Installments": "10", "DownPCT": "12   ", "InstallPCT": "8.8  ", "ValidInStates": "IL,IN,MD,TN,TX,VA"}, "Annual1AndHalfDown", 5, "12", "10", "12   ", "8.8  ", "IL,IN,MD,TN,TX,VA"}
	paymentPlanAnnual2Down                   = EnumPaymentPlanItem{paymentPlanAnnual2DownID, "Monthly, Auto Pay", map[string]string{"TermLength": "12", "Installments": "10", "DownPCT": "16.7 ", "InstallPCT": "8.33 ", "ValidInStates": "IL,IN,MD,TN,TX,VA"}, "Annual2Down", 6, "12", "10", "16.7 ", "8.33 ", "IL,IN,MD,TN,TX,VA"}
	paymentPlanAnnual6Down6Payment           = EnumPaymentPlanItem{paymentPlanAnnual6Down6PaymentID, "6 Months Down with Monthly Payments", map[string]string{"TermLength": "12", "Installments": "6 ", "DownPCT": "50   ", "InstallPCT": "8.33 ", "ValidInStates": "IL,IN,MD,TN,TX,VA"}, "Annual6Down6Payment", 7, "12", "6 ", "50   ", "8.33 ", "IL,IN,MD,TN,TX,VA"}
	paymentPlanAnnual6Down1PaymentBillMe     = EnumPaymentPlanItem{paymentPlanAnnual6Down1PaymentBillMeID, "Two Payments, Bill Me", map[string]string{"TermLength": "12", "Installments": "1 ", "DownPCT": "50   ", "InstallPCT": "50   ", "ValidInStates": "IL,IN,MD,TN,TX,VA"}, "Annual6Down1PaymentBillMe", 8, "12", "1 ", "50   ", "50   ", "IL,IN,MD,TN,TX,VA"}
	paymentPlanAnnual6Down1PaymentAuto       = EnumPaymentPlanItem{paymentPlanAnnual6Down1PaymentAutoID, "Two Payments, Auto Pay", map[string]string{"TermLength": "12", "Installments": "1 ", "DownPCT": "50   ", "InstallPCT": "50   ", "ValidInStates": "IL,IN,MD,TN,TX,VA"}, "Annual6Down1PaymentAuto", 9, "12", "1 ", "50   ", "50   ", "IL,IN,MD,TN,TX,VA"}
	paymentPlanAnnualPayInFull               = EnumPaymentPlanItem{paymentPlanAnnualPayInFullID, "Pay In Full", map[string]string{"TermLength": "12", "Installments": "0 ", "DownPCT": "100  ", "InstallPCT": "0    ", "ValidInStates": "IL,IN,MD,TN,TX,VA"}, "AnnualPayInFull", 10, "12", "0 ", "100  ", "0    ", "IL,IN,MD,TN,TX,VA"}
	paymentPlanSemiHalfDown5Payment          = EnumPaymentPlanItem{paymentPlanSemiHalfDown5PaymentID, "Half Month Down - 6M", map[string]string{"TermLength": "6 ", "Installments": "5 ", "DownPCT": "8.35 ", "InstallPCT": "18.33", "ValidInStates": "GA,IL,IN,MD,TN,TX,VA"}, "SemiHalfDown5Payment", 11, "6 ", "5 ", "8.35 ", "18.33", "GA,IL,IN,MD,TN,TX,VA"}
	paymentPlanSemi1Down                     = EnumPaymentPlanItem{paymentPlanSemi1DownID, "1 Month Down - 6M", map[string]string{"TermLength": "6 ", "Installments": "5 ", "DownPCT": "16.75", "InstallPCT": "16.65", "ValidInStates": "GA,IL,IN,MD,OH,TN,TX,VA"}, "Semi1Down", 12, "6 ", "5 ", "16.75", "16.65", "GA,IL,IN,MD,OH,TN,TX,VA"}
	paymentPlanSemi1DownSelectDueDate        = EnumPaymentPlanItem{paymentPlanSemi1DownSelectDueDateID, "1 Month Down - 6M - Select Due Date", map[string]string{"TermLength": "6 ", "Installments": "5 ", "DownPCT": "25   ", "InstallPCT": "15   ", "ValidInStates": "GA,IL,IN,MD,OH,TN,TX,VA"}, "Semi1DownSelectDueDate", 13, "6 ", "5 ", "25   ", "15   ", "GA,IL,IN,MD,OH,TN,TX,VA"}
	paymentPlanSemi2Down                     = EnumPaymentPlanItem{paymentPlanSemi2DownID, "2 Month Down - 6M", map[string]string{"TermLength": "6 ", "Installments": "4 ", "DownPCT": "33.36", "InstallPCT": "16.66", "ValidInStates": "GA,IL,IN,MD,OH,TN,TX,VA"}, "Semi2Down", 14, "6 ", "4 ", "33.36", "16.66", "GA,IL,IN,MD,OH,TN,TX,VA"}
	paymentPlanSemi3Down1Payment             = EnumPaymentPlanItem{paymentPlanSemi3Down1PaymentID, "Two Pay - 6M", map[string]string{"TermLength": "6 ", "Installments": "1 ", "DownPCT": "50   ", "InstallPCT": "50   ", "ValidInStates": "GA,IL,IN,MD,OH,TN,TX,VA"}, "Semi3Down1Payment", 15, "6 ", "1 ", "50   ", "50   ", "GA,IL,IN,MD,OH,TN,TX,VA"}
	paymentPlanSemiPayInFull                 = EnumPaymentPlanItem{paymentPlanSemiPayInFullID, "Pay In Full - 6M", map[string]string{"TermLength": "6 ", "Installments": "0 ", "DownPCT": "100  ", "InstallPCT": "0    ", "ValidInStates": "GA,IL,IN,MD,OH,TN,TX,VA"}, "SemiPayInFull", 16, "6 ", "0 ", "100  ", "0    ", "GA,IL,IN,MD,OH,TN,TX,VA"}
	paymentPlanSemi1DownSelectDefaultDueDate = EnumPaymentPlanItem{paymentPlanSemi1DownSelectDefaultDueDateID, "1M Default Due Date - 6M", map[string]string{"TermLength": "6 ", "Installments": "5 ", "DownPCT": "16.75", "InstallPCT": "16.65", "ValidInStates": "VA"}, "Semi1DownSelectDefaultDueDate", 17, "6 ", "5 ", "16.75", "16.65", "VA"}
	paymentPlanSemi2DownDefaultDueDate       = EnumPaymentPlanItem{paymentPlanSemi2DownDefaultDueDateID, "2M Default Due Date - 6M", map[string]string{"TermLength": "6 ", "Installments": "4 ", "DownPCT": "33.36", "InstallPCT": "16.66", "ValidInStates": "VA"}, "Semi2DownDefaultDueDate", 18, "6 ", "4 ", "33.36", "16.66", "VA"}
)

// EnumPaymentPlan is a collection of PaymentPlan items
type EnumPaymentPlan struct {
	Description string
	Items       []*EnumPaymentPlanItem
	Name        string

	AnnualRenewalCatchUp          *EnumPaymentPlanItem
	AnnualRenewalStandard         *EnumPaymentPlanItem
	Annual1Down                   *EnumPaymentPlanItem
	Annual1DownSelectDueDate      *EnumPaymentPlanItem
	Annual1AndHalfDown            *EnumPaymentPlanItem
	Annual2Down                   *EnumPaymentPlanItem
	Annual6Down6Payment           *EnumPaymentPlanItem
	Annual6Down1PaymentBillMe     *EnumPaymentPlanItem
	Annual6Down1PaymentAuto       *EnumPaymentPlanItem
	AnnualPayInFull               *EnumPaymentPlanItem
	SemiHalfDown5Payment          *EnumPaymentPlanItem
	Semi1Down                     *EnumPaymentPlanItem
	Semi1DownSelectDueDate        *EnumPaymentPlanItem
	Semi2Down                     *EnumPaymentPlanItem
	Semi3Down1Payment             *EnumPaymentPlanItem
	SemiPayInFull                 *EnumPaymentPlanItem
	Semi1DownSelectDefaultDueDate *EnumPaymentPlanItem
	Semi2DownDefaultDueDate       *EnumPaymentPlanItem

	itemDict map[string]*EnumPaymentPlanItem
}

// PaymentPlan is a public singleton instance of EnumPaymentPlan
// representing payment plans
var PaymentPlan = &EnumPaymentPlan{
	Description: "payment plans",
	Items: []*EnumPaymentPlanItem{
		&paymentPlanAnnualRenewalCatchUp,
		&paymentPlanAnnualRenewalStandard,
		&paymentPlanAnnual1Down,
		&paymentPlanAnnual1DownSelectDueDate,
		&paymentPlanAnnual1AndHalfDown,
		&paymentPlanAnnual2Down,
		&paymentPlanAnnual6Down6Payment,
		&paymentPlanAnnual6Down1PaymentBillMe,
		&paymentPlanAnnual6Down1PaymentAuto,
		&paymentPlanAnnualPayInFull,
		&paymentPlanSemiHalfDown5Payment,
		&paymentPlanSemi1Down,
		&paymentPlanSemi1DownSelectDueDate,
		&paymentPlanSemi2Down,
		&paymentPlanSemi3Down1Payment,
		&paymentPlanSemiPayInFull,
		&paymentPlanSemi1DownSelectDefaultDueDate,
		&paymentPlanSemi2DownDefaultDueDate,
	},
	Name:                          "EnumPaymentPlan",
	AnnualRenewalCatchUp:          &paymentPlanAnnualRenewalCatchUp,
	AnnualRenewalStandard:         &paymentPlanAnnualRenewalStandard,
	Annual1Down:                   &paymentPlanAnnual1Down,
	Annual1DownSelectDueDate:      &paymentPlanAnnual1DownSelectDueDate,
	Annual1AndHalfDown:            &paymentPlanAnnual1AndHalfDown,
	Annual2Down:                   &paymentPlanAnnual2Down,
	Annual6Down6Payment:           &paymentPlanAnnual6Down6Payment,
	Annual6Down1PaymentBillMe:     &paymentPlanAnnual6Down1PaymentBillMe,
	Annual6Down1PaymentAuto:       &paymentPlanAnnual6Down1PaymentAuto,
	AnnualPayInFull:               &paymentPlanAnnualPayInFull,
	SemiHalfDown5Payment:          &paymentPlanSemiHalfDown5Payment,
	Semi1Down:                     &paymentPlanSemi1Down,
	Semi1DownSelectDueDate:        &paymentPlanSemi1DownSelectDueDate,
	Semi2Down:                     &paymentPlanSemi2Down,
	Semi3Down1Payment:             &paymentPlanSemi3Down1Payment,
	SemiPayInFull:                 &paymentPlanSemiPayInFull,
	Semi1DownSelectDefaultDueDate: &paymentPlanSemi1DownSelectDefaultDueDate,
	Semi2DownDefaultDueDate:       &paymentPlanSemi2DownDefaultDueDate,

	itemDict: map[string]*EnumPaymentPlanItem{
		strings.ToLower(string(paymentPlanAnnualRenewalCatchUpID)):          &paymentPlanAnnualRenewalCatchUp,
		strings.ToLower(string(paymentPlanAnnualRenewalStandardID)):         &paymentPlanAnnualRenewalStandard,
		strings.ToLower(string(paymentPlanAnnual1DownID)):                   &paymentPlanAnnual1Down,
		strings.ToLower(string(paymentPlanAnnual1DownSelectDueDateID)):      &paymentPlanAnnual1DownSelectDueDate,
		strings.ToLower(string(paymentPlanAnnual1AndHalfDownID)):            &paymentPlanAnnual1AndHalfDown,
		strings.ToLower(string(paymentPlanAnnual2DownID)):                   &paymentPlanAnnual2Down,
		strings.ToLower(string(paymentPlanAnnual6Down6PaymentID)):           &paymentPlanAnnual6Down6Payment,
		strings.ToLower(string(paymentPlanAnnual6Down1PaymentBillMeID)):     &paymentPlanAnnual6Down1PaymentBillMe,
		strings.ToLower(string(paymentPlanAnnual6Down1PaymentAutoID)):       &paymentPlanAnnual6Down1PaymentAuto,
		strings.ToLower(string(paymentPlanAnnualPayInFullID)):               &paymentPlanAnnualPayInFull,
		strings.ToLower(string(paymentPlanSemiHalfDown5PaymentID)):          &paymentPlanSemiHalfDown5Payment,
		strings.ToLower(string(paymentPlanSemi1DownID)):                     &paymentPlanSemi1Down,
		strings.ToLower(string(paymentPlanSemi1DownSelectDueDateID)):        &paymentPlanSemi1DownSelectDueDate,
		strings.ToLower(string(paymentPlanSemi2DownID)):                     &paymentPlanSemi2Down,
		strings.ToLower(string(paymentPlanSemi3Down1PaymentID)):             &paymentPlanSemi3Down1Payment,
		strings.ToLower(string(paymentPlanSemiPayInFullID)):                 &paymentPlanSemiPayInFull,
		strings.ToLower(string(paymentPlanSemi1DownSelectDefaultDueDateID)): &paymentPlanSemi1DownSelectDefaultDueDate,
		strings.ToLower(string(paymentPlanSemi2DownDefaultDueDateID)):       &paymentPlanSemi2DownDefaultDueDate,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumPaymentPlan) ByID(id PaymentPlanIdentifier) *EnumPaymentPlanItem {
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
func (e *EnumPaymentPlan) ByIDString(idx string) *EnumPaymentPlanItem {
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
func (e *EnumPaymentPlan) ByIndex(idx int) *EnumPaymentPlanItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedPaymentPlanID is a struct that is designed to replace a *PaymentPlanID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *PaymentPlanID it contains while being a better JSON citizen.
type ValidatedPaymentPlanID struct {
	// id will point to a valid PaymentPlanID, if possible
	// If id is nil, then ValidatedPaymentPlanID.Valid() will return false.
	id *PaymentPlanID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedPaymentPlanID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedPaymentPlanID
func (vi *ValidatedPaymentPlanID) Clone() *ValidatedPaymentPlanID {
	if vi == nil {
		return nil
	}

	var cid *PaymentPlanID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedPaymentPlanID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedPaymentPlanIds represent the same PaymentPlan
func (vi *ValidatedPaymentPlanID) Equals(vj *ValidatedPaymentPlanID) bool {
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

// Valid returns true if and only if the ValidatedPaymentPlanID corresponds to a recognized PaymentPlan
func (vi *ValidatedPaymentPlanID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedPaymentPlanID) ID() *PaymentPlanID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedPaymentPlanID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedPaymentPlanID) ValidatedID() *ValidatedPaymentPlanID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedPaymentPlanID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedPaymentPlanID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedPaymentPlanID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedPaymentPlanID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedPaymentPlanID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := PaymentPlanID(capString)
	item := PaymentPlan.ByID(&id)
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

func (vi ValidatedPaymentPlanID) String() string {
	return vi.ToIDString()
}

type PaymentPlanIdentifier interface {
	ID() *PaymentPlanID
	Valid() bool
}
