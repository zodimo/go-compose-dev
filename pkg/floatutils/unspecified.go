package floatutils

import "math"

type Float interface {
	~float32 | ~float64
}

var Float64Unspecified = math.NaN()
var Float32Unspecified = float32(math.NaN())

func TakeOrElse[T Float](v T, defaultValue T) T {
	if math.IsNaN(float64(v)) {
		return defaultValue
	}
	return v
}
