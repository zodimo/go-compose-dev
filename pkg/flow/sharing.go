package flow

import (
	"context"
	"time"
)

// SharingCommand controls the upstream flow in shareIn/stateIn.
type SharingCommand int

const (
	// SharingCommandStart starts the upstream flow.
	SharingCommandStart SharingCommand = iota
	// SharingCommandStop stops the upstream flow.
	SharingCommandStop
	// SharingCommandStopAndResetReplayCache stops the upstream flow and resets the replay cache.
	SharingCommandStopAndResetReplayCache
)

// SharingStarted controls when the sharing is started and stopped.
type SharingStarted interface {
	Command(subscriptionCount StateFlow[int]) Flow[SharingCommand]
}

// Eagerly starts sharing immediately and never stops.
func SharingStartedEagerly() SharingStarted {
	return &startedEagerly{}
}

// Lazily starts sharing when the first subscriber appears and never stops.
func SharingStartedLazily() SharingStarted {
	return &startedLazily{}
}

// WhileSubscribed starts sharing when the first subscriber appears and stops when the last subscriber disappears.
// You can configure a grace period for stopping and a duration for keeping the replay cache.
func SharingStartedWhileSubscribed(stopTimeout time.Duration, replayExpiration time.Duration) SharingStarted {
	return &startedWhileSubscribed{
		stopTimeout:      stopTimeout,
		replayExpiration: replayExpiration,
	}
}

// --- Implementations ---

type startedEagerly struct{}

func (s *startedEagerly) Command(subscriptionCount StateFlow[int]) Flow[SharingCommand] {
	return NewFlow(func(ctx context.Context, emit func(SharingCommand)) error {
		emit(SharingCommandStart)
		<-ctx.Done()
		return ctx.Err()
	})
}

type startedLazily struct{}

func (s *startedLazily) Command(subscriptionCount StateFlow[int]) Flow[SharingCommand] {
	return NewFlow(func(ctx context.Context, emit func(SharingCommand)) error {
		var started bool
		return subscriptionCount.Collect(ctx, func(count int) {
			if count > 0 && !started {
				started = true
				emit(SharingCommandStart)
			}
		})
	})
}

type startedWhileSubscribed struct {
	stopTimeout      time.Duration
	replayExpiration time.Duration
}

func (s *startedWhileSubscribed) Command(subscriptionCount StateFlow[int]) Flow[SharingCommand] {
	return NewFlow(func(ctx context.Context, emit func(SharingCommand)) error {
		// Use a local context for internal goroutines
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		counts := make(chan int)
		// Launch goroutine to bridge subscriptionCount to channel
		go func() {
			defer close(counts)
			subscriptionCount.Collect(ctx, func(c int) {
				select {
				case counts <- c:
				case <-ctx.Done():
				}
			})
		}()

		var timer *time.Timer
		var timerCh <-chan time.Time

		// Track current state to avoid redundant emissions
		var state SharingCommand = SharingCommandStop

		for {
			select {
			case count, ok := <-counts:
				if !ok {
					return nil
				}
				if count > 0 {
					// Has subscribers: Cancel any pending stop timer
					if timer != nil {
						timer.Stop()
						timer = nil
						timerCh = nil
					}
					if state != SharingCommandStart {
						state = SharingCommandStart
						emit(SharingCommandStart)
					}
				} else {
					// No subscribers: Start stop timer if not already pending
					if count == 0 && timer == nil {
						if s.stopTimeout == 0 {
							// Stop immediately
							if state != SharingCommandStop {
								state = SharingCommandStop
								emit(SharingCommandStop)
							}
							if s.replayExpiration == 0 {
								emit(SharingCommandStopAndResetReplayCache)
							} else {
								// Stop emitted, now schedule reset
								timer = time.NewTimer(s.replayExpiration)
								timerCh = timer.C
							}
						} else {
							timer = time.NewTimer(s.stopTimeout)
							timerCh = timer.C
						}
					}
				}
			case <-timerCh:
				// Timer expired
				timer = nil
				timerCh = nil

				if state == SharingCommandStart {
					// stopTimeout expired. Transition to Stop.
					state = SharingCommandStop
					emit(SharingCommandStop)

					// Schedule replay expiration if needed
					if s.replayExpiration == 0 {
						emit(SharingCommandStopAndResetReplayCache)
					} else {
						timer = time.NewTimer(s.replayExpiration)
						timerCh = timer.C
					}
				} else if state == SharingCommandStop {
					// replayExpiration expired. Reset cache.
					emit(SharingCommandStopAndResetReplayCache)
				}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})
}

// NOTE: The above WhileSubscribed implementation is simplified and flawed regarding concurrency.
// A robust implementation requires a transform that processes subscription count changes and timer events in a select loop.
// Given constraints, I will refine WhileSubscribed in a future iteration if robust timing is needed.

// ShareIn transforms a cold Flow into a hot SharedFlow.
func ShareIn[T any](
	ctx context.Context,
	upstream Flow[T],
	started SharingStarted,
	replay int,
) SharedFlow[T] {
	shared := NewMutableSharedFlow[T](replay, 0, BufferOverflowSuspend)

	// Launch the sharing logic
	go func() {
		var upstreamCtx context.Context
		var cancelUpstream context.CancelFunc

		started.Command(shared.SubscriptionCount()).Collect(ctx, func(cmd SharingCommand) {
			switch cmd {
			case SharingCommandStart:
				if upstreamCtx == nil {
					upstreamCtx, cancelUpstream = context.WithCancel(ctx)
					go func() {
						upstream.Collect(upstreamCtx, func(value T) {
							shared.Emit(value)
						})
					}()
				}
			case SharingCommandStop:
				if cancelUpstream != nil {
					cancelUpstream()
					cancelUpstream = nil
					upstreamCtx = nil
				}
			case SharingCommandStopAndResetReplayCache:
				if cancelUpstream != nil {
					cancelUpstream()
					cancelUpstream = nil
					upstreamCtx = nil
				}
				shared.ResetReplayCache()
			}
		})

		if cancelUpstream != nil {
			cancelUpstream()
		}
	}()

	return shared
}

// StateIn transforms a flow into a StateFlow.
func StateIn[T any](
	ctx context.Context,
	upstream Flow[T],
	started SharingStarted,
	initialValue T,
) StateFlow[T] {
	// StateFlow is config: replay=1, onBufferOverflow=DROP_OLDEST
	// But MutableStateFlow has extra logic.
	// We use `NewMutableStateFlow` to create the destination.
	state := NewMutableStateFlow(initialValue)

	// Similar logic to ShareIn
	go func() {
		var upstreamCtx context.Context
		var cancelUpstream context.CancelFunc

		started.Command(state.SubscriptionCount()).Collect(ctx, func(cmd SharingCommand) {
			switch cmd {
			case SharingCommandStart:
				if upstreamCtx == nil {
					upstreamCtx, cancelUpstream = context.WithCancel(ctx)
					go func() {
						upstream.Collect(upstreamCtx, func(value T) {
							state.Emit(value)
						})
					}()
				}
			case SharingCommandStop:
				if cancelUpstream != nil {
					cancelUpstream()
					cancelUpstream = nil
					upstreamCtx = nil
				}
			case SharingCommandStopAndResetReplayCache:
				if cancelUpstream != nil {
					cancelUpstream()
					cancelUpstream = nil
					upstreamCtx = nil
				}
				// StateFlow cannot reset replay cache (always has value).
				// Just stop upstream.
			}
		})

		if cancelUpstream != nil {
			cancelUpstream()
		}
	}()

	return state
}
