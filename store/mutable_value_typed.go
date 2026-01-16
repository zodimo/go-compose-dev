package store

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/zodimo/go-compose/state"
)

type TypedMutableValueInterface[T any] = state.TypedMutableValue[T]

var _ TypedMutableValueInterface[any] = &MutableValueTypedWrapper[any]{}
var _ MutableValueInterface = &MutableValueTypedWrapper[any]{}
var _ state.StateChangeNotifier = &MutableValueTypedWrapper[any]{}

var _ TypedMutableValueInterface[any] = &MutableValueTyped[any]{}
var _ MutableValueInterface = &MutableValueTyped[any]{}
var _ state.StateChangeNotifier = &MutableValueTyped[any]{}

// MutableValue is a state container that notifies subscribers when its value changes.
type MutableValueTyped[T any] struct {
	cell           T
	changeNotifier func(T)
	mu             sync.RWMutex // RWMutex for thread-safe access (following go-frp Behavior pattern)
	compare        func(T, T) bool

	// Subscription support for push-based invalidation
	subscribers *state.SubscriptionManager
}

func NewMutableValueTyped[T any](initial T, changeNotifier func(T), compare func(T, T) bool) *MutableValueTyped[T] {

	if changeNotifier == nil {
		changeNotifier = func(T) {}
	}

	if compare == nil {
		compare = func(t1, t2 T) bool {
			return reflect.DeepEqual(t1, t2)
		}
	}
	return &MutableValueTyped[T]{
		cell:           initial,
		changeNotifier: changeNotifier,
		compare:        compare,
		subscribers:    state.NewSubscriptionManager(),
	}
}

func (mv *MutableValueTyped[T]) Get() T {
	state.NotifyRead(mv)
	mv.mu.RLock()
	value := mv.cell
	mv.mu.RUnlock()
	return value
}

func (mv *MutableValueTyped[T]) Set(value T) {
	mv.mu.Lock()
	changed := !mv.compare(mv.cell, value)
	if changed {
		mv.cell = value
	}
	changeNotifier := mv.changeNotifier
	mv.mu.Unlock()

	if changed {
		// Notify legacy change notifier
		if changeNotifier != nil {
			changeNotifier(value)
		}

		// Notify all subscribers (push invalidation to derived states)
		mv.subscribers.NotifyAll()
	}
}

func (mv *MutableValueTyped[T]) Unwrap() MutableValueInterface {
	return MutableValueTypedToUntyped(mv)
}

// Subscribe registers a callback to be invoked when the value changes.
// Returns a Subscription that can be used to stop receiving notifications.
func (mv *MutableValueTyped[T]) Subscribe(callback func()) state.Subscription {
	return mv.subscribers.Subscribe(callback)
}

// Wrapper

type MutableValueTypedWrapper[T any] struct {
	mv *MutableValue
}

func MutableValueToTyped[T any](mv *MutableValue) (TypedMutableValueInterface[T], error) {

	_, ok := mv.cell.(T)
	if !ok {
		var zero T
		return nil, fmt.Errorf("cell is not of type %T, got %T", zero, mv.cell)
	}

	return &MutableValueTypedWrapper[T]{
		mv: mv,
	}, nil
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

func (w *MutableValueTypedWrapper[T]) Unwrap() MutableValueInterface {
	return w.mv
}
