package text

import (
	"github.com/zodimo/go-compose/compose/ui/next/text"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/pkg/api"
)

type Composable = api.Composable
type Composer = api.Composer

type TextStyle = text.TextStyle

type uiTextLayoutResult = text.TextLayoutResult
type uiPlaceholder = text.Placeholder
type uiTextUnit = unit.TextUnit

var uiTextUnitSp = unit.Sp

type uiAnnotatedString = text.AnnotatedString

type uiConstraints = unit.Constraints
