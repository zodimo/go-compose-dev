package text

import (
	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/widget/material"

	"git.sr.ht/~schnwalter/gio-mw/token"
	"github.com/zodimo/go-compose/theme"
)

func DefaultTextStyleOptions() *TextStyleOptions {
	Options := &TextStyleOptions{
		Color:          theme.ColorHelper.ColorSelector().SurfaceRoles.OnSurface,
		SelectionColor: theme.ColorHelper.ColorSelector().PrimaryRoles.Primary.SetOpacity(token.OpacityLevel8),
		TextSize:       14, // Default text size in SP
	}
	return Options
}

func DefaultTextOptions() TextOptions {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	textStyleOptions := DefaultTextStyleOptions()
	textStyleOptions.Font.Typeface = th.Face

	return TextOptions{
		Modifier:         EmptyModifier,
		TextStyleOptions: textStyleOptions,
		Shaper:           th.Shaper,
	}
}
