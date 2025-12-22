package style

import "testing"

func TestTextOverFlowValues(t *testing.T) {
	// Verify the values match Kotlin's explicit constants
	tests := []struct {
		name     string
		overflow TextOverFlow
		want     int
	}{
		{"Clip", OverFlowClip, 1},
		{"Ellipsis", OverFlowEllipsis, 2},
		{"Visible", OverFlowVisible, 3},
		{"StartEllipsis", OverFlowStartEllipsis, 4},
		{"MiddleEllipsis", OverFlowMiddleEllipsis, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.overflow) != tt.want {
				t.Errorf("%s = %d, want %d", tt.name, tt.overflow, tt.want)
			}
		})
	}
}

func TestTextOverFlowString(t *testing.T) {
	tests := []struct {
		overflow TextOverFlow
		want     string
	}{
		{OverFlowClip, "Clip"},
		{OverFlowEllipsis, "Ellipsis"},
		{OverFlowVisible, "Visible"},
		{OverFlowStartEllipsis, "StartEllipsis"},
		{OverFlowMiddleEllipsis, "MiddleEllipsis"},
		{TextOverFlow(0), "Invalid"},
		{TextOverFlow(99), "Invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.overflow.String(); got != tt.want {
				t.Errorf("TextOverFlow(%d).String() = %q, want %q", tt.overflow, got, tt.want)
			}
		})
	}
}
