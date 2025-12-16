package animation

import (
	"github.com/zodimo/go-compose/internal/animation"
	"github.com/zodimo/go-compose/internal/modifier"
	"github.com/zodimo/go-compose/modifiers/background"
	"github.com/zodimo/go-compose/theme"
)

type AnimatedBackgroundElement struct {
	Anim  *animation.VisibilityAnimation
	Color theme.ColorDescriptor
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
	if !ok {
		return false
	}
	// Basic implementation - likely need proper comparison for ColorDescriptor if it's complex
	// For now assume pointer or simple struct comparison works or use Compare method if consistent.
	// ColorDescriptor has Compare method.
	return o.Anim == e.Anim && e.Color.Compare(o.Color)
}

// AnimatedBackground creates a modifier that paints a background with alpha scaled by animation.
func AnimatedBackground(anim *animation.VisibilityAnimation, color theme.ColorDescriptor, shape background.Shape) Modifier {
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
