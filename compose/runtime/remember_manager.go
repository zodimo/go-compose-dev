// Package runtime provides core composition runtime interfaces ported from Jetpack Compose.
package runtime

// RememberManager is an interface used during ControlledComposition.ApplyChanges
// and Composition.Dispose to track when RememberObserver instances enter and
// leave the composition, and also allows recording SideEffect calls.
type RememberManager interface {
	// Remembering is called when the RememberObserver is being remembered
	// by a slot in the slot table.
	Remembering(instance *RememberObserverHolder)

	// Forgetting is called when the RememberObserver is being forgotten
	// by a slot in the slot table.
	Forgetting(instance *RememberObserverHolder)

	// SideEffect schedules the effect to be called when changes are being applied
	// but after the remember/forget notifications are sent.
	SideEffect(effect func())

	// Deactivating is called when the ComposeNodeLifecycleCallback is being deactivated.
	Deactivating(instance ComposeNodeLifecycleCallback)

	// Releasing is called when the ComposeNodeLifecycleCallback is being released.
	Releasing(instance ComposeNodeLifecycleCallback)

	// RememberPausingScope is called when the restart scope is pausing.
	RememberPausingScope(scope RecomposeScope)

	// StartResumingScope is called when the restart scope is resuming.
	StartResumingScope(scope RecomposeScope)

	// EndResumingScope is called when the restart scope is finished resuming.
	EndResumingScope(scope RecomposeScope)
}
