[![Build Status](https://travis-ci.org/smyrman/units.svg?branch=master)](https://travis-ci.org/smyrman/units)
[![GoDoc](https://godoc.org/github.com/smyrman/units?status.svg)](http://godoc.org/github.com/smyrman/units)

# Units

Units is a small Go package that utilizes Go's strict type system to achieve
better type safety when storing and doing math on physical units.

The concept is very similar to that of the time.Duration type in the Go
standard library.

The following physical types are supported:

* Two-dimensional coordinates
* Distance
* Velocity
* Angle
* Angular velocity

These are basically the types we cared about when we implemented this library
back in 2012. Pull-requests for implementing additional Physical types, are
welcome.

## Why would I want to use the units package?

The [Mars Climate Orbiter](https://en.wikipedia.org/wiki/Mars_Climate_Orbiter)
crashed into Mars due to software producing the non-SI units of pound-seconds
was used as input to a program expecting newton-seconds (N s). While the units
package maybe would not have prevented that particular issue, it would help
make sure that within a program, you don't mess things up.

With the units package, you don't need to worry about which unit the underlying
value is stored as. Like with `duration.Time` from the standard library, you
multiply with a unit when you first set the value:

```go
var a units.Angle = 90 * units.Degrees
```

We found this package surprisingly useful back in 2012, when we where
developing a autonomous robot using Go for the high-level programming. This
package helped us verify that we never did something stupid in the code, like
passing a distance variable to a function that wanted the velocity.


## Examples
This example Shows that you can divide `units.Distance` with a `time.Duration`
(standard library) to get a velocity.

```go
import(
	"fmt"
	"time"
	"github.com/smyrman/units"
)

// For 'd', the type is units.Distance, the underlying type is flot64, and
// the underlying value is 6000.0:
d := 6*uints.Meter

// For 't', the type is time.Duration, the underlying type is int64, and
// the underlying value is 2e9:
t := 2*time.Second

// For 'v', the type will be units.Velocity, the underlying type is float64, and
// the underlying value will be 3000.0:
v := dist.DivideWithDuration(t)

// v.MetersPerSecond() returns the float64 value 3.0.
fmt.Printf("v = %.2f\n", v.MetersPerSecond())
```

## Origin

The units package was originally developed for use in the robot Loke
(eng:Loki), that participated in the Eurobot competition in France in 2012
(http://eurobot-ntnu.no).

