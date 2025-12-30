package graphics

// PointMode defines the manner in which a sequence of points should be drawn.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/PointMode.kt
type PointMode int

const (
	// PointModePoints draws each point as a dot at the specified position.
	PointModePoints PointMode = iota

	// PointModeLines draws each pair of points as a line segment.
	// If there are an odd number of points, the last point is ignored.
	PointModeLines

	// PointModePolygon draws the list of points as a polygon, connecting consecutive points.
	// Unlike Lines, this connects each point to the next rather than treating pairs separately.
	PointModePolygon
)

// String returns the string representation of the PointMode.
func (p PointMode) String() string {
	switch p {
	case PointModePoints:
		return "Points"
	case PointModeLines:
		return "Lines"
	case PointModePolygon:
		return "Polygon"
	default:
		return "Unknown"
	}
}
