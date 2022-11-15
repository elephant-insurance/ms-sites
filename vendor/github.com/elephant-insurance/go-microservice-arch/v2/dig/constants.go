package dig

const (
	alertMsgMemPeakAllocOneLastDay                 = "RAM allocated to the application heap achieved a new peak in the past 24 hours."
	alertMsgMemPeakAllocMultiLastDay               = "RAM allocated to the application heap achieved a new peak more than once in the past 24 hours."
	alertMsgMemPeakHeapObjOneLastDay               = "Objects in the application heap achieved a new peak in the past 24 hours."
	alertMsgMemPeakHeapObjMultiLastDay             = "Objects in the application heap achieved a new peak more than once in the past 24 hours."
	alertMsgMemPeakSysOneLastDay                   = "The operating system has reserved additional RAM for this app in the past 24 hours."
	alertMsgMemPeakSysMultiLastDay                 = "The operating system has reserved additional RAM for this app more than once in the past 24 hours."
	alertMsgMemGrowthAlloc50PercentOverBaseline    = "Heap RAM allocation has increased at least 100% over baseline"
	alertMsgMemGrowthAlloc100PercentOverBaseline   = "Heap RAM allocation has increased 50% over baseline"
	alertMsgMemGrowthHashObj50PercentOverBaseline  = "Heap object allocation has increased at least 100% over baseline"
	alertMsgMemGrowthHashObj100PercentOverBaseline = "Heap object allocation has increased 50% over baseline"
	alertMsgMemGrowthSys100PercentOverBaseline     = "System reserved RAM has increased at least 100% over baseline"
	alertMsgMemGrowthSys50PercentOverBaseline      = "System reserved RAM has increased 50% over baseline"
	alertMsgMemGrowthSysAnyOverBaseline            = "System reserved RAM has increased over baseline"
	alertMsgMemGrowthAllocMonotonicLastDay         = "RAM allocated to the heap has increased but not decreased in the past 24 hours"
	alertMsgMemGrowthHashObjMonotonicLastDay       = "Count of objects on the heap has increased but not decreased in the past 24 hours"

	ErrorMessageCannotDetermineWorkingDirectory        = "Could not determine working directory for running process"
	ErrorMessageCannotLoadFileHandle                   = "Could not load file handle for running file"
	ErrorMessageCannotLoadFileHandleWarnDT             = "Bypassing crash in DEV/TEST environment.\n * IF THIS APP IS RUNNING IN A CONTAINER THEN IT IS MISCONFIGURED AND WILL FAIL WHEN DEPLOYED *"
	ErrorMessageInvalidPath                            = "Invalid path for running file"
	ErrorMsgCannotDetermineHostname                    = "Could not determine hostname"
	ErrorMsgEmptyBranch                                = "Branch cannot be empty"
	ErrorMsgEmptyBuild                                 = "Build cannot be empty"
	ErrorMsgEmptyCommit                                = "Commit cannot be empty"
	ErrorMsgEmptyDateBuilt                             = "DateBuilt cannot be empty"
	ErrorMsgEmptyEnvironment                           = "Environment cannot be empty"
	ErrorMsgEmptyImageTag                              = "ImageTag cannot be empty"
	ErrorMsgEmptyProcName                              = "processName cannot be empty"
	ErrorMsgSerializeResults                           = "Error serializing result set"
	ErrorMsgWritingResult                              = "Error writing result set"
	ErrorMsgNilURL                                     = "Nil URL sent to ServeHTTP"
	ErrorMsgNilRequest                                 = "Nil request sent to ServeHTTP"
	ErrorMsgFailedToCreateMemMaxBuffer                 = "error returned when attempting to create MemMaxBuffer"
	ErrorMsgFailedToCreateMemStatBuffer                = "error returned when attempting to create memStatBuffer"
	ErrorMsgFailedToCreateTimingBuffer                 = "error returned when attempting to create TimingBuffer"
	ErrorMsgGinDebugMode                               = "Gin cannot be in debug mode if the environment is production"
	ErrorMsgGinTestMode                                = "Gin cannot be in test mode unless the environment is dev or test"
	dummyPeakAlloc                              uint64 = 1234
	dummyPeakHeapObjects                        uint64 = 2345
	dummyPeakSequenceNumber                     uint64 = 3456
	dummyPeakSys                                uint64 = 4567
	minutesToWaitForMemBaseline                        = 60
	mimeTypeCSV                                        = "text/plain"
	paramNameAfter                                     = "after"
	paramNameCSV                                       = "csv"
)

// peakJumpBufferSize is how many memory peaks we keep a record of
const peakJumpBufferSize = 1000

const EnvironmentSettingPrefix string = `MSVCDIG_`
