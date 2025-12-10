package helpers

import "image/color"

func ToNRGBA(input color.Color) color.NRGBA {
	nrgbaModel := color.NRGBAModel
	return nrgbaModel.Convert(input).(color.NRGBA)
}

func Clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
