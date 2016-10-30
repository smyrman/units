//go:generate units-codegen -o acceleration_generated.go -t Acceleration generate.json
//go:generate gofmt -w acceleration_generated.go

package linear

// Acceleration in one dimension or one direction, stored as float64 millimeters per second squared.
type Acceleration float64
