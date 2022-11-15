package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// PolicyStatusID uniquely identifies a particular PolicyStatus
type PolicyStatusID string

// Clone creates a safe, independent copy of a PolicyStatusID
func (i *PolicyStatusID) Clone() *PolicyStatusID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two PolicyStatusIds are equivalent
func (i *PolicyStatusID) Equals(j *PolicyStatusID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *PolicyStatusID that is either valid or nil
func (i *PolicyStatusID) ID() *PolicyStatusID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *PolicyStatusID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the PolicyStatusID corresponds to a recognized PolicyStatus
func (i *PolicyStatusID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return PolicyStatus.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *PolicyStatusID) ValidatedID() *ValidatedPolicyStatusID {
	if i != nil {
		return &ValidatedPolicyStatusID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *PolicyStatusID) MarshalJSON() ([]byte, error) {
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

func (i *PolicyStatusID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := PolicyStatusID(dataString)
	item := PolicyStatus.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	policyStatusDraftID  PolicyStatusID = "Draft"
	policyStatusQuotedID PolicyStatusID = "Quoted"
)

// EnumPolicyStatusItem describes an entry in an enumeration of PolicyStatus
type EnumPolicyStatusItem struct {
	ID        PolicyStatusID    `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	policyStatusDraft  = EnumPolicyStatusItem{policyStatusDraftID, "Draft", nil, "Draft", 1}
	policyStatusQuoted = EnumPolicyStatusItem{policyStatusQuotedID, "Quoted", nil, "Quoted", 2}
)

// EnumPolicyStatus is a collection of PolicyStatus items
type EnumPolicyStatus struct {
	Description string
	Items       []*EnumPolicyStatusItem
	Name        string

	Draft  *EnumPolicyStatusItem
	Quoted *EnumPolicyStatusItem

	itemDict map[string]*EnumPolicyStatusItem
}

// PolicyStatus is a public singleton instance of EnumPolicyStatus
// representing statuses of policies, draft or quoted
var PolicyStatus = &EnumPolicyStatus{
	Description: "statuses of policies, draft or quoted",
	Items: []*EnumPolicyStatusItem{
		&policyStatusDraft,
		&policyStatusQuoted,
	},
	Name:   "EnumPolicyStatus",
	Draft:  &policyStatusDraft,
	Quoted: &policyStatusQuoted,

	itemDict: map[string]*EnumPolicyStatusItem{
		strings.ToLower(string(policyStatusDraftID)):  &policyStatusDraft,
		strings.ToLower(string(policyStatusQuotedID)): &policyStatusQuoted,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumPolicyStatus) ByID(id PolicyStatusIdentifier) *EnumPolicyStatusItem {
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
func (e *EnumPolicyStatus) ByIDString(idx string) *EnumPolicyStatusItem {
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
func (e *EnumPolicyStatus) ByIndex(idx int) *EnumPolicyStatusItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedPolicyStatusID is a struct that is designed to replace a *PolicyStatusID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *PolicyStatusID it contains while being a better JSON citizen.
type ValidatedPolicyStatusID struct {
	// id will point to a valid PolicyStatusID, if possible
	// If id is nil, then ValidatedPolicyStatusID.Valid() will return false.
	id *PolicyStatusID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedPolicyStatusID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedPolicyStatusID
func (vi *ValidatedPolicyStatusID) Clone() *ValidatedPolicyStatusID {
	if vi == nil {
		return nil
	}

	var cid *PolicyStatusID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedPolicyStatusID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedPolicyStatusIds represent the same PolicyStatus
func (vi *ValidatedPolicyStatusID) Equals(vj *ValidatedPolicyStatusID) bool {
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

// Valid returns true if and only if the ValidatedPolicyStatusID corresponds to a recognized PolicyStatus
func (vi *ValidatedPolicyStatusID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedPolicyStatusID) ID() *PolicyStatusID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedPolicyStatusID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedPolicyStatusID) ValidatedID() *ValidatedPolicyStatusID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedPolicyStatusID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedPolicyStatusID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedPolicyStatusID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedPolicyStatusID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedPolicyStatusID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := PolicyStatusID(capString)
	item := PolicyStatus.ByID(&id)
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

func (vi ValidatedPolicyStatusID) String() string {
	return vi.ToIDString()
}

type PolicyStatusIdentifier interface {
	ID() *PolicyStatusID
	Valid() bool
}
