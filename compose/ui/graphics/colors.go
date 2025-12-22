package graphics

import (
	"image/color"

	"github.com/zodimo/go-compose/theme"
)

var (
	Transparent = theme.ColorHelper.SpecificColor(color.RGBA{A: 0})
	Black       = theme.ColorHelper.SpecificColor(color.RGBA{R: 0, G: 0, B: 0, A: 255})
)
