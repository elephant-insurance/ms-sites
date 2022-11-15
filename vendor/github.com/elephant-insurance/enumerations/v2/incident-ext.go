package enumerations

func init() {
	// The enumeration itemDict is only used for unmarshaling enumeration values
	// therefore we can easily create an "alias" id by simply adding it here
	// these IDs will UNMARSHAL to the correct enumeration items
	// they will then MARSHAL using the "canonical" (original) item ID
	// TODO: work this into codegen to check for duplicate IDs
	Incident.itemDict[incidentAliasAtFault_atfault] = &incidentAccidentAtFault
	Incident.itemDict[incidentAliasNotAtFault_notatfault] = &incidentAccidentNotAtFault
}

// IncidentCategory collects all the possible values for the incident category meta value
// this can be turned into a full-fledged enumeration if needed, but for now this is better than a list of constants
var IncidentCategory = struct {
	AccidentOrClaim string
	MinorViolation  string
	MajorViolation  string
	OtherViolation  string
}{
	"Accidents/Claims",
	"Minor Violations",
	"Major Violations",
	"Other Violations",
}

// IncidentClass collects all the possible values for the incident class meta value
// this can be turned into a full-fledged enumeration if needed, but for now this is better than a list of constants
var IncidentClass = struct {
	AtFault                 string
	DUI                     string
	Minor                   string
	Major                   string
	NonChargeableAccident   string
	NonChargeableConviction string
}{
	"AFA",
	"DUI",
	"MIN",
	"MAJ",
	"NCA",
	"NCC",
}

const (
	IncidentMetaCategoryKey       string = `Category`
	IncidentMetaClassificationKey string = `Classification`

	incidentAliasAtFault_atfault       string = `atfault`
	incidentAliasNotAtFault_notatfault string = `notatfault`
)
