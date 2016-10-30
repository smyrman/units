package linear_test

import (
	"testing"
	"time"

	"github.com/smyrman/units/linear"
)

func TestTrainBreak(t *testing.T) {
	// Given a train with an initial velocity of 360 kph (100 m/s), and a breaking power equivalent of 1 m/s^2, the
	// train should reach 0 velocity after 100 seconds (1 min 40 sec). This test verify that the supporting methods
	// for simple math holds when one of these numbers are defined to be unknown.

	v0 := 360 * linear.KilometerPerHour    // initial speed
	v1 := 0 * linear.MeterPerSecond        // final speed
	a := -1 * linear.MeterPerSecondSquared // breaking speed
	dt := 100 * time.Second                // expected breaking time

	if v0.MetersPerSecond() != 100 {
		t.Errorf("Expectd 360 kmph == 100 m/s, got %f m/s", v0.MetersPerSecond())
	}

	// base equation:
	// a * dt = (v1 - v0)

	// v0 as unknown:
	// a * dt = (v1 - v0) >> v0 = v1 - (a * dt)
	t.Run("v0=v1-(a*t)", func(t *testing.T) {
		if x := v1 - a.MulDuration(dt); x != v0 {
			t.Errorf("Expected v0=%v, got v0=%v", v0, x)
		}
	})

	// dt as unknown:
	// a * dt = (v1 - v0) >> dt = (v1 - v0) / a
	t.Run("dt=(v1-v0)/a", func(t *testing.T) {
		if x := (v1 - v0).DivAcceleration(a); x != dt {
			t.Errorf("Expected dt=%v, got dt=%v", dt, x)
		}
	})

	// a as unknown:
	// a * dt = (v1 - v0) >> a = (v1 - v0) / dt
	t.Run("a=(v1-v0)/dt", func(t *testing.T) {
		if x := (v1 - v0).DivDuration(dt); x != a {
			t.Errorf("Expected a=%v, got a=%v", a, x)
		}
	})
}
