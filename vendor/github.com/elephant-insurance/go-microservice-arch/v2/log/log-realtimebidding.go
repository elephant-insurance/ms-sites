package log

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/elephant-insurance/go-microservice-arch/v2/mbuf"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
)

// RealTimeBidLog is a specialized Log type for recording real-time bidding results.
// RealTimeBidLog is designed to be very simple and does not support many of the features of the microservice log, such as levels.

type RealTimeBidLog struct {
	// appName is required and can't change, so we save it here
	appName string
	// ImmediateHandler is a message handler that runs on the calling thread.
	// This is primarily for logging to the console in development.
	// Set this to nil to suppress logging to the console.
	ImmediateHandler RTBMessageHandler
	// Relay is a message relay that forwards our messages to another service.
	// Set this to nil to avoid sending anything.
	Relay mbuf.MessageRelay
	// settings keeps a copy of the values we were initialized with
	// this saves copying each setting to a struct field
	settings RTBSettings
}

// RTBSettings is a log.Settings struct for RTBLogs.
// A type alias lets us share code, but also to differentiate the types where needed,
// such as when overriding log settings with environment variables.
type RTBSettings Settings

// RTBLogType is an Azure log type (table name) for an RTBLog.
// Using a type alias makes it a bit more difficult for a dev to accidentally mis-name a table.
type RTBLogType string

// We give each app its own table name to avoid JSON conflicts.
const (
	RTBEverQuoteLogType  RTBLogType = `rtb_everquote_log`
	RTBMediaAlphaLogType RTBLogType = `rtb_mediaalpha_log`
	RTBQuinStreetLogType RTBLogType = `rtb_quinstreet_log`
	RTBQuoteWizardBid    RTBLogType = `rtb_quotewizard_log`
)

// GetEnvironmentSetting implements the Overridable interface for the configuration package.
// Given a field name, it returns the name of the corresponding environment variable.
// This makes it possible to override an RTBSettings field using an environment variable.
func (l RTBSettings) GetEnvironmentSetting(fieldName string) string {
	return EnvironmentSettingPrefixRTB + fieldName
}

// NewRealTimeBiddingLog takes its required settings and validates them.
// It sets up the handlers for immediate (console) logging and for Azure logging.
func NewRealTimeBiddingLog(appName string, logType RTBLogType, s *RTBSettings) (*RealTimeBidLog, error) {
	if err := validateForRTB(appName, logType, s); err != nil {
		return nil, err
	}

	rtn := &RealTimeBidLog{
		appName:  appName,
		settings: *s,
	}

	// do we need a console handler? Default (nil) to YES
	if s.ConsoleMode == nil || *s.ConsoleMode == ConsoleModeStandardOut {
		// for now just create a simple default handler
		h := &rtbMessageHandler{}
		rtn.ImmediateHandler = h
	}

	r := mbuf.NewAzureLogAnalyticsMessageRelay(*s.AzureWorkspaceID, *s.AzureSharedKey, string(logType), azureRTBLogTimeField, s.RelaySettings())
	rtn.Relay = r

	return rtn, nil
}

// Write takes an RTBLogEntry and submits it to the RealTimeBidLog's configured
// handlers for immediate and remote logging.
func (l *RealTimeBidLog) Write(entry *RTBLogEntry) {
	if l == nil || entry == nil {
		return
	}

	entry.Timestamp = uf.Pointer.ToNow()
	entry.AppName = uf.Pointer.ToString(l.appName)

	if l.ImmediateHandler != nil {
		l.ImmediateHandler.HandleMessage(entry)
	}

	if l.Relay != nil {
		l.Relay.Add(entry)
	}
}

// validateForRTB performs a basic sanity-check of submitted startup parameters.
func validateForRTB(appName string, logType RTBLogType, s *RTBSettings) error {
	if s == nil {
		return errors.New(ErrorNilSettings)
	}

	if appName == "" {
		return errors.New(ErrorEmptyInstanceName)
	}

	if logType == "" {
		return errors.New(ErrorEmptyLogType)
	}

	// can't have an RTB log without Azure settings
	if s.AzureSharedKey == nil || *s.AzureSharedKey == "" || s.AzureWorkspaceID == nil || *s.AzureWorkspaceID == "" {
		return errors.New(ErrorAzureSettingsMissing)
	}

	return nil
}

// RelaySettings is just a convenient little hack to generate mbuf settings for the remote handler.
// This is already implemented for log.Settings, and RTBSettings is just a type alias for log.Settings.
func (s *RTBSettings) RelaySettings() *mbuf.Settings {
	if s == nil {
		return nil
	}

	// cast our "RTBSettings" back to its true form and run the same method
	sobj := Settings(*s)

	return (&sobj).RelaySettings()
}

// rtbMessageHandler is the default immediate (console) handler for RealTimeBidLog.
type rtbMessageHandler struct{}

func (h *rtbMessageHandler) HandleMessage(msg *RTBLogEntry) {
	// TODO: make this nice for local debugging
	// we have a strong type here, so we can do whatever we want
	// for now, just marshal it and shoot it to stdout
	bytes, err := json.Marshal(msg)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
