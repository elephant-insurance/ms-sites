package cfg

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
)

// methods for crawling the structure of a config struct applying rules and effects
// according to struct tags

// makeDiagnosticMap walks a config struct and all its sub-structs.
// Any property marked public is added to the return map as FieldName: Value.
func makeDiagnosticMap(cfgStruct reflect.Value) (map[string]interface{}, error) {
	if !cfgStruct.IsValid() || cfgStruct.Kind() != reflect.Struct {
		return nil, errors.New(errorArgMustBeValidStruct)
	}

	cfgType := cfgStruct.Type()
	numFields := cfgType.NumField()

	rtn := map[string]interface{}{}

	for i := 0; i < numFields; i++ {
		thisStructField := cfgType.Field(i)
		thisFieldName := thisStructField.Name
		thisFieldValue := cfgStruct.Field(i)

		if !thisFieldValue.IsValid() {
			continue
		}

		thisFieldKind := thisStructField.Type.Kind()

		thisFieldPublic := false
		if cfgTag, ok := thisStructField.Tag.Lookup(FieldTagName); ok {
			thisFieldPublic = strings.Contains(cfgTag, FieldTagPublic)
		}

		switch thisFieldKind {
		case reflect.Struct:
			subMap, err := makeDiagnosticMap(thisFieldValue)
			if err != nil {
				return nil, err
			}

			if len(subMap) == 0 {
				rtn[thisFieldName] = `(no public fields)`
			} else {
				rtn[thisFieldName] = subMap
			}
		case reflect.Ptr:
			if thisStructField.Type.Elem().Kind() != reflect.Struct {
				if thisFieldPublic {
					rtn[thisFieldName] = thisFieldValue.Interface()
				}
				break
			}

			if thisFieldValue.IsNil() {
				rtn[thisFieldName] = `(nil)`
			} else {
				subMap, err := makeDiagnosticMap(thisFieldValue.Elem())
				if err != nil {
					return nil, err
				}

				if len(subMap) == 0 {
					rtn[thisFieldName] = `(no public fields)`
				} else {
					rtn[thisFieldName] = subMap
				}
			}

		default:
			if thisFieldPublic {
				rtn[thisFieldName] = thisFieldValue.Interface()
			}
		}
	}

	return rtn, nil
}

// verifyReferenceFields verifies that every struct field in the submitted Type can take a nil value.
// This is necessary to fully distinguish empty (nil) values
// from values that happen to have the zero value for their type (string "", integer 0, bool false).
// The func returns an error if it finds any field in any struct or substruct of the submitted type
// with a Kind other than struct, ptr, map, slice, chan, func, or interface
// A field of Kind struct (not *struct) is considered an extension of the base type structure
// unless its Type is contained in the structValueTypes array.
// Any field that is of a Type in the structValueTypes array (not a pointer to the Type) will raise an error.
func verifyReferenceFields(cfgStructType reflect.Type) error {
	return nil
}

