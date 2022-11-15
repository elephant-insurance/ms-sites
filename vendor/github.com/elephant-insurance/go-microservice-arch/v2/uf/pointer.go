package uf

import "time"

type pointerUtil struct{}

var Pointer = &pointerUtil{}

// These "To" methods make it easy to create a pointer inline by hiding the necessary creation of an intermediate variable.
// Feel free to add a "To" method for any Golang built-in type.
// If you need a "To" method for an app-specific type, please create it in your application code.

func (p *pointerUtil) ToString(s string) *string {
	return &s
}

func (p *pointerUtil) ToInt(i int) *int {
	return &i
}

func (p *pointerUtil) ToInt64(i int64) *int64 {
	return &i
}

func (p *pointerUtil) ToFloat32(f float32) *float32 {
	return &f
}

func (p *pointerUtil) ToFloat64(f float64) *float64 {
	return &f
}

func (p *pointerUtil) ToTime(t time.Time) *time.Time {
	return &t
}

// ToNow returns a pointer to a time.Time set to the current system time
func (p *pointerUtil) ToNow() *time.Time {
	jetzt := time.Now()
	return &jetzt
}

func (p *pointerUtil) ToBool(b bool) *bool {
	return &b
}

func (p *pointerUtil) ToTrue() *bool {
	foo := true
	return &foo
}

func (p *pointerUtil) ToFalse() *bool {
	foo := false
	return &foo
}

// CloneInt creates a copy of the underlying int value and returns a pointer to it
func (p *pointerUtil) CloneInt(src *int) *int {
	if src == nil {
		return nil
	}

	rtn := *src

	return &rtn
}

// CompareInt returns true if both pointers are nil, or if they point to equal values
func (p *pointerUtil) CompareInt(i1, i2 *int) bool {
	if i1 == nil && i2 == nil {
		// both nil are equal
		return true
	}

	if i1 == nil || i2 == nil {
		return false
	}

	return *i1 == *i2
}
