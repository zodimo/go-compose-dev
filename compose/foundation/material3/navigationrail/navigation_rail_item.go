package navigationrail

import (
	"fmt"
	"image/color"

	"go-compose-dev/compose/foundation/layout/box"
	"go-compose-dev/compose/foundation/layout/column"
	"go-compose-dev/compose/ui/graphics/shape"
	"go-compose-dev/internal/modifiers/clickable"
	"go-compose-dev/internal/modifiers/clip"
	"go-compose-dev/internal/modifiers/padding"
	"go-compose-dev/internal/modifiers/size"
	"go-compose-dev/internal/theme"

	"go-compose-dev/pkg/api"

	"go-compose-dev/compose/foundation/layout/spacer"
	"go-compose-dev/compose/foundation/material3/surface"

	"gioui.org/unit"

	"gioui.org/widget"
)

// NavigationRailItem represents an item within a NavigationRail.
//
// Material 3 Specs:
// - Height: Can vary, usually minimum touch target.
// - Indicator: Pill shape around icon when selected.
// - Label: Text below icon.
func NavigationRailItem(
	selected bool,
	onClick func(),
	icon Composable,
	label Composable,
	modifier Modifier,
	// colors, enabled, interactionSource... omitted for brevity
) Composable {
	return func(c Composer) Composer {
		// State for click interaction
		key := c.GenerateID()
		path := c.GetPath()
		clickStatePath := fmt.Sprintf("%d/%s/railitem_click", key, path)
		clickValue := c.State(clickStatePath, func() any { return &widget.Clickable{} })
		clickWidget := clickValue.Get().(*widget.Clickable)

		// Define indicator styling (pill shape)
		// Usually 56x32dp or similar for indicator.
		// We'll wrap the icon in a Surface that acts as the indicator.

		return column.Column(
			func(c Composer) Composer {
				// Icon Container (Indicator)
				surface.Surface(
					icon,
					surface.WithColor(func() color.Color {
						if selected {
							tm := theme.GetThemeManager()
							m3 := tm.GetMaterial3Theme()
							// Secondary Container color for selected state
							return m3.Scheme.SecondaryContainer.Color.AsNRGBA()
						}
						return color.NRGBA{A: 0} // Transparent
					}()),
					surface.WithShape(shape.RoundedCornerShape{Radius: unit.Dp(12)}), // Pill shape (approx)
					surface.WithModifier(
						api.EmptyModifier.
							Then(size.FillMaxWidth()).
							Then(size.Height(32)).
							Then(clip.Clip(shape.RoundedCornerShape{Radius: unit.Dp(12)})). // Clip to pill shape
							Then(padding.Padding(4, 4, 4, 4)),                              // Padding inside indicator? Or just center icon.
					),
					surface.WithAlignment(box.Center),
				)(c)

				// Label
				if label != nil {
					spacer.Spacer(4)(c)
					label(c) // Should be styled with smaller font-size usually
				}
				return c
			},
			column.WithAlignment(column.Middle),
			column.WithModifier(
				modifier.
					Then(size.FillMaxWidth()).
					Then(clickable.OnClick(func() {
						if onClick != nil {
							onClick()
						}
					}, clickable.WithClickable(clickWidget))).
					Then(padding.Padding(0, 4, 0, 4)),
			),
		)(c)
	}
}
