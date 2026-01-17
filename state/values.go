package state

type Value interface {
	Get() any
	Subscribe(callback func()) Subscription
}

type TypedValue[T any] interface {
	Get() T
	Subscribe(callback func()) Subscription
}

type MutableValue interface {
	Value
	Set(value any)
}

type TypedMutableValue[T any] interface {
	TypedValue[T]
	Set(value T)

	Unwrap() MutableValue
}
