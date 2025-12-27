package selection

import (
	"sync"
	"sync/atomic"

	"github.com/zodimo/go-compose/compose/ui/geometry"
	"github.com/zodimo/go-compose/compose/ui/layout"
)

// SelectionRegistrarImpl is the implementation of SelectionRegistrar.
//
// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/foundation/foundation/src/commonMain/kotlin/androidx/compose/foundation/text/selection/SelectionRegistrarImpl.kt
type SelectionRegistrarImpl struct {
	mu sync.RWMutex

	// sorted indicates if the selectables have already been sorted.
	sorted bool

	// selectables is the list of registered components that want to handle text selection.
	selectables []Selectable

	// selectableMap is a map from selectable keys to subscribed selectables.
	selectableMap map[int64]Selectable

	// incrementId is the incremental id to be assigned to each selectable.
	// It starts from 1 and 0 is used to denote an invalid id.
	incrementId atomic.Int64

	// subselections stores the current selection information on each Selectable.
	subselections map[int64]*Selection

	// Callbacks

	// OnPositionChangeCallback is called when a position change was triggered.
	OnPositionChangeCallback func(selectableId int64)

	// OnSelectionUpdateStartCallback is called when selection is initiated.
	OnSelectionUpdateStartCallback func(isInTouchMode bool, layoutCoordinates layout.LayoutCoordinates, position geometry.Offset, adjustment SelectionAdjustment)

	// OnSelectionUpdateSelectAllCallback is called when selection is initiated with selectAll.
	OnSelectionUpdateSelectAllCallback func(isInTouchMode bool, selectableId int64)

	// OnSelectionUpdateCallback is called when selection is updated.
	OnSelectionUpdateCallback func(isInTouchMode bool, layoutCoordinates layout.LayoutCoordinates, newPosition, previousPosition geometry.Offset, isStartHandle bool, adjustment SelectionAdjustment) bool

	// OnSelectionUpdateEndCallback is called when selection update finished.
	OnSelectionUpdateEndCallback func()

	// OnSelectableChangeCallback is called when one of the selectable has changed.
	OnSelectableChangeCallback func(selectableId int64)

	// AfterSelectableUnsubscribe is called after a selectable is unsubscribed.
	AfterSelectableUnsubscribe func(selectableId int64)
}

// NewSelectionRegistrarImpl creates a new SelectionRegistrarImpl.
func NewSelectionRegistrarImpl() *SelectionRegistrarImpl {
	impl := &SelectionRegistrarImpl{
		selectables:   make([]Selectable, 0),
		selectableMap: make(map[int64]Selectable),
		subselections: make(map[int64]*Selection),
	}
	impl.incrementId.Store(1)
	return impl
}

// NewSelectionRegistrarImplWithInitialId creates a new SelectionRegistrarImpl with an initial ID.
// This is used for state restoration.
func NewSelectionRegistrarImplWithInitialId(initialId int64) *SelectionRegistrarImpl {
	impl := NewSelectionRegistrarImpl()
	impl.incrementId.Store(initialId)
	return impl
}

// Subselections implements SelectionRegistrar.
func (r *SelectionRegistrarImpl) Subselections() map[int64]*Selection {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.subselections
}

// SetSubselections updates the subselections map.
func (r *SelectionRegistrarImpl) SetSubselections(subs map[int64]*Selection) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.subselections = subs
}

