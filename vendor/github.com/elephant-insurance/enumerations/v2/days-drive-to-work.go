package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// DaysDriveToWorkID uniquely identifies a particular DaysDriveToWork
type DaysDriveToWorkID string

// Clone creates a safe, independent copy of a DaysDriveToWorkID
func (i *DaysDriveToWorkID) Clone() *DaysDriveToWorkID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two DaysDriveToWorkIds are equivalent
func (i *DaysDriveToWorkID) Equals(j *DaysDriveToWorkID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *DaysDriveToWorkID that is either valid or nil
func (i *DaysDriveToWorkID) ID() *DaysDriveToWorkID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *DaysDriveToWorkID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the DaysDriveToWorkID corresponds to a recognized DaysDriveToWork
func (i *DaysDriveToWorkID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return DaysDriveToWork.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *DaysDriveToWorkID) ValidatedID() *ValidatedDaysDriveToWorkID {
	if i != nil {
		return &ValidatedDaysDriveToWorkID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *DaysDriveToWorkID) MarshalJSON() ([]byte, error) {
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

func (i *DaysDriveToWorkID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := DaysDriveToWorkID(dataString)
	item := DaysDriveToWork.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	daysDriveToWorkZeroID  DaysDriveToWorkID = "0"
	daysDriveToWorkOneID   DaysDriveToWorkID = "1"
	daysDriveToWorkTwoID   DaysDriveToWorkID = "2"
	daysDriveToWorkThreeID DaysDriveToWorkID = "3"
	daysDriveToWorkFourID  DaysDriveToWorkID = "4"
	daysDriveToWorkFiveID  DaysDriveToWorkID = "5OrMore"
)

// EnumDaysDriveToWorkItem describes an entry in an enumeration of DaysDriveToWork
type EnumDaysDriveToWorkItem struct {
	ID        DaysDriveToWorkID `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	daysDriveToWorkZero  = EnumDaysDriveToWorkItem{daysDriveToWorkZeroID, "0", nil, "Zero", 1}
	daysDriveToWorkOne   = EnumDaysDriveToWorkItem{daysDriveToWorkOneID, "1", nil, "One", 2}
	daysDriveToWorkTwo   = EnumDaysDriveToWorkItem{daysDriveToWorkTwoID, "2", nil, "Two", 3}
	daysDriveToWorkThree = EnumDaysDriveToWorkItem{daysDriveToWorkThreeID, "3", nil, "Three", 4}
	daysDriveToWorkFour  = EnumDaysDriveToWorkItem{daysDriveToWorkFourID, "4", nil, "Four", 5}
	daysDriveToWorkFive  = EnumDaysDriveToWorkItem{daysDriveToWorkFiveID, "5 or more", nil, "Five", 6}
)

// EnumDaysDriveToWork is a collection of DaysDriveToWork items
type EnumDaysDriveToWork struct {
	Description string
	Items       []*EnumDaysDriveToWorkItem
	Name        string

	Zero  *EnumDaysDriveToWorkItem
	One   *EnumDaysDriveToWorkItem
	Two   *EnumDaysDriveToWorkItem
	Three *EnumDaysDriveToWorkItem
	Four  *EnumDaysDriveToWorkItem
	Five  *EnumDaysDriveToWorkItem

	itemDict map[string]*EnumDaysDriveToWorkItem
}

// DaysDriveToWork is a public singleton instance of EnumDaysDriveToWork
// representing ranges of days per week a customer drives a vehicle
var DaysDriveToWork = &EnumDaysDriveToWork{
	Description: "ranges of days per week a customer drives a vehicle",
	Items: []*EnumDaysDriveToWorkItem{
		&daysDriveToWorkZero,
		&daysDriveToWorkOne,
		&daysDriveToWorkTwo,
		&daysDriveToWorkThree,
		&daysDriveToWorkFour,
		&daysDriveToWorkFive,
	},
	Name:  "EnumDaysDriveToWork",
	Zero:  &daysDriveToWorkZero,
	One:   &daysDriveToWorkOne,
	Two:   &daysDriveToWorkTwo,
	Three: &daysDriveToWorkThree,
	Four:  &daysDriveToWorkFour,
	Five:  &daysDriveToWorkFive,

	itemDict: map[string]*EnumDaysDriveToWorkItem{
		strings.ToLower(string(daysDriveToWorkZeroID)):  &daysDriveToWorkZero,
		strings.ToLower(string(daysDriveToWorkOneID)):   &daysDriveToWorkOne,
		strings.ToLower(string(daysDriveToWorkTwoID)):   &daysDriveToWorkTwo,
		strings.ToLower(string(daysDriveToWorkThreeID)): &daysDriveToWorkThree,
		strings.ToLower(string(daysDriveToWorkFourID)):  &daysDriveToWorkFour,
		strings.ToLower(string(daysDriveToWorkFiveID)):  &daysDriveToWorkFive,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumDaysDriveToWork) ByID(id DaysDriveToWorkIdentifier) *EnumDaysDriveToWorkItem {
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
func (e *EnumDaysDriveToWork) ByIDString(idx string) *EnumDaysDriveToWorkItem {
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
func (e *EnumDaysDriveToWork) ByIndex(idx int) *EnumDaysDriveToWorkItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedDaysDriveToWorkID is a struct that is designed to replace a *DaysDriveToWorkID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *DaysDriveToWorkID it contains while being a better JSON citizen.
type ValidatedDaysDriveToWorkID struct {
	// id will point to a valid DaysDriveToWorkID, if possible
	// If id is nil, then ValidatedDaysDriveToWorkID.Valid() will return false.
	id *DaysDriveToWorkID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedDaysDriveToWorkID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedDaysDriveToWorkID
func (vi *ValidatedDaysDriveToWorkID) Clone() *ValidatedDaysDriveToWorkID {
	if vi == nil {
		return nil
	}

	var cid *DaysDriveToWorkID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedDaysDriveToWorkID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedDaysDriveToWorkIds represent the same DaysDriveToWork
func (vi *ValidatedDaysDriveToWorkID) Equals(vj *ValidatedDaysDriveToWorkID) bool {
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

// Valid returns true if and only if the ValidatedDaysDriveToWorkID corresponds to a recognized DaysDriveToWork
func (vi *ValidatedDaysDriveToWorkID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedDaysDriveToWorkID) ID() *DaysDriveToWorkID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedDaysDriveToWorkID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedDaysDriveToWorkID) ValidatedID() *ValidatedDaysDriveToWorkID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedDaysDriveToWorkID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedDaysDriveToWorkID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedDaysDriveToWorkID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedDaysDriveToWorkID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedDaysDriveToWorkID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := DaysDriveToWorkID(capString)
	item := DaysDriveToWork.ByID(&id)
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

func (vi ValidatedDaysDriveToWorkID) String() string {
	return vi.ToIDString()
}

type DaysDriveToWorkIdentifier interface {
	ID() *DaysDriveToWorkID
	Valid() bool
}
