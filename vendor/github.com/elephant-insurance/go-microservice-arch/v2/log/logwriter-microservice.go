package log

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
	"github.com/elephant-insurance/go-microservice-arch/v2/timing"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
)

// microserviceLogwriter is the fundamental type for microservice logging
type microserviceLogwriter struct {
	// active controls whether the logwriter is switched off by a failed condition or not.
	// A logwriter that is not active will not log anything, nor will it update its fields or entry.
	// An inactive logwriter will be switched back on as soon as it executes any logging command (e.g., Info()).
	active bool
	// context, if added with SetContext, can contain log values
	context msrqc.Context
	// elsable remembers whether this logwriter recently failed a conditional.
	elsable bool
	// entry contains all of the fields that we want to send to the relays (e.g. Azure)
	entry *MicroserviceLogEntry
	// event keeps track of an event that we want to log
	events []*MSEventLogEntry
	// fieldMap contains all of the fields that we have added using WithConsoleField
	// these fields are sent to the immediateHandlers, but not to the relays
	fieldMap map[string]interface{}
	// log is a reference back to the Log that created this Logwriter
	// messages are sent back to the Log for handling
	log *MicroserviceLog
}

// Logging methods
func (lw *microserviceLogwriter) Trace(msg string, args ...interface{}) MicroserviceLogger {
	// log will check this, as well, but checking here saves the work of copying args
	if len(args) > 0 && lw.log.level.ID.Equals(&enum.LogLevel.Trace.ID) {
		lw.WithConsoleFields(args...)
	}

	lw.dispatch(enum.LogLevel.Trace, &msg)

	return lw
}

func (lw *microserviceLogwriter) Debug(msg string, args ...interface{}) MicroserviceLogger {
	if len(args) > 0 && (lw.log.level.ID.Equals(&enum.LogLevel.Debug.ID) || lw.log.level.ID.Equals(&enum.LogLevel.Trace.ID)) {
		lw.WithConsoleFields(args...)
	}

	lw.dispatch(enum.LogLevel.Debug, &msg)

	return lw
}

func (lw *microserviceLogwriter) Info(msg string, args ...interface{}) MicroserviceLogger {
	if len(args) > 0 && !lw.log.level.IsMoreUrgentThan(enum.LogLevel.Info) {
		lw.WithConsoleFields(args...)
	}

	lw.dispatch(enum.LogLevel.Info, &msg)

	return lw
}

func (lw *microserviceLogwriter) Warn(msg string, args ...interface{}) MicroserviceLogger {
	if len(args) > 0 && !lw.log.level.IsMoreUrgentThan(enum.LogLevel.Warn) {
		lw.WithConsoleFields(args...)
	}

	lw.dispatch(enum.LogLevel.Warn, &msg)

	return lw
}

func (lw *microserviceLogwriter) Error(msg string, args ...interface{}) MicroserviceLogger {
	if len(args) > 0 && !lw.log.level.IsMoreUrgentThan(enum.LogLevel.Error) {
		lw.WithConsoleFields(args...)
	}

	lw.dispatch(enum.LogLevel.Error, &msg)

	return lw
}

func (lw *microserviceLogwriter) Fatal(msg string, args ...interface{}) MicroserviceLogger {
	if len(args) > 0 {
		lw.WithConsoleFields(args...)
	}

	lw.dispatch(enum.LogLevel.Fatal, &msg)

	return lw
}

func (lw *microserviceLogwriter) Panic(msg string, args ...interface{}) MicroserviceLogger {
	if len(args) > 0 {
		lw.WithConsoleFields(args...)
	}

	lw.dispatch(enum.LogLevel.Panic, &msg)

	return lw
}

// dispatch calls the logwriter's log.dispatch. All logging functions pass through here.
func (lw *microserviceLogwriter) dispatch(level *enum.EnumLogLevelItem, msg *string) {
	if lw.active {
		lw.log.dispatch(lw, level, msg)
		lw.clearEphemeralFields()
	}

	// this makes conditional logging work: a failed condition only turns me off until I try to log something
	lw.active = true
}

