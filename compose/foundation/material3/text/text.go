package text

import (
	"github.com/zodimo/go-compose/compose/foundation/text"

	"github.com/zodimo/go-compose/internal/layoutnode"
	"github.com/zodimo/go-compose/internal/modifier"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/theme"

	"git.sr.ht/~schnwalter/gio-mw/token"
	"git.sr.ht/~schnwalter/gio-mw/wdk"
)

// Text displays text with the given style from the Material Theme.
// It retrieves the theme from the layout context at runtime.
func Text(value string, style Typestyle, options ...text.TextOption) api.Composable {
	return func(c api.Composer) api.Composer {
		// Resolve options
		opts := text.DefaultTextOptions()
		for _, option := range options {
			if option == nil {
				continue
			}
			option(&opts)
		}

		c.StartBlock("Material3Text")
		c.Modifier(func(modifier modifier.Modifier) modifier.Modifier {
			return modifier.Then(opts.Modifier)
		})

		c.SetWidgetConstructor(layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
			return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {

				// Map options to LabelStyle
				labelStyle := wdk.LabelStyle{
					Typestyle:  style,
					Alignment:  opts.Alignment,
					MaxLines:   opts.MaxLines,
					WrapPolicy: opts.WrapPolicy,
				}

				// Resolve ColorDescriptor to NRGBA, then to MatColor
				if opts.TextStyleOptions != nil && opts.TextStyleOptions.Color != nil {
					tm := theme.GetThemeManager()
					resolvedColor := tm.ResolveColorDescriptor(opts.TextStyleOptions.Color).AsNRGBA()
					labelStyle.Color = token.MatColor(resolvedColor)
				}

				return wdk.LayoutLabel(gtx, labelStyle, value)
			}
		}))
		return c.EndBlock()
	}
}
