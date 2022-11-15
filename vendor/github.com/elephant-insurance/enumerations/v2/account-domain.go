package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// AccountDomainID uniquely identifies a particular AccountDomain
type AccountDomainID string

// Clone creates a safe, independent copy of a AccountDomainID
func (i *AccountDomainID) Clone() *AccountDomainID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two AccountDomainIds are equivalent
func (i *AccountDomainID) Equals(j *AccountDomainID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *AccountDomainID that is either valid or nil
func (i *AccountDomainID) ID() *AccountDomainID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *AccountDomainID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the AccountDomainID corresponds to a recognized AccountDomain
func (i *AccountDomainID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return AccountDomain.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *AccountDomainID) ValidatedID() *ValidatedAccountDomainID {
	if i != nil {
		return &ValidatedAccountDomainID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *AccountDomainID) MarshalJSON() ([]byte, error) {
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

func (i *AccountDomainID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := AccountDomainID(dataString)
	item := AccountDomain.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	accountDomainDirectID        AccountDomainID = "direct"
	accountDomainLibertyMutualID AccountDomainID = "liberty_mutual"
	accountDomainAgencyID        AccountDomainID = "agency"
)

// EnumAccountDomainItem describes an entry in an enumeration of AccountDomain
type EnumAccountDomainItem struct {
	ID        AccountDomainID   `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	accountDomainDirect        = EnumAccountDomainItem{accountDomainDirectID, "Direct insurance", nil, "Direct", 1}
	accountDomainLibertyMutual = EnumAccountDomainItem{accountDomainLibertyMutualID, "Liberty mutual only agency - subjected to go away", nil, "LibertyMutual", 2}
	accountDomainAgency        = EnumAccountDomainItem{accountDomainAgencyID, "All Agencies", nil, "Agency", 3}
)

// EnumAccountDomain is a collection of AccountDomain items
type EnumAccountDomain struct {
	Description string
	Items       []*EnumAccountDomainItem
	Name        string

	Direct        *EnumAccountDomainItem
	LibertyMutual *EnumAccountDomainItem
	Agency        *EnumAccountDomainItem

	itemDict map[string]*EnumAccountDomainItem
}

// AccountDomain is a public singleton instance of EnumAccountDomain
// representing independent account domains at Elephant
var AccountDomain = &EnumAccountDomain{
	Description: "independent account domains at Elephant",
	Items: []*EnumAccountDomainItem{
		&accountDomainDirect,
		&accountDomainLibertyMutual,
		&accountDomainAgency,
	},
	Name:          "EnumAccountDomain",
	Direct:        &accountDomainDirect,
	LibertyMutual: &accountDomainLibertyMutual,
	Agency:        &accountDomainAgency,

	itemDict: map[string]*EnumAccountDomainItem{
		strings.ToLower(string(accountDomainDirectID)):        &accountDomainDirect,
		strings.ToLower(string(accountDomainLibertyMutualID)): &accountDomainLibertyMutual,
		strings.ToLower(string(accountDomainAgencyID)):        &accountDomainAgency,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumAccountDomain) ByID(id AccountDomainIdentifier) *EnumAccountDomainItem {
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
func (e *EnumAccountDomain) ByIDString(idx string) *EnumAccountDomainItem {
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
func (e *EnumAccountDomain) ByIndex(idx int) *EnumAccountDomainItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedAccountDomainID is a struct that is designed to replace a *AccountDomainID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *AccountDomainID it contains while being a better JSON citizen.
type ValidatedAccountDomainID struct {
	// id will point to a valid AccountDomainID, if possible
	// If id is nil, then ValidatedAccountDomainID.Valid() will return false.
	id *AccountDomainID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedAccountDomainID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedAccountDomainID
func (vi *ValidatedAccountDomainID) Clone() *ValidatedAccountDomainID {
	if vi == nil {
		return nil
	}

	var cid *AccountDomainID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedAccountDomainID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedAccountDomainIds represent the same AccountDomain
func (vi *ValidatedAccountDomainID) Equals(vj *ValidatedAccountDomainID) bool {
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

// Valid returns true if and only if the ValidatedAccountDomainID corresponds to a recognized AccountDomain
func (vi *ValidatedAccountDomainID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedAccountDomainID) ID() *AccountDomainID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedAccountDomainID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedAccountDomainID) ValidatedID() *ValidatedAccountDomainID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedAccountDomainID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedAccountDomainID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedAccountDomainID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedAccountDomainID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedAccountDomainID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := AccountDomainID(capString)
	item := AccountDomain.ByID(&id)
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

func (vi ValidatedAccountDomainID) String() string {
	return vi.ToIDString()
}

type AccountDomainIdentifier interface {
	ID() *AccountDomainID
	Valid() bool
}
