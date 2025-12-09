package runtime

type Runtime interface {
	Run(LayoutContext, LayoutNode)
}
