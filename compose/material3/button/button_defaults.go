package button

import "github.com/zodimo/go-compose/compose/ui"

func DefaultButtonOptions() ButtonOptions {
	return ButtonOptions{
		Modifier: ui.EmptyModifier,
		Enabled:  true,
	}
}
