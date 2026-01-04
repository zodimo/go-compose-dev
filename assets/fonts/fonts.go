// Package fonts provides embedded fonts for the go-compose framework.
// This includes both the Go fonts (for text) and Noto Color Emoji (for emoji).
package fonts

import (
	"embed"
	"sync"

	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
)

//go:embed NotoColorEmoji.ttf
var emojiFont embed.FS

var (
	cachedCollection []font.FontFace
	collectionOnce   sync.Once
)

// Collection returns a font collection that includes both Go fonts and Noto Color Emoji.
// This provides full unicode and emoji support in a portable, self-contained way
// without relying on system fonts.
func Collection() []font.FontFace {
	collectionOnce.Do(func() {
		// Start with the Go font collection
		cachedCollection = gofont.Collection()

		// Load and add the emoji font
		emojiData, err := emojiFont.ReadFile("NotoColorEmoji.ttf")
		if err != nil {
			// If emoji font fails to load, keep just the Go fonts
			return
		}

		// ParseCollection returns []font.FontFace directly
		emojiFaces, err := opentype.ParseCollection(emojiData)
		if err != nil {
			return
		}

		// Append emoji faces to the collection
		cachedCollection = append(cachedCollection, emojiFaces...)
	})

	return cachedCollection
}

// GoFontsOnly returns just the Go font collection without emoji.
// Use this if you don't need emoji support and want smaller binary size.
func GoFontsOnly() []font.FontFace {
	return gofont.Collection()
}
