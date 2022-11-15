package log

import (
	"context"
	"fmt"
	olog "log"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/davecgh/go-spew/spew"
	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/cfg"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
	"github.com/elephant-insurance/go-microservice-arch/v2/timing"
)

// DebugDump dumps the full structure of whatever is passed in to stderror.
// This is for diagnosing issues on a development machine.
// It will not do anything unless the environment is DEV or TEST
func DebugDump(args ...interface{}) {
	if defaultLog != nil && defaultLog.settings.Environment != nil && (*defaultLog.settings.Environment == enum.ServiceEnvironment.Development.ID || *defaultLog.settings.Environment == enum.ServiceEnvironment.Testing.ID) {
		spew.Dump(args...)
	}
}

// ForFunc is the primary and preferred way to create a MicroserviceLogger.
// It creates a new logger, sets its context, and sets its "func" field to the name of the calling func.
func ForFunc(c context.Context, args ...interface{}) MicroserviceLogger {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fn := trimFuncName(frame.Func.Name())

	return New().SetContext(c).SetFunction(fn).Trace(MessageCalled)
}

// Initialize sets up the package default Microservice Log and returns it.
// This is the function that you should use to set up logging in your main.go.
// This method will kill the application if it finds an error, so it should only be run at startup.
func Initialize(rq cfg.Configurator, s *Settings) *MicroserviceLog {
	level := rq.GetLogLevel()
	env := rq.GetEnvironment()
	instanceName := rq.GetInstanceName()
	serviceID := rq.GetServiceID()

	l, err := newMicroserviceLog(&level, instanceName, &env, &serviceID, s)
	if err != nil {
		olog.Fatal(err.Error())
	}
	defaultLog = l

	tl, tlerr := NewMicroserviceTimingLog(instanceName, &serviceID, s)
	if tlerr != nil {
		olog.Fatal(tlerr.Error())
	}
	defaultTimingLog = tl

	return l
}

// New returns a new MicroserviceLogger without setting the func field or context.
func New() MicroserviceLogger {
	return defaultLog.NewMicroserviceLogger()
}

// Diagnostics returns the diagnostic output for the first relay in the default log, if any.
// This is almost always what we want.
func Diagnostics() map[string]interface{} {
	rtn := map[string]interface{}{}

	if defaultLog != nil && len(defaultLog.relays) > 0 && len(defaultLog.relays[defaultLog.level.ID]) > 0 {
		rtn[`app-log`] = defaultLog.relays[defaultLog.level.ID][0].Diagnostics()
		if defaultLog.eventRelay.active {
			rtn[`event-log`] = defaultLog.eventRelay.Diagnostics()
		} else {
			rtn[`event-log`] = `disabled`
		}
		if defaultTimingLog != nil && defaultTimingLog.Relay != nil {
			rtn[`timing-log`] = defaultTimingLog.Relay.Diagnostics()
		}
	} else {
		rtn["RelayEnabled"] = false
	}

	return rtn
}

// WriteTiming sends a timing log entry to the Azure microservice timing log
func WriteTiming(c msrqc.Context, t *timing.Timing) {
	te := NewTimingEntry(c, t)
	defaultTimingLog.Write(c, te)
}

// defaultLog and defaultTimingLog are the Logs that carry out package-level commands.
// If these are nil, then all our package-level funcs will fail, so we initialize it with a trivial log.
// This will get you through most tests without having to call Initialize(...), which requires parameters.
var (
	defaultLog       = newEmptyMicroserviceLog()
	defaultTimingLog = newEmptyTimingLog()
)

// set this to true to see diagnostic output from various logging functions
var debugLogging = false

// debugLog prints to sderr so that we can see it during testing
func debugLog(a ...interface{}) {
	if debugLogging {
		fmt.Fprintln(os.Stderr, a...)
	}
}

// debugLogf printfs with a linefeed to sderr so that we can see it during testing
func debugLogf(format string, a ...interface{}) {
	if debugLogging {
		fmt.Fprintf(os.Stderr, format+"\n", a...)
	}
}

func debugSpew(a ...interface{}) {
	if debugLogging && !defaultLog.level.IsMoreUrgentThan(enum.LogLevel.Debug) {
		spew.Dump(a...)
	}
}

func debugStack() {
	if debugLogging && !defaultLog.level.IsMoreUrgentThan(enum.LogLevel.Debug) {
		debug.PrintStack()
	}
}

// trimFuncName removes all but the last bit of a func name, since we don't really need the entire path.
func trimFuncName(funcName string) string {
	strlen := len(funcName)
	if strlen == 0 {
		return unknownFuncName
	}

	lastSlash := strings.LastIndex(funcName, `/`)

	if lastSlash < 1 {
		return funcName
	}

	if (strlen - lastSlash) < 2 {
		return unknownFuncName
	}

	return funcName[lastSlash+1:]
}
