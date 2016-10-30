[![Build Status](https://travis-ci.org/smyrman/units.svg?branch=master)](https://travis-ci.org/smyrman/units)
[![GoDoc](https://godoc.org/github.com/smyrman/units?status.svg)](http://godoc.org/github.com/smyrman/units)

# Units

## Use at your own risk!
Note that this library is NOT production ready. If you want to use it anyway, contributions and bug reports are welcome!

## Breaking changes!
This branch of the units library is introduces breaking changes to the API. To use the old API, please checkout
tag `v0.0.1`. If you download the units library via `go get`, consider using tool such as [Glide](http://glide.sh/)
instead to help you install/vendor specific versions of your Go dependencies.

## TODO

- [x] Code generation for sub-packages
- [x] Sub-package for angle, angular velocity and acceleration
- [x] Sub-package for linear distance, velocity and acceleration
- [x] Sub-package for planar coordinates based on distance and angle
- [ ] Support for volume, volumetric flow rate (velocity) and acceleration
- [ ] Support for mass, mass flow rate (velocity) and acceleration
- [ ] Support for force, work and torque
- [x] Metric units
- [x] Modern Imperial units
- [ ] Appropriate test coverage per package.

## What has changed?
The package layout is being completely changed. Type and constant definitions have been moved into sub-packages to allow
for shorter names, and a more restricted scope from which package symmetry can arise.

Here are the type renames:
- `units.Distance` -> `linear.Distance`
- `units.Velocity` -> `linear.Velocity`
- `units.Angle` -> `angular.Angle`
- `units.AngularVelocity` -> `angular.Velocity`
- `units.Coordinate` -> `planar.Point`

There are also some new types for these packages:
- `linear.Acceleration`
- `angular.Acceleration`
- `planar.Area`

The new sub-packages forms a symmetry of related types and rely on code generation for easier maintenance.

## About
Units is a small Go package that utilizes Go's strict type system to achieve better type safety when storing and doing
simple mathematical operations on physical units. The concept is very similar to that of the `time.Duration` type in the
Go standard library.

It currently has the following sub-packages:

- [angular](http://godoc.org/github.com/smyrman/units/angular) defines the types:
	- Angle
	- Velocity
	- Acceleration
- [linear](https://godoc.org/github.com/smyrman/units/linear) define the types:
	- Distance
	- Velocity
	- Acceleration
- [planar](https://godoc.org/github.com/smyrman/units/planar) defines the types:
	- Point
	- Area

## Why would I want to use the units package?
The [Mars Climate Orbiter](https://en.wikipedia.org/wiki/Mars_Climate_Orbiter) crashed into Mars due to software
producing the non-SI units of pound-seconds was used as input to a program expecting newton-seconds (N s). While the
units package would not have prevented that particular issue of communication between programs, the goal is that it
will at least help you to avoid unit related mistakes within your own program, and to make working with various units
as straight-forward and pain-free as possible.

With the units package, you don't need to worry about which unit the underlying value is stored as. Like with
`time.Duration` from the standard library, you multiply with a unit when you first set the value:

```go
var a angular.Angle = 90 * angular.Degrees
```

## Our experience
We found an the original version of this code surprisingly useful back in 2012, when we where developing a autonomous
robot using Go for the high-level programming.

While we still used float64 values in the matrices that calculated the next system states, we found using the units
package for the input and output values of various blocks extremely useful. This package helped us verify that we never
did something stupid when passing the values around in the robot code, like passing a distance to a function that wanted
the velocity. It was also very easy to make sure we always Normalized our angles - which is easy to forget, but pretty
important for navigational features!

## Example
This example shows that you can divide `linear.Distance` with a `time.Duration` (standard library) to get a velocity.

```go
package main

import(
	"fmt"
	"time"
	"github.com/smyrman/units/linear"
)

func main() {
	d := 66 * linear.Feet
	t := 2 * time.Second
	v := d.DivDuration(t)
	fmt.Printf("v = %.2f m/s^2\n", v.MetersPerSecond())
	fmt.Printf("v = %.2f kmph\n", v.KilometersPerHour())
	fmt.Printf("v = %.2f mph\n", v.MilesPerHour())
}
```

## Origin
The units package was originally developed for use in the robot Loke (eng:Loki), that participated in the Eurobot
competition in France in 2012 (http://eurobot-ntnu.no).
