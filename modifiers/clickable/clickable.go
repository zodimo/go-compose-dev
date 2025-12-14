package clickable

type ClickableData struct {
	OnClick   func()
	Clickable *GioClickable
}

type ClickableElement struct {
	clickableData ClickableData
}

// Create creates a new Chain Node instance
func (e ClickableElement) Create() Node {
	//chainNode
	return NewClickableNode(e)

}

// Update updates an existing Chain node for efficiency
func (e ClickableElement) Update(node Node) {
	if node == nil {
		panic("node cannot be nil")
	}

	n := node.(ClickableNode)
	n.clickableData = e.clickableData

}

// Equals checks if this element is equivalent to another
// used during filter operations like Modifier.Any
func (e ClickableElement) Equals(other Element) bool {
	return false
}

func (e ClickableElement) ClickableData() ClickableData {
	return e.clickableData
}
