package store

import (
	"fmt"

	"github.com/zodimo/go-compose/state"
)

func Remember[T any](c state.SupportState, key string, calc func() T) (T, error) {
	anyCalc := func() any { return calc() }
	anyValue := c.Remember(key, anyCalc)

	tValue, ok := anyValue.(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("value is not of type %T", zero)
	}
	return tValue, nil
}

func RememberUnsafe[T any](c state.SupportState, key string, calc func() T) T {
	anyCalc := func() any { return calc() }
	anyValue := c.Remember(key, anyCalc)

	tValue, ok := anyValue.(T)
	if !ok {
		var zero T
		panic(fmt.Errorf("value is not of type %T", zero))
	}
	return tValue
}

func MustRemember[T any](c state.SupportState, key string, calc func() T) T {
	return RememberUnsafe[T](c, key, calc)
}

func State[T any](c state.SupportState, key string, initial func() T) (TypedMutableValueInterface[T], error) {
	mv := c.State(key, func() any { return initial() })
	anyMv, ok := mv.(*MutableValue)
	if !ok {
		return nil, fmt.Errorf("mutable value is not of type %T", MutableValue{})
	}
	return WrapMutableValue[T](anyMv)
}

func StateUnsafe[T any](c state.SupportState, key string, initial func() T) TypedMutableValueInterface[T] {
	mv := c.State(key, func() any { return initial() })
	anyMv, ok := mv.(*MutableValue)
	if !ok {
		panic(fmt.Errorf("mutable value is not of type %T", MutableValue{}))
	}
	wrapped, err := WrapMutableValue[T](anyMv)
	if err != nil {
		panic(err)
	}
	return wrapped
}

func MustState[T any](c state.SupportState, key string, initial func() T) TypedMutableValueInterface[T] {
	return StateUnsafe[T](c, key, initial)
}
