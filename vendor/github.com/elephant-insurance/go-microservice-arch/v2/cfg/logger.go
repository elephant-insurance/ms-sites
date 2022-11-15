package cfg

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// Logger is a simple logger designed to dump messages to the console BEFORE the ELephant logging package has loaded.
// We don't want references to our logging package here because our logging package needs values from the app configuration in order to initialize itself at app start.
type configLogger interface {
	Info(string)
	Warn(string)
	Error(string)
	Panic(string)
}

type logger struct{}

func (l *logger) Info(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

func (l *logger) Warn(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

func (l *logger) Error(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

func (l *logger) Panic(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

var debugLogging = false

func debugSpew(args ...interface{}) {
	if debugLogging {
		spew.Dump(args...)
	}
}
