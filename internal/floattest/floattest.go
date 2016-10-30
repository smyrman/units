package floattest

import (
	"math"
	"testing"
)

// AlmostEqual returns true if the diff between x and y is less than 1e-11.
func AlmostEqual(x, y float64) bool {
	const tolerance = 1e-11

	if math.Abs(x-y) < tolerance {
		return true
	}
	return false
}

// CompareCase should be used for Go >= 1.7 sub-tests that compares a returned float64 value from a function (Func)
// against an expected value (Expect).
type CompareCase struct {
	Name   string
	Expect float64
	Func   func() float64
}

// Run logs an error on t if the result of c.Func differs more than 1e-11 from c.Expect. the test will be run in
// parallel.
func (c CompareCase) Run(t *testing.T) {
	t.Run(c.Name, func(t *testing.T) {
		t.Parallel()
		if actual := c.Func(); !AlmostEqual(c.Expect, actual) {
			t.Errorf("%f (expected) != %f (actual)", c.Expect, actual)
		}
	})
}
