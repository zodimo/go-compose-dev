package navigationdrawer

import (
	"go-compose-dev/compose/foundation/layout/box"
	"go-compose-dev/compose/foundation/material3/surface"
	"go-compose-dev/compose/ui/graphics/shape"
	"go-compose-dev/internal/modifier"
	"go-compose-dev/internal/modifiers/background"
	"go-compose-dev/internal/modifiers/clickable"
	"go-compose-dev/internal/modifiers/size"
	"go-compose-dev/internal/theme"
	"go-compose-dev/pkg/api"

	"gioui.org/unit"
	"git.sr.ht/~schnwalter/gio-mw/token"
)

type Modifier = modifier.Modifier
type Composable = api.Composable
type Composer = api.Composer

// ModalNavigationDrawer implements a navigation drawer that overlays the content.
// It uses a generic Box layout to stack a scrim and the drawer sheet over the content.
func ModalNavigationDrawer(
	drawerContent Composable,
	content Composable,
	options ...ModalNavigationDrawerOption,

) Composable {
	opts := DefaultModalNavigationDrawerOptions()
	for _, opt := range options {
		opt(&opts)
	}

	return func(c Composer) Composer {
		tm := theme.GetThemeManager()
		m3 := tm.GetMaterial3Theme()

		scrimColor := m3.Scheme.Scrim.SetOpacity(token.OpacityLevel8).AsNRGBA() // 32% opacity

		drawerContainerColor := m3.Scheme.SurfaceContainerLow.AsNRGBA()

		return box.Box(
			func(c Composer) Composer {
				// 1. Main Content (Background)
				content(c)

				// 2. Overlay (Scrim + Drawer)
				if opts.IsOpen {
					// Scrim
					box.Box(
						func(c Composer) Composer { return c },
						box.WithModifier(
							modifier.EmptyModifier.
								Then(size.FillMax()).
								Then(background.Background(scrimColor)).
								Then(clickable.OnClick(func() {
									if opts.OnClose != nil {
										opts.OnClose()
									}
								})),
						),
					)(c)

					// Drawer Sheet
					surface.Surface(
						drawerContent,
						surface.WithColor(drawerContainerColor),
						surface.WithShape(shape.RoundedCornerShape{Radius: unit.Dp(16)}), // Uniform radius for now
						surface.WithModifier(
							modifier.EmptyModifier.
								Then(size.Width(360)).
								Then(size.FillMaxHeight()),
						),
					)(c)
				}
				return c
			},
		)(c)
	}
}
