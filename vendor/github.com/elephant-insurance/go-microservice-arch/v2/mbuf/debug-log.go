package mbuf

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// debugLog is a simple console logger for debugging this package ONLY
// if enabled, it will print diagnostic info to the console at run time
// we cannot use the regular logging package here because it would create a circular dependency

var debugLogging = false

func debugLog(msg string) {
	if debugLogging {
		fmt.Fprintln(os.Stderr, msg)
	}
}

func debugError(err error) {
	if debugLogging {
		fmt.Println(err.Error())
	}
}

func debugSpew(args ...interface{}) {
	if debugLogging {
		spew.Dump(args...)
	}
}
