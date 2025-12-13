package navigationdrawer

import (
	"go-compose-dev/compose/foundation/layout/box"
	"go-compose-dev/compose/foundation/material3/surface"
	"go-compose-dev/compose/ui/graphics/shape"
	"go-compose-dev/internal/animation"
	animMod "go-compose-dev/internal/modifiers/animation"
	"go-compose-dev/internal/modifiers/clickable"
	"go-compose-dev/internal/modifiers/size"
	"go-compose-dev/internal/theme"
	"time"

	"gioui.org/unit"
	"git.sr.ht/~schnwalter/gio-mw/token"
)

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

		drawerContainerColor := m3.Scheme.SurfaceContainerLow.AsNRGBA()

		// Animation state
		anim := c.State(c.GenerateID().String()+"/anim", func() any {
			return &animation.VisibilityAnimation{
				Duration: time.Millisecond * 300,
				State:    animation.Invisible,
			}
		}).Get().(*animation.VisibilityAnimation)

		// Sync animation state
		if opts.IsOpen {
			anim.Appear(time.Now())
		} else {
			anim.Disappear(time.Now())
		}

		baseScrim := m3.Scheme.Scrim.SetOpacity(token.OpacityLevel8).AsNRGBA()
		// We use AnimatedBackground for scrim opacity

		return box.Box(
			func(c Composer) Composer {
				// 1. Main Content (Background)
				content(c)

				// 2. Scrim (Animated Opacity)
				// We always render scrim, but AnimatedBackground handles transparency.
				// However, if fully invisible, we might want to skip click handling.
				// For now, let's rely on AnimatedBackground to just hide it visually.
				// But click handling: if opacity is 0, click should ideally pass through?
				// Gio Clickable intercepts even if invisible?
				// Yes.
				// So if !anim.Visible(), we should NOT render the scrim box.
				// But anim.Visible() is stale.
				// We can just render it. The click intercept requires Modal to blocking.
				// If drawer is closed, scrim should be GONE.
				// If we rely on stale `anim.Visible()`, it might flicker.
				// But `opts.IsOpen` is the source of truth for "Goal".
				// If `!opts.IsOpen` AND `anim.State == Invisible`, then we can skip.

				shouldRender := opts.IsOpen || anim.Visible()

				if shouldRender {
					box.Box(
						func(c Composer) Composer { return c },
						box.WithModifier(
							EmptyModifier.
								Then(size.FillMax()).
								Then(animMod.AnimatedBackground(anim, baseScrim, shape.ShapeRectangle)). // RectangleShape singleton
								Then(clickable.OnClick(func() {
									if opts.OnClose != nil {
										opts.OnClose()
									}
								})),
						),
					)(c)

					// 3. Drawer Sheet
					surface.Surface(
						drawerContent,
						surface.WithColor(drawerContainerColor),
						surface.WithShape(shape.RoundedCornerShape{Radius: unit.Dp(16)}),
						surface.WithModifier(
							EmptyModifier.
								Then(animMod.AnimatedWidth(anim, 360)). // Animate width
								Then(size.FillMaxHeight()),
						),
					)(c)
				}

				return c
			},
		)(c)
	}
}
