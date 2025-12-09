package modifier

var _ Element = (*element)(nil)

type element struct {
	createFunc func() Node
	updateFunc func(Node)
	equalsFunc func(Element) bool
}

// Create creates a new Chain Node instance
func (e element) Create() Node {
	return e.createFunc()
}

// Update updates an existing Chain node for efficiency
func (e element) Update(node Node) {
	if node == nil {
		panic("element update: node cannot be nil")
	}
	e.updateFunc(node)
}

// Equals checks if this element is equivalent to another
// used during filter operations like Modifier.Any
func (e element) Equals(other Element) bool {
	if other == nil {
		return false
	}
	return e.equalsFunc(other)
}
