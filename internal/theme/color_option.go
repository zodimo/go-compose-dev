package theme

import (
	"image/color"
)

type ColorReader = func(themeColor ThemeColor) color.Color

type ThemeColorSet struct {
	ThemeColor ColorReader
}
