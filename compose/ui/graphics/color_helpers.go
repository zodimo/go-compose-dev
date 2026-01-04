package graphics

import "image/color"

func SetOpacity(c Color, opacity float32) Color {
	return c.Copy(CopyWithAlpha(opacity))
}
func ColorToNRGBA(c Color) color.NRGBA {
	return color.NRGBA{
		R: uint8(c.Red() * 255),
		G: uint8(c.Green() * 255),
		B: uint8(c.Blue() * 255),
		A: uint8(c.Alpha() * 255),
	}
}

func FromNRGBA(c color.NRGBA) Color {
	return NewColorSrgb(int(c.R), int(c.G), int(c.B), int(c.A))
}

func (c Color) ToNRGBA() color.NRGBA {
	return color.NRGBA{
		R: uint8(c.Red() * 255),
		G: uint8(c.Green() * 255),
		B: uint8(c.Blue() * 255),
		A: uint8(c.Alpha() * 255),
	}
}

func (c Color) SetOpacity(opacity float32) Color {
	return c.Copy(CopyWithAlpha(opacity))
}
