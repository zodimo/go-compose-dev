package theme

import (
	"github.com/zodimo/go-compose/assets/fonts"

	"gioui.org/text"
	"gioui.org/widget/material"
)

func defaultMaterialTheme() *material.Theme {
	materialTheme := material.NewTheme()
	materialTheme.Shaper = text.NewShaper(text.WithCollection(fonts.Collection()))

	return materialTheme
}
