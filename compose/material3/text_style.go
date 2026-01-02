package material3

import (
	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/compose/ui/text"
)

var LocalTextStyle = compose.CompositionLocalOf(func() *text.TextStyle {
	return nil
})
