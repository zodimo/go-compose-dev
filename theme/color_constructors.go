package theme

import (
	"image/color"

	"github.com/zodimo/go-compose/compose/ui/graphics"
)

func SpecificColor(color graphics.Color) ColorDescriptor {
	return colorDescriptor{
		color:   color,
		isColor: true,
	}
}

func ThemeRoleColor(colorRole ColorRole) ColorDescriptor {
	return colorDescriptor{
		colorRole: colorRole,
	}
}

func ThemeColorFromGraphicsColor(color graphics.Color) ThemeColor {
	return themeColor{
		tokenColor: TokenColor(graphics.ColorToNRGBA(color)),
	}
}

func ThemeColorFromColor(color color.Color) ThemeColor {
	return themeColor{
		tokenColor: TokenColor(toNRGBA(color)),
	}
}
func ThemeColorFromNRGBA(color color.NRGBA) ThemeColor {
	return themeColor{
		tokenColor: TokenColor(color),
	}
}

func ThemeColorFromTokenColor(color TokenColor) ThemeColor {
	return themeColor{
		tokenColor: color,
	}
}
