package main

import (
	"go-compose-dev/compose"
	"go-compose-dev/compose/foundation/layout/box"
	"go-compose-dev/compose/foundation/layout/column"
	"go-compose-dev/compose/foundation/material3/button"
	"go-compose-dev/compose/foundation/material3/snackbar"
	m3Text "go-compose-dev/compose/foundation/material3/text"
	"go-compose-dev/internal/modifiers/padding"
	"go-compose-dev/internal/modifiers/size"
	"go-compose-dev/pkg/api"
	"log"
)

func UI(c api.Composer) api.LayoutNode {

	// Create snackbar host state outside the loop to persist state
	snackbarHostState := c.State("snackbarHostState", func() any { return snackbar.NewSnackbarHostState() }).Get().(*snackbar.SnackbarHostState)

	// Define UI inline
	return box.Box(
		func(c compose.Composer) compose.Composer {
			// Content
			column.Column(
				func(c compose.Composer) compose.Composer {
					// Headline
					m3Text.Text("Snackbars!", m3Text.TypestyleHeadlineLarge)(c)

					// Body
					m3Text.Text("This is a simple body message, click for debug info.", m3Text.TypestyleBodyLarge)(c)

					// Button
					button.Elevated(func() {
						log.Println("Showing Snackbar")
						snackbarHostState.ShowSnackbar("Hi!")
					}, "Say Hi!",
						button.WithModifier(padding.Vertical(20, 0)),
					)(c)

					return c
				},
				column.WithModifier(padding.All(16)),
			)(c)

			// SnackbarHost overlay
			// Since Box stacks children, this will be drawn on top (last child)
			snackbar.SnackbarHost(snackbarHostState)(c)

			return c
		},
		box.WithModifier(
			size.FillMax(),
		),
	)(c).Build()
}
