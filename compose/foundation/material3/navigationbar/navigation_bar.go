package navigationbar

import (
	"go-compose-dev/compose/foundation/layout/row"
	"go-compose-dev/compose/foundation/material3/surface"
	"go-compose-dev/internal/modifiers/size"
)

// NavigationBar displays a navigation bar at the bottom of the screen.
//
// Material 3 Specs:
// - Height: 80dp
// - Layout: Items are equally distributed.
func NavigationBar(
	content Composable,
	options ...NavigationBarOption,
) Composable {
	return func(c Composer) Composer {
		opts := DefaultNavigationBarOptions()
		for _, option := range options {
			option(&opts)
		}

		return surface.Surface(
			func(c Composer) Composer {
				return row.Row(
					func(c Composer) Composer {
						if content != nil {
							content(c)
						}
						return c
					},
					row.WithModifier(size.FillMaxWidth()),
					row.WithModifier(size.FillMaxHeight()),
					// Items utilize weight to distribute space evenly
					row.WithAlignment(row.Middle), // Vertically centered
				)(c)
			},
			surface.WithModifier(
				opts.Modifier.
					Then(size.FillMaxWidth()).
					Then(size.Height(int(opts.Height))),
			),
			surface.WithColor(opts.ContainerColor),
			surface.WithContentColor(opts.ContentColor),
			surface.WithTonalElevation(opts.TonalElevation),
		)(c)
	}
}
