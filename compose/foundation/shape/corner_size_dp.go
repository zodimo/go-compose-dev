package shape

import (
	"fmt"

	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

var _ CornerSize = (*DpCornerSize)(nil)

type DpCornerSize struct {
	size unit.Dp
}

func NewDpCornerSize(size unit.Dp) *DpCornerSize {
	return &DpCornerSize{size: size}
}

func (d *DpCornerSize) ToPx(shapeSize geometry.Size, density unit.Density) float32 {
	return 0
}

func (d *DpCornerSize) stringCornerSize() string {
	return fmt.Sprintf("CornerSize{size = %s}", d.size)
}
func (d *DpCornerSize) isZero() bool {
	return d.size == unit.Dp(0)
}
