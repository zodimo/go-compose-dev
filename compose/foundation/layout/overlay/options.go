package overlay

import (
	"github.com/zodimo/go-compose/theme"
)

type OverlayOptions struct {
	Modifier   Modifier
	OnDismiss  func()
	ScrimColor theme.ColorDescriptor
}

type OverlayOption func(*OverlayOptions)

func DefaultOverlayOptions() OverlayOptions {
	return OverlayOptions{
		Modifier:   EmptyModifier,
		ScrimColor: theme.ColorHelper.ColorSelector().ScrimRoles.Scrim,
	}
}

func WithModifier(m Modifier) OverlayOption {
	return func(o *OverlayOptions) {
		o.Modifier = m
	}
}

func WithOnDismiss(f func()) OverlayOption {
	return func(o *OverlayOptions) {
		o.OnDismiss = f
	}
}

func WithScrimColor(c theme.ColorDescriptor) OverlayOption {
	return func(o *OverlayOptions) {
		o.ScrimColor = c
	}
}
