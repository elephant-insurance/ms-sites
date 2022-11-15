package uf

import "github.com/shopspring/decimal"

type decimalUtil struct{}

var Decimal = &decimalUtil{}

// RePrecision turns an int64 into a decimal with 4 digits after the decimal
// We use this to write large sets of decimals, all with the same precision, to a file
// This is for PARSING RAW SCORE FILES
func (du *decimalUtil) RePrecision(numerator int64, precision int) decimal.Decimal {
	rtn := decimal.New(numerator, -1*int32(precision))
	return rtn
}

// DePrecision removes the decimal from a decimal and returns it as a big integer
// This only works if we use the same precision throughout!
// This is for BUILDING A RAW SCORE FILE
func (du *decimalUtil) DePrecision(d *decimal.Decimal) int64 {
	if d == nil || d.IsZero() {
		return 0
	}

	rtnDec := *d

	factor := -1 * d.Exponent() // d.Exponent should usually be -4, but Decimal might trim int64 part for trailing zeroes?
	for i := int32(0); i < factor; i++ {
		rtnDec = rtnDec.Mul(ten)
	}

	return rtnDec.IntPart()
}

var ten = decimal.NewFromInt(10)
