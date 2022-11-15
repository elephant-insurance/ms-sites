package uf

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const minInt = -int(^uint(0)>>1) - 1

// minDateLength specifies the minimum length required for date string value = yyyy-mm-dd = 10
const minDateLength = 10

// Date stores a date without a time
type Date struct {
	Day   *int
	Month *int
	Year  *int
}

// Datable abstracts the needed func from the Time struct so that we can use our type interchangably
type Datable interface {
	Date() (int, time.Month, int)
}

type dateFactory struct{}

var DateFactory = &dateFactory{}

// FromTime converts from Time to Date object
func (df *dateFactory) FromTime(t *time.Time) *Date {
	if t == nil {
		return nil
	}

	y, md, d := t.Date()
	m := int(md)

	return &Date{
		Day:   &d,
		Month: &m,
		Year:  &y,
	}
}

// New returns new Date object based on date, month, and year
func (df *dateFactory) New(d, m, y int) (*Date, error) {
	// just some basic sanity-checking
	if d < 1 || d > 31 {
		return nil, errors.New("day must be between 1 and 31")
	}

	if m < 1 || m > 12 {
		return nil, errors.New("month must be between 1 and 12")
	}

	return &Date{
		Day:   &d,
		Month: &m,
		Year:  &y,
	}, nil
}

// Today returns current date of Date object
func (df *dateFactory) Today() *Date {
	tn := time.Now()

	y, md, d := tn.Date()
	m := int(md)

	return &Date{
		Day:   &d,
		Month: &m,
		Year:  &y,
	}
}

// Datable interface member func
// Date mimics the time.Time func of the same name
func (d *Date) Date() (year int, month time.Month, day int) {
	if d != nil {
		if d.Day != nil && *d.Day > 0 && *d.Day < 32 {
			day = *d.Day
		}

		if d.Month != nil && *d.Month > 0 && *d.Month < 13 {
			month = time.Month(*d.Month)
		}

		if d.Year != nil {
			year = *d.Year
		}
	}

	return
}

// Clone deep copy and created duplicate Date object
func (d *Date) Clone() *Date {
	if d == nil {
		return nil
	}

	return &Date{
		Day:   Pointer.CloneInt(d.Day),
		Month: Pointer.CloneInt(d.Month),
		Year:  Pointer.CloneInt(d.Year),
	}
}

// Equals compares two Date objects against their equality
func (d *Date) Equals(d2 *Date) bool {
	if d == nil && d2 == nil {
		return true
	}

	if d == nil || d2 == nil {
		return false
	}

	return Pointer.CompareInt(d.Day, d2.Day) && Pointer.CompareInt(d.Month, d2.Month) && Pointer.CompareInt(d.Year, d2.Year)
}

// After check target(d2) is after the main Date(d)
func (d *Date) After(d2 *Date) bool {
	// nils always return false for this one
	if !d.Valid() || !d2.Valid() {
		return false
	}

	return *d.Year > *d2.Year ||
		(*d.Year == *d2.Year && *d.Month > *d2.Month) ||
		(*d.Year == *d2.Year && *d.Month == *d2.Month && *d.Day > *d2.Day)
}

// Before check target(d2) is before the main Date(d)
func (d *Date) Before(d2 *Date) bool {
	// nils always return false for this one
	if !d.Valid() || !d2.Valid() {
		return false
	}

	return *d.Year < *d2.Year ||
		(*d.Year == *d2.Year && *d.Month < *d2.Month) ||
		(*d.Year == *d2.Year && *d.Month == *d2.Month && *d.Day < *d2.Day)
}

/*
func (d *Date) DaysSince(asOf *Date) (int, error) {
	if asOf == nil {
		asOf = Today()
	}

	if !d.Valid || !asOf.Valid {
		return minInt, errors.New("invalid date submitted to DaysSince")
	}

	if d.Equals(asOf) { return 0 }

	// Goofy to have to do it this way, but it's really the recommended method
	st := d.ToUTCTime()
	et := asOf.ToUTCTime()
	days := 0

	for dateNow := *st; dateNow.Before(*et); dateNow = dateNow.AddDate(0,0,1) {
		days++
	}

}
*/

// ToUTCTime converts the Date object to UTC time format of Time object
func (d *Date) ToUTCTime() *time.Time {
	if !d.Valid() {
		return nil
	}

	m := time.Month(*d.Month)

	rtn := time.Date(*d.Year, m, *d.Day, 0, 0, 0, 0, time.UTC)

	return &rtn
}

// Valid checks and validate, values of Date Object
func (d *Date) Valid() bool {
	return d != nil && d.Day != nil && *d.Day > 0 && *d.Day < 32 && d.Month != nil && *d.Month > 0 && *d.Month < 13 && d.Year != nil
}

// maxAge is the max age we care about
// it should never matter, but just in case, we need a place to stop
const maxAge int = 120

// AgeInYears calculates Age based on current date and main Date (d) object
func (d *Date) AgeInYears() int {
	if d == nil || d.After(DateFactory.Today()) {
		return 0
	}
	rightNow := time.Now().UTC()
	today := rightNow.Truncate(time.Hour * 24)
	dob := d.ToUTCTime().Add(time.Hour)

	for i := 0; i < maxAge; i++ {
		if dob.AddDate(i, 0, 0).After(today) {
			return i
		}
	}

	return 0
}

