package theme

import (
	"image/color"
	"math"
)

// rgbaToHSL converts an RGBA color to HSL color space
// Returns hue (0-360), saturation (0-1), lightness (0-1)
func rgbaToHSL(c color.NRGBA) (h, s, l float32) {
	// Normalize RGB to 0-1
	r := float32(c.R) / 255.0
	g := float32(c.G) / 255.0
	b := float32(c.B) / 255.0

	// Find min and max values
	max := r
	if g > max {
		max = g
	}
	if b > max {
		max = b
	}

	min := r
	if g < min {
		min = g
	}
	if b < min {
		min = b
	}

	// Calculate lightness
	l = (max + min) / 2.0

	// If max == min, it's a shade of gray (no saturation)
	if max == min {
		h = 0.0
		s = 0.0
		return h, s, l
	}

	// Calculate saturation
	if l > 0.5 {
		s = (max - min) / (2.0 - max - min)
	} else {
		s = (max - min) / (max + min)
	}

	// Calculate hue
	switch max {
	case r:
		h = (g - b) / (max - min)
		if g < b {
			h += 6.0
		}
	case g:
		h = 2.0 + (b-r)/(max-min)
	case b:
		h = 4.0 + (r-g)/(max-min)
	}

	h *= 60.0 // Convert to degrees

	return h, s, l
}

// hslToRGBA converts HSL color space to RGBA
// h: hue (0-360), s: saturation (0-1), l: lightness (0-1), a: alpha (0-255)
func hslToRGBA(h, s, l float32, a uint8) color.NRGBA {
	var r, g, b float32

	// If saturation is 0, it's a shade of gray
	if s == 0.0 {
		r = l
		g = l
		b = l
	} else {
		var q float32
		if l < 0.5 {
			q = l * (1.0 + s)
		} else {
			q = l + s - (l * s)
		}
		p := 2.0*l - q

		// Normalize hue to 0-1 range
		hNorm := h / 360.0

		r = hueToRGB(p, q, hNorm+1.0/3.0)
		g = hueToRGB(p, q, hNorm)
		b = hueToRGB(p, q, hNorm-1.0/3.0)
	}

	return color.NRGBA{
		R: uint8(clamp(r*255.0, 0.0, 255.0)),
		G: uint8(clamp(g*255.0, 0.0, 255.0)),
		B: uint8(clamp(b*255.0, 0.0, 255.0)),
		A: a,
	}
}

// hueToRGB is a helper function for HSL to RGB conversion
func hueToRGB(p, q, t float32) float32 {
	// Wrap t to 0-1 range
	if t < 0.0 {
		t += 1.0
	}
	if t > 1.0 {
		t -= 1.0
	}

	if t < 1.0/6.0 {
		return p + (q-p)*6.0*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6.0
	}
	return p
}

// clamp restricts a value to a given range
func clamp(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// lightenColor increases the lightness of a color by the given percentage
// percentage should be in range 0.0 to 1.0 (e.g., 0.2 = 20% lighter)
func lightenColor(c color.NRGBA, percentage float32) color.NRGBA {
	h, s, l := rgbaToHSL(c)

	// Increase lightness
	l = clamp(l+percentage, 0.0, 1.0)

	return hslToRGBA(h, s, l, c.A)
}

// darkenColor decreases the lightness of a color by the given percentage
// percentage should be in range 0.0 to 1.0 (e.g., 0.2 = 20% darker)
func darkenColor(c color.NRGBA, percentage float32) color.NRGBA {
	h, s, l := rgbaToHSL(c)

	// Decrease lightness
	l = clamp(l-percentage, 0.0, 1.0)

	return hslToRGBA(h, s, l, c.A)
}

// saturateColor increases the saturation of a color by the given percentage
// percentage should be in range 0.0 to 1.0 (e.g., 0.2 = 20% more saturated)
func saturateColor(c color.NRGBA, percentage float32) color.NRGBA {
	h, s, l := rgbaToHSL(c)

	// Increase saturation
	s = clamp(s+percentage, 0.0, 1.0)

	return hslToRGBA(h, s, l, c.A)
}

// desaturateColor decreases the saturation of a color by the given percentage
// percentage should be in range 0.0 to 1.0 (e.g., 0.2 = 20% less saturated)
func desaturateColor(c color.NRGBA, percentage float32) color.NRGBA {
	h, s, l := rgbaToHSL(c)

	// Decrease saturation
	s = clamp(s-percentage, 0.0, 1.0)

	return hslToRGBA(h, s, l, c.A)
}

// rotateHue rotates the hue of a color by the given degrees
// degrees can be positive or negative (e.g., 120.0 for complementary color)
func rotateHue(c color.NRGBA, degrees float32) color.NRGBA {
	h, s, l := rgbaToHSL(c)

	// Rotate hue and wrap to 0-360 range
	h = float32(math.Mod(float64(h+degrees), 360.0))
	if h < 0 {
		h += 360.0
	}

	return hslToRGBA(h, s, l, c.A)
}
