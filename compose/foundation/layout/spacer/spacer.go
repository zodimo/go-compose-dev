package spacer

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/internal/modifiers/size"
	"github.com/zodimo/go-compose/pkg/api"
)

func Spacer(d int) api.Composable {
	return func(c api.Composer) api.Composer {
		return box.Box(
			func(c api.Composer) api.Composer { return c },
			box.WithModifier(
				size.Size(d, d),
			),
		)(c)
	}
}

func SpacerWidth(d int) api.Composable {
	return func(c api.Composer) api.Composer {
		return box.Box(
			func(c api.Composer) api.Composer { return c },
			box.WithModifier(
				size.Width(d),
			),
		)(c)
	}
}

func SpacerHeight(d int) api.Composable {
	return func(c api.Composer) api.Composer {
		return box.Box(
			func(c api.Composer) api.Composer { return c },
			box.WithModifier(
				size.Height(d),
			),
		)(c)
	}
}
