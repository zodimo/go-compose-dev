package animation

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/internal/animation"
	"github.com/zodimo/go-compose/internal/modifier"
)

type AnimatedHeightElement struct {
	Anim      *animation.VisibilityAnimation
	MaxHeight int // Optional max height constraint, 0 means unconstrained (or fill max)
	// Actually, for bottom sheet, we just want to animate the height from 0 to content height.
}

func (e AnimatedHeightElement) Create() Node {
	return NewAnimatedHeightNode(e)
}

func (e AnimatedHeightElement) Update(node Node) {
	n := node.(*AnimatedHeightNode)
	n.element = e
}

func (e AnimatedHeightElement) Equals(other Element) bool {
	o, ok := other.(AnimatedHeightElement)
	return ok && o.Anim == e.Anim && o.MaxHeight == e.MaxHeight
}

// AnimatedHeight animates the height of the content based on the VisibilityAnimation.
// It clips the content as it grows/shrinks.
func AnimatedHeight(anim *animation.VisibilityAnimation, maxHeight int) ui.Modifier {
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			AnimatedHeightElement{
				Anim:      anim,
				MaxHeight: maxHeight,
			},
		),
		modifier.NewInspectorInfo(
			"AnimatedHeight",
			map[string]any{
				"Anim":      anim,
				"MaxHeight": maxHeight,
			},
		),
	)
}
