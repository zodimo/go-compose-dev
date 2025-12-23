package theme

import (
	"image/color"
)

func toNRGBA(c color.Color) color.NRGBA {
	if nrgba, ok := c.(color.NRGBA); ok {
		return nrgba
	}
	nrgbaModel := color.NRGBAModel
	return nrgbaModel.Convert(c).(color.NRGBA)
}
