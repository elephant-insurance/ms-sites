package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// CurrentInsuranceLimitID uniquely identifies a particular CurrentInsuranceLimit
type CurrentInsuranceLimitID string

// Clone creates a safe, independent copy of a CurrentInsuranceLimitID
func (i *CurrentInsuranceLimitID) Clone() *CurrentInsuranceLimitID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two CurrentInsuranceLimitIds are equivalent
func (i *CurrentInsuranceLimitID) Equals(j *CurrentInsuranceLimitID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *CurrentInsuranceLimitID that is either valid or nil
func (i *CurrentInsuranceLimitID) ID() *CurrentInsuranceLimitID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *CurrentInsuranceLimitID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the CurrentInsuranceLimitID corresponds to a recognized CurrentInsuranceLimit
func (i *CurrentInsuranceLimitID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return CurrentInsuranceLimit.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *CurrentInsuranceLimitID) ValidatedID() *ValidatedCurrentInsuranceLimitID {
	if i != nil {
		return &ValidatedCurrentInsuranceLimitID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *CurrentInsuranceLimitID) MarshalJSON() ([]byte, error) {
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

func (i *CurrentInsuranceLimitID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := CurrentInsuranceLimitID(dataString)
	item := CurrentInsuranceLimit.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	currentInsuranceLimitDontKnowID        CurrentInsuranceLimitID = "unknown"
	currentInsuranceLimitMinimumID         CurrentInsuranceLimitID = "minimum_limit"
	currentInsuranceLimitLessThan50_100ID  CurrentInsuranceLimitID = "less_50_100"
	currentInsuranceLimitLessThan100_300ID CurrentInsuranceLimitID = "less_100_300"
	currentInsuranceLimitMoreThan100_300ID CurrentInsuranceLimitID = "100_300_greater"
)

// EnumCurrentInsuranceLimitItem describes an entry in an enumeration of CurrentInsuranceLimit
type EnumCurrentInsuranceLimitItem struct {
	ID        CurrentInsuranceLimitID `json:"Value"`
	Desc      string                  `json:"Description,omitempty"`
	Meta      map[string]string       `json:",omitempty"`
	Name      string                  `json:"Name"`
	SortOrder int
}

var (
	currentInsuranceLimitDontKnow        = EnumCurrentInsuranceLimitItem{currentInsuranceLimitDontKnowID, "Don't Know", nil, "DontKnow", 1}
	currentInsuranceLimitMinimum         = EnumCurrentInsuranceLimitItem{currentInsuranceLimitMinimumID, "Minimum Limit", nil, "Minimum", 2}
	currentInsuranceLimitLessThan50_100  = EnumCurrentInsuranceLimitItem{currentInsuranceLimitLessThan50_100ID, "More Than Minimum but Less Than 50/100", nil, "LessThan50_100", 3}
	currentInsuranceLimitLessThan100_300 = EnumCurrentInsuranceLimitItem{currentInsuranceLimitLessThan100_300ID, "50/100 or More but Less Than 100/300", nil, "LessThan100_300", 4}
	currentInsuranceLimitMoreThan100_300 = EnumCurrentInsuranceLimitItem{currentInsuranceLimitMoreThan100_300ID, "100/300 or More", nil, "MoreThan100_300", 5}
)

// EnumCurrentInsuranceLimit is a collection of CurrentInsuranceLimit items
type EnumCurrentInsuranceLimit struct {
	Description string
	Items       []*EnumCurrentInsuranceLimitItem
	Name        string

	DontKnow        *EnumCurrentInsuranceLimitItem
	Minimum         *EnumCurrentInsuranceLimitItem
	LessThan50_100  *EnumCurrentInsuranceLimitItem
	LessThan100_300 *EnumCurrentInsuranceLimitItem
	MoreThan100_300 *EnumCurrentInsuranceLimitItem

	itemDict map[string]*EnumCurrentInsuranceLimitItem
}

// CurrentInsuranceLimit is a public singleton instance of EnumCurrentInsuranceLimit
// representing coverage limits on a customer's current policy
var CurrentInsuranceLimit = &EnumCurrentInsuranceLimit{
	Description: "coverage limits on a customer's current policy",
	Items: []*EnumCurrentInsuranceLimitItem{
		&currentInsuranceLimitDontKnow,
		&currentInsuranceLimitMinimum,
		&currentInsuranceLimitLessThan50_100,
		&currentInsuranceLimitLessThan100_300,
		&currentInsuranceLimitMoreThan100_300,
	},
	Name:            "EnumCurrentInsuranceLimit",
	DontKnow:        &currentInsuranceLimitDontKnow,
	Minimum:         &currentInsuranceLimitMinimum,
	LessThan50_100:  &currentInsuranceLimitLessThan50_100,
	LessThan100_300: &currentInsuranceLimitLessThan100_300,
	MoreThan100_300: &currentInsuranceLimitMoreThan100_300,

	itemDict: map[string]*EnumCurrentInsuranceLimitItem{
		strings.ToLower(string(currentInsuranceLimitDontKnowID)):        &currentInsuranceLimitDontKnow,
		strings.ToLower(string(currentInsuranceLimitMinimumID)):         &currentInsuranceLimitMinimum,
		strings.ToLower(string(currentInsuranceLimitLessThan50_100ID)):  &currentInsuranceLimitLessThan50_100,
		strings.ToLower(string(currentInsuranceLimitLessThan100_300ID)): &currentInsuranceLimitLessThan100_300,
		strings.ToLower(string(currentInsuranceLimitMoreThan100_300ID)): &currentInsuranceLimitMoreThan100_300,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumCurrentInsuranceLimit) ByID(id CurrentInsuranceLimitIdentifier) *EnumCurrentInsuranceLimitItem {
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
func (e *EnumCurrentInsuranceLimit) ByIDString(idx string) *EnumCurrentInsuranceLimitItem {
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
func (e *EnumCurrentInsuranceLimit) ByIndex(idx int) *EnumCurrentInsuranceLimitItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedCurrentInsuranceLimitID is a struct that is designed to replace a *CurrentInsuranceLimitID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *CurrentInsuranceLimitID it contains while being a better JSON citizen.
type ValidatedCurrentInsuranceLimitID struct {
	// id will point to a valid CurrentInsuranceLimitID, if possible
	// If id is nil, then ValidatedCurrentInsuranceLimitID.Valid() will return false.
	id *CurrentInsuranceLimitID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedCurrentInsuranceLimitID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedCurrentInsuranceLimitID
func (vi *ValidatedCurrentInsuranceLimitID) Clone() *ValidatedCurrentInsuranceLimitID {
	if vi == nil {
		return nil
	}

	var cid *CurrentInsuranceLimitID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedCurrentInsuranceLimitID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedCurrentInsuranceLimitIds represent the same CurrentInsuranceLimit
func (vi *ValidatedCurrentInsuranceLimitID) Equals(vj *ValidatedCurrentInsuranceLimitID) bool {
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

// Valid returns true if and only if the ValidatedCurrentInsuranceLimitID corresponds to a recognized CurrentInsuranceLimit
func (vi *ValidatedCurrentInsuranceLimitID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedCurrentInsuranceLimitID) ID() *CurrentInsuranceLimitID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedCurrentInsuranceLimitID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedCurrentInsuranceLimitID) ValidatedID() *ValidatedCurrentInsuranceLimitID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedCurrentInsuranceLimitID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedCurrentInsuranceLimitID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedCurrentInsuranceLimitID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedCurrentInsuranceLimitID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedCurrentInsuranceLimitID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := CurrentInsuranceLimitID(capString)
	item := CurrentInsuranceLimit.ByID(&id)
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

func (vi ValidatedCurrentInsuranceLimitID) String() string {
	return vi.ToIDString()
}

type CurrentInsuranceLimitIdentifier interface {
	ID() *CurrentInsuranceLimitID
	Valid() bool
}
