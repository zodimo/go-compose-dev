package main

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/material3/badge"
	"github.com/zodimo/go-compose/compose/material3/text"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/theme"
)

func UI() api.Composable {
	return func(c api.Composer) api.Composer {
		return column.Column(
			c.Sequence(
				text.TextWithStyle("Badge Components", text.TypestyleHeadlineMedium),
				spacer.Height(16),

				// Small badge (dot)
				text.TextWithStyle("Small Badge (Dot)", text.TypestyleTitleSmall),
				spacer.Height(8),
				badge.Badge(),
				spacer.Height(16),

				// Large badge with number
				text.TextWithStyle("Large Badge (Number)", text.TypestyleTitleSmall),
				spacer.Height(8),
				badge.Badge(badge.WithText("999+")),

				spacer.Height(16),
				// BadgedBox examples
				text.TextWithStyle("BadgedBox Examples", text.TypestyleTitleSmall),
				spacer.Height(8),
				row.Row(
					c.Sequence(
						// Dot badge on text
						badge.BadgedBox(
							badge.Badge(),
							text.TextWithStyle("🔔", text.TypestyleHeadlineMedium),
						),
						spacer.Width(24),

						// Number badge on icon
						badge.BadgedBox(
							badge.Badge(badge.WithText("5")),
							text.TextWithStyle("✉️", text.TypestyleHeadlineMedium),
						),

						spacer.Width(24),

						// Custom color badge
						badge.BadgedBox(
							badge.Badge(
								badge.WithText("!"),
								badge.WithContainerColor(theme.ColorHelper.SpecificColor(graphics.NewColorSrgb(0, 200, 0, 255))),
							),
							text.TextWithStyle("📦", text.TypestyleHeadlineMedium),
						),
					),
				),
			),
		)(c)
	}
}
