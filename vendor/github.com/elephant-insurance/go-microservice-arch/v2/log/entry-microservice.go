package log

import (
	"reflect"
	"runtime/debug"
	"strings"
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
)

// MicroserviceLogEntry describes the schema of the Mocroservice Log table in Azure Log Analytics
// You can use this type directly, if you want
// But normally it is used as a "filter" to remove log fields that should not be pushed to Azure
// Every field of this struct will appear as an appropriately-typed column in Azure
// Every log field that does NOT appear here will be removed before sending to Azure
// PLEASE NOTE: the most important element here is the value of the `json:"NAME"` tag!
// The NAME in the JSON tag will be the NAME of the field in your log message
// and also the NAME of the corresponding column in Azure
// Omitting the JSON NAME tag will prevent the field from going to Azure
// Additional JSON tags (like omitempty) may be added, but will have no real effect.
type MicroserviceLogEntry struct {
	// "Set" fields persist after logging
	AccountID             *string                    `json:"account,omitempty"`
	BusinessType          *enum.BusinessTypeID       `json:"BusinessType,omitempty"`
	Channel               *string                    `json:"channel,omitempty"`
	Code                  *string                    `json:"code,omitempty"`
	Coverage              *string                    `json:"coverage,omitempty"`
	Count                 *int                       `json:"count,omitempty"`
	Date                  *string                    `json:"date,omitempty"`
	Detail                *string                    `json:"detail,omitempty"`
	Driver                *string                    `json:"driver,omitempty"`
	Environment           *enum.ServiceEnvironmentID `json:"env"`
	HostName              *string                    `json:"host"`
	ID                    *string                    `json:"id,omitempty"`
	InstanceName          *string                    `json:"InstanceName"`
	IPAddress             *string                    `json:"ip-address,omitempty"`
	JobID                 *string                    `json:"job,omitempty"`
	Level                 *enum.LogLevelID           `json:"level"`
	Name                  *string                    `json:"Name,omitempty"`
	Number                *int                       `json:"number,omitempty"`
	PolicyID              *string                    `json:"policy,omitempty"`
	PublicID              *string                    `json:"public-id,omitempty"`
	QuoteID               *string                    `json:"quote-id,omitempty"`
	QuoteNumber           *string                    `json:"quote-number,omitempty"`
	ServiceID             *enum.ServiceID            `json:"app"`
	State                 *enum.StateID              `json:"state,omitempty"`
	TransactionBrand      *enum.BrandID              `json:"txbrand,omitempty"`
	TransactionDomain     *enum.AccountDomainID      `json:"txdomain,omitempty"`
	TransactionID         *string                    `json:"txid,omitempty"`
	TransactionInstance   *string                    `json:"txinstance,omitempty"`
	TransactionIntegrator *enum.IntegrationPartnerID `json:"txintegrator,omitempty"`
	TransactionIPAddress  *string                    `json:"txip,omitempty"`
	TransactionSource     *enum.SourceOfBusinessID   `json:"txsource,omitempty"`
	TransactionType       *string                    `json:"txtype,omitempty"`
	Vehicle               *string                    `json:"vehicle,omitempty"`
	VIN                   *string                    `json:"vin,omitempty"`

	// "With" fields get logged once, then clear
	CommonErrorCode     *enum.CommonErrorCodeID `json:"error-code,omitempty"`
	ElapsedMicroseconds *int64                  `json:"microseconds,omitempty"`
	Error               *string                 `json:"err,omitempty"`
	Function            *string                 `json:"func,omitempty"`
	HTTPMethod          *string                 `json:"http-method,omitempty"`
	HTTPStatus          *int                    `json:"http-status,omitempty"`
	Message             *string                 `json:"msg"`
	Stack               *string                 `json:"stack,omitempty"`
	Time                *time.Time              `json:"time"`
	URL                 *string                 `json:"url,omitempty"`
}