// applyEnvironmentOverrides checks the runtime environment for any settings
// that can override configuration fields.
// Matching environment settings are detected by their name
// The naming pattern must be provided by an Overridable interface on the app config struct.
// If Overridable is not implemented, then environment variables will not affect configuration.
func applyEnvironmentOverrides(cfgStructPtr reflect.Value) error {
	if !cfgStructPtr.IsValid() {
		return errors.New(errorArgMustBeValidStructPtr)
	}
	if cfgStructPtr.Kind() != reflect.Ptr {
		return errors.New(errorArgMustBeValidStructPtr)
	}
	if cfgStructPtr.IsNil() {
		return errors.New(errorArgMustBeValidStructPtr)
	}

	cfgPtrType := cfgStructPtr.Type()
	if cfgPtrType.Elem().Kind() != reflect.Struct {
		return errors.New(errorArgMustBeValidStructPtr)
	}

	cfgType := cfgPtrType.Elem()
	cfgStruct := cfgStructPtr.Elem()

	var covr Overridable
	var ok bool

	ovrtype := reflect.TypeOf((*Overridable)(nil)).Elem()
	overridable := !cfgStructPtr.IsZero() && cfgStructPtr.Type().Implements(ovrtype)
	if overridable {
		covr, ok = cfgStructPtr.Interface().(Overridable)
		if !ok {
			overridable = false
		}
	}

	numFields := cfgType.NumField()
	for i := 0; i < numFields; i++ {
		thisFieldValue := cfgStruct.Field(i)
		if !thisFieldValue.IsValid() || !thisFieldValue.CanSet() {
			continue
		}

		thisStructField := cfgType.Field(i)
		thisFieldIsFinal := false
		if fieldTag, ok := thisStructField.Tag.Lookup(FieldTagName); ok {
			thisFieldIsFinal = strings.Contains(fieldTag, FieldTagFinal)
		}

		if thisFieldIsFinal {
			// we don't override final fields, even from the environment
			continue
		}

		thisFieldKind := thisStructField.Type.Kind()
		thisFieldEnvironmentName := EnvironmentVariableNamePrefix + thisStructField.Name
		if overridable {
			thisFieldEnvironmentName = covr.GetEnvironmentSetting(thisStructField.Name)
		}

		switch thisFieldKind {
		case reflect.Struct:
			if err := applyEnvironmentOverrides(thisFieldValue.Addr()); err != nil {
				return err
			}
			continue
		case reflect.Ptr:
			if thisFieldValue.IsNil() || thisStructField.Type.Elem().Kind() != reflect.Struct {
				break
			}

			if err := applyEnvironmentOverrides(thisFieldValue); err != nil {
				return err
			}
			continue
		}

		valueFromEnvironment(thisFieldValue, thisFieldEnvironmentName)
	}

	return nil
}

const missingRequiredMsgTemplate string = `%v is nil or empty but not tagged optional`

// Overridable is an interface for configuration settings objects that can be overridden with environment variables
// If a struct implements Overridabale, the name of each field is run through the GetEnvironmentSetting func
// If an environment variable with the matching name is set, then the cvalue of that variable overrides any other setting
type Overridable interface {
	GetEnvironmentSetting(fieldName string) string
}

