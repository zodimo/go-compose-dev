package main

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	m3text "github.com/zodimo/go-compose/compose/material3/text"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/pkg/api"
)

// TypographyScreen shows typography
func TypographyScreen(c api.Composer) api.Composable {
	return func(c api.Composer) api.Composer {
		return column.Column(
			c.Sequence(
				SectionTitle("Typography"),
				spacer.Height(8),
				m3text.LabelSmall("LabelSmall"),
				m3text.LabelMedium("LabelMedium"),
				m3text.LabelLarge("LabelLarge"),
				m3text.BodySmall("BodySmall"),
				m3text.BodyMedium("BodyMedium"),
				m3text.BodyLarge("BodyLarge"),
				m3text.TitleSmall("TitleSmall"),
				m3text.TitleMedium("TitleMedium"),
				m3text.TitleLarge("TitleLarge"),
				m3text.HeadlineSmall("HeadlineSmall"),
				m3text.HeadlineMedium("HeadlineMedium"),
				m3text.HeadlineLarge("HeadlineLarge"),
				m3text.DisplaySmall("DisplaySmall"),
				m3text.DisplayMedium("DisplayMedium"),
				m3text.DisplayLarge("DisplayLarge"),

				m3text.LabelSmallEmphasized("LabelSmallEmphasized"),
				m3text.LabelMediumEmphasized("LabelMediumEmphasized"),
				m3text.LabelLargeEmphasized("LabelLargeEmphasized"),
				m3text.BodySmallEmphasized("BodySmallEmphasized"),
				m3text.BodyMediumEmphasized("BodyMediumEmphasized"),
				m3text.BodyLargeEmphasized("BodyLargeEmphasized"),
				m3text.TitleSmallEmphasized("TitleSmallEmphasized"),
				m3text.TitleMediumEmphasized("TitleMediumEmphasized"),
				m3text.TitleLargeEmphasized("TitleLargeEmphasized"),
				m3text.HeadlineSmallEmphasized("HeadlineSmallEmphasized"),
				m3text.HeadlineMediumEmphasized("HeadlineMediumEmphasized"),
				m3text.HeadlineLargeEmphasized("HeadlineLargeEmphasized"),
				m3text.DisplaySmallEmphasized("DisplaySmallEmphasized"),
				m3text.DisplayMediumEmphasized("DisplayMediumEmphasized"),
				m3text.DisplayLargeEmphasized("DisplayLargeEmphasized"),
			),
			column.WithModifier(padding.All(16)),
		)(c)
	}
}
