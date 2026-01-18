package store

import (
	"context"
	"testing"

	"github.com/zodimo/go-compose/state"
)

// mockRememberObserver tracks lifecycle calls for testing
type mockRememberObserver struct {
	onRememberedCalled bool
	onForgottenCalled  bool
	onAbandonedCalled  bool
}

func (m *mockRememberObserver) OnRemembered() { m.onRememberedCalled = true }
func (m *mockRememberObserver) OnForgotten()  { m.onForgottenCalled = true }
func (m *mockRememberObserver) OnAbandoned()  { m.onAbandonedCalled = true }

// mockViewModel tracks OnCleared calls for testing
type mockViewModel struct {
	onClearedCalled bool
}

func (m *mockViewModel) OnCleared() { m.onClearedCalled = true }

func TestGC_StateRetainedIfAccessed(t *testing.T) {
	store := NewPersistentState(map[string]state.ScopedValue{})

	store.StartFrame()
	store.State("key1", func() any { return "value1" })
	store.EndFrame()

	// Access again in next frame
	store.StartFrame()
	mv := store.State("key1", func() any { return "should not be called" })
	store.EndFrame()

	if mv.Get() != "value1" {
		t.Errorf("expected state to be retained, got %v", mv.Get())
	}
}

func TestGC_StateRemovedIfNotAccessed(t *testing.T) {
	store := NewPersistentState(map[string]state.ScopedValue{})

	store.StartFrame()
	store.State("key1", func() any { return "value1" })
	store.EndFrame()

	// Frame 2: don't access key1
	store.StartFrame()
	store.EndFrame()

	// Frame 3: key1 should be gone, so initial should be called again
	store.StartFrame()
	called := false
	mv := store.State("key1", func() any {
		called = true
		return "new_value"
	})
	store.EndFrame()

	if !called {
		t.Error("expected initial func to be called because key was garbage collected")
	}
	if mv.Get() != "new_value" {
		t.Errorf("expected new_value, got %v", mv.Get())
	}
}

func TestGC_OnForgottenCalledForRememberObserver(t *testing.T) {
	store := NewPersistentState(map[string]state.ScopedValue{})

	observer := &mockRememberObserver{}

	store.StartFrame()
	store.State("observer_key", func() any { return observer })
	store.EndFrame()

	if !observer.onRememberedCalled {
		t.Error("expected OnRemembered to be called")
	}

	// Frame 2: don't access observer_key -> should be GC'd
	store.StartFrame()
	store.EndFrame()

	if !observer.onForgottenCalled {
		t.Error("expected OnForgotten to be called on GC")
	}
}

func TestGC_OnClearedCalledForViewModel(t *testing.T) {
	store := NewPersistentState(map[string]state.ScopedValue{})

	vm := &mockViewModel{}

	store.StartFrame()
	store.State("vm_key", func() any { return vm })
	store.EndFrame()

	// Frame 2: don't access vm_key -> should be GC'd
	store.StartFrame()
	store.EndFrame()

	if !vm.onClearedCalled {
		t.Error("expected OnCleared to be called on GC")
	}
}

func TestGC_MultipleKeysPartialRetention(t *testing.T) {
	store := NewPersistentState(map[string]state.ScopedValue{})

	store.StartFrame()
	store.State("keep", func() any { return "keep_value" })
	store.State("remove", func() any { return "remove_value" })
	store.EndFrame()

	// Frame 2: only access "keep"
	store.StartFrame()
	store.State("keep", func() any { return "should not call" })
	store.EndFrame()

	// Frame 3: verify "keep" exists and "remove" was GC'd
	store.StartFrame()
	keepMv := store.State("keep", func() any { return "BAD" })
	removeCalled := false
	removeMv := store.State("remove", func() any {
		removeCalled = true
		return "new_remove_value"
	})
	store.EndFrame()

	if keepMv.Get() != "keep_value" {
		t.Errorf("expected keep_value, got %v", keepMv.Get())
	}
	if !removeCalled {
		t.Error("expected remove key to be garbage collected")
	}
	if removeMv.Get() != "new_remove_value" {
		t.Errorf("expected new_remove_value, got %v", removeMv.Get())
	}
}

func TestPanic_GetStateOutsideFrame(t *testing.T) {
	store := NewPersistentState(map[string]state.ScopedValue{})

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when GetState called outside frame")
		}
	}()

	// This should panic
	store.State("key", func() any { return "value" })
}

func TestPanic_DoubleStartFrame(t *testing.T) {
	store := NewPersistentState(map[string]state.ScopedValue{})

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on double StartFrame")
		}
	}()

	store.StartFrame()
	store.StartFrame() // This should panic
}

func TestPanic_EndFrameWithoutStart(t *testing.T) {
	store := NewPersistentState(map[string]state.ScopedValue{})

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on EndFrame without StartFrame")
		}
	}()

	store.EndFrame() // This should panic
}

func TestWithFrame_Success(t *testing.T) {
	store := NewPersistentState(map[string]state.ScopedValue{})

	state.WithFrame(store, func() {
		store.State("key", func() any { return "value" })
	})

	// Should not panic, and state should be accessible in next frame
	state.WithFrame(store, func() {
		mv := store.State("key", func() any { return "should not call" })
		if mv.Get() != "value" {
			t.Errorf("expected value, got %v", mv.Get())
		}
	})
}

// mockViewModelWithScope tracks context lifecycle
type mockViewModelWithScope struct {
	ctx             context.Context
	onClearedCalled bool
}

func (m *mockViewModelWithScope) SetViewModelScope(ctx context.Context) {
	m.ctx = ctx
}

func (m *mockViewModelWithScope) OnCleared() {
	m.onClearedCalled = true
}

func TestViewModelScope_ContextCancelledBeforeOnCleared(t *testing.T) {
	st := NewPersistentState(map[string]state.ScopedValue{})

	vm := &mockViewModelWithScope{}

	// Frame 1: create VM with scope
	state.WithFrame(st, func() {
		st.State("vm_scope_key", func() any { return vm })
	})

	if vm.ctx == nil {
		t.Fatal("expected context to be set via SetViewModelScope")
	}

	select {
	case <-vm.ctx.Done():
		t.Error("context should not be cancelled yet")
	default:
		// Good, context is still active
	}

	// Frame 2: don't access VM -> should be GC'd
	state.WithFrame(st, func() {
		// Don't access the ViewModel
	})

	select {
	case <-vm.ctx.Done():
		// Good, context was cancelled
	default:
		t.Error("expected context to be cancelled when VM is GC'd")
	}

	if !vm.onClearedCalled {
		t.Error("expected OnCleared to be called after context cancellation")
	}
}
