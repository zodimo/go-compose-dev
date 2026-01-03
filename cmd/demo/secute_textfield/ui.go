package main

import (
	"gioui.org/unit"
	"github.com/zodimo/go-compose/compose/foundation/layout/column"
	"github.com/zodimo/go-compose/compose/foundation/layout/spacer"
	"github.com/zodimo/go-compose/compose/material3/textfield"
	"github.com/zodimo/go-compose/pkg/api"
)

func UI(c api.Composer) api.LayoutNode {
	password := c.State("password", func() any { return "" })
	passwordVal := password.Get().(string)

	root := column.Column(
		c.Sequence(
			spacer.Height(int(unit.Dp(20))),

			textfield.SecureTextField(
				passwordVal,
				func(newVal string) {
					password.Set(newVal)
				},
				textfield.WithLabel("Password (Filled)"),
			),

			spacer.Height(int(unit.Dp(20))),

			textfield.OutlinedSecureTextField(
				passwordVal,
				func(newVal string) {
					password.Set(newVal)
				},
				textfield.WithLabel("Password (Outlined)"),
			),

			spacer.Height(int(unit.Dp(20))),

			textfield.OutlinedSecureTextField(
				passwordVal,
				func(newVal string) {
					password.Set(newVal)
				},
				textfield.WithLabel("Custom Mask (*)"),
				textfield.WithMask('*'),
			),
		),
	)

	return root(c).Build()
}
