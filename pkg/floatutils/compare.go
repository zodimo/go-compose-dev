package floatutils

import "math"

const (
	Float64EqualityThreshold = 1e-9
	Float32EqualityThreshold = 1e-6
)

// floatEquals compares two float32 values with absolute epsilon tolerance.
func Float32Equals(a, b, epsilon float32) bool {
	return math.Abs(float64(a-b)) <= float64(epsilon)
}

// floatEquals compares two float32 values with absolute epsilon tolerance.
func Float64Equals(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}
