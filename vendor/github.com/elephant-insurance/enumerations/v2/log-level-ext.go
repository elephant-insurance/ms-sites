package enumerations

// IsMoreUrgentThan is a convenience function for logging
// we set a level for the logger and a level for the message
// if the logger's level IsMoreUrgentThan the message's level, we don't log it
func (l *EnumLogLevelItem) IsMoreUrgentThan(otherLevel *EnumLogLevelItem) bool {
	if l == nil {
		// nil is not more urgent than anything
		return false
	}
	if otherLevel == nil {
		// anything is more urgent than nil
		return true
	}

	return l.SortOrder < otherLevel.SortOrder
}
