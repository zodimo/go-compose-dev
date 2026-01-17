package state

type Value interface {
	Get() any
	Subscribe(callback func()) Subscription
}

type ValueTyped[T any] interface {
	Get() T
	Subscribe(callback func()) Subscription
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
