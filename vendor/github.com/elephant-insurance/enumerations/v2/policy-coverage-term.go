package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// PolicyCoverageTermID uniquely identifies a particular PolicyCoverageTerm
type PolicyCoverageTermID string

// Clone creates a safe, independent copy of a PolicyCoverageTermID
func (i *PolicyCoverageTermID) Clone() *PolicyCoverageTermID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two PolicyCoverageTermIds are equivalent
func (i *PolicyCoverageTermID) Equals(j *PolicyCoverageTermID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *PolicyCoverageTermID that is either valid or nil
func (i *PolicyCoverageTermID) ID() *PolicyCoverageTermID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *PolicyCoverageTermID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the PolicyCoverageTermID corresponds to a recognized PolicyCoverageTerm
func (i *PolicyCoverageTermID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return PolicyCoverageTerm.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *PolicyCoverageTermID) ValidatedID() *ValidatedPolicyCoverageTermID {
	if i != nil {
		return &ValidatedPolicyCoverageTermID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *PolicyCoverageTermID) MarshalJSON() ([]byte, error) {
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

func (i *PolicyCoverageTermID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := PolicyCoverageTermID(dataString)
	item := PolicyCoverageTerm.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	policyCoverageTermBodilyInjuryLimitsID                    PolicyCoverageTermID = "PABodilyInjury"
	policyCoverageTermPropertyDamageLimitsID                  PolicyCoverageTermID = "PAPropertyDamage"
	policyCoverageTermUninsuredMotoristBodilyInjuryLimitsID   PolicyCoverageTermID = "PAUMBI"
	policyCoverageTermMedPayLimitsID                          PolicyCoverageTermID = "PAMedLimit"
	policyCoverageTermUninsuredMotoristPropertyDamageLimitsID PolicyCoverageTermID = "PAUMPDLimit"
	policyCoverageTermUninsuredMotoristGeorgiaAddOnID         PolicyCoverageTermID = "UM_GA_AddOn"
	policyCoverageTermUMPDIndianaDeductibleID                 PolicyCoverageTermID = "UMPD_IN"
	policyCoverageTermUMPDGeorgiaDeductibleID                 PolicyCoverageTermID = "UMPD_GA"
	policyCoverageTermPIPMarylandLimitsID                     PolicyCoverageTermID = "PAPIPMD_LIMIT"
	policyCoverageTermPIPMarylandGuestID                      PolicyCoverageTermID = "PAPIPMD_GUEST"
	policyCoverageTermPIPTexasLimitsID                        PolicyCoverageTermID = "PAPIPTX_LIMIT"
	policyCoverageTermLegalPlanFamilyID                       PolicyCoverageTermID = "PAFamilyLegalPlan"
	policyCoverageTermLegalPlanIndividualID                   PolicyCoverageTermID = "PASingleLegalPlan"
	policyCoverageTermIncomeLossLimitID                       PolicyCoverageTermID = "IncomeLossLimit"
)

// EnumPolicyCoverageTermItem describes an entry in an enumeration of PolicyCoverageTerm
type EnumPolicyCoverageTermItem struct {
	ID        PolicyCoverageTermID `json:"Value"`
	Desc      string               `json:"Description,omitempty"`
	Meta      map[string]string    `json:",omitempty"`
	Name      string               `json:"Name"`
	SortOrder int
}

var (
	policyCoverageTermBodilyInjuryLimits                    = EnumPolicyCoverageTermItem{policyCoverageTermBodilyInjuryLimitsID, "Bodily Injury Limits", nil, "BodilyInjuryLimits", 1}
	policyCoverageTermPropertyDamageLimits                  = EnumPolicyCoverageTermItem{policyCoverageTermPropertyDamageLimitsID, "Property Damage Limits", nil, "PropertyDamageLimits", 2}
	policyCoverageTermUninsuredMotoristBodilyInjuryLimits   = EnumPolicyCoverageTermItem{policyCoverageTermUninsuredMotoristBodilyInjuryLimitsID, "Uninsured Motorist Bodily Injury Limits", nil, "UninsuredMotoristBodilyInjuryLimits", 3}
	policyCoverageTermMedPayLimits                          = EnumPolicyCoverageTermItem{policyCoverageTermMedPayLimitsID, "Medical Payment Limits", nil, "MedPayLimits", 4}
	policyCoverageTermUninsuredMotoristPropertyDamageLimits = EnumPolicyCoverageTermItem{policyCoverageTermUninsuredMotoristPropertyDamageLimitsID, "Uninsured Motorist Bodily Injury Limits", nil, "UninsuredMotoristPropertyDamageLimits", 5}
	policyCoverageTermUninsuredMotoristGeorgiaAddOn         = EnumPolicyCoverageTermItem{policyCoverageTermUninsuredMotoristGeorgiaAddOnID, "Uninsured Motorist Georgia Add-On", nil, "UninsuredMotoristGeorgiaAddOn", 6}
	policyCoverageTermUMPDIndianaDeductible                 = EnumPolicyCoverageTermItem{policyCoverageTermUMPDIndianaDeductibleID, "Indiana Uninsured Motorist Property Damage Deductible", nil, "UMPDIndianaDeductible", 7}
	policyCoverageTermUMPDGeorgiaDeductible                 = EnumPolicyCoverageTermItem{policyCoverageTermUMPDGeorgiaDeductibleID, "Georgia Uninsured Motorist Property Damage Deductible", nil, "UMPDGeorgiaDeductible", 8}
	policyCoverageTermPIPMarylandLimits                     = EnumPolicyCoverageTermItem{policyCoverageTermPIPMarylandLimitsID, "Maryland PIP Limits", nil, "PIPMarylandLimits", 9}
	policyCoverageTermPIPMarylandGuest                      = EnumPolicyCoverageTermItem{policyCoverageTermPIPMarylandGuestID, "Maryland PIP Guest Coverage", nil, "PIPMarylandGuest", 10}
	policyCoverageTermPIPTexasLimits                        = EnumPolicyCoverageTermItem{policyCoverageTermPIPTexasLimitsID, "Texas PIP Limits", nil, "PIPTexasLimits", 11}
	policyCoverageTermLegalPlanFamily                       = EnumPolicyCoverageTermItem{policyCoverageTermLegalPlanFamilyID, "Legal Plan - Family", nil, "LegalPlanFamily", 12}
	policyCoverageTermLegalPlanIndividual                   = EnumPolicyCoverageTermItem{policyCoverageTermLegalPlanIndividualID, "Legal Plan - Individual", nil, "LegalPlanIndividual", 13}
	policyCoverageTermIncomeLossLimit                       = EnumPolicyCoverageTermItem{policyCoverageTermIncomeLossLimitID, "IncomeLossLimit", nil, "IncomeLossLimit", 14}
)

// EnumPolicyCoverageTerm is a collection of PolicyCoverageTerm items
type EnumPolicyCoverageTerm struct {
	Description string
	Items       []*EnumPolicyCoverageTermItem
	Name        string

	BodilyInjuryLimits                    *EnumPolicyCoverageTermItem
	PropertyDamageLimits                  *EnumPolicyCoverageTermItem
	UninsuredMotoristBodilyInjuryLimits   *EnumPolicyCoverageTermItem
	MedPayLimits                          *EnumPolicyCoverageTermItem
	UninsuredMotoristPropertyDamageLimits *EnumPolicyCoverageTermItem
	UninsuredMotoristGeorgiaAddOn         *EnumPolicyCoverageTermItem
	UMPDIndianaDeductible                 *EnumPolicyCoverageTermItem
	UMPDGeorgiaDeductible                 *EnumPolicyCoverageTermItem
	PIPMarylandLimits                     *EnumPolicyCoverageTermItem
	PIPMarylandGuest                      *EnumPolicyCoverageTermItem
	PIPTexasLimits                        *EnumPolicyCoverageTermItem
	LegalPlanFamily                       *EnumPolicyCoverageTermItem
	LegalPlanIndividual                   *EnumPolicyCoverageTermItem
	IncomeLossLimit                       *EnumPolicyCoverageTermItem

	itemDict map[string]*EnumPolicyCoverageTermItem
}

// PolicyCoverageTerm is a public singleton instance of EnumPolicyCoverageTerm
// representing term codes for policy coverages
var PolicyCoverageTerm = &EnumPolicyCoverageTerm{
	Description: "term codes for policy coverages",
	Items: []*EnumPolicyCoverageTermItem{
		&policyCoverageTermBodilyInjuryLimits,
		&policyCoverageTermPropertyDamageLimits,
		&policyCoverageTermUninsuredMotoristBodilyInjuryLimits,
		&policyCoverageTermMedPayLimits,
		&policyCoverageTermUninsuredMotoristPropertyDamageLimits,
		&policyCoverageTermUninsuredMotoristGeorgiaAddOn,
		&policyCoverageTermUMPDIndianaDeductible,
		&policyCoverageTermUMPDGeorgiaDeductible,
		&policyCoverageTermPIPMarylandLimits,
		&policyCoverageTermPIPMarylandGuest,
		&policyCoverageTermPIPTexasLimits,
		&policyCoverageTermLegalPlanFamily,
		&policyCoverageTermLegalPlanIndividual,
		&policyCoverageTermIncomeLossLimit,
	},
	Name:                                  "EnumPolicyCoverageTerm",
	BodilyInjuryLimits:                    &policyCoverageTermBodilyInjuryLimits,
	PropertyDamageLimits:                  &policyCoverageTermPropertyDamageLimits,
	UninsuredMotoristBodilyInjuryLimits:   &policyCoverageTermUninsuredMotoristBodilyInjuryLimits,
	MedPayLimits:                          &policyCoverageTermMedPayLimits,
	UninsuredMotoristPropertyDamageLimits: &policyCoverageTermUninsuredMotoristPropertyDamageLimits,
	UninsuredMotoristGeorgiaAddOn:         &policyCoverageTermUninsuredMotoristGeorgiaAddOn,
	UMPDIndianaDeductible:                 &policyCoverageTermUMPDIndianaDeductible,
	UMPDGeorgiaDeductible:                 &policyCoverageTermUMPDGeorgiaDeductible,
	PIPMarylandLimits:                     &policyCoverageTermPIPMarylandLimits,
	PIPMarylandGuest:                      &policyCoverageTermPIPMarylandGuest,
	PIPTexasLimits:                        &policyCoverageTermPIPTexasLimits,
	LegalPlanFamily:                       &policyCoverageTermLegalPlanFamily,
	LegalPlanIndividual:                   &policyCoverageTermLegalPlanIndividual,
	IncomeLossLimit:                       &policyCoverageTermIncomeLossLimit,

	itemDict: map[string]*EnumPolicyCoverageTermItem{
		strings.ToLower(string(policyCoverageTermBodilyInjuryLimitsID)):                    &policyCoverageTermBodilyInjuryLimits,
		strings.ToLower(string(policyCoverageTermPropertyDamageLimitsID)):                  &policyCoverageTermPropertyDamageLimits,
		strings.ToLower(string(policyCoverageTermUninsuredMotoristBodilyInjuryLimitsID)):   &policyCoverageTermUninsuredMotoristBodilyInjuryLimits,
		strings.ToLower(string(policyCoverageTermMedPayLimitsID)):                          &policyCoverageTermMedPayLimits,
		strings.ToLower(string(policyCoverageTermUninsuredMotoristPropertyDamageLimitsID)): &policyCoverageTermUninsuredMotoristPropertyDamageLimits,
		strings.ToLower(string(policyCoverageTermUninsuredMotoristGeorgiaAddOnID)):         &policyCoverageTermUninsuredMotoristGeorgiaAddOn,
		strings.ToLower(string(policyCoverageTermUMPDIndianaDeductibleID)):                 &policyCoverageTermUMPDIndianaDeductible,
		strings.ToLower(string(policyCoverageTermUMPDGeorgiaDeductibleID)):                 &policyCoverageTermUMPDGeorgiaDeductible,
		strings.ToLower(string(policyCoverageTermPIPMarylandLimitsID)):                     &policyCoverageTermPIPMarylandLimits,
		strings.ToLower(string(policyCoverageTermPIPMarylandGuestID)):                      &policyCoverageTermPIPMarylandGuest,
		strings.ToLower(string(policyCoverageTermPIPTexasLimitsID)):                        &policyCoverageTermPIPTexasLimits,
		strings.ToLower(string(policyCoverageTermLegalPlanFamilyID)):                       &policyCoverageTermLegalPlanFamily,
		strings.ToLower(string(policyCoverageTermLegalPlanIndividualID)):                   &policyCoverageTermLegalPlanIndividual,
		strings.ToLower(string(policyCoverageTermIncomeLossLimitID)):                       &policyCoverageTermIncomeLossLimit,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumPolicyCoverageTerm) ByID(id PolicyCoverageTermIdentifier) *EnumPolicyCoverageTermItem {
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
func (e *EnumPolicyCoverageTerm) ByIDString(idx string) *EnumPolicyCoverageTermItem {
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
func (e *EnumPolicyCoverageTerm) ByIndex(idx int) *EnumPolicyCoverageTermItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedPolicyCoverageTermID is a struct that is designed to replace a *PolicyCoverageTermID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *PolicyCoverageTermID it contains while being a better JSON citizen.
type ValidatedPolicyCoverageTermID struct {
	// id will point to a valid PolicyCoverageTermID, if possible
	// If id is nil, then ValidatedPolicyCoverageTermID.Valid() will return false.
	id *PolicyCoverageTermID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedPolicyCoverageTermID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedPolicyCoverageTermID
func (vi *ValidatedPolicyCoverageTermID) Clone() *ValidatedPolicyCoverageTermID {
	if vi == nil {
		return nil
	}

	var cid *PolicyCoverageTermID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedPolicyCoverageTermID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedPolicyCoverageTermIds represent the same PolicyCoverageTerm
func (vi *ValidatedPolicyCoverageTermID) Equals(vj *ValidatedPolicyCoverageTermID) bool {
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

// Valid returns true if and only if the ValidatedPolicyCoverageTermID corresponds to a recognized PolicyCoverageTerm
func (vi *ValidatedPolicyCoverageTermID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedPolicyCoverageTermID) ID() *PolicyCoverageTermID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedPolicyCoverageTermID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedPolicyCoverageTermID) ValidatedID() *ValidatedPolicyCoverageTermID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedPolicyCoverageTermID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedPolicyCoverageTermID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedPolicyCoverageTermID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedPolicyCoverageTermID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedPolicyCoverageTermID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := PolicyCoverageTermID(capString)
	item := PolicyCoverageTerm.ByID(&id)
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

func (vi ValidatedPolicyCoverageTermID) String() string {
	return vi.ToIDString()
}

type PolicyCoverageTermIdentifier interface {
	ID() *PolicyCoverageTermID
	Valid() bool
}
