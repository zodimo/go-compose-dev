package snackbar

import (
	"go-compose-dev/internal/layoutnode"

	"gioui.org/layout"
	"git.sr.ht/~schnwalter/gio-mw/wdk/block"
	"git.sr.ht/~schnwalter/gio-mw/widget/overlay"
	"git.sr.ht/~schnwalter/gio-mw/widget/snackbar"
)

const SnackbarHostNodeID = "SnackbarHost"

type SnackbarHostState struct {
	overlay *overlay.Overlay
}

func NewSnackbarHostState() *SnackbarHostState {
	return &SnackbarHostState{
		overlay: &overlay.Overlay{},
	}
}

func (s *SnackbarHostState) ShowSnackbar(message string, options ...SnackbarOption) {
	opts := DefaultOptions()
	for _, opt := range options {
		opt(&opts)
	}
	// Simple text snackbar for now
	snackStyle := snackbar.Plain(message)
	// Default duration of 4 seconds
	item := overlay.NewItem(snackStyle.Layout, block.GravityBottomCenter).WithDuration(opts.Duration)
	s.overlay.Show(item)
}

func SnackbarHost(hostState *SnackbarHostState) Composable {
	return func(c Composer) Composer {
		c.StartBlock(SnackbarHostNodeID)

		c.SetWidgetConstructor(snackbarHostWidgetConstructor(hostState))
		return c.EndBlock()
	}
}

func snackbarHostWidgetConstructor(hostState *SnackbarHostState) layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {
			// We need to call Update on the overlay
			hostState.overlay.Update(gtx)

			// Then Layout
			hostState.overlay.Layout(gtx)

			// Overlay usually takes up the max constraints or just renders on top.
			// gio-mw overlay.Layout doc says it iterates and layouts items.
			// It doesn't seem to return Dimensions, so we return empty or max dimensions?
			// Looking at gio-mw code, Overlay.Layout signature is `func (o *Overlay) Layout(gtx layout.Context)`.
			// It doesn't return Dimensions.
			// However, our widget constructor must return Dimensions.
			// Since it is an overlay, it likely draws on top and doesn't affect flow layout size,
			// but usually it is placed in a stack or similar.
			// For the purpose of the Host, we probably want it to take up available space if it's handling positioning,
			// or just 0 size if it's purely painting.
			// But gio-mw overlay items handle their own gravity/positioning.

			return layout.Dimensions{}
		}
	})
}
