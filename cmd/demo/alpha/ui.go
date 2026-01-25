package main

import (
	"fmt"

	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/material3/slider"
	"github.com/zodimo/go-compose/compose/material3/text"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/modifiers/alpha"
	"github.com/zodimo/go-compose/modifiers/background"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/state"
)

func UI() api.Composable {
	return func(c api.Composer) api.Composer {
		alphaState := state.MustState(c, "alpha", func() float32 {
			return 1.0
		})

		return column.Column(
			c.Sequence(
				text.HeadlineMedium("Alpha Modifier Demo"),
				spacer.Height(16),
				text.BodyMedium(fmt.Sprintf("Alpha: %.2f", alphaState.Get())),
				slider.Slider(
					alphaState.Get(),
					func(val float32) { alphaState.Set(val) },
				),
				spacer.Height(32),
				box.Box(
					func(c api.Composer) api.Composer { return c },
					box.WithModifier(
						alpha.Alpha(alphaState.Get()).
							Then(background.Background(graphics.ColorRed)).
							Then(size.Size(200, 200)),
					),
				),
			),
			column.WithSpacing(column.SpaceSides),
			column.WithAlignment(column.Middle),
			column.WithModifier(size.FillMax()),
		)(c)
	}
}
