package cfg

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
)

// Validate is the outer context for the recursive validateStruct function
// It tests the config against known validation rules and returns a list of problems it finds
// Currently the only rules checked are:
//  empty fields must be marked "optional"
// Validate will NOT detect bad values in leaf nodes of map values
// If you need any kind of validation of map values, you must implement
//  one of the validation interfaces (PreValidated or PostValidated) appropriately
func Validate(conf interface{}) {
	rtn := []string{}
	confValue := reflect.ValueOf(conf)
	if !confValue.IsValid() {
		rtn = append(rtn, "Configuration passed to validate is not valid.")
		cfgCrasher.Fatal(strings.Join(rtn, `, `))
		return
	}

	if confValue.Kind() != reflect.Ptr {
		// gotta have a pointer here
		rtn = append(rtn, `must pass a pointer type to cfg.Validate()`)
		cfgCrasher.Fatal(strings.Join(rtn, `, `))
		return
	}

	confTarget := reflect.Indirect(confValue)

	if confTarget.Kind() != reflect.Struct {
		rtn = append(rtn, fmt.Sprintf("Configuration passed to validate is not a struct, kind is \"%v\".", confTarget.Kind()))
		cfgCrasher.Fatal(strings.Join(rtn, `, `))
		return
	}

	// if the config type has a PreValidate method, run that first
	if ifc, ok := conf.(PreValidated); ok {
		rtn = ifc.PreValidate()
	}

	// Here is where we run any global validation rules
	if problems, verr := uf.Reflect.FindEmptyTaggedFields(confValue, hasNoOptionalTag); verr != nil {
		cfgCrasher.Fatal(verr.Error())
		return
	} else if len(problems) > 0 {
		for _, fieldName := range problems {
			msg := fmt.Sprintf(`INVALID CONFIG: field %v is empty, but required, crashing...`, fieldName)
			rtn = append(rtn, msg)
		}
	}

	// if the config type has a PostValidate method, run that last
	if ifc, ok := conf.(PostValidated); ok {
		rtn = ifc.PostValidate(rtn)
	}

	// if we picked up any error messages along the way, dump them and quit
	if len(rtn) > 0 {
		cfgCrasher.Fatal(strings.Join(rtn, `, `))
	}
}

/*

REPLACED by crawler funcs

func validateStruct(s reflect.Value) []string {
	rtn := []string{}
	structType := s.Type()

	var myFieldDict map[string]mergo.FieldInfo
	var ok bool
	if myFieldDict, ok = mergo.StructFieldDict[structType]; !ok {
		rtn = append(rtn, fmt.Sprintf("validate has no fieldInfo for type \"%v\".", structType))
		return rtn
	}

	for i := 0; i < structType.NumField(); i++ {
		thisStructField := structType.Field(i)
		thisConfField := s.Field(i)
		if thisStructField.PkgPath != "" {
			// not an exported field
			continue
		}

		thisFieldName := thisStructField.Name
		var thisFieldInfo mergo.FieldInfo
		if thisFieldInfo, ok = myFieldDict[thisFieldName]; !ok {
			rtn = append(rtn, fmt.Sprintf("validate has no fieldInfo for field \"%v\" of type \"%v\".", thisFieldName, structType))
			return rtn
		}

		// the problem children: structs and pointers to structs
		rightNow := time.Now()
		timeType := reflect.TypeOf(rightNow)
		if thisConfField.Kind() == reflect.Struct {
			// special handling for Time type
			if thisStructField.Type == timeType {
				// we'll just check for a zero value and quit
				zeroTime := time.Time{}
				if !thisFieldInfo.Optional && (!thisConfField.IsValid() || thisConfField.Interface() == zeroTime) {
					rtn = append(rtn, fmt.Sprintf("Field \"%v\" of struct type \"%v\" is empty but not optional.", thisFieldName, structType.Name()))
				}
				continue
			} else {
				rtn = append(rtn, validateStruct(thisConfField)...)
			}
		} else if thisConfField.Kind() == reflect.Ptr {
			if !thisConfField.IsValid() || thisConfField.IsNil() {
				if !thisFieldInfo.Optional {
					rtn = append(rtn, fmt.Sprintf("Field \"%v\" of struct type \"%v\" is empty but not optional.", thisFieldName, structType.Name()))
					continue
				}
			} else {
				targetKind := thisConfField.Elem().Kind()
				targetType := thisConfField.Elem().Type()
				if targetKind == reflect.Struct && targetType != timeType {
					//rtn = append(rtn, fmt.Sprintf("Field \"%v\" of struct type \"%v\" is empty but not optional.", thisFieldName, structType.Name()))
					continue
				}
			}
		}

		// check for empty value in a non-optional field
		if isEmptyValue(thisConfField) && !thisFieldInfo.Optional {
			rtn = append(rtn, fmt.Sprintf("Field \"%v\" of struct type \"%v\" is empty but not optional.", thisFieldName, structType.Name()))
		}

		// check for non-replaced token value
		if thisFieldInfo.Token != "" {
			myStringValue, err := valueToString(thisConfField)
			if err != nil {
				rtn = append(rtn, fmt.Sprintf("validate unable to generate string value for field \"%v\" of type \"%v\": %v", thisFieldName, structType, err.Error()))
			} else if myStringValue == thisFieldInfo.Token {
				rtn = append(rtn, fmt.Sprintf("Field \"%v\" of struct type \"%v\" matches its token value of \"%v\".", thisFieldName, structType.Name(), thisFieldInfo.Token))
			}
		}
	}
	return rtn
} */

func valueToString(v reflect.Value) (string, error) {
	t := v.Kind()

	// follow the pointer trail til we get to a scalar
	for t == reflect.Ptr {
		v = v.Elem()
		t = v.Kind()
	}

	switch t {
	case reflect.String:
		return v.String(), nil
	case reflect.Bool:
		return fmt.Sprintf("%v", v.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%v", v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return fmt.Sprintf("%v", v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%v", v.Float()), nil
	}
	return "", fmt.Errorf("valueToString received invalid kind \"%v\"", t)
}
