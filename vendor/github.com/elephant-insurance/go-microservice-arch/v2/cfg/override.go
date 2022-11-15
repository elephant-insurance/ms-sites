package cfg

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
	"github.com/imdario/mergo"
)

func overwrite(base, override interface{}) error {
	// can't work with nils:
	if base == nil {
		return errors.New("overwriteFields: base config must not be nil")
	}

	// if nothing in the override, just return the original
	if override == nil {
		return nil
	}

	rcType := reflect.TypeOf(RequiredConfig{})
	baseVal := reflect.ValueOf(base)
	overVal := reflect.ValueOf(override)
	if baseVal.Kind() != reflect.Ptr || overVal.Kind() != reflect.Ptr {
		return errors.New("both values submitted to overwrite must be pointers to structs")
	}
	//bve := baseVal.Elem()
	bvt := baseVal.Type().Elem()
	ovt := overVal.Type().Elem()

	if bvt != ovt {
		return errors.New("src and dst must be of same type")
	}

	// spew.Dump(bvt)
	composes := false

	// check that we've embedded the RequiredConfig
	for i := 0; i < bvt.NumField(); i++ {
		if bvt.Field(i).Type == rcType {
			composes = true
			break
		}
	}

	if !composes {
		baseType := reflect.TypeOf(base)
		return fmt.Errorf("overwriteFields: config type %v.%v does not embed RequiredConfig", baseType.PkgPath(), baseType.Name())
	}

	// clear the value of any mustoverride fields before we merge.
	// if these are required and not overridden, the app will fatal out at startup.
	if clearerr := uf.Reflect.ClearTaggedFields(baseVal, hasMustOverrideTag); clearerr != nil {
		return clearerr
	}

	// now clear the value of any "final" fields in the override.
	// These should be empty, anyway, but clearing them before we merge makes sure.
	if clearerr := uf.Reflect.ClearTaggedFields(overVal, hasFinalTag); clearerr != nil {
		return clearerr
	}

	//spew.Dump("BEFORE", base, override)
	if mergeerr := mergo.Merge(base, override, mergo.WithOverride); mergeerr != nil {
		return mergeerr
	}

	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return isEmptyValue(v.Elem())
	case reflect.Func:
		return v.IsNil()
	case reflect.Invalid:
		return true
	}
	return false
}
