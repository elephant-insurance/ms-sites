package log

import (
	"context"
	"errors"
	olog "log"
	"os"
	"runtime"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/crasher"
	"github.com/elephant-insurance/go-microservice-arch/v2/mbuf"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
)

// A "Log" is a complete description of a logging method for a microservice
// It is the highest-level object in the logging namespace
// It encapsulates:
//  * a specific log message structure (broadly speaking, a table in Azure Log Analytics)
//  * a message struct type that encapsulates the message structure
//  * a set of rules for printing log messages to the console
//  * a set of zero or more message relay structs for forwarding messages to analysis services
//  * a log level for each message relay
//  * any app-specific plug-ins (renderers, marshalers, flushers) for the message relays
//  * a factory method for making loggers

// The default Log for the logging package is the Microservice Log
// When the logging package loads, it automatically creates a microservice log called default
// Most package-level logging methods (e.g., logging.Debug(msg)) call the matching method on the default Log

// MicroserviceLog is a top-level struct in the logging namespace
// It holds settings for a microservice log
type MicroserviceLog struct {
	// crasher handles panics and fatals
	crasher crasher.Crasher
	// eventRelay is a separate relay for event messages
	eventRelay relay
	// immediateHandlers are message handlers that run on the calling thread
	// this is primarily for logging to the console in development
	// use a map by level so that we can find the relevant handlers as fast as possible
	immediateHandlers map[enum.LogLevelID][]MSMessageHandler
	// level is the base level for the log. Nothing less urgent than this will be logged
	level *enum.EnumLogLevelItem
	// relays are message relays that forward our messages to other services
	relays map[enum.LogLevelID][]relay
	// settings keeps a copy of the values we were initialized with
	// this save copying each setting to a struct field
	settings Settings
}

// NewMicroserviceLog sets up a new Microservice Log and returns it. It does not set the default log for the package.
func NewMicroserviceLog(level *enum.LogLevelID, instanceName string, environment *enum.ServiceEnvironmentID, serviceID *enum.ServiceID, s *Settings) (*MicroserviceLog, error) {
	return newMicroserviceLog(level, instanceName, environment, serviceID, s)
}

// NewTestMicroserviceLog returns a Log with a TestPrinter for observing what is logged
func NewTestMicroserviceLog(level *enum.LogLevelID, instanceName string, environment *enum.ServiceEnvironmentID, s *Settings) (*MicroserviceLog, *TestConsolePrinter, *mbuf.TestMessageRelay, error) {
	rtn, err := newMicroserviceLog(level, instanceName, environment, &enum.Service.Test.ID, s)
	if err != nil {
		return nil, nil, nil, err
	}

	llitem := enum.LogLevel.ByID(level)
	rtn.immediateHandlers = emptyImmediates()
	th, tp := NewTestConsoleLogger()
	rtn.AddMessageHandler(th, llitem)
	tr := mbuf.NewTestMessageRelay()
	rtn.AddMessageRelay(tr, llitem)
	rtn.eventRelay = relay{mbuf.NewTestMessageRelay(), false}
	defaultLog = rtn

	return rtn, tp, tr, nil
}

