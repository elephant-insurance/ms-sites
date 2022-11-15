package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// RelationshipToInsuredID uniquely identifies a particular RelationshipToInsured
type RelationshipToInsuredID string

// Clone creates a safe, independent copy of a RelationshipToInsuredID
func (i *RelationshipToInsuredID) Clone() *RelationshipToInsuredID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two RelationshipToInsuredIds are equivalent
func (i *RelationshipToInsuredID) Equals(j *RelationshipToInsuredID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *RelationshipToInsuredID that is either valid or nil
func (i *RelationshipToInsuredID) ID() *RelationshipToInsuredID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *RelationshipToInsuredID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the RelationshipToInsuredID corresponds to a recognized RelationshipToInsured
func (i *RelationshipToInsuredID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return RelationshipToInsured.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *RelationshipToInsuredID) ValidatedID() *ValidatedRelationshipToInsuredID {
	if i != nil {
		return &ValidatedRelationshipToInsuredID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *RelationshipToInsuredID) MarshalJSON() ([]byte, error) {
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

func (i *RelationshipToInsuredID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := RelationshipToInsuredID(dataString)
	item := RelationshipToInsured.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	relationshipToInsuredApplicantID         RelationshipToInsuredID = "applicant"
	relationshipToInsuredSpouseID            RelationshipToInsuredID = "spouse"
	relationshipToInsuredChildID             RelationshipToInsuredID = "child"
	relationshipToInsuredParentID            RelationshipToInsuredID = "parent"
	relationshipToInsuredOtherRelativeID     RelationshipToInsuredID = "otherrelative"
	relationshipToInsuredNonResidentDriverID RelationshipToInsuredID = "nonresidentdriver"
	relationshipToInsuredNoneOfTheAboveID    RelationshipToInsuredID = "noneoftheabove"
)

// EnumRelationshipToInsuredItem describes an entry in an enumeration of RelationshipToInsured
type EnumRelationshipToInsuredItem struct {
	ID        RelationshipToInsuredID `json:"Value"`
	Desc      string                  `json:"Description,omitempty"`
	Meta      map[string]string       `json:",omitempty"`
	Name      string                  `json:"Name"`
	SortOrder int

	// Meta Properties
	AllowedWeb string
}

var (
	relationshipToInsuredApplicant         = EnumRelationshipToInsuredItem{relationshipToInsuredApplicantID, "Self", map[string]string{"AllowedWeb": "false"}, "Applicant", 1, "false"}
	relationshipToInsuredSpouse            = EnumRelationshipToInsuredItem{relationshipToInsuredSpouseID, "Spouse", map[string]string{"AllowedWeb": "false"}, "Spouse", 2, "false"}
	relationshipToInsuredChild             = EnumRelationshipToInsuredItem{relationshipToInsuredChildID, "Child", map[string]string{"AllowedWeb": "true"}, "Child", 3, "true"}
	relationshipToInsuredParent            = EnumRelationshipToInsuredItem{relationshipToInsuredParentID, "Parent", map[string]string{"AllowedWeb": "true"}, "Parent", 4, "true"}
	relationshipToInsuredOtherRelative     = EnumRelationshipToInsuredItem{relationshipToInsuredOtherRelativeID, "Other Relative", map[string]string{"AllowedWeb": "true"}, "OtherRelative", 5, "true"}
	relationshipToInsuredNonResidentDriver = EnumRelationshipToInsuredItem{relationshipToInsuredNonResidentDriverID, "Non resident driver", map[string]string{"AllowedWeb": "false"}, "NonResidentDriver", 6, "false"}
	relationshipToInsuredNoneOfTheAbove    = EnumRelationshipToInsuredItem{relationshipToInsuredNoneOfTheAboveID, "Other Non-Relative", map[string]string{"AllowedWeb": "true"}, "NoneOfTheAbove", 7, "true"}
)

// EnumRelationshipToInsured is a collection of RelationshipToInsured items
type EnumRelationshipToInsured struct {
	Description string
	Items       []*EnumRelationshipToInsuredItem
	Name        string

	Applicant         *EnumRelationshipToInsuredItem
	Spouse            *EnumRelationshipToInsuredItem
	Child             *EnumRelationshipToInsuredItem
	Parent            *EnumRelationshipToInsuredItem
	OtherRelative     *EnumRelationshipToInsuredItem
	NonResidentDriver *EnumRelationshipToInsuredItem
	NoneOfTheAbove    *EnumRelationshipToInsuredItem

	itemDict map[string]*EnumRelationshipToInsuredItem
}

// RelationshipToInsured is a public singleton instance of EnumRelationshipToInsured
// representing relationships of drivers to the primary insured
var RelationshipToInsured = &EnumRelationshipToInsured{
	Description: "relationships of drivers to the primary insured",
	Items: []*EnumRelationshipToInsuredItem{
		&relationshipToInsuredApplicant,
		&relationshipToInsuredSpouse,
		&relationshipToInsuredChild,
		&relationshipToInsuredParent,
		&relationshipToInsuredOtherRelative,
		&relationshipToInsuredNonResidentDriver,
		&relationshipToInsuredNoneOfTheAbove,
	},
	Name:              "EnumRelationshipToInsured",
	Applicant:         &relationshipToInsuredApplicant,
	Spouse:            &relationshipToInsuredSpouse,
	Child:             &relationshipToInsuredChild,
	Parent:            &relationshipToInsuredParent,
	OtherRelative:     &relationshipToInsuredOtherRelative,
	NonResidentDriver: &relationshipToInsuredNonResidentDriver,
	NoneOfTheAbove:    &relationshipToInsuredNoneOfTheAbove,

	itemDict: map[string]*EnumRelationshipToInsuredItem{
		strings.ToLower(string(relationshipToInsuredApplicantID)):         &relationshipToInsuredApplicant,
		strings.ToLower(string(relationshipToInsuredSpouseID)):            &relationshipToInsuredSpouse,
		strings.ToLower(string(relationshipToInsuredChildID)):             &relationshipToInsuredChild,
		strings.ToLower(string(relationshipToInsuredParentID)):            &relationshipToInsuredParent,
		strings.ToLower(string(relationshipToInsuredOtherRelativeID)):     &relationshipToInsuredOtherRelative,
		strings.ToLower(string(relationshipToInsuredNonResidentDriverID)): &relationshipToInsuredNonResidentDriver,
		strings.ToLower(string(relationshipToInsuredNoneOfTheAboveID)):    &relationshipToInsuredNoneOfTheAbove,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumRelationshipToInsured) ByID(id RelationshipToInsuredIdentifier) *EnumRelationshipToInsuredItem {
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
func (e *EnumRelationshipToInsured) ByIDString(idx string) *EnumRelationshipToInsuredItem {
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
func (e *EnumRelationshipToInsured) ByIndex(idx int) *EnumRelationshipToInsuredItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedRelationshipToInsuredID is a struct that is designed to replace a *RelationshipToInsuredID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *RelationshipToInsuredID it contains while being a better JSON citizen.
type ValidatedRelationshipToInsuredID struct {
	// id will point to a valid RelationshipToInsuredID, if possible
	// If id is nil, then ValidatedRelationshipToInsuredID.Valid() will return false.
	id *RelationshipToInsuredID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedRelationshipToInsuredID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedRelationshipToInsuredID
func (vi *ValidatedRelationshipToInsuredID) Clone() *ValidatedRelationshipToInsuredID {
	if vi == nil {
		return nil
	}

	var cid *RelationshipToInsuredID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedRelationshipToInsuredID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedRelationshipToInsuredIds represent the same RelationshipToInsured
func (vi *ValidatedRelationshipToInsuredID) Equals(vj *ValidatedRelationshipToInsuredID) bool {
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

// Valid returns true if and only if the ValidatedRelationshipToInsuredID corresponds to a recognized RelationshipToInsured
func (vi *ValidatedRelationshipToInsuredID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedRelationshipToInsuredID) ID() *RelationshipToInsuredID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedRelationshipToInsuredID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedRelationshipToInsuredID) ValidatedID() *ValidatedRelationshipToInsuredID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedRelationshipToInsuredID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedRelationshipToInsuredID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedRelationshipToInsuredID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedRelationshipToInsuredID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedRelationshipToInsuredID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := RelationshipToInsuredID(capString)
	item := RelationshipToInsured.ByID(&id)
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

func (vi ValidatedRelationshipToInsuredID) String() string {
	return vi.ToIDString()
}

type RelationshipToInsuredIdentifier interface {
	ID() *RelationshipToInsuredID
	Valid() bool
}
