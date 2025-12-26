package graphics

import (
	"math"

	"github.com/zodimo/go-compose/compose/ui/geometry"
)

// RadialGradient Brush implementation.
type RadialGradient struct {
	Colors   []Color
	Stops    []float32
	Center   geometry.Offset
	Radius   float32
	TileMode TileMode
}

func (r RadialGradient) isBrush() {}

func (r RadialGradient) ApplyTo(size geometry.Size, p *Paint, alpha float32) {
	applyToShaderBrush(r, size, p, alpha)
}

func (r RadialGradient) IntrinsicSize() geometry.Size {
	if !math.IsInf(float64(r.Radius), 0) && !math.IsNaN(float64(r.Radius)) {
		return geometry.NewSize(r.Radius*2, r.Radius*2)
	}
	return geometry.SizeUnspecified
}

func (r RadialGradient) CreateShader(size geometry.Size) Shader {
	centerX := r.Center.X()
	centerY := r.Center.Y()
	if r.Center.IsUnspecified() {
		center := size.Center()
		centerX = center.X()
		centerY = center.Y()
	} else {
		if centerX == float32(math.Inf(1)) {
			centerX = size.Width()
		}
		if centerY == float32(math.Inf(1)) {
			centerY = size.Height()
		}
	}
	radius := r.Radius
	if radius == float32(math.Inf(1)) {
		radius = size.MinDimension() / 2
	}

	return RadialGradientShader{
		Colors:     r.Colors,
		ColorStops: r.Stops,
		Center:     geometry.NewOffset(centerX, centerY),
		Radius:     radius,
		TileMode:   r.TileMode,
	}
}

func (r RadialGradient) Equal(other Brush) bool {
	o, ok := other.(RadialGradient)
	if !ok {
		return false
	}
	if len(r.Colors) != len(o.Colors) {
		return false
	}
	for i := range r.Colors {
		if r.Colors[i] != o.Colors[i] {
			return false
		}
	}
	if !float32SliceEqual(r.Stops, o.Stops) {
		return false
	}
	if !r.Center.Equal(o.Center) {
		return false
	}
	if r.Radius != o.Radius {
		return false
	}
	if r.TileMode != o.TileMode {
		return false
	}
	return true
}

func RadialGradientBrush(colors []Color, center geometry.Offset, radius float32, tileMode TileMode) RadialGradient {
	// Defaults: center=Unspecified, radius=Infinite, tileMode=Clamp
	return RadialGradient{
		Colors:   colors,
		Center:   center,
		Radius:   radius,
		TileMode: tileMode,
	}
}
