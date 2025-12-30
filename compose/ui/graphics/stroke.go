package graphics

import "math"

// Stroke is a DrawStyle that provides information for drawing content with a stroke.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/drawscope/DrawScope.kt;l=960
type Stroke struct {
	// Width configures the width of the stroke in pixels.
	Width float32

	// Miter is the stroke miter value. Used to control the behavior of miter joins
	// when the joins angle is sharp. A value of 4.0 is the default.
	Miter float32

	// Cap is the treatment applied to the beginning and end of stroked lines and paths.
	Cap StrokeCap

	// Join is the treatment applied where lines and curve segments join on a stroked path.
	Join StrokeJoin

	// PathEffect is an optional effect to apply to the stroke.
	// Currently a placeholder - PathEffect interface needs separate implementation.
	PathEffect interface{}
}

// Default stroke constants
const (
	// DefaultStrokeWidth is the default stroke width (hairline).
	DefaultStrokeWidth float32 = 0.0

	// DefaultStrokeMiter is the default miter limit.
	DefaultStrokeMiter float32 = 4.0
)

// DefaultStrokeCap is the default stroke cap.
var DefaultStrokeCap = StrokeCapButt

// DefaultStrokeJoin is the default stroke join.
var DefaultStrokeJoin = StrokeJoinMiter

// NewStroke creates a new Stroke with the given width and default values.
func NewStroke(width float32) *Stroke {
	return &Stroke{
		Width:      width,
		Miter:      DefaultStrokeMiter,
		Cap:        DefaultStrokeCap,
		Join:       DefaultStrokeJoin,
		PathEffect: nil,
	}
}

// NewStrokeWithOptions creates a new Stroke with all customizable options.
func NewStrokeWithOptions(width, miter float32, cap StrokeCap, join StrokeJoin, pathEffect interface{}) *Stroke {
	return &Stroke{
		Width:      width,
		Miter:      miter,
		Cap:        cap,
		Join:       join,
		PathEffect: pathEffect,
	}
}

// Hairline represents a hairline stroke (the thinnest line that can be drawn).
var Hairline = NewStroke(DefaultStrokeWidth)

func (s *Stroke) isDrawStyle() {}

// Equal checks if two Stroke instances are equal.
func (s *Stroke) Equal(other *Stroke) bool {
	if s == other {
		return true
	}
	if s == nil || other == nil {
		return false
	}
	// Use tolerance for float comparison
	const epsilon = 1e-6
	widthEq := math.Abs(float64(s.Width-other.Width)) < epsilon
	miterEq := math.Abs(float64(s.Miter-other.Miter)) < epsilon
	return widthEq && miterEq && s.Cap == other.Cap && s.Join == other.Join
	// Note: PathEffect comparison would need separate handling
}

// String returns a string representation of the Stroke.
func (s *Stroke) String() string {
	return "Stroke(width=" + formatFloat(s.Width) + ", miter=" + formatFloat(s.Miter) +
		", cap=" + s.Cap.String() + ", join=" + s.Join.String() + ")"
}

func formatFloat(f float32) string {
	// Simple float formatting
	return string(rune(int(f*10)/10)) + "." + string(rune(int(f*10)%10))
}
