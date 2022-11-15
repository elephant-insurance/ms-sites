package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// InsuranceLapseID uniquely identifies a particular InsuranceLapse
type InsuranceLapseID string

// Clone creates a safe, independent copy of a InsuranceLapseID
func (i *InsuranceLapseID) Clone() *InsuranceLapseID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two InsuranceLapseIds are equivalent
func (i *InsuranceLapseID) Equals(j *InsuranceLapseID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *InsuranceLapseID that is either valid or nil
func (i *InsuranceLapseID) ID() *InsuranceLapseID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *InsuranceLapseID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the InsuranceLapseID corresponds to a recognized InsuranceLapse
func (i *InsuranceLapseID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return InsuranceLapse.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *InsuranceLapseID) ValidatedID() *ValidatedInsuranceLapseID {
	if i != nil {
		return &ValidatedInsuranceLapseID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *InsuranceLapseID) MarshalJSON() ([]byte, error) {
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

func (i *InsuranceLapseID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := InsuranceLapseID(dataString)
	item := InsuranceLapse.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	insuranceLapseNoLapseID           InsuranceLapseID = "no_lapse"
	insuranceLapseYes30DaysOrLessID   InsuranceLapseID = "lapse_within_30days"
	insuranceLapseYesMoreThan30DaysID InsuranceLapseID = "lapse_more_than_30days"
)

// EnumInsuranceLapseItem describes an entry in an enumeration of InsuranceLapse
type EnumInsuranceLapseItem struct {
	ID        InsuranceLapseID  `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	insuranceLapseNoLapse           = EnumInsuranceLapseItem{insuranceLapseNoLapseID, "No", nil, "NoLapse", 1}
	insuranceLapseYes30DaysOrLess   = EnumInsuranceLapseItem{insuranceLapseYes30DaysOrLessID, "Yes, 30 Days or Less", nil, "Yes30DaysOrLess", 2}
	insuranceLapseYesMoreThan30Days = EnumInsuranceLapseItem{insuranceLapseYesMoreThan30DaysID, "Yes, More than 30 Days", nil, "YesMoreThan30Days", 3}
)

// EnumInsuranceLapse is a collection of InsuranceLapse items
type EnumInsuranceLapse struct {
	Description string
	Items       []*EnumInsuranceLapseItem
	Name        string

	NoLapse           *EnumInsuranceLapseItem
	Yes30DaysOrLess   *EnumInsuranceLapseItem
	YesMoreThan30Days *EnumInsuranceLapseItem

	itemDict map[string]*EnumInsuranceLapseItem
}

// InsuranceLapse is a public singleton instance of EnumInsuranceLapse
// representing lapses in insurance coverage
var InsuranceLapse = &EnumInsuranceLapse{
	Description: "lapses in insurance coverage",
	Items: []*EnumInsuranceLapseItem{
		&insuranceLapseNoLapse,
		&insuranceLapseYes30DaysOrLess,
		&insuranceLapseYesMoreThan30Days,
	},
	Name:              "EnumInsuranceLapse",
	NoLapse:           &insuranceLapseNoLapse,
	Yes30DaysOrLess:   &insuranceLapseYes30DaysOrLess,
	YesMoreThan30Days: &insuranceLapseYesMoreThan30Days,

	itemDict: map[string]*EnumInsuranceLapseItem{
		strings.ToLower(string(insuranceLapseNoLapseID)):           &insuranceLapseNoLapse,
		strings.ToLower(string(insuranceLapseYes30DaysOrLessID)):   &insuranceLapseYes30DaysOrLess,
		strings.ToLower(string(insuranceLapseYesMoreThan30DaysID)): &insuranceLapseYesMoreThan30Days,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumInsuranceLapse) ByID(id InsuranceLapseIdentifier) *EnumInsuranceLapseItem {
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
func (e *EnumInsuranceLapse) ByIDString(idx string) *EnumInsuranceLapseItem {
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
func (e *EnumInsuranceLapse) ByIndex(idx int) *EnumInsuranceLapseItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedInsuranceLapseID is a struct that is designed to replace a *InsuranceLapseID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *InsuranceLapseID it contains while being a better JSON citizen.
type ValidatedInsuranceLapseID struct {
	// id will point to a valid InsuranceLapseID, if possible
	// If id is nil, then ValidatedInsuranceLapseID.Valid() will return false.
	id *InsuranceLapseID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedInsuranceLapseID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedInsuranceLapseID
func (vi *ValidatedInsuranceLapseID) Clone() *ValidatedInsuranceLapseID {
	if vi == nil {
		return nil
	}

	var cid *InsuranceLapseID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedInsuranceLapseID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedInsuranceLapseIds represent the same InsuranceLapse
func (vi *ValidatedInsuranceLapseID) Equals(vj *ValidatedInsuranceLapseID) bool {
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

// Valid returns true if and only if the ValidatedInsuranceLapseID corresponds to a recognized InsuranceLapse
func (vi *ValidatedInsuranceLapseID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedInsuranceLapseID) ID() *InsuranceLapseID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedInsuranceLapseID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedInsuranceLapseID) ValidatedID() *ValidatedInsuranceLapseID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedInsuranceLapseID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedInsuranceLapseID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedInsuranceLapseID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedInsuranceLapseID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedInsuranceLapseID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := InsuranceLapseID(capString)
	item := InsuranceLapse.ByID(&id)
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

func (vi ValidatedInsuranceLapseID) String() string {
	return vi.ToIDString()
}

type InsuranceLapseIdentifier interface {
	ID() *InsuranceLapseID
	Valid() bool
}
