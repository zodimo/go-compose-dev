package immap

type ImmutableMap[T any] map[string]T

// EmptyImmutableMap is the canonical empty map.
var EmptyImmutableMapAny = ImmutableMap[any]{}

func EmptyImmutableMap[T any]() ImmutableMap[T] {
	return ImmutableMap[T]{}
}

// Assoc returns a new map with (k,v) added. The receiver is unchanged.
func (m ImmutableMap[T]) Assoc(k string, v T) ImmutableMap[T] {
	out := make(ImmutableMap[T], len(m)+1)
	for kk, vv := range m {
		out[kk] = vv
	}
	out[k] = v
	return out
}

// Find returns the value for key k and whether it was present.
// Mirrors the real slot-table lookup semantics.
func (m ImmutableMap[T]) Find(k string) (T, bool) {
	v, ok := m[k]
	return v, ok
}
