package box

import (
	"github.com/zodimo/go-compose/internal/modifier"
	"github.com/zodimo/go-compose/internal/modifiers/box"

	"github.com/zodimo/go-compose/pkg/api"

	"gioui.org/layout"
)

type Modifier = modifier.Modifier

var EmptyModifier = modifier.EmptyModifier

type Composable = api.Composable
type Composer = api.Composer

var MatchParentSizeKey = box.MatchParentSizeKey

// Direction is the alignment of widgets relative to a containing
// space.
type Direction = layout.Direction

const (
	NW     Direction = layout.NW
	N      Direction = layout.N
	NE     Direction = layout.NE
	E      Direction = layout.E
	SE     Direction = layout.SE
	S      Direction = layout.S
	SW     Direction = layout.SW
	W      Direction = layout.W
	Center Direction = layout.Center
)

type Stack = layout.Stack
type StackChild = layout.StackChild
type LayoutContext = layout.Context
type LayoutDimensions = layout.Dimensions
