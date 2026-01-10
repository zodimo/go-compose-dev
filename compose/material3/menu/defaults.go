package menu

import (
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

// MenuDefaults contains default values used for DropdownMenu and DropdownMenuItem.
// Reference: https://m3.material.io/components/menus/specs
var MenuDefaults = menuDefaults{}

type menuDefaults struct{}

// Size constants from Material3 spec (Menu.kt lines 1426-1436)

// MenuVerticalMargin is the vertical margin around the menu.
const MenuVerticalMargin unit.Dp = 48

// MenuHorizontalMargin is the horizontal margin around the menu.
const MenuHorizontalMargin unit.Dp = 48

// MenuListItemContainerHeight is the default height for menu items.
const MenuListItemContainerHeight unit.Dp = 48

// DropdownMenuItemHorizontalPadding is the horizontal padding for menu items.
const DropdownMenuItemHorizontalPadding unit.Dp = 12

// DropdownMenuVerticalPadding is the vertical padding inside the menu container.
const DropdownMenuVerticalPadding unit.Dp = 8

// DropdownMenuItemDefaultMinWidth is the minimum width for menu items.
const DropdownMenuItemDefaultMinWidth unit.Dp = 112

// DropdownMenuItemDefaultMaxWidth is the maximum width for menu items.
const DropdownMenuItemDefaultMaxWidth unit.Dp = 280

// Elevation constants

// TonalElevation is the default tonal elevation for a menu.
const TonalElevation unit.Dp = 0 // ElevationTokens.Level0

// ShadowElevation is the default shadow elevation for a menu.
const ShadowElevation unit.Dp = 3 // MenuTokens.ContainerElevation

// Shape returns the default shape for a menu (4dp rounded corners - Extra Small).
func (menuDefaults) Shape() shape.Shape {
	return &shape.RoundedCornerShape{Radius: 4}
}

// ItemColors returns the default colors for a DropdownMenuItem.
func (menuDefaults) ItemColors(c Composer) DropdownMenuItemColors {
	theme := material3.Theme(c)
	colorScheme := theme.ColorScheme()

	return DropdownMenuItemColors{
		TextColor:                 colorScheme.OnSurface,
		LeadingIconColor:          colorScheme.OnSurfaceVariant,
		TrailingIconColor:         colorScheme.OnSurfaceVariant,
		DisabledTextColor:         graphics.SetOpacity(colorScheme.OnSurface, 0.38),
		DisabledLeadingIconColor:  graphics.SetOpacity(colorScheme.OnSurface, 0.38),
		DisabledTrailingIconColor: graphics.SetOpacity(colorScheme.OnSurface, 0.38),
	}
}

// ContainerColor returns the default container color for a menu.
func (menuDefaults) ContainerColor(c Composer) graphics.Color {
	return material3.Theme(c).ColorScheme().SurfaceContainer
}

// DropdownMenuItemContentPadding returns the default content padding for menu items.
// start:12, top:0, end:12, bottom:0
func (menuDefaults) DropdownMenuItemContentPadding() ContentPadding {
	return ContentPadding{
		Start:  DropdownMenuItemHorizontalPadding,
		Top:    0,
		End:    DropdownMenuItemHorizontalPadding,
		Bottom: 0,
	}
}

// ContentPadding represents padding values for content.
type ContentPadding struct {
	Start  unit.Dp
	Top    unit.Dp
	End    unit.Dp
	Bottom unit.Dp
}
