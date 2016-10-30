//go:generate units-codegen -o velocity_generated.go -t Velocity generate.json
//go:generate gofmt -w velocity_generated.go

package angular

// Velocity describes an angular velocity in planar geometry.
type Velocity float64
