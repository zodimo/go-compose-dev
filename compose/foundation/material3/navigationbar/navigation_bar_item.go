package navigationbar

import (
	"fmt"
	"image/color"

	"go-compose-dev/compose/foundation/layout/box"
	"go-compose-dev/compose/foundation/layout/column"
	"go-compose-dev/compose/foundation/layout/spacer"
	"go-compose-dev/compose/foundation/material3/surface"
	"go-compose-dev/compose/ui/graphics/shape"
	"go-compose-dev/internal/modifiers/clickable"
	"go-compose-dev/internal/modifiers/clip"
	"go-compose-dev/internal/modifiers/size"
	"go-compose-dev/internal/modifiers/weight"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

// NavigationBarItem represents an item within a NavigationBar.
//
// Material 3 Specs:
// - Height: Fill container (80dp).
// - Indicator: Pill shape (64x32dp) around icon when selected.
// - Label: Text below icon.
func NavigationBarItem(
	selected bool,
	onClick func(),
	icon Composable,
	label Composable,
	modifier Modifier,
) Composable {
	return func(c Composer) Composer {
		// State for click interaction
		key := c.GenerateID()
		path := c.GetPath()
		clickStatePath := fmt.Sprintf("%d/%s/navitem_click", key, path)
		clickValue := c.State(clickStatePath, func() any { return &widget.Clickable{} })
		clickWidget := clickValue.Get().(*widget.Clickable)

		// Defaults
		colors := NavigationBarDefaults.Colors()

		return box.Box(
			column.Column(
				func(c Composer) Composer {
					spacer.Spacer(12)(c) // Top padding

					// Icon Container (Indicator)
					surface.Surface(
						func(c Composer) Composer {
							return box.Box(
								icon,
								box.WithAlignment(layout.Center),
							)(c)
						},
						surface.WithColor(func() color.Color {
							if selected {
								return colors.IndicatorColor
							}
							return color.NRGBA{A: 0} // Transparent
						}()),
						surface.WithShape(shape.RoundedCornerShape{Radius: unit.Dp(16)}),
						surface.WithModifier(
							EmptyModifier.
								Then(size.Width(64)).
								Then(size.Height(32)).
								Then(clip.Clip(shape.RoundedCornerShape{Radius: unit.Dp(16)})),
						),
						surface.WithAlignment(layout.Center),
					)(c)

					// Label
					if label != nil {
						spacer.Spacer(4)(c)
						label(c)
					}

					spacer.Spacer(12)(c) // Bottom padding
					return c
				},
				column.WithAlignment(column.Middle), // Center children horizontally
			),
			box.WithModifier(
				modifier.
					Then(weight.Weight(1)).     // Fill share of the Row
					Then(size.FillMaxHeight()). // Fill height of the Row
					Then(clickable.OnClick(func() {
						if onClick != nil {
							onClick()
						}
					}, clickable.WithClickable(clickWidget))),
			),
			box.WithAlignment(layout.Center), // Center the Column within the allocated slot
		)(c)
	}
}
