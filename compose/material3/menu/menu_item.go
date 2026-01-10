package menu

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	fText "github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/compose/material3/text"
	"github.com/zodimo/go-compose/compose/ui"
	uiText "github.com/zodimo/go-compose/compose/ui/text"
	"github.com/zodimo/go-compose/modifiers/clickable"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
)

// DropdownMenuItem Composable
func DropdownMenuItem(
	textStr string,
	onClick func(),
	options ...DropdownMenuItemOption,
) api.Composable {
	return func(c Composer) Composer {
		opts := DefaultDropdownMenuItemOptions()
		for _, option := range options {
			if option == nil {
				continue
			}
			option(&opts)
		}

		colors := MenuDefaults.ItemColors(c)

		// Determine colors based on enabled state using helper methods
		textColor := colors.TextColorFor(opts.Enabled)

		// TODO: Apply ripple indicator in Clickable when available

		return box.Box(
			func(c Composer) Composer {
				return row.Row(
					func(c Composer) Composer {
						// Leading Icon
						if opts.LeadingIcon != nil {
							leadingIconMod := ui.Modifier(padding.Padding(
								padding.NotSet, padding.NotSet,
								int(DropdownMenuItemHorizontalPadding), padding.NotSet,
							))
							box.Box(
								opts.LeadingIcon,
								box.WithModifier(leadingIconMod),
							)(c)
						}

						// Text - M3 spec: Label Large
						text.LabelLarge(
							textStr,
							fText.WithTextStyleOptions(
								uiText.WithColor(textColor),
							),
						)(c)

						// Trailing Icon
						if opts.TrailingIcon != nil {
							trailingIconMod := ui.Modifier(padding.Padding(
								int(DropdownMenuItemHorizontalPadding), padding.NotSet,
								padding.NotSet, padding.NotSet,
							))
							box.Box(
								opts.TrailingIcon,
								box.WithModifier(trailingIconMod),
							)(c)
						}
						return c
					},
					row.WithAlignment(row.Middle),
				)(c)
			},
			box.WithModifier(
				size.WrapContentWidth().
					Then(size.MinWidth(int(DropdownMenuItemDefaultMinWidth))).
					Then(size.Height(int(MenuListItemContainerHeight))).
					Then(clickable.OnClick(onClick)).
					Then(padding.Horizontal(
						int(DropdownMenuItemHorizontalPadding),
						int(DropdownMenuItemHorizontalPadding),
					)).
					Then(opts.Modifier),
			),
			box.WithAlignment(box.W), // Align Content to Center Start (West)
		)(c)
	}
}
