package flow

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/zodimo/go-compose/state"
)

// StateFlow defines the read-only behavior
type StateFlow[T any] interface {
	Value() T
	Flow[T]
}

// Compile-time check that MutableStateFlow implements state.StateChangeNotifier
var _ state.StateChangeNotifier = (*MutableStateFlow[any])(nil)

// StateFlowOption configures a MutableStateFlow
type StateFlowOption[T any] func(*stateFlowConfig[T])

// stateFlowConfig holds configuration for MutableStateFlow
type stateFlowConfig[T any] struct {
	policy state.MutationPolicy[T]
}

// WithPolicy sets a custom mutation policy for the StateFlow.
// If not set, StructuralEqualityPolicy is used by default.
func WithPolicy[T any](policy state.MutationPolicy[T]) StateFlowOption[T] {
	return func(c *stateFlowConfig[T]) {
		c.policy = policy
	}
}

// MutableStateFlow is the hot, stateful producer.
// It matches Kotlin's MutableStateFlow semantics:
// - Equality-based conflation: updates with equal values are ignored
// - Thread-safe atomic operations
// - CAS-based update methods
type MutableStateFlow[T any] struct {
	mu          sync.RWMutex
	value       atomic.Value
	subscribers []chan T
	policy      state.MutationPolicy[T]

	// Subscription manager for state change notifications (push-based invalidation)
	stateSubscribers *state.SubscriptionManager
}

// NewMutableStateFlow creates a new MutableStateFlow with the given initial value.
// Options can be used to customize behavior (e.g., WithPolicy for custom equality).
func NewMutableStateFlow[T any](initial T, opts ...StateFlowOption[T]) *MutableStateFlow[T] {
	// Apply configuration
	config := &stateFlowConfig[T]{
		policy: state.StructuralEqualityPolicy[T](),
	}
	for _, opt := range opts {
		opt(config)
	}

	flow := &MutableStateFlow[T]{
		subscribers:      make([]chan T, 0),
		stateSubscribers: state.NewSubscriptionManager(),
		policy:           config.policy,
	}
	flow.value.Store(initial)
	return flow
}

// Value returns the current state (Thread-safe).
// Also notifies the read observer to enable derived state dependency tracking.
func (s *MutableStateFlow[T]) Value() T {
	state.NotifyRead(s)
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.value.Load().(T)
}

// CompareAndSet atomically compares the current value with expect and sets it to update if equal.
// Returns true if the value was set to update, false otherwise.
// If both expect and update equal the current value, returns true but does not notify subscribers.
// This matches Kotlin's MutableStateFlow.compareAndSet semantics.
func (s *MutableStateFlow[T]) CompareAndSet(expect, update T) bool {
	s.mu.Lock()
	current := s.value.Load().(T)

	// Check if current matches expected
	if !s.policy.Equivalent(current, expect) {
		s.mu.Unlock()
		return false
	}

	// Check if update is same as current (no change needed)
	if s.policy.Equivalent(current, update) {
		s.mu.Unlock()
		return true // CAS succeeded but no notification needed
	}

	// Perform the update
	s.value.Store(update)
	subs := make([]chan T, len(s.subscribers))
	copy(subs, s.subscribers)
	s.mu.Unlock()

	s.notifySubscribers(update, subs)
	s.stateSubscribers.NotifyAll()
	return true
}

// Emit updates the value and notifies all collectors and state subscribers.
// If the new value equals the current value (using the comparator), no notification occurs.
// This matches Kotlin's equality-based conflation behavior.
func (s *MutableStateFlow[T]) Emit(value T) {
	s.mu.Lock()
	current := s.value.Load().(T)

	// Equality-based conflation: skip if value hasn't changed
	if s.policy.Equivalent(current, value) {
		s.mu.Unlock()
		return
	}

	s.value.Store(value)
	subs := make([]chan T, len(s.subscribers))
	copy(subs, s.subscribers)
	s.mu.Unlock()
	s.notifySubscribers(value, subs)
	s.stateSubscribers.NotifyAll()
}

func (s *MutableStateFlow[T]) notifySubscribers(value T, subscribers []chan T) {
	for _, ch := range subscribers {
		// Non-blocking send (Conflation)
		// If the subscriber is slow, we drain the old value and send the new one
		select {
		case ch <- value:
		default:
			// Buffer is full (slow collector); drain the old and replace with new
			select {
			case <-ch:
			default:
			}
			ch <- value
		}
	}

}

// Collect follows the Kotlin pattern: it blocks until the context is cancelled
func (s *MutableStateFlow[T]) Collect(ctx context.Context, collector func(T)) error {
	ch := make(chan T, 1)

	s.mu.Lock()
	// Capture the value AND register the channel in one atomic step
	current := s.value.Load().(T)
	s.subscribers = append(s.subscribers, ch)
	s.mu.Unlock()

	collector(current)

	defer func() {
		s.mu.Lock()
		for i, sub := range s.subscribers {
			if sub == ch {
				s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
				break
			}
		}
		s.mu.Unlock()
		close(ch)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case val := <-ch:
			collector(val)
		}
	}
}

// Update atomically updates the value using the given function.
// Uses a CAS loop internally, so the function may be called multiple times
// if there are concurrent modifications.
// This matches Kotlin's MutableStateFlow.update extension function.
func (s *MutableStateFlow[T]) Update(f func(current T) T) {
	for {
		current := s.Value()
		newValue := f(current)
		if s.CompareAndSet(current, newValue) {
			return
		}
	}
}

// UpdateAndGet atomically updates the value and returns the new value.
// Uses a CAS loop internally, so the function may be called multiple times.
// This matches Kotlin's MutableStateFlow.updateAndGet extension function.
func (s *MutableStateFlow[T]) UpdateAndGet(f func(current T) T) T {
	for {
		current := s.Value()
		newValue := f(current)
		if s.CompareAndSet(current, newValue) {
			return newValue
		}
	}
}

// GetAndUpdate atomically updates the value and returns the previous value.
// Uses a CAS loop internally, so the function may be called multiple times.
// This matches Kotlin's MutableStateFlow.getAndUpdate extension function.
func (s *MutableStateFlow[T]) GetAndUpdate(f func(current T) T) T {
	for {
		current := s.Value()
		newValue := f(current)
		if s.CompareAndSet(current, newValue) {
			return current
		}
	}
}

func (s *MutableStateFlow[T]) AsStateFlow() StateFlow[T] {
	return s
}

// Subscribe registers a callback to be invoked when the flow's value changes.
// This implements state.StateChangeNotifier, enabling MutableStateFlow to be
// used as a dependency for DerivedState.
func (s *MutableStateFlow[T]) Subscribe(callback func()) state.Subscription {
	return s.stateSubscribers.Subscribe(callback)
}
