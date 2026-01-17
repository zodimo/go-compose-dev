package state_test

import (
	"testing"

	"github.com/zodimo/go-compose/state"
)

func TestMutationPolicy_Structural(t *testing.T) {
	policy := state.StructuralEqualityPolicy[[]int]()

	t.Run("equivalent slices are equal", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		if !policy.Equivalent(a, b) {
			t.Error("Equivalent slices should be equal")
		}
	})

	t.Run("different slices are not equal", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 4}
		if policy.Equivalent(a, b) {
			t.Error("Different slices should not be equal")
		}
	})

	t.Run("merge returns false", func(t *testing.T) {
		a := []int{1, 2, 3}
		_, ok := policy.Merge(a, a, a)
		if ok {
			t.Error("Structural policy should not support merge")
		}
	})
}

func TestMutationPolicy_Referential(t *testing.T) {
	policy := state.ReferentialEqualityPolicy[int]()

	t.Run("same values are equal", func(t *testing.T) {
		if !policy.Equivalent(42, 42) {
			t.Error("Same int values should be equal")
		}
	})

	t.Run("different values are not equal", func(t *testing.T) {
		if policy.Equivalent(42, 43) {
			t.Error("Different int values should not be equal")
		}
	})

	t.Run("merge returns false", func(t *testing.T) {
		_, ok := policy.Merge(1, 2, 3)
		if ok {
			t.Error("Referential policy should not support merge")
		}
	})
}

func TestMutationPolicy_NeverEqual(t *testing.T) {
	policy := state.NeverEqualPolicy[int]()

	t.Run("same values are not equal", func(t *testing.T) {
		if policy.Equivalent(42, 42) {
			t.Error("NeverEqual policy should never return true")
		}
	})

	t.Run("merge returns false", func(t *testing.T) {
		_, ok := policy.Merge(1, 2, 3)
		if ok {
			t.Error("NeverEqual policy should not support merge")
		}
	})
}

func TestMutationPolicy_Custom(t *testing.T) {
	type Item struct {
		ID   int
		Name string
	}

	// Custom policy: compare only by ID
	policy := state.NewMutationPolicy(
		func(a, b Item) bool { return a.ID == b.ID },
		nil, // no merge support
	)

	t.Run("same ID different name are equal", func(t *testing.T) {
		a := Item{ID: 1, Name: "Alice"}
		b := Item{ID: 1, Name: "Bob"}
		if !policy.Equivalent(a, b) {
			t.Error("Items with same ID should be equal")
		}
	})

	t.Run("different ID are not equal", func(t *testing.T) {
		a := Item{ID: 1, Name: "Alice"}
		b := Item{ID: 2, Name: "Alice"}
		if policy.Equivalent(a, b) {
			t.Error("Items with different ID should not be equal")
		}
	})
}

func TestMutationPolicy_WithMerge(t *testing.T) {
	// Counter policy: merge by summing deltas
	policy := state.NewMutationPolicy(
		func(a, b int) bool { return a == b },
		func(previous, current, applied int) (int, bool) {
			// Merge: add the delta from applied to current
			delta := applied - previous
			return current + delta, true
		},
	)

	t.Run("merge succeeds", func(t *testing.T) {
		// previous=10, current=15 (added 5), applied=12 (added 2)
		// merged should be 15 + 2 = 17
		merged, ok := policy.Merge(10, 15, 12)
		if !ok {
			t.Error("Merge should succeed")
		}
		if merged != 17 {
			t.Errorf("Expected merged value 17, got %d", merged)
		}
	})
}
