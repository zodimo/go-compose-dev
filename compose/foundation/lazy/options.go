package lazy

import (
	"go-compose-dev/internal/modifier"
)

type LazyListOption func(*LazyListOptions)

type LazyListOptions struct {
	Modifier modifier.Modifier
	State    *LazyListState
}

func DefaultLazyListOptions() LazyListOptions {
	return LazyListOptions{
		Modifier: modifier.EmptyModifier,
		State:    nil,
	}
}

func WithModifier(m modifier.Modifier) LazyListOption {
	return func(o *LazyListOptions) {
		o.Modifier = m
	}
}

func WithState(state *LazyListState) LazyListOption {
	return func(o *LazyListOptions) {
		o.State = state
	}
}
