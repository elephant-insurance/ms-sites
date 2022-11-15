package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// funcs and utilities for testing application routes

type RouteTest struct {
	Method         string
	URL            string
	ExpectedRoute  string
	ExpectedParams map[string]string
}

// Run uses g to resolve its route, then asks r what the last route it chose was
// this won't work in a multi-threaded test! The route still resolves fine, but the last route chosen changes before we test it.
// so for parallel benchmarks ONLY, set the last param to true.
func (rt *RouteTest) Run(r *Router, g http.Handler, t testing.TB, ignoreErrors ...interface{}) {
	if rt == nil || r == nil {
		t.Fatal(`attempt to run nil route test`)
	}

	ignoringErrors := false

	if len(ignoreErrors) > 0 {
		if ie, ok := ignoreErrors[0].(bool); ok && ie {
			// fmt.Fprintln(os.Stderr, `Test is IGNORING ROUTING ERRORS. Only do this if you are running a parallel benchmark!`)
			ignoringErrors = true
		}
	}

	rq, _ := http.NewRequest(rt.Method, rt.URL, nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, rq)

	lr := r.LastRouteSelected
	if (lr == nil || lr.Label != rt.ExpectedRoute) && !ignoringErrors {
		rs := `<nil>`
		if lr != nil && lr.Label != `` {
			rs = lr.Label
		}
		t.Errorf("router failed to resolve %v route for %v %v, chose %v instead", rt.ExpectedRoute, rt.Method, rt.URL, rs)
	}

	lp := r.LastRouteParameters
	for k, v := range rt.ExpectedParams {
		if val, ok := lp[k]; (!ok || val != v) && !ignoringErrors {
			t.Errorf(`failed to resolve %v parameter, expected %v, got %v`, k, v, val)
		}
	}

}
