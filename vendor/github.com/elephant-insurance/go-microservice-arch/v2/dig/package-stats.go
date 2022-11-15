package dig

// PackageStats is a func type for other packages to use
// to report relevant stats to the diagnostics page
type PackageStats func() map[string]interface{}

var packageStats map[string]PackageStats

// AddPackageStats adds a package stats function to the diagnostics output
// Use this to report additional diagnostics from other packages in your app
func AddPackageStats(label string, f PackageStats) {
	packageStats[label] = f
}
