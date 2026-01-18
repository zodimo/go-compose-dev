package state

func WithCompare(compare func(any, any) bool) StateOption {
	if compare == nil {
		panic("WithCompare: compare cannot be nil")
	}
	return func(opts *StateOptions) {
		opts.Compare = compare
	}
}

type PersistentState interface {
	StateChangeNotifier

	// createScopedMemo(scope string) MemoState[any]

	// Remember(key string, calc func() any) any

	State(key string, initial func() any, options ...StateOption) MutableValue

	// StartFrame marks the beginning of a composition frame.
	// All state keys accessed after this call and before EndFrame will be retained.
	StartFrame()

	// EndFrame marks the end of a composition frame.
	// Any state keys that were NOT accessed during the frame will be garbage collected.
	// If the value implements RememberObserver (OnForgotten) or ViewModel (OnCleared),
	// the appropriate cleanup method will be called before removal.
	EndFrame()
}

// WithFrame wraps a composition function with StartFrame/EndFrame calls.
// This is the recommended way to ensure proper lifecycle management.]]
func WithFrame(store PersistentState, compose func()) {
	store.StartFrame()
	defer store.EndFrame()
	compose()
}
