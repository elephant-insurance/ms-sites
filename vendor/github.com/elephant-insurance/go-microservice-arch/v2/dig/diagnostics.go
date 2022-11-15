package dig

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	olog "log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var client *http.Client

const defaultHTTPTimeout = 5000 * time.Millisecond

// DiagnosticTest is a function type for running arbitrary diagnostic tests
// Any function that can be cast to this type may be run as a diagnostic test
type DiagnosticTest func() (DiagnosticResult, error)

// AddDiagnosticTest appends a new diagnostic test to be run with each invocation.
// The function will run if the runTests query string parameter is specified.
// It should return a useful result, which will be displayed on the diagnostics page.
func AddDiagnosticTest(name string, t DiagnosticTest) {
	if diagnosticTests == nil {
		olog.Fatal("attempted to add a test before initializing diagnostics")
	}
	testArray = append(testArray, name)
	diagnosticTests[name] = t
}

func runDiagnostic(n string, t DiagnosticTest, r *DiagnosticResult, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if recov := recover(); recov != nil {
			f := false
			d := fmt.Sprintf("Recovered from PANIC: %v", recov)
			// printf("Recovered from PANIC: %v", recov)
			pr := DiagnosticResult{
				Name:        n,
				Success:     &f,
				Description: &d,
			}
			*r = pr
		}
	}()
	startTime := time.Now()
	// we could just have the test func set result.Error,
	// but returning it separately is more idiomatic
	result, err := t()
	result.Name = n
	if err != nil {
		result.Error = err
	}
	elapsed := time.Since(startTime)
	result.Elapsed = fmt.Sprintf("%v", elapsed.Microseconds())
	*r = result
}

// AddConnectionTest adds a simple connectivity test
// The url parameter should start with a protocol specifier (e.g., "http://").
// It will attempt to GET the specified URL. If the GET returns a 200, the result will have Success = true.
// Any code other than 200 will result in Success = false.
// This method is deprecated. For new code, please use the equivalent AddGetConnectionTest(url string)
func AddConnectionTest(url string) {
	AddConnectionTestWithHeaders(url, http.MethodGet, nil)
}

// AddGetConnectionTest adds a simple connectivity test
// The url parameter should start with a protocol specifier (e.g., "http://").
// It will attempt to GET the specified URL. If the GET returns a 200, the result will have Success = true.
// Any code other than 200 will result in Success = false.
func AddGetConnectionTest(url string) {
	AddConnectionTestWithHeaders(url, http.MethodGet, nil)
}

// AddHeadConnectionTest adds a simple connectivity test
// The url parameter should start with a protocol specifier (e.g., "http://").
// It will attempt to HEAD the specified URL. If the HEAD request returns a 200, the result will have Success = true.
// Any code other than 200 will result in Success = false.
func AddHeadConnectionTest(url string) {
	AddConnectionTestWithHeaders(url, http.MethodHead, nil)
}

// AddNamedConnectionTest adds a simple connectivity test with sensitive values
// DO NOT use this method for tests without sensitive URL parameters, use AddConnectionTest instead
// The url parameter should start with a protocol specifier (e.g., "http://").
// It will attempt to GET the specified URL. If the GET returns a 200, the result will have Success = true.
// Any code other than 200 will result in Success = false.
// The name parameter overrides the automatic naming of the test, useful in case there are
// sensitive values in the connection URL. Supplying a name will suppress output of the URL
// whether the test is run or not.
// This method is deprecated. For new code, please use the equivalent AddNamedGetConnectionTest(url, name string).
func AddNamedConnectionTest(url, name string) {
	AddNamedConnectionTestWithHeaders(url, http.MethodGet, nil, name)
}

// AddNamedGetConnectionTest adds a simple connectivity test with sensitive values
// DO NOT use this method for tests without sensitive URL parameters, use AddGetConnectionTest instead
// The url parameter should start with a protocol specifier (e.g., "http://").
// It will attempt to GET the specified URL. If the GET returns a 200, the result will have Success = true.
// Any code other than 200 will result in Success = false.
// The name parameter overrides the automatic naming of the test, useful in case there are
// sensitive values in the connection URL. Supplying a name will suppress output of the URL
// whether the test is run or not.
func AddNamedGetConnectionTest(url, name string) {
	AddNamedConnectionTestWithHeaders(url, http.MethodGet, nil, name)
}

// AddNamedHeadConnectionTest adds a simple connectivity test with sensitive values
// DO NOT use this method for tests without sensitive URL parameters, use AddGetConnectionTest instead
// The url parameter should start with a protocol specifier (e.g., "http://").
// It will attempt to HEAD the specified URL. If the HEAD request returns a 200, the result will have Success = true.
// Any code other than 200 will result in Success = false.
// The name parameter overrides the automatic naming of the test, useful in case there are
// sensitive values in the connection URL. Supplying a name will suppress output of the URL
// whether the test is run or not.
func AddNamedHeadConnectionTest(url, name string) {
	AddNamedConnectionTestWithHeaders(url, http.MethodHead, nil, name)
}

