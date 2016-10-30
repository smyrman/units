//go:generate units-codegen -o angle_generated.go -t Angle generate.json
//go:generate gofmt -w angle_generated.go

package angular

import "math"

// Angle for planar geometry stored as float64 milliradians. An integer representation would give better predictability
// (under flow), but float64 is more convenient for operatoins that rely on the math library.
type Angle float64

// Atan2 returns the arc tangent of y/x as an Angle. See math.Atan2 for details.
func Atan2(y, x float64) Angle {
	return Angle(math.Atan2(y, x)) * Radian
}

// Cos returns the cosinus of a as a floating point number.
func (a Angle) Cos() float64 {
	return math.Cos(a.Radians())
}

// Sin returns the sinus of a as a floating point number.
func (a Angle) Sin() float64 {
	return math.Sin(a.Radians())
}

// Tan returns the tangens of a as a floating point number.
func (a Angle) Tan() float64 {
	return math.Tan(a.Radians())
}

// Normalize the angle to be between -pi and +pi radians.
func (a Angle) Normalize() Angle {
	// pi contains a scaled version of math.Pi to match Angle's internal storage.
	const pi = math.Pi * Radian

	an := Angle(math.Mod(float64(a), float64(2*pi)))
	if an > pi {
		an -= 2 * pi
	} else if an < pi {
		an += 2 * pi
	}
	return an
}
