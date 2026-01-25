package progress

import (
	"image"
	"testing"

	"gioui.org/layout"
)

func TestCalculateLoadingDiameter(t *testing.T) {
	defaultDiameter := 48

	// Case 1: Default size (48dp) with loose constraints
	constraints := layout.Constraints{
		Min: kIntPoint(0, 0),
		Max: kIntPoint(1000, 1000),
	}
	d := calculateDiameter(constraints, defaultDiameter)
	if d != 48 {
		t.Errorf("Expected default size 48, got %d", d)
	}

	// Case 2: Explicit size 200x200 (Min=200, Max=200)
	constraints = layout.Constraints{
		Min: kIntPoint(200, 200),
		Max: kIntPoint(200, 200),
	}
	d = calculateDiameter(constraints, defaultDiameter)
	if d != 200 {
		t.Errorf("Expected size 200, got %d", d)
	}

	// Case 3: Constrained Max (Max=20, Min=0)
	constraints = layout.Constraints{
		Min: kIntPoint(0, 0),
		Max: kIntPoint(20, 20),
	}
	d = calculateDiameter(constraints, defaultDiameter)
	if d != 20 {
		t.Errorf("Expected constrained size 20, got %d", d)
	}

	// Case 4: Min > Default (Min=100, Max=1000) -> Should grow to Min
	constraints = layout.Constraints{
		Min: kIntPoint(100, 100),
		Max: kIntPoint(1000, 1000),
	}
	d = calculateDiameter(constraints, defaultDiameter)
	if d != 100 {
		t.Errorf("Expected min size 100, got %d", d)
	}

	// Case 5: Min width 200, Min height 50. Max width 200, Max height 50.
	// Should pick 200 then clamp to 50 due to max height?
	// Trace: minW=200 (>48) -> d=200. minH=50 (>48) -> d=200.
	// MaxW=200 -> d=200. MaxH=50 -> d=50.
	// Result 50.
	constraints = layout.Constraints{
		Min: kIntPoint(200, 50),
		Max: kIntPoint(200, 50),
	}
	d = calculateDiameter(constraints, defaultDiameter)
	if d != 50 {
		t.Errorf("Expected constrained size 50, got %d", d)
	}

	// Case 6: Min width 50, Min height 200.
	// Trace: minW=50 -> d=50. minH=200 -> d=200.
	// MaxW=50 -> d=50. MaxH=200 -> d=50.
	// Result 50.
	constraints = layout.Constraints{
		Min: kIntPoint(50, 200),
		Max: kIntPoint(50, 200),
	}
	d = calculateDiameter(constraints, defaultDiameter)
	if d != 50 {
		t.Errorf("Expected constrained size 50, got %d", d)
	}
}

func kIntPoint(x, y int) (p image.Point) {
	p.X = x
	p.Y = y
	return
}
