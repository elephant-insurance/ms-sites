package uf

import (
	enum "github.com/elephant-insurance/enumerations/v2"
)

type eventUtil struct{}

var EventFactory = &eventUtil{}

// Event represents a condition or occurrence that we wnat to keep track of with automation.
// It may represent an error or critical condition requiring immediate attention,
// a diagnostic message for tracking statistics over time, etc.
type Event struct {
	// Description should be a human-readable description of what happened.
	// It should be short, but does not need to be constant or machine-readable.
	Description string
	// ID is a unique identifier for this event.
	// It should not be specific to any source (e.g., the app raising it)
	// or to any target (e.g., the destination of a failed request event)
	ID *enum.EventID
	// Severity identifies the importance of the event:
	// Developmental, Diagnostic, Informational, Error, and Critical
	Severity *enum.EventSeverityID
	// TargetService identifies the target of the event identified by ID, where applicable.
	// It should be a const for reliable machine processing.
	TargetService *enum.ServiceID
}

// New returns an Event with the severity set to its default by the EventID
func (eu *eventUtil) New(id *enum.EventID, targetService *enum.ServiceID, description string) *Event {
	evType := enum.Event.ByID(id)
	if evType == nil {
		return nil
	}

	return &Event{
		Description:   description,
		ID:            &evType.ID,
		Severity:      &evType.Severity.ID,
		TargetService: targetService,
	}
}

// OverrideSeverity sets the severity of the event, in case the default severity is not appropriate
func (e *Event) OverrideSeverity(id *enum.EventSeverityID) {
	if id != nil {
		e.Severity = id
	}
}
