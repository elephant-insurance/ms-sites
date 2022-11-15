package mbuf

import (
	"fmt"
	"strings"

	enum "github.com/elephant-insurance/enumerations/v2"
)

func newLogglyHTTPMarshaller(logglyKey, appName string, env enum.ServiceEnvironmentID) HTTPMarshaler {
	logglyURL := fmt.Sprintf(logglyURLPattern, logglyKey, strings.ToLower(appName), env)
	return NewDefaultHTTPMarshaler(logglyURL)
}
