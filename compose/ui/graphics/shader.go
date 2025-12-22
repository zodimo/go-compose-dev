package graphics

// Shader corresponds to the Android/Skia Shader class.
// It is used to draw gradients and bitmaps.
type Shader interface {
	isShader()
}

type LinearGradientShader struct {
	Colors     []Color
	ColorStops []float32
	From       Offset
	To         Offset
	TileMode   TileMode
}

func (s LinearGradientShader) isShader() {}

type RadialGradientShader struct {
	Colors     []Color
	ColorStops []float32
	Center     Offset
	Radius     float32
	TileMode   TileMode
}

func (s RadialGradientShader) isShader() {}

type SweepGradientShader struct {
	Center     Offset
	Colors     []Color
	ColorStops []float32
}

func (s SweepGradientShader) isShader() {}

type CompositeShader struct {
	Dst       Shader
	Src       Shader
	BlendMode BlendMode
}

func (s CompositeShader) isShader() {}
