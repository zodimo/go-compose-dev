package theme

// GradientDirection specifies the direction of a gradient
type GradientDirection int

const (
	// GradientVertical creates a top-to-bottom gradient
	GradientVertical GradientDirection = iota
	// GradientHorizontal creates a left-to-right gradient
	GradientHorizontal
	// GradientDiagonal creates a top-left to bottom-right gradient
	GradientDiagonal
)

// GradientStop represents a color point in a gradient
type GradientStop struct {
	// Position in the gradient from 0.0 (start) to 1.0 (end)
	Position float32
	// Color descriptor for this stop, will be resolved at runtime
	Color ColorDescriptor
}

// GradientConfig defines a theme-aware gradient
// All colors are ColorDescriptors that will be resolved at runtime,
// allowing gradients to adapt when themes change
type GradientConfig struct {
	// Stops define the color points in the gradient
	Stops []GradientStop
	// Direction specifies how the gradient flows
	Direction GradientDirection
}

// NewGradientConfig creates a new gradient configuration
func NewGradientConfig(stops []GradientStop, direction GradientDirection) GradientConfig {
	return GradientConfig{
		Stops:     stops,
		Direction: direction,
	}
}

// VerticalGradient creates a top-to-bottom gradient
func VerticalGradient(stops []GradientStop) GradientConfig {
	return NewGradientConfig(stops, GradientVertical)
}

// HorizontalGradient creates a left-to-right gradient
func HorizontalGradient(stops []GradientStop) GradientConfig {
	return NewGradientConfig(stops, GradientHorizontal)
}

// DiagonalGradient creates a top-left to bottom-right gradient
func DiagonalGradient(stops []GradientStop) GradientConfig {
	return NewGradientConfig(stops, GradientDiagonal)
}
