package state

import (
	"sync"
	"testing"
)

func TestMutableValue_AtomicUpdates(t *testing.T) {
	mv := NewMutableValue(10, nil, nil)

	// CompareAndSet
	if !mv.CompareAndSet(10, 20) {
		t.Error("CompareAndSet(10, 20) failed")
	}
	if mv.Get() != 20 {
		t.Errorf("Expected 20, got %v", mv.Get())
	}
	if mv.CompareAndSet(10, 30) {
		t.Error("CompareAndSet(10, 30) succeeded unexpectedly")
	}
	if mv.Get() != 20 {
		t.Errorf("Expected 20, got %v", mv.Get())
	}

	// Update
	mv.Update(func(v any) any {
		return v.(int) + 5
	})
	if mv.Get() != 25 {
		t.Errorf("Expected 25, got %v", mv.Get())
	}

	// UpdateAndGet
	res := mv.UpdateAndGet(func(v any) any {
		return v.(int) * 2
	})
	if res != 50 || mv.Get() != 50 {
		t.Errorf("UpdateAndGet failed. Res: %v, Value: %v", res, mv.Get())
	}

	// GetAndUpdate
	prev := mv.GetAndUpdate(func(v any) any {
		return v.(int) + 1
	})
	if prev != 50 || mv.Get() != 51 {
		t.Errorf("GetAndUpdate failed. Prev: %v, Value: %v", prev, mv.Get())
	}
}

func TestMutableValueTyped_AtomicUpdates(t *testing.T) {
	mv := NewMutableState(10)

	// CompareAndSet
	if !mv.CompareAndSet(10, 20) {
		t.Error("CompareAndSet(10, 20) failed")
	}
	if mv.Get() != 20 {
		t.Errorf("Expected 20, got %v", mv.Get())
	}
	if mv.CompareAndSet(10, 30) {
		t.Error("CompareAndSet(10, 30) succeeded unexpectedly")
	}
	if mv.Get() != 20 {
		t.Errorf("Expected 20, got %v", mv.Get())
	}

	// Update
	mv.Update(func(v int) int {
		return v + 5
	})
	if mv.Get() != 25 {
		t.Errorf("Expected 25, got %v", mv.Get())
	}

	// UpdateAndGet
	res := mv.UpdateAndGet(func(v int) int {
		return v * 2
	})
	if res != 50 || mv.Get() != 50 {
		t.Errorf("UpdateAndGet failed. Res: %v, Value: %v", res, mv.Get())
	}

	// GetAndUpdate
	prev := mv.GetAndUpdate(func(v int) int {
		return v + 1
	})
	if prev != 50 || mv.Get() != 51 {
		t.Errorf("GetAndUpdate failed. Prev: %v, Value: %v", prev, mv.Get())
	}
}

func TestMutableValue_Concurrency(t *testing.T) {
	mv := NewMutableState(0)
	var wg sync.WaitGroup
	routines := 50
	increments := 100

	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				mv.Update(func(c int) int {
					return c + 1
				})
			}
		}()
	}
	wg.Wait()

	expected := routines * increments
	if mv.Get() != expected {
		t.Errorf("Concurrent updates failed. Expected %d, got %d", expected, mv.Get())
	}
}

func TestWrappers(t *testing.T) {
	// Typed wrapper (MutableValueTyped wrapped as MutableValueTypedWrapper which implements MutableValueTyped? No, wrapper is ... implicit?)
	// Let's test the wrappers explicitly.

	// 1. Typed -> Untyped (Unwrap)
	typed := NewMutableState(10)
	untyped := typed.Unwrap() // returns MutableValue

	untyped.Update(func(current any) any {
		return current.(int) + 10
	})
	if typed.Get() != 20 {
		t.Errorf("Wrappers: Untyped Update failed to affect underlying typed value")
	}

	// 2. Untyped -> Typed (Wrapper)
	// Create untyped
	baseUntyped := NewMutableValue(100, nil, nil)
	// Wrap as typed
	wrappedTyped, err := MutableValueToTyped[int](baseUntyped)
	if err != nil {
		t.Fatalf("MutableValueToTyped failed: %v", err)
	}

	wrappedTyped.Update(func(current int) int {
		return current + 50
	})
	if baseUntyped.Get() != 150 {
		t.Errorf("Wrappers: Typed wrapper Update failed to affect underlying untyped value")
	}
}