// Subscribe implements SelectionRegistrar.
func (r *SelectionRegistrarImpl) Subscribe(selectable Selectable) Selectable {
	id := selectable.SelectableId()
	if id == InvalidSelectableId {
		panic("The selectable contains an invalid id")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.selectableMap[id]; exists {
		panic("Another selectable with the same id has already subscribed")
	}

	r.selectableMap[id] = selectable
	r.selectables = append(r.selectables, selectable)
	r.sorted = false
	return selectable
}

// Unsubscribe implements SelectionRegistrar.
func (r *SelectionRegistrarImpl) Unsubscribe(selectable Selectable) {
	id := selectable.SelectableId()

	r.mu.Lock()
	if _, exists := r.selectableMap[id]; !exists {
		r.mu.Unlock()
		return
	}

	// Remove from slice
	for i, s := range r.selectables {
		if s.SelectableId() == id {
			r.selectables = append(r.selectables[:i], r.selectables[i+1:]...)
			break
		}
	}
	delete(r.selectableMap, id)
	r.mu.Unlock()

	if r.AfterSelectableUnsubscribe != nil {
		r.AfterSelectableUnsubscribe(id)
	}
}

// NextSelectableId implements SelectionRegistrar.
func (r *SelectionRegistrarImpl) NextSelectableId() int64 {
	for {
		id := r.incrementId.Add(1) - 1
		if id != InvalidSelectableId {
			return id
		}
	}
}

// Sort sorts the list of registered Selectables.
// Currently the order is geometric-based (y-coordinate first, then x-coordinate).
func (r *SelectionRegistrarImpl) Sort(containerLayoutCoordinates layout.LayoutCoordinates) []Selectable {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.sorted {
		return r.selectables
	}

	// Sort selectables by y-coordinate first, then x-coordinate
	// to match English hand-writing habit
	// TODO: Implement proper sorting when LayoutCoordinates is fully available
	r.sorted = true
	return r.selectables
}

// Selectables returns the list of registered selectables.
func (r *SelectionRegistrarImpl) Selectables() []Selectable {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.selectables
}

// SelectableMap returns the map of selectables.
func (r *SelectionRegistrarImpl) SelectableMap() map[int64]Selectable {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.selectableMap
}

// NotifyPositionChange implements SelectionRegistrar.
func (r *SelectionRegistrarImpl) NotifyPositionChange(selectableId int64) {
	r.mu.Lock()
	r.sorted = false
	r.mu.Unlock()

	if r.OnPositionChangeCallback != nil {
		r.OnPositionChangeCallback(selectableId)
	}
}

// NotifySelectionUpdateStart implements SelectionRegistrar.
func (r *SelectionRegistrarImpl) NotifySelectionUpdateStart(
	layoutCoordinates layout.LayoutCoordinates,
	startPosition geometry.Offset,
	adjustment SelectionAdjustment,
	isInTouchMode bool,
) {
	if r.OnSelectionUpdateStartCallback != nil {
		r.OnSelectionUpdateStartCallback(isInTouchMode, layoutCoordinates, startPosition, adjustment)
	}
}

// NotifySelectionUpdateSelectAll implements SelectionRegistrar.
func (r *SelectionRegistrarImpl) NotifySelectionUpdateSelectAll(selectableId int64, isInTouchMode bool) {
	if r.OnSelectionUpdateSelectAllCallback != nil {
		r.OnSelectionUpdateSelectAllCallback(isInTouchMode, selectableId)
	}
}

// NotifySelectionUpdate implements SelectionRegistrar.
func (r *SelectionRegistrarImpl) NotifySelectionUpdate(
	layoutCoordinates layout.LayoutCoordinates,
	newPosition geometry.Offset,
	previousPosition geometry.Offset,
	isStartHandle bool,
	adjustment SelectionAdjustment,
	isInTouchMode bool,
) bool {
	if r.OnSelectionUpdateCallback != nil {
		return r.OnSelectionUpdateCallback(isInTouchMode, layoutCoordinates, newPosition, previousPosition, isStartHandle, adjustment)
	}
	return true
}

// NotifySelectionUpdateEnd implements SelectionRegistrar.
func (r *SelectionRegistrarImpl) NotifySelectionUpdateEnd() {
	if r.OnSelectionUpdateEndCallback != nil {
		r.OnSelectionUpdateEndCallback()
	}
}

// NotifySelectableChange implements SelectionRegistrar.
func (r *SelectionRegistrarImpl) NotifySelectableChange(selectableId int64) {
	if r.OnSelectableChangeCallback != nil {
		r.OnSelectableChangeCallback(selectableId)
	}
}

// GetIncrementId returns the current increment ID value.
// This is useful for state saving.
func (r *SelectionRegistrarImpl) GetIncrementId() int64 {
	return r.incrementId.Load()
}

// Verify SelectionRegistrarImpl implements SelectionRegistrar at compile time.
var _ SelectionRegistrar = (*SelectionRegistrarImpl)(nil)
