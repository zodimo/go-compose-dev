package size

var _ Element = (*SizeElement)(nil)

type SizeElement struct {
	size SizeData
}

func CompareSize(a, b SizeData) bool {
	return a.Width == b.Width && a.Height == b.Height && a.Required == b.Required
}

// Create creates a new Chain Node instance
func (be SizeElement) Create() Node {
	//chainNode
	return NewSizeNode(be.size)

}

// Update updates an existing Chain node for efficiency
func (be SizeElement) Update(node Node) {
	if node == nil {
		panic("node cannot be nil")
	}

	bn := node.(SizeNode)
	bn.size = be.size

}

// Equals checks if this element is equivalent to another
// used during filter operations like Modifier.Any
func (se SizeElement) Equals(other Element) bool {
	if otherElement, ok := other.(SizeElement); ok {
		return CompareSize(se.size, otherElement.size)
	}
	return false
}

func (se SizeElement) Size() SizeData {
	return se.size
}
