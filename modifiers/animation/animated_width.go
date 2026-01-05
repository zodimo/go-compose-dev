package animation

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/internal/animation"
	"github.com/zodimo/go-compose/internal/modifier"
)

type AnimatedWidthElement struct {
	Anim     *animation.VisibilityAnimation
	MaxWidth int
}

func (e AnimatedWidthElement) Create() Node {
	return NewAnimatedWidthNode(e)
}

func (e AnimatedWidthElement) Update(node Node) {
	n := node.(*AnimatedWidthNode)
	n.element = e
}

func (e AnimatedWidthElement) Equals(other Element) bool {
	o, ok := other.(AnimatedWidthElement)
	return ok && o.Anim == e.Anim && o.MaxWidth == e.MaxWidth
}

func AnimatedWidth(anim *animation.VisibilityAnimation, maxWidth int) ui.Modifier {
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			AnimatedWidthElement{
				Anim:     anim,
				MaxWidth: maxWidth,
			},
		),
		modifier.NewInspectorInfo(
			"AnimatedWidth",
			map[string]any{
				"Anim":     anim,
				"MaxWidth": maxWidth,
			},
		),
	)
}
