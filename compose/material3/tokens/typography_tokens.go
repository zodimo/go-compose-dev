package tokens

import (
	"github.com/zodimo/go-compose/compose/ui/text"
)

var (
	BodyLarge = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleBodyLargeFont),
		text.WithFontWeight(TypeScaleBodyLargeWeight),
		text.WithFontSize(TypeScaleBodyLargeSize),
		text.WithLineHeight(TypeScaleBodyLargeLineHeight),
		text.WithLetterSpacing(TypeScaleBodyLargeTracking),
	)
	BodyMedium = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleBodyMediumFont),
		text.WithFontWeight(TypeScaleBodyMediumWeight),
		text.WithFontSize(TypeScaleBodyMediumSize),
		text.WithLineHeight(TypeScaleBodyMediumLineHeight),
		text.WithLetterSpacing(TypeScaleBodyMediumTracking),
	)
	BodySmall = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleBodySmallFont),
		text.WithFontWeight(TypeScaleBodySmallWeight),
		text.WithFontSize(TypeScaleBodySmallSize),
		text.WithLineHeight(TypeScaleBodySmallLineHeight),
		text.WithLetterSpacing(TypeScaleBodySmallTracking),
	)
	DisplayLarge = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleDisplayLargeFont),
		text.WithFontWeight(TypeScaleDisplayLargeWeight),
		text.WithFontSize(TypeScaleDisplayLargeSize),
		text.WithLineHeight(TypeScaleDisplayLargeLineHeight),
		text.WithLetterSpacing(TypeScaleDisplayLargeTracking),
	)
	DisplayMedium = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleDisplayMediumFont),
		text.WithFontWeight(TypeScaleDisplayMediumWeight),
		text.WithFontSize(TypeScaleDisplayMediumSize),
		text.WithLineHeight(TypeScaleDisplayMediumLineHeight),
		text.WithLetterSpacing(TypeScaleDisplayMediumTracking),
	)
	DisplaySmall = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleDisplaySmallFont),
		text.WithFontWeight(TypeScaleDisplaySmallWeight),
		text.WithFontSize(TypeScaleDisplaySmallSize),
		text.WithLineHeight(TypeScaleDisplaySmallLineHeight),
		text.WithLetterSpacing(TypeScaleDisplaySmallTracking),
	)
	HeadlineLarge = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleHeadlineLargeFont),
		text.WithFontWeight(TypeScaleHeadlineLargeWeight),
		text.WithFontSize(TypeScaleHeadlineLargeSize),
		text.WithLineHeight(TypeScaleHeadlineLargeLineHeight),
		text.WithLetterSpacing(TypeScaleHeadlineLargeTracking),
	)
	HeadlineMedium = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleHeadlineMediumFont),
		text.WithFontWeight(TypeScaleHeadlineMediumWeight),
		text.WithFontSize(TypeScaleHeadlineMediumSize),
		text.WithLineHeight(TypeScaleHeadlineMediumLineHeight),
		text.WithLetterSpacing(TypeScaleHeadlineMediumTracking),
	)
	HeadlineSmall = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleHeadlineSmallFont),
		text.WithFontWeight(TypeScaleHeadlineSmallWeight),
		text.WithFontSize(TypeScaleHeadlineSmallSize),
		text.WithLineHeight(TypeScaleHeadlineSmallLineHeight),
		text.WithLetterSpacing(TypeScaleHeadlineSmallTracking),
	)
	LabelLarge = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleLabelLargeFont),
		text.WithFontWeight(TypeScaleLabelLargeWeight),
		text.WithFontSize(TypeScaleLabelLargeSize),
		text.WithLineHeight(TypeScaleLabelLargeLineHeight),
		text.WithLetterSpacing(TypeScaleLabelLargeTracking),
	)
	LabelMedium = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleLabelMediumFont),
		text.WithFontWeight(TypeScaleLabelMediumWeight),
		text.WithFontSize(TypeScaleLabelMediumSize),
		text.WithLineHeight(TypeScaleLabelMediumLineHeight),
		text.WithLetterSpacing(TypeScaleLabelMediumTracking),
	)
	LabelSmall = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleLabelSmallFont),
		text.WithFontWeight(TypeScaleLabelSmallWeight),
		text.WithFontSize(TypeScaleLabelSmallSize),
		text.WithLineHeight(TypeScaleLabelSmallLineHeight),
		text.WithLetterSpacing(TypeScaleLabelSmallTracking),
	)
	TitleLarge = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleTitleLargeFont),
		text.WithFontWeight(TypeScaleTitleLargeWeight),
		text.WithFontSize(TypeScaleTitleLargeSize),
		text.WithLineHeight(TypeScaleTitleLargeLineHeight),
		text.WithLetterSpacing(TypeScaleTitleLargeTracking),
	)
	TitleMedium = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleTitleMediumFont),
		text.WithFontWeight(TypeScaleTitleMediumWeight),
		text.WithFontSize(TypeScaleTitleMediumSize),
		text.WithLineHeight(TypeScaleTitleMediumLineHeight),
		text.WithLetterSpacing(TypeScaleTitleMediumTracking),
	)
	TitleSmall = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleTitleSmallFont),
		text.WithFontWeight(TypeScaleTitleSmallWeight),
		text.WithFontSize(TypeScaleTitleSmallSize),
		text.WithLineHeight(TypeScaleTitleSmallLineHeight),
		text.WithLetterSpacing(TypeScaleTitleSmallTracking),
	)
	// Emphasized
	BodyLargeEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleBodyLargeEmphasizedFont),
		text.WithFontWeight(TypeScaleBodyLargeEmphasizedWeight),
		text.WithFontSize(TypeScaleBodyLargeEmphasizedSize),
		text.WithLineHeight(TypeScaleBodyLargeEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleBodyLargeEmphasizedTracking),
	)
	BodyMediumEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleBodyMediumEmphasizedFont),
		text.WithFontWeight(TypeScaleBodyMediumEmphasizedWeight),
		text.WithFontSize(TypeScaleBodyMediumEmphasizedSize),
		text.WithLineHeight(TypeScaleBodyMediumEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleBodyMediumEmphasizedTracking),
	)
	BodySmallEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleBodySmallEmphasizedFont),
		text.WithFontWeight(TypeScaleBodySmallEmphasizedWeight),
		text.WithFontSize(TypeScaleBodySmallEmphasizedSize),
		text.WithLineHeight(TypeScaleBodySmallEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleBodySmallEmphasizedTracking),
	)
	DisplayLargeEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleDisplayLargeEmphasizedFont),
		text.WithFontWeight(TypeScaleDisplayLargeEmphasizedWeight),
		text.WithFontSize(TypeScaleDisplayLargeEmphasizedSize),
		text.WithLineHeight(TypeScaleDisplayLargeEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleDisplayLargeEmphasizedTracking),
	)
	DisplayMediumEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleDisplayMediumEmphasizedFont),
		text.WithFontWeight(TypeScaleDisplayMediumEmphasizedWeight),
		text.WithFontSize(TypeScaleDisplayMediumEmphasizedSize),
		text.WithLineHeight(TypeScaleDisplayMediumEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleDisplayMediumEmphasizedTracking),
	)
	DisplaySmallEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleDisplaySmallEmphasizedFont),
		text.WithFontWeight(TypeScaleDisplaySmallEmphasizedWeight),
		text.WithFontSize(TypeScaleDisplaySmallEmphasizedSize),
		text.WithLineHeight(TypeScaleDisplaySmallEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleDisplaySmallEmphasizedTracking),
	)
	HeadlineLargeEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleHeadlineLargeEmphasizedFont),
		text.WithFontWeight(TypeScaleHeadlineLargeEmphasizedWeight),
		text.WithFontSize(TypeScaleHeadlineLargeEmphasizedSize),
		text.WithLineHeight(TypeScaleHeadlineLargeEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleHeadlineLargeEmphasizedTracking),
	)
	HeadlineMediumEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleHeadlineMediumEmphasizedFont),
		text.WithFontWeight(TypeScaleHeadlineMediumEmphasizedWeight),
		text.WithFontSize(TypeScaleHeadlineMediumEmphasizedSize),
		text.WithLineHeight(TypeScaleHeadlineMediumEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleHeadlineMediumEmphasizedTracking),
	)
	HeadlineSmallEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleHeadlineSmallEmphasizedFont),
		text.WithFontWeight(TypeScaleHeadlineSmallEmphasizedWeight),
		text.WithFontSize(TypeScaleHeadlineSmallEmphasizedSize),
		text.WithLineHeight(TypeScaleHeadlineSmallEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleHeadlineSmallEmphasizedTracking),
	)
	LabelLargeEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleLabelLargeEmphasizedFont),
		text.WithFontWeight(TypeScaleLabelLargeEmphasizedWeight),
		text.WithFontSize(TypeScaleLabelLargeEmphasizedSize),
		text.WithLineHeight(TypeScaleLabelLargeEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleLabelLargeEmphasizedTracking),
	)
	LabelMediumEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleLabelMediumEmphasizedFont),
		text.WithFontWeight(TypeScaleLabelMediumEmphasizedWeight),
		text.WithFontSize(TypeScaleLabelMediumEmphasizedSize),
		text.WithLineHeight(TypeScaleLabelMediumEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleLabelMediumEmphasizedTracking),
	)
	LabelSmallEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleLabelSmallEmphasizedFont),
		text.WithFontWeight(TypeScaleLabelSmallEmphasizedWeight),
		text.WithFontSize(TypeScaleLabelSmallEmphasizedSize),
		text.WithLineHeight(TypeScaleLabelSmallEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleLabelSmallEmphasizedTracking),
	)
	TitleLargeEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleTitleLargeEmphasizedFont),
		text.WithFontWeight(TypeScaleTitleLargeEmphasizedWeight),
		text.WithFontSize(TypeScaleTitleLargeEmphasizedSize),
		text.WithLineHeight(TypeScaleTitleLargeEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleTitleLargeEmphasizedTracking),
	)
	TitleMediumEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleTitleMediumEmphasizedFont),
		text.WithFontWeight(TypeScaleTitleMediumEmphasizedWeight),
		text.WithFontSize(TypeScaleTitleMediumEmphasizedSize),
		text.WithLineHeight(TypeScaleTitleMediumEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleTitleMediumEmphasizedTracking),
	)
	TitleSmallEmphasized = text.TextStyleFromOptions(
		text.WithFontFamily(TypeScaleTitleSmallEmphasizedFont),
		text.WithFontWeight(TypeScaleTitleSmallEmphasizedWeight),
		text.WithFontSize(TypeScaleTitleSmallEmphasizedSize),
		text.WithLineHeight(TypeScaleTitleSmallEmphasizedLineHeight),
		text.WithLetterSpacing(TypeScaleTitleSmallEmphasizedTracking),
	)
)
