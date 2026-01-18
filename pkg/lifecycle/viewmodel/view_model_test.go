package viewmodel

import (
	"testing"

	"github.com/zodimo/go-compose/state"
	"github.com/zodimo/go-compose/store"
)

// mockComposer implements api.Composer for testing
type mockComposer struct {
	store state.PersistentState
}

func (m *mockComposer) State(key string, initial func() any, options ...state.StateOption) state.MutableValue {
	return m.store.State(key, initial, options...)
}

// mockViewModel tracks lifecycle
type mockViewModel struct {
	initialized bool
	cleared     bool
}

func NewMockViewModel() *mockViewModel {
	return &mockViewModel{initialized: true}
}

func (m *mockViewModel) OnCleared() {
	m.cleared = true
}

func TestViewModel_PersistsAcrossFrames(t *testing.T) {
	st := store.NewPersistentState(map[string]state.ScopedValue{})

	var vm1, vm2 *mockViewModel

	// Frame 1
	state.WithFrame(st, func() {
		mv := st.State("viewmodel_test", func() any { return NewMockViewModel() })
		vm1 = mv.Get().(*mockViewModel)
	})

	if vm1 == nil {
		t.Error("expected ViewModel vm1 to be created")
	}

	// Frame 2
	state.WithFrame(st, func() {
		mv := st.State("viewmodel_test", func() any {
			t.Error("factory should not be called on second access")
			return NewMockViewModel()
		})
		vm2 = mv.Get().(*mockViewModel)
	})

	if vm2 == nil {
		t.Error("expected ViewModel vm2 to be created")
	}

	if vm1 != vm2 {
		t.Error("expected same ViewModel instance across frames")
	}
}

func TestViewModel_OnClearedCalledWhenGCd(t *testing.T) {
	st := store.NewPersistentState(map[string]state.ScopedValue{})

	var vm *mockViewModel

	// Frame 1: create VM
	state.WithFrame(st, func() {
		mv := st.State("viewmodel_test", func() any { return NewMockViewModel() })
		vm = mv.Get().(*mockViewModel)
	})

	if vm == nil {
		t.Error("expected ViewModel to be created")
	}

	if vm.cleared {
		t.Error("OnCleared should not be called yet")
	}

	// Frame 2: don't access VM -> should be GC'd
	state.WithFrame(st, func() {
		// Don't access the ViewModel
	})

	if !vm.cleared {
		t.Error("expected OnCleared to be called when ViewModel is GC'd")
	}
}
