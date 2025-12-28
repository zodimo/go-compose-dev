package icon

import (
	"image/color"

	"github.com/zodimo/go-compose/internal/modifier"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/theme"

	"gioui.org/layout"
)

type Modifier = modifier.Modifier

var EmptyModifier = modifier.EmptyModifier

type Composable = api.Composable
type Composer = api.Composer

type layoutContext = layout.Context
type layoutDimensions = layout.Dimensions
type IconWidget = func(gtx layoutContext, foreground color.NRGBA) layoutDimensions

type ColorDescriptor = theme.ColorDescriptor

var colorHelper = theme.ColorHelper
var themeManager = theme.GetThemeManager()
