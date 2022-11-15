package cfg

import enum "github.com/elephant-insurance/enumerations/v2"

// configurator is an interface type for wrapping RequiredConfig.
// By forcing devs to submit a configurator instead of a simple config value, we make it more difficult to sidestep environment checking.

// !!! IMPORTANT !!!
// THIS INTERFACE MUST ALWAYS RETURN VALUES, NOT REFERENCES
// Returning a reference could allow a developer to overwrite a RequiredConfig setting.
// For example, assume GetEnvironment returned a pointer instead of a value:
// 		appEnv := rc.GetEnvironment()
// 		(*appEnv) = enum.ServiceEnvironment.Development.ID // this updates the target, so the app environment is now DEV
type Configurator interface {
	GetAllowCompressedRequests() bool
	GetAllowedHeaders() []string
	GetAllowedMethods() []string
	GetAllowedOrigins() []string
	GetAppAbbreviation() string
	GetEnvironment() enum.ServiceEnvironmentID
	GetExposedHeaders() []string
	GetInstanceName() string
	GetListenPort() string
	GetLogLevel() enum.LogLevelID
	GetServiceID() enum.ServiceID
}
