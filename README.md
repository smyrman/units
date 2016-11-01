[![Build Status](https://travis-ci.org/smyrman/units.svg?branch=master)](https://travis-ci.org/smyrman/units)
[![GoDoc](https://godoc.org/github.com/smyrman/units?status.svg)](http://godoc.org/github.com/smyrman/units)

# Units

## Use at your own risk
Note that this library is NOT production ready. If you want to use it anyway, contributions and bug reports are welcome!

## Breaking changes!
This branch of the units library is introduces breaking changes to the API. To use the old API, please checkout tag
`v0.0.1`. If you download the units library via `go get`, consider using a tool such as [Glide](http://glide.sh/)
instead to help you install/vendor specific versions of your Go dependencies.

To instead use the new API, see this [short migration guide](#migrating-from-v001) for more information.

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

Each type has defined constants for relevant units. In general, you don't need to worry about which unit the underlying
values is stored as, as long as you **always multiply with a relevant uint** when you first set your value. E.g.:

```go
var a angular.Angle = 90 * angular.Degrees
```

## Why would I want to use the units package?
The [Mars Climate Orbiter](https://en.wikipedia.org/wiki/Mars_Climate_Orbiter) crashed into Mars when software producing
the non-SI units of pound-seconds was used to generate input to a program expecting newton-seconds (N s). While the
units package would not have prevented that particular issue of communication between programs, the goal is that it will
help you to avoid unit related mistakes within your own program, and to make working with various units as
straight-forward and painless as possible.


## Our experience
We found an the original version of this code surprisingly useful back in 2012, when we where developing a autonomous
robot using Go for the high-level programming.

While we still used float64 values in the matrices that calculated the next system states, we found using the units
package for the input and output values of various blocks extremely useful. This package helped us verify that we never
did something stupid like passing a distance to a function that wanted the velocity. The simplisity of using the return
value from`.Normalize()`on an angle, also made it very easy for us to ensure normalized angles for our navigation
system.

## Simple example
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

See the [godoc reference docs](http://godoc.org/github.com/smyrman/units) for a full list of methods per type.


## Origin
The units package was originally developed for use in the robot Loke (eng:Loki), that participated in the Eurobot
competition in France in 2012.


## Migrating from v0.0.1
Since v0.0.1, the package layout has been completely changed. Type and constant definitions have been moved
into sub-packages to allow for shorter names, and a more restricted scope from which package symmetry can arise.

To migrate, search and replace the following:
- `units.Distance` -> `linear.Distance`
- `units.Velocity` -> `linear.Velocity`
- `units.Angle` -> `angular.Angle`
- `units.AngularVelocity` -> `angular.Velocity`
- `units.Coordinate` -> `planar.Point`

Then, assure the propper imports, most conveniently through the use of
[goimports](https://godoc.org/golang.org/x/tools/cmd/goimports), or manually by replacing the
"github.com/smyrman/units" import with the appropriate import(s):

- "github.com/smyrman/units/linear"
- "github.com/smyrman/units/angular"
- "github.com/smyrman/units/planar"
