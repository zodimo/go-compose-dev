package lazy

import (
	"go-compose-dev/compose"
	"go-compose-dev/internal/layoutnode"
	"go-compose-dev/internal/modifier"

	"gioui.org/layout"
)

type C = layout.Context
type D = layout.Dimensions

// LazyColumn is a vertically scrolling list that only composes and lays out currently visible items.
// Note: Current implementation purely implements Lazy Layout (eager composition).
func LazyColumn(mod modifier.Modifier, state *LazyListState, content func(LazyListScope)) compose.Composable {
	return lazyList(mod, state, layout.Vertical, content)
}

// LazyRow is a horizontally scrolling list that only composes and lays out currently visible items.
// Note: Current implementation purely implements Lazy Layout (eager composition).
func LazyRow(mod modifier.Modifier, state *LazyListState, content func(LazyListScope)) compose.Composable {
	return lazyList(mod, state, layout.Horizontal, content)
}

func lazyList(mod modifier.Modifier, state *LazyListState, axis layout.Axis, content func(LazyListScope)) compose.Composable {
	return func(c compose.Composer) compose.Composer {
		c.StartBlock("LazyList")
		c.Modifier(func(m modifier.Modifier) modifier.Modifier {
			return m.Then(mod)
		})

		scope := &lazyListScopeImpl{}
		content(scope)

		// Emit all items as children
		// This creates the LayoutNodes for all items in the list.
		// Layout will use layout.List to only measure and draw visible ones.
		for _, item := range scope.items {
			c.WithComposable(item.Content)
		}

		c.SetWidgetConstructor(lazyListWidgetConstructor(state, axis))
		return c.EndBlock()
	}
}

func lazyListWidgetConstructor(state *LazyListState, axis layout.Axis) layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {
			// Update axis configuration
			state.List.List.Axis = axis

			return state.List.List.Layout(gtx, len(node.Children()), func(gtx C, i int) D {
				if i < 0 || i >= len(node.Children()) {
					return D{}
				}
				child := node.Children()[i].(layoutnode.NodeCoordinator)
				return child.Layout(gtx)
			})
		}
	})
}
