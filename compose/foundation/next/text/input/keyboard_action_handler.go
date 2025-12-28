package input

// KeyboardActionHandler handles IME action events.
//
// IME actions are triggered by software keyboard buttons like "Done", "Go",
// "Search", "Send", "Next", or "Previous". On single-line text fields,
// the Enter key on a hardware keyboard also triggers this handler.
//
// This is a port of androidx.compose.foundation.text.input.KeyboardActionHandler.
type KeyboardActionHandler interface {
	// OnKeyboardAction is called when an IME action is performed.
	//
	// The performDefaultAction function executes the default behavior for
	// the action (e.g., closing the keyboard, moving focus). Call it if
	// you want to include the default behavior along with custom handling.
	//
	// If you want to prevent the default behavior, don't call performDefaultAction.
	OnKeyboardAction(performDefaultAction func())
}

// KeyboardActionHandlerFunc is a function type that implements KeyboardActionHandler.
type KeyboardActionHandlerFunc func(performDefaultAction func())

// OnKeyboardAction implements KeyboardActionHandler.
func (f KeyboardActionHandlerFunc) OnKeyboardAction(performDefaultAction func()) {
	f(performDefaultAction)
}

// DefaultKeyboardActionHandler performs only the default action.
var DefaultKeyboardActionHandler KeyboardActionHandler = KeyboardActionHandlerFunc(func(performDefaultAction func()) {
	performDefaultAction()
})

// NoOpKeyboardActionHandler does nothing when an action is triggered.
var NoOpKeyboardActionHandler KeyboardActionHandler = KeyboardActionHandlerFunc(func(performDefaultAction func()) {
	// Do nothing
})
