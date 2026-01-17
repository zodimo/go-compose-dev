package flow_test

import (
	"context"
	"testing"
	"time"

	"github.com/zodimo/go-compose/pkg/flow"
)

func TestSharedFlow_Multicast(t *testing.T) {
	// Create a shared flow
	shared := flow.NewMutableSharedFlow[int](0, 0, flow.BufferOverflowSuspend)

	// Launch two collectors
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	v1 := make(chan int)
	v2 := make(chan int)

	go shared.Collect(ctx, func(v int) { v1 <- v })
	go shared.Collect(ctx, func(v int) { v2 <- v })

	// Wait for subscriptions to be active (simplistic wait)
	time.Sleep(50 * time.Millisecond)

	go shared.Emit(1)

	// Both should receive
	select {
	case v := <-v1:
		if v != 1 {
			t.Errorf("v1 got %d, want 1", v)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("v1 timed out")
	}

	select {
	case v := <-v2:
		if v != 1 {
			t.Errorf("v2 got %d, want 1", v)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("v2 timed out")
	}
}

func TestSharedFlow_Replay(t *testing.T) {
	// Replay = 2
	shared := flow.NewMutableSharedFlow[int](2, 0, flow.BufferOverflowSuspend)
	shared.Emit(1)
	shared.Emit(2)
	shared.Emit(3) // 1 should be dropped from replay

	// New collector should get 2 and 3
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make(chan int, 2)
	go shared.Collect(ctx, func(v int) { results <- v })

	// Verify 2
	select {
	case v := <-results:
		if v != 2 {
			t.Errorf("got %d, want 2", v)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout 1")
	}

	// Verify 3
	select {
	case v := <-results:
		if v != 3 {
			t.Errorf("got %d, want 3", v)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timeout 2")
	}
}

func TestSubscriptionCount(t *testing.T) {
	shared := flow.NewMutableSharedFlow[int](0, 0, flow.BufferOverflowSuspend)

	// Check initial count
	if shared.SubscriptionCount().Value() != 0 {
		t.Errorf("initial count %d, want 0", shared.SubscriptionCount().Value())
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Collect
	ready := make(chan struct{})
	go func() {
		close(ready)
		shared.Collect(ctx, func(int) {})
	}()
	<-ready
	time.Sleep(50 * time.Millisecond) // Allow propagation

	if shared.SubscriptionCount().Value() != 1 {
		t.Errorf("count %d, want 1", shared.SubscriptionCount().Value())
	}

	// Cancel
	cancel()
	time.Sleep(50 * time.Millisecond) // Allow cleanup

	if shared.SubscriptionCount().Value() != 0 {
		t.Errorf("end count %d, want 0", shared.SubscriptionCount().Value())
	}
}

func TestShareIn_Lazily(t *testing.T) {
	ctx := context.Background()
	upstream := flow.NewFlow(func(ctx context.Context, emit func(int)) error {
		emit(100)
		<-ctx.Done()
		return nil
	})

	// Lazily: starts when collected
	shared := flow.ShareIn(ctx, upstream, flow.SharingStartedLazily(), 1)

	// Should not have started yet (no subscribers)
	// We can't easily peek upstream start without side effects.

	// Subscribe
	subCtx, subCancel := context.WithCancel(ctx)
	defer subCancel()

	results := make(chan int)
	go shared.Collect(subCtx, func(v int) { results <- v })

	select {
	case v := <-results:
		if v != 100 {
			t.Errorf("got %d, want 100", v)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("timeout")
	}
}
