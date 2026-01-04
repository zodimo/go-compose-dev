package material3

import (
	"github.com/zodimo/go-compose/compose/material3/tokens"
	"github.com/zodimo/go-compose/compose/ui/graphics"
)

// Immutable
type ColorSchemeNext struct {
	Primary                 graphics.Color
	OnPrimary               graphics.Color
	PrimaryContainer        graphics.Color
	OnPrimaryContainer      graphics.Color
	InversePrimary          graphics.Color
	Secondary               graphics.Color
	OnSecondary             graphics.Color
	SecondaryContainer      graphics.Color
	OnSecondaryContainer    graphics.Color
	Tertiary                graphics.Color
	OnTertiary              graphics.Color
	TertiaryContainer       graphics.Color
	OnTertiaryContainer     graphics.Color
	Background              graphics.Color
	OnBackground            graphics.Color
	Surface                 graphics.Color
	OnSurface               graphics.Color
	SurfaceVariant          graphics.Color
	OnSurfaceVariant        graphics.Color
	SurfaceTint             graphics.Color
	InverseSurface          graphics.Color
	InverseOnSurface        graphics.Color
	Error                   graphics.Color
	OnError                 graphics.Color
	ErrorContainer          graphics.Color
	OnErrorContainer        graphics.Color
	Outline                 graphics.Color
	OutlineVariant          graphics.Color
	Scrim                   graphics.Color
	SurfaceBright           graphics.Color
	SurfaceContainer        graphics.Color
	SurfaceContainerHigh    graphics.Color
	SurfaceContainerHighest graphics.Color
	SurfaceContainerLow     graphics.Color
	SurfaceContainerLowest  graphics.Color
	SurfaceDim              graphics.Color
	PrimaryFixed            graphics.Color
	PrimaryFixedDim         graphics.Color
	OnPrimaryFixed          graphics.Color
	OnPrimaryFixedVariant   graphics.Color
	SecondaryFixed          graphics.Color
	SecondaryFixedDim       graphics.Color
	OnSecondaryFixed        graphics.Color
	OnSecondaryFixedVariant graphics.Color
	TertiaryFixed           graphics.Color
	TertiaryFixedDim        graphics.Color
	OnTertiaryFixed         graphics.Color
	OnTertiaryFixedVariant  graphics.Color
}

func (c ColorSchemeNext) ContentFor(backgroundColor graphics.Color) graphics.Color {
	switch backgroundColor {
	case c.Primary:
		return c.OnPrimary
	case c.Secondary:
		return c.OnSecondary
	case c.Tertiary:
		return c.OnTertiary
	case c.Background:
		return c.OnBackground
	case c.Error:
		return c.OnError

	case c.PrimaryContainer:
		return c.OnPrimaryContainer
	case c.SecondaryContainer:
		return c.OnSecondaryContainer
	case c.TertiaryContainer:
		return c.OnTertiaryContainer
	case c.ErrorContainer:
		return c.OnErrorContainer

	case c.InverseSurface:
		return c.InverseOnSurface
	case c.Surface:
		return c.OnSurface
	case c.SurfaceVariant:
		return c.OnSurfaceVariant
	case c.SurfaceBright:
		return c.OnSurface
	case c.SurfaceContainer:
		return c.OnSurface
	case c.SurfaceContainerHigh:
		return c.OnSurface
	case c.SurfaceContainerHighest:
		return c.OnSurface
	case c.SurfaceContainerLow:
		return c.OnSurface
	case c.SurfaceContainerLowest:
		return c.OnSurface
	case c.SurfaceDim:
		return c.OnSurface

	case c.PrimaryFixed:
		return c.OnPrimaryFixed
	case c.PrimaryFixedDim:
		return c.OnPrimaryFixed
	case c.SecondaryFixed:
		return c.OnSecondaryFixed
	case c.SecondaryFixedDim:
		return c.OnSecondaryFixed
	case c.TertiaryFixed:
		return c.OnTertiaryFixed
	case c.TertiaryFixedDim:
		return c.OnTertiaryFixed

	default:
		return graphics.ColorUnspecified

	}
}