// valueFromEnvironment checks the submitted environment variable name for a value
// if a value is found, the submitted value is overwritten with the value of the environment variable
// this only works for certain scalar types: string, int, bool, and float64
// returns true if and only if the value is overwritten.
// If we can't find the environment variable with the name as submitted, then we try again
// with the name of the variable in ALL-CAPS.
// there is probably a very elegant way to handle the pointer case with a recursive call
// but I didn't feel like figuring it out after getting this here to work 8^)
func valueFromEnvironment(fieldValue reflect.Value, envVarName string) bool {
	// zero value
	// var z reflect.Value
	fieldType := fieldValue.Type()

	switch fieldType.Kind() {
	case reflect.Ptr:
		// this field is a pointer
		// if it is a pointer to a simple scalar type, we can check the environment for override variables
		pointedToType := fieldType.Elem()
		pointedToKind := fieldType.Elem().Kind()

		switch pointedToKind {
		case reflect.Bool:
			if envVal := getEnvironmentBool(envVarName); envVal != nil {
				if fieldValue.Elem().IsValid() {
					// the pointer in the base struct IS NOT NIL, so we can overwrite its target directly
					fieldValue.Elem().SetBool(*envVal)
					return true
				}
				// the pointer in the base struct IS NIL, so we can't address its target
				fieldValue.Set(reflect.ValueOf(envVal))
				return true
			}
		case reflect.Int:
			if envVal := getEnvironmentInt(envVarName); envVal != nil {
				if fieldValue.Elem().IsValid() {
					// the pointer in the base struct IS NOT NIL, so we can overwrite its target directly
					fieldValue.Elem().SetInt(int64(*envVal))
					return true
				}
				// the pointer in the base struct IS NIL, so we can't address its target
				fieldValue.Set(reflect.ValueOf(envVal))
				return true
			}
		case reflect.Float64:
			if envVal := getEnvironmentFloat64(envVarName); envVal != nil {
				if fieldValue.Elem().IsValid() {
					// the pointer in the base struct IS NOT NIL, so we can overwrite its target directly
					fieldValue.Elem().SetFloat(*envVal)
					return true
				}
				// the pointer in the base struct IS NIL, so we can't address its target
				fieldValue.Set(reflect.ValueOf(envVal))
				return true
			}
		case reflect.String:
			if envVal := getEnvironmentString(envVarName); envVal != nil {
				if pointedToType == reflect.TypeOf(`foo`) {
					// this is a regular string, so easy:
					if fieldValue.Elem().IsValid() {
						// the pointer in the base struct IS NOT NIL, so we can overwrite its target directly
						fieldValue.Elem().SetString(*envVal)
						return true
					}
					// the pointer in the base struct IS NIL, so we can't address its target
					fieldValue.Set(reflect.ValueOf(envVal))
					return true
				} else {
					envVal := reflect.ValueOf(*envVal)
					castedVal := envVal.Convert(pointedToType)
					if !fieldValue.Elem().IsValid() {
						// the pointer in the base struct IS NIL, so we have to set it
						empty := reflect.New(pointedToType)
						fieldValue = empty
					}
					// the pointer in the base struct IS NOT NIL, so we can overwrite its target directly
					fieldValue.Elem().Set(castedVal)
					return true
				}

			}
		default:
			// not a type we can work with
			// TODO: log an error
			//return z
			return false
		} // END pointer target kind switch

	case reflect.Bool:
		if envVal := getEnvironmentBool(envVarName); envVal != nil {
			fieldValue.SetBool(*envVal)
			return true
			//return reflect.ValueOf(*envVal)
		}
	case reflect.Int:
		if envVal := getEnvironmentInt(envVarName); envVal != nil {
			fieldValue.SetInt(int64(*envVal))
			return true
			//return reflect.ValueOf(*envVal)
		}
	case reflect.Float64:
		if envVal := getEnvironmentFloat64(envVarName); envVal != nil {
			fieldValue.SetFloat(*envVal)
			return true
			//return reflect.ValueOf(*envVal)
		}
	case reflect.String:
		if envVal := getEnvironmentString(envVarName); envVal != nil {
			fieldValue.SetString(*envVal)
			return true
			//return reflect.ValueOf(*envVal)
		}
	default:
		// not a type we can override with environment variables
		// TODO: log an error
		return false
	} // END field kind switch

	capECName := strings.ToUpper(envVarName)
	if capECName != envVarName {
		return valueFromEnvironment(fieldValue, capECName)
	}
	return false
}

func getEnvironmentString(vn string) *string {
	var rtn *string
	if value := os.Getenv(vn); value != "" {
		rtn = &value
	}

	return rtn
}

func getEnvironmentInt(vn string) *int {
	var rtn *int
	if strvalue := os.Getenv(vn); strvalue != "" {
		if value, err := strconv.Atoi(strvalue); err == nil {
			rtn = &value
		}
	}
	return rtn
}

func getEnvironmentBool(vn string) *bool {
	var rtn *bool
	if strvalue := os.Getenv(vn); strvalue != "" {
		if value, err := strconv.ParseBool(strvalue); err == nil {
			rtn = &value
		}
	}
	return rtn
}

func getEnvironmentFloat64(vn string) *float64 {
	var rtn *float64
	if strvalue := os.Getenv(vn); strvalue != "" {
		if value, err := strconv.ParseFloat(strvalue, 64); err == nil {
			rtn = &value
		}
	}
	return rtn
}

var (
	hasNoOptionalTag   uf.TagSelector = uf.Reflect.ByNamespaceAndValue(FieldTagName, FieldTagOptional, false)
	hasFinalTag        uf.TagSelector = uf.Reflect.ByNamespaceAndValue(FieldTagName, FieldTagFinal, true)
	hasMustOverrideTag uf.TagSelector = uf.Reflect.ByNamespaceAndValue(FieldTagName, FieldTagMustOverride, true)
)
