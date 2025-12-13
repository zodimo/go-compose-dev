package text

import (
	"go-compose-dev/compose/foundation/text"
	"go-compose-dev/internal/layoutnode"

	"git.sr.ht/~schnwalter/gio-mw/token"
	"git.sr.ht/~schnwalter/gio-mw/wdk"
)

const Material3TextNodeID = "Material3Text"

// Text displays text with the given style from the Material Theme.
// It retrieves the theme from the layout context at runtime.
func Text(value string, style Typestyle, options ...text.TextOption) Composable {
	return func(c Composer) Composer {
		// Resolve options
		opts := text.DefaultTextOptions()
		for _, option := range options {
			if option == nil {
				continue
			}
			option(&opts)
		}

		c.StartBlock(Material3TextNodeID)
		c.Modifier(func(modifier Modifier) Modifier {
			return modifier.Then(opts.Modifier)
		})

		constructorArgs := TextConstructorArgs{
			Value: value,
			Style: style,
			Opts:  opts,
		}

		c.SetWidgetConstructor(textWidgetConstructor(constructorArgs))
		return c.EndBlock()
	}
}

type TextConstructorArgs struct {
	Value string
	Style Typestyle
	Opts  text.TextOptions
}

func textWidgetConstructor(args TextConstructorArgs) layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {

			// Map options to LabelStyle
			labelStyle := wdk.LabelStyle{
				Typestyle:  args.Style,
				Alignment:  args.Opts.Alignment,
				MaxLines:   args.Opts.MaxLines,
				WrapPolicy: args.Opts.WrapPolicy,
				// Color is tricky: foundation text options uses NRGBA, wdk uses token.MatColor.
				// If the user specified a color, use it. Otherwise wdk.LayoutLabel defaults to theme color.
			}

			if args.Opts.TextStyleOptions != nil && args.Opts.TextStyleOptions.Color.A > 0 {
				labelStyle.Color = token.MatColor(args.Opts.TextStyleOptions.Color)
			}

			return wdk.LayoutLabel(gtx, labelStyle, args.Value)
		}
	})
}
