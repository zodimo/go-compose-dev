package sentinel

import (
	"fmt"
	"math"
)

type IntValue = int

// 1. Sentinel - IntValueUnspecified
// IntValue is the type for sentinel int pattern.
// The sentinel math.MinInt is used.
// Note: This means math.MinInt cannot be used as a valid value.

const IntValueUnspecified IntValue = math.MinInt

// Deprecated: Use IntValueUnspecified instead
const IntUnspecified IntValue = IntValueUnspecified

// 2. IsSpecified - predicate (package-level function)
func IsSpecifiedIntValue(i IntValue) bool {
	return i != IntValueUnspecified
}

// IsUnspecifiedIntValue - convenience predicate
func IsUnspecifiedIntValue(i IntValue) bool {
	return i == IntValueUnspecified
}

// 3. TakeOrElse - 2-param fallback (package-level function)
func TakeOrElseIntValue(a, b IntValue) IntValue {
	if a != IntValueUnspecified {
		return a
	}
	return b
}

// 4. Merge - composition merge (package-level function)
// Prefers incoming specified values over current values
func MergeIntValue(a, b IntValue) IntValue {
	if b != IntValueUnspecified {
		return b
	}
	return a
}

// 5. String - stringification (package-level function)
func StringIntValue(i IntValue) string {
	if i == IntValueUnspecified {
		return "IntValue{Unspecified}"
	}
	return fmt.Sprintf("IntValue{%d}", i)
}

// 6. Coalesce - N/A for int type (int is a value type, no nil possible)
// Not applicable

// 7. Same - identity (package-level function)
func SameIntValue(a, b IntValue) bool {
	return a == b
}

// 8. SemanticEqual - semantic equality (package-level function)
// For ints, this is the same as Same
func SemanticEqualIntValue(a, b IntValue) bool {
	return a == b
}

// 9. Equal - equality check (package-level function)
func EqualIntValue(a, b IntValue) bool {
	return a == b
}

// 10. Copy - identity for immutable value types (package-level function)
func CopyIntValue(i IntValue) IntValue {
	return i
}