func LightColorScheme() ColorSchemeNext {
	return ColorSchemeNext{
		Primary:                 tokens.ColorLightTokens.Primary,
		OnPrimary:               tokens.ColorLightTokens.OnPrimary,
		PrimaryContainer:        tokens.ColorLightTokens.PrimaryContainer,
		OnPrimaryContainer:      tokens.ColorLightTokens.OnPrimaryContainer,
		InversePrimary:          tokens.ColorLightTokens.InversePrimary,
		Secondary:               tokens.ColorLightTokens.Secondary,
		OnSecondary:             tokens.ColorLightTokens.OnSecondary,
		SecondaryContainer:      tokens.ColorLightTokens.SecondaryContainer,
		OnSecondaryContainer:    tokens.ColorLightTokens.OnSecondaryContainer,
		Tertiary:                tokens.ColorLightTokens.Tertiary,
		OnTertiary:              tokens.ColorLightTokens.OnTertiary,
		TertiaryContainer:       tokens.ColorLightTokens.TertiaryContainer,
		OnTertiaryContainer:     tokens.ColorLightTokens.OnTertiaryContainer,
		Background:              tokens.ColorLightTokens.Background,
		OnBackground:            tokens.ColorLightTokens.OnBackground,
		Surface:                 tokens.ColorLightTokens.Surface,
		OnSurface:               tokens.ColorLightTokens.OnSurface,
		SurfaceVariant:          tokens.ColorLightTokens.SurfaceVariant,
		OnSurfaceVariant:        tokens.ColorLightTokens.OnSurfaceVariant,
		SurfaceTint:             tokens.ColorLightTokens.SurfaceTint,
		InverseSurface:          tokens.ColorLightTokens.InverseSurface,
		InverseOnSurface:        tokens.ColorLightTokens.InverseOnSurface,
		Error:                   tokens.ColorLightTokens.Error,
		OnError:                 tokens.ColorLightTokens.OnError,
		ErrorContainer:          tokens.ColorLightTokens.ErrorContainer,
		OnErrorContainer:        tokens.ColorLightTokens.OnErrorContainer,
		Outline:                 tokens.ColorLightTokens.Outline,
		OutlineVariant:          tokens.ColorLightTokens.OutlineVariant,
		Scrim:                   tokens.ColorLightTokens.Scrim,
		SurfaceBright:           tokens.ColorLightTokens.SurfaceBright,
		SurfaceContainer:        tokens.ColorLightTokens.SurfaceContainer,
		SurfaceContainerHigh:    tokens.ColorLightTokens.SurfaceContainerHigh,
		SurfaceContainerHighest: tokens.ColorLightTokens.SurfaceContainerHighest,
		SurfaceContainerLow:     tokens.ColorLightTokens.SurfaceContainerLow,
		SurfaceContainerLowest:  tokens.ColorLightTokens.SurfaceContainerLowest,
		SurfaceDim:              tokens.ColorLightTokens.SurfaceDim,
		PrimaryFixed:            tokens.ColorLightTokens.PrimaryFixed,
		PrimaryFixedDim:         tokens.ColorLightTokens.PrimaryFixedDim,
		OnPrimaryFixed:          tokens.ColorLightTokens.OnPrimaryFixed,
		OnPrimaryFixedVariant:   tokens.ColorLightTokens.OnPrimaryFixedVariant,
		SecondaryFixed:          tokens.ColorLightTokens.SecondaryFixed,
		SecondaryFixedDim:       tokens.ColorLightTokens.SecondaryFixedDim,
		OnSecondaryFixed:        tokens.ColorLightTokens.OnSecondaryFixed,
		OnSecondaryFixedVariant: tokens.ColorLightTokens.OnSecondaryFixedVariant,
		TertiaryFixed:           tokens.ColorLightTokens.TertiaryFixed,
		TertiaryFixedDim:        tokens.ColorLightTokens.TertiaryFixedDim,
		OnTertiaryFixed:         tokens.ColorLightTokens.OnTertiaryFixed,
		OnTertiaryFixedVariant:  tokens.ColorLightTokens.OnTertiaryFixedVariant,
	}
}

