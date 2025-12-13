package divider

import (
	"go-compose-dev/internal/layoutnode"
)

const Material3DivideNodeID = "Material3Divider"

func Divider(options ...DividerOption) Composable {
	return func(c Composer) Composer {

		opts := DefaultDividerOptions()
		for _, option := range options {
			if option == nil {
				continue
			}
			option(&opts)
		}

		c.StartBlock(Material3DivideNodeID)
		c.Modifier(func(modifier Modifier) Modifier {
			return modifier.Then(opts.Modifier)
		})
		c.SetWidgetConstructor(widgetConstructor(opts))

		return c.EndBlock()
	}
}

func widgetConstructor(options DividerOptions) layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {
			return M3Divider().Layout(gtx)
		}
	})

}
