package input

import "math"

// TextFieldLineLimits specifies text wrapping and height behavior.
//
// This is a port of androidx.compose.foundation.text.input.TextFieldLineLimits.
type TextFieldLineLimits interface {
	isTextFieldLineLimits()
}

// SingleLine makes the field always single-line, ignoring newlines.
// The field scrolls horizontally when text overflows.
type SingleLine struct{}

func (SingleLine) isTextFieldLineLimits() {}

// IsSingleLine returns true if the line limits enforce single-line behavior.
func IsSingleLine(limits TextFieldLineLimits) bool {
	_, ok := limits.(SingleLine)
	return ok
}

// MultiLine allows multiple lines with configurable height limits.
// The field grows from MinHeightInLines to MaxHeightInLines, then scrolls.
type MultiLine struct {
	// MinHeightInLines is the minimum height in lines. Default is 1.
	MinHeightInLines int

	// MaxHeightInLines is the maximum height in lines. Default is MaxInt.
	MaxHeightInLines int
}

func (MultiLine) isTextFieldLineLimits() {}

// NewMultiLine creates a MultiLine with default values (1 to MaxInt lines).
func NewMultiLine() MultiLine {
	return MultiLine{
		MinHeightInLines: 1,
		MaxHeightInLines: math.MaxInt,
	}
}

// NewMultiLineWithLimits creates a MultiLine with specified limits.
// Panics if minLines < 1 or maxLines < minLines.
func NewMultiLineWithLimits(minLines, maxLines int) MultiLine {
	if minLines < 1 {
		panic("minLines must be at least 1")
	}
	if maxLines < minLines {
		panic("maxLines must be >= minLines")
	}
	return MultiLine{
		MinHeightInLines: minLines,
		MaxHeightInLines: maxLines,
	}
}

// TextFieldLineLimitsDefault is the default line limits (multi-line with no restrictions).
var TextFieldLineLimitsDefault TextFieldLineLimits = NewMultiLine()

// TextFieldLineLimitsSingleLine is a convenience constant for single-line fields.
var TextFieldLineLimitsSingleLine TextFieldLineLimits = SingleLine{}

// ShouldWrap returns whether the text field should wrap text.
func ShouldWrap(limits TextFieldLineLimits) bool {
	// Single line should not wrap
	return !IsSingleLine(limits)
}
