package state

type StateOption func(*StateOptions)

type StateOptions struct {
	Compare func(any, any) bool
}

type SupportState interface {
	Remember(key string, calc func() any) any                                  // transient state
	State(key string, initial func() any, options ...StateOption) MutableValue // persistent state
}
