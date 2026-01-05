package navigationbar

import "github.com/zodimo/go-compose/compose/ui"

type NavigationBarItemOptions struct {
	Modifier ui.Modifier
}

type NavigationBarItemOption func(*NavigationBarItemOptions)

func ItemWithModifier(m ui.Modifier) NavigationBarItemOption {
	return func(o *NavigationBarItemOptions) {
		o.Modifier = m
	}
}
