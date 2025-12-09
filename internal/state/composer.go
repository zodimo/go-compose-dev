package state

type StatefulComposer interface {
	Remember(key string, calc func() any) any          // transient state
	State(key string, initial func() any) MutableValue // persistent state
}
