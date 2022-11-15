package cfg

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/davecgh/go-spew/spew"
	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/crasher"
	"github.com/imdario/mergo"
	yaml "gopkg.in/yaml.v2"
)

const (
	errNoBase              = "no base config file found"
	errBaseNotParseable    = "base config file could not be parsed"
	errTokensNotFound      = "tokenized config file could not be read at location"
	errOverrideNotReadable = "override config file could not be read at location"
	errMustPassPointer     = "must pass a pointer to config object to LoadConfig"
	errMustPassLogger      = "must pass a valid logger to LoadConfig"
	errBase64Decode        = `error attempting to base64 decode override file`
)

var (
	// cfgCrasher crashes the app when it is not properly configured.
	// It may be substituted with a test Crasher for testing.
	cfgCrasher crasher.Crasher = crasher.NewRealCrasher()

	// Base64OverrideToken => if an override filename contains this string, then we will first attempt to base64 decode the file before parsing it
	Base64OverrideToken string = `-base64`

	// BaseConfigPath is the config path for the base configuration file.
	BaseConfigPath = "./config.yml"

	// TokenizedConfigPath is the path to the yaml file with
	//  automation-friendly override token values
	// If present, config will ensure that all tokens have been replaced.
	TokenizedConfigPath = "./config-tokens.yml"

	// CheckEnvironment defaults to true.
	// If true, config will look for environment variable overrides.
	CheckEnvironment = true

	tokenDict map[interface{}]interface{}

	log logger

	// BaseConfig contains a pointer to the original base config after LoadCofig is called
	BaseConfig interface{}

	// OverrideConfig contains a pointer to the override config (if any) after LoadCofig is called
	OverrideConfig interface{}
)

// LoadConfig attempts to load the app configuration located in './config.yml'
// and to override its values with any it finds in the override config or in
// environment variables.
// LoadConfig crashes the application if it finds any problems.
func LoadConfig(target interface{}) {
	loadConfig(target)
}

