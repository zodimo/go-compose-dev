package theme

import (
	"image/color"
	"testing"
)

// Test HSL conversion round-trip
func TestRGBAToHSLRoundTrip(t *testing.T) {
	testCases := []struct {
		name  string
		color color.NRGBA
	}{
		{"Red", color.NRGBA{R: 255, G: 0, B: 0, A: 255}},
		{"Green", color.NRGBA{R: 0, G: 255, B: 0, A: 255}},
		{"Blue", color.NRGBA{R: 0, G: 0, B: 255, A: 255}},
		{"White", color.NRGBA{R: 255, G: 255, B: 255, A: 255}},
		{"Black", color.NRGBA{R: 0, G: 0, B: 0, A: 255}},
		{"Gray", color.NRGBA{R: 128, G: 128, B: 128, A: 255}},
		{"Purple", color.NRGBA{R: 128, G: 0, B: 128, A: 255}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h, s, l := rgbaToHSL(tc.color)
			result := hslToRGBA(h, s, l, tc.color.A)

			// Allow small rounding errors (±1 per channel)
			if diff := abs(int(tc.color.R) - int(result.R)); diff > 1 {
				t.Errorf("Red channel mismatch: expected %d, got %d (diff=%d)", tc.color.R, result.R, diff)
			}
			if diff := abs(int(tc.color.G) - int(result.G)); diff > 1 {
				t.Errorf("Green channel mismatch: expected %d, got %d (diff=%d)", tc.color.G, result.G, diff)
			}
			if diff := abs(int(tc.color.B) - int(result.B)); diff > 1 {
				t.Errorf("Blue channel mismatch: expected %d, got %d (diff=%d)", tc.color.B, result.B, diff)
			}
			if tc.color.A != result.A {
				t.Errorf("Alpha channel mismatch: expected %d, got %d", tc.color.A, result.A)
			}
		})
	}
}

// Test lighten function
func TestLightenColor(t *testing.T) {
	baseColor := color.NRGBA{R: 100, G: 100, B: 100, A: 255}

	t.Run("Lighten by 0.2", func(t *testing.T) {
		result := lightenColor(baseColor, 0.2)
		h, s, l := rgbaToHSL(result)
		baseH, baseS, baseL := rgbaToHSL(baseColor)

		if h != baseH {
			t.Errorf("Hue should not change: expected %f, got %f", baseH, h)
		}
		if s != baseS {
			t.Errorf("Saturation should not change: expected %f, got %f", baseS, s)
		}
		if l <= baseL {
			t.Errorf("Lightness should increase: base=%f, result=%f", baseL, l)
		}
		if diff := abs32(l - (baseL + 0.2)); diff > 0.01 {
			t.Errorf("Lightness should increase by 0.2: expected %f, got %f", baseL+0.2, l)
		}
	})

	t.Run("Lighten white (clamp test)", func(t *testing.T) {
		white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		result := lightenColor(white, 0.5)

		// Should remain white
		if result.R != 255 || result.G != 255 || result.B != 255 {
			t.Errorf("White should remain white after lightening: got RGB(%d,%d,%d)", result.R, result.G, result.B)
		}
	})

	t.Run("Preserve alpha channel", func(t *testing.T) {
		colorWithAlpha := color.NRGBA{R: 100, G: 100, B: 100, A: 128}
		result := lightenColor(colorWithAlpha, 0.2)
		if result.A != 128 {
			t.Errorf("Alpha should be preserved: expected 128, got %d", result.A)
		}
	})
}

// Test darken function
func TestDarkenColor(t *testing.T) {
	baseColor := color.NRGBA{R: 150, G: 150, B: 150, A: 255}

	t.Run("Darken by 0.2", func(t *testing.T) {
		result := darkenColor(baseColor, 0.2)
		h, s, l := rgbaToHSL(result)
		baseH, baseS, baseL := rgbaToHSL(baseColor)

		if h != baseH {
			t.Errorf("Hue should not change: expected %f, got %f", baseH, h)
		}
		if s != baseS {
			t.Errorf("Saturation should not change: expected %f, got %f", baseS, s)
		}
		if l >= baseL {
			t.Errorf("Lightness should decrease: base=%f, result=%f", baseL, l)
		}
		if diff := abs32(l - (baseL - 0.2)); diff > 0.01 {
			t.Errorf("Lightness should decrease by 0.2: expected %f, got %f", baseL-0.2, l)
		}
	})

	t.Run("Darken black (clamp test)", func(t *testing.T) {
		black := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
		result := darkenColor(black, 0.5)

		// Should remain black
		if result.R != 0 || result.G != 0 || result.B != 0 {
			t.Errorf("Black should remain black after darkening: got RGB(%d,%d,%d)", result.R, result.G, result.B)
		}
	})

	t.Run("Preserve alpha channel", func(t *testing.T) {
		colorWithAlpha := color.NRGBA{R: 150, G: 150, B: 150, A: 64}
		result := darkenColor(colorWithAlpha, 0.2)
		if result.A != 64 {
			t.Errorf("Alpha should be preserved: expected 64, got %d", result.A)
		}
	})
}

