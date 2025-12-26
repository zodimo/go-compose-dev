package text

var PlatformTextStyleUnspecified *PlatformTextStyle = &PlatformTextStyle{
	SpanStyle:      PlatformSpanStyleUnspecified,
	ParagraphStyle: PlatformParagraphStyleUnspecified,
}
var PlatformSpanStyleUnspecified *PlatformSpanStyle = &PlatformSpanStyle{}
var PlatformParagraphStyleUnspecified *PlatformParagraphStyle = &PlatformParagraphStyle{}

// PlatformSpanStyle contains platform-specific text styling.
type PlatformSpanStyle struct {
}

// PlatformParagraphStyle contains platform-specific paragraph styling.
type PlatformParagraphStyle struct {
}

// PlatformTextStyle contains platform-specific text styling combining span and paragraph styles.
type PlatformTextStyle struct {
	SpanStyle      *PlatformSpanStyle
	ParagraphStyle *PlatformParagraphStyle
}

func IsSpecifiedPlatformTextStyle(s *PlatformTextStyle) bool {
	return s != nil && s != PlatformTextStyleUnspecified
}

func IsSpecifiedPlatformSpanStyle(s *PlatformSpanStyle) bool {
	return s != nil && s != PlatformSpanStyleUnspecified
}

func IsSpecifiedPlatformParagraphStyle(s *PlatformParagraphStyle) bool {
	return s != nil && s != PlatformParagraphStyleUnspecified
}

func TakeOrElsePlatformTextStyle(s, def *PlatformTextStyle) *PlatformTextStyle {
	if !IsSpecifiedPlatformTextStyle(s) {
		return def
	}
	return s
}

func TakeOrElsePlatformSpanStyle(s, def *PlatformSpanStyle) *PlatformSpanStyle {
	if !IsSpecifiedPlatformSpanStyle(s) {
		return def
	}
	return s
}

func TakeOrElsePlatformParagraphStyle(s, def *PlatformParagraphStyle) *PlatformParagraphStyle {
	if !IsSpecifiedPlatformParagraphStyle(s) {
		return def
	}
	return s
}

func EqualPlatformTextStyle(a, b *PlatformTextStyle) bool {
	panic("EqualPlatformTextStyle not implemented")
}

func EqualPlatformSpanStyle(a, b *PlatformSpanStyle) bool {
	panic("EqualPlatformSpanStyle not implemented")
}

func EqualPlatformParagraphStyle(a, b *PlatformParagraphStyle) bool {
	panic("EqualPlatformParagraphStyle not implemented")
}
