package icon

import (
	"image/color"

	"github.com/zodimo/go-compose/pkg/api"

	"gioui.org/layout"
	// . "golang.org/x/exp/shiny/materialdesign/icons"
)

type Composable = api.Composable
type Composer = api.Composer

type layoutContext = layout.Context
type layoutDimensions = layout.Dimensions
type IconWidget = func(gtx layoutContext, foreground color.NRGBA) layoutDimensions
