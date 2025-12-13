package lazy

import (
	"fmt"
	"go-compose-dev/compose"

	"gioui.org/layout"
	"gioui.org/widget"
)

type LazyListState struct {
	List widget.List
}

func NewLazyListState() *LazyListState {
	return &LazyListState{
		List: widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}
}

func RememberLazyListState(c compose.Composer) *LazyListState {
	id := c.GenerateID()
	key := fmt.Sprintf("lazyListState-%v", id)
	return c.State(key, func() any {
		return NewLazyListState()
	}).Get().(*LazyListState)
}
