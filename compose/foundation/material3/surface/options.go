package surface

import (
	"image/color"

	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
	"github.com/zodimo/go-compose/internal/modifier"
)

type SurfaceOptions struct {
	Modifier        Modifier
	Shape           Shape
	Color           color.Color
	ContentColor    color.Color
	TonalElevation  Dp
	ShadowElevation Dp
	BorderWidth     Dp
	BorderColor     color.Color
	Alignment       box.Direction // Optional alignment for content inside surface
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
		BorderWidth:     0,
		BorderColor:     color.NRGBA{A: 0},
		Alignment:       box.NW, // Default to NW to match Box default, or Center? Box defaults to NW.
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

func WithBorder(width Dp, color color.Color) SurfaceOption {
	return func(o *SurfaceOptions) {
		o.BorderWidth = width
		o.BorderColor = color
	}
}

func WithAlignment(alignment box.Direction) SurfaceOption {
	return func(o *SurfaceOptions) {
		o.Alignment = alignment
	}
}
