package theme

import (
	"go-compose-dev/internal/theme/colorrole"

	"gioui.org/widget/material"
	"git.sr.ht/~schnwalter/gio-mw/token"
)

type ColorRole = colorrole.ColorRole

type Theme = token.Theme
type BasicTheme = material.Theme

type TokenColor = token.MatColor

type OpacityLevel = token.OpacityLevel
