package store

import (
	"reflect"

	"github.com/zodimo/go-compose/state"
)

type MutableValueInterface = state.MutableValue
type PersistentStateInterface = state.PersistentState
type TypedMutableValueInterface[T any] = state.TypedMutableValue[T]

var _ MutableValueInterface = &MutableValue{}
var _ MutableValueInterface = &MutableValueTypedWrapper[any]{}

var _ state.StateChangeNotifier = &MutableValue{}
var _ state.StateChangeNotifier = &MutableValueTypedWrapper[any]{}

var _ TypedMutableValueInterface[any] = &MutableValueTypedWrapper[any]{}

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
