package layoutnode

import "gioui.org/op"

type LayoutWidget interface {
	Map(mapFun func(LayoutWidget) LayoutWidget) LayoutWidget
	Layout(gtx LayoutContext) LayoutDimensions
}

type layoutWidget struct {
	innerWidget GioLayoutWidget
}

func (lw layoutWidget) IsEmpty() bool {
	return lw.innerWidget == nil
}
func (lw layoutWidget) Layout(gtx LayoutContext) LayoutDimensions {
	if lw.IsEmpty() {
		// log that we using the identity widget
		return IdentityLayoutWidget.Layout(gtx)
	}
	return lw.innerWidget(gtx)
}

// This is how the Modifier chain is applied
func (lw layoutWidget) Map(mapFun func(LayoutWidget) LayoutWidget) LayoutWidget {
	return mapFun(lw)
}

type DrawFunc = func(gtx LayoutContext, node LayoutNode) DrawOp
type DrawWidget interface {
	Map(mapFun func(DrawWidget) DrawWidget) DrawWidget
	Draw(gtx LayoutContext, node LayoutNode) DrawOp
}

var _ DrawWidget = (*drawWidget)(nil)

type drawWidget struct {
	drawFunc DrawFunc
}

func (dw drawWidget) Draw(gtx LayoutContext, node LayoutNode) DrawOp {
	return dw.drawFunc(gtx, node)
}

// This is how the Modifier chain is applied
func (lw drawWidget) Map(mapFun func(DrawWidget) DrawWidget) DrawWidget {
	return mapFun(lw)
}

type PointerWidget interface {
	Map(mapFun func(PointerWidget) PointerWidget) PointerWidget
	Update(gtx LayoutContext)
}

var _ PointerWidget = (*pointerWidget)(nil)

type pointerWidget struct {
	innerWidget GioLayoutWidget
}

func (dw pointerWidget) Update(gtx LayoutContext) {
	//capture and discard if draw operations where to happen here
	defer op.Record(gtx.Ops).Stop()
	dw.innerWidget(gtx)
}

// This is how the Modifier chain is applied
func (lw pointerWidget) Map(mapFun func(PointerWidget) PointerWidget) PointerWidget {
	return mapFun(lw)
}
