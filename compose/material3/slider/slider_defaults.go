package slider

import (
	gioUnit "gioui.org/unit"
	"github.com/zodimo/go-compose/compose/ui/graphics"
)

// SliderDefaults holds default values for the Slider component.
var SliderDefaults = sliderDefaults{}

type sliderDefaults struct{}

// Colors returns the default SliderColors.
func (d sliderDefaults) Colors() SliderColors {
	return SliderColors{
		ThumbColor:            graphics.ColorUnspecified,
		ActiveTrackColor:      graphics.ColorUnspecified,
		ActiveTickColor:       graphics.ColorUnspecified,
		InactiveTrackColor:    graphics.ColorUnspecified,
		InactiveTickColor:     graphics.ColorUnspecified,
		DisabledThumbColor:    graphics.ColorUnspecified,
		DisabledActiveTrack:   graphics.ColorUnspecified,
		DisabledActiveTick:    graphics.ColorUnspecified,
		DisabledInactiveTrack: graphics.ColorUnspecified,
		DisabledInactiveTick:  graphics.ColorUnspecified,
	}
}

// Dimensions constants
// @TODO this should be compose.ui.unit
var (
	TrackHeight     = gioUnit.Dp(4)
	ThumbSize       = gioUnit.Dp(20)
	ActiveThumbSize = gioUnit.Dp(28) // M3 State Layer/Enlarged handle
	TickSize        = gioUnit.Dp(2)
	ThumbTrackGap   = gioUnit.Dp(6) // Approximate
)