// emit returns a MicroserviceLogEntry containing all of the fields from this logWriter that we want to relay
func (lw *microserviceLogwriter) emit() *MicroserviceLogEntry {
	// TODO: copy matching fields from map for backward compatibility with WithConsoleFields -- ??? Maybe not.
	// Until this is done, a microservice log field set using WithConsoleField will not forward its value to the relay

	rtn := lw.entry.Clone()

	// cycle through context and update entry with matching fields
	// the key to the context value must match the JSON tag name for the field
	if lw.context != nil {
		ffc := fieldsFromContext(lw.context)

		for k, v := range ffc {
			if propName, ok := microserviceLogFields[k]; ok && propName != `` {
				rtn.setPropertyByName(propName, v)
			}
		}
	}

	return rtn
}

// makeMap consolidates the data held in the logwriter and returns it as a map.
// It also returns a slice of sorted keys for the map.
// We have three sources of data: values stored in the Context have the highest priority.
// This is because a value set in the Context should be relevant throughout the lifetime of the Context;
// it should not be overwritten unless the developer explicitly overwrites it in the Context.
// "Set" fields (properties on lw.entry) have middle priority,
// and "With" fields (stored in lw.fieldMap) have the lowest priority.
// Priority only matters when a key value in the context map or fieldMap matches a key in the other map,
// or matches the JSON fieldname of a property on lw.entry.
// Example: set "code"="A" in the context, set lw.WithConsoleField("code", "B"), and call lw.SetCode("C")
// Regardless of the order called, "A" wins.
func (lw *microserviceLogwriter) makeMap() (dataMap map[string]interface{}, sortedKeys []string) {
	// I don't like abusing JSON this way, but with all but the essential fields set to omitempty,
	// this is the most efficient way to go from a big struct to a much smaller map[string]interface{}
	var objMap map[string]json.RawMessage
	entryBytes, err := json.Marshal(*lw.entry)
	if err != nil {
		debugLog(`microserviceLogwriter.makeMap error attempting to Marshal MicroserviceLogEntry: ` + err.Error())
		return
	}

	err = json.Unmarshal(entryBytes, &objMap)
	if err != nil {
		debugLog(`microserviceLogwriter.makeMap error attempting to Unmarshal MicroserviceLogEntry: ` + err.Error())
		return
	}

	var contextMap map[string]interface{}
	if lw.context != nil {
		contextMap = fieldsFromContext(lw.context)
	} else {
		contextMap = map[string]interface{}{}
	}

	// figure out how big our target is so we don't have to resize it
	totalLen := len(lw.fieldMap) + len(objMap) + len(contextMap)

	dataMap = make(map[string]interface{}, totalLen)

	// "With" fields from fieldMap have lowest priority
	for k, v := range lw.fieldMap {
		dataMap[k] = v
	}

	// debugLog(`AFTER fieldMap`)
	// debugSpew(dataMap)

	// next come the "Set" fields in lw.entry
	for k, v := range objMap {
		// if shouldSkip, ok := microserviceLogSkipMapEntries[k]; !ok || !shouldSkip {
		dataMap[k] = v
		// }
	}

	// debugLog(`AFTER objMap`)
	// debugSpew(dataMap)

	// fields from context have highest priority, so we add them last
	for k, v := range contextMap {
		dataMap[k] = v
	}

	keycount := len(dataMap)
	if keycount > 0 {
		sortedKeys = make([]string, 0, keycount)
		for k := range dataMap {
			sortedKeys = append(sortedKeys, k)
		}

		sort.Strings(sortedKeys)
	}

	debugSpew(dataMap)
	return
}

// clearEphemeralFields clears the values of any ephemaral fields immediately after they are written.
// This ensures that fields with very specific application (like Error and HTTPStatusCode) don't stick around
// and make confusing messages later.
func (lw *microserviceLogwriter) clearEphemeralFields() {
	if lw == nil {
		return
	}

	if lw.entry == nil {
		return
	}

	lw.entry.CommonErrorCode = nil
	lw.entry.Message = nil
	lw.entry.Level = nil
	lw.entry.Time = nil
	lw.entry.ElapsedMicroseconds = nil
	lw.entry.Error = nil
	lw.entry.HTTPMethod = nil
	lw.entry.HTTPStatus = nil
	lw.entry.Stack = nil
	lw.entry.URL = nil

	lw.fieldMap = map[string]interface{}{}
}

// Microservice Log methods set properties on the contained entry

