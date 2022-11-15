package dig

import "time"

// Settings provides all needed settings for the diagnostics package
// It is designed to be added as a member of the app config struct
type Settings struct {
	AppName    string    `json:"AppName,omitempty" yaml:"AppName" config:"optional"`
	AppModTime time.Time `json:"AppModTime,omitempty" yaml:"AppModTime" config:"optional"`
	Branch     string    `json:"Branch,omitempty" yaml:"Branch" config:"optional"`
	// BufferMax if set sets the maximum size for memstats and timings buffers
	BufferMax        *int32                    `json:"BufferMax" yaml:"BufferMax" config:"optional"`
	Build            string                    `json:"Build,omitempty" yaml:"Build" config:"optional"`
	Commit           string                    `json:"Commit,omitempty" yaml:"Commit" config:"optional"`
	DateBuilt        string                    `json:"DateBuilt,omitempty" yaml:"DateBuilt" config:"optional"`
	Diagnostics      map[string]DiagnosticTest `json:"-" yaml:"-" config:"optional"`
	Hostname         string                    `json:"Hostname,omitempty" yaml:"Hostname" config:"optional"`
	ImageTag         string                    `json:"ImageTag,omitempty" yaml:"ImageTag" config:"optional"`
	InstanceName     string                    `json:"InstanceName,omitempty" yaml:"InstanceName" config:"optional"`
	MemStatBufferMax *int32                    `json:"MemStatBufferMax" yaml:"MemStatBufferMax" config:"optional"`
	ProcName         string                    `json:"ProcName,omitempty" yaml:"ProcName" config:"optional"`
	StartTime        *time.Time                `json:"StartTime,omitempty" yaml:"-" config:"optional"`
	TestArray        []string                  `json:"-" yaml:"-" config:"optional"`
	TimingBufferMax  *int32                    `json:"TimingBufferMax" yaml:"TimingBufferMax" config:"optional"`
	VerboseLog       bool                      `json:"VerboseLog,omitempty" yaml:"VerboseLog" config:"optional"`
}

// MemStatIntervalSeconds is how often we should take a snapshot of our RAM usage
var MemStatIntervalSeconds = 60

// ForTesting fills out a Settings struct for quick, easy testing without a lot of setup
func (s *Settings) ForTesting() *Settings {
	rn := time.Now()
	s.AppModTime = rn
	s.Branch = "test_branch"
	s.Build = "test_build"
	s.Commit = "test_commit"
	s.DateBuilt = rn.Format("2006-01-02")
	s.Hostname = "test_hostname"
	s.ImageTag = "test_image_tag"
	s.ProcName = "test_proc_name"
	s.StartTime = &rn

	return s
}

// GetEnvironmentSetting implements the Overridable interface for the configuration package
// Given a field name, it returns the name of the corresponding environment variable
func (s Settings) GetEnvironmentSetting(fieldName string) string {
	return EnvironmentSettingPrefix + fieldName
}
