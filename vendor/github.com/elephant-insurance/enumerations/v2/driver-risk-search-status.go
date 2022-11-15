package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// DriverRiskSearchStatusID uniquely identifies a particular DriverRiskSearchStatus
type DriverRiskSearchStatusID string

// Clone creates a safe, independent copy of a DriverRiskSearchStatusID
func (i *DriverRiskSearchStatusID) Clone() *DriverRiskSearchStatusID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two DriverRiskSearchStatusIds are equivalent
func (i *DriverRiskSearchStatusID) Equals(j *DriverRiskSearchStatusID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *DriverRiskSearchStatusID that is either valid or nil
func (i *DriverRiskSearchStatusID) ID() *DriverRiskSearchStatusID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *DriverRiskSearchStatusID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the DriverRiskSearchStatusID corresponds to a recognized DriverRiskSearchStatus
func (i *DriverRiskSearchStatusID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return DriverRiskSearchStatus.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *DriverRiskSearchStatusID) ValidatedID() *ValidatedDriverRiskSearchStatusID {
	if i != nil {
		return &ValidatedDriverRiskSearchStatusID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *DriverRiskSearchStatusID) MarshalJSON() ([]byte, error) {
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

func (i *DriverRiskSearchStatusID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := DriverRiskSearchStatusID(dataString)
	item := DriverRiskSearchStatus.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	driverRiskSearchStatusOriginalHitID                   DriverRiskSearchStatusID = "originalHit"
	driverRiskSearchStatusOriginalHitClearID              DriverRiskSearchStatusID = "originalHitClear"
	driverRiskSearchStatusOriginalHitRatableID            DriverRiskSearchStatusID = "originalHitRatable"
	driverRiskSearchStatusOriginalHitNotRatableID         DriverRiskSearchStatusID = "originalHitNotRatable"
	driverRiskSearchStatusOriginalNotFoundID              DriverRiskSearchStatusID = "originalNotFound"
	driverRiskSearchStatusDuplicateHitID                  DriverRiskSearchStatusID = "duplicateHit"
	driverRiskSearchStatusDuplicateHitClearID             DriverRiskSearchStatusID = "duplicateHitClear"
	driverRiskSearchStatusDuplicateHitArchiveID           DriverRiskSearchStatusID = "duplicateHitArchive"
	driverRiskSearchStatusDuplicateHitArchiveClearID      DriverRiskSearchStatusID = "duplicateHitArchiveClear"
	driverRiskSearchStatusDuplicateHitRatableID           DriverRiskSearchStatusID = "duplicateHitRatable"
	driverRiskSearchStatusDuplicateHitNotRatableID        DriverRiskSearchStatusID = "duplicateHitNotRatable"
	driverRiskSearchStatusDuplicateNotFoundID             DriverRiskSearchStatusID = "duplicateNotFound"
	driverRiskSearchStatusDuplicateHitArchiveRatableID    DriverRiskSearchStatusID = "duplicateHitArchiveRatable"
	driverRiskSearchStatusDuplicateHitArchiveNotRatableID DriverRiskSearchStatusID = "duplicateHitArchiveNotRatable"
	driverRiskSearchStatusDuplicateNotFoundArchiveID      DriverRiskSearchStatusID = "duplicateNotFoundArchive"
)

// EnumDriverRiskSearchStatusItem describes an entry in an enumeration of DriverRiskSearchStatus
type EnumDriverRiskSearchStatusItem struct {
	ID        DriverRiskSearchStatusID `json:"Value"`
	Desc      string                   `json:"Description,omitempty"`
	Meta      map[string]string        `json:",omitempty"`
	Name      string                   `json:"Name"`
	SortOrder int
}

var (
	driverRiskSearchStatusOriginalHit                   = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusOriginalHitID, "OriginalHit", nil, "OriginalHit", 1}
	driverRiskSearchStatusOriginalHitClear              = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusOriginalHitClearID, "OriginalHitClear", nil, "OriginalHitClear", 2}
	driverRiskSearchStatusOriginalHitRatable            = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusOriginalHitRatableID, "OriginalHitRatable", nil, "OriginalHitRatable", 3}
	driverRiskSearchStatusOriginalHitNotRatable         = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusOriginalHitNotRatableID, "OriginalHitNotRatable", nil, "OriginalHitNotRatable", 4}
	driverRiskSearchStatusOriginalNotFound              = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusOriginalNotFoundID, "OriginalNotFound", nil, "OriginalNotFound", 5}
	driverRiskSearchStatusDuplicateHit                  = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusDuplicateHitID, "DuplicateHit", nil, "DuplicateHit", 6}
	driverRiskSearchStatusDuplicateHitClear             = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusDuplicateHitClearID, "DuplicateHitClear", nil, "DuplicateHitClear", 7}
	driverRiskSearchStatusDuplicateHitArchive           = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusDuplicateHitArchiveID, "DuplicateHitArchive", nil, "DuplicateHitArchive", 8}
	driverRiskSearchStatusDuplicateHitArchiveClear      = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusDuplicateHitArchiveClearID, "DuplicateHitArchiveClear", nil, "DuplicateHitArchiveClear", 9}
	driverRiskSearchStatusDuplicateHitRatable           = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusDuplicateHitRatableID, "DuplicateHitRatable", nil, "DuplicateHitRatable", 10}
	driverRiskSearchStatusDuplicateHitNotRatable        = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusDuplicateHitNotRatableID, "DuplicateHitNotRatable", nil, "DuplicateHitNotRatable", 11}
	driverRiskSearchStatusDuplicateNotFound             = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusDuplicateNotFoundID, "DuplicateNotFound", nil, "DuplicateNotFound", 12}
	driverRiskSearchStatusDuplicateHitArchiveRatable    = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusDuplicateHitArchiveRatableID, "DuplicateHitArchiveRatable", nil, "DuplicateHitArchiveRatable", 13}
	driverRiskSearchStatusDuplicateHitArchiveNotRatable = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusDuplicateHitArchiveNotRatableID, "DuplicateHitArchiveNotRatable", nil, "DuplicateHitArchiveNotRatable", 14}
	driverRiskSearchStatusDuplicateNotFoundArchive      = EnumDriverRiskSearchStatusItem{driverRiskSearchStatusDuplicateNotFoundArchiveID, "DuplicateNotFoundArchive", nil, "DuplicateNotFoundArchive", 15}
)

