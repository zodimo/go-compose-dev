package menu

import (
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/ui/graphics"
)

type DropdownMenuColors struct {
	ContainerColor graphics.Color
}

type DropdownMenuItemColors struct {
	TextColor                 graphics.Color
	LeadingIconColor          graphics.Color
	TrailingIconColor         graphics.Color
	DisabledTextColor         graphics.Color
	DisabledLeadingIconColor  graphics.Color
	DisabledTrailingIconColor graphics.Color
}

func DefaultDropdownMenuColors(c Composer) DropdownMenuColors {
	return DropdownMenuColors{
		ContainerColor: material3.Theme(c).ColorScheme().SurfaceContainer, //theme.ColorHelper.ColorSelector().SurfaceRoles.Container, // Elevation Level 2 default
	}
}

func DefaultDropdownMenuItemColors(c Composer) DropdownMenuItemColors {

	return DropdownMenuItemColors{
		TextColor:                 material3.Theme(c).ColorScheme().OnSurface,                            //theme.ColorHelper.ColorSelector().SurfaceRoles.OnSurface,                  // m3.Scheme.Surface.OnColor,
		LeadingIconColor:          material3.Theme(c).ColorScheme().OnSurfaceVariant,                     //theme.ColorHelper.ColorSelector().SurfaceRoles.OnVariant,                  //m3.Scheme.SurfaceVariant.OnColor,
		TrailingIconColor:         material3.Theme(c).ColorScheme().OnSurfaceVariant,                     // theme.ColorHelper.ColorSelector().SurfaceRoles.OnVariant,                  //m3.Scheme.SurfaceVariant.OnColor,
		DisabledTextColor:         graphics.SetOpacity(material3.Theme(c).ColorScheme().OnSurface, 0.38), //theme.ColorHelper.ColorSelector().SurfaceRoles.OnSurface.SetOpacity(0.38), //m3.Scheme.Surface.OnColor.SetOpacity(0.38),
		DisabledLeadingIconColor:  graphics.SetOpacity(material3.Theme(c).ColorScheme().OnSurface, 0.38), //theme.ColorHelper.ColorSelector().SurfaceRoles.OnSurface.SetOpacity(0.38), //m3.Scheme.Surface.OnColor.SetOpacity(0.38),
		DisabledTrailingIconColor: graphics.SetOpacity(material3.Theme(c).ColorScheme().OnSurface, 0.38), // theme.ColorHelper.ColorSelector().SurfaceRoles.OnSurface.SetOpacity(0.38), //m3.Scheme.Surface.OnColor.SetOpacity(0.38),
	}
}
