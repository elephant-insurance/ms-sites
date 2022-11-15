package bc

import (
	"fmt"
	"net/http"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
)

// CountResponsesByStatus determines whether we keep track of response statuses here in the controller
// note that the diagnostics/v2 package tracks this data, also, if we are using web-services/routes
var CountResponsesByStatus = false

// responseCounter keeps a count of all response codes
// The mutex is just there in case we need to add new response code
var responseCounters = map[int]*clicker.Clicker{
	http.StatusOK:                  {},
	http.StatusNotFound:            {},
	http.StatusBadRequest:          {},
	http.StatusUnauthorized:        {},
	http.StatusUnprocessableEntity: {},
}

// Diagnostics returns a map of response codes to counts for display in diagnostics results
func Diagnostics() map[string]interface{} {
	rtn := map[string]interface{}{}

	if CountResponsesByStatus {
		for code, count := range responseCounters {
			key := keyForStatus(code)
			rtn[key] = count.Clicks
		}
	}

	return rtn
}

// RegisterResponse simply increments the counter for the given response code
// This can be called by controller methods that do not use the base controller's
// render functions so that the stats reflect all controller responses
func RegisterResponse(status int) {
	if CountResponsesByStatus {
		if _, ok := responseCounters[status]; !ok {
			responseCounters[status] = &clicker.Clicker{}
		}

		responseCounters[status].Click(1)
	}
}

// keyForStatus formats the diagnostics output keys
func keyForStatus(status int) string {
	return fmt.Sprintf("%v (%v)", status, http.StatusText(status))
}
