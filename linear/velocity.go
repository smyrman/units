//go:generate units-codegen -o velocity_generated.go -t Velocity generate.json
//go:generate gofmt -w velocity_generated.go

package linear

// Velocity in one dimension stored as float64 millimeters per second.
type Velocity float64
