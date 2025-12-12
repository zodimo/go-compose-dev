package surface

import (
	"go-compose-dev/compose/foundation/layout/box"
	"go-compose-dev/internal/modifiers/background"
	"go-compose-dev/internal/modifiers/border"
	"go-compose-dev/internal/modifiers/clip"
	"go-compose-dev/internal/modifiers/shadow"
)

// Surface is a layout composable that represents a Material surface.
// It handles clipping, background color, and elevation.
func Surface(
	content Composable,
	options ...SurfaceOption,
) Composable {
	return func(c Composer) Composer {
		opts := DefaultSurfaceOptions()
		for _, option := range options {
			option(&opts)
		}

		// Apply modifiers: Clip then Background.
		// Shadow should be behind everything.

		surfaceModifier := opts.Modifier.
			Then(shadow.Simple(opts.ShadowElevation, opts.Shape)).
			Then(clip.Clip(opts.Shape)).
			Then(background.Background(opts.Color, func(o *background.BackgroundOptions) { o.Shape = opts.Shape })).
			Then(border.Border(opts.BorderWidth, opts.BorderColor, opts.Shape))

		// Use Box to hold content and apply modifiers
		return box.Box(content, box.WithModifier(surfaceModifier), box.WithAlignment(opts.Alignment))(c)
	}
}
