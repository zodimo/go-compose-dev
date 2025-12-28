package icon

import (
	"github.com/zodimo/go-compose/theme"
)

type IconOptions struct {
	Modifier Modifier
	Color    theme.ColorDescriptor
}

type IconOption func(*IconOptions)

func DefaultIconOptions() IconOptions {
	return IconOptions{
		Modifier: EmptyModifier,
		Color:    theme.ColorUnspecified,
	}
}

func WithModifier(m Modifier) IconOption {
	return func(o *IconOptions) {
		o.Modifier = m
	}
}

func WithColor(col theme.ColorDescriptor) IconOption {
	return func(o *IconOptions) {
		o.Color = col
	}
}
