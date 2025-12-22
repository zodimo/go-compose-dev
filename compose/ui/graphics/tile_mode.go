package graphics

// TileMode defines what happens at the edge of the gradient.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/TileMode.kt
type TileMode int

const (
	// Edge is clamped to the final color.
	TileModeClamp TileMode = 0
	// Edge is repeated from first color to last.
	TileModeRepeated TileMode = 1
	// Edge is mirrored from last color to first.
	TileModeMirror TileMode = 2
	// Render the shader's image pixels only within its original bounds.
	TileModeDecal TileMode = 3
)

func (t TileMode) String() string {
	switch t {
	case TileModeClamp:
		return "Clamp"
	case TileModeRepeated:
		return "Repeated"
	case TileModeMirror:
		return "Mirror"
	case TileModeDecal:
		return "Decal"
	default:
		return "Unknown"
	}
}

// IsSupported returns true if the TileMode is supported.
// Clamp, Repeated, and Mirror are guaranteed to be supported.
func (t TileMode) IsSupported() bool {
	// Assuming all are supported for now, or match Kotlin's expect/actual logic which implies specific platform support.
	// For Go port, we can default to true or restrict Decal if needed.
	// Kotlin docs say: Clamp, Repeated, Mirror are guaranteed. Decal might not be.
	return true
}
