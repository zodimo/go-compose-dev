package store

import (
	"sync"

	"github.com/zodimo/go-compose/state"
)

// MutableValue is a state container that notifies subscribers when its value changes.
type MutableValue struct {
	cell           any
	changeNotifier func(any)
	mu             sync.RWMutex // RWMutex for thread-safe access (following go-frp Behavior pattern)
	compare        func(any, any) bool

	// Subscription support for push-based invalidation
	subscribers *state.SubscriptionManager
}

func NewMutableValue(initial any, changeNotifier func(any), compare func(any, any) bool) *MutableValue {
	return &MutableValue{
		cell:           initial,
		changeNotifier: changeNotifier,
		compare:        compare,
		subscribers:    state.NewSubscriptionManager(),
	}
}

func (mv *MutableValue) Get() any {
	state.NotifyRead(mv)
	mv.mu.RLock()
	value := mv.cell
	mv.mu.RUnlock()
	return value
}

func (mv *MutableValue) Set(value any) {
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
func (mv *MutableValue) Subscribe(callback func()) state.Subscription {
	return mv.subscribers.Subscribe(callback)
}
