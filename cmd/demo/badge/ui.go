package main

import (
	"image/color"

	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/foundation/material3/badge"
	"github.com/zodimo/go-compose/compose/foundation/material3/text"
	"github.com/zodimo/go-compose/pkg/api"
)

func UI() api.Composable {
	return func(c api.Composer) api.Composer {
		return column.Column(
			func(c api.Composer) api.Composer {
				text.Text("Badge Components", text.TypestyleHeadlineMedium)(c)
				spacer.Height(16)(c)

				// Small badge (dot)
				text.Text("Small Badge (Dot)", text.TypestyleTitleSmall)(c)
				spacer.Height(8)(c)
				badge.Badge()(c)

				spacer.Height(16)(c)

				// Large badge with number
				text.Text("Large Badge (Number)", text.TypestyleTitleSmall)(c)
				spacer.Height(8)(c)
				badge.Badge(badge.WithText("999+"))(c)

				spacer.Height(16)(c)

				// BadgedBox examples
				text.Text("BadgedBox Examples", text.TypestyleTitleSmall)(c)
				spacer.Height(8)(c)

				row.Row(func(c api.Composer) api.Composer {
					// Dot badge on text
					badge.BadgedBox(
						badge.Badge(),
						text.Text("🔔", text.TypestyleHeadlineMedium),
					)(c)

					spacer.Width(24)(c)

					// Number badge on icon
					badge.BadgedBox(
						badge.Badge(badge.WithText("5")),
						text.Text("✉️", text.TypestyleHeadlineMedium),
					)(c)

					spacer.Width(24)(c)

					// Custom color badge
					badge.BadgedBox(
						badge.Badge(
							badge.WithText("!"),
							badge.WithContainerColor(color.NRGBA{R: 0, G: 200, B: 0, A: 255}),
						),
						text.Text("📦", text.TypestyleHeadlineMedium),
					)(c)

					return c
				})(c)

				return c
			},
		)(c)
	}
}
