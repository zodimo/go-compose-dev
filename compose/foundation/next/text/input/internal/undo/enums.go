package undo

// TextEditType categorizes the nature of a text change.
//
// This is used to determine whether adjacent undo operations can be merged.
type TextEditType int

const (
	// TextEditTypeInsert indicates text was inserted (zero-length range replaced with text).
	TextEditTypeInsert TextEditType = iota

	// TextEditTypeDelete indicates text was deleted (non-zero range replaced with nothing).
	TextEditTypeDelete

	// TextEditTypeReplace indicates text was replaced (non-zero range replaced with text).
	TextEditTypeReplace
)

// String returns a string representation of TextEditType.
func (t TextEditType) String() string {
	switch t {
	case TextEditTypeInsert:
		return "Insert"
	case TextEditTypeDelete:
		return "Delete"
	case TextEditTypeReplace:
		return "Replace"
	default:
		return "Unknown"
	}
}

// TextDeleteType identifies the direction of deletion for merge logic.
//
// This is used to determine if consecutive deletions can be merged.
// For example, consecutive backspace presses can be merged, but a mix
// of backspace and delete cannot.
type TextDeleteType int

const (
	// TextDeleteTypeStart represents backspace behavior (delete toward 0).
	// "abcd|efg" → "abc|efg"
	TextDeleteTypeStart TextDeleteType = iota

	// TextDeleteTypeEnd represents delete key behavior (delete toward length).
	// "abcd|efg" → "abcd|fg"
	TextDeleteTypeEnd

	// TextDeleteTypeInner represents selection deletion (directionless).
	// "ab|cde|fg" → "ab|fg"
	TextDeleteTypeInner

	// TextDeleteTypeNotByUser represents programmatic deletion not adjacent to cursor.
	// This cannot be merged with user-initiated deletions.
	TextDeleteTypeNotByUser
)

// String returns a string representation of TextDeleteType.
func (t TextDeleteType) String() string {
	switch t {
	case TextDeleteTypeStart:
		return "Start"
	case TextDeleteTypeEnd:
		return "End"
	case TextDeleteTypeInner:
		return "Inner"
	case TextDeleteTypeNotByUser:
		return "NotByUser"
	default:
		return "Unknown"
	}
}

// TextFieldEditUndoBehavior controls undo stack behavior for edits.
//
// This determines how new edits interact with the undo history.
type TextFieldEditUndoBehavior int

const (
	// TextFieldEditUndoBehaviorMergeIfPossible attempts to merge with the previous undo entry.
	// This is the default behavior for normal typing.
	TextFieldEditUndoBehaviorMergeIfPossible TextFieldEditUndoBehavior = iota

	// TextFieldEditUndoBehaviorClearHistory clears all undo/redo history.
	// Used for programmatic updates that shouldn't be undoable.
	TextFieldEditUndoBehaviorClearHistory

	// TextFieldEditUndoBehaviorNeverMerge always creates a new undo entry.
	// Used for cut, paste, and other atomic operations.
	TextFieldEditUndoBehaviorNeverMerge
)

// String returns a string representation of TextFieldEditUndoBehavior.
func (t TextFieldEditUndoBehavior) String() string {
	switch t {
	case TextFieldEditUndoBehaviorMergeIfPossible:
		return "MergeIfPossible"
	case TextFieldEditUndoBehaviorClearHistory:
		return "ClearHistory"
	case TextFieldEditUndoBehaviorNeverMerge:
		return "NeverMerge"
	default:
		return "Unknown"
	}
}
