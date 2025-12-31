package material3

import (
	"git.sr.ht/~schnwalter/gio-mw/token"
	"github.com/zodimo/go-compose/compose/ui/graphics"
)

type ColorScheme struct {
	Primary             *ColorSet
	PrimaryContainer    *ColorSet
	PrimaryFixed        *ColorSet
	PrimaryFixedVariant *ColorSet

	Secondary             *ColorSet
	SecondaryContainer    *ColorSet
	SecondaryFixed        *ColorSet
	SecondaryFixedVariant *ColorSet

	Tertiary             *ColorSet
	TertiaryContainer    *ColorSet
	TertiaryFixed        *ColorSet
	TertiaryFixedVariant *ColorSet

	Surface                 *ColorSet
	SurfaceVariant          *ColorSet
	SurfaceDim              Color
	SurfaceBright           Color
	SurfaceContainerLowest  Color
	SurfaceContainerLow     Color
	SurfaceContainer        Color
	SurfaceContainerHigh    Color
	SurfaceContainerHighest Color

	InverseSurface *ColorSet
	InversePrimary Color

	Background *ColorSet

	Outline        Color
	OutlineVariant Color

	Error          *ColorSet
	ErrorContainer *ColorSet

	Scrim  Color
	Shadow Color
}

func (c *ColorScheme) Copy(options ...ColorSchemeOption) *ColorScheme {
	copy := *c
	for _, option := range options {
		option(&copy)
	}
	return &copy
}

func (c *ColorScheme) ToTokens() *token.Scheme {
	return &token.Scheme{
		Primary:             c.Primary.ToTokens(),
		PrimaryContainer:    c.PrimaryContainer.ToTokens(),
		PrimaryFixed:        c.PrimaryFixed.ToTokens(),
		PrimaryFixedVariant: c.PrimaryFixedVariant.ToTokens(),

		Secondary:             c.Secondary.ToTokens(),
		SecondaryContainer:    c.SecondaryContainer.ToTokens(),
		SecondaryFixed:        c.SecondaryFixed.ToTokens(),
		SecondaryFixedVariant: c.SecondaryFixedVariant.ToTokens(),

		Tertiary:             c.Tertiary.ToTokens(),
		TertiaryContainer:    c.TertiaryContainer.ToTokens(),
		TertiaryFixed:        c.TertiaryFixed.ToTokens(),
		TertiaryFixedVariant: c.TertiaryFixedVariant.ToTokens(),

		Surface:                 c.Surface.ToTokens(),
		SurfaceVariant:          c.SurfaceVariant.ToTokens(),
		SurfaceDim:              ColorToToken(c.SurfaceDim),
		SurfaceBright:           ColorToToken(c.SurfaceBright),
		SurfaceContainerLowest:  ColorToToken(c.SurfaceContainerLowest),
		SurfaceContainerLow:     ColorToToken(c.SurfaceContainerLow),
		SurfaceContainer:        ColorToToken(c.SurfaceContainer),
		SurfaceContainerHigh:    ColorToToken(c.SurfaceContainerHigh),
		SurfaceContainerHighest: ColorToToken(c.SurfaceContainerHighest),

		InverseSurface: c.InverseSurface.ToTokens(),
		InversePrimary: ColorToToken(c.InversePrimary),

		Background: c.Background.ToTokens(),

		Outline:        ColorToToken(c.Outline),
		OutlineVariant: ColorToToken(c.OutlineVariant),

		Error:          c.Error.ToTokens(),
		ErrorContainer: c.ErrorContainer.ToTokens(),

		Scrim:  ColorToToken(c.Scrim),
		Shadow: ColorToToken(c.Shadow),
	}
}

func ColorSchemeFromTokens(tokenScheme *token.Scheme) *ColorScheme {
	return &ColorScheme{
		Primary:             ColorSetFromTokens(tokenScheme.Primary),
		PrimaryContainer:    ColorSetFromTokens(tokenScheme.PrimaryContainer),
		PrimaryFixed:        ColorSetFromTokens(tokenScheme.PrimaryFixed),
		PrimaryFixedVariant: ColorSetFromTokens(tokenScheme.PrimaryFixedVariant),

		Secondary:             ColorSetFromTokens(tokenScheme.Secondary),
		SecondaryContainer:    ColorSetFromTokens(tokenScheme.SecondaryContainer),
		SecondaryFixed:        ColorSetFromTokens(tokenScheme.SecondaryFixed),
		SecondaryFixedVariant: ColorSetFromTokens(tokenScheme.SecondaryFixedVariant),

		Tertiary:             ColorSetFromTokens(tokenScheme.Tertiary),
		TertiaryContainer:    ColorSetFromTokens(tokenScheme.TertiaryContainer),
		TertiaryFixed:        ColorSetFromTokens(tokenScheme.TertiaryFixed),
		TertiaryFixedVariant: ColorSetFromTokens(tokenScheme.TertiaryFixedVariant),

		Surface:                 ColorSetFromTokens(tokenScheme.Surface),
		SurfaceVariant:          ColorSetFromTokens(tokenScheme.SurfaceVariant),
		SurfaceDim:              ColorFromTokens(tokenScheme.SurfaceDim),
		SurfaceBright:           ColorFromTokens(tokenScheme.SurfaceBright),
		SurfaceContainerLowest:  ColorFromTokens(tokenScheme.SurfaceContainerLowest),
		SurfaceContainerLow:     ColorFromTokens(tokenScheme.SurfaceContainerLow),
		SurfaceContainer:        ColorFromTokens(tokenScheme.SurfaceContainer),
		SurfaceContainerHigh:    ColorFromTokens(tokenScheme.SurfaceContainerHigh),
		SurfaceContainerHighest: ColorFromTokens(tokenScheme.SurfaceContainerHighest),

		InverseSurface: ColorSetFromTokens(tokenScheme.InverseSurface),
		InversePrimary: ColorFromTokens(tokenScheme.InversePrimary),

		Background: ColorSetFromTokens(tokenScheme.Background),

		Outline:        ColorFromTokens(tokenScheme.Outline),
		OutlineVariant: ColorFromTokens(tokenScheme.OutlineVariant),

		Error:          ColorSetFromTokens(tokenScheme.Error),
		ErrorContainer: ColorSetFromTokens(tokenScheme.ErrorContainer),

		Scrim:  ColorFromTokens(tokenScheme.Scrim),
		Shadow: ColorFromTokens(tokenScheme.Shadow),
	}
}

func ColorFromTokens(token token.MatColor) Color {
	return graphics.FromNRGBA(token.AsNRGBA())
}

func ColorToToken(color Color) token.MatColor {
	return token.MatColor(graphics.ColorToNRGBA(color))
}
