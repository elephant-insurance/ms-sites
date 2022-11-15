package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// PrimarilyParkedID uniquely identifies a particular PrimarilyParked
type PrimarilyParkedID string

// Clone creates a safe, independent copy of a PrimarilyParkedID
func (i *PrimarilyParkedID) Clone() *PrimarilyParkedID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two PrimarilyParkedIds are equivalent
func (i *PrimarilyParkedID) Equals(j *PrimarilyParkedID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *PrimarilyParkedID that is either valid or nil
func (i *PrimarilyParkedID) ID() *PrimarilyParkedID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *PrimarilyParkedID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the PrimarilyParkedID corresponds to a recognized PrimarilyParked
func (i *PrimarilyParkedID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return PrimarilyParked.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *PrimarilyParkedID) ValidatedID() *ValidatedPrimarilyParkedID {
	if i != nil {
		return &ValidatedPrimarilyParkedID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *PrimarilyParkedID) MarshalJSON() ([]byte, error) {
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

func (i *PrimarilyParkedID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := PrimarilyParkedID(dataString)
	item := PrimarilyParked.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	primarilyParkedCarportID    PrimarilyParkedID = "Carport"
	primarilyParkedDrivewayID   PrimarilyParkedID = "Driveway"
	primarilyParkedIndoorLotID  PrimarilyParkedID = "Indoor_parking_lot"
	primarilyParkedGarageID     PrimarilyParkedID = "Garage"
	primarilyParkedOutdoorLotID PrimarilyParkedID = "Outdoor_parking_lot"
	primarilyParkedStreetID     PrimarilyParkedID = "Street"
	primarilyParkedOtherID      PrimarilyParkedID = "Other"
)

// EnumPrimarilyParkedItem describes an entry in an enumeration of PrimarilyParked
type EnumPrimarilyParkedItem struct {
	ID        PrimarilyParkedID `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	primarilyParkedCarport    = EnumPrimarilyParkedItem{primarilyParkedCarportID, "Carport", nil, "Carport", 1}
	primarilyParkedDriveway   = EnumPrimarilyParkedItem{primarilyParkedDrivewayID, "Driveway", nil, "Driveway", 2}
	primarilyParkedIndoorLot  = EnumPrimarilyParkedItem{primarilyParkedIndoorLotID, "Parking Garage", nil, "IndoorLot", 3}
	primarilyParkedGarage     = EnumPrimarilyParkedItem{primarilyParkedGarageID, "Private Garage", nil, "Garage", 4}
	primarilyParkedOutdoorLot = EnumPrimarilyParkedItem{primarilyParkedOutdoorLotID, "Parking Lot", nil, "OutdoorLot", 5}
	primarilyParkedStreet     = EnumPrimarilyParkedItem{primarilyParkedStreetID, "Street", nil, "Street", 6}
	primarilyParkedOther      = EnumPrimarilyParkedItem{primarilyParkedOtherID, "Other", nil, "Other", 7}
)

// EnumPrimarilyParked is a collection of PrimarilyParked items
type EnumPrimarilyParked struct {
	Description string
	Items       []*EnumPrimarilyParkedItem
	Name        string

	Carport    *EnumPrimarilyParkedItem
	Driveway   *EnumPrimarilyParkedItem
	IndoorLot  *EnumPrimarilyParkedItem
	Garage     *EnumPrimarilyParkedItem
	OutdoorLot *EnumPrimarilyParkedItem
	Street     *EnumPrimarilyParkedItem
	Other      *EnumPrimarilyParkedItem

	itemDict map[string]*EnumPrimarilyParkedItem
}

// PrimarilyParked is a public singleton instance of EnumPrimarilyParked
// representing primary parking locations for vehicles
var PrimarilyParked = &EnumPrimarilyParked{
	Description: "primary parking locations for vehicles",
	Items: []*EnumPrimarilyParkedItem{
		&primarilyParkedCarport,
		&primarilyParkedDriveway,
		&primarilyParkedIndoorLot,
		&primarilyParkedGarage,
		&primarilyParkedOutdoorLot,
		&primarilyParkedStreet,
		&primarilyParkedOther,
	},
	Name:       "EnumPrimarilyParked",
	Carport:    &primarilyParkedCarport,
	Driveway:   &primarilyParkedDriveway,
	IndoorLot:  &primarilyParkedIndoorLot,
	Garage:     &primarilyParkedGarage,
	OutdoorLot: &primarilyParkedOutdoorLot,
	Street:     &primarilyParkedStreet,
	Other:      &primarilyParkedOther,

	itemDict: map[string]*EnumPrimarilyParkedItem{
		strings.ToLower(string(primarilyParkedCarportID)):    &primarilyParkedCarport,
		strings.ToLower(string(primarilyParkedDrivewayID)):   &primarilyParkedDriveway,
		strings.ToLower(string(primarilyParkedIndoorLotID)):  &primarilyParkedIndoorLot,
		strings.ToLower(string(primarilyParkedGarageID)):     &primarilyParkedGarage,
		strings.ToLower(string(primarilyParkedOutdoorLotID)): &primarilyParkedOutdoorLot,
		strings.ToLower(string(primarilyParkedStreetID)):     &primarilyParkedStreet,
		strings.ToLower(string(primarilyParkedOtherID)):      &primarilyParkedOther,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumPrimarilyParked) ByID(id PrimarilyParkedIdentifier) *EnumPrimarilyParkedItem {
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
func (e *EnumPrimarilyParked) ByIDString(idx string) *EnumPrimarilyParkedItem {
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
func (e *EnumPrimarilyParked) ByIndex(idx int) *EnumPrimarilyParkedItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedPrimarilyParkedID is a struct that is designed to replace a *PrimarilyParkedID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *PrimarilyParkedID it contains while being a better JSON citizen.
type ValidatedPrimarilyParkedID struct {
	// id will point to a valid PrimarilyParkedID, if possible
	// If id is nil, then ValidatedPrimarilyParkedID.Valid() will return false.
	id *PrimarilyParkedID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedPrimarilyParkedID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedPrimarilyParkedID
func (vi *ValidatedPrimarilyParkedID) Clone() *ValidatedPrimarilyParkedID {
	if vi == nil {
		return nil
	}

	var cid *PrimarilyParkedID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedPrimarilyParkedID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedPrimarilyParkedIds represent the same PrimarilyParked
func (vi *ValidatedPrimarilyParkedID) Equals(vj *ValidatedPrimarilyParkedID) bool {
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

// Valid returns true if and only if the ValidatedPrimarilyParkedID corresponds to a recognized PrimarilyParked
func (vi *ValidatedPrimarilyParkedID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedPrimarilyParkedID) ID() *PrimarilyParkedID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedPrimarilyParkedID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedPrimarilyParkedID) ValidatedID() *ValidatedPrimarilyParkedID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedPrimarilyParkedID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedPrimarilyParkedID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedPrimarilyParkedID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedPrimarilyParkedID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedPrimarilyParkedID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := PrimarilyParkedID(capString)
	item := PrimarilyParked.ByID(&id)
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

func (vi ValidatedPrimarilyParkedID) String() string {
	return vi.ToIDString()
}

type PrimarilyParkedIdentifier interface {
	ID() *PrimarilyParkedID
	Valid() bool
}
