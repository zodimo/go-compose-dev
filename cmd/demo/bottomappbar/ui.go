package main

import (
	"image/color"

	"gioui.org/layout"

	"go-compose-dev/compose/foundation/layout/box"
	"go-compose-dev/compose/foundation/layout/column"
	"go-compose-dev/compose/foundation/layout/row"
	"go-compose-dev/compose/foundation/material3/bottomappbar"
	"go-compose-dev/compose/foundation/material3/iconbutton"
	m3text "go-compose-dev/compose/foundation/material3/text"
	"go-compose-dev/compose/foundation/text"
	"go-compose-dev/pkg/api" // Import api for Composable/Composer

	"go-compose-dev/internal/modifiers/size"
	"go-compose-dev/internal/modifiers/weight"

	mdicons "golang.org/x/exp/shiny/materialdesign/icons"
)

func UI() api.Composable {
	return func(c api.Composer) api.Composer {
		return column.Column(
			func(c api.Composer) api.Composer {
				// Content Area
				box.Box(
					func(c api.Composer) api.Composer {
						return m3text.Text(
							"Content Area",
							m3text.TypestyleBodyLarge, // valid style
							text.WithTextStyleOptions(
								text.StyleWithColor(color.NRGBA{A: 255}),
							),
						)(c)
					},
					box.WithModifier(weight.Weight(1)),
					box.WithAlignment(layout.Center),
				)(c)

				// Bottom App Bar with FAB
				bottomappbar.BottomAppBar(
					func(c api.Composer) api.Composer {
						return row.Row(
							func(c api.Composer) api.Composer {
								IconButton(mdicons.NavigationMenu)(c)
								IconButton(mdicons.ActionSearch)(c)
								IconButton(mdicons.ContentSave)(c)
								IconButton(mdicons.SocialShare)(c)
								return c
							},
						)(c)
					},
					bottomappbar.WithFloatingActionButton(
						func(c api.Composer) api.Composer {
							// Simulate FAB
							return iconbutton.Filled(
								func() { /* click */ },
								mdicons.ContentAdd,
								"Add",
							)(c)
						},
					),
				)(c)
				return c
			},
			column.WithModifier(size.FillMax()),
		)(c)
	}
}

func IconButton(iconData []byte) api.Composable {
	return func(c api.Composer) api.Composer {
		return iconbutton.Standard(
			func() {
				// OnClick
			},
			iconData,
			"Icon",
		)(c)
	}
}
