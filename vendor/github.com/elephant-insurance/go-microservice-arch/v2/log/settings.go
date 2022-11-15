package log

import (
	"errors"
	"os"
	"strings"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/mbuf"
)

// Settings is a set of initialization settings for a Log
// Settings allows for a single Console logger, a single Azure message relay, and a single Loggly message relay
// This is probably enough for most purposes, but additional relays may be added at runtime
type Settings struct {
	// AZURE Settings
	// If BOTH AzureSharedKey and AzureWorkspaceID are specified and valid, we will send messages to Azure

	// AzureSharedKey is the Azure shared key for sending direct to Azure, if so configured
	AzureSharedKey *string `yaml:"AzureSharedKey" config:"optional"`
	// AzureWorkspaceID is the Azure WorkspaceID for sending direct to Azure, if so configured
	AzureWorkspaceID *string `yaml:"AzureWorkspaceID" config:"optional"`
	// AzureLogLevel sets the level for the Azure relay, if it is different from the global log level
	AzureLogLevel *enum.LogLevelID `yaml:"AzureLogLevel" config:"optional"`

	// CONSOLE Settings

	// ConsoleMode determines whether the package logs to stdout
	// Valid values are: ConsoleModeNone, ConsoleModeStandardOut, and ConsoleModeSendOrStandardOut (default)
	ConsoleMode *ConsoleMode `yaml:"ConsoleMode" config:"optional"`
	// ConsoleLogLevel sets the level for the console, if it is different from the global log level
	ConsoleLogLevel *enum.LogLevelID `yaml:"ConsoleLogLevel" config:"optional"`

	// LOGGLY Settings
	// If LogglyToken is specified and valid, we will attempt to send to Loggly

	// LogglyToken is the token to use for sending direct to Loggly, if we are so configured
	LogglyToken *string `yaml:"LogglyToken" config:"optional"`
	// LogglyLogLevel sets the level for the Loggly relay, if it is different from the global log level
	LogglyLogLevel *enum.LogLevelID `yaml:"LogglyLogLevel" config:"optional"`

	// RELAY (shared) Settings
	// These settings are applied to both the Azure and Loggly relays, if these are configured

	// ClientTimeoutSecs is how many seconds to allow the log relayer to respond to log requests
	// We are using Docker networking, so this can be very fast if desired
	ClientTimeoutSecs *int `yaml:"ClientTimeoutSecs" config:"optional"`
	// Environment keeps track of what environment we are running in
	Environment *enum.ServiceEnvironmentID `yaml:"Environment" config:"optional"`
	// FailuresToShutdown is how many times we have to fail to send to the relay
	// before we stop trying to send for a while
	FailuresToShutdown *int `yaml:"FailuresToShutdown" config:"optional"`
	// HostName
	HostName *string `yaml:"HostName" config:"optional"`
	// HostURL, if not nil, will set the package to relay to this URL
	// if nil, the package will not try to relay log messages
	// Valid values include: SendDirectToAzure SendDirectToLoggly SendToBoth SendToNone
	HostURL *string `yaml:"HostURL" config:"optional"`
	// InstanceName sets the name of this instance of the application, if it is different from the Service Name.
	// The InstanceName field is optional, and if omitted it defaults to the string value of ServiceID, which is required.
	// This field is used primarily for routing, to distinguish different instances of the same service, where needed.
	InstanceName *string `yaml:"InstanceName" config:"optional"`
	// Level sets the default level for this Log
	// Relays and console handlers may have different log levels, but if those are not set
	// they will default to this value
	Level *enum.LogLevelID `yaml:"Level" config:"optional"`
	// MaxBufferLength is how many log messages to buffer before sending
	MaxBufferLength *int `yaml:"MaxBufferLength" config:"optional"`
	// MaxRetries is the max number of times we try to send a buffer of messages after a network error
	MaxRetries *int `yaml:"MaxRetries" config:"optional"`
	// PollIntervalSecs controls how often we send all available messages to the relay
	PollIntervalSecs *int `yaml:"PollIntervalSecs" config:"optional"`
	// RetryWaitSecs is how many seconds to wait to retry a failed connection
	RetryWaitSecs *int `yaml:"RetryWaitSecs" config:"optional"`
	// ServiceID sets the app field for this Log
	// This field references the name of the application itself, which is also the name of the GitHub repo for the app,
	// This is not necessarily the same as the name of the running instance, which is used for routing
	ServiceID *enum.ServiceID `yaml:"ServiceID" config:"optional"`
	// Shutdown1DurationSecs is how many seconds to atop sending the first time the
	// relay is shut down due to connection failures
	Shutdown1DurationSecs *int `yaml:"Shutdown1DurationSecs" config:"optional"`
	// Shutdown2DurationSecs is how many seconds to atop sending the second time the
	// relay is shut down due to connection failures
	Shutdown2DurationSecs *int `yaml:"Shutdown2DurationSecs" config:"optional"`
	// Shutdown3DurationSecs is how many seconds to atop sending the third and subsequent times the
	// relay is shut down due to connection failures
	Shutdown3DurationSecs *int `yaml:"Shutdown3DurationSecs" config:"optional"`
}

