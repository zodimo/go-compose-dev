package selection

// SelectionAdjustment defines how a selection should be adjusted.
//
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/foundation/foundation/src/commonMain/kotlin/androidx/compose/foundation/text/selection/SelectionAdjustment.kt
type SelectionAdjustment int

const (
	// SelectionAdjustmentNone makes no adjustment to the selection.
	SelectionAdjustmentNone SelectionAdjustment = iota

	// SelectionAdjustmentCharacter adjusts selection to the character level.
	SelectionAdjustmentCharacter

	// SelectionAdjustmentWord adjusts selection to the word level.
	SelectionAdjustmentWord

	// SelectionAdjustmentParagraph adjusts selection to the paragraph level.
	SelectionAdjustmentParagraph

	// SelectionAdjustmentCharacterWithWordAccelerate starts with character adjustment
	// and accelerates to word level when dragging fast.
	SelectionAdjustmentCharacterWithWordAccelerate
)

// Adjust applies this adjustment to the selection layout and returns the adjusted selection.
// This is a placeholder - full implementation requires SelectionLayout.
func (adj SelectionAdjustment) Adjust(layout SelectionLayout) *Selection {
	// TODO: Implement full adjustment logic based on SelectionLayout
	return nil
}

// SelectionLayout contains layout information for computing selection adjustments.
// This is a placeholder interface that will be expanded as needed.
//
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/foundation/foundation/src/commonMain/kotlin/androidx/compose/foundation/text/selection/SelectionLayout.kt
type SelectionLayout interface {
	// ShouldRecomputeSelection returns true if selection should be recomputed
	// based on the previous layout.
	ShouldRecomputeSelection(previousLayout SelectionLayout) bool
}
