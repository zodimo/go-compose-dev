package graphics

import (
	"github.com/zodimo/go-compose/compose/ui/geometry"
)

// DrawTransform defines transformations that can be applied to a drawing environment.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/drawscope/DrawTransform.kt
type DrawTransform interface {
	// Size returns the current size of the drawing environment.
	Size() geometry.Size

	// Center returns the center offset of the current transformation.
	Center() geometry.Offset

	// Inset simultaneously translates the coordinate space and modifies the drawing bounds.
	// The width becomes width - (left + right), height becomes height - (top + bottom).
	Inset(left, top, right, bottom float32)

	// ClipRect reduces the clip region to the intersection of the current clip
	// and the given rectangle.
	ClipRect(left, top, right, bottom float32, clipOp ClipOp)

	// ClipPath reduces the clip region to the intersection of the current clip
	// and the given path.
	ClipPath(path Path, clipOp ClipOp)

	// Translate translates the coordinate space by the given delta.
	Translate(left, top float32)

	// Rotate adds a rotation (in degrees clockwise) to the current transform
	// at the given pivot point.
	Rotate(degrees float32, pivot geometry.Offset)

	// Scale adds an axis-aligned scale to the current transform at the given pivot point.
	Scale(scaleX, scaleY float32, pivot geometry.Offset)

	// Transform applies the given transformation matrix to the drawing environment.
	Transform(matrix Matrix)
}

// InsetSymmetric is a convenience method that insets both left/right by horizontal
// and top/bottom by vertical.
func InsetSymmetric(t DrawTransform, horizontal, vertical float32) {
	t.Inset(horizontal, vertical, horizontal, vertical)
}

// InsetAll is a convenience method that insets all sides by the same amount.
func InsetAll(t DrawTransform, inset float32) {
	t.Inset(inset, inset, inset, inset)
}

// RotateRad is a convenience method that rotates by radians instead of degrees.
func RotateRad(t DrawTransform, radians float32, pivot geometry.Offset) {
	t.Rotate(Degrees(radians), pivot)
}

// ScaleUniform is a convenience method that scales uniformly in both directions.
func ScaleUniform(t DrawTransform, scale float32, pivot geometry.Offset) {
	t.Scale(scale, scale, pivot)
}

// Degrees converts radians to degrees.
func Degrees(radians float32) float32 {
	return radians * 180.0 / 3.14159265358979323846
}

// Radians converts degrees to radians.
func Radians(degrees float32) float32 {
	return degrees * 3.14159265358979323846 / 180.0
}
