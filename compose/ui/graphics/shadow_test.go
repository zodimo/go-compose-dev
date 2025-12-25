package graphics

import (
	"testing"

	"github.com/zodimo/go-compose/compose/ui/geometry"
)

func TestShadow_SentinelPattern(t *testing.T) {
	// Test EmptyShadow singleton
	if ShadowUnspecified == nil {
		t.Error("EmptyShadow should not be nil")
	}
	if IsShadow(ShadowUnspecified) {
		t.Error("EmptyShadow should not be specified")
	}

	// Test TakeOrElse
	t.Run("TakeOrElse", func(t *testing.T) {
		s1 := NewShadow(ColorBlack, geometry.NewOffset(1, 1), 10)
		defaultShadow := NewShadow(ColorRed, geometry.NewOffset(2, 2), 5)

		// Case 1: s is specified
		res := TakeOrElseShadow(s1, defaultShadow)
		if res != s1 {
			t.Error("TakeOrElse should return s when s is specified")
		}

		// Case 2: s is nil
		res = TakeOrElseShadow(nil, defaultShadow)
		if res != defaultShadow {
			t.Error("TakeOrElse should return default when s is nil")
		}

		// Case 3: s is EmptyShadow
		res = TakeOrElseShadow(ShadowUnspecified, defaultShadow)
		if res != defaultShadow {
			t.Error("TakeOrElse should return default when s is EmptyShadow")
		}
	})

	// Test String()
	t.Run("String", func(t *testing.T) {
		if StringShadow(ShadowUnspecified) != "EmptyShadow" {
			t.Errorf("Expected 'EmptyShadow', got '%s'", StringShadow(ShadowUnspecified))
		}
	})
}
