package state

// InvalidationNotifier is an optional interface for derived states that support
// a separate subscription for invalidation events. This allows derived states
// to subscribe to invalidation events (for chain propagation) rather than
// value-change events (for user callbacks).
type InvalidationNotifier interface {
	// SubscribeForInvalidation registers a callback to be invoked when this
	// state is invalidated (not when its value changes).
	SubscribeForInvalidation(callback func()) Subscription
}
