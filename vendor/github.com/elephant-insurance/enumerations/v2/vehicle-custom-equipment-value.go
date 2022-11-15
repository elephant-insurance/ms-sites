package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// VehicleCustomEquipmentValueID uniquely identifies a particular VehicleCustomEquipmentValue
type VehicleCustomEquipmentValueID string

// Clone creates a safe, independent copy of a VehicleCustomEquipmentValueID
func (i *VehicleCustomEquipmentValueID) Clone() *VehicleCustomEquipmentValueID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two VehicleCustomEquipmentValueIds are equivalent
func (i *VehicleCustomEquipmentValueID) Equals(j *VehicleCustomEquipmentValueID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *VehicleCustomEquipmentValueID that is either valid or nil
func (i *VehicleCustomEquipmentValueID) ID() *VehicleCustomEquipmentValueID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *VehicleCustomEquipmentValueID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the VehicleCustomEquipmentValueID corresponds to a recognized VehicleCustomEquipmentValue
func (i *VehicleCustomEquipmentValueID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return VehicleCustomEquipmentValue.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *VehicleCustomEquipmentValueID) ValidatedID() *ValidatedVehicleCustomEquipmentValueID {
	if i != nil {
		return &ValidatedVehicleCustomEquipmentValueID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *VehicleCustomEquipmentValueID) MarshalJSON() ([]byte, error) {
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

func (i *VehicleCustomEquipmentValueID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := VehicleCustomEquipmentValueID(dataString)
	item := VehicleCustomEquipmentValue.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	vehicleCustomEquipmentValueUpTo1500ID       VehicleCustomEquipmentValueID = "1500R"
	vehicleCustomEquipmentValueFrom1501To2000ID VehicleCustomEquipmentValueID = "2000R"
	vehicleCustomEquipmentValueFrom2001To2500ID VehicleCustomEquipmentValueID = "2500R"
	vehicleCustomEquipmentValueFrom2501To3000ID VehicleCustomEquipmentValueID = "3000R"
	vehicleCustomEquipmentValueFrom3001To3500ID VehicleCustomEquipmentValueID = "3500R"
	vehicleCustomEquipmentValueFrom3501To4000ID VehicleCustomEquipmentValueID = "4000R"
	vehicleCustomEquipmentValueFrom4001To4500ID VehicleCustomEquipmentValueID = "4500R"
	vehicleCustomEquipmentValueFrom4501To5000ID VehicleCustomEquipmentValueID = "5000R"
	vehicleCustomEquipmentValueMoreThan5000ID   VehicleCustomEquipmentValueID = "5000Plus"
)

// EnumVehicleCustomEquipmentValueItem describes an entry in an enumeration of VehicleCustomEquipmentValue
type EnumVehicleCustomEquipmentValueItem struct {
	ID        VehicleCustomEquipmentValueID `json:"Value"`
	Desc      string                        `json:"Description,omitempty"`
	Meta      map[string]string             `json:",omitempty"`
	Name      string                        `json:"Name"`
	SortOrder int

	// Meta Properties
	StateCodes string
}

var (
	vehicleCustomEquipmentValueUpTo1500       = EnumVehicleCustomEquipmentValueItem{vehicleCustomEquipmentValueUpTo1500ID, "$1,001 - $1,500", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "UpTo1500", 1, "VA,TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentValueFrom1501To2000 = EnumVehicleCustomEquipmentValueItem{vehicleCustomEquipmentValueFrom1501To2000ID, "$1,501 - $2,000", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "From1501To2000", 2, "VA,TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentValueFrom2001To2500 = EnumVehicleCustomEquipmentValueItem{vehicleCustomEquipmentValueFrom2001To2500ID, "$2,001 - $2,500", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "From2001To2500", 3, "VA,TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentValueFrom2501To3000 = EnumVehicleCustomEquipmentValueItem{vehicleCustomEquipmentValueFrom2501To3000ID, "$2,501 - $3,000", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "From2501To3000", 4, "VA,TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentValueFrom3001To3500 = EnumVehicleCustomEquipmentValueItem{vehicleCustomEquipmentValueFrom3001To3500ID, "$3,001 - $3,500", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "From3001To3500", 5, "VA,TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentValueFrom3501To4000 = EnumVehicleCustomEquipmentValueItem{vehicleCustomEquipmentValueFrom3501To4000ID, "$3,501 - $4,000", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "From3501To4000", 6, "VA,TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentValueFrom4001To4500 = EnumVehicleCustomEquipmentValueItem{vehicleCustomEquipmentValueFrom4001To4500ID, "$4,001 - $4,500", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "From4001To4500", 7, "VA,TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentValueFrom4501To5000 = EnumVehicleCustomEquipmentValueItem{vehicleCustomEquipmentValueFrom4501To5000ID, "$4,500 - $5,000", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "From4501To5000", 8, "VA,TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentValueMoreThan5000   = EnumVehicleCustomEquipmentValueItem{vehicleCustomEquipmentValueMoreThan5000ID, "Greater than $5000", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "MoreThan5000", 9, "VA,TX,MD,IL,IN,TN,GA,OH"}
)

// EnumVehicleCustomEquipmentValue is a collection of VehicleCustomEquipmentValue items
type EnumVehicleCustomEquipmentValue struct {
	Description string
	Items       []*EnumVehicleCustomEquipmentValueItem
	Name        string

	UpTo1500       *EnumVehicleCustomEquipmentValueItem
	From1501To2000 *EnumVehicleCustomEquipmentValueItem
	From2001To2500 *EnumVehicleCustomEquipmentValueItem
	From2501To3000 *EnumVehicleCustomEquipmentValueItem
	From3001To3500 *EnumVehicleCustomEquipmentValueItem
	From3501To4000 *EnumVehicleCustomEquipmentValueItem
	From4001To4500 *EnumVehicleCustomEquipmentValueItem
	From4501To5000 *EnumVehicleCustomEquipmentValueItem
	MoreThan5000   *EnumVehicleCustomEquipmentValueItem

	itemDict map[string]*EnumVehicleCustomEquipmentValueItem
}

// VehicleCustomEquipmentValue is a public singleton instance of EnumVehicleCustomEquipmentValue
// representing value levels for vehicle custom equipment
var VehicleCustomEquipmentValue = &EnumVehicleCustomEquipmentValue{
	Description: "value levels for vehicle custom equipment",
	Items: []*EnumVehicleCustomEquipmentValueItem{
		&vehicleCustomEquipmentValueUpTo1500,
		&vehicleCustomEquipmentValueFrom1501To2000,
		&vehicleCustomEquipmentValueFrom2001To2500,
		&vehicleCustomEquipmentValueFrom2501To3000,
		&vehicleCustomEquipmentValueFrom3001To3500,
		&vehicleCustomEquipmentValueFrom3501To4000,
		&vehicleCustomEquipmentValueFrom4001To4500,
		&vehicleCustomEquipmentValueFrom4501To5000,
		&vehicleCustomEquipmentValueMoreThan5000,
	},
	Name:           "EnumVehicleCustomEquipmentValue",
	UpTo1500:       &vehicleCustomEquipmentValueUpTo1500,
	From1501To2000: &vehicleCustomEquipmentValueFrom1501To2000,
	From2001To2500: &vehicleCustomEquipmentValueFrom2001To2500,
	From2501To3000: &vehicleCustomEquipmentValueFrom2501To3000,
	From3001To3500: &vehicleCustomEquipmentValueFrom3001To3500,
	From3501To4000: &vehicleCustomEquipmentValueFrom3501To4000,
	From4001To4500: &vehicleCustomEquipmentValueFrom4001To4500,
	From4501To5000: &vehicleCustomEquipmentValueFrom4501To5000,
	MoreThan5000:   &vehicleCustomEquipmentValueMoreThan5000,

	itemDict: map[string]*EnumVehicleCustomEquipmentValueItem{
		strings.ToLower(string(vehicleCustomEquipmentValueUpTo1500ID)):       &vehicleCustomEquipmentValueUpTo1500,
		strings.ToLower(string(vehicleCustomEquipmentValueFrom1501To2000ID)): &vehicleCustomEquipmentValueFrom1501To2000,
		strings.ToLower(string(vehicleCustomEquipmentValueFrom2001To2500ID)): &vehicleCustomEquipmentValueFrom2001To2500,
		strings.ToLower(string(vehicleCustomEquipmentValueFrom2501To3000ID)): &vehicleCustomEquipmentValueFrom2501To3000,
		strings.ToLower(string(vehicleCustomEquipmentValueFrom3001To3500ID)): &vehicleCustomEquipmentValueFrom3001To3500,
		strings.ToLower(string(vehicleCustomEquipmentValueFrom3501To4000ID)): &vehicleCustomEquipmentValueFrom3501To4000,
		strings.ToLower(string(vehicleCustomEquipmentValueFrom4001To4500ID)): &vehicleCustomEquipmentValueFrom4001To4500,
		strings.ToLower(string(vehicleCustomEquipmentValueFrom4501To5000ID)): &vehicleCustomEquipmentValueFrom4501To5000,
		strings.ToLower(string(vehicleCustomEquipmentValueMoreThan5000ID)):   &vehicleCustomEquipmentValueMoreThan5000,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumVehicleCustomEquipmentValue) ByID(id VehicleCustomEquipmentValueIdentifier) *EnumVehicleCustomEquipmentValueItem {
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
func (e *EnumVehicleCustomEquipmentValue) ByIDString(idx string) *EnumVehicleCustomEquipmentValueItem {
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
func (e *EnumVehicleCustomEquipmentValue) ByIndex(idx int) *EnumVehicleCustomEquipmentValueItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedVehicleCustomEquipmentValueID is a struct that is designed to replace a *VehicleCustomEquipmentValueID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *VehicleCustomEquipmentValueID it contains while being a better JSON citizen.
type ValidatedVehicleCustomEquipmentValueID struct {
	// id will point to a valid VehicleCustomEquipmentValueID, if possible
	// If id is nil, then ValidatedVehicleCustomEquipmentValueID.Valid() will return false.
	id *VehicleCustomEquipmentValueID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedVehicleCustomEquipmentValueID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedVehicleCustomEquipmentValueID
func (vi *ValidatedVehicleCustomEquipmentValueID) Clone() *ValidatedVehicleCustomEquipmentValueID {
	if vi == nil {
		return nil
	}

	var cid *VehicleCustomEquipmentValueID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedVehicleCustomEquipmentValueID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedVehicleCustomEquipmentValueIds represent the same VehicleCustomEquipmentValue
func (vi *ValidatedVehicleCustomEquipmentValueID) Equals(vj *ValidatedVehicleCustomEquipmentValueID) bool {
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

// Valid returns true if and only if the ValidatedVehicleCustomEquipmentValueID corresponds to a recognized VehicleCustomEquipmentValue
func (vi *ValidatedVehicleCustomEquipmentValueID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedVehicleCustomEquipmentValueID) ID() *VehicleCustomEquipmentValueID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedVehicleCustomEquipmentValueID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedVehicleCustomEquipmentValueID) ValidatedID() *ValidatedVehicleCustomEquipmentValueID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedVehicleCustomEquipmentValueID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedVehicleCustomEquipmentValueID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedVehicleCustomEquipmentValueID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedVehicleCustomEquipmentValueID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedVehicleCustomEquipmentValueID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := VehicleCustomEquipmentValueID(capString)
	item := VehicleCustomEquipmentValue.ByID(&id)
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

func (vi ValidatedVehicleCustomEquipmentValueID) String() string {
	return vi.ToIDString()
}

type VehicleCustomEquipmentValueIdentifier interface {
	ID() *VehicleCustomEquipmentValueID
	Valid() bool
}
