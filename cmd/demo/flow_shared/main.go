package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zodimo/go-compose/pkg/flow"
)

func main() {
	fmt.Println("=== SharedFlow Demo ===")
	fmt.Println("Demonstrating: Broadcast, Replay, and BufferOverflowDropOldest")

	// 1. Create a SharedFlow
	// Replay: 1 (New subscribers get the last 1 value)
	// ExtraBuffer: 2 (Buffer size = 1 + 2 = 3)
	// Overflow: DropOldest (If buffer full, drop the oldest value to make room)
	shared := flow.NewMutableSharedFlow[int](1, 2, flow.BufferOverflowDropOldest)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 2. Launch Subscriber 1 (Fast)
	go func() {
		fmt.Println("[Sub 1] Connecting...")
		err := shared.Collect(ctx, func(val int) {
			fmt.Printf("[Sub 1] Received: %d\n", val)
		})
		if err != nil && err != context.Canceled {
			fmt.Printf("[Sub 1] Error: %v\n", err)
		}
	}()

	// Give Sub 1 time to connect
	time.Sleep(100 * time.Millisecond)

	// 3. Emit values
	fmt.Println("[Main] Emitting 1, 2, 3")
	shared.Emit(1)
	shared.Emit(2)
	shared.Emit(3)
	time.Sleep(100 * time.Millisecond)

	// 4. Launch Subscriber 2 (Late) - Should see Replay (3) then new values
	go func() {
		fmt.Println("[Sub 2] Connecting (Late)...")
		err := shared.Collect(ctx, func(val int) {
			fmt.Printf("[Sub 2] Received: %d\n", val)
			// Simulate slow processing to force buffer overflow for Sub 2?
			// Actually, DropOldest happens on Emit. If Sub 2 is blocked here,
			// the *next* Emit might drop values for Sub 2.
			time.Sleep(200 * time.Millisecond)
		})
		if err != nil && err != context.Canceled {
			fmt.Printf("[Sub 2] Error: %v\n", err)
		}
	}()

	// Give Sub 2 time to receive replay
	time.Sleep(100 * time.Millisecond)

	// 5. Emit more values rapidly to test Overflow
	// Sub 2 is sleeping 200ms per value. We emit faster.
	// Buffer capacity is 3.
	// Values: 4, 5, 6, 7, 8
	fmt.Println("[Main] Emitting 4..8 rapidly")
	for i := 4; i <= 8; i++ {
		fmt.Printf("[Main] Emitting %d\n", i)
		shared.Emit(i)
		time.Sleep(50 * time.Millisecond)
	}

	// Wait for processing
	time.Sleep(1 * time.Second)
	fmt.Println("=== Done ===")
}
