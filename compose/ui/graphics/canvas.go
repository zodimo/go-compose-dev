package graphics

import (
	"github.com/zodimo/go-compose/compose/ui/geometry"
)

// Canvas provides a low-level interface for drawing operations.
// This is an interface that will be implemented by platform-specific backends (e.g., Skia).
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/Canvas.kt
type Canvas interface {
	// Save saves a copy of the current transform and clip on the save stack.
	Save()

	// Restore pops the current save stack, restoring the previous transform and clip.
	Restore()

	// SaveLayer saves the current state and creates a new group for subsequent operations.
	// When restored, the layer is composited into the previous layer using the paint's blend mode.
	SaveLayer(bounds geometry.Rect, paint *Paint)

	// Translate shifts the coordinate space by the given delta.
	Translate(dx, dy float32)

	// Scale scales the coordinate space by the given factors.
	Scale(sx, sy float32)

	// Rotate rotates the coordinate space by the given degrees clockwise.
	Rotate(degrees float32)

	// Skew applies an axis-aligned skew transformation.
	Skew(sx, sy float32)

	// Concat multiplies the current transform by the specified matrix.
	Concat(matrix Matrix)

	// ClipRect reduces the clip region to the intersection of the current clip and the given rectangle.
	ClipRect(left, top, right, bottom float32, clipOp ClipOp)

	// ClipPath reduces the clip region to the intersection of the current clip and the given path.
	ClipPath(path Path, clipOp ClipOp)

	// DrawLine draws a line between the given points.
	DrawLine(p1, p2 geometry.Offset, paint *Paint)

	// DrawRect draws a rectangle.
	DrawRect(left, top, right, bottom float32, paint *Paint)

	// DrawRoundRect draws a rounded rectangle.
	DrawRoundRect(left, top, right, bottom, radiusX, radiusY float32, paint *Paint)

	// DrawOval draws an axis-aligned oval that fills the given rectangle.
	DrawOval(left, top, right, bottom float32, paint *Paint)

	// DrawCircle draws a circle at the given center with the given radius.
	DrawCircle(center geometry.Offset, radius float32, paint *Paint)

	// DrawArc draws an arc scaled to fit inside the given rectangle.
	DrawArc(left, top, right, bottom, startAngle, sweepAngle float32, useCenter bool, paint *Paint)

	// DrawPath draws the given path.
	DrawPath(path Path, paint *Paint)

	// DrawImage draws an image at the given offset.
	DrawImage(image ImageBitmap, topLeftOffset geometry.Offset, paint *Paint)

	// DrawImageRect draws a portion of an image into a destination rectangle.
	DrawImageRect(image ImageBitmap, srcOffset IntOffset, srcSize IntSize, dstOffset IntOffset, dstSize IntSize, paint *Paint)

	// DrawPoints draws a sequence of points according to the given PointMode.
	DrawPoints(pointMode PointMode, points []geometry.Offset, paint *Paint)

	// EnableZ enables Z-ordering for 3D-like effects (Android-specific, may be no-op on other platforms).
	EnableZ()

	// DisableZ disables Z-ordering.
	DisableZ()
}

// IntOffset represents an integer 2D offset.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-unit/src/commonMain/kotlin/androidx/compose/ui/unit/IntOffset.kt
type IntOffset struct {
	X int
	Y int
}

// IntOffsetZero is an IntOffset with zero values.
var IntOffsetZero = IntOffset{X: 0, Y: 0}

// IntSize represents an integer 2D size.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-unit/src/commonMain/kotlin/androidx/compose/ui/unit/IntSize.kt
type IntSize struct {
	Width  int
	Height int
}

// IntSizeZero is an IntSize with zero dimensions.
var IntSizeZero = IntSize{Width: 0, Height: 0}

// Matrix represents a 4x4 transformation matrix.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/Matrix.kt
type Matrix [16]float32

// MatrixIdentity is the identity matrix.
var MatrixIdentity = Matrix{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
}

// WithSave executes the given block with a save/restore pair.
func WithSave(canvas Canvas, block func()) {
	canvas.Save()
	defer canvas.Restore()
	block()
}

// WithSaveLayer executes the given block with a saveLayer/restore pair.
func WithSaveLayer(canvas Canvas, bounds geometry.Rect, paint *Paint, block func()) {
	canvas.SaveLayer(bounds, paint)
	defer canvas.Restore()
	block()
}
