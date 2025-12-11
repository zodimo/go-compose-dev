package border

import (
	node "go-compose-dev/internal/Node"
	"go-compose-dev/internal/layoutnode"

	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type BorderNode struct {
	ChainNode
	borderData BorderData
}

func NewBorderNode(element BorderElement) *BorderNode {
	n := &BorderNode{
		borderData: element.borderData,
	}
	n.ChainNode = node.NewChainNode(
		node.NewNodeID(),
		node.NodeKindDraw,
		node.DrawPhase,
		func(t TreeNode) {
			no := t.(DrawModifierNode)
			no.AttachDrawModifier(func(widget LayoutWidget) LayoutWidget {
				return layoutnode.NewLayoutWidget(func(gtx LayoutContext) LayoutDimensions {
					// Layout content first
					dims := widget.Layout(gtx)

					width := n.borderData.Width
					if width <= 0 {
						return dims
					}

					// Draw border on top? Or behind?
					// Usually border is on top of content if it's "inside", or it expands size if "outside".
					// In Compose, Border modifier draws ON TOP of the content, inside the bounds.
					// So acts like an overlay.

					// We need the outline path.
					// Shape logic uses CreateOutline(size, metric)
					// The generic generic Outline interface now supports Path(ops *op.Ops) clip.PathSpec.

					outline := n.borderData.Shape.CreateOutline(dims.Size, gtx.Metric)

					macro := op.Record(gtx.Ops)

					strokeWidth := float32(gtx.Metric.Dp(width))

					pathSpec := outline.Path(gtx.Ops)

					// Create stroke op
					strokeOp := clip.Stroke{
						Path:  pathSpec,
						Width: strokeWidth,
					}.Op()

					// Paint the stroke
					paint.FillShape(gtx.Ops, ToNRGBA(n.borderData.Color), strokeOp)

					call := macro.Stop()
					call.Add(gtx.Ops)

					return dims
				})
			})
		},
	)
	return n
}
