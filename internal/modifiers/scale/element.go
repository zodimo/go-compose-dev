package scale

import (
	node "github.com/zodimo/go-compose/internal/Node"
	"github.com/zodimo/go-compose/internal/modifier"
)

type ScaleElement struct {
	data ScaleData
}

func (e *ScaleElement) Create() node.Node {
	return NewScaleNode(e.data)
}

func (e *ScaleElement) Update(n node.Node) {
	no := n.(*ScaleNode)
	no.data = e.data
}

func (e *ScaleElement) Equals(other modifier.Element) bool {
	o, ok := other.(*ScaleElement)
	if !ok {
		return false
	}
	return e.data == o.data
}
