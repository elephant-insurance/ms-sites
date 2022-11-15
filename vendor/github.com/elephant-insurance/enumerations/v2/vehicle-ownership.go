package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// VehicleOwnershipID uniquely identifies a particular VehicleOwnership
type VehicleOwnershipID string

// Clone creates a safe, independent copy of a VehicleOwnershipID
func (i *VehicleOwnershipID) Clone() *VehicleOwnershipID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two VehicleOwnershipIds are equivalent
func (i *VehicleOwnershipID) Equals(j *VehicleOwnershipID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *VehicleOwnershipID that is either valid or nil
func (i *VehicleOwnershipID) ID() *VehicleOwnershipID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *VehicleOwnershipID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the VehicleOwnershipID corresponds to a recognized VehicleOwnership
func (i *VehicleOwnershipID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return VehicleOwnership.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *VehicleOwnershipID) ValidatedID() *ValidatedVehicleOwnershipID {
	if i != nil {
		return &ValidatedVehicleOwnershipID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *VehicleOwnershipID) MarshalJSON() ([]byte, error) {
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

func (i *VehicleOwnershipID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := VehicleOwnershipID(dataString)
	item := VehicleOwnership.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	vehicleOwnershipPaidOffID      VehicleOwnershipID = "PaidOff"
	vehicleOwnershipMakePaymentsID VehicleOwnershipID = "Other"
)

// EnumVehicleOwnershipItem describes an entry in an enumeration of VehicleOwnership
type EnumVehicleOwnershipItem struct {
	ID        VehicleOwnershipID `json:"Value"`
	Desc      string             `json:"Description,omitempty"`
	Meta      map[string]string  `json:",omitempty"`
	Name      string             `json:"Name"`
	SortOrder int
}

var (
	vehicleOwnershipPaidOff      = EnumVehicleOwnershipItem{vehicleOwnershipPaidOffID, "Own and do not make payments", nil, "PaidOff", 3}
	vehicleOwnershipMakePayments = EnumVehicleOwnershipItem{vehicleOwnershipMakePaymentsID, "Make payments", nil, "MakePayments", 4}
)

// EnumVehicleOwnership is a collection of VehicleOwnership items
type EnumVehicleOwnership struct {
	Description string
	Items       []*EnumVehicleOwnershipItem
	Name        string

	PaidOff      *EnumVehicleOwnershipItem
	MakePayments *EnumVehicleOwnershipItem

	itemDict map[string]*EnumVehicleOwnershipItem
}

// VehicleOwnership is a public singleton instance of EnumVehicleOwnership
// representing ownership statuses for vehicles
var VehicleOwnership = &EnumVehicleOwnership{
	Description: "ownership statuses for vehicles",
	Items: []*EnumVehicleOwnershipItem{
		&vehicleOwnershipPaidOff,
		&vehicleOwnershipMakePayments,
	},
	Name:         "EnumVehicleOwnership",
	PaidOff:      &vehicleOwnershipPaidOff,
	MakePayments: &vehicleOwnershipMakePayments,

	itemDict: map[string]*EnumVehicleOwnershipItem{
		strings.ToLower(string(vehicleOwnershipPaidOffID)):      &vehicleOwnershipPaidOff,
		strings.ToLower(string(vehicleOwnershipMakePaymentsID)): &vehicleOwnershipMakePayments,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumVehicleOwnership) ByID(id VehicleOwnershipIdentifier) *EnumVehicleOwnershipItem {
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
func (e *EnumVehicleOwnership) ByIDString(idx string) *EnumVehicleOwnershipItem {
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
func (e *EnumVehicleOwnership) ByIndex(idx int) *EnumVehicleOwnershipItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedVehicleOwnershipID is a struct that is designed to replace a *VehicleOwnershipID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *VehicleOwnershipID it contains while being a better JSON citizen.
type ValidatedVehicleOwnershipID struct {
	// id will point to a valid VehicleOwnershipID, if possible
	// If id is nil, then ValidatedVehicleOwnershipID.Valid() will return false.
	id *VehicleOwnershipID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedVehicleOwnershipID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedVehicleOwnershipID
func (vi *ValidatedVehicleOwnershipID) Clone() *ValidatedVehicleOwnershipID {
	if vi == nil {
		return nil
	}

	var cid *VehicleOwnershipID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedVehicleOwnershipID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedVehicleOwnershipIds represent the same VehicleOwnership
func (vi *ValidatedVehicleOwnershipID) Equals(vj *ValidatedVehicleOwnershipID) bool {
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

// Valid returns true if and only if the ValidatedVehicleOwnershipID corresponds to a recognized VehicleOwnership
func (vi *ValidatedVehicleOwnershipID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedVehicleOwnershipID) ID() *VehicleOwnershipID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedVehicleOwnershipID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedVehicleOwnershipID) ValidatedID() *ValidatedVehicleOwnershipID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedVehicleOwnershipID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedVehicleOwnershipID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedVehicleOwnershipID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedVehicleOwnershipID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedVehicleOwnershipID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := VehicleOwnershipID(capString)
	item := VehicleOwnership.ByID(&id)
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

func (vi ValidatedVehicleOwnershipID) String() string {
	return vi.ToIDString()
}

type VehicleOwnershipIdentifier interface {
	ID() *VehicleOwnershipID
	Valid() bool
}
