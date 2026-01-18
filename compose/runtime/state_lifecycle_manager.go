package runtime

// LifecycleToken is proof that memo lifecycle was processed.
// It can only be created by ProcessMemoLifecycle.
// The private field prevents external instantiation.
type LifecycleToken struct {
	_ struct{}
}

// ProcessMemoLifecycle processes RememberObserver callbacks and returns a token.
// Composers MUST call this and pass the token to Build().
//
// This function:
// 1. Compares currentMemo with previousMemo
// 2. Calls OnForgotten() for entries that were in previous but not in current
// 3. Calls OnRemembered() for entries that are in current but not in previous
// 4. Stores the current memo for the next frame
func ProcessMemoLifecycle(previous, current map[string]any) LifecycleToken {
	// OnForgotten: in previous but not in current
	for key, val := range previous {
		if _, exists := current[key]; !exists {
			if observer, ok := val.(RememberObserver); ok {
				observer.OnForgotten()
			}
		}
	}

	// OnRemembered: in current but not in previous
	for key, val := range current {
		if _, exists := previous[key]; !exists {
			if observer, ok := val.(RememberObserver); ok {
				observer.OnRemembered()
			}
		}
	}

	return LifecycleToken{}
}
