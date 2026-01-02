package button

import (
	"github.com/zodimo/go-compose/compose/foundation/text"
	"github.com/zodimo/go-compose/compose/material3"
)

const Material3ButtonNodeID = "Material3Button"

func Outlined(onClick func(), label string, options ...ButtonOption) Composable {
	return OutlinedButton(onClick, func(c Composer) Composer {
		style := material3.LocalTextStyle.Current(c)
		return text.Text(label, text.WithTextStyle(style))(c)
	}, options...)
}
