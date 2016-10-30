// This code is generated by units-codegen; do not manualy edit this file.

package angular

import (
	"time"
)

// Units for Velocity values. Always multiply with a unit when setting the initial value like you would for
// time.Time. This prevents you from having to worry about the internal storage format.
const (
	MilliradianPerSecond Velocity = Velocity(Milliradian) / Velocity(time.Second)
	RadianPerSecond      Velocity = Velocity(Radian) / Velocity(time.Second)
	DegreePerSecond      Velocity = Velocity(Degree) / Velocity(time.Second)
	GradianPerSecond     Velocity = Velocity(Gradian) / Velocity(time.Second)
)

// MilliradiansPerSecond returns v as a floating point number of milliradianspersecond.
func (v Velocity) MilliradiansPerSecond() float64 {
	return float64(v / MilliradianPerSecond)
}

// RadiansPerSecond returns v as a floating point number of radianspersecond.
func (v Velocity) RadiansPerSecond() float64 {
	return float64(v / RadianPerSecond)
}

// DegreeesPerSecond returns v as a floating point number of degreeespersecond.
func (v Velocity) DegreeesPerSecond() float64 {
	return float64(v / DegreePerSecond)
}

// GradiansPerSecond returns v as a floating point number of gradianspersecond.
func (v Velocity) GradiansPerSecond() float64 {
	return float64(v / GradianPerSecond)
}

// Abs returns the absolute value of v as a copy.
func (v Velocity) Abs() Velocity {
	if v < 0 {
		return -v
	}
	return v
}

// Mul returns the product of v * x as a new Velocity.
func (v Velocity) Mul(x float64) Velocity {
	return v * Velocity(x)
}

// Div returns the quotient of v / x as a new Velocity.
func (v Velocity) Div(x float64) Velocity {
	return v / Velocity(x)
}

// DivVelocity returns the quotient of v / x as a floating point number.
func (v Velocity) DivVelocity(x Velocity) float64 {
	return float64(v / x)
}

// DivDuration returns the quotient of v / t as a Acceleration.
func (v Velocity) DivDuration(t time.Duration) Acceleration {
	return Acceleration(float64(v) / float64(t))
}

// DivAcceleration returns the quotient of v / x as a time.Duration.
func (v Velocity) DivAcceleration(x Acceleration) time.Duration {
	return time.Duration(float64(v) / float64(x))
}

// MulDuration returns the product of v * t as a Angle.
func (v Velocity) MulDuration(t time.Duration) Angle {
	return Angle(float64(v) * float64(t))
}