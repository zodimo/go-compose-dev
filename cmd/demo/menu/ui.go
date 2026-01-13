package main

import (
	"fmt"

	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/material3/button"
	"github.com/zodimo/go-compose/compose/material3/icon"
	"github.com/zodimo/go-compose/compose/material3/menu"
	"github.com/zodimo/go-compose/compose/material3/scaffold"
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/pkg/api"

	mdicons "golang.org/x/exp/shiny/materialdesign/icons"
)

func UI() api.Composable {
	return func(c api.Composer) api.Composer {

		theme := material3.Theme(c)
		// State for menus
		expanded1 := c.State("menu1", func() any { return false })
		expanded2 := c.State("menu2", func() any { return false })

		return scaffold.Scaffold(
			func(c api.Composer) api.Composer {
				return column.Column(
					c.Sequence(
						// Demo 1: Simple Dropdown
						box.Box(
							func(c api.Composer) api.Composer {
								button.Filled(
									func() {
										expanded1.Set(true)
									},
									"Open Menu",
								)(c)

								menu.DropdownMenu(
									expanded1.Get().(bool),
									func() { expanded1.Set(false) },
									menu.MenuItems(
										menu.DropdownMenuItem(
											"Item 1 (Selected)",
											func() {
												fmt.Println("Item 1 Clicked")
												expanded1.Set(false)
											},
											menu.WithLeadingIcon(func(c api.Composer) api.Composer {
												return icon.Icon(
													mdicons.ActionDone,
													icon.WithColor(theme.ColorScheme().Primary),
												)(c)
											}),
										),
										menu.DropdownMenuItem(
											"Item 2 (Disabled)",
											func() {},
											menu.WithEnabled(false),
										),
										menu.DropdownMenuItem(
											"Item 3",
											func() {
												fmt.Println("Item 3 Clicked")
												expanded1.Set(false)
											},
										),
									),

									// Options
									menu.WithModifier(ui.EmptyModifier),
									menu.WithOffset(unit.Dp(0), unit.Dp(40)), // Position below button
								)(c)
								return c
							},
						),
						spacer.Height(32),
						// Demo 2: Leading/Trailing Icons
						box.Box(
							func(c api.Composer) api.Composer {
								button.FilledTonal(
									func() {
										expanded2.Set(true)
									},
									"Menu with Icons",
								)(c)

								menu.DropdownMenu(
									expanded2.Get().(bool),
									func() { expanded2.Set(false) },
									menu.MenuItems(
										menu.DropdownMenuItem(
											"Edit",
											func() { expanded2.Set(false) },
											menu.WithLeadingIcon(func(c api.Composer) api.Composer {
												return icon.Icon(
													mdicons.ContentCreate,
													icon.WithColor(theme.ColorScheme().OnSurface),
												)(c)
											}),
										),
										menu.DropdownMenuItem(
											"Settings",
											func() { expanded2.Set(false) },
											menu.WithLeadingIcon(func(c api.Composer) api.Composer {
												return icon.Icon(
													mdicons.ActionSettings,
													icon.WithColor(theme.ColorScheme().OnSurface),
												)(c)
											}),
										),
										menu.DropdownMenuItem(
											"Share",
											func() { expanded2.Set(false) },
											menu.WithLeadingIcon(func(c api.Composer) api.Composer {
												return icon.Icon(
													mdicons.SocialShare,
													icon.WithColor(theme.ColorScheme().OnSurface),
												)(c)
											}),
										),
									),
									// Options
									menu.WithOffset(unit.Dp(0), unit.Dp(48)), // Position below button
								)(c)
								return c
							},
						),
					),
					column.WithModifier(padding.All(32)),
				)(c)
			},
			// Scaffold Options
		)(c)
	}
}
