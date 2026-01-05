package button

import (
	"git.sr.ht/~schnwalter/gio-mw/widget/button"
	"github.com/zodimo/go-compose/compose/foundation"
	"github.com/zodimo/go-compose/compose/foundation/layout"
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
)

type ButtonOptions struct {
	Modifier       ui.Modifier
	Button         *button.Button
	Enabled        bool
	Shape          shape.Shape   // Single shape (legacy/simple API)
	Shapes         *ButtonShapes // Shape and PressedShape for expressive shape morphing
	Colors         *ButtonColors
	Elevation      *ButtonElevation
	Border         *foundation.BorderStroke
	ContentPadding layout.PaddingValues
}

type ButtonOption func(o *ButtonOptions)

func WithModifier(m ui.Modifier) ButtonOption {
	return func(o *ButtonOptions) {
		o.Modifier = m
	}
}

func WithButton(button *button.Button) ButtonOption {
	return func(o *ButtonOptions) {
		o.Button = button
	}
}

func WithEnabled(enabled bool) ButtonOption {
	return func(o *ButtonOptions) {
		o.Enabled = enabled
	}
}

func WithShape(s shape.Shape) ButtonOption {
	return func(o *ButtonOptions) {
		o.Shape = s
	}
}

func WithShapes(s *ButtonShapes) ButtonOption {
	return func(o *ButtonOptions) {
		o.Shapes = s
	}
}

func WithColors(c *ButtonColors) ButtonOption {
	return func(o *ButtonOptions) {
		o.Colors = c
	}
}

func WithElevation(e *ButtonElevation) ButtonOption {
	return func(o *ButtonOptions) {
		o.Elevation = e
	}
}

func WithBorder(b *foundation.BorderStroke) ButtonOption {
	return func(o *ButtonOptions) {
		o.Border = b
	}
}

func WithContentPadding(p layout.PaddingValues) ButtonOption {
	return func(o *ButtonOptions) {
		o.ContentPadding = p
	}
}
