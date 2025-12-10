package layoutnode

import (
	"fmt"
	"go-compose-dev/internal/modifier"

	"gioui.org/op"
)

var _ NodeCoordinator = (*nodeCoordinator)(nil)

type nodeCoordinator struct {
	LayoutNode
	layoutCallChain  LayoutWidget
	drawCallChain    DrawWidget
	pointerCallChain LayoutWidget
	elementStore     ElementStore
}

func (nc *nodeCoordinator) Expand() {
	modifierChain := nc.LayoutNode.UnwrapModifier().AsChain()
	*nc = *modifier.FoldOut(modifierChain, nc, func(nc *nodeCoordinator, mod Modifier) *nodeCoordinator {

		fmt.Println("==========DEBUG ===========BELOW =")
		fmt.Println(DebugLayoutNode(nc))

		if inspectable, ok := mod.(InspectableModifier); ok {
			mod = inspectable.Unwrap()
		}
		modifierElement := mod.(ModifierElement)

		modifierNode := modifierElement.Create()
		modifierChainNode := modifierNode.(ChainNode)
		fmt.Printf("%s\n", modifierChainNode.Kind())

		modifierChainNode.Attach(nc)

		return nc
	})

}

func (nc *nodeCoordinator) AttachLayoutModifier(attach func(gtx LayoutContext, widget LayoutWidget) LayoutWidget) {

	nc.layoutCallChain = nc.layoutCallChain.Map(func(in LayoutWidget) LayoutWidget {
		return NewLayoutWidget(func(gtx LayoutContext) LayoutDimensions {
			return attach(gtx, in).Layout(gtx)
		})
	})
}
func (nc *nodeCoordinator) AttachDrawModifier(attach func() DrawWidget) {
	nc.drawCallChain = nc.drawCallChain.Map(func(in DrawWidget) DrawWidget {
		return NewDrawWidget(func(gtx LayoutContext, node LayoutNode) op.CallOp {
			return attach().Draw(gtx, node)
		})
	})
}
func (nc *nodeCoordinator) AttachPointerModifier(attach func(gtx LayoutContext, widget LayoutWidget) LayoutWidget) {
	nc.pointerCallChain = nc.pointerCallChain.Map(func(in LayoutWidget) LayoutWidget {
		return NewLayoutWidget(func(gtx LayoutContext) LayoutDimensions {
			return attach(gtx, in).Layout(gtx)
		})
	})
}

func (nc *nodeCoordinator) LayoutPhase(gtx LayoutContext) {

	maybeLayoutResult := nc.GetLayoutResult()
	if maybeLayoutResult.IsSome() {
		return
	}

	macro := op.Record(gtx.Ops)
	dimensions := nc.LayoutSelf(gtx)
	drawOp := macro.Stop()
	result := LayoutResult{
		Dimensions: dimensions,
		DrawOp:     drawOp,
	}
	nc.SetLayoutResult(result)
}

func (nc *nodeCoordinator) PointerPhase(gtx LayoutContext) {
	defer op.Record(gtx.Ops).Stop()
	nc.pointerCallChain.Layout(gtx)
}

func (nc *nodeCoordinator) DrawPhase(gtx LayoutContext) DrawOp {
	return nc.drawCallChain.Draw(gtx, nc)
}

func (nc *nodeCoordinator) Elements() ElementStore {
	return nc.elementStore
}

func (nc *nodeCoordinator) LayoutSelf(gtx LayoutContext) LayoutDimensions {
	maybeLayoutResult := nc.GetLayoutResult()
	if maybeLayoutResult.IsSome() {
		return maybeLayoutResult.UnwrapUnsafe().Dimensions
	}
	return nc.layoutCallChain.Layout(gtx)
}

type LayoutContextReceiver = func(gtx LayoutContext)

type LayoutResult struct {
	Dimensions LayoutDimensions
	DrawOp     op.CallOp
}
