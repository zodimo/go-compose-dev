package style

import (
	"github.com/zodimo/go-compose/compose/ui/unit"
)

//https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-text/src/commonMain/kotlin/androidx/compose/ui/text/style/TextIndent.kt;drc=4970f6e96cdb06089723da0ab8ec93ae3f067c7a;l=32

// TextIndent specifies the indentation of the first line and the following lines.
type TextIndent struct {
	FirstLine unit.TextUnit
	RestLine  unit.TextUnit
}

// NewTextIndent creates a new TextIndent.
func NewTextIndent(firstLine, restLine unit.TextUnit) TextIndent {
	return TextIndent{
		FirstLine: firstLine,
		RestLine:  restLine,
	}
}

// LerpTextIndent interpolates between two TextIndents.
func LerpTextIndent(start, stop TextIndent, fraction float32) TextIndent {
	return TextIndent{
		FirstLine: unit.LerpTextUnit(start.FirstLine, stop.FirstLine, fraction),
		RestLine:  unit.LerpTextUnit(start.RestLine, stop.RestLine, fraction),
	}
}

var TextIndentNone = TextIndent{FirstLine: unit.Sp(0), RestLine: unit.Sp(0)}
