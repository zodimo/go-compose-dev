package weight

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/internal/modifier"
)

func Weight(weight int) ui.Modifier {

	if weight < 0 {
		panic("weight cannot be negative")
	}

	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&WeightElement{
				weight: WeightData{
					Weight: float32(weight),
				},
			},
		),
		modifier.NewInspectorInfo(
			"weight",
			map[string]any{
				"weight": weight,
			},
		),
	)
}
