package runtime

var _ Runtime = (*runtime)(nil)

type runtime struct {
}

func (r *runtime) Run(ctx LayoutContext, node LayoutNode) {
}
