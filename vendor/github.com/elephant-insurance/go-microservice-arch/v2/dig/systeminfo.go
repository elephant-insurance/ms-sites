package dig

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	enum "github.com/elephant-insurance/enumerations/v2"
	"github.com/elephant-insurance/go-microservice-arch/v2/alert"
	"github.com/elephant-insurance/go-microservice-arch/v2/log"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
)

// SystemInfo is the output type of the diagnostics page. With this struct, the results
// of a diagnostic page load can be loaded into memory and analyzed by another process.
type SystemInfo struct {
	Alerts             []*uf.Event                        `json:"Alerts,omitempty"`            // Alerts for any bad conditions detected while running diagnostics
	AppName            string                             `json:"AppName"`                     // The name of the application (codebase).
	Branch             string                             `json:"Branch"`                      // The branch the app was built from (set at initialization).
	Build              string                             `json:"Build"`                       // The Jenkins build that built the image (set at initialization).
	Commit             string                             `json:"Commit"`                      // The commit the app was built from (set at initialization).
	DateBuilt          string                             `json:"DateBuilt"`                   // The date the image was built (set at initialization).
	DiagnosticResults  []*DiagnosticResult                `json:"DiagnosticResults,omitempty"` // The full results of any DiagnosticTests that ran, if any.
	Environment        enum.ServiceEnvironmentID          `json:"Environment"`                 // The runtime environment of the consuming app.
	FailedTests        *int                               `json:"FailedTests,omitempty"`       // The number of DiagnosticTests that failed, if any.
	HostName           string                             `json:"HostName"`                    // The hostname (or container ID) for the machine we're running on.
	ImageTag           string                             `json:"ImageTag"`                    // The tag of the docker image (set at initialization).
	InstanceName       string                             `json:"InstanceName"`                // The actual instance name of the running application
	MemStatBaseline    *memStat                           `json:"MemStatBaseline,omitempty"`   // Memory usage baseline statistics.
	MemStats           *memStat                           `json:"MemStats"`                    // Memory usage statistics.
	Modified           time.Time                          `json:"Modified"`                    // The modified time of the running executable.
	Links              *map[string]string                 `json:"Links,omitempty"`
	LogLevel           string                             `json:"LogLevel"`                // Links to other diagnostics pages.
	OpenFileCount      *int                               `json:"OpenFileCount,omitempty"` // How many open file descriptors are in /var/<pid>/fd?
	PackageStatResults *map[string]map[string]interface{} `json:"PackageStats,omitempty"`
	SuccessfulTests    *int                               `json:"SuccessfulTests,omitempty"` // The number of DiagnosticTests that succeeded, if any.
	Uptime             string                             `json:"Uptime"`                    // How long it's been since the app started.
}

func getCurrentSystemInfo(runTests bool, r *http.Request) SystemInfo {
	lw := log.ForFunc(nil, `run-tests`, runTests)
	env := enum.ServiceEnvironment.Testing.ID
	if environment != nil {
		env = environment.ID
	}
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	shortStats := abbreviate(&stats)
	rtn := SystemInfo{
		AppName:         appName,
		InstanceName:    instanceName,
		Branch:          branch,
		Build:           build,
		Commit:          commit,
		DateBuilt:       datebuilt,
		Environment:     env,
		HostName:        hostname,
		ImageTag:        imagetag,
		LogLevel:        loglevel,
		MemStatBaseline: memStats.GetBaseline().Copy(),
		MemStats:        shortStats,
		Modified:        appModTime,
		Uptime:          formatDuration(time.Since(startTime)),
	}

	if fdc, err := getOpenFileDescriptorCount(); err == nil && fdc != nil {
		rtn.OpenFileCount = fdc
	}

	// Do we have any stats from other packages?
	if len(packageStats) > 0 {
		psr := map[string]map[string]interface{}{}
		// run through our map of stat generators
		for k, v := range packageStats {
			psr[k] = v()
		}
		rtn.PackageStatResults = &psr
	}

	if runTests {
		rtn.Alerts = alert.RunAllAnalytics()
		runAllTests(&rtn)
	}

	// try to get a base url from the request
	var basePath string
	if r != nil && r.URL != nil && r.URL.Host != `` {
		u := r.URL
		// log.DebugDump(rq)
		basePath = fmt.Sprintf(baseLinkPattern, u.Scheme, u.Host, instanceName)
	} else if r != nil && r.Host != `` {
		proto := `https`
		if environment.ID.Equals(&enum.ServiceEnvironment.Testing.ID) || environment.ID.Equals(&enum.ServiceEnvironment.Development.ID) {
			proto = `http`
		}
		basePath = fmt.Sprintf(baseLinkPattern, proto, r.Host, instanceName)
	} else {
		basePath = fmt.Sprintf(relativeLinkPattern, instanceName)
	}

	n := generateDiagnosticsLinks(basePath)
	rtn.Links = &n

	// only log if we're in verbose mode OR running tests. This prevents the logs from filling up with
	//  worthless records of simple pings.
	if verboseLogging || runTests {
		lw.WithConsoleField(`diagnostics`, rtn).Info(`service diagnostics`)
	}

	return rtn
}

// runAllTests runs all the tests in our test array and updates the submitted SystemInfo accordingly
func runAllTests(si *SystemInfo) {
	numTests := len(testArray)
	if numTests > 0 {
		results := make([]*DiagnosticResult, numTests)
		var wg sync.WaitGroup
		wg.Add(numTests)
		for i := 0; i < numTests; i++ {
			thisTestName := testArray[i]
			thisTest := diagnosticTests[thisTestName]
			thisResult := DiagnosticResult{}
			results[i] = &thisResult
			go runDiagnostic(thisTestName, thisTest, results[i], &wg)
		}

		wg.Wait()

		var ns, nf int
		for i := 0; i < numTests; i++ {
			if results[i] != nil {
				s := results[i].Success
				if s != nil {
					if *s {
						ns++
					} else {
						nf++
					}
				}
			}
		}

		if ns > 0 {
			si.SuccessfulTests = &ns
		}
		if nf > 0 {
			si.FailedTests = &nf
		}

		si.DiagnosticResults = results
	}
}

func getOpenFileDescriptorCount() (*int, error) {
	myPID := os.Getpid()
	myFDPath := fmt.Sprintf("/proc/%v/fd", myPID)

	files, err := ioutil.ReadDir(myFDPath)
	if err != nil {
		return nil, err
	}

	openFiles := len(files)

	return &openFiles, nil
}

func formatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	dd := d / (time.Hour * 24)
	d -= dd * (time.Hour * 24)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%03d:%02d:%02d:%02d", dd, h, m, s)
}

var notesBase string = `View memory stats: `
