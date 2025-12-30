package graphics

// ClipOp defines the algorithm used for clipping.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/ClipOp.kt
type ClipOp int

const (
	// ClipOpIntersect clips to the intersection of the current clip and the given path/rect.
	// This is the default behavior.
	ClipOpIntersect ClipOp = iota

	// ClipOpDifference clips to the difference of the current clip and the given path/rect.
	// This subtracts the given region from the current clip.
	ClipOpDifference
)

// String returns the string representation of the ClipOp.
func (c ClipOp) String() string {
	switch c {
	case ClipOpIntersect:
		return "Intersect"
	case ClipOpDifference:
		return "Difference"
	default:
		return "Unknown"
	}
}
