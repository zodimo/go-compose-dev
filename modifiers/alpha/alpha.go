package alpha

type AlphaElement struct {
	Alpha float32
}

var _ Element = (*AlphaElement)(nil)

// Create creates a new Chain Node instance
func (e AlphaElement) Create() Node {
	return NewAlphaNode(e)
}

// Update updates an existing Chain node for efficiency
func (e AlphaElement) Update(node Node) {
	if node == nil {
		panic("node cannot be nil")
	}

	n := node.(*AlphaNode)
	n.state.Alpha = e.Alpha

}

// Equals checks if this element is equivalent to another
func (e AlphaElement) Equals(other Element) bool {
	if otherElement, ok := other.(AlphaElement); ok {
		return e.Alpha == otherElement.Alpha
	}
	return false
}

func (e AlphaElement) GetAlpha() float32 {
	return e.Alpha
}
