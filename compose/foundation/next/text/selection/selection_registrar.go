package selection

import (
	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/layout"
)

// InvalidSelectableId represents an invalid ID for Selectable.
const InvalidSelectableId int64 = 0

// SelectionRegistrar allows a composable to subscribe and unsubscribe to selection changes.
//
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/foundation/foundation/src/commonMain/kotlin/androidx/compose/foundation/text/selection/SelectionRegistrar.kt
type SelectionRegistrar interface {
	// Subselections returns the map storing current selection information on each Selectable.
	// A selectable can query its selected range using its selectableId.
	Subselections() map[int64]*Selection

	// Subscribe registers a Selectable to this SelectionRegistrar.
	// Returns the Selectable for use with Unsubscribe.
	Subscribe(selectable Selectable) Selectable

	// Unsubscribe removes a Selectable from this SelectionRegistrar.
	Unsubscribe(selectable Selectable)

	// NextSelectableId returns a unique ID for a Selectable.
	NextSelectableId() int64

	// NotifyPositionChange is called when the global position of a subscribed Selectable changes.
	NotifyPositionChange(selectableId int64)

	// NotifySelectionUpdateStart is called when selection has been initiated.
	// layoutCoordinates: LayoutCoordinates of the Selectable.
	// startPosition: coordinates where selection is initiated.
	// adjustment: how selection should be adjusted.
	// isInTouchMode: whether the update is from a touch pointer.
	NotifySelectionUpdateStart(
		layoutCoordinates layout.LayoutCoordinates,
		startPosition geometry.Offset,
		adjustment SelectionAdjustment,
		isInTouchMode bool,
	)

	// NotifySelectionUpdateSelectAll is called when selection is initiated with selectAll.
	// selectableId: the selectableId of the Selectable.
	// isInTouchMode: whether the update is from a touch pointer.
	NotifySelectionUpdateSelectAll(selectableId int64, isInTouchMode bool)

	// NotifySelectionUpdate is called when a selection handle has moved.
	// Returns true if the selection handle movement is consumed.
	NotifySelectionUpdate(
		layoutCoordinates layout.LayoutCoordinates,
		newPosition geometry.Offset,
		previousPosition geometry.Offset,
		isStartHandle bool,
		adjustment SelectionAdjustment,
		isInTouchMode bool,
	) bool

	// NotifySelectionUpdateEnd is called when selection update has stopped.
	NotifySelectionUpdateEnd()

	// NotifySelectableChange is called when the content of a selectable has changed.
	NotifySelectableChange(selectableId int64)
}

// HasSelection checks if there is a selection on the CoreText with the given selectableId.
func HasSelection(registrar SelectionRegistrar, selectableId int64) bool {
	if registrar == nil {
		return false
	}
	_, exists := registrar.Subselections()[selectableId]
	return exists
}

// LocalSelectionRegistrar is a CompositionLocal for SelectionRegistrar.
// Composables that implement selection logic can use this to get a SelectionRegistrar
// in order to subscribe and unsubscribe to SelectionRegistrar.
//
// The default value is nil, meaning selection is not enabled.
var LocalSelectionRegistrar = compose.CompositionLocalOf[SelectionRegistrar](func() SelectionRegistrar {
	return nil
})
