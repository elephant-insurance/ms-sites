package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// RideSharingHoursID uniquely identifies a particular RideSharingHours
type RideSharingHoursID string

// Clone creates a safe, independent copy of a RideSharingHoursID
func (i *RideSharingHoursID) Clone() *RideSharingHoursID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two RideSharingHoursIds are equivalent
func (i *RideSharingHoursID) Equals(j *RideSharingHoursID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *RideSharingHoursID that is either valid or nil
func (i *RideSharingHoursID) ID() *RideSharingHoursID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *RideSharingHoursID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the RideSharingHoursID corresponds to a recognized RideSharingHours
func (i *RideSharingHoursID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return RideSharingHours.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *RideSharingHoursID) ValidatedID() *ValidatedRideSharingHoursID {
	if i != nil {
		return &ValidatedRideSharingHoursID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *RideSharingHoursID) MarshalJSON() ([]byte, error) {
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

func (i *RideSharingHoursID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := RideSharingHoursID(dataString)
	item := RideSharingHours.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	rideSharingHoursOneToFiveID             RideSharingHoursID = "1to5"
	rideSharingHoursSixToTenID              RideSharingHoursID = "6to10"
	rideSharingHoursElevenToFifteenID       RideSharingHoursID = "11to15"
	rideSharingHoursSixteenToTwentyID       RideSharingHoursID = "16to20"
	rideSharingHoursTwentyOneToTwentyFiveID RideSharingHoursID = "21to25"
	rideSharingHoursTwentysixToThirtyID     RideSharingHoursID = "26to30"
	rideSharingHoursThirtyOneToThirtyFiveID RideSharingHoursID = "31to35"
	rideSharingHoursThirtySixToFortyID      RideSharingHoursID = "36to40"
	rideSharingHoursFortyOnePlusID          RideSharingHoursID = "41andup"
)

// EnumRideSharingHoursItem describes an entry in an enumeration of RideSharingHours
type EnumRideSharingHoursItem struct {
	ID        RideSharingHoursID `json:"Value"`
	Desc      string             `json:"Description,omitempty"`
	Meta      map[string]string  `json:",omitempty"`
	Name      string             `json:"Name"`
	SortOrder int
}

var (
	rideSharingHoursOneToFive             = EnumRideSharingHoursItem{rideSharingHoursOneToFiveID, "1-5", nil, "OneToFive", 1}
	rideSharingHoursSixToTen              = EnumRideSharingHoursItem{rideSharingHoursSixToTenID, "6-10", nil, "SixToTen", 2}
	rideSharingHoursElevenToFifteen       = EnumRideSharingHoursItem{rideSharingHoursElevenToFifteenID, "11-15", nil, "ElevenToFifteen", 3}
	rideSharingHoursSixteenToTwenty       = EnumRideSharingHoursItem{rideSharingHoursSixteenToTwentyID, "16-20", nil, "SixteenToTwenty", 4}
	rideSharingHoursTwentyOneToTwentyFive = EnumRideSharingHoursItem{rideSharingHoursTwentyOneToTwentyFiveID, "21-25", nil, "TwentyOneToTwentyFive", 5}
	rideSharingHoursTwentysixToThirty     = EnumRideSharingHoursItem{rideSharingHoursTwentysixToThirtyID, "26-30", nil, "TwentysixToThirty", 6}
	rideSharingHoursThirtyOneToThirtyFive = EnumRideSharingHoursItem{rideSharingHoursThirtyOneToThirtyFiveID, "31-35", nil, "ThirtyOneToThirtyFive", 7}
	rideSharingHoursThirtySixToForty      = EnumRideSharingHoursItem{rideSharingHoursThirtySixToFortyID, "36-40", nil, "ThirtySixToForty", 8}
	rideSharingHoursFortyOnePlus          = EnumRideSharingHoursItem{rideSharingHoursFortyOnePlusID, "41+", nil, "FortyOnePlus", 9}
)

// EnumRideSharingHours is a collection of RideSharingHours items
type EnumRideSharingHours struct {
	Description string
	Items       []*EnumRideSharingHoursItem
	Name        string

	OneToFive             *EnumRideSharingHoursItem
	SixToTen              *EnumRideSharingHoursItem
	ElevenToFifteen       *EnumRideSharingHoursItem
	SixteenToTwenty       *EnumRideSharingHoursItem
	TwentyOneToTwentyFive *EnumRideSharingHoursItem
	TwentysixToThirty     *EnumRideSharingHoursItem
	ThirtyOneToThirtyFive *EnumRideSharingHoursItem
	ThirtySixToForty      *EnumRideSharingHoursItem
	FortyOnePlus          *EnumRideSharingHoursItem

	itemDict map[string]*EnumRideSharingHoursItem
}

// RideSharingHours is a public singleton instance of EnumRideSharingHours
// representing ranges of hours spent ridesharing per week
var RideSharingHours = &EnumRideSharingHours{
	Description: "ranges of hours spent ridesharing per week",
	Items: []*EnumRideSharingHoursItem{
		&rideSharingHoursOneToFive,
		&rideSharingHoursSixToTen,
		&rideSharingHoursElevenToFifteen,
		&rideSharingHoursSixteenToTwenty,
		&rideSharingHoursTwentyOneToTwentyFive,
		&rideSharingHoursTwentysixToThirty,
		&rideSharingHoursThirtyOneToThirtyFive,
		&rideSharingHoursThirtySixToForty,
		&rideSharingHoursFortyOnePlus,
	},
	Name:                  "EnumRideSharingHours",
	OneToFive:             &rideSharingHoursOneToFive,
	SixToTen:              &rideSharingHoursSixToTen,
	ElevenToFifteen:       &rideSharingHoursElevenToFifteen,
	SixteenToTwenty:       &rideSharingHoursSixteenToTwenty,
	TwentyOneToTwentyFive: &rideSharingHoursTwentyOneToTwentyFive,
	TwentysixToThirty:     &rideSharingHoursTwentysixToThirty,
	ThirtyOneToThirtyFive: &rideSharingHoursThirtyOneToThirtyFive,
	ThirtySixToForty:      &rideSharingHoursThirtySixToForty,
	FortyOnePlus:          &rideSharingHoursFortyOnePlus,

	itemDict: map[string]*EnumRideSharingHoursItem{
		strings.ToLower(string(rideSharingHoursOneToFiveID)):             &rideSharingHoursOneToFive,
		strings.ToLower(string(rideSharingHoursSixToTenID)):              &rideSharingHoursSixToTen,
		strings.ToLower(string(rideSharingHoursElevenToFifteenID)):       &rideSharingHoursElevenToFifteen,
		strings.ToLower(string(rideSharingHoursSixteenToTwentyID)):       &rideSharingHoursSixteenToTwenty,
		strings.ToLower(string(rideSharingHoursTwentyOneToTwentyFiveID)): &rideSharingHoursTwentyOneToTwentyFive,
		strings.ToLower(string(rideSharingHoursTwentysixToThirtyID)):     &rideSharingHoursTwentysixToThirty,
		strings.ToLower(string(rideSharingHoursThirtyOneToThirtyFiveID)): &rideSharingHoursThirtyOneToThirtyFive,
		strings.ToLower(string(rideSharingHoursThirtySixToFortyID)):      &rideSharingHoursThirtySixToForty,
		strings.ToLower(string(rideSharingHoursFortyOnePlusID)):          &rideSharingHoursFortyOnePlus,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumRideSharingHours) ByID(id RideSharingHoursIdentifier) *EnumRideSharingHoursItem {
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
func (e *EnumRideSharingHours) ByIDString(idx string) *EnumRideSharingHoursItem {
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
func (e *EnumRideSharingHours) ByIndex(idx int) *EnumRideSharingHoursItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedRideSharingHoursID is a struct that is designed to replace a *RideSharingHoursID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *RideSharingHoursID it contains while being a better JSON citizen.
type ValidatedRideSharingHoursID struct {
	// id will point to a valid RideSharingHoursID, if possible
	// If id is nil, then ValidatedRideSharingHoursID.Valid() will return false.
	id *RideSharingHoursID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedRideSharingHoursID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedRideSharingHoursID
func (vi *ValidatedRideSharingHoursID) Clone() *ValidatedRideSharingHoursID {
	if vi == nil {
		return nil
	}

	var cid *RideSharingHoursID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedRideSharingHoursID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedRideSharingHoursIds represent the same RideSharingHours
func (vi *ValidatedRideSharingHoursID) Equals(vj *ValidatedRideSharingHoursID) bool {
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

// Valid returns true if and only if the ValidatedRideSharingHoursID corresponds to a recognized RideSharingHours
func (vi *ValidatedRideSharingHoursID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedRideSharingHoursID) ID() *RideSharingHoursID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedRideSharingHoursID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedRideSharingHoursID) ValidatedID() *ValidatedRideSharingHoursID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedRideSharingHoursID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedRideSharingHoursID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedRideSharingHoursID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedRideSharingHoursID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedRideSharingHoursID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := RideSharingHoursID(capString)
	item := RideSharingHours.ByID(&id)
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

func (vi ValidatedRideSharingHoursID) String() string {
	return vi.ToIDString()
}

type RideSharingHoursIdentifier interface {
	ID() *RideSharingHoursID
	Valid() bool
}
