package shape

import (
	"fmt"
	"image"

	"gioui.org/op/clip"
	gioUnit "gioui.org/unit"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

var RoundedCornerShapeUnspecified = &RoundedCornerShape{
	Radius: unit.DpUnspecified,

	TopStart:    unit.DpUnspecified,
	TopEnd:      unit.DpUnspecified,
	BottomEnd:   unit.DpUnspecified,
	BottomStart: unit.DpUnspecified,
}

// RoundedCornerShape supports uniform radius via Radius, or per-corner radius
// via TopStart (NW), TopEnd (NE), BottomEnd (SE), BottomStart (SW).
// Per-corner fields take precedence when any are non-zero.
// This follows Jetpack Compose's RoundedCornerShape API.
type RoundedCornerShape struct {
	// Uniform radius applied to all corners (used when per-corner fields are all zero)
	Radius unit.Dp

	// Per-corner radius (following LTR layout direction):
	// TopStart = NW (top-left), TopEnd = NE (top-right)
	// BottomEnd = SE (bottom-right), BottomStart = SW (bottom-left)
	TopStart    unit.Dp
	TopEnd      unit.Dp
	BottomEnd   unit.Dp
	BottomStart unit.Dp
}

func (r *RoundedCornerShape) CreateOutline(size image.Point, metric gioUnit.Metric) Outline {
	rValid := coalesceRoundedCornerShape(r, RoundedCornerShapeUnspecified)

	// Determine if per-corner radius is being used
	hasPerCorner := rValid.TopStart.IsSpecified() || rValid.TopEnd.IsSpecified() || rValid.BottomEnd.IsSpecified() || rValid.BottomStart.IsSpecified()

	var nw, ne, se, sw int
	if hasPerCorner {
		nw = metric.Dp(unit.DpToGioUnit(rValid.TopStart.TakeOrElse(unit.Dp(0))))
		ne = metric.Dp(unit.DpToGioUnit(rValid.TopEnd.TakeOrElse(unit.Dp(0))))
		se = metric.Dp(unit.DpToGioUnit(rValid.BottomEnd.TakeOrElse(unit.Dp(0))))
		sw = metric.Dp(unit.DpToGioUnit(rValid.BottomStart.TakeOrElse(unit.Dp(0))))
	} else {
		radius := metric.Dp(unit.DpToGioUnit(rValid.Radius.TakeOrElse(unit.Dp(0))))
		if radius == 0 {
			return rectOutline{clip.Rect{Max: size}}
		}
		nw, ne, se, sw = radius, radius, radius, radius
	}

	return rrectOutline{clip.RRect{
		Rect: image.Rectangle{Max: size},
		NW:   nw,
		NE:   ne,
		SE:   se,
		SW:   sw,
	}}
}

func (r *RoundedCornerShape) mergeShape(other Shape) Shape {
	if otherRoundedCorner, ok := other.(*RoundedCornerShape); ok {
		return &RoundedCornerShape{
			Radius:      otherRoundedCorner.Radius.TakeOrElse(r.Radius),
			TopStart:    otherRoundedCorner.TopStart.TakeOrElse(r.TopStart),
			TopEnd:      otherRoundedCorner.TopEnd.TakeOrElse(r.TopEnd),
			BottomEnd:   otherRoundedCorner.BottomEnd.TakeOrElse(r.BottomEnd),
			BottomStart: otherRoundedCorner.BottomStart.TakeOrElse(r.BottomStart),
		}
	}
	panic(fmt.Sprintf("RoundedCornerShape.mergeShape: cannot merge with %s", other.stringShape()))
}
func (r *RoundedCornerShape) sameShape(other Shape) bool {
	if _, ok := other.(*RoundedCornerShape); ok {
		return true
	}
	return false
}
func (r *RoundedCornerShape) semanticEqualShape(other Shape) bool {
	if otherShape, ok := other.(*RoundedCornerShape); ok {
		return otherShape.Radius == r.Radius
	}
	return false
}
func (r *RoundedCornerShape) copyShape(options ...ShapeOption) Shape {

	rValid := coalesceRoundedCornerShape(r, RoundedCornerShapeUnspecified)

	copy := *rValid
	for _, option := range options {
		option(&copy)
	}
	return &copy
}

func (r *RoundedCornerShape) stringShape() string {
	if !isSpecifiedRoundedCornerShape(r) {
		return "RoundedCornerShape{Unspecified}"
	}
	return fmt.Sprintf(
		"RoundedCornerShape{Radius: %s, TopStart: %s, TopEnd: %s, BottomEnd: %s, BottomStart: %s}",
		r.Radius.String(), r.TopStart.String(), r.TopEnd.String(), r.BottomEnd.String(), r.BottomStart.String(),
	)
}

func isSpecifiedRoundedCornerShape(r *RoundedCornerShape) bool {
	return r != nil && r != RoundedCornerShapeUnspecified
}

func coalesceRoundedCornerShape(ptr, def *RoundedCornerShape) *RoundedCornerShape {
	if ptr == nil {
		return def
	}
	return ptr
}
