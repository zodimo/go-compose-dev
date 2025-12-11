package shadow

import (
	"go-compose-dev/compose/ui/graphics/shape"
	node "go-compose-dev/internal/Node"
	"go-compose-dev/internal/layoutnode"
	"go-compose-dev/internal/modifier"
	"go-compose-dev/internal/modifiers/helpers"
	"image/color"

	"gioui.org/unit"
)

type Modifier = modifier.Modifier
type Element = modifier.Element
type InspectableModifier = modifier.InspectableModifier
type ModifierInspectorInfo = modifier.InspectorInfo

type Node = node.Node
type TreeNode = node.TreeNode
type ChainNode = node.ChainNode

type DrawModifierNode = layoutnode.DrawModifierNode
type LayoutWidget = layoutnode.LayoutWidget
type LayoutContext = layoutnode.LayoutContext
type LayoutDimensions = layoutnode.LayoutDimensions

var ToNRGBA = helpers.ToNRGBA

type Shape = shape.Shape
type Dp = unit.Dp
type Color = color.Color
