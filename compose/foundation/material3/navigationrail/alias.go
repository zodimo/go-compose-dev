package navigationrail

import (
	"github.com/zodimo/go-compose/internal/modifier"
	"github.com/zodimo/go-compose/internal/state"
	"github.com/zodimo/go-compose/internal/theme"
	"github.com/zodimo/go-compose/pkg/api"
)

type Modifier = modifier.Modifier

var EmptyModifier = modifier.EmptyModifier

type Composable = api.Composable
type Composer = api.Composer

type ThemeManager = theme.ThemeManager

type MutableValue = state.MutableValue
