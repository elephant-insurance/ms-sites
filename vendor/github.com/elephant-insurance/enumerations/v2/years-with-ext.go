package enumerations

import "strings"

// it is in a separate file to allow re-generating the enumeration without losing this code

var alternativeYearsWith *map[string]*YearsWithID

func initializeAlternativeYearsWith() {
	if alternativeYearsWith != nil {
		return
	}

	newMap := map[string]*YearsWithID{}

	for _, thisItem := range YearsWith.Items {
		newMap[string(thisItem.ID)] = &thisItem.ID
		if thisItem.AlternativeKeys == "" {
			continue
		}

		tokens := strings.Split(thisItem.AlternativeKeys, ",")
		for _, thisToken := range tokens {
			thisToken = strings.ToLower(strings.TrimSpace(thisToken))
			if thisToken != "" {
				newMap[thisToken] = thisItem.ID.Clone()
			}
		}
	}

	alternativeYearsWith = &newMap
}

// ByAlternativeKey returns YearsWithID based on alternate key
func (e *EnumYearsWith) ByAlternativeKey(key *string) *YearsWithID {
	if e == nil || len(e.Items) < 1 || key == nil || *key == "" {
		return nil
	}

	initializeAlternativeYearsWith()

	itemKey := strings.ToLower(strings.TrimSpace(*key))

	if rtn, ok := (*alternativeYearsWith)[itemKey]; ok {
		return rtn
	}

	return nil
}
