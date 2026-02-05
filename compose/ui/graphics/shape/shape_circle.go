package shape

import (
	"image"

	"gioui.org/op"
	"gioui.org/op/clip"
	gioUnit "gioui.org/unit"
)

// Deprecated: Use CircleShape instead
var ShapeCircle Shape = &circleShape{}

// CircleShape is a shape describing a circle.
var CircleShape Shape = &circleShape{}

// CircleShape
type circleShape struct{}

func (c *circleShape) CreateOutline(size image.Point, metric gioUnit.Metric) Outline {
	radius := min(size.X, size.Y) / 2
	return rrectOutline{clip.RRect{
		Rect: image.Rectangle{Max: size},
		SE:   radius,
		SW:   radius,
		NW:   radius,
		NE:   radius,
	}}
}

// Sentinel Private methods
func (c *circleShape) mergeShape(other Shape) Shape {
	if otherCircle, ok := other.(*circleShape); ok {
		return otherCircle
	}
	return c
}
func (c *circleShape) sameShape(other Shape) bool {
	if _, ok := other.(*circleShape); ok {
		return true
	}
	return false
}
func (c *circleShape) semanticEqualShape(other Shape) bool {
	if _, ok := other.(*circleShape); ok {
		return true
	}
	return false
}
func (c *circleShape) copyShape(options ...ShapeOption) Shape {
	copy := *c
	if len(options) > 0 {
		// should we panic here ?
		return &copy
	}
	return &copy
}
func (c *circleShape) stringShape() string {
	return "CircleShape"
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

// Ellipse.Path takes ops argument? No, Ellipse.Path(ops) returns PathSpec.
// checking docs/source... Ellipse.Path(ops) -> PathSpec.
func (e ellipseOutline) Path(ops *op.Ops) clip.PathSpec {
	return e.Ellipse.Path(ops)
}
