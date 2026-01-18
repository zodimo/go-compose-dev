package state

import (
	"context"
)

type ScopedValue struct {
	id           string
	mv           MutableValue
	subscription Subscription
	cancelFunc   context.CancelFunc // Context cancel functions for ViewModelScope

	onForgotten func()
	onCleared   func()
}

func NewScopedValue(
	id string,
	mv MutableValue,
	subscription Subscription,
	cancelFunc context.CancelFunc,
	onForgotten func(),
	onCleared func(),
) ScopedValue {
	return ScopedValue{
		id:           id,
		mv:           mv,
		subscription: subscription,
		cancelFunc:   cancelFunc,
		onForgotten:  onForgotten,
		onCleared:    onCleared,
	}
}

func (sv ScopedValue) Cleanup() {
	sv.cancelFunc()
	sv.subscription.Unsubscribe()
	sv.onForgotten()
	sv.onCleared()
}

func (sv ScopedValue) MutableValue() MutableValue {
	return sv.mv
}
