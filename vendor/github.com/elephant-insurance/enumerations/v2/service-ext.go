package enumerations

// ServiceLogArea collects all the possible log areas for a service
// A log area is a collection of log tables for a set of related services
// These values are used by the Elephant log package to name log tables in Azure.
var ServiceLogArea = struct {
	Microservice    string
	None            string
	RealTimeBidding string
	SingleSearch    string
	Wallboard       string
}{
	`microservice`,
	`none`,
	`rtb`,
	`singlesearch`,
	`wallboard`,
}

const (
	ServiceMetaLogAreaKey string = `LogArea`
)
