package main

import (
	"go-compose-dev/compose/foundation/layout/column"
	"go-compose-dev/compose/foundation/layout/spacer"
	"go-compose-dev/compose/foundation/material3/bottomsheet"
	"go-compose-dev/compose/foundation/material3/button"
	"go-compose-dev/compose/foundation/material3/scaffold"
	"go-compose-dev/compose/foundation/material3/text"
	"go-compose-dev/internal/modifier"
	"go-compose-dev/internal/modifiers/padding"
	"go-compose-dev/pkg/api"
)

func UI(c api.Composer) api.LayoutNode {
	// State for managing sheet visibility
	showSheet := c.State("showSheet", func() any { return false })
	isOpen := showSheet.Get().(bool)

	return bottomsheet.ModalBottomSheet(
		// Sheet Content
		func(c api.Composer) api.Composer {
			return column.Column(
				func(c api.Composer) api.Composer {
					text.Text("Bottom Sheet Title", text.TypestyleTitleLarge)(c)
					spacer.SpacerHeight(16)(c)
					text.Text("This is the content of the bottom sheet. It slides up from the bottom.", text.TypestyleBodyMedium)(c)
					spacer.SpacerHeight(16)(c)
					button.Filled(
						func() {
							showSheet.Set(false)
						},
						"Close Sheet",
					)(c)
					return c
				},
				column.WithModifier(modifier.EmptyModifier.Then(padding.All(24))),
			)(c)
		},
		// Screen Content
		func(c api.Composer) api.Composer {
			return scaffold.Scaffold(
				// Content
				func(c api.Composer) api.Composer {
					return column.Column(
						func(c api.Composer) api.Composer {
							text.Text("Main Screen Content", text.TypestyleBodyMedium)(c)
							spacer.SpacerHeight(16)(c)
							button.Filled(
								func() {
									showSheet.Set(true)
								},
								"Show Bottom Sheet",
							)(c)
							return c
						},
						column.WithModifier(modifier.EmptyModifier.Then(padding.All(24))),
					)(c)
				},
			)(c)
		},
		bottomsheet.WithIsOpen(isOpen),
		bottomsheet.WithOnDismissRequest(func() {
			showSheet.Set(false)
		}),
	)(c).Build()
}
