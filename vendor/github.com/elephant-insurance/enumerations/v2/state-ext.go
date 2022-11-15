package enumerations

func (e *EnumState) IsInsurableState(id *StateID) bool {
	if e == nil || id == nil || string(*id) == "" {
		return false
	}

	lookupState := e.ByID(id)
	if lookupState == nil {
		return false
	}

	switch lookupState.ID {
	case stateVirginiaID:
		return true
	case stateTexasID:
		return true
	case stateMarylandID:
		return true
	case stateIllinoisID:
		return true
	case stateIndianaID:
		return true
	case stateTennesseeID:
		return true
	case stateOhioID:
		return true
	case stateGeorgiaID:
		return true
	default:
		return false
	}
}
