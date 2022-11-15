package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// VehicleCustomEquipmentv6ValueID uniquely identifies a particular VehicleCustomEquipmentv6Value
type VehicleCustomEquipmentv6ValueID string

// Clone creates a safe, independent copy of a VehicleCustomEquipmentv6ValueID
func (i *VehicleCustomEquipmentv6ValueID) Clone() *VehicleCustomEquipmentv6ValueID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two VehicleCustomEquipmentv6ValueIds are equivalent
func (i *VehicleCustomEquipmentv6ValueID) Equals(j *VehicleCustomEquipmentv6ValueID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *VehicleCustomEquipmentv6ValueID that is either valid or nil
func (i *VehicleCustomEquipmentv6ValueID) ID() *VehicleCustomEquipmentv6ValueID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *VehicleCustomEquipmentv6ValueID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the VehicleCustomEquipmentv6ValueID corresponds to a recognized VehicleCustomEquipmentv6Value
func (i *VehicleCustomEquipmentv6ValueID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return VehicleCustomEquipmentv6Value.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *VehicleCustomEquipmentv6ValueID) ValidatedID() *ValidatedVehicleCustomEquipmentv6ValueID {
	if i != nil {
		return &ValidatedVehicleCustomEquipmentv6ValueID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *VehicleCustomEquipmentv6ValueID) MarshalJSON() ([]byte, error) {
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

func (i *VehicleCustomEquipmentv6ValueID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := VehicleCustomEquipmentv6ValueID(dataString)
	item := VehicleCustomEquipmentv6Value.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	vehicleCustomEquipmentv6ValueUpto1000ID         VehicleCustomEquipmentv6ValueID = "1000"
	vehicleCustomEquipmentv6ValueFrom1001To1500ID   VehicleCustomEquipmentv6ValueID = "1500R"
	vehicleCustomEquipmentv6ValueFrom1501To2000ID   VehicleCustomEquipmentv6ValueID = "2000R"
	vehicleCustomEquipmentv6ValueFrom2001To2500ID   VehicleCustomEquipmentv6ValueID = "2500R"
	vehicleCustomEquipmentv6ValueFrom2501To3000ID   VehicleCustomEquipmentv6ValueID = "3000R"
	vehicleCustomEquipmentv6ValueFrom3001To3500ID   VehicleCustomEquipmentv6ValueID = "3500R"
	vehicleCustomEquipmentv6ValueFrom3501To4000ID   VehicleCustomEquipmentv6ValueID = "4000R"
	vehicleCustomEquipmentv6ValueFrom4001To4500ID   VehicleCustomEquipmentv6ValueID = "4500R"
	vehicleCustomEquipmentv6ValueFrom4501To5000ID   VehicleCustomEquipmentv6ValueID = "5000R"
	vehicleCustomEquipmentv6ValueUpto2000ID         VehicleCustomEquipmentv6ValueID = "2000"
	vehicleCustomEquipmentv6ValueUpto3000ID         VehicleCustomEquipmentv6ValueID = "3000"
	vehicleCustomEquipmentv6ValueUpto4000ID         VehicleCustomEquipmentv6ValueID = "4000"
	vehicleCustomEquipmentv6ValueUpto5000ID         VehicleCustomEquipmentv6ValueID = "5000"
	vehicleCustomEquipmentv6ValueUpto10000ID        VehicleCustomEquipmentv6ValueID = "10000"
	vehicleCustomEquipmentv6ValueUpto15000ID        VehicleCustomEquipmentv6ValueID = "15000"
	vehicleCustomEquipmentv6ValueUpto20000ID        VehicleCustomEquipmentv6ValueID = "20000"
	vehicleCustomEquipmentv6ValueGreaterThanAboveID VehicleCustomEquipmentv6ValueID = "5000Plus"
)

// EnumVehicleCustomEquipmentv6ValueItem describes an entry in an enumeration of VehicleCustomEquipmentv6Value
type EnumVehicleCustomEquipmentv6ValueItem struct {
	ID        VehicleCustomEquipmentv6ValueID `json:"Value"`
	Desc      string                          `json:"Description,omitempty"`
	Meta      map[string]string               `json:",omitempty"`
	Name      string                          `json:"Name"`
	SortOrder int

	// Meta Properties
	StateCodes string
}

var (
	vehicleCustomEquipmentv6ValueUpto1000         = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueUpto1000ID, "Up to $1,000", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "Upto1000", 1, "VA,TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentv6ValueFrom1001To1500   = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueFrom1001To1500ID, "$1,001 - $1,500", map[string]string{"StateCodes": "VA"}, "From1001To1500", 2, "VA"}
	vehicleCustomEquipmentv6ValueFrom1501To2000   = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueFrom1501To2000ID, "$1,501 - $2,000", map[string]string{"StateCodes": "VA"}, "From1501To2000", 3, "VA"}
	vehicleCustomEquipmentv6ValueFrom2001To2500   = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueFrom2001To2500ID, "$2,001 - $2,500", map[string]string{"StateCodes": "VA"}, "From2001To2500", 4, "VA"}
	vehicleCustomEquipmentv6ValueFrom2501To3000   = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueFrom2501To3000ID, "$2,501 - $3,000", map[string]string{"StateCodes": "VA"}, "From2501To3000", 5, "VA"}
	vehicleCustomEquipmentv6ValueFrom3001To3500   = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueFrom3001To3500ID, "$3,001 - $3,500", map[string]string{"StateCodes": "VA"}, "From3001To3500", 6, "VA"}
	vehicleCustomEquipmentv6ValueFrom3501To4000   = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueFrom3501To4000ID, "$3,501 - $4,000", map[string]string{"StateCodes": "VA"}, "From3501To4000", 7, "VA"}
	vehicleCustomEquipmentv6ValueFrom4001To4500   = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueFrom4001To4500ID, "$4,001 - $4,500", map[string]string{"StateCodes": "VA"}, "From4001To4500", 8, "VA"}
	vehicleCustomEquipmentv6ValueFrom4501To5000   = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueFrom4501To5000ID, "$4,501 - $5,000", map[string]string{"StateCodes": "VA"}, "From4501To5000", 9, "VA"}
	vehicleCustomEquipmentv6ValueUpto2000         = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueUpto2000ID, "$2,000", map[string]string{"StateCodes": "TX,MD,IL,IN,TN,GA,OH"}, "Upto2000", 10, "TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentv6ValueUpto3000         = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueUpto3000ID, "$3,000", map[string]string{"StateCodes": "TX,MD,IL,IN,TN,GA,OH"}, "Upto3000", 11, "TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentv6ValueUpto4000         = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueUpto4000ID, "$4,000", map[string]string{"StateCodes": "TX,MD,IL,IN,TN,GA,OH"}, "Upto4000", 12, "TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentv6ValueUpto5000         = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueUpto5000ID, "$5,000", map[string]string{"StateCodes": "TX,MD,IL,IN,TN,GA,OH"}, "Upto5000", 13, "TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentv6ValueUpto10000        = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueUpto10000ID, "$10,000", map[string]string{"StateCodes": "TX,MD,IL,IN,TN,GA,OH"}, "Upto10000", 14, "TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentv6ValueUpto15000        = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueUpto15000ID, "$15,000", map[string]string{"StateCodes": "TX,MD,IL,IN,TN,GA,OH"}, "Upto15000", 15, "TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentv6ValueUpto20000        = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueUpto20000ID, "$20,000", map[string]string{"StateCodes": "TX,MD,IL,IN,TN,GA,OH"}, "Upto20000", 16, "TX,MD,IL,IN,TN,GA,OH"}
	vehicleCustomEquipmentv6ValueGreaterThanAbove = EnumVehicleCustomEquipmentv6ValueItem{vehicleCustomEquipmentv6ValueGreaterThanAboveID, "Greater than above", map[string]string{"StateCodes": "VA,TX,MD,IL,IN,TN,GA,OH"}, "GreaterThanAbove", 17, "VA,TX,MD,IL,IN,TN,GA,OH"}
)

