package card

import "github.com/zodimo/go-compose/compose/ui"

func DefaultCardOptions() CardOptions {
	return CardOptions{
		Modifier: ui.EmptyModifier,
	}
}
