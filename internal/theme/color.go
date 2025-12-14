package theme

import (
	"image/color"
)

type ThemeColor interface {
	AsHex() string
	AsNRGBA() color.NRGBA
	SetOpacity(opacity OpacityLevel) ThemeColor
	AsTokenColor() TokenColor
}

type ThemeColorDescriptor struct {
	color     color.NRGBA
	colorRole ColorRole
	isColor   bool
}
