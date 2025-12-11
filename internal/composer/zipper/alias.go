package zipper

import (
	node "go-compose-dev/internal/Node"
	"go-compose-dev/internal/identity"
	"go-compose-dev/internal/layoutnode"
	"go-compose-dev/internal/modifier"
	"go-compose-dev/internal/state"
	"go-compose-dev/pkg/api"
	idApi "go-compose-dev/pkg/compose-identifier/api"
)

type LayoutNode = layoutnode.LayoutNode
type Memo = state.MemoTyped[any]
type ElementMemo = state.MemoTyped[Element]

type PersistentState = state.PersistentState

var EmptyMemo = state.EmptyMemo[any]()
var EmptyElementMemo = state.EmptyMemo[Element]()

type IdentityManager = identity.IdentityManager

var GetScopedIdentityManager = identity.GetScopedIdentityManager

// compose-identifier.api.Identifier
type Identifier = idApi.Identifier // Public API of the composer
type NodePath = node.NodePath
type Modifier = modifier.Modifier
type Element = modifier.Element
type MutableValue = state.MutableValue
type RootNode = node.TreeNode

type Composable = api.Composable
type Composer = api.Composer
