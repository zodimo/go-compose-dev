package menu

import (
	"github.com/zodimo/go-compose/theme"
)

type DropdownMenuColors struct {
	ContainerColor theme.ColorDescriptor
}

type DropdownMenuItemColors struct {
	TextColor                 theme.ColorDescriptor
	LeadingIconColor          theme.ColorDescriptor
	TrailingIconColor         theme.ColorDescriptor
	DisabledTextColor         theme.ColorDescriptor
	DisabledLeadingIconColor  theme.ColorDescriptor
	DisabledTrailingIconColor theme.ColorDescriptor
}

func DefaultDropdownMenuColors() DropdownMenuColors {
	return DropdownMenuColors{
		ContainerColor: theme.ColorHelper.ColorSelector().SurfaceRoles.Container, // Elevation Level 2 default
	}
}

func DefaultDropdownMenuItemColors() DropdownMenuItemColors {

	return DropdownMenuItemColors{
		TextColor:                 theme.ColorHelper.ColorSelector().SurfaceRoles.OnSurface,                  // m3.Scheme.Surface.OnColor,
		LeadingIconColor:          theme.ColorHelper.ColorSelector().SurfaceRoles.OnVariant,                  //m3.Scheme.SurfaceVariant.OnColor,
		TrailingIconColor:         theme.ColorHelper.ColorSelector().SurfaceRoles.OnVariant,                  //m3.Scheme.SurfaceVariant.OnColor,
		DisabledTextColor:         theme.ColorHelper.ColorSelector().SurfaceRoles.OnSurface.SetOpacity(0.38), //m3.Scheme.Surface.OnColor.SetOpacity(0.38),
		DisabledLeadingIconColor:  theme.ColorHelper.ColorSelector().SurfaceRoles.OnSurface.SetOpacity(0.38), //m3.Scheme.Surface.OnColor.SetOpacity(0.38),
		DisabledTrailingIconColor: theme.ColorHelper.ColorSelector().SurfaceRoles.OnSurface.SetOpacity(0.38), //m3.Scheme.Surface.OnColor.SetOpacity(0.38),
	}
}
