package internal

// KeyEvent represents a hardware keyboard event.
// This is a stub type that would be provided by the platform layer.
type KeyEvent struct {
	// Key is the key that was pressed (e.g., KeyA, KeyBackspace).
	Key Key

	// Type is whether this is a key down or key up event.
	Type KeyEventType

	// Modifiers indicates which modifier keys were held.
	Modifiers KeyModifiers
}

// Key represents a keyboard key.
type Key int

// Common key constants
const (
	KeyUnknown Key = iota
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeyBackspace
	KeyDelete
	KeyTab
	KeyEnter
	KeyEscape
	KeyLeft
	KeyRight
	KeyUp
	KeyDown
	KeyHome
	KeyEnd
	KeyPageUp
	KeyPageDown
)

// KeyEventType indicates key press or release.
type KeyEventType int

const (
	KeyEventTypeDown KeyEventType = iota
	KeyEventTypeUp
)

// KeyModifiers represents modifier key states.
type KeyModifiers struct {
	Ctrl  bool
	Shift bool
	Alt   bool
	Meta  bool // Command on Mac, Windows key on Windows
}

// TextFieldKeyEventHandler processes keyboard events for text fields.
//
// This is a stub interface that documents the operational semantics.
// A full implementation requires platform-specific key mapping and
// integration with the text editing commands.
//
// Operational Semantics:
//
// Character Input:
//   - Printable characters are inserted at the cursor
//   - If selection exists, it is replaced
//
// Navigation Keys:
//   - Left/Right: Move cursor one character
//   - Shift+Left/Right: Extend selection by one character
//   - Ctrl+Left/Right: Move cursor by word
//   - Ctrl+Shift+Left/Right: Extend selection by word
//   - Home/End: Move to line start/end
//   - Ctrl+Home/End: Move to text start/end
//   - Up/Down: Move to previous/next line (multi-line only)
//
// Deletion Keys:
//   - Backspace: Delete character before cursor (or selection)
//   - Delete: Delete character after cursor (or selection)
//   - Ctrl+Backspace: Delete word before cursor
//   - Ctrl+Delete: Delete word after cursor
//
// Selection and Clipboard:
//   - Ctrl+A: Select all
//   - Ctrl+C: Copy selection
//   - Ctrl+X: Cut selection
//   - Ctrl+V: Paste from clipboard
//
// Undo/Redo:
//   - Ctrl+Z: Undo
//   - Ctrl+Shift+Z or Ctrl+Y: Redo
//
// Platform Variations:
//   - On Mac, Cmd replaces Ctrl for most shortcuts
//   - Some platforms have additional shortcuts
//
// This is a port of key handling logic from
// androidx.compose.foundation.text.input.internal.TextFieldDecoratorModifier.
type TextFieldKeyEventHandler interface {
	// HandleKeyEvent processes a key event and returns whether it was consumed.
	//
	// Parameters:
	//   - event: The keyboard event
	//   - state: The transformed text field state
	//   - textLayoutState: The text layout state for cursor positioning
	//
	// Returns true if the event was handled, false to propagate it.
	HandleKeyEvent(
		event KeyEvent,
		state TransformedTextFieldStateInterface,
		textLayoutState TextLayoutState,
	) bool
}

// KeyCommand represents a text editing command derived from a key event.
type KeyCommand int

const (
	KeyCommandNone KeyCommand = iota

	// Character manipulation
	KeyCommandDeletePrevChar
	KeyCommandDeleteNextChar
	KeyCommandDeletePrevWord
	KeyCommandDeleteNextWord
	KeyCommandDeleteToLineStart
	KeyCommandDeleteToLineEnd

	// Cursor movement
	KeyCommandMovePrevChar
	KeyCommandMoveNextChar
	KeyCommandMovePrevWord
	KeyCommandMoveNextWord
	KeyCommandMoveLineStart
	KeyCommandMoveLineEnd
	KeyCommandMoveTextStart
	KeyCommandMoveTextEnd
	KeyCommandMovePrevLine
	KeyCommandMoveNextLine
	KeyCommandMovePrevPage
	KeyCommandMoveNextPage

	// Selection
	KeyCommandSelectPrevChar
	KeyCommandSelectNextChar
	KeyCommandSelectPrevWord
	KeyCommandSelectNextWord
	KeyCommandSelectLineStart
	KeyCommandSelectLineEnd
	KeyCommandSelectTextStart
	KeyCommandSelectTextEnd
	KeyCommandSelectPrevLine
	KeyCommandSelectNextLine
	KeyCommandSelectPrevPage
	KeyCommandSelectNextPage
	KeyCommandSelectAll

	// Clipboard
	KeyCommandCopy
	KeyCommandCut
	KeyCommandPaste

	// Undo/Redo
	KeyCommandUndo
	KeyCommandRedo

	// Misc
	KeyCommandNewLine
	KeyCommandTab
)
