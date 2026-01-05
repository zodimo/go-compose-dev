package shape

import (
	"fmt"

	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

var _ CornerSize = (*PxCornerSize)(nil)

type PxCornerSize struct {
	size float32
}

func NewPxCornerSize(size float32) *PxCornerSize {
	return &PxCornerSize{size: size}
}

func (p *PxCornerSize) ToPx(shapeSize geometry.Size, density unit.Density) float32 {
	return p.size
}

func (p *PxCornerSize) stringCornerSize() string {
	return fmt.Sprintf("CornerSize{size = %.1fpx}", p.size)
}

func (d *PxCornerSize) isZero() bool {
	return d.size == 0
}
