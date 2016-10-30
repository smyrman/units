//go:generate units-codegen -o acceleration_generated.go -t Acceleration generate.json
//go:generate gofmt -w acceleration_generated.go

package angular

// Acceleration describes an angular acceleration in planar geometry.
type Acceleration float64
