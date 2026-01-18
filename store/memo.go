package store

import (
	"github.com/zodimo/go-compose/internal/immap"
	"github.com/zodimo/go-compose/state"
	"github.com/zodimo/go-maybe"
)

var _ state.MemoState[any] = (*Memo)(nil)
var _ state.MemoState[any] = (*MemoTyped[any])(nil)

type Memo struct {
	innerMap immap.ImmutableMap[any]
}

func (m *Memo) Get(key string) maybe.Maybe[any] {
	value, ok := m.innerMap.Find(key)
	if !ok {
		return maybe.None[any]()
	}
	return maybe.Some(value)
}

func (m *Memo) Set(key string, value any) state.MemoState[any] {
	return &Memo{
		innerMap: m.innerMap.Assoc(key, value),
	}
}

type MemoTyped[T any] struct {
	innerMap immap.ImmutableMap[T]
}

func (m *MemoTyped[T]) Get(key string) maybe.Maybe[T] {
	value, ok := m.innerMap.Find(key)
	if !ok {
		return maybe.None[T]()
	}
	return maybe.Some(value)
}

func (m *MemoTyped[T]) Set(key string, value T) state.MemoState[T] {
	return &MemoTyped[T]{
		innerMap: m.innerMap.Assoc(key, value),
	}
}

func newMemoTyped[T any]() state.MemoState[T] {
	return &MemoTyped[T]{
		innerMap: immap.EmptyImmutableMap[T](),
	}
}

func EmptyMemo[T any]() state.MemoState[T] {
	return newMemoTyped[T]()
}
