// Package input provides text field state management and editing infrastructure.
//
// This package is a Go port of Android Jetpack Compose's foundation text input module,
// providing types for managing text field state, undo/redo, input/output transformations,
// and selection handling.
//
// Key types:
//   - TextFieldState: Main state holder for text field contents
//   - TextFieldBuffer: Mutable text buffer for editing operations
//   - TextFieldCharSequence: Immutable snapshot of text field contents
//   - InputTransformation: Filter/transform user input
//   - OutputTransformation: Transform displayed text
//
// This package builds on compose/ui/text for TextRange, AnnotatedString, and styling types.
package input
