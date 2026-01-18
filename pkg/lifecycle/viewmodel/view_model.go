// Package viewmodel provides helpers for creating ViewModels in go-compose.
package viewmodel

import (
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/state"
)

// ViewModel creates or retrieves a ViewModel that persists across recompositions.
// The factory is only called once when the ViewModel is first created.
// When the component leaves composition (not accessed during a frame),
// the ViewModel's OnCleared() method is called if it implements lifecycle.ViewModel.
//
// Usage:
//
//	vm := viewmodel.ViewModel(c, "myScreen", func() *MyViewModel { return NewMyViewModel() })
func ViewModel[T any](c api.Composer, key string, factory func() T) T {
	mv := state.StateUnsafe[T](c, "viewmodel_"+key, factory)
	return mv.Get()
}
