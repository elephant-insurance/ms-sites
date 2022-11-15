package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"

	enum "github.com/elephant-insurance/enumerations/v2"
)

// AbbreviateConsoleOutput if true, causes certain fields that are not
// very useful for local debugging to be omitted from the console log
var AbbreviateConsoleOutput = true

var ShowNilsInConsole = false

// ConsoleMode describes an output behavior
type ConsoleMode string

// ConsoleModeNone suppresses output to stdout
const ConsoleModeNone ConsoleMode = "ConsoleModeNone"

// ConsoleModeStandardOut sends nicely-formatted log output to stdout
const ConsoleModeStandardOut ConsoleMode = "ConsoleModeStandardOut"

// ConsoleModeSendOrStandardOut suppresses output to stdout as long as we are relaying
// if relaying fails, rough json will be printed to stdout
const ConsoleModeSendOrStandardOut ConsoleMode = "ConsoleModeSendOrStandardOut"

type defaultConsoleLogger struct {
	active    bool
	formatter consoleFormatter
	printer   consolePrinter
}

type defaultConsoleFormatter struct{}

func (f *defaultConsoleFormatter) Format(m map[string]interface{}, sortedFields []string) string {
	var msg string
	lvlstr := `NONE`
	buf := bytes.Buffer{}

	for i := 0; i < len(sortedFields); i++ {
		key := sortedFields[i]
		if key == "" {
			continue
		}

		// format the value
		if val, ok := m[key]; ok {
			// keep dereferencing until we either get to a real value or a nil:
			for reflect.ValueOf(val).Kind() == reflect.Ptr {
				if val != nil {
					val = reflect.ValueOf(val).Elem().Interface()
				} else {
					if !ShowNilsInConsole {
						continue
					}
					val = []byte(`<nil>`)
				}
			}

			var valStr string

			switch v := val.(type) {
			case []byte:
				valStr = string(v)
			case json.RawMessage:
				valStr = string(v)
			case string:
				valStr = v
			default:
				/*this switch will give us something useful from type aliases:
				switch reflect.ValueOf(val).Kind() {
				case reflect.String:
					case
				}
				debugLog(`failed to find type match for field: ` + key)
				// debugStack()
				debugSpew(v)*/
				valStr = fmt.Sprintf(`%v`, val)
			}
			//valBytes = []byte(fmt.Sprintf(`%v`, val))

			// decide what to do with this field based on key
			switch key {
			case `level`:
				lvlstr = valStr
			case `msg`:
				msg = valStr
			default:
				if AbbreviateConsoleOutput {
					if skip, ok := consoleSuppressKeys[key]; ok && skip {
						break
					}
				}
				buf.WriteString(key)
				buf.WriteString(`=`)
				buf.WriteString(valStr)
				buf.WriteString(` `)
			}
		} else {
			debugLog(`KEY NOT FOUND: ` + key)
		}
	}

	lvlAbb := strings.ToUpper(lvlstr[1:4])

	var mapstring string
	if buf.Len() > 0 {
		mapstring = "\n   " + buf.String() + "\n"
	}

	logColor, ok := colorForLevel[lvlAbb]

	if !ok {
		logColor = color.New(color.FgWhite)
	}

	rtn := logColor.Sprintf("%v | ", lvlAbb) + msg[1:len(msg)-1] + logColor.Sprint(mapstring)

	debugLog(rtn)
	return rtn
}

func NewDefaultConsoleLogger() MSMessageHandler {
	return &defaultConsoleLogger{true, &defaultConsoleFormatter{}, &realConsolePrinter{}}
}

func NewTestConsoleLogger() (MSMessageHandler, *TestConsolePrinter) {
	cp := &TestConsolePrinter{}
	rtn := &defaultConsoleLogger{true, &defaultConsoleFormatter{}, cp}

	return rtn, cp
}

// HandleMessage dumps our log message to stdout, if active
// this is primarily for local debugging on the developer's machine
// so messages are printed in two lines for best visibility
func (l *defaultConsoleLogger) HandleMessage(logger MicroserviceLogger) {
	if l == nil || !l.active || l.printer == nil {
		return
	}

	lw, ok := logger.(*microserviceLogwriter)

	if !ok || lw == nil || lw.entry == nil || lw.entry.Level == nil {
		return
	}

	lwmap, sortedKeys := lw.makeMap()
	// spew.Dump(lwmap)
	if len(lwmap) == 0 {
		return
	}

	consoleMsg := l.formatter.Format(lwmap, sortedKeys)
	// spew.Dump("CMSG: " + consoleMsg)
	l.printer.Print(consoleMsg)
}

func (l *defaultConsoleLogger) Active() bool {
	return l != nil && l.active
}

func (l *defaultConsoleLogger) ToggleActive(a bool) {
	if l != nil {
		l.active = a
	}
}

type realConsolePrinter struct{}

func (cp *realConsolePrinter) Print(msg string) {
	fmt.Print(msg)
}

type TestConsolePrinter struct {
	buf bytes.Buffer
}

func (cp *TestConsolePrinter) Print(msg string) {
	fmt.Fprint(&cp.buf, msg)
}

func (cp *TestConsolePrinter) SetBuf() *bytes.Buffer {
	if cp != nil {
		cp.buf = bytes.Buffer{}
		return &cp.buf
	}

	return nil
}

var colorForLevel = map[string]*color.Color{
	strings.ToUpper(string(enum.LogLevel.Trace.ID))[0:3]: color.New(color.FgMagenta, color.Faint),
	strings.ToUpper(string(enum.LogLevel.Debug.ID))[0:3]: color.New(color.FgWhite, color.Faint),
	strings.ToUpper(string(enum.LogLevel.Info.ID))[0:3]:  color.New(color.FgGreen),
	strings.ToUpper(string(enum.LogLevel.Warn.ID))[0:3]:  color.New(color.FgHiYellow),
	strings.ToUpper(string(enum.LogLevel.Error.ID))[0:3]: color.New(color.FgHiRed, color.Bold),
	strings.ToUpper(string(enum.LogLevel.Fatal.ID))[0:3]: color.New(color.FgHiRed, color.Bold, color.Underline),
	strings.ToUpper(string(enum.LogLevel.Panic.ID))[0:3]: color.New(color.FgHiRed, color.Underline),
}

// consoleSuppressKeys contains log keys that are not very useful for local logging
var consoleSuppressKeys = map[string]bool{
	fieldNameMSInstanceName: true,
	fieldNameMSServiceID:    true,
	fieldNameMSEnvironment:  true,
	fieldNameMSHostName:     true,
	fieldNameMSIPAddress:    true,
	fieldNameMSTime:         true,
}
