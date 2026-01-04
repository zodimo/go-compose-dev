package tokens

import (
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

// Version: androidx 1.4.1 (ShapeTokens.kt)
var ShapeTokens = ShapeTokensData{
	CornerExtraExtraLarge:     &shape.RoundedCornerShape{Radius: unit.Dp(48)},
	CornerExtraLarge:          &shape.RoundedCornerShape{Radius: unit.Dp(28)},
	CornerExtraLargeIncreased: &shape.RoundedCornerShape{Radius: unit.Dp(32)},
	CornerExtraLargeTop: &shape.RoundedCornerShape{
		TopStart:    unit.Dp(28),
		TopEnd:      unit.Dp(28),
		BottomEnd:   unit.Dp(0),
		BottomStart: unit.Dp(0),
	},
	CornerExtraSmall: &shape.RoundedCornerShape{Radius: unit.Dp(4)},
	CornerExtraSmallTop: &shape.RoundedCornerShape{
		TopStart:    unit.Dp(4),
		TopEnd:      unit.Dp(4),
		BottomEnd:   unit.Dp(0),
		BottomStart: unit.Dp(0),
	},
	CornerFull:           shape.CircleShape,
	CornerLarge:          &shape.RoundedCornerShape{Radius: unit.Dp(16)},
	CornerLargeEnd:       &shape.RoundedCornerShape{TopEnd: unit.Dp(16), BottomEnd: unit.Dp(16), TopStart: unit.Dp(0), BottomStart: unit.Dp(0)},
	CornerLargeIncreased: &shape.RoundedCornerShape{Radius: unit.Dp(20)},
	CornerLargeStart:     &shape.RoundedCornerShape{TopStart: unit.Dp(16), BottomStart: unit.Dp(16), TopEnd: unit.Dp(0), BottomEnd: unit.Dp(0)},
	CornerLargeTop: &shape.RoundedCornerShape{
		TopStart:    unit.Dp(16),
		TopEnd:      unit.Dp(16),
		BottomEnd:   unit.Dp(0),
		BottomStart: unit.Dp(0),
	},
	CornerMedium:                   &shape.RoundedCornerShape{Radius: unit.Dp(12)},
	CornerNone:                     shape.ShapeRectangle,
	CornerSmall:                    &shape.RoundedCornerShape{Radius: unit.Dp(8)},
	CornerValueExtraExtraLarge:     unit.Dp(48),
	CornerValueExtraLarge:          unit.Dp(28),
	CornerValueExtraLargeIncreased: unit.Dp(32),
	CornerValueExtraSmall:          unit.Dp(4),
	CornerValueLarge:               unit.Dp(16),
	CornerValueLargeIncreased:      unit.Dp(20),
	CornerValueMedium:              unit.Dp(12),
	CornerValueNone:                unit.Dp(0),
	CornerValueSmall:               unit.Dp(8),
}

type ShapeTokensData struct {
	CornerExtraExtraLarge          shape.Shape
	CornerExtraLarge               shape.Shape
	CornerExtraLargeIncreased      shape.Shape
	CornerExtraLargeTop            shape.Shape
	CornerExtraSmall               shape.Shape
	CornerExtraSmallTop            shape.Shape
	CornerFull                     shape.Shape
	CornerLarge                    shape.Shape
	CornerLargeEnd                 shape.Shape
	CornerLargeIncreased           shape.Shape
	CornerLargeStart               shape.Shape
	CornerLargeTop                 shape.Shape
	CornerMedium                   shape.Shape
	CornerNone                     shape.Shape
	CornerSmall                    shape.Shape
	CornerValueExtraExtraLarge     unit.Dp
	CornerValueExtraLarge          unit.Dp
	CornerValueExtraLargeIncreased unit.Dp
	CornerValueExtraSmall          unit.Dp
	CornerValueLarge               unit.Dp
	CornerValueLargeIncreased      unit.Dp
	CornerValueMedium              unit.Dp
	CornerValueNone                unit.Dp
	CornerValueSmall               unit.Dp
}
