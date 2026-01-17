package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zodimo/go-compose/pkg/flow"
)

// TimerFlow emits a tick every interval
func TimerFlow(interval time.Duration) flow.Flow[int64] {
	return flow.NewFlow(func(ctx context.Context, emit func(int64)) error {
		fmt.Println("[Timer] STARTING upstream...")
		ticker := time.NewTicker(interval)
		defer func() {
			ticker.Stop()
			fmt.Println("[Timer] STOPPING upstream...")
		}()

		var counter int64 = 0
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-ticker.C:
				counter++
				fmt.Printf("[Timer] Emit %d\n", counter)
				emit(counter)
			}
		}
	})
}

func main() {
	fmt.Println("=== ShareIn Demo ===")
	fmt.Println("Demonstrating: SharingStarted.WhileSubscribed")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Create a Cold Flow
	coldFlow := TimerFlow(200 * time.Millisecond)

	// 2. Share it!
	// Replay: 1
	// Started: WhileSubscribed (Start when 1st sub appears, Stop 1s after last sub leaves)
	// We use 1s stop timeout to demonstrate the "grace period".
	hotFlow := flow.ShareIn(
		ctx, // Scope for the shared flow
		coldFlow,
		flow.SharingStartedWhileSubscribed(1000*time.Millisecond, 0),
		1, // Replay
	)

	fmt.Println("[Main] Hot flow created. Upstream should NOT run yet.")
	time.Sleep(500 * time.Millisecond)

	// 3. First Subscriber
	fmt.Println("[Main] --> Adding Subscriber 1")
	sub1Ctx, sub1Cancel := context.WithCancel(ctx)
	sub1Done := make(chan struct{})
	go func() {
		hotFlow.Collect(sub1Ctx, func(val int64) {
			fmt.Printf("[Sub 1] Got %d\n", val)
		})
		close(sub1Done)
	}()

	// Let it run for 2 seconds (approx 10 ticks)
	time.Sleep(1100 * time.Millisecond)

	// 4. Remove Subscriber 1
	fmt.Println("[Main] <-- Removing Subscriber 1 (Expect stop in 1s)")
	sub1Cancel()
	<-sub1Done // Wait for Collect to return

	// 5. Watch for Stop
	// Upstream should stop after ~1000ms grace period.
	time.Sleep(1500 * time.Millisecond)

	// 6. Resubscribe
	fmt.Println("[Main] --> Adding Subscriber 2 (Expect restart)")
	sub2Ctx, sub2Cancel := context.WithCancel(ctx)
	go func() {
		hotFlow.Collect(sub2Ctx, func(val int64) {
			fmt.Printf("[Sub 2] Got %d\n", val)
		})
	}()

	time.Sleep(1000 * time.Millisecond)
	sub2Cancel()
	fmt.Println("=== Done ===")
}