// loadConfig is the main entrypoint for the Config library
// It takes a pointer to an empty config struct and fills it out from the base config file
// It then loads the override config file, if any, and merges the two structs
// TODO: this method should be shorter and call some subfuncs to make it easier to follow.
func loadConfig(target interface{}) {
	log := &logger{}

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		cfgCrasher.Fatal(errMustPassPointer)
		return
	}
	targetActual := targetValue.Elem()
	configType := targetActual.Type()
	baseptr := reflect.New(configType)
	// tokenptr := reflect.New(configType)

	// load base config
	if _, err := os.Stat(BaseConfigPath); !os.IsNotExist(err) {
		raw, err := ioutil.ReadFile(BaseConfigPath)
		if err != nil {
			cfgCrasher.Fatal(fmt.Sprintf("base config file could not be read at location \"%v\": %v", BaseConfigPath, err.Error()))
			return
		}

		err = yaml.UnmarshalStrict(raw, targetValue.Interface())
		if err != nil {
			cfgCrasher.Fatal(errBaseNotParseable + ": " + err.Error())
			return
		}
	} else {
		cfgCrasher.Fatal(errNoBase)
		return
	}

	// make a copy of the base config for later reference:
	if merr := mergo.Merge(baseptr.Interface(), targetValue.Interface(), mergo.WithOverride); merr != nil {
		cfgCrasher.Fatal("Error from mergo while saving base config: " + merr.Error())
	}

	BaseConfig = baseptr.Interface()

	// find override file path
	targetVal := reflect.ValueOf(targetValue.Elem().Interface())
	requiredConfig := targetVal.FieldByName(FieldNameRequiredConfig)
	overrideField := requiredConfig.FieldByName(FieldNameOverridePath).Interface().(string)

	rci := requiredConfig.Interface()
	if _, ok := rci.(RequiredConfig); !ok {
		cfgCrasher.Fatal(`error casting base RequiredConfig`)
		return
	}

	// load override config, if any
	if overrideField != "" {
		b64encoded := strings.Contains(overrideField, Base64OverrideToken)
		if _, err := os.Stat(overrideField); !os.IsNotExist(err) {
			ovrptr := reflect.New(configType)
			var raw []byte
			raw, err = ioutil.ReadFile(overrideField)
			if err != nil {
				cfgCrasher.Fatal(fmt.Sprintf(errOverrideNotReadable+" \"%v\": %v", overrideField, err.Error()))
				return
			}

			if b64encoded {
				raw, err = base64Decode(raw)
				if err != nil {
					cfgCrasher.Fatal(fmt.Sprintf(errBase64Decode+" \"%v\": %v", overrideField, err.Error()))
					return
				}
			}

			if yamlerr := yaml.Unmarshal(raw, ovrptr.Interface()); yamlerr != nil {
				cfgCrasher.Fatal("Override config file could not be parsed! \n" + yamlerr.Error())
				return
			}

			// make a copy of the override config
			ovrsvptr := reflect.New(configType)
			if ovrsaverr := mergo.MergeWithOverwrite(ovrsvptr.Interface(), ovrptr.Interface()); ovrsaverr != nil {
				cfgCrasher.Fatal(`error override configuration files: ` + ovrsaverr.Error())
				return
			}
			OverrideConfig = ovrsvptr.Interface()

			// process override
			if mergoerr := overwrite(targetValue.Interface(), ovrptr.Interface()); mergoerr != nil {
				cfgCrasher.Fatal(`error merging configuration files: ` + mergoerr.Error())
				return
			}

			OverrideConfig = ovrptr.Interface()
		} else {
			// an override was specified, but doesn't exist
			log.Warn(fmt.Sprintf("Specified config override file \"%v\" was not found: %v", overrideField, err.Error()))
			log.Warn("Unless this a development build, PLEASE STOP THIS SERVICE IMMEDIATELY and correct the app configuration!")
		}
	} else {
		// Warn if no config override is specified
		log.Warn("WARNING: No override configuration specified!")
	}

	// process any environment variables
	mergedPtr := reflect.ValueOf(targetValue.Interface())
	if enverr := applyEnvironmentOverrides(mergedPtr); enverr != nil {
		cfgCrasher.Fatal(`error merging applying environment variables: ` + enverr.Error())
		return
	}

	// pull the required config fields into their private stores
	mergedVal := reflect.ValueOf(targetValue.Elem().Interface())
	requiredConfig = mergedVal.FieldByName(FieldNameRequiredConfig)
	rci = requiredConfig.Interface()
	rqc, ok := rci.(RequiredConfig)

	if !ok {
		cfgCrasher.Fatal(`failed to cast merged RequiredConfig`)
		return
	}

	if sferr := setRequiredConfigFields(requiredConfig); sferr != nil {
		cfgCrasher.Fatal(`error loading required config: ` + sferr.Error())
		return
	}

	// deep-dump the config to stderr if we are in dev:
	if rqc.Environment.Equals(&enum.ServiceEnvironment.Development.ID) {
		spew.Dump(targetValue.Elem().Interface())
	}

	return
}

func base64Decode(cryptText []byte) (clearText []byte, err error) {
	var decodedLength int
	clearText = make([]byte, base64.StdEncoding.DecodedLen(len(cryptText)))
	decodedLength, err = base64.StdEncoding.Decode(clearText, cryptText)
	if err != nil {
		return
	}
	return clearText[:decodedLength], nil
}

