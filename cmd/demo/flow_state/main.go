package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/zodimo/go-compose/pkg/flow"
)

func main() {
	fmt.Println("=== StateFlow Demo ===")
	fmt.Println("Demonstrating: Conflation, Latest Value, and Atomic Updates")

	// 1. Create a StateFlow with initial value 0
	stateFlow := flow.NewMutableStateFlow(0)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 2. Launch a Slow Subscriber
	// StateFlow is conflated. If we emit faster than this subscriber can collect,
	// it should skip intermediate values and always get the latest.
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("[Sub] Connecting...")
		count := 0
		err := stateFlow.Collect(ctx, func(val int) {
			fmt.Printf("[Sub] Received: %d\n", val)
			count++
			// Simulate processing time
			time.Sleep(200 * time.Millisecond)
		})
		if err != nil && err != context.Canceled {
			fmt.Printf("[Sub] Error: %v\n", err)
		}
	}()

	// 3. Update State Rapidly
	// We will update 1..10 in 100ms (10ms per update).
	// The subscriber takes 200ms per value.
	// Expected: Sub sees initial (0), maybe one or two intermediates, and definitely final (10).
	time.Sleep(100 * time.Millisecond)
	fmt.Println("[Main] Updating state 1..10 rapidly...")
	for i := 1; i <= 10; i++ {
		stateFlow.Emit(i) // Or stateFlow.Value = i equivalent
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("[Main] Updates done. Current Value:", stateFlow.Value())

	// 4. Atomic Updates (CompareAndSet)
	// Simulate concurrent increment
	fmt.Println("[Main] Performing concurrent atomic increments...")
	var atomicWg sync.WaitGroup
	for i := 0; i < 5; i++ {
		atomicWg.Add(1)
		go func(id int) {
			defer atomicWg.Done()
			// Update: retry loop provided by helper methods
			stateFlow.Update(func(current int) int {
				return current + 1
			})
			fmt.Printf("[Worker %d] Incremented\n", id)
		}(i)
	}
	atomicWg.Wait()
	fmt.Println("[Main] Final Atomic Value (Should be 15):", stateFlow.Value())

	// Allow subscriber to catch up with the *latest* value (15)
	time.Sleep(500 * time.Millisecond)
	cancel()
	wg.Wait()
	fmt.Println("=== Done ===")
}
