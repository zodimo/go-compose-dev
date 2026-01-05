package icon

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/compose/ui/graphics"
)

type IconOptions struct {
	Modifier ui.Modifier
	Color    graphics.Color
}

type IconOption func(*IconOptions)

func DefaultIconOptions() IconOptions {
	return IconOptions{
		Modifier: ui.EmptyModifier,
		Color:    graphics.ColorUnspecified,
	}
}

func WithModifier(m ui.Modifier) IconOption {
	return func(o *IconOptions) {
		o.Modifier = m
	}
}

func WithColor(col graphics.Color) IconOption {
	return func(o *IconOptions) {
		o.Color = col
	}
}
