package resp

// PCResponseErrorer ...
type PCResponseErrorer interface {
	Error() string
	JobNumber() string
}

// PCResponseError ...
type PCResponseError struct {
	// Err       error
	message   string
	jobNumber string
}

// NewPCResponseError is created
func NewPCResponseError(message string, jobNumber string) *PCResponseError {
	return &PCResponseError{message: message, jobNumber: jobNumber}
}

// PCResponseError ...
func (e *PCResponseError) Error() string {
	return e.message
}

// JobNumber ...
func (e *PCResponseError) JobNumber() string {
	return e.jobNumber
}

// JsonApiError ... replaces the type from the old gson api package
type JsonApiError struct {
	ID          string      `json:"id,omitempty"`
	Status      string      `json:"status,omitempty"`
	Code        string      `json:"code,omitempty"`
	Title       string      `json:"title,omitempty"`
	Detail      string      `json:"detail,omitempty"`
	Source      string      `json:"source,omitempty"`
	ErrorDetail interface{} `json:"ErrorDetail,omitempty"`
}

// ValidationError ... Simple struct to store the Message & Key of a validation error
type ValidationError struct {
	Message, Key string
}

// Error ...
func (e *JsonApiError) Error() string {
	return e.Detail
}
