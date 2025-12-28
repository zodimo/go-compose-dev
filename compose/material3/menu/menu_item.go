package menu

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	baseText "github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/compose/material3/text"
	"github.com/zodimo/go-compose/internal/modifier"
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
	return func(c api.Composer) api.Composer {
		opts := DefaultDropdownMenuItemOptions()
		for _, option := range options {
			if option == nil {
				continue
			}
			option(&opts)
		}

		colors := DefaultDropdownMenuItemColors()

		// Determine colors based on enabled state
		textColor := colors.TextColor
		if !opts.Enabled {
			textColor = colors.DisabledTextColor
		}

		// TODO: Apply ripple info in Clickable when available

		return box.Box(
			func(c api.Composer) api.Composer {
				return row.Row(
					func(c api.Composer) api.Composer {
						// Leading Icon
						// Fix alignment manually via Box around icon if Padding modifier acts weird.
						// But let's try Padding(NotSet, NotSet, 12, NotSet) for End padding.
						// Padding(start, top, end, bottom).
						leadingIconMod := modifier.Modifier(padding.Padding(padding.NotSet, padding.NotSet, 12, padding.NotSet))

						if opts.LeadingIcon != nil {
							box.Box(
								opts.LeadingIcon,
								box.WithModifier(leadingIconMod),
							)(c)
						}

						// Text
						// M3 spec: Label Large
						// We wrap text in Box to allow weight/grow if needed, but Row handles simple layout well.
						// Text
						text.Text(
							textStr,
							text.TypestyleLabelLarge, // Correct usage
							baseText.WithTextStyleOptions(baseText.StyleWithColor(textColor)),
						)(c)

						// Spacer to push trailing icon?
						// Usually MenuItem fills width, so we might want Weight(1) on Text.
						// But Row here aligns Middle.

						if opts.TrailingIcon != nil {
							// Spacer? Or just padding?
							// Use Box with weight if we want space between.
							// For now just layout next to it.
							trailingIconMod := modifier.Modifier(padding.Padding(12, padding.NotSet, padding.NotSet, padding.NotSet))
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
				size.FillMaxWidth().
					Then(size.Height(48)).
					Then(clickable.OnClick(onClick)).
					Then(padding.Horizontal(12, 12)).
					Then(opts.Modifier),
			),
			box.WithAlignment(box.W), // Align Content to Center Start (West)
		)(c)
	}
}