// EnumVehicleCustomEquipmentv6Value is a collection of VehicleCustomEquipmentv6Value items
type EnumVehicleCustomEquipmentv6Value struct {
	Description string
	Items       []*EnumVehicleCustomEquipmentv6ValueItem
	Name        string

	Upto1000         *EnumVehicleCustomEquipmentv6ValueItem
	From1001To1500   *EnumVehicleCustomEquipmentv6ValueItem
	From1501To2000   *EnumVehicleCustomEquipmentv6ValueItem
	From2001To2500   *EnumVehicleCustomEquipmentv6ValueItem
	From2501To3000   *EnumVehicleCustomEquipmentv6ValueItem
	From3001To3500   *EnumVehicleCustomEquipmentv6ValueItem
	From3501To4000   *EnumVehicleCustomEquipmentv6ValueItem
	From4001To4500   *EnumVehicleCustomEquipmentv6ValueItem
	From4501To5000   *EnumVehicleCustomEquipmentv6ValueItem
	Upto2000         *EnumVehicleCustomEquipmentv6ValueItem
	Upto3000         *EnumVehicleCustomEquipmentv6ValueItem
	Upto4000         *EnumVehicleCustomEquipmentv6ValueItem
	Upto5000         *EnumVehicleCustomEquipmentv6ValueItem
	Upto10000        *EnumVehicleCustomEquipmentv6ValueItem
	Upto15000        *EnumVehicleCustomEquipmentv6ValueItem
	Upto20000        *EnumVehicleCustomEquipmentv6ValueItem
	GreaterThanAbove *EnumVehicleCustomEquipmentv6ValueItem

	itemDict map[string]*EnumVehicleCustomEquipmentv6ValueItem
}

