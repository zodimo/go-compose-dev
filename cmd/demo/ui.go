package main

import (
	"go-compose-dev/compose"
	"go-compose-dev/compose/foundation/layout/column"
	"go-compose-dev/internal/modifiers/background"
	"go-compose-dev/internal/modifiers/size"
	"go-compose-dev/pkg/api"
	"image/color"
)

func UI(c api.Composer) api.LayoutNode {

	c = column.Column(
		compose.Sequence(),
		column.WithModifier(size.Size(100, 100)),
		column.WithModifier(background.Background(color.NRGBA{R: 255, G: 0, B: 0, A: 255})),
	)(c)

	return c.Build()

}
