//go:generate units-codegen -o distance_generated.go -t Distance generate.json
//go:generate gofmt -w distance_generated.go

package linear

import "math"

// Distance in one dimension or one direction stored as float64 millimeters.
type Distance float64

// Hypot returns the hypotenuse of a right-angled triangle with sides x and y as a new Distance.
func Hypot(x, y Distance) Distance {
	return Distance(math.Hypot(float64(x), float64(y)))
}
