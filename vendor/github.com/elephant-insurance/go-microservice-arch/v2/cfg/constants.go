package cfg

import (
	"reflect"
	"time"

	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
	"github.com/shopspring/decimal"
)

const (
	// FieldTagName : All config field tags begin with this identifier
	FieldTagName = "config"

	// FieldTagFinal : Fields marked with this tag will not be overwritten
	FieldTagFinal = "final"

	// FieldTagMustOverride : Fields marked with this tag are blanked out before they're overridden
	FieldTagMustOverride = "mustoverride"

	// FieldTagOptional : Fields marked with this tag need not have values
	FieldTagOptional = "optional"

	// FieldTagPublic : Fields marked with this tag may be displayed in diagnostic output
	FieldTagPublic = `public`

	// FieldNameRequiredConfig : This is the name of the embedded required config type
	FieldNameRequiredConfig = "RequiredConfig"

	// FieldNameOverridePath : Within the RequiredConfig struct, this is the name of the field
	//  specifying the full path of the override file
	FieldNameOverridePath = "OverrideConfigPath"

	fieldNameAllowCompressedRequests     = `AllowCompressedRequests`
	fieldNameAllowedHeaders              = `AllowedHeaders`
	fieldNameAllowedMethods              = `AllowedMethods`
	fieldNameAllowedOrigins              = `AllowedOrigins`
	fieldNameAppAbbreviation             = `AppAbbreviation`
	fieldNameEnvironment                 = `Environment`
	fieldNameExposedHeaders              = `ExposedHeaders`
	fieldNameInstanceName                = `InstanceName`
	fieldNameListenPort                  = `DevTestListenPort`
	fieldNameLogLevel                    = `LogLevel`
	fieldNameServiceID                   = `ServiceID`
	defaultListenPort                int = 4000

	errorArgMustBeValidStruct    = `argument must be a valid struct`
	errorArgMustBeValidStructPtr = `argument must be a pointer to a valid struct`

	// EnvironmentVariableNamePrefix : Any environment variable that should override a config value must be named
	// EnvironmentVariableNamePrefix + strings.ToUpper(fieldName)
	EnvironmentVariableNamePrefix = "MSVC_"
)

// structValueTypes contains the reflect.Types of structs that are not treated as structs by this package
// Fields with struct types that are understood to describe a single value using mutiple fields
// (e.g., time, fraction, date, etc.) should be added here IF they are needed for microservice configuration

// If a field has a Type listed here,
// 1. It must be a pointer to the type, or verifyReferenceFields will raise an error.
// 2. "Crawler" methods will not recurse into these structures or examine their fields.
var structValueTypes = []reflect.Type{
	reflect.TypeOf(time.Now()),
	reflect.TypeOf(decimal.Zero),
	reflect.TypeOf(uf.DateFactory.Today()),
}
