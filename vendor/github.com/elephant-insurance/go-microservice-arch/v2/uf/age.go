package uf

import (
	"fmt"
	"time"
)

type ageUtil struct{}

var Age = &ageUtil{}

func (p *ageUtil) Today(dateOfBirth time.Time) int {
	return p.At(dateOfBirth, time.Now())
}

// At gets the age of an entity at a certain time.
func (p *ageUtil) At(dateOfBirth, now time.Time) int {
	age := 1
	for dateOfBirth.AddDate(age, 0, 0).Before(now) {
		age++
	}

	return age - 1
}

// FromString gets the age of someone given their birthdate in a standard date format and the date of evaluation
// These formats are recognized: "01/02/2006", "01-02-2006", "2006/01/02", "2006-01-02"
func (p *ageUtil) FromString(dob string, asOf time.Time) (int, error) {
	birthDate, err := DateTime.FromString(dob)
	if err != nil {
		return -1, fmt.Errorf("Could not determine driver age from birthdate %v: \r\n\t%v", dob, err.Error())
	}
	return p.At(birthDate, asOf), nil
}

// NowFromString gets the current age of someone given their birthdate in a standard date format
func (p *ageUtil) NowFromYMDString(dob string) (int, error) {
	return p.FromString(dob, time.Now())
}
