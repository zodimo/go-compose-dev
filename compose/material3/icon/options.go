package icon

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

type IconOptions struct {
	Modifier ui.Modifier
	Color    graphics.Color
	FontSize unit.TextUnit
}

type IconOption func(*IconOptions)

func DefaultIconOptions() IconOptions {
	return IconOptions{
		Modifier: ui.EmptyModifier,
		Color:    graphics.ColorUnspecified,
		FontSize: unit.TextUnitUnspecified,
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

func WithSymbolSize(size unit.TextUnit) IconOption {
	return func(o *IconOptions) {
		o.FontSize = size
	}
}
