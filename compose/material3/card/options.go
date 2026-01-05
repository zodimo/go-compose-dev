package card

import "github.com/zodimo/go-compose/compose/ui"

type CardOptions struct {
	Modifier ui.Modifier
}

type CardOption func(o *CardOptions)

func WithModifier(m ui.Modifier) CardOption {
	return func(o *CardOptions) {
		o.Modifier = m
	}
}
