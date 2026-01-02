package shape

import (
	"fmt"

	"github.com/zodimo/go-compose/compose/ui/unit"
)

// RoundedCornerShape and CutCornerShape Options

func WithRadius(radius unit.Dp) ShapeOption {
	return func(s Shape) {
		switch shape := s.(type) {
		case *RoundedCornerShape:
			shape.Radius = radius
		case *CutCornerShape:
			shape.Radius = radius
		default:
			panic(fmt.Sprintf("WithRadius: not supported on shape type %s", s.stringShape()))
		}
	}
}

// RoundedCornerShape Options
func WithTopEndRadius(radius unit.Dp) ShapeOption {
	return func(s Shape) {
		if shape, ok := s.(*RoundedCornerShape); ok {
			shape.TopEnd = radius
		}
		panic(fmt.Sprintf("WithTopEndRadius: not supported on shape type %s", s.stringShape()))
	}
}

func WithBottomStartRadius(radius unit.Dp) ShapeOption {
	return func(s Shape) {
		if shape, ok := s.(*RoundedCornerShape); ok {
			shape.BottomStart = radius
		}
		panic(fmt.Sprintf("WithBottomStartRadius: not supported on shape type %s", s.stringShape()))
	}
}

func WithBottomEndRadius(radius unit.Dp) ShapeOption {
	return func(s Shape) {
		if shape, ok := s.(*RoundedCornerShape); ok {
			shape.BottomEnd = radius
		}
		panic(fmt.Sprintf("WithBottomEndRadius: not supported on shape type %s", s.stringShape()))
	}
}
