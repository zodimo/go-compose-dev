package background

import (
	"github.com/zodimo/go-compose/internal/modifier"
	"image/color"
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

func Background(color color.Color, options ...BackgroundOption) Modifier {

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
					Color: color,
					Shape: opt.Shape,
				},
			},
		),
		modifier.NewInspectorInfo(
			"background",
			map[string]any{
				"color":   color,
				"options": opt,
			},
		),
	)
}
