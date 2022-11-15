package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// RatingStatusID uniquely identifies a particular RatingStatus
type RatingStatusID string

// Clone creates a safe, independent copy of a RatingStatusID
func (i *RatingStatusID) Clone() *RatingStatusID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two RatingStatusIds are equivalent
func (i *RatingStatusID) Equals(j *RatingStatusID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *RatingStatusID that is either valid or nil
func (i *RatingStatusID) ID() *RatingStatusID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *RatingStatusID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the RatingStatusID corresponds to a recognized RatingStatus
func (i *RatingStatusID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return RatingStatus.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *RatingStatusID) ValidatedID() *ValidatedRatingStatusID {
	if i != nil {
		return &ValidatedRatingStatusID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *RatingStatusID) MarshalJSON() ([]byte, error) {
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

func (i *RatingStatusID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := RatingStatusID(dataString)
	item := RatingStatus.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	ratingStatusUnknownID                 RatingStatusID = "Unknown"
	ratingStatusRatedID                   RatingStatusID = "rated"
	ratingStatusInsuredElsewhereID        RatingStatusID = "insuredelsewhere"
	ratingStatusListedID                  RatingStatusID = "listed"
	ratingStatusUnratedResidentRelativeID RatingStatusID = "unratedresidentrelative"
	ratingStatusUnacceptableRiskID        RatingStatusID = "unacceptablerisk"
	ratingStatusExcludedDriverID          RatingStatusID = "excludeddriver"
)

// EnumRatingStatusItem describes an entry in an enumeration of RatingStatus
type EnumRatingStatusItem struct {
	ID        RatingStatusID    `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	ratingStatusUnknown                 = EnumRatingStatusItem{ratingStatusUnknownID, "Unknown", nil, "Unknown", 1}
	ratingStatusRated                   = EnumRatingStatusItem{ratingStatusRatedID, "Rated", nil, "Rated", 2}
	ratingStatusInsuredElsewhere        = EnumRatingStatusItem{ratingStatusInsuredElsewhereID, "InsuredElsewhere", nil, "InsuredElsewhere", 3}
	ratingStatusListed                  = EnumRatingStatusItem{ratingStatusListedID, "Listed", nil, "Listed", 4}
	ratingStatusUnratedResidentRelative = EnumRatingStatusItem{ratingStatusUnratedResidentRelativeID, "UnratedResidentRelative", nil, "UnratedResidentRelative", 5}
	ratingStatusUnacceptableRisk        = EnumRatingStatusItem{ratingStatusUnacceptableRiskID, "UnacceptableRisk", nil, "UnacceptableRisk", 6}
	ratingStatusExcludedDriver          = EnumRatingStatusItem{ratingStatusExcludedDriverID, "ExcludedDriver", nil, "ExcludedDriver", 7}
)

// EnumRatingStatus is a collection of RatingStatus items
type EnumRatingStatus struct {
	Description string
	Items       []*EnumRatingStatusItem
	Name        string

	Unknown                 *EnumRatingStatusItem
	Rated                   *EnumRatingStatusItem
	InsuredElsewhere        *EnumRatingStatusItem
	Listed                  *EnumRatingStatusItem
	UnratedResidentRelative *EnumRatingStatusItem
	UnacceptableRisk        *EnumRatingStatusItem
	ExcludedDriver          *EnumRatingStatusItem

	itemDict map[string]*EnumRatingStatusItem
}

// RatingStatus is a public singleton instance of EnumRatingStatus
// representing rating statuses of drivers
var RatingStatus = &EnumRatingStatus{
	Description: "rating statuses of drivers",
	Items: []*EnumRatingStatusItem{
		&ratingStatusUnknown,
		&ratingStatusRated,
		&ratingStatusInsuredElsewhere,
		&ratingStatusListed,
		&ratingStatusUnratedResidentRelative,
		&ratingStatusUnacceptableRisk,
		&ratingStatusExcludedDriver,
	},
	Name:                    "EnumRatingStatus",
	Unknown:                 &ratingStatusUnknown,
	Rated:                   &ratingStatusRated,
	InsuredElsewhere:        &ratingStatusInsuredElsewhere,
	Listed:                  &ratingStatusListed,
	UnratedResidentRelative: &ratingStatusUnratedResidentRelative,
	UnacceptableRisk:        &ratingStatusUnacceptableRisk,
	ExcludedDriver:          &ratingStatusExcludedDriver,

	itemDict: map[string]*EnumRatingStatusItem{
		strings.ToLower(string(ratingStatusUnknownID)):                 &ratingStatusUnknown,
		strings.ToLower(string(ratingStatusRatedID)):                   &ratingStatusRated,
		strings.ToLower(string(ratingStatusInsuredElsewhereID)):        &ratingStatusInsuredElsewhere,
		strings.ToLower(string(ratingStatusListedID)):                  &ratingStatusListed,
		strings.ToLower(string(ratingStatusUnratedResidentRelativeID)): &ratingStatusUnratedResidentRelative,
		strings.ToLower(string(ratingStatusUnacceptableRiskID)):        &ratingStatusUnacceptableRisk,
		strings.ToLower(string(ratingStatusExcludedDriverID)):          &ratingStatusExcludedDriver,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumRatingStatus) ByID(id RatingStatusIdentifier) *EnumRatingStatusItem {
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
func (e *EnumRatingStatus) ByIDString(idx string) *EnumRatingStatusItem {
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
func (e *EnumRatingStatus) ByIndex(idx int) *EnumRatingStatusItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedRatingStatusID is a struct that is designed to replace a *RatingStatusID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *RatingStatusID it contains while being a better JSON citizen.
type ValidatedRatingStatusID struct {
	// id will point to a valid RatingStatusID, if possible
	// If id is nil, then ValidatedRatingStatusID.Valid() will return false.
	id *RatingStatusID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedRatingStatusID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedRatingStatusID
func (vi *ValidatedRatingStatusID) Clone() *ValidatedRatingStatusID {
	if vi == nil {
		return nil
	}

	var cid *RatingStatusID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedRatingStatusID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedRatingStatusIds represent the same RatingStatus
func (vi *ValidatedRatingStatusID) Equals(vj *ValidatedRatingStatusID) bool {
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

// Valid returns true if and only if the ValidatedRatingStatusID corresponds to a recognized RatingStatus
func (vi *ValidatedRatingStatusID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedRatingStatusID) ID() *RatingStatusID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedRatingStatusID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedRatingStatusID) ValidatedID() *ValidatedRatingStatusID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedRatingStatusID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedRatingStatusID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedRatingStatusID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedRatingStatusID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedRatingStatusID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := RatingStatusID(capString)
	item := RatingStatus.ByID(&id)
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

func (vi ValidatedRatingStatusID) String() string {
	return vi.ToIDString()
}

type RatingStatusIdentifier interface {
	ID() *RatingStatusID
	Valid() bool
}
