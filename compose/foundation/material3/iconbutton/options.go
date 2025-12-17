package iconbutton

import (
	"git.sr.ht/~schnwalter/gio-mw/widget/button"
)

type IconButtonOptions struct {
	Modifier Modifier
	Button   *button.Button
}

type IconButtonOption func(o *IconButtonOptions)

func WithModifier(m Modifier) IconButtonOption {
	return func(o *IconButtonOptions) {
		o.Modifier = m
	}
}

func DefaultIconButtonOptions() IconButtonOptions {

	return IconButtonOptions{
		Modifier: EmptyModifier,
	}
}
