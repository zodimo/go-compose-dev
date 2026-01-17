package state

import (
	"fmt"
	"reflect"
)

func Remember[T any](c SupportState, key string, calc func() T) (T, error) {
	anyCalc := func() any { return calc() }
	anyValue := c.Remember(key, anyCalc)

	tValue, ok := anyValue.(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("value is not of type %T", zero)
	}
	return tValue, nil
}

func RememberUnsafe[T any](c SupportState, key string, calc func() T) T {
	anyCalc := func() any { return calc() }
	anyValue := c.Remember(key, anyCalc)

	tValue, ok := anyValue.(T)
	if !ok {
		var zero T
		panic(fmt.Errorf("value is not of type %T", zero))
	}
	return tValue
}

func MustRemember[T any](c SupportState, key string, calc func() T) T {
	return RememberUnsafe[T](c, key, calc)
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

func State[T any](c SupportState, key string, initial func() T, options ...StateTypedOption[T]) (TypedMutableValueInterface[T], error) {
	opts := StateTypedOptions[T]{
		Compare: func(t1, t2 T) bool {
			return reflect.DeepEqual(t1, t2)
		},
	}
	for _, option := range options {
		option(&opts)
	}

	mv := c.State(key, func() any { return initial() }, WithCompare(func(a, b any) bool {
		return opts.Compare(a.(T), b.(T))
	}))
	anyMv, ok := mv.(*mutableValue)
	if !ok {
		return nil, fmt.Errorf("mutable value is not of type %T", mutableValue{})
	}
	return MutableValueToTyped[T](anyMv)
}

func StateUnsafe[T any](c SupportState, key string, initial func() T, options ...StateTypedOption[T]) TypedMutableValueInterface[T] {
	mv, err := State[T](c, key, initial, options...)
	if err != nil {
		panic(err)
	}
	return mv
}

func MustState[T any](c SupportState, key string, initial func() T, options ...StateTypedOption[T]) TypedMutableValueInterface[T] {
	return StateUnsafe[T](c, key, initial, options...)
}
