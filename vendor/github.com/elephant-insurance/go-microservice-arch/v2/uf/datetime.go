package uf

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type dateTimeUtil struct{}

var DateTime = &dateTimeUtil{}

// FromString converts date strings to Times
// These formats are recognized: "01/02/2006", "01-02-2006", "2006/01/02", "2006-01-02"
func (dtu *dateTimeUtil) FromString(strDate string) (time.Time, error) {
	if strDate != "" {
		datePattern := dtu.getFormat(strDate)
		if datePattern == `` {
			return time.Now(), fmt.Errorf("date string %v not recognized as a valid date", strDate)
		}

		t, err := time.ParseInLocation(datePattern, strDate, time.UTC)
		if err != nil {
			panic(err)
		}
		return t, err
	}
	return time.Now(), errors.New("no date string to convert")
}

// PCTimestampToWeb converts a pc timestamp (2006-01-02 15:04:05 -0700) to a web date (MM-DD-YYYY)
func (dtu *dateTimeUtil) PCTimestampToWeb(pcDate string, asUTC bool) (string, error) {
	var rtn time.Time
	var err error

	if rtn, err = time.Parse(datePattern_PCTimestamp, pcDate); err != nil {
		return "", err
	}

	// return as UTC date?
	if asUTC {
		rtnUtc := rtn.UTC()
		return dtu.ToWebDate(&rtnUtc), nil
	}

	// return as non-UTC date
	return dtu.ToWebDate(&rtn), nil
}

// PCToWeb converts a pc date (YYYY/MM/DD) to a web date (MM-DD-YYYY)
func (dtu *dateTimeUtil) PCToWeb(pcDate string) (string, error) {
	var rtn time.Time
	var err error

	if rtn, err = time.Parse(datePattern_PC, pcDate); err != nil {
		return "", err
	}

	return dtu.ToWebDate(&rtn), nil
}

// ToPCDate converts a Time into a PC date (YYYY/MM/DD)
func (dtu *dateTimeUtil) ToPCDate(t *time.Time) string {
	return dtu.formatDateTime(t, datePattern_PC)
}

// ToWebDate converts a Time into a web date (MM-DD-YYYY)
func (dtu *dateTimeUtil) ToWebDate(t *time.Time) string {
	return dtu.formatDateTime(t, datePattern_Web)
}

// WebToPC converts a web date (MM-DD-YYYY)/ (YYYY-MM-DD)  to a pc date (YYYY/MM/DD)
func (dtu *dateTimeUtil) WebToPC(webDate string) (string, error) {
	var rtn time.Time
	var err error

	// length of date
	if len(webDate) > 10 {
		webDate = webDate[:10]
	}

	if rtn, err = time.Parse(datePattern_Web, webDate); err != nil {
		if rtn, err = time.Parse(datePatten_WebToPC, webDate); err != nil {
			return "", err
		}
	}

	return dtu.ToPCDate(&rtn), nil
}

// WebToPCPatten converts a web date (YYYY-MM-DD) to a pc date (YYYY/MM/DD)
func (dtu *dateTimeUtil) WebToPCPatten(webDate string) (string, error) {
	var rtn time.Time
	var err error

	// length of date
	if len(webDate) > 10 {
		webDate = webDate[:10]
	}

	if rtn, err = time.Parse(datePatten_WebToPC, webDate); err != nil {
		return "", err
	}

	return dtu.ToPCDate(&rtn), nil
}

// YMDStringToInts => parsed year (int), month (int), and day from Date (string)
func (dtu *dateTimeUtil) YMDStringToInts(date string) (year, month, day int) {
	year, month, day = 0, 0, 0

	if date != "" {
		if dt, fail := dtu.FromString(date); fail == nil {
			year = dt.Year()
			month = int(dt.Month())
			day = dt.Day()
		}
	}

	return
}

// InvertDateString => converts between MDY and YMD with either slashes or dashes
// example: 01/02/2006 => 2006/01/02 or 2006-01-02 => 01-02-2006
func (dtu *dateTimeUtil) InvertDateString(date string) string {
	invertedPattern := ""
	datePattern := dtu.getFormat(date)
	dateTime, _ := dtu.FromString(date)
	switch {
	case datePattern == datePattern_MDY4_Slashes:
		invertedPattern = datePattern_Y4MD_Slashes
	case datePattern == datePattern_Y4MD_Slashes:
		invertedPattern = datePattern_MDY4_Slashes
	case datePattern == datePattern_MDY4_Dashes:
		invertedPattern = datePattern_Y4MD_Dashes
	case datePattern == datePattern_Y4MD_Dashes:
		invertedPattern = datePattern_MDY4_Dashes
	}

	return dtu.formatDateTime(&dateTime, invertedPattern)
}

//getFormat => returns matching date pattern const
func (dtu *dateTimeUtil) getFormat(strDate string) string {
	idxs := strings.IndexRune(strDate, '/')
	idxd := strings.IndexRune(strDate, '-')
	switch {
	case idxs == 2 && idxd < 0:
		return datePattern_MDY4_Slashes
	case idxs == 4 && idxd < 0:
		return datePattern_Y4MD_Slashes
	case idxd == 2 && idxs < 0:
		return datePattern_MDY4_Dashes
	case idxd == 4 && idxs < 0:
		return datePattern_Y4MD_Dashes
	default:
		return ""
	}
}

// formatDateTime converts Time to provided format string
func (dtu *dateTimeUtil) formatDateTime(t *time.Time, dateFormat string) string {
	if t != nil {
		return t.Format(dateFormat)
	}
	return ""
}

const (
	datePattern_MDY4_Dashes  string = `01-02-2006`
	datePattern_MDY4_Slashes string = `01/02/2006`
	datePattern_Y4MD_Dashes  string = `2006-01-02`
	datePattern_Y4MD_Slashes string = `2006/01/02`

	datePattern_PC          = datePattern_Y4MD_Slashes
	datePattern_Web         = datePattern_MDY4_Dashes
	datePattern_PCTimestamp = `2006-01-02 15:04:05 -0700`
	datePatten_WebToPC      = datePattern_Y4MD_Dashes
)

/*
utils.WebDateToPCDate
utils.PCDateToWebDate
utils.PCTimestampToWebDate
utils.DateStringToYearMonth
utils.FormatQuoteRetrieveKey
*/
