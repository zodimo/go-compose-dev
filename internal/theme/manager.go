package theme

import (
	"sync"

	"gioui.org/widget/material"
)

var themeManagerSingleton ThemeManager

type ThemeManager interface {
	MaterialTheme() *material.Theme
	SetMaterialTheme(theme *material.Theme)
}

var _ ThemeManager = (*themeManager)(nil)

type themeManager struct {
	mu            sync.RWMutex
	materialTheme *material.Theme
}

func newThemeManager(materialTheme *material.Theme) ThemeManager {
	return &themeManager{
		materialTheme: materialTheme,
	}
}

func (tm *themeManager) MaterialTheme() *material.Theme {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return tm.materialTheme

}
func (tm *themeManager) SetMaterialTheme(theme *material.Theme) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.materialTheme = theme
}

func GetThemeManager() ThemeManager {
	return themeManagerSingleton
}

func init() {
	themeManagerSingleton = newThemeManager(defaultMaterialTheme())

}
