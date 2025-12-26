package graphics

import (
	"math"

	"github.com/zodimo/go-compose/compose/ui/geometry"
)

// SweepGradient Brush implementation.
type SweepGradient struct {
	Center geometry.Offset
	Colors []Color
	Stops  []float32
}

func (s SweepGradient) isBrush() {}

func (s SweepGradient) ApplyTo(size geometry.Size, p *Paint, alpha float32) {
	applyToShaderBrush(s, size, p, alpha)
}

func (s SweepGradient) IntrinsicSize() geometry.Size {
	return geometry.SizeUnspecified
}

func (s SweepGradient) CreateShader(size geometry.Size) Shader {
	centerX := s.Center.X()
	centerY := s.Center.Y()
	if s.Center.IsUnspecified() {
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
	return SweepGradientShader{
		Center:     geometry.NewOffset(centerX, centerY),
		Colors:     s.Colors,
		ColorStops: s.Stops,
	}
}

func (s SweepGradient) Equal(other Brush) bool {
	o, ok := other.(SweepGradient)
	if !ok {
		return false
	}
	if !s.Center.Equal(o.Center) {
		return false
	}
	if len(s.Colors) != len(o.Colors) {
		return false
	}
	for i := range s.Colors {
		if s.Colors[i] != o.Colors[i] {
			return false
		}
	}
	if !float32SliceEqual(s.Stops, o.Stops) {
		return false
	}
	return true
}

func SweepGradientBrush(colors []Color, center geometry.Offset) SweepGradient {
	return SweepGradient{
		Colors: colors,
		Center: center,
	}
}
