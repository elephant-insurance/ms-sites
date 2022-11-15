package dig

import (
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/rbuf"
)

// this handler effectively allows us to inject dig functions into the log execution path
// without creating a circular dependency

/*

TODO:

add *alert.Alert as member of lw
create lw.WithAlert(alert) to set this field
might need to add lw methods to extract needed values
dig.init:  add dig immediate handler to default microservice log
HandleMessage: fire off sub-funcs:
	write message to dig.buffer
	send alert if any to ala
flesh out alert struct
dig.CreateAlert
Alert.Send?

*/

var logBufferHandler log.MSMessageHandler = newLogHandler()

const defaultLogBufferSize int = 1000

func newLogHandler() log.MSMessageHandler {
	// it's safe to swallow the error because it's only not nil if size is too small
	buf, _ := rbuf.NewRingBuffer(defaultLogBufferSize, dummyLogMessageMaker)
	return &logHandlerType{
		true,
		buf,
	}
}

type logHandlerType struct {
	active bool
	buffer rbuf.RingBuffer
}

func (lh *logHandlerType) Active() bool {
	if lh == nil {
		return false
	}

	return lh.active
}

func (lh *logHandlerType) ToggleActive(val bool) {
	if lh != nil {
		lh.active = val
	}
}

func (lh *logHandlerType) HandleMessage(logWriter log.MicroserviceLogger) {
}

type logMessageType struct {
	dmap map[string]interface{}
	keys []string
}

func dummyLogMessageMaker() interface{} {
	return &logMessageType{
		dmap: map[string]interface{}{
			`foo`: `bar`,
		},
		keys: []string{`foo`},
	}
}
