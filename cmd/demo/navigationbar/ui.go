package main

import (
	"fmt"
	"image/color"

	"gioui.org/layout"

	"go-compose-dev/compose/foundation/icon"
	"go-compose-dev/compose/foundation/layout/box"
	"go-compose-dev/compose/foundation/layout/column"
	"go-compose-dev/compose/foundation/material3/navigationbar"
	m3text "go-compose-dev/compose/foundation/material3/text"
	"go-compose-dev/compose/foundation/text"
	"go-compose-dev/internal/modifiers/size"
	"go-compose-dev/internal/modifiers/weight"
	"go-compose-dev/pkg/api"

	mdicons "golang.org/x/exp/shiny/materialdesign/icons"
)

func UI() api.Composable {
	return func(c api.Composer) api.Composer {
		// State for selected item
		// We use a pointer to int for the state value
		selectedIdxVal := c.State("nav_selected_index", func() any { return 0 })
		selectedIdx := selectedIdxVal.Get().(int)

		items := []struct {
			Label string
			Icon  []byte
		}{
			{"Home", mdicons.ActionHome},
			{"Search", mdicons.ActionSearch},
			{"Settings", mdicons.ActionSettings},
		}

		return column.Column(
			func(c api.Composer) api.Composer {
				// Content Area
				box.Box(
					func(c api.Composer) api.Composer {
						return m3text.Text(
							fmt.Sprintf("Selected: %s", items[selectedIdx].Label),
							m3text.TypestyleDisplayMedium,
							text.WithTextStyleOptions(
								text.StyleWithColor(color.NRGBA{A: 255}),
							),
						)(c)
					},
					box.WithModifier(weight.Weight(1)),
					box.WithAlignment(layout.Center),
				)(c)

				// Navigation Bar
				navigationbar.NavigationBar(
					func(c api.Composer) api.Composer {
						for i, item := range items {
							idx := i // capture loop variable
							navigationbar.NavigationBarItem(
								selectedIdx == idx,
								func() {
									selectedIdxVal.Set(idx)
								},
								func(c api.Composer) api.Composer {
									return icon.Icon(item.Icon)(c)
								},
								func(c api.Composer) api.Composer {
									return m3text.Text(item.Label, m3text.TypestyleLabelMedium)(c)
								},
								navigationbar.EmptyModifier,
							)(c)
						}
						return c
					},
				)(c)
				return c
			},
			column.WithModifier(size.FillMax()),
		)(c)
	}
}
