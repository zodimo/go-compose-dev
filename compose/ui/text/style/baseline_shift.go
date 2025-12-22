package style

import "math"

// BaselineShift represents the amount by which text is shifted up or down from the current baseline.
// The shift is calculated as: multiplier * (baseline - ascent)
// This is a 32-bit float value class equivalent to Kotlin's BaselineShift.
type BaselineShift float32

// Predefined baseline shift constants
const (
	// Superscript shifts text upward by 0.5 * (baseline - ascent)
	BaselineShiftSuperscript BaselineShift = 0.5
	// Subscript shifts text downward by -0.5 * (baseline - ascent)
	BaselineShiftSubscript BaselineShift = -0.5
	// None applies no baseline shift
	BaselineShiftNone BaselineShift = 0.0
)

// Unspecified represents an unset baseline shift (NaN value)
var Unspecified = BaselineShift(math.NaN())

// IsSpecified returns true if this baseline shift is not Unspecified (i.e., not NaN)
func (bs BaselineShift) IsSpecified() bool {
	return !math.IsNaN(float64(bs))
}

// TakeOrElse returns this BaselineShift if it's specified, otherwise returns the result of block.
// This is equivalent to Kotlin's inline function with lambda.
func (bs BaselineShift) TakeOrElse(block func() BaselineShift) BaselineShift {
	if bs.IsSpecified() {
		return bs
	}
	return block()
}

// Multiplier returns the underlying multiplier value for use in calculations
func (bs BaselineShift) Multiplier() float32 {
	return float32(bs)
}

// NewBaselineShift creates a new BaselineShift with the given multiplier
func NewBaselineShift(multiplier float32) BaselineShift {
	return BaselineShift(multiplier)
}

// LerpBaselineShift interpolates between two BaselineShifts.
func LerpBaselineShift(start, stop BaselineShift, fraction float32) BaselineShift {
	if !start.IsSpecified() || !stop.IsSpecified() {
		// Fallback or discrete or if one is unspecified return implementation specific
		// Usually lerp with unspecified means returning unspecified or stop?
		// Kotlin implementation: return lerp(start.multiplier, stop.multiplier, fraction)
		// If one is NaN, the result will contain NaN which is Unspecified.
		// But we should verify. For now, simple float lerp.
	}
	return BaselineShift(float32(start) + (float32(stop)-float32(start))*fraction)
}
