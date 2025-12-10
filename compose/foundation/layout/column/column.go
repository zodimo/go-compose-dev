package column

import (
	"go-compose-dev/internal/layoutnode"

	"gioui.org/layout"
)

func DefaultColumnOptions() ColumnOptions {
	return ColumnOptions{
		Modifier:  EmptyModifier,
		Spacing:   SpaceEnd, // 0
		Alignment: Start,    // 0
	}
}

func Column(content Composable, options ...ColumnOption) Composable {
	opts := DefaultColumnOptions()
	for _, option := range options {
		option(&opts)
	}
	return func(c Composer) Composer {
		c.StartBlock("Column")
		c.Modifier(func(modifier Modifier) Modifier {
			return modifier.Then(opts.Modifier)
		})
		c.WithComposable(content)
		c.SetWidgetConstructor(columnWidgetConstructor())

		return c.EndBlock()
	}
}

func columnWidgetConstructor() layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {

			flexedChildren := []layout.FlexChild{}
			for _, child := range node.Children() {

				childLayoutNode := child.(layoutnode.NodeCoordinator)

				// elementStore := childLayoutNode.Elements()

				var weight float32 = -1 // get child weight
				if weight == -1 {
					flexedChildren = append(flexedChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return childLayoutNode.LayoutSelf(gtx)
					}))
				} else {
					flexedChildren = append(flexedChildren, layout.Flexed(weight, func(gtx layout.Context) layout.Dimensions {
						return childLayoutNode.LayoutSelf(gtx)
					}))
				}
			}

			return layout.Flex{
				Axis: layout.Vertical,
				// Spacing:   constructorArgs.Options.Spacing,
				// Alignment: constructorArgs.Options.Alignment,
			}.Layout(gtx, flexedChildren...)
		}
	})

}
