package log

import (
	"errors"
	"fmt"

	enum "github.com/elephant-insurance/enumerations/v2"
)

type logTableType string

const (
	logTypeAppLog    logTableType = `log`
	logTypeEventLog  logTableType = `event`
	logTypeTimingLog logTableType = `timing`
)

func azureTableName(logType logTableType, service *enum.ServiceID) (string, error) {
	if logType == `` {
		return ``, errors.New(`you must supply a valid log type (log, event, or timing)`)
	}

	svc := enum.Service.ByID(service)
	if svc == nil || svc.LogArea == `` || (!service.Equals(&enum.Service.Test.ID) && svc.LogArea == enum.ServiceLogArea.None) {
		return ``, errors.New(ErrorBadServiceID)
	}
	return fmt.Sprintf(`%v_%v`, svc.LogArea, logType), nil
}
