package radiobutton

import "github.com/zodimo/go-compose/compose/ui"

type RadioButtonOptions struct {
	Modifier ui.Modifier
	Enabled  bool
	Colors   RadioButtonColors
}

type RadioButtonOption func(*RadioButtonOptions)

func DefaultRadioButtonOptions(c Composer) RadioButtonOptions {
	return RadioButtonOptions{
		Modifier: ui.EmptyModifier,
		Enabled:  true,
		Colors:   Defaults.Colors(c), // Use nil/defaults
	}
}

func WithModifier(m ui.Modifier) RadioButtonOption {
	return func(o *RadioButtonOptions) {
		o.Modifier = m
	}
}

func WithEnabled(enabled bool) RadioButtonOption {
	return func(o *RadioButtonOptions) {
		o.Enabled = enabled
	}
}

func WithColors(colors RadioButtonColors) RadioButtonOption {
	return func(o *RadioButtonOptions) {
		o.Colors = colors
	}
}
