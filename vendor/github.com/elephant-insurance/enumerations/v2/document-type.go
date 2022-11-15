package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// DocTypeID uniquely identifies a particular DocType
type DocTypeID string

// Clone creates a safe, independent copy of a DocTypeID
func (i *DocTypeID) Clone() *DocTypeID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two DocTypeIds are equivalent
func (i *DocTypeID) Equals(j *DocTypeID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *DocTypeID that is either valid or nil
func (i *DocTypeID) ID() *DocTypeID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *DocTypeID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the DocTypeID corresponds to a recognized DocType
func (i *DocTypeID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return DocType.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *DocTypeID) ValidatedID() *ValidatedDocTypeID {
	if i != nil {
		return &ValidatedDocTypeID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *DocTypeID) MarshalJSON() ([]byte, error) {
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

func (i *DocTypeID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := DocTypeID(dataString)
	item := DocType.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	docTypeIDCardID            DocTypeID = "idcard"
	docTypePolicyDeclarationID DocTypeID = "decpage"
	docTypeRoadSideIDCardID    DocTypeID = "roadsideltr"
)

// EnumDocTypeItem describes an entry in an enumeration of DocType
type EnumDocTypeItem struct {
	ID        DocTypeID         `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	docTypeIDCard            = EnumDocTypeItem{docTypeIDCardID, "IDCard", nil, "IDCard", 1}
	docTypePolicyDeclaration = EnumDocTypeItem{docTypePolicyDeclarationID, "PolicyDeclaration", nil, "PolicyDeclaration", 2}
	docTypeRoadSideIDCard    = EnumDocTypeItem{docTypeRoadSideIDCardID, "RoadSideIDCard", nil, "RoadSideIDCard", 3}
)

// EnumDocType is a collection of DocType items
type EnumDocType struct {
	Description string
	Items       []*EnumDocTypeItem
	Name        string

	IDCard            *EnumDocTypeItem
	PolicyDeclaration *EnumDocTypeItem
	RoadSideIDCard    *EnumDocTypeItem

	itemDict map[string]*EnumDocTypeItem
}

// DocType is a public singleton instance of EnumDocType
// representing doctype
var DocType = &EnumDocType{
	Description: "doctype",
	Items: []*EnumDocTypeItem{
		&docTypeIDCard,
		&docTypePolicyDeclaration,
		&docTypeRoadSideIDCard,
	},
	Name:              "EnumDocType",
	IDCard:            &docTypeIDCard,
	PolicyDeclaration: &docTypePolicyDeclaration,
	RoadSideIDCard:    &docTypeRoadSideIDCard,

	itemDict: map[string]*EnumDocTypeItem{
		strings.ToLower(string(docTypeIDCardID)):            &docTypeIDCard,
		strings.ToLower(string(docTypePolicyDeclarationID)): &docTypePolicyDeclaration,
		strings.ToLower(string(docTypeRoadSideIDCardID)):    &docTypeRoadSideIDCard,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumDocType) ByID(id DocTypeIdentifier) *EnumDocTypeItem {
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
func (e *EnumDocType) ByIDString(idx string) *EnumDocTypeItem {
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
func (e *EnumDocType) ByIndex(idx int) *EnumDocTypeItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedDocTypeID is a struct that is designed to replace a *DocTypeID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *DocTypeID it contains while being a better JSON citizen.
type ValidatedDocTypeID struct {
	// id will point to a valid DocTypeID, if possible
	// If id is nil, then ValidatedDocTypeID.Valid() will return false.
	id *DocTypeID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedDocTypeID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedDocTypeID
func (vi *ValidatedDocTypeID) Clone() *ValidatedDocTypeID {
	if vi == nil {
		return nil
	}

	var cid *DocTypeID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedDocTypeID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedDocTypeIds represent the same DocType
func (vi *ValidatedDocTypeID) Equals(vj *ValidatedDocTypeID) bool {
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

// Valid returns true if and only if the ValidatedDocTypeID corresponds to a recognized DocType
func (vi *ValidatedDocTypeID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedDocTypeID) ID() *DocTypeID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedDocTypeID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedDocTypeID) ValidatedID() *ValidatedDocTypeID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedDocTypeID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedDocTypeID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedDocTypeID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedDocTypeID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedDocTypeID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := DocTypeID(capString)
	item := DocType.ByID(&id)
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

func (vi ValidatedDocTypeID) String() string {
	return vi.ToIDString()
}

type DocTypeIdentifier interface {
	ID() *DocTypeID
	Valid() bool
}
