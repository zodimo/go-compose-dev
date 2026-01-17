package flow

import (
	"context"
	"fmt"

	"github.com/zodimo/go-compose/compose/effect"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/state"
)

// CollectAsState collects values from a Flow and represents it as a State.
// The initial value is used until the first value is emitted by the flow.
func CollectAsState[T any](c api.Composer, key string, flow Flow[T], initial T, options ...state.StateTypedOption[T]) state.ValueTyped[T] {

	// uniqueness
	genID := c.GenerateID()

	// 1. Create a state to hold the latest value
	stateValue := state.StateUnsafe[T](c, fmt.Sprintf("flow_state_%s_%s", genID, key), func() T { return initial }, options...)

	// 2. Launch a side-effect to collect the flow
	// We use the flow itself as a key so if the flow instance changes, we resubscribe
	effect.LaunchedEffect(func(ctx context.Context) {
		flow.Collect(ctx, func(value T) {
			stateValue.Set(value)
		})
	}, flow)(c)

	return stateValue
}

// CollectStateFlowAsState collects values from a StateFlow.
// It uses `flow.Value()` as the initial value.
func CollectStateFlowAsState[T any](c api.Composer, key string, flow StateFlow[T], options ...state.StateTypedOption[T]) state.ValueTyped[T] {
	return CollectAsState(c, key, flow, flow.Value(), options...)
}
