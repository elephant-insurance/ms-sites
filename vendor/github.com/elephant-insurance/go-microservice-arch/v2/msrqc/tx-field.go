package msrqc

import (
	"errors"
	"reflect"

	"github.com/elephant-insurance/enumerations/v2"
)

// TXField represents an Elephant TX-header field
type TXField struct {
	Key           string
	HeaderKey     string
	ReceivedValue *string
	StringValue   *string
	Value         interface{}
}

func (txf *TXField) Equals(otherTXField *TXField) bool {
	if txf == nil && otherTXField == nil {
		return true
	}
	if txf == nil || otherTXField == nil {
		return false
	}

	if txf.Key != otherTXField.Key || txf.HeaderKey != otherTXField.HeaderKey {
		return false
	}

	// These are complicated looking but they are all the same:
	// if either field is nil AND both fields are not nil return FALSE
	// else if one field is not nil and the pointed-to values are different return FALSE
	// Net result: true if both are nil, false if only one is nil, false if the values don't match
	if ((txf.ReceivedValue == nil || otherTXField.ReceivedValue == nil) &&
		!(txf.ReceivedValue == nil && otherTXField.ReceivedValue == nil)) ||
		(txf.ReceivedValue != nil && *txf.ReceivedValue != *otherTXField.ReceivedValue) {
		return false
	}

	if ((txf.StringValue == nil || otherTXField.StringValue == nil) &&
		!(txf.StringValue == nil && otherTXField.StringValue == nil)) ||
		(txf.StringValue != nil && *txf.StringValue != *otherTXField.StringValue) {
		return false
	}

	if !reflect.DeepEqual(txf.Value, otherTXField.Value) {
		return false
	}

	return true
}

func getTXFields(c Context) (map[string]TXField, error) {
	var (
		rtn                              map[string]TXField
		fieldNameMSTransactionBrand      = string(enumerations.TXHeader.Brand.ID)
		fieldNameMSTransactionDomain     = string(enumerations.TXHeader.Domain.ID)
		fieldNameMSTransactionID         = string(enumerations.TXHeader.ID.ID)
		fieldNameMSTransactionInstance   = string(enumerations.TXHeader.Instance.ID)
		fieldNameMSTransactionIntegrator = string(enumerations.TXHeader.Integrator.ID)
		fieldNameMSTransactionIPAddress  = string(enumerations.TXHeader.IPAddress.ID)
		fieldNameMSTransactionSource     = string(enumerations.TXHeader.Source.ID)
		fieldNameMSTransactionType       = string(enumerations.TXHeader.Type.ID)
	)

	if c == nil {
		return nil, errors.New(`nil Context`)
	}

	if val, ok := c.Get(namespaceKeyConfig); ok {
		if ns, stillOk := val.(*namespaceType); stillOk && ns.valDict != nil {
			// we could call Get over and over but that locks the namespace each time
			// instead we lock it once, but must be careful not to call anything else that locks!
			rtn = make(map[string]TXField)
			ns.Lock()
			defer ns.Unlock()

			// TXBrand
			if value, exists := ns.valDict[enumerations.TXHeader.Brand.ID.ToIDString()]; exists {
				if txHeaderVal, stillOK := value.(*TXField); stillOK && txHeaderVal != nil {
					rtn[fieldNameMSTransactionBrand] = *txHeaderVal
				}
			}

			// TXDomain
			if value, exists := ns.valDict[fieldNameMSTransactionDomain]; exists {
				if txHeaderVal, stillOK := value.(*TXField); stillOK && txHeaderVal != nil {
					rtn[fieldNameMSTransactionDomain] = *txHeaderVal
				}
			}

			// TXID
			if value, exists := ns.valDict[fieldNameMSTransactionID]; exists {
				if stringVal, stillOK := value.(*string); stillOK {
					rtn[fieldNameMSTransactionID] = TXField{
						Key:         fieldNameMSTransactionID,
						HeaderKey:   enumerations.TXHeader.ID.HeaderKey,
						StringValue: stringVal,
						Value:       stringVal,
					}
				}
			}

			// TXInstance
			if value, exists := ns.valDict[fieldNameMSTransactionInstance]; exists {
				if stringVal, stillOK := value.(*string); stillOK {
					rtn[fieldNameMSTransactionInstance] = TXField{
						Key:         fieldNameMSTransactionInstance,
						HeaderKey:   enumerations.TXHeader.Instance.HeaderKey,
						StringValue: stringVal,
						Value:       stringVal,
					}
				}
			}

			// TXIntegrator
			if value, exists := ns.valDict[fieldNameMSTransactionIntegrator]; exists {
				if txHeaderVal, stillOK := value.(*TXField); stillOK && txHeaderVal != nil {
					rtn[fieldNameMSTransactionIntegrator] = *txHeaderVal
				}
			}

			// TXIPAdress
			if value, exists := ns.valDict[fieldNameMSTransactionIPAddress]; exists {
				if stringVal, stillOK := value.(*string); stillOK {
					rtn[fieldNameMSTransactionIPAddress] = TXField{
						Key:         fieldNameMSTransactionIPAddress,
						HeaderKey:   enumerations.TXHeader.IPAddress.HeaderKey,
						StringValue: stringVal,
						Value:       stringVal,
					}
				}
			}

			// TXSource
			if value, exists := ns.valDict[fieldNameMSTransactionSource]; exists {
				if txHeaderVal, stillOK := value.(*TXField); stillOK && txHeaderVal != nil {
					rtn[fieldNameMSTransactionSource] = *txHeaderVal
				}
			}

			// TXType
			if value, exists := ns.valDict[fieldNameMSTransactionType]; exists {
				if stringVal, stillOK := value.(*string); stillOK {
					rtn[fieldNameMSTransactionType] = TXField{
						Key:         fieldNameMSTransactionType,
						HeaderKey:   enumerations.TXHeader.Type.HeaderKey,
						StringValue: stringVal,
						Value:       stringVal,
					}
				}
			}
		} else {
			return nil, errors.New(`invalid config namespace`)
		}
	} else {
		return nil, errors.New(`missing config namespace`)
	}

	return rtn, nil
}
