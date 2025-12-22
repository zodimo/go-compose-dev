package style

// TextMotion configuration.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-text/src/commonMain/kotlin/androidx/compose/ui/text/style/TextMotion.kt
type TextMotion int

const (
	TextMotionStatic TextMotion = iota
	TextMotionAnimated
)
