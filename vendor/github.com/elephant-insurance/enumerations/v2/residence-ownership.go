package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// ResidenceOwnershipID uniquely identifies a particular ResidenceOwnership
type ResidenceOwnershipID string

// Clone creates a safe, independent copy of a ResidenceOwnershipID
func (i *ResidenceOwnershipID) Clone() *ResidenceOwnershipID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two ResidenceOwnershipIds are equivalent
func (i *ResidenceOwnershipID) Equals(j *ResidenceOwnershipID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *ResidenceOwnershipID that is either valid or nil
func (i *ResidenceOwnershipID) ID() *ResidenceOwnershipID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *ResidenceOwnershipID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the ResidenceOwnershipID corresponds to a recognized ResidenceOwnership
func (i *ResidenceOwnershipID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return ResidenceOwnership.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *ResidenceOwnershipID) ValidatedID() *ValidatedResidenceOwnershipID {
	if i != nil {
		return &ValidatedResidenceOwnershipID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *ResidenceOwnershipID) MarshalJSON() ([]byte, error) {
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

func (i *ResidenceOwnershipID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := ResidenceOwnershipID(dataString)
	item := ResidenceOwnership.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	residenceOwnershipRentID       ResidenceOwnershipID = "rent"
	residenceOwnershipHomeID       ResidenceOwnershipID = "home"
	residenceOwnershipOtherID      ResidenceOwnershipID = "other"
	residenceOwnershipCondoID      ResidenceOwnershipID = "condo"
	residenceOwnershipMobileHomeID ResidenceOwnershipID = "mobile_home"
)

// EnumResidenceOwnershipItem describes an entry in an enumeration of ResidenceOwnership
type EnumResidenceOwnershipItem struct {
	ID        ResidenceOwnershipID `json:"Value"`
	Desc      string               `json:"Description,omitempty"`
	Meta      map[string]string    `json:",omitempty"`
	Name      string               `json:"Name"`
	SortOrder int
}

var (
	residenceOwnershipRent       = EnumResidenceOwnershipItem{residenceOwnershipRentID, "Rent", nil, "Rent", 1}
	residenceOwnershipHome       = EnumResidenceOwnershipItem{residenceOwnershipHomeID, "Own Home", nil, "Home", 2}
	residenceOwnershipOther      = EnumResidenceOwnershipItem{residenceOwnershipOtherID, "Live with Parents/Other", nil, "Other", 3}
	residenceOwnershipCondo      = EnumResidenceOwnershipItem{residenceOwnershipCondoID, "Own Condo", nil, "Condo", 4}
	residenceOwnershipMobileHome = EnumResidenceOwnershipItem{residenceOwnershipMobileHomeID, "Own Mobile Home", nil, "MobileHome", 5}
)

// EnumResidenceOwnership is a collection of ResidenceOwnership items
type EnumResidenceOwnership struct {
	Description string
	Items       []*EnumResidenceOwnershipItem
	Name        string

	Rent       *EnumResidenceOwnershipItem
	Home       *EnumResidenceOwnershipItem
	Other      *EnumResidenceOwnershipItem
	Condo      *EnumResidenceOwnershipItem
	MobileHome *EnumResidenceOwnershipItem

	itemDict map[string]*EnumResidenceOwnershipItem
}

// ResidenceOwnership is a public singleton instance of EnumResidenceOwnership
// representing ownership statuses of customer residences
var ResidenceOwnership = &EnumResidenceOwnership{
	Description: "ownership statuses of customer residences",
	Items: []*EnumResidenceOwnershipItem{
		&residenceOwnershipRent,
		&residenceOwnershipHome,
		&residenceOwnershipOther,
		&residenceOwnershipCondo,
		&residenceOwnershipMobileHome,
	},
	Name:       "EnumResidenceOwnership",
	Rent:       &residenceOwnershipRent,
	Home:       &residenceOwnershipHome,
	Other:      &residenceOwnershipOther,
	Condo:      &residenceOwnershipCondo,
	MobileHome: &residenceOwnershipMobileHome,

	itemDict: map[string]*EnumResidenceOwnershipItem{
		strings.ToLower(string(residenceOwnershipRentID)):       &residenceOwnershipRent,
		strings.ToLower(string(residenceOwnershipHomeID)):       &residenceOwnershipHome,
		strings.ToLower(string(residenceOwnershipOtherID)):      &residenceOwnershipOther,
		strings.ToLower(string(residenceOwnershipCondoID)):      &residenceOwnershipCondo,
		strings.ToLower(string(residenceOwnershipMobileHomeID)): &residenceOwnershipMobileHome,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumResidenceOwnership) ByID(id ResidenceOwnershipIdentifier) *EnumResidenceOwnershipItem {
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
func (e *EnumResidenceOwnership) ByIDString(idx string) *EnumResidenceOwnershipItem {
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
func (e *EnumResidenceOwnership) ByIndex(idx int) *EnumResidenceOwnershipItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedResidenceOwnershipID is a struct that is designed to replace a *ResidenceOwnershipID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *ResidenceOwnershipID it contains while being a better JSON citizen.
type ValidatedResidenceOwnershipID struct {
	// id will point to a valid ResidenceOwnershipID, if possible
	// If id is nil, then ValidatedResidenceOwnershipID.Valid() will return false.
	id *ResidenceOwnershipID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedResidenceOwnershipID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedResidenceOwnershipID
func (vi *ValidatedResidenceOwnershipID) Clone() *ValidatedResidenceOwnershipID {
	if vi == nil {
		return nil
	}

	var cid *ResidenceOwnershipID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedResidenceOwnershipID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedResidenceOwnershipIds represent the same ResidenceOwnership
func (vi *ValidatedResidenceOwnershipID) Equals(vj *ValidatedResidenceOwnershipID) bool {
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

// Valid returns true if and only if the ValidatedResidenceOwnershipID corresponds to a recognized ResidenceOwnership
func (vi *ValidatedResidenceOwnershipID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedResidenceOwnershipID) ID() *ResidenceOwnershipID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedResidenceOwnershipID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedResidenceOwnershipID) ValidatedID() *ValidatedResidenceOwnershipID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedResidenceOwnershipID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedResidenceOwnershipID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedResidenceOwnershipID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedResidenceOwnershipID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedResidenceOwnershipID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := ResidenceOwnershipID(capString)
	item := ResidenceOwnership.ByID(&id)
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

func (vi ValidatedResidenceOwnershipID) String() string {
	return vi.ToIDString()
}

type ResidenceOwnershipIdentifier interface {
	ID() *ResidenceOwnershipID
	Valid() bool
}
