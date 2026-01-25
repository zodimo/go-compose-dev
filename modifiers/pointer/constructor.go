package pointer

import (
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/internal/modifier"
)

func BlockPointer() ui.Modifier {
	return modifier.NewModifier(InputBlockerElement{})
}
