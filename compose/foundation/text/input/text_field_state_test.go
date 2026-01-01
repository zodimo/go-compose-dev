package input

import (
	"testing"
)

func TestTextFieldState_New(t *testing.T) {
	state := NewTextFieldState("Hello")

	if state.Text() != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", state.Text())
	}
	if state.Length() != 5 {
		t.Errorf("expected length 5, got %d", state.Length())
	}
	// Cursor should be at end
	sel := state.Selection()
	if sel.Start != 5 || sel.End != 5 {
		t.Errorf("expected selection (5,5), got (%d,%d)", sel.Start, sel.End)
	}
}

func TestTextFieldState_Edit(t *testing.T) {
	state := NewTextFieldState("Hello")

	state.Edit(func(buffer *TextFieldBuffer) {
		buffer.Append(" World")
	})

	if state.Text() != "Hello World" {
		t.Errorf("expected 'Hello World', got '%s'", state.Text())
	}
}

func TestTextFieldState_EditMultiple(t *testing.T) {
	state := NewTextFieldState("")

	state.Edit(func(buffer *TextFieldBuffer) {
		buffer.Append("Hello")
	})
	state.Edit(func(buffer *TextFieldBuffer) {
		buffer.Append(" ")
	})
	state.Edit(func(buffer *TextFieldBuffer) {
		buffer.Append("World")
	})

	if state.Text() != "Hello World" {
		t.Errorf("expected 'Hello World', got '%s'", state.Text())
	}
}

func TestTextFieldState_SetTextAndPlaceCursorAtEnd(t *testing.T) {
	state := NewTextFieldState("")
	state.SetTextAndPlaceCursorAtEnd("Hello")

	if state.Text() != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", state.Text())
	}
	sel := state.Selection()
	if sel.Start != 5 || sel.End != 5 {
		t.Errorf("expected selection (5,5), got (%d,%d)", sel.Start, sel.End)
	}
}

func TestTextFieldState_SetTextAndSelectAll(t *testing.T) {
	state := NewTextFieldState("")
	state.SetTextAndSelectAll("Hello")

	if state.Text() != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", state.Text())
	}
	sel := state.Selection()
	if sel.Start != 0 || sel.End != 5 {
		t.Errorf("expected selection (0,5), got (%d,%d)", sel.Start, sel.End)
	}
}

func TestTextFieldState_ClearText(t *testing.T) {
	state := NewTextFieldState("Hello")
	state.ClearText()

	if state.Text() != "" {
		t.Errorf("expected empty string, got '%s'", state.Text())
	}
	if state.Length() != 0 {
		t.Errorf("expected length 0, got %d", state.Length())
	}
}

func TestTextFieldState_IsEmpty(t *testing.T) {
	state := NewTextFieldState("")
	if !state.IsEmpty() {
		t.Error("expected IsEmpty true for empty state")
	}

	state.SetTextAndPlaceCursorAtEnd("Hello")
	if state.IsEmpty() {
		t.Error("expected IsEmpty false for non-empty state")
	}
}

func TestTextFieldState_SelectAll(t *testing.T) {
	state := NewTextFieldState("Hello")
	state.SelectAll()

	sel := state.Selection()
	if sel.Start != 0 || sel.End != 5 {
		t.Errorf("expected selection (0,5), got (%d,%d)", sel.Start, sel.End)
	}
}

func TestTextFieldState_PlaceCursorAtEnd(t *testing.T) {
	state := NewTextFieldState("Hello")
	// Force cursor to start first
	state.Edit(func(buffer *TextFieldBuffer) {
		buffer.PlaceCursorBeforeCharAt(0)
	})

	state.PlaceCursorAtEnd()

	sel := state.Selection()
	if sel.Start != 5 || sel.End != 5 {
		t.Errorf("expected selection (5,5), got (%d,%d)", sel.Start, sel.End)
	}
}
