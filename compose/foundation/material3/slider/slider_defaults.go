package slider

import (
	"image/color"

	"gioui.org/unit"
)

// SliderDefaults holds default values for the Slider component.
var SliderDefaults = sliderDefaults{}

type sliderDefaults struct{}

// Colors returns the default SliderColors.
// TODO: Hook this up to the actual MaterialTheme when available globally or pass it in.
// For now we use hardcoded Material 3 baseline colors for reference, expecting them to be
// retrieved from the theme in the actual composable if not provided.
func (d sliderDefaults) Colors() SliderColors {
	return SliderColors{
		ThumbColor:            color.NRGBA{R: 103, G: 80, B: 164, A: 255},  // Primary
		ActiveTrackColor:      color.NRGBA{R: 103, G: 80, B: 164, A: 255},  // Primary
		ActiveTickColor:       color.NRGBA{R: 208, G: 188, B: 255, A: 255}, // Inverse Surface / On Primary
		InactiveTrackColor:    color.NRGBA{R: 231, G: 224, B: 236, A: 255}, // Surface Variant
		InactiveTickColor:     color.NRGBA{R: 73, G: 69, B: 79, A: 255},    // On Surface Variant
		DisabledThumbColor:    color.NRGBA{R: 28, G: 27, B: 31, A: 97},     // On Surface (38%)
		DisabledActiveTrack:   color.NRGBA{R: 28, G: 27, B: 31, A: 97},     // On Surface (38%)
		DisabledActiveTick:    color.NRGBA{R: 28, G: 27, B: 31, A: 97},     // On Surface (38%) // Actually usually clearer
		DisabledInactiveTrack: color.NRGBA{R: 28, G: 27, B: 31, A: 31},     // On Surface (12%)
		DisabledInactiveTick:  color.NRGBA{R: 28, G: 27, B: 31, A: 31},     // On Surface (12%)
	}
}

// Dimensions constants
var (
	TrackHeight     = unit.Dp(4)
	ThumbSize       = unit.Dp(20)
	ActiveThumbSize = unit.Dp(28) // M3 State Layer/Enlarged handle
	TickSize        = unit.Dp(2)
	ThumbTrackGap   = unit.Dp(6) // Approximate
)
