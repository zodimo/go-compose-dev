package layoutnode

// The base Node for the Tree
type LayoutNode interface {
	TreeNode

	LayoutNodeChildren() []LayoutNode

	WithChildren(children []LayoutNode) LayoutNode

	Modifier(func(modifier Modifier) Modifier)

	WithSlotsAssoc(k string, v any) LayoutNode

	IsEmpty() bool
}

type NodeCoordinator interface {
	LayoutNode
	LayoutModifier(attach func(gtx LayoutContext, widget LayoutWidget) LayoutDimensions)
	DrawModifier(attach func(gtx LayoutContext) DrawOp)
}
