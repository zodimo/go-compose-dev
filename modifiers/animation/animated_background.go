package animation

import (
	"image/color"

	"github.com/zodimo/go-compose/internal/animation"
	"github.com/zodimo/go-compose/internal/modifier"
	"github.com/zodimo/go-compose/modifiers/background"
)

type AnimatedBackgroundElement struct {
	Anim  *animation.VisibilityAnimation
	Color color.Color
	Shape background.Shape
}

func (e AnimatedBackgroundElement) Create() Node {
	return NewAnimatedBackgroundNode(e)
}

func (e AnimatedBackgroundElement) Update(node Node) {
	n := node.(*AnimatedBackgroundNode)
	n.element = e
}

func (e AnimatedBackgroundElement) Equals(other Element) bool {
	o, ok := other.(AnimatedBackgroundElement)
	return ok && o.Anim == e.Anim && o.Color == e.Color
}

// AnimatedBackground creates a modifier that paints a background with alpha scaled by animation.
func AnimatedBackground(anim *animation.VisibilityAnimation, color color.Color, shape background.Shape) Modifier {
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			AnimatedBackgroundElement{
				Anim:  anim,
				Color: color,
				Shape: shape,
			},
		),
		modifier.NewInspectorInfo(
			"AnimatedBackground",
			map[string]any{
				"Anim":  anim,
				"Color": color,
			},
		),
	)
}
