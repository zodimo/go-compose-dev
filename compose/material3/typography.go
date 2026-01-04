package material3

import "github.com/zodimo/go-compose/compose/ui/text"

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