// Clone makes a complete copy of the entry
// Use this to copy fields from a logwriter right before sending them
func (le *MicroserviceLogEntry) Clone() *MicroserviceLogEntry {
	if le == nil {
		return nil
	}

	return &MicroserviceLogEntry{
		AccountID:             le.AccountID,
		BusinessType:          le.BusinessType,
		Channel:               le.Channel,
		Code:                  le.Code,
		CommonErrorCode:       le.CommonErrorCode,
		Count:                 le.Count,
		Date:                  le.Date,
		Detail:                le.Detail,
		Driver:                le.Driver,
		ElapsedMicroseconds:   le.ElapsedMicroseconds,
		Environment:           le.Environment,
		Error:                 le.Error,
		Function:              le.Function,
		HostName:              le.HostName,
		HTTPMethod:            le.HTTPMethod,
		HTTPStatus:            le.HTTPStatus,
		ID:                    le.ID,
		InstanceName:          le.InstanceName,
		IPAddress:             le.IPAddress,
		JobID:                 le.JobID,
		Level:                 le.Level,
		Message:               le.Message,
		Name:                  le.Name,
		Number:                le.Number,
		PolicyID:              le.PolicyID,
		PublicID:              le.PublicID,
		QuoteID:               le.QuoteID,
		QuoteNumber:           le.QuoteNumber,
		Stack:                 le.Stack,
		State:                 le.State,
		Time:                  le.Time,
		TransactionBrand:      le.TransactionBrand,
		TransactionDomain:     le.TransactionDomain,
		TransactionID:         le.TransactionID,
		TransactionInstance:   le.TransactionInstance,
		TransactionIntegrator: le.TransactionIntegrator,
		TransactionIPAddress:  le.TransactionIPAddress,
		TransactionSource:     le.TransactionSource,
		TransactionType:       le.TransactionType,
		URL:                   le.URL,
		Vehicle:               le.Vehicle,
		VIN:                   le.VIN,
	}
}

func (le *MicroserviceLogEntry) equals(other *MicroserviceLogEntry) bool {
	if le == nil || other == nil {
		return le == other
	}

	// STRING fields
	stringtests := []struct {
		k *string
		v *string
	}{
		{le.AccountID, other.AccountID},
		{le.InstanceName, other.InstanceName},
		{le.Channel, other.Channel},
		{le.Code, other.Code},
		{le.Date, other.Date},
		{le.Detail, other.Detail},
		{le.Driver, other.Driver},
		{le.Error, other.Error},
		{le.Function, other.Function},
		{le.HostName, other.HostName},
		{le.HTTPMethod, other.HTTPMethod},
		{le.ID, other.ID},
		{le.IPAddress, other.IPAddress},
		{le.JobID, other.JobID},
		{le.Message, other.Message},
		{le.Name, other.Name},
		{le.PolicyID, other.PolicyID},
		{le.PublicID, other.PublicID},
		{le.QuoteID, other.QuoteID},
		{le.QuoteNumber, other.QuoteNumber},
		{le.Stack, other.Stack},
		{le.TransactionID, other.TransactionID},
		{le.TransactionInstance, other.TransactionInstance},
		{le.TransactionIPAddress, other.TransactionIPAddress},
		{le.TransactionType, other.TransactionType},
		{le.URL, other.URL},
		{le.Vehicle, other.Vehicle},
		{le.VIN, other.VIN},
	}

	for _, thisTest := range stringtests {
		if thisTest.k == nil || thisTest.v == nil {
			if thisTest.k != thisTest.v {
				// one is nil, the other is not
				return false
			}
		} else if *thisTest.k != *thisTest.v {
			return false
		}
	}

	// INT Fields
	inttests := []struct {
		k *int
		v *int
	}{
		{le.Count, other.Count},
		{le.HTTPStatus, other.HTTPStatus},
		{le.Number, other.Number},
	}

	for _, thisTest := range inttests {
		if thisTest.k == nil || thisTest.v == nil {
			if thisTest.k != thisTest.v {
				// one is nil, the other is not
				return false
			}
		} else if *thisTest.k != *thisTest.v {
			return false
		}
	}

	// INT64
	if le.ElapsedMicroseconds == nil || other.ElapsedMicroseconds == nil {
		if le.ElapsedMicroseconds != other.ElapsedMicroseconds {
			// one is nil, the other is not
			return false
		}
	} else if *le.ElapsedMicroseconds != *other.ElapsedMicroseconds {
		return false
	}

	// ENUM Fields
	if le.BusinessType == nil || other.BusinessType == nil {
		if le.BusinessType != other.BusinessType {
			return false
		}
	} else if *le.BusinessType != *other.BusinessType {
		return false
	}

	if le.TransactionBrand == nil || other.TransactionBrand == nil {
		if le.TransactionBrand != other.TransactionBrand {
			return false
		}
	} else if *le.TransactionBrand != *other.TransactionBrand {
		return false
	}

	if le.CommonErrorCode == nil || other.CommonErrorCode == nil {
		if le.CommonErrorCode != other.CommonErrorCode {
			return false
		}
	} else if *le.CommonErrorCode != *other.CommonErrorCode {
		return false
	}

	if le.Environment == nil || other.Environment == nil {
		if le.Environment != other.Environment {
			return false
		}
	} else if *le.Environment != *other.Environment {
		return false
	}

	if le.Level == nil || other.Level == nil {
		if le.Level != other.Level {
			return false
		}
	} else if *le.Level != *other.Level {
		return false
	}

	if le.State == nil || other.State == nil {
		if le.State != other.State {
			return false
		}
	} else if *le.State != *other.State {
		return false
	}

	if le.TransactionDomain == nil || other.TransactionDomain == nil {
		if le.TransactionDomain != other.TransactionDomain {
			return false
		}
	} else if *le.TransactionDomain != *other.TransactionDomain {
		return false
	}

	if le.TransactionIntegrator == nil || other.TransactionIntegrator == nil {
		if le.TransactionIntegrator != other.TransactionIntegrator {
			return false
		}
	} else if *le.TransactionIntegrator != *other.TransactionIntegrator {
		return false
	}

	if le.TransactionSource == nil || other.TransactionSource == nil {
		if le.TransactionSource != other.TransactionSource {
			return false
		}
	} else if *le.TransactionSource != *other.TransactionSource {
		return false
	}

	// TIME Fields

	if le.Time == nil || other.Time == nil {
		if le.Time != other.Time {
			return false
		}
	} else if *le.Time != *other.Time {
		return false
	}

	return true

}

