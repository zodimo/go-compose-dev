package geometry

import (
	"fmt"

	"github.com/zodimo/go-compose/pkg/floatutils"
	"github.com/zodimo/go-compose/pkg/floatutils/lerp"
)

// CornerRadius represents the radii for corners of a rounded rectangle.
// It is a packed value where the x radius is in the high 32 bits and the y radius is in the low 32 bits.
type CornerRadius int64

// CornerRadiusZero is a CornerRadius with both x and y radii equal to zero (sharp corners).
var CornerRadiusZero = NewCornerRadius(0, 0)

// CornerRadiusUnspecified represents an unspecified corner radius.
var CornerRadiusUnspecified = NewCornerRadius(floatutils.Float32Unspecified, floatutils.Float32Unspecified)

// NewCornerRadius creates a CornerRadius from the given x and y radii.
// x is the radius of the corners along the x-axis.
// y is the radius of the corners along the y-axis (defaults to x if you want circular corners).
func NewCornerRadius(x, y float32) CornerRadius {
	return CornerRadius(floatutils.PackFloats(x, y))
}

func NewCornerRadiusUniform(radius float32) CornerRadius {
	return NewCornerRadius(radius, radius)
}

// NewCircularCornerRadius creates a CornerRadius with equal x and y radii (circular corners).
func NewCircularCornerRadius(radius float32) CornerRadius {
	return NewCornerRadius(radius, radius)
}

// X returns the x component of the corner radius.
func (c CornerRadius) X() float32 {
	return floatutils.UnpackFloat1(int64(c))
}

// Y returns the y component of the corner radius.
func (c CornerRadius) Y() float32 {
	return floatutils.UnpackFloat2(int64(c))
}

// String returns a string representation of the CornerRadius.
func (c CornerRadius) String() string {
	if c == CornerRadiusZero {
		return "CornerRadius.Zero"
	}
	if c == CornerRadiusUnspecified {
		return "CornerRadius.Unspecified"
	}
	return fmt.Sprintf("CornerRadius(%.1f, %.1f)", c.X(), c.Y())
}

// Equal checks equality with another CornerRadius.
func (c CornerRadius) Equal(other CornerRadius) bool {
	if c == other {
		return true
	}
	return floatutils.Float32Equals(c.X(), other.X(), floatutils.Float32EqualityThreshold) &&
		floatutils.Float32Equals(c.Y(), other.Y(), floatutils.Float32EqualityThreshold)
}

// IsSpecified returns true if this is not CornerRadius.Unspecified.
func (c CornerRadius) IsSpecified() bool {
	return c != CornerRadiusUnspecified
}

// IsUnspecified returns true if this is CornerRadius.Unspecified.
func (c CornerRadius) IsUnspecified() bool {
	return c == CornerRadiusUnspecified
}

// TakeOrElse returns this corner radius if Specified, otherwise returns the fallback.
func (c CornerRadius) TakeOrElse(fallback CornerRadius) CornerRadius {
	if c == CornerRadiusUnspecified {
		return fallback
	}
	return c
}

// Plus returns a CornerRadius with radii that are the sum of both corner radii.
func (c CornerRadius) Plus(other CornerRadius) CornerRadius {
	return NewCornerRadius(c.X()+other.X(), c.Y()+other.Y())
}

// Minus returns a CornerRadius with radii that are the difference of both corner radii.
func (c CornerRadius) Minus(other CornerRadius) CornerRadius {
	return NewCornerRadius(c.X()-other.X(), c.Y()-other.Y())
}

// Times returns a CornerRadius with radii scaled by the given factor.
func (c CornerRadius) Times(operand float32) CornerRadius {
	return NewCornerRadius(c.X()*operand, c.Y()*operand)
}

// Div returns a CornerRadius with radii divided by the given factor.
func (c CornerRadius) Div(operand float32) CornerRadius {
	return NewCornerRadius(c.X()/operand, c.Y()/operand)
}

// LerpCornerRadius linearly interpolates between two corner radii.
func LerpCornerRadius(start, stop CornerRadius, fraction float32) CornerRadius {
	return NewCornerRadius(
		lerp.Between32(start.X(), stop.X(), fraction),
		lerp.Between32(start.Y(), stop.Y(), fraction),
	)
}
