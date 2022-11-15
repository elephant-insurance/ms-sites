package log

import (
	"reflect"

	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
)

// The functions allow us to add logging values to an HTTP Context.
// Once a logging value has been added to a context, any "C" (context-aware) logging calls
// will include all of the key-value pairs that have been added to that context
// by this logging package.

// When a field is added to the context, the field's key/name is also added to the fieldKeyList
// stored in the context.

// fieldsFromContext reads through the submitted context for fields stored
// to it by this package. It returns a list of label->value pairs that can
// be added to any logger instance.
// Also included are the standard TX fields, although these are now managed by the msrqc package.
func fieldsFromContext(c msrqc.Context) map[string]interface{} {
	rtn, exists := msrqc.NamespaceDump(c, &msrqc.NamespaceKeyLog)
	if !exists || rtn == nil {
		rtn = map[string]interface{}{}
	}

	// add the tx-fields
	if txfields, txferr := msrqc.GetTXFields(c); txferr == nil && txfields != nil {
		for k, v := range txfields {
			if k != `` {
				rtn[k] = v.Value
			}
		}
	}
	return rtn
}

// setContextValue adds a value to the context
// Any future logging using this context will include this field
// with the given label.
// This method will not overwrite an existing value unless force = true.
func setContextValue(c msrqc.Context, key string, value interface{}, forceOverwrite bool) {
	if c == nil {
		return
	}

	if reflect.ValueOf(value).Kind() != reflect.Ptr {
		debugLog(`attempt to set context field to non-pointer value: ` + key)
		debugStack()
		return
	}

	msrqc.NamespaceSet(c, &msrqc.NamespaceKeyLog, &key, value, forceOverwrite)
}
