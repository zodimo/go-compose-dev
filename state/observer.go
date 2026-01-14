package state

type ReadObserver func(source any)

var activeReadObserver ReadObserver

// NotifyRead notifies the active read observer that a state object has been read.
// Source is typically the state object itself (e.g. *MutableValue).
func NotifyRead(source any) {
	if activeReadObserver != nil {
		activeReadObserver(source)
	}
}

// WithReadObserver executes the block with the given read observer.
// It restores the previous observer after the block finishes.
func WithReadObserver(observer ReadObserver, block func()) {
	prev := activeReadObserver
	activeReadObserver = observer
	defer func() {
		activeReadObserver = prev
	}()
	block()
}
