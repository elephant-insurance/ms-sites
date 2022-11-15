package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// CurrentInsuranceStatusID uniquely identifies a particular CurrentInsuranceStatus
type CurrentInsuranceStatusID string

// Clone creates a safe, independent copy of a CurrentInsuranceStatusID
func (i *CurrentInsuranceStatusID) Clone() *CurrentInsuranceStatusID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two CurrentInsuranceStatusIds are equivalent
func (i *CurrentInsuranceStatusID) Equals(j *CurrentInsuranceStatusID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *CurrentInsuranceStatusID that is either valid or nil
func (i *CurrentInsuranceStatusID) ID() *CurrentInsuranceStatusID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *CurrentInsuranceStatusID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the CurrentInsuranceStatusID corresponds to a recognized CurrentInsuranceStatus
func (i *CurrentInsuranceStatusID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return CurrentInsuranceStatus.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *CurrentInsuranceStatusID) ValidatedID() *ValidatedCurrentInsuranceStatusID {
	if i != nil {
		return &ValidatedCurrentInsuranceStatusID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *CurrentInsuranceStatusID) MarshalJSON() ([]byte, error) {
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

func (i *CurrentInsuranceStatusID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := CurrentInsuranceStatusID(dataString)
	item := CurrentInsuranceStatus.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	currentInsuranceStatusOwnPolicyID           CurrentInsuranceStatusID = "own_policy"
	currentInsuranceStatusAnothersPolicyID      CurrentInsuranceStatusID = "on_anothers_policy"
	currentInsuranceStatusDeployedOverseasID    CurrentInsuranceStatusID = "military_overseas"
	currentInsuranceStatusExpiredWithin30DaysID CurrentInsuranceStatusID = "policy_expired_within_30days"
	currentInsuranceStatusExpiredOver30DaysID   CurrentInsuranceStatusID = "policy_expired_over_30days"
	currentInsuranceStatusNoInsuranceRequiredID CurrentInsuranceStatusID = "just_acquired_auto"
)

// EnumCurrentInsuranceStatusItem describes an entry in an enumeration of CurrentInsuranceStatus
type EnumCurrentInsuranceStatusItem struct {
	ID        CurrentInsuranceStatusID `json:"Value"`
	Desc      string                   `json:"Description,omitempty"`
	Meta      map[string]string        `json:",omitempty"`
	Name      string                   `json:"Name"`
	SortOrder int
}

var (
	currentInsuranceStatusOwnPolicy           = EnumCurrentInsuranceStatusItem{currentInsuranceStatusOwnPolicyID, "Yes, I have my own Policy", nil, "OwnPolicy", 1}
	currentInsuranceStatusAnothersPolicy      = EnumCurrentInsuranceStatusItem{currentInsuranceStatusAnothersPolicyID, "Yes, on anothers policy", nil, "AnothersPolicy", 2}
	currentInsuranceStatusDeployedOverseas    = EnumCurrentInsuranceStatusItem{currentInsuranceStatusDeployedOverseasID, "Deployed Overseas with Military", nil, "DeployedOverseas", 3}
	currentInsuranceStatusExpiredWithin30Days = EnumCurrentInsuranceStatusItem{currentInsuranceStatusExpiredWithin30DaysID, "My Policy Expired 30 Days Ago or Less", nil, "ExpiredWithin30Days", 4}
	currentInsuranceStatusExpiredOver30Days   = EnumCurrentInsuranceStatusItem{currentInsuranceStatusExpiredOver30DaysID, "My Policy Expired More Than 30 Days Ago", nil, "ExpiredOver30Days", 5}
	currentInsuranceStatusNoInsuranceRequired = EnumCurrentInsuranceStatusItem{currentInsuranceStatusNoInsuranceRequiredID, "No Insurance Required", nil, "NoInsuranceRequired", 6}
)

// EnumCurrentInsuranceStatus is a collection of CurrentInsuranceStatus items
type EnumCurrentInsuranceStatus struct {
	Description string
	Items       []*EnumCurrentInsuranceStatusItem
	Name        string

	OwnPolicy           *EnumCurrentInsuranceStatusItem
	AnothersPolicy      *EnumCurrentInsuranceStatusItem
	DeployedOverseas    *EnumCurrentInsuranceStatusItem
	ExpiredWithin30Days *EnumCurrentInsuranceStatusItem
	ExpiredOver30Days   *EnumCurrentInsuranceStatusItem
	NoInsuranceRequired *EnumCurrentInsuranceStatusItem

	itemDict map[string]*EnumCurrentInsuranceStatusItem
}

// CurrentInsuranceStatus is a public singleton instance of EnumCurrentInsuranceStatus
// representing current insurance statuses
var CurrentInsuranceStatus = &EnumCurrentInsuranceStatus{
	Description: "current insurance statuses",
	Items: []*EnumCurrentInsuranceStatusItem{
		&currentInsuranceStatusOwnPolicy,
		&currentInsuranceStatusAnothersPolicy,
		&currentInsuranceStatusDeployedOverseas,
		&currentInsuranceStatusExpiredWithin30Days,
		&currentInsuranceStatusExpiredOver30Days,
		&currentInsuranceStatusNoInsuranceRequired,
	},
	Name:                "EnumCurrentInsuranceStatus",
	OwnPolicy:           &currentInsuranceStatusOwnPolicy,
	AnothersPolicy:      &currentInsuranceStatusAnothersPolicy,
	DeployedOverseas:    &currentInsuranceStatusDeployedOverseas,
	ExpiredWithin30Days: &currentInsuranceStatusExpiredWithin30Days,
	ExpiredOver30Days:   &currentInsuranceStatusExpiredOver30Days,
	NoInsuranceRequired: &currentInsuranceStatusNoInsuranceRequired,

	itemDict: map[string]*EnumCurrentInsuranceStatusItem{
		strings.ToLower(string(currentInsuranceStatusOwnPolicyID)):           &currentInsuranceStatusOwnPolicy,
		strings.ToLower(string(currentInsuranceStatusAnothersPolicyID)):      &currentInsuranceStatusAnothersPolicy,
		strings.ToLower(string(currentInsuranceStatusDeployedOverseasID)):    &currentInsuranceStatusDeployedOverseas,
		strings.ToLower(string(currentInsuranceStatusExpiredWithin30DaysID)): &currentInsuranceStatusExpiredWithin30Days,
		strings.ToLower(string(currentInsuranceStatusExpiredOver30DaysID)):   &currentInsuranceStatusExpiredOver30Days,
		strings.ToLower(string(currentInsuranceStatusNoInsuranceRequiredID)): &currentInsuranceStatusNoInsuranceRequired,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumCurrentInsuranceStatus) ByID(id CurrentInsuranceStatusIdentifier) *EnumCurrentInsuranceStatusItem {
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
func (e *EnumCurrentInsuranceStatus) ByIDString(idx string) *EnumCurrentInsuranceStatusItem {
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
func (e *EnumCurrentInsuranceStatus) ByIndex(idx int) *EnumCurrentInsuranceStatusItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedCurrentInsuranceStatusID is a struct that is designed to replace a *CurrentInsuranceStatusID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *CurrentInsuranceStatusID it contains while being a better JSON citizen.
type ValidatedCurrentInsuranceStatusID struct {
	// id will point to a valid CurrentInsuranceStatusID, if possible
	// If id is nil, then ValidatedCurrentInsuranceStatusID.Valid() will return false.
	id *CurrentInsuranceStatusID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedCurrentInsuranceStatusID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedCurrentInsuranceStatusID
func (vi *ValidatedCurrentInsuranceStatusID) Clone() *ValidatedCurrentInsuranceStatusID {
	if vi == nil {
		return nil
	}

	var cid *CurrentInsuranceStatusID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedCurrentInsuranceStatusID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedCurrentInsuranceStatusIds represent the same CurrentInsuranceStatus
func (vi *ValidatedCurrentInsuranceStatusID) Equals(vj *ValidatedCurrentInsuranceStatusID) bool {
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

// Valid returns true if and only if the ValidatedCurrentInsuranceStatusID corresponds to a recognized CurrentInsuranceStatus
func (vi *ValidatedCurrentInsuranceStatusID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedCurrentInsuranceStatusID) ID() *CurrentInsuranceStatusID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedCurrentInsuranceStatusID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedCurrentInsuranceStatusID) ValidatedID() *ValidatedCurrentInsuranceStatusID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedCurrentInsuranceStatusID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedCurrentInsuranceStatusID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedCurrentInsuranceStatusID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedCurrentInsuranceStatusID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedCurrentInsuranceStatusID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := CurrentInsuranceStatusID(capString)
	item := CurrentInsuranceStatus.ByID(&id)
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

func (vi ValidatedCurrentInsuranceStatusID) String() string {
	return vi.ToIDString()
}

type CurrentInsuranceStatusIdentifier interface {
	ID() *CurrentInsuranceStatusID
	Valid() bool
}
