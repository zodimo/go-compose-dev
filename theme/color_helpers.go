package theme

import "image/color"

func toNRGBA(c color.Color) color.NRGBA {
	return c.(color.NRGBA)
}
