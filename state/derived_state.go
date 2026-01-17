package state

import (
	"reflect"
	"sync"
	"sync/atomic"
)

var _ ValueTyped[any] = (*DerivedState[any])(nil)
var _ StateChangeNotifier = (*DerivedState[any])(nil)
var _ InvalidationNotifier = (*DerivedState[any])(nil)

// DerivedState computes a value from other state objects and caches the result.
// It uses push-based invalidation: when dependencies change, the derived state
// is marked invalid but not recalculated until Get() is called.
type DerivedState[T any] struct {
	calculation func() T
	value       T
	compare     func(T, T) bool

	// Push-based invalidation
	invalid       atomic.Bool
	initialized   atomic.Bool
	subscriptions []Subscription
	subMutex      sync.Mutex

	// Value access protection (following go-frp Behavior pattern)
	valueMu sync.RWMutex

	// Prevent concurrent recalculations
	recalcMu sync.Mutex

	// Internal invalidation propagation - notifies downstream derived states
	// that they need to recalculate. This happens during invalidation.
	invalidationSubs *SubscriptionManager

	// External value-change notifications - notifies user callbacks only
	// when the computed value actually changes. NOT called during invalidation.
	subscribers *SubscriptionManager
}

// DerivedStateOf creates a new DerivedState with the given calculation function.
// Uses reflect.DeepEqual for value comparison by default.
func DerivedStateOf[T any](calculation func() T) *DerivedState[T] {
	ds := &DerivedState[T]{
		calculation:      calculation,
		compare:          func(a, b T) bool { return reflect.DeepEqual(a, b) },
		invalidationSubs: NewSubscriptionManager(),
		subscribers:      NewSubscriptionManager(),
	}
	ds.invalid.Store(true) // Start as invalid to trigger initial calculation
	return ds
}

// DerivedStateOfCustom creates a new DerivedState with a custom comparison function.
func DerivedStateOfCustom[T any](calculation func() T, compare func(T, T) bool) *DerivedState[T] {
	ds := &DerivedState[T]{
		calculation:      calculation,
		compare:          compare,
		invalidationSubs: NewSubscriptionManager(),
		subscribers:      NewSubscriptionManager(),
	}
	ds.invalid.Store(true)
	return ds
}

// DerivedStateWithPolicy creates a new DerivedState with a custom EqualityPolicy.
func DerivedStateWithPolicy[T any](calculation func() T, policy EqualityPolicy[T]) *DerivedState[T] {
	return DerivedStateOfCustom(calculation, policy.Equivalent)
}

// Get returns the cached value, recalculating if invalid.
func (ds *DerivedState[T]) Get() T {
	// Check if we need to recalculate (push-based: invalid flag was set by dependency change)
	if !ds.initialized.Load() || ds.invalid.Load() {
		ds.recalculate()
	}

	NotifyRead(ds)

	// Read value with lock protection
	ds.valueMu.RLock()
	value := ds.value
	ds.valueMu.RUnlock()
	return value
}

// invalidate marks this derived state as needing recalculation.
// Called by dependency subscriptions when their values change.
//
// This method propagates invalidation through the chain via invalidationSubs.
// For example, if A -> B -> C:
// 1. A changes → B.invalidate() called → B marked invalid
// 2. B.invalidationSubs notifies C.invalidate() → C marked invalid
// 3. When C.Get() is called, it recalculates (calling B.Get() which also recalculates)
//
// Note: Only invalidationSubs is notified here (for downstream derived states).
// User-registered subscribers (via Subscribe()) are NOT notified during invalidation;
// they are only notified in recalculate() if the value actually changes.
func (ds *DerivedState[T]) invalidate() {
	// Only invalidate if not already invalid (prevents infinite loops in diamond deps)
	if ds.invalid.CompareAndSwap(false, true) {
		// Propagate invalidation to downstream derived states only
		ds.invalidationSubs.NotifyAll()

		// If there are direct subscribers to this derived state, they expect
		// value change notifications. We must eagerly recalculate to determine
		// if the value actually changed and notify them.
		if ds.subscribers.Count() > 0 {
			ds.recalculate()
		}
	}
}

// unsubscribeFromDeps removes all subscriptions to dependencies.
func (ds *DerivedState[T]) unsubscribeFromDeps() {
	ds.subMutex.Lock()
	defer ds.subMutex.Unlock()

	for _, sub := range ds.subscriptions {
		sub.Unsubscribe()
	}
	ds.subscriptions = nil
}

// recalculate re-executes the calculation and updates dependencies.
// Thread-safe: uses recalcMu to serialize concurrent recalculations.
func (ds *DerivedState[T]) recalculate() {
	// Serialize recalculations to prevent races when multiple goroutines
	// call Get() simultaneously on an invalid derived state
	ds.recalcMu.Lock()
	defer ds.recalcMu.Unlock()

	// Double-check: another goroutine may have already recalculated
	if ds.initialized.Load() && !ds.invalid.Load() {
		return
	}

	// Unsubscribe from old dependencies
	ds.unsubscribeFromDeps()

	var newSubs []Subscription

	// Track reads during calculation
	var newValue T
	WithReadObserver(func(source StateChangeNotifier) {
		// For derived states, subscribe to invalidation events (chain propagation).
		// For other state types, subscribe to value-change events.
		var sub Subscription
		if inv, ok := source.(InvalidationNotifier); ok {
			sub = inv.SubscribeForInvalidation(ds.invalidate)
		} else {
			sub = source.Subscribe(ds.invalidate)
		}
		newSubs = append(newSubs, sub)
	}, func() {
		newValue = ds.calculation()
	})

	// Store new subscriptions
	ds.subMutex.Lock()
	ds.subscriptions = newSubs
	ds.subMutex.Unlock()

	// Compare and update value with lock protection
	ds.valueMu.Lock()
	valueChanged := !ds.initialized.Load() || !ds.compare(ds.value, newValue)
	if valueChanged {
		ds.value = newValue
	}
	ds.valueMu.Unlock()

	ds.invalid.Store(false)
	ds.initialized.Store(true)

	// Notify user subscribers AFTER recalculation, only if value changed
	// This enables nested caching: user callbacks won't be notified
	// if the computed value didn't actually change.
	if valueChanged {
		ds.subscribers.NotifyAll()
	}
}

// Subscribe registers a callback to be invoked when this derived state's
// computed value actually changes (not just when it's invalidated).
// This is the user-facing subscription for value-change notifications.
func (ds *DerivedState[T]) Subscribe(callback func()) Subscription {
	// User callbacks go to subscribers (value-change notifications)
	// NOT to invalidationSubs (internal invalidation propagation)
	return ds.subscribers.Subscribe(callback)
}

// SubscribeForInvalidation registers a callback to be invoked when this
// derived state is invalidated. Used internally for derived state chains.
// Implements InvalidationNotifier interface.
func (ds *DerivedState[T]) SubscribeForInvalidation(callback func()) Subscription {
	return ds.invalidationSubs.Subscribe(callback)
}

// IsInvalid returns true if the derived state needs recalculation.
// Primarily for testing and debugging.
func (ds *DerivedState[T]) IsInvalid() bool {
	return ds.invalid.Load()
}