func newMicroserviceLog(level *enum.LogLevelID, instanceName string, environment *enum.ServiceEnvironmentID, serviceID *enum.ServiceID, s *Settings) (*MicroserviceLog, error) {
	if level == nil || instanceName == "" || environment == nil || serviceID == nil {
		return nil, errors.New(ErrorBadNewLogArgs)
	}

	var cs Settings

	if s != nil {
		err := s.validate()
		if err == nil {
			cs = *s
		} else {
			return nil, err
		}
	} else {
		cs = defaultSettings()
	}

	// we require these fields to be passed in separately because they are essential to the functioning of the package
	// it doesn't matter what is in the settings object, the params rule
	cs.InstanceName = &instanceName
	cs.Environment = environment
	cs.Level = level
	cs.ServiceID = serviceID

	if cs.HostName == nil {
		if s == nil || s.HostName == nil || *s.HostName == `` {
			hn, err := os.Hostname()
			if err != nil {
				return nil, err
			}
			cs.HostName = &hn
		} else {
			hn := *s.HostName
			cs.HostName = &hn
		}
	}

	rtn := MicroserviceLog{
		crasher:           crasher.NewRealCrasher(),
		eventRelay:        relay{nil, false},
		immediateHandlers: emptyImmediates(),
		relays:            emptyRelays(),
		settings:          cs,
	}

	// build our dispatch maps

	// first set the levels for all our potential receivers
	// everything gets set to the default level, then overridden as specified
	defaultLevel := enum.LogLevel.ByID(level)
	azureLevel := defaultLevel
	consoleLevel := defaultLevel
	logglyLevel := defaultLevel
	rtn.level = defaultLevel

	if cs.AzureLogLevel != nil {
		azureLevel = enum.LogLevel.ByID(cs.AzureLogLevel)
	}

	if cs.ConsoleLogLevel != nil {
		consoleLevel = enum.LogLevel.ByID(cs.ConsoleLogLevel)
	}

	if cs.LogglyLogLevel != nil {
		logglyLevel = enum.LogLevel.ByID(cs.LogglyLogLevel)
	}

	// now figure out what receivers we need for what levels

	// do we need a console handler? Default (nil) to YES
	if cs.ConsoleMode == nil || *cs.ConsoleMode == ConsoleModeStandardOut {
		// for now just create a simple default handler
		h := NewDefaultConsoleLogger()
		rtn.AddMessageHandler(h, consoleLevel)
	}

	// do we need Azure relays?
	if cs.AzureSharedKey != nil && *cs.AzureSharedKey != "" && cs.AzureWorkspaceID != nil && *cs.AzureWorkspaceID != "" {
		logTableName, atnerr := azureTableName(logTypeAppLog, serviceID)
		if atnerr != nil {
			return nil, atnerr
		}

		eventTableName, etnerr := azureTableName(logTypeEventLog, serviceID)
		if etnerr != nil {
			return nil, etnerr
		}

		r := mbuf.NewAzureLogAnalyticsMessageRelay(*cs.AzureWorkspaceID, *cs.AzureSharedKey, logTableName, azureMicroserviceLogTimeField, s.RelaySettings())
		rtn.AddMessageRelay(r, azureLevel)
		e := mbuf.NewAzureLogAnalyticsMessageRelay(*cs.AzureWorkspaceID, *cs.AzureSharedKey, eventTableName, azureMicroserviceEventLogTimeField, s.RelaySettings())
		rtn.eventRelay = relay{e, true}
	}

	// do we need a Loggly relay?
	if cs.LogglyToken != nil && *cs.LogglyToken != "" {
		r := mbuf.NewLogglyMessageRelay(*cs.LogglyToken, instanceName, *environment, s.RelaySettings())
		rtn.AddMessageRelay(r, logglyLevel)
	}

	return &rtn, nil
}

// newEmptyMicroserviceLog creates a MicroserviceLog that will not do much of anything, but will also not crash when it is asked to do something. Use for testing.
// This method will kill the app with a fatal error if it encounters any problems.
func newEmptyMicroserviceLog() *MicroserviceLog {
	level := &enum.LogLevel.Panic.ID
	instanceName := `unknown-app`
	environment := &enum.ServiceEnvironment.Testing.ID
	s := (&Settings{}).ForTesting()

	l, err := newMicroserviceLog(level, instanceName, environment, &enum.Service.Test.ID, s)
	if err != nil {
		olog.Fatal(err.Error())
	}

	return l
}

// NewMicroserviceLogger is the primary factory method for creating a Logger
func (l *MicroserviceLog) NewMicroserviceLogger() MicroserviceLogger {
	return l.newMicroserviceLogger()
}

func (l *MicroserviceLog) newMicroserviceLogger() *microserviceLogwriter {
	return &microserviceLogwriter{
		active:   true,
		elsable:  false,
		events:   []*MSEventLogEntry{},
		fieldMap: map[string]interface{}{},
		log:      l,
		entry:    l.NewEntry(),
	}
}

// NewEntry creates a new entry, prepoulated with fields we know now
func (l *MicroserviceLog) NewEntry() *MicroserviceLogEntry {
	return &MicroserviceLogEntry{
		// thse aren't user-settable, so it's fine to copy the pointer
		InstanceName: l.settings.InstanceName,
		HostName:     l.settings.HostName,
		ServiceID:    l.settings.ServiceID,
	}
}

// ForFunc creates a new MicroserviceLogger, sets its Function field to fieldName, and emits a Trace message.
// Optionally, additional parameter values may be passed in for logging, as { param1name, param1val, param2name, param2val, ...},
// where each name is a string, and each val can be anything.
// Use ForFunc at the top of any function in which you will want to write microservice log messages.
func (l *MicroserviceLog) ForFunc(c context.Context, args ...interface{}) MicroserviceLogger {
	lw := l.newMicroserviceLogger().SetContext(c)

	if len(args) > 1 {
		lw = lw.WithConsoleFields(args)
	}

	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fn := trimFuncName(frame.Func.Name())
	lw.SetFunction(fn)
	lw.Trace(MessageCalled)

	return lw
}

