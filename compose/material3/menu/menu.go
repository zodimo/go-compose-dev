package menu

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/material3/surface"
	"github.com/zodimo/go-compose/compose/ui/window" // Import internal modifier
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
)

func MenuItems(items ...api.Composable) []api.Composable {
	return items
}

// DropdownMenu Composable
// Uses a Popup to display content on top of other content.
func DropdownMenu(
	expanded bool,
	onDismissRequest func(),
	menuItems []api.Composable,
	options ...DropdownMenuOption,
) api.Composable {
	return func(c api.Composer) api.Composer {
		if !expanded {
			return c
		}

		opts := DefaultDropdownMenuOptions()
		for _, option := range options {
			if option == nil {
				continue
			}
			option(&opts)
		}

		return window.Popup(
			surface.Surface(
				column.Column(
					c.Sequence(
						menuItems...,
					),
					column.WithModifier(
						padding.Vertical(
							int(DropdownMenuVerticalPadding),
							int(DropdownMenuVerticalPadding),
						),
					),
				),
				// M3 Menu Spec: Shape Extra Small (4dp)
				surface.WithShape(MenuDefaults.Shape()),
				// Container Color: Surface (default)
				// Elevation: Level 2 (3.dp)
				surface.WithShadowElevation(ShadowElevation),
				surface.WithModifier(
					// M3 Specs:
					// Min width: 112.dp
					// Max width: 280.dp
					// TODO: Implement MinWidth/MaxWidth modifiers in size package.
					// For now, we depend on content width.
					size.
						WrapContentWidth().
						Then(opts.Modifier),
				),
			),
			window.WithOffset(opts.OffsetX, opts.OffsetY),
			window.WithOnDismissRequest(onDismissRequest),
		)(c)
	}
}
