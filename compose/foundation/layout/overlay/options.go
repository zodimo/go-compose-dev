package overlay

import (
	"image/color"
)

type OverlayOptions struct {
	Modifier   Modifier
	OnDismiss  func()
	ScrimColor color.NRGBA
}

type OverlayOption func(*OverlayOptions)

func DefaultOverlayOptions() OverlayOptions {
	return OverlayOptions{
		Modifier:   EmptyModifier,
		ScrimColor: color.NRGBA{R: 0, G: 0, B: 0, A: 128}, // Semi-transparent black
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

func WithScrimColor(c color.NRGBA) OverlayOption {
	return func(o *OverlayOptions) {
		o.ScrimColor = c
	}
}
