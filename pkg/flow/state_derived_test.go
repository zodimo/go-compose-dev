package flow_test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/zodimo/go-compose/pkg/flow"
	"github.com/zodimo/go-compose/state"
	"github.com/zodimo/go-compose/store"
)

// TestFlowAsStateSource verifies that MutableStateFlow can be used as a
// dependency for DerivedState via the StateChangeNotifier interface.
func TestFlowAsStateSource(t *testing.T) {
	flowValue := flow.NewMutableStateFlow(10)

	calculatedCalls := 0
	derived := state.DerivedStateOf(func() int {
		calculatedCalls++
		return flowValue.Value() * 2
	})

	// Initial get should calculate
	if got := derived.Get(); got != 20 {
		t.Errorf("Expected 20, got %d", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// Second get should be cached
	if got := derived.Get(); got != 20 {
		t.Errorf("Expected 20, got %d", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation (cached), got %d", calculatedCalls)
	}

	// Update the flow via Emit
	flowValue.Emit(15)

	// Should recalculate
	if got := derived.Get(); got != 30 {
		t.Errorf("Expected 30, got %d", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations, got %d", calculatedCalls)
	}
}

// TestFlowAsStateSource_Update verifies that Update() also triggers derived state recalculation.
func TestFlowAsStateSource_Update(t *testing.T) {
	flowValue := flow.NewMutableStateFlow(5)

	calculatedCalls := 0
	derived := state.DerivedStateOf(func() int {
		calculatedCalls++
		return flowValue.Value() + 100
	})

	// Initial get
	if got := derived.Get(); got != 105 {
		t.Errorf("Expected 105, got %d", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// Update via Update method
	flowValue.Update(func(current int) int {
		return current + 10
	})

	// Should recalculate
	if got := derived.Get(); got != 115 {
		t.Errorf("Expected 115, got %d", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations, got %d", calculatedCalls)
	}
}

// TestFlowAsStateSource_MultipleFlows verifies derived state with multiple flow dependencies.
func TestFlowAsStateSource_MultipleFlows(t *testing.T) {
	flowA := flow.NewMutableStateFlow(10)
	flowB := flow.NewMutableStateFlow(20)

	calculatedCalls := 0
	derived := state.DerivedStateOf(func() int {
		calculatedCalls++
		return flowA.Value() + flowB.Value()
	})

	// Initial calculation
	if got := derived.Get(); got != 30 {
		t.Errorf("Expected 30, got %d", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// Update flow A
	flowA.Emit(15)
	if got := derived.Get(); got != 35 {
		t.Errorf("Expected 35, got %d", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations, got %d", calculatedCalls)
	}

	// Update flow B
	flowB.Emit(30)
	if got := derived.Get(); got != 45 {
		t.Errorf("Expected 45, got %d", got)
	}
	if calculatedCalls != 3 {
		t.Errorf("Expected 3 calculations, got %d", calculatedCalls)
	}
}

// TestFlowAsStateSource_DependencySwitching verifies dynamic dependency tracking.
func TestFlowAsStateSource_DependencySwitching(t *testing.T) {
	toggle := flow.NewMutableStateFlow(true)
	flowA := flow.NewMutableStateFlow("A")
	flowB := flow.NewMutableStateFlow("B")

	calculatedCalls := 0
	derived := state.DerivedStateOf(func() string {
		calculatedCalls++
		if toggle.Value() {
			return flowA.Value()
		}
		return flowB.Value()
	})

	// Initial: toggle=true, depends on flowA
	if got := derived.Get(); got != "A" {
		t.Errorf("Expected 'A', got '%s'", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// Modify flowB (should NOT trigger recalculation since it's not a dependency)
	flowB.Emit("B2")
	if derived.IsInvalid() {
		t.Error("Derived should NOT be invalid after updating unused dependency")
	}
	if got := derived.Get(); got != "A" {
		t.Errorf("Expected 'A', got '%s'", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation (no change), got %d", calculatedCalls)
	}

	// Toggle to false: now depends on flowB
	toggle.Emit(false)
	if got := derived.Get(); got != "B2" {
		t.Errorf("Expected 'B2', got '%s'", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations, got %d", calculatedCalls)
	}

	// Modify flowA (should NOT trigger recalculation, no longer a dependency)
	flowA.Emit("A2")
	if derived.IsInvalid() {
		t.Error("Derived should NOT be invalid after updating unused dependency")
	}
	if got := derived.Get(); got != "B2" {
		t.Errorf("Expected 'B2', got '%s'", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations (no change), got %d", calculatedCalls)
	}
}

// TestFlowAsStateSource_ChainedDerived verifies flow -> derived -> derived chains.
func TestFlowAsStateSource_ChainedDerived(t *testing.T) {
	root := flow.NewMutableStateFlow(1)

	calcB := 0
	derivedB := state.DerivedStateOf(func() int {
		calcB++
		return root.Value() + 10
	})

	calcC := 0
	derivedC := state.DerivedStateOf(func() int {
		calcC++
		return derivedB.Get() + 100
	})

	// Initial: 1 + 10 + 100 = 111
	if got := derivedC.Get(); got != 111 {
		t.Errorf("Expected 111, got %d", got)
	}
	if calcB != 1 || calcC != 1 {
		t.Errorf("Expected calcB=1, calcC=1, got calcB=%d, calcC=%d", calcB, calcC)
	}

	// Update root
	root.Emit(2)

	// 2 + 10 + 100 = 112
	if got := derivedC.Get(); got != 112 {
		t.Errorf("Expected 112, got %d", got)
	}
	if calcB != 2 || calcC != 2 {
		t.Errorf("Expected calcB=2, calcC=2, got calcB=%d, calcC=%d", calcB, calcC)
	}
}

// TestFlowAsStateSource_DiamondDependency verifies diamond dependency pattern with flows.
func TestFlowAsStateSource_DiamondDependency(t *testing.T) {
	// root -> derivedA, root -> derivedB -> derivedFinal
	root := flow.NewMutableStateFlow(1)

	derivedA := state.DerivedStateOf(func() int {
		return root.Value() * 2
	})

	derivedB := state.DerivedStateOf(func() int {
		return root.Value() * 3
	})

	calcFinal := 0
	derivedFinal := state.DerivedStateOf(func() int {
		calcFinal++
		return derivedA.Get() + derivedB.Get()
	})

	// Initial: (1*2) + (1*3) = 5
	if got := derivedFinal.Get(); got != 5 {
		t.Errorf("Expected 5, got %d", got)
	}
	if calcFinal != 1 {
		t.Errorf("Expected 1 calculation, got %d", calcFinal)
	}

	// Update root
	root.Emit(2)

	// (2*2) + (2*3) = 10
	if got := derivedFinal.Get(); got != 10 {
		t.Errorf("Expected 10, got %d", got)
	}
}

// TestFlowAsStateSource_ConcurrentAccess verifies thread-safety with concurrent reads/writes.
func TestFlowAsStateSource_ConcurrentAccess(t *testing.T) {
	flowValue := flow.NewMutableStateFlow(0)

	derived := state.DerivedStateOf(func() int {
		return flowValue.Value() * 2
	})

	var wg sync.WaitGroup
	iterations := 100

	// Multiple readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				_ = derived.Get()
			}
		}()
	}

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < iterations; j++ {
			flowValue.Emit(j)
		}
	}()

	wg.Wait()

	// Should complete without race or panic
	// Final value should be (iterations-1) * 2
	expected := (iterations - 1) * 2
	if got := derived.Get(); got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

// TestFlowAsStateSource_NestedCaching verifies that derived states with unchanged values
// don't trigger downstream subscriber notifications.
func TestFlowAsStateSource_NestedCaching(t *testing.T) {
	root := flow.NewMutableStateFlow(10)

	calcB := 0
	derivedB := state.DerivedStateOf(func() int {
		calcB++
		v := root.Value()
		if v > 5 {
			return 1 // Always 1 if > 5
		}
		return 0
	})

	calcC := 0
	derivedC := state.DerivedStateOf(func() int {
		calcC++
		return derivedB.Get() + 100
	})

	// Track C's subscriber notifications
	cNotified := false
	derivedC.Subscribe(func() {
		cNotified = true
	})

	// Initial: root=10 -> B=1 -> C=101
	if got := derivedC.Get(); got != 101 {
		t.Errorf("Expected 101, got %d", got)
	}
	if calcB != 1 || calcC != 1 {
		t.Errorf("Expected calcB=1, calcC=1, got calcB=%d, calcC=%d", calcB, calcC)
	}

	// Update root to 20 (B still returns 1)
	cNotified = false
	root.Emit(20)

	if got := derivedC.Get(); got != 101 {
		t.Errorf("Expected 101, got %d", got)
	}
	if calcB != 2 {
		t.Errorf("B should have recalculated, calcB=%d", calcB)
	}
	if calcC != 2 {
		t.Errorf("C should have recalculated to verify value, calcC=%d", calcC)
	}
	// C's value didn't change, so subscribers should NOT be notified
	if cNotified {
		t.Error("C's subscribers should NOT be notified when C's value didn't change")
	}
}

// TestFlowAsStateSource_Subscribe verifies direct subscription to flow value changes.
func TestFlowCollectedAsState_Subscribe(t *testing.T) {
	flowValue := flow.NewMutableStateFlow(1)

	notifyCount := 0
	sub := flowValue.Subscribe(func() {
		notifyCount++
	})

	// Initial subscribe should not trigger (no emission yet)
	if notifyCount != 0 {
		t.Errorf("Expected 0 notifications, got %d", notifyCount)
	}

	// Emit should notify
	flowValue.Emit(2)
	if notifyCount != 1 {
		t.Errorf("Expected 1 notification, got %d", notifyCount)
	}

	// Update should notify
	flowValue.Update(func(v int) int { return v + 1 })
	if notifyCount != 2 {
		t.Errorf("Expected 2 notifications, got %d", notifyCount)
	}

	// Unsubscribe
	sub.Unsubscribe()

	// Should not notify after unsubscribe
	flowValue.Emit(100)
	if notifyCount != 2 {
		t.Errorf("Expected 2 notifications after unsubscribe, got %d", notifyCount)
	}
}

// End-to-end tests simulating CollectStateFlowAsState.
// These tests use actual MutableStateFlow and simulate the collection into MutableValue.

// collectFlowToState simulates what CollectStateFlowAsState does:
// subscribes to a flow and updates a MutableValue when the flow emits.
func collectFlowToState[T any](flowValue *flow.MutableStateFlow[T], initialValue T) *store.MutableValue {
	mv := store.NewMutableValue(initialValue, nil, func(a, b any) bool {
		return reflect.DeepEqual(a, b)
	})

	// Subscribe to flow changes and update the mutable value
	flowValue.Subscribe(func() {
		mv.Set(flowValue.Value())
	})

	return mv
}

// TestFlowToCollectedState_EndToEnd tests the actual flow -> collected state -> derived chain.
func TestFlowToCollectedState_EndToEnd(t *testing.T) {
	// Create an actual MutableStateFlow
	flowValue := flow.NewMutableStateFlow(10)

	// Simulate CollectStateFlowAsState: collect flow into mutable value
	collectedState := collectFlowToState(flowValue, flowValue.Value())

	// Create derived state that reads from the collected state
	calculatedCalls := 0
	derived := state.DerivedStateOf(func() int {
		calculatedCalls++
		return collectedState.Get().(int) * 2
	})

	// Initial calculation
	if got := derived.Get(); got != 20 {
		t.Errorf("Expected 20, got %d", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// Cached
	if got := derived.Get(); got != 20 {
		t.Errorf("Expected 20 (cached), got %d", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation (cached), got %d", calculatedCalls)
	}

	// Emit from flow - this should update the collected state and invalidate derived
	flowValue.Emit(15)

	// Derived should recalculate
	if got := derived.Get(); got != 30 {
		t.Errorf("Expected 30, got %d", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations, got %d", calculatedCalls)
	}
}

// TestFlowToCollectedState_MultipleFlows tests multiple flows collected into states.
func TestFlowToCollectedState_MultipleFlows(t *testing.T) {
	flowA := flow.NewMutableStateFlow(10)
	flowB := flow.NewMutableStateFlow(20)

	stateA := collectFlowToState(flowA, flowA.Value())
	stateB := collectFlowToState(flowB, flowB.Value())

	calculatedCalls := 0
	derived := state.DerivedStateOf(func() int {
		calculatedCalls++
		return stateA.Get().(int) + stateB.Get().(int)
	})

	// Initial: 10 + 20 = 30
	if got := derived.Get(); got != 30 {
		t.Errorf("Expected 30, got %d", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// Flow A emits
	flowA.Emit(15)
	if got := derived.Get(); got != 35 {
		t.Errorf("Expected 35, got %d", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations, got %d", calculatedCalls)
	}

	// Flow B emits
	flowB.Emit(30)
	if got := derived.Get(); got != 45 {
		t.Errorf("Expected 45, got %d", got)
	}
	if calculatedCalls != 3 {
		t.Errorf("Expected 3 calculations, got %d", calculatedCalls)
	}
}

// TestFlowToCollectedState_ChainedDerived tests flow -> state -> derived -> derived.
func TestFlowToCollectedState_ChainedDerived(t *testing.T) {
	flowValue := flow.NewMutableStateFlow(1)
	collectedState := collectFlowToState(flowValue, flowValue.Value())

	derivedA := state.DerivedStateOf(func() int {
		return collectedState.Get().(int) + 10
	})

	derivedB := state.DerivedStateOf(func() int {
		return derivedA.Get() + 100
	})

	// Initial: (1 + 10) + 100 = 111
	if got := derivedB.Get(); got != 111 {
		t.Errorf("Expected 111, got %d", got)
	}

	// Flow emits
	flowValue.Emit(2)

	// Should propagate: (2 + 10) + 100 = 112
	if got := derivedB.Get(); got != 112 {
		t.Errorf("Expected 112, got %d", got)
	}
}

// TestFlowToCollectedState_Update tests flow.Update() triggering derived recalculation.
func TestFlowToCollectedState_Update(t *testing.T) {
	flowValue := flow.NewMutableStateFlow(5)
	collectedState := collectFlowToState(flowValue, flowValue.Value())

	calculatedCalls := 0
	derived := state.DerivedStateOf(func() int {
		calculatedCalls++
		return collectedState.Get().(int) * 10
	})

	// Initial
	if got := derived.Get(); got != 50 {
		t.Errorf("Expected 50, got %d", got)
	}
	if calculatedCalls != 1 {
		t.Errorf("Expected 1 calculation, got %d", calculatedCalls)
	}

	// Use Update (not Emit)
	flowValue.Update(func(v int) int { return v + 5 })

	// Should recalculate
	if got := derived.Get(); got != 100 {
		t.Errorf("Expected 100, got %d", got)
	}
	if calculatedCalls != 2 {
		t.Errorf("Expected 2 calculations, got %d", calculatedCalls)
	}
}
