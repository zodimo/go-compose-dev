package material3

import (
	"github.com/zodimo/go-compose/compose/material3/tokens"
	"github.com/zodimo/go-compose/compose/ui/text/font"
)

var (
	// Default Font Family (System Default)
	// Should this exist?
	DefaultFontFamily = font.FontFamilySansSerif
)

// DefaultTypography returns the standard Material 3 type scale.
var DefaultTypography = &Typography{
	DisplayLarge:   tokens.TypographyTokens.DisplayLarge,
	DisplayMedium:  tokens.TypographyTokens.DisplayMedium,
	DisplaySmall:   tokens.TypographyTokens.DisplaySmall,
	HeadlineLarge:  tokens.TypographyTokens.HeadlineLarge,
	HeadlineMedium: tokens.TypographyTokens.HeadlineMedium,
	HeadlineSmall:  tokens.TypographyTokens.HeadlineSmall,
	TitleLarge:     tokens.TypographyTokens.TitleLarge,
	TitleMedium:    tokens.TypographyTokens.TitleMedium,
	TitleSmall:     tokens.TypographyTokens.TitleSmall,
	BodyLarge:      tokens.TypographyTokens.BodyLarge,
	BodyMedium:     tokens.TypographyTokens.BodyMedium,
	BodySmall:      tokens.TypographyTokens.BodySmall,
	LabelLarge:     tokens.TypographyTokens.LabelLarge,
	LabelMedium:    tokens.TypographyTokens.LabelMedium,
	LabelSmall:     tokens.TypographyTokens.LabelSmall,

	DisplayLargeEmphasized:   tokens.TypographyTokens.DisplayLargeEmphasized,
	DisplayMediumEmphasized:  tokens.TypographyTokens.DisplayMediumEmphasized,
	DisplaySmallEmphasized:   tokens.TypographyTokens.DisplaySmallEmphasized,
	HeadlineLargeEmphasized:  tokens.TypographyTokens.HeadlineLargeEmphasized,
	HeadlineMediumEmphasized: tokens.TypographyTokens.HeadlineMediumEmphasized,
	HeadlineSmallEmphasized:  tokens.TypographyTokens.HeadlineSmallEmphasized,
	TitleLargeEmphasized:     tokens.TypographyTokens.TitleLargeEmphasized,
	TitleMediumEmphasized:    tokens.TypographyTokens.TitleMediumEmphasized,
	TitleSmallEmphasized:     tokens.TypographyTokens.TitleSmallEmphasized,
	BodyLargeEmphasized:      tokens.TypographyTokens.BodyLargeEmphasized,
	BodyMediumEmphasized:     tokens.TypographyTokens.BodyMediumEmphasized,
	BodySmallEmphasized:      tokens.TypographyTokens.BodySmallEmphasized,
	LabelLargeEmphasized:     tokens.TypographyTokens.LabelLargeEmphasized,
	LabelMediumEmphasized:    tokens.TypographyTokens.LabelMediumEmphasized,
	LabelSmallEmphasized:     tokens.TypographyTokens.LabelSmallEmphasized,
}
