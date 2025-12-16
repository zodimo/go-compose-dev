package spacer

import (
	"github.com/zodimo/go-compose/compose"
	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-compose/modifiers/weight"
)

func Spacer(width, height int) Composable {
	return box.Box(
		func(c Composer) Composer { return c },
		box.WithModifier(
			size.Size(width, height),
		),
	)
}

func Uniform(d int) Composable {
	return box.Box(
		func(c Composer) Composer { return c },
		box.WithModifier(
			size.Size(d, d),
		),
	)
}

func Width(d int) Composable {
	return box.Box(
		func(c Composer) Composer { return c },
		box.WithModifier(
			size.Width(d),
		),
	)
}

func Height(d int) Composable {
	return box.Box(
		func(c Composer) Composer { return c },
		box.WithModifier(
			size.Height(d),
		),
	)
}

func Weight(d int) Composable {
	return box.Box(
		compose.Id(),
		box.WithModifier(
			weight.Weight(1),
		),
	)
}
