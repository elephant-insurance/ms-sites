package uf

import "strconv"

type strConvUtil struct{}

var ConvertString = &strConvUtil{}

// ToInt parses the string for an int value.
// If the string cannot be parsed, will return defaultVal, if it is specified and an int (default 0)
func (p *strConvUtil) ToInt(s string, defaultVal ...interface{}) int {
	def := 0

	if len(defaultVal) > 0 && defaultVal[0] != nil {
		if iv, ok := defaultVal[0].(int); ok {
			def = iv
		}
	}

	if numericValue, err := strconv.Atoi(s); err == nil {
		return numericValue
	}

	return def
}
