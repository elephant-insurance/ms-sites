package log

import (
	"context"
	"net/http"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/timing"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
)

type MicroserviceLogger interface {
	//
	// Logging Methods
	//

	// Trace logs a message at Trace level, if enabled.
	// If args are specified, this is equivalent to WithConsoleFields(args).Trace(msg).
	Trace(msg string, args ...interface{}) MicroserviceLogger

	// Debug logs a message at Debug level, if enabled.
	// If args are specified, this is equivalent to WithConsoleFields(args).Debug(msg).
	Debug(msg string, args ...interface{}) MicroserviceLogger

	// Info logs a message at Info level, if enabled.
	// If args are specified, this is equivalent to WithConsoleFields(args).Info(msg).
	Info(msg string, args ...interface{}) MicroserviceLogger

	// Warn logs a message at Warn level, if enabled.
	// If args are specified, this is equivalent to WithConsoleFields(args).Warn(msg).
	Warn(msg string, args ...interface{}) MicroserviceLogger

	// Error logs a message at Error level, if enabled.
	// If args are specified, this is equivalent to WithConsoleFields(args).Error(msg).
	Error(msg string, args ...interface{}) MicroserviceLogger

	// Fatal logs a message at Fatal level, then exits the application.
	// If args are specified, this is equivalent to WithConsoleFields(args).Fatal(msg).
	Fatal(msg string, args ...interface{}) MicroserviceLogger

	// Panic logs a message at Panic level, then issues a system panic.
	// If args are specified, this is equivalent to WithConsoleFields(args).Panic(msg).
	Panic(msg string, args ...interface{}) MicroserviceLogger

	//
	// Conditional Methods
	//

	// If disables the logwriter temporarily if condition is false
	If(condition bool) MicroserviceLogger

	// IfError disables the logwriter temporarily if err is nil
	IfError(err error) MicroserviceLogger

	// IfError disables the logwriter temporarily if evaluator returns false
	IfEval(evaluator Evaluator) MicroserviceLogger

	// Else re-activates a logwriter de-activated by a failed If
	Else() MicroserviceLogger

	// ElseIf re-activates a logwriter de-activated by a failed If, if and only if condition is true
	ElseIf(condition bool) MicroserviceLogger

	// SetContext permanently sets the context for this MicroserviceLogger
	SetContext(c context.Context) MicroserviceLogger

	/*
				Log field setting methods

				Interface Rules:
				* Any method starting with "Set" sets a property permanently (for the lifetime of the logger or the context).
				* Any method starting with "With" sets a property until the next time a logging method is called on the logger.
				* Any method ending with "Field(s)" sets a property that will NOT be sent to remote handlers (e.g., Loggly, Azure).
				* Every Set/With method that does not end in "Field(s)" sets a property that WILL be sent to remote handlers.
				* Every logging method (Debug, Info, etc.) that accepts optional field arguments works like WithConsoleFields(args...).

				Context scope:
				* Only "Set" methods can set values in the context.
				* Once set in a context, all loggers using that context will include it whenever they write an entry.
				* Once set in a context, a value will remain set until it is explicitly overwritten.

				Transaction Header (TX) fields:
		        * Any method starting with "SetTransaction" sets a Transaction Header (TX) field.
		        * All TX fields always have context scope.
		        * Once set, a TX field cannot be overwritten.
		        * The GLOG package handles setting TX fields from headers in received messages, and setting headers on outgoing messages.
				* All received values are set (by GLOG) before any app-specific code runs, so app code cannot change received values.

	*/

	// Microservice Log Field SET (persistent field) funcs
	// These fields remain set until overwritten.
	// Most may also be added to the context by adding bool true as an optional second parameter.
	// All of these fields are sent to log relays.

	// SetAccountID sets the account field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetAccountID(v string, addToContext ...interface{}) MicroserviceLogger

	// SetBusinessType sets the business type for the transaction that we are logging.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetBusinessType(v *enum.BusinessTypeID, addToContext ...interface{}) MicroserviceLogger

	// SetChannel sets the channel field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetChannel(v string, addToContext ...interface{}) MicroserviceLogger

	// SetCode sets the code field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetCode(v string, addToContext ...interface{}) MicroserviceLogger

	// SetCount sets the count field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetCount(v int, addToContext ...interface{}) MicroserviceLogger

	// SetDate sets the Date field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetDate(v uf.Datable, addToContext ...interface{}) MicroserviceLogger

	// SetDetail sets the "detail" string field, which may be used for logging any sort of relevant data.
	// The length of this field is capped at 1000 characters.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetDetail(v string, addToContext ...interface{}) MicroserviceLogger

	// SetDriver sets the driver field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetDriver(v string, addToContext ...interface{}) MicroserviceLogger

	// SetFunction sets the func field of the microservice log.
	// Since the name of the running function is by definition function scoped, this value cannot be added to the context.
	SetFunction(v string) MicroserviceLogger

	// SetID sets the id field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetID(v string, addToContext ...interface{}) MicroserviceLogger

	// SetIPAddress sets the ip-address field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetIPAddress(v string, addToContext ...interface{}) MicroserviceLogger

	// SetJobID sets the job field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetJobID(v string, addToContext ...interface{}) MicroserviceLogger

	// SetName sets the name field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetName(v string, addToContext ...interface{}) MicroserviceLogger

	// SetNumber sets the number field of the microservice log.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetNumber(v int, addToContext ...interface{}) MicroserviceLogger

	// SetPolicyID sets the policy field of the microservice logger.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetPolicyID(v string, addToContext ...interface{}) MicroserviceLogger

	// SetPublicID sets the public-id field of the microservice logger.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetPublicID(v string, addToContext ...interface{}) MicroserviceLogger

	// SetQuoteID sets the quote-id field of the microservice logger.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetQuoteID(v string, addToContext ...interface{}) MicroserviceLogger

	// SetQuoteNumber sets the quote-number field of the microservice logger.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetQuoteNumber(v string, addToContext ...interface{}) MicroserviceLogger

	// SetState sets the state (of the USA) field of the microservice logger.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetState(v *enum.StateID, addToContext ...interface{}) MicroserviceLogger

	// SetTransactionBrand sets the txbrand field of the microservice logger.
	// Because our business rules require us to forward its value, this field is added to the MicroserviceLogger's context, if not nil, automatically.
	SetTransactionBrand(v *enum.BrandID) MicroserviceLogger

	// SetTransactionDomain sets the txdomaihn field of the microservice logger.
	// Because our business rules require us to forward its value, this field is added to the MicroserviceLogger's context, if not nil, automatically.
	SetTransactionDomain(v *enum.AccountDomainID) MicroserviceLogger

	// SetTransactionID sets the txid field of the microservice logger.
	// Because our business rules require us to forward its value, this field is added to the MicroserviceLogger's context, if not nil, automatically.
	SetTransactionID(v string) MicroserviceLogger

	// SetTransactionIntegrator sets the txintegrator field of the microservice logger.
	// Because our business rules require us to forward its value, this field is added to the MicroserviceLogger's context, if not nil, automatically.
	SetTransactionIntegrator(v *enum.IntegrationPartnerID) MicroserviceLogger

	// SetTransactionSource sets the txsource (source of business) field of the microservice logger.
	// Because our business rules require us to forward its value, this field is added to the MicroserviceLogger's context, if not nil, automatically.
	SetTransactionSource(v *enum.SourceOfBusinessID) MicroserviceLogger

	// SetTransactionType sets the txtype field of the microservice logger.
	// Because our business rules require us to forward its value, this field is added to the MicroserviceLogger's context, if not nil, automatically.
	SetTransactionType(v string) MicroserviceLogger

	// SetVehicle sets the vehicle field of the microservice logger.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetVehicle(v string, addToContext ...interface{}) MicroserviceLogger

	// SetVIN sets the vin field of the microservice logger.
	// If the optional second parameter is bool true, then the value is also added to the MicroserviceLogger's context, if not nil.
	SetVIN(v string, addToContext ...interface{}) MicroserviceLogger

	// Microservice Log Field WITH (ephemeral field) funcs.
	// All of these fields are sent to log relays.
	// They are only written once, and erased after the first Logging function that fires.
	// These fields may not be added to the context directly.

	// WithCommonErrorCode sets the error-code field of the microservice log
	WithCommonErrorCode(errCode *enum.CommonErrorCodeID) MicroserviceLogger

	// WithElapsedMicroseconds sets the microseconds field of the microservice log
	WithElapsedMicroseconds(v int64) MicroserviceLogger

	// WithError sets the err field of the microservice log
	WithError(err error) MicroserviceLogger

	// WithHTTPMethod sets the http-method field of the microservice log
	WithHTTPMethod(v string) MicroserviceLogger

	// WithHTTPStatus sets the http-status field of the microservice log
	WithHTTPStatus(v int) MicroserviceLogger

	// WithStack generates a stack trace and sets the stack field of the microservice log
	WithStack() MicroserviceLogger

	// WithURL sets the url field of the microservice log
	WithURL(v string) MicroserviceLogger

	// Debug Log Field funcs
	// These do NOT get sent to relays, just printed to the console for debugging.
	// You can log just about anything you like here.

	// SetField sets a deubg log value in the logger's context,
	// if the context is not nil and the value for key is not already set.
	// If the context is nil, then it will set the value on a matching Set property, if any.
	// Failing this, it will set it as an ephemeral value with WithConsoleField.
	SetField(key string, value interface{}) MicroserviceLogger

	// WithConsoleFields(arg1, arg2, arg3, arg4,...) is equivalent to WithConsoleField(arg1, arg2).WithConsoleField(arg3, arg4)...
	// Each odd-numbered arg must be able to be cast to a string, or it will be ignored.
	WithConsoleFields(args ...interface{}) MicroserviceLogger

	// WithConsoleField sets an arbitrary ephemeral field on the MicroserviceLogger.
	WithConsoleField(key string, value interface{}, addToContext ...interface{}) MicroserviceLogger

	// GetTransactionID returns the transaction ID stored in this logger's context
	GetTransactionID() string

	// Event Handling

	// ClearEvents empties this logger's list of events and returns any that it had
	ClearEvents() []*MSEventLogEntry

	// WithEvent adds an event to be sent to the microservice event log.
	// Events are strictly-formatted messages optimized for automation.
	WithEvent(event *uf.Event) MicroserviceLogger

	// ServiceResponseSuccess is a one-stop logger for service responses
	// It will return true IF the following are both TRUE:
	// 1. err is nil.
	// 2. resp is not nil.
	// If either of these conditions fails, we log both an event and a log message at the Error level.
	// Use this version when the service may return an error status code (>300) under normal conditions (not a "real" error that we need to deal with).
	ServiceResponseSuccess(resp *http.Response, err error, respTime *timing.Timing, target *enum.ServiceID, description string) bool

	// ServiceResponseStrictSuccess works exactly like ServiceResponseSuccess
	// except that it also checks the response.StatusCode.
	// If the status code indicates an error (>300), we log an error and return false.
	// Use this version when only a 2xx response code is acceptable from the service.
	ServiceResponseStrictSuccess(resp *http.Response, err error, respTime *timing.Timing, target *enum.ServiceID, description string) bool

	// ServiceResponseWithoutError is a one-stop logger for service responses
	// It will log the response and return true if err is nil
	ServiceResponseWithoutError(resp *http.Response, err error, respTime *timing.Timing, target *enum.ServiceID, description string) bool

	/* TODO? Please let me know if you think of more!
	WithConsoleFieldBody(body string) MicroserviceLogger
	WithConsoleFieldCoverageCode(code string) MicroserviceLogger
	WithConsoleFieldErrorCode(code string) MicroserviceLogger
	WithConsoleFieldIndex(index int) MicroserviceLogger
	WithConsoleFieldKey(key string) MicroserviceLogger
	WithConsoleFieldResponse(resp string) MicroserviceLogger
	WithConsoleFieldSelectedValue(val string) MicroserviceLogger
	WithConsoleFieldTerm(term string) MicroserviceLogger
	WithConsoleFieldValidationCode(code string) MicroserviceLogger
	WithConsoleFieldValue(value string) MicroserviceLogger
	*/
}
