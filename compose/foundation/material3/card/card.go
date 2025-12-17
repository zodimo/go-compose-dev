package card

import (
	"github.com/zodimo/go-compose/compose/foundation/layout/box"
	"github.com/zodimo/go-compose/compose/ui/graphics/shape"
	"github.com/zodimo/go-compose/internal/layoutnode"
	"github.com/zodimo/go-compose/modifiers/background"
	"github.com/zodimo/go-compose/modifiers/border"
	"github.com/zodimo/go-compose/modifiers/clip"
	"github.com/zodimo/go-compose/modifiers/padding"
	"github.com/zodimo/go-compose/modifiers/shadow"
	"github.com/zodimo/go-compose/theme"

	"gioui.org/unit"
	"git.sr.ht/~schnwalter/gio-mw/wdk/block"
)

const Material3CardNodeID = "Material3Card"

// Card corner radius (Material 3 Medium = 12dp)
var cardCornerShape = shape.RoundedCornerShape{Radius: unit.Dp(12)}

func Elevated(contents CardContentContainer, options ...CardOption) Composable {
	return cardComposable(cardElevated, contents, options...)
}

func Filled(contents CardContentContainer, options ...CardOption) Composable {
	return cardComposable(cardFilled, contents, options...)
}

func Outlined(contents CardContentContainer, options ...CardOption) Composable {
	return cardComposable(cardOutlined, contents, options...)
}

func cardComposable(kind cardKind, contents CardContentContainer, options ...CardOption) Composable {
	return func(c Composer) Composer {
		opts := DefaultCardOptions()
		for _, option := range options {
			if option == nil {
				continue
			}
			option(&opts)
		}

		args := CardConstructorArgs{
			Kind:     kind,
			Contents: contents,
		}

		c.StartBlock(Material3CardNodeID)

		// Build modifier chain: user modifier + card styling modifiers
		c.Modifier(func(modifier Modifier) Modifier {
			// Start with user-provided modifier
			m := modifier.Then(opts.Modifier)

			// Get card colors from theme
			colorRoles := theme.ColorHelper.ColorSelector()

			// Add background and clip based on card kind
			switch kind {
			case cardElevated:
				// Elevated card: shadow + background + rounded clip
				m = m.Then(shadow.Simple(unit.Dp(2), cardCornerShape))
				m = m.Then(background.Background(colorRoles.SurfaceRoles.ContainerLow, background.WithShape(cardCornerShape)))
				m = m.Then(clip.Clip(cardCornerShape))
			case cardOutlined:
				// Outlined card: background + border + rounded clip
				m = m.Then(background.Background(colorRoles.SurfaceRoles.Surface, background.WithShape(cardCornerShape)))
				m = m.Then(border.Border(unit.Dp(1), colorRoles.OutlineRoles.OutlineVariant, cardCornerShape))
				m = m.Then(clip.Clip(cardCornerShape))
			default: // Filled
				// Filled card: background + rounded clip
				m = m.Then(background.Background(colorRoles.SurfaceRoles.ContainerHighest, background.WithShape(cardCornerShape)))
				m = m.Then(clip.Clip(cardCornerShape))
			}

			return m
		})

		for _, child := range contents.children {
			if child.cover {
				c.WithComposable(child.composable)
			} else {
				// add padding
				c.WithComposable(box.Box(child.composable, box.WithModifier(padding.All(16))))
			}
		}

		c.SetWidgetConstructor(cardWidgetConstructor(args))

		return c.EndBlock()
	}
}

type CardConstructorArgs struct {
	Kind     cardKind
	Contents CardContentContainer
}

func cardWidgetConstructor(args CardConstructorArgs) layoutnode.LayoutNodeWidgetConstructor {
	return layoutnode.NewLayoutNodeWidgetConstructor(func(node layoutnode.LayoutNode) layoutnode.GioLayoutWidget {
		return func(gtx layoutnode.LayoutContext) layoutnode.LayoutDimensions {

			// Build layout widgets for children
			var widgets []block.Segment
			for _, child := range node.Children() {
				// content
				widgets = append(widgets, block.NewSegment(child.(layoutnode.NodeCoordinator).Layout))

			}

			// Simple layout: just lay out children vertically
			// NO Clickable wrapper - this fixes focus navigation!
			return block.Line{
				Axis:     block.AxisVertical,
				Overflow: block.OverflowClip,
			}.Layout(gtx, widgets...)
		}
	})
}
