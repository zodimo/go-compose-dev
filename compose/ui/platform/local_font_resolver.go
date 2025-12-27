package platform

import (
	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/compose/ui/text/font"
)

// LocalFontFamilyResolver is a CompositionLocal that provides the FontFamilyResolver to the composition.
var LocalFontFamilyResolver = compose.StaticCompositionLocalOf[font.FontFamilyResolver](func() font.FontFamilyResolver {
	panic("CompositionLocal LocalFontFamilyResolver not present")
})
