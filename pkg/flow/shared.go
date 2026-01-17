package flow

import (
	"context"
	"sync"
)

// SharedFlow is a hot flow that shares emitted values among all collectors in a broadcast fashion.
type SharedFlow[T any] interface {
	Flow[T]
	// ReplayCache returns the current replay cache.
	ReplayCache() []T
}

// MutableSharedFlow is a SharedFlow that can also emit values.
type MutableSharedFlow[T any] interface {
	SharedFlow[T]
	// Emit emits a value to the shared flow.
	Emit(value T)
	// TryEmit attempts to emit a value without suspending. Returns true if successful.
	TryEmit(value T) bool
	// SubscriptionCount returns a StateFlow that tracks the number of active subscribers.
	SubscriptionCount() StateFlow[int]
	// ResetReplayCache resets the replay cache to an empty state.
	ResetReplayCache()
}

// implementation constants
const (
	defaultExtraBufferCapacity = 0 // Default extra buffer (on top of replay)
)

type sharedFlowImpl[T any] struct {
	mu                  sync.RWMutex
	replay              int
	extraBufferCapacity int
	onBufferOverflow    BufferOverflow

	replayCache []T
	subscribers map[chan T]struct{}

	subscriptionCount *MutableStateFlow[int]
}

// NewMutableSharedFlow creates a new MutableSharedFlow.
func NewMutableSharedFlow[T any](
	replay int,
	extraBufferCapacity int,
	onBufferOverflow BufferOverflow,
) MutableSharedFlow[T] {
	// Validate inputs matching Kotlin's require() semantics
	if replay < 0 {
		panic("replay cannot be negative")
	}
	if extraBufferCapacity < 0 {
		panic("extraBufferCapacity cannot be negative")
	}
	// Fallback to Suspend if invalid overflow strategy is somehow provided?
	// Enum is int, so technically possible. But we assume valid enum or handle in switch.

	return &sharedFlowImpl[T]{
		replay:              replay,
		extraBufferCapacity: extraBufferCapacity,
		onBufferOverflow:    onBufferOverflow,
		replayCache:         make([]T, 0, replay),
		subscribers:         make(map[chan T]struct{}),
		subscriptionCount:   NewMutableStateFlow(0, WithoutSubscriptionCount[int]()),
	}
}

func (s *sharedFlowImpl[T]) ReplayCache() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// Return a copy to avoid race conditions
	cache := make([]T, len(s.replayCache))
	copy(cache, s.replayCache)
	return cache
}

func (s *sharedFlowImpl[T]) SubscriptionCount() StateFlow[int] {
	return s.subscriptionCount
}

func (s *sharedFlowImpl[T]) ResetReplayCache() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.replayCache = make([]T, 0, s.replay)
}

func (s *sharedFlowImpl[T]) Collect(ctx context.Context, collector func(T)) error {
	// 1. Subscribe
	// Determine channel capacity: replay + extraBufferCapacity?
	// In Kotlin, the "buffer" is shared. Here, we give each subscriber a channel.
	// If we want to support SUSPEND properly, channel size should probably be 0 or small,
	// and we rely on the sender blocking.
	// But if we want to support replay, we need to send replay values immediately.
	// Let's use a buffer size equal to replay + extraBufferCapacity.
	capacity := s.replay + s.extraBufferCapacity
	if capacity < 0 {
		capacity = 0 // Should not happen based on types, but safecheck
	}

	ch := make(chan T, capacity)

	s.mu.Lock()
	// Send replay cache immediately to the channel
	for _, val := range s.replayCache {
		select {
		case ch <- val:
		default:
			// Buffer full. This can happen if capacity < replay, which shouldn't happen
			// with current sizing logic (capacity = replay + extra, extra >= 0).
			// If it does happen, we drop the oldest replay value for this subscriber.
		}
	}
	s.subscribers[ch] = struct{}{}
	count := len(s.subscribers)
	s.mu.Unlock()

	// Update subscription count
	s.subscriptionCount.Emit(count)

	// 2. Cleanup on exit
	defer func() {
		s.mu.Lock()
		delete(s.subscribers, ch)
		newCount := len(s.subscribers)
		s.mu.Unlock()
		s.subscriptionCount.Emit(newCount)
		close(ch)
	}()

	// 3. Process values
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case val, ok := <-ch:
			if !ok {
				return nil
			}
			collector(val)
		}
	}
}

func (s *sharedFlowImpl[T]) Emit(value T) {
	s.emit(value, false)
}

func (s *sharedFlowImpl[T]) TryEmit(value T) bool {
	// TryEmit in Kotlin returns false if it would suspend.
	// Implementing via common emit helper.
	return s.emit(value, true)
}

func (s *sharedFlowImpl[T]) emit(value T, tryOnly bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 1. Update Replay Cache
	if s.replay > 0 {
		if len(s.replayCache) >= s.replay {
			// Remove oldest (first)
			s.replayCache = s.replayCache[1:]
		}
		s.replayCache = append(s.replayCache, value)
	}

	// 2. Distribute to subscribers
	for ch := range s.subscribers {
		// Attempt to send
		sent := false
		select {
		case ch <- value:
			sent = true
		default:
			// Channel is full. Handle Overflow.
			switch s.onBufferOverflow {
			case BufferOverflowSuspend:
				if tryOnly {
					return false // Would block
				}
				// Must suspend (block) until this subscriber accepts.
				// WARNING: This holds the lock! This blocks ALL subscribers if one is slow.
				// This is how SharedFlow works if you use SUSPEND?
				// Kotlin docs: "The emitter suspends until all subscribers receive the value."
				// Wait, strictly speaking, Kotlin's shared buffer implementation is more clever.
				// Doing this with a lock is dangerous for throughput but correct for "Must deliver to all".
				// To avoid holding the lock for everyone, we might need a finer grained lock
				// or unlock/wait/lock loop, but that messes up order.
				// Given Go channels, we can't easily "peek" capacity without trying.
				// For now, blocking with lock held ensures strict order and "multicast" semantics.
				// User beware: BufferOverflowSuspend with slow subscribers halts the flow.
				ch <- value
				sent = true

			case BufferOverflowDropOldest:
				// If capacity is 0, we can't drop 'oldest' because there is no buffer.
				// Effectively, we must drop the 'latest' (the current value) because we can't force the subscriber to take it exactly now without blocking.
				if cap(ch) == 0 {
					sent = false
					break
				}

				// Drop oldest in channel -> read one, then write.
				select {
				case <-ch: // Consume oldest
				default:
					// Channel might be empty if consumer racing, or if we just missed the window.
				}

				// Retry send. If still full (racing), we drop the current value (DropLatest fallback).
				select {
				case ch <- value:
					sent = true
				default:
					// Failed to make space or raced. Drop current.
					sent = false
				}

			case BufferOverflowDropLatest:
				// Do nothing, just drop this value for this subscriber.
				sent = false
			}
		}
		_ = sent // used for debugging or metrics if needed
	}
	return true
}
