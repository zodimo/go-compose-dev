package store

import (
	"context"
	"reflect"
	"sync"

	"github.com/zodimo/go-compose/compose/runtime"
	"github.com/zodimo/go-compose/pkg/lifecycle"
	"github.com/zodimo/go-compose/state"
)

type PersistentStateInterface = state.PersistentState

type PersistentState struct {
	mu     sync.Mutex
	scopes map[string]state.ScopedValue
	// onRemeber lifecycle
	accessedKeys map[string]struct{} // Keys accessed during current frame
	frameActive  bool                // True when inside a StartFrame/EndFrame block
	previousMemo map[string]any      // Memo from previous frame for lifecycle

	subscribers *state.SubscriptionManager
}

func NewPersistentState(scopes map[string]state.ScopedValue) PersistentStateInterface {
	return &PersistentState{
		scopes:       scopes,
		accessedKeys: make(map[string]struct{}),
		frameActive:  false,
		subscribers:  state.NewSubscriptionManager(),
	}
}

func (ps *PersistentState) Subscribe(callback func()) state.Subscription {
	return ps.subscribers.Subscribe(callback)
}

// StartFrame marks the beginning of a composition frame.
// It resets the set of accessed keys to prepare for tracking during the frame.
func (ps *PersistentState) StartFrame() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if ps.frameActive {
		panic("StartFrame called while frame already active - use state.WithFrame() instead of manual StartFrame/EndFrame")
	}
	ps.frameActive = true
	//remeber observer onremembered
	ps.accessedKeys = make(map[string]struct{})
}

// EndFrame marks the end of a composition frame.
// It garbage collects any state that was not accessed during this frame.
// Lifecycle callbacks (OnForgotten for RememberObserver, OnCleared for ViewModel)
// are called before removal.
func (ps *PersistentState) EndFrame() {
	ps.mu.Lock()
	if !ps.frameActive {
		ps.mu.Unlock()
		panic("EndFrame called without StartFrame - use state.WithFrame() to ensure correct lifecycle")
	}
	ps.frameActive = false

	keysToRemove := []string{}

	for key := range ps.scopes {
		if _, accessed := ps.accessedKeys[key]; !accessed {
			keysToRemove = append(keysToRemove, key)
		}
	}

	toGC := make([]state.ScopedValue, 0, len(keysToRemove))

	for _, key := range keysToRemove {
		scopedValue := ps.scopes[key]

		delete(ps.scopes, key)
		toGC = append(toGC, scopedValue)
	}
	ps.mu.Unlock()

	// Call lifecycle callbacks outside the lock to avoid deadlocks
	for _, scopedValue := range toGC {
		scopedValue.Cleanup()
	}
}

func (ps *PersistentState) State(id string, initial func() any, options ...state.StateOption) state.MutableValue {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if !ps.frameActive {
		panic("GetState called outside composition frame - wrap your composition in state.WithFrame()")
	}

	opts := state.StateOptions{
		Compare: reflect.DeepEqual,
	}
	for _, option := range options {
		option(&opts)
	}

	// Mark this key as accessed for GC tracking
	ps.accessedKeys[id] = struct{}{}

	if v, ok := ps.scopes[id]; ok {
		return v.MutableValue()
	}

	newValue := initial()

	ps.initScopedValue(id, newValue, opts.Compare)

	return ps.scopes[id].MutableValue()
}

func (ps *PersistentState) initScopedValue(id string, initialValue any, compare func(any, any) bool) {

	scopedCancelFunc := func() {}
	onForgotten := func() {}
	onCleared := func() {}

	// If implements HasViewModelScope, create and provide a cancellable context
	if scopeHolder, ok := initialValue.(lifecycle.HasViewModelScope); ok {
		ctx, cancelFunc := context.WithCancel(context.Background())
		scopeHolder.SetViewModelScope(ctx)
		scopedCancelFunc = cancelFunc
	}

	if observer, ok := initialValue.(runtime.RememberObserver); ok {
		onForgotten = observer.OnForgotten
	}
	if vm, ok := initialValue.(lifecycle.ViewModel); ok {
		onCleared = vm.OnCleared
	}

	mv := state.NewMutableValue(initialValue, compare)

	scopedValue := state.NewScopedValue(
		id,
		mv,
		ps.subscribers.Subscribe(func() {
			ps.subscribers.NotifyAll()
		}),
		scopedCancelFunc,
		onForgotten,
		onCleared,
	)

	ps.scopes[id] = scopedValue

	// Call RememberObserver.OnRemembered if implemented
	// TODO: review lifecycle event
	if observer, ok := initialValue.(runtime.RememberObserver); ok {
		observer.OnRemembered()
	}

}

// StoreMemo saves the current frame's memo for lifecycle comparison in the next frame.
func (ps *PersistentState) StoreMemo(memo map[string]any) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.previousMemo = memo
}

// GetPreviousMemo returns the memo from the previous frame.
func (ps *PersistentState) GetPreviousMemo() map[string]any {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if ps.previousMemo == nil {
		return make(map[string]any)
	}
	return ps.previousMemo
}
