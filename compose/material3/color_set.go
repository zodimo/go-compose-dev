package material3

var ColorSetUnspecified = &ColorSet{
	Color:   ColorUnspecified,
	OnColor: ColorUnspecified,
}

type ColorSet struct {
	Color   Color
	OnColor Color
}

func CoalesceColorSet(ptr, def *ColorSet) *ColorSet {
	if ptr == nil {
		return def
	}
	return ptr
}

func IsSpecifiedColorSet(s *ColorSet) bool {
	return s != nil && s != ColorSetUnspecified
}

func TakeOrElseColorSet(s, def *ColorSet) *ColorSet {
	if !IsSpecifiedColorSet(s) {
		return def
	}
	return s
}

// Identity (2 ns)
func SameColorSet(a, b *ColorSet) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return b == ColorSetUnspecified
	}
	if b == nil {
		return a == ColorSetUnspecified
	}
	return a == b
}

// Semantic equality (field-by-field, 20 ns)
func SemanticEqualColorSet(a, b *ColorSet) bool {

	a = CoalesceColorSet(a, ColorSetUnspecified)
	b = CoalesceColorSet(b, ColorSetUnspecified)

	return a.Color == b.Color &&
		a.OnColor == b.OnColor
}

func EqualColorSet(a, b *ColorSet) bool {
	if !SameColorSet(a, b) {
		return SemanticEqualColorSet(a, b)
	}
	return true
}

func MergeColorSet(a, b *ColorSet) *ColorSet {
	if b == nil {
		return a
	}
	if a == nil {
		return b
	}

	a = CoalesceColorSet(a, ColorSetUnspecified)
	b = CoalesceColorSet(b, ColorSetUnspecified)

	return &ColorSet{
		Color:   b.Color.TakeOrElse(a.Color),
		OnColor: b.OnColor.TakeOrElse(a.OnColor),
	}
}
