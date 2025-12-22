package style

// LineHeightStyle configuration for line height.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-text/src/commonMain/kotlin/androidx/compose/ui/text/style/LineHeightStyle.kt
type LineHeightStyle struct {
	Alignment LineHeightStyleAlignment
	Trim      LineHeightStyleTrim
}

type LineHeightStyleAlignment int

const (
	LineHeightStyleAlignmentTop LineHeightStyleAlignment = iota
	LineHeightStyleAlignmentCenter
	LineHeightStyleAlignmentProportional
	LineHeightStyleAlignmentBottom
)

type LineHeightStyleTrim int

const (
	LineHeightStyleTrimFirstLineTop   LineHeightStyleTrim = 1 << 0
	LineHeightStyleTrimLastLineBottom LineHeightStyleTrim = 1 << 1
	LineHeightStyleTrimBoth           LineHeightStyleTrim = LineHeightStyleTrimFirstLineTop | LineHeightStyleTrimLastLineBottom
	LineHeightStyleTrimNone           LineHeightStyleTrim = 0
)

var DefaultLineHeightStyle = LineHeightStyle{
	Alignment: LineHeightStyleAlignmentProportional,
	Trim:      LineHeightStyleTrimBoth,
}

func (l LineHeightStyle) Equals(other LineHeightStyle) bool {
	return l.Alignment == other.Alignment && l.Trim == other.Trim
}
