// This code is generated by units-codegen; do not manualy edit this file.

package angular

import (
	"math"
	"time"
)

// Units for Angle values. Always multiply with a unit when setting the initial value like you would for
// time.Time. This prevents you from having to worry about the internal storage format.
const (
	Milliradian Angle = 1
	Radian      Angle = 1e3
	Degree      Angle = 2 * math.Pi * (Radian / 360)
	Gradian     Angle = 2 * math.Pi * (Radian / 400)
)

// Milliradians returns ng as a floating point number of milliradians.
func (ng Angle) Milliradians() float64 {
	return float64(ng / Milliradian)
}

// Radians returns ng as a floating point number of radians.
func (ng Angle) Radians() float64 {
	return float64(ng / Radian)
}

// Degreees returns ng as a floating point number of degreees.
func (ng Angle) Degreees() float64 {
	return float64(ng / Degree)
}

// Gradians returns ng as a floating point number of gradians.
func (ng Angle) Gradians() float64 {
	return float64(ng / Gradian)
}

// Abs returns the absolute value of ng as a copy.
func (ng Angle) Abs() Angle {
	if ng < 0 {
		return -ng
	}
	return ng
}

// Mul returns the product of ng * x as a new Angle.
func (ng Angle) Mul(x float64) Angle {
	return ng * Angle(x)
}

// Div returns the quotient of ng / x as a new Angle.
func (ng Angle) Div(x float64) Angle {
	return ng / Angle(x)
}

// DivAngle returns the quotient of ng / x as a floating point number.
func (ng Angle) DivAngle(x Angle) float64 {
	return float64(ng / x)
}

// DivDuration returns the quotient of ng / t as a Velocity.
func (ng Angle) DivDuration(t time.Duration) Velocity {
	return Velocity(float64(ng) / float64(t))
}

// DivVelocity returns the quotient of ng / x as a time.Duration.
func (ng Angle) DivVelocity(x Velocity) time.Duration {
	return time.Duration(float64(ng) / float64(x))
}
