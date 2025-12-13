package bottomappbar

import (
	"image/color"

	"gioui.org/unit"
)

// BottomAppBarColors represents the colors used by a BottomAppBar.
type BottomAppBarColors struct {
	ContainerColor color.Color
	ContentColor   color.Color
}

// BottomAppBarDefaults holds the default values for BottomAppBar.
var BottomAppBarDefaults = bottomAppBarDefaults{}

type bottomAppBarDefaults struct{}

// Colors returns the default colors for a BottomAppBar.
func (d bottomAppBarDefaults) Colors() BottomAppBarColors {
	return BottomAppBarColors{
		// Surface Container: R: 241, G: 244, B: 249
		ContainerColor: color.NRGBA{R: 241, G: 244, B: 249, A: 255},
		// On Surface Variant: R: 67, G: 71, B: 78
		ContentColor: color.NRGBA{R: 67, G: 71, B: 78, A: 255},
	}
}

// ContainerElevation returns the default elevation for a BottomAppBar.
func (d bottomAppBarDefaults) ContainerElevation() unit.Dp {
	return unit.Dp(3) // Level 2: 3.0dp
}

// ContentPadding returns the default content padding for a BottomAppBar.
// Standard is Horizontal 16dp, Vertical 12dp (for tall bar 80dp) or 16dp (for short 64dp).
// Let's use internal values in implementation, but expose a default here if we define PaddingValues.
// For now, I'll return specific values.
func (d bottomAppBarDefaults) ContentPadding() (start, top, end, bottom unit.Dp) {
	return unit.Dp(16), unit.Dp(12), unit.Dp(16), unit.Dp(12)
}
