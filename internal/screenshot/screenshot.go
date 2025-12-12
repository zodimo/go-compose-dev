package screenshot

import (
	"bytes"
	"image"
	"image/png"

	"gioui.org/gpu/headless"
	"gioui.org/op"
)

func TakeScreenshot(width, height int, drawOps op.CallOp) image.Image {

	cap := image.NewRGBA(image.Rect(0, 0, width, height))
	w, _ := headless.NewWindow(width, height)
	defer w.Release()

	ops := new(op.Ops)

	drawOps.Add(ops)

	w.Frame(ops)
	w.Screenshot(cap)
	b := new(bytes.Buffer)
	_ = png.Encode(b, cap)

	return cap
}
