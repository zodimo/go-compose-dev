package image

import (
	"github.com/zodimo/go-compose/compose/ui"
	uilayout "github.com/zodimo/go-compose/compose/ui/layout"
	"github.com/zodimo/go-compose/internal/modifier"
	"github.com/zodimo/go-compose/modifiers/size"
)

type ImageOptions struct {
	Modifier     ui.Modifier
	Alignment    size.Alignment
	ContentScale uilayout.ContentScale
	Alpha        float32
	// colorFilter is deferred
}

func DefaultImageOptions() ImageOptions {
	return ImageOptions{
		Modifier:     modifier.EmptyModifier,
		Alignment:    size.Center,
		ContentScale: uilayout.ContentScaleFit,
		Alpha:        1.0,
	}
}

type ImageOption func(*ImageOptions)

func WithModifier(m ui.Modifier) ImageOption {
	return func(o *ImageOptions) {
		o.Modifier = m
	}
}

func WithAlignment(alignment size.Alignment) ImageOption {
	return func(o *ImageOptions) {
		o.Alignment = alignment
	}
}

func WithContentScale(contentScale uilayout.ContentScale) ImageOption {
	return func(o *ImageOptions) {
		o.ContentScale = contentScale
	}
}

func WithAlpha(alpha float32) ImageOption {
	return func(o *ImageOptions) {
		o.Alpha = alpha
	}
}
