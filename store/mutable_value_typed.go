package store

import (
	"fmt"

	"github.com/zodimo/go-compose/state"
)

type MutableValueTyped[T any] struct {
	cell           T
	changeNotifier func(T)
	compare        func(T, T) bool
}

func NewMutableValueTyped[T any](initial T, changeNotifier func(T), compare func(T, T) bool) *MutableValueTyped[T] {
	return &MutableValueTyped[T]{
		cell:           initial,
		changeNotifier: changeNotifier,
		compare:        compare,
	}
}

type MutableValueTypedWrapper[T any] struct {
	mv *MutableValue
}

func (w *MutableValueTypedWrapper[T]) Get() T {
	return w.mv.Get().(T)
}

func (w *MutableValueTypedWrapper[T]) Set(value T) {
	w.mv.Set(value)
}

func (w *MutableValueTypedWrapper[T]) Subscribe(callback func()) state.Subscription {
	return w.mv.Subscribe(callback)
}

func WrapMutableValue[T any](mv *MutableValue) (TypedMutableValueInterface[T], error) {

	_, ok := mv.cell.(T)
	if !ok {
		var zero T
		return nil, fmt.Errorf("cell is not of type %T, got %T", zero, mv.cell)
	}

	return &MutableValueTypedWrapper[T]{
		mv: mv,
	}, nil
}

func (w *MutableValueTypedWrapper[T]) Unwrap() MutableValueInterface {
	return w.mv
}
