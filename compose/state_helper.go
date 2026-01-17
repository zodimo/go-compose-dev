package compose

import "github.com/zodimo/go-compose/state"

// local alias for state.Remember
func Remember[T any](c state.SupportState, key string, calc func() T) (T, error) {
	return state.Remember(c, key, calc)
}

// local alias for state.MustRemember
func MustRemember[T any](c state.SupportState, key string, calc func() T) T {
	return state.MustRemember(c, key, calc)
}

// local alias for state.State
func State[T any](c state.SupportState, key string, initial func() T, options ...state.StateTypedOption[T]) (state.MutableValueTyped[T], error) {
	return state.State[T](c, key, initial, options...)
}

// local alias for state.MustState
func MustState[T any](c state.SupportState, key string, initial func() T, options ...state.StateTypedOption[T]) state.MutableValueTyped[T] {
	return state.MustState[T](c, key, initial, options...)
}
