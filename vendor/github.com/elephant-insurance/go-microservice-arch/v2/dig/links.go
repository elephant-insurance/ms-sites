package dig

// generateDiagnosticsLinks uses the request URL (if any) to generate links to dig pages
func generateDiagnosticsLinks(basePath string) map[string]string {
	//lw := log.ForFunc(context.Background())

	return map[string]string{
		`Base diagnostics: `: basePath,
		`Run tests: `:        basePath + `?runTests=true`,
		`Memory stats: `:     basePath + `/` + memStatsPath,
		`Timing stats: `:     basePath + `/` + timingsPath,
		`Profile: `:          basePath + `/` + pprofPath,
	}
}

const (
	baseLinkPattern     = `%v://%v/%v/diagnostics`
	relativeLinkPattern = `/%v/diagnostics`
)
