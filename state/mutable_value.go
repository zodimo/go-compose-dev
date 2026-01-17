package state

import (
	"fmt"
	"reflect"
	"sync"
)

var _ MutableValue = &mutableValue{}
var _ StateChangeNotifier = &mutableValue{}

var _ MutableValue = &MutableValueWrapper[any]{}
var _ StateChangeNotifier = &MutableValueWrapper[any]{}

// MutableValue is a state container that notifies subscribers when its value changes.
type mutableValue struct {
	cell           any
	changeNotifier func(any)
	mu             sync.RWMutex // RWMutex for thread-safe access (following go-frp Behavior pattern)
	compare        func(any, any) bool

	// Subscription support for push-based invalidation
	subscribers *SubscriptionManager
}

func NewMutableValue(initial any, changeNotifier func(any), compare func(any, any) bool) MutableValue {

	if changeNotifier == nil {
		changeNotifier = func(any) {}
	}

	if compare == nil {
		compare = func(v1, v2 any) bool {
			return reflect.DeepEqual(v1, v2)
		}
	}

	return &mutableValue{
		cell:           initial,
		changeNotifier: changeNotifier,
		compare:        compare,
		subscribers:    NewSubscriptionManager(),
	}
}

func (mv *mutableValue) Get() any {
	NotifyRead(mv)
	mv.mu.RLock()
	value := mv.cell
	mv.mu.RUnlock()
	return value
}

func (mv *mutableValue) Set(value any) {
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

// Subscribe registers a callback to be invoked when the value changes.
// Returns a Subscription that can be used to stop receiving notifications.
func (mv *mutableValue) Subscribe(callback func()) Subscription {
	return mv.subscribers.Subscribe(callback)
}

// Wrapper

type MutableValueWrapper[T any] struct {
	mv *mutableValueTyped[T]
}

func MutableValueTypedToUntyped[T any](mv MutableValueTyped[T]) MutableValue {
	mvTyped, ok := mv.(*mutableValueTyped[T])
	if !ok {
		panic(fmt.Sprintf("MutableValueTypedToUntyped: expected *MutableValueTyped[T], got %T", mv))
	}

	return &MutableValueWrapper[T]{
		mv: mvTyped,
	}
}

func (w *MutableValueWrapper[T]) Get() any {
	return w.mv.Get()
}

func (w *MutableValueWrapper[T]) Set(value any) {
	tVal, ok := value.(T)
	if !ok {
		var zero T
		panic(fmt.Sprintf("MutableValueWrapper: Set expected value of type %T, got %T", zero, value))
	}
	w.mv.Set(tVal)
}

func (w *MutableValueWrapper[T]) Subscribe(callback func()) Subscription {
	return w.mv.Subscribe(callback)
}
