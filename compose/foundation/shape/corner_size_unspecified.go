package shape

import (
	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

var CornerSizeUnspecified CornerSize = &CornerSizeUnspecifiedImpl{}

var _ CornerSize = (*CornerSizeUnspecifiedImpl)(nil)

type CornerSizeUnspecifiedImpl struct{}

func (z *CornerSizeUnspecifiedImpl) ToPx(shapeSize geometry.Size, density unit.Density) float32 {
	panic("unspecified corner size")
}

func (z *CornerSizeUnspecifiedImpl) stringCornerSize() string {
	return "CornerSize{unspecified}"
}

func (d *CornerSizeUnspecifiedImpl) isZero() bool {
	panic("unspecified corner size")
}
