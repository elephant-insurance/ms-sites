package mbuf

import (
	"errors"
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
)

// NewWithDefaults sets up a simple JSON MessageRelay with all default values.
func NewWithDefaults(hostURL string) MessageRelay {
	rtn := newMBUF(&Settings{RelayURL: hostURL})
	rtn.start()

	return rtn
}

// New creates a new message relay using the settings supplied.
// New will return nil if Settings is nil or RelayURL is empty.
func New(s *Settings) MessageRelay {
	if s == nil || s.RelayURL == "" {
		return nil
	}

	rtn := newMBUF(s)
	rtn.start()

	return rtn
}

// NewAzureLogAnalyticsMessageRelay takes all the fields required for building a complete Azure Log Analytics MessageRelay.
func NewAzureLogAnalyticsMessageRelay(workspaceID, sharedKey, logType, timeStampField string, s *Settings) MessageRelay {
	var settings Settings
	if s == nil {
		settings = Settings{}
	} else {
		settings = *s
	}

	settings.Marshaler = newAzureLogMarshaler(workspaceID, sharedKey, logType, timeStampField)

	rtn := newMBUF(&settings)
	rtn.start()

	return rtn
}

// NewAzureServiceBusMessageRelay takes all the fields required for building a complete Azure Service Bus MessageRelay.
func NewAzureServiceBusMessageRelay(s *AzureServiceBusSettings) (MessageRelay, error) {
	if s == nil || s.ResourceName == `` || s.QueueOrTopic == `` || s.SecurityKeyName == `` || s.SecurityKey == `` {
		return nil, errors.New(`one or more empty args to NewAzureSerivceBusMessageRelay`)
	}

	var settings Settings
	if s.MBUFSettings == nil {
		settings = Settings{}
	} else {
		settings = *s.MBUFSettings
	}

	var ttl *time.Duration
	if s.TTLSeconds != nil && *s.TTLSeconds > 0 {
		ttld := time.Second * time.Duration(*s.TTLSeconds)
		ttl = &ttld
	}

	marshlr, merr := newAzureServiceBusMarshaler(s.ResourceName, s.QueueOrTopic, s.SecurityKeyName, s.SecurityKey, ttl)
	if merr != nil {
		return nil, merr
	}
	settings.Marshaler = marshlr

	rtn := newMBUF(&settings)
	rtn.start()

	return rtn, nil
}

// NewAzureLogAnalyticsMessageRelay takes all the fields required for building a complete Loggly MessageRelay.
func NewLogglyMessageRelay(logglyKey, appName string, env enum.ServiceEnvironmentID, s *Settings) MessageRelay {
	var settings Settings
	if s == nil {
		settings = Settings{}
	} else {
		settings = *s
	}

	settings.Marshaler = newLogglyHTTPMarshaller(logglyKey, appName, env)

	rtn := newMBUF(&settings)
	rtn.start()

	return rtn
}
