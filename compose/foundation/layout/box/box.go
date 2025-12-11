package box

import (
	"go-compose-dev/internal/layoutnode"

	"gioui.org/layout"
)

func Box(content Composable, options ...BoxOption) Composable {
	opts := DefaultBoxOptions()
	for _, option := range options {
		option(&opts)
	}
	return func(c Composer) Composer {
		c.StartBlock("Box")
		c.Modifier(func(modifier Modifier) Modifier {
			return modifier.Then(opts.Modifier)
		})
		c.WithComposable(content)
		c.SetWidgetConstructor(boxWidgetConstructor(opts))

		return c.EndBlock()
	}
}

func boxWidgetConstructor(options BoxOptions) layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {

			stackChildren := []StackChild{}
			for _, child := range node.Children() {

				childLayoutNode := child.(layoutnode.NodeCoordinator)

				matchParent := childLayoutNode.Elements().GetElement(MatchParentSizeKey)

				if matchParent.IsSome() {
					stackChildren = append(stackChildren, layout.Expanded(func(gtx LayoutContext) LayoutDimensions {
						return childLayoutNode.Layout(gtx)
					}))
				} else {
					stackChildren = append(stackChildren, layout.Stacked(func(gtx LayoutContext) LayoutDimensions {
						return childLayoutNode.Layout(gtx)
					}))
				}
			}

			return Stack{
				Alignment: options.Alignment,
			}.Layout(gtx, stackChildren...)
		}
	})

}
