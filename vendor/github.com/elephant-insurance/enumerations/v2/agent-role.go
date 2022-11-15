package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// AgentRoleID uniquely identifies a particular AgentRole
type AgentRoleID string

// Clone creates a safe, independent copy of a AgentRoleID
func (i *AgentRoleID) Clone() *AgentRoleID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two AgentRoleIds are equivalent
func (i *AgentRoleID) Equals(j *AgentRoleID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *AgentRoleID that is either valid or nil
func (i *AgentRoleID) ID() *AgentRoleID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *AgentRoleID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the AgentRoleID corresponds to a recognized AgentRole
func (i *AgentRoleID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return AgentRole.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *AgentRoleID) ValidatedID() *ValidatedAgentRoleID {
	if i != nil {
		return &ValidatedAgentRoleID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *AgentRoleID) MarshalJSON() ([]byte, error) {
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

func (i *AgentRoleID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := AgentRoleID(dataString)
	item := AgentRole.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	agentRoleElephantAgentID      AgentRoleID = "Customer-Portal-Agent"
	agentRoleCustomerPortalUserID AgentRoleID = "Customer-Portal-User"
	agentRoleAgencyAgentID        AgentRoleID = "Customer-Portal-AgencyAgent"
	agentRoleAgencySupportID      AgentRoleID = "Customer-Portal-AgencySupport"
	agentRoleAgencyAdminID        AgentRoleID = "Customer-Portal-AgencyAdmin"
	agentRoleSalesAgentID         AgentRoleID = "SalesAgent"
	agentRoleAllID                AgentRoleID = "All"
)

// EnumAgentRoleItem describes an entry in an enumeration of AgentRole
type EnumAgentRoleItem struct {
	ID        AgentRoleID       `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	agentRoleElephantAgent      = EnumAgentRoleItem{agentRoleElephantAgentID, "ElephantAgent", nil, "ElephantAgent", 1}
	agentRoleCustomerPortalUser = EnumAgentRoleItem{agentRoleCustomerPortalUserID, "CustomerPortalUser", nil, "CustomerPortalUser", 2}
	agentRoleAgencyAgent        = EnumAgentRoleItem{agentRoleAgencyAgentID, "AgencyAgent", nil, "AgencyAgent", 3}
	agentRoleAgencySupport      = EnumAgentRoleItem{agentRoleAgencySupportID, "AgencySupport", nil, "AgencySupport", 4}
	agentRoleAgencyAdmin        = EnumAgentRoleItem{agentRoleAgencyAdminID, "AgencyAdmin", nil, "AgencyAdmin", 5}
	agentRoleSalesAgent         = EnumAgentRoleItem{agentRoleSalesAgentID, "SalesAgent", nil, "SalesAgent", 6}
	agentRoleAll                = EnumAgentRoleItem{agentRoleAllID, "All", nil, "All", 7}
)

// EnumAgentRole is a collection of AgentRole items
type EnumAgentRole struct {
	Description string
	Items       []*EnumAgentRoleItem
	Name        string

	ElephantAgent      *EnumAgentRoleItem
	CustomerPortalUser *EnumAgentRoleItem
	AgencyAgent        *EnumAgentRoleItem
	AgencySupport      *EnumAgentRoleItem
	AgencyAdmin        *EnumAgentRoleItem
	SalesAgent         *EnumAgentRoleItem
	All                *EnumAgentRoleItem

	itemDict map[string]*EnumAgentRoleItem
}

// AgentRole is a public singleton instance of EnumAgentRole
// representing Roles for CSP And Agency
var AgentRole = &EnumAgentRole{
	Description: "Roles for CSP And Agency",
	Items: []*EnumAgentRoleItem{
		&agentRoleElephantAgent,
		&agentRoleCustomerPortalUser,
		&agentRoleAgencyAgent,
		&agentRoleAgencySupport,
		&agentRoleAgencyAdmin,
		&agentRoleSalesAgent,
		&agentRoleAll,
	},
	Name:               "EnumAgentRole",
	ElephantAgent:      &agentRoleElephantAgent,
	CustomerPortalUser: &agentRoleCustomerPortalUser,
	AgencyAgent:        &agentRoleAgencyAgent,
	AgencySupport:      &agentRoleAgencySupport,
	AgencyAdmin:        &agentRoleAgencyAdmin,
	SalesAgent:         &agentRoleSalesAgent,
	All:                &agentRoleAll,

	itemDict: map[string]*EnumAgentRoleItem{
		strings.ToLower(string(agentRoleElephantAgentID)):      &agentRoleElephantAgent,
		strings.ToLower(string(agentRoleCustomerPortalUserID)): &agentRoleCustomerPortalUser,
		strings.ToLower(string(agentRoleAgencyAgentID)):        &agentRoleAgencyAgent,
		strings.ToLower(string(agentRoleAgencySupportID)):      &agentRoleAgencySupport,
		strings.ToLower(string(agentRoleAgencyAdminID)):        &agentRoleAgencyAdmin,
		strings.ToLower(string(agentRoleSalesAgentID)):         &agentRoleSalesAgent,
		strings.ToLower(string(agentRoleAllID)):                &agentRoleAll,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumAgentRole) ByID(id AgentRoleIdentifier) *EnumAgentRoleItem {
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
func (e *EnumAgentRole) ByIDString(idx string) *EnumAgentRoleItem {
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
func (e *EnumAgentRole) ByIndex(idx int) *EnumAgentRoleItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedAgentRoleID is a struct that is designed to replace a *AgentRoleID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *AgentRoleID it contains while being a better JSON citizen.
type ValidatedAgentRoleID struct {
	// id will point to a valid AgentRoleID, if possible
	// If id is nil, then ValidatedAgentRoleID.Valid() will return false.
	id *AgentRoleID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedAgentRoleID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedAgentRoleID
func (vi *ValidatedAgentRoleID) Clone() *ValidatedAgentRoleID {
	if vi == nil {
		return nil
	}

	var cid *AgentRoleID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedAgentRoleID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedAgentRoleIds represent the same AgentRole
func (vi *ValidatedAgentRoleID) Equals(vj *ValidatedAgentRoleID) bool {
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

// Valid returns true if and only if the ValidatedAgentRoleID corresponds to a recognized AgentRole
func (vi *ValidatedAgentRoleID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedAgentRoleID) ID() *AgentRoleID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedAgentRoleID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedAgentRoleID) ValidatedID() *ValidatedAgentRoleID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedAgentRoleID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedAgentRoleID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedAgentRoleID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedAgentRoleID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedAgentRoleID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := AgentRoleID(capString)
	item := AgentRole.ByID(&id)
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

func (vi ValidatedAgentRoleID) String() string {
	return vi.ToIDString()
}

type AgentRoleIdentifier interface {
	ID() *AgentRoleID
	Valid() bool
}
