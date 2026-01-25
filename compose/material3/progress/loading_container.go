package progress

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/ui"
	"github.com/zodimo/go-compose/modifiers/alpha"
	boxModifier "github.com/zodimo/go-compose/modifiers/box"
	"github.com/zodimo/go-compose/modifiers/pointer"
	"github.com/zodimo/go-compose/modifiers/size"
	"github.com/zodimo/go-ternary"
)

// LoadingContainer wraps content with a loading overlay when isLoading is true.
// The overlay blocks input and displays a loading indicator.
func LoadingContainer(
	isLoading bool,
	content Composable,
	options ...LoadingContainerOption,
) Composable {
	opts := DefaultLoadingContainerOptions()
	for _, opt := range options {
		if opt != nil {
			opt(&opts)
		}
	}

	return func(c Composer) Composer {
		return box.Box(
			c.Sequence(
				box.Box(
					content,
					box.WithModifier(
						ternary.Ternary(
							isLoading,
							alpha.Alpha(0.6),
							ui.EmptyModifier,
						),
					),
				),

				c.When(
					isLoading,
					c.If(
						opts.LoadingIndicator != nil,
						box.Box(
							opts.LoadingIndicator,
							box.WithAlignment(box.Center),
							box.WithModifier(boxModifier.MatchParentSize()),
						),
						LoadingIndicator(
							WithModifier(boxModifier.MatchParentSize()),
						),
					),
				),
			),
			box.WithModifier(size.WrapContentSize().
				Then(
					ternary.Ternary(
						isLoading,
						pointer.BlockPointer(),
						ui.EmptyModifier,
					),
				),
			),
		)(c)
	}
}
