package background

import (
	node "go-compose-dev/internal/Node"
	"go-compose-dev/internal/layoutnode"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

var _ ChainNode = (*BackgroundNode)(nil)

// var _ DrawModifierNode = (*BackgroundNode)(nil)

// NodeKind should also implement the interface of the LayoutNode for that phase

func NewBackGroundNode(background BackgroundData) ChainNode {
	return BackgroundNode{
		ChainNode: node.NewChainNode(
			node.NewNodeID(),
			node.NodeKindDraw,
			node.DrawPhase,
			//OnAttach
			func(n TreeNode) {
				// how should the tree now be updated when attached
				// tree nde is the layout tree

				no := n.(layoutnode.DrawModifierNode)
				// we can now work with the layoutNode
				no.AttachDrawModifier(func() layoutnode.DrawWidget {

					return layoutnode.NewDrawWidget(func(gtx layoutnode.LayoutContext, node layoutnode.LayoutNode) layoutnode.DrawOp {
						macro := op.Record(gtx.Ops)
						layoutResult := node.GetLayoutResult().UnwrapUnsafe()

						layout.Background{}.Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								// shape
								// color
								defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()

								paint.Fill(gtx.Ops, ToNRGBA(background.Color))

								return layout.Dimensions{Size: gtx.Constraints.Min}

							},
							func(gtx layout.Context) layout.Dimensions {
								layoutResult.DrawOp.Add(gtx.Ops)
								return layoutResult.Dimensions
							},
						)

						return macro.Stop()
					})
				})

			},
		),
		background: background,
	}
}

type BackgroundNode struct {
	ChainNode
	background BackgroundData
}
