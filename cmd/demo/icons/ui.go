package main

import (
	"go-compose-dev/compose"
	"go-compose-dev/compose/foundation/icon"
	"go-compose-dev/compose/foundation/layout/column"
	"go-compose-dev/compose/foundation/layout/row"

	"go-compose-dev/pkg/api"
)

func UI(c api.Composer) api.LayoutNode {
	var rows []api.Composable
	chunkSize := 40

	for i := 0; i < len(List); i += chunkSize {
		end := i + chunkSize
		if end > len(List) {
			end = len(List)
		}

		var rowItems []api.Composable
		for _, def := range List[i:end] {
			rowItems = append(rowItems, icon.Icon(def.Data))
		}
		rows = append(rows, row.Row(compose.Sequence(rowItems...)))
	}

	c = column.Column(
		compose.Sequence(rows...),
	)(c)

	return c.Build()
}
