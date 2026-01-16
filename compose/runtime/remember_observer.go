// Package runtime provides core composition runtime interfaces ported from Jetpack Compose.
package runtime

// RememberObserver is notified when objects are remembered/forgotten in composition.
//
// An object is remembered by the composition if it is remembered in at least one
// place in a composition. An object is forgotten when it is no longer remembered
// anywhere in that composition.
//
// When objects implementing this interface are remembered and forgotten together,
// the order of OnForgotten is guaranteed to be called in the opposite order of
// OnRemembered.
//
// An object remembered in only one place in only one composition is guaranteed to:
//  1. have either OnRemembered or OnAbandoned called
//  2. if OnRemembered is called, OnForgotten will eventually be called
type RememberObserver interface {
	// OnRemembered is called when this object is successfully remembered by a composition.
	// This method is called on the composition's apply thread.
	OnRemembered()

	// OnForgotten is called when this object is forgotten by a composition.
	// This method is called on the composition's apply thread.
	OnForgotten()

	// OnAbandoned is called when this object is returned by the callback to remember
	// but is not successfully remembered by a composition.
	OnAbandoned()
}

// RememberObserverHolder wraps a RememberObserver with slot table position info.
type RememberObserverHolder struct {
	Wrapped         RememberObserver
	AfterGroupIndex int
}

// ReusableRememberObserverHolder is a RememberObserver which is not removed during
// reuse/deactivate of the group. Used to preserve composition locals between
// group deactivation.
type ReusableRememberObserverHolder struct {
	RememberObserverHolder
}

// NewRememberObserverHolder creates a new RememberObserverHolder.
func NewRememberObserverHolder(wrapped RememberObserver, afterGroupIndex int) *RememberObserverHolder {
	return &RememberObserverHolder{
		Wrapped:         wrapped,
		AfterGroupIndex: afterGroupIndex,
	}
}

// NewReusableRememberObserverHolder creates a new ReusableRememberObserverHolder.
func NewReusableRememberObserverHolder(wrapped RememberObserver, afterGroupIndex int) *ReusableRememberObserverHolder {
	return &ReusableRememberObserverHolder{
		RememberObserverHolder: RememberObserverHolder{
			Wrapped:         wrapped,
			AfterGroupIndex: afterGroupIndex,
		},
	}
}
