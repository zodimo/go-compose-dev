package alpha

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/internal/modifier"
)

func Alpha(alpha float32) ui.Modifier {
	if alpha < 0 || alpha > 1 {
		panic("alpha must be between 0 and 1")
	}
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&AlphaElement{
				Alpha: alpha,
			},
		),
		modifier.NewInspectorInfo(
			"alpha",
			map[string]any{
				"alpha": alpha,
			},
		),
	)
}
