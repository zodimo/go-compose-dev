package icon

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/compose/material3"
	uitext "github.com/zodimo/go-compose/compose/ui/text"
	"github.com/zodimo/go-compose/compose/ui/text/font"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/modifiers/size"
)

// MaterialSymbolsFontFamily is the font family for Material Symbols.
// It assumes the font "Material Symbols Outlined" has been loaded into the shaper.
var MaterialSymbolsFontFamily = font.NewGenericFontFamily("Material Symbols Outlined", "Material Symbols Outlined")

// Symbol renders a Material Symbol icon using the Material Symbols Outlined font.
// Use SymbolName constants (e.g. SymbolSearch, SymbolHome) to avoid typos.
func Symbol(name SymbolName, options ...IconOption) Composable {
	opts := DefaultIconOptions()
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&opts)
	}

	return func(c Composer) Composer {
		// Resolve color
		contentColor := material3.LocalContentColor.Current(c)
		iconColor := opts.Color.TakeOrElse(contentColor)

		// Resolve size
		// Material Symbols usually default to 24dp
		fontSize := opts.FontSize.TakeOrElse(unit.Sp(24))

		return box.Box(text.Text(
			string(name),
			text.WithModifier(opts.Modifier),
			text.WithTextStyle(
				uitext.TextStyleFromOptions(
					uitext.WithFontFamily(MaterialSymbolsFontFamily),
					uitext.WithFontSize(fontSize),
					uitext.WithColor(iconColor),
				),
			),
		),
			box.WithModifier(size.WrapContentSize()),
		)(c)
	}
}
