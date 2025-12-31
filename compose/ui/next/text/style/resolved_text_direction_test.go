package style

import "testing"

func TestResolvedTextDirection_String(t *testing.T) {
	tests := []struct {
		direction ResolvedTextDirection
		want      string
	}{
		{ResolvedTextDirectionLtr, "Ltr"},
		{ResolvedTextDirectionRtl, "Rtl"},
		{ResolvedTextDirection(99), "Unknown"},
	}

	for _, tt := range tests {
		if got := tt.direction.String(); got != tt.want {
			t.Errorf("ResolvedTextDirection(%d).String() = %q, want %q", tt.direction, got, tt.want)
		}
	}
}

func TestResolvedTextDirection_Values(t *testing.T) {
	// Verify the enum values
	if ResolvedTextDirectionLtr != 0 {
		t.Errorf("ResolvedTextDirectionLtr should be 0, got %d", ResolvedTextDirectionLtr)
	}
	if ResolvedTextDirectionRtl != 1 {
		t.Errorf("ResolvedTextDirectionRtl should be 1, got %d", ResolvedTextDirectionRtl)
	}
}
