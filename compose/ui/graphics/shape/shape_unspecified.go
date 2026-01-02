package shape

import (
	"image"

	gioUnit "gioui.org/unit"
)

var _ Shape = (*ShapeUnspecifiedImpl)(nil)
var ShapeUnspecified Shape = &ShapeUnspecifiedImpl{}

type ShapeUnspecifiedImpl struct {
}

func (s ShapeUnspecifiedImpl) CreateOutline(size image.Point, metric gioUnit.Metric) Outline {
	return nil
}
