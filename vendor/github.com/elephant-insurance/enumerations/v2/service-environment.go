package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// ServiceEnvironmentID uniquely identifies a particular ServiceEnvironment
type ServiceEnvironmentID string

// Clone creates a safe, independent copy of a ServiceEnvironmentID
func (i *ServiceEnvironmentID) Clone() *ServiceEnvironmentID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two ServiceEnvironmentIds are equivalent
func (i *ServiceEnvironmentID) Equals(j *ServiceEnvironmentID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *ServiceEnvironmentID that is either valid or nil
func (i *ServiceEnvironmentID) ID() *ServiceEnvironmentID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *ServiceEnvironmentID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the ServiceEnvironmentID corresponds to a recognized ServiceEnvironment
func (i *ServiceEnvironmentID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return ServiceEnvironment.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *ServiceEnvironmentID) ValidatedID() *ValidatedServiceEnvironmentID {
	if i != nil {
		return &ValidatedServiceEnvironmentID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *ServiceEnvironmentID) MarshalJSON() ([]byte, error) {
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

func (i *ServiceEnvironmentID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := ServiceEnvironmentID(dataString)
	item := ServiceEnvironment.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	serviceEnvironmentDevelopmentID ServiceEnvironmentID = "dev"
	serviceEnvironmentTestingID     ServiceEnvironmentID = "test"
	serviceEnvironmentQA1ID         ServiceEnvironmentID = "qa1"
	serviceEnvironmentQA2ID         ServiceEnvironmentID = "qa2"
	serviceEnvironmentQA3ID         ServiceEnvironmentID = "qa3"
	serviceEnvironmentQA4ID         ServiceEnvironmentID = "qa4"
	serviceEnvironmentQA5ID         ServiceEnvironmentID = "qa5"
	serviceEnvironmentQA6ID         ServiceEnvironmentID = "qa6"
	serviceEnvironmentQACloud1ID    ServiceEnvironmentID = "qac1"
	serviceEnvironmentQATestID      ServiceEnvironmentID = "qat"
	serviceEnvironmentCI1ID         ServiceEnvironmentID = "ci1"
	serviceEnvironmentCI2ID         ServiceEnvironmentID = "ci2"
	serviceEnvironmentSandboxID     ServiceEnvironmentID = "sbx"
	serviceEnvironmentUATID         ServiceEnvironmentID = "uat"
	serviceEnvironmentPartnerTestID ServiceEnvironmentID = "pnt"
	serviceEnvironmentProdSupID     ServiceEnvironmentID = "prs"
	serviceEnvironmentProductionID  ServiceEnvironmentID = "prd"
)

// EnumServiceEnvironmentItem describes an entry in an enumeration of ServiceEnvironment
type EnumServiceEnvironmentItem struct {
	ID        ServiceEnvironmentID `json:"Value"`
	Desc      string               `json:"Description,omitempty"`
	Meta      map[string]string    `json:",omitempty"`
	Name      string               `json:"Name"`
	SortOrder int

	// Meta Properties
	Prod string
}

var (
	serviceEnvironmentDevelopment = EnumServiceEnvironmentItem{serviceEnvironmentDevelopmentID, "Development", map[string]string{"Prod": "false"}, "Development", 1, "false"}
	serviceEnvironmentTesting     = EnumServiceEnvironmentItem{serviceEnvironmentTestingID, "Testing", map[string]string{"Prod": "false"}, "Testing", 2, "false"}
	serviceEnvironmentQA1         = EnumServiceEnvironmentItem{serviceEnvironmentQA1ID, "QA1", map[string]string{"Prod": "false"}, "QA1", 3, "false"}
	serviceEnvironmentQA2         = EnumServiceEnvironmentItem{serviceEnvironmentQA2ID, "QA2", map[string]string{"Prod": "false"}, "QA2", 4, "false"}
	serviceEnvironmentQA3         = EnumServiceEnvironmentItem{serviceEnvironmentQA3ID, "QA3", map[string]string{"Prod": "false"}, "QA3", 5, "false"}
	serviceEnvironmentQA4         = EnumServiceEnvironmentItem{serviceEnvironmentQA4ID, "QA4", map[string]string{"Prod": "false"}, "QA4", 6, "false"}
	serviceEnvironmentQA5         = EnumServiceEnvironmentItem{serviceEnvironmentQA5ID, "QA5", map[string]string{"Prod": "false"}, "QA5", 7, "false"}
	serviceEnvironmentQA6         = EnumServiceEnvironmentItem{serviceEnvironmentQA6ID, "QA6", map[string]string{"Prod": "false"}, "QA6", 8, "false"}
	serviceEnvironmentQACloud1    = EnumServiceEnvironmentItem{serviceEnvironmentQACloud1ID, "QACloud1", map[string]string{"Prod": "false"}, "QACloud1", 9, "false"}
	serviceEnvironmentQATest      = EnumServiceEnvironmentItem{serviceEnvironmentQATestID, "QATest", map[string]string{"Prod": "false"}, "QATest", 10, "false"}
	serviceEnvironmentCI1         = EnumServiceEnvironmentItem{serviceEnvironmentCI1ID, "CI1", map[string]string{"Prod": "false"}, "CI1", 11, "false"}
	serviceEnvironmentCI2         = EnumServiceEnvironmentItem{serviceEnvironmentCI2ID, "CI2", map[string]string{"Prod": "false"}, "CI2", 12, "false"}
	serviceEnvironmentSandbox     = EnumServiceEnvironmentItem{serviceEnvironmentSandboxID, "Sandbox", map[string]string{"Prod": "false"}, "Sandbox", 13, "false"}
	serviceEnvironmentUAT         = EnumServiceEnvironmentItem{serviceEnvironmentUATID, "UAT", map[string]string{"Prod": "false"}, "UAT", 14, "false"}
	serviceEnvironmentPartnerTest = EnumServiceEnvironmentItem{serviceEnvironmentPartnerTestID, "PartnerTest", map[string]string{"Prod": "false"}, "PartnerTest", 15, "false"}
	serviceEnvironmentProdSup     = EnumServiceEnvironmentItem{serviceEnvironmentProdSupID, "ProdSup", map[string]string{"Prod": "true"}, "ProdSup", 16, "true"}
	serviceEnvironmentProduction  = EnumServiceEnvironmentItem{serviceEnvironmentProductionID, "Production", map[string]string{"Prod": "true"}, "Production", 17, "true"}
)

// EnumServiceEnvironment is a collection of ServiceEnvironment items
type EnumServiceEnvironment struct {
	Description string
	Items       []*EnumServiceEnvironmentItem
	Name        string

	Development *EnumServiceEnvironmentItem
	Testing     *EnumServiceEnvironmentItem
	QA1         *EnumServiceEnvironmentItem
	QA2         *EnumServiceEnvironmentItem
	QA3         *EnumServiceEnvironmentItem
	QA4         *EnumServiceEnvironmentItem
	QA5         *EnumServiceEnvironmentItem
	QA6         *EnumServiceEnvironmentItem
	QACloud1    *EnumServiceEnvironmentItem
	QATest      *EnumServiceEnvironmentItem
	CI1         *EnumServiceEnvironmentItem
	CI2         *EnumServiceEnvironmentItem
	Sandbox     *EnumServiceEnvironmentItem
	UAT         *EnumServiceEnvironmentItem
	PartnerTest *EnumServiceEnvironmentItem
	ProdSup     *EnumServiceEnvironmentItem
	Production  *EnumServiceEnvironmentItem

	itemDict map[string]*EnumServiceEnvironmentItem
}

// ServiceEnvironment is a public singleton instance of EnumServiceEnvironment
// representing environments where we run microservices
var ServiceEnvironment = &EnumServiceEnvironment{
	Description: "environments where we run microservices",
	Items: []*EnumServiceEnvironmentItem{
		&serviceEnvironmentDevelopment,
		&serviceEnvironmentTesting,
		&serviceEnvironmentQA1,
		&serviceEnvironmentQA2,
		&serviceEnvironmentQA3,
		&serviceEnvironmentQA4,
		&serviceEnvironmentQA5,
		&serviceEnvironmentQA6,
		&serviceEnvironmentQACloud1,
		&serviceEnvironmentQATest,
		&serviceEnvironmentCI1,
		&serviceEnvironmentCI2,
		&serviceEnvironmentSandbox,
		&serviceEnvironmentUAT,
		&serviceEnvironmentPartnerTest,
		&serviceEnvironmentProdSup,
		&serviceEnvironmentProduction,
	},
	Name:        "EnumServiceEnvironment",
	Development: &serviceEnvironmentDevelopment,
	Testing:     &serviceEnvironmentTesting,
	QA1:         &serviceEnvironmentQA1,
	QA2:         &serviceEnvironmentQA2,
	QA3:         &serviceEnvironmentQA3,
	QA4:         &serviceEnvironmentQA4,
	QA5:         &serviceEnvironmentQA5,
	QA6:         &serviceEnvironmentQA6,
	QACloud1:    &serviceEnvironmentQACloud1,
	QATest:      &serviceEnvironmentQATest,
	CI1:         &serviceEnvironmentCI1,
	CI2:         &serviceEnvironmentCI2,
	Sandbox:     &serviceEnvironmentSandbox,
	UAT:         &serviceEnvironmentUAT,
	PartnerTest: &serviceEnvironmentPartnerTest,
	ProdSup:     &serviceEnvironmentProdSup,
	Production:  &serviceEnvironmentProduction,

	itemDict: map[string]*EnumServiceEnvironmentItem{
		strings.ToLower(string(serviceEnvironmentDevelopmentID)): &serviceEnvironmentDevelopment,
		strings.ToLower(string(serviceEnvironmentTestingID)):     &serviceEnvironmentTesting,
		strings.ToLower(string(serviceEnvironmentQA1ID)):         &serviceEnvironmentQA1,
		strings.ToLower(string(serviceEnvironmentQA2ID)):         &serviceEnvironmentQA2,
		strings.ToLower(string(serviceEnvironmentQA3ID)):         &serviceEnvironmentQA3,
		strings.ToLower(string(serviceEnvironmentQA4ID)):         &serviceEnvironmentQA4,
		strings.ToLower(string(serviceEnvironmentQA5ID)):         &serviceEnvironmentQA5,
		strings.ToLower(string(serviceEnvironmentQA6ID)):         &serviceEnvironmentQA6,
		strings.ToLower(string(serviceEnvironmentQACloud1ID)):    &serviceEnvironmentQACloud1,
		strings.ToLower(string(serviceEnvironmentQATestID)):      &serviceEnvironmentQATest,
		strings.ToLower(string(serviceEnvironmentCI1ID)):         &serviceEnvironmentCI1,
		strings.ToLower(string(serviceEnvironmentCI2ID)):         &serviceEnvironmentCI2,
		strings.ToLower(string(serviceEnvironmentSandboxID)):     &serviceEnvironmentSandbox,
		strings.ToLower(string(serviceEnvironmentUATID)):         &serviceEnvironmentUAT,
		strings.ToLower(string(serviceEnvironmentPartnerTestID)): &serviceEnvironmentPartnerTest,
		strings.ToLower(string(serviceEnvironmentProdSupID)):     &serviceEnvironmentProdSup,
		strings.ToLower(string(serviceEnvironmentProductionID)):  &serviceEnvironmentProduction,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumServiceEnvironment) ByID(id ServiceEnvironmentIdentifier) *EnumServiceEnvironmentItem {
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
func (e *EnumServiceEnvironment) ByIDString(idx string) *EnumServiceEnvironmentItem {
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
func (e *EnumServiceEnvironment) ByIndex(idx int) *EnumServiceEnvironmentItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedServiceEnvironmentID is a struct that is designed to replace a *ServiceEnvironmentID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *ServiceEnvironmentID it contains while being a better JSON citizen.
type ValidatedServiceEnvironmentID struct {
	// id will point to a valid ServiceEnvironmentID, if possible
	// If id is nil, then ValidatedServiceEnvironmentID.Valid() will return false.
	id *ServiceEnvironmentID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedServiceEnvironmentID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedServiceEnvironmentID
func (vi *ValidatedServiceEnvironmentID) Clone() *ValidatedServiceEnvironmentID {
	if vi == nil {
		return nil
	}

	var cid *ServiceEnvironmentID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedServiceEnvironmentID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedServiceEnvironmentIds represent the same ServiceEnvironment
func (vi *ValidatedServiceEnvironmentID) Equals(vj *ValidatedServiceEnvironmentID) bool {
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

// Valid returns true if and only if the ValidatedServiceEnvironmentID corresponds to a recognized ServiceEnvironment
func (vi *ValidatedServiceEnvironmentID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedServiceEnvironmentID) ID() *ServiceEnvironmentID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedServiceEnvironmentID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedServiceEnvironmentID) ValidatedID() *ValidatedServiceEnvironmentID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedServiceEnvironmentID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedServiceEnvironmentID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedServiceEnvironmentID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedServiceEnvironmentID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedServiceEnvironmentID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := ServiceEnvironmentID(capString)
	item := ServiceEnvironment.ByID(&id)
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

func (vi ValidatedServiceEnvironmentID) String() string {
	return vi.ToIDString()
}

type ServiceEnvironmentIdentifier interface {
	ID() *ServiceEnvironmentID
	Valid() bool
}
