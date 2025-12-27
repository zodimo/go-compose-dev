package text

import "gioui.org/text"

// Re-Export GIO Enums

type Alignment = text.Alignment

const (
	Start  Alignment = text.Start
	End    Alignment = text.End
	Middle Alignment = text.Middle
)

// WrapPolicy configures strategies for choosing where to break lines of text for line
// wrapping.
type WrapPolicy = text.WrapPolicy

const (
	// WrapHeuristically tries to minimize breaking within words (UAX#14 text segments)
	// while also ensuring that text fits within the given MaxWidth. It will only break
	// a line within a word (on a UAX#29 grapheme cluster boundary) when that word cannot
	// fit on a line by itself. Additionally, when the final word of a line is being
	// truncated, this policy will preserve as many symbols of that word as
	// possible before the truncator.
	WrapHeuristically WrapPolicy = text.WrapHeuristically
	// WrapWords does not permit words (UAX#14 text segments) to be broken across lines.
	// This means that sometimes long words will exceed the MaxWidth they are wrapped with.
	WrapWords WrapPolicy = text.WrapWords
	// WrapGraphemes will maximize the amount of text on each line at the expense of readability,
	// breaking any word across lines on UAX#29 grapheme cluster boundaries to maximize the number of
	// grapheme clusters on each line.
	WrapGraphemes WrapPolicy = text.WrapGraphemes
)
