package dig

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/alert"
	"github.com/elephant-insurance/go-microservice-arch/v2/cfg"
	"github.com/elephant-insurance/go-microservice-arch/v2/crasher"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/timing"
	"github.com/gin-gonic/gin"
)

const (
	runTestsParameterName = "runTests"
)

const (
	// DefaultPath is a convenience constant.
	// Add the diagnostics handler to this route in your app for consistency
	baseDigPath  = `/diagnostics`
	DefaultPath  = baseDigPath + "/*params"
	timingsPath  = "timings"
	memStatsPath = "memstats"
	pprofPath    = "pprof"
)

var (
	diagnosticTests                                           map[string]DiagnosticTest
	environment                                               *enum.EnumServiceEnvironmentItem
	lastLogEntry                                              string
	procName, appName, instanceName, hostname, branch, commit string
	imagetag, build, loglevel, datebuilt                      string
	startTime, appModTime                                     time.Time
	testArray                                                 []string
	verboseLogging                                            bool
)

var digCrasher crasher.Crasher = crasher.NewRealCrasher()

// Initialize sets up the diagnostic namespace.
// Environment should be set by the config file.
// The returned handler should be assigned to a route in the consuming application.
func Initialize(requiredConfig cfg.Configurator, s *Settings, g *gin.Engine) {
	lw := log.ForFunc(context.Background())

	processName := os.Args[0]
	instanceName = requiredConfig.GetInstanceName()
	s.InstanceName = instanceName
	appName = string(requiredConfig.GetServiceID())
	s.AppName = appName
	loglevel = string(requiredConfig.GetLogLevel())

	diagnosticTests = map[string]DiagnosticTest{}
	envID := requiredConfig.GetEnvironment()
	isProd, isDevTest := false, false
	if !envID.Valid() {
		digCrasher.Fatal(ErrorMsgEmptyEnvironment)
	} else {
		environment = enum.ServiceEnvironment.ByID(&envID)
		isProd = environment.IsProduction()
		isDevTest = environment.ID.Equals(&enum.ServiceEnvironment.Testing.ID) || environment.ID.Equals(&enum.ServiceEnvironment.Development.ID)
	}

	if processName == "" {
		digCrasher.Fatal(ErrorMsgEmptyProcName)
	}
	if !isDevTest && gin.Mode() == gin.TestMode {
		digCrasher.Fatal(ErrorMsgGinTestMode)
	}
	if isProd && gin.Mode() == gin.DebugMode {
		digCrasher.Fatal(ErrorMsgGinDebugMode)
	}
	if s.Branch == "" {
		digCrasher.Fatal(ErrorMsgEmptyBranch)
	}
	if s.Commit == "" {
		digCrasher.Fatal(ErrorMsgEmptyCommit)
	}
	if s.ImageTag == "" {
		digCrasher.Fatal(ErrorMsgEmptyImageTag)
	}
	if s.Build == "" {
		digCrasher.Fatal(ErrorMsgEmptyBuild)
	}
	if s.DateBuilt == "" {
		digCrasher.Fatal(ErrorMsgEmptyDateBuilt)
	}

	config := s

	config.ProcName = getBaseFileName(processName)

	// TODO: make this configurable, though it is unlikely that we'll ever want this true
	config.VerboseLog = false

	host, err := os.Hostname()
	if err != nil {
		digCrasher.Fatal(ErrorMsgCannotDetermineHostname)
	}
	config.Hostname = host

	workingDirectory, err := os.Getwd()
	if err != nil {
		digCrasher.Fatal(ErrorMessageCannotDetermineWorkingDirectory)
	}

	appFilepath := fmt.Sprintf("%v/%v", workingDirectory, config.ProcName)
	appFileHandle, err := os.Stat(appFilepath)
	if err != nil {
		// hack for bare-metal run on Windoze
		if !isDevTest {
			digCrasher.Fatal(ErrorMessageCannotLoadFileHandle)
		}
		lw.Warn(ErrorMessageCannotLoadFileHandleWarnDT)
	}

	if appFileHandle == nil || appFileHandle.Sys() == nil || appFileHandle.Mode().IsDir() {
		if !isDevTest {
			digCrasher.Fatal(ErrorMessageInvalidPath)
		}
		lw.Warn(ErrorMessageCannotLoadFileHandleWarnDT)
	} else {
		config.AppModTime = appFileHandle.ModTime()
	}

	initializeFromConfig(config, g)
}

// InitializeFromConfig makes it possible to set up diagnostics without runtime calls, primarily for testing.
func initializeFromConfig(c *Settings, g *gin.Engine) {
	var err error

	startTime = time.Now()
	diagnosticTests = map[string]DiagnosticTest{}
	packageStats = map[string]PackageStats{}

	// copy the settings into local private variables
	procName = c.ProcName
	appName = c.AppName
	instanceName = c.InstanceName
	branch = c.Branch
	commit = c.Commit
	imagetag = c.ImageTag
	build = c.Build
	datebuilt = c.DateBuilt
	hostname = c.Hostname
	if c.StartTime != nil {
		startTime = *c.StartTime
	}
	appModTime = c.AppModTime
	if c.Diagnostics != nil {
		diagnosticTests = c.Diagnostics
	}
	testArray = c.TestArray
	verboseLogging = c.VerboseLog

	// default initial capacity with no limit set: 1024 because why not
	var tbCap, msbCap int32
	tbCap, msbCap = 16*1024, 16*1024

	if c.BufferMax != nil {
		msbCap = *c.BufferMax
		tbCap = *c.BufferMax
	}

	if c.MemStatBufferMax != nil {
		msbCap = *c.MemStatBufferMax
	}

	if c.TimingBufferMax != nil {
		tbCap = *c.TimingBufferMax
	}

	timing.ApplicationTimings, err = timing.NewTimingBuffer(int(tbCap))
	if err != nil {
		digCrasher.Fatal(ErrorMsgFailedToCreateTimingBuffer)
	}

	memStats, err = NewMemStatBuffer(int(msbCap))
	if err != nil {
		digCrasher.Fatal(ErrorMsgFailedToCreateMemStatBuffer)
	}

	if MemStatIntervalSeconds > 0 {
		RecordMemStats()
	}

	peakJumpBuffer, err = NewMemPeakJumpBuffer(peakJumpBufferSize)
	if err != nil {
		digCrasher.Fatal(ErrorMsgFailedToCreateMemMaxBuffer)
	}

	// start with our default analytical functions
	memStatAnalytics = []MemStatAnalytic{
		MemStatAnalyticNewPeaksLast24Hours,
		MemStatAnalyticGrowthSinceBaseline,
		MemStatAnalyticMonotonicGrowthLast24Hours,
	}

	alert.AddAnalytic(runMemStatAnalytics)
}

// Tokenize on path separators. Starting from the end and working toward the start,
//  return the first token with positive length
func getBaseFileName(p string) string {
	reg := regexp.MustCompile(`[/\\]`)
	tokens := reg.Split(p, -1)
	nt := len(tokens)
	for i := 1; i < nt+1; i++ {
		if i == nt || len(tokens[nt-i]) > 0 {
			return tokens[nt-i]
		}
	}
	return ""
}
