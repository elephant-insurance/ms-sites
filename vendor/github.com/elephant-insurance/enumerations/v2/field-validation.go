package enumerations

import (
	"errors"
	"fmt"
	"reflect"
)

// ValidatedID contains all the non-type-specific methods implemented by all validated ID types
// By checking for this interface, we can quickly determine whether a reflect.Value is a validated ID type
type ValidatedID interface {
	CapturedValue() *string
	MarshalJSON() ([]byte, error)
	ToIDString() string
	UnmarshalJSON(data []byte) error
	Valid() bool
}

// ValidateFIelds is just a wrapper for validateFields
// It performs basic type testing
func ValidateFields(docPtr interface{}) (map[string]*string, error) {
	var (
		docStructPtr   = reflect.ValueOf(docPtr)
		ptrValid       = docStructPtr.IsValid()
		ptrKind        = docStructPtr.Kind()
		ptrKindPointer = ptrKind == reflect.Ptr
		ptrNil         = ptrKindPointer && docStructPtr.IsNil()
		ptrElemKind    reflect.Kind
	)

	if ptrKindPointer && !ptrNil {
		ptrElemKind = docStructPtr.Elem().Kind()
	}

	if !ptrValid || ptrKind != reflect.Ptr || ptrNil || ptrElemKind != reflect.Struct {
		return nil, errors.New(`argument must be pointer to struct`)
	}

	return validateFields(docStructPtr)
}

// validateFields walks the submitted struct, checking each field to see whether it is a Validated ID type
// Any validated ID type fields are checked
// The argument docStructPtr must be the reflect.ValueOf a pointer to a struct.
// The returned map consists of fieldName: capturedValue for any validated ID type field where Valid() returns false
func validateFields(docStructPtr reflect.Value) (map[string]*string, error) {
	var (
		ptrValid = docStructPtr.IsValid()
		ptrKind  = docStructPtr.Kind()
		ptrNil   = (ptrKind == reflect.Ptr) && docStructPtr.IsNil()
	)

	// we don't return an error here because this could be a valid pointer to a sub-struct
	// other than a validated ID
	if !ptrValid || ptrKind != reflect.Ptr || ptrNil || docStructPtr.Elem().Kind() != reflect.Struct {
		return nil, nil
	}

	var (
		vidType    = reflect.TypeOf((*ValidatedID)(nil)).Elem()
		docPtrType = docStructPtr.Type()
		docType    = docPtrType.Elem()
		docStruct  = docStructPtr.Elem()
		numFields  = docType.NumField()
		rtn        = map[string]*string{}
	)

	for i := 0; i < numFields; i++ {
		var (
			thisFieldValue  = docStruct.Field(i)
			thisStructField = docType.Field(i)
			thisFieldName   = thisStructField.Name
			thisFieldType   = thisStructField.Type
			thisFieldKind   = thisFieldType.Kind()
			thisFieldIsVID  = thisFieldType.Implements(vidType)
		)

		// this field is a validated ID, so add it to the return if it is not valid
		if thisFieldIsVID {
			fieldInterface, ok := thisFieldValue.Interface().(ValidatedID)
			if ok && fieldInterface != nil {
				if !fieldInterface.Valid() {
					rtn[thisFieldName] = fieldInterface.CapturedValue()
				}
			}

			continue
		}

		// this field was not itself an ID, but it might be a struct with IDs in it
		var (
			subMap map[string]*string
			subErr error
		)

		// Is it a struct value or struct pointer?
		if thisFieldKind == reflect.Struct {
			// make a recursive call
			subMap, subErr = validateFields(thisFieldValue.Addr())
		} else if thisFieldKind == reflect.Ptr && thisFieldType.Elem().Kind() == reflect.Struct {
			subMap, subErr = validateFields(thisFieldValue)
		}

		if subErr != nil {
			return nil, subErr
		}

		// if we found any invalid ids in sub-documents, add them with fully-qualified field name
		if len(subMap) > 0 {
			for k, v := range subMap {
				if k != `` {
					subKey := fmt.Sprintf("%v.%v", thisFieldName, k)
					rtn[subKey] = v
				}
			}
		}
	}

	return rtn, nil
}
