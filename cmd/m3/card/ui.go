package main

import (
	"fmt"
	"go-compose-dev/compose"
	"go-compose-dev/compose/foundation/layout/column"
	"go-compose-dev/compose/foundation/material3/button"
	"go-compose-dev/compose/foundation/material3/card"
	"go-compose-dev/compose/foundation/material3/divider"
	"go-compose-dev/compose/foundation/text"
	"go-compose-dev/internal/modifiers/padding"
	"go-compose-dev/pkg/api"
)

func UI(c api.Composer) api.LayoutNode {

	c = column.Column(compose.Sequence(
		button.FilledTonal(
			func() {
				fmt.Println("M3 Filled Tonal Button clicked!")
			},
			"Hello M3 FilledTonal Button",
			button.WithModifier(padding.All(20)),
		),
		divider.Divider(),
		card.Filled(card.CardContents(
			card.Content(text.Text("Filled")),
		)),
		card.Elevated(card.CardContents(
			card.Content(text.Text("Elevated")),
		)),
		card.Outlined(card.CardContents(
			card.Content(text.Text("Outlined")),
		)),
	),
	)(c)

	return c.Build()
}
