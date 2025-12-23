package floatutils

import "math"

type Float interface {
	~float32 | ~float64
}

var FloatUnspecified = math.NaN()
var Float32Unspecified = float32(math.NaN())
