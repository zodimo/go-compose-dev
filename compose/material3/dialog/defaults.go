package dialog

import (
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
	"github.com/zodimo/go-compose/compose/ui/unit"
	"github.com/zodimo/go-compose/pkg/api"
)

// DialogDefaults contains default values for AlertDialog and BasicAlertDialog styling.
// Follows Material Design 3 dialog specifications.
var DialogDefaults = struct {
	// Shape is the default shape for alert dialogs (ExtraLarge - 28dp radius).
	Shape shape.Shape
	// MinWidth is the minimum dialog width.
	MinWidth unit.Dp
	// MaxWidth is the maximum dialog width.
	MaxWidth unit.Dp
	// TonalElevation is the default tonal elevation (0dp for dialogs).
	TonalElevation unit.Dp
	// ShadowElevation is the default shadow elevation (Level 3).
	ShadowElevation unit.Dp
}{
	Shape:           &shape.RoundedCornerShape{Radius: unit.Dp(28)},
	MinWidth:        unit.Dp(280),
	MaxWidth:        unit.Dp(560),
	TonalElevation:  unit.Dp(0),
	ShadowElevation: unit.Dp(6), // Elevation Level 3
}

// DialogPadding contains the spacing values for dialog content.
var DialogPadding = struct {
	// All is the uniform padding around dialog content.
	All unit.Dp
	// IconBottom is the spacing below the icon.
	IconBottom unit.Dp
	// TitleBottom is the spacing below the title.
	TitleBottom unit.Dp
	// TextBottom is the spacing below the supporting text.
	TextBottom unit.Dp
	// ButtonSpacing is the spacing between buttons.
	ButtonSpacing unit.Dp
}{
	All:           unit.Dp(24),
	IconBottom:    unit.Dp(16),
	TitleBottom:   unit.Dp(16),
	TextBottom:    unit.Dp(24),
	ButtonSpacing: unit.Dp(8),
}

// ContainerColor returns the default container color for dialogs.
func ContainerColor(c api.Composer) graphics.Color {
	return material3.Theme(c).ColorScheme().SurfaceContainerHigh
}

// IconContentColor returns the default icon content color for dialogs.
func IconContentColor(c api.Composer) graphics.Color {
	return material3.Theme(c).ColorScheme().Secondary
}

// TitleContentColor returns the default title content color for dialogs.
func TitleContentColor(c api.Composer) graphics.Color {
	return material3.Theme(c).ColorScheme().OnSurface
}

// TextContentColor returns the default supporting text content color for dialogs.
func TextContentColor(c api.Composer) graphics.Color {
	return material3.Theme(c).ColorScheme().OnSurfaceVariant
}
