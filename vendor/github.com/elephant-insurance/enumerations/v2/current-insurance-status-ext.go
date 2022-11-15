package enumerations

// CurrentInsuranceStatusBasic is a sub-collection of EnumCurrentInsuranceStatus for display
var CurrentInsuranceStatusBasic = &EnumCurrentInsuranceStatus{
	Items: []*EnumCurrentInsuranceStatusItem{
		&currentInsuranceStatusOwnPolicy,
		&currentInsuranceStatusAnothersPolicy,
	},
	Name:           "EnumCurrentInsuranceStatusBasic",
	OwnPolicy:      &currentInsuranceStatusOwnPolicy,
	AnothersPolicy: &currentInsuranceStatusAnothersPolicy,

	itemDict: map[string]*EnumCurrentInsuranceStatusItem{
		string(currentInsuranceStatusOwnPolicyID):      &currentInsuranceStatusOwnPolicy,
		string(currentInsuranceStatusAnothersPolicyID): &currentInsuranceStatusAnothersPolicy,
	},
}

// CurrentInsuranceReason is a sub-collection of EnumCurrentInsuranceStatus for display
var CurrentInsuranceReason = &EnumCurrentInsuranceStatus{
	Items: []*EnumCurrentInsuranceStatusItem{
		&currentInsuranceStatusDeployedOverseas,
		&currentInsuranceStatusExpiredWithin30Days,
		&currentInsuranceStatusExpiredOver30Days,
		&currentInsuranceStatusNoInsuranceRequired,
	},
	Name:                "EnumCurrentInsuranceReason",
	DeployedOverseas:    &currentInsuranceStatusDeployedOverseas,
	ExpiredWithin30Days: &currentInsuranceStatusExpiredWithin30Days,
	ExpiredOver30Days:   &currentInsuranceStatusExpiredOver30Days,
	NoInsuranceRequired: &currentInsuranceStatusNoInsuranceRequired,

	itemDict: map[string]*EnumCurrentInsuranceStatusItem{
		string(currentInsuranceStatusDeployedOverseasID):    &currentInsuranceStatusDeployedOverseas,
		string(currentInsuranceStatusExpiredWithin30DaysID): &currentInsuranceStatusExpiredWithin30Days,
		string(currentInsuranceStatusExpiredOver30DaysID):   &currentInsuranceStatusExpiredOver30Days,
		string(currentInsuranceStatusNoInsuranceRequiredID): &currentInsuranceStatusNoInsuranceRequired,
	},
}
