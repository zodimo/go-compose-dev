package graphics

import (
	"github.com/zodimo/go-compose/compose/ui/geometry"
)

// ImageBitmap represents an image that can be drawn onto a canvas.
// This is an interface that will be implemented by platform-specific image types.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/ImageBitmap.kt
type ImageBitmap interface {
	// Width returns the width of the image in pixels.
	Width() int

	// Height returns the height of the image in pixels.
	Height() int

	// Size returns the size of the image.
	Size() geometry.Size
}
