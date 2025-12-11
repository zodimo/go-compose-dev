package shadow

import (
	"go-compose-dev/internal/modifier"
	"image/color"
)

type ShadowData struct {
	Elevation    Dp
	Shape        Shape
	AmbientColor Color
	SpotColor    Color
}

type ShadowElement struct {
	shadowData ShadowData
}

func (e *ShadowElement) Create() Node {
	return NewShadowNode(*e)
}

func (e *ShadowElement) Update(node Node) {
	n := node.(*ShadowNode)
	n.shadowData = e.shadowData
}

func (e *ShadowElement) Equals(other Element) bool {
	if otherEle, ok := other.(*ShadowElement); ok {
		return e.shadowData.Elevation == otherEle.shadowData.Elevation &&
			e.shadowData.Shape == otherEle.shadowData.Shape // Pointer comparison for shape might be enough if not changed
	}
	return false
}

func Shadow(elevation Dp, shape Shape, ambientColor, spotColor Color) Modifier {
	return modifier.NewInspectableModifier(
		modifier.NewModifier(
			&ShadowElement{
				shadowData: ShadowData{
					Elevation:    elevation,
					Shape:        shape,
					AmbientColor: ambientColor,
					SpotColor:    spotColor,
				},
			},
		),
		modifier.NewInspectorInfo(
			"shadow",
			map[string]any{
				"elevation":    elevation,
				"shape":        shape,
				"ambientColor": ambientColor,
				"spotColor":    spotColor,
			},
		),
	)
}

// Simple Shadow with defaults
func Simple(elevation Dp, shape Shape) Modifier {
	// Default shadow colors usually black or standard shadow color?
	// gio-mw uses a shadow color from theme.
	// For generic modifier, we might default to black with low opacity if not specified,
	// but here we let caller specify or use these defaults.
	black := color.NRGBA{A: 255}
	return Shadow(elevation, shape, black, black)
}
