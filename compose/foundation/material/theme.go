package material

import (
	"sync"

	"gioui.org/font/gofont"
	"gioui.org/text"
	"gioui.org/widget/material"
)

var materialThemeSingleton *material.Theme
var materialThemeMutex sync.Mutex

func defaultTheme() *material.Theme {
	if materialThemeSingleton == nil {
		materialThemeSingleton = material.NewTheme()
		materialThemeSingleton.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	}
	return materialThemeSingleton
}

func SetTheme(theme *material.Theme) {
	materialThemeMutex.Lock()
	defer materialThemeMutex.Unlock()
	materialThemeSingleton = theme
}

func GetTheme() *material.Theme {
	materialThemeMutex.Lock()
	defer materialThemeMutex.Unlock()
	return materialThemeSingleton
}

func init() {
	materialThemeSingleton = defaultTheme()
}
