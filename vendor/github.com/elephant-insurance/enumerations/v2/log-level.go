package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// LogLevelID uniquely identifies a particular LogLevel
type LogLevelID string

// Clone creates a safe, independent copy of a LogLevelID
func (i *LogLevelID) Clone() *LogLevelID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two LogLevelIds are equivalent
func (i *LogLevelID) Equals(j *LogLevelID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *LogLevelID that is either valid or nil
func (i *LogLevelID) ID() *LogLevelID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *LogLevelID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the LogLevelID corresponds to a recognized LogLevel
func (i *LogLevelID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return LogLevel.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *LogLevelID) ValidatedID() *ValidatedLogLevelID {
	if i != nil {
		return &ValidatedLogLevelID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *LogLevelID) MarshalJSON() ([]byte, error) {
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

func (i *LogLevelID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := LogLevelID(dataString)
	item := LogLevel.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	logLevelPanicID LogLevelID = "PANIC"
	logLevelFatalID LogLevelID = "FATAL"
	logLevelErrorID LogLevelID = "ERROR"
	logLevelWarnID  LogLevelID = "WARN"
	logLevelInfoID  LogLevelID = "INFO"
	logLevelDebugID LogLevelID = "DEBUG"
	logLevelTraceID LogLevelID = "TRACE"
)

// EnumLogLevelItem describes an entry in an enumeration of LogLevel
type EnumLogLevelItem struct {
	ID        LogLevelID        `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	logLevelPanic = EnumLogLevelItem{logLevelPanicID, "Panic", nil, "Panic", 0}
	logLevelFatal = EnumLogLevelItem{logLevelFatalID, "Fatal", nil, "Fatal", 1}
	logLevelError = EnumLogLevelItem{logLevelErrorID, "Error", nil, "Error", 2}
	logLevelWarn  = EnumLogLevelItem{logLevelWarnID, "Warn", nil, "Warn", 3}
	logLevelInfo  = EnumLogLevelItem{logLevelInfoID, "Info", nil, "Info", 4}
	logLevelDebug = EnumLogLevelItem{logLevelDebugID, "Debug", nil, "Debug", 5}
	logLevelTrace = EnumLogLevelItem{logLevelTraceID, "Trace", nil, "Trace", 6}
)

// EnumLogLevel is a collection of LogLevel items
type EnumLogLevel struct {
	Description string
	Items       []*EnumLogLevelItem
	Name        string

	Panic *EnumLogLevelItem
	Fatal *EnumLogLevelItem
	Error *EnumLogLevelItem
	Warn  *EnumLogLevelItem
	Info  *EnumLogLevelItem
	Debug *EnumLogLevelItem
	Trace *EnumLogLevelItem

	itemDict map[string]*EnumLogLevelItem
}

// LogLevel is a public singleton instance of EnumLogLevel
// representing log urgency levels
var LogLevel = &EnumLogLevel{
	Description: "log urgency levels",
	Items: []*EnumLogLevelItem{
		&logLevelPanic,
		&logLevelFatal,
		&logLevelError,
		&logLevelWarn,
		&logLevelInfo,
		&logLevelDebug,
		&logLevelTrace,
	},
	Name:  "EnumLogLevel",
	Panic: &logLevelPanic,
	Fatal: &logLevelFatal,
	Error: &logLevelError,
	Warn:  &logLevelWarn,
	Info:  &logLevelInfo,
	Debug: &logLevelDebug,
	Trace: &logLevelTrace,

	itemDict: map[string]*EnumLogLevelItem{
		strings.ToLower(string(logLevelPanicID)): &logLevelPanic,
		strings.ToLower(string(logLevelFatalID)): &logLevelFatal,
		strings.ToLower(string(logLevelErrorID)): &logLevelError,
		strings.ToLower(string(logLevelWarnID)):  &logLevelWarn,
		strings.ToLower(string(logLevelInfoID)):  &logLevelInfo,
		strings.ToLower(string(logLevelDebugID)): &logLevelDebug,
		strings.ToLower(string(logLevelTraceID)): &logLevelTrace,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumLogLevel) ByID(id LogLevelIdentifier) *EnumLogLevelItem {
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
func (e *EnumLogLevel) ByIDString(idx string) *EnumLogLevelItem {
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
func (e *EnumLogLevel) ByIndex(idx int) *EnumLogLevelItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedLogLevelID is a struct that is designed to replace a *LogLevelID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *LogLevelID it contains while being a better JSON citizen.
type ValidatedLogLevelID struct {
	// id will point to a valid LogLevelID, if possible
	// If id is nil, then ValidatedLogLevelID.Valid() will return false.
	id *LogLevelID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedLogLevelID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedLogLevelID
func (vi *ValidatedLogLevelID) Clone() *ValidatedLogLevelID {
	if vi == nil {
		return nil
	}

	var cid *LogLevelID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedLogLevelID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedLogLevelIds represent the same LogLevel
func (vi *ValidatedLogLevelID) Equals(vj *ValidatedLogLevelID) bool {
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

// Valid returns true if and only if the ValidatedLogLevelID corresponds to a recognized LogLevel
func (vi *ValidatedLogLevelID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedLogLevelID) ID() *LogLevelID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedLogLevelID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedLogLevelID) ValidatedID() *ValidatedLogLevelID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedLogLevelID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedLogLevelID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedLogLevelID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedLogLevelID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedLogLevelID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := LogLevelID(capString)
	item := LogLevel.ByID(&id)
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

func (vi ValidatedLogLevelID) String() string {
	return vi.ToIDString()
}

type LogLevelIdentifier interface {
	ID() *LogLevelID
	Valid() bool
}
