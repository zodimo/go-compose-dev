package graphics

// StrokeCap defines the shape to be used at the ends of open subpaths when they are stroked.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/StrokeCap.kt
type StrokeCap int

const (
	// StrokeCapButt ends with a flat edge and no extension.
	StrokeCapButt StrokeCap = iota

	// StrokeCapRound ends with a semicircle with diameter equal to the stroke width.
	StrokeCapRound

	// StrokeCapSquare ends with a half-square with side equal to the stroke width.
	StrokeCapSquare
)

// String returns the string representation of the StrokeCap.
func (s StrokeCap) String() string {
	switch s {
	case StrokeCapButt:
		return "Butt"
	case StrokeCapRound:
		return "Round"
	case StrokeCapSquare:
		return "Square"
	default:
		return "Unknown"
	}
}
