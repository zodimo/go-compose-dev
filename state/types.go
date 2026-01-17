package state

type StateOption func(*StateOptions)

type StateOptions struct {
	Compare func(any, any) bool
}

type StateTypedOptions[T any] struct {
	Compare func(T, T) bool
}

type StateTypedOption[T any] func(*StateTypedOptions[T])

func WithTypedCompare[T any](compare func(T, T) bool) StateTypedOption[T] {
	if compare == nil {
		panic("compare cannot be nil")
	}
	return func(o *StateTypedOptions[T]) {
		o.Compare = compare
	}
}

type SupportState interface {
	Remember(key string, calc func() any) any                                  // transient state
	State(key string, initial func() any, options ...StateOption) MutableValue // persistent state
}
