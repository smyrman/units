//go:generate units-codegen -o area_generated.go -t Area generate.json
//go:generate gofmt -w area_generated.go

package planar

import "github.com/smyrman/units/linear"

// Area stores a planar area as float64 squear millimeters.
type Area float64

// MulDistance returns the product of x * y as a new Area.
func MulDistance(x, y linear.Distance) Area {
	return Area(x * y)
}
