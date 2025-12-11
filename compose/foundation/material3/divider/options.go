package divider

import (
	"image/color"
)

type DividerOptions struct {
	Modifier  Modifier
	Thickness float64
	Color     color.Color
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

func WithColor(value color.Color) DividerOption {
	return func(o *DividerOptions) {
		o.Color = value
	}
}
