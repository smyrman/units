package linear_test

import (
	"testing"
	"time"

	"github.com/smyrman/units/internal/floattest"
	"github.com/smyrman/units/linear"
)

func TestDistanceDivDuration(t *testing.T) {
	cases := []floattest.CompareCase{
		{
			Name:   "1m:1sec==1m/sec",
			Expect: 1,
			Func:   linear.Meter.DivDuration(1 * time.Second).MetersPerSecond,
		},
		{
			Name:   "100m:4sec==25m/sec",
			Expect: 25,
			Func:   (100 * linear.Meter).DivDuration(4 * time.Second).MetersPerSecond,
		},
		{
			Name:   "100km:5h==20km/h",
			Expect: 20,
			Func:   (100 * linear.Kilometer).DivDuration(5 * time.Hour).KilometersPerHour,
		},
		{
			Name:   "100miles:5h==20mph",
			Expect: 20,
			Func:   (100 * linear.Mile).DivDuration(5 * time.Hour).MilesPerHour,
		},
		{
			Name:   "100nmi:5h==20knot",
			Expect: 20,
			Func:   (100 * linear.NauticalMile).DivDuration(5 * time.Hour).Knots,
		},
	}
	for i := range cases {
		cases[i].Run(t)
	}
}

func TestMetricDistanceUnits(t *testing.T) {
	d := 2 * linear.Meter
	cases := []floattest.CompareCase{
		{Name: "Nanometers()", Expect: 2e9, Func: d.Nanometers},
		{Name: "Micrometers()", Expect: 2e6, Func: d.Micrometers},
		{Name: "Millimeters()", Expect: 2e3, Func: d.Millimeters},
		{Name: "Centimeters()", Expect: 2e2, Func: d.Centimeters},
		{Name: "Decimeters()", Expect: 2e1, Func: d.Decimeters},
		{Name: "Meters()", Expect: 2, Func: d.Meters},
		{Name: "Kilometers()", Expect: 2e-3, Func: d.Kilometers},
	}
	for i := range cases {
		cases[i].Run(t)
	}
}

func TestImperialDistanceUnits(t *testing.T) {
	// Test cases are based on list of units from https://en.wikipedia.org/wiki/Imperial_units.
	cases := []struct {
		Name                      string
		Unit, Expect              linear.Distance
		Feet, Millimeters, Meters float64
	}{
		{
			Name:        "Thou",
			Unit:        linear.Thou,
			Expect:      25.4 * linear.Micrometer,
			Feet:        1 / 12000,
			Millimeters: 0.0254,
			Meters:      0.0000254,
		},
		{
			Name:        "Inch",
			Unit:        linear.Inch,
			Expect:      1000 * linear.Thou,
			Feet:        1 / 12,
			Millimeters: 25.4,
			Meters:      0.0254,
		},
		{
			Name:        "Foot",
			Unit:        linear.Foot,
			Expect:      12 * linear.Inch,
			Feet:        1,
			Millimeters: 304.8,
			Meters:      0.3048,
		},
		{
			Name:        "Yard",
			Unit:        linear.Yard,
			Expect:      3 * linear.Foot,
			Feet:        3,
			Millimeters: 914.4,
			Meters:      0.9144,
		},
		{
			Name:        "Chain",
			Unit:        linear.Chain,
			Expect:      22 * linear.Yard,
			Feet:        66,
			Millimeters: 20116.8,
			Meters:      20.1168,
		},
		{
			Name:   "Furlong",
			Unit:   linear.Furlong,
			Expect: 10 * linear.Chain,
			Feet:   660,
			Meters: 201.168,
		},
		{
			Name:   "Mile",
			Unit:   linear.Mile,
			Expect: 8 * linear.Furlong,
			Feet:   5280,
			Meters: 1609.344,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			var f1, f2 float64

			f1, f2 = float64(tc.Unit), float64(tc.Expect)
			if !floattest.AlmostEqual(f1, f2) {
				t.Errorf("Unit [%.11f] != Expect [%.11f]", f1, f2)
			}

			if tc.Feet != 0 {
				f1, f2 = tc.Unit.Feet(), tc.Feet
				if !floattest.AlmostEqual(f1, f2) {
					t.Errorf("Unit.Feet() [%.11f] != Expect [%.11f]", f1, f2)
				}
			}

			if tc.Meters != 0 {
				f1, f2 = tc.Unit.Meters(), tc.Meters
				if !floattest.AlmostEqual(f1, f2) {
					t.Errorf("Unit.Meters() [%.11f] != Expect [%.11f]", f1, f2)
				}
			}

			if tc.Millimeters != 0 {
				f1, f2 = tc.Unit.Millimeters(), tc.Millimeters
				if !floattest.AlmostEqual(f1, f2) {
					t.Errorf("Unit.Millimeters() [%.11f] != Expect [%.11f]", f1, f2)
				}
			}
		})
	}
}