// SetAccountID
func (lw *microserviceLogwriter) SetAccountID(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.AccountID = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSAccountID, &v, true)
		}
	}

	return lw
}

// SetBusinessType sets the type of business for the transaction we are logging
func (lw *microserviceLogwriter) SetBusinessType(v *enum.BusinessTypeID, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.BusinessType = v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSBusinessType, v, true)
		}
	}

	return lw
}

// SetChannel
func (lw *microserviceLogwriter) SetChannel(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.Channel = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSChannel, &v, true)
		}
	}

	return lw
}

// SetCode can be used for coverage code, error code, whatever
func (lw *microserviceLogwriter) SetCode(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.Code = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSCode, &v, true)
		}
	}

	return lw
}

func (lw *microserviceLogwriter) WithCommonErrorCode(v *enum.CommonErrorCodeID) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.CommonErrorCode = v
		msg := v.ToString()
		lw.entry.Error = &msg
	}

	return lw
}

// SetCount can be used to log counts of things
func (lw *microserviceLogwriter) SetCount(v int, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.Count = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSCount, &v, true)
		}
	}

	return lw
}

// SetDate sets the Date field of the microservice log.
// For simplicity we store and render the date as a string in YYYY-MM-DD format.
// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
func (lw *microserviceLogwriter) SetDate(v uf.Datable, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active && v != nil {
		y, mm, d := v.Date()
		m := int(mm)
		dateString := fmt.Sprintf(`%04d-%02d-%02d`, y, m, d)
		lw.entry.Date = &dateString

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSDate, &dateString, true)
		}
	}

	return lw
}

// SetDetail can be used to set any arbitrary string value, up to 1000 chars.
func (lw *microserviceLogwriter) SetDetail(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		if len(v) > 1000 {
			v = string(v[:1000])
		}

		lw.entry.Detail = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSDetail, &v, true)
		}
	}

	return lw
}

// SetDriver can be used to set a driver id or a driver name
func (lw *microserviceLogwriter) SetDriver(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.Driver = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSDriver, &v, true)
		}
	}

	return lw
}

// WithElapsedMicroseconds should be used to record timings.
// Normalizing to microseconds (not seconds, nanoseconds, or milliseconds) gives us a useful number in just about every context,
// and ensures that numbers from different sources are directly comparable.
func (lw *microserviceLogwriter) WithElapsedMicroseconds(v int64) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.ElapsedMicroseconds = &v
	}

	return lw
}

// SetFunction should be used to log the name of the running function
func (lw *microserviceLogwriter) SetFunction(v string) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.Function = &v
	}

	return lw
}

// WithHTTPMethod should be used to record the HTTP method for a request
func (lw *microserviceLogwriter) WithHTTPMethod(v string) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.HTTPMethod = &v
	}

	return lw
}

// WithHTTPStatus can be used to log HTTP status codes
func (lw *microserviceLogwriter) WithHTTPStatus(v int) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.HTTPStatus = &v
	}

	return lw
}

// SetID
func (lw *microserviceLogwriter) SetID(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.ID = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSID, &v, true)
		}
	}

	return lw
}

// SetIPAddress should be used to record a Guidewire Job Number
func (lw *microserviceLogwriter) SetIPAddress(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.IPAddress = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSIPAddress, &v, true)
		}
	}

	return lw
}

// SetJobID should be used to record a Guidewire Job Number
func (lw *microserviceLogwriter) SetJobID(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.JobID = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSJobID, &v, true)
		}
	}

	return lw
}

// SetName
func (lw *microserviceLogwriter) SetName(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.Name = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSName, &v, true)
		}
	}

	return lw
}

// SetNumber can be used to log counts of things
func (lw *microserviceLogwriter) SetNumber(v int, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.Number = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSNumber, &v, true)
		}
	}

	return lw
}

// SetPolicyID
func (lw *microserviceLogwriter) SetPolicyID(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.PolicyID = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSPolicyID, &v, true)
		}
	}

	return lw
}

// SetPublicID
func (lw *microserviceLogwriter) SetPublicID(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.PublicID = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSPublicID, &v, true)
		}
	}

	return lw
}

// SetQuoteID
func (lw *microserviceLogwriter) SetQuoteID(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.QuoteID = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSQuoteID, &v, true)
		}
	}

	return lw
}

