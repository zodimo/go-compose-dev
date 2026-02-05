package clip

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
	"github.com/zodimo/go-compose/internal/modifier"
)

func Clip(shape shape.Shape) ui.Modifier {

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&ClipElement{
				clipData: ClipData{
					Shape: shape,
				},
			},
		),
		modifier.NewInspectorInfo(
			"clip",
			map[string]any{
				"shape": shape,
			},
		),
	)
}

func ClipToBounds() ui.Modifier {

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&ClipElement{
				clipData: ClipData{
					Shape:        ShapeRectangle,
					ClipToBounds: true,
				},
			},
		),
		modifier.NewInspectorInfo(
			"clip",
			map[string]any{
				"shape":        ShapeRectangle,
				"clipToBounds": true,
			},
		),
	)
}
