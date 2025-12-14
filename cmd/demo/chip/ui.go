package main

import (
	"fmt"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/foundation/material3/chip"
	"github.com/zodimo/go-compose/compose/foundation/material3/text"
	"github.com/zodimo/go-compose/pkg/api"
)

func UI() api.Composable {
	return func(c api.Composer) api.Composer {
		// State for Filter chips
		selectedState := c.State("filter_selected", func() any { return false })
		selected := selectedState.Get().(bool)

		return column.Column(
			func(c api.Composer) api.Composer {
				// Title
				text.Text(
					"Chip Components",
					text.TypestyleHeadlineMedium,
				)(c)

				spacer.SpacerHeight(16)(c)

				// Assist Chip Row
				row.Row(
					func(c api.Composer) api.Composer {
						chip.AssistChip(func() { fmt.Println("Assist Chip Clicked") }, "Assist Chip")(c)
						spacer.SpacerWidth(16)(c)
						chip.AssistChip(func() { fmt.Println("Assist with Icon") }, "With Icon",
							chip.WithLeadingIcon(func(c api.Composer) api.Composer {
								return text.Text("★", text.TypestyleBodyMedium)(c)
							}),
						)(c)
						return c
					},
				)(c)

				spacer.SpacerHeight(16)(c)

				// Filter Chip Row
				row.Row(
					func(c api.Composer) api.Composer {
						label := "Filter Chip"
						if selected {
							label = "Selected"
						}

						var leadingIcon api.Composable
						if selected {
							leadingIcon = func(c api.Composer) api.Composer { return text.Text("✓", text.TypestyleBodyMedium)(c) }
						}

						chip.FilterChip(
							func() { selectedState.Set(!selected) },
							label,
							chip.WithSelected(selected),
							chip.WithLeadingIcon(leadingIcon),
						)(c)
						return c
					},
				)(c)

				spacer.SpacerHeight(16)(c)

				// Input Chip
				row.Row(
					func(c api.Composer) api.Composer {
						chip.InputChip(func() { fmt.Println("Input Clicked") }, "Input Chip",
							chip.WithTrailingIcon(func(c api.Composer) api.Composer {
								return text.Text("×", text.TypestyleBodyMedium)(c)
							}),
						)(c)
						return c
					},
				)(c)

				spacer.SpacerHeight(16)(c)

				// Suggestion Chip
				row.Row(
					func(c api.Composer) api.Composer {
						chip.SuggestionChip(func() { fmt.Println("Suggestion 1") }, "Suggestion 1")(c)
						spacer.SpacerWidth(8)(c)
						chip.SuggestionChip(func() { fmt.Println("Suggestion 2") }, "Suggestion 2")(c)
						return c
					},
				)(c)

				return c
			},
		)(c)
	}
}