// FromString Convert from string to Date object
// The date string must be at the start of the submitted string, and may be quoted
// Formats recognized: YYYYMMDD, YYYY-MM-DD, YYYY/MM/DD, MM-DD-YYYY, MM/DD/YYYY
func (df *dateFactory) FromString(dateYMD string) (*Date, error) {
	// Marshaling data is getting " (double quotes) are around value of date
	ymdString := strings.ReplaceAll(dateYMD, "\"", "")
	var selectedReggie *regexp.Regexp
	var match []string

	for i := 0; i < len(dateFormatReggies) && selectedReggie == nil; i++ {
		thisReg := dateFormatReggies[i]
		match = thisReg.FindStringSubmatch(ymdString)
		if len(match) == 4 {
			selectedReggie = &thisReg
		}
	}

	if selectedReggie == nil {
		// try looking up a date func by name
		dsl := strings.ToLower(ymdString)
		if dfunc, ok := dateStringMap[dsl]; ok {
			return dfunc(), nil
		}
		return nil, errors.New(`submitted string did not start with a recognized date string`)
	}

	nameMap := make(map[string]string, 3)
	names := selectedReggie.SubexpNames()
	for i := 1; i < len(names); i++ {
		thisName := names[i]
		if names[i] != "" {
			nameMap[thisName] = match[i]
		}
	}

	var ok bool
	var year, month, day int
	var ystr, mstr, dstr string

	ystr, ok = nameMap[regTagYear]
	if ok {
		mstr, ok = nameMap[regTagMonth]
	}
	if ok {
		dstr, ok = nameMap[regTagDay]
	}
	if !ok {
		return nil, errors.New(`failed to find a date element in submitted string`)
	}

	var parseErr error
	year, parseErr = strconv.Atoi(ystr)
	if parseErr == nil {
		month, parseErr = strconv.Atoi(mstr)
	}
	if parseErr == nil {
		day, parseErr = strconv.Atoi(dstr)
	}
	if parseErr != nil {
		return nil, errors.New(`failed to convert an element to integer`)
	}

	return df.New(day, month, year)
}

// ToYMDCompact  Converts the Date object in YYYYMMDD format
func (d *Date) ToYMDCompact() string {
	if d == nil || !d.Valid() {
		return ""
	}

	return fmt.Sprintf("%d%02d%02d", *d.Year, *d.Month, *d.Day)
}

// ToYMDFormat  Converts the Date object in YYYY-MM-DD format
func (d *Date) ToYMDFormat() string {
	if d == nil || !d.Valid() {
		return ""
	}

	return fmt.Sprintf("%d-%02d-%02d", *d.Year, *d.Month, *d.Day)
}

// ToMDYFormat converts the Date object in MM-DD-YYYY format
func (d *Date) ToMDYFormat() string {
	if d == nil || !d.Valid() {
		return ""
	}

	return fmt.Sprintf("%02d-%02d-%d", *d.Month, *d.Day, *d.Year)
}

// MarshalJSON custom JSON parsing for Date - return YYYY-MM-DD format
func (d *Date) MarshalJSON() ([]byte, error) {
	if d == nil || !d.Valid() {
		return nil, nil
	}

	return json.Marshal(d.ToYMDFormat())
}

// UnmarshalJSON unmarshal JSON to Date - from YYYY-MM-DD format only
func (d *Date) UnmarshalJSON(data []byte) error {
	dateObj, err := DateFactory.FromString(string(data))

	if dateObj.Valid() {
		*d = *dateObj
	}

	return err
}

// MarshalXML custom XML parsing for Date - return YYYY-MM-DD format
func (d *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d.Valid() {
		return e.EncodeElement(d.ToYMDFormat(), start)
	}

	return errors.New("invalid date object")
}

// UnmarshalXML unmarshal XML to Date - from YYYY-MM-DD format only
func (d *Date) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}
	dateObj, err := DateFactory.FromString(xmlData)

	if dateObj.Valid() {
		*d = *dateObj
	}

	return err
}

var dateFormatReggies = []regexp.Regexp{
	*regexp.MustCompile(`^(?P<` + regTagYear + `>\d\d\d\d)-(?P<` + regTagMonth + `>\d\d)-(?P<` + regTagDay + `>\d\d)`),   // YYYY-MM-DD
	*regexp.MustCompile(`^(?P<` + regTagYear + `>\d\d\d\d)\/(?P<` + regTagMonth + `>\d\d)\/(?P<` + regTagDay + `>\d\d)`), // YYYY/MM/DD
	*regexp.MustCompile(`^(?P<` + regTagYear + `>\d\d\d\d)(?P<` + regTagMonth + `>\d\d)(?P<` + regTagDay + `>\d\d)`),     // YYYYMMDD
	*regexp.MustCompile(`^(?P<` + regTagMonth + `>\d\d)-(?P<` + regTagDay + `>\d\d)-(?P<` + regTagYear + `>\d\d\d\d)`),   // MM-DD-YYYY
	*regexp.MustCompile(`^(?P<` + regTagMonth + `>\d\d)\/(?P<` + regTagDay + `>\d\d)\/(?P<` + regTagYear + `>\d\d\d\d)`), // MM/DD/YYYY
}

const (
	regTagYear  = `year`
	regTagMonth = `month`
	regTagDay   = `day`
)
