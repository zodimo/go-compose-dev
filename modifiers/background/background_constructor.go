package background

import (
	"github.com/zodimo/go-compose/theme"

	"github.com/zodimo/go-compose/internal/modifier"
)

type BackgroundOptions struct {
	Shape Shape
}

func DefaultBackgroundOptions() BackgroundOptions {
	return BackgroundOptions{
		Shape: ShapeRectangle,
	}
}

type BackgroundOption func(options *BackgroundOptions)

// WithShape sets the shape for the background (e.g., rounded corners)
func WithShape(shape Shape) BackgroundOption {
	return func(options *BackgroundOptions) {
		options.Shape = shape
	}
}

func Background(colorDesc theme.ColorDescriptor, options ...BackgroundOption) Modifier {

	opt := DefaultBackgroundOptions()
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&opt)
	}
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&BackgroundElement{
				background: BackgroundData{
					Color: colorDesc,
					Shape: opt.Shape,
				},
			},
		),
		modifier.NewInspectorInfo(
			"background",
			map[string]any{
				"color":   colorDesc,
				"options": opt,
			},
		),
	)
}
