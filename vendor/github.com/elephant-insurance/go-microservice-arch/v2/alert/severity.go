package alert

type AlertSeverity int

const (
	AlertSeverityUnknown AlertSeverity = iota
	AlertSeverityLow
	AlertSeverityMedium
	AlertSeverityHigh
	AlertSeverityCritical
)

// String gives a human readable severity string
func (s AlertSeverity) String() string {
	return [...]string{"Unknown", "Low", "Medium", "High", "Critical"}[s]
}
