package cfg

import enum "github.com/elephant-insurance/enumerations/v2"

// RequiredConfig should be added as an anonymous child struct of
// the main configuration struct for the app. This ensures that
// certain basic settings are always available
// NOTE: These fields have been permanently moved to the Dignostics.Settings struct
// and should be set there, instead: Branch, Build, Commit, DateBuilt, ImageTag, Version
type RequiredConfig struct {
	AllowCompressedRequests *bool                     `yaml:"AllowCompressedRequests" config:"mustoverride,optional"`
	AllowedHeaders          string                    `yaml:"AllowedHeaders" config:"optional"`
	AllowedMethods          string                    `yaml:"AllowedMethods" config:"optional"`
	AllowedOrigins          string                    `yaml:"AllowedOrigins"`
	AppAbbreviation         string                    `yaml:"AppAbbreviation" config:"public"`
	DevTestListenPort       *int                      `yaml:"DevTestListenPort" config:"optional"`
	Environment             enum.ServiceEnvironmentID `yaml:"Environment" config:"public,mustoverride"`
	ExposedHeaders          string                    `yaml:"ExposedHeaders" config:"optional"`
	// InstanceName is the running name of this instance, which determines its name for routing.
	// Default: same as ServiceID
	InstanceName       string          `yaml:"InstanceName" config:"public,optional"`
	LogLevel           enum.LogLevelID `yaml:"LogLevel" config:"public"`
	OverrideConfigPath string          `config:"final,optional" yaml:"OverrideConfigPath"`
	// ServiceID uniquely identifies the name of the service, which corresponds to a GitHub repository
	ServiceID enum.ServiceID `yaml:"ServiceID"`
}

// These variables hold the values that we loaded in at startup. Regardless of what settings
// a RequiredConfig may itself have, it will always return these startup values.
// This makes it infeasible to spoof a configuration by simply creating a RequiredConfig inline.
var (
	requiredAllowCompressedRequests bool
	requiredAllowedHeaders          []string
	requiredAllowedMethods          []string
	requiredAllowedOrigins          []string
	requiredAppAbbreviation         string
	requiredConfigLoaded            = false
	requiredEnvironment             enum.ServiceEnvironmentID
	requiredExposedHeaders          []string
	requiredInstanceName            string
	requiredListenPort              string
	requiredLogLevel                enum.LogLevelID
	requiredServiceID               enum.ServiceID
)

func (otc RequiredConfig) GetEnvironmentSetting(fieldName string) string {
	return EnvironmentVariableNamePrefix + fieldName
}

func (otc RequiredConfig) GetAllowCompressedRequests() bool {
	return requiredAllowCompressedRequests
}

func (otc RequiredConfig) GetAllowedHeaders() []string {
	if requiredAllowedHeaders == nil {
		requiredAllowedHeaders = []string{}
	}

	return requiredAllowedHeaders
}

func (otc RequiredConfig) GetAllowedMethods() []string {
	if requiredAllowedMethods == nil {
		requiredAllowedMethods = []string{}
	}

	return requiredAllowedMethods
}

func (otc RequiredConfig) GetAllowedOrigins() []string {
	if requiredAllowedOrigins == nil {
		requiredAllowedOrigins = []string{}
	}

	return requiredAllowedOrigins
}

func (otc RequiredConfig) GetAppAbbreviation() string {
	return requiredAppAbbreviation
}

func (otc RequiredConfig) GetEnvironment() enum.ServiceEnvironmentID {
	return requiredEnvironment
}

func (otc RequiredConfig) GetExposedHeaders() []string {
	if requiredExposedHeaders == nil {
		requiredExposedHeaders = []string{}
	}

	return requiredExposedHeaders
}

func (otc RequiredConfig) GetInstanceName() string {
	return requiredInstanceName
}

func (otc RequiredConfig) GetListenPort() string {
	return requiredListenPort
}

func (otc RequiredConfig) GetLogLevel() enum.LogLevelID {
	return requiredLogLevel
}

func (otc RequiredConfig) GetServiceID() enum.ServiceID {
	return requiredServiceID
}