func DarkColorScheme() ColorSchemeNext {
	return ColorSchemeNext{
		Primary:                 tokens.ColorDarkTokens.Primary,
		OnPrimary:               tokens.ColorDarkTokens.OnPrimary,
		PrimaryContainer:        tokens.ColorDarkTokens.PrimaryContainer,
		OnPrimaryContainer:      tokens.ColorDarkTokens.OnPrimaryContainer,
		InversePrimary:          tokens.ColorDarkTokens.InversePrimary,
		Secondary:               tokens.ColorDarkTokens.Secondary,
		OnSecondary:             tokens.ColorDarkTokens.OnSecondary,
		SecondaryContainer:      tokens.ColorDarkTokens.SecondaryContainer,
		OnSecondaryContainer:    tokens.ColorDarkTokens.OnSecondaryContainer,
		Tertiary:                tokens.ColorDarkTokens.Tertiary,
		OnTertiary:              tokens.ColorDarkTokens.OnTertiary,
		TertiaryContainer:       tokens.ColorDarkTokens.TertiaryContainer,
		OnTertiaryContainer:     tokens.ColorDarkTokens.OnTertiaryContainer,
		Background:              tokens.ColorDarkTokens.Background,
		OnBackground:            tokens.ColorDarkTokens.OnBackground,
		Surface:                 tokens.ColorDarkTokens.Surface,
		OnSurface:               tokens.ColorDarkTokens.OnSurface,
		SurfaceVariant:          tokens.ColorDarkTokens.SurfaceVariant,
		OnSurfaceVariant:        tokens.ColorDarkTokens.OnSurfaceVariant,
		SurfaceTint:             tokens.ColorDarkTokens.SurfaceTint,
		InverseSurface:          tokens.ColorDarkTokens.InverseSurface,
		InverseOnSurface:        tokens.ColorDarkTokens.InverseOnSurface,
		Error:                   tokens.ColorDarkTokens.Error,
		OnError:                 tokens.ColorDarkTokens.OnError,
		ErrorContainer:          tokens.ColorDarkTokens.ErrorContainer,
		OnErrorContainer:        tokens.ColorDarkTokens.OnErrorContainer,
		Outline:                 tokens.ColorDarkTokens.Outline,
		OutlineVariant:          tokens.ColorDarkTokens.OutlineVariant,
		Scrim:                   tokens.ColorDarkTokens.Scrim,
		SurfaceBright:           tokens.ColorDarkTokens.SurfaceBright,
		SurfaceContainer:        tokens.ColorDarkTokens.SurfaceContainer,
		SurfaceContainerHigh:    tokens.ColorDarkTokens.SurfaceContainerHigh,
		SurfaceContainerHighest: tokens.ColorDarkTokens.SurfaceContainerHighest,
		SurfaceContainerLow:     tokens.ColorDarkTokens.SurfaceContainerLow,
		SurfaceContainerLowest:  tokens.ColorDarkTokens.SurfaceContainerLowest,
		SurfaceDim:              tokens.ColorDarkTokens.SurfaceDim,
		PrimaryFixed:            tokens.ColorDarkTokens.PrimaryFixed,
		PrimaryFixedDim:         tokens.ColorDarkTokens.PrimaryFixedDim,
		OnPrimaryFixed:          tokens.ColorDarkTokens.OnPrimaryFixed,
		OnPrimaryFixedVariant:   tokens.ColorDarkTokens.OnPrimaryFixedVariant,
		SecondaryFixed:          tokens.ColorDarkTokens.SecondaryFixed,
		SecondaryFixedDim:       tokens.ColorDarkTokens.SecondaryFixedDim,
		OnSecondaryFixed:        tokens.ColorDarkTokens.OnSecondaryFixed,
		OnSecondaryFixedVariant: tokens.ColorDarkTokens.OnSecondaryFixedVariant,
		TertiaryFixed:           tokens.ColorDarkTokens.TertiaryFixed,
		TertiaryFixedDim:        tokens.ColorDarkTokens.TertiaryFixedDim,
		OnTertiaryFixed:         tokens.ColorDarkTokens.OnTertiaryFixed,
		OnTertiaryFixedVariant:  tokens.ColorDarkTokens.OnTertiaryFixedVariant,
	}
}
