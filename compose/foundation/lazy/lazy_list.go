package lazy

import (
	"fmt"
	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/internal/layoutnode"
	"github.com/zodimo/go-compose/internal/modifier"

	"gioui.org/layout"
)

type C = layout.Context
type D = layout.Dimensions

// LazyColumn is a vertically scrolling list that only composes and lays out currently visible items.
// Note: Current implementation purely implements Lazy Layout (eager composition).
func LazyColumn(content func(LazyListScope), options ...LazyListOption) compose.Composable {
	return lazyList(layout.Vertical, content, options...)
}

// LazyRow is a horizontally scrolling list that only composes and lays out currently visible items.
// Note: Current implementation purely implements Lazy Layout (eager composition).
func LazyRow(content func(LazyListScope), options ...LazyListOption) compose.Composable {
	return lazyList(layout.Horizontal, content, options...)
}

func lazyList(axis layout.Axis, content func(LazyListScope), options ...LazyListOption) compose.Composable {
	return func(c compose.Composer) compose.Composer {
		opts := DefaultLazyListOptions()
		for _, opt := range options {
			opt(&opts)
		}

		// Ensure state is initialized
		// Note: Ideally state should be passed by user. If not, we create a local one,
		// but since we can't persist it easily without ID generation here (or using Remember internally),
		// we should encourage passing state.
		// For now, if state is nil, we create a new one but it won't be persisted across recompositions
		// unless the user provided a state from RememberLazyListState.
		// To fix this auto-persistence if nil, we would need to look it up in state.
		if opts.State == nil {
			id := c.GenerateID()
			path := c.GetPath()
			key := fmt.Sprintf("%d/%s/lazyListState", id, path)
			// Try to get existing or create new
			opts.State = c.State(key, func() any { return NewLazyListState() }).Get().(*LazyListState)
		}

		c.StartBlock("LazyList")
		c.Modifier(func(m modifier.Modifier) modifier.Modifier {
			return m.Then(opts.Modifier)
		})

		scope := &lazyListScopeImpl{}
		content(scope)

		// Emit all items as children
		for _, item := range scope.items {
			c.WithComposable(item.Content)
		}

		c.SetWidgetConstructor(lazyListWidgetConstructor(opts.State, axis))
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
