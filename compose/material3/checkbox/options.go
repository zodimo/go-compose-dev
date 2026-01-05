package checkbox

import "github.com/zodimo/go-compose/compose/ui"

type CheckboxOptions struct {
	Modifier ui.Modifier
}

type CheckboxOption func(*CheckboxOptions)

func DefaultCheckboxOptions() CheckboxOptions {
	return CheckboxOptions{
		Modifier: ui.EmptyModifier,
	}
}

func WithModifier(m ui.Modifier) CheckboxOption {
	return func(o *CheckboxOptions) {
		o.Modifier = m
	}
}
