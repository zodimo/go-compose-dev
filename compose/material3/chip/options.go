package chip

import (
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
	"github.com/zodimo/go-compose/pkg/api"

	"github.com/zodimo/go-compose/compose/ui/unit"
)

type ChipOptions struct {
	Modifier     ui.Modifier
	Shape        shape.Shape
	Color        graphics.Color
	BorderColor  graphics.Color
	BorderWidth  unit.Dp
	Elevation    unit.Dp
	Height       unit.Dp
	LeadingIcon  api.Composable
	TrailingIcon api.Composable
	Enabled      bool
	Selected     bool
}

type ChipOption func(*ChipOptions)

type Composable = api.Composable
type Composer = api.Composer

func DefaultChipOptions(c Composer) ChipOptions {
	theme := material3.Theme(c)
	return ChipOptions{
		Modifier:    ui.EmptyModifier,
		Shape:       &shape.RoundedCornerShape{Radius: unit.Dp(8)}, // Material 3 small rounding usually
		Color:       theme.ColorScheme().Surface,                   //theme.ColorHelper.ColorSelector().SurfaceRoles.Surface, // Default to Surface
		BorderWidth: unit.Dp(1),
		BorderColor: theme.ColorScheme().OutlineVariant, //theme.ColorHelper.ColorSelector().OutlineRoles.OutlineVariant, // Outline variant
		Height:      unit.Dp(32),
		Enabled:     true,
		Selected:    false,
	}
}

func WithModifier(m ui.Modifier) ChipOption {
	return func(o *ChipOptions) {
		o.Modifier = m
	}
}

func WithEnabled(enabled bool) ChipOption {
	return func(o *ChipOptions) {
		o.Enabled = enabled
	}
}

func WithSelected(selected bool) ChipOption {
	return func(o *ChipOptions) {
		o.Selected = selected
	}
}

func WithLeadingIcon(icon api.Composable) ChipOption {
	return func(o *ChipOptions) {
		o.LeadingIcon = icon
	}
}

func WithTrailingIcon(icon api.Composable) ChipOption {
	return func(o *ChipOptions) {
		o.TrailingIcon = icon
	}
}

func WithColor(col graphics.Color) ChipOption {
	return func(o *ChipOptions) {
		o.Color = col
	}
}

func WithBorder(width unit.Dp, col graphics.Color) ChipOption {
	return func(o *ChipOptions) {
		o.BorderWidth = width
		o.BorderColor = col
	}
}

func WithShape(s shape.Shape) ChipOption {
	return func(o *ChipOptions) {
		o.Shape = s
	}
}

func WithElevation(elevation unit.Dp) ChipOption {
	return func(o *ChipOptions) {
		o.Elevation = elevation
	}
}
