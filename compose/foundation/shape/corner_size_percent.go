package shape

import (
	"fmt"

	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

var _ CornerSize = (*PercentCornerSize)(nil)

type PercentCornerSize struct {
	percent float32
}

func NewPercentCornerSize(percent float32) *PercentCornerSize {
	return &PercentCornerSize{percent: percent}
}

func (p *PercentCornerSize) ToPx(shapeSize geometry.Size, density unit.Density) float32 {
	return (p.percent / 100) * shapeSize.MinDimension()
}

func (p *PercentCornerSize) stringCornerSize() string {
	return fmt.Sprintf("CornerSize{percent = %.1f%%}", p.percent)
}

func (d *PercentCornerSize) isZero() bool {
	return d.percent == 0
}
