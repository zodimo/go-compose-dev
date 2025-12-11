package button

import (
	"go-compose-dev/compose/foundation/material"
)

func DefaultButtonOptions() ButtonOptions {
	return ButtonOptions{
		Modifier: EmptyModifier,
		Theme:    material.GetTheme(),
	}
}
