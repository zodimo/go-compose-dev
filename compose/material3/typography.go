package material3

import (
	"fmt"

	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/compose/material3/tokens"
	"github.com/zodimo/go-compose/compose/ui/text"
)

var TypographyUnspecified = &Typography{
	BodyLarge:  text.TextStyleUnspecified,
	BodyMedium: text.TextStyleUnspecified,
	BodySmall:  text.TextStyleUnspecified,

	DisplayLarge:  text.TextStyleUnspecified,
	DisplayMedium: text.TextStyleUnspecified,
	DisplaySmall:  text.TextStyleUnspecified,

	HeadlineLarge:  text.TextStyleUnspecified,
	HeadlineMedium: text.TextStyleUnspecified,
	HeadlineSmall:  text.TextStyleUnspecified,

	LabelLarge:  text.TextStyleUnspecified,
	LabelMedium: text.TextStyleUnspecified,
	LabelSmall:  text.TextStyleUnspecified,

	TitleLarge:  text.TextStyleUnspecified,
	TitleMedium: text.TextStyleUnspecified,
	TitleSmall:  text.TextStyleUnspecified,

	// Emphasized
	BodyLargeEmphasized:  text.TextStyleUnspecified,
	BodyMediumEmphasized: text.TextStyleUnspecified,
	BodySmallEmphasized:  text.TextStyleUnspecified,

	DisplayLargeEmphasized:  text.TextStyleUnspecified,
	DisplayMediumEmphasized: text.TextStyleUnspecified,
	DisplaySmallEmphasized:  text.TextStyleUnspecified,

	HeadlineLargeEmphasized:  text.TextStyleUnspecified,
	HeadlineMediumEmphasized: text.TextStyleUnspecified,
	HeadlineSmallEmphasized:  text.TextStyleUnspecified,

	LabelLargeEmphasized:  text.TextStyleUnspecified,
	LabelMediumEmphasized: text.TextStyleUnspecified,
	LabelSmallEmphasized:  text.TextStyleUnspecified,

	TitleLargeEmphasized:  text.TextStyleUnspecified,
	TitleMediumEmphasized: text.TextStyleUnspecified,
	TitleSmallEmphasized:  text.TextStyleUnspecified,
}

type Typography struct {
	BodyLarge  *text.TextStyle
	BodyMedium *text.TextStyle
	BodySmall  *text.TextStyle

	DisplayLarge  *text.TextStyle
	DisplayMedium *text.TextStyle
	DisplaySmall  *text.TextStyle

	HeadlineLarge  *text.TextStyle
	HeadlineMedium *text.TextStyle
	HeadlineSmall  *text.TextStyle

	LabelLarge  *text.TextStyle
	LabelMedium *text.TextStyle
	LabelSmall  *text.TextStyle

	TitleLarge  *text.TextStyle
	TitleMedium *text.TextStyle
	TitleSmall  *text.TextStyle

	// Emphasized
	BodyLargeEmphasized  *text.TextStyle
	BodyMediumEmphasized *text.TextStyle
	BodySmallEmphasized  *text.TextStyle

	DisplayLargeEmphasized  *text.TextStyle
	DisplayMediumEmphasized *text.TextStyle
	DisplaySmallEmphasized  *text.TextStyle

	HeadlineLargeEmphasized  *text.TextStyle
	HeadlineMediumEmphasized *text.TextStyle
	HeadlineSmallEmphasized  *text.TextStyle

	LabelLargeEmphasized  *text.TextStyle
	LabelMediumEmphasized *text.TextStyle
	LabelSmallEmphasized  *text.TextStyle

	TitleLargeEmphasized  *text.TextStyle
	TitleMediumEmphasized *text.TextStyle
	TitleSmallEmphasized  *text.TextStyle
}

func (t *Typography) FromToken(value tokens.TypographyTokenKey) *text.TextStyle {
	switch value {
	case tokens.TypographyTokenKeyUnspecified:
		return text.TextStyleUnspecified
	case tokens.TypographyTokenKeyBodyLarge:
		return t.BodyLarge
	case tokens.TypographyTokenKeyBodyMedium:
		return t.BodyMedium
	case tokens.TypographyTokenKeyBodySmall:
		return t.BodySmall
	case tokens.TypographyTokenKeyDisplayLarge:
		return t.DisplayLarge
	case tokens.TypographyTokenKeyDisplayMedium:
		return t.DisplayMedium
	case tokens.TypographyTokenKeyDisplaySmall:
		return t.DisplaySmall
	case tokens.TypographyTokenKeyHeadlineLarge:
		return t.HeadlineLarge
	case tokens.TypographyTokenKeyHeadlineMedium:
		return t.HeadlineMedium
	case tokens.TypographyTokenKeyHeadlineSmall:
		return t.HeadlineSmall
	case tokens.TypographyTokenKeyLabelLarge:
		return t.LabelLarge
	case tokens.TypographyTokenKeyLabelMedium:
		return t.LabelMedium
	case tokens.TypographyTokenKeyLabelSmall:
		return t.LabelSmall
	case tokens.TypographyTokenKeyTitleLarge:
		return t.TitleLarge
	case tokens.TypographyTokenKeyTitleMedium:
		return t.TitleMedium
	case tokens.TypographyTokenKeyTitleSmall:
		return t.TitleSmall
	case tokens.TypographyTokenKeyBodyLargeEmphasized:
		return t.BodyLargeEmphasized
	case tokens.TypographyTokenKeyBodyMediumEmphasized:
		return t.BodyMediumEmphasized
	case tokens.TypographyTokenKeyBodySmallEmphasized:
		return t.BodySmallEmphasized
	case tokens.TypographyTokenKeyDisplayLargeEmphasized:
		return t.DisplayLargeEmphasized
	case tokens.TypographyTokenKeyDisplayMediumEmphasized:
		return t.DisplayMediumEmphasized
	case tokens.TypographyTokenKeyDisplaySmallEmphasized:
		return t.DisplaySmallEmphasized
	case tokens.TypographyTokenKeyHeadlineLargeEmphasized:
		return t.HeadlineLargeEmphasized
	case tokens.TypographyTokenKeyHeadlineMediumEmphasized:
		return t.HeadlineMediumEmphasized
	case tokens.TypographyTokenKeyHeadlineSmallEmphasized:
		return t.HeadlineSmallEmphasized
	case tokens.TypographyTokenKeyLabelLargeEmphasized:
		return t.LabelLargeEmphasized
	case tokens.TypographyTokenKeyLabelMediumEmphasized:
		return t.LabelMediumEmphasized
	case tokens.TypographyTokenKeyLabelSmallEmphasized:
		return t.LabelSmallEmphasized
	case tokens.TypographyTokenKeyTitleLargeEmphasized:
		return t.TitleLargeEmphasized
	case tokens.TypographyTokenKeyTitleMediumEmphasized:
		return t.TitleMediumEmphasized
	case tokens.TypographyTokenKeyTitleSmallEmphasized:
		return t.TitleSmallEmphasized
	default:
		panic(fmt.Sprintf("unknown typography token: %s", value))
	}
}

var LocalTypography = compose.CompositionLocalOf(func() *Typography {
	return DefaultTypography
})
