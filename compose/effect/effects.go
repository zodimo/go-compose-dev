package effect

import (
	"context"
	"fmt"
	"go-compose-dev/internal/layoutnode"
	"go-compose-dev/pkg/api"
	"reflect"
)

// LaunchedEffect runs a side-effect in a goroutine.
// The effect is restarted if any of the keys change.
func LaunchedEffect(block func(context.Context), keys ...any) api.Composable {
	return func(c api.Composer) api.Composer {
		c.StartBlock("LaunchedEffect")

		// Key for the persistent state, ensuring uniqueness per node using its ID.
		// Since State() might be global in this implementation, we need to scope it manually.
		stateKey := fmt.Sprintf("launched_effect_%d", c.GetID().Value())

		effectState := c.State(stateKey, func() any {
			return &launchEffectState{}
		})

		state := effectState.Get().(*launchEffectState)

		// Check if keys changed
		keysChanged := !reflect.DeepEqual(state.lastKeys, keys)

		if keysChanged {
			// Cancel previous
			if state.cancel != nil {
				state.cancel()
			}

			// Start new
			ctx, cancel := context.WithCancel(context.Background())
			state.cancel = cancel
			// Copy keys to ensure we store a snapshot (though variadic slice is usually fresh)
			keysCopy := make([]any, len(keys))
			copy(keysCopy, keys)
			state.lastKeys = keysCopy

			go func() {
				block(ctx)
			}()
		}

		// Set a dummy widget constructor that does nothing (zero size)
		c.SetWidgetConstructor(layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
			return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {
				return layoutnode.LayoutDimensions{}
			}
		}))

		return c.EndBlock()
	}
}

type launchEffectState struct {
	cancel   context.CancelFunc
	lastKeys []any
}
