package divider

import (
	"github.com/zodimo/go-compose/theme"
)

func DefaultDividerOptions() DividerOptions {
	return DividerOptions{
		Modifier:  EmptyModifier,
		Thickness: 1,
		Color:     theme.ColorHelper.ColorSelector().OutlineRoles.OutlineVariant,
	}
}
