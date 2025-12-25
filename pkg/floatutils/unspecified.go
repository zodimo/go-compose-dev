package floatutils

import "math"

type Float interface {
	~float32 | ~float64
}

var FloatUnspecified = math.NaN()
var Float32Unspecified = float32(math.NaN())

func TakeOrElseFloat32(v float32, defaultValue float32) float32 {
	if v == Float32Unspecified {
		return defaultValue
	}
	return v
}
