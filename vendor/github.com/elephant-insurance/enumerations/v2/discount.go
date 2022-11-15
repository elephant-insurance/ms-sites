package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// DiscountID uniquely identifies a particular Discount
type DiscountID string

// Clone creates a safe, independent copy of a DiscountID
func (i *DiscountID) Clone() *DiscountID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two DiscountIds are equivalent
func (i *DiscountID) Equals(j *DiscountID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *DiscountID that is either valid or nil
func (i *DiscountID) ID() *DiscountID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *DiscountID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the DiscountID corresponds to a recognized Discount
func (i *DiscountID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return Discount.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *DiscountID) ValidatedID() *ValidatedDiscountID {
	if i != nil {
		return &ValidatedDiscountID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *DiscountID) MarshalJSON() ([]byte, error) {
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

func (i *DiscountID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := DiscountID(dataString)
	item := Discount.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	discountEarlyBirdID           DiscountID = "earlybird"
	discountElectronicSignatureID DiscountID = "electronicsignature"
	discountGoodStudentID         DiscountID = "goodstudent"
	discountHomeOwnerID           DiscountID = "homeowner"
	discountMultiCarID            DiscountID = "multicar"
	discountOnlineID              DiscountID = "online"
	discountPaperlessID           DiscountID = "paperless"
	discountReferralID            DiscountID = "referralCode"
	discountResponsibleDriverID   DiscountID = "responsibledriver"
	discountSafetyFeatureID       DiscountID = "safetyfeature"
	discountStudentAwayID         DiscountID = "studentaway"
	discountWorkFromHomeID        DiscountID = "workFromHome"
	discountAntiTheftID           DiscountID = "antitheft"
	discountCompareID             DiscountID = "compare"
	discountMatureDriverID        DiscountID = "maturedriver"
	discountMultiPolicyID         DiscountID = "multipolicy"
	discountHomeRentersID         DiscountID = "homerenters"
	discountLifeMotoID            DiscountID = "lifemoto"
	discountPIFID                 DiscountID = "pif"
	discountReapplicationID       DiscountID = "reapplication"
	discountSafeCarID             DiscountID = "safecar"
	discountAdmiralEmployeeID     DiscountID = "admiralemployee"
	discountAceableID             DiscountID = "aceable"
)

// EnumDiscountItem describes an entry in an enumeration of Discount
type EnumDiscountItem struct {
	ID        DiscountID        `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	StateCodes string
}

var (
	discountEarlyBird           = EnumDiscountItem{discountEarlyBirdID, "Early Bird Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "EarlyBird", 1, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountElectronicSignature = EnumDiscountItem{discountElectronicSignatureID, "Electronic Signature Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "ElectronicSignature", 2, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountGoodStudent         = EnumDiscountItem{discountGoodStudentID, "Good Student Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "GoodStudent", 3, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountHomeOwner           = EnumDiscountItem{discountHomeOwnerID, "Homeowner Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "HomeOwner", 4, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountMultiCar            = EnumDiscountItem{discountMultiCarID, "Multi-Car Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "MultiCar", 5, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountOnline              = EnumDiscountItem{discountOnlineID, "Online Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "Online", 6, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountPaperless           = EnumDiscountItem{discountPaperlessID, "Paperless Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "Paperless", 7, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountReferral            = EnumDiscountItem{discountReferralID, "Referral Discount", map[string]string{"StateCodes": "VA, MD, TX, IL"}, "Referral", 8, "VA, MD, TX, IL"}
	discountResponsibleDriver   = EnumDiscountItem{discountResponsibleDriverID, "Responsible Driver Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "ResponsibleDriver", 9, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountSafetyFeature       = EnumDiscountItem{discountSafetyFeatureID, "Safety Feature Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "SafetyFeature", 10, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountStudentAway         = EnumDiscountItem{discountStudentAwayID, "Student Away Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "StudentAway", 11, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountWorkFromHome        = EnumDiscountItem{discountWorkFromHomeID, "Work From Home Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,OH"}, "WorkFromHome", 12, "IN,TN,IL,MD,TX,VA,OH"}
	discountAntiTheft           = EnumDiscountItem{discountAntiTheftID, "Anti-Theft Discount", map[string]string{"StateCodes": "IL"}, "AntiTheft", 13, "IL"}
	discountCompare             = EnumDiscountItem{discountCompareID, "Compare.com Discount", map[string]string{"StateCodes": "IL,TX"}, "Compare", 14, "IL,TX"}
	discountMatureDriver        = EnumDiscountItem{discountMatureDriverID, "Mature Driver Discount", map[string]string{"StateCodes": "VA,IL,TN,OH"}, "MatureDriver", 15, "VA,IL,TN,OH"}
	discountMultiPolicy         = EnumDiscountItem{discountMultiPolicyID, "Multi-Policy Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "MultiPolicy", 16, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountHomeRenters         = EnumDiscountItem{discountHomeRentersID, "Home/Renters Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "HomeRenters", 17, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountLifeMoto            = EnumDiscountItem{discountLifeMotoID, "Life/Moto Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "LifeMoto", 18, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountPIF                 = EnumDiscountItem{discountPIFID, "PIF Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "PIF", 19, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountReapplication       = EnumDiscountItem{discountReapplicationID, "Reapplication Discount", map[string]string{"StateCodes": "IN,TN,IL,TX,VA,GA,OH"}, "Reapplication", 20, "IN,TN,IL,TX,VA,GA,OH"}
	discountSafeCar             = EnumDiscountItem{discountSafeCarID, "Safe Car Discount", map[string]string{"StateCodes": "IN,TN,IL,MD,TX,VA,GA,OH"}, "SafeCar", 21, "IN,TN,IL,MD,TX,VA,GA,OH"}
	discountAdmiralEmployee     = EnumDiscountItem{discountAdmiralEmployeeID, "Admiral Employee Discount", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,GA,OH"}, "AdmiralEmployee", 22, "VA,TX,MD,IL,IN,GA,OH"}
	discountAceable             = EnumDiscountItem{discountAceableID, "Aceable Discount", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,GA,OH"}, "Aceable", 23, "VA,TX,MD,IL,IN,GA,OH"}
)

// EnumDiscount is a collection of Discount items
type EnumDiscount struct {
	Description string
	Items       []*EnumDiscountItem
	Name        string

	EarlyBird           *EnumDiscountItem
	ElectronicSignature *EnumDiscountItem
	GoodStudent         *EnumDiscountItem
	HomeOwner           *EnumDiscountItem
	MultiCar            *EnumDiscountItem
	Online              *EnumDiscountItem
	Paperless           *EnumDiscountItem
	Referral            *EnumDiscountItem
	ResponsibleDriver   *EnumDiscountItem
	SafetyFeature       *EnumDiscountItem
	StudentAway         *EnumDiscountItem
	WorkFromHome        *EnumDiscountItem
	AntiTheft           *EnumDiscountItem
	Compare             *EnumDiscountItem
	MatureDriver        *EnumDiscountItem
	MultiPolicy         *EnumDiscountItem
	HomeRenters         *EnumDiscountItem
	LifeMoto            *EnumDiscountItem
	PIF                 *EnumDiscountItem
	Reapplication       *EnumDiscountItem
	SafeCar             *EnumDiscountItem
	AdmiralEmployee     *EnumDiscountItem
	Aceable             *EnumDiscountItem

	itemDict map[string]*EnumDiscountItem
}

// Discount is a public singleton instance of EnumDiscount
// representing discounts available to our customers
var Discount = &EnumDiscount{
	Description: "discounts available to our customers",
	Items: []*EnumDiscountItem{
		&discountEarlyBird,
		&discountElectronicSignature,
		&discountGoodStudent,
		&discountHomeOwner,
		&discountMultiCar,
		&discountOnline,
		&discountPaperless,
		&discountReferral,
		&discountResponsibleDriver,
		&discountSafetyFeature,
		&discountStudentAway,
		&discountWorkFromHome,
		&discountAntiTheft,
		&discountCompare,
		&discountMatureDriver,
		&discountMultiPolicy,
		&discountHomeRenters,
		&discountLifeMoto,
		&discountPIF,
		&discountReapplication,
		&discountSafeCar,
		&discountAdmiralEmployee,
		&discountAceable,
	},
	Name:                "EnumDiscount",
	EarlyBird:           &discountEarlyBird,
	ElectronicSignature: &discountElectronicSignature,
	GoodStudent:         &discountGoodStudent,
	HomeOwner:           &discountHomeOwner,
	MultiCar:            &discountMultiCar,
	Online:              &discountOnline,
	Paperless:           &discountPaperless,
	Referral:            &discountReferral,
	ResponsibleDriver:   &discountResponsibleDriver,
	SafetyFeature:       &discountSafetyFeature,
	StudentAway:         &discountStudentAway,
	WorkFromHome:        &discountWorkFromHome,
	AntiTheft:           &discountAntiTheft,
	Compare:             &discountCompare,
	MatureDriver:        &discountMatureDriver,
	MultiPolicy:         &discountMultiPolicy,
	HomeRenters:         &discountHomeRenters,
	LifeMoto:            &discountLifeMoto,
	PIF:                 &discountPIF,
	Reapplication:       &discountReapplication,
	SafeCar:             &discountSafeCar,
	AdmiralEmployee:     &discountAdmiralEmployee,
	Aceable:             &discountAceable,

	itemDict: map[string]*EnumDiscountItem{
		strings.ToLower(string(discountEarlyBirdID)):           &discountEarlyBird,
		strings.ToLower(string(discountElectronicSignatureID)): &discountElectronicSignature,
		strings.ToLower(string(discountGoodStudentID)):         &discountGoodStudent,
		strings.ToLower(string(discountHomeOwnerID)):           &discountHomeOwner,
		strings.ToLower(string(discountMultiCarID)):            &discountMultiCar,
		strings.ToLower(string(discountOnlineID)):              &discountOnline,
		strings.ToLower(string(discountPaperlessID)):           &discountPaperless,
		strings.ToLower(string(discountReferralID)):            &discountReferral,
		strings.ToLower(string(discountResponsibleDriverID)):   &discountResponsibleDriver,
		strings.ToLower(string(discountSafetyFeatureID)):       &discountSafetyFeature,
		strings.ToLower(string(discountStudentAwayID)):         &discountStudentAway,
		strings.ToLower(string(discountWorkFromHomeID)):        &discountWorkFromHome,
		strings.ToLower(string(discountAntiTheftID)):           &discountAntiTheft,
		strings.ToLower(string(discountCompareID)):             &discountCompare,
		strings.ToLower(string(discountMatureDriverID)):        &discountMatureDriver,
		strings.ToLower(string(discountMultiPolicyID)):         &discountMultiPolicy,
		strings.ToLower(string(discountHomeRentersID)):         &discountHomeRenters,
		strings.ToLower(string(discountLifeMotoID)):            &discountLifeMoto,
		strings.ToLower(string(discountPIFID)):                 &discountPIF,
		strings.ToLower(string(discountReapplicationID)):       &discountReapplication,
		strings.ToLower(string(discountSafeCarID)):             &discountSafeCar,
		strings.ToLower(string(discountAdmiralEmployeeID)):     &discountAdmiralEmployee,
		strings.ToLower(string(discountAceableID)):             &discountAceable,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumDiscount) ByID(id DiscountIdentifier) *EnumDiscountItem {
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
func (e *EnumDiscount) ByIDString(idx string) *EnumDiscountItem {
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
func (e *EnumDiscount) ByIndex(idx int) *EnumDiscountItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedDiscountID is a struct that is designed to replace a *DiscountID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *DiscountID it contains while being a better JSON citizen.
type ValidatedDiscountID struct {
	// id will point to a valid DiscountID, if possible
	// If id is nil, then ValidatedDiscountID.Valid() will return false.
	id *DiscountID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedDiscountID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedDiscountID
func (vi *ValidatedDiscountID) Clone() *ValidatedDiscountID {
	if vi == nil {
		return nil
	}

	var cid *DiscountID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedDiscountID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedDiscountIds represent the same Discount
func (vi *ValidatedDiscountID) Equals(vj *ValidatedDiscountID) bool {
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

// Valid returns true if and only if the ValidatedDiscountID corresponds to a recognized Discount
func (vi *ValidatedDiscountID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedDiscountID) ID() *DiscountID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedDiscountID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedDiscountID) ValidatedID() *ValidatedDiscountID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedDiscountID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedDiscountID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedDiscountID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedDiscountID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedDiscountID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := DiscountID(capString)
	item := Discount.ByID(&id)
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

func (vi ValidatedDiscountID) String() string {
	return vi.ToIDString()
}

type DiscountIdentifier interface {
	ID() *DiscountID
	Valid() bool
}
