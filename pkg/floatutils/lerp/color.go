package lerp

import "math"

// ============================================================================
// Color Interpolation
// ============================================================================

// TODO: Move to color/lerp package to eliminate coupling with internal packages

// Color performs standard RGBA interpolation.
//
// Example:
//
//	mid := lerp.Color(color1, color2, 0.5)
func Color(a, b struct{ R, G, B, A float32 }, p float32) struct{ R, G, B, A float32 } {
	invP := 1 - p
	return struct{ R, G, B, A float32 }{
		R: a.R*invP + b.R*p,
		G: a.G*invP + b.G*p,
		B: a.B*invP + b.B*p,
		A: a.A*invP + b.A*p,
	}
}

// ColorLerpPrecise uses squared interpolation for gamma-correct color blending.
// ~3x slower but produces perceptually better results.
func ColorLerpPrecise(a, b struct{ R, G, B, A float32 }, p float32) struct{ R, G, B, A float32 } {
	invP := 1 - p
	return struct{ R, G, B, A float32 }{
		R: sqrt(a.R*a.R*invP + b.R*b.R*p),
		G: sqrt(a.G*a.G*invP + b.G*b.G*p),
		B: sqrt(a.B*a.B*invP + b.B*b.B*p),
		A: a.A*invP + b.A*p,
	}
}

func sqrt(v float32) float32 {
	return float32(math.Sqrt(float64(v)))
}

func ColorList(a, b []struct{ R, G, B, A float32 }, t float32) []struct{ R, G, B, A float32 } {
	n := len(a)
	if len(b) > n {
		n = len(b)
	}
	res := make([]struct{ R, G, B, A float32 }, n)
	for i := 0; i < n; i++ {
		c1 := a[min(i, len(a)-1)]
		c2 := b[min(i, len(b)-1)]
		res[i] = Color(c1, c2, t)
	}
	return res
}

func ColorPreciceList(a, b []struct{ R, G, B, A float32 }, t float32) []struct{ R, G, B, A float32 } {
	n := len(a)
	if len(b) > n {
		n = len(b)
	}
	res := make([]struct{ R, G, B, A float32 }, n)
	for i := 0; i < n; i++ {
		c1 := a[min(i, len(a)-1)]
		c2 := b[min(i, len(b)-1)]
		res[i] = ColorLerpPrecise(c1, c2, t)
	}
	return res
}