// ForTesting fills out required elements of a Settings object for quick and simple testing
func (s *Settings) ForTesting() *Settings {
	om := ConsoleModeStandardOut
	s.ConsoleMode = &om
	s.AzureSharedKey = nil // make sure we don't try to send to Azure during testing
	s.LogglyToken = nil    // make sure we don't try to send to Loggly during testing
	s.HostURL = nil

	return s
}

// GetEnvironmentSetting implements the Overridable interface for the configuration package
// Given a field name, it returns the name of the corresponding environment variable
func (s Settings) GetEnvironmentSetting(fieldName string) string {
	return EnvironmentSettingPrefixMSL + fieldName
}

func defaultSettings() Settings {
	hostname, _ := os.Hostname()
	return Settings{
		HostName: &hostname,
	}
}

// validate checks submitted settings for internal consistency
func (s *Settings) validate() error {
	if s == nil {
		return nil
	}

	errs := []string{}

	var url string
	if s.HostURL != nil && *s.HostURL == `` {
		url = *s.HostURL
	}

	azureIntended := s.AzureLogLevel != nil || s.AzureSharedKey != nil || s.AzureWorkspaceID != nil || url == SettingSendDirectToAzure || url == SettingSendToBoth
	logglyIntended := s.LogglyLogLevel != nil || s.LogglyToken != nil || url == SettingSendDirectToLoggly || url == SettingSendToBoth
	noneIntended := url == SettingSendToNone

	if azureIntended && (s.AzureSharedKey == nil || s.AzureWorkspaceID == nil) {
		errs = append(errs, `logging to Azure requires both a shared key and a workspace id`)
	}

	if logglyIntended && s.LogglyToken == nil {
		errs = append(errs, `logging to Loggly requires a Loggly key`)
	}

	if noneIntended && (azureIntended || logglyIntended) {
		errs = append(errs, `inconsistent log relay settings`)
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, `,`))
	}

	return nil
}

func (s *Settings) RelaySettings() *mbuf.Settings {
	if s == nil {
		return nil
	}

	rtn := &mbuf.Settings{
		ClientTimeoutSecs: s.ClientTimeoutSecs,
		MaxLength:         s.MaxBufferLength,
		MaxRetries:        s.MaxRetries,
		PollIntervalSecs:  s.PollIntervalSecs,
		RetryWaitSeconds:  s.RetryWaitSecs,
	}

	if s.HostURL != nil {
		rtn.RelayURL = *s.HostURL
	}

	return rtn
}

/*

TODO:

* Test intialization
* Make sure sendorstandardout works

* Add ringbuffer to MessageRelay for better network stats

*/
