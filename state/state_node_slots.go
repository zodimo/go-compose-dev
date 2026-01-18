package state

import "github.com/zodimo/go-maybe"

type NodeSlotsState[T any] interface {
	Get(key string) maybe.Maybe[T]
	Set(key string, value T) NodeSlotsState[T]
	// should be part of debug interface
	AsMap() map[string]T
}
