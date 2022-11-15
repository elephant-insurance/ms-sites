package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// ProcessingEnvironmentNameID uniquely identifies a particular ProcessingEnvironmentName
type ProcessingEnvironmentNameID string

// Clone creates a safe, independent copy of a ProcessingEnvironmentNameID
func (i *ProcessingEnvironmentNameID) Clone() *ProcessingEnvironmentNameID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two ProcessingEnvironmentNameIds are equivalent
func (i *ProcessingEnvironmentNameID) Equals(j *ProcessingEnvironmentNameID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *ProcessingEnvironmentNameID that is either valid or nil
func (i *ProcessingEnvironmentNameID) ID() *ProcessingEnvironmentNameID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *ProcessingEnvironmentNameID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the ProcessingEnvironmentNameID corresponds to a recognized ProcessingEnvironmentName
func (i *ProcessingEnvironmentNameID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return ProcessingEnvironmentName.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *ProcessingEnvironmentNameID) ValidatedID() *ValidatedProcessingEnvironmentNameID {
	if i != nil {
		return &ValidatedProcessingEnvironmentNameID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *ProcessingEnvironmentNameID) MarshalJSON() ([]byte, error) {
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

func (i *ProcessingEnvironmentNameID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := ProcessingEnvironmentNameID(dataString)
	item := ProcessingEnvironmentName.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	processingEnvironmentNameProductionID   ProcessingEnvironmentNameID = "production"
	processingEnvironmentNameStandardTestID ProcessingEnvironmentNameID = "standardTest"
	processingEnvironmentNameEnhancedTestID ProcessingEnvironmentNameID = "enhancedTest"
)

// EnumProcessingEnvironmentNameItem describes an entry in an enumeration of ProcessingEnvironmentName
type EnumProcessingEnvironmentNameItem struct {
	ID        ProcessingEnvironmentNameID `json:"Value"`
	Desc      string                      `json:"Description,omitempty"`
	Meta      map[string]string           `json:",omitempty"`
	Name      string                      `json:"Name"`
	SortOrder int
}

var (
	processingEnvironmentNameProduction   = EnumProcessingEnvironmentNameItem{processingEnvironmentNameProductionID, "Production", nil, "Production", 1}
	processingEnvironmentNameStandardTest = EnumProcessingEnvironmentNameItem{processingEnvironmentNameStandardTestID, "StandardTest", nil, "StandardTest", 2}
	processingEnvironmentNameEnhancedTest = EnumProcessingEnvironmentNameItem{processingEnvironmentNameEnhancedTestID, "EnhancedTest", nil, "EnhancedTest", 3}
)

// EnumProcessingEnvironmentName is a collection of ProcessingEnvironmentName items
type EnumProcessingEnvironmentName struct {
	Description string
	Items       []*EnumProcessingEnvironmentNameItem
	Name        string

	Production   *EnumProcessingEnvironmentNameItem
	StandardTest *EnumProcessingEnvironmentNameItem
	EnhancedTest *EnumProcessingEnvironmentNameItem

	itemDict map[string]*EnumProcessingEnvironmentNameItem
}

// ProcessingEnvironmentName is a public singleton instance of EnumProcessingEnvironmentName
// representing name of processing environment
var ProcessingEnvironmentName = &EnumProcessingEnvironmentName{
	Description: "name of processing environment",
	Items: []*EnumProcessingEnvironmentNameItem{
		&processingEnvironmentNameProduction,
		&processingEnvironmentNameStandardTest,
		&processingEnvironmentNameEnhancedTest,
	},
	Name:         "EnumProcessingEnvironmentName",
	Production:   &processingEnvironmentNameProduction,
	StandardTest: &processingEnvironmentNameStandardTest,
	EnhancedTest: &processingEnvironmentNameEnhancedTest,

	itemDict: map[string]*EnumProcessingEnvironmentNameItem{
		strings.ToLower(string(processingEnvironmentNameProductionID)):   &processingEnvironmentNameProduction,
		strings.ToLower(string(processingEnvironmentNameStandardTestID)): &processingEnvironmentNameStandardTest,
		strings.ToLower(string(processingEnvironmentNameEnhancedTestID)): &processingEnvironmentNameEnhancedTest,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumProcessingEnvironmentName) ByID(id ProcessingEnvironmentNameIdentifier) *EnumProcessingEnvironmentNameItem {
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
func (e *EnumProcessingEnvironmentName) ByIDString(idx string) *EnumProcessingEnvironmentNameItem {
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
func (e *EnumProcessingEnvironmentName) ByIndex(idx int) *EnumProcessingEnvironmentNameItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedProcessingEnvironmentNameID is a struct that is designed to replace a *ProcessingEnvironmentNameID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *ProcessingEnvironmentNameID it contains while being a better JSON citizen.
type ValidatedProcessingEnvironmentNameID struct {
	// id will point to a valid ProcessingEnvironmentNameID, if possible
	// If id is nil, then ValidatedProcessingEnvironmentNameID.Valid() will return false.
	id *ProcessingEnvironmentNameID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedProcessingEnvironmentNameID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedProcessingEnvironmentNameID
func (vi *ValidatedProcessingEnvironmentNameID) Clone() *ValidatedProcessingEnvironmentNameID {
	if vi == nil {
		return nil
	}

	var cid *ProcessingEnvironmentNameID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedProcessingEnvironmentNameID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedProcessingEnvironmentNameIds represent the same ProcessingEnvironmentName
func (vi *ValidatedProcessingEnvironmentNameID) Equals(vj *ValidatedProcessingEnvironmentNameID) bool {
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

// Valid returns true if and only if the ValidatedProcessingEnvironmentNameID corresponds to a recognized ProcessingEnvironmentName
func (vi *ValidatedProcessingEnvironmentNameID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedProcessingEnvironmentNameID) ID() *ProcessingEnvironmentNameID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedProcessingEnvironmentNameID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedProcessingEnvironmentNameID) ValidatedID() *ValidatedProcessingEnvironmentNameID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedProcessingEnvironmentNameID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedProcessingEnvironmentNameID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedProcessingEnvironmentNameID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedProcessingEnvironmentNameID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedProcessingEnvironmentNameID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := ProcessingEnvironmentNameID(capString)
	item := ProcessingEnvironmentName.ByID(&id)
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

func (vi ValidatedProcessingEnvironmentNameID) String() string {
	return vi.ToIDString()
}

type ProcessingEnvironmentNameIdentifier interface {
	ID() *ProcessingEnvironmentNameID
	Valid() bool
}
