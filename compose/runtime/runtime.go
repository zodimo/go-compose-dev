package runtime

import "gioui.org/op"

type Runtime interface {
	Run(LayoutContext, LayoutNode) op.CallOp
}
