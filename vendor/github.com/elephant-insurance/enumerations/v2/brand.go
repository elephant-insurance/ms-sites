package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// BrandID uniquely identifies a particular Brand
type BrandID string

// Clone creates a safe, independent copy of a BrandID
func (i *BrandID) Clone() *BrandID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two BrandIds are equivalent
func (i *BrandID) Equals(j *BrandID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *BrandID that is either valid or nil
func (i *BrandID) ID() *BrandID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *BrandID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the BrandID corresponds to a recognized Brand
func (i *BrandID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return Brand.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *BrandID) ValidatedID() *ValidatedBrandID {
	if i != nil {
		return &ValidatedBrandID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *BrandID) MarshalJSON() ([]byte, error) {
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

func (i *BrandID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := BrandID(dataString)
	item := Brand.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	brandElephantID BrandID = "elephant"
	brandApparentID BrandID = "apparent"
)

// EnumBrandItem describes an entry in an enumeration of Brand
type EnumBrandItem struct {
	ID        BrandID           `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	brandElephant = EnumBrandItem{brandElephantID, "Elephant Auto Insurance", nil, "Elephant", 1}
	brandApparent = EnumBrandItem{brandApparentID, "Apparent Auto Insurance", nil, "Apparent", 2}
)

// EnumBrand is a collection of Brand items
type EnumBrand struct {
	Description string
	Items       []*EnumBrandItem
	Name        string

	Elephant *EnumBrandItem
	Apparent *EnumBrandItem

	itemDict map[string]*EnumBrandItem
}

// Brand is a public singleton instance of EnumBrand
// representing brands of Elephant insurance
var Brand = &EnumBrand{
	Description: "brands of Elephant insurance",
	Items: []*EnumBrandItem{
		&brandElephant,
		&brandApparent,
	},
	Name:     "EnumBrand",
	Elephant: &brandElephant,
	Apparent: &brandApparent,

	itemDict: map[string]*EnumBrandItem{
		strings.ToLower(string(brandElephantID)): &brandElephant,
		strings.ToLower(string(brandApparentID)): &brandApparent,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumBrand) ByID(id BrandIdentifier) *EnumBrandItem {
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
func (e *EnumBrand) ByIDString(idx string) *EnumBrandItem {
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
func (e *EnumBrand) ByIndex(idx int) *EnumBrandItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedBrandID is a struct that is designed to replace a *BrandID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *BrandID it contains while being a better JSON citizen.
type ValidatedBrandID struct {
	// id will point to a valid BrandID, if possible
	// If id is nil, then ValidatedBrandID.Valid() will return false.
	id *BrandID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedBrandID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedBrandID
func (vi *ValidatedBrandID) Clone() *ValidatedBrandID {
	if vi == nil {
		return nil
	}

	var cid *BrandID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedBrandID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedBrandIds represent the same Brand
func (vi *ValidatedBrandID) Equals(vj *ValidatedBrandID) bool {
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

// Valid returns true if and only if the ValidatedBrandID corresponds to a recognized Brand
func (vi *ValidatedBrandID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedBrandID) ID() *BrandID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedBrandID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedBrandID) ValidatedID() *ValidatedBrandID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedBrandID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedBrandID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedBrandID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedBrandID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedBrandID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := BrandID(capString)
	item := Brand.ByID(&id)
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

func (vi ValidatedBrandID) String() string {
	return vi.ToIDString()
}

type BrandIdentifier interface {
	ID() *BrandID
	Valid() bool
}
