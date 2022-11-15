package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// VehicleMileageID uniquely identifies a particular VehicleMileage
type VehicleMileageID string

// Clone creates a safe, independent copy of a VehicleMileageID
func (i *VehicleMileageID) Clone() *VehicleMileageID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two VehicleMileageIds are equivalent
func (i *VehicleMileageID) Equals(j *VehicleMileageID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *VehicleMileageID that is either valid or nil
func (i *VehicleMileageID) ID() *VehicleMileageID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *VehicleMileageID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the VehicleMileageID corresponds to a recognized VehicleMileage
func (i *VehicleMileageID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return VehicleMileage.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *VehicleMileageID) ValidatedID() *ValidatedVehicleMileageID {
	if i != nil {
		return &ValidatedVehicleMileageID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *VehicleMileageID) MarshalJSON() ([]byte, error) {
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

func (i *VehicleMileageID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := VehicleMileageID(dataString)
	item := VehicleMileage.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	vehicleMileageLessThan4000ID     VehicleMileageID = "less_than_4000"
	vehicleMileageFrom4000To5999ID   VehicleMileageID = "4000_to_5999"
	vehicleMileageFrom6000To7999ID   VehicleMileageID = "6000_to_7999"
	vehicleMileageFrom8000To9999ID   VehicleMileageID = "8000_to_9999"
	vehicleMileageFrom10000To11999ID VehicleMileageID = "10000_to_11999"
	vehicleMileageFrom12000To14999ID VehicleMileageID = "12000_to_14999"
	vehicleMileageFrom15000To19999ID VehicleMileageID = "15000_to_19999"
	vehicleMileageMoreThan20000ID    VehicleMileageID = "20000_or_more"
)

// EnumVehicleMileageItem describes an entry in an enumeration of VehicleMileage
type EnumVehicleMileageItem struct {
	ID        VehicleMileageID  `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	Max string
	Min string
}

var (
	vehicleMileageLessThan4000     = EnumVehicleMileageItem{vehicleMileageLessThan4000ID, "Less than 4,000", map[string]string{"Max": "3999", "Min": "0"}, "LessThan4000", 1, "3999", "0"}
	vehicleMileageFrom4000To5999   = EnumVehicleMileageItem{vehicleMileageFrom4000To5999ID, "4,000-5,999", map[string]string{"Max": "5999", "Min": "4000"}, "From4000To5999", 2, "5999", "4000"}
	vehicleMileageFrom6000To7999   = EnumVehicleMileageItem{vehicleMileageFrom6000To7999ID, "6,000-7,999", map[string]string{"Max": "7999", "Min": "6000"}, "From6000To7999", 3, "7999", "6000"}
	vehicleMileageFrom8000To9999   = EnumVehicleMileageItem{vehicleMileageFrom8000To9999ID, "8,000-9,999", map[string]string{"Max": "9999", "Min": "8000"}, "From8000To9999", 4, "9999", "8000"}
	vehicleMileageFrom10000To11999 = EnumVehicleMileageItem{vehicleMileageFrom10000To11999ID, "10,000-11,999", map[string]string{"Max": "11999", "Min": "10000"}, "From10000To11999", 5, "11999", "10000"}
	vehicleMileageFrom12000To14999 = EnumVehicleMileageItem{vehicleMileageFrom12000To14999ID, "12,000-14,999", map[string]string{"Max": "14999", "Min": "12000"}, "From12000To14999", 6, "14999", "12000"}
	vehicleMileageFrom15000To19999 = EnumVehicleMileageItem{vehicleMileageFrom15000To19999ID, "15,000-19,999", map[string]string{"Max": "19999", "Min": "15000"}, "From15000To19999", 7, "19999", "15000"}
	vehicleMileageMoreThan20000    = EnumVehicleMileageItem{vehicleMileageMoreThan20000ID, "20,000 or more", map[string]string{"Max": "1000000", "Min": "20000"}, "MoreThan20000", 8, "1000000", "20000"}
)

// EnumVehicleMileage is a collection of VehicleMileage items
type EnumVehicleMileage struct {
	Description string
	Items       []*EnumVehicleMileageItem
	Name        string

	LessThan4000     *EnumVehicleMileageItem
	From4000To5999   *EnumVehicleMileageItem
	From6000To7999   *EnumVehicleMileageItem
	From8000To9999   *EnumVehicleMileageItem
	From10000To11999 *EnumVehicleMileageItem
	From12000To14999 *EnumVehicleMileageItem
	From15000To19999 *EnumVehicleMileageItem
	MoreThan20000    *EnumVehicleMileageItem

	itemDict map[string]*EnumVehicleMileageItem
}

// VehicleMileage is a public singleton instance of EnumVehicleMileage
// representing ranges of vehicle mileage
var VehicleMileage = &EnumVehicleMileage{
	Description: "ranges of vehicle mileage",
	Items: []*EnumVehicleMileageItem{
		&vehicleMileageLessThan4000,
		&vehicleMileageFrom4000To5999,
		&vehicleMileageFrom6000To7999,
		&vehicleMileageFrom8000To9999,
		&vehicleMileageFrom10000To11999,
		&vehicleMileageFrom12000To14999,
		&vehicleMileageFrom15000To19999,
		&vehicleMileageMoreThan20000,
	},
	Name:             "EnumVehicleMileage",
	LessThan4000:     &vehicleMileageLessThan4000,
	From4000To5999:   &vehicleMileageFrom4000To5999,
	From6000To7999:   &vehicleMileageFrom6000To7999,
	From8000To9999:   &vehicleMileageFrom8000To9999,
	From10000To11999: &vehicleMileageFrom10000To11999,
	From12000To14999: &vehicleMileageFrom12000To14999,
	From15000To19999: &vehicleMileageFrom15000To19999,
	MoreThan20000:    &vehicleMileageMoreThan20000,

	itemDict: map[string]*EnumVehicleMileageItem{
		strings.ToLower(string(vehicleMileageLessThan4000ID)):     &vehicleMileageLessThan4000,
		strings.ToLower(string(vehicleMileageFrom4000To5999ID)):   &vehicleMileageFrom4000To5999,
		strings.ToLower(string(vehicleMileageFrom6000To7999ID)):   &vehicleMileageFrom6000To7999,
		strings.ToLower(string(vehicleMileageFrom8000To9999ID)):   &vehicleMileageFrom8000To9999,
		strings.ToLower(string(vehicleMileageFrom10000To11999ID)): &vehicleMileageFrom10000To11999,
		strings.ToLower(string(vehicleMileageFrom12000To14999ID)): &vehicleMileageFrom12000To14999,
		strings.ToLower(string(vehicleMileageFrom15000To19999ID)): &vehicleMileageFrom15000To19999,
		strings.ToLower(string(vehicleMileageMoreThan20000ID)):    &vehicleMileageMoreThan20000,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumVehicleMileage) ByID(id VehicleMileageIdentifier) *EnumVehicleMileageItem {
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
func (e *EnumVehicleMileage) ByIDString(idx string) *EnumVehicleMileageItem {
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
func (e *EnumVehicleMileage) ByIndex(idx int) *EnumVehicleMileageItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedVehicleMileageID is a struct that is designed to replace a *VehicleMileageID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *VehicleMileageID it contains while being a better JSON citizen.
type ValidatedVehicleMileageID struct {
	// id will point to a valid VehicleMileageID, if possible
	// If id is nil, then ValidatedVehicleMileageID.Valid() will return false.
	id *VehicleMileageID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedVehicleMileageID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedVehicleMileageID
func (vi *ValidatedVehicleMileageID) Clone() *ValidatedVehicleMileageID {
	if vi == nil {
		return nil
	}

	var cid *VehicleMileageID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedVehicleMileageID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedVehicleMileageIds represent the same VehicleMileage
func (vi *ValidatedVehicleMileageID) Equals(vj *ValidatedVehicleMileageID) bool {
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

// Valid returns true if and only if the ValidatedVehicleMileageID corresponds to a recognized VehicleMileage
func (vi *ValidatedVehicleMileageID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedVehicleMileageID) ID() *VehicleMileageID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedVehicleMileageID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedVehicleMileageID) ValidatedID() *ValidatedVehicleMileageID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedVehicleMileageID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedVehicleMileageID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedVehicleMileageID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedVehicleMileageID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedVehicleMileageID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := VehicleMileageID(capString)
	item := VehicleMileage.ByID(&id)
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

func (vi ValidatedVehicleMileageID) String() string {
	return vi.ToIDString()
}

type VehicleMileageIdentifier interface {
	ID() *VehicleMileageID
	Valid() bool
}