// VehicleCustomEquipmentv6Value is a public singleton instance of EnumVehicleCustomEquipmentv6Value
// representing value levels for vehicle custom equipment v6
var VehicleCustomEquipmentv6Value = &EnumVehicleCustomEquipmentv6Value{
	Description: "value levels for vehicle custom equipment v6",
	Items: []*EnumVehicleCustomEquipmentv6ValueItem{
		&vehicleCustomEquipmentv6ValueUpto1000,
		&vehicleCustomEquipmentv6ValueFrom1001To1500,
		&vehicleCustomEquipmentv6ValueFrom1501To2000,
		&vehicleCustomEquipmentv6ValueFrom2001To2500,
		&vehicleCustomEquipmentv6ValueFrom2501To3000,
		&vehicleCustomEquipmentv6ValueFrom3001To3500,
		&vehicleCustomEquipmentv6ValueFrom3501To4000,
		&vehicleCustomEquipmentv6ValueFrom4001To4500,
		&vehicleCustomEquipmentv6ValueFrom4501To5000,
		&vehicleCustomEquipmentv6ValueUpto2000,
		&vehicleCustomEquipmentv6ValueUpto3000,
		&vehicleCustomEquipmentv6ValueUpto4000,
		&vehicleCustomEquipmentv6ValueUpto5000,
		&vehicleCustomEquipmentv6ValueUpto10000,
		&vehicleCustomEquipmentv6ValueUpto15000,
		&vehicleCustomEquipmentv6ValueUpto20000,
		&vehicleCustomEquipmentv6ValueGreaterThanAbove,
	},
	Name:             "EnumVehicleCustomEquipmentv6Value",
	Upto1000:         &vehicleCustomEquipmentv6ValueUpto1000,
	From1001To1500:   &vehicleCustomEquipmentv6ValueFrom1001To1500,
	From1501To2000:   &vehicleCustomEquipmentv6ValueFrom1501To2000,
	From2001To2500:   &vehicleCustomEquipmentv6ValueFrom2001To2500,
	From2501To3000:   &vehicleCustomEquipmentv6ValueFrom2501To3000,
	From3001To3500:   &vehicleCustomEquipmentv6ValueFrom3001To3500,
	From3501To4000:   &vehicleCustomEquipmentv6ValueFrom3501To4000,
	From4001To4500:   &vehicleCustomEquipmentv6ValueFrom4001To4500,
	From4501To5000:   &vehicleCustomEquipmentv6ValueFrom4501To5000,
	Upto2000:         &vehicleCustomEquipmentv6ValueUpto2000,
	Upto3000:         &vehicleCustomEquipmentv6ValueUpto3000,
	Upto4000:         &vehicleCustomEquipmentv6ValueUpto4000,
	Upto5000:         &vehicleCustomEquipmentv6ValueUpto5000,
	Upto10000:        &vehicleCustomEquipmentv6ValueUpto10000,
	Upto15000:        &vehicleCustomEquipmentv6ValueUpto15000,
	Upto20000:        &vehicleCustomEquipmentv6ValueUpto20000,
	GreaterThanAbove: &vehicleCustomEquipmentv6ValueGreaterThanAbove,

	itemDict: map[string]*EnumVehicleCustomEquipmentv6ValueItem{
		strings.ToLower(string(vehicleCustomEquipmentv6ValueUpto1000ID)):         &vehicleCustomEquipmentv6ValueUpto1000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueFrom1001To1500ID)):   &vehicleCustomEquipmentv6ValueFrom1001To1500,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueFrom1501To2000ID)):   &vehicleCustomEquipmentv6ValueFrom1501To2000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueFrom2001To2500ID)):   &vehicleCustomEquipmentv6ValueFrom2001To2500,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueFrom2501To3000ID)):   &vehicleCustomEquipmentv6ValueFrom2501To3000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueFrom3001To3500ID)):   &vehicleCustomEquipmentv6ValueFrom3001To3500,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueFrom3501To4000ID)):   &vehicleCustomEquipmentv6ValueFrom3501To4000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueFrom4001To4500ID)):   &vehicleCustomEquipmentv6ValueFrom4001To4500,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueFrom4501To5000ID)):   &vehicleCustomEquipmentv6ValueFrom4501To5000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueUpto2000ID)):         &vehicleCustomEquipmentv6ValueUpto2000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueUpto3000ID)):         &vehicleCustomEquipmentv6ValueUpto3000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueUpto4000ID)):         &vehicleCustomEquipmentv6ValueUpto4000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueUpto5000ID)):         &vehicleCustomEquipmentv6ValueUpto5000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueUpto10000ID)):        &vehicleCustomEquipmentv6ValueUpto10000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueUpto15000ID)):        &vehicleCustomEquipmentv6ValueUpto15000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueUpto20000ID)):        &vehicleCustomEquipmentv6ValueUpto20000,
		strings.ToLower(string(vehicleCustomEquipmentv6ValueGreaterThanAboveID)): &vehicleCustomEquipmentv6ValueGreaterThanAbove,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumVehicleCustomEquipmentv6Value) ByID(id VehicleCustomEquipmentv6ValueIdentifier) *EnumVehicleCustomEquipmentv6ValueItem {
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
func (e *EnumVehicleCustomEquipmentv6Value) ByIDString(idx string) *EnumVehicleCustomEquipmentv6ValueItem {
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
func (e *EnumVehicleCustomEquipmentv6Value) ByIndex(idx int) *EnumVehicleCustomEquipmentv6ValueItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedVehicleCustomEquipmentv6ValueID is a struct that is designed to replace a *VehicleCustomEquipmentv6ValueID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *VehicleCustomEquipmentv6ValueID it contains while being a better JSON citizen.
type ValidatedVehicleCustomEquipmentv6ValueID struct {
	// id will point to a valid VehicleCustomEquipmentv6ValueID, if possible
	// If id is nil, then ValidatedVehicleCustomEquipmentv6ValueID.Valid() will return false.
	id *VehicleCustomEquipmentv6ValueID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedVehicleCustomEquipmentv6ValueID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedVehicleCustomEquipmentv6ValueID
func (vi *ValidatedVehicleCustomEquipmentv6ValueID) Clone() *ValidatedVehicleCustomEquipmentv6ValueID {
	if vi == nil {
		return nil
	}

	var cid *VehicleCustomEquipmentv6ValueID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedVehicleCustomEquipmentv6ValueID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedVehicleCustomEquipmentv6ValueIds represent the same VehicleCustomEquipmentv6Value
func (vi *ValidatedVehicleCustomEquipmentv6ValueID) Equals(vj *ValidatedVehicleCustomEquipmentv6ValueID) bool {
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

// Valid returns true if and only if the ValidatedVehicleCustomEquipmentv6ValueID corresponds to a recognized VehicleCustomEquipmentv6Value
func (vi *ValidatedVehicleCustomEquipmentv6ValueID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedVehicleCustomEquipmentv6ValueID) ID() *VehicleCustomEquipmentv6ValueID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedVehicleCustomEquipmentv6ValueID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedVehicleCustomEquipmentv6ValueID) ValidatedID() *ValidatedVehicleCustomEquipmentv6ValueID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedVehicleCustomEquipmentv6ValueID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedVehicleCustomEquipmentv6ValueID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedVehicleCustomEquipmentv6ValueID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedVehicleCustomEquipmentv6ValueID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedVehicleCustomEquipmentv6ValueID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := VehicleCustomEquipmentv6ValueID(capString)
	item := VehicleCustomEquipmentv6Value.ByID(&id)
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

func (vi ValidatedVehicleCustomEquipmentv6ValueID) String() string {
	return vi.ToIDString()
}

type VehicleCustomEquipmentv6ValueIdentifier interface {
	ID() *VehicleCustomEquipmentv6ValueID
	Valid() bool
}
