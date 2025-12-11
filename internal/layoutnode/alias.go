package layoutnode

import (
	node "go-compose-dev/internal/Node"
	"go-compose-dev/internal/identity"
	"go-compose-dev/internal/immap"
	"go-compose-dev/internal/modifier"
	"go-compose-dev/internal/state"

	"gioui.org/layout"
	"gioui.org/op"
)

type TreeNode = node.TreeNode
type ChainNode = node.ChainNode

type LayoutContext = layout.Context
type LayoutDimensions = layout.Dimensions
type LayoutConstraints = layout.Constraints
type GioLayoutWidget = layout.Widget

type DrawOp = op.CallOp

type NodeID = node.NodeID

type Modifier = modifier.Modifier
type ModifierElement = modifier.ModifierElement
type InspectableModifier = modifier.InspectableModifier

type Element = modifier.Element

type ElementStore = modifier.ElementStore

var EmptyElementStore = modifier.EmptyElementStore

var EmptyModifier = modifier.EmptyModifier

type Identifier = identity.Identifier
type IdentityManager = identity.IdentityManager

var GetScopedIdentityManager = identity.GetScopedIdentityManager

type SupportState = state.SupportState
type Memo = state.Memo
type PersistentState = state.PersistentState
type MutableValue = state.MutableValue

type Slots = immap.ImmutableMap[any]
