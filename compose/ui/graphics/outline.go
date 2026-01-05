package graphics

import "github.com/zodimo/go-compose/compose/ui/geometry"

type Outline interface {
	private()
	Bounds() geometry.Rect
	Equals(other Outline) bool
}

var _ Outline = (*RectangleOutline)(nil)

type RectangleOutline struct {
	rect geometry.Rect
}

func NewRectangleOutline(rect geometry.Rect) *RectangleOutline {
	return &RectangleOutline{
		rect: rect,
	}
}

func (r *RectangleOutline) private() {}

func (r *RectangleOutline) Bounds() geometry.Rect {
	return r.rect
}

func (r *RectangleOutline) Equals(other Outline) bool {
	return r == other
}

var _ Outline = (*RoundedOutline)(nil)

type RoundedOutline struct {
	roundRect geometry.RoundRect
}

func NewRoundedOutline(roundRect geometry.RoundRect) *RoundedOutline {
	return &RoundedOutline{
		roundRect: roundRect,
	}
}

func (r *RoundedOutline) private() {}

func (r *RoundedOutline) Bounds() geometry.Rect {
	return geometry.Rect{}
}

func (r *RoundedOutline) Equals(other Outline) bool {
	return r == other
}

var _ Outline = (*GenericOutline)(nil)

type GenericOutline struct {
	path Path
}

func (g *GenericOutline) private() {}

func (g *GenericOutline) Bounds() geometry.Rect {
	return g.path.GetBounds()
}

func (g *GenericOutline) Equals(other Outline) bool {
	return false
}
