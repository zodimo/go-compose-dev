package material3

import (
	"github.com/zodimo/go-compose/compose/ui/text"
	"github.com/zodimo/go-compose/compose/ui/text/font"
	"github.com/zodimo/go-compose/compose/ui/unit"
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
		DisplayLarge: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightNormal), // 400
			text.WithFontSize(unit.Sp(57)),
			text.WithLineHeight(unit.Sp(64)),
			text.WithLetterSpacing(unit.Sp(-0.25)),
		),
		DisplayMedium: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightNormal), // 400
			text.WithFontSize(unit.Sp(45)),
			text.WithLineHeight(unit.Sp(52)),
			text.WithLetterSpacing(unit.Sp(0)),
		),
		DisplaySmall: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightNormal), // 400
			text.WithFontSize(unit.Sp(36)),
			text.WithLineHeight(unit.Sp(44)),
			text.WithLetterSpacing(unit.Sp(0)),
		),

		HeadlineLarge: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightNormal), // 400
			text.WithFontSize(unit.Sp(32)),
			text.WithLineHeight(unit.Sp(40)),
			text.WithLetterSpacing(unit.Sp(0)),
		),
		HeadlineMedium: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightNormal), // 400
			text.WithFontSize(unit.Sp(28)),
			text.WithLineHeight(unit.Sp(36)),
			text.WithLetterSpacing(unit.Sp(0)),
		),
		HeadlineSmall: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightNormal), // 400
			text.WithFontSize(unit.Sp(24)),
			text.WithLineHeight(unit.Sp(32)),
			text.WithLetterSpacing(unit.Sp(0)),
		),

		TitleLarge: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightNormal), // 400
			text.WithFontSize(unit.Sp(22)),
			text.WithLineHeight(unit.Sp(28)),
			text.WithLetterSpacing(unit.Sp(0)),
		),
		TitleMedium: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightMedium), // 500
			text.WithFontSize(unit.Sp(16)),
			text.WithLineHeight(unit.Sp(24)),
			text.WithLetterSpacing(unit.Sp(0.15)),
		),
		TitleSmall: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightMedium), // 500
			text.WithFontSize(unit.Sp(14)),
			text.WithLineHeight(unit.Sp(20)),
			text.WithLetterSpacing(unit.Sp(0.1)),
		),

		LabelLarge: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightMedium), // 500
			text.WithFontSize(unit.Sp(14)),
			text.WithLineHeight(unit.Sp(20)),
			text.WithLetterSpacing(unit.Sp(0.1)),
		),
		LabelMedium: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightMedium), // 500
			text.WithFontSize(unit.Sp(12)),
			text.WithLineHeight(unit.Sp(16)),
			text.WithLetterSpacing(unit.Sp(0.5)),
		),
		LabelSmall: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightMedium), // 500
			text.WithFontSize(unit.Sp(11)),
			text.WithLineHeight(unit.Sp(16)),
			text.WithLetterSpacing(unit.Sp(0.5)),
		),

		BodyLarge: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightNormal), // 400
			text.WithFontSize(unit.Sp(16)),
			text.WithLineHeight(unit.Sp(24)),
			text.WithLetterSpacing(unit.Sp(0.5)),
		),
		BodyMedium: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightNormal), // 400
			text.WithFontSize(unit.Sp(14)),
			text.WithLineHeight(unit.Sp(20)),
			text.WithLetterSpacing(unit.Sp(0.25)),
		),
		BodySmall: text.TextStyleFromOptions(
			text.WithFontFamily(defaultFontFamily),
			text.WithFontWeight(font.FontWeightNormal), // 400
			text.WithFontSize(unit.Sp(12)),
			text.WithLineHeight(unit.Sp(16)),
			text.WithLetterSpacing(unit.Sp(0.4)),
		),
	}
}
