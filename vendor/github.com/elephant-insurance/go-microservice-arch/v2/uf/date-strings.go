package uf

import "time"

const (
	dateStringToday         string = `today`
	dateStringTomorrow      string = `tomorrow`
	dateStringTodayWeek     string = `todayweek`
	dateStringOneWeek       string = `oneweek`
	dateStringWeek          string = `week`
	dateStringTomorrowWeek  string = `tomorrowweek`
	dateStringYesterday     string = `yesterday`
	dateStringSunday        string = `sunday`
	dateStringMonday        string = `monday`
	dateStringTuesday       string = `tuesday`
	dateStringWednesday     string = `wednesday`
	dateStringThursday      string = `thursday`
	dateStringFriday        string = `friday`
	dateStringSaturday      string = `saturday`
	dateStringSundayWeek    string = `sundayweek`
	dateStringMondayWeek    string = `mondayweek`
	dateStringTuesdayWeek   string = `tuesdayweek`
	dateStringWednesdayWeek string = `wednesdayweek`
	dateStringThursdayWeek  string = `thursdayweek`
	dateStringFridayWeek    string = `fridayweek`
	dateStringSaturdayWeek  string = `saturdayweek`
	dateStringTwoWeeks      string = `twoweeks`
	dateStringThreeWeeks    string = `threeweeks`
	dateStringFourWeeks     string = `fourweeks`
	dateStringMonth         string = `month`
	dateStringOneMonth      string = `onemonth`
	dateStringTwoMonths     string = `twomonths`
	dateStringThreeMonths   string = `threemonths`
	dateStringSixMonths     string = `sixmonths`
)

var dateStringMap = map[string]func() *Date{
	dateStringToday:         addDateFunc(0, 0, 0),
	dateStringTomorrow:      addDateFunc(0, 0, 1),
	dateStringTodayWeek:     addDateFunc(0, 0, 7),
	dateStringTomorrowWeek:  addDateFunc(0, 0, 8),
	dateStringYesterday:     addDateFunc(0, 0, -1),
	dateStringSunday:        weekdayFunc(time.Sunday),
	dateStringMonday:        weekdayFunc(time.Monday),
	dateStringTuesday:       weekdayFunc(time.Tuesday),
	dateStringWednesday:     weekdayFunc(time.Wednesday),
	dateStringThursday:      weekdayFunc(time.Thursday),
	dateStringFriday:        weekdayFunc(time.Friday),
	dateStringSaturday:      weekdayFunc(time.Saturday),
	dateStringSundayWeek:    weekdayWeekFunc(time.Sunday),
	dateStringMondayWeek:    weekdayWeekFunc(time.Monday),
	dateStringTuesdayWeek:   weekdayWeekFunc(time.Tuesday),
	dateStringWednesdayWeek: weekdayWeekFunc(time.Wednesday),
	dateStringThursdayWeek:  weekdayWeekFunc(time.Thursday),
	dateStringFridayWeek:    weekdayWeekFunc(time.Friday),
	dateStringSaturdayWeek:  weekdayWeekFunc(time.Saturday),
	dateStringOneWeek:       addDateFunc(0, 0, 7),
	dateStringWeek:          addDateFunc(0, 0, 7),
	dateStringTwoWeeks:      addDateFunc(0, 0, 14),
	dateStringThreeWeeks:    addDateFunc(0, 0, 21),
	dateStringFourWeeks:     addDateFunc(0, 0, 28),
	dateStringMonth:         addDateFunc(0, 1, 0),
	dateStringOneMonth:      addDateFunc(0, 1, 0),
	dateStringTwoMonths:     addDateFunc(0, 2, 0),
	dateStringThreeMonths:   addDateFunc(0, 3, 0),
	dateStringSixMonths:     addDateFunc(0, 6, 0),
}

func nextWeekday(wd time.Weekday) *time.Time {
	tryDate := time.Now()
	for i := 0; i < 7; i++ {
		tryDate = tryDate.AddDate(0, 0, 1)
		if tryDate.Weekday() == wd {
			return &tryDate
		}
	}
	return nil
}

func nextWeekdayDate(wd time.Weekday) *Date {
	return DateFactory.FromTime(nextWeekday(wd))
}

func nextWeekdayWeekDate(wd time.Weekday) *Date {
	nwdw := nextWeekday(wd).AddDate(0, 0, 7)
	return DateFactory.FromTime(&nwdw)
}

func weekdayFunc(wd time.Weekday) func() *Date {
	return func() *Date {
		return nextWeekdayDate(wd)
	}
}

func weekdayWeekFunc(wd time.Weekday) func() *Date {
	return func() *Date {
		return nextWeekdayWeekDate(wd)
	}
}

func addDateFunc(years, months, days int) func() *Date {
	return func() *Date {
		target := time.Now().AddDate(years, months, days)
		return DateFactory.FromTime(&target)
	}
}
