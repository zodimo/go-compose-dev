package shape

import (
	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-ternary"
)

var _ graphics.Shape = (*RoundedCornerShape)(nil)

type RoundedCornerShape struct {
	TopStart    CornerSize
	TopEnd      CornerSize
	BottomStart CornerSize
	BottomEnd   CornerSize
}

func (s RoundedCornerShape) CreateOutline(size geometry.Size, layoutDirection unit.LayoutDirection, density unit.Density) graphics.Outline {
	if s.TopStart.isZero() && s.TopEnd.isZero() && s.BottomStart.isZero() && s.BottomEnd.isZero() {
		return graphics.NewRectangleOutline(size.ToRect())
	}
	return graphics.NewRoundedOutline(
		geometry.NewRoundRect(
			size.ToRect(),
			geometry.NewCornerRadiusUniform(ternary.Ternary(
				layoutDirection == unit.LayoutDirectionLtr,
				s.TopStart.ToPx(size, density),
				s.TopEnd.ToPx(size, density),
			)),
			geometry.NewCornerRadiusUniform(ternary.Ternary(
				layoutDirection == unit.LayoutDirectionLtr,
				s.TopEnd.ToPx(size, density),
				s.TopStart.ToPx(size, density),
			)),
			geometry.NewCornerRadiusUniform(ternary.Ternary(
				layoutDirection == unit.LayoutDirectionLtr,
				s.BottomEnd.ToPx(size, density),
				s.BottomStart.ToPx(size, density),
			)),
			geometry.NewCornerRadiusUniform(ternary.Ternary(
				layoutDirection == unit.LayoutDirectionLtr,
				s.BottomStart.ToPx(size, density),
				s.BottomEnd.ToPx(size, density),
			)),
		),
	)

}

func RoundedCornerShapeAll(cornerSize CornerSize) RoundedCornerShape {
	return RoundedCornerShape{
		TopStart:    cornerSize,
		TopEnd:      cornerSize,
		BottomStart: cornerSize,
		BottomEnd:   cornerSize,
	}
}

func RoundedCornerShapeDpAll(size unit.Dp) RoundedCornerShape {
	return RoundedCornerShapeAll(NewDpCornerSize(size))
}

func RoundedCornerShapePxAll(size float32) RoundedCornerShape {
	return RoundedCornerShapeAll(NewPxCornerSize(size))
}

func RoundedCornerShapePercentAll(percent float32) RoundedCornerShape {
	return RoundedCornerShapeAll(NewPercentCornerSize(percent))
}

func RoundedCornerShapeDp(topStart, topEnd, bottomEnd, bottomStart unit.Dp) RoundedCornerShape {
	return RoundedCornerShape{
		TopStart:    NewDpCornerSize(topStart),
		TopEnd:      NewDpCornerSize(topEnd),
		BottomEnd:   NewDpCornerSize(bottomEnd),
		BottomStart: NewDpCornerSize(bottomStart),
	}
}

func RoundedCornerShapePx(topStart, topEnd, bottomEnd, bottomStart float32) RoundedCornerShape {
	return RoundedCornerShape{
		TopStart:    NewPxCornerSize(topStart),
		TopEnd:      NewPxCornerSize(topEnd),
		BottomEnd:   NewPxCornerSize(bottomEnd),
		BottomStart: NewPxCornerSize(bottomStart),
	}
}

/**
 * Creates [RoundedCornerShape] with sizes defined in percents of the shape's smaller side.
 *
 * @param topStartPercent The top start corner radius as a percentage of the smaller side, with a
 *   range of 0 - 100.
 * @param topEndPercent The top end corner radius as a percentage of the smaller side, with a range
 *   of 0 - 100.
 * @param bottomEndPercent The bottom end corner radius as a percentage of the smaller side, with a
 *   range of 0 - 100.
 * @param bottomStartPercent The bottom start corner radius as a percentage of the smaller side,
 *   with a range of 0 - 100.
 */
func RoundedCornerShapePercent(topStartPercent, topEndPercent, bottomEndPercent, bottomStartPercent float32) RoundedCornerShape {
	return RoundedCornerShape{
		TopStart:    NewPercentCornerSize(topStartPercent),
		TopEnd:      NewPercentCornerSize(topEndPercent),
		BottomEnd:   NewPercentCornerSize(bottomEndPercent),
		BottomStart: NewPercentCornerSize(bottomStartPercent),
	}
}
