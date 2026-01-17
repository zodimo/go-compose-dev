package store

import (
	"reflect"

	"github.com/zodimo/go-compose/state"
)

type PersistentStateInterface = state.PersistentState

type PersistentState struct {
	scopes        map[string]MutableValueInterface
	onStateChange func()
}

func NewPersistentState(scopes map[string]MutableValueInterface) PersistentStateInterface {
	return &PersistentState{scopes: scopes}
}

func (ps *PersistentState) SetOnStateChange(callback func()) {
	ps.onStateChange = callback
}

func (ps *PersistentState) GetState(id string, initial func() any, options ...state.StateOption) MutableValueInterface {

	opts := state.StateOptions{
		Compare: reflect.DeepEqual,
	}
	for _, option := range options {
		option(&opts)
	}

	if v, ok := ps.scopes[id]; ok {
		return v
	}
	ps.scopes[id] = &MutableValue{
		cell: initial(),
		changeNotifier: func(any) {
			if ps.onStateChange != nil {
				ps.onStateChange()
			}
		},
		compare:     opts.Compare,
		subscribers: state.NewSubscriptionManager(),
	}
	return ps.scopes[id]
}
