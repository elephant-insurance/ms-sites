package cfg

import (
	"fmt"
	"strings"

	enum "github.com/elephant-insurance/enumerations/v2"
)

type testConfigurator struct {
	allowCompressedRequests bool
	allowedMethods          []string
	allowedHeaders          []string
	allowedOrigins          []string
	appAbbreviation         string
	environment             enum.ServiceEnvironmentID
	exposedHeaders          []string
	instanceName            string
	listenPort              string
	logLevel                enum.LogLevelID
	serviceID               enum.ServiceID
}

func NewTestConfigurator(rc RequiredConfig) Configurator {
	verifyEnvironment()
	ah, am, ao, eh := []string{}, []string{}, []string{}, []string{}
	lp := defaultListenPort

	if len(rc.AllowedHeaders) > 0 {
		ah = strings.Split(rc.AllowedHeaders, ",")
	}
	if len(rc.AllowedMethods) > 0 {
		am = strings.Split(rc.AllowedMethods, ",")
	}
	if len(rc.AllowedOrigins) > 0 {
		ao = strings.Split(rc.AllowedOrigins, ",")
	}
	if len(rc.ExposedHeaders) > 0 {
		eh = strings.Split(rc.ExposedHeaders, ",")
	}

	if rc.DevTestListenPort != nil && *rc.DevTestListenPort != 0 {
		lp = *rc.DevTestListenPort
	}

	return &testConfigurator{
		allowedHeaders:  ah,
		allowedMethods:  am,
		allowedOrigins:  ao,
		appAbbreviation: rc.AppAbbreviation,
		environment:     rc.Environment,
		exposedHeaders:  eh,
		instanceName:    rc.InstanceName,
		listenPort:      fmt.Sprintf(`:%v`, lp),
		logLevel:        rc.LogLevel,
		serviceID:       enum.Service.Test.ID,
	}
}

func (tc *testConfigurator) GetAllowCompressedRequests() bool {
	verifyEnvironment()
	return tc.allowCompressedRequests
}

func (tc *testConfigurator) GetAllowedHeaders() []string {
	verifyEnvironment()
	return tc.allowedHeaders
}

func (tc *testConfigurator) GetAllowedMethods() []string {
	verifyEnvironment()
	return tc.allowedMethods
}

func (tc *testConfigurator) GetAllowedOrigins() []string {
	verifyEnvironment()
	return tc.allowedOrigins
}

func (tc *testConfigurator) GetAppAbbreviation() string {
	verifyEnvironment()
	return tc.appAbbreviation
}

func (tc *testConfigurator) GetEnvironment() enum.ServiceEnvironmentID {
	verifyEnvironment()
	return tc.environment
}

func (tc *testConfigurator) GetExposedHeaders() []string {
	verifyEnvironment()
	return tc.exposedHeaders
}

func (tc *testConfigurator) GetInstanceName() string {
	verifyEnvironment()
	return tc.instanceName
}

func (tc *testConfigurator) GetServiceID() enum.ServiceID {
	verifyEnvironment()
	return tc.serviceID
}

func (tc *testConfigurator) GetListenPort() string {
	verifyEnvironment()
	return tc.listenPort
}

func (tc *testConfigurator) GetLogLevel() enum.LogLevelID {
	verifyEnvironment()
	return tc.logLevel
}

// verifyEnvironment crashes if we have loaded a real RequiredConfig
func verifyEnvironment() {
	if requiredConfigLoaded {
		cfgCrasher.Panic(`attempted to load test config after real config has been loaded`)
	}
}
