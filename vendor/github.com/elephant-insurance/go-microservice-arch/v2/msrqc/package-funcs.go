package msrqc

import (
	"context"
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
)

// Functions for getting and setting transaction (TX) headers in the request context

// New takes in a parent context and returns a full msrqc.Context.
// If the parent context passed in is *msrqcContext or *gin.Context, it is returned unchanged.
// If the parent context passed in is context.Context, it is wrapped in a new msrqc.Context and returned.
// If the parent context is nil or not a valid context, a new, empty msrqc.Context is returned.
// This enables inline conversion of virtually anything to an msrqc.Context:
//
//	result := doSomething(msrq.New(myContext)) // Doesn't matter whether myContext is *gin.Context or just context.Context
func New(parent interface{}) Context {
	if parent != nil {
		// if the parent context does what we want, just return it unchanged
		msrqcContext, ok := parent.(Context)
		if ok {
			return msrqcContext
		}

		// if the parent context is a valid context, wrap it in an msrqcContext
		baseContext, ok := parent.(context.Context)
		if ok {
			return newMSRQContext(baseContext)
		}
	}

	return newMSRQContext(context.Background())
}

// NewContext returns the same thing as New only as a plain context.Context, which isn't very useful.
// It is here to make sure that we're still fulfilling the context.Context interface.
// If we ever break that interface, this will fail to compile.
func NewContext(c context.Context) context.Context {
	return New(c)
}

// TODO: replace all context get/set with namespaced get/set if possible
// This will prevent future naming conflicts as more is added to the context

// TODO: investigate threading issues. Maybe add sync.Mutex to context directly?
// Threading is not a problem in log package because the logwriter keeps context and
// logwriter is explicitly NOT thread-safe, so should almost never be passed as a func param
// However if the same context is sent to N threads and log tries to set context value in each of the threads, may panic
// Currently we have a Mutex on the namespace (map), which may be sufficient for thread safety
// AS LONG AS every copy of the context already has the namespace itself created

// TODO: move ALL TX getters and setters into here.
// Unfortunately these will have to be package funcs, like msrqc.GetTransactionID(context), since we need
// to pass gin.Context as msrqc.Context, and we can't do that if we add more funcs to the interface.
// Putting TX getters and setter in here will make it much easier to use these values throughout the arch
// and in app code, where appropriate, instead of relying on the log package to set and retrieve values

func NamespaceDump(c Context, namespace *string) (valMap map[string]interface{}, exists bool) {
	exists = false
	if c == nil || namespace == nil || *namespace == `` {
		return
	}

	if val, ok := c.Get(*namespace); ok {
		if ns, stillOk := val.(*namespaceType); stillOk {
			exists = true
			valMap = ns.Dump()
		}
	}

	return
}

// NamespaceGet returns a value that has been set in the context with a Namespace
// The Namespace is used to distinguish the keys set by different packages
func NamespaceGet(c Context, namespace, key *string) (value interface{}, exists bool) {
	exists = false
	if c == nil || namespace == nil || key == nil || *namespace == `` || *key == `` {
		return
	}

	if val, ok := c.Get(*namespace); ok {
		if ns, stillOk := val.(*namespaceType); stillOk {
			value, exists = ns.Get(*key)
		}
	}

	return
}

// NamespaceSet sets a value in the context with a Namespace
// The Namespace is used to distinguish the keys set by different packages
// This func is intended primarily for use in packages
// Applications may use the c.Set method so long as they are careful with their own keys.
func NamespaceSet(c Context, namespace, key *string, value interface{}, overwrite bool) {
	if c == nil || namespace == nil || key == nil || *namespace == `` || *key == `` {
		return
	}

	if nsval, ok := c.Get(*namespace); !ok {
		newNS := &namespaceType{
			valDict: map[string]interface{}{
				*key: value,
			},
		}

		c.Set(*namespace, newNS)
	} else if ns, stillOk := nsval.(*namespaceType); stillOk {
		ns.Set(*key, value, overwrite)
	}
}

// GetTransactionBrand gets the txbrand field from the context.
func GetTransactionBrand(c Context) (brandID *enum.BrandID) {
	if c != nil {
		if brandVal, nsok := NamespaceGet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Brand.ID)); nsok && brandVal != nil {
			if brandField, ok := brandVal.(*TXField); ok {
				brandID = brandField.Value.(*enum.BrandID)
			}
		}
	}
	return
}

// SetTransactionBrand sets the txbrand field in the context, if it is not already set.
func SetTransactionBrand(c Context, v *enum.BrandID) {
	if c == nil || v == nil {
		return
	}

	foo := string(*v)
	txf := &TXField{
		Key:         enum.TXHeader.Brand.ID.ToIDString(),
		HeaderKey:   enum.TXHeader.Brand.HeaderKey,
		StringValue: &foo,
		Value:       v,
	}

	NamespaceSet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Brand.ID), txf, false)
}

// GetTransactionDomain gets the txdomain field from the context.
func GetTransactionDomain(c Context) (domainID *enum.AccountDomainID) {
	if c != nil {
		if domainVal, nsok := NamespaceGet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Domain.ID)); nsok && domainVal != nil {
			if domainField, ok := domainVal.(*TXField); ok {
				domainID = domainField.Value.(*enum.AccountDomainID)
			}
		}
	}

	return
}

// SetTransactionDomain sets the txdomaihn field in the context, if it is not already set.
func SetTransactionDomain(c Context, v *enum.AccountDomainID) {
	if c == nil || v == nil {
		return
	}

	foo := string(*v)
	txf := &TXField{
		Key:         enum.TXHeader.Domain.ID.ToIDString(),
		HeaderKey:   enum.TXHeader.Domain.HeaderKey,
		StringValue: &foo,
		Value:       v,
	}

	NamespaceSet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Domain.ID), txf, false)
}

