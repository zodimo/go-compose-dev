package navigationbar

import (
	"image/color"

	"gioui.org/unit"
)

// NavigationBarColors represents the colors used by a NavigationBar.
type NavigationBarColors struct {
	ContainerColor color.Color
	ContentColor   color.Color
	IndicatorColor color.Color
}

// NavigationBarDefaults holds the default values for NavigationBar.
var NavigationBarDefaults = navigationBarDefaults{}

type navigationBarDefaults struct{}

// Colors returns the default colors for a NavigationBar.
func (d navigationBarDefaults) Colors() NavigationBarColors {
	return NavigationBarColors{
		// Surface Container: R: 242, G: 237, B: 246 (Light Theme approximation)
		// Ideally this should come from the theme, but for defaults we often hardcode or provide a way to resolve against theme.
		// Matching BottomAppBar's approach of hardcoded defaults for now.
		ContainerColor: color.NRGBA{R: 242, G: 237, B: 246, A: 255},
		// On Surface Variant: R: 73, G: 69, B: 79
		ContentColor: color.NRGBA{R: 73, G: 69, B: 79, A: 255},
		// Secondary Container: R: 232, G: 222, B: 248
		IndicatorColor: color.NRGBA{R: 232, G: 222, B: 248, A: 255},
	}
}

// ContainerElevation returns the default elevation for a NavigationBar.
func (d navigationBarDefaults) ContainerElevation() unit.Dp {
	return unit.Dp(3) // Level 2: 3.0dp
}

// Height returns the default height for a NavigationBar.
func (d navigationBarDefaults) Height() unit.Dp {
	return unit.Dp(80)
}
