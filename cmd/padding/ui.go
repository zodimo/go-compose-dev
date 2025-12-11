package main

import (
	"go-compose-dev/compose"
	"go-compose-dev/compose/foundation/layout/column"

	"go-compose-dev/compose/foundation/text"
	"go-compose-dev/internal/modifiers/background"
	"go-compose-dev/internal/modifiers/padding"
	"go-compose-dev/pkg/api"
	"image/color"
)

func UI(c api.Composer) api.LayoutNode {

	c = column.Column(
		compose.Sequence(
			text.Text("hello world",
				text.WithModifier(background.Background(color.NRGBA{R: 255, G: 0, B: 0, A: 150})),
				text.WithModifier(padding.All(20)),
				text.WithModifier(background.Background(color.NRGBA{R: 0, G: 255, B: 0, A: 255})),
			),
		),
		column.WithModifier(background.Background(color.NRGBA{R: 0, G: 0, B: 200, A: 255})),

		column.WithAlignment(column.Middle),
	)(c)

	return c.Build()

}
