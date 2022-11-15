package routes

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/elephant-insurance/go-microservice-arch/v2/log"
)

// AllowTrailingSlash controls whether paths match with a trailing slash
// if true, each path path regex is appended to match a trailing slash
var AllowTrailingSlash = true

const RouteLabelKey = "elephant-route-label"

// Route encapsulates a single URL => handler pair
type Route struct {
	Handler    gin.HandlerFunc
	Label      string
	PathParams []string
	PathRegex  *regexp.Regexp
	RawPath    string
	// Silent marks this route as one for which we do not wish to log requests (static files, etc.)
	Silent bool
}

// handles checks whether the route can handle the submitted URL
// If the URL matches the Route's PathRegex, set the appropriate path variables in the context
func (r *Route) handles(c *gin.Context) (bool, map[string]string) {
	if r == nil || r.PathRegex == nil || c == nil || c.Request == nil || c.Request.URL == nil || c.Request.URL.Path == "" {
		return false, nil
	}
	parmMap := map[string]string{}
	url := c.Request.URL.Path

	match := r.PathRegex.FindStringSubmatch(url)
	rtn := false

	if match != nil {
		rtn = true
		for i, name := range r.PathRegex.SubexpNames() {
			if i != 0 && name != "" {
				// we assume that nothing else has been adding params to the context
				c.Params = append(c.Params, gin.Param{Key: name, Value: match[i]})
				parmMap[name] = match[i]
			}
		}
		c.Params = append(c.Params, gin.Param{Key: RouteLabelKey, Value: r.Label})
		parmMap[RouteLabelKey] = r.Label
	}

	return rtn, parmMap
}

func makeRoute(label, path string, handler gin.HandlerFunc) *Route {
	// lw := log.ForFunc("web-services.routes.makeRoute").WithConsoleField("label", label).WithConsoleField("path", path)
	regex, params := buildRegexp(path)

	return &Route{
		Handler:    handler,
		Label:      label,
		PathParams: params,
		PathRegex:  regex,
		RawPath:    path,
	}
}

// buildRegexp builds a regular expression from a path specification
// Named path parameters are replaced by named regexp captures
// Path parameter names, if any, are returned as an array
// So that we can extract any named parameters at the same time that we match it
// THIS METHOD WILL THROW A FATAL ERROR IF IT CANNOT BUILD A REGEXP
// Therefore, DO NOT call it after app startup
func buildRegexp(path string) (*regexp.Regexp, []string) {
	lw := log.ForFunc(nil).WithConsoleField("path", path)
	lw.Debug("building regex for path")

	if path == "" {
		return nil, nil
	}

	paramNames := []string{}

	// break the path into individual path elements
	pathTokens := strings.Split(path, slash)

	// this keeps track of our regex as we build it
	pathRegexStr := "(?i)"

	// now we cycle over the path elements
	// empty elements are ignored
	// parameter elements (with a ":" prefix) are converted to a named capture
	// regular elements (strings not starting with ":") are added to the regex as-is
	for i := 0; i < len(pathTokens); i++ {
		thisToken := pathTokens[i]
		if thisToken == "" {
			continue
		}

		if len(thisToken) > 1 && strings.HasPrefix(thisToken, ":") {
			// add named capture
			paramName := strings.ToLower(thisToken[1:])
			pathRegexStr += fmt.Sprintf("/(?P<%v>[^/]*)", paramName)
			paramNames = append(paramNames, paramName)
			continue
		}

		if len(path) > 1 && strings.HasPrefix(thisToken, "*") {
			// greedy capture means we wrap up the rest into one big param:
			paramName := strings.ToLower(thisToken[1:])
			pathRegexStr += fmt.Sprintf("/(?P<%v>.*)", paramName)
			paramNames = append(paramNames, paramName)
			break
		}

		pathRegexStr += strings.ToLower(slash + thisToken)
	}

	// add "end of string" token to make sure we only match paths that end where our regex ends
	// otherwise the path spec /customers/:id would match the request path "/customers/123/invoices/445"
	// and the extra path parameters are lost
	if AllowTrailingSlash {
		pathRegexStr += "[/]?$"
	} else {
		pathRegexStr += "$"
	}

	// the regex is compiled before returning, or we die
	rtn, err := regexp.Compile(pathRegexStr)
	lw.IfError(err).Fatal("failed to compile route regexp")

	return rtn, paramNames
}

// timingParams is a convenience function that returns pointers to the string values we need for a route timing
func (r *Route) timingParams() (*string, *string) {
	if r == nil {
		return nil, nil
	}

	label := r.Label
	rawPath := r.RawPath

	return &label, &rawPath
}
