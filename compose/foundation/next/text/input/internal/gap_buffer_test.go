package internal

import (
	"testing"
)

func TestGapBuffer_NewGapBuffer(t *testing.T) {
	gb := NewGapBuffer("Hello")
	if gb.Length() != 5 {
		t.Errorf("expected length 5, got %d", gb.Length())
	}
	if gb.String() != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", gb.String())
	}
}

func TestGapBuffer_Get(t *testing.T) {
	gb := NewGapBuffer("Hello")
	tests := []struct {
		index int
		want  rune
	}{
		{0, 'H'},
		{1, 'e'},
		{4, 'o'},
	}
	for _, tt := range tests {
		got := gb.Get(tt.index)
		if got != tt.want {
			t.Errorf("Get(%d) = %q, want %q", tt.index, got, tt.want)
		}
	}
}

func TestGapBuffer_Replace(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		start    int
		end      int
		text     string
		expected string
	}{
		{"insert at start", "World", 0, 0, "Hello ", "Hello World"},
		{"insert at end", "Hello", 5, 5, " World", "Hello World"},
		{"insert in middle", "Hllo", 1, 1, "e", "Hello"},
		{"replace single", "Hxllo", 1, 2, "e", "Hello"},
		{"replace multiple", "Hxxxo", 1, 4, "ell", "Hello"},
		{"delete", "Helllo", 3, 4, "", "Hello"},
		{"replace all", "Goodbye", 0, 7, "Hello", "Hello"},
		{"empty to text", "", 0, 0, "Hello", "Hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gb := NewGapBuffer(tt.initial)
			gb.Replace(tt.start, tt.end, tt.text)
			got := gb.String()
			if got != tt.expected {
				t.Errorf("Replace(%d, %d, %q) = %q, want %q", tt.start, tt.end, tt.text, got, tt.expected)
			}
		})
	}
}

func TestGapBuffer_Append(t *testing.T) {
	gb := NewGapBuffer("Hello")
	gb.Append(" World")
	if gb.String() != "Hello World" {
		t.Errorf("expected 'Hello World', got '%s'", gb.String())
	}
}

func TestGapBuffer_Insert(t *testing.T) {
	gb := NewGapBuffer("Hllo")
	gb.Insert(1, "e")
	if gb.String() != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", gb.String())
	}
}

func TestGapBuffer_Delete(t *testing.T) {
	gb := NewGapBuffer("Helllo")
	gb.Delete(3, 4)
	if gb.String() != "Hello" {
		t.Errorf("expected 'Hello', got '%s'", gb.String())
	}
}

func TestGapBuffer_SubSequence(t *testing.T) {
	gb := NewGapBuffer("Hello World")
	tests := []struct {
		start int
		end   int
		want  string
	}{
		{0, 5, "Hello"},
		{6, 11, "World"},
		{0, 11, "Hello World"},
		{2, 4, "ll"},
	}
	for _, tt := range tests {
		got := gb.SubSequence(tt.start, tt.end)
		if got != tt.want {
			t.Errorf("SubSequence(%d, %d) = %q, want %q", tt.start, tt.end, got, tt.want)
		}
	}
}

func TestGapBuffer_Unicode(t *testing.T) {
	gb := NewGapBuffer("Hello 世界")
	if gb.Length() != 8 {
		t.Errorf("expected length 8, got %d", gb.Length())
	}
	if gb.Get(6) != '世' {
		t.Errorf("expected '世', got %q", gb.Get(6))
	}
	gb.Replace(6, 8, "🌍")
	if gb.String() != "Hello 🌍" {
		t.Errorf("expected 'Hello 🌍', got '%s'", gb.String())
	}
}

func TestGapBuffer_ContentEquals(t *testing.T) {
	gb := NewGapBuffer("Hello")
	if !gb.ContentEquals("Hello") {
		t.Error("expected ContentEquals to be true")
	}
	if gb.ContentEquals("World") {
		t.Error("expected ContentEquals to be false for different text")
	}
	if gb.ContentEquals("Hello World") {
		t.Error("expected ContentEquals to be false for different length")
	}
}

func TestPartialGapBuffer_Basic(t *testing.T) {
	pgb := NewPartialGapBuffer("Hello World")
	if pgb.Length() != 11 {
		t.Errorf("expected length 11, got %d", pgb.Length())
	}
	pgb.Replace(6, 11, "Go")
	if pgb.String() != "Hello Go" {
		t.Errorf("expected 'Hello Go', got '%s'", pgb.String())
	}
}

func TestPartialGapBuffer_MultipleEdits(t *testing.T) {
	pgb := NewPartialGapBuffer("Hello")
	pgb.Replace(5, 5, " ")
	pgb.Replace(6, 6, "World")
	if pgb.String() != "Hello World" {
		t.Errorf("expected 'Hello World', got '%s'", pgb.String())
	}
}

func TestPartialGapBuffer_Get(t *testing.T) {
	pgb := NewPartialGapBuffer("Hello")
	pgb.Replace(5, 5, " World")
	// Test getting characters from different regions
	tests := []struct {
		index int
		want  rune
	}{
		{0, 'H'},  // Before buffer
		{5, ' '},  // In buffer
		{6, 'W'},  // In buffer
		{10, 'd'}, // After buffer
	}
	for _, tt := range tests {
		got := pgb.Get(tt.index)
		if got != tt.want {
			t.Errorf("Get(%d) = %q, want %q", tt.index, got, tt.want)
		}
	}
}
