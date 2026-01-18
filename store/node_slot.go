package store

import (
	"github.com/zodimo/go-compose/internal/immap"
	"github.com/zodimo/go-compose/state"
	"github.com/zodimo/go-maybe"
)

var _ state.NodeSlotsState[any] = (*NodeSlots[any])(nil)

type NodeSlots[T any] struct {
	innerMap immap.ImmutableMap[T]
}

func EmptyNodeSlots[T any]() state.NodeSlotsState[T] {
	return &NodeSlots[T]{
		innerMap: immap.EmptyImmutableMap[T](),
	}
}

func (ns *NodeSlots[T]) Get(key string) maybe.Maybe[T] {
	value, ok := ns.innerMap.Find(key)
	if !ok {
		return maybe.None[T]()
	}
	return maybe.Some(value)
}

func (ns *NodeSlots[T]) Set(key string, value T) state.NodeSlotsState[T] {
	return &NodeSlots[T]{
		innerMap: ns.innerMap.Assoc(key, value),
	}
}
func (ns *NodeSlots[T]) AsMap() map[string]T {
	return ns.innerMap
}
