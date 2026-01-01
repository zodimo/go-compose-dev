package text

import (
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/compose/ui/next/text/style"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

var DefaultColor = graphics.ColorBlack
var DefaultColorForegroundStyle = style.TextForegroundStyleFromColor(DefaultColor)

var DefaultFontSize = unit.Sp(14)
var DefaultLetterSpacing = unit.Sp(0)
var DefaultBackgroundColor = graphics.ColorTransparent
