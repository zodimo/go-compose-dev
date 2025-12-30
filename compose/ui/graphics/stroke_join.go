package graphics

// StrokeJoin defines the shape to be used at the corners of stroked paths.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/StrokeJoin.kt
type StrokeJoin int

const (
	// StrokeJoinMiter creates a sharp corner where line segments join.
	// The miter limit determines when the corner is beveled if too sharp.
	StrokeJoinMiter StrokeJoin = iota

	// StrokeJoinRound creates a rounded corner with radius equal to half the stroke width.
	StrokeJoinRound

	// StrokeJoinBevel creates a beveled corner where the outer edges meet with a straight line.
	StrokeJoinBevel
)

// String returns the string representation of the StrokeJoin.
func (s StrokeJoin) String() string {
	switch s {
	case StrokeJoinMiter:
		return "Miter"
	case StrokeJoinRound:
		return "Round"
	case StrokeJoinBevel:
		return "Bevel"
	default:
		return "Unknown"
	}
}
