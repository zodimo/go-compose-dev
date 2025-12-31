package text

import (
	"testing"
)

func TestNewTextRange(t *testing.T) {
	r := NewTextRange(5, 10)
	if r.Start != 5 {
		t.Errorf("expected Start=5, got %d", r.Start)
	}
	if r.End != 10 {
		t.Errorf("expected End=10, got %d", r.End)
	}
}

func TestNewTextRange_PanicsOnNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on negative start")
		}
	}()
	NewTextRange(-1, 10)
}

func TestTextRange_MinMax(t *testing.T) {
	tests := []struct {
		name    string
		start   int
		end     int
		wantMin int
		wantMax int
	}{
		{"normal", 5, 10, 5, 10},
		{"reversed", 10, 5, 5, 10},
		{"same", 5, 5, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := TextRange{Start: tt.start, End: tt.end}
			if got := r.Min(); got != tt.wantMin {
				t.Errorf("Min() = %d, want %d", got, tt.wantMin)
			}
			if got := r.Max(); got != tt.wantMax {
				t.Errorf("Max() = %d, want %d", got, tt.wantMax)
			}
		})
	}
}

func TestTextRange_Collapsed(t *testing.T) {
	if r := (TextRange{Start: 5, End: 5}); !r.Collapsed() {
		t.Error("expected collapsed for equal start/end")
	}
	if r := (TextRange{Start: 5, End: 10}); r.Collapsed() {
		t.Error("expected not collapsed for different start/end")
	}
}

func TestTextRange_Reversed(t *testing.T) {
	if r := (TextRange{Start: 10, End: 5}); !r.Reversed() {
		t.Error("expected reversed when start > end")
	}
	if r := (TextRange{Start: 5, End: 10}); r.Reversed() {
		t.Error("expected not reversed when start < end")
	}
}

func TestTextRange_Length(t *testing.T) {
	tests := []struct {
		start int
		end   int
		want  int
	}{
		{5, 10, 5},
		{10, 5, 5},
		{5, 5, 0},
	}

	for _, tt := range tests {
		r := TextRange{Start: tt.start, End: tt.end}
		if got := r.Length(); got != tt.want {
			t.Errorf("TextRange(%d, %d).Length() = %d, want %d", tt.start, tt.end, got, tt.want)
		}
	}
}

func TestTextRange_Intersects(t *testing.T) {
	tests := []struct {
		name   string
		r1     TextRange
		r2     TextRange
		expect bool
	}{
		{"overlapping", TextRange{5, 10}, TextRange{8, 15}, true},
		{"adjacent", TextRange{5, 10}, TextRange{10, 15}, false},
		{"no overlap", TextRange{5, 10}, TextRange{15, 20}, false},
		{"contained", TextRange{5, 15}, TextRange{8, 12}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r1.Intersects(tt.r2); got != tt.expect {
				t.Errorf("Intersects() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestTextRange_Contains(t *testing.T) {
	tests := []struct {
		name   string
		r1     TextRange
		r2     TextRange
		expect bool
	}{
		{"contains", TextRange{5, 15}, TextRange{8, 12}, true},
		{"equal", TextRange{5, 10}, TextRange{5, 10}, true},
		{"partial overlap", TextRange{5, 10}, TextRange{8, 15}, false},
		{"outside", TextRange{5, 10}, TextRange{15, 20}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r1.Contains(tt.r2); got != tt.expect {
				t.Errorf("Contains() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestTextRange_ContainsOffset(t *testing.T) {
	r := TextRange{Start: 5, End: 10}

	if !r.ContainsOffset(5) {
		t.Error("expected 5 to be contained (inclusive)")
	}
	if !r.ContainsOffset(7) {
		t.Error("expected 7 to be contained")
	}
	if r.ContainsOffset(10) {
		t.Error("expected 10 to not be contained (exclusive)")
	}
	if r.ContainsOffset(4) {
		t.Error("expected 4 to not be contained")
	}
}

func TestTextRange_CoerceIn(t *testing.T) {
	r := TextRange{Start: 2, End: 20}
	coerced := r.CoerceIn(5, 15)

	if coerced.Start != 5 {
		t.Errorf("expected Start=5, got %d", coerced.Start)
	}
	if coerced.End != 15 {
		t.Errorf("expected End=15, got %d", coerced.End)
	}
}

func TestTextRange_String(t *testing.T) {
	r := TextRange{Start: 5, End: 10}
	expected := "TextRange(5, 10)"
	if got := r.String(); got != expected {
		t.Errorf("String() = %q, want %q", got, expected)
	}
}

func TestTextRange_Substring(t *testing.T) {
	r := TextRange{Start: 6, End: 11}
	s := "Hello World"
	if got := r.Substring(s); got != "World" {
		t.Errorf("Substring() = %q, want %q", got, "World")
	}
}

func TestTextRangeZero(t *testing.T) {
	if TextRangeZero.Start != 0 || TextRangeZero.End != 0 {
		t.Error("TextRangeZero should have Start=0 and End=0")
	}
}
