package text

/** The annotation tag used by inline content. */
const INLINE_CONTENT_TAG = "compose.foundation.text.inlineContent"

// A string that contains a replacement character specified by unicode. It's used as the default
// value of alternate text.
const REPLACEMENT_CHAR = "\uFFFD"

// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/foundation/foundation/src/commonMain/kotlin/androidx/compose/foundation/text/InlineTextContent.kt;drc=4970f6e96cdb06089723da0ab8ec93ae3f067c7a;l=72
type InlineTextContent struct {
	/**
	 * The setting object that defines the size and vertical alignment of this composable in the
	 * text line. This is different from the measure of Layout
	 *
	 * @see Placeholder
	 */
	PlaceHolder uiPlaceholder
	/**
	 * The composable to be inserted into the text layout. The string parameter passed to it will
	 * the alternateText given to [appendInlineContent].
	 */
	Children []func(string) Composable
}
