package uf

import (
	"errors"
	"reflect"
	"strings"
)

// Methods that use reflection to do magical things.
// These methods tend to be complex, tedious, and
// difficult to debug, so once we get them working,
// save them here so we never have to write them again.

type reflector struct{}

// TagSelector is a simple func that evaluates a struct tag
// and returns a bool indicating whether the tag for a struct
// field meets a certain condition.
// reflect.StructTag is just a type alias for string with some
// helpful methods added, so these are easy to write.
// fieldType is included in case you want your selector to act
// differently for different field types.
type TagSelector func(fieldTag reflect.StructTag, fieldType reflect.Type) bool

// Reflect implements reflection utility methods.
var Reflect = &reflector{}

// ClearTaggedFields requires the reflect.Value of a pointer to a struct.
// Any field marked with the submitted tag in the struct,
// or in any sub-struct, is set to is zero/nil value
// selector should return true for any field that should be cleared.
func (r *reflector) ClearTaggedFields(cfgStructPtr reflect.Value, selector TagSelector) error {
	// don't have log package to help us out with the func name so add this for clarity:
	const funcName = `uf.Reflect.ClearTaggedFields: `
	if !cfgStructPtr.IsValid() {
		return errors.New(funcName + errorArgMustBeValidStructPtr)
	}
	if cfgStructPtr.Kind() != reflect.Ptr {
		return errors.New(funcName + errorArgMustBeValidStructPtr)
	}
	if cfgStructPtr.IsNil() {
		return errors.New(funcName + errorArgMustBeValidStructPtr)
	}

	cfgPtrType := cfgStructPtr.Type()
	if cfgPtrType.Elem().Kind() != reflect.Struct {
		return errors.New(funcName + errorArgMustBeValidStructPtr)
	}

	cfgType := cfgPtrType.Elem()
	cfgStruct := cfgStructPtr.Elem()

	numFields := cfgType.NumField()
	for i := 0; i < numFields; i++ {
		thisFieldValue := cfgStruct.Field(i)
		if !thisFieldValue.IsValid() || !thisFieldValue.CanSet() {
			continue
		}

		thisStructField := cfgType.Field(i)

		if selector(thisStructField.Tag, thisStructField.Type) {
			nilval := reflect.New(thisStructField.Type)
			thisFieldValue.Set(nilval.Elem())
			continue
		}

		thisFieldKind := thisStructField.Type.Kind()

		switch thisFieldKind {
		case reflect.Struct:
			if err := r.ClearTaggedFields(thisFieldValue.Addr(), selector); err != nil {
				return err
			}
			continue
		case reflect.Ptr:
			if thisFieldValue.IsNil() || thisStructField.Type.Elem().Kind() != reflect.Struct {
				continue
			}

			if err := r.ClearTaggedFields(thisFieldValue, selector); err != nil {
				return err
			}

			continue
		default:
			continue
		}
	}

	return nil
}

// ByNamespaceAndValue returns a TagSelector that finds tags with the given
// tagNamespace (e.g., "json") and tagValue (e.g., "omitempty").
// For a submitted StructTag, the returned TagSelector will return
// the value of returnIfPresent if the tagValue is found within the
// tagNamespace of the submitted StructTag, and the opposite value if
// it is not found.
func (r *reflector) ByNamespaceAndValue(tagNamespace, tagValue string, returnIfPresent bool) TagSelector {
	return func(fieldTag reflect.StructTag, fieldType reflect.Type) bool {
		if myTag, ok := fieldTag.Lookup(tagNamespace); ok {
			if strings.Contains(myTag, tagValue) {
				return returnIfPresent
			}
		}

		return !returnIfPresent
	}
}

// FindEmptyTaggedFields walks the config struct and detects fields that are empty and tagged.
// A field is considered "tagged" if selector returns true for its StructTag.
// It returns an array of dot-separated field names if it finds any missing fields.
// It returns an error only if it is unable to walk the struct for any reason.
// A field is "missing" if it is a nil pointer field or a value field that is not a struct type but is Zero.
// Be wary of fields of type string, int, or bool, which could validly equal their zero values ("", 0, false).
// For this reason it is strongly recommended to use reference values, such as pointers, in structs used with this func.
func (r *reflector) FindEmptyTaggedFields(cfgStructPtr reflect.Value, selector TagSelector) (missingFieldNames []string, err error) {
	const funcName = `uf.Reflect.FindEmptyTaggedFields: `
	if !cfgStructPtr.IsValid() {
		err = errors.New(funcName + errorArgMustBeValidStructPtr)
		return
	}
	if cfgStructPtr.Kind() != reflect.Ptr {
		err = errors.New(funcName + errorArgMustBeValidStructPtr)
		return
	}
	if cfgStructPtr.IsNil() {
		err = errors.New(funcName + errorArgMustBeValidStructPtr)
		return
	}

	cfgPtrType := cfgStructPtr.Type()
	if cfgPtrType.Elem().Kind() != reflect.Struct {
		err = errors.New(funcName + errorArgMustBeValidStructPtr)
		return
	}

	cfgType := cfgPtrType.Elem()
	cfgStruct := cfgStructPtr.Elem()
	numFields := cfgType.NumField()

	for i := 0; i < numFields; i++ {
		var (
			thisStructField = cfgType.Field(i)
			thisFieldName   = thisStructField.Name
			thisFieldValue  = cfgStruct.Field(i)
			thisFieldType   = thisStructField.Type
			thisFieldKind   = thisFieldType.Kind()
		)

		if selector(thisStructField.Tag, thisFieldType) {
			if !thisFieldValue.IsValid() {
				missingFieldNames = append(missingFieldNames, thisFieldName)
				continue
			}

			// if field is struct or pointer to struct and not empty/nil, process it recursively
			// otherwise regardless of type/kind if required and empty add an error message
			switch thisFieldKind {
			case reflect.Struct:
				subVals, subErr := r.FindEmptyTaggedFields(thisFieldValue.Addr(), selector)
				if subErr != nil {
					err = subErr
					return
				}

				for vcnt := 0; vcnt < len(subVals); vcnt++ {
					missingFieldNames = append(missingFieldNames, thisFieldName+`.`+subVals[vcnt])
				}

				continue

			case reflect.Ptr:
				if thisFieldValue.IsNil() {
					missingFieldNames = append(missingFieldNames, thisFieldName)
					continue
				}
				if thisStructField.Type.Elem().Kind() == reflect.Struct {
					subVals, subErr := r.FindEmptyTaggedFields(thisFieldValue, selector)
					if subErr != nil {
						err = subErr
						return
					}

					for vcnt := 0; vcnt < len(subVals); vcnt++ {
						missingFieldNames = append(missingFieldNames, thisFieldName+`.`+subVals[vcnt])
					}

					continue
				}

			default:
				if thisFieldValue.IsZero() {
					missingFieldNames = append(missingFieldNames, thisFieldName)
				}
				continue
			}
		}
	}

	return
}

const (
	errorArgMustBeValidStructPtr = `argument must be a pointer to a valid struct`
)
