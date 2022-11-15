package cfg

import (
	"github.com/elephant-insurance/go-microservice-arch/v2/cfg"
	"github.com/elephant-insurance/go-microservice-arch/v2/dig"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/sec"
)

// Config for the ms-sites web service
var Config *AppConfig

// Initialize loads the configuration
func Initialize(testConfig *AppConfig) {
	if testConfig != nil {
		Config = testConfig
		return
	}

	Config = &AppConfig{}
	cfg.LoadConfig(Config)

	// this needs to be a pointer:
	cfg.Validate(Config)
}

// AppConfig contains all the config information we load at startup
type AppConfig struct {
	cfg.RequiredConfig `yaml:"RequiredConfig"`
	Diagnostics        dig.Settings `yaml:"Diagnostics" config:"public"`
	Logging            log.Settings `yaml:"Logging"`
	Security           sec.Settings `yaml:"Security"`
	StorageAccountName string       `yaml:"StorageAccountName"`
	BlobContainer      string       `yaml:"BlobContainer"`
	StorageAccountKey  string       `yaml:"StorageAccountKey"`
}

func (config *AppConfig) PreValidate() []string {
	// Pre-validate here, if you need to
	return []string{}
}

func (config *AppConfig) PostValidate(previousErrors []string) []string {
	// Post-validate here, if you need to

	return previousErrors
}
