package selection

import (
	"github.com/zodimo/go-compose/compose"
)

// DisableSelection disables text selection for its direct or indirect children.
// To use this, simply add this to wrap one or more text composables.
//
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/foundation/foundation/src/commonMain/kotlin/androidx/compose/foundation/text/selection/SelectionContainer.kt
func DisableSelection(content Composable) Composable {
	return compose.CompositionLocalProvider1(
		LocalSelectionRegistrar,
		nil,
		content,
	)
}
