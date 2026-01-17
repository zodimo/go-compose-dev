package state

import (
	"sync"
)

var _ MutableValue = (*mutableValue)(nil)

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
