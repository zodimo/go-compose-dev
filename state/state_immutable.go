package state

import "github.com/zodimo/go-maybe"

// Higher-Kinded Type
// EndGame - use it
type ImmutableState[T, U any] interface {
	Get(key string) maybe.Maybe[U]
	Set(key string, value U) ImmutableState[T, U]
}
