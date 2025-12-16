package radiobutton

import (
	"github.com/zodimo/go-compose/theme"

	"git.sr.ht/~schnwalter/gio-mw/token"
)

// RadioButtonColors represents the colors used by a RadioButton in different states.
type RadioButtonColors struct {
	SelectedColor   theme.ColorDescriptor
	UnselectedColor theme.ColorDescriptor
	DisabledColor   theme.ColorDescriptor
}

// Defaults contains the default values for RadioButton.
var Defaults = radioButtonDefaults{}

type radioButtonDefaults struct{}

// Colors returns the default colors for a RadioButton.
func (d radioButtonDefaults) Colors() RadioButtonColors {
	return RadioButtonColors{
		SelectedColor:   theme.ColorHelper.ColorSelector().PrimaryRoles.Primary,
		UnselectedColor: theme.ColorHelper.ColorSelector().SurfaceRoles.OnVariant,
		DisabledColor:   theme.ColorHelper.ColorSelector().SurfaceRoles.OnSurface.SetOpacity(token.OpacityLevel5),
	}
}

// Color returns the resolved color based on state (used internally for drawing).
func (c RadioButtonColors) Color(enabled, selected bool) theme.ColorDescriptor {
	if !enabled {
		return c.DisabledColor
	}
	if selected {
		return c.SelectedColor
	}
	return c.UnselectedColor
}
