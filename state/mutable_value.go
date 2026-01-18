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
	cell    any
	mu      sync.RWMutex // RWMutex for thread-safe access (following go-frp Behavior pattern)
	compare func(any, any) bool

	// Subscription support for push-based invalidation
	subscribers *SubscriptionManager
}

func NewMutableValue(initial any, compare func(any, any) bool) MutableValue {

	if compare == nil {
		compare = func(v1, v2 any) bool {
			return reflect.DeepEqual(v1, v2)
		}
	}

	return &mutableValue{
		cell:        initial,
		compare:     compare,
		subscribers: NewSubscriptionManager(),
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
	mv.mu.Unlock()

	if changed {
		// Notify all subscribers (push invalidation to derived states)
		mv.subscribers.NotifyAll()
	}
}

func (mv *mutableValue) CompareAndSet(expect, update any) bool {
	mv.mu.Lock()
	current := mv.cell

	// Check if current matches expected
	if !mv.compare(current, expect) {
		mv.mu.Unlock()
		return false
	}

	// Check if update is same as current (no change needed)
	if mv.compare(current, update) {
		mv.mu.Unlock()
		return true // CAS succeeded but no notification needed
	}

	// Perform the update
	mv.cell = update
	mv.mu.Unlock()

	// Notify all subscribers
	mv.subscribers.NotifyAll()
	return true
}

func (mv *mutableValue) Update(f func(any) any) {
	for {
		current := mv.Get()
		newValue := f(current)
		if mv.CompareAndSet(current, newValue) {
			return
		}
	}
}

// update then get, return new
func (mv *mutableValue) UpdateAndGet(f func(any) any) any {
	for {
		current := mv.Get()
		newValue := f(current)
		if mv.CompareAndSet(current, newValue) {
			return newValue
		}
	}
}

// get then update, return old
func (mv *mutableValue) GetAndUpdate(f func(any) any) any {
	for {
		current := mv.Get()
		newValue := f(current)
		if mv.CompareAndSet(current, newValue) {
			return current
		}
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

func (w *MutableValueWrapper[T]) CompareAndSet(expect, update any) bool {
	tExpect, ok := expect.(T)
	if !ok {
		var zero T
		panic(fmt.Sprintf("MutableValueWrapper: CompareAndSet expected expect value of type %T, got %T", zero, expect))
	}
	tUpdate, ok := update.(T)
	if !ok {
		var zero T
		panic(fmt.Sprintf("MutableValueWrapper: CompareAndSet expected update value of type %T, got %T", zero, update))
	}
	return w.mv.CompareAndSet(tExpect, tUpdate)
}

func (w *MutableValueWrapper[T]) Update(f func(any) any) {
	w.mv.Update(func(current T) T {
		res := f(current)
		tRes, ok := res.(T)
		if !ok {
			var zero T
			panic(fmt.Sprintf("MutableValueWrapper: Update function returned value of type %T, expected %T", res, zero))
		}
		return tRes
	})
}

func (w *MutableValueWrapper[T]) UpdateAndGet(f func(any) any) any {
	return w.mv.UpdateAndGet(func(current T) T {
		res := f(current)
		tRes, ok := res.(T)
		if !ok {
			var zero T
			panic(fmt.Sprintf("MutableValueWrapper: UpdateAndGet function returned value of type %T, expected %T", res, zero))
		}
		return tRes
	})
}

func (w *MutableValueWrapper[T]) GetAndUpdate(f func(any) any) any {
	return w.mv.GetAndUpdate(func(current T) T {
		res := f(current)
		tRes, ok := res.(T)
		if !ok {
			var zero T
			panic(fmt.Sprintf("MutableValueWrapper: GetAndUpdate function returned value of type %T, expected %T", res, zero))
		}
		return tRes
	})
}
