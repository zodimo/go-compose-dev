package divider

import (
	"github.com/zodimo/go-compose/theme"
)

type DividerOptions struct {
	Modifier  Modifier
	Thickness float64
	Color     theme.ColorDescriptor
}

type DividerOption func(*DividerOptions)

func WithModifier(modifier Modifier) DividerOption {
	return func(o *DividerOptions) {
		o.Modifier = o.Modifier.Then(modifier)
	}
}

func WithThickness(value int) DividerOption {
	return func(o *DividerOptions) {
		o.Thickness = float64(value)
	}
}

func WithColor(value theme.ColorDescriptor) DividerOption {
	return func(o *DividerOptions) {
		o.Color = value
	}
}