// EnumDriverRiskSearchStatus is a collection of DriverRiskSearchStatus items
type EnumDriverRiskSearchStatus struct {
	Description string
	Items       []*EnumDriverRiskSearchStatusItem
	Name        string

	OriginalHit                   *EnumDriverRiskSearchStatusItem
	OriginalHitClear              *EnumDriverRiskSearchStatusItem
	OriginalHitRatable            *EnumDriverRiskSearchStatusItem
	OriginalHitNotRatable         *EnumDriverRiskSearchStatusItem
	OriginalNotFound              *EnumDriverRiskSearchStatusItem
	DuplicateHit                  *EnumDriverRiskSearchStatusItem
	DuplicateHitClear             *EnumDriverRiskSearchStatusItem
	DuplicateHitArchive           *EnumDriverRiskSearchStatusItem
	DuplicateHitArchiveClear      *EnumDriverRiskSearchStatusItem
	DuplicateHitRatable           *EnumDriverRiskSearchStatusItem
	DuplicateHitNotRatable        *EnumDriverRiskSearchStatusItem
	DuplicateNotFound             *EnumDriverRiskSearchStatusItem
	DuplicateHitArchiveRatable    *EnumDriverRiskSearchStatusItem
	DuplicateHitArchiveNotRatable *EnumDriverRiskSearchStatusItem
	DuplicateNotFoundArchive      *EnumDriverRiskSearchStatusItem

	itemDict map[string]*EnumDriverRiskSearchStatusItem
}

