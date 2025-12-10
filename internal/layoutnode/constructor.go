package layoutnode

import (
	"go-compose-dev/internal/immap"

	"github.com/zodimo/go-maybe"
)

func NewLayoutNode(id NodeID, key string, slotStore immap.ImmutableMap[Element]) LayoutNode {
	return &layoutNode{
		id:           id,
		key:          key,
		children:     []LayoutNode{},
		modifier:     EmptyModifier,
		slots:        slotStore,
		innerWidget:  IdentityGioLayoutWidget,
		layoutResult: maybe.None[LayoutResult](),
	}
}

func NewNodeCoordinator(node LayoutNode) NodeCoordinator {

	return &nodeCoordinator{
		LayoutNode:       node,
		layoutCallChain:  NewLayoutWidget(node.GetWidget()),
		pointerCallChain: NewLayoutWidget(node.GetWidget()),
		drawCallChain:    node.GetDrawWidget(),
		elementStore:     EmptyElementStore,
	}
}

var IdentityGioLayoutWidget = func(gtx LayoutContext) LayoutDimensions {
	return LayoutDimensions{
		Size: gtx.Constraints.Min,
	}
}
var IdentityLayoutWidget = NewLayoutWidget(IdentityGioLayoutWidget)

func NewLayoutWidget(innerWidget GioLayoutWidget) LayoutWidget {
	return layoutWidget{
		innerWidget: innerWidget,
	}
}

func NewDrawWidget(drawFunc DrawFunc) DrawWidget {
	return drawWidget{
		drawFunc: drawFunc,
	}
}

// func NewPointerWidget(innerWidget LayoutContextReceiver) PointerWidget {
// 	return pointerWidget{
// 		innerWidget: innerWidget,
// 	}
// }

var _ LayoutNodeWidgetConstructor = (*layoutNodeWidgetConstructor)(nil)

type layoutNodeWidgetConstructor struct {
	MakeFunc func(node LayoutNode) GioLayoutWidget
}

func NewLayoutNodeWidgetConstructor(makeFunc func(node LayoutNode) GioLayoutWidget) layoutNodeWidgetConstructor {
	return layoutNodeWidgetConstructor{
		MakeFunc: makeFunc,
	}
}

func (c layoutNodeWidgetConstructor) Make(node LayoutNode) GioLayoutWidget {
	return c.MakeFunc(node)
}

func NewLayoutContextReceiver(widget GioLayoutWidget) LayoutContextReceiver {
	return func(gtx LayoutContext) {
		widget(gtx)
	}
}