func setRequiredConfigFields(rc reflect.Value) error {
	// Allow Compressed Requests
	// A malicious user could send a bad request that decompresses infititely
	// So by default we leave this set to false
	// This field must be overridden in order to work
	requiredAllowCompressedRequests = false
	allowCompressedRequestsBoolPtr := rc.FieldByName(fieldNameAllowCompressedRequests)
	if allowCompressedRequestsBoolPtr.IsValid() {
		allowCompressedRequests, ok := allowCompressedRequestsBoolPtr.Interface().(*bool)
		if ok && allowCompressedRequests != nil && *allowCompressedRequests {
			requiredAllowCompressedRequests = true
		}
	}

	// Allowed Methods
	allowedMethodsField := rc.FieldByName(fieldNameAllowedMethods)
	if allowedMethodsField.IsValid() {
		allowedMethodsString, ok := allowedMethodsField.Interface().(string)
		if !ok || allowedMethodsString == `` {
			requiredAllowedMethods = []string{}
		} else {
			requiredAllowedMethods = strings.Split(allowedMethodsString, `,`)
		}
	} else {
		requiredAllowedMethods = []string{}
	}

	// Allowed Headers
	allowedHeadersField, ok := rc.FieldByName(fieldNameAllowedHeaders).Interface().(string)
	if !ok || allowedHeadersField == `` {
		requiredAllowedHeaders = []string{}
	} else {
		requiredAllowedHeaders = strings.Split(allowedHeadersField, `,`)
	}

	// Allowed Origins
	allowedOriginsField, ok := rc.FieldByName(fieldNameAllowedOrigins).Interface().(string)
	if !ok {
		return errors.New(`could not cast submitted allowed origins list to type string`)
	} else if allowedOriginsField == `` {
		return errors.New(`missing required configuration field AllowedOrigins`)
	}
	requiredAllowedOrigins = strings.Split(allowedOriginsField, `,`)

	// App Abbreviation
	appAbbreviationField, ok := rc.FieldByName(fieldNameAppAbbreviation).Interface().(string)
	if !ok {
		return errors.New(`could not cast submitted app abbreviation to type string`)
	} else if appAbbreviationField == `` {
		return errors.New(`missing required configuration field AppAbbreviation`)
	}
	requiredAppAbbreviation = appAbbreviationField

	// Environment
	environmentField, ok := rc.FieldByName(fieldNameEnvironment).Interface().(enum.ServiceEnvironmentID)
	if !ok {
		return errors.New(`could not cast submitted service environment id to type enum.ServiceEnvironmentID`)
	} else if !(&environmentField).Valid() {
		return errors.New(`invalid environment`)
	}
	requiredEnvironment = environmentField

	// Exposed Headers
	exposedHeadersField, ok := rc.FieldByName(fieldNameExposedHeaders).Interface().(string)
	if !ok || exposedHeadersField == `` {
		requiredExposedHeaders = []string{}
	} else {
		requiredExposedHeaders = strings.Split(exposedHeadersField, `,`)
	}

	// Service ID & Instance Name
	serviceNameField, ok := rc.FieldByName(fieldNameServiceID).Interface().(enum.ServiceID)
	if !ok {
		return errors.New(`could not cast submitted service id to type enum.ServiceID`)
	} else if !(&serviceNameField).Valid() {
		return errors.New(`invalid service id`)
	}

	// The instance name comes from the set service ID:
	requiredInstanceName = serviceNameField.ToIDString()

	// The "real" serviceid is the serviceid of the parent service, if any:
	service := serviceNameField.Parent()
	for !service.Equals(service.Parent()) {
		service = service.Parent()
	}
	requiredServiceID = *service

	// Instance Name can be overridden:
	instanceNameField, ok := rc.FieldByName(fieldNameInstanceName).Interface().(string)
	if ok && instanceNameField != `` {
		requiredInstanceName = instanceNameField
	}

	// Listen Port ONLY works in DEV and TEST
	listenPortField := rc.FieldByName(fieldNameListenPort)
	lp := defaultListenPort
	if (requiredEnvironment.Equals(&enum.ServiceEnvironment.Development.ID) || requiredEnvironment.Equals(&enum.ServiceEnvironment.Testing.ID)) && listenPortField.IsValid() {
		listenPortPtr, ok2 := listenPortField.Interface().(*int)
		if ok2 && listenPortPtr != nil {
			lp = *listenPortPtr
		}
	}

	requiredListenPort = fmt.Sprintf(`:%v`, lp)

	// Log Level
	logLevelField, ok := rc.FieldByName(fieldNameLogLevel).Interface().(enum.LogLevelID)
	if !ok {
		return errors.New(`could not cast submitted log level id to type enum.LogLevelID`)
	} else if !(&logLevelField).Valid() {
		return errors.New(`invalid log level`)
	}
	requiredLogLevel = logLevelField

	requiredConfigLoaded = true

	return nil
}
