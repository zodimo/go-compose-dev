package iconbutton

import (
	"git.sr.ht/~schnwalter/gio-mw/widget/button"
	"github.com/zodimo/go-compose/compose/ui"
)

type IconButtonOptions struct {
	Modifier ui.Modifier
	Button   *button.Button
}

type IconButtonOption func(o *IconButtonOptions)

func WithModifier(m ui.Modifier) IconButtonOption {
	return func(o *IconButtonOptions) {
		o.Modifier = m
	}
}

func DefaultIconButtonOptions() IconButtonOptions {

	return IconButtonOptions{
		Modifier: ui.EmptyModifier,
	}
}
