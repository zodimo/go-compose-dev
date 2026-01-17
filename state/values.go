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
}

type MutableValueTyped[T any] interface {
	ValueTyped[T]
	Set(value T)

	Unwrap() MutableValue
}
