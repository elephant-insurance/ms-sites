package crasher

import (
	"fmt"
	"os"
	"strings"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
)

// crasher is a simple interface for crashing the application.
// In production, we use a RealCrasher, which behaves like a normal system Panic or Fatal error.
// For testing we can use a TestCrasher, which simply tracks the crashes it would have performed if it were "real".

type Crasher interface {
	// Fatal prints its messsage to stderr then immediately exits the application with an error code (1).
	Fatal(msg string)
	// Panic writes its message to stderr and then calls panic(). Panics "bubble up" and may be caught by consuming code.
	Panic(msg string)
}

type realCrasher struct{}

func (rc *realCrasher) Fatal(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func (rc *realCrasher) Panic(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	panic(msg)
}

func NewRealCrasher() Crasher {
	return &realCrasher{}
}

type TestCrasher struct {
	FatalMessages []string
	Fatals        *clicker.Clicker
	PanicMessages []string
	Panics        *clicker.Clicker
}

func (tc *TestCrasher) Fatal(msg string) {
	tc.Fatals.Click(1)
	tc.FatalMessages = append(tc.FatalMessages, msg)
}

func (tc *TestCrasher) Panic(msg string) {
	tc.Panics.Click(1)
	tc.PanicMessages = append(tc.PanicMessages, msg)
}

func NewTestCrasher() *TestCrasher {
	return &TestCrasher{
		FatalMessages: []string{},
		Fatals:        &clicker.Clicker{},
		PanicMessages: []string{},
		Panics:        &clicker.Clicker{},
	}
}

// LastFatalMessageContains checks the last fatal message to make sure it matches what we expect.
// Use this to test Fatal errors during initialization.
func (tc *TestCrasher) LastFatalMessageContains(msg string) bool {
	if tc.Fatals.Clicks < 1 || len(tc.FatalMessages) < 1 {
		return false
	}

	return strings.Contains(tc.LastFatalMessage(), msg)
}

func (tc *TestCrasher) LastFatalMessage() string {
	if len(tc.FatalMessages) > 0 {
		return tc.FatalMessages[len(tc.FatalMessages)-1]
	}

	return ``
}
