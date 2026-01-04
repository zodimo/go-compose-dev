package textfield

import (
	"github.com/zodimo/go-compose/compose/material3"
	"github.com/zodimo/go-compose/compose/ui/graphics"
)

type TextFieldColors struct {
	TextColor                   graphics.Color
	DisabledTextColor           graphics.Color
	CursorColor                 graphics.Color
	ErrorCursorColor            graphics.Color
	SelectionColor              graphics.Color
	FocusedIndicatorColor       graphics.Color
	UnfocusedIndicatorColor     graphics.Color
	DisabledIndicatorColor      graphics.Color
	ErrorIndicatorColor         graphics.Color
	HoveredIndicatorColor       graphics.Color
	LeadingIconColor            graphics.Color
	TrailingIconColor           graphics.Color
	DisabledLeadingIconColor    graphics.Color
	DisabledTrailingIconColor   graphics.Color
	LabelColor                  graphics.Color
	UnfocusedLabelColor         graphics.Color
	FocusedLabelColor           graphics.Color // Added
	DisabledLabelColor          graphics.Color
	ErrorLabelColor             graphics.Color
	PlaceholderColor            graphics.Color
	DisabledPlaceholderColor    graphics.Color
	SupportingTextColor         graphics.Color
	DisabledSupportingTextColor graphics.Color
	ErrorSupportingTextColor    graphics.Color
	ContainerColor              graphics.Color // Added
	DisabledContainerColor      graphics.Color // Added
}

func DefaultTextFieldColors() TextFieldColors {
	return TextFieldColors{
		TextColor:                   graphics.ColorUnspecified,
		DisabledTextColor:           graphics.ColorUnspecified,
		CursorColor:                 graphics.ColorUnspecified,
		ErrorCursorColor:            graphics.ColorUnspecified,
		SelectionColor:              graphics.ColorUnspecified,
		FocusedIndicatorColor:       graphics.ColorUnspecified,
		UnfocusedIndicatorColor:     graphics.ColorUnspecified,
		DisabledIndicatorColor:      graphics.ColorUnspecified,
		ErrorIndicatorColor:         graphics.ColorUnspecified,
		HoveredIndicatorColor:       graphics.ColorUnspecified,
		LeadingIconColor:            graphics.ColorUnspecified,
		TrailingIconColor:           graphics.ColorUnspecified,
		DisabledLeadingIconColor:    graphics.ColorUnspecified,
		DisabledTrailingIconColor:   graphics.ColorUnspecified,
		LabelColor:                  graphics.ColorUnspecified,
		UnfocusedLabelColor:         graphics.ColorUnspecified,
		FocusedLabelColor:           graphics.ColorUnspecified, // Added
		DisabledLabelColor:          graphics.ColorUnspecified,
		ErrorLabelColor:             graphics.ColorUnspecified,
		PlaceholderColor:            graphics.ColorUnspecified,
		DisabledPlaceholderColor:    graphics.ColorUnspecified,
		SupportingTextColor:         graphics.ColorUnspecified,
		DisabledSupportingTextColor: graphics.ColorUnspecified,
		ErrorSupportingTextColor:    graphics.ColorUnspecified,
		ContainerColor:              graphics.ColorUnspecified, // Added
		DisabledContainerColor:      graphics.ColorUnspecified, // Added
	}
}

func ResolveTextFieldColors(c Composer, colors TextFieldColors) TextFieldColors {

	theme := material3.Theme(c)
	return TextFieldColors{
		TextColor:                   colors.TextColor.TakeOrElse(theme.ColorScheme().OnSurface),
		DisabledTextColor:           colors.DisabledTextColor.TakeOrElse(graphics.SetOpacity(theme.ColorScheme().OnSurface, 0.38)),
		CursorColor:                 colors.CursorColor.TakeOrElse(theme.ColorScheme().Primary),
		ErrorCursorColor:            colors.ErrorCursorColor.TakeOrElse(theme.ColorScheme().Error),
		SelectionColor:              colors.SelectionColor.TakeOrElse(theme.ColorScheme().Primary),
		FocusedIndicatorColor:       colors.FocusedIndicatorColor.TakeOrElse(theme.ColorScheme().Primary),
		UnfocusedIndicatorColor:     colors.UnfocusedIndicatorColor.TakeOrElse(theme.ColorScheme().Outline),
		DisabledIndicatorColor:      colors.DisabledIndicatorColor.TakeOrElse(graphics.SetOpacity(theme.ColorScheme().OnSurface, 0.12)),
		ErrorIndicatorColor:         colors.ErrorIndicatorColor.TakeOrElse(theme.ColorScheme().Error),
		HoveredIndicatorColor:       colors.HoveredIndicatorColor.TakeOrElse(theme.ColorScheme().OnSurface),
		LeadingIconColor:            colors.LeadingIconColor.TakeOrElse(theme.ColorScheme().OnSurfaceVariant),
		TrailingIconColor:           colors.TrailingIconColor.TakeOrElse(theme.ColorScheme().OnSurfaceVariant),
		DisabledLeadingIconColor:    colors.DisabledLeadingIconColor.TakeOrElse(graphics.SetOpacity(theme.ColorScheme().OnSurface, 0.38)),
		DisabledTrailingIconColor:   colors.DisabledTrailingIconColor.TakeOrElse(graphics.SetOpacity(theme.ColorScheme().OnSurface, 0.38)),
		LabelColor:                  colors.LabelColor.TakeOrElse(theme.ColorScheme().OnSurfaceVariant),
		UnfocusedLabelColor:         colors.UnfocusedLabelColor.TakeOrElse(theme.ColorScheme().OnSurfaceVariant),
		FocusedLabelColor:           colors.FocusedLabelColor.TakeOrElse(theme.ColorScheme().Primary), // Added
		DisabledLabelColor:          colors.DisabledLabelColor.TakeOrElse(graphics.SetOpacity(theme.ColorScheme().OnSurface, 0.38)),
		ErrorLabelColor:             colors.ErrorLabelColor.TakeOrElse(theme.ColorScheme().Error),
		PlaceholderColor:            colors.PlaceholderColor.TakeOrElse(theme.ColorScheme().OnSurfaceVariant),
		DisabledPlaceholderColor:    colors.DisabledPlaceholderColor.TakeOrElse(graphics.SetOpacity(theme.ColorScheme().OnSurface, 0.38)),
		SupportingTextColor:         colors.SupportingTextColor.TakeOrElse(theme.ColorScheme().OnSurfaceVariant),
		DisabledSupportingTextColor: colors.DisabledSupportingTextColor.TakeOrElse(graphics.SetOpacity(theme.ColorScheme().OnSurface, 0.38)),
		ErrorSupportingTextColor:    colors.ErrorSupportingTextColor.TakeOrElse(theme.ColorScheme().Error),
		ContainerColor:              colors.ContainerColor.TakeOrElse(theme.ColorScheme().SurfaceVariant),                               // Added - default for filled
		DisabledContainerColor:      colors.DisabledContainerColor.TakeOrElse(graphics.SetOpacity(theme.ColorScheme().OnSurface, 0.12)), // Added - approximate default? Or SurfaceVariant with opacity
	}
}
