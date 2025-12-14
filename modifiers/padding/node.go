package padding

var _ ChainNode = (*PaddingNode)(nil)

// NodeKind should also implement the interface of the LayoutNode for that phase

type PaddingNode struct {
	ChainNode
	padding PaddingData
}
