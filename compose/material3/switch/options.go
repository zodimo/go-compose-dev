package mswitch

import "github.com/zodimo/go-compose/compose/ui"

type SwitchOptions struct {
	Modifier ui.Modifier
}

type SwitchOption func(*SwitchOptions)

func DefaultSwitchOptions() SwitchOptions {
	return SwitchOptions{
		Modifier: ui.EmptyModifier,
	}
}

func WithModifier(m ui.Modifier) SwitchOption {
	return func(o *SwitchOptions) {
		o.Modifier = m
	}
}
