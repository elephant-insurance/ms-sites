package log

const (
	EnvironmentSettingPrefixMSL string = `MSVCLOG_`
	EnvironmentSettingPrefixRTB string = `MSRTBLOG_`

	MessageCalled    string = `called`
	MessageCompleted string = `completed`
	MessageStarted   string = `started`

	azureMicroserviceLogTimeField      string = `time`
	azureMicroserviceEventLogTimeField string = `Time`
	azureRTBLogTimeField               string = `timestamp`
	azureTestSharedKey                 string = `eJdtXPBmKLynulNXzv0zqiN5tXPa6L6K19qpbmQs4UeonOqDKUcWD3RaV0huKlYDRF/K8AzbIjxq/Zi64M7y7g==`
	azureTestWorkspaceID               string = `a4c56fcf-27b5-47d2-b207-3d28d35f57c9`

	ErrorBadNewLogArgs        string = `you must supply a valid Instance Name, Level, and Environment to create a new Microservice Log`
	ErrorBadServiceID         string = `you must supply a valid service id with a valid log table area`
	ErrorNilEntry             string = `attempted to log a nil log entry`
	ErrorNilSettings          string = `attempted to initialize a Log with nil Settings`
	ErrorAzureSettingsMissing string = `you must supply a valid Azure WorkspaceID and Shared Key`
	ErrorEmptyInstanceName    string = `instance name cannot be empty`
	ErrorEmptyLogType         string = `log type cannot be empty`

	unknownFuncName string = `<unknown>`

	// max length for string fields to be sent to Azure
	stringLengthLimitShort int = 64
	stringLengthLimitLong  int = 1000

	// These are no longer really required but are here for backward compatibility
	SettingSendDirectToAzure  string = `SendDirectToAzure`
	SettingSendDirectToLoggly string = `SendDirectToLoggly`
	SettingSendToNone         string = `SendToNone`
	SettingSendToBoth         string = `SendToBoth`
)

var (
	// use these for making bool pointers inline
	truthy, falsy         = true, false
	contextListKey string = `contextListKey`
	// making these private for now
	fieldNameMSAccountID    string = `account`
	fieldNameMSBrand        string = `brand`
	fieldNameMSBusinessType string = `BusinessType`
	fieldNameMSChannel      string = `channel`
	fieldNameMSCode         string = `code`
	fieldNameMSCount        string = `count`
	fieldNameMSCoverage     string = `coverage`
	fieldNameMSDate         string = `date`
	fieldNameMSDetail       string = `detail`
	fieldNameMSDriver       string = `driver`
	fieldNameMSEnvironment  string = `env`
	fieldNameMSHostName     string = `host`
	fieldNameMSID           string = `id`
	fieldNameMSInstanceName string = `InstanceName`
	fieldNameMSIPAddress    string = `ip-address`
	fieldNameMSJobID        string = `job`
	fieldNameMSLevel        string = `level`
	fieldNameMSName         string = `name`
	fieldNameMSNumber       string = `number`
	fieldNameMSPartner      string = `partner`
	fieldNameMSPolicyID     string = `policy`
	fieldNameMSPublicID     string = `public-id`
	fieldNameMSQuoteID      string = `quote-id`
	fieldNameMSQuoteNumber  string = `quote-number`
	fieldNameMSServiceID    string = `app`
	fieldNameMSState        string = `state`
	fieldNameMSVehicle      string = `vehicle`
	fieldNameMSVIN          string = `vin`

	fieldNameMSCommonErrorCode     string = `error-code`
	fieldNameMSElapsedMicroseconds string = `microseconds`
	fieldNameMSError               string = `err`
	fieldNameMSFunction            string = `func`
	fieldNameMSHTTPMethod          string = `http-method`
	fieldNameMSHTTPStatus          string = `http-status`
	fieldNameMSMessage             string = `msg`
	fieldNameMSStack               string = `stack`
	fieldNameMSTime                string = `time`
	fieldNameMSURL                 string = `url`
)
