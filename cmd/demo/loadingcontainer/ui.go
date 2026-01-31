package main

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	fText "github.com/zodimo/go-compose/compose/foundation/text"
	mButton "github.com/zodimo/go-compose/compose/material3/button"
	"github.com/zodimo/go-compose/compose/material3/progress"
	mwSwitch "github.com/zodimo/go-compose/compose/material3/switch"
	"github.com/zodimo/go-compose/compose/material3/text"
	"github.com/zodimo/go-compose/compose/ui/graphics"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/state"
)

func UI() api.Composable {
	return func(c api.Composer) api.Composer {
		isLoadingState := state.MustState(c, "isLoading", func() bool {
			return false
		})
		return column.Column(
			c.Sequence(
				text.HeadlineMedium("Loading Container Demo"),
				spacer.Height(16),
				mwSwitch.Switch(isLoadingState.Get(), func(b bool) { isLoadingState.Set(b) }),
				spacer.Height(16),
				progress.LoadingContainer(
					isLoadingState.Get(),
					mButton.Filled(
						func() {},
						"load data",
					),
				),
				spacer.Height(16),
				progress.LoadingContainer(
					isLoadingState.Get(),
					mButton.Filled(
						func() {},
						"load data",
					),
					progress.WithLoadingIndicator(
						fText.Text("Loading...", fText.WithColor(graphics.ColorRed)),
					),
				),
				spacer.Height(16),
				progress.LoadingContainer(
					isLoadingState.Get(),
					mButton.Filled(
						func() {},
						"load data",
					),
					progress.WithLoadingIndicator(
						text.TitleLargeEmphasized("Loading..."),
					),
				),
			),
		)(c)
	}
}
