package graphics

// DrawStyle defines a way to draw something.
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-graphics/src/commonMain/kotlin/androidx/compose/ui/graphics/drawscope/DrawScope.kt;l=939
type DrawStyle interface {
	isDrawStyle()
}

// Fill is the default DrawStyle indicating shapes should be drawn completely filled.
var Fill DrawStyle = fillStyle{}

type fillStyle struct{}

func (fillStyle) isDrawStyle() {}

// DrawStyleFill is an alias for Fill for backwards compatibility.
var DrawStyleFill = Fill

func EqualDrawStyle(a, b DrawStyle) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	// Check if both are Fill
	if _, aFill := a.(fillStyle); aFill {
		if _, bFill := b.(fillStyle); bFill {
			return true
		}
		return false
	}
	// Check if both are Stroke
	if aStroke, ok := a.(*Stroke); ok {
		if bStroke, ok := b.(*Stroke); ok {
			return aStroke.Equal(bStroke)
		}
	}
	return false
}

func TakeOrElseDrawStyle(a, b DrawStyle) DrawStyle {
	if a != nil {
		return a
	}
	return b
}
