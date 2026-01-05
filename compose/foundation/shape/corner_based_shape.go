package shape

import (
	"fmt"

	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

var _ graphics.Shape = (*CornerBasedShape)(nil)

// base class
type CornerBasedShape struct {
	TopStart    CornerSize
	TopEnd      CornerSize
	BottomStart CornerSize
	BottomEnd   CornerSize
}

func (s CornerBasedShape) CreateOutline(size geometry.Size, layoutDirection unit.LayoutDirection, density unit.Density) graphics.Outline {
	topStart := s.TopStart.ToPx(size, density)
	topEnd := s.TopEnd.ToPx(size, density)
	bottomEnd := s.BottomEnd.ToPx(size, density)
	bottomStart := s.BottomStart.ToPx(size, density)
	minDimension := size.MinDimension()

	if topStart+bottomStart > minDimension {
		scale := minDimension / (topStart + bottomStart)
		topStart *= scale
		bottomStart *= scale
	}
	if topEnd+bottomEnd > minDimension {
		scale := minDimension / (topEnd + bottomEnd)
		topEnd *= scale
		bottomEnd *= scale
	}
	requirePrecondition(
		topStart >= 0 && topEnd >= 0 && bottomEnd >= 0 && bottomStart >= 0,
		fmt.Sprintf("Corner size in Px can't be negative(topStart = %f, topEnd = %f, bottomEnd = %f, bottomStart = %f)", topStart, topEnd, bottomEnd, bottomStart),
	)

	panic("implementation not complete")

}

func requirePrecondition(predicate bool, message string) {
	if !predicate {
		panic(message)
	}
}

// var topStart = topStart.toPx(size, density)
// var topEnd = topEnd.toPx(size, density)
// var bottomEnd = bottomEnd.toPx(size, density)
// var bottomStart = bottomStart.toPx(size, density)
// val minDimension = size.minDimension
// if (topStart + bottomStart > minDimension) {
// 	val scale = minDimension / (topStart + bottomStart)
// 	topStart *= scale
// 	bottomStart *= scale
// }
// if (topEnd + bottomEnd > minDimension) {
// 	val scale = minDimension / (topEnd + bottomEnd)
// 	topEnd *= scale
// 	bottomEnd *= scale
// }
// requirePrecondition(
// 	topStart >= 0.0f && topEnd >= 0.0f && bottomEnd >= 0.0f && bottomStart >= 0.0f
// ) {
// 	"Corner size in Px can't be negative(topStart = $topStart, topEnd = $topEnd, " +
// 		"bottomEnd = $bottomEnd, bottomStart = $bottomStart)!"
// }
// return createOutline(
// 	size = size,
// 	topStart = topStart,
// 	topEnd = topEnd,
// 	bottomEnd = bottomEnd,
// 	bottomStart = bottomStart,
// 	layoutDirection = layoutDirection,
// )
