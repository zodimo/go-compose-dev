package sentinel

import "fmt"

type StringValue = string

// 1. Sentinel - StringUnspecified
// StringValue is the type for sentinel string pattern.
// The sentinel "\x00unspecified" is used when empty string is meaningful.
var StringUnspecified StringValue = "\x00unspecified"

// 2. IsSpecified - predicate (package-level function)
func IsSpecifiedString(s StringValue) bool {
	return s != StringUnspecified
}

// IsUnspecifiedString - convenience predicate
func IsUnspecifiedString(s StringValue) bool {
	return s == StringUnspecified
}

// 3. TakeOrElse - 2-param fallback (package-level function)
func TakeOrElseString(a, b StringValue) StringValue {
	if a != StringUnspecified {
		return a
	}
	return b
}

// 4. Merge - composition merge (package-level function)
// Prefers incoming specified values over current values
func MergeString(a, b StringValue) StringValue {
	if b != StringUnspecified {
		return b
	}
	return a
}

// 5. String - stringification (package-level function)
func StringString(s StringValue) string {
	if s == StringUnspecified {
		return "StringValue{Unspecified}"
	}
	return fmt.Sprintf("StringValue{%q}", s)
}

// 6. Coalesce - N/A for string type (string is a value type, no nil possible)
// Not applicable

// 7. Same - identity (package-level function)
func SameString(a, b StringValue) bool {
	return a == b
}

// 8. SemanticEqual - semantic equality (package-level function)
// For strings, this is the same as Same
func SemanticEqualString(a, b StringValue) bool {
	return a == b
}

// 9. Equal - equality check (package-level function)
func EqualString(a, b StringValue) bool {
	return a == b
}

// 10. Copy - identity for immutable value types (package-level function)
func CopyString(s StringValue) StringValue {
	return s
}