// SetQuoteNumber
func (lw *microserviceLogwriter) SetQuoteNumber(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.QuoteNumber = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSQuoteNumber, &v, true)
		}
	}

	return lw
}

// WithStack writes a stack trace to the log entry
func (lw *microserviceLogwriter) WithStack() MicroserviceLogger {
	stack := stack(2) // skip two frames to get to the immediate calling func
	lw.entry.Stack = &stack
	return lw
}

// SetState should be used to log a state of the USA
func (lw *microserviceLogwriter) SetState(v *enum.StateID, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.State = v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSState, v, true)
		}
	}

	return lw
}

// WithURL
func (lw *microserviceLogwriter) WithURL(v string) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.URL = &v
	}

	return lw
}

// SetTransactionBrand sets the TransactionBrand property of the logger's entry.
// This method will not overwrite a value that has already been set, nor will it write an empty value.
func (lw *microserviceLogwriter) SetTransactionBrand(v *enum.BrandID) MicroserviceLogger {
	if lw == nil || !lw.active || v == nil {
		return lw
	}

	if lw.context != nil {
		msrqc.SetTransactionBrand(lw.context, v)
	} else {
		lw.entry.TransactionBrand = v
	}

	return lw
}

// SetTransactionDomain sets the TransactionDomain property of the logger's entry.
// This method will not overwrite a value that has already been set, nor will it write an empty value.
func (lw *microserviceLogwriter) SetTransactionDomain(v *enum.AccountDomainID) MicroserviceLogger {
	if lw == nil || !lw.active || v == nil {
		return lw
	}

	if lw.context != nil {
		msrqc.SetTransactionDomain(lw.context, v)
	} else {
		lw.entry.TransactionDomain = v
	}

	return lw
}

// SetTransactionID sets the TransactionID property of the logger's entry.
// This method will not overwrite a value that has already been set, nor will it write an empty value.
func (lw *microserviceLogwriter) SetTransactionID(v string) MicroserviceLogger {
	if lw == nil || !lw.active || v == `` {
		return lw
	}

	if lw.context != nil {
		msrqc.SetTransactionID(lw.context, v)
	} else {
		lw.entry.TransactionID = &v
	}

	return lw
}

func (lw *microserviceLogwriter) GetTransactionID() string {
	var rtn string
	if lw != nil {
		rtn = msrqc.GetTransactionID(lw.context)
		if rtn == `` {
			if lw.entry != nil && lw.entry.TransactionID != nil {
				rtn = *lw.entry.TransactionID
			}
		}
	}

	return rtn
}

// SetTransactionIntegrator sets the TransactionIntegrator (integration partner) property of the logger's entry.
// This method will not overwrite a value that has already been set, nor will it write an empty value.
func (lw *microserviceLogwriter) SetTransactionIntegrator(v *enum.IntegrationPartnerID) MicroserviceLogger {
	if lw == nil || !lw.active || v == nil {
		return lw
	}

	if lw.context != nil {
		msrqc.SetTransactionIntegrator(lw.context, v)
	} else {
		lw.entry.TransactionIntegrator = v
	}

	return lw
}

// SetTransactionSource sets the TransactionSource (source of business) property of the logger's entry.
// This method will not overwrite a value that has already been set, nor will it write an empty value.
func (lw *microserviceLogwriter) SetTransactionSource(v *enum.SourceOfBusinessID) MicroserviceLogger {
	if lw == nil || !lw.active || v == nil {
		return lw
	}

	if lw.context != nil {
		msrqc.SetTransactionSource(lw.context, v)
	} else {
		lw.entry.TransactionSource = v
	}

	return lw
}

// SetTransactionType sets the TransactionType property of the logger's entry.
// This method will not overwrite a value that has already been set, nor will it write an empty value.
func (lw *microserviceLogwriter) SetTransactionType(v string) MicroserviceLogger {
	if lw == nil || !lw.active || v == `` {
		return lw
	}

	if lw.context != nil {
		msrqc.SetTransactionType(lw.context, v)
	} else {
		lw.entry.TransactionType = &v
	}

	return lw
}

// SetVehicle should be used to identify a vehicle, usually bu ID
func (lw *microserviceLogwriter) SetVehicle(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.Vehicle = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSVehicle, &v, true)
		}
	}

	return lw
}

