package cfg

// PreValidated : A config implementing this interface will have its
// PreValidate method run before the basic validate method
// Returns an array of error messages, if any errors are detected.
// The results of PreValidate will appear at the top of the output from validate()
// Use this to run custom validation of your config that can be as complex as needed.
// It's also useful for building required config settings that are constructed from
//  others, like: cfg.QuoteServiceURL = cfg.ServiceBaseURL + cfg.QuoteServicePath
// Return an empty array to signal that everything passed.
type PreValidated interface {
	PreValidate() []string
}

// PostValidated : A config implementing this interface will have its
// PostValidate method run after the basic validate method
// Any results from PreValidate and validate() will be passed
//  in to PostValidate so that the list can be analyzed or altered
// Use append to add new error messages, if appropriate, to the list passed in, and then return it.
// No validation takes place after this function exits, so BE CAREFUL what you do here!
// A poorly-coded PostValidate func can effectively bypass the built-in checks.
type PostValidated interface {
	PostValidate(previousErrors []string) []string
}
