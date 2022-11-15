package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// CreditRatingID uniquely identifies a particular CreditRating
type CreditRatingID string

// Clone creates a safe, independent copy of a CreditRatingID
func (i *CreditRatingID) Clone() *CreditRatingID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two CreditRatingIds are equivalent
func (i *CreditRatingID) Equals(j *CreditRatingID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *CreditRatingID that is either valid or nil
func (i *CreditRatingID) ID() *CreditRatingID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *CreditRatingID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the CreditRatingID corresponds to a recognized CreditRating
func (i *CreditRatingID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return CreditRating.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *CreditRatingID) ValidatedID() *ValidatedCreditRatingID {
	if i != nil {
		return &ValidatedCreditRatingID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *CreditRatingID) MarshalJSON() ([]byte, error) {
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

func (i *CreditRatingID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := CreditRatingID(dataString)
	item := CreditRating.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	creditRatingPoorID      CreditRatingID = "poor"
	creditRatingFairID      CreditRatingID = "fair"
	creditRatingGoodID      CreditRatingID = "good"
	creditRatingVeryGoodID  CreditRatingID = "verygood"
	creditRatingExcellentID CreditRatingID = "excellent"
)

// EnumCreditRatingItem describes an entry in an enumeration of CreditRating
type EnumCreditRatingItem struct {
	ID        CreditRatingID    `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	creditRatingPoor      = EnumCreditRatingItem{creditRatingPoorID, "Poor", nil, "Poor", 1}
	creditRatingFair      = EnumCreditRatingItem{creditRatingFairID, "Fair", nil, "Fair", 2}
	creditRatingGood      = EnumCreditRatingItem{creditRatingGoodID, "Good", nil, "Good", 3}
	creditRatingVeryGood  = EnumCreditRatingItem{creditRatingVeryGoodID, "Very Good", nil, "VeryGood", 4}
	creditRatingExcellent = EnumCreditRatingItem{creditRatingExcellentID, "Excellent", nil, "Excellent", 5}
)

// EnumCreditRating is a collection of CreditRating items
type EnumCreditRating struct {
	Description string
	Items       []*EnumCreditRatingItem
	Name        string

	Poor      *EnumCreditRatingItem
	Fair      *EnumCreditRatingItem
	Good      *EnumCreditRatingItem
	VeryGood  *EnumCreditRatingItem
	Excellent *EnumCreditRatingItem

	itemDict map[string]*EnumCreditRatingItem
}

// CreditRating is a public singleton instance of EnumCreditRating
// representing credit rating levels
var CreditRating = &EnumCreditRating{
	Description: "credit rating levels",
	Items: []*EnumCreditRatingItem{
		&creditRatingPoor,
		&creditRatingFair,
		&creditRatingGood,
		&creditRatingVeryGood,
		&creditRatingExcellent,
	},
	Name:      "EnumCreditRating",
	Poor:      &creditRatingPoor,
	Fair:      &creditRatingFair,
	Good:      &creditRatingGood,
	VeryGood:  &creditRatingVeryGood,
	Excellent: &creditRatingExcellent,

	itemDict: map[string]*EnumCreditRatingItem{
		strings.ToLower(string(creditRatingPoorID)):      &creditRatingPoor,
		strings.ToLower(string(creditRatingFairID)):      &creditRatingFair,
		strings.ToLower(string(creditRatingGoodID)):      &creditRatingGood,
		strings.ToLower(string(creditRatingVeryGoodID)):  &creditRatingVeryGood,
		strings.ToLower(string(creditRatingExcellentID)): &creditRatingExcellent,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumCreditRating) ByID(id CreditRatingIdentifier) *EnumCreditRatingItem {
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
func (e *EnumCreditRating) ByIDString(idx string) *EnumCreditRatingItem {
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
func (e *EnumCreditRating) ByIndex(idx int) *EnumCreditRatingItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedCreditRatingID is a struct that is designed to replace a *CreditRatingID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *CreditRatingID it contains while being a better JSON citizen.
type ValidatedCreditRatingID struct {
	// id will point to a valid CreditRatingID, if possible
	// If id is nil, then ValidatedCreditRatingID.Valid() will return false.
	id *CreditRatingID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedCreditRatingID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedCreditRatingID
func (vi *ValidatedCreditRatingID) Clone() *ValidatedCreditRatingID {
	if vi == nil {
		return nil
	}

	var cid *CreditRatingID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedCreditRatingID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedCreditRatingIds represent the same CreditRating
func (vi *ValidatedCreditRatingID) Equals(vj *ValidatedCreditRatingID) bool {
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

// Valid returns true if and only if the ValidatedCreditRatingID corresponds to a recognized CreditRating
func (vi *ValidatedCreditRatingID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedCreditRatingID) ID() *CreditRatingID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedCreditRatingID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedCreditRatingID) ValidatedID() *ValidatedCreditRatingID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedCreditRatingID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedCreditRatingID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedCreditRatingID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedCreditRatingID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedCreditRatingID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := CreditRatingID(capString)
	item := CreditRating.ByID(&id)
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

func (vi ValidatedCreditRatingID) String() string {
	return vi.ToIDString()
}

type CreditRatingIdentifier interface {
	ID() *CreditRatingID
	Valid() bool
}
