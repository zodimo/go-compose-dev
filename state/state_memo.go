package state

import (
	"github.com/zodimo/go-maybe"
)

type MemoState[T any] interface {
	Get(key string) maybe.Maybe[T]
	Set(key string, value T) MemoState[T]
}
