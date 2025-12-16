package slider

import (
	"gioui.org/unit"

	"github.com/zodimo/go-compose/theme"
)

// SliderDefaults holds default values for the Slider component.
var SliderDefaults = sliderDefaults{}

type sliderDefaults struct{}

// Colors returns the default SliderColors.
func (d sliderDefaults) Colors() SliderColors {
	selector := theme.ColorHelper.ColorSelector()
	return SliderColors{
		ThumbColor:            selector.PrimaryRoles.Primary,
		ActiveTrackColor:      selector.PrimaryRoles.Primary,
		ActiveTickColor:       selector.PrimaryRoles.OnPrimary.SetOpacity(0.38), // M3 Spec says Outline Variant usually, but let's stick to OnPrimary for contrast on Primary track
		InactiveTrackColor:    selector.SurfaceRoles.ContainerHighest,           // SurfaceContainerHighest
		InactiveTickColor:     selector.SurfaceRoles.OnVariant.SetOpacity(0.38),
		DisabledThumbColor:    selector.SurfaceRoles.OnSurface.SetOpacity(0.38),
		DisabledActiveTrack:   selector.SurfaceRoles.OnSurface.SetOpacity(0.38),
		DisabledActiveTick:    selector.SurfaceRoles.OnSurface.SetOpacity(0.38),
		DisabledInactiveTrack: selector.SurfaceRoles.OnSurface.SetOpacity(0.12),
		DisabledInactiveTick:  selector.SurfaceRoles.OnSurface.SetOpacity(0.12),
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
