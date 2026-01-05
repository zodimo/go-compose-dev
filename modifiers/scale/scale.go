package scale

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/internal/modifier"
)

type ScaleData struct {
	ScaleX float32
	ScaleY float32
}

func Scale(scale float32) ui.Modifier {
	return modifier.NewModifier(
		&ScaleElement{
			data: ScaleData{
				ScaleX: scale,
				ScaleY: scale,
			},
		},
	)
}

func ScaleXY(x, y float32) ui.Modifier {
	return modifier.NewModifier(
		&ScaleElement{
			data: ScaleData{
				ScaleX: x,
				ScaleY: y,
			},
		},
	)
}