// SetVIN
func (lw *microserviceLogwriter) SetVIN(v string, addToContext ...interface{}) MicroserviceLogger {
	if lw != nil && lw.active {
		lw.entry.VIN = &v

		if ifSet(addToContext) && lw.context != nil {
			setContextValue(lw.context, fieldNameMSVIN, &v, true)
		}
	}

	return lw
}

/*
	Conditional logging methods

	The following methods are intended to be used in fluent chains of statements, like:

	lw.If(x == 0).Warn(`x should never be zero`)

	When a conditional logging statement evaluates to false, it cuases the output of the
	logwriter's next logging statement to be suppressed.

	Because of the way the logic works, it is permissible to string together multiple
	conditional logging statements. However, since any true result resets the logwriter,
	only the last conditional in the chain will actually matter.

*/

// If evaluates condition, and sets its active flag to equal condition. elsable is set to the opposite of condition.
func (lw *microserviceLogwriter) If(condition bool) MicroserviceLogger {
	lw.active = condition
	lw.elsable = !condition
	return lw
}

// IfError is equivalent to If(err != nil).SetError(err)
func (lw *microserviceLogwriter) IfError(err error) MicroserviceLogger {
	if err != nil {
		lw.WithError(err)
	}

	return lw.If(err != nil)
}

// IfEval works just like IF, except it takes a func() bool as its argument
func (lw *microserviceLogwriter) IfEval(evaluator Evaluator) MicroserviceLogger {
	return lw.If(evaluator())
}

// Else
func (lw *microserviceLogwriter) Else() MicroserviceLogger {
	lw.active = lw.elsable
	lw.elsable = false
	return lw
}

// ElseIf
func (lw *microserviceLogwriter) ElseIf(condition bool) MicroserviceLogger {
	if lw.elsable {
		return lw.If(condition)
	}
	lw.active = false
	lw.elsable = false
	return lw
}

func (lw *microserviceLogwriter) SetContext(c context.Context) MicroserviceLogger {
	lw.context = msrqc.New(c)

	return lw
}

func (lw *microserviceLogwriter) SetField(key string, value interface{}) MicroserviceLogger {
	if lw.context != nil {
		setContextValue(lw.context, key, value, false)
		return lw
	}

	var setValue interface{}

	if reflect.ValueOf(value).Kind() != reflect.Ptr {
		setValue = &value
	} else {
		setValue = value
	}

	if _, ok := microserviceLogFields[key]; ok {
		lw.entry.setPropertyByJSONName(key, setValue)
		return lw
	}

	return lw.WithConsoleField(key, value)
}

// WithError
func (lw *microserviceLogwriter) WithError(err error) MicroserviceLogger {
	if err != nil {
		errmsg := err.Error()
		lw.entry.Error = &errmsg
	}

	return lw
}

func (lw *microserviceLogwriter) WithConsoleField(key string, value interface{}, addToContext ...interface{}) MicroserviceLogger {
	if key != "" {
		lw.fieldMap[key] = value
	}

	return lw
}

// WithConsoleFields is a convenience method for setting multiple "with" fields at once.
// This function expects every odd-numbered argument to be a string.
// It is equivalent to WithConsoleField(`arg1`, arg2).WithConsoleField(`arg3`, arg4)...
func (lw *microserviceLogwriter) WithConsoleFields(args ...interface{}) MicroserviceLogger {
	if len(args) > 0 {
		// we expect each odd-numbered arg to be a string field name
		// and each even-numbered arg to be a parameter value
		for i := 0; i < len(args)-1; i += 2 {
			argName := args[i]
			argVal := args[i+1]
			if argName == nil || argVal == nil {
				continue
			}
			nameString, ok := argName.(string)
			if !ok || nameString == "" {
				continue
			}
			lw.fieldMap[nameString] = argVal
		}
	}

	return lw
}

// ClearEvents removes any events from the logwriter and returns them to the caller.
// REMEMBER: log writers are NOT thread safe!
func (lw *microserviceLogwriter) ClearEvents() []*MSEventLogEntry {
	rtn := []*MSEventLogEntry{}

	if lw != nil {
		foo := lw.events
		lw.events = rtn
		rtn = foo
	}

	return rtn
}

const (
	unsetHostname     = `(not set)`
	unsetInstanceName = `(not set)`
	unsetTXType       = `(not set)`
)

