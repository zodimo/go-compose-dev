package shape

import (
	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

var _ CornerSize = (*ZeroCornerSize)(nil)

type ZeroCornerSize struct {
}

func NewZeroCornerSize() *ZeroCornerSize {
	return &ZeroCornerSize{}
}

func (z *ZeroCornerSize) ToPx(shapeSize geometry.Size, density unit.Density) float32 {
	return 0
}

func (z *ZeroCornerSize) stringCornerSize() string {
	return "CornerSize{zero}"
}

func (d *ZeroCornerSize) isZero() bool {
	return true
}