// DriverRiskSearchStatus is a public singleton instance of EnumDriverRiskSearchStatus
// representing Status for driver risk search
var DriverRiskSearchStatus = &EnumDriverRiskSearchStatus{
	Description: "Status for driver risk search",
	Items: []*EnumDriverRiskSearchStatusItem{
		&driverRiskSearchStatusOriginalHit,
		&driverRiskSearchStatusOriginalHitClear,
		&driverRiskSearchStatusOriginalHitRatable,
		&driverRiskSearchStatusOriginalHitNotRatable,
		&driverRiskSearchStatusOriginalNotFound,
		&driverRiskSearchStatusDuplicateHit,
		&driverRiskSearchStatusDuplicateHitClear,
		&driverRiskSearchStatusDuplicateHitArchive,
		&driverRiskSearchStatusDuplicateHitArchiveClear,
		&driverRiskSearchStatusDuplicateHitRatable,
		&driverRiskSearchStatusDuplicateHitNotRatable,
		&driverRiskSearchStatusDuplicateNotFound,
		&driverRiskSearchStatusDuplicateHitArchiveRatable,
		&driverRiskSearchStatusDuplicateHitArchiveNotRatable,
		&driverRiskSearchStatusDuplicateNotFoundArchive,
	},
	Name:                          "EnumDriverRiskSearchStatus",
	OriginalHit:                   &driverRiskSearchStatusOriginalHit,
	OriginalHitClear:              &driverRiskSearchStatusOriginalHitClear,
	OriginalHitRatable:            &driverRiskSearchStatusOriginalHitRatable,
	OriginalHitNotRatable:         &driverRiskSearchStatusOriginalHitNotRatable,
	OriginalNotFound:              &driverRiskSearchStatusOriginalNotFound,
	DuplicateHit:                  &driverRiskSearchStatusDuplicateHit,
	DuplicateHitClear:             &driverRiskSearchStatusDuplicateHitClear,
	DuplicateHitArchive:           &driverRiskSearchStatusDuplicateHitArchive,
	DuplicateHitArchiveClear:      &driverRiskSearchStatusDuplicateHitArchiveClear,
	DuplicateHitRatable:           &driverRiskSearchStatusDuplicateHitRatable,
	DuplicateHitNotRatable:        &driverRiskSearchStatusDuplicateHitNotRatable,
	DuplicateNotFound:             &driverRiskSearchStatusDuplicateNotFound,
	DuplicateHitArchiveRatable:    &driverRiskSearchStatusDuplicateHitArchiveRatable,
	DuplicateHitArchiveNotRatable: &driverRiskSearchStatusDuplicateHitArchiveNotRatable,
	DuplicateNotFoundArchive:      &driverRiskSearchStatusDuplicateNotFoundArchive,

	itemDict: map[string]*EnumDriverRiskSearchStatusItem{
		strings.ToLower(string(driverRiskSearchStatusOriginalHitID)):                   &driverRiskSearchStatusOriginalHit,
		strings.ToLower(string(driverRiskSearchStatusOriginalHitClearID)):              &driverRiskSearchStatusOriginalHitClear,
		strings.ToLower(string(driverRiskSearchStatusOriginalHitRatableID)):            &driverRiskSearchStatusOriginalHitRatable,
		strings.ToLower(string(driverRiskSearchStatusOriginalHitNotRatableID)):         &driverRiskSearchStatusOriginalHitNotRatable,
		strings.ToLower(string(driverRiskSearchStatusOriginalNotFoundID)):              &driverRiskSearchStatusOriginalNotFound,
		strings.ToLower(string(driverRiskSearchStatusDuplicateHitID)):                  &driverRiskSearchStatusDuplicateHit,
		strings.ToLower(string(driverRiskSearchStatusDuplicateHitClearID)):             &driverRiskSearchStatusDuplicateHitClear,
		strings.ToLower(string(driverRiskSearchStatusDuplicateHitArchiveID)):           &driverRiskSearchStatusDuplicateHitArchive,
		strings.ToLower(string(driverRiskSearchStatusDuplicateHitArchiveClearID)):      &driverRiskSearchStatusDuplicateHitArchiveClear,
		strings.ToLower(string(driverRiskSearchStatusDuplicateHitRatableID)):           &driverRiskSearchStatusDuplicateHitRatable,
		strings.ToLower(string(driverRiskSearchStatusDuplicateHitNotRatableID)):        &driverRiskSearchStatusDuplicateHitNotRatable,
		strings.ToLower(string(driverRiskSearchStatusDuplicateNotFoundID)):             &driverRiskSearchStatusDuplicateNotFound,
		strings.ToLower(string(driverRiskSearchStatusDuplicateHitArchiveRatableID)):    &driverRiskSearchStatusDuplicateHitArchiveRatable,
		strings.ToLower(string(driverRiskSearchStatusDuplicateHitArchiveNotRatableID)): &driverRiskSearchStatusDuplicateHitArchiveNotRatable,
		strings.ToLower(string(driverRiskSearchStatusDuplicateNotFoundArchiveID)):      &driverRiskSearchStatusDuplicateNotFoundArchive,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumDriverRiskSearchStatus) ByID(id DriverRiskSearchStatusIdentifier) *EnumDriverRiskSearchStatusItem {
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
func (e *EnumDriverRiskSearchStatus) ByIDString(idx string) *EnumDriverRiskSearchStatusItem {
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
func (e *EnumDriverRiskSearchStatus) ByIndex(idx int) *EnumDriverRiskSearchStatusItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedDriverRiskSearchStatusID is a struct that is designed to replace a *DriverRiskSearchStatusID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *DriverRiskSearchStatusID it contains while being a better JSON citizen.
type ValidatedDriverRiskSearchStatusID struct {
	// id will point to a valid DriverRiskSearchStatusID, if possible
	// If id is nil, then ValidatedDriverRiskSearchStatusID.Valid() will return false.
	id *DriverRiskSearchStatusID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedDriverRiskSearchStatusID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedDriverRiskSearchStatusID
func (vi *ValidatedDriverRiskSearchStatusID) Clone() *ValidatedDriverRiskSearchStatusID {
	if vi == nil {
		return nil
	}

	var cid *DriverRiskSearchStatusID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedDriverRiskSearchStatusID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedDriverRiskSearchStatusIds represent the same DriverRiskSearchStatus
func (vi *ValidatedDriverRiskSearchStatusID) Equals(vj *ValidatedDriverRiskSearchStatusID) bool {
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

// Valid returns true if and only if the ValidatedDriverRiskSearchStatusID corresponds to a recognized DriverRiskSearchStatus
func (vi *ValidatedDriverRiskSearchStatusID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedDriverRiskSearchStatusID) ID() *DriverRiskSearchStatusID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedDriverRiskSearchStatusID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedDriverRiskSearchStatusID) ValidatedID() *ValidatedDriverRiskSearchStatusID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedDriverRiskSearchStatusID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedDriverRiskSearchStatusID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedDriverRiskSearchStatusID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedDriverRiskSearchStatusID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedDriverRiskSearchStatusID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := DriverRiskSearchStatusID(capString)
	item := DriverRiskSearchStatus.ByID(&id)
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

func (vi ValidatedDriverRiskSearchStatusID) String() string {
	return vi.ToIDString()
}

type DriverRiskSearchStatusIdentifier interface {
	ID() *DriverRiskSearchStatusID
	Valid() bool
}
