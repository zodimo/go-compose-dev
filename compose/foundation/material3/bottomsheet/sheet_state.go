package bottomsheet

import (
	"time"

	"github.com/zodimo/go-compose/internal/animation"
)

type SheetState struct {
	visibleAnim *animation.VisibilityAnimation
}

func NewSheetState() *SheetState {
	return &SheetState{
		visibleAnim: &animation.VisibilityAnimation{
			Duration: time.Millisecond * 300,
			State:    animation.Invisible,
		},
	}
}

func (s *SheetState) Show() {
	s.visibleAnim.Appear(time.Now())
}

func (s *SheetState) Hide() {
	s.visibleAnim.Disappear(time.Now())
}

func (s *SheetState) IsVisible() bool {
	return s.visibleAnim.Visible()
}
