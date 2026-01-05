package appbar

import "github.com/zodimo/go-compose/compose/ui"

// TopAppBarOptions configuration
// TopAppBarOptions configuration
type TopAppBarOptions struct {
	Modifier       ui.Modifier
	NavigationIcon Composable
	Actions        []Composable
	Colors         TopAppBarColors
}

type TopAppBarOption func(*TopAppBarOptions)

func DefaultTopAppBarOptions(c Composer) TopAppBarOptions {
	return TopAppBarOptions{
		Modifier: ui.EmptyModifier,
		Colors:   TopAppBarDefaults.Colors(c),
	}
}

func WithModifier(m ui.Modifier) TopAppBarOption {
	return func(o *TopAppBarOptions) {
		o.Modifier = m
	}
}

func WithNavigationIcon(icon Composable) TopAppBarOption {
	return func(o *TopAppBarOptions) {
		o.NavigationIcon = icon
	}
}

func WithActions(actions ...Composable) TopAppBarOption {
	return func(o *TopAppBarOptions) {
		o.Actions = actions
	}
}

func WithColors(colors TopAppBarColors) TopAppBarOption {
	return func(o *TopAppBarOptions) {
		o.Colors = colors
	}
}
