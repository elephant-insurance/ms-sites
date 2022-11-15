package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// IntegrationPartnerID uniquely identifies a particular IntegrationPartner
type IntegrationPartnerID string

// Clone creates a safe, independent copy of a IntegrationPartnerID
func (i *IntegrationPartnerID) Clone() *IntegrationPartnerID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two IntegrationPartnerIds are equivalent
func (i *IntegrationPartnerID) Equals(j *IntegrationPartnerID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *IntegrationPartnerID that is either valid or nil
func (i *IntegrationPartnerID) ID() *IntegrationPartnerID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *IntegrationPartnerID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the IntegrationPartnerID corresponds to a recognized IntegrationPartner
func (i *IntegrationPartnerID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return IntegrationPartner.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *IntegrationPartnerID) ValidatedID() *ValidatedIntegrationPartnerID {
	if i != nil {
		return &ValidatedIntegrationPartnerID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *IntegrationPartnerID) MarshalJSON() ([]byte, error) {
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

func (i *IntegrationPartnerID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := IntegrationPartnerID(dataString)
	item := IntegrationPartner.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	integrationPartnerAllWebID     IntegrationPartnerID = "AllWeb"
	integrationPartnerBoltID       IntegrationPartnerID = "Bolt"
	integrationPartnerEzLynxID     IntegrationPartnerID = "EzLynx"
	integrationPartnerLeadCloudID  IntegrationPartnerID = "LeadCloud"
	integrationPartnerMediaAlphaID IntegrationPartnerID = "MediaAlpha"
	integrationPartnerCompareNowID IntegrationPartnerID = "CompareNow"
	integrationPartnerDirectID     IntegrationPartnerID = "Direct"
	integrationPartnerRedVentureID IntegrationPartnerID = "RedVenture"
	integrationPartnerITCID        IntegrationPartnerID = "ITC"
)

// EnumIntegrationPartnerItem describes an entry in an enumeration of IntegrationPartner
type EnumIntegrationPartnerItem struct {
	ID        IntegrationPartnerID `json:"Value"`
	Desc      string               `json:"Description,omitempty"`
	Meta      map[string]string    `json:",omitempty"`
	Name      string               `json:"Name"`
	SortOrder int
}

var (
	integrationPartnerAllWeb     = EnumIntegrationPartnerItem{integrationPartnerAllWebID, "All Web", nil, "AllWeb", 1}
	integrationPartnerBolt       = EnumIntegrationPartnerItem{integrationPartnerBoltID, "Bolt", nil, "Bolt", 2}
	integrationPartnerEzLynx     = EnumIntegrationPartnerItem{integrationPartnerEzLynxID, "EzLynx", nil, "EzLynx", 3}
	integrationPartnerLeadCloud  = EnumIntegrationPartnerItem{integrationPartnerLeadCloudID, "Lead Cloud", nil, "LeadCloud", 4}
	integrationPartnerMediaAlpha = EnumIntegrationPartnerItem{integrationPartnerMediaAlphaID, "Media Alpha", nil, "MediaAlpha", 5}
	integrationPartnerCompareNow = EnumIntegrationPartnerItem{integrationPartnerCompareNowID, "Compare", nil, "CompareNow", 6}
	integrationPartnerDirect     = EnumIntegrationPartnerItem{integrationPartnerDirectID, "Direct", nil, "Direct", 7}
	integrationPartnerRedVenture = EnumIntegrationPartnerItem{integrationPartnerRedVentureID, "RedVenture", nil, "RedVenture", 8}
	integrationPartnerITC        = EnumIntegrationPartnerItem{integrationPartnerITCID, "ITC", nil, "ITC", 9}
)

// EnumIntegrationPartner is a collection of IntegrationPartner items
type EnumIntegrationPartner struct {
	Description string
	Items       []*EnumIntegrationPartnerItem
	Name        string

	AllWeb     *EnumIntegrationPartnerItem
	Bolt       *EnumIntegrationPartnerItem
	EzLynx     *EnumIntegrationPartnerItem
	LeadCloud  *EnumIntegrationPartnerItem
	MediaAlpha *EnumIntegrationPartnerItem
	CompareNow *EnumIntegrationPartnerItem
	Direct     *EnumIntegrationPartnerItem
	RedVenture *EnumIntegrationPartnerItem
	ITC        *EnumIntegrationPartnerItem

	itemDict map[string]*EnumIntegrationPartnerItem
}

// IntegrationPartner is a public singleton instance of EnumIntegrationPartner
// representing integration partners
var IntegrationPartner = &EnumIntegrationPartner{
	Description: "integration partners",
	Items: []*EnumIntegrationPartnerItem{
		&integrationPartnerAllWeb,
		&integrationPartnerBolt,
		&integrationPartnerEzLynx,
		&integrationPartnerLeadCloud,
		&integrationPartnerMediaAlpha,
		&integrationPartnerCompareNow,
		&integrationPartnerDirect,
		&integrationPartnerRedVenture,
		&integrationPartnerITC,
	},
	Name:       "EnumIntegrationPartner",
	AllWeb:     &integrationPartnerAllWeb,
	Bolt:       &integrationPartnerBolt,
	EzLynx:     &integrationPartnerEzLynx,
	LeadCloud:  &integrationPartnerLeadCloud,
	MediaAlpha: &integrationPartnerMediaAlpha,
	CompareNow: &integrationPartnerCompareNow,
	Direct:     &integrationPartnerDirect,
	RedVenture: &integrationPartnerRedVenture,
	ITC:        &integrationPartnerITC,

	itemDict: map[string]*EnumIntegrationPartnerItem{
		strings.ToLower(string(integrationPartnerAllWebID)):     &integrationPartnerAllWeb,
		strings.ToLower(string(integrationPartnerBoltID)):       &integrationPartnerBolt,
		strings.ToLower(string(integrationPartnerEzLynxID)):     &integrationPartnerEzLynx,
		strings.ToLower(string(integrationPartnerLeadCloudID)):  &integrationPartnerLeadCloud,
		strings.ToLower(string(integrationPartnerMediaAlphaID)): &integrationPartnerMediaAlpha,
		strings.ToLower(string(integrationPartnerCompareNowID)): &integrationPartnerCompareNow,
		strings.ToLower(string(integrationPartnerDirectID)):     &integrationPartnerDirect,
		strings.ToLower(string(integrationPartnerRedVentureID)): &integrationPartnerRedVenture,
		strings.ToLower(string(integrationPartnerITCID)):        &integrationPartnerITC,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumIntegrationPartner) ByID(id IntegrationPartnerIdentifier) *EnumIntegrationPartnerItem {
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
func (e *EnumIntegrationPartner) ByIDString(idx string) *EnumIntegrationPartnerItem {
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
func (e *EnumIntegrationPartner) ByIndex(idx int) *EnumIntegrationPartnerItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedIntegrationPartnerID is a struct that is designed to replace a *IntegrationPartnerID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *IntegrationPartnerID it contains while being a better JSON citizen.
type ValidatedIntegrationPartnerID struct {
	// id will point to a valid IntegrationPartnerID, if possible
	// If id is nil, then ValidatedIntegrationPartnerID.Valid() will return false.
	id *IntegrationPartnerID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedIntegrationPartnerID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedIntegrationPartnerID
func (vi *ValidatedIntegrationPartnerID) Clone() *ValidatedIntegrationPartnerID {
	if vi == nil {
		return nil
	}

	var cid *IntegrationPartnerID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedIntegrationPartnerID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedIntegrationPartnerIds represent the same IntegrationPartner
func (vi *ValidatedIntegrationPartnerID) Equals(vj *ValidatedIntegrationPartnerID) bool {
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

// Valid returns true if and only if the ValidatedIntegrationPartnerID corresponds to a recognized IntegrationPartner
func (vi *ValidatedIntegrationPartnerID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedIntegrationPartnerID) ID() *IntegrationPartnerID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedIntegrationPartnerID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedIntegrationPartnerID) ValidatedID() *ValidatedIntegrationPartnerID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedIntegrationPartnerID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedIntegrationPartnerID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedIntegrationPartnerID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedIntegrationPartnerID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedIntegrationPartnerID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := IntegrationPartnerID(capString)
	item := IntegrationPartner.ByID(&id)
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

func (vi ValidatedIntegrationPartnerID) String() string {
	return vi.ToIDString()
}

type IntegrationPartnerIdentifier interface {
	ID() *IntegrationPartnerID
	Valid() bool
}
