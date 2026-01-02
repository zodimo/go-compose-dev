package shape

import (
	"fmt"
	"image"

	"gioui.org/op"
	"gioui.org/op/clip"
	gioUnit "gioui.org/unit"
)

// https://developer.android.com/reference/kotlin/androidx/compose/ui/graphics/Shape

type ShapeOption func(Shape)

type Shape interface {
	CreateOutline(size image.Point, metric gioUnit.Metric) Outline
	mergeShape(other Shape) Shape
	sameShape(other Shape) bool
	semanticEqualShape(other Shape) bool
	copyShape(options ...ShapeOption) Shape
	stringShape() string
}

type Outline interface {
	Push(ops *op.Ops) clip.Stack
	Op(ops *op.Ops) clip.Op
	Path(ops *op.Ops) clip.PathSpec
}

func IsSpecifiedShape(s Shape) bool {
	return s != nil && s != ShapeUnspecified
}

func TakeOrElseShape(s, defaultStyle Shape) Shape {
	if s == nil || s == ShapeUnspecified {
		return defaultStyle
	}
	return s
}

func CoalesceShape(ptr, def Shape) Shape {
	if ptr == nil {
		return def
	}
	return ptr
}

func MergeShape(a, b Shape) Shape {
	a = CoalesceShape(a, ShapeUnspecified)
	b = CoalesceShape(b, ShapeUnspecified)

	if a == ShapeUnspecified {
		return b
	}
	if b == ShapeUnspecified {
		return a
	}

	return a.mergeShape(b)
}

func SameShape(a, b Shape) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == ShapeUnspecified
	}
	if b == nil {
		return a == ShapeUnspecified
	}
	return a.sameShape(b)
}

func SemanticEqualShape(a, b Shape) bool {

	a = CoalesceShape(a, ShapeUnspecified)
	b = CoalesceShape(b, ShapeUnspecified)

	return a.semanticEqualShape(b)
}

func EqualShape(a, b Shape) bool {
	if !SameShape(a, b) {
		return SemanticEqualShape(a, b)
	}
	return true
}

func CopyShape(s Shape, options ...ShapeOption) Shape {
	switch shape := s.(type) {
	case *circleShape:
		copy := *shape
		for _, option := range options {
			option(&copy)
		}
		return &copy
	case *rectangleShape:
		copy := *shape
		for _, option := range options {
			option(&copy)
		}
		return &copy
	case *RoundedCornerShape:
		copy := *shape
		for _, option := range options {
			option(&copy)
		}
		return &copy
	case *CutCornerShape:
		copy := *shape
		for _, option := range options {
			option(&copy)
		}
		return &copy

	default:
		panic(fmt.Sprintf("CopyShape: unknown shape type %s", s.stringShape()))
	}
}

func StringShape(shape Shape) string {
	return shape.stringShape()
}
