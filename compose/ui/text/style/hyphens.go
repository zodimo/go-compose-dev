package style

// Hyphens configuration.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-text/src/commonMain/kotlin/androidx/compose/ui/text/style/Hyphens.kt
type Hyphens int

const (
	HyphensNone Hyphens = iota
	HyphensAuto
	HyphensUnspecified
)

var DefaultHyphens = HyphensNone
