package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// WidgetID uniquely identifies a particular Widget
type WidgetID string

// Clone creates a safe, independent copy of a WidgetID
func (i *WidgetID) Clone() *WidgetID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two WidgetIds are equivalent
func (i *WidgetID) Equals(j *WidgetID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *WidgetID that is either valid or nil
func (i *WidgetID) ID() *WidgetID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

func (i *WidgetID) Parent() *WidgetID {
	if this := Widget.ByIDString(string(*i)); this != nil {
		if this.Parent != nil {
			return &this.Parent.ID
		}

		return i
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *WidgetID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the WidgetID corresponds to a recognized Widget
func (i *WidgetID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return Widget.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *WidgetID) ValidatedID() *ValidatedWidgetID {
	if i != nil {
		return &ValidatedWidgetID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *WidgetID) MarshalJSON() ([]byte, error) {
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

func (i *WidgetID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := WidgetID(dataString)
	item := Widget.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	widgetBarneyID WidgetID = "barney_rubble"
	widgetBettyID  WidgetID = "betty_rubble"
	widgetFredID   WidgetID = "fred_flintstone"
	widgetWilmaID  WidgetID = "wilma_flintstone"
	widgetElroyID  WidgetID = "elroy_jetson"
	widgetGeorgeID WidgetID = "george_jetson"
	widgetJaneID   WidgetID = "jane_jetson"
	widgetJudyID   WidgetID = "judy_jetson"
	widgetBartID   WidgetID = "bart_simpson"
	widgetHomerID  WidgetID = "homer_simpson"
	widgetLisaID   WidgetID = "lisa_simpson"
	widgetMargeID  WidgetID = "marge_simpson"
)

// EnumWidgetItem describes an entry in an enumeration of Widget
type EnumWidgetItem struct {
	ID        WidgetID          `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	Parent    *EnumWidgetItem   `json:"-"`
	SortOrder int

	// Meta Properties
	CostumeColor string
	Role         string
}

var (
	widgetBarney = EnumWidgetItem{widgetBarneyID, "Barney Rubble", map[string]string{"CostumeColor": "red", "Role": "neighbor"}, "Barney", nil, 1, "red", "neighbor"}
	widgetBetty  = EnumWidgetItem{widgetBettyID, "Betty Rubble", map[string]string{"CostumeColor": "blue", "Role": "neighbor"}, "Betty", nil, 2, "blue", "neighbor"}
	widgetFred   = EnumWidgetItem{widgetFredID, "Fred Flintstone", map[string]string{"CostumeColor": "orange", "Role": "father"}, "Fred", nil, 3, "orange", "father"}
	widgetWilma  = EnumWidgetItem{widgetWilmaID, "Wilma Flintstone", map[string]string{"CostumeColor": "white", "Role": "mother"}, "Wilma", nil, 4, "white", "mother"}
	widgetElroy  = EnumWidgetItem{widgetElroyID, "Elroy Jetson", map[string]string{"CostumeColor": "brown", "Role": "son"}, "Elroy", &widgetGeorge, 5, "brown", "son"}
	widgetGeorge = EnumWidgetItem{widgetGeorgeID, "George Jetson", map[string]string{"CostumeColor": "white", "Role": "father"}, "George", nil, 6, "white", "father"}
	widgetJane   = EnumWidgetItem{widgetJaneID, "Jane Jetson", map[string]string{"CostumeColor": "green", "Role": "mother"}, "Jane", nil, 7, "green", "mother"}
	widgetJudy   = EnumWidgetItem{widgetJudyID, "Judy Jetson", map[string]string{"CostumeColor": "red", "Role": "daughter"}, "Judy", &widgetGeorge, 8, "red", "daughter"}
	widgetBart   = EnumWidgetItem{widgetBartID, "Bart Simpson", map[string]string{"CostumeColor": "red", "Role": "son"}, "Bart", nil, 9, "red", "son"}
	widgetHomer  = EnumWidgetItem{widgetHomerID, "Homer Simpson", map[string]string{"CostumeColor": "white", "Role": "father"}, "Homer", nil, 10, "white", "father"}
	widgetLisa   = EnumWidgetItem{widgetLisaID, "Lisa Simpson", map[string]string{"CostumeColor": "red", "Role": "daughter"}, "Lisa", nil, 11, "red", "daughter"}
	widgetMarge  = EnumWidgetItem{widgetMargeID, "Marge Simpson", map[string]string{"CostumeColor": "green", "Role": "mother"}, "Marge", nil, 12, "green", "mother"}
)

// EnumWidget is a collection of Widget items
type EnumWidget struct {
	Description string
	Items       []*EnumWidgetItem
	Name        string

	Barney *EnumWidgetItem
	Betty  *EnumWidgetItem
	Fred   *EnumWidgetItem
	Wilma  *EnumWidgetItem
	Elroy  *EnumWidgetItem
	George *EnumWidgetItem
	Jane   *EnumWidgetItem
	Judy   *EnumWidgetItem
	Bart   *EnumWidgetItem
	Homer  *EnumWidgetItem
	Lisa   *EnumWidgetItem
	Marge  *EnumWidgetItem

	itemDict map[string]*EnumWidgetItem
}

// Widget is a public singleton instance of EnumWidget
// representing widgets, which are awesome
var Widget = &EnumWidget{
	Description: "widgets, which are awesome",
	Items: []*EnumWidgetItem{
		&widgetBarney,
		&widgetBetty,
		&widgetFred,
		&widgetWilma,
		&widgetElroy,
		&widgetGeorge,
		&widgetJane,
		&widgetJudy,
		&widgetBart,
		&widgetHomer,
		&widgetLisa,
		&widgetMarge,
	},
	Name:   "EnumWidget",
	Barney: &widgetBarney,
	Betty:  &widgetBetty,
	Fred:   &widgetFred,
	Wilma:  &widgetWilma,
	Elroy:  &widgetElroy,
	George: &widgetGeorge,
	Jane:   &widgetJane,
	Judy:   &widgetJudy,
	Bart:   &widgetBart,
	Homer:  &widgetHomer,
	Lisa:   &widgetLisa,
	Marge:  &widgetMarge,

	itemDict: map[string]*EnumWidgetItem{
		strings.ToLower(string(widgetBarneyID)): &widgetBarney,
		strings.ToLower(string(widgetBettyID)):  &widgetBetty,
		strings.ToLower(string(widgetFredID)):   &widgetFred,
		strings.ToLower(string(widgetWilmaID)):  &widgetWilma,
		strings.ToLower(string(widgetElroyID)):  &widgetElroy,
		strings.ToLower(string(widgetGeorgeID)): &widgetGeorge,
		strings.ToLower(string(widgetJaneID)):   &widgetJane,
		strings.ToLower(string(widgetJudyID)):   &widgetJudy,
		strings.ToLower(string(widgetBartID)):   &widgetBart,
		strings.ToLower(string(widgetHomerID)):  &widgetHomer,
		strings.ToLower(string(widgetLisaID)):   &widgetLisa,
		strings.ToLower(string(widgetMargeID)):  &widgetMarge,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumWidget) ByID(id WidgetIdentifier) *EnumWidgetItem {
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
func (e *EnumWidget) ByIDString(idx string) *EnumWidgetItem {
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
func (e *EnumWidget) ByIndex(idx int) *EnumWidgetItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedWidgetID is a struct that is designed to replace a *WidgetID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *WidgetID it contains while being a better JSON citizen.
type ValidatedWidgetID struct {
	// id will point to a valid WidgetID, if possible
	// If id is nil, then ValidatedWidgetID.Valid() will return false.
	id *WidgetID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedWidgetID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedWidgetID
func (vi *ValidatedWidgetID) Clone() *ValidatedWidgetID {
	if vi == nil {
		return nil
	}

	var cid *WidgetID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedWidgetID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedWidgetIds represent the same Widget
func (vi *ValidatedWidgetID) Equals(vj *ValidatedWidgetID) bool {
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

func (vi *ValidatedWidgetID) Parent() *ValidatedWidgetID {
	if vi == nil || vi.id == nil {
		pid := vi.id.Parent()
		if pid != nil {
			return pid.ValidatedID()
		}
	}

	return nil
}

// Valid returns true if and only if the ValidatedWidgetID corresponds to a recognized Widget
func (vi *ValidatedWidgetID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedWidgetID) ID() *WidgetID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedWidgetID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedWidgetID) ValidatedID() *ValidatedWidgetID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedWidgetID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedWidgetID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedWidgetID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedWidgetID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedWidgetID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := WidgetID(capString)
	item := Widget.ByID(&id)
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

func (vi ValidatedWidgetID) String() string {
	return vi.ToIDString()
}

type WidgetIdentifier interface {
	ID() *WidgetID
	Valid() bool
}
