package surface

import (
	"go-compose-dev/compose/ui/graphics/shape"
	"go-compose-dev/internal/modifier"
	"image/color"
)

type SurfaceOptions struct {
	Modifier        Modifier
	Shape           Shape
	Color           color.Color
	ContentColor    color.Color
	TonalElevation  Dp
	ShadowElevation Dp
	Border          Dp
}

type SurfaceOption func(*SurfaceOptions)

func DefaultSurfaceOptions() SurfaceOptions {
	return SurfaceOptions{
		Modifier:        modifier.EmptyModifier,
		Shape:           shape.ShapeRectangle,
		Color:           color.NRGBA{R: 255, G: 255, B: 255, A: 255}, // Default white? Or theme?
		ContentColor:    color.NRGBA{A: 255},                         // Black?
		TonalElevation:  0,
		ShadowElevation: 0,
		Border:          0,
	}
}

func WithModifier(m Modifier) SurfaceOption {
	return func(o *SurfaceOptions) {
		o.Modifier = m
	}
}

func WithShape(s Shape) SurfaceOption {
	return func(o *SurfaceOptions) {
		o.Shape = s
	}
}

func WithColor(c color.Color) SurfaceOption {
	return func(o *SurfaceOptions) {
		o.Color = c
	}
}

func WithContentColor(c color.Color) SurfaceOption {
	return func(o *SurfaceOptions) {
		o.ContentColor = c
	}
}

func WithTonalElevation(elevation Dp) SurfaceOption {
	return func(o *SurfaceOptions) {
		o.TonalElevation = elevation
	}
}

func WithShadowElevation(elevation Dp) SurfaceOption {
	return func(o *SurfaceOptions) {
		o.ShadowElevation = elevation
	}
}

func WithBorder(border Dp) SurfaceOption {
	return func(o *SurfaceOptions) {
		o.Border = border
	}
}
