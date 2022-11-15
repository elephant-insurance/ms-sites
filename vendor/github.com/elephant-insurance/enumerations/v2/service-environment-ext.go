package enumerations

func (i *EnumServiceEnvironmentItem) IsProduction() bool {
	return i != nil && i.Prod == "true"
}
