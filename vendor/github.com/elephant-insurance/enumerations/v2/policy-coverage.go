package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// PolicyCoverageID uniquely identifies a particular PolicyCoverage
type PolicyCoverageID string

// Clone creates a safe, independent copy of a PolicyCoverageID
func (i *PolicyCoverageID) Clone() *PolicyCoverageID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two PolicyCoverageIds are equivalent
func (i *PolicyCoverageID) Equals(j *PolicyCoverageID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *PolicyCoverageID that is either valid or nil
func (i *PolicyCoverageID) ID() *PolicyCoverageID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *PolicyCoverageID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the PolicyCoverageID corresponds to a recognized PolicyCoverage
func (i *PolicyCoverageID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return PolicyCoverage.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *PolicyCoverageID) ValidatedID() *ValidatedPolicyCoverageID {
	if i != nil {
		return &ValidatedPolicyCoverageID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *PolicyCoverageID) MarshalJSON() ([]byte, error) {
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

func (i *PolicyCoverageID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := PolicyCoverageID(dataString)
	item := PolicyCoverage.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	policyCoverageBodilyInjuryID                    PolicyCoverageID = "EISPABodilyInjuryCov"
	policyCoveragePropertyDamageID                  PolicyCoverageID = "EISPAPropertyDamageCov"
	policyCoverageUninsuredMotoristBodilyInjuryID   PolicyCoverageID = "PAUMBICov"
	policyCoverageMedPayID                          PolicyCoverageID = "PAMedPayCov"
	policyCoverageUninsuredMotoristPropertyDamageID PolicyCoverageID = "PAUMPDCov"
	policyCoveragePIPMarylandID                     PolicyCoverageID = "PAPIP_MD"
	policyCoveragePIPTexasID                        PolicyCoverageID = "PAPIP_TX"
	policyCoverageLegalPlanID                       PolicyCoverageID = "EISPALegalPlanCov"
	policyCoverageIncomeLossID                      PolicyCoverageID = "EISPAIncomeLossCov"
	policyCoverageUninsuredMotoristGeorgiaID        PolicyCoverageID = "PAUM_GACov"
)

// EnumPolicyCoverageItem describes an entry in an enumeration of PolicyCoverage
type EnumPolicyCoverageItem struct {
	ID        PolicyCoverageID  `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	policyCoverageBodilyInjury                    = EnumPolicyCoverageItem{policyCoverageBodilyInjuryID, "Bodily Injury Liability", nil, "BodilyInjury", 1}
	policyCoveragePropertyDamage                  = EnumPolicyCoverageItem{policyCoveragePropertyDamageID, "Property Damage Liability", nil, "PropertyDamage", 2}
	policyCoverageUninsuredMotoristBodilyInjury   = EnumPolicyCoverageItem{policyCoverageUninsuredMotoristBodilyInjuryID, "Uninsured/Underinsured Motorist - Bodily Injury", nil, "UninsuredMotoristBodilyInjury", 3}
	policyCoverageMedPay                          = EnumPolicyCoverageItem{policyCoverageMedPayID, "Medical Payments", nil, "MedPay", 4}
	policyCoverageUninsuredMotoristPropertyDamage = EnumPolicyCoverageItem{policyCoverageUninsuredMotoristPropertyDamageID, "Uninsured/Underinsured Motorist - Property Damage", nil, "UninsuredMotoristPropertyDamage", 5}
	policyCoveragePIPMaryland                     = EnumPolicyCoverageItem{policyCoveragePIPMarylandID, "PIP - Maryland", nil, "PIPMaryland", 6}
	policyCoveragePIPTexas                        = EnumPolicyCoverageItem{policyCoveragePIPTexasID, "PIP - Texas", nil, "PIPTexas", 7}
	policyCoverageLegalPlan                       = EnumPolicyCoverageItem{policyCoverageLegalPlanID, "Legal Plan", nil, "LegalPlan", 8}
	policyCoverageIncomeLoss                      = EnumPolicyCoverageItem{policyCoverageIncomeLossID, "Income Loss", nil, "IncomeLoss", 9}
	policyCoverageUninsuredMotoristGeorgia        = EnumPolicyCoverageItem{policyCoverageUninsuredMotoristGeorgiaID, "Georgia Added-On or Reduced Coverage", nil, "UninsuredMotoristGeorgia", 10}
)

// EnumPolicyCoverage is a collection of PolicyCoverage items
type EnumPolicyCoverage struct {
	Description string
	Items       []*EnumPolicyCoverageItem
	Name        string

	BodilyInjury                    *EnumPolicyCoverageItem
	PropertyDamage                  *EnumPolicyCoverageItem
	UninsuredMotoristBodilyInjury   *EnumPolicyCoverageItem
	MedPay                          *EnumPolicyCoverageItem
	UninsuredMotoristPropertyDamage *EnumPolicyCoverageItem
	PIPMaryland                     *EnumPolicyCoverageItem
	PIPTexas                        *EnumPolicyCoverageItem
	LegalPlan                       *EnumPolicyCoverageItem
	IncomeLoss                      *EnumPolicyCoverageItem
	UninsuredMotoristGeorgia        *EnumPolicyCoverageItem

	itemDict map[string]*EnumPolicyCoverageItem
}

// PolicyCoverage is a public singleton instance of EnumPolicyCoverage
// representing codes for policy coverages
var PolicyCoverage = &EnumPolicyCoverage{
	Description: "codes for policy coverages",
	Items: []*EnumPolicyCoverageItem{
		&policyCoverageBodilyInjury,
		&policyCoveragePropertyDamage,
		&policyCoverageUninsuredMotoristBodilyInjury,
		&policyCoverageMedPay,
		&policyCoverageUninsuredMotoristPropertyDamage,
		&policyCoveragePIPMaryland,
		&policyCoveragePIPTexas,
		&policyCoverageLegalPlan,
		&policyCoverageIncomeLoss,
		&policyCoverageUninsuredMotoristGeorgia,
	},
	Name:                            "EnumPolicyCoverage",
	BodilyInjury:                    &policyCoverageBodilyInjury,
	PropertyDamage:                  &policyCoveragePropertyDamage,
	UninsuredMotoristBodilyInjury:   &policyCoverageUninsuredMotoristBodilyInjury,
	MedPay:                          &policyCoverageMedPay,
	UninsuredMotoristPropertyDamage: &policyCoverageUninsuredMotoristPropertyDamage,
	PIPMaryland:                     &policyCoveragePIPMaryland,
	PIPTexas:                        &policyCoveragePIPTexas,
	LegalPlan:                       &policyCoverageLegalPlan,
	IncomeLoss:                      &policyCoverageIncomeLoss,
	UninsuredMotoristGeorgia:        &policyCoverageUninsuredMotoristGeorgia,

	itemDict: map[string]*EnumPolicyCoverageItem{
		strings.ToLower(string(policyCoverageBodilyInjuryID)):                    &policyCoverageBodilyInjury,
		strings.ToLower(string(policyCoveragePropertyDamageID)):                  &policyCoveragePropertyDamage,
		strings.ToLower(string(policyCoverageUninsuredMotoristBodilyInjuryID)):   &policyCoverageUninsuredMotoristBodilyInjury,
		strings.ToLower(string(policyCoverageMedPayID)):                          &policyCoverageMedPay,
		strings.ToLower(string(policyCoverageUninsuredMotoristPropertyDamageID)): &policyCoverageUninsuredMotoristPropertyDamage,
		strings.ToLower(string(policyCoveragePIPMarylandID)):                     &policyCoveragePIPMaryland,
		strings.ToLower(string(policyCoveragePIPTexasID)):                        &policyCoveragePIPTexas,
		strings.ToLower(string(policyCoverageLegalPlanID)):                       &policyCoverageLegalPlan,
		strings.ToLower(string(policyCoverageIncomeLossID)):                      &policyCoverageIncomeLoss,
		strings.ToLower(string(policyCoverageUninsuredMotoristGeorgiaID)):        &policyCoverageUninsuredMotoristGeorgia,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumPolicyCoverage) ByID(id PolicyCoverageIdentifier) *EnumPolicyCoverageItem {
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
func (e *EnumPolicyCoverage) ByIDString(idx string) *EnumPolicyCoverageItem {
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
func (e *EnumPolicyCoverage) ByIndex(idx int) *EnumPolicyCoverageItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedPolicyCoverageID is a struct that is designed to replace a *PolicyCoverageID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *PolicyCoverageID it contains while being a better JSON citizen.
type ValidatedPolicyCoverageID struct {
	// id will point to a valid PolicyCoverageID, if possible
	// If id is nil, then ValidatedPolicyCoverageID.Valid() will return false.
	id *PolicyCoverageID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedPolicyCoverageID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedPolicyCoverageID
func (vi *ValidatedPolicyCoverageID) Clone() *ValidatedPolicyCoverageID {
	if vi == nil {
		return nil
	}

	var cid *PolicyCoverageID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedPolicyCoverageID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedPolicyCoverageIds represent the same PolicyCoverage
func (vi *ValidatedPolicyCoverageID) Equals(vj *ValidatedPolicyCoverageID) bool {
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

// Valid returns true if and only if the ValidatedPolicyCoverageID corresponds to a recognized PolicyCoverage
func (vi *ValidatedPolicyCoverageID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedPolicyCoverageID) ID() *PolicyCoverageID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedPolicyCoverageID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedPolicyCoverageID) ValidatedID() *ValidatedPolicyCoverageID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedPolicyCoverageID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedPolicyCoverageID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedPolicyCoverageID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedPolicyCoverageID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedPolicyCoverageID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := PolicyCoverageID(capString)
	item := PolicyCoverage.ByID(&id)
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

func (vi ValidatedPolicyCoverageID) String() string {
	return vi.ToIDString()
}

type PolicyCoverageIdentifier interface {
	ID() *PolicyCoverageID
	Valid() bool
}
