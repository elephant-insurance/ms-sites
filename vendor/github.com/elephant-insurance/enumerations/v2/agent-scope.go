package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// AgentScopeID uniquely identifies a particular AgentScope
type AgentScopeID string

// Clone creates a safe, independent copy of a AgentScopeID
func (i *AgentScopeID) Clone() *AgentScopeID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two AgentScopeIds are equivalent
func (i *AgentScopeID) Equals(j *AgentScopeID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *AgentScopeID that is either valid or nil
func (i *AgentScopeID) ID() *AgentScopeID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *AgentScopeID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the AgentScopeID corresponds to a recognized AgentScope
func (i *AgentScopeID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return AgentScope.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *AgentScopeID) ValidatedID() *ValidatedAgentScopeID {
	if i != nil {
		return &ValidatedAgentScopeID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *AgentScopeID) MarshalJSON() ([]byte, error) {
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

func (i *AgentScopeID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := AgentScopeID(dataString)
	item := AgentScope.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	agentScopeAllID  AgentScopeID = "All"
	agentScopeFR44ID AgentScopeID = "FR44"
)

// EnumAgentScopeItem describes an entry in an enumeration of AgentScope
type EnumAgentScopeItem struct {
	ID        AgentScopeID      `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	agentScopeAll  = EnumAgentScopeItem{agentScopeAllID, "All", nil, "All", 1}
	agentScopeFR44 = EnumAgentScopeItem{agentScopeFR44ID, "FR44", nil, "FR44", 2}
)

// EnumAgentScope is a collection of AgentScope items
type EnumAgentScope struct {
	Description string
	Items       []*EnumAgentScopeItem
	Name        string

	All  *EnumAgentScopeItem
	FR44 *EnumAgentScopeItem

	itemDict map[string]*EnumAgentScopeItem
}

// AgentScope is a public singleton instance of EnumAgentScope
// representing Scopes For Agency
var AgentScope = &EnumAgentScope{
	Description: "Scopes For Agency",
	Items: []*EnumAgentScopeItem{
		&agentScopeAll,
		&agentScopeFR44,
	},
	Name: "EnumAgentScope",
	All:  &agentScopeAll,
	FR44: &agentScopeFR44,

	itemDict: map[string]*EnumAgentScopeItem{
		strings.ToLower(string(agentScopeAllID)):  &agentScopeAll,
		strings.ToLower(string(agentScopeFR44ID)): &agentScopeFR44,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumAgentScope) ByID(id AgentScopeIdentifier) *EnumAgentScopeItem {
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
func (e *EnumAgentScope) ByIDString(idx string) *EnumAgentScopeItem {
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
func (e *EnumAgentScope) ByIndex(idx int) *EnumAgentScopeItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedAgentScopeID is a struct that is designed to replace a *AgentScopeID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *AgentScopeID it contains while being a better JSON citizen.
type ValidatedAgentScopeID struct {
	// id will point to a valid AgentScopeID, if possible
	// If id is nil, then ValidatedAgentScopeID.Valid() will return false.
	id *AgentScopeID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedAgentScopeID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedAgentScopeID
func (vi *ValidatedAgentScopeID) Clone() *ValidatedAgentScopeID {
	if vi == nil {
		return nil
	}

	var cid *AgentScopeID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedAgentScopeID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedAgentScopeIds represent the same AgentScope
func (vi *ValidatedAgentScopeID) Equals(vj *ValidatedAgentScopeID) bool {
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

// Valid returns true if and only if the ValidatedAgentScopeID corresponds to a recognized AgentScope
func (vi *ValidatedAgentScopeID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedAgentScopeID) ID() *AgentScopeID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedAgentScopeID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedAgentScopeID) ValidatedID() *ValidatedAgentScopeID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedAgentScopeID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedAgentScopeID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedAgentScopeID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedAgentScopeID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedAgentScopeID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := AgentScopeID(capString)
	item := AgentScope.ByID(&id)
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

func (vi ValidatedAgentScopeID) String() string {
	return vi.ToIDString()
}

type AgentScopeIdentifier interface {
	ID() *AgentScopeID
	Valid() bool
}
