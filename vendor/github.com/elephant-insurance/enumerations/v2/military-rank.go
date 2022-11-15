package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// MilitaryRankID uniquely identifies a particular MilitaryRank
type MilitaryRankID string

// Clone creates a safe, independent copy of a MilitaryRankID
func (i *MilitaryRankID) Clone() *MilitaryRankID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two MilitaryRankIds are equivalent
func (i *MilitaryRankID) Equals(j *MilitaryRankID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *MilitaryRankID that is either valid or nil
func (i *MilitaryRankID) ID() *MilitaryRankID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *MilitaryRankID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the MilitaryRankID corresponds to a recognized MilitaryRank
func (i *MilitaryRankID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return MilitaryRank.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *MilitaryRankID) ValidatedID() *ValidatedMilitaryRankID {
	if i != nil {
		return &ValidatedMilitaryRankID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *MilitaryRankID) MarshalJSON() ([]byte, error) {
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

func (i *MilitaryRankID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := MilitaryRankID(dataString)
	item := MilitaryRank.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	militaryRankEnlistedE1ID             MilitaryRankID = "E1"
	militaryRankEnlistedE2ID             MilitaryRankID = "E2"
	militaryRankEnlistedE3ID             MilitaryRankID = "E3"
	militaryRankEnlistedE4ID             MilitaryRankID = "E4"
	militaryRankNonCommissionedE5ID      MilitaryRankID = "E5"
	militaryRankNonCommissionedE6ID      MilitaryRankID = "E6"
	militaryRankNonCommissionedE7ID      MilitaryRankID = "E7"
	militaryRankNonCommissionedE8ID      MilitaryRankID = "E8"
	militaryRankNonCommissionedE9ID      MilitaryRankID = "E9"
	militaryRankWarrantOfficerOW1ID      MilitaryRankID = "OW1"
	militaryRankWarrantOfficerOW2ID      MilitaryRankID = "OW2"
	militaryRankWarrantOfficerOW3ID      MilitaryRankID = "OW3"
	militaryRankWarrantOfficerOW4ID      MilitaryRankID = "OW4"
	militaryRankWarrantOfficerOW5ID      MilitaryRankID = "OW5"
	militaryRankCommissionedOfficerO1ID  MilitaryRankID = "O1"
	militaryRankCommissionedOfficerO2ID  MilitaryRankID = "O2"
	militaryRankCommissionedOfficerO3ID  MilitaryRankID = "O3"
	militaryRankCommissionedOfficerO4ID  MilitaryRankID = "O4"
	militaryRankCommissionedOfficerO5ID  MilitaryRankID = "O5"
	militaryRankCommissionedOfficerO6ID  MilitaryRankID = "O6"
	militaryRankCommissionedOfficerO7ID  MilitaryRankID = "O7"
	militaryRankCommissionedOfficerO8ID  MilitaryRankID = "O8"
	militaryRankCommissionedOfficerO9ID  MilitaryRankID = "O9"
	militaryRankCommissionedOfficerO10ID MilitaryRankID = "O10"
)

// EnumMilitaryRankItem describes an entry in an enumeration of MilitaryRank
type EnumMilitaryRankItem struct {
	ID        MilitaryRankID    `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	Classification string
}

var (
	militaryRankEnlistedE1             = EnumMilitaryRankItem{militaryRankEnlistedE1ID, "E-1", map[string]string{"Classification": MilitaryRankClassification.Enlisted}, "EnlistedE1", 1, MilitaryRankClassification.Enlisted}
	militaryRankEnlistedE2             = EnumMilitaryRankItem{militaryRankEnlistedE2ID, "E-2", map[string]string{"Classification": MilitaryRankClassification.Enlisted}, "EnlistedE2", 2, MilitaryRankClassification.Enlisted}
	militaryRankEnlistedE3             = EnumMilitaryRankItem{militaryRankEnlistedE3ID, "E-3", map[string]string{"Classification": MilitaryRankClassification.Enlisted}, "EnlistedE3", 3, MilitaryRankClassification.Enlisted}
	militaryRankEnlistedE4             = EnumMilitaryRankItem{militaryRankEnlistedE4ID, "E-4", map[string]string{"Classification": MilitaryRankClassification.Enlisted}, "EnlistedE4", 4, MilitaryRankClassification.Enlisted}
	militaryRankNonCommissionedE5      = EnumMilitaryRankItem{militaryRankNonCommissionedE5ID, "E-5", map[string]string{"Classification": MilitaryRankClassification.NonCommissionedOfficer}, "NonCommissionedE5", 5, MilitaryRankClassification.NonCommissionedOfficer}
	militaryRankNonCommissionedE6      = EnumMilitaryRankItem{militaryRankNonCommissionedE6ID, "E-6", map[string]string{"Classification": MilitaryRankClassification.NonCommissionedOfficer}, "NonCommissionedE6", 6, MilitaryRankClassification.NonCommissionedOfficer}
	militaryRankNonCommissionedE7      = EnumMilitaryRankItem{militaryRankNonCommissionedE7ID, "E-7", map[string]string{"Classification": MilitaryRankClassification.NonCommissionedOfficer}, "NonCommissionedE7", 7, MilitaryRankClassification.NonCommissionedOfficer}
	militaryRankNonCommissionedE8      = EnumMilitaryRankItem{militaryRankNonCommissionedE8ID, "E-8", map[string]string{"Classification": MilitaryRankClassification.NonCommissionedOfficer}, "NonCommissionedE8", 8, MilitaryRankClassification.NonCommissionedOfficer}
	militaryRankNonCommissionedE9      = EnumMilitaryRankItem{militaryRankNonCommissionedE9ID, "E-9", map[string]string{"Classification": MilitaryRankClassification.NonCommissionedOfficer}, "NonCommissionedE9", 9, MilitaryRankClassification.NonCommissionedOfficer}
	militaryRankWarrantOfficerOW1      = EnumMilitaryRankItem{militaryRankWarrantOfficerOW1ID, "OW-1", map[string]string{"Classification": MilitaryRankClassification.WarrantOfficer}, "WarrantOfficerOW1", 10, MilitaryRankClassification.WarrantOfficer}
	militaryRankWarrantOfficerOW2      = EnumMilitaryRankItem{militaryRankWarrantOfficerOW2ID, "OW-2", map[string]string{"Classification": MilitaryRankClassification.WarrantOfficer}, "WarrantOfficerOW2", 11, MilitaryRankClassification.WarrantOfficer}
	militaryRankWarrantOfficerOW3      = EnumMilitaryRankItem{militaryRankWarrantOfficerOW3ID, "OW-3", map[string]string{"Classification": MilitaryRankClassification.WarrantOfficer}, "WarrantOfficerOW3", 12, MilitaryRankClassification.WarrantOfficer}
	militaryRankWarrantOfficerOW4      = EnumMilitaryRankItem{militaryRankWarrantOfficerOW4ID, "OW-4", map[string]string{"Classification": MilitaryRankClassification.WarrantOfficer}, "WarrantOfficerOW4", 13, MilitaryRankClassification.WarrantOfficer}
	militaryRankWarrantOfficerOW5      = EnumMilitaryRankItem{militaryRankWarrantOfficerOW5ID, "OW-5", map[string]string{"Classification": MilitaryRankClassification.WarrantOfficer}, "WarrantOfficerOW5", 14, MilitaryRankClassification.WarrantOfficer}
	militaryRankCommissionedOfficerO1  = EnumMilitaryRankItem{militaryRankCommissionedOfficerO1ID, "O-1", map[string]string{"Classification": MilitaryRankClassification.CommissionedOfficer}, "CommissionedOfficerO1", 15, MilitaryRankClassification.CommissionedOfficer}
	militaryRankCommissionedOfficerO2  = EnumMilitaryRankItem{militaryRankCommissionedOfficerO2ID, "O-2", map[string]string{"Classification": MilitaryRankClassification.CommissionedOfficer}, "CommissionedOfficerO2", 16, MilitaryRankClassification.CommissionedOfficer}
	militaryRankCommissionedOfficerO3  = EnumMilitaryRankItem{militaryRankCommissionedOfficerO3ID, "O-3", map[string]string{"Classification": MilitaryRankClassification.CommissionedOfficer}, "CommissionedOfficerO3", 17, MilitaryRankClassification.CommissionedOfficer}
	militaryRankCommissionedOfficerO4  = EnumMilitaryRankItem{militaryRankCommissionedOfficerO4ID, "O-4", map[string]string{"Classification": MilitaryRankClassification.CommissionedOfficer}, "CommissionedOfficerO4", 18, MilitaryRankClassification.CommissionedOfficer}
	militaryRankCommissionedOfficerO5  = EnumMilitaryRankItem{militaryRankCommissionedOfficerO5ID, "O-5", map[string]string{"Classification": MilitaryRankClassification.CommissionedOfficer}, "CommissionedOfficerO5", 19, MilitaryRankClassification.CommissionedOfficer}
	militaryRankCommissionedOfficerO6  = EnumMilitaryRankItem{militaryRankCommissionedOfficerO6ID, "O-6", map[string]string{"Classification": MilitaryRankClassification.CommissionedOfficer}, "CommissionedOfficerO6", 20, MilitaryRankClassification.CommissionedOfficer}
	militaryRankCommissionedOfficerO7  = EnumMilitaryRankItem{militaryRankCommissionedOfficerO7ID, "O-7", map[string]string{"Classification": MilitaryRankClassification.CommissionedOfficer}, "CommissionedOfficerO7", 21, MilitaryRankClassification.CommissionedOfficer}
	militaryRankCommissionedOfficerO8  = EnumMilitaryRankItem{militaryRankCommissionedOfficerO8ID, "O-8", map[string]string{"Classification": MilitaryRankClassification.CommissionedOfficer}, "CommissionedOfficerO8", 22, MilitaryRankClassification.CommissionedOfficer}
	militaryRankCommissionedOfficerO9  = EnumMilitaryRankItem{militaryRankCommissionedOfficerO9ID, "O-9", map[string]string{"Classification": MilitaryRankClassification.CommissionedOfficer}, "CommissionedOfficerO9", 23, MilitaryRankClassification.CommissionedOfficer}
	militaryRankCommissionedOfficerO10 = EnumMilitaryRankItem{militaryRankCommissionedOfficerO10ID, "O-10", map[string]string{"Classification": MilitaryRankClassification.CommissionedOfficer}, "CommissionedOfficerO10", 24, MilitaryRankClassification.CommissionedOfficer}
)

// EnumMilitaryRank is a collection of MilitaryRank items
type EnumMilitaryRank struct {
	Description string
	Items       []*EnumMilitaryRankItem
	Name        string

	EnlistedE1             *EnumMilitaryRankItem
	EnlistedE2             *EnumMilitaryRankItem
	EnlistedE3             *EnumMilitaryRankItem
	EnlistedE4             *EnumMilitaryRankItem
	NonCommissionedE5      *EnumMilitaryRankItem
	NonCommissionedE6      *EnumMilitaryRankItem
	NonCommissionedE7      *EnumMilitaryRankItem
	NonCommissionedE8      *EnumMilitaryRankItem
	NonCommissionedE9      *EnumMilitaryRankItem
	WarrantOfficerOW1      *EnumMilitaryRankItem
	WarrantOfficerOW2      *EnumMilitaryRankItem
	WarrantOfficerOW3      *EnumMilitaryRankItem
	WarrantOfficerOW4      *EnumMilitaryRankItem
	WarrantOfficerOW5      *EnumMilitaryRankItem
	CommissionedOfficerO1  *EnumMilitaryRankItem
	CommissionedOfficerO2  *EnumMilitaryRankItem
	CommissionedOfficerO3  *EnumMilitaryRankItem
	CommissionedOfficerO4  *EnumMilitaryRankItem
	CommissionedOfficerO5  *EnumMilitaryRankItem
	CommissionedOfficerO6  *EnumMilitaryRankItem
	CommissionedOfficerO7  *EnumMilitaryRankItem
	CommissionedOfficerO8  *EnumMilitaryRankItem
	CommissionedOfficerO9  *EnumMilitaryRankItem
	CommissionedOfficerO10 *EnumMilitaryRankItem

	itemDict map[string]*EnumMilitaryRankItem
}

// MilitaryRank is a public singleton instance of EnumMilitaryRank
// representing military ranks
var MilitaryRank = &EnumMilitaryRank{
	Description: "military ranks",
	Items: []*EnumMilitaryRankItem{
		&militaryRankEnlistedE1,
		&militaryRankEnlistedE2,
		&militaryRankEnlistedE3,
		&militaryRankEnlistedE4,
		&militaryRankNonCommissionedE5,
		&militaryRankNonCommissionedE6,
		&militaryRankNonCommissionedE7,
		&militaryRankNonCommissionedE8,
		&militaryRankNonCommissionedE9,
		&militaryRankWarrantOfficerOW1,
		&militaryRankWarrantOfficerOW2,
		&militaryRankWarrantOfficerOW3,
		&militaryRankWarrantOfficerOW4,
		&militaryRankWarrantOfficerOW5,
		&militaryRankCommissionedOfficerO1,
		&militaryRankCommissionedOfficerO2,
		&militaryRankCommissionedOfficerO3,
		&militaryRankCommissionedOfficerO4,
		&militaryRankCommissionedOfficerO5,
		&militaryRankCommissionedOfficerO6,
		&militaryRankCommissionedOfficerO7,
		&militaryRankCommissionedOfficerO8,
		&militaryRankCommissionedOfficerO9,
		&militaryRankCommissionedOfficerO10,
	},
	Name:                   "EnumMilitaryRank",
	EnlistedE1:             &militaryRankEnlistedE1,
	EnlistedE2:             &militaryRankEnlistedE2,
	EnlistedE3:             &militaryRankEnlistedE3,
	EnlistedE4:             &militaryRankEnlistedE4,
	NonCommissionedE5:      &militaryRankNonCommissionedE5,
	NonCommissionedE6:      &militaryRankNonCommissionedE6,
	NonCommissionedE7:      &militaryRankNonCommissionedE7,
	NonCommissionedE8:      &militaryRankNonCommissionedE8,
	NonCommissionedE9:      &militaryRankNonCommissionedE9,
	WarrantOfficerOW1:      &militaryRankWarrantOfficerOW1,
	WarrantOfficerOW2:      &militaryRankWarrantOfficerOW2,
	WarrantOfficerOW3:      &militaryRankWarrantOfficerOW3,
	WarrantOfficerOW4:      &militaryRankWarrantOfficerOW4,
	WarrantOfficerOW5:      &militaryRankWarrantOfficerOW5,
	CommissionedOfficerO1:  &militaryRankCommissionedOfficerO1,
	CommissionedOfficerO2:  &militaryRankCommissionedOfficerO2,
	CommissionedOfficerO3:  &militaryRankCommissionedOfficerO3,
	CommissionedOfficerO4:  &militaryRankCommissionedOfficerO4,
	CommissionedOfficerO5:  &militaryRankCommissionedOfficerO5,
	CommissionedOfficerO6:  &militaryRankCommissionedOfficerO6,
	CommissionedOfficerO7:  &militaryRankCommissionedOfficerO7,
	CommissionedOfficerO8:  &militaryRankCommissionedOfficerO8,
	CommissionedOfficerO9:  &militaryRankCommissionedOfficerO9,
	CommissionedOfficerO10: &militaryRankCommissionedOfficerO10,

	itemDict: map[string]*EnumMilitaryRankItem{
		strings.ToLower(string(militaryRankEnlistedE1ID)):             &militaryRankEnlistedE1,
		strings.ToLower(string(militaryRankEnlistedE2ID)):             &militaryRankEnlistedE2,
		strings.ToLower(string(militaryRankEnlistedE3ID)):             &militaryRankEnlistedE3,
		strings.ToLower(string(militaryRankEnlistedE4ID)):             &militaryRankEnlistedE4,
		strings.ToLower(string(militaryRankNonCommissionedE5ID)):      &militaryRankNonCommissionedE5,
		strings.ToLower(string(militaryRankNonCommissionedE6ID)):      &militaryRankNonCommissionedE6,
		strings.ToLower(string(militaryRankNonCommissionedE7ID)):      &militaryRankNonCommissionedE7,
		strings.ToLower(string(militaryRankNonCommissionedE8ID)):      &militaryRankNonCommissionedE8,
		strings.ToLower(string(militaryRankNonCommissionedE9ID)):      &militaryRankNonCommissionedE9,
		strings.ToLower(string(militaryRankWarrantOfficerOW1ID)):      &militaryRankWarrantOfficerOW1,
		strings.ToLower(string(militaryRankWarrantOfficerOW2ID)):      &militaryRankWarrantOfficerOW2,
		strings.ToLower(string(militaryRankWarrantOfficerOW3ID)):      &militaryRankWarrantOfficerOW3,
		strings.ToLower(string(militaryRankWarrantOfficerOW4ID)):      &militaryRankWarrantOfficerOW4,
		strings.ToLower(string(militaryRankWarrantOfficerOW5ID)):      &militaryRankWarrantOfficerOW5,
		strings.ToLower(string(militaryRankCommissionedOfficerO1ID)):  &militaryRankCommissionedOfficerO1,
		strings.ToLower(string(militaryRankCommissionedOfficerO2ID)):  &militaryRankCommissionedOfficerO2,
		strings.ToLower(string(militaryRankCommissionedOfficerO3ID)):  &militaryRankCommissionedOfficerO3,
		strings.ToLower(string(militaryRankCommissionedOfficerO4ID)):  &militaryRankCommissionedOfficerO4,
		strings.ToLower(string(militaryRankCommissionedOfficerO5ID)):  &militaryRankCommissionedOfficerO5,
		strings.ToLower(string(militaryRankCommissionedOfficerO6ID)):  &militaryRankCommissionedOfficerO6,
		strings.ToLower(string(militaryRankCommissionedOfficerO7ID)):  &militaryRankCommissionedOfficerO7,
		strings.ToLower(string(militaryRankCommissionedOfficerO8ID)):  &militaryRankCommissionedOfficerO8,
		strings.ToLower(string(militaryRankCommissionedOfficerO9ID)):  &militaryRankCommissionedOfficerO9,
		strings.ToLower(string(militaryRankCommissionedOfficerO10ID)): &militaryRankCommissionedOfficerO10,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumMilitaryRank) ByID(id MilitaryRankIdentifier) *EnumMilitaryRankItem {
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
func (e *EnumMilitaryRank) ByIDString(idx string) *EnumMilitaryRankItem {
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
func (e *EnumMilitaryRank) ByIndex(idx int) *EnumMilitaryRankItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedMilitaryRankID is a struct that is designed to replace a *MilitaryRankID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *MilitaryRankID it contains while being a better JSON citizen.
type ValidatedMilitaryRankID struct {
	// id will point to a valid MilitaryRankID, if possible
	// If id is nil, then ValidatedMilitaryRankID.Valid() will return false.
	id *MilitaryRankID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedMilitaryRankID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedMilitaryRankID
func (vi *ValidatedMilitaryRankID) Clone() *ValidatedMilitaryRankID {
	if vi == nil {
		return nil
	}

	var cid *MilitaryRankID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedMilitaryRankID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedMilitaryRankIds represent the same MilitaryRank
func (vi *ValidatedMilitaryRankID) Equals(vj *ValidatedMilitaryRankID) bool {
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

// Valid returns true if and only if the ValidatedMilitaryRankID corresponds to a recognized MilitaryRank
func (vi *ValidatedMilitaryRankID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedMilitaryRankID) ID() *MilitaryRankID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedMilitaryRankID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedMilitaryRankID) ValidatedID() *ValidatedMilitaryRankID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedMilitaryRankID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedMilitaryRankID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedMilitaryRankID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedMilitaryRankID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedMilitaryRankID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := MilitaryRankID(capString)
	item := MilitaryRank.ByID(&id)
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

func (vi ValidatedMilitaryRankID) String() string {
	return vi.ToIDString()
}

type MilitaryRankIdentifier interface {
	ID() *MilitaryRankID
	Valid() bool
}
