package store

import (
	"reflect"

	"github.com/zodimo/go-compose/state"
)

type Versionable interface {
	Version() int64
}

type DerivedState[T any] struct {
	calculation func() T
	value       T
	deps        map[any]int64
	initialized bool
	version     int64
	compare     func(T, T) bool
}

func DerivedStateOf[T any](calculation func() T) *DerivedState[T] {
	return &DerivedState[T]{
		calculation: calculation,
		deps:        make(map[any]int64),
		compare:     func(a, b T) bool { return reflect.DeepEqual(a, b) }, // Default deep equal
	}
}

// DerivedStateOfCustom allows providing a custom comparison function
func DerivedStateOfCustom[T any](calculation func() T, compare func(T, T) bool) *DerivedState[T] {
	return &DerivedState[T]{
		calculation: calculation,
		deps:        make(map[any]int64),
		compare:     compare,
	}
}

func (ds *DerivedState[T]) Get() any {
	// Check if we need to recalculate
	if !ds.initialized || ds.isStale() {
		ds.recalculate()
	}

	state.NotifyRead(ds)
	return ds.value
}

func (ds *DerivedState[T]) isStale() bool {
	for dep, savedVersion := range ds.deps {
		if v, ok := dep.(Versionable); ok {
			if v.Version() != savedVersion {
				return true
			}
		}
	}
	return false
}

func (ds *DerivedState[T]) recalculate() {
	newDeps := make(map[any]int64)

	var newValue T
	state.WithReadObserver(func(source any) {
		if v, ok := source.(Versionable); ok {
			newDeps[source] = v.Version()
		}
	}, func() {
		newValue = ds.calculation()
	})

	// Only increment version if the value actually changed
	if !ds.initialized || !ds.compare(ds.value, newValue) {
		ds.version++
		ds.value = newValue
	}

	ds.deps = newDeps
	ds.initialized = true
}

func (ds *DerivedState[T]) Version() int64 {
	// Ensure up-to-date before returning version, mainly for nested derived states
	if !ds.initialized || ds.isStale() {
		ds.recalculate()
	}
	return ds.version
}
