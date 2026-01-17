package flow_test

import (
	"context"
	"testing"
	"time"

	"github.com/zodimo/go-compose/pkg/flow"
)

func TestSharedFlow_EdgeCases(t *testing.T) {
	// Test 1: DropOldest with Zero Capacity (Rendezvous)
	// Should behave like DropLatest because we can't "make space"
	t.Run("DropOldest with Zero Capacity", func(t *testing.T) {
		shared := flow.NewMutableSharedFlow[int](0, 0, flow.BufferOverflowDropOldest)

		// Subscribe with slow consumer
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		received := make(chan int, 10)
		go shared.Collect(ctx, func(v int) {
			time.Sleep(10 * time.Millisecond) // Slow
			received <- v
		})

		// Give time to subscribe
		time.Sleep(20 * time.Millisecond)

		// Emit. Subscriber is waiting, so first should go through?
		// Actually, if subscriber is "waiting" on channel receive, ch <- val succeeds.
		// Slow consumer logic here means it takes time TO PROCESS.
		// Collect loop: receives -> processes -> loops back.
		// If processing takes time, it's NOT reading from channel.

		shared.TryEmit(1) // Should go through if subscriber is ready at select
		time.Sleep(5 * time.Millisecond)
		shared.TryEmit(2) // Subscriber likely busy processing 1, so channel blocked. capacity=0. DropOldest->DropLatest.
		shared.TryEmit(3)

		time.Sleep(50 * time.Millisecond)
		cancel()

		// Verify what we got.
		// Expect 1 to be received. 2 and 3 potentially dropped.
		// Note: "TryEmit" calls Check, but "Emit" calls Check.
		// With DropOldest, Emit should not suspend, it should drop.

		count := 0
		close(received)
		for v := range received {
			count++
			t.Logf("Received: %d", v)
		}

		if count == 0 {
			t.Error("Specific test logic depends on timing, but expected at least 1 value")
		}
		// Exact count is brittle with timing, but we verify no panic and behavior is non-blocking.
	})

	// Test 2: Invalid Inputs should panic
	t.Run("Invalid Inputs", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("The code did not panic")
			}
		}()
		_ = flow.NewMutableSharedFlow[int](-1, -1, flow.BufferOverflowSuspend)
	})

	// Test 3: DropOldest with Capacity
	t.Run("DropOldest with Capacity", func(t *testing.T) {
		// Replay=0, Extra=1. Buffer size 1.
		shared := flow.NewMutableSharedFlow[int](0, 1, flow.BufferOverflowDropOldest)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Dont read anything
		go shared.Collect(ctx, func(v int) {
			time.Sleep(1 * time.Second)
		})
		time.Sleep(20 * time.Millisecond)

		shared.Emit(1) // Fits in buffer
		shared.Emit(2) // Should drop 1, put 2? Or drop 1 from buffer?
		// DropOldest means: Buffer is [1]. Full.
		// Pop [1]. Buffer empty.
		// Push 2. Buffer [2].

		// To verify this, we need to read properly eventually.
		// Let's create a new test setup where we control the read.
	})
}

func TestSharedFlow_DropOldest_Explicit(t *testing.T) {
	// Replay=0, Extra=1.
	shared := flow.NewMutableSharedFlow[int](0, 1, flow.BufferOverflowDropOldest)

	// Use a channel we can control
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// We want to verify what's strictly in the channel.
	// We can't peek the internal channel. We have to consume.

	// Trick: The 'Collect' creates the channel.
	// We can't access it.

	// Launch flow.
	values := make(chan int, 10)
	ready := make(chan struct{})

	go shared.Collect(ctx, func(v int) {
		// Block here until we allow it
		<-ready
		values <- v
	})

	time.Sleep(20 * time.Millisecond) // Wait for sub

	shared.Emit(1) // In buffer
	shared.Emit(2) // Should drop 1, insert 2
	shared.Emit(3) // Should drop 2, insert 3

	close(ready) // Unleash subscriber

	// We expect 1 (which was picked up and held by the subscriber)
	// Then we expect 3 (which replaced 2 in the buffer).
	// 2 should be dropped.

	// First value: 1
	select {
	case v := <-values:
		if v != 1 {
			t.Errorf("Expected 1, got %d", v)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Timeout waiting for value 1")
	}

	// Second value: 3
	select {
	case v := <-values:
		if v != 3 {
			t.Errorf("Expected 3, got %d", v)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Timeout waiting for value 3")
	}
}
