package flow_test

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/zodimo/go-compose/pkg/flow"
	"github.com/zodimo/go-compose/state"
)

// TestMutableStateFlow_CompareAndSet tests the CompareAndSet method
func TestMutableStateFlow_CompareAndSet(t *testing.T) {
	t.Run("succeeds when expect matches current", func(t *testing.T) {
		sf := flow.NewMutableStateFlow(10)

		ok := sf.CompareAndSet(10, 20)
		if !ok {
			t.Error("CompareAndSet should return true when expect matches")
		}
		if got := sf.Value(); got != 20 {
			t.Errorf("Expected 20, got %d", got)
		}
	})

	t.Run("fails when expect does not match current", func(t *testing.T) {
		sf := flow.NewMutableStateFlow(10)

		ok := sf.CompareAndSet(5, 20)
		if ok {
			t.Error("CompareAndSet should return false when expect doesn't match")
		}
		if got := sf.Value(); got != 10 {
			t.Errorf("Value should remain 10, got %d", got)
		}
	})

	t.Run("succeeds without notification when expect equals update equals current", func(t *testing.T) {
		sf := flow.NewMutableStateFlow(10)

		notified := false
		sf.Subscribe(func() {
			notified = true
		})

		// All three equal - should return true but not notify
		ok := sf.CompareAndSet(10, 10)
		if !ok {
			t.Error("CompareAndSet should return true when all values are equal")
		}
		if notified {
			t.Error("Should not notify when update equals current")
		}
	})

	t.Run("notifies subscribers on actual change", func(t *testing.T) {
		sf := flow.NewMutableStateFlow(10)

		notifyCount := 0
		sf.Subscribe(func() {
			notifyCount++
		})

		sf.CompareAndSet(10, 20)
		if notifyCount != 1 {
			t.Errorf("Expected 1 notification, got %d", notifyCount)
		}
	})
}

// TestMutableStateFlow_EqualityConflation tests that Emit skips equal values
func TestMutableStateFlow_EqualityConflation(t *testing.T) {
	t.Run("Emit skips notification for equal value", func(t *testing.T) {
		sf := flow.NewMutableStateFlow(10)

		notifyCount := 0
		sf.Subscribe(func() {
			notifyCount++
		})

		// Emit same value - should not notify
		sf.Emit(10)
		if notifyCount != 0 {
			t.Errorf("Expected 0 notifications for same value, got %d", notifyCount)
		}

		// Emit different value - should notify
		sf.Emit(20)
		if notifyCount != 1 {
			t.Errorf("Expected 1 notification, got %d", notifyCount)
		}

		// Emit same value again - should not notify
		sf.Emit(20)
		if notifyCount != 1 {
			t.Errorf("Expected 1 notification still, got %d", notifyCount)
		}
	})

	t.Run("Update skips notification for equal value", func(t *testing.T) {
		sf := flow.NewMutableStateFlow(10)

		notifyCount := 0
		sf.Subscribe(func() {
			notifyCount++
		})

		// Update returning same value - should not notify
		sf.Update(func(current int) int { return current })
		if notifyCount != 0 {
			t.Errorf("Expected 0 notifications, got %d", notifyCount)
		}

		// Update returning different value - should notify
		sf.Update(func(current int) int { return current + 1 })
		if notifyCount != 1 {
			t.Errorf("Expected 1 notification, got %d", notifyCount)
		}
	})
}

// TestMutableStateFlow_UpdateAndGet tests the UpdateAndGet method
func TestMutableStateFlow_UpdateAndGet(t *testing.T) {
	sf := flow.NewMutableStateFlow(10)

	newValue := sf.UpdateAndGet(func(current int) int {
		return current * 2
	})

	if newValue != 20 {
		t.Errorf("UpdateAndGet should return new value 20, got %d", newValue)
	}
	if got := sf.Value(); got != 20 {
		t.Errorf("Value should be 20, got %d", got)
	}
}

// TestMutableStateFlow_GetAndUpdate tests the GetAndUpdate method
func TestMutableStateFlow_GetAndUpdate(t *testing.T) {
	sf := flow.NewMutableStateFlow(10)

	oldValue := sf.GetAndUpdate(func(current int) int {
		return current * 2
	})

	if oldValue != 10 {
		t.Errorf("GetAndUpdate should return old value 10, got %d", oldValue)
	}
	if got := sf.Value(); got != 20 {
		t.Errorf("Value should be 20, got %d", got)
	}
}

// TestMutableStateFlow_ConcurrentCAS tests thread-safety of CAS operations
func TestMutableStateFlow_ConcurrentCAS(t *testing.T) {
	sf := flow.NewMutableStateFlow(0)

	var wg sync.WaitGroup
	iterations := 100
	goroutines := 10

	// Multiple goroutines incrementing using Update
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				sf.Update(func(current int) int {
					return current + 1
				})
			}
		}()
	}

	wg.Wait()

	expected := goroutines * iterations
	if got := sf.Value(); got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

// TestMutableStateFlow_ConcurrentCAS_GetAndUpdate tests GetAndUpdate concurrency
func TestMutableStateFlow_ConcurrentCAS_GetAndUpdate(t *testing.T) {
	sf := flow.NewMutableStateFlow(0)

	var wg sync.WaitGroup
	var sum atomic.Int64
	iterations := 100
	goroutines := 10

	// Sum all previous values using GetAndUpdate
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				oldVal := sf.GetAndUpdate(func(current int) int {
					return current + 1
				})
				sum.Add(int64(oldVal))
			}
		}()
	}

	wg.Wait()

	// Sum of 0+1+2+...+(n-1) = n*(n-1)/2
	total := goroutines * iterations
	expectedSum := int64(total * (total - 1) / 2)
	if got := sum.Load(); got != expectedSum {
		t.Errorf("Expected sum %d, got %d", expectedSum, got)
	}
}

// TestMutableStateFlow_WithPolicy tests custom mutation policy
func TestMutableStateFlow_WithPolicy(t *testing.T) {
	// Custom comparator: only compare the "ID" field
	type Item struct {
		ID   int
		Name string
	}

	compareByID := func(a, b Item) bool {
		return a.ID == b.ID
	}

	policy := state.NewMutationPolicy(compareByID, nil)

	sf := flow.NewMutableStateFlow(Item{ID: 1, Name: "Alice"}, flow.WithPolicy(policy))

	notifyCount := 0
	sf.Subscribe(func() {
		notifyCount++
	})

	// Same ID, different name - should not notify (using custom comparator)
	sf.Emit(Item{ID: 1, Name: "Bob"})
	if notifyCount != 0 {
		t.Errorf("Expected 0 notifications (same ID), got %d", notifyCount)
	}

	// Different ID - should notify
	sf.Emit(Item{ID: 2, Name: "Charlie"})
	if notifyCount != 1 {
		t.Errorf("Expected 1 notification (different ID), got %d", notifyCount)
	}
}
