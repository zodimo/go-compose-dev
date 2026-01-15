package store

import (
	"fmt"
	"reflect"

	"github.com/zodimo/go-compose/state"
)

type MutableValueInterface = state.MutableValue
type PersistentStateInterface = state.PersistentState
type TypedMutableValueInterface[T any] = state.TypedMutableValue[T]

var _ MutableValueInterface = &MutableValue{}
var _ MutableValueInterface = &MutableValueTypedWrapper[any]{}

var _ state.StateChangeNotifier = &MutableValue{}
var _ state.StateChangeNotifier = &MutableValueTypedWrapper[any]{}

var _ TypedMutableValueInterface[any] = &MutableValueTypedWrapper[any]{}

type PersistentState struct {
	scopes        map[string]MutableValueInterface
	onStateChange func()
}

func NewPersistentState(scopes map[string]MutableValueInterface) PersistentStateInterface {
	return &PersistentState{scopes: scopes}
}

func (ps *PersistentState) SetOnStateChange(callback func()) {
	ps.onStateChange = callback
}

func (ps *PersistentState) GetState(id string, initial func() any, options ...state.StateOption) MutableValueInterface {

	opts := state.StateOptions{
		Compare: reflect.DeepEqual,
	}
	for _, option := range options {
		option(&opts)
	}

	if v, ok := ps.scopes[id]; ok {
		return v
	}
	ps.scopes[id] = &MutableValue{
		cell: initial(),
		changeNotifier: func(any) {
			if ps.onStateChange != nil {
				ps.onStateChange()
			}
		},
		compare:     opts.Compare,
		subscribers: state.NewSubscriptionManager(),
	}
	return ps.scopes[id]
}

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
