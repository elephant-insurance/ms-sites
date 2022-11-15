package enumerations

import (
	"regexp"
	"strings"
)

// IDByDescription gets an occupation id by its description
func (e *EnumOccupation) IDByDescription(title string) *OccupationID {
	titleKey := normalizeOccupationMapKey(title)
	if rtn, found := occupationDescriptionToIDMap[titleKey]; found {
		return rtn
	}
	return nil
}

// IDByKeyword gets an occupation id by one of its keywords
// NOTE: we don't do any checking for duplicate keys in this map, so you will get the last match we processed on init()
func (e *EnumOccupation) IDByKeyword(keyword string) *OccupationID {
	keywordKey := normalizeOccupationMapKey(keyword)
	if rtn, found := occupationKeywordToIDMap[keywordKey]; found {
		return rtn
	}

	return nil
}

func normalizeOccupationMapKey(t string) string {
	krx := regexp.MustCompile(`\w`)
	return strings.ToLower(strings.Join(krx.FindAllString(t, -1), ""))
}

var (
	occupationDescriptionToIDMap map[string]*OccupationID
	occupationKeywordToIDMap     map[string]*OccupationID
)

func init() {
	occupationDescriptionToIDMap = map[string]*OccupationID{}
	occupationKeywordToIDMap = map[string]*OccupationID{}
	for i := 0; i < len(Occupation.Items); i++ {
		thisOcc := Occupation.Items[i]
		if thisOcc == nil {
			continue
		}
		normIdx := normalizeOccupationMapKey(thisOcc.Desc)
		occupationDescriptionToIDMap[normIdx] = &thisOcc.ID
		kw := thisOcc.Keywords
		if kw == "" {
			continue
		}
		tokens := strings.Split(kw, ",")
		for _, thisToken := range tokens {
			kwkey := normalizeOccupationMapKey(thisToken)
			if kwkey != "" {
				occupationKeywordToIDMap[kwkey] = &thisOcc.ID
			}
		}
	}
}
