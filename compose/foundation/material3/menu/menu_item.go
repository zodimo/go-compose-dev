package menu

import (
	"go-compose-dev/compose/foundation/layout/row"
	"go-compose-dev/internal/modifiers/clickable"
	"go-compose-dev/internal/modifiers/padding"
	"go-compose-dev/internal/modifiers/size"
)

// DropdownMenuItem Composable
func DropdownMenuItem(
	onClick func(),
	content Composable,
) Composable {
	return func(c Composer) Composer {
		return row.Row(
			content,
			row.WithModifier(clickable.OnClick(onClick)),
			row.WithModifier(padding.Horizontal(16, 16)),
			row.WithModifier(padding.Vertical(8, 8)),
			row.WithModifier(size.FillMaxWidth()),
			row.WithModifier(size.Height(48)),
			row.WithAlignment(row.Middle),
		)(c)
	}
}