// microserviceLogFields maps the JSON NAME of every recognized field to its PROPERTY NAME
var microserviceLogFields map[string]string

// init builds the microserviceLogFields array using reflection
func init() {
	buildMicroserviceLogFields()
}

// nameFromTag finds the JSON name associated with this field
func nameFromTag(tag string) string {
	rtn := tag
	if idx := strings.Index(tag, ","); idx != -1 {
		rtn = tag[:idx]
	}

	if rtn == `-` {
		rtn = ``
	}

	return rtn
}

// buildMicroserviceLogFields uses reflection to build a map[string]string
// describing the fields of the MicroserviceLogEntry type
// the key value is the JSON tag name of the field (e.g., "app")
// the value is the name of the actual field in the struct (e.g., ServiceID)
func buildMicroserviceLogFields() {
	foo := map[string]string{}
	bar := MicroserviceLogEntry{InstanceName: uf.Pointer.ToString(`nothing`)}
	msltype := reflect.TypeOf(bar)
	for i := 0; i < msltype.NumField(); i++ {
		field := msltype.Field(i)
		tag := field.Tag.Get(`json`)
		jsonName := nameFromTag(tag)
		fieldName := field.Name
		foo[jsonName] = fieldName
	}

	microserviceLogFields = foo
}

// put any fields we DON'T want to set with WithConsoleField(string, interface{}) here
// we are setting and printing these fields explicitly by name,
// so we don't also want them printing out with the rest of the map
var microserviceLogSkipMapEntries = map[string]bool{
	`app`:   true,
	`level`: true,
	`msg`:   true,
}

func (le *MicroserviceLogEntry) setPropertyByName(propertyName string, value interface{}) {
	// can't do anything with a nil target, so fail silently
	if le == nil {
		debugLog(`attempt to set property on nil MicroserviceLogEntry`)
		return
	}

	// all the properties of MicroserviceLogEntry are pointers,
	// so the value passed in must also be a pointer
	valueVal := reflect.ValueOf(value)
	if valueVal.Kind() != reflect.Ptr {
		valueVal = reflect.ValueOf(&value)
	}
	if valueVal.Kind() != reflect.Ptr {
		debugLogf(`attempt to set property to non-pointer value: %v`, propertyName)
		debug.PrintStack()
		return
	}

	// the type of thing value is pointing to
	valTargetType := valueVal.Type().Elem()

	// pointer value
	lePtrVal := reflect.ValueOf(le)

	// target MicroserviceLogEntry value
	leTargetVal := lePtrVal.Elem()

	// lookup the field by name in the MicroserviceLogEntry type
	theFieldVal := leTargetVal.FieldByName(propertyName)

	// the type pointed to by the field
	theFieldValTargetType := theFieldVal.Type().Elem()

	// the type of pointer passed in has to match the type of pointer in this field of MicroserviceLogEntry, or we fail silently
	if theFieldValTargetType != valTargetType {
		debugLogf(`property type does not match submitted type: %v`, propertyName)
		return
	}

	if theFieldVal.IsValid() {
		if theFieldVal.CanSet() {
			theFieldVal.Set(valueVal)
		} else {
			debugLogf(`field value not settable: %v`, propertyName)
		}
	} else {
		debugLogf(`field value not valid: %v`, propertyName)
	}
}

func (le *MicroserviceLogEntry) setPropertyByJSONName(jsonName string, value interface{}) {
	if le == nil || jsonName == `` {
		return
	}

	if propertyName, ok := microserviceLogFields[jsonName]; ok && propertyName != `` {
		le.setPropertyByName(propertyName, value)
	}
}
