package shape

import (
	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

type CornerSize interface {
	ToPx(shapeSize geometry.Size, density unit.Density) float32
	stringCornerSize() string
	isZero() bool
}

func StringCornerSize(size CornerSize) string {
	size = CoalesceCornerSize(size, CornerSizeUnspecified)
	return size.stringCornerSize()
}

func IsSpecifiedCornerSize(size CornerSize) bool {
	return size != nil && size != CornerSizeUnspecified
}

func TakeOrElseCornerSize(s, defaultStyle CornerSize) CornerSize {
	if s == nil || s == CornerSizeUnspecified {
		return defaultStyle
	}
	return s
}

func CoalesceCornerSize(ptr, def CornerSize) CornerSize {
	if ptr == nil {
		return def
	}
	return ptr
}
