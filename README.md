Units
=====

Units is a small go package that implements types, constants, converter
functions and some mathematics for the following physical concepts:

* Distance
* Velocity
* Angle
* Angular velocity

The concept is very similar to that of the time.Duration type in the Go
standard library.

Why would I need a units package?
---------------------------------

When operating with variables that hold physical units, it is sometimes hard to
see what kinds of data a variable holds, and what kind of units it represents.
Does it store a velocity, or a distance? Given it is a velocity, does it store
the value in meters per second, millimeteres per second, or in some other kind
of unit? Miles perhaps? Is the angle measured in degrees or radians?

By using the units package, all variables with a certain physical type (e.g.
distance, angle or velocity), is stored in the same format with the same unit.
Even though the numbers are always stored in the same way, you can specify them
anyway you want, and represent them anyway you want.  E.g. even though an angle
is stored as milliradians, you can specify it as degrees, and display it as
gradians.

The fact that Go only allows explicit type casts, has proven useful when using
the units package. E.g. it is hard to use a variable storing a velocity as a
distance by a mistake.

Note, that for best practice, types from the units package should only be
converted to float64 values by use of the converter functions (unless you have
a good reason not to).

Examples
--------

Import and const definitions::

	import(
		"fmt"
		"time"
		"github.com/smyrman/units"
	)

	// Define some shorcuts (optional)
	const(
		// The following has type time.Duration:
		sec   = time.Second
		msec  = time.Millisecond

		// The following has type units.Distance:
		m     = units.Meter
		cm    = units.Centimeter
		mm    = units.Millimeter

		// The following has type units.Velocity:
		kmph  = units.KilometerPerSecond
		mpsec = units.MeterPerSecond

		// The following has type units.Angle:
		deg = units.Degree
		rad = units.Radian
	)

Distance to velocity example::

    // For 'd', the type is units.Distance, the underlying type is flot64, and
    the underlying value is 6000.0:
    d := 6*m

    // For 't', the type is time.Duration, the underlying type is int64, and
    the underlying value is 2e9:
    t := 2*sec

    // For 'v', the type will be units.Velocity, the underlying type is float64, and
    the underlying value will be 3000.0:
    v := dist.DivideWithTime(t)

    // v.MetersPerSecond() returns the float64 value 3.0.
    fmt.Printf("v = %.2f\n", v.MetersPerSecond())

Origin
------

The units package was originally developed for use in the robot Loke
(eng:Loki), that participated in the Eurobot competition in France in 2012
(http://eurobot-ntnu.no).

