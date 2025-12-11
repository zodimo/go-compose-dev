package clickable

import (
	node "go-compose-dev/internal/Node"
	"go-compose-dev/internal/layoutnode"

	"gioui.org/widget"
	"gioui.org/widget/material"
)

var _ ChainNode = (*ClickableNode)(nil)

// NodeKind should also implement the interface of the LayoutNode for that phase

func NewClickableNode(element ClickableElement) ChainNode {
	return ClickableNode{
		ChainNode: node.NewChainNode(
			node.NewNodeID(),
			node.NodeKindPointerInput,
			node.PointerInputPhase, // bit mask and && node.DrawPhase,
			//OnAttach
			func(n TreeNode) {
				// how should the tree now be updated when attached
				// tree nde is the layout tree

				// we need persistent storage
				// lno := n.(layoutnode.LayoutNode)

				widgetClickable := widget.Clickable{}

				no := n.(layoutnode.PointerInputModifierNode)
				// we can now work with the layoutNode
				no.AttachPointerInputModifier(func(widget LayoutWidget) layoutnode.LayoutWidget {

					return layoutnode.NewLayoutWidget(func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {
						if widgetClickable.Clicked(gtx) {
							element.clickable.OnClick()
						}
						return layoutnode.LayoutDimensions{}
					})
				})

				dno := n.(layoutnode.DrawModifierNode)

				dno.AttachDrawModifier(func(widget LayoutWidget) layoutnode.LayoutWidget {
					return layoutnode.NewLayoutWidget(func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {
						return material.Clickable(gtx, &widgetClickable, widget.Layout)
					})
				})

			},
		),
		clickable: element.clickable,
	}
}

type ClickableNode struct {
	ChainNode
	clickable ClickableData
}
