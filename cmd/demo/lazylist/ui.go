package main

import (
	"fmt"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/lazy"
	"github.com/zodimo/go-compose/compose/foundation/material3/text"
	ftext "github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/internal/modifiers/padding"
	"github.com/zodimo/go-compose/internal/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
)

func UI() api.Composable {
	return func(c api.Composer) api.Composer {
		state := lazy.RememberLazyListState(c)

		return column.Column(
			func(c api.Composer) api.Composer {
				lazy.LazyColumn(
					func(scope lazy.LazyListScope) {
						// Header
						scope.Item(nil, func(c api.Composer) api.Composer {
							text.Text("Lazy List Demo", text.TypestyleHeadlineMedium, ftext.WithModifier(padding.All(16)))(c)
							return c
						})

						// 100 items
						scope.Items(100, nil, func(index int) api.Composable {
							return func(c api.Composer) api.Composer {
								text.Text(fmt.Sprintf("Item %d", index), text.TypestyleBodyLarge, ftext.WithModifier(padding.All(8)))(c)
								return c
							}
						})
					},
					lazy.WithModifier(size.FillMax()),
					lazy.WithState(state),
				)(c)
				return c
			},
			column.WithModifier(size.FillMax()),
		)(c)
	}
}
