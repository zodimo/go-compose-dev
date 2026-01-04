package iconbutton

import (
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/material3/tokens"
	"github.com/zodimo/go-compose/pkg/api"
)

// StandardIconButtonTokens - Ported from androidx/compose/material3/tokens
func StandardIconButtonTokens(c api.Composer) tokens.IconButtonTokens {
	theme := material3.Theme(c)
	return tokens.IconButtonTokens{
		DisabledColor:   theme.ColorScheme().OnSurface,
		DisabledOpacity: 0.38,

		FocusedColor:           theme.ColorScheme().OnSurfaceVariant,
		HoveredColor:           theme.ColorScheme().OnSurfaceVariant,
		Color:                  theme.ColorScheme().OnSurfaceVariant,
		PressedColor:           theme.ColorScheme().OnSurfaceVariant,
		SelectedFocusedColor:   theme.ColorScheme().Primary,
		SelectedHoveredColor:   theme.ColorScheme().Primary,
		SelectedColor:          theme.ColorScheme().Primary,
		SelectedPressedColor:   theme.ColorScheme().Primary,
		UnselectedFocusedColor: theme.ColorScheme().OnSurfaceVariant,
		UnselectedHoveredColor: theme.ColorScheme().OnSurfaceVariant,
		UnselectedColor:        theme.ColorScheme().OnSurfaceVariant,
		UnselectedPressedColor: theme.ColorScheme().OnSurfaceVariant,
	}
}

// ElevatedIconButtonTokens - Ported from gio-mw BuildElevatedTheme
// Uses Primary.Color for enabled label/icon (merged into Color)
func ElevatedIconButtonTokens(c api.Composer) tokens.IconButtonTokens {
	theme := material3.Theme(c)
	return tokens.IconButtonTokens{
		DisabledColor:   theme.ColorScheme().OnSurface,
		DisabledOpacity: 0.64,

		FocusedColor: theme.ColorScheme().Primary,
		HoveredColor: theme.ColorScheme().Primary,
		Color:        theme.ColorScheme().Primary,
		PressedColor: theme.ColorScheme().Primary,

		SelectedFocusedColor:   theme.ColorScheme().Primary,
		SelectedHoveredColor:   theme.ColorScheme().Primary,
		SelectedColor:          theme.ColorScheme().Primary,
		SelectedPressedColor:   theme.ColorScheme().Primary,
		UnselectedFocusedColor: theme.ColorScheme().OnSurfaceVariant,
		UnselectedHoveredColor: theme.ColorScheme().OnSurfaceVariant,
		UnselectedColor:        theme.ColorScheme().OnSurfaceVariant,
		UnselectedPressedColor: theme.ColorScheme().OnSurfaceVariant,
	}
}

// FilledIconButtonTokens - Ported from gio-mw BuildFilledTheme
// Uses Primary.OnColor for enabled label/icon (merged into Color)
func FilledIconButtonTokens(c api.Composer) tokens.IconButtonTokens {
	theme := material3.Theme(c)
	return tokens.IconButtonTokens{
		DisabledColor:   theme.ColorScheme().OnSurface,
		DisabledOpacity: 0.64, // OpacityLevel9

		FocusedColor: theme.ColorScheme().OnPrimary,
		HoveredColor: theme.ColorScheme().OnPrimary,
		Color:        theme.ColorScheme().OnPrimary,
		PressedColor: theme.ColorScheme().OnPrimary,

		SelectedFocusedColor:   theme.ColorScheme().OnPrimary,
		SelectedHoveredColor:   theme.ColorScheme().OnPrimary,
		SelectedColor:          theme.ColorScheme().OnPrimary,
		SelectedPressedColor:   theme.ColorScheme().OnPrimary,
		UnselectedFocusedColor: theme.ColorScheme().OnSurfaceVariant,
		UnselectedHoveredColor: theme.ColorScheme().OnSurfaceVariant,
		UnselectedColor:        theme.ColorScheme().OnSurfaceVariant,
		UnselectedPressedColor: theme.ColorScheme().OnSurfaceVariant,
	}
}

// FilledTonalIconButtonTokens - Ported from gio-mw BuildFilledTonalTheme
// Uses SecondaryContainer.OnColor for enabled label/icon (merged into Color)
func FilledTonalIconButtonTokens(c api.Composer) tokens.IconButtonTokens {
	theme := material3.Theme(c)
	return tokens.IconButtonTokens{
		DisabledColor:   theme.ColorScheme().OnSurface,
		DisabledOpacity: 0.64, // OpacityLevel9

		FocusedColor: theme.ColorScheme().OnSecondaryContainer,
		HoveredColor: theme.ColorScheme().OnSecondaryContainer,
		Color:        theme.ColorScheme().OnSecondaryContainer,
		PressedColor: theme.ColorScheme().OnSecondaryContainer,

		SelectedFocusedColor:   theme.ColorScheme().OnSecondaryContainer,
		SelectedHoveredColor:   theme.ColorScheme().OnSecondaryContainer,
		SelectedColor:          theme.ColorScheme().OnSecondaryContainer,
		SelectedPressedColor:   theme.ColorScheme().OnSecondaryContainer,
		UnselectedFocusedColor: theme.ColorScheme().OnSurfaceVariant,
		UnselectedHoveredColor: theme.ColorScheme().OnSurfaceVariant,
		UnselectedColor:        theme.ColorScheme().OnSurfaceVariant,
		UnselectedPressedColor: theme.ColorScheme().OnSurfaceVariant,
	}
}

// OutlinedIconButtonTokens - Ported from gio-mw BuildOutlinedTheme
// Uses Primary.Color for enabled label/icon (merged into Color)
func OutlinedIconButtonTokens(c api.Composer) tokens.IconButtonTokens {
	theme := material3.Theme(c)
	return tokens.IconButtonTokens{
		DisabledColor:   theme.ColorScheme().OnSurface,
		DisabledOpacity: 0.64, // OpacityLevel9

		FocusedColor: theme.ColorScheme().Primary,
		HoveredColor: theme.ColorScheme().Primary,
		Color:        theme.ColorScheme().Primary,
		PressedColor: theme.ColorScheme().Primary,

		SelectedFocusedColor:   theme.ColorScheme().InverseSurface,
		SelectedHoveredColor:   theme.ColorScheme().InverseSurface,
		SelectedColor:          theme.ColorScheme().InverseSurface,
		SelectedPressedColor:   theme.ColorScheme().InverseSurface,
		UnselectedFocusedColor: theme.ColorScheme().OnSurfaceVariant,
		UnselectedHoveredColor: theme.ColorScheme().OnSurfaceVariant,
		UnselectedColor:        theme.ColorScheme().OnSurfaceVariant,
		UnselectedPressedColor: theme.ColorScheme().OnSurfaceVariant,
	}
}
