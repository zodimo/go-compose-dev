package style

// LineBreak configuration for line breaking.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-text/src/commonMain/kotlin/androidx/compose/ui/text/style/LineBreak.kt
type LineBreak int

const (
	LineBreakSimple LineBreak = iota
	LineBreakHeading
	LineBreakParagraph
	LineBreakUnspecified
)

var DefaultLineBreak = LineBreakSimple
