package mbuf

const (
	AzureServiceBusURIPattern string = `https://%v.servicebus.windows.net/%v/messages`

	absoluteMaxRelayRetries int = 50

	azureAuthBasePattern       string = `SharedKey %v:`
	azureURLPattern            string = `https://%v.ods.opinsights.azure.com/api/logs?api-version=2016-04-01`
	azureDefaultLogType        string = `microservice_log`
	azureServiceBusContentType string = `application/vnd.microsoft.servicebus.json`

	logglyURLPattern string = `https://logs-01.loggly.com/bulk/%v/tag/%v,%v`

	// Increased by Sizemore, 20 July 2022, to improve our logging reliability.
	// These values should get us over more network issues while still being safe.
	// The total outage length the service can endure without losing messages is:
	// (ClientTimeoutSecs * (MaxRetries + 1)) + (RetryWaitSecs * MaxRetries)
	// Prior to 22 July: (20 * 4) + (10 * 3) = 110 seconds outage
	// Now: (20 * 6) + (60 * 5) = 420 seconds outage
	defaultClientTimeoutSecs int = 20
	defaultMaxBufferLength   int = 50 // Increasing this means fewer, larger requests. Was: 25
	defaultMaxRetries        int = 5  // Increasing this gives the network longer to recover at the expense of a little RAM during outages. Was: 3
	defaultPollIntervalSecs  int = 30
	defaultRetryWaitSecs     int = 60 // Increasing this gives the network longer to recover at the expense of a little RAM during outages. Was: 10

	diagnosticsFieldAzureQueueOrTopic string = `AzureQueueOrTopic`
	diagnosticsFieldAzureResourceName string = `AzureResourceName`
	diagnosticsFieldAzureWorkspaceID  string = `workspaceID`
	diagnosticsFieldBadResponses      string = `badResponses`
	diagnosticsFieldErrorResponses    string = `errorResponses`
	diagnosticsFieldFailuresToMarshal string = `failuresToMarshal`
	diagnosticsFieldFailuresToRender  string = `failuresToRender`
	diagnosticsFieldLogType           string = `logType`
	diagnosticsFieldMessagesFlushed   string = `messagesFlushed`
	diagnosticsFieldSuccessfulSends   string = `successfulSends`
	diagnosticsFieldTotalFailures     string = `totalFailures`
	diagnosticsFieldTotalRetries      string = `totalRetries`
	diagnosticsFieldTotalSendAttempts string = `totalSendAttempts`
	diagnosticsFieldTotalShutdowns    string = `totalShutdowns`

	errMsgFailedToCreateRequest string = `failed to create new http request`
)
