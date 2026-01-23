package tree

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/row"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/foundation/lazy"
	fText "github.com/zodimo/go-compose/compose/foundation/text"
	m3Text "github.com/zodimo/go-compose/compose/material3/text"
	"github.com/zodimo/go-compose/modifiers/clickable"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
)

// TreeScope defines the operations available within a Tree.
type TreeScope interface {
	// Node adds a leaf node to the tree.
	Node(key any, content api.Composable)

	// Branch adds a collapsible branch node to the tree.
	// header is the content displayed for the branch itself.
	// children is a function that defines the children of this branch.
	Branch(key any, header api.Composable, children func(TreeScope))
}

// Tree creates a tree component that efficiently renders hierarchical data using LazyColumn.
// It supports expanding and collapsing branches.
func Tree(
	state *TreeState,
	content func(TreeScope),
	options ...lazy.LazyListOption,
) api.Composable {
	return lazy.LazyColumn(
		func(scope lazy.LazyListScope) {
			tScope := &treeScopeImpl{
				listScope: scope,
				state:     state,
				depth:     0,
			}
			content(tScope)
		},
		options...,
	)
}

type treeScopeImpl struct {
	listScope lazy.LazyListScope
	state     *TreeState
	depth     int
}

func (s *treeScopeImpl) Node(key any, content api.Composable) {
	s.listScope.Item(key, func(c api.Composer) api.Composer {
		return row.Row(
			func(c api.Composer) api.Composer {
				// Indentation
				spacer.Width(s.depth * 24)(c)

				// Spacer for the expander icon alignment
				spacer.Width(24)(c)

				// Node content
				content(c)
				return c
			},
			row.WithModifier(size.FillMaxWidth().Then(padding.All(4))),
		)(c)
	})
}

func (s *treeScopeImpl) Branch(key any, header api.Composable, children func(TreeScope)) {
	isExpanded := s.state.IsExpanded(key)

	// Branch Header
	s.listScope.Item(key, func(c api.Composer) api.Composer {
		return row.Row(
			func(c api.Composer) api.Composer {
				// Indentation
				spacer.Width(s.depth * 24)(c)

				// Expander Icon
				icon := "▶" // Right pointer
				if isExpanded {
					icon = "▼" // Down pointer
				}

				// Toggle Button
				// We wrap it in a clickable modifier
				m3Text.TextWithStyle(
					icon,
					m3Text.TypestyleBodyMedium,
					fText.WithModifier(
						clickable.OnClick(func() {
							s.state.Toggle(key)
						}).Then(size.Width(24)),
					),
				)(c)

				// Header Content
				header(c)
				return c
			},
			row.WithModifier(
				size.FillMaxWidth().
					Then(padding.All(4)).
					Then(clickable.OnClick(func() {
						s.state.Toggle(key)
					})),
			),
		)(c)
	})

	// Children
	if isExpanded {
		childScope := &treeScopeImpl{
			listScope: s.listScope,
			state:     s.state,
			depth:     s.depth + 1,
		}
		children(childScope)
	}
}

// TreeFromData creates a Tree from a data structure, similar to fyne's data-driven Tree.
// roots: List of root IDs.
// childUIDs: Function to get children IDs for a given ID.
// isBranch: Function to strict check if a node is a branch (optional, if nil, checks if children > 0).
// createNode: Composable factory for a node ID.
func TreeFromData(
	state *TreeState,
	roots []any,
	childUIDs func(any) []any,
	isBranch func(any) bool,
	createNode func(any) api.Composable,
	options ...lazy.LazyListOption,
) api.Composable {
	return Tree(state, func(scope TreeScope) {
		var traverse func(s TreeScope, ids []any)
		traverse = func(s TreeScope, ids []any) {
			for _, id := range ids {
				isB := false
				if isBranch != nil {
					isB = isBranch(id)
				} else {
					children := childUIDs(id)
					isB = len(children) > 0
				}

				if isB {
					s.Branch(id, createNode(id), func(innerS TreeScope) {
						traverse(innerS, childUIDs(id))
					})
				} else {
					s.Node(id, createNode(id))
				}
			}
		}

		traverse(scope, roots)
	}, options...)
}
