package lazy

import (
	"go-compose-dev/compose"
)

type LazyListScope interface {
	Item(key any, content compose.Composable)
	Items(count int, key func(index int) any, itemContent func(index int) compose.Composable)
}

type lazyListScopeImpl struct {
	items []lazyItem
}

type lazyItem struct {
	Key     any
	Content compose.Composable
}

func (s *lazyListScopeImpl) Item(key any, content compose.Composable) {
	s.items = append(s.items, lazyItem{Key: key, Content: content})
}

func (s *lazyListScopeImpl) Items(count int, key func(index int) any, itemContent func(index int) compose.Composable) {
	for i := 0; i < count; i++ {
		var k any
		if key != nil {
			k = key(i)
		}
		s.items = append(s.items, lazyItem{Key: k, Content: itemContent(i)})
	}
}
