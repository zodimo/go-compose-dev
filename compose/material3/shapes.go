package material3

import (
	"fmt"

	"github.com/zodimo/go-compose/compose"
	fShape "github.com/zodimo/go-compose/compose/foundation/shape"
	"github.com/zodimo/go-compose/compose/material3/tokens"
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
)

var DefaultShapes = &Shapes{}

type Shapes struct {
	ExtraSmall          shape.Shape
	Small               shape.Shape
	Medium              shape.Shape
	Large               shape.Shape
	LargeIncreased      shape.Shape
	ExtraLarge          shape.Shape
	ExtraLargeIncreased shape.Shape
	ExtraExtraLarge     shape.Shape

	CornerNone                fShape.CornerSize
	CornerExtraSmall          fShape.CornerSize
	CornerSmall               fShape.CornerSize
	CornerMedium              fShape.CornerSize
	CornerLarge               fShape.CornerSize
	CornerLargeIncreased      fShape.CornerSize
	CornerExtraLarge          fShape.CornerSize
	CornerExtraLargeIncreased fShape.CornerSize
	CornerExtraExtraLarge     fShape.CornerSize
	CornerFull                fShape.CornerSize
}

func (s *Shapes) Top(bottomSize fShape.CornerSize) shape.Shape {
	bottomSize = fShape.TakeOrElseCornerSize(bottomSize, s.CornerNone)
	return nil
}
func (s *Shapes) Bottom(topSize fShape.CornerSize) shape.Shape {
	topSize = fShape.TakeOrElseCornerSize(topSize, s.CornerNone)
	return nil
}
func (s *Shapes) Start(endSize fShape.CornerSize) shape.Shape {
	endSize = fShape.TakeOrElseCornerSize(endSize, s.CornerNone)
	return nil
}
func (s *Shapes) End(startSize fShape.CornerSize) shape.Shape {
	startSize = fShape.TakeOrElseCornerSize(startSize, s.CornerNone)
	return nil
}

func (s *Shapes) FromToken(value tokens.ShapeTokenKey) shape.Shape {
	switch value {
	case tokens.ShapeTokenKeyCornerExtraLarge:
		return s.ExtraLarge
	case tokens.ShapeTokenKeyCornerExtraLargeIncreased:
		return s.ExtraLargeIncreased
	case tokens.ShapeTokenKeyCornerExtraExtraLarge:
		return s.ExtraExtraLarge
	// case tokens.ShapeTokenKeyCornerExtraLargeTop:
	// 	return s.ExtraLarge.Top()
	case tokens.ShapeTokenKeyCornerExtraSmall:
		return s.ExtraSmall
	// case tokens.ShapeTokenKeyCornerExtraSmallTop:
	// 	return s.ExtraSmall.Top()
	case tokens.ShapeTokenKeyCornerFull:
		return shape.CircleShape
	case tokens.ShapeTokenKeyCornerLarge:
		return s.Large
	case tokens.ShapeTokenKeyCornerLargeIncreased:
		return s.LargeIncreased
	// case tokens.ShapeTokenKeyCornerLargeEnd:
	// 	return s.Large.End()
	// case tokens.ShapeTokenKeyCornerLargeTop:
	// 	return s.Large.Top()
	case tokens.ShapeTokenKeyCornerMedium:
		return s.Medium
	case tokens.ShapeTokenKeyCornerNone:
		return shape.RectangleShape
	case tokens.ShapeTokenKeyCornerSmall:
		return s.Small
	// case tokens.ShapeTokenKeyCornerLargeStart:
	// 	return s.Large.Start()

	default:
		panic(fmt.Sprintf("unknown shape token key: %s", value))
	}

}

var LocalShapes = compose.CompositionLocalOf(func() *Shapes {
	return DefaultShapes
})
