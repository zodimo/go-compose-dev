package slider

import (
	"image/color"

	"gioui.org/layout"
)

// SliderColors represents the colors used by a Slider in different states.
type SliderColors struct {
	ThumbColor            color.NRGBA
	ActiveTrackColor      color.NRGBA
	ActiveTickColor       color.NRGBA
	InactiveTrackColor    color.NRGBA
	InactiveTickColor     color.NRGBA
	DisabledThumbColor    color.NRGBA
	DisabledActiveTrack   color.NRGBA
	DisabledActiveTick    color.NRGBA
	DisabledInactiveTrack color.NRGBA
	DisabledInactiveTick  color.NRGBA
}

// ThumbColor returns the color of the thumb based on the enabled state.
func (c SliderColors) Thumb(enabled bool) color.NRGBA {
	if enabled {
		return c.ThumbColor
	}
	return c.DisabledThumbColor
}

// TrackColor returns the color of the track based on the enabled and active state.
func (c SliderColors) Track(enabled, active bool) color.NRGBA {
	if enabled {
		if active {
			return c.ActiveTrackColor
		}
		return c.InactiveTrackColor
	}
	if active {
		return c.DisabledActiveTrack
	}
	return c.DisabledInactiveTrack
}

// TickColor returns the color of the ticks based on the enabled and active state.
func (c SliderColors) Tick(enabled, active bool) color.NRGBA {
	if enabled {
		if active {
			return c.ActiveTickColor
		}
		return c.InactiveTickColor
	}
	if active {
		return c.DisabledActiveTick
	}
	return c.DisabledInactiveTick
}

func (c SliderColors) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Dimensions{}
}
