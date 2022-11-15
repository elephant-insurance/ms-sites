package log

import (
	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/mbuf"
)

// Evaluator is any func that takes no args and returns a bool
type Evaluator func() bool

// Log is the highest-level inteerface in this package. It encapsulates a complete logging path.
// It manages console printing, remote receivers, and levels, and creates Loggers.
type Log interface {
	// AddMessageHandler lets us add a new message handler to the Log. MessageHandlers
	// fire on the calling thread immediately when the Logger logs.
	// MessageHandlers do things like write messages to stdout.
	AddMessageHandler(handler MSMessageHandler, level *enum.EnumLogLevelItem)
	// AddMessageRelay lets us add a new message relay for logs submitted with this Log.
	// MessageRelays send log messages to remote receivers.
	AddMessageRelay(messageRelay mbuf.MessageRelay, level *enum.EnumLogLevelItem)
}

// MSMessageHandler is a generic interface for structs that process microservice log messages without returning anything
// Create these and add them to the immediateHandlers collection of a logger to have it do work on the calling thread
// like write the message to stdout.
// A MSMessageHandler must run very fast, as it runs on the caller's thread while the caller waits
type MSMessageHandler interface {
	// Active returns false if the handler has been set to inactive
	Active() bool
	// HandleMessage does something with a message, like writing it to the console
	HandleMessage(lw MicroserviceLogger)
	// ToggleActive turns this handler on and off
	ToggleActive(bool)
}

// RTBMessageHandler is a generic interface for structs that process RTB log messages without returning anything
// Create these and add them to the immediateHandlers collection of a logger to have it do work on the calling thread
// like write the message to stdout.
type RTBMessageHandler interface {
	// HandleMessage does something with a message, like writing it to the console
	HandleMessage(msg *RTBLogEntry)
}

// PRIVATE Interfaces

// consoleFormatter receives ALL the data in the logwriter and formats it for the console
type consoleFormatter interface {
	Format(map[string]interface{}, []string) string
}

// relay is a simple internal struct for keeping track of a messagerelay
type relay struct {
	mbuf.MessageRelay
	active bool
}

// consolePrinter allows us to easily override where stuff gets printed. Use this for testing.
type consolePrinter interface {
	Print(msg string)
}
