package overlay

import (
	"image"

	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/internal/layoutnode"

	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func Overlay(content Composable, options ...OverlayOption) Composable {
	opts := DefaultOverlayOptions()
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&opts)
	}
	return func(c Composer) Composer {

		theme := material3.Theme(c)
		opts.ScrimColor = opts.ScrimColor.TakeOrElse(graphics.SetOpacity(theme.ColorScheme().Scrim, 0.32))

		c.StartBlock("Overlay")
		c.Modifier(func(modifier ui.Modifier) ui.Modifier {
			return modifier.Then(opts.Modifier)
		})
		c.WithComposable(content)
		c.SetWidgetConstructor(overlayWidgetConstructor(opts))

		return c.EndBlock()
	}
}

func overlayWidgetConstructor(options OverlayOptions) layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {
			// Persistent gesture for scrim click detection
			scrimClick := node.State("scrimClick", func() any { return &gesture.Click{} }).Get().(*gesture.Click)

			// Track content dimensions for positioning scrim click areas
			var contentDims layout.Dimensions
			var contentOffset image.Point

			// Resolve ScrimColor to NRGBA
			scrimColor := graphics.ColorToNRGBA(options.ScrimColor)

			parentSize := gtx.Constraints.Max

			// Layout with Stack - center content
			dims := layout.Stack{Alignment: layout.Center}.Layout(gtx,
				// Layer 1: Scrim background (drawn first, behind content)
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					paint.Fill(gtx.Ops, scrimColor)
					return layout.Dimensions{Size: gtx.Constraints.Max}
				}),
				// Layer 2: Content (drawn on top)
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					children := node.Children()
					if len(children) > 0 {
						child := children[0].(layoutnode.NodeCoordinator)
						contentDims = child.Layout(gtx)
						// Calculate content offset (centered in parent)
						contentOffset = image.Point{
							X: (parentSize.X - contentDims.Size.X) / 2,
							Y: (parentSize.Y - contentDims.Size.Y) / 2,
						}
						return contentDims
					}
					return layout.Dimensions{}
				}),
			)

			// Handle scrim clicks ONLY for the scrim area (not content area)
			// We create click handler areas around the content using 4 rectangles
			if options.OnDismiss != nil {
				// Process any pending click events from scrim areas
				for {
					e, ok := scrimClick.Update(gtx.Source)
					if !ok {
						break
					}
					if e.Kind == gesture.KindClick {
						options.OnDismiss()
					}
				}

				// Register click handlers for the 4 scrim areas around content
				// This avoids covering the content area, so content handlers work normally

				// Top region
				if contentOffset.Y > 0 {
					topRect := clip.Rect{Max: image.Point{X: parentSize.X, Y: contentOffset.Y}}
					func() {
						defer topRect.Push(gtx.Ops).Pop()
						scrimClick.Add(gtx.Ops)
						event.Op(gtx.Ops, scrimClick)
					}()
				}

				// Bottom region
				contentBottom := contentOffset.Y + contentDims.Size.Y
				if contentBottom < parentSize.Y {
					op.Offset(image.Point{X: 0, Y: contentBottom}).Add(gtx.Ops)
					bottomMacro := op.Record(gtx.Ops)
					bottomRect := clip.Rect{Max: image.Point{X: parentSize.X, Y: parentSize.Y - contentBottom}}
					func() {
						defer bottomRect.Push(gtx.Ops).Pop()
						scrimClick.Add(gtx.Ops)
						event.Op(gtx.Ops, scrimClick)
					}()
					bottomMacro.Stop()
				}

				// Left region (only the strip between top and bottom)
				if contentOffset.X > 0 {
					leftMacro := op.Record(gtx.Ops)
					op.Offset(image.Point{X: 0, Y: contentOffset.Y}).Add(gtx.Ops)
					leftRect := clip.Rect{Max: image.Point{X: contentOffset.X, Y: contentDims.Size.Y}}
					func() {
						defer leftRect.Push(gtx.Ops).Pop()
						scrimClick.Add(gtx.Ops)
						event.Op(gtx.Ops, scrimClick)
					}()
					leftMacro.Stop()
				}

				// Right region (only the strip between top and bottom)
				contentRight := contentOffset.X + contentDims.Size.X
				if contentRight < parentSize.X {
					rightMacro := op.Record(gtx.Ops)
					op.Offset(image.Point{X: contentRight, Y: contentOffset.Y}).Add(gtx.Ops)
					rightRect := clip.Rect{Max: image.Point{X: parentSize.X - contentRight, Y: contentDims.Size.Y}}
					func() {
						defer rightRect.Push(gtx.Ops).Pop()
						scrimClick.Add(gtx.Ops)
						event.Op(gtx.Ops, scrimClick)
					}()
					rightMacro.Stop()
				}
			}

			return dims
		}
	})
}
