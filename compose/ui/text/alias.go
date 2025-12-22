package text

import (
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/compose/ui/text/font"
	"github.com/zodimo/go-compose/compose/ui/text/intl"
	"github.com/zodimo/go-compose/compose/ui/text/style"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/theme"
)

type Color = theme.ColorDescriptor

type uiTextUnit = unit.TextUnit
type uiIntSize = unit.IntSize
type uiTextGeometricTransform = style.TextGeometricTransform
type uiShadow = graphics.Shadow
type uiLocaleList = intl.LocaleList
type uiTextDecoration = style.TextDecoration
type uiDirection = style.TextDirection
type uiIndent = style.TextIndent
type uiFontSynthesis = font.FontSynthesis
type uiGraphicsBrush = graphics.Brush
