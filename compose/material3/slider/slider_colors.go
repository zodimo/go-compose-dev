package slider

import (
	"gioui.org/layout"

	"github.com/zodimo/go-compose/theme"
)

// SliderColors represents the colors used by a Slider in different states.
type SliderColors struct {
	ThumbColor            theme.ColorDescriptor
	ActiveTrackColor      theme.ColorDescriptor
	ActiveTickColor       theme.ColorDescriptor
	InactiveTrackColor    theme.ColorDescriptor
	InactiveTickColor     theme.ColorDescriptor
	DisabledThumbColor    theme.ColorDescriptor
	DisabledActiveTrack   theme.ColorDescriptor
	DisabledActiveTick    theme.ColorDescriptor
	DisabledInactiveTrack theme.ColorDescriptor
	DisabledInactiveTick  theme.ColorDescriptor
}

// ThumbColor returns the color of the thumb based on the enabled state.
func (c SliderColors) Thumb(enabled bool) theme.ColorDescriptor {
	if enabled {
		return c.ThumbColor
	}
	return c.DisabledThumbColor
}

// TrackColor returns the color of the track based on the enabled and active state.
func (c SliderColors) Track(enabled, active bool) theme.ColorDescriptor {
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
func (c SliderColors) Tick(enabled, active bool) theme.ColorDescriptor {
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
