package text

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"
)

type TextOptions struct {
	Modifier Modifier

	// Alignment specifies the text alignment.
	Alignment Alignment
	// MaxLines limits the number of lines. Zero means no limit.
	MaxLines int
	// Truncator is the text that will be shown at the end of the final
	// line if MaxLines is exceeded. Defaults to "…" if empty.
	Truncator string
	// WrapPolicy configures how displayed text will be broken into lines.
	WrapPolicy WrapPolicy
	// LineHeight controls the distance between the baselines of lines of text.
	// If zero, a sensible default will be used.
	LineHeight unit.Sp
	// LineHeightScale applies a scaling factor to the LineHeight. If zero, a
	// sensible default will be used.
	LineHeightScale float32

	// Shaper is the text shaper used to display this labe. This field is automatically
	// set using by all constructor functions. If constructing a LabelStyle literal, you
	// must provide a Shaper or displaying text will panic.
	Shaper *text.Shaper

	// Selectable provides text selection state for the label. If not set, the label cannot
	// be selected or copied interactively.
	Selectable bool

	TextStyleOptions *TextStyleOptions
}

type TextOption func(*TextOptions)

func TextWithModifier(modifier Modifier) TextOption {
	return func(o *TextOptions) {
		o.Modifier = o.Modifier.Then(modifier)
	}
}

func TextWithAlignment(alignment Alignment) TextOption {
	return func(o *TextOptions) {
		o.Alignment = alignment
	}
}

func TextWithMaxLines(maxLines int) TextOption {
	return func(o *TextOptions) {
		o.MaxLines = maxLines
	}
}

func TextWithTruncator(truncator string) TextOption {
	return func(o *TextOptions) {
		o.Truncator = truncator
	}
}

func TextWithWrapPolicy(wrapPolicy WrapPolicy) TextOption {

	return func(o *TextOptions) {
		o.WrapPolicy = wrapPolicy
	}
}

func TextWithLineHeight(lineHeightInSP float32) TextOption {

	return func(o *TextOptions) {
		o.LineHeight = unit.Sp(lineHeightInSP)
	}
}

func TextWithLineHeightScale(lineHeightScale float32) TextOption {

	return func(o *TextOptions) {
		o.LineHeightScale = lineHeightScale
	}
}

func TextWithShaper(shaper *text.Shaper) TextOption {

	return func(o *TextOptions) {
		o.Shaper = shaper
	}
}

func TextWithTextStyleOptions(textStyleOptions ...TextStyleOption) TextOption {
	return func(o *TextOptions) {
		for _, textStyleOption := range textStyleOptions {
			textStyleOption(o.TextStyleOptions)
		}
	}
}

// TextSelectable sets the text to be selectable.
// stores *widget.Selectable in runtime memoization
func TextSelectable() TextOption {
	return func(o *TextOptions) {
		o.Selectable = true
	}
}

type TextStyleOptions struct {
	// Face defines the text style.
	Font font.Font
	// Color is the text color.
	Color color.NRGBA
	// SelectionColor is the color of the background for selected text.
	SelectionColor color.NRGBA
	// TextSize determines the size of the text glyphs.
	TextSize unit.Sp
}

type TextStyleOption func(*TextStyleOptions)

func TextStyleWithFont(font font.Font) TextStyleOption {
	return func(o *TextStyleOptions) {
		o.Font = font
	}
}

func TextStyleWithColor(color color.NRGBA) TextStyleOption {
	return func(o *TextStyleOptions) {
		o.Color = color
	}
}

func TextStyleWithTextSize(sizeInSP float32) TextStyleOption {
	return func(o *TextStyleOptions) {
		o.TextSize = unit.Sp(sizeInSP)
	}
}
