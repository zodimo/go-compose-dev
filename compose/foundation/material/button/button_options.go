package button

import (
	"gioui.org/widget/material"
)

type ButtonOptions struct {
	Modifier  Modifier
	Theme     *material.Theme
	Clickable *GioClickable
}

type ButtonOption func(o *ButtonOptions)

func WithTheme(theme *material.Theme) ButtonOption {
	return func(o *ButtonOptions) {
		o.Theme = theme
	}
}

func WithModifier(m Modifier) ButtonOption {
	return func(o *ButtonOptions) {
		o.Modifier = m
	}
}

func WithClickable(clickable *GioClickable) ButtonOption {
	return func(o *ButtonOptions) {
		o.Clickable = clickable
	}
}
