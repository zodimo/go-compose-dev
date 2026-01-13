package window

import (
	"image"

	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/internal/layoutnode"
	"github.com/zodimo/go-compose/pkg/api"

	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
)

// PopupAlignment determines how the popup is aligned relative to its anchor point.
// For now, we assume the anchor point is the top-left of where this Popup is called.
type PopupAlignment int

const (
	AlignTopLeft PopupAlignment = iota
)

// Popup shows content overlaid on top of other content.
// It uses op.Defer to effectively "break out" of the z-order (though it respects parent clipping).
// The content is laid out with loose constraints (0 to Max).
func Popup(
	content api.Composable,
	options ...PopupOption,
) api.Composable {
	return func(c api.Composer) api.Composer {
		opts := DefaultPopupOptions()
		for _, option := range options {
			if option == nil {
				continue
			}
			option(&opts)
		}

		c.StartBlock("Popup")
		c.WithComposable(content)
		c.SetWidgetConstructor(popupWidgetConstructor(opts))
		return c.EndBlock()
	}
}

func popupWidgetConstructor(opts PopupOptions) layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {
			// Persistent click gesture for blocking pointer events from passing through
			popupClick := node.State("popupClick", func() any { return &gesture.Click{} }).Get().(*gesture.Click)

			// 1. Record the content layout
			macro := op.Record(gtx.Ops)

			// 2. Prepare context for popup content
			pGtx := gtx
			pGtx.Constraints.Min = image.Point{}

			// Apply offset if needed
			xPx := pGtx.Dp(unit.DpToGioUnit(opts.OffsetX))
			yPx := pGtx.Dp(unit.DpToGioUnit(opts.OffsetY))

			op.Offset(image.Pt(xPx, yPx)).Add(pGtx.Ops)

			// Layout children and track content size
			var contentSize image.Point
			for _, child := range node.Children() {
				childLayoutNode := child.(layoutnode.NodeCoordinator)
				dims := childLayoutNode.Layout(pGtx)
				if dims.Size.X > contentSize.X {
					contentSize.X = dims.Size.X
				}
				if dims.Size.Y > contentSize.Y {
					contentSize.Y = dims.Size.Y
				}
			}

			// Add blocking handler at the popup content area
			// This handler catches clicks and prevents them from passing through to
			// elements below the popup. Since it's registered AFTER content handlers
			// (added during child layout), the content handlers take priority.
			// We use PassOp so this doesn't block the content handlers we just added.
			if contentSize.X > 0 && contentSize.Y > 0 {
				// Push PassOp - handlers added in this scope won't block siblings
				passStack := pointer.PassOp{}.Push(pGtx.Ops)

				// Add blocking click handler covering content area
				stack := clip.Rect{Max: contentSize}.Push(pGtx.Ops)
				popupClick.Add(pGtx.Ops)
				event.Op(pGtx.Ops, node)
				stack.Pop()

				passStack.Pop()
			}

			// 3. Stop recording
			call := macro.Stop()

			// 4. Defer the execution
			op.Defer(gtx.Ops, call)

			// 5. Return empty dimensions
			return layout.Dimensions{}
		}
	})
}
