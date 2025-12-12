package main

import (
	"go-compose-dev/compose"
	"go-compose-dev/compose/foundation/icon"
	"go-compose-dev/compose/foundation/layout/column"
	"go-compose-dev/compose/foundation/layout/row"

	"golang.org/x/exp/shiny/materialdesign/icons"

	"go-compose-dev/pkg/api"
)

func UI(c api.Composer) api.LayoutNode {

	c = column.Column(
		compose.Sequence(
			row.Row(
				compose.Sequence(
					icon.Icon(icons.ContentMail),
					icon.Icon(icons.SocialNotifications),
					icon.Icon(icons.ActionAccountCircle),
				),
			),
			row.Row(
				compose.Sequence(
					icon.Icon(icons.ContentMail),
					icon.Icon(icons.SocialNotifications),
					icon.Icon(icons.ActionAccountCircle),
				),
			),
		),
	)(c)

	return c.Build()

}
