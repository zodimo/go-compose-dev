package text

import (
	"github.com/zodimo/go-compose/internal/color/f32color"

	"gioui.org/font/gofont"

	"gioui.org/text"
	"gioui.org/widget/material"
)

func DefaultTextStyleOptions(theme *material.Theme) *TextStyleOptions {
	Options := &TextStyleOptions{
		Color:          theme.Fg,
		SelectionColor: f32color.MulAlpha(theme.ContrastBg, 0x60),
		TextSize:       theme.TextSize,
	}
	Options.Font.Typeface = theme.Face
	return Options
}
func DefaultTextOptions() TextOptions {

	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	textStyleOptions := DefaultTextStyleOptions(th)

	return TextOptions{
		Modifier:         EmptyModifier,
		TextStyleOptions: textStyleOptions,
		Shaper:           th.Shaper,
	}
}
