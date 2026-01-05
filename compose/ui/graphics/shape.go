package graphics

import (
	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

type Shape interface {
	/**
	 * Creates [Outline] of this shape for the given [size].
	 *
	 * @param size the size of the shape boundary.
	 * @param layoutDirection the current layout direction.
	 * @param density the current density of the screen.
	 * @return [Outline] of this shape for the given [size].
	 */
	CreateOutline(size geometry.Size, layoutDirection unit.LayoutDirection, density unit.Density) Outline
}
