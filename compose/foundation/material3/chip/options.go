package chip

import (
	"go-compose-dev/compose/ui/graphics/shape"
	"go-compose-dev/internal/modifier"
	"go-compose-dev/pkg/api"
	"image/color"

	"gioui.org/unit"
)

type ChipOptions struct {
	Modifier     modifier.Modifier
	Shape        shape.Shape
	Color        color.NRGBA
	BorderColor  color.NRGBA
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
type Modifier = modifier.Modifier

func DefaultChipOptions() ChipOptions {
	return ChipOptions{
		Modifier:    modifier.EmptyModifier,
		Shape:       shape.RoundedCornerShape{Radius: unit.Dp(8)}, // Material 3 small rounding usually
		Color:       color.NRGBA{R: 0, G: 0, B: 0, A: 0},          // Transparent by default, surface handles it
		BorderWidth: unit.Dp(1),
		BorderColor: color.NRGBA{R: 0x79, G: 0x74, B: 0x7E, A: 0xFF}, // Outline variant
		Height:      unit.Dp(32),
		Enabled:     true,
		Selected:    false,
	}
}

func WithModifier(m Modifier) ChipOption {
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

func WithColor(c color.NRGBA) ChipOption {
	return func(o *ChipOptions) {
		o.Color = c
	}
}

func WithBorder(width unit.Dp, color color.NRGBA) ChipOption {
	return func(o *ChipOptions) {
		o.BorderWidth = width
		o.BorderColor = color
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
