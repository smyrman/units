package planar

import (
	"github.com/smyrman/units/angular"
	"github.com/smyrman/units/linear"
)

// Point describes a planar point, stored as cartesian cooridinates.
type Point struct {
	X, Y linear.Distance
}

//PolarPoint returns a new Point based on the passed in polar coordinates.
func PolarPoint(r linear.Distance, phi angular.Angle) Point {
	return Point{
		X: r.Mul(phi.Cos()),
		Y: r.Mul(phi.Sin()),
	}
}

// Polar returns the polar representation of p.
func (p Point) Polar() (linear.Distance, angular.Angle) {
	r := linear.Hypot(p.X, p.Y)
	phi := angular.Atan2(float64(p.Y), float64(p.X)).Normalize()
	return r, phi
}

// Add returns the sum of p and p2 as a new point.
func (p Point) Add(p2 Point) Point {
	dx := p.X + p2.X
	dy := p.Y + p2.Y
	return Point{dx, dy}
}

// Sub returns the difference between p - p2 as a new point.
func (p Point) Sub(p2 Point) Point {
	dx := p.X - p2.X
	dy := p.Y - p2.Y
	return Point{dx, dy}
}

// Mul returns the product of p * k as a new Point.
func (p Point) Mul(k float64) Point {
	return Point{p.X * linear.Distance(k), p.Y * linear.Distance(k)}
}

// Div returns the quotient of p / k as a new Point.
func (p Point) Div(k float64) Point {
	return Point{p.X / linear.Distance(k), p.Y / linear.Distance(k)}
}

// Distance returns the linear distance between p and p2.
func (p Point) Distance(p2 Point) linear.Distance {
	dx := p2.X - p.X
	dy := p2.Y - p.Y
	return linear.Hypot(dx, dy)
}

// Angle return the Atan2 between dx/dy where dx and dy are given by p2 - p. In other words, you get the direction from
// p to p2 in a cartesial cooridinate system with 0 angles.
func (p Point) Angle(p2 Point) angular.Angle {
	dx := float64(p2.X - p.X)
	dy := float64(p2.Y - p.Y)
	return angular.Atan2(dy, dx).Normalize()
}
