// Package runtime provides core composition runtime interfaces ported from Jetpack Compose.
package runtime

// RecomposeScope represents a recomposable scope or section of the composition hierarchy.
// Can be used to manually invalidate the scope to schedule it for recomposition.
type RecomposeScope interface {
	// Invalidate the corresponding scope, requesting the composer recompose this scope.
	// This method is thread safe.
	Invalidate()
}

// RecomposeScopeOwner is implemented by compositions that own recompose scopes.
type RecomposeScopeOwner interface {
	// Invalidate schedules the scope to be recomposed.
	Invalidate(scope RecomposeScope, instance any) InvalidationResult

	// RecomposeScopeReleased is called when the recompose scope has been removed
	// from the composition.
	RecomposeScopeReleased(scope RecomposeScope)

	// RecordReadOf records that a value has been read in the current scope.
	RecordReadOf(value any)
}

// InvalidationResult represents the result of invalidating a recompose scope.
type InvalidationResult int

const (
	// InvalidationIgnored indicates the invalidation was ignored because the
	// associated recompose scope is no longer part of the composition or has
	// yet to be entered in the composition.
	InvalidationIgnored InvalidationResult = iota

	// InvalidationScheduled indicates the composition is not currently composing
	// and the invalidation was recorded for a future composition.
	InvalidationScheduled

	// InvalidationDeferred indicates the composition is actively composing but
	// the scope has already been composed or is in the process of composing.
	InvalidationDeferred

	// InvalidationImminent indicates the composition is actively composing and
	// the invalidated scope has not been composed yet but will be recomposed
	// before the composition completes.
	InvalidationImminent
)

// String returns a string representation of the InvalidationResult.
func (r InvalidationResult) String() string {
	switch r {
	case InvalidationIgnored:
		return "IGNORED"
	case InvalidationScheduled:
		return "SCHEDULED"
	case InvalidationDeferred:
		return "DEFERRED"
	case InvalidationImminent:
		return "IMMINENT"
	default:
		return "UNKNOWN"
	}
}
