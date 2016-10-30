// This code is generated by units-codegen; do not manualy edit this file.

package planar

import (
	"github.com/smyrman/units/linear"
)

// Units for Area values. Always multiply with a unit when setting the initial value like you would for
// time.Time. This prevents you from having to worry about the internal storage format.
const (
	SquearNanometer  Area = 1e-36
	SquearMicrometer Area = 1e-9
	SquearMillimeter Area = 1
	SquearCentimeter Area = 1e2
	SquearDecimeter  Area = 1e4
	SquearMeter      Area = 1e9
	SquearKilometer  Area = 1e36
	SquearInch       Area = Area(linear.Inch * linear.Inch)
	SquearFoot       Area = Area(linear.Foot * linear.Foot)
	SquearChain      Area = Area(linear.Chain * linear.Chain)
	Acre             Area = Area(linear.Chain * linear.Furlong)
	SquearMile       Area = Area(linear.Mile * linear.Mile)
)

// SquearNanometers returns a as a floating point number of squearnanometers.
func (a Area) SquearNanometers() float64 {
	return float64(a / SquearNanometer)
}

// SquearMicrometers returns a as a floating point number of squearmicrometers.
func (a Area) SquearMicrometers() float64 {
	return float64(a / SquearMicrometer)
}

// SquearMillimeters returns a as a floating point number of squearmillimeters.
func (a Area) SquearMillimeters() float64 {
	return float64(a / SquearMillimeter)
}

// SquearCentimeters returns a as a floating point number of squearcentimeters.
func (a Area) SquearCentimeters() float64 {
	return float64(a / SquearCentimeter)
}

// SquearDecimeters returns a as a floating point number of squeardecimeters.
func (a Area) SquearDecimeters() float64 {
	return float64(a / SquearDecimeter)
}

// SquearMeters returns a as a floating point number of squearmeters.
func (a Area) SquearMeters() float64 {
	return float64(a / SquearMeter)
}

// SquearKilometers returns a as a floating point number of squearkilometers.
func (a Area) SquearKilometers() float64 {
	return float64(a / SquearKilometer)
}

// SquearInches returns a as a floating point number of squearinches.
func (a Area) SquearInches() float64 {
	return float64(a / SquearInch)
}

// SquearFeet returns a as a floating point number of squearfeet.
func (a Area) SquearFeet() float64 {
	return float64(a / SquearFoot)
}

// Chains returns a as a floating point number of chains.
func (a Area) Chains() float64 {
	return float64(a / SquearChain)
}

// Acre returns a as a floating point number of acre.
func (a Area) Acre() float64 {
	return float64(a / Acre)
}

// SquearMiles returns a as a floating point number of squearmiles.
func (a Area) SquearMiles() float64 {
	return float64(a / SquearMile)
}

// Abs returns the absolute value of a as a copy.
func (a Area) Abs() Area {
	if a < 0 {
		return -a
	}
	return a
}

// Mul returns the product of a * x as a new Area.
func (a Area) Mul(x float64) Area {
	return a * Area(x)
}

// Div returns the quotient of a / x as a new Area.
func (a Area) Div(x float64) Area {
	return a / Area(x)
}

// DivArea returns the quotient of a / x as a floating point number.
func (a Area) DivArea(x Area) float64 {
	return float64(a / x)
}
