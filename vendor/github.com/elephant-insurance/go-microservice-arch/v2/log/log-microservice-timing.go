package log

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	olog "log"
	"os"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/mbuf"
)

// MicroserviceTimingLog is a specialized Log type for recording diagnostic timings.
// MicroserviceTimingLog is designed to be very simple and does not support many of the features of the microservice log, such as levels.

type MicroserviceTimingLog struct {
	// instanceName is required and can't change, so we save it here
	instanceName string
	// hostName is required and can't change, so we save it here
	hostName string
	// Relay is a message relay that forwards our messages to another service.
	// Set this to nil to avoid sending anything.
	Relay mbuf.MessageRelay
	// serviceID keeps track of the actual program that is running, as well as the correct log table names
	serviceID *enum.ServiceID
	// settings keeps a copy of the values we were initialized with
	// this saves copying each setting to a struct field
	settings Settings
}

// NewMicroserviceTimingLog takes its required settings and validates them.
// It sets up the handlers for immediate (console) logging and for Azure logging.
func NewMicroserviceTimingLog(instanceName string, serviceID *enum.ServiceID, s *Settings) (*MicroserviceTimingLog, error) {
	err := validateForTiming(instanceName, serviceID, s)
	if err != nil {
		return nil, err
	}

	var hname string
	if s == nil || s.HostName == nil || *s.HostName == `` {
		hname, err = os.Hostname()
		if err != nil {
			return nil, err
		}
	} else {
		hname = *s.HostName
	}

	timingTableName, ttnerr := azureTableName(logTypeTimingLog, serviceID)
	if ttnerr != nil {
		return nil, ttnerr
	}

	s.ServiceID = serviceID

	rtn := &MicroserviceTimingLog{
		instanceName: instanceName,
		hostName:     hname,
		serviceID:    serviceID,
		settings:     *s,
	}

	// For now I don't see any real use for an immediate handler

	if s.AzureSharedKey != nil && *s.AzureSharedKey != "" && s.AzureWorkspaceID != nil && *s.AzureWorkspaceID != "" {
		r := mbuf.NewAzureLogAnalyticsMessageRelay(*s.AzureWorkspaceID, *s.AzureSharedKey, timingTableName, azureRTBLogTimeField, s.RelaySettings())
		rtn.Relay = r
	}
	return rtn, nil
}

// validateForTiming performs a basic sanity-check of submitted startup parameters.
func validateForTiming(instanceName string, serviceID *enum.ServiceID, s *Settings) error {
	if s == nil {
		return errors.New(ErrorNilSettings)
	}

	if serviceID == nil || !serviceID.Valid() {
		return errors.New(ErrorBadServiceID)
	}

	svc := enum.Service.ByID(serviceID)
	if svc == nil || svc.LogArea == `` || (!serviceID.Equals(&enum.Service.Test.ID) && svc.LogArea == enum.ServiceLogArea.None) {
		return errors.New(ErrorBadServiceID)
	}

	if instanceName == "" {
		return errors.New(ErrorEmptyInstanceName)
	}

	// can't have an Azure log without Azure settings
	/*if s.AzureSharedKey == nil || *s.AzureSharedKey == "" || s.AzureWorkspaceID == nil || *s.AzureWorkspaceID == "" {
		return errors.New(ErrorAzureSettingsMissing)
	}*/

	return nil
}

// Write takes a Timing and submits it to the MicroserviceTimingLog's configured
// handlers for immediate and remote logging.
func (l *MicroserviceTimingLog) Write(c context.Context, t *MSTimingLogEntry) {
	if l == nil || t == nil {
		return
	}

	if l.Relay != nil {
		l.Relay.Add(t)
	}
}

// newEmptyTimingLog creates a MicroserviceTimingLog that will not do much of anything, but will also not crash when it is asked to do something. Use for testing.
// This method will kill the app with a fatal error if it encounters any problems.
func newEmptyTimingLog() *MicroserviceTimingLog {
	instanceName := `unknown-app`
	s := (&Settings{}).ForTesting()

	l, err := NewMicroserviceTimingLog(instanceName, &enum.Service.Test.ID, s)
	if err != nil {
		olog.Fatal(err.Error())
	}

	return l
}

// timingMessageHandler is the default immediate (console) handler for MicroserviceTimingLog.
type timingMessageHandler struct{}

func (h *timingMessageHandler) HandleMessage(msg *MSTimingLogEntry) {
	// TODO: make this nice for local debugging
	// we have a strong type here, so we can do whatever we want
	// for now, just marshal it and shoot it to stdout
	bytes, err := json.Marshal(msg)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
