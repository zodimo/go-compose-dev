package image

import (
	"image"

	"github.com/zodimo/go-compose/compose/ui/layout"
	"github.com/zodimo/go-compose/internal/layoutnode"
	"github.com/zodimo/go-compose/modifiers/size"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// ImageWidget is a widget that displays an image.
type ImageWidget struct {
	// Src is the image to display.
	Src paint.ImageOp

	// ContentScale specifies how to scale the image to the constraints.
	ContentScale layout.ContentScale

	// Alignment specifies where to position the image within the constraints.
	Alignment size.Alignment

	// Alpha applies opacity to the image.
	Alpha float32
}

func (im ImageWidget) Layout(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {
	srcSize := im.Src.Size()

	// Determine destination size for scaling calculations.
	// We use the maximum constraints as the target "destination" box.
	dstSize := gtx.Constraints.Max

	srcDim := layoutnode.LayoutDimensions{Size: srcSize}
	dstDim := layoutnode.LayoutDimensions{Size: dstSize}

	// Calculate scale factor using ContentScale
	scaleFactor := im.ContentScale.Scale(srcDim, dstDim)

	// Calculate the scaled size of the content (image)
	scaledW := int(float32(srcSize.X) * scaleFactor.ScaleX)
	scaledH := int(float32(srcSize.Y) * scaleFactor.ScaleY)
	scaledSize := image.Point{X: scaledW, Y: scaledH}

	// Determine the reported size of the widget.
	// This usually depends on the ContentScale and constraints.
	// For Fit/Contain, we report the scaled size constrained.
	// For Crop/Cover, we report the constraints size (filling it).
	// However, a generic way is to constrain the scaled size ???
	// Wait, if we use FillBounds, scaledSize == dstSize.
	// If we use Fit, scaledSize <= dstSize.
	// So constraining scaledSize seems correct for "Wrap Content" behavior relative to Dst?
	// But if Dst is huge (FillMax), we want to fill it?

	// The `Fit.scale` logic used `cs.Constrain(scaledSize)`.
	// If scaledSize is smaller than Min, it grows to Min.
	// If larger than Max, it shrinks to Max.

	// Let's resolve the reported size.
	// If ContentScaleFillBounds, scaledSize is Max.
	// If ContentScaleFit, scaledSize fits inside Max.

	// Issue: if Constraints.Min is distinct from Max.
	// e.g. Min(0,0), Max(100,100).
	// Fit: image becomes e.g. 100x50. Reported size 100x50. Align Center -> Centered in what?
	// If we return 100x50, the parent might center this 100x50 box.
	// But `Image` composable often implies filling the allotted space if modifiers say so.
	// BUT `ImageWidget` here is a leaf widget.
	// In Compose, `Image` layouts usually take the size of the image (scaled) unless modifier overrides size.
	// If `modifier.Size(100)` is used, constraints are Max=100, Min=100.
	// Then `reportedSize` will be 100x100.

	reportedSize := gtx.Constraints.Constrain(scaledSize)

	// Calculate offset for alignment
	offset := im.Alignment.Align(scaledSize, reportedSize, layoutnode.LayoutDirectionLTR)

	// Clip to reported bounds
	defer clip.Rect{Max: reportedSize}.Push(gtx.Ops).Pop()

	// Gio Affine2D matrix:
	// | a  b  c |   where a=scaleX, b=shearX, c=offsetX
	// | d  e  f |   where d=shearY, e=scaleY, f=offsetY
	// NewAffine2D(a, b, c, d, e, f)
	trans := f32.NewAffine2D(
		scaleFactor.ScaleX, 0, float32(offset.X),
		0, scaleFactor.ScaleY, float32(offset.Y),
	)
	defer op.Affine(trans).Push(gtx.Ops).Pop()

	// Draw Image
	im.Src.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return layoutnode.LayoutDimensions{Size: reportedSize}
}
