[![GitHub Actions](https://github.com/smyrman/units/workflows/Go/badge.svg?branch=master)](https://github.com/smyrman/units/actions?query=workflow%3AGo+branch%master)
[![Go Reference](https://pkg.go.dev/badge/github.com/smyrman/units.svg)](https://pkg.go.dev/github.com/smyrman/units)

# Units

Units is a small Go package that utilizes Go's strict type system to achieve better type safety when storing and doing simple mathematical operations on physical units. The concept is very similar to that of the `time.Duration` type in the Go standard library.

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

Each type has defined constants for relevant units. In general, you don't need to worry about which unit the underlying values is stored in, as long as you **always multiply with a relevant uint** when you first set your value. E.g.:

```go
var a angular.Angle = 90 * angular.Degrees
```

## Use at your own risk ⚠️

This library has seen low usage and has little testing. If you still prefer to use this package over other packages, contributions are still welcome.

## Alternatives

Alternative units pacakges:

- https://pkg.go.dev/gonum.org/v1/gonum/unit
- https://pkg.go.dev/github.com/bcicen/go-units

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

## Why would I want to use a units package?

The [Mars Climate Orbiter](https://en.wikipedia.org/wiki/Mars_Climate_Orbiter) crashed into Mars when software producing the non-SI units of pound-seconds was used to generate input to a program expecting newton-seconds (N s). While the units package would not have prevented that particular issue of communication between programs, the goal is that it will help you to avoid unit related mistakes within your own program, and to make working with various units as straight-forward and painless as possible.

Considering the low amount of usage of _this_ units package in particular, you may want to consider one that's more used though! Or make sure to add the number of tests that you need to be confident in your solution.

## Our experience

We found an earlier version of this code surprisingly useful back in 2012, when we where developing an autonomous robot using Go for the high-level programming.

While we still used plain float64 values in the matrices that calculated the next system states, we found using the units package for the input and output values of various blocks extremely useful. This package helped us verify that we never did something stupid like passing a distance to a function that expected a velocity. The simplicity of using the return value from `.Normalize()`on an angle, also made it very easy for us to ensure normalized angles for our navigation system. Forgetting to normalize an angle could lead to spinning robots...

We used metric values for both internal and displaying purposes, but the aim is also that this package will help those who wish to provide non-metric units as input or display values for their programs.

## Simple example

This example shows that you can divide `linear.Distance` with a `time.Duration` (standard library) to get a velocity, as well as use arbitrary units for input and display purposes.

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
	fmt.Printf("v = %.2f m/s\n", v.MetersPerSecond())
	fmt.Printf("v = %.2f kmph\n", v.KilometersPerHour())
	fmt.Printf("v = %.2f mph\n", v.MilesPerHour())
}
```
