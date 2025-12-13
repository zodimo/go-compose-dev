package navigationbar

import (
	"image/color"

	"gioui.org/unit"
)

// NavigationBarOptions configuration
type NavigationBarOptions struct {
	Modifier       Modifier
	ContainerColor color.Color
	ContentColor   color.Color
	TonalElevation unit.Dp
	Height         unit.Dp
}

// NavigationBarOption is a function that configures a NavigationBar.
type NavigationBarOption func(*NavigationBarOptions)

// DefaultNavigationBarOptions returns the default options.
func DefaultNavigationBarOptions() NavigationBarOptions {
	return NavigationBarOptions{
		Modifier:       EmptyModifier,
		ContainerColor: NavigationBarDefaults.Colors().ContainerColor,
		ContentColor:   NavigationBarDefaults.Colors().ContentColor,
		TonalElevation: NavigationBarDefaults.ContainerElevation(),
		Height:         NavigationBarDefaults.Height(),
	}
}

func WithModifier(m Modifier) NavigationBarOption {
	return func(o *NavigationBarOptions) {
		o.Modifier = m
	}
}

func WithContainerColor(c color.Color) NavigationBarOption {
	return func(o *NavigationBarOptions) {
		o.ContainerColor = c
	}
}

func WithContentColor(c color.Color) NavigationBarOption {
	return func(o *NavigationBarOptions) {
		o.ContentColor = c
	}
}

func WithTonalElevation(e unit.Dp) NavigationBarOption {
	return func(o *NavigationBarOptions) {
		o.TonalElevation = e
	}
}

func WithHeight(h unit.Dp) NavigationBarOption {
	return func(o *NavigationBarOptions) {
		o.Height = h
	}
}
