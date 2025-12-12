package spacer

import (
	"go-compose-dev/compose/foundation/layout/box"
	"go-compose-dev/internal/modifiers/size"
	"go-compose-dev/pkg/api"
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
