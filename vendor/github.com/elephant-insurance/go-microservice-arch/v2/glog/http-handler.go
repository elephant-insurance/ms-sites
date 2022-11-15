package glog

import (
	"net/http"
	"time"

	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/msrqc"
	"github.com/gin-gonic/gin"
)

// TODO:
// Add required settings for RequiredHeaders []string and ValidatedHeaders []string
// Values should be set in config.yml by header key (e.g.: ValidatedHeaders: admiral-txbrand,admiral-txdomain)
// If a header is listed in RequiredHeaders, then it must be non-empty in the initiating request
// If a header is listed in ValidatedHeaders, then it must unmarshal to a valid enumeration
// If any header doesn't pass these tests, then glog short-circuits the handler stack, returning bad request

// LogRequestStart controls whether to write a message when a request is received
//
//	in addition to writing one when the response is written
//
// Set to false to reduce some logfile fluff
var LogRequestStart = false

// LogDiagnosticsRequests controls whether we log completion of diagnostics calls
// Generally, there is no reason not to turn this on UNLESS there is an
// application that "pings" the diagnostics page. The proper way to ping is
// with a HEAD request to /, not a GET to /diagnostics
var LogDiagnosticsRequests = false

var appPrefix = "UNK"

var suppressMethods = map[string]bool{
	http.MethodGet:     false,
	http.MethodHead:    true,
	http.MethodPost:    false,
	http.MethodPut:     false,
	http.MethodPatch:   false,
	http.MethodDelete:  false,
	http.MethodConnect: true,
	http.MethodOptions: true,
	http.MethodTrace:   true,
}

// SuppressLogForMethod prevents the logger from creating entries for the given method
// This can reduce log fluff from pulse checks and other automated processes
// By default the methods HEAD, CONNECT, OPTIONS, and TRACE are suppressed
func SuppressLogForMethod(m string) {
	suppressMethods[m] = true
}

// EnableLogForMethod forces the logger to create entries for the given method
// This can reduce lead to worthless log fluff from pulse checks and other automated processes
// By default the methods HEAD, CONNECT, OPTIONS, and TRACE are suppressed
func EnableLogForMethod(m string) {
	suppressMethods[m] = false
}

func serveHTTP(c *gin.Context) {
	lw := log.ForFunc(c)
	req := c.Request
	//spew.Dump("START:", c.Keys)

	start := time.Now()

	msrqc.SetHeaders(c, req.Header, appPrefix)

	// If we've received a client IP, copy it to the IP Address field
	if txipptr := msrqc.GetTransactionIPAddress(c); txipptr != nil {
		lw.SetIPAddress(*txipptr, true)
	}

	suppress, ok := suppressMethods[req.Method]
	skipLogging := (ok && suppress)

	if skipLogging {
		c.Next()
		return
	}

	lw.WithHTTPMethod(req.Method).WithURL(req.URL.Path)
	// spew.Dump(lw)

	if LogRequestStart {
		lw.Info(log.MessageStarted).WithHTTPMethod(req.Method).WithURL(req.URL.Path)
	}

	c.Next()

	res := c.Writer

	lw.WithHTTPStatus(res.Status())
	lw.WithElapsedMicroseconds(time.Since(start).Microseconds())
	lw.Info(log.MessageCompleted)
	//spew.Dump("END", c.Keys)
}
