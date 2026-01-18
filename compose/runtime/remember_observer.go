// Package runtime provides core compose runtime interfaces and utilities.
package runtime

// RememberObserver is an interface for objects that need lifecycle callbacks
// when remembered in composition. This is the Go equivalent of
// androidx.compose.runtime.RememberObserver.
//
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/runtime/runtime/src/commonMain/kotlin/androidx/compose/runtime/RememberObserver.kt
type RememberObserver interface {
	// OnRemembered is called when the object is successfully stored by remember.
	// This is called after the composition where the object was remembered completes.
	OnRemembered()

	// OnForgotten is called when the object is no longer being remembered.
	// This happens when the remember call is removed from composition.
	OnForgotten()

	// OnAbandoned is called when the remember call was not committed to the composition.
	// This can happen if the composition is abandoned before it is applied.
	OnAbandoned()
}
