package input

// TextHighlightType defines highlight styles for IME handwriting gestures.
// Used to visually preview gestures before they're committed.
//
// This is a port of androidx.compose.foundation.text.input.TextHighlightType.
type TextHighlightType int

const (
	// TextHighlightTypeHandwritingSelectPreview previews text that would be selected
	// by an ongoing stylus handwriting select gesture.
	TextHighlightTypeHandwritingSelectPreview TextHighlightType = iota

	// TextHighlightTypeHandwritingDeletePreview previews text that would be deleted
	// by an ongoing stylus handwriting delete gesture.
	TextHighlightTypeHandwritingDeletePreview
)

// String returns a string representation of the TextHighlightType.
func (t TextHighlightType) String() string {
	switch t {
	case TextHighlightTypeHandwritingSelectPreview:
		return "HandwritingSelectPreview"
	case TextHighlightTypeHandwritingDeletePreview:
		return "HandwritingDeletePreview"
	default:
		return "Unknown"
	}
}
