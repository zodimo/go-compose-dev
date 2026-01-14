package store

import (
	"testing"
)

func TestDerivedState_BasicCalculation(t *testing.T) {
	mv1 := NewMutableValue(10, nil, func(a, b any) bool { return a == b })
	mv2 := NewMutableValue(20, nil, func(a, b any) bool { return a == b })

	calculatedCalls := 0
	derived := DerivedStateOf(func() int {
		calculatedCalls++
		return mv1.Get().(int) + mv2.Get().(int)
	})

	// Initial get
	if got := derived.Get(); got != 30 {
		t.Errorf("Expected 30, got %v", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// Second get, should be cached
	if got := derived.Get(); got != 30 {
		t.Errorf("Expected 30, got %v", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// Update dependency
	mv1.Set(15)

	// Third get, should recalculate
	if got := derived.Get(); got != 35 {
		t.Errorf("Expected 35, got %v", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations, got %d", calculatedCalls)
	}
}

func TestDerivedState_DependencySwitching(t *testing.T) {
	// Setup: A -> B or C
	toggle := NewMutableValue(true, nil, func(a, b any) bool { return a == b })
	valB := NewMutableValue("B", nil, func(a, b any) bool { return a == b })
	valC := NewMutableValue("C", nil, func(a, b any) bool { return a == b })

	calculatedCalls := 0
	derived := DerivedStateOf(func() string {
		calculatedCalls++
		if toggle.Get().(bool) {
			return valB.Get().(string)
		}
		return valC.Get().(string)
	})

	// 1. Initial: true -> depends on toggle, valB
	if got := derived.Get(); got != "B" {
		t.Errorf("Expected B, got %v", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// 2. Modify valC (should NOT trigger recalc on next Get, as it's not a dependency)
	valC.Set("C2")
	if got := derived.Get(); got != "B" {
		t.Errorf("Expected B, got %v", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// 3. Toggle to false -> depends on toggle, valC
	toggle.Set(false)
	if got := derived.Get(); got != "C2" {
		t.Errorf("Expected C2, got %v", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations, got %d", calculatedCalls)
	}

	// 4. Modify valB (should NOT trigger recalc on next Get, as it's no longer a dependency)
	valB.Set("B2")
	if got := derived.Get(); got != "C2" {
		t.Errorf("Expected C2, got %v", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations, got %d", calculatedCalls)
	}
}

func TestDerivedState_DiamondDependency(t *testing.T) {
	// Root -> A, Root -> B
	// Derived -> A + B
	root := NewMutableValue(1, nil, func(a, b any) bool { return a == b })

	derivedA := DerivedStateOf(func() int {
		return root.Get().(int) * 2
	})

	derivedB := DerivedStateOf(func() int {
		return root.Get().(int) * 3
	})

	calculatedCalls := 0
	derivedFinal := DerivedStateOf(func() int {
		calculatedCalls++
		return derivedA.Get().(int) + derivedB.Get().(int)
	})

	// Initial: (1*2) + (1*3) = 5
	if got := derivedFinal.Get(); got != 5 {
		t.Errorf("Expected 5, got %v", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// Update root
	root.Set(2)
	// DerivedA -> 4, DerivedB -> 6
	// DerivedFinal -> 10

	if got := derivedFinal.Get(); got != 10 {
		t.Errorf("Expected 10, got %v", got)
	}
	// Depending on implementation, might calculate more than once if not optimized, but should be at least correct.
	// Since we haven't implemented eager invalidation, this relies on pull-based lazy evaluation.
	// derivedFinal.Get() -> checks derivedA.Version() and derivedB.Version()
	// derivedA.Version() -> checks root.Version(), sees change, recalculates -> returns new version
	// derivedB.Version() -> checks root.Version(), sees change, recalculates -> returns new version
	// derivedFinal sees change in A and B versions -> recalculates

	// So it should work correctly now that DerivedState implements Versionable.
}

func TestDerivedState_NestedCaching(t *testing.T) {
	// A -> Derived B -> Derived C
	// If A changes but B ends up same value, C should NOT recalculate!

	valA := NewMutableValue(10, nil, func(a, b any) bool { return a == b })

	calcB := 0
	derivedB := DerivedStateOf(func() int {
		calcB++
		a := valA.Get().(int)
		if a > 5 {
			return 1 // Always 1 if A > 5
		}
		return 0
	})

	calcC := 0
	derivedC := DerivedStateOf(func() int {
		calcC++
		return derivedB.Get().(int) + 100
	})

	// Initial: A=10 -> B=1 -> C=101
	if got := derivedC.Get(); got != 101 {
		t.Errorf("Got %d", got)
	}
	if calcB != 1 {
		t.Errorf("B calc %d", calcB)
	}
	if calcC != 1 {
		t.Errorf("C calc %d", calcC)
	}

	// Update A to 20. B is still 1. C should NOT recalculate.
	valA.Set(20)

	if got := derivedC.Get(); got != 101 {
		t.Errorf("Got %d", got)
	}

	// derivedC.Get() calls derivedB.Version()
	// derivedB.Version() checks valA -> recalc B. B is still 1. Version should NOT change?
	// If derivedB version doesn't change, derivedC doesn't recalc.

	if calcB != 2 {
		t.Errorf("B should have recalculated, got %d", calcB)
	}
	if calcC != 1 {
		t.Errorf("C shouldn't have recalculated, got %d", calcC)
	}
}
