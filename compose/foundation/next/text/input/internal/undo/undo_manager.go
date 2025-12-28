package undo

// DefaultUndoCapacity is the default maximum number of undo operations to keep.
const DefaultUndoCapacity = 100

// UndoManager is a generic undo/redo stack manager.
//
// It maintains two stacks: undo and redo. Recording a new action clears the
// redo stack. Undo moves the top of undo to redo, and redo does the reverse.
//
// The total capacity limits the combined size of both stacks. When capacity
// is exceeded, the oldest undo entries are discarded.
//
// This is a port of androidx.compose.foundation.text.input.internal.undo.UndoManager.
type UndoManager[T any] struct {
	undoStack []T
	redoStack []T
	capacity  int
}

// NewUndoManager creates a new UndoManager with the given capacity.
func NewUndoManager[T any](capacity int) *UndoManager[T] {
	if capacity <= 0 {
		capacity = DefaultUndoCapacity
	}
	return &UndoManager[T]{
		undoStack: make([]T, 0, 16),
		redoStack: make([]T, 0, 16),
		capacity:  capacity,
	}
}

// NewUndoManagerWithStacks creates an UndoManager with initial stacks.
// This is used for restoring from saved state.
func NewUndoManagerWithStacks[T any](undoStack, redoStack []T, capacity int) *UndoManager[T] {
	if capacity <= 0 {
		capacity = DefaultUndoCapacity
	}
	um := &UndoManager[T]{
		capacity: capacity,
	}
	um.undoStack = make([]T, len(undoStack))
	copy(um.undoStack, undoStack)
	um.redoStack = make([]T, len(redoStack))
	copy(um.redoStack, redoStack)
	return um
}

// CanUndo returns true if there are operations to undo.
func (u *UndoManager[T]) CanUndo() bool {
	return len(u.undoStack) > 0
}

// CanRedo returns true if there are operations to redo.
func (u *UndoManager[T]) CanRedo() bool {
	return len(u.redoStack) > 0
}

// Size returns the total number of operations in both stacks.
func (u *UndoManager[T]) Size() int {
	return len(u.undoStack) + len(u.redoStack)
}

// UndoStackSize returns the number of operations in the undo stack.
func (u *UndoManager[T]) UndoStackSize() int {
	return len(u.undoStack)
}

// RedoStackSize returns the number of operations in the redo stack.
func (u *UndoManager[T]) RedoStackSize() int {
	return len(u.redoStack)
}

// Record adds a new undoable action to the stack.
//
// This clears the redo stack and adds the action to the undo stack.
// If capacity is exceeded, the oldest undo entries are discarded.
func (u *UndoManager[T]) Record(action T) {
	// Clear redo stack
	u.redoStack = u.redoStack[:0]

	// Make room if needed
	for u.Size() >= u.capacity {
		if len(u.undoStack) > 0 {
			// Remove oldest undo entry
			u.undoStack = u.undoStack[1:]
		} else {
			break
		}
	}

	// Add new action
	u.undoStack = append(u.undoStack, action)
}

// Undo pops the top operation from the undo stack and moves it to redo.
//
// Panics if CanUndo() is false.
func (u *UndoManager[T]) Undo() T {
	if !u.CanUndo() {
		panic("cannot undo: undo stack is empty - check CanUndo() before calling")
	}

	// Pop from undo stack
	lastIdx := len(u.undoStack) - 1
	action := u.undoStack[lastIdx]
	u.undoStack = u.undoStack[:lastIdx]

	// Push to redo stack
	u.redoStack = append(u.redoStack, action)

	return action
}

// Redo pops the top operation from the redo stack and moves it to undo.
//
// Panics if CanRedo() is false.
func (u *UndoManager[T]) Redo() T {
	if !u.CanRedo() {
		panic("cannot redo: redo stack is empty - check CanRedo() before calling")
	}

	// Pop from redo stack
	lastIdx := len(u.redoStack) - 1
	action := u.redoStack[lastIdx]
	u.redoStack = u.redoStack[:lastIdx]

	// Push to undo stack
	u.undoStack = append(u.undoStack, action)

	return action
}

// ClearHistory removes all undo and redo operations.
func (u *UndoManager[T]) ClearHistory() {
	u.undoStack = u.undoStack[:0]
	u.redoStack = u.redoStack[:0]
}

// PeekUndo returns the top of the undo stack without removing it.
// Panics if the stack is empty.
func (u *UndoManager[T]) PeekUndo() T {
	if !u.CanUndo() {
		panic("cannot peek: undo stack is empty")
	}
	return u.undoStack[len(u.undoStack)-1]
}

// PeekRedo returns the top of the redo stack without removing it.
// Panics if the stack is empty.
func (u *UndoManager[T]) PeekRedo() T {
	if !u.CanRedo() {
		panic("cannot peek: redo stack is empty")
	}
	return u.redoStack[len(u.redoStack)-1]
}

// ReplaceTop replaces the top of the undo stack without clearing redo.
// This is used for merging operations.
// Panics if the stack is empty.
func (u *UndoManager[T]) ReplaceTop(action T) {
	if !u.CanUndo() {
		panic("cannot replace: undo stack is empty")
	}
	u.undoStack[len(u.undoStack)-1] = action
}

// GetUndoStack returns a copy of the undo stack (for serialization).
func (u *UndoManager[T]) GetUndoStack() []T {
	result := make([]T, len(u.undoStack))
	copy(result, u.undoStack)
	return result
}

// GetRedoStack returns a copy of the redo stack (for serialization).
func (u *UndoManager[T]) GetRedoStack() []T {
	result := make([]T, len(u.redoStack))
	copy(result, u.redoStack)
	return result
}

// Capacity returns the maximum number of operations that can be stored.
func (u *UndoManager[T]) Capacity() int {
	return u.capacity
}
