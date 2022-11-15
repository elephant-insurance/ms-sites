package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// LicenseStatusID uniquely identifies a particular LicenseStatus
type LicenseStatusID string

// Clone creates a safe, independent copy of a LicenseStatusID
func (i *LicenseStatusID) Clone() *LicenseStatusID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two LicenseStatusIds are equivalent
func (i *LicenseStatusID) Equals(j *LicenseStatusID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *LicenseStatusID that is either valid or nil
func (i *LicenseStatusID) ID() *LicenseStatusID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *LicenseStatusID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the LicenseStatusID corresponds to a recognized LicenseStatus
func (i *LicenseStatusID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return LicenseStatus.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *LicenseStatusID) ValidatedID() *ValidatedLicenseStatusID {
	if i != nil {
		return &ValidatedLicenseStatusID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *LicenseStatusID) MarshalJSON() ([]byte, error) {
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

func (i *LicenseStatusID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := LicenseStatusID(dataString)
	item := LicenseStatus.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	licenseStatusValidID       LicenseStatusID = "valid"
	licenseStatusExpiredID     LicenseStatusID = "expired"
	licenseStatusNonLicensedID LicenseStatusID = "nonlicensed"
	licenseStatusForeignID     LicenseStatusID = "foreign"
	licenseStatusPermitID      LicenseStatusID = "permit"
	licenseStatusSurrenderedID LicenseStatusID = "surrendered"
	licenseStatusSuspendedID   LicenseStatusID = "suspended"
	licenseStatusRestrictedID  LicenseStatusID = "restricted"
	licenseStatusRevokedID     LicenseStatusID = "revoked"
)

// EnumLicenseStatusItem describes an entry in an enumeration of LicenseStatus
type EnumLicenseStatusItem struct {
	ID        LicenseStatusID   `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	StopPHSession string
}

var (
	licenseStatusValid       = EnumLicenseStatusItem{licenseStatusValidID, "Valid", map[string]string{"StopPHSession": ""}, "Valid", 1, ""}
	licenseStatusExpired     = EnumLicenseStatusItem{licenseStatusExpiredID, "Expired", map[string]string{"StopPHSession": "IN,TN,VA,TX,MD,IL,GA,OH"}, "Expired", 2, "IN,TN,VA,TX,MD,IL,GA,OH"}
	licenseStatusNonLicensed = EnumLicenseStatusItem{licenseStatusNonLicensedID, "Non-licensed", map[string]string{"StopPHSession": "IN,TN,VA,TX,MD,IL,GA,OH"}, "NonLicensed", 3, "IN,TN,VA,TX,MD,IL,GA,OH"}
	licenseStatusForeign     = EnumLicenseStatusItem{licenseStatusForeignID, "Non-US", map[string]string{"StopPHSession": ""}, "Foreign", 4, ""}
	licenseStatusPermit      = EnumLicenseStatusItem{licenseStatusPermitID, "Permit", map[string]string{"StopPHSession": "IN,TN,VA,TX,MD,IL,GA,OH"}, "Permit", 5, "IN,TN,VA,TX,MD,IL,GA,OH"}
	licenseStatusSurrendered = EnumLicenseStatusItem{licenseStatusSurrenderedID, "Surrendered", map[string]string{"StopPHSession": "IN,TN,VA,TX,MD,IL,GA,OH"}, "Surrendered", 6, "IN,TN,VA,TX,MD,IL,GA,OH"}
	licenseStatusSuspended   = EnumLicenseStatusItem{licenseStatusSuspendedID, "Suspended", map[string]string{"StopPHSession": "IN,TN,VA,TX,MD,IL,GA,OH"}, "Suspended", 7, "IN,TN,VA,TX,MD,IL,GA,OH"}
	licenseStatusRestricted  = EnumLicenseStatusItem{licenseStatusRestrictedID, "Restricted", map[string]string{"StopPHSession": ""}, "Restricted", 8, ""}
	licenseStatusRevoked     = EnumLicenseStatusItem{licenseStatusRevokedID, "Revoked", map[string]string{"StopPHSession": "IN,TN,VA,TX,MD,IL,GA,OH"}, "Revoked", 9, "IN,TN,VA,TX,MD,IL,GA,OH"}
)

// EnumLicenseStatus is a collection of LicenseStatus items
type EnumLicenseStatus struct {
	Description string
	Items       []*EnumLicenseStatusItem
	Name        string

	Valid       *EnumLicenseStatusItem
	Expired     *EnumLicenseStatusItem
	NonLicensed *EnumLicenseStatusItem
	Foreign     *EnumLicenseStatusItem
	Permit      *EnumLicenseStatusItem
	Surrendered *EnumLicenseStatusItem
	Suspended   *EnumLicenseStatusItem
	Restricted  *EnumLicenseStatusItem
	Revoked     *EnumLicenseStatusItem

	itemDict map[string]*EnumLicenseStatusItem
}

// LicenseStatus is a public singleton instance of EnumLicenseStatus
// representing legal statuses of drivers' licences
var LicenseStatus = &EnumLicenseStatus{
	Description: "legal statuses of drivers' licences",
	Items: []*EnumLicenseStatusItem{
		&licenseStatusValid,
		&licenseStatusExpired,
		&licenseStatusNonLicensed,
		&licenseStatusForeign,
		&licenseStatusPermit,
		&licenseStatusSurrendered,
		&licenseStatusSuspended,
		&licenseStatusRestricted,
		&licenseStatusRevoked,
	},
	Name:        "EnumLicenseStatus",
	Valid:       &licenseStatusValid,
	Expired:     &licenseStatusExpired,
	NonLicensed: &licenseStatusNonLicensed,
	Foreign:     &licenseStatusForeign,
	Permit:      &licenseStatusPermit,
	Surrendered: &licenseStatusSurrendered,
	Suspended:   &licenseStatusSuspended,
	Restricted:  &licenseStatusRestricted,
	Revoked:     &licenseStatusRevoked,

	itemDict: map[string]*EnumLicenseStatusItem{
		strings.ToLower(string(licenseStatusValidID)):       &licenseStatusValid,
		strings.ToLower(string(licenseStatusExpiredID)):     &licenseStatusExpired,
		strings.ToLower(string(licenseStatusNonLicensedID)): &licenseStatusNonLicensed,
		strings.ToLower(string(licenseStatusForeignID)):     &licenseStatusForeign,
		strings.ToLower(string(licenseStatusPermitID)):      &licenseStatusPermit,
		strings.ToLower(string(licenseStatusSurrenderedID)): &licenseStatusSurrendered,
		strings.ToLower(string(licenseStatusSuspendedID)):   &licenseStatusSuspended,
		strings.ToLower(string(licenseStatusRestrictedID)):  &licenseStatusRestricted,
		strings.ToLower(string(licenseStatusRevokedID)):     &licenseStatusRevoked,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumLicenseStatus) ByID(id LicenseStatusIdentifier) *EnumLicenseStatusItem {
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
func (e *EnumLicenseStatus) ByIDString(idx string) *EnumLicenseStatusItem {
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
func (e *EnumLicenseStatus) ByIndex(idx int) *EnumLicenseStatusItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedLicenseStatusID is a struct that is designed to replace a *LicenseStatusID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *LicenseStatusID it contains while being a better JSON citizen.
type ValidatedLicenseStatusID struct {
	// id will point to a valid LicenseStatusID, if possible
	// If id is nil, then ValidatedLicenseStatusID.Valid() will return false.
	id *LicenseStatusID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedLicenseStatusID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedLicenseStatusID
func (vi *ValidatedLicenseStatusID) Clone() *ValidatedLicenseStatusID {
	if vi == nil {
		return nil
	}

	var cid *LicenseStatusID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedLicenseStatusID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedLicenseStatusIds represent the same LicenseStatus
func (vi *ValidatedLicenseStatusID) Equals(vj *ValidatedLicenseStatusID) bool {
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

// Valid returns true if and only if the ValidatedLicenseStatusID corresponds to a recognized LicenseStatus
func (vi *ValidatedLicenseStatusID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedLicenseStatusID) ID() *LicenseStatusID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedLicenseStatusID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedLicenseStatusID) ValidatedID() *ValidatedLicenseStatusID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedLicenseStatusID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedLicenseStatusID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedLicenseStatusID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedLicenseStatusID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedLicenseStatusID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := LicenseStatusID(capString)
	item := LicenseStatus.ByID(&id)
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

func (vi ValidatedLicenseStatusID) String() string {
	return vi.ToIDString()
}

type LicenseStatusIdentifier interface {
	ID() *LicenseStatusID
	Valid() bool
}
