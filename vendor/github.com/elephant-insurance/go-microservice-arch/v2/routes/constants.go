package routes

const (
	// this will work for apps with an app folder next to main.go:
	defaultTestPathPrefix string = `../.`

	errMessageNilRouter          string = `attempted to call method on nil Router`
	errMessageRouterFinalized    string = `attempted to initialize a Router that is already finalized`
	errMessageRouterNotFinalized string = `attempted to use a Router that has not been finalized`

	routeLabelDiagnostics string = `diagnostics`
	routeLabelStatic      string = `static`
)

var (
	HTMLTemplateLeftDelimiter  string = `{{`
	HTMLTemplateRightDelimiter string = `}}`
)