// Test ColorDescriptor Lighten method
func TestColorDescriptorLighten(t *testing.T) {
	colorHelper := ColorHelper
	baseColor := colorHelper.ColorSelector().PrimaryRoles.Primary

	t.Run("Lighten adds update", func(t *testing.T) {
		lightened := baseColor.Lighten(0.3)

		updates := lightened.Updates()
		if len(updates) != 1 {
			t.Fatalf("Should have one update, got %d", len(updates))
		}
		if updates[0].Action() != LightenColorUpdateAction {
			t.Errorf("Expected LightenColorUpdateAction, got %d", updates[0].Action())
		}
		if GetLighten(updates[0]) != 0.3 {
			t.Errorf("Expected percentage 0.3, got %f", GetLighten(updates[0]))
		}
	})

	t.Run("Multiple lightens chain", func(t *testing.T) {
		lightened := baseColor.Lighten(0.2).Lighten(0.1)

		updates := lightened.Updates()
		if len(updates) != 2 {
			t.Fatalf("Should have two updates, got %d", len(updates))
		}
		if updates[0].Action() != LightenColorUpdateAction {
			t.Errorf("First update should be Lighten")
		}
		if updates[1].Action() != LightenColorUpdateAction {
			t.Errorf("Second update should be Lighten")
		}
	})
}

// Test ColorDescriptor Darken method
func TestColorDescriptorDarken(t *testing.T) {
	colorHelper := ColorHelper
	baseColor := colorHelper.ColorSelector().SecondaryRoles.Secondary

	t.Run("Darken adds update", func(t *testing.T) {
		darkened := baseColor.Darken(0.4)

		updates := darkened.Updates()
		if len(updates) != 1 {
			t.Fatalf("Should have one update, got %d", len(updates))
		}
		if updates[0].Action() != DarkenColorUpdateAction {
			t.Errorf("Expected DarkenColorUpdateAction, got %d", updates[0].Action())
		}
		if GetDarken(updates[0]) != 0.4 {
			t.Errorf("Expected percentage 0.4, got %f", GetDarken(updates[0]))
		}
	})

	t.Run("Darken then lighten", func(t *testing.T) {
		modified := baseColor.Darken(0.3).Lighten(0.1)

		updates := modified.Updates()
		if len(updates) != 2 {
			t.Fatalf("Should have two updates, got %d", len(updates))
		}
		if updates[0].Action() != DarkenColorUpdateAction {
			t.Errorf("First update should be Darken")
		}
		if updates[1].Action() != LightenColorUpdateAction {
			t.Errorf("Second update should be Lighten")
		}
	})
}

// Test clamp function
func TestClamp(t *testing.T) {
	if result := clamp(0.5, 0.0, 1.0); result != 0.5 {
		t.Errorf("Expected 0.5, got %f", result)
	}
	if result := clamp(-0.5, 0.0, 1.0); result != 0.0 {
		t.Errorf("Expected 0.0, got %f", result)
	}
	if result := clamp(1.5, 0.0, 1.0); result != 1.0 {
		t.Errorf("Expected 1.0, got %f", result)
	}
}

// Test hue rotation
func TestRotateHue(t *testing.T) {
	red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}

	t.Run("Rotate by 120 degrees", func(t *testing.T) {
		result := rotateHue(red, 120.0)
		h, _, _ := rgbaToHSL(result)

		// Red is 0°, rotating by 120° should give green (120°)
		if diff := abs32(h - 120.0); diff > 5.0 {
			t.Errorf("Hue should be approximately 120°, got %f", h)
		}
	})

	t.Run("Rotate by -120 degrees", func(t *testing.T) {
		result := rotateHue(red, -120.0)
		h, _, _ := rgbaToHSL(result)

		// Red is 0°, rotating by -120° should give blue (240°)
		if diff := abs32(h - 240.0); diff > 5.0 {
			t.Errorf("Hue should be approximately 240°, got %f", h)
		}
	})

	t.Run("Rotate by 360 degrees (full circle)", func(t *testing.T) {
		result := rotateHue(red, 360.0)

		// Should be back to red (allow small rounding errors)
		if diff := abs(int(red.R) - int(result.R)); diff > 5 {
			t.Errorf("Red channel should be preserved, expected %d, got %d", red.R, result.R)
		}
	})
}

// Helper functions
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func abs32(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

// Benchmark color transformations
func BenchmarkLightenColor(b *testing.B) {
	c := color.NRGBA{R: 100, G: 150, B: 200, A: 255}
	for i := 0; i < b.N; i++ {
		_ = lightenColor(c, 0.2)
	}
}

func BenchmarkDarkenColor(b *testing.B) {
	c := color.NRGBA{R: 100, G: 150, B: 200, A: 255}
	for i := 0; i < b.N; i++ {
		_ = darkenColor(c, 0.2)
	}
}

func BenchmarkRGBAToHSL(b *testing.B) {
	c := color.NRGBA{R: 100, G: 150, B: 200, A: 255}
	for i := 0; i < b.N; i++ {
		_, _, _ = rgbaToHSL(c)
	}
}

func BenchmarkHSLToRGBA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hslToRGBA(210.0, 0.5, 0.6, 255)
	}
}
