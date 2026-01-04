package material3

import (
	"github.com/zodimo/go-compose/compose/material3/tokens"
	"github.com/zodimo/go-compose/compose/ui/text"
	"github.com/zodimo/go-compose/compose/ui/text/font"
)

var (
	// Default Font Family (System Default)
	defaultFontFamily = font.FontFamilySansSerif
)

var TypographyUnspecified = &Typography{
	DisplaySmall:  text.TextStyleUnspecified,
	DisplayMedium: text.TextStyleUnspecified,
	DisplayLarge:  text.TextStyleUnspecified,

	HeadlineSmall:  text.TextStyleUnspecified,
	HeadlineMedium: text.TextStyleUnspecified,
	HeadlineLarge:  text.TextStyleUnspecified,

	TitleSmall:  text.TextStyleUnspecified,
	TitleMedium: text.TextStyleUnspecified,
	TitleLarge:  text.TextStyleUnspecified,

	BodySmall:  text.TextStyleUnspecified,
	BodyMedium: text.TextStyleUnspecified,
	BodyLarge:  text.TextStyleUnspecified,

	LabelSmall:  text.TextStyleUnspecified,
	LabelMedium: text.TextStyleUnspecified,
	LabelLarge:  text.TextStyleUnspecified,

	DisplaySmallEmphasized:  text.TextStyleUnspecified,
	DisplayMediumEmphasized: text.TextStyleUnspecified,
	DisplayLargeEmphasized:  text.TextStyleUnspecified,

	HeadlineSmallEmphasized:  text.TextStyleUnspecified,
	HeadlineMediumEmphasized: text.TextStyleUnspecified,
	HeadlineLargeEmphasized:  text.TextStyleUnspecified,

	TitleSmallEmphasized:  text.TextStyleUnspecified,
	TitleMediumEmphasized: text.TextStyleUnspecified,
	TitleLargeEmphasized:  text.TextStyleUnspecified,

	BodySmallEmphasized:  text.TextStyleUnspecified,
	BodyMediumEmphasized: text.TextStyleUnspecified,
	BodyLargeEmphasized:  text.TextStyleUnspecified,

	LabelSmallEmphasized:  text.TextStyleUnspecified,
	LabelMediumEmphasized: text.TextStyleUnspecified,
	LabelLargeEmphasized:  text.TextStyleUnspecified,
}

// DefaultTypography returns the standard Material 3 type scale.
func DefaultTypography() *Typography {
	return &Typography{
		DisplayLarge:   tokens.DisplayLarge,
		DisplayMedium:  tokens.DisplayMedium,
		DisplaySmall:   tokens.DisplaySmall,
		HeadlineLarge:  tokens.HeadlineLarge,
		HeadlineMedium: tokens.HeadlineMedium,
		HeadlineSmall:  tokens.HeadlineSmall,
		TitleLarge:     tokens.TitleLarge,
		TitleMedium:    tokens.TitleMedium,
		TitleSmall:     tokens.TitleSmall,
		BodyLarge:      tokens.BodyLarge,
		BodyMedium:     tokens.BodyMedium,
		BodySmall:      tokens.BodySmall,
		LabelLarge:     tokens.LabelLarge,
		LabelMedium:    tokens.LabelMedium,
		LabelSmall:     tokens.LabelSmall,

		DisplayLargeEmphasized:   tokens.DisplayLargeEmphasized,
		DisplayMediumEmphasized:  tokens.DisplayMediumEmphasized,
		DisplaySmallEmphasized:   tokens.DisplaySmallEmphasized,
		HeadlineLargeEmphasized:  tokens.HeadlineLargeEmphasized,
		HeadlineMediumEmphasized: tokens.HeadlineMediumEmphasized,
		HeadlineSmallEmphasized:  tokens.HeadlineSmallEmphasized,
		TitleLargeEmphasized:     tokens.TitleLargeEmphasized,
		TitleMediumEmphasized:    tokens.TitleMediumEmphasized,
		TitleSmallEmphasized:     tokens.TitleSmallEmphasized,
		BodyLargeEmphasized:      tokens.BodyLargeEmphasized,
		BodyMediumEmphasized:     tokens.BodyMediumEmphasized,
		BodySmallEmphasized:      tokens.BodySmallEmphasized,
		LabelLargeEmphasized:     tokens.LabelLargeEmphasized,
		LabelMediumEmphasized:    tokens.LabelMediumEmphasized,
		LabelSmallEmphasized:     tokens.LabelSmallEmphasized,
	}
}
