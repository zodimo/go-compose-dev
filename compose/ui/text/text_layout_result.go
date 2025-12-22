package text

import (
	"github.com/zodimo/go-compose/compose/ui/text/style"
	"github.com/zodimo/go-compose/compose/ui/unit"
)

// TextLayoutInput holds the parameters used for computing a text layout result.
type TextLayoutInput struct {
	Text               AnnotatedString
	Style              TextStyle
	Overflow           style.TextOverFlow
	SoftWrap           bool
	MaxLines           int
	Density            unit.Density
	FontFamilyResolver interface{} // Placeholder
}

type MultiParagraph interface {
	// Placeholder for now
	LineCount() int
	IsLineEllipsized(lineIndex int) bool
	DidOverflowWidth() bool
	DidOverflowHeight() bool
}

// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-text/src/commonMain/kotlin/androidx/compose/ui/text/TextLayoutResult.kt;drc=4970f6e96cdb06089723da0ab8ec93ae3f067c7a;l=291
type TextLayoutResult struct {
	layoutInput    TextLayoutInput
	multiParagraph MultiParagraph
	size           unit.IntSize
}

func NewTextLayoutResult(layoutInput TextLayoutInput, multiparagraph MultiParagraph, size unit.IntSize) TextLayoutResult {
	return TextLayoutResult{
		layoutInput:    layoutInput,
		multiParagraph: multiparagraph,
		size:           size,
	}
}

func (r TextLayoutResult) LayoutInput() TextLayoutInput {
	return r.layoutInput
}

func (r TextLayoutResult) DidOverflowWidth() bool {
	if r.multiParagraph == nil {
		return false
	}
	return r.multiParagraph.DidOverflowWidth()
}

func (r TextLayoutResult) DidOverflowHeight() bool {
	if r.multiParagraph == nil {
		return false
	}
	return r.multiParagraph.DidOverflowHeight()
}

func (r TextLayoutResult) LineCount() int {
	if r.multiParagraph == nil {
		return 0
	}
	return r.multiParagraph.LineCount()
}

func (r TextLayoutResult) IsLineEllipsized(lineIndex int) bool {
	if r.multiParagraph == nil {
		return false
	}
	return r.multiParagraph.IsLineEllipsized(lineIndex)
}

func (r TextLayoutResult) DidOverflowBounds() bool {
	return r.DidOverflowWidth() || r.DidOverflowHeight()
}
