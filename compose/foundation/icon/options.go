package icon

import (
	"go-compose-dev/internal/theme"
	"image/color"

	"github.com/zodimo/go-maybe"
)

type ThemeColorSet struct {
	ThemeColor func(themeColor theme.ThemeColor) color.Color
}

type IconOptions struct {
	Modifier   Modifier
	ThemeColor maybe.Maybe[ThemeColorSet]
	Color      color.Color
}

type IconOption func(*IconOptions)

func DefaultIconOptions() IconOptions {
	return IconOptions{
		Modifier:   EmptyModifier,
		ThemeColor: maybe.None[ThemeColorSet](),
		// Default Fallback is black
		Color: color.NRGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		},
	}
}

func WithModifier(m Modifier) IconOption {
	return func(o *IconOptions) {
		o.Modifier = m
	}
}

func WithThemeColor(reader func(themeColor theme.ThemeColor) color.Color) IconOption {
	return func(o *IconOptions) {
		o.ThemeColor = maybe.Some(ThemeColorSet{
			ThemeColor: reader,
		})
	}
}

func WithColor(color color.Color) IconOption {
	return func(o *IconOptions) {
		o.Color = color
	}
}
