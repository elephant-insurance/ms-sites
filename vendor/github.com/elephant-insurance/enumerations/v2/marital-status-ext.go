package enumerations

import "strings"

// This extension adds properties to the marital status enumeration and its members
// it is in a separate file to allow re-generating the enumeration without losing this code

var alternativeMaritalStatusMap *map[string]*MaritalStatusID

func initializeAlternativeMaritalStatusMap() {
	if alternativeMaritalStatusMap != nil {
		return
	}

	newMap := map[string]*MaritalStatusID{}

	for i := 0; i < len(MaritalStatus.Items); i++ {
		thisItem := MaritalStatus.Items[i]
		newMap[string(thisItem.ID)] = &thisItem.ID
		if thisItem.AlternativeKeys == "" {
			continue
		}

		tokens := strings.Split(thisItem.AlternativeKeys, ",")
		for j := 0; j < len(tokens); j++ {
			thisToken := strings.ToLower(strings.TrimSpace(tokens[j]))
			if thisToken != "" {
				newMap[thisToken] = thisItem.ID.Clone()
			}
		}
	}

	alternativeMaritalStatusMap = &newMap
}

// ByAlternativeKey get matrial status using alternative keys
func (e *EnumMaritalStatus) ByAlternativeKey(key *string) *MaritalStatusID {
	if e == nil || len(e.Items) < 1 || key == nil || *key == "" {
		return nil
	}

	initializeAlternativeMaritalStatusMap()

	itemKey := strings.ToLower(strings.TrimSpace(*key))
	if rtn, ok := (*alternativeMaritalStatusMap)[itemKey]; ok {
		return rtn
	}

	return nil
}
