package state

import (
	"fmt"
	"reflect"
	"sync"
)

type TypedMutableValueInterface[T any] = MutableValueTyped[T]

var _ TypedMutableValueInterface[any] = &MutableValueTypedWrapper[any]{}
var _ MutableValue = &MutableValueTypedWrapper[any]{}
var _ StateChangeNotifier = &MutableValueTypedWrapper[any]{}

var _ TypedMutableValueInterface[any] = &mutableValueTyped[any]{}
var _ MutableValue = &mutableValueTyped[any]{}
var _ StateChangeNotifier = &mutableValueTyped[any]{}

// MutableValue is a state container that notifies subscribers when its value changes.
type mutableValueTyped[T any] struct {
	cell           T
	changeNotifier func(T)
	mu             sync.RWMutex // RWMutex for thread-safe access (following go-frp Behavior pattern)
	compare        func(T, T) bool

	// Subscription support for push-based invalidation
	subscribers *SubscriptionManager
}

func NewMutableValueTyped[T any](initial T, changeNotifier func(T), compare func(T, T) bool) MutableValueTyped[T] {

	if changeNotifier == nil {
		changeNotifier = func(T) {}
	}

	if compare == nil {
		compare = func(t1, t2 T) bool {
			return reflect.DeepEqual(t1, t2)
		}
	}
	return &mutableValueTyped[T]{
		cell:           initial,
		changeNotifier: changeNotifier,
		compare:        compare,
		subscribers:    NewSubscriptionManager(),
	}
}

func (mv *mutableValueTyped[T]) Get() T {
	NotifyRead(mv)
	mv.mu.RLock()
	value := mv.cell
	mv.mu.RUnlock()
	return value
}

func (mv *mutableValueTyped[T]) Set(value T) {
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

func (mv *mutableValueTyped[T]) Unwrap() MutableValue {
	return MutableValueTypedToUntyped(mv)
}

// Subscribe registers a callback to be invoked when the value changes.
// Returns a Subscription that can be used to stop receiving notifications.
func (mv *mutableValueTyped[T]) Subscribe(callback func()) Subscription {
	return mv.subscribers.Subscribe(callback)
}

// Wrapper

type MutableValueTypedWrapper[T any] struct {
	mv *mutableValue
}

func MutableValueToTyped[T any](mv MutableValue) (MutableValueTyped[T], error) {
	mvTyped, ok := mv.(*mutableValue)
	if !ok {
		return nil, fmt.Errorf("cell is not of type %T, got %T", mvTyped, mv)
	}

	_, ok = mvTyped.cell.(T)
	if !ok {
		var zero T
		return nil, fmt.Errorf("cell is not of type %T, got %T", zero, mvTyped.cell)
	}

	return &MutableValueTypedWrapper[T]{
		mv: mvTyped,
	}, nil
}

func (w *MutableValueTypedWrapper[T]) Get() T {
	return w.mv.Get().(T)
}

func (w *MutableValueTypedWrapper[T]) Set(value T) {
	w.mv.Set(value)
}

func (w *MutableValueTypedWrapper[T]) Subscribe(callback func()) Subscription {
	return w.mv.Subscribe(callback)
}

func (w *MutableValueTypedWrapper[T]) Unwrap() MutableValue {
	return w.mv
}
