package textfield

import (
	"github.com/zodimo/go-compose/compose/foundation/text/input"
	"github.com/zodimo/go-compose/compose/ui"
)

// DefaultTextFieldOptions returns the default options for BasicTextField.
func DefaultTextFieldOptions() TextFieldOptions {
	return TextFieldOptions{
		Modifier:             ui.EmptyModifier,
		Enabled:              true,
		ReadOnly:             false,
		TextStyle:            nil,
		InputTransformation:  nil,
		OutputTransformation: nil,
		LineLimits:           input.TextFieldLineLimitsDefault,
		// OnTextLayout:         nil,
		// KeyboardActionHandler: nil,
		Decorator: input.NoDecorationDecorator,
		// CursorColor:           nil,
	}
}
