package zipper

import (
	node "github.com/zodimo/go-compose/internal/Node"
	"github.com/zodimo/go-compose/internal/identity"
	"github.com/zodimo/go-compose/internal/layoutnode"
	"github.com/zodimo/go-compose/internal/modifier"
	"github.com/zodimo/go-compose/internal/sequence"
	"github.com/zodimo/go-compose/pkg/api"
	idApi "github.com/zodimo/go-compose/pkg/compose-identifier/api"
	"github.com/zodimo/go-compose/state"
	"github.com/zodimo/go-compose/store"
)

type LayoutNode = layoutnode.LayoutNode
type Memo = state.MemoState[any]
type ElementMemo = state.MemoState[Element]

type PersistentState = state.PersistentState

var EmptyMemo = store.EmptyMemo[any]()
var EmptySlots = store.EmptyNodeSlots[any]()

type IdentityManager = identity.IdentityManager

var GetScopedIdentityManager = identity.GetScopedIdentityManager

// compose-identifier.api.Identifier
type Identifier = idApi.Identifier // Public API of the composer
type NodePath = node.NodePath

type Element = modifier.Element
type MutableValue = state.MutableValue
type StateOption = state.StateOption
type StateOptions = state.StateOptions
type RootNode = node.TreeNode

type Composable = api.Composable
type Composer = api.Composer

type ProvidedValue = api.ProvidedValue

var Sequence = sequence.Sequence
