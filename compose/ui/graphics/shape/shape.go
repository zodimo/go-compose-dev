package shape

import (
	"image"

	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/unit"
)

// https://developer.android.com/reference/kotlin/androidx/compose/ui/graphics/Shape

type Outline interface {
	Push(ops *op.Ops) clip.Stack
	Op(ops *op.Ops) clip.Op
}

type Shape interface {
	CreateOutline(size image.Point, metric unit.Metric) Outline
}

// RectangleShape
type rectangleShape struct{}

func (r rectangleShape) CreateOutline(size image.Point, metric unit.Metric) Outline {
	return rectOutline{clip.Rect{Max: size}}
}

type rectOutline struct {
	clip.Rect
}

func (r rectOutline) Push(ops *op.Ops) clip.Stack {
	return r.Rect.Push(ops)
}

func (r rectOutline) Op(ops *op.Ops) clip.Op {
	return r.Rect.Op()
}

var ShapeRectangle Shape = rectangleShape{}

// CircleShape
type circleShape struct{}

func (c circleShape) CreateOutline(size image.Point, metric unit.Metric) Outline {
	return ellipseOutline{clip.Ellipse{Max: size}}
}

type ellipseOutline struct {
	clip.Ellipse
}

func (e ellipseOutline) Push(ops *op.Ops) clip.Stack {
	return e.Ellipse.Push(ops)
}

func (e ellipseOutline) Op(ops *op.Ops) clip.Op {
	return e.Ellipse.Op(ops)
}

var ShapeCircle Shape = circleShape{}

// RoundedCornerShape
type RoundedCornerShape struct {
	Radius unit.Dp
}

func (r RoundedCornerShape) CreateOutline(size image.Point, metric unit.Metric) Outline {
	radius := metric.Dp(r.Radius)
	if radius == 0 {
		return rectOutline{clip.Rect{Max: size}}
	}
	return rrectOutline{clip.RRect{
		Rect: image.Rectangle{Max: size},
		SE:   radius,
		SW:   radius,
		NW:   radius,
		NE:   radius,
	}}
}

type rrectOutline struct {
	clip.RRect
}

func (r rrectOutline) Push(ops *op.Ops) clip.Stack {
	return r.RRect.Push(ops)
}

func (r rrectOutline) Op(ops *op.Ops) clip.Op {
	return r.RRect.Op(ops)
}