// dispatch determines what to do with a message and does it
func (l *MicroserviceLog) dispatch(lw *microserviceLogwriter, level *enum.EnumLogLevelItem, msg *string) {
	if l == nil || l.level == nil || lw == nil || lw.entry == nil || level == nil || msg == nil {
		return
	}

	lw.entry.Time = uf.Pointer.ToNow()

	// regardless of log level, clear any events that have been attached to this logwriter:
	events := lw.ClearEvents()

	// send the events, if we are sending:
	if l.eventRelay.active {
		for i := 0; i < len(events); i++ {
			thisEvent := events[i]
			// fill out the memstats
			thisEvent.addMemStats()
			// make sure the event(s) and the log message have the exact same timestamp
			thisEvent.Time = *lw.entry.Time
			l.eventRelay.Add(thisEvent)
		}
	}

	// never do anything else if we are logging at a more urgent level than this message
	if l.level.IsMoreUrgentThan(level) {
		// make sure nothing we added just for this message sticks around
		lw.clearEphemeralFields()
		return
	}

	lw.entry.Message = msg
	lw.entry.Level = &level.ID
	lw.entry.InstanceName = l.settings.InstanceName
	lw.entry.Environment = l.settings.Environment
	lw.entry.HostName = l.settings.HostName
	lw.entry.ServiceID = l.settings.ServiceID

	activeImmediates := l.immediateHandlers[level.ID]
	for i := 0; i < len(activeImmediates); i++ {
		handler := activeImmediates[i]
		if handler == nil || !handler.Active() {
			continue
		}

		handler.HandleMessage(lw)
	}

	// before copying all my fields with emit, make sure that I will need them
	// check my relays to see whether they're active and valid
	relaysForLevel := l.relays[level.ID]

	if len(relaysForLevel) > 0 {
		activeRelays := make([]relay, 0, len(relaysForLevel))
		for i := 0; i < len(relaysForLevel); i++ {
			thisRelay := relaysForLevel[i]
			if !thisRelay.active {
				continue
			}
			activeRelays = append(activeRelays, thisRelay)
		}

		if len(activeRelays) > 0 {
			myEntry := lw.emit()
			for i := 0; i < len(activeRelays); i++ {
				activeRelays[i].Add(myEntry)
			}
		}
	}

	// we have written everything we were going to write, so time to clear out any ephemeral values
	// so that we start fresh for next time:
	lw.clearEphemeralFields()

	if level.ID == enum.LogLevel.Fatal.ID {
		l.fatal(msg)
	}

	if level.ID == enum.LogLevel.Panic.ID {
		l.panic(msg)
	}
}

func emptyImmediates() map[enum.LogLevelID][]MSMessageHandler {
	rtn := map[enum.LogLevelID][]MSMessageHandler{}

	for i := 0; i < len(enum.LogLevel.Items); i++ {
		thisLevel := enum.LogLevel.Items[i]
		rtn[thisLevel.ID] = []MSMessageHandler{}
	}

	return rtn
}

func emptyRelays() map[enum.LogLevelID][]relay {
	rtn := map[enum.LogLevelID][]relay{}

	for i := 0; i < len(enum.LogLevel.Items); i++ {
		thisLevel := enum.LogLevel.Items[i]
		rtn[thisLevel.ID] = []relay{}
	}

	return rtn
}

func (l *MicroserviceLog) AddMessageRelay(messageRelay mbuf.MessageRelay, level *enum.EnumLogLevelItem) {
	if l == nil || messageRelay == nil || level == nil {
		return
	}

	for i := 0; i < len(enum.LogLevel.Items); i++ {
		thisLevel := enum.LogLevel.Items[i]
		if !level.IsMoreUrgentThan(thisLevel) {
			l.relays[thisLevel.ID] = append(l.relays[thisLevel.ID], relay{messageRelay, true})
		}
	}
}

func (l *MicroserviceLog) AddMessageHandler(handler MSMessageHandler, level *enum.EnumLogLevelItem) {
	if l == nil || handler == nil || level == nil {
		return
	}

	for i := 0; i < len(enum.LogLevel.Items); i++ {
		thisLevel := enum.LogLevel.Items[i]
		if !level.IsMoreUrgentThan(thisLevel) {
			l.immediateHandlers[thisLevel.ID] = append(l.immediateHandlers[thisLevel.ID], handler)
		}
	}
}

// sendAllNow attempts to send any messages that we have cached, usually before quitting
func (l *MicroserviceLog) sendAllNow() {
	for _, v := range l.relays[enum.LogLevel.Fatal.ID] {
		if v.active {
			v.SendNowNoRetry()
		}
	}
}

func (l *MicroserviceLog) fatal(msg *string) {
	sm := ""
	if msg != nil {
		sm = *msg
	}
	l.sendAllNow()
	l.crasher.Fatal(sm)
}

func (l *MicroserviceLog) panic(msg *string) {
	sm := ""
	if msg != nil {
		sm = *msg
	}
	l.sendAllNow()
	l.crasher.Panic(sm)
}
