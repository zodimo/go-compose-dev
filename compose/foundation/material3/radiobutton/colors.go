package radiobutton

import (
	"image/color"
)

// RadioButtonColors represents the colors used by a RadioButton in different states.
type RadioButtonColors struct {
	SelectedColor   color.Color
	UnselectedColor color.Color
	DisabledColor   color.Color
}

// Defaults contains the default values for RadioButton.
var Defaults = radioButtonDefaults{}

type radioButtonDefaults struct{}

// Colors returns the default colors for a RadioButton.
// The nil values allow the component to resolve theme colors from the context.
func (d radioButtonDefaults) Colors() RadioButtonColors {
	return RadioButtonColors{
		// Leaving fields as nil/zero-value to imply "use theme defaults"
	}
}

// Composite function to resolve color based on state
func (c RadioButtonColors) Color(enabled, selected bool) color.Color {
	if !enabled {
		return c.DisabledColor
	}
	if selected {
		return c.SelectedColor
	}
	return c.UnselectedColor
}
