package geometry

import "github.com/zodimo/go-compose/pkg/floatutils"

type Offset struct {
	X float32
	Y float32
}

func (o Offset) Equal(other Offset) bool {

	return floatutils.Float32Equals(o.X, other.X, floatutils.Float32EqualityThreshold) && floatutils.Float32Equals(o.Y, other.Y, floatutils.Float32EqualityThreshold)
}

func OffsetZero() Offset {
	return Offset{
		X: 0,
		Y: 0,
	}
}