// AddConnectionTestWithHeaders adds a simple connectivity test
// The url parameter should start with a protocol specifier (e.g., "http://").
// It will attempt to contact the specified URL using the HTTP method specified.
// Note that while any HTTP method may be valid, the method does not support passing
// a request body. So methods that depend on the request body (PUT, POST, PATCH)
// will likely not work as desired. GET, OPTIONS, and HEAD should work fine.
// If the request returns a 200, the result will have Success = true.
// Any code other than 200 will result in Success = false.
// If requestHeaders is not nil, headers will be added to the request when testing
func AddConnectionTestWithHeaders(url, method string, requestHeaders map[string][]string) {
	testName := fmt.Sprintf("Verify GET %v", url)
	AddNamedConnectionTestWithHeaders(url, method, requestHeaders, testName)
}

// AddNamedConnectionTestWithHeaders adds a simple connectivity test with sensitive values
// DO NOT use this method for tests without sensitive URL parameters, use AddConnectionTestWithHeaders instead
// The url parameter should start with a protocol specifier (e.g., "http://").
// It will attempt to contact the specified URL using the HTTP method specified.
// Note that while any HTTP method may be valid, the method does not support passing
// a request body. So methods that depend on the request body (PUT, POST, PATCH)
// will likely not work as desired. GET, OPTIONS, and HEAD should work fine.
// If the request returns a 200, the result will have Success = true.
// Any code other than 200 will result in Success = false.
// If requestHeaders is not nil, headers will be added to the request when testing
// The name parameter overrides the automatic naming of the test, useful in case there are
// sensitive values in the connection URL. Supplying a name will suppress output of the URL
// whether the test is run or not.
func AddNamedConnectionTestWithHeaders(url, method string, requestHeaders map[string][]string, name string) {
	test := buildConnectionTestFunc(url, name, method, requestHeaders)

	AddDiagnosticTest(name, test)
}

// buildConnectionTestFunc builds a simple connection tester func
func buildConnectionTestFunc(url, name, method string, requestHeaders map[string][]string) DiagnosticTest {
	test := func() (DiagnosticResult, error) {
		testDesc := url
		suppressError := false
		if len(name) > 0 {
			testDesc = name
			suppressError = true
		}
		rtn := NewResult()
		rq, _ := http.NewRequest(method, url, nil)
		rq.Close = true
		if len(requestHeaders) > 0 {
			for thisKey, thisVal := range requestHeaders {
				if thisKey != "" && len(thisVal) > 0 {
					rq.Header.Set(thisKey, strings.Join(thisVal, ","))
				}
			}
		}

		if client == nil {
			client = &http.Client{
				Timeout: defaultHTTPTimeout,
			}
		}

		res, err := client.Do(rq)
		if res != nil {
			defer res.Body.Close()
		}

		if err != nil {
			if !suppressError {
				rtn.Fail().SetDescriptionf("Error attempting %v %v: \n%v", method, testDesc, err.Error())
			} else {
				rtn.Fail().SetDescriptionf("Error attempting %v %v, see log for error details", method, testDesc)
				// printf("Error attempting %v %v: \n%v", method, testDesc, err.Error())
				err = nil
			}
			return *rtn, err
		}

		if res.StatusCode != http.StatusOK {
			rtn.Fail().SetDescriptionf("%v %v returned non-200 code: %v", method, testDesc, res.StatusCode)
			return *rtn, nil
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			rtn.Fail().SetDescriptionf("Error attempting to read response from %v %v: \n%v", method, testDesc, err.Error())
			return *rtn, err
		}

		rtn.Succeed().SetDescriptionf("%v %v returned 200 OK", method, testDesc)
		return *rtn, nil
	}

	return test
}

// AddDatabasePingTest adds a database ping to the list of diagnostic tests
func AddDatabasePingTest(db *sql.DB, dbname string) {
	testName := fmt.Sprintf("Verify database ping to %v", dbname)
	test := func() (DiagnosticResult, error) {
		rtn := NewResult()
		if db == nil {
			rtn.Fail().SetDescriptionf("Error attempting to ping database %v: db handle is nil", dbname)
			return *rtn, fmt.Errorf("error attempting to ping database %v: db handle is nil", dbname)
		}

		err := db.Ping()
		if err != nil {
			rtn.Fail().SetDescriptionf("Error attempting to ping database %v: \n%v", dbname, err.Error())
			return *rtn, err
		}

		rtn.Succeed().SetDescriptionf("Ping of database %v returned no error", dbname)
		return *rtn, nil
	}

	AddDiagnosticTest(testName, test)
}

// DBGetter is a simple func type that returns a pointer to a sql.db
type DBGetter func() (*sql.DB, error)

// AddDatabasePingTestFunction adds a simple database ping test
// The getter param is a function that returns a pointer to the sql.db that we want to test
// The dbname is a human-readable identifier for clarity in logging
func AddDatabasePingTestFunction(getter DBGetter, dbname string) {
	testName := fmt.Sprintf("Verify database ping to %v", dbname)
	test := func() (DiagnosticResult, error) {
		rtn := NewResult()
		db, err := getter()
		if db == nil || err != nil {
			rtn.Fail().SetDescriptionf("Error attempting to ping database %v: db handle is invalid\n Error: %v", dbname, err)
			return *rtn, fmt.Errorf("error attempting to ping database %v: db handle is invalid\n Error: %v", dbname, err)
		}

		err = db.Ping()

		if err != nil {
			rtn.Fail().SetDescriptionf("Error attempting to ping database %v: \n%v", dbname, err.Error())
			return *rtn, err
		}

		rtn.Succeed().SetDescriptionf("Ping of database %v returned no error", dbname)
		return *rtn, nil
	}

	AddDiagnosticTest(testName, test)
}
