package graphics

import (
	"github.com/zodimo/go-compose/compose/ui/geometry"
)

// WithTransform performs transformations and executes drawing commands within the transformed space.
// After the block completes, the transformation is restored.
func WithTransform(scope DrawScope, transformBlock func(DrawTransform), drawBlock func(DrawScope)) {
	ctx := scope.DrawContext()
	canvas := ctx.Canvas()
	originalSize := ctx.Size()

	canvas.Save()
	transformBlock(ctx.Transform())
	drawBlock(scope)
	canvas.Restore()

	// Restore original size in case inset was called
	ctx.SetSize(originalSize)
}

// WithInset translates the coordinate space and modifies the drawing bounds.
// The width becomes width - (left + right), height becomes height - (top + bottom).
func WithInset(scope DrawScope, left, top, right, bottom float32, block func(DrawScope)) {
	WithTransform(scope, func(t DrawTransform) {
		t.Inset(left, top, right, bottom)
	}, block)
}

// WithInsetSymmetric is a convenience function for symmetric horizontal and vertical insets.
func WithInsetSymmetric(scope DrawScope, horizontal, vertical float32, block func(DrawScope)) {
	WithInset(scope, horizontal, vertical, horizontal, vertical, block)
}

// WithInsetAll is a convenience function that insets all sides equally.
func WithInsetAll(scope DrawScope, inset float32, block func(DrawScope)) {
	WithInset(scope, inset, inset, inset, inset, block)
}

// WithTranslate translates the coordinate space.
func WithTranslate(scope DrawScope, left, top float32, block func(DrawScope)) {
	WithTransform(scope, func(t DrawTransform) {
		t.Translate(left, top)
	}, block)
}

// WithRotate rotates the coordinate space around the given pivot point.
func WithRotate(scope DrawScope, degrees float32, pivot geometry.Offset, block func(DrawScope)) {
	WithTransform(scope, func(t DrawTransform) {
		t.Rotate(degrees, pivot)
	}, block)
}

// WithRotateCenter rotates the coordinate space around the center.
func WithRotateCenter(scope DrawScope, degrees float32, block func(DrawScope)) {
	WithRotate(scope, degrees, scope.Center(), block)
}

// WithRotateRad rotates the coordinate space by radians around the given pivot point.
func WithRotateRad(scope DrawScope, radians float32, pivot geometry.Offset, block func(DrawScope)) {
	WithRotate(scope, Degrees(radians), pivot, block)
}

// WithScale scales the coordinate space around the given pivot point.
func WithScale(scope DrawScope, scaleX, scaleY float32, pivot geometry.Offset, block func(DrawScope)) {
	WithTransform(scope, func(t DrawTransform) {
		t.Scale(scaleX, scaleY, pivot)
	}, block)
}

// WithScaleCenter scales uniformly around the center.
func WithScaleCenter(scope DrawScope, scale float32, block func(DrawScope)) {
	WithScale(scope, scale, scale, scope.Center(), block)
}

// WithClipRect reduces the clip region to the given rectangle.
func WithClipRect(scope DrawScope, left, top, right, bottom float32, clipOp ClipOp, block func(DrawScope)) {
	WithTransform(scope, func(t DrawTransform) {
		t.ClipRect(left, top, right, bottom, clipOp)
	}, block)
}

// WithClipRectBounds clips to the current drawing bounds.
func WithClipRectBounds(scope DrawScope, clipOp ClipOp, block func(DrawScope)) {
	size := scope.Size()
	WithClipRect(scope, 0, 0, size.Width(), size.Height(), clipOp, block)
}

// WithClipPath reduces the clip region to the given path.
func WithClipPath(scope DrawScope, path Path, clipOp ClipOp, block func(DrawScope)) {
	WithTransform(scope, func(t DrawTransform) {
		t.ClipPath(path, clipOp)
	}, block)
}
