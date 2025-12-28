package undo

import (
	"testing"
)

func TestUndoManager_Empty(t *testing.T) {
	um := NewUndoManager[string](10)
	if um.CanUndo() {
		t.Error("expected CanUndo false for empty manager")
	}
	if um.CanRedo() {
		t.Error("expected CanRedo false for empty manager")
	}
	if um.Size() != 0 {
		t.Errorf("expected Size 0, got %d", um.Size())
	}
}

func TestUndoManager_Record(t *testing.T) {
	um := NewUndoManager[string](10)
	um.Record("action1")

	if !um.CanUndo() {
		t.Error("expected CanUndo true after Record")
	}
	if um.CanRedo() {
		t.Error("expected CanRedo false after Record")
	}
	if um.Size() != 1 {
		t.Errorf("expected Size 1, got %d", um.Size())
	}
}

func TestUndoManager_UndoRedo(t *testing.T) {
	um := NewUndoManager[string](10)
	um.Record("action1")
	um.Record("action2")

	// Undo
	action := um.Undo()
	if action != "action2" {
		t.Errorf("expected 'action2', got '%s'", action)
	}
	if !um.CanUndo() {
		t.Error("expected CanUndo true")
	}
	if !um.CanRedo() {
		t.Error("expected CanRedo true")
	}

	// Redo
	action = um.Redo()
	if action != "action2" {
		t.Errorf("expected 'action2', got '%s'", action)
	}
	if !um.CanUndo() {
		t.Error("expected CanUndo true after redo")
	}
	if um.CanRedo() {
		t.Error("expected CanRedo false after redo")
	}
}

func TestUndoManager_RecordClearsRedo(t *testing.T) {
	um := NewUndoManager[string](10)
	um.Record("action1")
	um.Undo()

	if !um.CanRedo() {
		t.Error("expected CanRedo true after undo")
	}

	um.Record("action2")

	if um.CanRedo() {
		t.Error("expected CanRedo false after new record")
	}
}

func TestUndoManager_Capacity(t *testing.T) {
	um := NewUndoManager[int](3)
	um.Record(1)
	um.Record(2)
	um.Record(3)
	um.Record(4) // Should evict 1

	if um.Size() != 3 {
		t.Errorf("expected Size 3, got %d", um.Size())
	}

	// Verify oldest was removed
	action1 := um.Undo() // 4
	action2 := um.Undo() // 3
	action3 := um.Undo() // 2

	if action1 != 4 || action2 != 3 || action3 != 2 {
		t.Errorf("unexpected actions: %d, %d, %d", action1, action2, action3)
	}

	if um.CanUndo() {
		t.Error("expected CanUndo false - oldest should be evicted")
	}
}

func TestUndoManager_ClearHistory(t *testing.T) {
	um := NewUndoManager[string](10)
	um.Record("action1")
	um.Record("action2")
	um.Undo()

	um.ClearHistory()

	if um.CanUndo() {
		t.Error("expected CanUndo false after clear")
	}
	if um.CanRedo() {
		t.Error("expected CanRedo false after clear")
	}
	if um.Size() != 0 {
		t.Errorf("expected Size 0, got %d", um.Size())
	}
}

func TestUndoManager_PeekUndo(t *testing.T) {
	um := NewUndoManager[string](10)
	um.Record("action1")
	um.Record("action2")

	peeked := um.PeekUndo()
	if peeked != "action2" {
		t.Errorf("expected PeekUndo 'action2', got '%s'", peeked)
	}

	// Peek shouldn't remove
	if um.UndoStackSize() != 2 {
		t.Errorf("expected UndoStackSize 2, got %d", um.UndoStackSize())
	}
}

func TestUndoManager_ReplaceTop(t *testing.T) {
	um := NewUndoManager[string](10)
	um.Record("action1")
	um.ReplaceTop("modified")

	action := um.Undo()
	if action != "modified" {
		t.Errorf("expected 'modified', got '%s'", action)
	}
}
