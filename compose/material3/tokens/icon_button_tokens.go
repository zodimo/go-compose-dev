package tokens

import (
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/pkg/api"
)

type IconButtonTokens struct {
	DisabledColor          graphics.Color
	DisabledOpacity        float32
	FocusedColor           graphics.Color
	HoveredColor           graphics.Color
	Color                  graphics.Color
	PressedColor           graphics.Color
	SelectedFocusedColor   graphics.Color
	SelectedHoveredColor   graphics.Color
	SelectedColor          graphics.Color
	SelectedPressedColor   graphics.Color
	UnselectedFocusedColor graphics.Color
	UnselectedHoveredColor graphics.Color
	UnselectedColor        graphics.Color
	UnselectedPressedColor graphics.Color
}

// from androidx/compose/material3/material3/src/commonMain/kotlin/androidx/compose/material3/tokens
func StandardIconButtonTokens(c api.Composer) IconButtonTokens {
	theme := material3.Theme(c)
	return IconButtonTokens{
		DisabledColor:   theme.ColorScheme().Surface.OnColor,
		DisabledOpacity: 0.38,

		FocusedColor:           theme.ColorScheme().SurfaceVariant.OnColor,
		HoveredColor:           theme.ColorScheme().SurfaceVariant.OnColor,
		Color:                  theme.ColorScheme().SurfaceVariant.OnColor,
		PressedColor:           theme.ColorScheme().SurfaceVariant.OnColor,
		SelectedFocusedColor:   theme.ColorScheme().Primary.Color,
		SelectedHoveredColor:   theme.ColorScheme().Primary.Color,
		SelectedColor:          theme.ColorScheme().Primary.Color,
		SelectedPressedColor:   theme.ColorScheme().Primary.Color,
		UnselectedFocusedColor: theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedHoveredColor: theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedColor:        theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedPressedColor: theme.ColorScheme().SurfaceVariant.OnColor,
	}
}

// ElevatedIconButtonTokens - Ported from gio-mw BuildElevatedTheme
// Uses Primary.Color for enabled label/icon (merged into Color)
func ElevatedIconButtonTokens(c api.Composer) IconButtonTokens {
	theme := material3.Theme(c)
	return IconButtonTokens{
		DisabledColor:   theme.ColorScheme().Surface.OnColor,
		DisabledOpacity: 0.64,

		FocusedColor: theme.ColorScheme().Primary.Color,
		HoveredColor: theme.ColorScheme().Primary.Color,
		Color:        theme.ColorScheme().Primary.Color,
		PressedColor: theme.ColorScheme().Primary.Color,

		SelectedFocusedColor:   theme.ColorScheme().Primary.Color,
		SelectedHoveredColor:   theme.ColorScheme().Primary.Color,
		SelectedColor:          theme.ColorScheme().Primary.Color,
		SelectedPressedColor:   theme.ColorScheme().Primary.Color,
		UnselectedFocusedColor: theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedHoveredColor: theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedColor:        theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedPressedColor: theme.ColorScheme().SurfaceVariant.OnColor,
	}
}

// FilledIconButtonTokens - Ported from gio-mw BuildFilledTheme
// Uses Primary.OnColor for enabled label/icon (merged into Color)
func FilledIconButtonTokens(c api.Composer) IconButtonTokens {
	theme := material3.Theme(c)
	return IconButtonTokens{
		DisabledColor:   theme.ColorScheme().Surface.OnColor,
		DisabledOpacity: 0.64, // OpacityLevel9

		FocusedColor: theme.ColorScheme().Primary.OnColor,
		HoveredColor: theme.ColorScheme().Primary.OnColor,
		Color:        theme.ColorScheme().Primary.OnColor,
		PressedColor: theme.ColorScheme().Primary.OnColor,

		SelectedFocusedColor:   theme.ColorScheme().Primary.OnColor,
		SelectedHoveredColor:   theme.ColorScheme().Primary.OnColor,
		SelectedColor:          theme.ColorScheme().Primary.OnColor,
		SelectedPressedColor:   theme.ColorScheme().Primary.OnColor,
		UnselectedFocusedColor: theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedHoveredColor: theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedColor:        theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedPressedColor: theme.ColorScheme().SurfaceVariant.OnColor,
	}
}

// FilledTonalIconButtonTokens - Ported from gio-mw BuildFilledTonalTheme
// Uses SecondaryContainer.OnColor for enabled label/icon (merged into Color)
func FilledTonalIconButtonTokens(c api.Composer) IconButtonTokens {
	theme := material3.Theme(c)
	return IconButtonTokens{
		DisabledColor:   theme.ColorScheme().Surface.OnColor,
		DisabledOpacity: 0.64, // OpacityLevel9

		FocusedColor: theme.ColorScheme().SecondaryContainer.OnColor,
		HoveredColor: theme.ColorScheme().SecondaryContainer.OnColor,
		Color:        theme.ColorScheme().SecondaryContainer.OnColor,
		PressedColor: theme.ColorScheme().SecondaryContainer.OnColor,

		SelectedFocusedColor:   theme.ColorScheme().SecondaryContainer.OnColor,
		SelectedHoveredColor:   theme.ColorScheme().SecondaryContainer.OnColor,
		SelectedColor:          theme.ColorScheme().SecondaryContainer.OnColor,
		SelectedPressedColor:   theme.ColorScheme().SecondaryContainer.OnColor,
		UnselectedFocusedColor: theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedHoveredColor: theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedColor:        theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedPressedColor: theme.ColorScheme().SurfaceVariant.OnColor,
	}
}

// OutlinedIconButtonTokens - Ported from gio-mw BuildOutlinedTheme
// Uses Primary.Color for enabled label/icon (merged into Color)
func OutlinedIconButtonTokens(c api.Composer) IconButtonTokens {
	theme := material3.Theme(c)
	return IconButtonTokens{
		DisabledColor:   theme.ColorScheme().Surface.OnColor,
		DisabledOpacity: 0.64, // OpacityLevel9

		FocusedColor: theme.ColorScheme().Primary.Color,
		HoveredColor: theme.ColorScheme().Primary.Color,
		Color:        theme.ColorScheme().Primary.Color,
		PressedColor: theme.ColorScheme().Primary.Color,

		SelectedFocusedColor:   theme.ColorScheme().InverseSurface.Color,
		SelectedHoveredColor:   theme.ColorScheme().InverseSurface.Color,
		SelectedColor:          theme.ColorScheme().InverseSurface.Color,
		SelectedPressedColor:   theme.ColorScheme().InverseSurface.Color,
		UnselectedFocusedColor: theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedHoveredColor: theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedColor:        theme.ColorScheme().SurfaceVariant.OnColor,
		UnselectedPressedColor: theme.ColorScheme().SurfaceVariant.OnColor,
	}
}
