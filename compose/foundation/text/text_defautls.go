package text

import (
	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/widget/material"

	"github.com/zodimo/go-compose/theme"
)

func DefaultTextStyleOptions() *TextStyleOptions {
	Options := &TextStyleOptions{
		Color:          theme.ColorHelper.UnspecifiedColor(),
		SelectionColor: theme.ColorHelper.UnspecifiedColor(),

		TextSize: 14, // Default text size in SP
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
