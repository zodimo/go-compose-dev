package modifiers

import (
	node "go-compose-dev/internal/Node"
	"go-compose-dev/internal/modifier"
)

type Modifier = modifier.Modifier
type Element = modifier.Element
type InspectableModifier = modifier.InspectableModifier
type ModifierInspectorInfo = modifier.InspectorInfo

type Node = node.Node
type TreeNode = node.TreeNode
type ChainNode = node.ChainNode

var EmptyModifier = modifier.EmptyModifier