// WithEvent adds an Event to the logger
func (lw *microserviceLogwriter) WithEvent(event *uf.Event) MicroserviceLogger {
	if lw == nil || lw.entry == nil || event == nil {
		return lw
	}

	hn := unsetHostname
	if lw.entry.HostName != nil {
		hn = *lw.entry.HostName
	}

	in := unsetInstanceName
	if lw.entry.InstanceName != nil {
		in = *lw.entry.InstanceName
	}

	tid := lw.GetTransactionID()

	// create entry from event
	evte := &MSEventLogEntry{
		Description:   event.Description,
		ID:            event.ID,
		Severity:      event.Severity,
		TargetService: event.TargetService,

		CommonErrorCode:      lw.entry.CommonErrorCode,
		DurationMicroseconds: lw.entry.ElapsedMicroseconds,
		Error:                lw.entry.Error,
		HostName:             hn,
		InstanceName:         in,
		Time:                 time.Now(),
		TransactionID:        &tid,
		TransactionType:      lw.entry.TransactionType,
	}

	lw.events = append(lw.events, evte)

	return lw
}

// ServiceResponseSuccess is a one-stop logger for service responses
// It will return true IF the follwing are both TRUE:
// 1. err is nil.
// 2. resp is not nil.
// If either of these conditions fails, we log both an event and a log message at the Error level.
// Use this version when the service may return an error status code (>300) under normal conditions (not a "real" error that we need to deal with).
func (lw *microserviceLogwriter) ServiceResponseSuccess(resp *http.Response, err error, respTime *timing.Timing, target *enum.ServiceID, description string) bool {
	if !lw.ServiceResponseWithoutError(resp, err, respTime, target, description) {
		return false
	}

	if resp == nil {
		msg := `HTTP service call returned nil response`
		if description == `` {
			description = msg
		}

		evt := uf.EventFactory.New(&enum.Event.ServiceRequestNilResponse.ID, target, description)
		lw.WithError(err).WithEvent(evt).Error(msg)
		return false
	}

	return true
}

// ServiceResponseStrictSuccess works exactly like ServiceResponseSuccess
// except that it also checks the response.StatusCode.
// If the status code indicates an error (>300), we log an error and return false.
// Use this version when only a 2xx response code is acceptable from the service.
func (lw *microserviceLogwriter) ServiceResponseStrictSuccess(resp *http.Response, err error, respTime *timing.Timing, target *enum.ServiceID, description string) bool {
	if !lw.ServiceResponseSuccess(resp, err, respTime, target, description) {
		return false
	}

	if resp.StatusCode >= http.StatusMultipleChoices {
		msg := `HTTP service call reurned bad status`
		if description == `` {
			description = msg
		}

		evt := uf.EventFactory.New(&enum.Event.ServiceRequestError.ID, target, description)
		lw.WithHTTPStatus(resp.StatusCode).WithError(err).WithEvent(evt).Error(msg)
		return false
	}

	return true
}

// ServiceResponseWithoutError is a one-stop logger for service responses
// It will return true IF the follwing is TRUE:
// 1. err is nil.
// If err is not nil, we log both an event and a log message at the Error level.
// Use this version when the service may return a nil response or an error status code (>300) under normal conditions (not a "real" error that we need to deal with).
func (lw *microserviceLogwriter) ServiceResponseWithoutError(resp *http.Response, err error, respTime *timing.Timing, target *enum.ServiceID, description string) bool {
	if respTime != nil {
		lw.WithElapsedMicroseconds(respTime.DurationMicroseconds)
	}

	if err != nil {
		msg := `HTTP service call returned error`
		if description == `` {
			description = msg
		}

		evt := uf.EventFactory.New(&enum.Event.ServiceRequestError.ID, target, description)
		lw.WithError(err).WithEvent(evt).Error(msg)
		return false
	}

	return true
}

// ifSet is a convenience method for determining whether a Set method was called with a true addToContext parameter.
func ifSet(addToContext ...interface{}) bool {
	if len(addToContext) > 0 && addToContext[0] != nil {
		atcifc, ok := addToContext[0].([]interface{})
		if ok && len(atcifc) > 0 {
			if atcbool, ok2 := atcifc[0].(bool); ok2 && atcbool {
				return true
			}
		}
	}

	return false
}