// GetTransactionID gets the txid field from the context.
func GetTransactionID(c Context) (txID string) {
	var (
		ok  bool
		ptr *string
	)

	if c != nil {
		if txIDVal, nsok := NamespaceGet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.ID.ID)); nsok && txIDVal != nil {
			if ptr, ok = txIDVal.(*string); !ok || ptr == nil {
				return ``
			}

			txID = *ptr
		}
	}

	return
}

// SetTransactionID sets the txid field in the context, if it is not already set.
// Note that although the func accepts a string value the context value is a *string
// Because all context values must be reference types
func SetTransactionID(c Context, v string) {
	if c == nil || v == `` {
		return
	}

	NamespaceSet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.ID.ID), &v, false)
}

// GetTransactionInstance gets the txinstance field from the context.
func GetTransactionInstance(c Context) (txInstance *string) {
	var (
		ok  bool
		ptr *string
	)

	if c != nil {
		if txInstanceVal, nsok := NamespaceGet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Instance.ID)); nsok && txInstanceVal != nil {
			if ptr, ok = txInstanceVal.(*string); !ok {
				return nil
			}

			txInstance = ptr
		}
	}

	return
}

// SetTransactionInstance sets the txinstance field in the context, if it is not already set.
// Note that although the func acceps a string value the context value is a *string
// Because all context values must be reference types
func SetTransactionInstance(c Context, v string) {
	if c == nil || v == `` {
		return
	}

	NamespaceSet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Instance.ID), &v, false)
}

// GetTransactionIntegrator gets the txintegrator field from the context.
func GetTransactionIntegrator(c Context) (integrator *enum.IntegrationPartnerID) {
	if c != nil {
		if integratorVal, nsok := NamespaceGet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Integrator.ID)); nsok && integratorVal != nil {
			if integratorField, ok := integratorVal.(*TXField); ok {
				integrator = integratorField.Value.(*enum.IntegrationPartnerID)
			}
		}
	}

	return
}

// SetTransactionIntegrator sets the txintegrator field in the context, if it is not already set.
func SetTransactionIntegrator(c Context, v *enum.IntegrationPartnerID) {
	if c == nil || v == nil {
		return
	}

	foo := string(*v)
	txf := &TXField{
		Key:         enum.TXHeader.Integrator.ID.ToIDString(),
		HeaderKey:   enum.TXHeader.Integrator.HeaderKey,
		StringValue: &foo,
		Value:       v,
	}

	NamespaceSet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Integrator.ID), txf, false)
}

// GetTransactionIPAddress gets the txip field from the context.
func GetTransactionIPAddress(c Context) (txIPAddress *string) {
	var (
		ok  bool
		ptr *string
	)

	if c != nil {
		if txIPVal, nsok := NamespaceGet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.IPAddress.ID)); nsok && txIPVal != nil {
			if ptr, ok = txIPVal.(*string); !ok {
				return nil
			}

			txIPAddress = ptr
		}
	}

	return
}

// SetTransactionIPAddress sets the txip field in the context, if it is not already set.
// Note that although the func acceps a string value the context value is a *string
// Because all context values must be reference types
func SetTransactionIPAddress(c Context, v string) {
	if c == nil || v == `` {
		return
	}

	NamespaceSet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.IPAddress.ID), &v, false)
}

// GetTransactionSource gets the txsource (source of business) field of the microservice logger.
func GetTransactionSource(c Context) (source *enum.SourceOfBusinessID) {
	if c != nil {
		if sobVal, nsok := NamespaceGet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Source.ID)); nsok && sobVal != nil {
			if sobField, ok := sobVal.(*TXField); ok {
				source = sobField.Value.(*enum.SourceOfBusinessID)
			}
		}
	}

	return
}

// SetTransactionSource sets the txsource (source of business) field of the microservice logger.
func SetTransactionSource(c Context, v *enum.SourceOfBusinessID) {
	if c == nil || v == nil {
		return
	}

	foo := string(*v)
	txf := &TXField{
		Key:         enum.TXHeader.Source.ID.ToIDString(),
		HeaderKey:   enum.TXHeader.Source.HeaderKey,
		StringValue: &foo,
		Value:       v,
	}

	NamespaceSet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Source.ID), txf, false)
}

// GetTransactionType gets the txtype field from the context.
func GetTransactionType(c Context) (txType *string) {
	var (
		ok  bool
		ptr *string
	)

	if c != nil {
		if txTypeVal, nsok := NamespaceGet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Type.ID)); nsok && txTypeVal != nil {
			if ptr, ok = txTypeVal.(*string); !ok {
				return nil
			}

			txType = ptr
		}
	}

	return
}

// SetTransactionType sets the txtype field in the context, if it is not already set.
// Note that although the func acceps a string value the context value is a *string
// Because all context values must be reference types
func SetTransactionType(c Context, v string) {
	if c == nil || v == `` {
		return
	}

	NamespaceSet(c, &namespaceKeyConfig, (*string)(&enum.TXHeader.Type.ID), &v, false)
}

// WithTimeout returns a microservice context with a timeout and a cancel func
// It is a simple wrapper around the API func of the same name
func WithTimeout(parent context.Context, timeout time.Duration) (Context, context.CancelFunc) {
	c, cancel := context.WithDeadline(parent, time.Now().Add(timeout))
	return New(c), cancel
}

func GetTXFields(c Context) (map[string]TXField, error) {
	return getTXFields(c)
}
