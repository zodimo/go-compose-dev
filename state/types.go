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

type Value interface {
	StateChangeNotifier
	Get() any
}

type ValueTyped[T any] interface {
	StateChangeNotifier
	Get() T
}

type MutableValue interface {
	Value
	Set(value any)
	CompareAndSet(expect, update any) bool
	Update(func(any) any)
	UpdateAndGet(func(any) any) any
	GetAndUpdate(func(any) any) any
}

type MutableValueTyped[T any] interface {
	ValueTyped[T]
	Set(value T)
	CompareAndSet(expect, update T) bool
	Update(func(T) T)
	UpdateAndGet(func(T) T) T
	GetAndUpdate(func(T) T) T

	Unwrap() MutableValue
}
